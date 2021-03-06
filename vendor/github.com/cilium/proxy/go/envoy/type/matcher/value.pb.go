// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/matcher/value.proto

package matcher

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/lyft/protoc-gen-validate/validate"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Specifies the way to match a ProtobufWkt::Value. Primitive values and ListValue are supported.
// StructValue is not supported and is always not matched.
type ValueMatcher struct {
	// Specifies how to match a value.
	//
	// Types that are valid to be assigned to MatchPattern:
	//	*ValueMatcher_NullMatch_
	//	*ValueMatcher_DoubleMatch
	//	*ValueMatcher_StringMatch
	//	*ValueMatcher_BoolMatch
	//	*ValueMatcher_PresentMatch
	//	*ValueMatcher_ListMatch
	MatchPattern         isValueMatcher_MatchPattern `protobuf_oneof:"match_pattern"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ValueMatcher) Reset()         { *m = ValueMatcher{} }
func (m *ValueMatcher) String() string { return proto.CompactTextString(m) }
func (*ValueMatcher) ProtoMessage()    {}
func (*ValueMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_145b36501d266253, []int{0}
}

func (m *ValueMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValueMatcher.Unmarshal(m, b)
}
func (m *ValueMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValueMatcher.Marshal(b, m, deterministic)
}
func (m *ValueMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValueMatcher.Merge(m, src)
}
func (m *ValueMatcher) XXX_Size() int {
	return xxx_messageInfo_ValueMatcher.Size(m)
}
func (m *ValueMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_ValueMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_ValueMatcher proto.InternalMessageInfo

type isValueMatcher_MatchPattern interface {
	isValueMatcher_MatchPattern()
}

type ValueMatcher_NullMatch_ struct {
	NullMatch *ValueMatcher_NullMatch `protobuf:"bytes,1,opt,name=null_match,json=nullMatch,proto3,oneof"`
}

type ValueMatcher_DoubleMatch struct {
	DoubleMatch *DoubleMatcher `protobuf:"bytes,2,opt,name=double_match,json=doubleMatch,proto3,oneof"`
}

type ValueMatcher_StringMatch struct {
	StringMatch *StringMatcher `protobuf:"bytes,3,opt,name=string_match,json=stringMatch,proto3,oneof"`
}

type ValueMatcher_BoolMatch struct {
	BoolMatch bool `protobuf:"varint,4,opt,name=bool_match,json=boolMatch,proto3,oneof"`
}

type ValueMatcher_PresentMatch struct {
	PresentMatch bool `protobuf:"varint,5,opt,name=present_match,json=presentMatch,proto3,oneof"`
}

type ValueMatcher_ListMatch struct {
	ListMatch *ListMatcher `protobuf:"bytes,6,opt,name=list_match,json=listMatch,proto3,oneof"`
}

func (*ValueMatcher_NullMatch_) isValueMatcher_MatchPattern() {}

func (*ValueMatcher_DoubleMatch) isValueMatcher_MatchPattern() {}

func (*ValueMatcher_StringMatch) isValueMatcher_MatchPattern() {}

func (*ValueMatcher_BoolMatch) isValueMatcher_MatchPattern() {}

func (*ValueMatcher_PresentMatch) isValueMatcher_MatchPattern() {}

func (*ValueMatcher_ListMatch) isValueMatcher_MatchPattern() {}

func (m *ValueMatcher) GetMatchPattern() isValueMatcher_MatchPattern {
	if m != nil {
		return m.MatchPattern
	}
	return nil
}

func (m *ValueMatcher) GetNullMatch() *ValueMatcher_NullMatch {
	if x, ok := m.GetMatchPattern().(*ValueMatcher_NullMatch_); ok {
		return x.NullMatch
	}
	return nil
}

func (m *ValueMatcher) GetDoubleMatch() *DoubleMatcher {
	if x, ok := m.GetMatchPattern().(*ValueMatcher_DoubleMatch); ok {
		return x.DoubleMatch
	}
	return nil
}

func (m *ValueMatcher) GetStringMatch() *StringMatcher {
	if x, ok := m.GetMatchPattern().(*ValueMatcher_StringMatch); ok {
		return x.StringMatch
	}
	return nil
}

func (m *ValueMatcher) GetBoolMatch() bool {
	if x, ok := m.GetMatchPattern().(*ValueMatcher_BoolMatch); ok {
		return x.BoolMatch
	}
	return false
}

func (m *ValueMatcher) GetPresentMatch() bool {
	if x, ok := m.GetMatchPattern().(*ValueMatcher_PresentMatch); ok {
		return x.PresentMatch
	}
	return false
}

func (m *ValueMatcher) GetListMatch() *ListMatcher {
	if x, ok := m.GetMatchPattern().(*ValueMatcher_ListMatch); ok {
		return x.ListMatch
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ValueMatcher) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ValueMatcher_OneofMarshaler, _ValueMatcher_OneofUnmarshaler, _ValueMatcher_OneofSizer, []interface{}{
		(*ValueMatcher_NullMatch_)(nil),
		(*ValueMatcher_DoubleMatch)(nil),
		(*ValueMatcher_StringMatch)(nil),
		(*ValueMatcher_BoolMatch)(nil),
		(*ValueMatcher_PresentMatch)(nil),
		(*ValueMatcher_ListMatch)(nil),
	}
}

func _ValueMatcher_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ValueMatcher)
	// match_pattern
	switch x := m.MatchPattern.(type) {
	case *ValueMatcher_NullMatch_:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NullMatch); err != nil {
			return err
		}
	case *ValueMatcher_DoubleMatch:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DoubleMatch); err != nil {
			return err
		}
	case *ValueMatcher_StringMatch:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StringMatch); err != nil {
			return err
		}
	case *ValueMatcher_BoolMatch:
		t := uint64(0)
		if x.BoolMatch {
			t = 1
		}
		b.EncodeVarint(4<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *ValueMatcher_PresentMatch:
		t := uint64(0)
		if x.PresentMatch {
			t = 1
		}
		b.EncodeVarint(5<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *ValueMatcher_ListMatch:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ListMatch); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ValueMatcher.MatchPattern has unexpected type %T", x)
	}
	return nil
}

func _ValueMatcher_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ValueMatcher)
	switch tag {
	case 1: // match_pattern.null_match
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ValueMatcher_NullMatch)
		err := b.DecodeMessage(msg)
		m.MatchPattern = &ValueMatcher_NullMatch_{msg}
		return true, err
	case 2: // match_pattern.double_match
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DoubleMatcher)
		err := b.DecodeMessage(msg)
		m.MatchPattern = &ValueMatcher_DoubleMatch{msg}
		return true, err
	case 3: // match_pattern.string_match
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StringMatcher)
		err := b.DecodeMessage(msg)
		m.MatchPattern = &ValueMatcher_StringMatch{msg}
		return true, err
	case 4: // match_pattern.bool_match
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.MatchPattern = &ValueMatcher_BoolMatch{x != 0}
		return true, err
	case 5: // match_pattern.present_match
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.MatchPattern = &ValueMatcher_PresentMatch{x != 0}
		return true, err
	case 6: // match_pattern.list_match
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ListMatcher)
		err := b.DecodeMessage(msg)
		m.MatchPattern = &ValueMatcher_ListMatch{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ValueMatcher_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ValueMatcher)
	// match_pattern
	switch x := m.MatchPattern.(type) {
	case *ValueMatcher_NullMatch_:
		s := proto.Size(x.NullMatch)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ValueMatcher_DoubleMatch:
		s := proto.Size(x.DoubleMatch)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ValueMatcher_StringMatch:
		s := proto.Size(x.StringMatch)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ValueMatcher_BoolMatch:
		n += 1 // tag and wire
		n += 1
	case *ValueMatcher_PresentMatch:
		n += 1 // tag and wire
		n += 1
	case *ValueMatcher_ListMatch:
		s := proto.Size(x.ListMatch)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// NullMatch is an empty message to specify a null value.
type ValueMatcher_NullMatch struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValueMatcher_NullMatch) Reset()         { *m = ValueMatcher_NullMatch{} }
func (m *ValueMatcher_NullMatch) String() string { return proto.CompactTextString(m) }
func (*ValueMatcher_NullMatch) ProtoMessage()    {}
func (*ValueMatcher_NullMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_145b36501d266253, []int{0, 0}
}

func (m *ValueMatcher_NullMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValueMatcher_NullMatch.Unmarshal(m, b)
}
func (m *ValueMatcher_NullMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValueMatcher_NullMatch.Marshal(b, m, deterministic)
}
func (m *ValueMatcher_NullMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValueMatcher_NullMatch.Merge(m, src)
}
func (m *ValueMatcher_NullMatch) XXX_Size() int {
	return xxx_messageInfo_ValueMatcher_NullMatch.Size(m)
}
func (m *ValueMatcher_NullMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_ValueMatcher_NullMatch.DiscardUnknown(m)
}

var xxx_messageInfo_ValueMatcher_NullMatch proto.InternalMessageInfo

// Specifies the way to match a list value.
type ListMatcher struct {
	// Types that are valid to be assigned to MatchPattern:
	//	*ListMatcher_OneOf
	MatchPattern         isListMatcher_MatchPattern `protobuf_oneof:"match_pattern"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ListMatcher) Reset()         { *m = ListMatcher{} }
func (m *ListMatcher) String() string { return proto.CompactTextString(m) }
func (*ListMatcher) ProtoMessage()    {}
func (*ListMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_145b36501d266253, []int{1}
}

func (m *ListMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMatcher.Unmarshal(m, b)
}
func (m *ListMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMatcher.Marshal(b, m, deterministic)
}
func (m *ListMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMatcher.Merge(m, src)
}
func (m *ListMatcher) XXX_Size() int {
	return xxx_messageInfo_ListMatcher.Size(m)
}
func (m *ListMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_ListMatcher proto.InternalMessageInfo

type isListMatcher_MatchPattern interface {
	isListMatcher_MatchPattern()
}

type ListMatcher_OneOf struct {
	OneOf *ValueMatcher `protobuf:"bytes,1,opt,name=one_of,json=oneOf,proto3,oneof"`
}

func (*ListMatcher_OneOf) isListMatcher_MatchPattern() {}

func (m *ListMatcher) GetMatchPattern() isListMatcher_MatchPattern {
	if m != nil {
		return m.MatchPattern
	}
	return nil
}

func (m *ListMatcher) GetOneOf() *ValueMatcher {
	if x, ok := m.GetMatchPattern().(*ListMatcher_OneOf); ok {
		return x.OneOf
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ListMatcher) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ListMatcher_OneofMarshaler, _ListMatcher_OneofUnmarshaler, _ListMatcher_OneofSizer, []interface{}{
		(*ListMatcher_OneOf)(nil),
	}
}

func _ListMatcher_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ListMatcher)
	// match_pattern
	switch x := m.MatchPattern.(type) {
	case *ListMatcher_OneOf:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.OneOf); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ListMatcher.MatchPattern has unexpected type %T", x)
	}
	return nil
}

func _ListMatcher_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ListMatcher)
	switch tag {
	case 1: // match_pattern.one_of
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ValueMatcher)
		err := b.DecodeMessage(msg)
		m.MatchPattern = &ListMatcher_OneOf{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ListMatcher_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ListMatcher)
	// match_pattern
	switch x := m.MatchPattern.(type) {
	case *ListMatcher_OneOf:
		s := proto.Size(x.OneOf)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*ValueMatcher)(nil), "envoy.type.matcher.ValueMatcher")
	proto.RegisterType((*ValueMatcher_NullMatch)(nil), "envoy.type.matcher.ValueMatcher.NullMatch")
	proto.RegisterType((*ListMatcher)(nil), "envoy.type.matcher.ListMatcher")
}

func init() { proto.RegisterFile("envoy/type/matcher/value.proto", fileDescriptor_145b36501d266253) }

var fileDescriptor_145b36501d266253 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0x4b, 0xfb, 0x30,
	0x18, 0xc7, 0xd7, 0xdf, 0x7e, 0x9b, 0xf6, 0xe9, 0x76, 0xc9, 0x41, 0x65, 0x07, 0x37, 0x07, 0xc2,
	0xf0, 0x90, 0x82, 0x9e, 0xbc, 0xc9, 0x10, 0x19, 0xf8, 0x0f, 0x26, 0x78, 0xf0, 0x52, 0x5b, 0xfb,
	0x4c, 0x0b, 0x59, 0x52, 0xd2, 0xb4, 0xb0, 0xb7, 0xe2, 0x2b, 0xf1, 0xe8, 0xdb, 0xf1, 0x5d, 0x48,
	0x93, 0xd4, 0x16, 0xec, 0xf0, 0x96, 0x3c, 0xcf, 0xe7, 0xfb, 0xc9, 0x93, 0xa6, 0x70, 0x88, 0xbc,
	0x10, 0x1b, 0x5f, 0x6d, 0x52, 0xf4, 0xd7, 0xa1, 0x7a, 0x79, 0x43, 0xe9, 0x17, 0x21, 0xcb, 0x91,
	0xa6, 0x52, 0x28, 0x41, 0x88, 0xee, 0xd3, 0xb2, 0x4f, 0x6d, 0x7f, 0x34, 0x6e, 0xc9, 0xf0, 0x7c,
	0x1d, 0xa1, 0x34, 0xa1, 0x56, 0x20, 0x53, 0x32, 0xe1, 0xaf, 0x16, 0xd8, 0x2f, 0x42, 0x96, 0xc4,
	0xa1, 0x42, 0xbf, 0x5a, 0x98, 0xc6, 0xf4, 0xbd, 0x0b, 0x83, 0xc7, 0xf2, 0xf8, 0x5b, 0x13, 0x23,
	0xd7, 0x00, 0x3c, 0x67, 0x2c, 0xd0, 0x9a, 0x03, 0x67, 0xe2, 0xcc, 0xbc, 0xd3, 0x13, 0xfa, 0x7b,
	0x28, 0xda, 0x4c, 0xd1, 0xbb, 0x9c, 0x31, 0xbd, 0x5e, 0x74, 0x96, 0x2e, 0xaf, 0x36, 0xe4, 0x0a,
	0x06, 0xb1, 0xc8, 0x23, 0x86, 0x56, 0xf7, 0x4f, 0xeb, 0x8e, 0xda, 0x74, 0x97, 0x9a, 0xb3, 0xbe,
	0x45, 0x67, 0xe9, 0xc5, 0x75, 0xa1, 0xf4, 0x98, 0xeb, 0x58, 0x4f, 0x77, 0xbb, 0xe7, 0x41, 0x73,
	0x0d, 0x4f, 0x56, 0x17, 0xc8, 0x18, 0x20, 0x12, 0xa2, 0xba, 0xdc, 0xff, 0x89, 0x33, 0xdb, 0x2d,
	0x07, 0x2e, 0x6b, 0x06, 0x38, 0x86, 0x61, 0x2a, 0x31, 0x43, 0xae, 0x2c, 0xd3, 0xb3, 0xcc, 0xc0,
	0x96, 0x0d, 0x76, 0x01, 0xc0, 0x92, 0xac, 0x62, 0xfa, 0x7a, 0x9a, 0x71, 0xdb, 0x34, 0x37, 0x49,
	0xa6, 0xea, 0x59, 0x5c, 0x56, 0x6d, 0x47, 0x1e, 0xb8, 0x3f, 0xdf, 0x6c, 0xbe, 0x07, 0x43, 0x1d,
	0x08, 0xd2, 0x50, 0x29, 0x94, 0x9c, 0xf4, 0x3e, 0xbe, 0x3e, 0xbb, 0xce, 0xf4, 0x19, 0xbc, 0x86,
	0x80, 0x9c, 0x43, 0x5f, 0x70, 0x0c, 0xc4, 0xca, 0x3e, 0xcb, 0xe4, 0xaf, 0x67, 0x59, 0x74, 0x96,
	0x3d, 0xc1, 0xf1, 0x7e, 0xb5, 0xed, 0x84, 0xb9, 0xfb, 0xb4, 0x63, 0x83, 0x51, 0x5f, 0xff, 0x10,
	0x67, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x39, 0x37, 0xa7, 0x44, 0xa1, 0x02, 0x00, 0x00,
}
