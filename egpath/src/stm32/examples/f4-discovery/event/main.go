package main

import (
	"rtos"

	"arch/cortexm/bitband"

	"stm32/hal/exti"
	"stm32/hal/gpio"
	"stm32/hal/irq"
	"stm32/hal/system"
	"stm32/hal/system/timer/systick"
)

const button = gpio.Pin0

var leds struct{ Red, Green, Blue, Orange bitband.Bit }

func init() {
	system.Setup168(8)
	systick.Setup()
	
	gpio.A.EnableClock(true)
	btnport := gpio.A
	gpio.D.EnableClock(false)
	ledport := gpio.D

	// LEDs

	cfg := gpio.Config{Mode: gpio.Out, Speed: gpio.Low}
	for pin := 12; pin <= 15; pin++ {
		ledport.SetupPin(pin, &cfg)
	}
	pins := ledport.OutPins()
	leds.Green = pins.Bit(12)
	leds.Orange = pins.Bit(13)
	leds.Red = pins.Bit(14)
	leds.Blue = pins.Bit(15)

	// Button

	btnport.Setup(button, &gpio.Config{Mode: gpio.In})
	line := exti.Lines(button)
	line.Connect(btnport)
	line.EnableRiseTrig()
	line.EnableInt()
	rtos.IRQ(irq.EXTI0).Enable()
}

var event = rtos.NewEvent()

func buttonISR() {
	exti.Lines(button).ClearPending()
	event.Send()
}

//emgo:const
//c:__attribute__((section(".ISRs")))
var ISRs = [...]func(){
	irq.EXTI0: buttonISR,
}

func main() {
	sec := func(s int) int64 {
		return rtos.Nanosec() + int64(s)*1e9
	}
	for {
		leds.Green.Clear()
		leds.Orange.Set()
		event.Wait(sec(2))
		leds.Orange.Clear()
		leds.Red.Set()
		event.Wait(sec(2))
		leds.Red.Clear()
		leds.Blue.Set()
		event.Wait(sec(2))
		leds.Blue.Clear()
		leds.Green.Set()
		event.Wait(sec(2))
	}
}