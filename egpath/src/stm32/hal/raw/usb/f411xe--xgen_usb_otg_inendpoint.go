// +build f411xe

package usb

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f411xe/mmap"
)

type USB_OTG_INEndpoint_Periph struct {
	DIEPCTL  RDIEPCTL
	_        uint32
	DIEPINT  RDIEPINT
	_        uint32
	DIEPTSIZ RDIEPTSIZ
	DIEPDMA  RDIEPDMA
	DTXFSTS  RDTXFSTS
}

func (p *USB_OTG_INEndpoint_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

type DIEPCTL uint32

func (b DIEPCTL) Field(mask DIEPCTL) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DIEPCTL) J(v int) DIEPCTL {
	return DIEPCTL(bits.Make32(v, uint32(mask)))
}

type RDIEPCTL struct{ mmio.U32 }

func (r *RDIEPCTL) Bits(mask DIEPCTL) DIEPCTL { return DIEPCTL(r.U32.Bits(uint32(mask))) }
func (r *RDIEPCTL) StoreBits(mask, b DIEPCTL) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDIEPCTL) SetBits(mask DIEPCTL)      { r.U32.SetBits(uint32(mask)) }
func (r *RDIEPCTL) ClearBits(mask DIEPCTL)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDIEPCTL) Load() DIEPCTL             { return DIEPCTL(r.U32.Load()) }
func (r *RDIEPCTL) Store(b DIEPCTL)           { r.U32.Store(uint32(b)) }

func (r *RDIEPCTL) AtomicStoreBits(mask, b DIEPCTL) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDIEPCTL) AtomicSetBits(mask DIEPCTL)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDIEPCTL) AtomicClearBits(mask DIEPCTL)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDIEPCTL struct{ mmio.UM32 }

func (rm RMDIEPCTL) Load() DIEPCTL   { return DIEPCTL(rm.UM32.Load()) }
func (rm RMDIEPCTL) Store(b DIEPCTL) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_INEndpoint_Periph) MPSIZ() RMDIEPCTL {
	return RMDIEPCTL{mmio.UM32{&p.DIEPCTL.U32, uint32(MPSIZ)}}
}

func (p *USB_OTG_INEndpoint_Periph) USBAEP() RMDIEPCTL {
	return RMDIEPCTL{mmio.UM32{&p.DIEPCTL.U32, uint32(USBAEP)}}
}

func (p *USB_OTG_INEndpoint_Periph) EONUM_DPID() RMDIEPCTL {
	return RMDIEPCTL{mmio.UM32{&p.DIEPCTL.U32, uint32(EONUM_DPID)}}
}

func (p *USB_OTG_INEndpoint_Periph) NAKSTS() RMDIEPCTL {
	return RMDIEPCTL{mmio.UM32{&p.DIEPCTL.U32, uint32(NAKSTS)}}
}

func (p *USB_OTG_INEndpoint_Periph) EPTYP() RMDIEPCTL {
	return RMDIEPCTL{mmio.UM32{&p.DIEPCTL.U32, uint32(EPTYP)}}
}

func (p *USB_OTG_INEndpoint_Periph) STALL() RMDIEPCTL {
	return RMDIEPCTL{mmio.UM32{&p.DIEPCTL.U32, uint32(STALL)}}
}

func (p *USB_OTG_INEndpoint_Periph) TXFNUM() RMDIEPCTL {
	return RMDIEPCTL{mmio.UM32{&p.DIEPCTL.U32, uint32(TXFNUM)}}
}

func (p *USB_OTG_INEndpoint_Periph) CNAK() RMDIEPCTL {
	return RMDIEPCTL{mmio.UM32{&p.DIEPCTL.U32, uint32(CNAK)}}
}

func (p *USB_OTG_INEndpoint_Periph) SNAK() RMDIEPCTL {
	return RMDIEPCTL{mmio.UM32{&p.DIEPCTL.U32, uint32(SNAK)}}
}

func (p *USB_OTG_INEndpoint_Periph) SD0PID_SEVNFRM() RMDIEPCTL {
	return RMDIEPCTL{mmio.UM32{&p.DIEPCTL.U32, uint32(SD0PID_SEVNFRM)}}
}

func (p *USB_OTG_INEndpoint_Periph) SODDFRM() RMDIEPCTL {
	return RMDIEPCTL{mmio.UM32{&p.DIEPCTL.U32, uint32(SODDFRM)}}
}

func (p *USB_OTG_INEndpoint_Periph) EPDIS() RMDIEPCTL {
	return RMDIEPCTL{mmio.UM32{&p.DIEPCTL.U32, uint32(EPDIS)}}
}

func (p *USB_OTG_INEndpoint_Periph) EPENA() RMDIEPCTL {
	return RMDIEPCTL{mmio.UM32{&p.DIEPCTL.U32, uint32(EPENA)}}
}

type DIEPINT uint32

func (b DIEPINT) Field(mask DIEPINT) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DIEPINT) J(v int) DIEPINT {
	return DIEPINT(bits.Make32(v, uint32(mask)))
}

type RDIEPINT struct{ mmio.U32 }

func (r *RDIEPINT) Bits(mask DIEPINT) DIEPINT { return DIEPINT(r.U32.Bits(uint32(mask))) }
func (r *RDIEPINT) StoreBits(mask, b DIEPINT) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDIEPINT) SetBits(mask DIEPINT)      { r.U32.SetBits(uint32(mask)) }
func (r *RDIEPINT) ClearBits(mask DIEPINT)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDIEPINT) Load() DIEPINT             { return DIEPINT(r.U32.Load()) }
func (r *RDIEPINT) Store(b DIEPINT)           { r.U32.Store(uint32(b)) }

func (r *RDIEPINT) AtomicStoreBits(mask, b DIEPINT) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDIEPINT) AtomicSetBits(mask DIEPINT)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDIEPINT) AtomicClearBits(mask DIEPINT)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDIEPINT struct{ mmio.UM32 }

func (rm RMDIEPINT) Load() DIEPINT   { return DIEPINT(rm.UM32.Load()) }
func (rm RMDIEPINT) Store(b DIEPINT) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_INEndpoint_Periph) XFRC() RMDIEPINT {
	return RMDIEPINT{mmio.UM32{&p.DIEPINT.U32, uint32(XFRC)}}
}

func (p *USB_OTG_INEndpoint_Periph) EPDISD() RMDIEPINT {
	return RMDIEPINT{mmio.UM32{&p.DIEPINT.U32, uint32(EPDISD)}}
}

func (p *USB_OTG_INEndpoint_Periph) TOC() RMDIEPINT {
	return RMDIEPINT{mmio.UM32{&p.DIEPINT.U32, uint32(TOC)}}
}

func (p *USB_OTG_INEndpoint_Periph) ITTXFE() RMDIEPINT {
	return RMDIEPINT{mmio.UM32{&p.DIEPINT.U32, uint32(ITTXFE)}}
}

func (p *USB_OTG_INEndpoint_Periph) INEPNE() RMDIEPINT {
	return RMDIEPINT{mmio.UM32{&p.DIEPINT.U32, uint32(INEPNE)}}
}

func (p *USB_OTG_INEndpoint_Periph) TXFE() RMDIEPINT {
	return RMDIEPINT{mmio.UM32{&p.DIEPINT.U32, uint32(TXFE)}}
}

func (p *USB_OTG_INEndpoint_Periph) TXFIFOUDRN() RMDIEPINT {
	return RMDIEPINT{mmio.UM32{&p.DIEPINT.U32, uint32(TXFIFOUDRN)}}
}

func (p *USB_OTG_INEndpoint_Periph) BNA() RMDIEPINT {
	return RMDIEPINT{mmio.UM32{&p.DIEPINT.U32, uint32(BNA)}}
}

func (p *USB_OTG_INEndpoint_Periph) PKTDRPSTS() RMDIEPINT {
	return RMDIEPINT{mmio.UM32{&p.DIEPINT.U32, uint32(PKTDRPSTS)}}
}

func (p *USB_OTG_INEndpoint_Periph) BERR() RMDIEPINT {
	return RMDIEPINT{mmio.UM32{&p.DIEPINT.U32, uint32(BERR)}}
}

func (p *USB_OTG_INEndpoint_Periph) NAK() RMDIEPINT {
	return RMDIEPINT{mmio.UM32{&p.DIEPINT.U32, uint32(NAK)}}
}

type DIEPTSIZ uint32

func (b DIEPTSIZ) Field(mask DIEPTSIZ) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DIEPTSIZ) J(v int) DIEPTSIZ {
	return DIEPTSIZ(bits.Make32(v, uint32(mask)))
}

type RDIEPTSIZ struct{ mmio.U32 }

func (r *RDIEPTSIZ) Bits(mask DIEPTSIZ) DIEPTSIZ { return DIEPTSIZ(r.U32.Bits(uint32(mask))) }
func (r *RDIEPTSIZ) StoreBits(mask, b DIEPTSIZ)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDIEPTSIZ) SetBits(mask DIEPTSIZ)       { r.U32.SetBits(uint32(mask)) }
func (r *RDIEPTSIZ) ClearBits(mask DIEPTSIZ)     { r.U32.ClearBits(uint32(mask)) }
func (r *RDIEPTSIZ) Load() DIEPTSIZ              { return DIEPTSIZ(r.U32.Load()) }
func (r *RDIEPTSIZ) Store(b DIEPTSIZ)            { r.U32.Store(uint32(b)) }

func (r *RDIEPTSIZ) AtomicStoreBits(mask, b DIEPTSIZ) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDIEPTSIZ) AtomicSetBits(mask DIEPTSIZ)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDIEPTSIZ) AtomicClearBits(mask DIEPTSIZ)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDIEPTSIZ struct{ mmio.UM32 }

func (rm RMDIEPTSIZ) Load() DIEPTSIZ   { return DIEPTSIZ(rm.UM32.Load()) }
func (rm RMDIEPTSIZ) Store(b DIEPTSIZ) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_INEndpoint_Periph) XFRSIZ() RMDIEPTSIZ {
	return RMDIEPTSIZ{mmio.UM32{&p.DIEPTSIZ.U32, uint32(XFRSIZ)}}
}

func (p *USB_OTG_INEndpoint_Periph) PKTCNT() RMDIEPTSIZ {
	return RMDIEPTSIZ{mmio.UM32{&p.DIEPTSIZ.U32, uint32(PKTCNT)}}
}

func (p *USB_OTG_INEndpoint_Periph) MULCNT() RMDIEPTSIZ {
	return RMDIEPTSIZ{mmio.UM32{&p.DIEPTSIZ.U32, uint32(MULCNT)}}
}

type DIEPDMA uint32

func (b DIEPDMA) Field(mask DIEPDMA) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DIEPDMA) J(v int) DIEPDMA {
	return DIEPDMA(bits.Make32(v, uint32(mask)))
}

type RDIEPDMA struct{ mmio.U32 }

func (r *RDIEPDMA) Bits(mask DIEPDMA) DIEPDMA { return DIEPDMA(r.U32.Bits(uint32(mask))) }
func (r *RDIEPDMA) StoreBits(mask, b DIEPDMA) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDIEPDMA) SetBits(mask DIEPDMA)      { r.U32.SetBits(uint32(mask)) }
func (r *RDIEPDMA) ClearBits(mask DIEPDMA)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDIEPDMA) Load() DIEPDMA             { return DIEPDMA(r.U32.Load()) }
func (r *RDIEPDMA) Store(b DIEPDMA)           { r.U32.Store(uint32(b)) }

func (r *RDIEPDMA) AtomicStoreBits(mask, b DIEPDMA) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDIEPDMA) AtomicSetBits(mask DIEPDMA)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDIEPDMA) AtomicClearBits(mask DIEPDMA)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDIEPDMA struct{ mmio.UM32 }

func (rm RMDIEPDMA) Load() DIEPDMA   { return DIEPDMA(rm.UM32.Load()) }
func (rm RMDIEPDMA) Store(b DIEPDMA) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_INEndpoint_Periph) DMAADDR() RMDIEPDMA {
	return RMDIEPDMA{mmio.UM32{&p.DIEPDMA.U32, uint32(DMAADDR)}}
}

type DTXFSTS uint32

func (b DTXFSTS) Field(mask DTXFSTS) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DTXFSTS) J(v int) DTXFSTS {
	return DTXFSTS(bits.Make32(v, uint32(mask)))
}

type RDTXFSTS struct{ mmio.U32 }

func (r *RDTXFSTS) Bits(mask DTXFSTS) DTXFSTS { return DTXFSTS(r.U32.Bits(uint32(mask))) }
func (r *RDTXFSTS) StoreBits(mask, b DTXFSTS) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDTXFSTS) SetBits(mask DTXFSTS)      { r.U32.SetBits(uint32(mask)) }
func (r *RDTXFSTS) ClearBits(mask DTXFSTS)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDTXFSTS) Load() DTXFSTS             { return DTXFSTS(r.U32.Load()) }
func (r *RDTXFSTS) Store(b DTXFSTS)           { r.U32.Store(uint32(b)) }

func (r *RDTXFSTS) AtomicStoreBits(mask, b DTXFSTS) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDTXFSTS) AtomicSetBits(mask DTXFSTS)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDTXFSTS) AtomicClearBits(mask DTXFSTS)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDTXFSTS struct{ mmio.UM32 }

func (rm RMDTXFSTS) Load() DTXFSTS   { return DTXFSTS(rm.UM32.Load()) }
func (rm RMDTXFSTS) Store(b DTXFSTS) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_INEndpoint_Periph) INEPTFSAV() RMDTXFSTS {
	return RMDTXFSTS{mmio.UM32{&p.DTXFSTS.U32, uint32(INEPTFSAV)}}
}
