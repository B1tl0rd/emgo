// Peripheral: FLASH_Periph  FLASH Registers.
// Instances:
//  FLASH  mmap.FLASH_R_BASE
// Registers:
//  0x00 32  ACR
//  0x04 32  KEYR
//  0x08 32  OPTKEYR
//  0x0C 32  SR
//  0x10 32  CR
//  0x14 32  AR
//  0x18 32  RESERVED
//  0x1C 32  OBR
//  0x20 32  WRPR
// Import:
//  stm32/o/f10x_md/mmap
package flash

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	LATENCY   ACR = 0x03 << 0 //+ LATENCY[2:0] bits (Latency).
	LATENCY_0 ACR = 0x00 << 0 //  Bit 0.
	LATENCY_1 ACR = 0x01 << 0 //  Bit 0.
	LATENCY_2 ACR = 0x02 << 0 //  Bit 1.
	HLFCYA    ACR = 0x01 << 3 //+ Flash Half Cycle Access Enable.
	PRFTBE    ACR = 0x01 << 4 //+ Prefetch Buffer Enable.
	PRFTBS    ACR = 0x01 << 5 //+ Prefetch Buffer Status.
)

const (
	LATENCYn = 0
	HLFCYAn  = 3
	PRFTBEn  = 4
	PRFTBSn  = 5
)

const (
	FKEYR KEYR = 0xFFFFFFFF << 0 //+ FPEC Key.
)

const (
	FKEYRn = 0
)

const (
	BSY      SR = 0x01 << 0 //+ Busy.
	PGERR    SR = 0x01 << 2 //+ Programming Error.
	WRPRTERR SR = 0x01 << 4 //+ Write Protection Error.
	EOP      SR = 0x01 << 5 //+ End of operation.
)

const (
	BSYn      = 0
	PGERRn    = 2
	WRPRTERRn = 4
	EOPn      = 5
)

const (
	PG     CR = 0x01 << 0  //+ Programming.
	PER    CR = 0x01 << 1  //+ Page Erase.
	MER    CR = 0x01 << 2  //+ Mass Erase.
	OPTPG  CR = 0x01 << 4  //+ Option Byte Programming.
	OPTER  CR = 0x01 << 5  //+ Option Byte Erase.
	STRT   CR = 0x01 << 6  //+ Start.
	LOCK   CR = 0x01 << 7  //+ Lock.
	OPTWRE CR = 0x01 << 9  //+ Option Bytes Write Enable.
	ERRIE  CR = 0x01 << 10 //+ Error Interrupt Enable.
	EOPIE  CR = 0x01 << 12 //+ End of operation interrupt enable.
)

const (
	PGn     = 0
	PERn    = 1
	MERn    = 2
	OPTPGn  = 4
	OPTERn  = 5
	STRTn   = 6
	LOCKn   = 7
	OPTWREn = 9
	ERRIEn  = 10
	EOPIEn  = 12
)

const (
	FAR AR = 0xFFFFFFFF << 0 //+ Flash Address.
)

const (
	FARn = 0
)

const (
	OPTERR     OBR = 0x01 << 0 //+ Option Byte Error.
	RDPRT      OBR = 0x01 << 1 //+ Read protection.
	USER       OBR = 0xFF << 2 //+ User Option Bytes.
	WDG_SW     OBR = 0x01 << 2 //  WDG_SW.
	nRST_STOP  OBR = 0x02 << 2 //  nRST_STOP.
	nRST_STDBY OBR = 0x04 << 2 //  nRST_STDBY.
	BFB2       OBR = 0x08 << 2 //  BFB2.
)

const (
	OPTERRn = 0
	RDPRTn  = 1
	USERn   = 2
)
