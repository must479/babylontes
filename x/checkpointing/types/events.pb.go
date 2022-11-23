// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babylon/checkpointing/events.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type EventCheckpointAccumulating struct {
	Checkpoint *RawCheckpointWithMeta `protobuf:"bytes,1,opt,name=checkpoint,proto3" json:"checkpoint,omitempty"`
}

func (m *EventCheckpointAccumulating) Reset()         { *m = EventCheckpointAccumulating{} }
func (m *EventCheckpointAccumulating) String() string { return proto.CompactTextString(m) }
func (*EventCheckpointAccumulating) ProtoMessage()    {}
func (*EventCheckpointAccumulating) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d41a0fa2283f67f, []int{0}
}
func (m *EventCheckpointAccumulating) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCheckpointAccumulating) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCheckpointAccumulating.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCheckpointAccumulating) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCheckpointAccumulating.Merge(m, src)
}
func (m *EventCheckpointAccumulating) XXX_Size() int {
	return m.Size()
}
func (m *EventCheckpointAccumulating) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCheckpointAccumulating.DiscardUnknown(m)
}

var xxx_messageInfo_EventCheckpointAccumulating proto.InternalMessageInfo

func (m *EventCheckpointAccumulating) GetCheckpoint() *RawCheckpointWithMeta {
	if m != nil {
		return m.Checkpoint
	}
	return nil
}

type EventCheckpointSealed struct {
	Checkpoint *RawCheckpointWithMeta `protobuf:"bytes,1,opt,name=checkpoint,proto3" json:"checkpoint,omitempty"`
}

func (m *EventCheckpointSealed) Reset()         { *m = EventCheckpointSealed{} }
func (m *EventCheckpointSealed) String() string { return proto.CompactTextString(m) }
func (*EventCheckpointSealed) ProtoMessage()    {}
func (*EventCheckpointSealed) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d41a0fa2283f67f, []int{1}
}
func (m *EventCheckpointSealed) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCheckpointSealed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCheckpointSealed.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCheckpointSealed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCheckpointSealed.Merge(m, src)
}
func (m *EventCheckpointSealed) XXX_Size() int {
	return m.Size()
}
func (m *EventCheckpointSealed) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCheckpointSealed.DiscardUnknown(m)
}

var xxx_messageInfo_EventCheckpointSealed proto.InternalMessageInfo

func (m *EventCheckpointSealed) GetCheckpoint() *RawCheckpointWithMeta {
	if m != nil {
		return m.Checkpoint
	}
	return nil
}

type EventCheckpointSubmitted struct {
	Checkpoint *RawCheckpointWithMeta `protobuf:"bytes,1,opt,name=checkpoint,proto3" json:"checkpoint,omitempty"`
}

func (m *EventCheckpointSubmitted) Reset()         { *m = EventCheckpointSubmitted{} }
func (m *EventCheckpointSubmitted) String() string { return proto.CompactTextString(m) }
func (*EventCheckpointSubmitted) ProtoMessage()    {}
func (*EventCheckpointSubmitted) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d41a0fa2283f67f, []int{2}
}
func (m *EventCheckpointSubmitted) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCheckpointSubmitted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCheckpointSubmitted.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCheckpointSubmitted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCheckpointSubmitted.Merge(m, src)
}
func (m *EventCheckpointSubmitted) XXX_Size() int {
	return m.Size()
}
func (m *EventCheckpointSubmitted) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCheckpointSubmitted.DiscardUnknown(m)
}

var xxx_messageInfo_EventCheckpointSubmitted proto.InternalMessageInfo

func (m *EventCheckpointSubmitted) GetCheckpoint() *RawCheckpointWithMeta {
	if m != nil {
		return m.Checkpoint
	}
	return nil
}

type EventCheckpointConfirmed struct {
	Checkpoint *RawCheckpointWithMeta `protobuf:"bytes,1,opt,name=checkpoint,proto3" json:"checkpoint,omitempty"`
}

func (m *EventCheckpointConfirmed) Reset()         { *m = EventCheckpointConfirmed{} }
func (m *EventCheckpointConfirmed) String() string { return proto.CompactTextString(m) }
func (*EventCheckpointConfirmed) ProtoMessage()    {}
func (*EventCheckpointConfirmed) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d41a0fa2283f67f, []int{3}
}
func (m *EventCheckpointConfirmed) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCheckpointConfirmed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCheckpointConfirmed.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCheckpointConfirmed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCheckpointConfirmed.Merge(m, src)
}
func (m *EventCheckpointConfirmed) XXX_Size() int {
	return m.Size()
}
func (m *EventCheckpointConfirmed) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCheckpointConfirmed.DiscardUnknown(m)
}

var xxx_messageInfo_EventCheckpointConfirmed proto.InternalMessageInfo

func (m *EventCheckpointConfirmed) GetCheckpoint() *RawCheckpointWithMeta {
	if m != nil {
		return m.Checkpoint
	}
	return nil
}

type EventCheckpointFinalized struct {
	Checkpoint *RawCheckpointWithMeta `protobuf:"bytes,1,opt,name=checkpoint,proto3" json:"checkpoint,omitempty"`
}

func (m *EventCheckpointFinalized) Reset()         { *m = EventCheckpointFinalized{} }
func (m *EventCheckpointFinalized) String() string { return proto.CompactTextString(m) }
func (*EventCheckpointFinalized) ProtoMessage()    {}
func (*EventCheckpointFinalized) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d41a0fa2283f67f, []int{4}
}
func (m *EventCheckpointFinalized) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCheckpointFinalized) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCheckpointFinalized.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCheckpointFinalized) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCheckpointFinalized.Merge(m, src)
}
func (m *EventCheckpointFinalized) XXX_Size() int {
	return m.Size()
}
func (m *EventCheckpointFinalized) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCheckpointFinalized.DiscardUnknown(m)
}

var xxx_messageInfo_EventCheckpointFinalized proto.InternalMessageInfo

func (m *EventCheckpointFinalized) GetCheckpoint() *RawCheckpointWithMeta {
	if m != nil {
		return m.Checkpoint
	}
	return nil
}

type EventCheckpointForgotten struct {
	Checkpoint *RawCheckpointWithMeta `protobuf:"bytes,1,opt,name=checkpoint,proto3" json:"checkpoint,omitempty"`
}

func (m *EventCheckpointForgotten) Reset()         { *m = EventCheckpointForgotten{} }
func (m *EventCheckpointForgotten) String() string { return proto.CompactTextString(m) }
func (*EventCheckpointForgotten) ProtoMessage()    {}
func (*EventCheckpointForgotten) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d41a0fa2283f67f, []int{5}
}
func (m *EventCheckpointForgotten) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCheckpointForgotten) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCheckpointForgotten.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCheckpointForgotten) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCheckpointForgotten.Merge(m, src)
}
func (m *EventCheckpointForgotten) XXX_Size() int {
	return m.Size()
}
func (m *EventCheckpointForgotten) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCheckpointForgotten.DiscardUnknown(m)
}

var xxx_messageInfo_EventCheckpointForgotten proto.InternalMessageInfo

func (m *EventCheckpointForgotten) GetCheckpoint() *RawCheckpointWithMeta {
	if m != nil {
		return m.Checkpoint
	}
	return nil
}

func init() {
	proto.RegisterType((*EventCheckpointAccumulating)(nil), "babylon.checkpointing.v1.EventCheckpointAccumulating")
	proto.RegisterType((*EventCheckpointSealed)(nil), "babylon.checkpointing.v1.EventCheckpointSealed")
	proto.RegisterType((*EventCheckpointSubmitted)(nil), "babylon.checkpointing.v1.EventCheckpointSubmitted")
	proto.RegisterType((*EventCheckpointConfirmed)(nil), "babylon.checkpointing.v1.EventCheckpointConfirmed")
	proto.RegisterType((*EventCheckpointFinalized)(nil), "babylon.checkpointing.v1.EventCheckpointFinalized")
	proto.RegisterType((*EventCheckpointForgotten)(nil), "babylon.checkpointing.v1.EventCheckpointForgotten")
}

func init() {
	proto.RegisterFile("babylon/checkpointing/events.proto", fileDescriptor_9d41a0fa2283f67f)
}

var fileDescriptor_9d41a0fa2283f67f = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4a, 0x4a, 0x4c, 0xaa,
	0xcc, 0xc9, 0xcf, 0xd3, 0x4f, 0xce, 0x48, 0x4d, 0xce, 0x2e, 0xc8, 0xcf, 0xcc, 0x2b, 0xc9, 0xcc,
	0x4b, 0xd7, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92,
	0x80, 0xaa, 0xd1, 0x43, 0x51, 0xa3, 0x57, 0x66, 0x28, 0xa5, 0x86, 0x5d, 0x37, 0x82, 0x07, 0x31,
	0x41, 0x29, 0x8f, 0x4b, 0xda, 0x15, 0x64, 0xa2, 0x33, 0x5c, 0xc2, 0x31, 0x39, 0xb9, 0x34, 0xb7,
	0x34, 0x27, 0x11, 0xa4, 0x5e, 0xc8, 0x9f, 0x8b, 0x0b, 0xa1, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83,
	0xdb, 0x48, 0x5f, 0x0f, 0x97, 0xad, 0x7a, 0x41, 0x89, 0xe5, 0x08, 0x83, 0xc2, 0x33, 0x4b, 0x32,
	0x7c, 0x53, 0x4b, 0x12, 0x83, 0x90, 0x8c, 0x50, 0xca, 0xe0, 0x12, 0x45, 0xb3, 0x2f, 0x38, 0x35,
	0x31, 0x27, 0x35, 0x85, 0xfa, 0x36, 0x65, 0x73, 0x49, 0xa0, 0xdb, 0x54, 0x9a, 0x94, 0x9b, 0x59,
	0x52, 0x42, 0x1f, 0xcb, 0x9c, 0xf3, 0xf3, 0xd2, 0x32, 0x8b, 0x72, 0xe9, 0x63, 0x99, 0x5b, 0x66,
	0x5e, 0x62, 0x4e, 0x66, 0x15, 0x9d, 0x2c, 0xcb, 0x2f, 0x4a, 0xcf, 0x2f, 0x29, 0x49, 0xcd, 0xa3,
	0xba, 0x65, 0x4e, 0xfe, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c,
	0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x65, 0x9a,
	0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f, 0xb5, 0x20, 0x39, 0x23, 0x31,
	0x33, 0x0f, 0xc6, 0xd1, 0xaf, 0x40, 0x4b, 0xe9, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0xe0,
	0x54, 0x6e, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x76, 0x0b, 0x2a, 0x4d, 0x03, 0x00, 0x00,
}

func (m *EventCheckpointAccumulating) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCheckpointAccumulating) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCheckpointAccumulating) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Checkpoint != nil {
		{
			size, err := m.Checkpoint.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvents(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventCheckpointSealed) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCheckpointSealed) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCheckpointSealed) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Checkpoint != nil {
		{
			size, err := m.Checkpoint.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvents(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventCheckpointSubmitted) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCheckpointSubmitted) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCheckpointSubmitted) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Checkpoint != nil {
		{
			size, err := m.Checkpoint.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvents(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventCheckpointConfirmed) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCheckpointConfirmed) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCheckpointConfirmed) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Checkpoint != nil {
		{
			size, err := m.Checkpoint.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvents(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventCheckpointFinalized) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCheckpointFinalized) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCheckpointFinalized) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Checkpoint != nil {
		{
			size, err := m.Checkpoint.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvents(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventCheckpointForgotten) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCheckpointForgotten) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCheckpointForgotten) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Checkpoint != nil {
		{
			size, err := m.Checkpoint.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvents(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventCheckpointAccumulating) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Checkpoint != nil {
		l = m.Checkpoint.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventCheckpointSealed) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Checkpoint != nil {
		l = m.Checkpoint.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventCheckpointSubmitted) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Checkpoint != nil {
		l = m.Checkpoint.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventCheckpointConfirmed) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Checkpoint != nil {
		l = m.Checkpoint.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventCheckpointFinalized) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Checkpoint != nil {
		l = m.Checkpoint.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventCheckpointForgotten) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Checkpoint != nil {
		l = m.Checkpoint.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventCheckpointAccumulating) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventCheckpointAccumulating: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCheckpointAccumulating: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checkpoint", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Checkpoint == nil {
				m.Checkpoint = &RawCheckpointWithMeta{}
			}
			if err := m.Checkpoint.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventCheckpointSealed) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventCheckpointSealed: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCheckpointSealed: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checkpoint", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Checkpoint == nil {
				m.Checkpoint = &RawCheckpointWithMeta{}
			}
			if err := m.Checkpoint.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventCheckpointSubmitted) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventCheckpointSubmitted: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCheckpointSubmitted: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checkpoint", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Checkpoint == nil {
				m.Checkpoint = &RawCheckpointWithMeta{}
			}
			if err := m.Checkpoint.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventCheckpointConfirmed) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventCheckpointConfirmed: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCheckpointConfirmed: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checkpoint", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Checkpoint == nil {
				m.Checkpoint = &RawCheckpointWithMeta{}
			}
			if err := m.Checkpoint.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventCheckpointFinalized) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventCheckpointFinalized: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCheckpointFinalized: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checkpoint", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Checkpoint == nil {
				m.Checkpoint = &RawCheckpointWithMeta{}
			}
			if err := m.Checkpoint.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventCheckpointForgotten) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventCheckpointForgotten: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCheckpointForgotten: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checkpoint", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Checkpoint == nil {
				m.Checkpoint = &RawCheckpointWithMeta{}
			}
			if err := m.Checkpoint.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvents
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
			if length < 0 {
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
