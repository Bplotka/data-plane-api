// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: capture.proto

/*
	Package v2 is a generated protocol buffer package.

	It is generated from these files:
		capture.proto

	It has these top-level messages:
		Connection
		Event
		Trace
*/
package v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_api_v2_core "envoy/api/v2/core"
import google_protobuf3 "google/protobuf"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Connection properties.
type Connection struct {
	// Global unique connection ID for Envoy session. Matches connection IDs used
	// in Envoy logs.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Local address.
	LocalAddress *envoy_api_v2_core.Address `protobuf:"bytes,2,opt,name=local_address,json=localAddress" json:"local_address,omitempty"`
	// Remote address.
	RemoteAddress *envoy_api_v2_core.Address `protobuf:"bytes,3,opt,name=remote_address,json=remoteAddress" json:"remote_address,omitempty"`
}

func (m *Connection) Reset()                    { *m = Connection{} }
func (m *Connection) String() string            { return proto.CompactTextString(m) }
func (*Connection) ProtoMessage()               {}
func (*Connection) Descriptor() ([]byte, []int) { return fileDescriptorCapture, []int{0} }

func (m *Connection) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Connection) GetLocalAddress() *envoy_api_v2_core.Address {
	if m != nil {
		return m.LocalAddress
	}
	return nil
}

func (m *Connection) GetRemoteAddress() *envoy_api_v2_core.Address {
	if m != nil {
		return m.RemoteAddress
	}
	return nil
}

// Event in a capture trace.
type Event struct {
	// Timestamp for event.
	Timestamp *google_protobuf3.Timestamp `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	// Read or write with content as bytes string.
	//
	// Types that are valid to be assigned to EventSelector:
	//	*Event_Read_
	//	*Event_Write_
	EventSelector isEvent_EventSelector `protobuf_oneof:"event_selector"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptorCapture, []int{1} }

type isEvent_EventSelector interface {
	isEvent_EventSelector()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Event_Read_ struct {
	Read *Event_Read `protobuf:"bytes,2,opt,name=read,oneof"`
}
type Event_Write_ struct {
	Write *Event_Write `protobuf:"bytes,3,opt,name=write,oneof"`
}

func (*Event_Read_) isEvent_EventSelector()  {}
func (*Event_Write_) isEvent_EventSelector() {}

func (m *Event) GetEventSelector() isEvent_EventSelector {
	if m != nil {
		return m.EventSelector
	}
	return nil
}

func (m *Event) GetTimestamp() *google_protobuf3.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Event) GetRead() *Event_Read {
	if x, ok := m.GetEventSelector().(*Event_Read_); ok {
		return x.Read
	}
	return nil
}

func (m *Event) GetWrite() *Event_Write {
	if x, ok := m.GetEventSelector().(*Event_Write_); ok {
		return x.Write
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Event) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Event_OneofMarshaler, _Event_OneofUnmarshaler, _Event_OneofSizer, []interface{}{
		(*Event_Read_)(nil),
		(*Event_Write_)(nil),
	}
}

func _Event_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Event)
	// event_selector
	switch x := m.EventSelector.(type) {
	case *Event_Read_:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Read); err != nil {
			return err
		}
	case *Event_Write_:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Write); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Event.EventSelector has unexpected type %T", x)
	}
	return nil
}

func _Event_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Event)
	switch tag {
	case 2: // event_selector.read
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Event_Read)
		err := b.DecodeMessage(msg)
		m.EventSelector = &Event_Read_{msg}
		return true, err
	case 3: // event_selector.write
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Event_Write)
		err := b.DecodeMessage(msg)
		m.EventSelector = &Event_Write_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Event_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Event)
	// event_selector
	switch x := m.EventSelector.(type) {
	case *Event_Read_:
		s := proto.Size(x.Read)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_Write_:
		s := proto.Size(x.Write)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Data read by Envoy from the transport socket.
type Event_Read struct {
	// Binary data read.
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Event_Read) Reset()                    { *m = Event_Read{} }
func (m *Event_Read) String() string            { return proto.CompactTextString(m) }
func (*Event_Read) ProtoMessage()               {}
func (*Event_Read) Descriptor() ([]byte, []int) { return fileDescriptorCapture, []int{1, 0} }

func (m *Event_Read) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// Data written by Envoy to the transport socket.
type Event_Write struct {
	// Binary data written.
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// Stream was half closed after this write.
	EndStream bool `protobuf:"varint,2,opt,name=end_stream,json=endStream,proto3" json:"end_stream,omitempty"`
}

func (m *Event_Write) Reset()                    { *m = Event_Write{} }
func (m *Event_Write) String() string            { return proto.CompactTextString(m) }
func (*Event_Write) ProtoMessage()               {}
func (*Event_Write) Descriptor() ([]byte, []int) { return fileDescriptorCapture, []int{1, 1} }

func (m *Event_Write) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Event_Write) GetEndStream() bool {
	if m != nil {
		return m.EndStream
	}
	return false
}

// Sequence of read/write events that constitute a captured trace on a socket.
// Multiple Trace messages might be emitted for a given connection ID, with the
// sink (e.g. file set, network) responsible for later reassembly.
type Trace struct {
	// Connection properties.
	Connection *Connection `protobuf:"bytes,1,opt,name=connection" json:"connection,omitempty"`
	// Sequence of observed events.
	Events []*Event `protobuf:"bytes,2,rep,name=events" json:"events,omitempty"`
}

func (m *Trace) Reset()                    { *m = Trace{} }
func (m *Trace) String() string            { return proto.CompactTextString(m) }
func (*Trace) ProtoMessage()               {}
func (*Trace) Descriptor() ([]byte, []int) { return fileDescriptorCapture, []int{2} }

func (m *Trace) GetConnection() *Connection {
	if m != nil {
		return m.Connection
	}
	return nil
}

func (m *Trace) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func init() {
	proto.RegisterType((*Connection)(nil), "envoy.data.tap.v2alpha.Connection")
	proto.RegisterType((*Event)(nil), "envoy.data.tap.v2alpha.Event")
	proto.RegisterType((*Event_Read)(nil), "envoy.data.tap.v2alpha.Event.Read")
	proto.RegisterType((*Event_Write)(nil), "envoy.data.tap.v2alpha.Event.Write")
	proto.RegisterType((*Trace)(nil), "envoy.data.tap.v2alpha.Trace")
}
func (m *Connection) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Connection) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintCapture(dAtA, i, uint64(m.Id))
	}
	if m.LocalAddress != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCapture(dAtA, i, uint64(m.LocalAddress.Size()))
		n1, err := m.LocalAddress.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.RemoteAddress != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCapture(dAtA, i, uint64(m.RemoteAddress.Size()))
		n2, err := m.RemoteAddress.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *Event) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Event) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Timestamp != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCapture(dAtA, i, uint64(m.Timestamp.Size()))
		n3, err := m.Timestamp.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.EventSelector != nil {
		nn4, err := m.EventSelector.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn4
	}
	return i, nil
}

func (m *Event_Read_) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Read != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCapture(dAtA, i, uint64(m.Read.Size()))
		n5, err := m.Read.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	return i, nil
}
func (m *Event_Write_) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Write != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCapture(dAtA, i, uint64(m.Write.Size()))
		n6, err := m.Write.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	return i, nil
}
func (m *Event_Read) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Event_Read) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCapture(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	return i, nil
}

func (m *Event_Write) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Event_Write) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCapture(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	if m.EndStream {
		dAtA[i] = 0x10
		i++
		if m.EndStream {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *Trace) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Trace) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Connection != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCapture(dAtA, i, uint64(m.Connection.Size()))
		n7, err := m.Connection.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	if len(m.Events) > 0 {
		for _, msg := range m.Events {
			dAtA[i] = 0x12
			i++
			i = encodeVarintCapture(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintCapture(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Connection) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovCapture(uint64(m.Id))
	}
	if m.LocalAddress != nil {
		l = m.LocalAddress.Size()
		n += 1 + l + sovCapture(uint64(l))
	}
	if m.RemoteAddress != nil {
		l = m.RemoteAddress.Size()
		n += 1 + l + sovCapture(uint64(l))
	}
	return n
}

func (m *Event) Size() (n int) {
	var l int
	_ = l
	if m.Timestamp != nil {
		l = m.Timestamp.Size()
		n += 1 + l + sovCapture(uint64(l))
	}
	if m.EventSelector != nil {
		n += m.EventSelector.Size()
	}
	return n
}

func (m *Event_Read_) Size() (n int) {
	var l int
	_ = l
	if m.Read != nil {
		l = m.Read.Size()
		n += 1 + l + sovCapture(uint64(l))
	}
	return n
}
func (m *Event_Write_) Size() (n int) {
	var l int
	_ = l
	if m.Write != nil {
		l = m.Write.Size()
		n += 1 + l + sovCapture(uint64(l))
	}
	return n
}
func (m *Event_Read) Size() (n int) {
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovCapture(uint64(l))
	}
	return n
}

func (m *Event_Write) Size() (n int) {
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovCapture(uint64(l))
	}
	if m.EndStream {
		n += 2
	}
	return n
}

func (m *Trace) Size() (n int) {
	var l int
	_ = l
	if m.Connection != nil {
		l = m.Connection.Size()
		n += 1 + l + sovCapture(uint64(l))
	}
	if len(m.Events) > 0 {
		for _, e := range m.Events {
			l = e.Size()
			n += 1 + l + sovCapture(uint64(l))
		}
	}
	return n
}

func sovCapture(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCapture(x uint64) (n int) {
	return sovCapture(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Connection) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCapture
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Connection: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Connection: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCapture
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LocalAddress", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCapture
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCapture
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LocalAddress == nil {
				m.LocalAddress = &envoy_api_v2_core.Address{}
			}
			if err := m.LocalAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemoteAddress", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCapture
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCapture
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RemoteAddress == nil {
				m.RemoteAddress = &envoy_api_v2_core.Address{}
			}
			if err := m.RemoteAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCapture(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCapture
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Event) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCapture
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Event: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Event: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCapture
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCapture
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timestamp == nil {
				m.Timestamp = &google_protobuf3.Timestamp{}
			}
			if err := m.Timestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Read", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCapture
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCapture
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Event_Read{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.EventSelector = &Event_Read_{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Write", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCapture
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCapture
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Event_Write{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.EventSelector = &Event_Write_{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCapture(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCapture
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Event_Read) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCapture
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Read: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Read: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCapture
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCapture
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCapture(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCapture
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Event_Write) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCapture
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Write: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Write: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCapture
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCapture
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndStream", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCapture
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EndStream = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipCapture(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCapture
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Trace) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCapture
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Trace: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Trace: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Connection", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCapture
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCapture
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Connection == nil {
				m.Connection = &Connection{}
			}
			if err := m.Connection.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCapture
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCapture
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Events = append(m.Events, &Event{})
			if err := m.Events[len(m.Events)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCapture(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCapture
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCapture(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCapture
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCapture
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCapture
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCapture
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCapture
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCapture(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCapture = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCapture   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("capture.proto", fileDescriptorCapture) }

var fileDescriptorCapture = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x4f, 0xcb, 0xd3, 0x40,
	0x10, 0xc6, 0xdf, 0xcd, 0x9b, 0x14, 0x3b, 0xfd, 0x43, 0xd9, 0x83, 0x94, 0x40, 0xff, 0x10, 0x2f,
	0x3d, 0x6d, 0x20, 0x22, 0x14, 0x3d, 0x48, 0x2b, 0x42, 0xcf, 0x6b, 0x41, 0xf0, 0x52, 0xb6, 0xd9,
	0xb1, 0x06, 0x92, 0x6c, 0xd8, 0x6c, 0x23, 0x5e, 0xfd, 0x24, 0xe2, 0xa7, 0xf1, 0xe8, 0x47, 0x90,
	0x7e, 0x12, 0xc9, 0x26, 0x6d, 0x3d, 0xd4, 0xbe, 0xb7, 0xcc, 0xe6, 0xf9, 0x3d, 0xf3, 0xcc, 0x0c,
	0x0c, 0x62, 0x51, 0x98, 0xa3, 0x46, 0x56, 0x68, 0x65, 0x14, 0x7d, 0x8e, 0x79, 0xa5, 0xbe, 0x31,
	0x29, 0x8c, 0x60, 0x46, 0x14, 0xac, 0x8a, 0x44, 0x5a, 0x7c, 0x11, 0xfe, 0xcc, 0xbe, 0x87, 0xa2,
	0x48, 0xc2, 0x2a, 0x0a, 0x63, 0xa5, 0x31, 0x14, 0x52, 0x6a, 0x2c, 0xcb, 0x06, 0xf4, 0x67, 0x07,
	0xa5, 0x0e, 0x29, 0x86, 0xb6, 0xda, 0x1f, 0x3f, 0x87, 0x26, 0xc9, 0xb0, 0x34, 0x22, 0x2b, 0x1a,
	0x41, 0xf0, 0x83, 0x00, 0xbc, 0x53, 0x79, 0x8e, 0xb1, 0x49, 0x54, 0x4e, 0x87, 0xe0, 0x24, 0x72,
	0x4c, 0xe6, 0x64, 0xe1, 0x72, 0x27, 0x91, 0xf4, 0x2d, 0x0c, 0x52, 0x15, 0x8b, 0x74, 0xd7, 0xda,
	0x8e, 0x9d, 0x39, 0x59, 0xf4, 0x22, 0x9f, 0x35, 0x81, 0x44, 0x91, 0xb0, 0x2a, 0x62, 0x75, 0x63,
	0xb6, 0x6a, 0x14, 0xbc, 0x6f, 0x81, 0xb6, 0xa2, 0x2b, 0x18, 0x6a, 0xcc, 0x94, 0xc1, 0x8b, 0xc3,
	0xe3, 0x93, 0x0e, 0x83, 0x86, 0x68, 0xcb, 0xe0, 0xa7, 0x03, 0xde, 0xfb, 0x0a, 0x73, 0x43, 0x97,
	0xd0, 0xbd, 0xe4, 0xb7, 0x21, 0x6b, 0x9f, 0x66, 0x42, 0x76, 0x9e, 0x90, 0x6d, 0xcf, 0x0a, 0x7e,
	0x15, 0xd3, 0x25, 0xb8, 0x1a, 0x85, 0x6c, 0xe3, 0x07, 0xec, 0xf6, 0x3e, 0x99, 0x6d, 0xc3, 0x38,
	0x0a, 0xb9, 0x79, 0xe0, 0x96, 0xa0, 0x6f, 0xc0, 0xfb, 0xaa, 0x13, 0x83, 0x6d, 0xee, 0x17, 0xf7,
	0xd1, 0x8f, 0xb5, 0x74, 0xf3, 0xc0, 0x1b, 0xc6, 0xf7, 0xc1, 0xad, 0xcd, 0x28, 0x05, 0xb7, 0x06,
	0x6c, 0xe6, 0x3e, 0xb7, 0xdf, 0xfe, 0x6b, 0xf0, 0xac, 0xfa, 0xd6, 0x4f, 0x3a, 0x01, 0xc0, 0x5c,
	0xee, 0x4a, 0xa3, 0x51, 0x64, 0x36, 0xf5, 0x33, 0xde, 0xc5, 0x5c, 0x7e, 0xb0, 0x0f, 0xeb, 0x11,
	0x0c, 0xb1, 0xee, 0xb7, 0x2b, 0x31, 0xc5, 0xd8, 0x28, 0x1d, 0x7c, 0x27, 0xe0, 0x6d, 0xb5, 0x88,
	0x91, 0xae, 0x01, 0xe2, 0xcb, 0x41, 0xdb, 0x2d, 0xfd, 0x77, 0xe0, 0xeb, 0xe9, 0xf9, 0x3f, 0x14,
	0x7d, 0x05, 0x1d, 0xeb, 0x5f, 0xdf, 0xfb, 0x71, 0xd1, 0x8b, 0x26, 0x77, 0xa7, 0xe6, 0xad, 0x78,
	0x3d, 0xfa, 0x75, 0x9a, 0x92, 0xdf, 0xa7, 0x29, 0xf9, 0x73, 0x9a, 0x92, 0x4f, 0x4e, 0x15, 0xed,
	0x3b, 0xf6, 0x2c, 0x2f, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x4e, 0x03, 0x75, 0xd0, 0x02,
	0x00, 0x00,
}
