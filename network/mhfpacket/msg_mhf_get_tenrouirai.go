package mhfpacket

import (
	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfGetTenrouirai represents the MSG_MHF_GET_TENROUIRAI
type MsgMhfGetTenrouirai struct {
	AckHandle uint32
	Unk0      uint16
	Unk1      uint32
	Unk2      uint16
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfGetTenrouirai) Opcode() network.PacketID {
	return network.MSG_MHF_GET_TENROUIRAI
}

// Parse parses the packet from binary
func (m *MsgMhfGetTenrouirai) Parse(bf *byteframe.ByteFrame) error {
	m.AckHandle = bf.ReadUint32()
	m.Unk0 = bf.ReadUint16()
	m.Unk1 = bf.ReadUint32()
	m.Unk2 = bf.ReadUint16()
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfGetTenrouirai) Build(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}
