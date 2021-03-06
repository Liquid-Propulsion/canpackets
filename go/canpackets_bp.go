// Code generated by bitproto. DO NOT EDIT.

package canpackets

import (
	"strconv"
	"encoding/json"

	bp "github.com/hit9/bitproto/lib/go"
)

// Avoid possible golang import not used error
var formatInt = strconv.FormatInt
var jsonMarshal = json.Marshal
var _ = bp.Useless

// The different types of nodes currently possible, will be exanded in the future.
type NodeType uint8 // 4bit

const (
	PRESSURE_TRANSDUCER_NODE NodeType = 0
	LOAD_CELL_NODE = 1
	THERMAL_COUPLE_NODE = 2
	SOLENOID_NODE = 3
)

func (m NodeType) BpProcessor() bp.Processor {
	return bp.NewEnumProcessor(bp.NewUint(4))
}

// String returns the name of this enum item.
func (v NodeType) String() string {
	switch v {
	case 0:
		return "PRESSURE_TRANSDUCER_NODE"
	case 1:
		return "LOAD_CELL_NODE"
	case 2:
		return "THERMAL_COUPLE_NODE"
	case 3:
		return "SOLENOID_NODE"
	default:
		return "NodeType(" + formatInt(int64(v), 10) + ")"
	}
}

// Must be repeated every 20ms, otherwise the solenoids will be closed. (Sent by the mainland.)
// ID: 0x01
type StagePacket struct {
	SolenoidState [64]bool `json:"solenoid_state"` // 64bit
}

// Number of bytes to serialize struct StagePacket
const BYTES_LENGTH_STAGE_PACKET uint32 = 8

func (m *StagePacket) Size() uint32 { return 8 }

// Returns string representation for struct StagePacket.
func (m *StagePacket) String() string {
	v, _ := jsonMarshal(m)
	return string(v)
}

// Encode struct StagePacket to bytes buffer.
func (m *StagePacket) Encode() []byte {
	ctx := bp.NewEncodeContext(int(m.Size()))
	m.BpProcessor().Process(ctx, nil, m)
	return ctx.Buffer()
}

func (m *StagePacket) Decode(s []byte) {
	ctx := bp.NewDecodeContext(s)
	m.BpProcessor().Process(ctx, nil, m)
}

func (m *StagePacket) BpProcessor() bp.Processor {
	fieldDescriptors := []*bp.MessageFieldProcessor{
		bp.NewMessageFieldProcessor(1, bp.NewArray(false, 64, bp.NewBool())),
	}
	return bp.NewMessageProcessor(false, 64, fieldDescriptors)
}

func (m *StagePacket) BpGetAccessor(di *bp.DataIndexer) bp.Accessor {
	switch di.F() {
	default:
		return nil  // Won't reached
	}
}

func (m *StagePacket) BpSetByte(di *bp.DataIndexer, lshift int, b byte) {
	switch di.F() {
		case 1:
			m.SolenoidState[di.I(0)] = bp.Byte2bool(b)
		default:
			return
	}
}

func (m *StagePacket) BpGetByte(di *bp.DataIndexer, rshift int) byte {
	switch di.F() {
		case 1:
			return bp.Bool2byte(m.SolenoidState[di.I(0)]) >> rshift
		default:
			return byte(0) // Won't reached
	}
}

// Must be repeated every 20ms, otherwise the power will be cut to all solenoids. (Sent by the mainland.)
// ID: 0x00
type PowerPacket struct {
	SystemPowered bool `json:"system_powered"` // 1bit
	Siren bool `json:"siren"` // 1bit
}

// Number of bytes to serialize struct PowerPacket
const BYTES_LENGTH_POWER_PACKET uint32 = 1

func (m *PowerPacket) Size() uint32 { return 1 }

// Returns string representation for struct PowerPacket.
func (m *PowerPacket) String() string {
	v, _ := jsonMarshal(m)
	return string(v)
}

// Encode struct PowerPacket to bytes buffer.
func (m *PowerPacket) Encode() []byte {
	ctx := bp.NewEncodeContext(int(m.Size()))
	m.BpProcessor().Process(ctx, nil, m)
	return ctx.Buffer()
}

func (m *PowerPacket) Decode(s []byte) {
	ctx := bp.NewDecodeContext(s)
	m.BpProcessor().Process(ctx, nil, m)
}

func (m *PowerPacket) BpProcessor() bp.Processor {
	fieldDescriptors := []*bp.MessageFieldProcessor{
		bp.NewMessageFieldProcessor(1, bp.NewBool()),
		bp.NewMessageFieldProcessor(2, bp.NewBool()),
	}
	return bp.NewMessageProcessor(false, 2, fieldDescriptors)
}

func (m *PowerPacket) BpGetAccessor(di *bp.DataIndexer) bp.Accessor {
	switch di.F() {
	default:
		return nil  // Won't reached
	}
}

func (m *PowerPacket) BpSetByte(di *bp.DataIndexer, lshift int, b byte) {
	switch di.F() {
		case 1:
			m.SystemPowered = bp.Byte2bool(b)
		case 2:
			m.Siren = bp.Byte2bool(b)
		default:
			return
	}
}

func (m *PowerPacket) BpGetByte(di *bp.DataIndexer, rshift int) byte {
	switch di.F() {
		case 1:
			return bp.Bool2byte(m.SystemPowered) >> rshift
		case 2:
			return bp.Bool2byte(m.Siren) >> rshift
		default:
			return byte(0) // Won't reached
	}
}

// Causes a Island node to blink it's USER LED for 5 seconds.
// ID: 0x06
type BlinkPacket struct {
	NodeId uint8 `json:"node_id"` // 8bit
}

// Number of bytes to serialize struct BlinkPacket
const BYTES_LENGTH_BLINK_PACKET uint32 = 1

func (m *BlinkPacket) Size() uint32 { return 1 }

// Returns string representation for struct BlinkPacket.
func (m *BlinkPacket) String() string {
	v, _ := jsonMarshal(m)
	return string(v)
}

// Encode struct BlinkPacket to bytes buffer.
func (m *BlinkPacket) Encode() []byte {
	ctx := bp.NewEncodeContext(int(m.Size()))
	m.BpProcessor().Process(ctx, nil, m)
	return ctx.Buffer()
}

func (m *BlinkPacket) Decode(s []byte) {
	ctx := bp.NewDecodeContext(s)
	m.BpProcessor().Process(ctx, nil, m)
}

func (m *BlinkPacket) BpProcessor() bp.Processor {
	fieldDescriptors := []*bp.MessageFieldProcessor{
		bp.NewMessageFieldProcessor(1, bp.NewUint(8)),
	}
	return bp.NewMessageProcessor(false, 8, fieldDescriptors)
}

func (m *BlinkPacket) BpGetAccessor(di *bp.DataIndexer) bp.Accessor {
	switch di.F() {
	default:
		return nil  // Won't reached
	}
}

func (m *BlinkPacket) BpSetByte(di *bp.DataIndexer, lshift int, b byte) {
	switch di.F() {
		case 1:
			m.NodeId |= (uint8(b) << lshift)
		default:
			return
	}
}

func (m *BlinkPacket) BpGetByte(di *bp.DataIndexer, rshift int) byte {
	switch di.F() {
		case 1:
			return byte(m.NodeId >> rshift)
		default:
			return byte(0) // Won't reached
	}
}

// ID: 0x03
type SensorDataPacket struct {
	SensorId uint8 `json:"sensor_id"` // 4bit
	SensorData uint32 `json:"sensor_data"` // 32bit
}

// Number of bytes to serialize struct SensorDataPacket
const BYTES_LENGTH_SENSOR_DATA_PACKET uint32 = 5

func (m *SensorDataPacket) Size() uint32 { return 5 }

// Returns string representation for struct SensorDataPacket.
func (m *SensorDataPacket) String() string {
	v, _ := jsonMarshal(m)
	return string(v)
}

// Encode struct SensorDataPacket to bytes buffer.
func (m *SensorDataPacket) Encode() []byte {
	ctx := bp.NewEncodeContext(int(m.Size()))
	m.BpProcessor().Process(ctx, nil, m)
	return ctx.Buffer()
}

func (m *SensorDataPacket) Decode(s []byte) {
	ctx := bp.NewDecodeContext(s)
	m.BpProcessor().Process(ctx, nil, m)
}

func (m *SensorDataPacket) BpProcessor() bp.Processor {
	fieldDescriptors := []*bp.MessageFieldProcessor{
		bp.NewMessageFieldProcessor(1, bp.NewUint(4)),
		bp.NewMessageFieldProcessor(2, bp.NewUint(32)),
	}
	return bp.NewMessageProcessor(false, 36, fieldDescriptors)
}

func (m *SensorDataPacket) BpGetAccessor(di *bp.DataIndexer) bp.Accessor {
	switch di.F() {
	default:
		return nil  // Won't reached
	}
}

func (m *SensorDataPacket) BpSetByte(di *bp.DataIndexer, lshift int, b byte) {
	switch di.F() {
		case 1:
			m.SensorId |= (uint8(b) << lshift)
		case 2:
			m.SensorData |= (uint32(b) << lshift)
		default:
			return
	}
}

func (m *SensorDataPacket) BpGetByte(di *bp.DataIndexer, rshift int) byte {
	switch di.F() {
		case 1:
			return byte(m.SensorId >> rshift)
		case 2:
			return byte(m.SensorData >> rshift)
		default:
			return byte(0) // Won't reached
	}
}

// Returned by all Island nodes, identifying their ID and their sensor type.
// ID: 0x05
type PongPacket struct {
	NodeId uint8 `json:"node_id"` // 8bit
	NodeType NodeType `json:"node_type"` // 4bit
}

// Number of bytes to serialize struct PongPacket
const BYTES_LENGTH_PONG_PACKET uint32 = 2

func (m *PongPacket) Size() uint32 { return 2 }

// Returns string representation for struct PongPacket.
func (m *PongPacket) String() string {
	v, _ := jsonMarshal(m)
	return string(v)
}

// Encode struct PongPacket to bytes buffer.
func (m *PongPacket) Encode() []byte {
	ctx := bp.NewEncodeContext(int(m.Size()))
	m.BpProcessor().Process(ctx, nil, m)
	return ctx.Buffer()
}

func (m *PongPacket) Decode(s []byte) {
	ctx := bp.NewDecodeContext(s)
	m.BpProcessor().Process(ctx, nil, m)
}

func (m *PongPacket) BpProcessor() bp.Processor {
	fieldDescriptors := []*bp.MessageFieldProcessor{
		bp.NewMessageFieldProcessor(1, bp.NewUint(8)),
		bp.NewMessageFieldProcessor(2, (NodeType(0)).BpProcessor()),
	}
	return bp.NewMessageProcessor(false, 12, fieldDescriptors)
}

func (m *PongPacket) BpGetAccessor(di *bp.DataIndexer) bp.Accessor {
	switch di.F() {
	default:
		return nil  // Won't reached
	}
}

func (m *PongPacket) BpSetByte(di *bp.DataIndexer, lshift int, b byte) {
	switch di.F() {
		case 1:
			m.NodeId |= (uint8(b) << lshift)
		case 2:
			m.NodeType |= (NodeType(b) << lshift)
		default:
			return
	}
}

func (m *PongPacket) BpGetByte(di *bp.DataIndexer, rshift int) byte {
	switch di.F() {
		case 1:
			return byte(m.NodeId >> rshift)
		case 2:
			return byte(m.NodeType >> rshift)
		default:
			return byte(0) // Won't reached
	}
}