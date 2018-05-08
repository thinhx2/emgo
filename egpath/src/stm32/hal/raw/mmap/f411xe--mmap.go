// +build f411xe

// Package mmap provides base memory adresses for all peripherals.
package mmap

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	FLASH_BASE      uintptr = 0x08000000 // FLASH(up to 1 MB) base address in the alias region
	SRAM1_BASE      uintptr = 0x20000000 // SRAM1(128 KB) base address in the alias region
	PERIPH_BASE     uintptr = 0x40000000 // Peripheral base address in the alias region
	SRAM1_BB_BASE   uintptr = 0x22000000 // SRAM1(128 KB) base address in the bit-band region
	PERIPH_BB_BASE  uintptr = 0x42000000 // Peripheral base address in the bit-band region
	BKPSRAM_BB_BASE uintptr = 0x42480000 // Backup SRAM(4 KB) base address in the bit-band region
	FLASH_END       uintptr = 0x0807FFFF // FLASH end address
	FLASH_OTP_BASE  uintptr = 0x1FFF7800 // Base address of : (up to 528 Bytes) embedded FLASH OTP Area
	FLASH_OTP_END   uintptr = 0x1FFF7A0F // End address of : (up to 528 Bytes) embedded FLASH OTP Area
	SRAM_BASE       uintptr = SRAM1_BASE
	SRAM_BB_BASE    uintptr = SRAM1_BB_BASE
)

// Peripheral memory map
const (
	APB1PERIPH_BASE uintptr = PERIPH_BASE
	APB2PERIPH_BASE uintptr = PERIPH_BASE + 0x00010000
	AHB1PERIPH_BASE uintptr = PERIPH_BASE + 0x00020000
	AHB2PERIPH_BASE uintptr = PERIPH_BASE + 0x10000000
)

// APB1 peripherals
const (
	TIM2_BASE    uintptr = APB1PERIPH_BASE + 0x0000
	TIM3_BASE    uintptr = APB1PERIPH_BASE + 0x0400
	TIM4_BASE    uintptr = APB1PERIPH_BASE + 0x0800
	TIM5_BASE    uintptr = APB1PERIPH_BASE + 0x0C00
	RTC_BASE     uintptr = APB1PERIPH_BASE + 0x2800
	WWDG_BASE    uintptr = APB1PERIPH_BASE + 0x2C00
	IWDG_BASE    uintptr = APB1PERIPH_BASE + 0x3000
	I2S2ext_BASE uintptr = APB1PERIPH_BASE + 0x3400
	SPI2_BASE    uintptr = APB1PERIPH_BASE + 0x3800
	SPI3_BASE    uintptr = APB1PERIPH_BASE + 0x3C00
	I2S3ext_BASE uintptr = APB1PERIPH_BASE + 0x4000
	USART2_BASE  uintptr = APB1PERIPH_BASE + 0x4400
	I2C1_BASE    uintptr = APB1PERIPH_BASE + 0x5400
	I2C2_BASE    uintptr = APB1PERIPH_BASE + 0x5800
	I2C3_BASE    uintptr = APB1PERIPH_BASE + 0x5C00
	PWR_BASE     uintptr = APB1PERIPH_BASE + 0x7000
)

// APB2 peripherals
const (
	TIM1_BASE        uintptr = APB2PERIPH_BASE + 0x0000
	USART1_BASE      uintptr = APB2PERIPH_BASE + 0x1000
	USART6_BASE      uintptr = APB2PERIPH_BASE + 0x1400
	ADC1_BASE        uintptr = APB2PERIPH_BASE + 0x2000
	ADC1_COMMON_BASE uintptr = APB2PERIPH_BASE + 0x2300
	ADC_BASE         uintptr = ADC1_COMMON_BASE
	SDIO_BASE        uintptr = APB2PERIPH_BASE + 0x2C00
	SPI1_BASE        uintptr = APB2PERIPH_BASE + 0x3000
	SPI4_BASE        uintptr = APB2PERIPH_BASE + 0x3400
	SYSCFG_BASE      uintptr = APB2PERIPH_BASE + 0x3800
	EXTI_BASE        uintptr = APB2PERIPH_BASE + 0x3C00
	TIM9_BASE        uintptr = APB2PERIPH_BASE + 0x4000
	TIM10_BASE       uintptr = APB2PERIPH_BASE + 0x4400
	TIM11_BASE       uintptr = APB2PERIPH_BASE + 0x4800
	SPI5_BASE        uintptr = APB2PERIPH_BASE + 0x5000
)

// AHB1 peripherals
const (
	GPIOA_BASE        uintptr = AHB1PERIPH_BASE + 0x0000
	GPIOB_BASE        uintptr = AHB1PERIPH_BASE + 0x0400
	GPIOC_BASE        uintptr = AHB1PERIPH_BASE + 0x0800
	GPIOD_BASE        uintptr = AHB1PERIPH_BASE + 0x0C00
	GPIOE_BASE        uintptr = AHB1PERIPH_BASE + 0x1000
	GPIOH_BASE        uintptr = AHB1PERIPH_BASE + 0x1C00
	CRC_BASE          uintptr = AHB1PERIPH_BASE + 0x3000
	RCC_BASE          uintptr = AHB1PERIPH_BASE + 0x3800
	FLASH_R_BASE      uintptr = AHB1PERIPH_BASE + 0x3C00
	DMA1_BASE         uintptr = AHB1PERIPH_BASE + 0x6000
	DMA1_Stream0_BASE uintptr = DMA1_BASE + 0x010
	DMA1_Stream1_BASE uintptr = DMA1_BASE + 0x028
	DMA1_Stream2_BASE uintptr = DMA1_BASE + 0x040
	DMA1_Stream3_BASE uintptr = DMA1_BASE + 0x058
	DMA1_Stream4_BASE uintptr = DMA1_BASE + 0x070
	DMA1_Stream5_BASE uintptr = DMA1_BASE + 0x088
	DMA1_Stream6_BASE uintptr = DMA1_BASE + 0x0A0
	DMA1_Stream7_BASE uintptr = DMA1_BASE + 0x0B8
	DMA2_BASE         uintptr = AHB1PERIPH_BASE + 0x6400
	DMA2_Stream0_BASE uintptr = DMA2_BASE + 0x010
	DMA2_Stream1_BASE uintptr = DMA2_BASE + 0x028
	DMA2_Stream2_BASE uintptr = DMA2_BASE + 0x040
	DMA2_Stream3_BASE uintptr = DMA2_BASE + 0x058
	DMA2_Stream4_BASE uintptr = DMA2_BASE + 0x070
	DMA2_Stream5_BASE uintptr = DMA2_BASE + 0x088
	DMA2_Stream6_BASE uintptr = DMA2_BASE + 0x0A0
	DMA2_Stream7_BASE uintptr = DMA2_BASE + 0x0B8
)

// Debug MCU registers base address
const (
	DBGMCU_BASE uintptr = 0xE0042000
)

// USB registers base address
const (
	USB_OTG_FS_PERIPH_BASE    uintptr = 0x50000000
	USB_OTG_GLOBAL_BASE       uintptr = 0x000
	USB_OTG_DEVICE_BASE       uintptr = 0x800
	USB_OTG_IN_ENDPOINT_BASE  uintptr = 0x900
	USB_OTG_OUT_ENDPOINT_BASE uintptr = 0xB00
	USB_OTG_EP_REG_SIZE       uintptr = 0x20
	USB_OTG_HOST_BASE         uintptr = 0x400
	USB_OTG_HOST_PORT_BASE    uintptr = 0x440
	USB_OTG_HOST_CHANNEL_BASE uintptr = 0x500
	USB_OTG_HOST_CHANNEL_SIZE uintptr = 0x20
	USB_OTG_PCGCCTL_BASE      uintptr = 0xE00
	USB_OTG_FIFO_BASE         uintptr = 0x1000
	USB_OTG_FIFO_SIZE         uintptr = 0x1000
	UID_BASE                  uintptr = 0x1FFF7A10 // Unique device ID register base address
	FLASHSIZE_BASE            uintptr = 0x1FFF7A22 // FLASH Size register base address
	PACKAGE_BASE              uintptr = 0x1FFF7BF0 // Package size register base address
)
