package bcmw

import (
	"delay"
	"encoding/binary/le"
	"fmt"
	"io"

	"sdcard"
	"sdcard/sdio"
)

func (d *Driver) debug(f string, args ...interface{}) {
	if d.error() {
		fmt.Printf("error: %v\n", d.firstErr())
	} else {
		fmt.Printf(f, args...)
	}
}

func (d *Driver) error() bool {
	return d.sd.Err(false) != nil ||
		d.ioStatus&^sdcard.IO_CURRENT_STATE != 0 ||
		d.err != 0
}

func (d *Driver) firstErr() error {
	if err := d.sd.Err(false); err != nil {
		return err
	}
	if d.ioStatus&^sdcard.IO_CURRENT_STATE != 0 {
		return ErrIOStatus
	}
	if d.err != 0 {
		return d.err
	}
	return nil
}

func (d *Driver) sdioRead8(f int, addr int) (b byte) {
	if d.error() {
		return 0
	}
	b, d.ioStatus = d.sd.SendCmd(sdcard.CMD52(f, addr, sdcard.Read, 0)).R5()
	return
}

func (d *Driver) sdioWrite8(f int, addr int, b byte) {
	if d.error() {
		return
	}
	_, d.ioStatus = d.sd.SendCmd(sdcard.CMD52(f, addr, sdcard.Write, b)).R5()
}

func (d *Driver) sdioEnableFunc(f, timeoutms int) {
	if d.error() {
		return
	}
	r := d.sdioRead8(cia, sdio.CCCR_IOEN)
	m := byte(1 << uint(f))
	d.sdioWrite8(cia, sdio.CCCR_IOEN, r|m)
	for retry := timeoutms >> 1; retry > 0; retry-- {
		r := d.sdioRead8(cia, sdio.CCCR_IORDY)
		if d.error() || r&m != 0 {
			return
		}
		delay.Millisec(2)
	}
	d.err = ErrTimeout
}

func (d *Driver) sdioDisableFunc(f int) {
	if d.error() {
		return
	}
	r := d.sdioRead8(cia, sdio.CCCR_IOEN)
	r &^= 1 << uint(f)
	d.sdioWrite8(cia, sdio.CCCR_IOEN, r)
}

func (d *Driver) sdioSetBlockSize(f, blksiz int) {
	if d.error() {
		return
	}
	d.sdioWrite8(cia, f<<8+sdio.FBR_BLKSIZE0, byte(blksiz))
	d.sdioWrite8(cia, f<<8+sdio.FBR_BLKSIZE1, byte(blksiz>>8))
	return
}

// The backplaneSetWindow, backplaneRead32, backplaneWrite32 are methods
// that allow to access core registers in the way specific to Sonics Silicon
// Backplane. More info: http://www.gc-linux.org/wiki/Wii:WLAN

func (d *Driver) backplaneSetWindow(addr uint32) {
	if d.error() {
		return
	}
	addr &^= 0x7FFF
	win := d.backplaneWindow
	if win == addr {
		return
	}
	d.backplaneWindow = addr
	for n := 0; n < 3; n++ {
		addr >>= 8
		win >>= 8
		if a := byte(addr); a != byte(win) {
			d.sdioWrite8(backplane, sbsdioFunc1SBAddrLow+n, a)
		}
	}
	if d.error() {
		d.backplaneWindow = 0
	}
}

func (d *Driver) backplaneRead8(addr uint32) byte {
	d.backplaneSetWindow(addr)
	return d.sdioRead8(backplane, int(addr&0x7FFF))
}

func (d *Driver) backplaneWrite8(addr uint32, b byte) {
	d.backplaneSetWindow(addr)
	d.sdioWrite8(backplane, int(addr&0x7FFF), b)
}

func (d *Driver) backplaneRead32(addr uint32) uint32 {
	d.backplaneSetWindow(addr)
	if d.error() {
		return 0
	}
	sd := d.sd
	var buf [1]uint64
	sd.SetupData(sdcard.Recv|sdcard.IO|sdcard.Block4, buf[:], 4)
	_, d.ioStatus = sd.SendCmd(sdcard.CMD53(
		backplane, int(addr&0x7FFF|sbsdioAccess32bit), sdcard.Read, 4,
	)).R5()
	return le.Decode32(sdcard.AsData(buf[:]).Bytes())
}

func (d *Driver) backplaneWrite32(addr, v uint32) {
	d.backplaneSetWindow(addr)
	if d.error() {
		return
	}
	sd := d.sd
	var buf [1]uint64
	le.Encode32(sdcard.AsData(buf[:]).Bytes(), v)
	sd.SetupData(sdcard.Send|sdcard.IO|sdcard.Block4, buf[:], 4)
	_, d.ioStatus = sd.SendCmd(sdcard.CMD53(
		backplane, int(addr&0x7FFF|sbsdioAccess32bit), sdcard.Write, 4,
	)).R5()
}

func (d *Driver) backplaneWrite(addr uint32, data []byte) {
	if d.error() || len(data) == 0 {
		return
	}
	head, aligned, tail := uint64slice(data)
	for _, b := range head {
		d.backplaneWrite8(addr, b)
		addr++
	}
	if d.error() {
		return
	}
	sd := d.sd
	for len(aligned) >= 8 {
		nbl := len(aligned) >> 3
		if nbl > 0x1FF {
			nbl = 0x1FF
		}
		d.backplaneSetWindow(addr)
		sd.SetupData(sdcard.Send|sdcard.IO|sdcard.Block64, aligned, nbl*64)
		_, d.ioStatus = sd.SendCmd(sdcard.CMD53(
			backplane, int(addr&0x7FFF), sdcard.BlockWrite|sdcard.IncAddr, nbl,
		)).R5()
		if d.error() {
			return
		}
		aligned = aligned[nbl*8:]
		addr += uint32(nbl) * 64
	}
	bms := [...]sdcard.DataMode{sdcard.Block32, sdcard.Block16, sdcard.Block8}
	for i, bm := range bms {
		n := 1 << uint(2-i)
		if n > len(aligned) {
			continue
		}
		siz := n * 8
		d.backplaneSetWindow(addr)
		sd.SetupData(sdcard.Send|sdcard.IO|bm, aligned, siz)
		_, d.ioStatus = sd.SendCmd(sdcard.CMD53(
			backplane, int(addr&0x7FFF), sdcard.Write|sdcard.IncAddr, siz,
		)).R5()
		if d.error() {
			return
		}
		aligned = aligned[n:]
		addr += uint32(siz)
	}
	for _, b := range tail {
		d.backplaneWrite8(addr, b)
		addr++
	}
}

func (d *Driver) backplaneUpload(addr uint32, r io.Reader) error {
	var buf [4 * 64]byte
	for {
		n, err := r.Read(buf[:])
		d.backplaneWrite(addr, buf[:n])
		if d.error() {
			return d.firstErr()
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		addr += uint32(n)
	}
	return nil
}

func (d *Driver) chipIsCoreUp(core int) bool {
	if d.error() {
		return false
	}
	base := wrapBase[core]
	r := d.backplaneRead32(base + agentIOCtl)
	if r&(ioCtlFGC|ioCtlClk) != ioCtlClk {
		return false
	}
	return d.backplaneRead32(base+agentResetCtl)&1 == 0
}

func (d *Driver) chipCoreDisable(core int, prereset, reset uint32) {
	if d.error() {
		return
	}
	base := wrapBase[core]
	if d.backplaneRead32(base+agentResetCtl)&1 == 0 {
		goto configure // Already in reset state.
	}
	d.backplaneWrite32(base+agentIOCtl, ioCtlFGC|ioCtlClk|prereset)
	d.backplaneRead32(base + agentIOCtl)
	d.backplaneWrite32(base+agentResetCtl, 1)
	delay.Millisec(1)
	if d.backplaneRead32(base+agentResetCtl)&1 == 0 {
		if d.err == 0 {
			d.err = ErrTimeout
		}
		return
	}
configure:
	d.backplaneWrite32(base+agentIOCtl, ioCtlFGC|ioCtlClk|reset)
	d.backplaneRead32(base + agentIOCtl)
}

func (d *Driver) chipCoreReset(core int, prereset, reset, postreset uint32) {
	if d.error() {
		return
	}
	d.chipCoreDisable(core, prereset, reset)
	base := wrapBase[core]
	for retry := 3; ; retry-- {
		d.backplaneWrite32(base+agentResetCtl, 0)
		delay.Millisec(1)
		r := d.backplaneRead32(base + agentResetCtl)
		if d.error() {
			return
		}
		if r&1 == 0 {
			break
		}
		if retry == 1 {
			d.err = ErrTimeout
			return
		}
	}
	d.backplaneWrite32(base+agentIOCtl, ioCtlClk|postreset)
	d.backplaneRead32(base + agentIOCtl)
}
