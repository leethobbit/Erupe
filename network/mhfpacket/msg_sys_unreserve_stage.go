package mhfpacket

import (
	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgSysUnreserveStage represents the MSG_SYS_UNRESERVE_STAGE
type MsgSysUnreserveStage struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysUnreserveStage) Opcode() network.PacketID {
	return network.MSG_SYS_UNRESERVE_STAGE
}

// Parse parses the packet from binary
func (m *MsgSysUnreserveStage) Parse(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgSysUnreserveStage) Build(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}
