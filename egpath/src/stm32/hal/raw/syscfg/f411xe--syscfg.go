// +build f411xe

// Peripheral: SYSCFG_Periph  System configuration controller.
// Instances:
//  SYSCFG  mmap.SYSCFG_BASE
// Registers:
//  0x00 32  MEMRMP    Memory remap register.
//  0x04 32  PMC       Peripheral mode configuration register.
//  0x08 32  EXTICR[4] External interrupt configuration registers.
//  0x20 32  CMPCR     Compensation cell control register.
// Import:
//  stm32/o/f411xe/mmap
package syscfg

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	MEM_MODE MEMRMP = 0x03 << 0 //+ SYSCFG_Memory Remap Config.
)

const (
	MEM_MODEn = 0
)

const (
	ADC1DC2 PMC = 0x01 << 16 //+ Refer to AN4073 on how to use this bit.
)

const (
	ADC1DC2n = 16
)

const (
	EXTI0    EXTICR = 0x0F << 0  //+ EXTI 0 configuration.
	EXTI1    EXTICR = 0x0F << 4  //+ EXTI 1 configuration.
	EXTI2    EXTICR = 0x0F << 8  //+ EXTI 2 configuration.
	EXTI3    EXTICR = 0x0F << 12 //+ EXTI 3 configuration.
	EXTI0_PA EXTICR = 0x00 << 12 //  PA[0] pin.
	EXTI0_PB EXTICR = 0x01 << 0  //  PB[0] pin.
	EXTI0_PC EXTICR = 0x02 << 0  //  PC[0] pin.
	EXTI0_PD EXTICR = 0x03 << 0  //  PD[0] pin.
	EXTI0_PE EXTICR = 0x04 << 0  //  PE[0] pin.
	EXTI0_PH EXTICR = 0x07 << 0  //  PH[0] pin.
	EXTI1_PA EXTICR = 0x00 << 12 //  PA[1] pin.
	EXTI1_PB EXTICR = 0x01 << 4  //  PB[1] pin.
	EXTI1_PC EXTICR = 0x02 << 4  //  PC[1] pin.
	EXTI1_PD EXTICR = 0x03 << 4  //  PD[1] pin.
	EXTI1_PE EXTICR = 0x04 << 4  //  PE[1] pin.
	EXTI1_PH EXTICR = 0x07 << 4  //  PH[1] pin.
	EXTI2_PA EXTICR = 0x00 << 12 //  PA[2] pin.
	EXTI2_PB EXTICR = 0x01 << 8  //  PB[2] pin.
	EXTI2_PC EXTICR = 0x02 << 8  //  PC[2] pin.
	EXTI2_PD EXTICR = 0x03 << 8  //  PD[2] pin.
	EXTI2_PE EXTICR = 0x04 << 8  //  PE[2] pin.
	EXTI2_PH EXTICR = 0x07 << 8  //  PH[2] pin.
	EXTI3_PA EXTICR = 0x00 << 12 //  PA[3] pin.
	EXTI3_PB EXTICR = 0x01 << 12 //  PB[3] pin.
	EXTI3_PC EXTICR = 0x02 << 12 //  PC[3] pin.
	EXTI3_PD EXTICR = 0x03 << 12 //  PD[3] pin.
	EXTI3_PE EXTICR = 0x04 << 12 //  PE[3] pin.
	EXTI3_PH EXTICR = 0x07 << 12 //  PH[3] pin.
)

const (
	EXTI0n = 0
	EXTI1n = 4
	EXTI2n = 8
	EXTI3n = 12
)

const (
	CMP_PD CMPCR = 0x01 << 0 //+ Compensation cell ready flag.
	READY  CMPCR = 0x01 << 8 //+ Compensation cell power-down.
)

const (
	CMP_PDn = 0
	READYn  = 8
)
