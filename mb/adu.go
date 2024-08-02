package mb

type pdu struct {
	fc     uint8
	data   []byte
	length uint8
}

type ADU struct {
	pdu
	addr uint8
}

func (a ADU) Build() ([]byte, error) {
	bb := append([]byte{a.addr, a.fc, a.length}, a.data...)
	var crc Crc
	crc.Init()
	crc.Add(bb)
	bb = append(bb, crc.Value()...)
	return bb, nil
}

func NewADU(adr uint8, fc uint8, data int16) ADU {
	adu := ADU{
		pdu: pdu{
			fc:     fc,
			data:   []byte{byte(data >> 8), byte(data)},
			length: 2,
		},
		addr: adr,
	}
	return adu
}
