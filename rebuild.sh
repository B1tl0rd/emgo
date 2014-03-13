#!/bin/bash

set -e

list="
builtin

sync/barrier
delay
mmio

cortexm
cortexm/irq
cortexm/systick
cortexm/sleep

runtime/noos
runtime

strconv
math/matrix32

stm32/stlink
stm32/f4/clock
stm32/f4/flash
stm32/f4/gpio
stm32/f4/periph
stm32/f4/setup

stm32/l1/clock
stm32/l1/flash
stm32/l1/gpio
stm32/l1/periph
stm32/l1/setup
"
for p in $list; do 
	echo $p
	egc $p
done
