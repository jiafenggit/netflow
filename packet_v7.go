package netflow

import (
	"fmt"
	"io"
)

// V7Header is a NetFlow version 7 (Catalyst 5000) header
//
// As specified at http://www.cisco.com/c/en/us/td/docs/net_mgmt/netflow_collection_engine/3-6/user/guide/format.html#wp1007543
type V7Header struct {
	V1Header
	FlowSequence uint32
	Reserved     uint32
}

func (h *V7Header) GetVersion() uint16 {
	return h.Version
}

func (h *V7Header) SetVersion(v uint16) {
	h.Version = v
}

// V7FlowRecord is a NetFlow version 7 (Catalyst 5000) Flow Record Format
//
// As specified at http://www.cisco.com/c/en/us/td/docs/net_mgmt/netflow_collection_engine/3-6/user/guide/format.html#wp1007604
type V7FlowRecord struct {
	// SrcAddr is the Source IP address
	SrcAddr LongIPv4
	// DstAddr is the Destination IP address
	DstAddr LongIPv4
	// NextHop is the IP address of next hop router
	NextHop LongIPv4
	// Input is the SNMP index of input interface
	Input uint16
	// Output is the SNMP index of output interface
	Output uint16
	// Packets is the number of packets in the flow
	Packets uint32
	// Octets is the number of bytes in the flow
	Octets uint32
	// First is the SysUptime at start of flow
	First uint32
	// Last is the SysUptime at end of flow
	Last uint32
	// SrcPort is the TCP/UDP source port number or equivalent
	SrcPort uint16
	// DstPort is the TCP/UDP destination port number or equivalent
	DstPort uint16
	// Pad0 are unused bytes
	Pad0 uint16
	// TCPFlags are the TCP header flags
	TCPFlags uint8
	// Protocol number (IP)
	Protocol uint8
	// ToS is the IP type of service
	ToS uint8
	// SrcAS is the Autonomous system number of the source, either origin or peer
	SrcAS uint16
	// DstAS is the Autonomous system number of the destination, either origin or peer
	DstAS uint16
	// SrcMask are the source address prefix mask bits
	SrcMask uint8
	// DstMask are the destination address prefix mask bits
	DstMask uint8
	// Flags indicating, among other things, what flows are invalid
	Flags uint16
	// RouterSC is the IP address of the router that is bypassed by the Catalyst 5000 series switch. This is the same address the router uses when it sends NetFlow export packets. This IP address is propagated to all switches bypassing the router through the FCP protocol.
	RouterSC LongIPv4
}

func (r *V7FlowRecord) Bytes() []byte {
	return structPack(r)
}

func (r *V7FlowRecord) Len() int {
	return structLen(r)
}

func (r *V7FlowRecord) String() string {
	return fmt.Sprintf("%s/%d:%d -> %s/%d:%d", r.SrcAddr, r.SrcMask, r.SrcPort, r.DstAddr, r.DstMask, r.DstPort)
}

func (r *V7FlowRecord) Unmarshal(h io.Reader) error {
	var err error
	if r.SrcAddr, err = readLongIPv4(h); err != nil {
		return err
	}
	if r.DstAddr, err = readLongIPv4(h); err != nil {
		return err
	}
	if r.NextHop, err = readLongIPv4(h); err != nil {
		return err
	}
	if r.Input, err = readUint16(h); err != nil {
		return err
	}
	if r.Output, err = readUint16(h); err != nil {
		return err
	}
	if r.Packets, err = readUint32(h); err != nil {
		return err
	}
	if r.Octets, err = readUint32(h); err != nil {
		return err
	}
	if r.First, err = readUint32(h); err != nil {
		return err
	}
	if r.Last, err = readUint32(h); err != nil {
		return err
	}
	if r.SrcPort, err = readUint16(h); err != nil {
		return err
	}
	if r.DstPort, err = readUint16(h); err != nil {
		return err
	}
	if r.Pad0, err = readUint16(h); err != nil {
		return err
	}
	if r.Protocol, err = readUint8(h); err != nil {
		return err
	}
	if r.ToS, err = readUint8(h); err != nil {
		return err
	}
	if r.SrcAS, err = readUint16(h); err != nil {
		return err
	}
	if r.DstAS, err = readUint16(h); err != nil {
		return err
	}
	if r.SrcMask, err = readUint8(h); err != nil {
		return err
	}
	if r.DstMask, err = readUint8(h); err != nil {
		return err
	}
	if r.Flags, err = readUint16(h); err != nil {
		return err
	}
	if r.RouterSC, err = readLongIPv4(h); err != nil {
		return err
	}
	return nil
}
