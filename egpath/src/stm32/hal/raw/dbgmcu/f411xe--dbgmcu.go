// +build f411xe

// Peripheral: DBGMCU_Periph  Debug MCU.
// Instances:
//  DBGMCU  mmap.DBGMCU_BASE
// Registers:
//  0x00 32  IDCODE MCU device ID code.
//  0x04 32  CR     Debug MCU configuration register.
//  0x08 32  APB1FZ Debug MCU APB1 freeze register.
//  0x0C 32  APB2FZ Debug MCU APB2 freeze register.
// Import:
//  stm32/o/f411xe/mmap
package dbgmcu

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	DEV_ID IDCODE = 0xFFF << 0   //+
	REV_ID IDCODE = 0xFFFF << 16 //+
)

const (
	DEV_IDn = 0
	REV_IDn = 16
)

const (
	DBG_SLEEP   CR = 0x01 << 0 //+
	DBG_STOP    CR = 0x01 << 1 //+
	DBG_STANDBY CR = 0x01 << 2 //+
	TRACE_IOEN  CR = 0x01 << 5 //+
	TRACE_MODE  CR = 0x03 << 6 //+
)

const (
	DBG_SLEEPn   = 0
	DBG_STOPn    = 1
	DBG_STANDBYn = 2
	TRACE_IOENn  = 5
	TRACE_MODEn  = 6
)
