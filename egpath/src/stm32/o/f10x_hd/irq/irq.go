// Package irq provides list of all defined external interrupts.
package irq

import "arch/cortexm/nvic"

const (
	WWDG            nvic.IRQ = 0  // Window WatchDog Interrupt.
	PVD             nvic.IRQ = 1  // PVD through EXTI Line detection Interrupt.
	TAMPER          nvic.IRQ = 2  // Tamper Interrupt.
	RTC             nvic.IRQ = 3  // RTC global Interrupt.
	FLASH           nvic.IRQ = 4  // FLASH global Interrupt.
	RCC             nvic.IRQ = 5  // RCC global Interrupt.
	EXTI0           nvic.IRQ = 6  // EXTI Line0 Interrupt.
	EXTI1           nvic.IRQ = 7  // EXTI Line1 Interrupt.
	EXTI2           nvic.IRQ = 8  // EXTI Line2 Interrupt.
	EXTI3           nvic.IRQ = 9  // EXTI Line3 Interrupt.
	EXTI4           nvic.IRQ = 10 // EXTI Line4 Interrupt.
	DMA1_Channel1   nvic.IRQ = 11 // DMA1 Channel 1 global Interrupt.
	DMA1_Channel2   nvic.IRQ = 12 // DMA1 Channel 2 global Interrupt.
	DMA1_Channel3   nvic.IRQ = 13 // DMA1 Channel 3 global Interrupt.
	DMA1_Channel4   nvic.IRQ = 14 // DMA1 Channel 4 global Interrupt.
	DMA1_Channel5   nvic.IRQ = 15 // DMA1 Channel 5 global Interrupt.
	DMA1_Channel6   nvic.IRQ = 16 // DMA1 Channel 6 global Interrupt.
	DMA1_Channel7   nvic.IRQ = 17 // DMA1 Channel 7 global Interrupt.
	ADC1_2          nvic.IRQ = 18 // ADC1 and ADC2 global Interrupt.
	USB_HP_CAN1_TX  nvic.IRQ = 19 // USB Device High Priority or CAN1 TX Interrupts.
	USB_LP_CAN1_RX0 nvic.IRQ = 20 // USB Device Low Priority or CAN1 RX0 Interrupts.
	CAN1_RX1        nvic.IRQ = 21 // CAN1 RX1 Interrupt.
	CAN1_SCE        nvic.IRQ = 22 // CAN1 SCE Interrupt.
	EXTI9_5         nvic.IRQ = 23 // External Line[9:5] Interrupts.
	TIM1_BRK        nvic.IRQ = 24 // TIM1 Break Interrupt.
	TIM1_UP         nvic.IRQ = 25 // TIM1 Update Interrupt.
	TIM1_TRG_COM    nvic.IRQ = 26 // TIM1 Trigger and Commutation Interrupt.
	TIM1_CC         nvic.IRQ = 27 // TIM1 Capture Compare Interrupt.
	TIM2            nvic.IRQ = 28 // TIM2 global Interrupt.
	TIM3            nvic.IRQ = 29 // TIM3 global Interrupt.
	TIM4            nvic.IRQ = 30 // TIM4 global Interrupt.
	I2C1_EV         nvic.IRQ = 31 // I2C1 Event Interrupt.
	I2C1_ER         nvic.IRQ = 32 // I2C1 Error Interrupt.
	I2C2_EV         nvic.IRQ = 33 // I2C2 Event Interrupt.
	I2C2_ER         nvic.IRQ = 34 // I2C2 Error Interrupt.
	SPI1            nvic.IRQ = 35 // SPI1 global Interrupt.
	SPI2            nvic.IRQ = 36 // SPI2 global Interrupt.
	USART1          nvic.IRQ = 37 // USART1 global Interrupt.
	USART2          nvic.IRQ = 38 // USART2 global Interrupt.
	USART3          nvic.IRQ = 39 // USART3 global Interrupt.
	EXTI15_10       nvic.IRQ = 40 // External Line[15:10] Interrupts.
	RTCAlarm        nvic.IRQ = 41 // RTC Alarm through EXTI Line Interrupt.
	USBWakeUp       nvic.IRQ = 42 // USB Device WakeUp from suspend through EXTI Line Interrupt.
	TIM8_BRK        nvic.IRQ = 43 // TIM8 Break Interrupt.
	TIM8_UP         nvic.IRQ = 44 // TIM8 Update Interrupt.
	TIM8_TRG_COM    nvic.IRQ = 45 // TIM8 Trigger and Commutation Interrupt.
	TIM8_CC         nvic.IRQ = 46 // TIM8 Capture Compare Interrupt.
	ADC3            nvic.IRQ = 47 // ADC3 global Interrupt.
	FSMC            nvic.IRQ = 48 // FSMC global Interrupt.
	SDIO            nvic.IRQ = 49 // SDIO global Interrupt.
	TIM5            nvic.IRQ = 50 // TIM5 global Interrupt.
	SPI3            nvic.IRQ = 51 // SPI3 global Interrupt.
	UART4           nvic.IRQ = 52 // UART4 global Interrupt.
	UART5           nvic.IRQ = 53 // UART5 global Interrupt.
	TIM6            nvic.IRQ = 54 // TIM6 global Interrupt.
	TIM7            nvic.IRQ = 55 // TIM7 global Interrupt.
	DMA2_Channel1   nvic.IRQ = 56 // DMA2 Channel 1 global Interrupt.
	DMA2_Channel2   nvic.IRQ = 57 // DMA2 Channel 2 global Interrupt.
	DMA2_Channel3   nvic.IRQ = 58 // DMA2 Channel 3 global Interrupt.
	DMA2_Channel4_5 nvic.IRQ = 59 // DMA2 Channel 4 and Channel 5 global Interrupt.
)