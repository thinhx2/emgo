// This example demonstrates usage of FTDI EVE based displays.
//
// It seems that FT800CB-HY50B display is unstable with fast SPI. If you have
// problems please reduce SPI speed or better desolder U1 and U2 (74LCX125
// buffers) and short the U1:2-3,5-6,11-2, U2:2-3,5-6 traces.
package main

import (
	"delay"
	"fmt"
	"rtos"

	"display/eve"
	"display/eve/evetest"
	"display/eve/ft81"

	"stm32/evedci"

	"stm32/hal/dma"
	"stm32/hal/exti"
	"stm32/hal/gpio"
	"stm32/hal/irq"
	"stm32/hal/spi"
	"stm32/hal/system"
	"stm32/hal/system/timer/systick"
)

var dci *evedci.SPI

func init() {
	system.Setup96(8)
	systick.Setup(2e6)

	// GPIO

	gpio.A.EnableClock(true)
	spiport, sck, miso, mosi := gpio.A, gpio.Pin5, gpio.Pin6, gpio.Pin7
	pdn := gpio.A.Pin(9)

	gpio.B.EnableClock(true)
	csn := gpio.B.Pin(6)

	gpio.C.EnableClock(true)
	irqn := gpio.C.Pin(7)

	// EVE control lines

	cfg := gpio.Config{Mode: gpio.Out, Speed: gpio.High}
	pdn.Setup(&cfg)
	csn.Setup(&cfg)
	irqn.Setup(&gpio.Config{Mode: gpio.In})
	irqline := exti.Lines(irqn.Mask())
	irqline.Connect(irqn.Port())
	irqline.EnableFallTrig()
	irqline.EnableIRQ()
	rtos.IRQ(irq.EXTI9_5).Enable()

	// EVE SPI

	spiport.Setup(sck|mosi, &gpio.Config{Mode: gpio.Alt, Speed: gpio.High})
	spiport.Setup(miso, &gpio.Config{Mode: gpio.AltIn})
	spiport.SetAltFunc(sck|miso|mosi, gpio.SPI1)
	d := dma.DMA2
	d.EnableClock(true)
	spidrv := spi.NewDriver(spi.SPI1, d.Channel(3, 3), d.Channel(2, 3))
	rtos.IRQ(irq.SPI1).Enable()
	rtos.IRQ(irq.DMA2_Stream2).Enable()
	rtos.IRQ(irq.DMA2_Stream3).Enable()

	dci = evedci.NewSPI(spidrv, csn, pdn)
	dci.Setup(11e6)
}

func curFreq(lcd *eve.Driver) uint32 {
	clk1 := lcd.ReadUint32(ft81.REG_CLOCK)
	t1 := rtos.Nanosec()
	delay.Millisec(8)
	clk2 := lcd.ReadUint32(ft81.REG_CLOCK)
	t2 := rtos.Nanosec()
	return uint32(int64(clk2-clk1) * 1e9 / (t2 - t1))
}

func main() {
	delay.Millisec(100) // For SWO output.

	spip := dci.SPI().Periph()
	fmt.Printf("\nSPI on %s (%d MHz).\n", spip.Bus(), spip.Bus().Clock()/1e6)
	fmt.Printf("SPI speed: %d bps.\n", spip.Baudrate(spip.Conf()))

	// Dithering causes distortion of vertical gradients on my FT811CB-HY50HD:
	// horizontal darker lines appear. The effect is dramatic if both Dither and
	// Spread are set (this is default after reset). Please write on Emgo forum
	// how it looks on your screen.

	lcd := eve.NewDriver(dci, 128)
	lcd.Init(&eve.Default800x480, &eve.Config{Dither: 1})

	fmt.Printf("EVE clock: %d Hz.\n", curFreq(lcd))
	dci.Setup(30e6)
	fmt.Printf("SPI speed: %d bps.\n", spip.Baudrate(spip.Conf()))

	if err := evetest.Run(lcd); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("End.\n")
}

func lcdSPIISR() {
	dci.SPI().ISR()
}

func lcdRxDMAISR() {
	dci.SPI().DMAISR(dci.SPI().RxDMA())
}

func lcdTxDMAISR() {
	dci.SPI().DMAISR(dci.SPI().TxDMA())
}

func exti9_5ISR() {
	pending := exti.Pending()
	pending &= exti.L5 | exti.L6 | exti.L7 | exti.L8 | exti.L9
	pending.ClearPending()
	if pending&exti.L7 != 0 {
		dci.ISR()
	}
}

//emgo:const
//c:__attribute__((section(".ISRs")))
var ISRs = [...]func(){
	irq.SPI1:         lcdSPIISR,
	irq.DMA2_Stream2: lcdRxDMAISR,
	irq.DMA2_Stream3: lcdTxDMAISR,
	irq.EXTI9_5:      exti9_5ISR,
}
