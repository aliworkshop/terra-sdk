package stakingv1beta1

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	v1beta1 "github.com/aliworkshop/terra-sdk/api/cosmos/base/v1beta1"
	types "github.com/aliworkshop/terra-sdk/api/tendermint/types"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var _ protoreflect.List = (*_HistoricalInfo_2_list)(nil)

type _HistoricalInfo_2_list struct {
	list *[]*Validator
}

func (x *_HistoricalInfo_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_HistoricalInfo_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_HistoricalInfo_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Validator)
	(*x.list)[i] = concreteValue
}

func (x *_HistoricalInfo_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Validator)
	*x.list = append(*x.list, concreteValue)
}

func (x *_HistoricalInfo_2_list) AppendMutable() protoreflect.Value {
	v := new(Validator)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_HistoricalInfo_2_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_HistoricalInfo_2_list) NewElement() protoreflect.Value {
	v := new(Validator)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_HistoricalInfo_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_HistoricalInfo        protoreflect.MessageDescriptor
	fd_HistoricalInfo_header protoreflect.FieldDescriptor
	fd_HistoricalInfo_valset protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_staking_proto_init()
	md_HistoricalInfo = File_cosmos_staking_v1beta1_staking_proto.Messages().ByName("HistoricalInfo")
	fd_HistoricalInfo_header = md_HistoricalInfo.Fields().ByName("header")
	fd_HistoricalInfo_valset = md_HistoricalInfo.Fields().ByName("valset")
}

var _ protoreflect.Message = (*fastReflection_HistoricalInfo)(nil)

type fastReflection_HistoricalInfo HistoricalInfo

func (x *HistoricalInfo) ProtoReflect() protoreflect.Message {
	return (*fastReflection_HistoricalInfo)(x)
}

func (x *HistoricalInfo) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_HistoricalInfo_messageType fastReflection_HistoricalInfo_messageType
var _ protoreflect.MessageType = fastReflection_HistoricalInfo_messageType{}

type fastReflection_HistoricalInfo_messageType struct{}

func (x fastReflection_HistoricalInfo_messageType) Zero() protoreflect.Message {
	return (*fastReflection_HistoricalInfo)(nil)
}
func (x fastReflection_HistoricalInfo_messageType) New() protoreflect.Message {
	return new(fastReflection_HistoricalInfo)
}
func (x fastReflection_HistoricalInfo_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_HistoricalInfo
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_HistoricalInfo) Descriptor() protoreflect.MessageDescriptor {
	return md_HistoricalInfo
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_HistoricalInfo) Type() protoreflect.MessageType {
	return _fastReflection_HistoricalInfo_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_HistoricalInfo) New() protoreflect.Message {
	return new(fastReflection_HistoricalInfo)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_HistoricalInfo) Interface() protoreflect.ProtoMessage {
	return (*HistoricalInfo)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_HistoricalInfo) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Header != nil {
		value := protoreflect.ValueOfMessage(x.Header.ProtoReflect())
		if !f(fd_HistoricalInfo_header, value) {
			return
		}
	}
	if len(x.Valset) != 0 {
		value := protoreflect.ValueOfList(&_HistoricalInfo_2_list{list: &x.Valset})
		if !f(fd_HistoricalInfo_valset, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_HistoricalInfo) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.HistoricalInfo.header":
		return x.Header != nil
	case "cosmos.staking.v1beta1.HistoricalInfo.valset":
		return len(x.Valset) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.HistoricalInfo"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.HistoricalInfo does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_HistoricalInfo) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.HistoricalInfo.header":
		x.Header = nil
	case "cosmos.staking.v1beta1.HistoricalInfo.valset":
		x.Valset = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.HistoricalInfo"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.HistoricalInfo does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_HistoricalInfo) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.HistoricalInfo.header":
		value := x.Header
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.staking.v1beta1.HistoricalInfo.valset":
		if len(x.Valset) == 0 {
			return protoreflect.ValueOfList(&_HistoricalInfo_2_list{})
		}
		listValue := &_HistoricalInfo_2_list{list: &x.Valset}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.HistoricalInfo"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.HistoricalInfo does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_HistoricalInfo) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.HistoricalInfo.header":
		x.Header = value.Message().Interface().(*types.Header)
	case "cosmos.staking.v1beta1.HistoricalInfo.valset":
		lv := value.List()
		clv := lv.(*_HistoricalInfo_2_list)
		x.Valset = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.HistoricalInfo"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.HistoricalInfo does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_HistoricalInfo) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.HistoricalInfo.header":
		if x.Header == nil {
			x.Header = new(types.Header)
		}
		return protoreflect.ValueOfMessage(x.Header.ProtoReflect())
	case "cosmos.staking.v1beta1.HistoricalInfo.valset":
		if x.Valset == nil {
			x.Valset = []*Validator{}
		}
		value := &_HistoricalInfo_2_list{list: &x.Valset}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.HistoricalInfo"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.HistoricalInfo does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_HistoricalInfo) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.HistoricalInfo.header":
		m := new(types.Header)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.staking.v1beta1.HistoricalInfo.valset":
		list := []*Validator{}
		return protoreflect.ValueOfList(&_HistoricalInfo_2_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.HistoricalInfo"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.HistoricalInfo does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_HistoricalInfo) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.HistoricalInfo", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_HistoricalInfo) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_HistoricalInfo) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_HistoricalInfo) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_HistoricalInfo) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*HistoricalInfo)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.Header != nil {
			l = options.Size(x.Header)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Valset) > 0 {
			for _, e := range x.Valset {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*HistoricalInfo)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Valset) > 0 {
			for iNdEx := len(x.Valset) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Valset[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x12
			}
		}
		if x.Header != nil {
			encoded, err := options.Marshal(x.Header)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*HistoricalInfo)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: HistoricalInfo: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: HistoricalInfo: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Header == nil {
					x.Header = &types.Header{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Header); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Valset", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Valset = append(x.Valset, &Validator{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Valset[len(x.Valset)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_CommissionRates                 protoreflect.MessageDescriptor
	fd_CommissionRates_rate            protoreflect.FieldDescriptor
	fd_CommissionRates_max_rate        protoreflect.FieldDescriptor
	fd_CommissionRates_max_change_rate protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_staking_proto_init()
	md_CommissionRates = File_cosmos_staking_v1beta1_staking_proto.Messages().ByName("CommissionRates")
	fd_CommissionRates_rate = md_CommissionRates.Fields().ByName("rate")
	fd_CommissionRates_max_rate = md_CommissionRates.Fields().ByName("max_rate")
	fd_CommissionRates_max_change_rate = md_CommissionRates.Fields().ByName("max_change_rate")
}

var _ protoreflect.Message = (*fastReflection_CommissionRates)(nil)

type fastReflection_CommissionRates CommissionRates

func (x *CommissionRates) ProtoReflect() protoreflect.Message {
	return (*fastReflection_CommissionRates)(x)
}

func (x *CommissionRates) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_CommissionRates_messageType fastReflection_CommissionRates_messageType
var _ protoreflect.MessageType = fastReflection_CommissionRates_messageType{}

type fastReflection_CommissionRates_messageType struct{}

func (x fastReflection_CommissionRates_messageType) Zero() protoreflect.Message {
	return (*fastReflection_CommissionRates)(nil)
}
func (x fastReflection_CommissionRates_messageType) New() protoreflect.Message {
	return new(fastReflection_CommissionRates)
}
func (x fastReflection_CommissionRates_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_CommissionRates
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_CommissionRates) Descriptor() protoreflect.MessageDescriptor {
	return md_CommissionRates
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_CommissionRates) Type() protoreflect.MessageType {
	return _fastReflection_CommissionRates_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_CommissionRates) New() protoreflect.Message {
	return new(fastReflection_CommissionRates)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_CommissionRates) Interface() protoreflect.ProtoMessage {
	return (*CommissionRates)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_CommissionRates) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Rate != "" {
		value := protoreflect.ValueOfString(x.Rate)
		if !f(fd_CommissionRates_rate, value) {
			return
		}
	}
	if x.MaxRate != "" {
		value := protoreflect.ValueOfString(x.MaxRate)
		if !f(fd_CommissionRates_max_rate, value) {
			return
		}
	}
	if x.MaxChangeRate != "" {
		value := protoreflect.ValueOfString(x.MaxChangeRate)
		if !f(fd_CommissionRates_max_change_rate, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_CommissionRates) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.CommissionRates.rate":
		return x.Rate != ""
	case "cosmos.staking.v1beta1.CommissionRates.max_rate":
		return x.MaxRate != ""
	case "cosmos.staking.v1beta1.CommissionRates.max_change_rate":
		return x.MaxChangeRate != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.CommissionRates"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.CommissionRates does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_CommissionRates) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.CommissionRates.rate":
		x.Rate = ""
	case "cosmos.staking.v1beta1.CommissionRates.max_rate":
		x.MaxRate = ""
	case "cosmos.staking.v1beta1.CommissionRates.max_change_rate":
		x.MaxChangeRate = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.CommissionRates"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.CommissionRates does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_CommissionRates) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.CommissionRates.rate":
		value := x.Rate
		return protoreflect.ValueOfString(value)
	case "cosmos.staking.v1beta1.CommissionRates.max_rate":
		value := x.MaxRate
		return protoreflect.ValueOfString(value)
	case "cosmos.staking.v1beta1.CommissionRates.max_change_rate":
		value := x.MaxChangeRate
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.CommissionRates"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.CommissionRates does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_CommissionRates) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.CommissionRates.rate":
		x.Rate = value.Interface().(string)
	case "cosmos.staking.v1beta1.CommissionRates.max_rate":
		x.MaxRate = value.Interface().(string)
	case "cosmos.staking.v1beta1.CommissionRates.max_change_rate":
		x.MaxChangeRate = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.CommissionRates"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.CommissionRates does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_CommissionRates) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.CommissionRates.rate":
		panic(fmt.Errorf("field rate of message cosmos.staking.v1beta1.CommissionRates is not mutable"))
	case "cosmos.staking.v1beta1.CommissionRates.max_rate":
		panic(fmt.Errorf("field max_rate of message cosmos.staking.v1beta1.CommissionRates is not mutable"))
	case "cosmos.staking.v1beta1.CommissionRates.max_change_rate":
		panic(fmt.Errorf("field max_change_rate of message cosmos.staking.v1beta1.CommissionRates is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.CommissionRates"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.CommissionRates does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_CommissionRates) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.CommissionRates.rate":
		return protoreflect.ValueOfString("")
	case "cosmos.staking.v1beta1.CommissionRates.max_rate":
		return protoreflect.ValueOfString("")
	case "cosmos.staking.v1beta1.CommissionRates.max_change_rate":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.CommissionRates"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.CommissionRates does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_CommissionRates) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.CommissionRates", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_CommissionRates) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_CommissionRates) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_CommissionRates) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_CommissionRates) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*CommissionRates)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Rate)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.MaxRate)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.MaxChangeRate)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*CommissionRates)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.MaxChangeRate) > 0 {
			i -= len(x.MaxChangeRate)
			copy(dAtA[i:], x.MaxChangeRate)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.MaxChangeRate)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.MaxRate) > 0 {
			i -= len(x.MaxRate)
			copy(dAtA[i:], x.MaxRate)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.MaxRate)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Rate) > 0 {
			i -= len(x.Rate)
			copy(dAtA[i:], x.Rate)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Rate)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*CommissionRates)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: CommissionRates: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: CommissionRates: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Rate", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Rate = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MaxRate", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.MaxRate = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MaxChangeRate", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.MaxChangeRate = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_Commission                  protoreflect.MessageDescriptor
	fd_Commission_commission_rates protoreflect.FieldDescriptor
	fd_Commission_update_time      protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_staking_proto_init()
	md_Commission = File_cosmos_staking_v1beta1_staking_proto.Messages().ByName("Commission")
	fd_Commission_commission_rates = md_Commission.Fields().ByName("commission_rates")
	fd_Commission_update_time = md_Commission.Fields().ByName("update_time")
}

var _ protoreflect.Message = (*fastReflection_Commission)(nil)

type fastReflection_Commission Commission

func (x *Commission) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Commission)(x)
}

func (x *Commission) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Commission_messageType fastReflection_Commission_messageType
var _ protoreflect.MessageType = fastReflection_Commission_messageType{}

type fastReflection_Commission_messageType struct{}

func (x fastReflection_Commission_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Commission)(nil)
}
func (x fastReflection_Commission_messageType) New() protoreflect.Message {
	return new(fastReflection_Commission)
}
func (x fastReflection_Commission_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Commission
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Commission) Descriptor() protoreflect.MessageDescriptor {
	return md_Commission
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Commission) Type() protoreflect.MessageType {
	return _fastReflection_Commission_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Commission) New() protoreflect.Message {
	return new(fastReflection_Commission)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Commission) Interface() protoreflect.ProtoMessage {
	return (*Commission)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Commission) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.CommissionRates != nil {
		value := protoreflect.ValueOfMessage(x.CommissionRates.ProtoReflect())
		if !f(fd_Commission_commission_rates, value) {
			return
		}
	}
	if x.UpdateTime != nil {
		value := protoreflect.ValueOfMessage(x.UpdateTime.ProtoReflect())
		if !f(fd_Commission_update_time, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Commission) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.Commission.commission_rates":
		return x.CommissionRates != nil
	case "cosmos.staking.v1beta1.Commission.update_time":
		return x.UpdateTime != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Commission"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Commission does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Commission) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.Commission.commission_rates":
		x.CommissionRates = nil
	case "cosmos.staking.v1beta1.Commission.update_time":
		x.UpdateTime = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Commission"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Commission does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Commission) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.Commission.commission_rates":
		value := x.CommissionRates
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.staking.v1beta1.Commission.update_time":
		value := x.UpdateTime
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Commission"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Commission does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Commission) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.Commission.commission_rates":
		x.CommissionRates = value.Message().Interface().(*CommissionRates)
	case "cosmos.staking.v1beta1.Commission.update_time":
		x.UpdateTime = value.Message().Interface().(*timestamppb.Timestamp)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Commission"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Commission does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Commission) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.Commission.commission_rates":
		if x.CommissionRates == nil {
			x.CommissionRates = new(CommissionRates)
		}
		return protoreflect.ValueOfMessage(x.CommissionRates.ProtoReflect())
	case "cosmos.staking.v1beta1.Commission.update_time":
		if x.UpdateTime == nil {
			x.UpdateTime = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.UpdateTime.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Commission"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Commission does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Commission) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.Commission.commission_rates":
		m := new(CommissionRates)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.staking.v1beta1.Commission.update_time":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Commission"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Commission does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Commission) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.Commission", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Commission) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Commission) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Commission) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Commission) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Commission)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.CommissionRates != nil {
			l = options.Size(x.CommissionRates)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.UpdateTime != nil {
			l = options.Size(x.UpdateTime)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Commission)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.UpdateTime != nil {
			encoded, err := options.Marshal(x.UpdateTime)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if x.CommissionRates != nil {
			encoded, err := options.Marshal(x.CommissionRates)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Commission)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Commission: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Commission: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field CommissionRates", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.CommissionRates == nil {
					x.CommissionRates = &CommissionRates{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.CommissionRates); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field UpdateTime", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.UpdateTime == nil {
					x.UpdateTime = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.UpdateTime); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_Description                  protoreflect.MessageDescriptor
	fd_Description_moniker          protoreflect.FieldDescriptor
	fd_Description_identity         protoreflect.FieldDescriptor
	fd_Description_website          protoreflect.FieldDescriptor
	fd_Description_security_contact protoreflect.FieldDescriptor
	fd_Description_details          protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_staking_proto_init()
	md_Description = File_cosmos_staking_v1beta1_staking_proto.Messages().ByName("Description")
	fd_Description_moniker = md_Description.Fields().ByName("moniker")
	fd_Description_identity = md_Description.Fields().ByName("identity")
	fd_Description_website = md_Description.Fields().ByName("website")
	fd_Description_security_contact = md_Description.Fields().ByName("security_contact")
	fd_Description_details = md_Description.Fields().ByName("details")
}

var _ protoreflect.Message = (*fastReflection_Description)(nil)

type fastReflection_Description Description

func (x *Description) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Description)(x)
}

func (x *Description) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Description_messageType fastReflection_Description_messageType
var _ protoreflect.MessageType = fastReflection_Description_messageType{}

type fastReflection_Description_messageType struct{}

func (x fastReflection_Description_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Description)(nil)
}
func (x fastReflection_Description_messageType) New() protoreflect.Message {
	return new(fastReflection_Description)
}
func (x fastReflection_Description_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Description
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Description) Descriptor() protoreflect.MessageDescriptor {
	return md_Description
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Description) Type() protoreflect.MessageType {
	return _fastReflection_Description_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Description) New() protoreflect.Message {
	return new(fastReflection_Description)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Description) Interface() protoreflect.ProtoMessage {
	return (*Description)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Description) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Moniker != "" {
		value := protoreflect.ValueOfString(x.Moniker)
		if !f(fd_Description_moniker, value) {
			return
		}
	}
	if x.Identity != "" {
		value := protoreflect.ValueOfString(x.Identity)
		if !f(fd_Description_identity, value) {
			return
		}
	}
	if x.Website != "" {
		value := protoreflect.ValueOfString(x.Website)
		if !f(fd_Description_website, value) {
			return
		}
	}
	if x.SecurityContact != "" {
		value := protoreflect.ValueOfString(x.SecurityContact)
		if !f(fd_Description_security_contact, value) {
			return
		}
	}
	if x.Details != "" {
		value := protoreflect.ValueOfString(x.Details)
		if !f(fd_Description_details, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Description) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.Description.moniker":
		return x.Moniker != ""
	case "cosmos.staking.v1beta1.Description.identity":
		return x.Identity != ""
	case "cosmos.staking.v1beta1.Description.website":
		return x.Website != ""
	case "cosmos.staking.v1beta1.Description.security_contact":
		return x.SecurityContact != ""
	case "cosmos.staking.v1beta1.Description.details":
		return x.Details != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Description"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Description does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Description) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.Description.moniker":
		x.Moniker = ""
	case "cosmos.staking.v1beta1.Description.identity":
		x.Identity = ""
	case "cosmos.staking.v1beta1.Description.website":
		x.Website = ""
	case "cosmos.staking.v1beta1.Description.security_contact":
		x.SecurityContact = ""
	case "cosmos.staking.v1beta1.Description.details":
		x.Details = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Description"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Description does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Description) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.Description.moniker":
		value := x.Moniker
		return protoreflect.ValueOfString(value)
	case "cosmos.staking.v1beta1.Description.identity":
		value := x.Identity
		return protoreflect.ValueOfString(value)
	case "cosmos.staking.v1beta1.Description.website":
		value := x.Website
		return protoreflect.ValueOfString(value)
	case "cosmos.staking.v1beta1.Description.security_contact":
		value := x.SecurityContact
		return protoreflect.ValueOfString(value)
	case "cosmos.staking.v1beta1.Description.details":
		value := x.Details
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Description"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Description does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Description) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.Description.moniker":
		x.Moniker = value.Interface().(string)
	case "cosmos.staking.v1beta1.Description.identity":
		x.Identity = value.Interface().(string)
	case "cosmos.staking.v1beta1.Description.website":
		x.Website = value.Interface().(string)
	case "cosmos.staking.v1beta1.Description.security_contact":
		x.SecurityContact = value.Interface().(string)
	case "cosmos.staking.v1beta1.Description.details":
		x.Details = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Description"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Description does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Description) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.Description.moniker":
		panic(fmt.Errorf("field moniker of message cosmos.staking.v1beta1.Description is not mutable"))
	case "cosmos.staking.v1beta1.Description.identity":
		panic(fmt.Errorf("field identity of message cosmos.staking.v1beta1.Description is not mutable"))
	case "cosmos.staking.v1beta1.Description.website":
		panic(fmt.Errorf("field website of message cosmos.staking.v1beta1.Description is not mutable"))
	case "cosmos.staking.v1beta1.Description.security_contact":
		panic(fmt.Errorf("field security_contact of message cosmos.staking.v1beta1.Description is not mutable"))
	case "cosmos.staking.v1beta1.Description.details":
		panic(fmt.Errorf("field details of message cosmos.staking.v1beta1.Description is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Description"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Description does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Description) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.Description.moniker":
		return protoreflect.ValueOfString("")
	case "cosmos.staking.v1beta1.Description.identity":
		return protoreflect.ValueOfString("")
	case "cosmos.staking.v1beta1.Description.website":
		return protoreflect.ValueOfString("")
	case "cosmos.staking.v1beta1.Description.security_contact":
		return protoreflect.ValueOfString("")
	case "cosmos.staking.v1beta1.Description.details":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Description"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Description does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Description) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.Description", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Description) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Description) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Description) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Description) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Description)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Moniker)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Identity)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Website)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.SecurityContact)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Details)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Description)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Details) > 0 {
			i -= len(x.Details)
			copy(dAtA[i:], x.Details)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Details)))
			i--
			dAtA[i] = 0x2a
		}
		if len(x.SecurityContact) > 0 {
			i -= len(x.SecurityContact)
			copy(dAtA[i:], x.SecurityContact)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.SecurityContact)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Website) > 0 {
			i -= len(x.Website)
			copy(dAtA[i:], x.Website)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Website)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Identity) > 0 {
			i -= len(x.Identity)
			copy(dAtA[i:], x.Identity)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Identity)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Moniker) > 0 {
			i -= len(x.Moniker)
			copy(dAtA[i:], x.Moniker)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Moniker)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Description)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Description: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Description: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Moniker", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Moniker = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Identity", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Identity = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Website", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Website = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SecurityContact", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.SecurityContact = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Details", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Details = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_Validator                     protoreflect.MessageDescriptor
	fd_Validator_operator_address    protoreflect.FieldDescriptor
	fd_Validator_consensus_pubkey    protoreflect.FieldDescriptor
	fd_Validator_jailed              protoreflect.FieldDescriptor
	fd_Validator_status              protoreflect.FieldDescriptor
	fd_Validator_tokens              protoreflect.FieldDescriptor
	fd_Validator_delegator_shares    protoreflect.FieldDescriptor
	fd_Validator_description         protoreflect.FieldDescriptor
	fd_Validator_unbonding_height    protoreflect.FieldDescriptor
	fd_Validator_unbonding_time      protoreflect.FieldDescriptor
	fd_Validator_commission          protoreflect.FieldDescriptor
	fd_Validator_min_self_delegation protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_staking_proto_init()
	md_Validator = File_cosmos_staking_v1beta1_staking_proto.Messages().ByName("Validator")
	fd_Validator_operator_address = md_Validator.Fields().ByName("operator_address")
	fd_Validator_consensus_pubkey = md_Validator.Fields().ByName("consensus_pubkey")
	fd_Validator_jailed = md_Validator.Fields().ByName("jailed")
	fd_Validator_status = md_Validator.Fields().ByName("status")
	fd_Validator_tokens = md_Validator.Fields().ByName("tokens")
	fd_Validator_delegator_shares = md_Validator.Fields().ByName("delegator_shares")
	fd_Validator_description = md_Validator.Fields().ByName("description")
	fd_Validator_unbonding_height = md_Validator.Fields().ByName("unbonding_height")
	fd_Validator_unbonding_time = md_Validator.Fields().ByName("unbonding_time")
	fd_Validator_commission = md_Validator.Fields().ByName("commission")
	fd_Validator_min_self_delegation = md_Validator.Fields().ByName("min_self_delegation")
}

var _ protoreflect.Message = (*fastReflection_Validator)(nil)

type fastReflection_Validator Validator

func (x *Validator) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Validator)(x)
}

func (x *Validator) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Validator_messageType fastReflection_Validator_messageType
var _ protoreflect.MessageType = fastReflection_Validator_messageType{}

type fastReflection_Validator_messageType struct{}

func (x fastReflection_Validator_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Validator)(nil)
}
func (x fastReflection_Validator_messageType) New() protoreflect.Message {
	return new(fastReflection_Validator)
}
func (x fastReflection_Validator_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Validator
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Validator) Descriptor() protoreflect.MessageDescriptor {
	return md_Validator
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Validator) Type() protoreflect.MessageType {
	return _fastReflection_Validator_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Validator) New() protoreflect.Message {
	return new(fastReflection_Validator)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Validator) Interface() protoreflect.ProtoMessage {
	return (*Validator)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Validator) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.OperatorAddress != "" {
		value := protoreflect.ValueOfString(x.OperatorAddress)
		if !f(fd_Validator_operator_address, value) {
			return
		}
	}
	if x.ConsensusPubkey != nil {
		value := protoreflect.ValueOfMessage(x.ConsensusPubkey.ProtoReflect())
		if !f(fd_Validator_consensus_pubkey, value) {
			return
		}
	}
	if x.Jailed != false {
		value := protoreflect.ValueOfBool(x.Jailed)
		if !f(fd_Validator_jailed, value) {
			return
		}
	}
	if x.Status != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.Status))
		if !f(fd_Validator_status, value) {
			return
		}
	}
	if x.Tokens != "" {
		value := protoreflect.ValueOfString(x.Tokens)
		if !f(fd_Validator_tokens, value) {
			return
		}
	}
	if x.DelegatorShares != "" {
		value := protoreflect.ValueOfString(x.DelegatorShares)
		if !f(fd_Validator_delegator_shares, value) {
			return
		}
	}
	if x.Description != nil {
		value := protoreflect.ValueOfMessage(x.Description.ProtoReflect())
		if !f(fd_Validator_description, value) {
			return
		}
	}
	if x.UnbondingHeight != int64(0) {
		value := protoreflect.ValueOfInt64(x.UnbondingHeight)
		if !f(fd_Validator_unbonding_height, value) {
			return
		}
	}
	if x.UnbondingTime != nil {
		value := protoreflect.ValueOfMessage(x.UnbondingTime.ProtoReflect())
		if !f(fd_Validator_unbonding_time, value) {
			return
		}
	}
	if x.Commission != nil {
		value := protoreflect.ValueOfMessage(x.Commission.ProtoReflect())
		if !f(fd_Validator_commission, value) {
			return
		}
	}
	if x.MinSelfDelegation != "" {
		value := protoreflect.ValueOfString(x.MinSelfDelegation)
		if !f(fd_Validator_min_self_delegation, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Validator) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.Validator.operator_address":
		return x.OperatorAddress != ""
	case "cosmos.staking.v1beta1.Validator.consensus_pubkey":
		return x.ConsensusPubkey != nil
	case "cosmos.staking.v1beta1.Validator.jailed":
		return x.Jailed != false
	case "cosmos.staking.v1beta1.Validator.status":
		return x.Status != 0
	case "cosmos.staking.v1beta1.Validator.tokens":
		return x.Tokens != ""
	case "cosmos.staking.v1beta1.Validator.delegator_shares":
		return x.DelegatorShares != ""
	case "cosmos.staking.v1beta1.Validator.description":
		return x.Description != nil
	case "cosmos.staking.v1beta1.Validator.unbonding_height":
		return x.UnbondingHeight != int64(0)
	case "cosmos.staking.v1beta1.Validator.unbonding_time":
		return x.UnbondingTime != nil
	case "cosmos.staking.v1beta1.Validator.commission":
		return x.Commission != nil
	case "cosmos.staking.v1beta1.Validator.min_self_delegation":
		return x.MinSelfDelegation != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Validator"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Validator does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Validator) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.Validator.operator_address":
		x.OperatorAddress = ""
	case "cosmos.staking.v1beta1.Validator.consensus_pubkey":
		x.ConsensusPubkey = nil
	case "cosmos.staking.v1beta1.Validator.jailed":
		x.Jailed = false
	case "cosmos.staking.v1beta1.Validator.status":
		x.Status = 0
	case "cosmos.staking.v1beta1.Validator.tokens":
		x.Tokens = ""
	case "cosmos.staking.v1beta1.Validator.delegator_shares":
		x.DelegatorShares = ""
	case "cosmos.staking.v1beta1.Validator.description":
		x.Description = nil
	case "cosmos.staking.v1beta1.Validator.unbonding_height":
		x.UnbondingHeight = int64(0)
	case "cosmos.staking.v1beta1.Validator.unbonding_time":
		x.UnbondingTime = nil
	case "cosmos.staking.v1beta1.Validator.commission":
		x.Commission = nil
	case "cosmos.staking.v1beta1.Validator.min_self_delegation":
		x.MinSelfDelegation = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Validator"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Validator does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Validator) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.Validator.operator_address":
		value := x.OperatorAddress
		return protoreflect.ValueOfString(value)
	case "cosmos.staking.v1beta1.Validator.consensus_pubkey":
		value := x.ConsensusPubkey
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.staking.v1beta1.Validator.jailed":
		value := x.Jailed
		return protoreflect.ValueOfBool(value)
	case "cosmos.staking.v1beta1.Validator.status":
		value := x.Status
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	case "cosmos.staking.v1beta1.Validator.tokens":
		value := x.Tokens
		return protoreflect.ValueOfString(value)
	case "cosmos.staking.v1beta1.Validator.delegator_shares":
		value := x.DelegatorShares
		return protoreflect.ValueOfString(value)
	case "cosmos.staking.v1beta1.Validator.description":
		value := x.Description
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.staking.v1beta1.Validator.unbonding_height":
		value := x.UnbondingHeight
		return protoreflect.ValueOfInt64(value)
	case "cosmos.staking.v1beta1.Validator.unbonding_time":
		value := x.UnbondingTime
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.staking.v1beta1.Validator.commission":
		value := x.Commission
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.staking.v1beta1.Validator.min_self_delegation":
		value := x.MinSelfDelegation
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Validator"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Validator does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Validator) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.Validator.operator_address":
		x.OperatorAddress = value.Interface().(string)
	case "cosmos.staking.v1beta1.Validator.consensus_pubkey":
		x.ConsensusPubkey = value.Message().Interface().(*anypb.Any)
	case "cosmos.staking.v1beta1.Validator.jailed":
		x.Jailed = value.Bool()
	case "cosmos.staking.v1beta1.Validator.status":
		x.Status = (BondStatus)(value.Enum())
	case "cosmos.staking.v1beta1.Validator.tokens":
		x.Tokens = value.Interface().(string)
	case "cosmos.staking.v1beta1.Validator.delegator_shares":
		x.DelegatorShares = value.Interface().(string)
	case "cosmos.staking.v1beta1.Validator.description":
		x.Description = value.Message().Interface().(*Description)
	case "cosmos.staking.v1beta1.Validator.unbonding_height":
		x.UnbondingHeight = value.Int()
	case "cosmos.staking.v1beta1.Validator.unbonding_time":
		x.UnbondingTime = value.Message().Interface().(*timestamppb.Timestamp)
	case "cosmos.staking.v1beta1.Validator.commission":
		x.Commission = value.Message().Interface().(*Commission)
	case "cosmos.staking.v1beta1.Validator.min_self_delegation":
		x.MinSelfDelegation = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Validator"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Validator does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Validator) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.Validator.consensus_pubkey":
		if x.ConsensusPubkey == nil {
			x.ConsensusPubkey = new(anypb.Any)
		}
		return protoreflect.ValueOfMessage(x.ConsensusPubkey.ProtoReflect())
	case "cosmos.staking.v1beta1.Validator.description":
		if x.Description == nil {
			x.Description = new(Description)
		}
		return protoreflect.ValueOfMessage(x.Description.ProtoReflect())
	case "cosmos.staking.v1beta1.Validator.unbonding_time":
		if x.UnbondingTime == nil {
			x.UnbondingTime = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.UnbondingTime.ProtoReflect())
	case "cosmos.staking.v1beta1.Validator.commission":
		if x.Commission == nil {
			x.Commission = new(Commission)
		}
		return protoreflect.ValueOfMessage(x.Commission.ProtoReflect())
	case "cosmos.staking.v1beta1.Validator.operator_address":
		panic(fmt.Errorf("field operator_address of message cosmos.staking.v1beta1.Validator is not mutable"))
	case "cosmos.staking.v1beta1.Validator.jailed":
		panic(fmt.Errorf("field jailed of message cosmos.staking.v1beta1.Validator is not mutable"))
	case "cosmos.staking.v1beta1.Validator.status":
		panic(fmt.Errorf("field status of message cosmos.staking.v1beta1.Validator is not mutable"))
	case "cosmos.staking.v1beta1.Validator.tokens":
		panic(fmt.Errorf("field tokens of message cosmos.staking.v1beta1.Validator is not mutable"))
	case "cosmos.staking.v1beta1.Validator.delegator_shares":
		panic(fmt.Errorf("field delegator_shares of message cosmos.staking.v1beta1.Validator is not mutable"))
	case "cosmos.staking.v1beta1.Validator.unbonding_height":
		panic(fmt.Errorf("field unbonding_height of message cosmos.staking.v1beta1.Validator is not mutable"))
	case "cosmos.staking.v1beta1.Validator.min_self_delegation":
		panic(fmt.Errorf("field min_self_delegation of message cosmos.staking.v1beta1.Validator is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Validator"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Validator does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Validator) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.Validator.operator_address":
		return protoreflect.ValueOfString("")
	case "cosmos.staking.v1beta1.Validator.consensus_pubkey":
		m := new(anypb.Any)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.staking.v1beta1.Validator.jailed":
		return protoreflect.ValueOfBool(false)
	case "cosmos.staking.v1beta1.Validator.status":
		return protoreflect.ValueOfEnum(0)
	case "cosmos.staking.v1beta1.Validator.tokens":
		return protoreflect.ValueOfString("")
	case "cosmos.staking.v1beta1.Validator.delegator_shares":
		return protoreflect.ValueOfString("")
	case "cosmos.staking.v1beta1.Validator.description":
		m := new(Description)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.staking.v1beta1.Validator.unbonding_height":
		return protoreflect.ValueOfInt64(int64(0))
	case "cosmos.staking.v1beta1.Validator.unbonding_time":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.staking.v1beta1.Validator.commission":
		m := new(Commission)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.staking.v1beta1.Validator.min_self_delegation":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Validator"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Validator does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Validator) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.Validator", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Validator) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Validator) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Validator) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Validator) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Validator)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.OperatorAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.ConsensusPubkey != nil {
			l = options.Size(x.ConsensusPubkey)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Jailed {
			n += 2
		}
		if x.Status != 0 {
			n += 1 + runtime.Sov(uint64(x.Status))
		}
		l = len(x.Tokens)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.DelegatorShares)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Description != nil {
			l = options.Size(x.Description)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.UnbondingHeight != 0 {
			n += 1 + runtime.Sov(uint64(x.UnbondingHeight))
		}
		if x.UnbondingTime != nil {
			l = options.Size(x.UnbondingTime)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Commission != nil {
			l = options.Size(x.Commission)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.MinSelfDelegation)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Validator)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.MinSelfDelegation) > 0 {
			i -= len(x.MinSelfDelegation)
			copy(dAtA[i:], x.MinSelfDelegation)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.MinSelfDelegation)))
			i--
			dAtA[i] = 0x5a
		}
		if x.Commission != nil {
			encoded, err := options.Marshal(x.Commission)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x52
		}
		if x.UnbondingTime != nil {
			encoded, err := options.Marshal(x.UnbondingTime)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x4a
		}
		if x.UnbondingHeight != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.UnbondingHeight))
			i--
			dAtA[i] = 0x40
		}
		if x.Description != nil {
			encoded, err := options.Marshal(x.Description)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x3a
		}
		if len(x.DelegatorShares) > 0 {
			i -= len(x.DelegatorShares)
			copy(dAtA[i:], x.DelegatorShares)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.DelegatorShares)))
			i--
			dAtA[i] = 0x32
		}
		if len(x.Tokens) > 0 {
			i -= len(x.Tokens)
			copy(dAtA[i:], x.Tokens)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Tokens)))
			i--
			dAtA[i] = 0x2a
		}
		if x.Status != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Status))
			i--
			dAtA[i] = 0x20
		}
		if x.Jailed {
			i--
			if x.Jailed {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x18
		}
		if x.ConsensusPubkey != nil {
			encoded, err := options.Marshal(x.ConsensusPubkey)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.OperatorAddress) > 0 {
			i -= len(x.OperatorAddress)
			copy(dAtA[i:], x.OperatorAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.OperatorAddress)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Validator)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Validator: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Validator: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field OperatorAddress", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.OperatorAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ConsensusPubkey", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.ConsensusPubkey == nil {
					x.ConsensusPubkey = &anypb.Any{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.ConsensusPubkey); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Jailed", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.Jailed = bool(v != 0)
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
				}
				x.Status = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Status |= BondStatus(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Tokens", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Tokens = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 6:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DelegatorShares", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.DelegatorShares = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 7:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Description == nil {
					x.Description = &Description{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Description); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 8:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field UnbondingHeight", wireType)
				}
				x.UnbondingHeight = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.UnbondingHeight |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 9:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field UnbondingTime", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.UnbondingTime == nil {
					x.UnbondingTime = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.UnbondingTime); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 10:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Commission", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Commission == nil {
					x.Commission = &Commission{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Commission); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 11:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MinSelfDelegation", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.MinSelfDelegation = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_ValAddresses_1_list)(nil)

type _ValAddresses_1_list struct {
	list *[]string
}

func (x *_ValAddresses_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_ValAddresses_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfString((*x.list)[i])
}

func (x *_ValAddresses_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_ValAddresses_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_ValAddresses_1_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message ValAddresses at list field Addresses as it is not of Message kind"))
}

func (x *_ValAddresses_1_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_ValAddresses_1_list) NewElement() protoreflect.Value {
	v := ""
	return protoreflect.ValueOfString(v)
}

func (x *_ValAddresses_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_ValAddresses           protoreflect.MessageDescriptor
	fd_ValAddresses_addresses protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_staking_proto_init()
	md_ValAddresses = File_cosmos_staking_v1beta1_staking_proto.Messages().ByName("ValAddresses")
	fd_ValAddresses_addresses = md_ValAddresses.Fields().ByName("addresses")
}

var _ protoreflect.Message = (*fastReflection_ValAddresses)(nil)

type fastReflection_ValAddresses ValAddresses

func (x *ValAddresses) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ValAddresses)(x)
}

func (x *ValAddresses) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ValAddresses_messageType fastReflection_ValAddresses_messageType
var _ protoreflect.MessageType = fastReflection_ValAddresses_messageType{}

type fastReflection_ValAddresses_messageType struct{}

func (x fastReflection_ValAddresses_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ValAddresses)(nil)
}
func (x fastReflection_ValAddresses_messageType) New() protoreflect.Message {
	return new(fastReflection_ValAddresses)
}
func (x fastReflection_ValAddresses_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ValAddresses
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ValAddresses) Descriptor() protoreflect.MessageDescriptor {
	return md_ValAddresses
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ValAddresses) Type() protoreflect.MessageType {
	return _fastReflection_ValAddresses_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ValAddresses) New() protoreflect.Message {
	return new(fastReflection_ValAddresses)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ValAddresses) Interface() protoreflect.ProtoMessage {
	return (*ValAddresses)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ValAddresses) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Addresses) != 0 {
		value := protoreflect.ValueOfList(&_ValAddresses_1_list{list: &x.Addresses})
		if !f(fd_ValAddresses_addresses, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_ValAddresses) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.ValAddresses.addresses":
		return len(x.Addresses) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.ValAddresses"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.ValAddresses does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValAddresses) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.ValAddresses.addresses":
		x.Addresses = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.ValAddresses"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.ValAddresses does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ValAddresses) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.ValAddresses.addresses":
		if len(x.Addresses) == 0 {
			return protoreflect.ValueOfList(&_ValAddresses_1_list{})
		}
		listValue := &_ValAddresses_1_list{list: &x.Addresses}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.ValAddresses"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.ValAddresses does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValAddresses) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.ValAddresses.addresses":
		lv := value.List()
		clv := lv.(*_ValAddresses_1_list)
		x.Addresses = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.ValAddresses"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.ValAddresses does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValAddresses) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.ValAddresses.addresses":
		if x.Addresses == nil {
			x.Addresses = []string{}
		}
		value := &_ValAddresses_1_list{list: &x.Addresses}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.ValAddresses"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.ValAddresses does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ValAddresses) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.ValAddresses.addresses":
		list := []string{}
		return protoreflect.ValueOfList(&_ValAddresses_1_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.ValAddresses"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.ValAddresses does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ValAddresses) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.ValAddresses", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ValAddresses) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValAddresses) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_ValAddresses) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ValAddresses) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ValAddresses)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if len(x.Addresses) > 0 {
			for _, s := range x.Addresses {
				l = len(s)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*ValAddresses)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Addresses) > 0 {
			for iNdEx := len(x.Addresses) - 1; iNdEx >= 0; iNdEx-- {
				i -= len(x.Addresses[iNdEx])
				copy(dAtA[i:], x.Addresses[iNdEx])
				i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Addresses[iNdEx])))
				i--
				dAtA[i] = 0xa
			}
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*ValAddresses)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ValAddresses: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ValAddresses: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Addresses", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Addresses = append(x.Addresses, string(dAtA[iNdEx:postIndex]))
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_DVPair                   protoreflect.MessageDescriptor
	fd_DVPair_delegator_address protoreflect.FieldDescriptor
	fd_DVPair_validator_address protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_staking_proto_init()
	md_DVPair = File_cosmos_staking_v1beta1_staking_proto.Messages().ByName("DVPair")
	fd_DVPair_delegator_address = md_DVPair.Fields().ByName("delegator_address")
	fd_DVPair_validator_address = md_DVPair.Fields().ByName("validator_address")
}

var _ protoreflect.Message = (*fastReflection_DVPair)(nil)

type fastReflection_DVPair DVPair

func (x *DVPair) ProtoReflect() protoreflect.Message {
	return (*fastReflection_DVPair)(x)
}

func (x *DVPair) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_DVPair_messageType fastReflection_DVPair_messageType
var _ protoreflect.MessageType = fastReflection_DVPair_messageType{}

type fastReflection_DVPair_messageType struct{}

func (x fastReflection_DVPair_messageType) Zero() protoreflect.Message {
	return (*fastReflection_DVPair)(nil)
}
func (x fastReflection_DVPair_messageType) New() protoreflect.Message {
	return new(fastReflection_DVPair)
}
func (x fastReflection_DVPair_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_DVPair
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_DVPair) Descriptor() protoreflect.MessageDescriptor {
	return md_DVPair
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_DVPair) Type() protoreflect.MessageType {
	return _fastReflection_DVPair_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_DVPair) New() protoreflect.Message {
	return new(fastReflection_DVPair)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_DVPair) Interface() protoreflect.ProtoMessage {
	return (*DVPair)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_DVPair) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.DelegatorAddress != "" {
		value := protoreflect.ValueOfString(x.DelegatorAddress)
		if !f(fd_DVPair_delegator_address, value) {
			return
		}
	}
	if x.ValidatorAddress != "" {
		value := protoreflect.ValueOfString(x.ValidatorAddress)
		if !f(fd_DVPair_validator_address, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_DVPair) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.DVPair.delegator_address":
		return x.DelegatorAddress != ""
	case "cosmos.staking.v1beta1.DVPair.validator_address":
		return x.ValidatorAddress != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.DVPair"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.DVPair does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DVPair) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.DVPair.delegator_address":
		x.DelegatorAddress = ""
	case "cosmos.staking.v1beta1.DVPair.validator_address":
		x.ValidatorAddress = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.DVPair"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.DVPair does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_DVPair) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.DVPair.delegator_address":
		value := x.DelegatorAddress
		return protoreflect.ValueOfString(value)
	case "cosmos.staking.v1beta1.DVPair.validator_address":
		value := x.ValidatorAddress
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.DVPair"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.DVPair does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DVPair) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.DVPair.delegator_address":
		x.DelegatorAddress = value.Interface().(string)
	case "cosmos.staking.v1beta1.DVPair.validator_address":
		x.ValidatorAddress = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.DVPair"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.DVPair does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DVPair) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.DVPair.delegator_address":
		panic(fmt.Errorf("field delegator_address of message cosmos.staking.v1beta1.DVPair is not mutable"))
	case "cosmos.staking.v1beta1.DVPair.validator_address":
		panic(fmt.Errorf("field validator_address of message cosmos.staking.v1beta1.DVPair is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.DVPair"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.DVPair does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_DVPair) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.DVPair.delegator_address":
		return protoreflect.ValueOfString("")
	case "cosmos.staking.v1beta1.DVPair.validator_address":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.DVPair"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.DVPair does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_DVPair) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.DVPair", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_DVPair) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DVPair) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_DVPair) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_DVPair) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*DVPair)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.DelegatorAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ValidatorAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*DVPair)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.ValidatorAddress) > 0 {
			i -= len(x.ValidatorAddress)
			copy(dAtA[i:], x.ValidatorAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ValidatorAddress)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.DelegatorAddress) > 0 {
			i -= len(x.DelegatorAddress)
			copy(dAtA[i:], x.DelegatorAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.DelegatorAddress)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*DVPair)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: DVPair: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: DVPair: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.DelegatorAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ValidatorAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_DVPairs_1_list)(nil)

type _DVPairs_1_list struct {
	list *[]*DVPair
}

func (x *_DVPairs_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_DVPairs_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_DVPairs_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*DVPair)
	(*x.list)[i] = concreteValue
}

func (x *_DVPairs_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*DVPair)
	*x.list = append(*x.list, concreteValue)
}

func (x *_DVPairs_1_list) AppendMutable() protoreflect.Value {
	v := new(DVPair)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_DVPairs_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_DVPairs_1_list) NewElement() protoreflect.Value {
	v := new(DVPair)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_DVPairs_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_DVPairs       protoreflect.MessageDescriptor
	fd_DVPairs_pairs protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_staking_proto_init()
	md_DVPairs = File_cosmos_staking_v1beta1_staking_proto.Messages().ByName("DVPairs")
	fd_DVPairs_pairs = md_DVPairs.Fields().ByName("pairs")
}

var _ protoreflect.Message = (*fastReflection_DVPairs)(nil)

type fastReflection_DVPairs DVPairs

func (x *DVPairs) ProtoReflect() protoreflect.Message {
	return (*fastReflection_DVPairs)(x)
}

func (x *DVPairs) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_DVPairs_messageType fastReflection_DVPairs_messageType
var _ protoreflect.MessageType = fastReflection_DVPairs_messageType{}

type fastReflection_DVPairs_messageType struct{}

func (x fastReflection_DVPairs_messageType) Zero() protoreflect.Message {
	return (*fastReflection_DVPairs)(nil)
}
func (x fastReflection_DVPairs_messageType) New() protoreflect.Message {
	return new(fastReflection_DVPairs)
}
func (x fastReflection_DVPairs_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_DVPairs
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_DVPairs) Descriptor() protoreflect.MessageDescriptor {
	return md_DVPairs
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_DVPairs) Type() protoreflect.MessageType {
	return _fastReflection_DVPairs_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_DVPairs) New() protoreflect.Message {
	return new(fastReflection_DVPairs)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_DVPairs) Interface() protoreflect.ProtoMessage {
	return (*DVPairs)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_DVPairs) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Pairs) != 0 {
		value := protoreflect.ValueOfList(&_DVPairs_1_list{list: &x.Pairs})
		if !f(fd_DVPairs_pairs, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_DVPairs) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.DVPairs.pairs":
		return len(x.Pairs) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.DVPairs"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.DVPairs does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DVPairs) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.DVPairs.pairs":
		x.Pairs = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.DVPairs"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.DVPairs does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_DVPairs) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.DVPairs.pairs":
		if len(x.Pairs) == 0 {
			return protoreflect.ValueOfList(&_DVPairs_1_list{})
		}
		listValue := &_DVPairs_1_list{list: &x.Pairs}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.DVPairs"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.DVPairs does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DVPairs) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.DVPairs.pairs":
		lv := value.List()
		clv := lv.(*_DVPairs_1_list)
		x.Pairs = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.DVPairs"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.DVPairs does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DVPairs) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.DVPairs.pairs":
		if x.Pairs == nil {
			x.Pairs = []*DVPair{}
		}
		value := &_DVPairs_1_list{list: &x.Pairs}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.DVPairs"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.DVPairs does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_DVPairs) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.DVPairs.pairs":
		list := []*DVPair{}
		return protoreflect.ValueOfList(&_DVPairs_1_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.DVPairs"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.DVPairs does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_DVPairs) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.DVPairs", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_DVPairs) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DVPairs) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_DVPairs) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_DVPairs) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*DVPairs)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if len(x.Pairs) > 0 {
			for _, e := range x.Pairs {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*DVPairs)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Pairs) > 0 {
			for iNdEx := len(x.Pairs) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Pairs[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0xa
			}
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*DVPairs)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: DVPairs: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: DVPairs: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Pairs", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Pairs = append(x.Pairs, &DVPair{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Pairs[len(x.Pairs)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_DVVTriplet                       protoreflect.MessageDescriptor
	fd_DVVTriplet_delegator_address     protoreflect.FieldDescriptor
	fd_DVVTriplet_validator_src_address protoreflect.FieldDescriptor
	fd_DVVTriplet_validator_dst_address protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_staking_proto_init()
	md_DVVTriplet = File_cosmos_staking_v1beta1_staking_proto.Messages().ByName("DVVTriplet")
	fd_DVVTriplet_delegator_address = md_DVVTriplet.Fields().ByName("delegator_address")
	fd_DVVTriplet_validator_src_address = md_DVVTriplet.Fields().ByName("validator_src_address")
	fd_DVVTriplet_validator_dst_address = md_DVVTriplet.Fields().ByName("validator_dst_address")
}

var _ protoreflect.Message = (*fastReflection_DVVTriplet)(nil)

type fastReflection_DVVTriplet DVVTriplet

func (x *DVVTriplet) ProtoReflect() protoreflect.Message {
	return (*fastReflection_DVVTriplet)(x)
}

func (x *DVVTriplet) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_DVVTriplet_messageType fastReflection_DVVTriplet_messageType
var _ protoreflect.MessageType = fastReflection_DVVTriplet_messageType{}

type fastReflection_DVVTriplet_messageType struct{}

func (x fastReflection_DVVTriplet_messageType) Zero() protoreflect.Message {
	return (*fastReflection_DVVTriplet)(nil)
}
func (x fastReflection_DVVTriplet_messageType) New() protoreflect.Message {
	return new(fastReflection_DVVTriplet)
}
func (x fastReflection_DVVTriplet_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_DVVTriplet
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_DVVTriplet) Descriptor() protoreflect.MessageDescriptor {
	return md_DVVTriplet
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_DVVTriplet) Type() protoreflect.MessageType {
	return _fastReflection_DVVTriplet_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_DVVTriplet) New() protoreflect.Message {
	return new(fastReflection_DVVTriplet)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_DVVTriplet) Interface() protoreflect.ProtoMessage {
	return (*DVVTriplet)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_DVVTriplet) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.DelegatorAddress != "" {
		value := protoreflect.ValueOfString(x.DelegatorAddress)
		if !f(fd_DVVTriplet_delegator_address, value) {
			return
		}
	}
	if x.ValidatorSrcAddress != "" {
		value := protoreflect.ValueOfString(x.ValidatorSrcAddress)
		if !f(fd_DVVTriplet_validator_src_address, value) {
			return
		}
	}
	if x.ValidatorDstAddress != "" {
		value := protoreflect.ValueOfString(x.ValidatorDstAddress)
		if !f(fd_DVVTriplet_validator_dst_address, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_DVVTriplet) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.DVVTriplet.delegator_address":
		return x.DelegatorAddress != ""
	case "cosmos.staking.v1beta1.DVVTriplet.validator_src_address":
		return x.ValidatorSrcAddress != ""
	case "cosmos.staking.v1beta1.DVVTriplet.validator_dst_address":
		return x.ValidatorDstAddress != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.DVVTriplet"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.DVVTriplet does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DVVTriplet) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.DVVTriplet.delegator_address":
		x.DelegatorAddress = ""
	case "cosmos.staking.v1beta1.DVVTriplet.validator_src_address":
		x.ValidatorSrcAddress = ""
	case "cosmos.staking.v1beta1.DVVTriplet.validator_dst_address":
		x.ValidatorDstAddress = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.DVVTriplet"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.DVVTriplet does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_DVVTriplet) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.DVVTriplet.delegator_address":
		value := x.DelegatorAddress
		return protoreflect.ValueOfString(value)
	case "cosmos.staking.v1beta1.DVVTriplet.validator_src_address":
		value := x.ValidatorSrcAddress
		return protoreflect.ValueOfString(value)
	case "cosmos.staking.v1beta1.DVVTriplet.validator_dst_address":
		value := x.ValidatorDstAddress
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.DVVTriplet"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.DVVTriplet does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DVVTriplet) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.DVVTriplet.delegator_address":
		x.DelegatorAddress = value.Interface().(string)
	case "cosmos.staking.v1beta1.DVVTriplet.validator_src_address":
		x.ValidatorSrcAddress = value.Interface().(string)
	case "cosmos.staking.v1beta1.DVVTriplet.validator_dst_address":
		x.ValidatorDstAddress = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.DVVTriplet"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.DVVTriplet does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DVVTriplet) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.DVVTriplet.delegator_address":
		panic(fmt.Errorf("field delegator_address of message cosmos.staking.v1beta1.DVVTriplet is not mutable"))
	case "cosmos.staking.v1beta1.DVVTriplet.validator_src_address":
		panic(fmt.Errorf("field validator_src_address of message cosmos.staking.v1beta1.DVVTriplet is not mutable"))
	case "cosmos.staking.v1beta1.DVVTriplet.validator_dst_address":
		panic(fmt.Errorf("field validator_dst_address of message cosmos.staking.v1beta1.DVVTriplet is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.DVVTriplet"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.DVVTriplet does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_DVVTriplet) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.DVVTriplet.delegator_address":
		return protoreflect.ValueOfString("")
	case "cosmos.staking.v1beta1.DVVTriplet.validator_src_address":
		return protoreflect.ValueOfString("")
	case "cosmos.staking.v1beta1.DVVTriplet.validator_dst_address":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.DVVTriplet"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.DVVTriplet does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_DVVTriplet) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.DVVTriplet", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_DVVTriplet) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DVVTriplet) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_DVVTriplet) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_DVVTriplet) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*DVVTriplet)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.DelegatorAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ValidatorSrcAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ValidatorDstAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*DVVTriplet)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.ValidatorDstAddress) > 0 {
			i -= len(x.ValidatorDstAddress)
			copy(dAtA[i:], x.ValidatorDstAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ValidatorDstAddress)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.ValidatorSrcAddress) > 0 {
			i -= len(x.ValidatorSrcAddress)
			copy(dAtA[i:], x.ValidatorSrcAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ValidatorSrcAddress)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.DelegatorAddress) > 0 {
			i -= len(x.DelegatorAddress)
			copy(dAtA[i:], x.DelegatorAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.DelegatorAddress)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*DVVTriplet)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: DVVTriplet: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: DVVTriplet: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.DelegatorAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ValidatorSrcAddress", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ValidatorSrcAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ValidatorDstAddress", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ValidatorDstAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_DVVTriplets_1_list)(nil)

type _DVVTriplets_1_list struct {
	list *[]*DVVTriplet
}

func (x *_DVVTriplets_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_DVVTriplets_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_DVVTriplets_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*DVVTriplet)
	(*x.list)[i] = concreteValue
}

func (x *_DVVTriplets_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*DVVTriplet)
	*x.list = append(*x.list, concreteValue)
}

func (x *_DVVTriplets_1_list) AppendMutable() protoreflect.Value {
	v := new(DVVTriplet)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_DVVTriplets_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_DVVTriplets_1_list) NewElement() protoreflect.Value {
	v := new(DVVTriplet)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_DVVTriplets_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_DVVTriplets          protoreflect.MessageDescriptor
	fd_DVVTriplets_triplets protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_staking_proto_init()
	md_DVVTriplets = File_cosmos_staking_v1beta1_staking_proto.Messages().ByName("DVVTriplets")
	fd_DVVTriplets_triplets = md_DVVTriplets.Fields().ByName("triplets")
}

var _ protoreflect.Message = (*fastReflection_DVVTriplets)(nil)

type fastReflection_DVVTriplets DVVTriplets

func (x *DVVTriplets) ProtoReflect() protoreflect.Message {
	return (*fastReflection_DVVTriplets)(x)
}

func (x *DVVTriplets) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_DVVTriplets_messageType fastReflection_DVVTriplets_messageType
var _ protoreflect.MessageType = fastReflection_DVVTriplets_messageType{}

type fastReflection_DVVTriplets_messageType struct{}

func (x fastReflection_DVVTriplets_messageType) Zero() protoreflect.Message {
	return (*fastReflection_DVVTriplets)(nil)
}
func (x fastReflection_DVVTriplets_messageType) New() protoreflect.Message {
	return new(fastReflection_DVVTriplets)
}
func (x fastReflection_DVVTriplets_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_DVVTriplets
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_DVVTriplets) Descriptor() protoreflect.MessageDescriptor {
	return md_DVVTriplets
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_DVVTriplets) Type() protoreflect.MessageType {
	return _fastReflection_DVVTriplets_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_DVVTriplets) New() protoreflect.Message {
	return new(fastReflection_DVVTriplets)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_DVVTriplets) Interface() protoreflect.ProtoMessage {
	return (*DVVTriplets)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_DVVTriplets) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Triplets) != 0 {
		value := protoreflect.ValueOfList(&_DVVTriplets_1_list{list: &x.Triplets})
		if !f(fd_DVVTriplets_triplets, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_DVVTriplets) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.DVVTriplets.triplets":
		return len(x.Triplets) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.DVVTriplets"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.DVVTriplets does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DVVTriplets) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.DVVTriplets.triplets":
		x.Triplets = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.DVVTriplets"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.DVVTriplets does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_DVVTriplets) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.DVVTriplets.triplets":
		if len(x.Triplets) == 0 {
			return protoreflect.ValueOfList(&_DVVTriplets_1_list{})
		}
		listValue := &_DVVTriplets_1_list{list: &x.Triplets}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.DVVTriplets"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.DVVTriplets does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DVVTriplets) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.DVVTriplets.triplets":
		lv := value.List()
		clv := lv.(*_DVVTriplets_1_list)
		x.Triplets = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.DVVTriplets"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.DVVTriplets does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DVVTriplets) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.DVVTriplets.triplets":
		if x.Triplets == nil {
			x.Triplets = []*DVVTriplet{}
		}
		value := &_DVVTriplets_1_list{list: &x.Triplets}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.DVVTriplets"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.DVVTriplets does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_DVVTriplets) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.DVVTriplets.triplets":
		list := []*DVVTriplet{}
		return protoreflect.ValueOfList(&_DVVTriplets_1_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.DVVTriplets"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.DVVTriplets does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_DVVTriplets) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.DVVTriplets", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_DVVTriplets) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DVVTriplets) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_DVVTriplets) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_DVVTriplets) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*DVVTriplets)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if len(x.Triplets) > 0 {
			for _, e := range x.Triplets {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*DVVTriplets)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Triplets) > 0 {
			for iNdEx := len(x.Triplets) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Triplets[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0xa
			}
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*DVVTriplets)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: DVVTriplets: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: DVVTriplets: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Triplets", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Triplets = append(x.Triplets, &DVVTriplet{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Triplets[len(x.Triplets)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_Delegation                   protoreflect.MessageDescriptor
	fd_Delegation_delegator_address protoreflect.FieldDescriptor
	fd_Delegation_validator_address protoreflect.FieldDescriptor
	fd_Delegation_shares            protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_staking_proto_init()
	md_Delegation = File_cosmos_staking_v1beta1_staking_proto.Messages().ByName("Delegation")
	fd_Delegation_delegator_address = md_Delegation.Fields().ByName("delegator_address")
	fd_Delegation_validator_address = md_Delegation.Fields().ByName("validator_address")
	fd_Delegation_shares = md_Delegation.Fields().ByName("shares")
}

var _ protoreflect.Message = (*fastReflection_Delegation)(nil)

type fastReflection_Delegation Delegation

func (x *Delegation) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Delegation)(x)
}

func (x *Delegation) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Delegation_messageType fastReflection_Delegation_messageType
var _ protoreflect.MessageType = fastReflection_Delegation_messageType{}

type fastReflection_Delegation_messageType struct{}

func (x fastReflection_Delegation_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Delegation)(nil)
}
func (x fastReflection_Delegation_messageType) New() protoreflect.Message {
	return new(fastReflection_Delegation)
}
func (x fastReflection_Delegation_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Delegation
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Delegation) Descriptor() protoreflect.MessageDescriptor {
	return md_Delegation
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Delegation) Type() protoreflect.MessageType {
	return _fastReflection_Delegation_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Delegation) New() protoreflect.Message {
	return new(fastReflection_Delegation)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Delegation) Interface() protoreflect.ProtoMessage {
	return (*Delegation)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Delegation) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.DelegatorAddress != "" {
		value := protoreflect.ValueOfString(x.DelegatorAddress)
		if !f(fd_Delegation_delegator_address, value) {
			return
		}
	}
	if x.ValidatorAddress != "" {
		value := protoreflect.ValueOfString(x.ValidatorAddress)
		if !f(fd_Delegation_validator_address, value) {
			return
		}
	}
	if x.Shares != "" {
		value := protoreflect.ValueOfString(x.Shares)
		if !f(fd_Delegation_shares, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Delegation) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.Delegation.delegator_address":
		return x.DelegatorAddress != ""
	case "cosmos.staking.v1beta1.Delegation.validator_address":
		return x.ValidatorAddress != ""
	case "cosmos.staking.v1beta1.Delegation.shares":
		return x.Shares != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Delegation"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Delegation does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Delegation) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.Delegation.delegator_address":
		x.DelegatorAddress = ""
	case "cosmos.staking.v1beta1.Delegation.validator_address":
		x.ValidatorAddress = ""
	case "cosmos.staking.v1beta1.Delegation.shares":
		x.Shares = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Delegation"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Delegation does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Delegation) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.Delegation.delegator_address":
		value := x.DelegatorAddress
		return protoreflect.ValueOfString(value)
	case "cosmos.staking.v1beta1.Delegation.validator_address":
		value := x.ValidatorAddress
		return protoreflect.ValueOfString(value)
	case "cosmos.staking.v1beta1.Delegation.shares":
		value := x.Shares
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Delegation"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Delegation does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Delegation) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.Delegation.delegator_address":
		x.DelegatorAddress = value.Interface().(string)
	case "cosmos.staking.v1beta1.Delegation.validator_address":
		x.ValidatorAddress = value.Interface().(string)
	case "cosmos.staking.v1beta1.Delegation.shares":
		x.Shares = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Delegation"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Delegation does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Delegation) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.Delegation.delegator_address":
		panic(fmt.Errorf("field delegator_address of message cosmos.staking.v1beta1.Delegation is not mutable"))
	case "cosmos.staking.v1beta1.Delegation.validator_address":
		panic(fmt.Errorf("field validator_address of message cosmos.staking.v1beta1.Delegation is not mutable"))
	case "cosmos.staking.v1beta1.Delegation.shares":
		panic(fmt.Errorf("field shares of message cosmos.staking.v1beta1.Delegation is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Delegation"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Delegation does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Delegation) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.Delegation.delegator_address":
		return protoreflect.ValueOfString("")
	case "cosmos.staking.v1beta1.Delegation.validator_address":
		return protoreflect.ValueOfString("")
	case "cosmos.staking.v1beta1.Delegation.shares":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Delegation"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Delegation does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Delegation) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.Delegation", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Delegation) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Delegation) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Delegation) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Delegation) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Delegation)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.DelegatorAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ValidatorAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Shares)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Delegation)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Shares) > 0 {
			i -= len(x.Shares)
			copy(dAtA[i:], x.Shares)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Shares)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.ValidatorAddress) > 0 {
			i -= len(x.ValidatorAddress)
			copy(dAtA[i:], x.ValidatorAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ValidatorAddress)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.DelegatorAddress) > 0 {
			i -= len(x.DelegatorAddress)
			copy(dAtA[i:], x.DelegatorAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.DelegatorAddress)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Delegation)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Delegation: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Delegation: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.DelegatorAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ValidatorAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Shares", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Shares = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_UnbondingDelegation_3_list)(nil)

type _UnbondingDelegation_3_list struct {
	list *[]*UnbondingDelegationEntry
}

func (x *_UnbondingDelegation_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_UnbondingDelegation_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_UnbondingDelegation_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*UnbondingDelegationEntry)
	(*x.list)[i] = concreteValue
}

func (x *_UnbondingDelegation_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*UnbondingDelegationEntry)
	*x.list = append(*x.list, concreteValue)
}

func (x *_UnbondingDelegation_3_list) AppendMutable() protoreflect.Value {
	v := new(UnbondingDelegationEntry)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_UnbondingDelegation_3_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_UnbondingDelegation_3_list) NewElement() protoreflect.Value {
	v := new(UnbondingDelegationEntry)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_UnbondingDelegation_3_list) IsValid() bool {
	return x.list != nil
}

var (
	md_UnbondingDelegation                   protoreflect.MessageDescriptor
	fd_UnbondingDelegation_delegator_address protoreflect.FieldDescriptor
	fd_UnbondingDelegation_validator_address protoreflect.FieldDescriptor
	fd_UnbondingDelegation_entries           protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_staking_proto_init()
	md_UnbondingDelegation = File_cosmos_staking_v1beta1_staking_proto.Messages().ByName("UnbondingDelegation")
	fd_UnbondingDelegation_delegator_address = md_UnbondingDelegation.Fields().ByName("delegator_address")
	fd_UnbondingDelegation_validator_address = md_UnbondingDelegation.Fields().ByName("validator_address")
	fd_UnbondingDelegation_entries = md_UnbondingDelegation.Fields().ByName("entries")
}

var _ protoreflect.Message = (*fastReflection_UnbondingDelegation)(nil)

type fastReflection_UnbondingDelegation UnbondingDelegation

func (x *UnbondingDelegation) ProtoReflect() protoreflect.Message {
	return (*fastReflection_UnbondingDelegation)(x)
}

func (x *UnbondingDelegation) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_UnbondingDelegation_messageType fastReflection_UnbondingDelegation_messageType
var _ protoreflect.MessageType = fastReflection_UnbondingDelegation_messageType{}

type fastReflection_UnbondingDelegation_messageType struct{}

func (x fastReflection_UnbondingDelegation_messageType) Zero() protoreflect.Message {
	return (*fastReflection_UnbondingDelegation)(nil)
}
func (x fastReflection_UnbondingDelegation_messageType) New() protoreflect.Message {
	return new(fastReflection_UnbondingDelegation)
}
func (x fastReflection_UnbondingDelegation_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_UnbondingDelegation
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_UnbondingDelegation) Descriptor() protoreflect.MessageDescriptor {
	return md_UnbondingDelegation
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_UnbondingDelegation) Type() protoreflect.MessageType {
	return _fastReflection_UnbondingDelegation_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_UnbondingDelegation) New() protoreflect.Message {
	return new(fastReflection_UnbondingDelegation)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_UnbondingDelegation) Interface() protoreflect.ProtoMessage {
	return (*UnbondingDelegation)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_UnbondingDelegation) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.DelegatorAddress != "" {
		value := protoreflect.ValueOfString(x.DelegatorAddress)
		if !f(fd_UnbondingDelegation_delegator_address, value) {
			return
		}
	}
	if x.ValidatorAddress != "" {
		value := protoreflect.ValueOfString(x.ValidatorAddress)
		if !f(fd_UnbondingDelegation_validator_address, value) {
			return
		}
	}
	if len(x.Entries) != 0 {
		value := protoreflect.ValueOfList(&_UnbondingDelegation_3_list{list: &x.Entries})
		if !f(fd_UnbondingDelegation_entries, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_UnbondingDelegation) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.UnbondingDelegation.delegator_address":
		return x.DelegatorAddress != ""
	case "cosmos.staking.v1beta1.UnbondingDelegation.validator_address":
		return x.ValidatorAddress != ""
	case "cosmos.staking.v1beta1.UnbondingDelegation.entries":
		return len(x.Entries) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.UnbondingDelegation"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.UnbondingDelegation does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_UnbondingDelegation) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.UnbondingDelegation.delegator_address":
		x.DelegatorAddress = ""
	case "cosmos.staking.v1beta1.UnbondingDelegation.validator_address":
		x.ValidatorAddress = ""
	case "cosmos.staking.v1beta1.UnbondingDelegation.entries":
		x.Entries = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.UnbondingDelegation"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.UnbondingDelegation does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_UnbondingDelegation) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.UnbondingDelegation.delegator_address":
		value := x.DelegatorAddress
		return protoreflect.ValueOfString(value)
	case "cosmos.staking.v1beta1.UnbondingDelegation.validator_address":
		value := x.ValidatorAddress
		return protoreflect.ValueOfString(value)
	case "cosmos.staking.v1beta1.UnbondingDelegation.entries":
		if len(x.Entries) == 0 {
			return protoreflect.ValueOfList(&_UnbondingDelegation_3_list{})
		}
		listValue := &_UnbondingDelegation_3_list{list: &x.Entries}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.UnbondingDelegation"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.UnbondingDelegation does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_UnbondingDelegation) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.UnbondingDelegation.delegator_address":
		x.DelegatorAddress = value.Interface().(string)
	case "cosmos.staking.v1beta1.UnbondingDelegation.validator_address":
		x.ValidatorAddress = value.Interface().(string)
	case "cosmos.staking.v1beta1.UnbondingDelegation.entries":
		lv := value.List()
		clv := lv.(*_UnbondingDelegation_3_list)
		x.Entries = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.UnbondingDelegation"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.UnbondingDelegation does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_UnbondingDelegation) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.UnbondingDelegation.entries":
		if x.Entries == nil {
			x.Entries = []*UnbondingDelegationEntry{}
		}
		value := &_UnbondingDelegation_3_list{list: &x.Entries}
		return protoreflect.ValueOfList(value)
	case "cosmos.staking.v1beta1.UnbondingDelegation.delegator_address":
		panic(fmt.Errorf("field delegator_address of message cosmos.staking.v1beta1.UnbondingDelegation is not mutable"))
	case "cosmos.staking.v1beta1.UnbondingDelegation.validator_address":
		panic(fmt.Errorf("field validator_address of message cosmos.staking.v1beta1.UnbondingDelegation is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.UnbondingDelegation"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.UnbondingDelegation does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_UnbondingDelegation) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.UnbondingDelegation.delegator_address":
		return protoreflect.ValueOfString("")
	case "cosmos.staking.v1beta1.UnbondingDelegation.validator_address":
		return protoreflect.ValueOfString("")
	case "cosmos.staking.v1beta1.UnbondingDelegation.entries":
		list := []*UnbondingDelegationEntry{}
		return protoreflect.ValueOfList(&_UnbondingDelegation_3_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.UnbondingDelegation"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.UnbondingDelegation does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_UnbondingDelegation) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.UnbondingDelegation", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_UnbondingDelegation) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_UnbondingDelegation) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_UnbondingDelegation) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_UnbondingDelegation) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*UnbondingDelegation)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.DelegatorAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ValidatorAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Entries) > 0 {
			for _, e := range x.Entries {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*UnbondingDelegation)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Entries) > 0 {
			for iNdEx := len(x.Entries) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Entries[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x1a
			}
		}
		if len(x.ValidatorAddress) > 0 {
			i -= len(x.ValidatorAddress)
			copy(dAtA[i:], x.ValidatorAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ValidatorAddress)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.DelegatorAddress) > 0 {
			i -= len(x.DelegatorAddress)
			copy(dAtA[i:], x.DelegatorAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.DelegatorAddress)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*UnbondingDelegation)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: UnbondingDelegation: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: UnbondingDelegation: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.DelegatorAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ValidatorAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Entries", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Entries = append(x.Entries, &UnbondingDelegationEntry{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Entries[len(x.Entries)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_UnbondingDelegationEntry                 protoreflect.MessageDescriptor
	fd_UnbondingDelegationEntry_creation_height protoreflect.FieldDescriptor
	fd_UnbondingDelegationEntry_completion_time protoreflect.FieldDescriptor
	fd_UnbondingDelegationEntry_initial_balance protoreflect.FieldDescriptor
	fd_UnbondingDelegationEntry_balance         protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_staking_proto_init()
	md_UnbondingDelegationEntry = File_cosmos_staking_v1beta1_staking_proto.Messages().ByName("UnbondingDelegationEntry")
	fd_UnbondingDelegationEntry_creation_height = md_UnbondingDelegationEntry.Fields().ByName("creation_height")
	fd_UnbondingDelegationEntry_completion_time = md_UnbondingDelegationEntry.Fields().ByName("completion_time")
	fd_UnbondingDelegationEntry_initial_balance = md_UnbondingDelegationEntry.Fields().ByName("initial_balance")
	fd_UnbondingDelegationEntry_balance = md_UnbondingDelegationEntry.Fields().ByName("balance")
}

var _ protoreflect.Message = (*fastReflection_UnbondingDelegationEntry)(nil)

type fastReflection_UnbondingDelegationEntry UnbondingDelegationEntry

func (x *UnbondingDelegationEntry) ProtoReflect() protoreflect.Message {
	return (*fastReflection_UnbondingDelegationEntry)(x)
}

func (x *UnbondingDelegationEntry) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_UnbondingDelegationEntry_messageType fastReflection_UnbondingDelegationEntry_messageType
var _ protoreflect.MessageType = fastReflection_UnbondingDelegationEntry_messageType{}

type fastReflection_UnbondingDelegationEntry_messageType struct{}

func (x fastReflection_UnbondingDelegationEntry_messageType) Zero() protoreflect.Message {
	return (*fastReflection_UnbondingDelegationEntry)(nil)
}
func (x fastReflection_UnbondingDelegationEntry_messageType) New() protoreflect.Message {
	return new(fastReflection_UnbondingDelegationEntry)
}
func (x fastReflection_UnbondingDelegationEntry_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_UnbondingDelegationEntry
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_UnbondingDelegationEntry) Descriptor() protoreflect.MessageDescriptor {
	return md_UnbondingDelegationEntry
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_UnbondingDelegationEntry) Type() protoreflect.MessageType {
	return _fastReflection_UnbondingDelegationEntry_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_UnbondingDelegationEntry) New() protoreflect.Message {
	return new(fastReflection_UnbondingDelegationEntry)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_UnbondingDelegationEntry) Interface() protoreflect.ProtoMessage {
	return (*UnbondingDelegationEntry)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_UnbondingDelegationEntry) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.CreationHeight != int64(0) {
		value := protoreflect.ValueOfInt64(x.CreationHeight)
		if !f(fd_UnbondingDelegationEntry_creation_height, value) {
			return
		}
	}
	if x.CompletionTime != nil {
		value := protoreflect.ValueOfMessage(x.CompletionTime.ProtoReflect())
		if !f(fd_UnbondingDelegationEntry_completion_time, value) {
			return
		}
	}
	if x.InitialBalance != "" {
		value := protoreflect.ValueOfString(x.InitialBalance)
		if !f(fd_UnbondingDelegationEntry_initial_balance, value) {
			return
		}
	}
	if x.Balance != "" {
		value := protoreflect.ValueOfString(x.Balance)
		if !f(fd_UnbondingDelegationEntry_balance, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_UnbondingDelegationEntry) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.UnbondingDelegationEntry.creation_height":
		return x.CreationHeight != int64(0)
	case "cosmos.staking.v1beta1.UnbondingDelegationEntry.completion_time":
		return x.CompletionTime != nil
	case "cosmos.staking.v1beta1.UnbondingDelegationEntry.initial_balance":
		return x.InitialBalance != ""
	case "cosmos.staking.v1beta1.UnbondingDelegationEntry.balance":
		return x.Balance != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.UnbondingDelegationEntry"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.UnbondingDelegationEntry does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_UnbondingDelegationEntry) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.UnbondingDelegationEntry.creation_height":
		x.CreationHeight = int64(0)
	case "cosmos.staking.v1beta1.UnbondingDelegationEntry.completion_time":
		x.CompletionTime = nil
	case "cosmos.staking.v1beta1.UnbondingDelegationEntry.initial_balance":
		x.InitialBalance = ""
	case "cosmos.staking.v1beta1.UnbondingDelegationEntry.balance":
		x.Balance = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.UnbondingDelegationEntry"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.UnbondingDelegationEntry does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_UnbondingDelegationEntry) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.UnbondingDelegationEntry.creation_height":
		value := x.CreationHeight
		return protoreflect.ValueOfInt64(value)
	case "cosmos.staking.v1beta1.UnbondingDelegationEntry.completion_time":
		value := x.CompletionTime
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.staking.v1beta1.UnbondingDelegationEntry.initial_balance":
		value := x.InitialBalance
		return protoreflect.ValueOfString(value)
	case "cosmos.staking.v1beta1.UnbondingDelegationEntry.balance":
		value := x.Balance
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.UnbondingDelegationEntry"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.UnbondingDelegationEntry does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_UnbondingDelegationEntry) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.UnbondingDelegationEntry.creation_height":
		x.CreationHeight = value.Int()
	case "cosmos.staking.v1beta1.UnbondingDelegationEntry.completion_time":
		x.CompletionTime = value.Message().Interface().(*timestamppb.Timestamp)
	case "cosmos.staking.v1beta1.UnbondingDelegationEntry.initial_balance":
		x.InitialBalance = value.Interface().(string)
	case "cosmos.staking.v1beta1.UnbondingDelegationEntry.balance":
		x.Balance = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.UnbondingDelegationEntry"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.UnbondingDelegationEntry does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_UnbondingDelegationEntry) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.UnbondingDelegationEntry.completion_time":
		if x.CompletionTime == nil {
			x.CompletionTime = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.CompletionTime.ProtoReflect())
	case "cosmos.staking.v1beta1.UnbondingDelegationEntry.creation_height":
		panic(fmt.Errorf("field creation_height of message cosmos.staking.v1beta1.UnbondingDelegationEntry is not mutable"))
	case "cosmos.staking.v1beta1.UnbondingDelegationEntry.initial_balance":
		panic(fmt.Errorf("field initial_balance of message cosmos.staking.v1beta1.UnbondingDelegationEntry is not mutable"))
	case "cosmos.staking.v1beta1.UnbondingDelegationEntry.balance":
		panic(fmt.Errorf("field balance of message cosmos.staking.v1beta1.UnbondingDelegationEntry is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.UnbondingDelegationEntry"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.UnbondingDelegationEntry does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_UnbondingDelegationEntry) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.UnbondingDelegationEntry.creation_height":
		return protoreflect.ValueOfInt64(int64(0))
	case "cosmos.staking.v1beta1.UnbondingDelegationEntry.completion_time":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.staking.v1beta1.UnbondingDelegationEntry.initial_balance":
		return protoreflect.ValueOfString("")
	case "cosmos.staking.v1beta1.UnbondingDelegationEntry.balance":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.UnbondingDelegationEntry"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.UnbondingDelegationEntry does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_UnbondingDelegationEntry) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.UnbondingDelegationEntry", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_UnbondingDelegationEntry) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_UnbondingDelegationEntry) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_UnbondingDelegationEntry) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_UnbondingDelegationEntry) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*UnbondingDelegationEntry)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.CreationHeight != 0 {
			n += 1 + runtime.Sov(uint64(x.CreationHeight))
		}
		if x.CompletionTime != nil {
			l = options.Size(x.CompletionTime)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.InitialBalance)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Balance)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*UnbondingDelegationEntry)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Balance) > 0 {
			i -= len(x.Balance)
			copy(dAtA[i:], x.Balance)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Balance)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.InitialBalance) > 0 {
			i -= len(x.InitialBalance)
			copy(dAtA[i:], x.InitialBalance)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.InitialBalance)))
			i--
			dAtA[i] = 0x1a
		}
		if x.CompletionTime != nil {
			encoded, err := options.Marshal(x.CompletionTime)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if x.CreationHeight != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.CreationHeight))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*UnbondingDelegationEntry)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: UnbondingDelegationEntry: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: UnbondingDelegationEntry: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field CreationHeight", wireType)
				}
				x.CreationHeight = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.CreationHeight |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field CompletionTime", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.CompletionTime == nil {
					x.CompletionTime = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.CompletionTime); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field InitialBalance", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.InitialBalance = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Balance = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_RedelegationEntry                 protoreflect.MessageDescriptor
	fd_RedelegationEntry_creation_height protoreflect.FieldDescriptor
	fd_RedelegationEntry_completion_time protoreflect.FieldDescriptor
	fd_RedelegationEntry_initial_balance protoreflect.FieldDescriptor
	fd_RedelegationEntry_shares_dst      protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_staking_proto_init()
	md_RedelegationEntry = File_cosmos_staking_v1beta1_staking_proto.Messages().ByName("RedelegationEntry")
	fd_RedelegationEntry_creation_height = md_RedelegationEntry.Fields().ByName("creation_height")
	fd_RedelegationEntry_completion_time = md_RedelegationEntry.Fields().ByName("completion_time")
	fd_RedelegationEntry_initial_balance = md_RedelegationEntry.Fields().ByName("initial_balance")
	fd_RedelegationEntry_shares_dst = md_RedelegationEntry.Fields().ByName("shares_dst")
}

var _ protoreflect.Message = (*fastReflection_RedelegationEntry)(nil)

type fastReflection_RedelegationEntry RedelegationEntry

func (x *RedelegationEntry) ProtoReflect() protoreflect.Message {
	return (*fastReflection_RedelegationEntry)(x)
}

func (x *RedelegationEntry) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_RedelegationEntry_messageType fastReflection_RedelegationEntry_messageType
var _ protoreflect.MessageType = fastReflection_RedelegationEntry_messageType{}

type fastReflection_RedelegationEntry_messageType struct{}

func (x fastReflection_RedelegationEntry_messageType) Zero() protoreflect.Message {
	return (*fastReflection_RedelegationEntry)(nil)
}
func (x fastReflection_RedelegationEntry_messageType) New() protoreflect.Message {
	return new(fastReflection_RedelegationEntry)
}
func (x fastReflection_RedelegationEntry_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_RedelegationEntry
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_RedelegationEntry) Descriptor() protoreflect.MessageDescriptor {
	return md_RedelegationEntry
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_RedelegationEntry) Type() protoreflect.MessageType {
	return _fastReflection_RedelegationEntry_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_RedelegationEntry) New() protoreflect.Message {
	return new(fastReflection_RedelegationEntry)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_RedelegationEntry) Interface() protoreflect.ProtoMessage {
	return (*RedelegationEntry)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_RedelegationEntry) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.CreationHeight != int64(0) {
		value := protoreflect.ValueOfInt64(x.CreationHeight)
		if !f(fd_RedelegationEntry_creation_height, value) {
			return
		}
	}
	if x.CompletionTime != nil {
		value := protoreflect.ValueOfMessage(x.CompletionTime.ProtoReflect())
		if !f(fd_RedelegationEntry_completion_time, value) {
			return
		}
	}
	if x.InitialBalance != "" {
		value := protoreflect.ValueOfString(x.InitialBalance)
		if !f(fd_RedelegationEntry_initial_balance, value) {
			return
		}
	}
	if x.SharesDst != "" {
		value := protoreflect.ValueOfString(x.SharesDst)
		if !f(fd_RedelegationEntry_shares_dst, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_RedelegationEntry) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.RedelegationEntry.creation_height":
		return x.CreationHeight != int64(0)
	case "cosmos.staking.v1beta1.RedelegationEntry.completion_time":
		return x.CompletionTime != nil
	case "cosmos.staking.v1beta1.RedelegationEntry.initial_balance":
		return x.InitialBalance != ""
	case "cosmos.staking.v1beta1.RedelegationEntry.shares_dst":
		return x.SharesDst != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.RedelegationEntry"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.RedelegationEntry does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_RedelegationEntry) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.RedelegationEntry.creation_height":
		x.CreationHeight = int64(0)
	case "cosmos.staking.v1beta1.RedelegationEntry.completion_time":
		x.CompletionTime = nil
	case "cosmos.staking.v1beta1.RedelegationEntry.initial_balance":
		x.InitialBalance = ""
	case "cosmos.staking.v1beta1.RedelegationEntry.shares_dst":
		x.SharesDst = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.RedelegationEntry"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.RedelegationEntry does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_RedelegationEntry) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.RedelegationEntry.creation_height":
		value := x.CreationHeight
		return protoreflect.ValueOfInt64(value)
	case "cosmos.staking.v1beta1.RedelegationEntry.completion_time":
		value := x.CompletionTime
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.staking.v1beta1.RedelegationEntry.initial_balance":
		value := x.InitialBalance
		return protoreflect.ValueOfString(value)
	case "cosmos.staking.v1beta1.RedelegationEntry.shares_dst":
		value := x.SharesDst
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.RedelegationEntry"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.RedelegationEntry does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_RedelegationEntry) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.RedelegationEntry.creation_height":
		x.CreationHeight = value.Int()
	case "cosmos.staking.v1beta1.RedelegationEntry.completion_time":
		x.CompletionTime = value.Message().Interface().(*timestamppb.Timestamp)
	case "cosmos.staking.v1beta1.RedelegationEntry.initial_balance":
		x.InitialBalance = value.Interface().(string)
	case "cosmos.staking.v1beta1.RedelegationEntry.shares_dst":
		x.SharesDst = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.RedelegationEntry"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.RedelegationEntry does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_RedelegationEntry) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.RedelegationEntry.completion_time":
		if x.CompletionTime == nil {
			x.CompletionTime = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.CompletionTime.ProtoReflect())
	case "cosmos.staking.v1beta1.RedelegationEntry.creation_height":
		panic(fmt.Errorf("field creation_height of message cosmos.staking.v1beta1.RedelegationEntry is not mutable"))
	case "cosmos.staking.v1beta1.RedelegationEntry.initial_balance":
		panic(fmt.Errorf("field initial_balance of message cosmos.staking.v1beta1.RedelegationEntry is not mutable"))
	case "cosmos.staking.v1beta1.RedelegationEntry.shares_dst":
		panic(fmt.Errorf("field shares_dst of message cosmos.staking.v1beta1.RedelegationEntry is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.RedelegationEntry"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.RedelegationEntry does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_RedelegationEntry) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.RedelegationEntry.creation_height":
		return protoreflect.ValueOfInt64(int64(0))
	case "cosmos.staking.v1beta1.RedelegationEntry.completion_time":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.staking.v1beta1.RedelegationEntry.initial_balance":
		return protoreflect.ValueOfString("")
	case "cosmos.staking.v1beta1.RedelegationEntry.shares_dst":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.RedelegationEntry"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.RedelegationEntry does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_RedelegationEntry) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.RedelegationEntry", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_RedelegationEntry) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_RedelegationEntry) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_RedelegationEntry) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_RedelegationEntry) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*RedelegationEntry)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.CreationHeight != 0 {
			n += 1 + runtime.Sov(uint64(x.CreationHeight))
		}
		if x.CompletionTime != nil {
			l = options.Size(x.CompletionTime)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.InitialBalance)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.SharesDst)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*RedelegationEntry)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.SharesDst) > 0 {
			i -= len(x.SharesDst)
			copy(dAtA[i:], x.SharesDst)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.SharesDst)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.InitialBalance) > 0 {
			i -= len(x.InitialBalance)
			copy(dAtA[i:], x.InitialBalance)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.InitialBalance)))
			i--
			dAtA[i] = 0x1a
		}
		if x.CompletionTime != nil {
			encoded, err := options.Marshal(x.CompletionTime)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if x.CreationHeight != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.CreationHeight))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*RedelegationEntry)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: RedelegationEntry: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: RedelegationEntry: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field CreationHeight", wireType)
				}
				x.CreationHeight = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.CreationHeight |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field CompletionTime", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.CompletionTime == nil {
					x.CompletionTime = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.CompletionTime); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field InitialBalance", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.InitialBalance = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SharesDst", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.SharesDst = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_Redelegation_4_list)(nil)

type _Redelegation_4_list struct {
	list *[]*RedelegationEntry
}

func (x *_Redelegation_4_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_Redelegation_4_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_Redelegation_4_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*RedelegationEntry)
	(*x.list)[i] = concreteValue
}

func (x *_Redelegation_4_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*RedelegationEntry)
	*x.list = append(*x.list, concreteValue)
}

func (x *_Redelegation_4_list) AppendMutable() protoreflect.Value {
	v := new(RedelegationEntry)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Redelegation_4_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_Redelegation_4_list) NewElement() protoreflect.Value {
	v := new(RedelegationEntry)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Redelegation_4_list) IsValid() bool {
	return x.list != nil
}

var (
	md_Redelegation                       protoreflect.MessageDescriptor
	fd_Redelegation_delegator_address     protoreflect.FieldDescriptor
	fd_Redelegation_validator_src_address protoreflect.FieldDescriptor
	fd_Redelegation_validator_dst_address protoreflect.FieldDescriptor
	fd_Redelegation_entries               protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_staking_proto_init()
	md_Redelegation = File_cosmos_staking_v1beta1_staking_proto.Messages().ByName("Redelegation")
	fd_Redelegation_delegator_address = md_Redelegation.Fields().ByName("delegator_address")
	fd_Redelegation_validator_src_address = md_Redelegation.Fields().ByName("validator_src_address")
	fd_Redelegation_validator_dst_address = md_Redelegation.Fields().ByName("validator_dst_address")
	fd_Redelegation_entries = md_Redelegation.Fields().ByName("entries")
}

var _ protoreflect.Message = (*fastReflection_Redelegation)(nil)

type fastReflection_Redelegation Redelegation

func (x *Redelegation) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Redelegation)(x)
}

func (x *Redelegation) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Redelegation_messageType fastReflection_Redelegation_messageType
var _ protoreflect.MessageType = fastReflection_Redelegation_messageType{}

type fastReflection_Redelegation_messageType struct{}

func (x fastReflection_Redelegation_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Redelegation)(nil)
}
func (x fastReflection_Redelegation_messageType) New() protoreflect.Message {
	return new(fastReflection_Redelegation)
}
func (x fastReflection_Redelegation_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Redelegation
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Redelegation) Descriptor() protoreflect.MessageDescriptor {
	return md_Redelegation
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Redelegation) Type() protoreflect.MessageType {
	return _fastReflection_Redelegation_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Redelegation) New() protoreflect.Message {
	return new(fastReflection_Redelegation)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Redelegation) Interface() protoreflect.ProtoMessage {
	return (*Redelegation)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Redelegation) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.DelegatorAddress != "" {
		value := protoreflect.ValueOfString(x.DelegatorAddress)
		if !f(fd_Redelegation_delegator_address, value) {
			return
		}
	}
	if x.ValidatorSrcAddress != "" {
		value := protoreflect.ValueOfString(x.ValidatorSrcAddress)
		if !f(fd_Redelegation_validator_src_address, value) {
			return
		}
	}
	if x.ValidatorDstAddress != "" {
		value := protoreflect.ValueOfString(x.ValidatorDstAddress)
		if !f(fd_Redelegation_validator_dst_address, value) {
			return
		}
	}
	if len(x.Entries) != 0 {
		value := protoreflect.ValueOfList(&_Redelegation_4_list{list: &x.Entries})
		if !f(fd_Redelegation_entries, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Redelegation) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.Redelegation.delegator_address":
		return x.DelegatorAddress != ""
	case "cosmos.staking.v1beta1.Redelegation.validator_src_address":
		return x.ValidatorSrcAddress != ""
	case "cosmos.staking.v1beta1.Redelegation.validator_dst_address":
		return x.ValidatorDstAddress != ""
	case "cosmos.staking.v1beta1.Redelegation.entries":
		return len(x.Entries) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Redelegation"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Redelegation does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Redelegation) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.Redelegation.delegator_address":
		x.DelegatorAddress = ""
	case "cosmos.staking.v1beta1.Redelegation.validator_src_address":
		x.ValidatorSrcAddress = ""
	case "cosmos.staking.v1beta1.Redelegation.validator_dst_address":
		x.ValidatorDstAddress = ""
	case "cosmos.staking.v1beta1.Redelegation.entries":
		x.Entries = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Redelegation"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Redelegation does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Redelegation) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.Redelegation.delegator_address":
		value := x.DelegatorAddress
		return protoreflect.ValueOfString(value)
	case "cosmos.staking.v1beta1.Redelegation.validator_src_address":
		value := x.ValidatorSrcAddress
		return protoreflect.ValueOfString(value)
	case "cosmos.staking.v1beta1.Redelegation.validator_dst_address":
		value := x.ValidatorDstAddress
		return protoreflect.ValueOfString(value)
	case "cosmos.staking.v1beta1.Redelegation.entries":
		if len(x.Entries) == 0 {
			return protoreflect.ValueOfList(&_Redelegation_4_list{})
		}
		listValue := &_Redelegation_4_list{list: &x.Entries}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Redelegation"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Redelegation does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Redelegation) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.Redelegation.delegator_address":
		x.DelegatorAddress = value.Interface().(string)
	case "cosmos.staking.v1beta1.Redelegation.validator_src_address":
		x.ValidatorSrcAddress = value.Interface().(string)
	case "cosmos.staking.v1beta1.Redelegation.validator_dst_address":
		x.ValidatorDstAddress = value.Interface().(string)
	case "cosmos.staking.v1beta1.Redelegation.entries":
		lv := value.List()
		clv := lv.(*_Redelegation_4_list)
		x.Entries = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Redelegation"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Redelegation does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Redelegation) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.Redelegation.entries":
		if x.Entries == nil {
			x.Entries = []*RedelegationEntry{}
		}
		value := &_Redelegation_4_list{list: &x.Entries}
		return protoreflect.ValueOfList(value)
	case "cosmos.staking.v1beta1.Redelegation.delegator_address":
		panic(fmt.Errorf("field delegator_address of message cosmos.staking.v1beta1.Redelegation is not mutable"))
	case "cosmos.staking.v1beta1.Redelegation.validator_src_address":
		panic(fmt.Errorf("field validator_src_address of message cosmos.staking.v1beta1.Redelegation is not mutable"))
	case "cosmos.staking.v1beta1.Redelegation.validator_dst_address":
		panic(fmt.Errorf("field validator_dst_address of message cosmos.staking.v1beta1.Redelegation is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Redelegation"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Redelegation does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Redelegation) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.Redelegation.delegator_address":
		return protoreflect.ValueOfString("")
	case "cosmos.staking.v1beta1.Redelegation.validator_src_address":
		return protoreflect.ValueOfString("")
	case "cosmos.staking.v1beta1.Redelegation.validator_dst_address":
		return protoreflect.ValueOfString("")
	case "cosmos.staking.v1beta1.Redelegation.entries":
		list := []*RedelegationEntry{}
		return protoreflect.ValueOfList(&_Redelegation_4_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Redelegation"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Redelegation does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Redelegation) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.Redelegation", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Redelegation) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Redelegation) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Redelegation) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Redelegation) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Redelegation)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.DelegatorAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ValidatorSrcAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ValidatorDstAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Entries) > 0 {
			for _, e := range x.Entries {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Redelegation)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Entries) > 0 {
			for iNdEx := len(x.Entries) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Entries[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x22
			}
		}
		if len(x.ValidatorDstAddress) > 0 {
			i -= len(x.ValidatorDstAddress)
			copy(dAtA[i:], x.ValidatorDstAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ValidatorDstAddress)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.ValidatorSrcAddress) > 0 {
			i -= len(x.ValidatorSrcAddress)
			copy(dAtA[i:], x.ValidatorSrcAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ValidatorSrcAddress)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.DelegatorAddress) > 0 {
			i -= len(x.DelegatorAddress)
			copy(dAtA[i:], x.DelegatorAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.DelegatorAddress)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Redelegation)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Redelegation: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Redelegation: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.DelegatorAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ValidatorSrcAddress", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ValidatorSrcAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ValidatorDstAddress", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ValidatorDstAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Entries", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Entries = append(x.Entries, &RedelegationEntry{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Entries[len(x.Entries)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_Params                     protoreflect.MessageDescriptor
	fd_Params_unbonding_time      protoreflect.FieldDescriptor
	fd_Params_max_validators      protoreflect.FieldDescriptor
	fd_Params_max_entries         protoreflect.FieldDescriptor
	fd_Params_historical_entries  protoreflect.FieldDescriptor
	fd_Params_bond_denom          protoreflect.FieldDescriptor
	fd_Params_min_commission_rate protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_staking_proto_init()
	md_Params = File_cosmos_staking_v1beta1_staking_proto.Messages().ByName("Params")
	fd_Params_unbonding_time = md_Params.Fields().ByName("unbonding_time")
	fd_Params_max_validators = md_Params.Fields().ByName("max_validators")
	fd_Params_max_entries = md_Params.Fields().ByName("max_entries")
	fd_Params_historical_entries = md_Params.Fields().ByName("historical_entries")
	fd_Params_bond_denom = md_Params.Fields().ByName("bond_denom")
	fd_Params_min_commission_rate = md_Params.Fields().ByName("min_commission_rate")
}

var _ protoreflect.Message = (*fastReflection_Params)(nil)

type fastReflection_Params Params

func (x *Params) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Params)(x)
}

func (x *Params) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Params_messageType fastReflection_Params_messageType
var _ protoreflect.MessageType = fastReflection_Params_messageType{}

type fastReflection_Params_messageType struct{}

func (x fastReflection_Params_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Params)(nil)
}
func (x fastReflection_Params_messageType) New() protoreflect.Message {
	return new(fastReflection_Params)
}
func (x fastReflection_Params_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Params
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Params) Descriptor() protoreflect.MessageDescriptor {
	return md_Params
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Params) Type() protoreflect.MessageType {
	return _fastReflection_Params_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Params) New() protoreflect.Message {
	return new(fastReflection_Params)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Params) Interface() protoreflect.ProtoMessage {
	return (*Params)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Params) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.UnbondingTime != nil {
		value := protoreflect.ValueOfMessage(x.UnbondingTime.ProtoReflect())
		if !f(fd_Params_unbonding_time, value) {
			return
		}
	}
	if x.MaxValidators != uint32(0) {
		value := protoreflect.ValueOfUint32(x.MaxValidators)
		if !f(fd_Params_max_validators, value) {
			return
		}
	}
	if x.MaxEntries != uint32(0) {
		value := protoreflect.ValueOfUint32(x.MaxEntries)
		if !f(fd_Params_max_entries, value) {
			return
		}
	}
	if x.HistoricalEntries != uint32(0) {
		value := protoreflect.ValueOfUint32(x.HistoricalEntries)
		if !f(fd_Params_historical_entries, value) {
			return
		}
	}
	if x.BondDenom != "" {
		value := protoreflect.ValueOfString(x.BondDenom)
		if !f(fd_Params_bond_denom, value) {
			return
		}
	}
	if x.MinCommissionRate != "" {
		value := protoreflect.ValueOfString(x.MinCommissionRate)
		if !f(fd_Params_min_commission_rate, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Params) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.Params.unbonding_time":
		return x.UnbondingTime != nil
	case "cosmos.staking.v1beta1.Params.max_validators":
		return x.MaxValidators != uint32(0)
	case "cosmos.staking.v1beta1.Params.max_entries":
		return x.MaxEntries != uint32(0)
	case "cosmos.staking.v1beta1.Params.historical_entries":
		return x.HistoricalEntries != uint32(0)
	case "cosmos.staking.v1beta1.Params.bond_denom":
		return x.BondDenom != ""
	case "cosmos.staking.v1beta1.Params.min_commission_rate":
		return x.MinCommissionRate != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Params"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Params does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.Params.unbonding_time":
		x.UnbondingTime = nil
	case "cosmos.staking.v1beta1.Params.max_validators":
		x.MaxValidators = uint32(0)
	case "cosmos.staking.v1beta1.Params.max_entries":
		x.MaxEntries = uint32(0)
	case "cosmos.staking.v1beta1.Params.historical_entries":
		x.HistoricalEntries = uint32(0)
	case "cosmos.staking.v1beta1.Params.bond_denom":
		x.BondDenom = ""
	case "cosmos.staking.v1beta1.Params.min_commission_rate":
		x.MinCommissionRate = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Params"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Params does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Params) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.Params.unbonding_time":
		value := x.UnbondingTime
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.staking.v1beta1.Params.max_validators":
		value := x.MaxValidators
		return protoreflect.ValueOfUint32(value)
	case "cosmos.staking.v1beta1.Params.max_entries":
		value := x.MaxEntries
		return protoreflect.ValueOfUint32(value)
	case "cosmos.staking.v1beta1.Params.historical_entries":
		value := x.HistoricalEntries
		return protoreflect.ValueOfUint32(value)
	case "cosmos.staking.v1beta1.Params.bond_denom":
		value := x.BondDenom
		return protoreflect.ValueOfString(value)
	case "cosmos.staking.v1beta1.Params.min_commission_rate":
		value := x.MinCommissionRate
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Params"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Params does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.Params.unbonding_time":
		x.UnbondingTime = value.Message().Interface().(*durationpb.Duration)
	case "cosmos.staking.v1beta1.Params.max_validators":
		x.MaxValidators = uint32(value.Uint())
	case "cosmos.staking.v1beta1.Params.max_entries":
		x.MaxEntries = uint32(value.Uint())
	case "cosmos.staking.v1beta1.Params.historical_entries":
		x.HistoricalEntries = uint32(value.Uint())
	case "cosmos.staking.v1beta1.Params.bond_denom":
		x.BondDenom = value.Interface().(string)
	case "cosmos.staking.v1beta1.Params.min_commission_rate":
		x.MinCommissionRate = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Params"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Params does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.Params.unbonding_time":
		if x.UnbondingTime == nil {
			x.UnbondingTime = new(durationpb.Duration)
		}
		return protoreflect.ValueOfMessage(x.UnbondingTime.ProtoReflect())
	case "cosmos.staking.v1beta1.Params.max_validators":
		panic(fmt.Errorf("field max_validators of message cosmos.staking.v1beta1.Params is not mutable"))
	case "cosmos.staking.v1beta1.Params.max_entries":
		panic(fmt.Errorf("field max_entries of message cosmos.staking.v1beta1.Params is not mutable"))
	case "cosmos.staking.v1beta1.Params.historical_entries":
		panic(fmt.Errorf("field historical_entries of message cosmos.staking.v1beta1.Params is not mutable"))
	case "cosmos.staking.v1beta1.Params.bond_denom":
		panic(fmt.Errorf("field bond_denom of message cosmos.staking.v1beta1.Params is not mutable"))
	case "cosmos.staking.v1beta1.Params.min_commission_rate":
		panic(fmt.Errorf("field min_commission_rate of message cosmos.staking.v1beta1.Params is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Params"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Params does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Params) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.Params.unbonding_time":
		m := new(durationpb.Duration)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.staking.v1beta1.Params.max_validators":
		return protoreflect.ValueOfUint32(uint32(0))
	case "cosmos.staking.v1beta1.Params.max_entries":
		return protoreflect.ValueOfUint32(uint32(0))
	case "cosmos.staking.v1beta1.Params.historical_entries":
		return protoreflect.ValueOfUint32(uint32(0))
	case "cosmos.staking.v1beta1.Params.bond_denom":
		return protoreflect.ValueOfString("")
	case "cosmos.staking.v1beta1.Params.min_commission_rate":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Params"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Params does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Params) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.Params", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Params) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Params) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Params) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Params)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.UnbondingTime != nil {
			l = options.Size(x.UnbondingTime)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.MaxValidators != 0 {
			n += 1 + runtime.Sov(uint64(x.MaxValidators))
		}
		if x.MaxEntries != 0 {
			n += 1 + runtime.Sov(uint64(x.MaxEntries))
		}
		if x.HistoricalEntries != 0 {
			n += 1 + runtime.Sov(uint64(x.HistoricalEntries))
		}
		l = len(x.BondDenom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.MinCommissionRate)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Params)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.MinCommissionRate) > 0 {
			i -= len(x.MinCommissionRate)
			copy(dAtA[i:], x.MinCommissionRate)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.MinCommissionRate)))
			i--
			dAtA[i] = 0x32
		}
		if len(x.BondDenom) > 0 {
			i -= len(x.BondDenom)
			copy(dAtA[i:], x.BondDenom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BondDenom)))
			i--
			dAtA[i] = 0x2a
		}
		if x.HistoricalEntries != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.HistoricalEntries))
			i--
			dAtA[i] = 0x20
		}
		if x.MaxEntries != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.MaxEntries))
			i--
			dAtA[i] = 0x18
		}
		if x.MaxValidators != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.MaxValidators))
			i--
			dAtA[i] = 0x10
		}
		if x.UnbondingTime != nil {
			encoded, err := options.Marshal(x.UnbondingTime)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Params)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Params: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field UnbondingTime", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.UnbondingTime == nil {
					x.UnbondingTime = &durationpb.Duration{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.UnbondingTime); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MaxValidators", wireType)
				}
				x.MaxValidators = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.MaxValidators |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MaxEntries", wireType)
				}
				x.MaxEntries = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.MaxEntries |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field HistoricalEntries", wireType)
				}
				x.HistoricalEntries = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.HistoricalEntries |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BondDenom", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.BondDenom = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 6:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MinCommissionRate", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.MinCommissionRate = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_DelegationResponse            protoreflect.MessageDescriptor
	fd_DelegationResponse_delegation protoreflect.FieldDescriptor
	fd_DelegationResponse_balance    protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_staking_proto_init()
	md_DelegationResponse = File_cosmos_staking_v1beta1_staking_proto.Messages().ByName("DelegationResponse")
	fd_DelegationResponse_delegation = md_DelegationResponse.Fields().ByName("delegation")
	fd_DelegationResponse_balance = md_DelegationResponse.Fields().ByName("balance")
}

var _ protoreflect.Message = (*fastReflection_DelegationResponse)(nil)

type fastReflection_DelegationResponse DelegationResponse

func (x *DelegationResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_DelegationResponse)(x)
}

func (x *DelegationResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_DelegationResponse_messageType fastReflection_DelegationResponse_messageType
var _ protoreflect.MessageType = fastReflection_DelegationResponse_messageType{}

type fastReflection_DelegationResponse_messageType struct{}

func (x fastReflection_DelegationResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_DelegationResponse)(nil)
}
func (x fastReflection_DelegationResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_DelegationResponse)
}
func (x fastReflection_DelegationResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_DelegationResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_DelegationResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_DelegationResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_DelegationResponse) Type() protoreflect.MessageType {
	return _fastReflection_DelegationResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_DelegationResponse) New() protoreflect.Message {
	return new(fastReflection_DelegationResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_DelegationResponse) Interface() protoreflect.ProtoMessage {
	return (*DelegationResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_DelegationResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Delegation != nil {
		value := protoreflect.ValueOfMessage(x.Delegation.ProtoReflect())
		if !f(fd_DelegationResponse_delegation, value) {
			return
		}
	}
	if x.Balance != nil {
		value := protoreflect.ValueOfMessage(x.Balance.ProtoReflect())
		if !f(fd_DelegationResponse_balance, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_DelegationResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.DelegationResponse.delegation":
		return x.Delegation != nil
	case "cosmos.staking.v1beta1.DelegationResponse.balance":
		return x.Balance != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.DelegationResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.DelegationResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DelegationResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.DelegationResponse.delegation":
		x.Delegation = nil
	case "cosmos.staking.v1beta1.DelegationResponse.balance":
		x.Balance = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.DelegationResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.DelegationResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_DelegationResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.DelegationResponse.delegation":
		value := x.Delegation
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.staking.v1beta1.DelegationResponse.balance":
		value := x.Balance
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.DelegationResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.DelegationResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DelegationResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.DelegationResponse.delegation":
		x.Delegation = value.Message().Interface().(*Delegation)
	case "cosmos.staking.v1beta1.DelegationResponse.balance":
		x.Balance = value.Message().Interface().(*v1beta1.Coin)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.DelegationResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.DelegationResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DelegationResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.DelegationResponse.delegation":
		if x.Delegation == nil {
			x.Delegation = new(Delegation)
		}
		return protoreflect.ValueOfMessage(x.Delegation.ProtoReflect())
	case "cosmos.staking.v1beta1.DelegationResponse.balance":
		if x.Balance == nil {
			x.Balance = new(v1beta1.Coin)
		}
		return protoreflect.ValueOfMessage(x.Balance.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.DelegationResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.DelegationResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_DelegationResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.DelegationResponse.delegation":
		m := new(Delegation)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.staking.v1beta1.DelegationResponse.balance":
		m := new(v1beta1.Coin)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.DelegationResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.DelegationResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_DelegationResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.DelegationResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_DelegationResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DelegationResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_DelegationResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_DelegationResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*DelegationResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.Delegation != nil {
			l = options.Size(x.Delegation)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Balance != nil {
			l = options.Size(x.Balance)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*DelegationResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Balance != nil {
			encoded, err := options.Marshal(x.Balance)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if x.Delegation != nil {
			encoded, err := options.Marshal(x.Delegation)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*DelegationResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: DelegationResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: DelegationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Delegation", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Delegation == nil {
					x.Delegation = &Delegation{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Delegation); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Balance == nil {
					x.Balance = &v1beta1.Coin{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Balance); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_RedelegationEntryResponse                    protoreflect.MessageDescriptor
	fd_RedelegationEntryResponse_redelegation_entry protoreflect.FieldDescriptor
	fd_RedelegationEntryResponse_balance            protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_staking_proto_init()
	md_RedelegationEntryResponse = File_cosmos_staking_v1beta1_staking_proto.Messages().ByName("RedelegationEntryResponse")
	fd_RedelegationEntryResponse_redelegation_entry = md_RedelegationEntryResponse.Fields().ByName("redelegation_entry")
	fd_RedelegationEntryResponse_balance = md_RedelegationEntryResponse.Fields().ByName("balance")
}

var _ protoreflect.Message = (*fastReflection_RedelegationEntryResponse)(nil)

type fastReflection_RedelegationEntryResponse RedelegationEntryResponse

func (x *RedelegationEntryResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_RedelegationEntryResponse)(x)
}

func (x *RedelegationEntryResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[17]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_RedelegationEntryResponse_messageType fastReflection_RedelegationEntryResponse_messageType
var _ protoreflect.MessageType = fastReflection_RedelegationEntryResponse_messageType{}

type fastReflection_RedelegationEntryResponse_messageType struct{}

func (x fastReflection_RedelegationEntryResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_RedelegationEntryResponse)(nil)
}
func (x fastReflection_RedelegationEntryResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_RedelegationEntryResponse)
}
func (x fastReflection_RedelegationEntryResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_RedelegationEntryResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_RedelegationEntryResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_RedelegationEntryResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_RedelegationEntryResponse) Type() protoreflect.MessageType {
	return _fastReflection_RedelegationEntryResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_RedelegationEntryResponse) New() protoreflect.Message {
	return new(fastReflection_RedelegationEntryResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_RedelegationEntryResponse) Interface() protoreflect.ProtoMessage {
	return (*RedelegationEntryResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_RedelegationEntryResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.RedelegationEntry != nil {
		value := protoreflect.ValueOfMessage(x.RedelegationEntry.ProtoReflect())
		if !f(fd_RedelegationEntryResponse_redelegation_entry, value) {
			return
		}
	}
	if x.Balance != "" {
		value := protoreflect.ValueOfString(x.Balance)
		if !f(fd_RedelegationEntryResponse_balance, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_RedelegationEntryResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.RedelegationEntryResponse.redelegation_entry":
		return x.RedelegationEntry != nil
	case "cosmos.staking.v1beta1.RedelegationEntryResponse.balance":
		return x.Balance != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.RedelegationEntryResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.RedelegationEntryResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_RedelegationEntryResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.RedelegationEntryResponse.redelegation_entry":
		x.RedelegationEntry = nil
	case "cosmos.staking.v1beta1.RedelegationEntryResponse.balance":
		x.Balance = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.RedelegationEntryResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.RedelegationEntryResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_RedelegationEntryResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.RedelegationEntryResponse.redelegation_entry":
		value := x.RedelegationEntry
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.staking.v1beta1.RedelegationEntryResponse.balance":
		value := x.Balance
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.RedelegationEntryResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.RedelegationEntryResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_RedelegationEntryResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.RedelegationEntryResponse.redelegation_entry":
		x.RedelegationEntry = value.Message().Interface().(*RedelegationEntry)
	case "cosmos.staking.v1beta1.RedelegationEntryResponse.balance":
		x.Balance = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.RedelegationEntryResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.RedelegationEntryResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_RedelegationEntryResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.RedelegationEntryResponse.redelegation_entry":
		if x.RedelegationEntry == nil {
			x.RedelegationEntry = new(RedelegationEntry)
		}
		return protoreflect.ValueOfMessage(x.RedelegationEntry.ProtoReflect())
	case "cosmos.staking.v1beta1.RedelegationEntryResponse.balance":
		panic(fmt.Errorf("field balance of message cosmos.staking.v1beta1.RedelegationEntryResponse is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.RedelegationEntryResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.RedelegationEntryResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_RedelegationEntryResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.RedelegationEntryResponse.redelegation_entry":
		m := new(RedelegationEntry)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.staking.v1beta1.RedelegationEntryResponse.balance":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.RedelegationEntryResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.RedelegationEntryResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_RedelegationEntryResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.RedelegationEntryResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_RedelegationEntryResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_RedelegationEntryResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_RedelegationEntryResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_RedelegationEntryResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*RedelegationEntryResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.RedelegationEntry != nil {
			l = options.Size(x.RedelegationEntry)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Balance)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*RedelegationEntryResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Balance) > 0 {
			i -= len(x.Balance)
			copy(dAtA[i:], x.Balance)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Balance)))
			i--
			dAtA[i] = 0x22
		}
		if x.RedelegationEntry != nil {
			encoded, err := options.Marshal(x.RedelegationEntry)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*RedelegationEntryResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: RedelegationEntryResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: RedelegationEntryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RedelegationEntry", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.RedelegationEntry == nil {
					x.RedelegationEntry = &RedelegationEntry{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.RedelegationEntry); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Balance = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var _ protoreflect.List = (*_RedelegationResponse_2_list)(nil)

type _RedelegationResponse_2_list struct {
	list *[]*RedelegationEntryResponse
}

func (x *_RedelegationResponse_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_RedelegationResponse_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_RedelegationResponse_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*RedelegationEntryResponse)
	(*x.list)[i] = concreteValue
}

func (x *_RedelegationResponse_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*RedelegationEntryResponse)
	*x.list = append(*x.list, concreteValue)
}

func (x *_RedelegationResponse_2_list) AppendMutable() protoreflect.Value {
	v := new(RedelegationEntryResponse)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_RedelegationResponse_2_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_RedelegationResponse_2_list) NewElement() protoreflect.Value {
	v := new(RedelegationEntryResponse)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_RedelegationResponse_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_RedelegationResponse              protoreflect.MessageDescriptor
	fd_RedelegationResponse_redelegation protoreflect.FieldDescriptor
	fd_RedelegationResponse_entries      protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_staking_proto_init()
	md_RedelegationResponse = File_cosmos_staking_v1beta1_staking_proto.Messages().ByName("RedelegationResponse")
	fd_RedelegationResponse_redelegation = md_RedelegationResponse.Fields().ByName("redelegation")
	fd_RedelegationResponse_entries = md_RedelegationResponse.Fields().ByName("entries")
}

var _ protoreflect.Message = (*fastReflection_RedelegationResponse)(nil)

type fastReflection_RedelegationResponse RedelegationResponse

func (x *RedelegationResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_RedelegationResponse)(x)
}

func (x *RedelegationResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[18]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_RedelegationResponse_messageType fastReflection_RedelegationResponse_messageType
var _ protoreflect.MessageType = fastReflection_RedelegationResponse_messageType{}

type fastReflection_RedelegationResponse_messageType struct{}

func (x fastReflection_RedelegationResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_RedelegationResponse)(nil)
}
func (x fastReflection_RedelegationResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_RedelegationResponse)
}
func (x fastReflection_RedelegationResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_RedelegationResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_RedelegationResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_RedelegationResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_RedelegationResponse) Type() protoreflect.MessageType {
	return _fastReflection_RedelegationResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_RedelegationResponse) New() protoreflect.Message {
	return new(fastReflection_RedelegationResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_RedelegationResponse) Interface() protoreflect.ProtoMessage {
	return (*RedelegationResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_RedelegationResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Redelegation != nil {
		value := protoreflect.ValueOfMessage(x.Redelegation.ProtoReflect())
		if !f(fd_RedelegationResponse_redelegation, value) {
			return
		}
	}
	if len(x.Entries) != 0 {
		value := protoreflect.ValueOfList(&_RedelegationResponse_2_list{list: &x.Entries})
		if !f(fd_RedelegationResponse_entries, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_RedelegationResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.RedelegationResponse.redelegation":
		return x.Redelegation != nil
	case "cosmos.staking.v1beta1.RedelegationResponse.entries":
		return len(x.Entries) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.RedelegationResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.RedelegationResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_RedelegationResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.RedelegationResponse.redelegation":
		x.Redelegation = nil
	case "cosmos.staking.v1beta1.RedelegationResponse.entries":
		x.Entries = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.RedelegationResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.RedelegationResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_RedelegationResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.RedelegationResponse.redelegation":
		value := x.Redelegation
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.staking.v1beta1.RedelegationResponse.entries":
		if len(x.Entries) == 0 {
			return protoreflect.ValueOfList(&_RedelegationResponse_2_list{})
		}
		listValue := &_RedelegationResponse_2_list{list: &x.Entries}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.RedelegationResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.RedelegationResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_RedelegationResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.RedelegationResponse.redelegation":
		x.Redelegation = value.Message().Interface().(*Redelegation)
	case "cosmos.staking.v1beta1.RedelegationResponse.entries":
		lv := value.List()
		clv := lv.(*_RedelegationResponse_2_list)
		x.Entries = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.RedelegationResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.RedelegationResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_RedelegationResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.RedelegationResponse.redelegation":
		if x.Redelegation == nil {
			x.Redelegation = new(Redelegation)
		}
		return protoreflect.ValueOfMessage(x.Redelegation.ProtoReflect())
	case "cosmos.staking.v1beta1.RedelegationResponse.entries":
		if x.Entries == nil {
			x.Entries = []*RedelegationEntryResponse{}
		}
		value := &_RedelegationResponse_2_list{list: &x.Entries}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.RedelegationResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.RedelegationResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_RedelegationResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.RedelegationResponse.redelegation":
		m := new(Redelegation)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.staking.v1beta1.RedelegationResponse.entries":
		list := []*RedelegationEntryResponse{}
		return protoreflect.ValueOfList(&_RedelegationResponse_2_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.RedelegationResponse"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.RedelegationResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_RedelegationResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.RedelegationResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_RedelegationResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_RedelegationResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_RedelegationResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_RedelegationResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*RedelegationResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.Redelegation != nil {
			l = options.Size(x.Redelegation)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Entries) > 0 {
			for _, e := range x.Entries {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*RedelegationResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Entries) > 0 {
			for iNdEx := len(x.Entries) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Entries[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x12
			}
		}
		if x.Redelegation != nil {
			encoded, err := options.Marshal(x.Redelegation)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*RedelegationResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: RedelegationResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: RedelegationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Redelegation", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Redelegation == nil {
					x.Redelegation = &Redelegation{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Redelegation); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Entries", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Entries = append(x.Entries, &RedelegationEntryResponse{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Entries[len(x.Entries)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_Pool                   protoreflect.MessageDescriptor
	fd_Pool_not_bonded_tokens protoreflect.FieldDescriptor
	fd_Pool_bonded_tokens     protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_staking_proto_init()
	md_Pool = File_cosmos_staking_v1beta1_staking_proto.Messages().ByName("Pool")
	fd_Pool_not_bonded_tokens = md_Pool.Fields().ByName("not_bonded_tokens")
	fd_Pool_bonded_tokens = md_Pool.Fields().ByName("bonded_tokens")
}

var _ protoreflect.Message = (*fastReflection_Pool)(nil)

type fastReflection_Pool Pool

func (x *Pool) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Pool)(x)
}

func (x *Pool) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[19]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Pool_messageType fastReflection_Pool_messageType
var _ protoreflect.MessageType = fastReflection_Pool_messageType{}

type fastReflection_Pool_messageType struct{}

func (x fastReflection_Pool_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Pool)(nil)
}
func (x fastReflection_Pool_messageType) New() protoreflect.Message {
	return new(fastReflection_Pool)
}
func (x fastReflection_Pool_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Pool
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Pool) Descriptor() protoreflect.MessageDescriptor {
	return md_Pool
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Pool) Type() protoreflect.MessageType {
	return _fastReflection_Pool_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Pool) New() protoreflect.Message {
	return new(fastReflection_Pool)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Pool) Interface() protoreflect.ProtoMessage {
	return (*Pool)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Pool) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.NotBondedTokens != "" {
		value := protoreflect.ValueOfString(x.NotBondedTokens)
		if !f(fd_Pool_not_bonded_tokens, value) {
			return
		}
	}
	if x.BondedTokens != "" {
		value := protoreflect.ValueOfString(x.BondedTokens)
		if !f(fd_Pool_bonded_tokens, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Pool) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.Pool.not_bonded_tokens":
		return x.NotBondedTokens != ""
	case "cosmos.staking.v1beta1.Pool.bonded_tokens":
		return x.BondedTokens != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Pool"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Pool does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Pool) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.Pool.not_bonded_tokens":
		x.NotBondedTokens = ""
	case "cosmos.staking.v1beta1.Pool.bonded_tokens":
		x.BondedTokens = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Pool"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Pool does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Pool) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.Pool.not_bonded_tokens":
		value := x.NotBondedTokens
		return protoreflect.ValueOfString(value)
	case "cosmos.staking.v1beta1.Pool.bonded_tokens":
		value := x.BondedTokens
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Pool"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Pool does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Pool) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.Pool.not_bonded_tokens":
		x.NotBondedTokens = value.Interface().(string)
	case "cosmos.staking.v1beta1.Pool.bonded_tokens":
		x.BondedTokens = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Pool"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Pool does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Pool) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.Pool.not_bonded_tokens":
		panic(fmt.Errorf("field not_bonded_tokens of message cosmos.staking.v1beta1.Pool is not mutable"))
	case "cosmos.staking.v1beta1.Pool.bonded_tokens":
		panic(fmt.Errorf("field bonded_tokens of message cosmos.staking.v1beta1.Pool is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Pool"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Pool does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Pool) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.Pool.not_bonded_tokens":
		return protoreflect.ValueOfString("")
	case "cosmos.staking.v1beta1.Pool.bonded_tokens":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.Pool"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.Pool does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Pool) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.Pool", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Pool) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Pool) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Pool) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Pool) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Pool)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.NotBondedTokens)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.BondedTokens)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Pool)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.BondedTokens) > 0 {
			i -= len(x.BondedTokens)
			copy(dAtA[i:], x.BondedTokens)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BondedTokens)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.NotBondedTokens) > 0 {
			i -= len(x.NotBondedTokens)
			copy(dAtA[i:], x.NotBondedTokens)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.NotBondedTokens)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Pool)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Pool: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Pool: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field NotBondedTokens", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.NotBondedTokens = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BondedTokens", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.BondedTokens = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: cosmos/staking/v1beta1/staking.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// BondStatus is the status of a validator.
type BondStatus int32

const (
	// UNSPECIFIED defines an invalid validator status.
	BondStatus_BOND_STATUS_UNSPECIFIED BondStatus = 0
	// UNBONDED defines a validator that is not bonded.
	BondStatus_BOND_STATUS_UNBONDED BondStatus = 1
	// UNBONDING defines a validator that is unbonding.
	BondStatus_BOND_STATUS_UNBONDING BondStatus = 2
	// BONDED defines a validator that is bonded.
	BondStatus_BOND_STATUS_BONDED BondStatus = 3
)

// Enum value maps for BondStatus.
var (
	BondStatus_name = map[int32]string{
		0: "BOND_STATUS_UNSPECIFIED",
		1: "BOND_STATUS_UNBONDED",
		2: "BOND_STATUS_UNBONDING",
		3: "BOND_STATUS_BONDED",
	}
	BondStatus_value = map[string]int32{
		"BOND_STATUS_UNSPECIFIED": 0,
		"BOND_STATUS_UNBONDED":    1,
		"BOND_STATUS_UNBONDING":   2,
		"BOND_STATUS_BONDED":      3,
	}
)

func (x BondStatus) Enum() *BondStatus {
	p := new(BondStatus)
	*p = x
	return p
}

func (x BondStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BondStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_cosmos_staking_v1beta1_staking_proto_enumTypes[0].Descriptor()
}

func (BondStatus) Type() protoreflect.EnumType {
	return &file_cosmos_staking_v1beta1_staking_proto_enumTypes[0]
}

func (x BondStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BondStatus.Descriptor instead.
func (BondStatus) EnumDescriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_staking_proto_rawDescGZIP(), []int{0}
}

// HistoricalInfo contains header and validator information for a given block.
// It is stored as part of staking module's state, which persists the `n` most
// recent HistoricalInfo
// (`n` is set by the staking module's `historical_entries` parameter).
type HistoricalInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *types.Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Valset []*Validator  `protobuf:"bytes,2,rep,name=valset,proto3" json:"valset,omitempty"`
}

func (x *HistoricalInfo) Reset() {
	*x = HistoricalInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoricalInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoricalInfo) ProtoMessage() {}

// Deprecated: Use HistoricalInfo.ProtoReflect.Descriptor instead.
func (*HistoricalInfo) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_staking_proto_rawDescGZIP(), []int{0}
}

func (x *HistoricalInfo) GetHeader() *types.Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *HistoricalInfo) GetValset() []*Validator {
	if x != nil {
		return x.Valset
	}
	return nil
}

// CommissionRates defines the initial commission rates to be used for creating
// a validator.
type CommissionRates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// rate is the commission rate charged to delegators, as a fraction.
	Rate string `protobuf:"bytes,1,opt,name=rate,proto3" json:"rate,omitempty"`
	// max_rate defines the maximum commission rate which validator can ever charge, as a fraction.
	MaxRate string `protobuf:"bytes,2,opt,name=max_rate,json=maxRate,proto3" json:"max_rate,omitempty"`
	// max_change_rate defines the maximum daily increase of the validator commission, as a fraction.
	MaxChangeRate string `protobuf:"bytes,3,opt,name=max_change_rate,json=maxChangeRate,proto3" json:"max_change_rate,omitempty"`
}

func (x *CommissionRates) Reset() {
	*x = CommissionRates{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommissionRates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommissionRates) ProtoMessage() {}

// Deprecated: Use CommissionRates.ProtoReflect.Descriptor instead.
func (*CommissionRates) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_staking_proto_rawDescGZIP(), []int{1}
}

func (x *CommissionRates) GetRate() string {
	if x != nil {
		return x.Rate
	}
	return ""
}

func (x *CommissionRates) GetMaxRate() string {
	if x != nil {
		return x.MaxRate
	}
	return ""
}

func (x *CommissionRates) GetMaxChangeRate() string {
	if x != nil {
		return x.MaxChangeRate
	}
	return ""
}

// Commission defines commission parameters for a given validator.
type Commission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// commission_rates defines the initial commission rates to be used for creating a validator.
	CommissionRates *CommissionRates `protobuf:"bytes,1,opt,name=commission_rates,json=commissionRates,proto3" json:"commission_rates,omitempty"`
	// update_time is the last time the commission rate was changed.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *Commission) Reset() {
	*x = Commission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Commission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commission) ProtoMessage() {}

// Deprecated: Use Commission.ProtoReflect.Descriptor instead.
func (*Commission) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_staking_proto_rawDescGZIP(), []int{2}
}

func (x *Commission) GetCommissionRates() *CommissionRates {
	if x != nil {
		return x.CommissionRates
	}
	return nil
}

func (x *Commission) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

// Description defines a validator description.
type Description struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// moniker defines a human-readable name for the validator.
	Moniker string `protobuf:"bytes,1,opt,name=moniker,proto3" json:"moniker,omitempty"`
	// identity defines an optional identity signature (ex. UPort or Keybase).
	Identity string `protobuf:"bytes,2,opt,name=identity,proto3" json:"identity,omitempty"`
	// website defines an optional website link.
	Website string `protobuf:"bytes,3,opt,name=website,proto3" json:"website,omitempty"`
	// security_contact defines an optional email for security contact.
	SecurityContact string `protobuf:"bytes,4,opt,name=security_contact,json=securityContact,proto3" json:"security_contact,omitempty"`
	// details define other optional details.
	Details string `protobuf:"bytes,5,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *Description) Reset() {
	*x = Description{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Description) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Description) ProtoMessage() {}

// Deprecated: Use Description.ProtoReflect.Descriptor instead.
func (*Description) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_staking_proto_rawDescGZIP(), []int{3}
}

func (x *Description) GetMoniker() string {
	if x != nil {
		return x.Moniker
	}
	return ""
}

func (x *Description) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *Description) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *Description) GetSecurityContact() string {
	if x != nil {
		return x.SecurityContact
	}
	return ""
}

func (x *Description) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

// Validator defines a validator, together with the total amount of the
// Validator's bond shares and their exchange rate to coins. Slashing results in
// a decrease in the exchange rate, allowing correct calculation of future
// undelegations without iterating over delegators. When coins are delegated to
// this validator, the validator is credited with a delegation whose number of
// bond shares is based on the amount of coins delegated divided by the current
// exchange rate. Voting power can be calculated as total bonded shares
// multiplied by exchange rate.
type Validator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// operator_address defines the address of the validator's operator; bech encoded in JSON.
	OperatorAddress string `protobuf:"bytes,1,opt,name=operator_address,json=operatorAddress,proto3" json:"operator_address,omitempty"`
	// consensus_pubkey is the consensus public key of the validator, as a Protobuf Any.
	ConsensusPubkey *anypb.Any `protobuf:"bytes,2,opt,name=consensus_pubkey,json=consensusPubkey,proto3" json:"consensus_pubkey,omitempty"`
	// jailed defined whether the validator has been jailed from bonded status or not.
	Jailed bool `protobuf:"varint,3,opt,name=jailed,proto3" json:"jailed,omitempty"`
	// status is the validator status (bonded/unbonding/unbonded).
	Status BondStatus `protobuf:"varint,4,opt,name=status,proto3,enum=cosmos.staking.v1beta1.BondStatus" json:"status,omitempty"`
	// tokens define the delegated tokens (incl. self-delegation).
	Tokens string `protobuf:"bytes,5,opt,name=tokens,proto3" json:"tokens,omitempty"`
	// delegator_shares defines total shares issued to a validator's delegators.
	DelegatorShares string `protobuf:"bytes,6,opt,name=delegator_shares,json=delegatorShares,proto3" json:"delegator_shares,omitempty"`
	// description defines the description terms for the validator.
	Description *Description `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	// unbonding_height defines, if unbonding, the height at which this validator has begun unbonding.
	UnbondingHeight int64 `protobuf:"varint,8,opt,name=unbonding_height,json=unbondingHeight,proto3" json:"unbonding_height,omitempty"`
	// unbonding_time defines, if unbonding, the min time for the validator to complete unbonding.
	UnbondingTime *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=unbonding_time,json=unbondingTime,proto3" json:"unbonding_time,omitempty"`
	// commission defines the commission parameters.
	Commission *Commission `protobuf:"bytes,10,opt,name=commission,proto3" json:"commission,omitempty"`
	// min_self_delegation is the validator's self declared minimum self delegation.
	MinSelfDelegation string `protobuf:"bytes,11,opt,name=min_self_delegation,json=minSelfDelegation,proto3" json:"min_self_delegation,omitempty"`
}

func (x *Validator) Reset() {
	*x = Validator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Validator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Validator) ProtoMessage() {}

// Deprecated: Use Validator.ProtoReflect.Descriptor instead.
func (*Validator) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_staking_proto_rawDescGZIP(), []int{4}
}

func (x *Validator) GetOperatorAddress() string {
	if x != nil {
		return x.OperatorAddress
	}
	return ""
}

func (x *Validator) GetConsensusPubkey() *anypb.Any {
	if x != nil {
		return x.ConsensusPubkey
	}
	return nil
}

func (x *Validator) GetJailed() bool {
	if x != nil {
		return x.Jailed
	}
	return false
}

func (x *Validator) GetStatus() BondStatus {
	if x != nil {
		return x.Status
	}
	return BondStatus_BOND_STATUS_UNSPECIFIED
}

func (x *Validator) GetTokens() string {
	if x != nil {
		return x.Tokens
	}
	return ""
}

func (x *Validator) GetDelegatorShares() string {
	if x != nil {
		return x.DelegatorShares
	}
	return ""
}

func (x *Validator) GetDescription() *Description {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *Validator) GetUnbondingHeight() int64 {
	if x != nil {
		return x.UnbondingHeight
	}
	return 0
}

func (x *Validator) GetUnbondingTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UnbondingTime
	}
	return nil
}

func (x *Validator) GetCommission() *Commission {
	if x != nil {
		return x.Commission
	}
	return nil
}

func (x *Validator) GetMinSelfDelegation() string {
	if x != nil {
		return x.MinSelfDelegation
	}
	return ""
}

// ValAddresses defines a repeated set of validator addresses.
type ValAddresses struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addresses []string `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
}

func (x *ValAddresses) Reset() {
	*x = ValAddresses{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValAddresses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValAddresses) ProtoMessage() {}

// Deprecated: Use ValAddresses.ProtoReflect.Descriptor instead.
func (*ValAddresses) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_staking_proto_rawDescGZIP(), []int{5}
}

func (x *ValAddresses) GetAddresses() []string {
	if x != nil {
		return x.Addresses
	}
	return nil
}

// DVPair is struct that just has a delegator-validator pair with no other data.
// It is intended to be used as a marshalable pointer. For example, a DVPair can
// be used to construct the key to getting an UnbondingDelegation from state.
type DVPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DelegatorAddress string `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty"`
	ValidatorAddress string `protobuf:"bytes,2,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
}

func (x *DVPair) Reset() {
	*x = DVPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DVPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DVPair) ProtoMessage() {}

// Deprecated: Use DVPair.ProtoReflect.Descriptor instead.
func (*DVPair) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_staking_proto_rawDescGZIP(), []int{6}
}

func (x *DVPair) GetDelegatorAddress() string {
	if x != nil {
		return x.DelegatorAddress
	}
	return ""
}

func (x *DVPair) GetValidatorAddress() string {
	if x != nil {
		return x.ValidatorAddress
	}
	return ""
}

// DVPairs defines an array of DVPair objects.
type DVPairs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pairs []*DVPair `protobuf:"bytes,1,rep,name=pairs,proto3" json:"pairs,omitempty"`
}

func (x *DVPairs) Reset() {
	*x = DVPairs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DVPairs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DVPairs) ProtoMessage() {}

// Deprecated: Use DVPairs.ProtoReflect.Descriptor instead.
func (*DVPairs) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_staking_proto_rawDescGZIP(), []int{7}
}

func (x *DVPairs) GetPairs() []*DVPair {
	if x != nil {
		return x.Pairs
	}
	return nil
}

// DVVTriplet is struct that just has a delegator-validator-validator triplet
// with no other data. It is intended to be used as a marshalable pointer. For
// example, a DVVTriplet can be used to construct the key to getting a
// Redelegation from state.
type DVVTriplet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DelegatorAddress    string `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty"`
	ValidatorSrcAddress string `protobuf:"bytes,2,opt,name=validator_src_address,json=validatorSrcAddress,proto3" json:"validator_src_address,omitempty"`
	ValidatorDstAddress string `protobuf:"bytes,3,opt,name=validator_dst_address,json=validatorDstAddress,proto3" json:"validator_dst_address,omitempty"`
}

func (x *DVVTriplet) Reset() {
	*x = DVVTriplet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DVVTriplet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DVVTriplet) ProtoMessage() {}

// Deprecated: Use DVVTriplet.ProtoReflect.Descriptor instead.
func (*DVVTriplet) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_staking_proto_rawDescGZIP(), []int{8}
}

func (x *DVVTriplet) GetDelegatorAddress() string {
	if x != nil {
		return x.DelegatorAddress
	}
	return ""
}

func (x *DVVTriplet) GetValidatorSrcAddress() string {
	if x != nil {
		return x.ValidatorSrcAddress
	}
	return ""
}

func (x *DVVTriplet) GetValidatorDstAddress() string {
	if x != nil {
		return x.ValidatorDstAddress
	}
	return ""
}

// DVVTriplets defines an array of DVVTriplet objects.
type DVVTriplets struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Triplets []*DVVTriplet `protobuf:"bytes,1,rep,name=triplets,proto3" json:"triplets,omitempty"`
}

func (x *DVVTriplets) Reset() {
	*x = DVVTriplets{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DVVTriplets) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DVVTriplets) ProtoMessage() {}

// Deprecated: Use DVVTriplets.ProtoReflect.Descriptor instead.
func (*DVVTriplets) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_staking_proto_rawDescGZIP(), []int{9}
}

func (x *DVVTriplets) GetTriplets() []*DVVTriplet {
	if x != nil {
		return x.Triplets
	}
	return nil
}

// Delegation represents the bond with tokens held by an account. It is
// owned by one delegator, and is associated with the voting power of one
// validator.
type Delegation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// delegator_address is the bech32-encoded address of the delegator.
	DelegatorAddress string `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty"`
	// validator_address is the bech32-encoded address of the validator.
	ValidatorAddress string `protobuf:"bytes,2,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	// shares define the delegation shares received.
	Shares string `protobuf:"bytes,3,opt,name=shares,proto3" json:"shares,omitempty"`
}

func (x *Delegation) Reset() {
	*x = Delegation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Delegation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Delegation) ProtoMessage() {}

// Deprecated: Use Delegation.ProtoReflect.Descriptor instead.
func (*Delegation) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_staking_proto_rawDescGZIP(), []int{10}
}

func (x *Delegation) GetDelegatorAddress() string {
	if x != nil {
		return x.DelegatorAddress
	}
	return ""
}

func (x *Delegation) GetValidatorAddress() string {
	if x != nil {
		return x.ValidatorAddress
	}
	return ""
}

func (x *Delegation) GetShares() string {
	if x != nil {
		return x.Shares
	}
	return ""
}

// UnbondingDelegation stores all of a single delegator's unbonding bonds
// for a single validator in an time-ordered list.
type UnbondingDelegation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// delegator_address is the bech32-encoded address of the delegator.
	DelegatorAddress string `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty"`
	// validator_address is the bech32-encoded address of the validator.
	ValidatorAddress string `protobuf:"bytes,2,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	// entries are the unbonding delegation entries.
	Entries []*UnbondingDelegationEntry `protobuf:"bytes,3,rep,name=entries,proto3" json:"entries,omitempty"` // unbonding delegation entries
}

func (x *UnbondingDelegation) Reset() {
	*x = UnbondingDelegation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnbondingDelegation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnbondingDelegation) ProtoMessage() {}

// Deprecated: Use UnbondingDelegation.ProtoReflect.Descriptor instead.
func (*UnbondingDelegation) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_staking_proto_rawDescGZIP(), []int{11}
}

func (x *UnbondingDelegation) GetDelegatorAddress() string {
	if x != nil {
		return x.DelegatorAddress
	}
	return ""
}

func (x *UnbondingDelegation) GetValidatorAddress() string {
	if x != nil {
		return x.ValidatorAddress
	}
	return ""
}

func (x *UnbondingDelegation) GetEntries() []*UnbondingDelegationEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

// UnbondingDelegationEntry defines an unbonding object with relevant metadata.
type UnbondingDelegationEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// creation_height is the height which the unbonding took place.
	CreationHeight int64 `protobuf:"varint,1,opt,name=creation_height,json=creationHeight,proto3" json:"creation_height,omitempty"`
	// completion_time is the unix time for unbonding completion.
	CompletionTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=completion_time,json=completionTime,proto3" json:"completion_time,omitempty"`
	// initial_balance defines the tokens initially scheduled to receive at completion.
	InitialBalance string `protobuf:"bytes,3,opt,name=initial_balance,json=initialBalance,proto3" json:"initial_balance,omitempty"`
	// balance defines the tokens to receive at completion.
	Balance string `protobuf:"bytes,4,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *UnbondingDelegationEntry) Reset() {
	*x = UnbondingDelegationEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnbondingDelegationEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnbondingDelegationEntry) ProtoMessage() {}

// Deprecated: Use UnbondingDelegationEntry.ProtoReflect.Descriptor instead.
func (*UnbondingDelegationEntry) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_staking_proto_rawDescGZIP(), []int{12}
}

func (x *UnbondingDelegationEntry) GetCreationHeight() int64 {
	if x != nil {
		return x.CreationHeight
	}
	return 0
}

func (x *UnbondingDelegationEntry) GetCompletionTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CompletionTime
	}
	return nil
}

func (x *UnbondingDelegationEntry) GetInitialBalance() string {
	if x != nil {
		return x.InitialBalance
	}
	return ""
}

func (x *UnbondingDelegationEntry) GetBalance() string {
	if x != nil {
		return x.Balance
	}
	return ""
}

// RedelegationEntry defines a redelegation object with relevant metadata.
type RedelegationEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// creation_height  defines the height which the redelegation took place.
	CreationHeight int64 `protobuf:"varint,1,opt,name=creation_height,json=creationHeight,proto3" json:"creation_height,omitempty"`
	// completion_time defines the unix time for redelegation completion.
	CompletionTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=completion_time,json=completionTime,proto3" json:"completion_time,omitempty"`
	// initial_balance defines the initial balance when redelegation started.
	InitialBalance string `protobuf:"bytes,3,opt,name=initial_balance,json=initialBalance,proto3" json:"initial_balance,omitempty"`
	// shares_dst is the amount of destination-validator shares created by redelegation.
	SharesDst string `protobuf:"bytes,4,opt,name=shares_dst,json=sharesDst,proto3" json:"shares_dst,omitempty"`
}

func (x *RedelegationEntry) Reset() {
	*x = RedelegationEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedelegationEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedelegationEntry) ProtoMessage() {}

// Deprecated: Use RedelegationEntry.ProtoReflect.Descriptor instead.
func (*RedelegationEntry) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_staking_proto_rawDescGZIP(), []int{13}
}

func (x *RedelegationEntry) GetCreationHeight() int64 {
	if x != nil {
		return x.CreationHeight
	}
	return 0
}

func (x *RedelegationEntry) GetCompletionTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CompletionTime
	}
	return nil
}

func (x *RedelegationEntry) GetInitialBalance() string {
	if x != nil {
		return x.InitialBalance
	}
	return ""
}

func (x *RedelegationEntry) GetSharesDst() string {
	if x != nil {
		return x.SharesDst
	}
	return ""
}

// Redelegation contains the list of a particular delegator's redelegating bonds
// from a particular source validator to a particular destination validator.
type Redelegation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// delegator_address is the bech32-encoded address of the delegator.
	DelegatorAddress string `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty"`
	// validator_src_address is the validator redelegation source operator address.
	ValidatorSrcAddress string `protobuf:"bytes,2,opt,name=validator_src_address,json=validatorSrcAddress,proto3" json:"validator_src_address,omitempty"`
	// validator_dst_address is the validator redelegation destination operator address.
	ValidatorDstAddress string `protobuf:"bytes,3,opt,name=validator_dst_address,json=validatorDstAddress,proto3" json:"validator_dst_address,omitempty"`
	// entries are the redelegation entries.
	Entries []*RedelegationEntry `protobuf:"bytes,4,rep,name=entries,proto3" json:"entries,omitempty"` // redelegation entries
}

func (x *Redelegation) Reset() {
	*x = Redelegation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Redelegation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Redelegation) ProtoMessage() {}

// Deprecated: Use Redelegation.ProtoReflect.Descriptor instead.
func (*Redelegation) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_staking_proto_rawDescGZIP(), []int{14}
}

func (x *Redelegation) GetDelegatorAddress() string {
	if x != nil {
		return x.DelegatorAddress
	}
	return ""
}

func (x *Redelegation) GetValidatorSrcAddress() string {
	if x != nil {
		return x.ValidatorSrcAddress
	}
	return ""
}

func (x *Redelegation) GetValidatorDstAddress() string {
	if x != nil {
		return x.ValidatorDstAddress
	}
	return ""
}

func (x *Redelegation) GetEntries() []*RedelegationEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

// Params defines the parameters for the staking module.
type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// unbonding_time is the time duration of unbonding.
	UnbondingTime *durationpb.Duration `protobuf:"bytes,1,opt,name=unbonding_time,json=unbondingTime,proto3" json:"unbonding_time,omitempty"`
	// max_validators is the maximum number of validators.
	MaxValidators uint32 `protobuf:"varint,2,opt,name=max_validators,json=maxValidators,proto3" json:"max_validators,omitempty"`
	// max_entries is the max entries for either unbonding delegation or redelegation (per pair/trio).
	MaxEntries uint32 `protobuf:"varint,3,opt,name=max_entries,json=maxEntries,proto3" json:"max_entries,omitempty"`
	// historical_entries is the number of historical entries to persist.
	HistoricalEntries uint32 `protobuf:"varint,4,opt,name=historical_entries,json=historicalEntries,proto3" json:"historical_entries,omitempty"`
	// bond_denom defines the bondable coin denomination.
	BondDenom string `protobuf:"bytes,5,opt,name=bond_denom,json=bondDenom,proto3" json:"bond_denom,omitempty"`
	// min_commission_rate is the chain-wide minimum commission rate that a validator can charge their delegators
	MinCommissionRate string `protobuf:"bytes,6,opt,name=min_commission_rate,json=minCommissionRate,proto3" json:"min_commission_rate,omitempty"`
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

// Deprecated: Use Params.ProtoReflect.Descriptor instead.
func (*Params) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_staking_proto_rawDescGZIP(), []int{15}
}

func (x *Params) GetUnbondingTime() *durationpb.Duration {
	if x != nil {
		return x.UnbondingTime
	}
	return nil
}

func (x *Params) GetMaxValidators() uint32 {
	if x != nil {
		return x.MaxValidators
	}
	return 0
}

func (x *Params) GetMaxEntries() uint32 {
	if x != nil {
		return x.MaxEntries
	}
	return 0
}

func (x *Params) GetHistoricalEntries() uint32 {
	if x != nil {
		return x.HistoricalEntries
	}
	return 0
}

func (x *Params) GetBondDenom() string {
	if x != nil {
		return x.BondDenom
	}
	return ""
}

func (x *Params) GetMinCommissionRate() string {
	if x != nil {
		return x.MinCommissionRate
	}
	return ""
}

// DelegationResponse is equivalent to Delegation except that it contains a
// balance in addition to shares which is more suitable for client responses.
type DelegationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Delegation *Delegation   `protobuf:"bytes,1,opt,name=delegation,proto3" json:"delegation,omitempty"`
	Balance    *v1beta1.Coin `protobuf:"bytes,2,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *DelegationResponse) Reset() {
	*x = DelegationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelegationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelegationResponse) ProtoMessage() {}

// Deprecated: Use DelegationResponse.ProtoReflect.Descriptor instead.
func (*DelegationResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_staking_proto_rawDescGZIP(), []int{16}
}

func (x *DelegationResponse) GetDelegation() *Delegation {
	if x != nil {
		return x.Delegation
	}
	return nil
}

func (x *DelegationResponse) GetBalance() *v1beta1.Coin {
	if x != nil {
		return x.Balance
	}
	return nil
}

// RedelegationEntryResponse is equivalent to a RedelegationEntry except that it
// contains a balance in addition to shares which is more suitable for client
// responses.
type RedelegationEntryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RedelegationEntry *RedelegationEntry `protobuf:"bytes,1,opt,name=redelegation_entry,json=redelegationEntry,proto3" json:"redelegation_entry,omitempty"`
	Balance           string             `protobuf:"bytes,4,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *RedelegationEntryResponse) Reset() {
	*x = RedelegationEntryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[17]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedelegationEntryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedelegationEntryResponse) ProtoMessage() {}

// Deprecated: Use RedelegationEntryResponse.ProtoReflect.Descriptor instead.
func (*RedelegationEntryResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_staking_proto_rawDescGZIP(), []int{17}
}

func (x *RedelegationEntryResponse) GetRedelegationEntry() *RedelegationEntry {
	if x != nil {
		return x.RedelegationEntry
	}
	return nil
}

func (x *RedelegationEntryResponse) GetBalance() string {
	if x != nil {
		return x.Balance
	}
	return ""
}

// RedelegationResponse is equivalent to a Redelegation except that its entries
// contain a balance in addition to shares which is more suitable for client
// responses.
type RedelegationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Redelegation *Redelegation                `protobuf:"bytes,1,opt,name=redelegation,proto3" json:"redelegation,omitempty"`
	Entries      []*RedelegationEntryResponse `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *RedelegationResponse) Reset() {
	*x = RedelegationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[18]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedelegationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedelegationResponse) ProtoMessage() {}

// Deprecated: Use RedelegationResponse.ProtoReflect.Descriptor instead.
func (*RedelegationResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_staking_proto_rawDescGZIP(), []int{18}
}

func (x *RedelegationResponse) GetRedelegation() *Redelegation {
	if x != nil {
		return x.Redelegation
	}
	return nil
}

func (x *RedelegationResponse) GetEntries() []*RedelegationEntryResponse {
	if x != nil {
		return x.Entries
	}
	return nil
}

// Pool is used for tracking bonded and not-bonded token supply of the bond
// denomination.
type Pool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NotBondedTokens string `protobuf:"bytes,1,opt,name=not_bonded_tokens,json=notBondedTokens,proto3" json:"not_bonded_tokens,omitempty"`
	BondedTokens    string `protobuf:"bytes,2,opt,name=bonded_tokens,json=bondedTokens,proto3" json:"bonded_tokens,omitempty"`
}

func (x *Pool) Reset() {
	*x = Pool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_staking_proto_msgTypes[19]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pool) ProtoMessage() {}

// Deprecated: Use Pool.ProtoReflect.Descriptor instead.
func (*Pool) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_staking_proto_rawDescGZIP(), []int{19}
}

func (x *Pool) GetNotBondedTokens() string {
	if x != nil {
		return x.NotBondedTokens
	}
	return ""
}

func (x *Pool) GetBondedTokens() string {
	if x != nil {
		return x.BondedTokens
	}
	return ""
}

var File_cosmos_staking_v1beta1_staking_proto protoreflect.FileDescriptor

var file_cosmos_staking_v1beta1_staking_proto_rawDesc = []byte{
	0x0a, 0x24, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73,
	0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x14,
	0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x74, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x0e, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x36, 0x0a, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x73, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74,
	0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x06, 0x76,
	0x61, 0x6c, 0x73, 0x65, 0x74, 0x22, 0xac, 0x02, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x73, 0x12, 0x50, 0x0a, 0x04, 0x72, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3c, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f,
	0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x63, 0xd2, 0xb4, 0x2d, 0x0a, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x44, 0x65, 0x63, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x12, 0x57, 0x0a, 0x08, 0x6d,
	0x61, 0x78, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3c, 0xc8,
	0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x63, 0xd2, 0xb4, 0x2d,
	0x0a, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x44, 0x65, 0x63, 0x52, 0x07, 0x6d, 0x61, 0x78,
	0x52, 0x61, 0x74, 0x65, 0x12, 0x64, 0x0a, 0x0f, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3c, 0xc8,
	0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x63, 0xd2, 0xb4, 0x2d,
	0x0a, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x44, 0x65, 0x63, 0x52, 0x0d, 0x6d, 0x61, 0x78,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x3a, 0x08, 0x98, 0xa0, 0x1f, 0x00,
	0xe8, 0xa0, 0x1f, 0x01, 0x22, 0xbb, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x5c, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x61, 0x74, 0x65, 0x73, 0x42, 0x08, 0xc8, 0xde, 0x1f, 0x00, 0xd0, 0xde, 0x1f, 0x01,
	0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65,
	0x73, 0x12, 0x45, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x08, 0xc8, 0xde, 0x1f, 0x00, 0x90, 0xdf, 0x1f, 0x01, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x3a, 0x08, 0x98, 0xa0, 0x1f, 0x00, 0xe8, 0xa0,
	0x1f, 0x01, 0x22, 0xac, 0x01, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x6f, 0x6e, 0x69, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x6f, 0x6e, 0x69, 0x6b, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x73,
	0x69, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69,
	0x74, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x3a, 0x08, 0x98, 0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f,
	0x01, 0x22, 0xc9, 0x06, 0x0a, 0x09, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x43, 0x0a, 0x10, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x52, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x59, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75,
	0x73, 0x5f, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x42, 0x18, 0xca, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x52, 0x0f,
	0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x6a, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x6a, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x3a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x42, 0x6f, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x54, 0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x3c, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x26, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x49, 0x6e, 0x74, 0xd2, 0xb4, 0x2d, 0x0a, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x49, 0x6e,
	0x74, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x67, 0x0a, 0x10, 0x64, 0x65, 0x6c,
	0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x3c, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x26, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x44, 0x65, 0x63, 0xd2, 0xb4, 0x2d, 0x0a, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x44, 0x65,
	0x63, 0x52, 0x0f, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x73, 0x12, 0x4b, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x04, 0xc8, 0xde,
	0x1f, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x29, 0x0a, 0x10, 0x75, 0x6e, 0x62, 0x6f, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x75, 0x6e, 0x62, 0x6f, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x4b, 0x0a, 0x0e, 0x75, 0x6e,
	0x62, 0x6f, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08,
	0xc8, 0xde, 0x1f, 0x00, 0x90, 0xdf, 0x1f, 0x01, 0x52, 0x0d, 0x75, 0x6e, 0x62, 0x6f, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x48, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42,
	0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x6c, 0x0a, 0x13, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x6c, 0x66, 0x5f, 0x64, 0x65,
	0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3c,
	0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0xd2, 0xb4,
	0x2d, 0x0a, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x11, 0x6d, 0x69,
	0x6e, 0x53, 0x65, 0x6c, 0x66, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a,
	0x0c, 0x88, 0xa0, 0x1f, 0x00, 0x98, 0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x00, 0x22, 0x50, 0x0a,
	0x0c, 0x56, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x36, 0x0a,
	0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x65, 0x73, 0x3a, 0x08, 0x98, 0xa0, 0x1f, 0x00, 0x80, 0xdc, 0x20, 0x01, 0x22,
	0xa4, 0x01, 0x0a, 0x06, 0x44, 0x56, 0x50, 0x61, 0x69, 0x72, 0x12, 0x45, 0x0a, 0x11, 0x64, 0x65,
	0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52,
	0x10, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x45, 0x0a, 0x11, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4,
	0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x10, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3a, 0x0c, 0x88, 0xa0, 0x1f, 0x00, 0x98, 0xa0,
	0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x00, 0x22, 0x45, 0x0a, 0x07, 0x44, 0x56, 0x50, 0x61, 0x69, 0x72,
	0x73, 0x12, 0x3a, 0x0a, 0x05, 0x70, 0x61, 0x69, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x56, 0x50, 0x61, 0x69, 0x72,
	0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x05, 0x70, 0x61, 0x69, 0x72, 0x73, 0x22, 0xfd, 0x01,
	0x0a, 0x0a, 0x44, 0x56, 0x56, 0x54, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x74, 0x12, 0x45, 0x0a, 0x11,
	0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x52, 0x10, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x4c, 0x0a, 0x15, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x5f, 0x73, 0x72, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x13, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x72, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x4c, 0x0a, 0x15, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x64,
	0x73, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x13, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x44, 0x73, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x3a,
	0x0c, 0x88, 0xa0, 0x1f, 0x00, 0x98, 0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x00, 0x22, 0x53, 0x0a,
	0x0b, 0x44, 0x56, 0x56, 0x54, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x74, 0x73, 0x12, 0x44, 0x0a, 0x08,
	0x74, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x56, 0x56, 0x54, 0x72, 0x69, 0x70, 0x6c,
	0x65, 0x74, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x08, 0x74, 0x72, 0x69, 0x70, 0x6c, 0x65,
	0x74, 0x73, 0x22, 0xfe, 0x01, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x45, 0x0a, 0x11, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4,
	0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x10, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f,
	0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x45, 0x0a, 0x11, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x10, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x54, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x3c, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x63, 0xd2,
	0xb4, 0x2d, 0x0a, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x44, 0x65, 0x63, 0x52, 0x06, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x73, 0x3a, 0x0c, 0x88, 0xa0, 0x1f, 0x00, 0x98, 0xa0, 0x1f, 0x00, 0xe8,
	0xa0, 0x1f, 0x00, 0x22, 0x83, 0x02, 0x0a, 0x13, 0x55, 0x6e, 0x62, 0x6f, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x11, 0x64,
	0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x52, 0x10, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x45, 0x0a, 0x11, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2,
	0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x10, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x50, 0x0a, 0x07, 0x65, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x55, 0x6e, 0x62, 0x6f, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x6c,
	0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x04, 0xc8, 0xde,
	0x1f, 0x00, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x3a, 0x0c, 0x88, 0xa0, 0x1f,
	0x00, 0x98, 0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x00, 0x22, 0xdb, 0x02, 0x0a, 0x18, 0x55, 0x6e,
	0x62, 0x6f, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x4d, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xc8, 0xde, 0x1f, 0x00, 0x90, 0xdf, 0x1f, 0x01, 0x52, 0x0e,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x65,
	0x0a, 0x0f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3c, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f,
	0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0xd2, 0xb4, 0x2d, 0x0a, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x0e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3c, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x26,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0xd2, 0xb4, 0x2d, 0x0a, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x49, 0x6e, 0x74, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x3a, 0x08, 0x98,
	0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x01, 0x22, 0xd9, 0x02, 0x0a, 0x11, 0x52, 0x65, 0x64, 0x65,
	0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x27, 0x0a,
	0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x4d, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xc8, 0xde, 0x1f,
	0x00, 0x90, 0xdf, 0x1f, 0x01, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x65, 0x0a, 0x0f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c,
	0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3c,
	0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0xd2, 0xb4,
	0x2d, 0x0a, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x0e, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x61, 0x6c, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x0a,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x5f, 0x64, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x3c, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x63,
	0xd2, 0xb4, 0x2d, 0x0a, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x44, 0x65, 0x63, 0x52, 0x09,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x44, 0x73, 0x74, 0x3a, 0x08, 0x98, 0xa0, 0x1f, 0x00, 0xe8,
	0xa0, 0x1f, 0x01, 0x22, 0xca, 0x02, 0x0a, 0x0c, 0x52, 0x65, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x11, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f,
	0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x10, 0x64, 0x65, 0x6c, 0x65, 0x67,
	0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x4c, 0x0a, 0x15, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x72, 0x63, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x52, 0x13, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x53,
	0x72, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x4c, 0x0a, 0x15, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x64, 0x73, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x52, 0x13, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x44, 0x73, 0x74,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x49, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x52, 0x65, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x3a, 0x0c, 0x88, 0xa0, 0x1f, 0x00, 0x98, 0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x00,
	0x22, 0xf2, 0x02, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x4a, 0x0a, 0x0e, 0x75,
	0x6e, 0x62, 0x6f, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08,
	0xc8, 0xde, 0x1f, 0x00, 0x98, 0xdf, 0x1f, 0x01, 0x52, 0x0d, 0x75, 0x6e, 0x62, 0x6f, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x5f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0d, 0x6d, 0x61, 0x78, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12,
	0x2d, 0x0a, 0x12, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x65, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x68, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x62, 0x6f, 0x6e, 0x64, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x62, 0x6f, 0x6e, 0x64, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x7c, 0x0a,
	0x13, 0x6d, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x72, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x4c, 0xc8, 0xde, 0x1f, 0x00,
	0xda, 0xde, 0x1f, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x63, 0xf2, 0xde, 0x1f, 0x1a, 0x79, 0x61,
	0x6d, 0x6c, 0x3a, 0x22, 0x6d, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x22, 0x52, 0x11, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x3a, 0x08, 0x98, 0xa0, 0x1f,
	0x00, 0xe8, 0xa0, 0x1f, 0x01, 0x22, 0xa3, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0a,
	0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x65,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f,
	0x69, 0x6e, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x3a, 0x08, 0x98, 0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x00, 0x22, 0xd9, 0x01, 0x0a, 0x19,
	0x52, 0x65, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x12, 0x72, 0x65, 0x64,
	0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73,
	0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52,
	0x65, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x11, 0x72, 0x65, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x56, 0x0a, 0x07, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3c, 0xc8, 0xde, 0x1f, 0x00,
	0xda, 0xde, 0x1f, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0xd2, 0xb4, 0x2d, 0x0a, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x3a, 0x04, 0xe8, 0xa0, 0x1f, 0x01, 0x22, 0xbf, 0x01, 0x0a, 0x14, 0x52, 0x65, 0x64, 0x65,
	0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4e, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x52, 0x65, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x04, 0xc8, 0xde,
	0x1f, 0x00, 0x52, 0x0c, 0x72, 0x65, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x51, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x31, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x64, 0x65, 0x6c,
	0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x3a, 0x04, 0xe8, 0xa0, 0x1f, 0x00, 0x22, 0x83, 0x02, 0x0a, 0x04, 0x50, 0x6f,
	0x6f, 0x6c, 0x12, 0x7d, 0x0a, 0x11, 0x6e, 0x6f, 0x74, 0x5f, 0x62, 0x6f, 0x6e, 0x64, 0x65, 0x64,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x51, 0xc8,
	0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0xea, 0xde, 0x1f,
	0x11, 0x6e, 0x6f, 0x74, 0x5f, 0x62, 0x6f, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0xd2, 0xb4, 0x2d, 0x0a, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x49, 0x6e, 0x74,
	0x52, 0x0f, 0x6e, 0x6f, 0x74, 0x42, 0x6f, 0x6e, 0x64, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x12, 0x72, 0x0a, 0x0d, 0x62, 0x6f, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x4d, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde,
	0x1f, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0xea, 0xde, 0x1f, 0x0d, 0x62, 0x6f, 0x6e, 0x64,
	0x65, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0xd2, 0xb4, 0x2d, 0x0a, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x0c, 0x62, 0x6f, 0x6e, 0x64, 0x65, 0x64, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x3a, 0x08, 0xe8, 0xa0, 0x1f, 0x01, 0xf0, 0xa0, 0x1f, 0x01, 0x2a,
	0xb6, 0x01, 0x0a, 0x0a, 0x42, 0x6f, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c,
	0x0a, 0x17, 0x42, 0x4f, 0x4e, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x1a, 0x0f, 0x8a, 0x9d, 0x20,
	0x0b, 0x55, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x26, 0x0a, 0x14,
	0x42, 0x4f, 0x4e, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x42, 0x4f,
	0x4e, 0x44, 0x45, 0x44, 0x10, 0x01, 0x1a, 0x0c, 0x8a, 0x9d, 0x20, 0x08, 0x55, 0x6e, 0x62, 0x6f,
	0x6e, 0x64, 0x65, 0x64, 0x12, 0x28, 0x0a, 0x15, 0x42, 0x4f, 0x4e, 0x44, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x42, 0x4f, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x1a,
	0x0d, 0x8a, 0x9d, 0x20, 0x09, 0x55, 0x6e, 0x62, 0x6f, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x22,
	0x0a, 0x12, 0x42, 0x4f, 0x4e, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x42, 0x4f,
	0x4e, 0x44, 0x45, 0x44, 0x10, 0x03, 0x1a, 0x0a, 0x8a, 0x9d, 0x20, 0x06, 0x42, 0x6f, 0x6e, 0x64,
	0x65, 0x64, 0x1a, 0x04, 0x88, 0xa3, 0x1e, 0x00, 0x42, 0xec, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0c, 0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x3b, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2,
	0x02, 0x03, 0x43, 0x53, 0x58, 0xaa, 0x02, 0x16, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x53,
	0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02,
	0x16, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5c,
	0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x22, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x5c, 0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x18, 0x43,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_staking_v1beta1_staking_proto_rawDescOnce sync.Once
	file_cosmos_staking_v1beta1_staking_proto_rawDescData = file_cosmos_staking_v1beta1_staking_proto_rawDesc
)

func file_cosmos_staking_v1beta1_staking_proto_rawDescGZIP() []byte {
	file_cosmos_staking_v1beta1_staking_proto_rawDescOnce.Do(func() {
		file_cosmos_staking_v1beta1_staking_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_staking_v1beta1_staking_proto_rawDescData)
	})
	return file_cosmos_staking_v1beta1_staking_proto_rawDescData
}

var file_cosmos_staking_v1beta1_staking_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cosmos_staking_v1beta1_staking_proto_msgTypes = make([]protoimpl.MessageInfo, 20)
var file_cosmos_staking_v1beta1_staking_proto_goTypes = []interface{}{
	(BondStatus)(0),                   // 0: cosmos.staking.v1beta1.BondStatus
	(*HistoricalInfo)(nil),            // 1: cosmos.staking.v1beta1.HistoricalInfo
	(*CommissionRates)(nil),           // 2: cosmos.staking.v1beta1.CommissionRates
	(*Commission)(nil),                // 3: cosmos.staking.v1beta1.Commission
	(*Description)(nil),               // 4: cosmos.staking.v1beta1.Description
	(*Validator)(nil),                 // 5: cosmos.staking.v1beta1.Validator
	(*ValAddresses)(nil),              // 6: cosmos.staking.v1beta1.ValAddresses
	(*DVPair)(nil),                    // 7: cosmos.staking.v1beta1.DVPair
	(*DVPairs)(nil),                   // 8: cosmos.staking.v1beta1.DVPairs
	(*DVVTriplet)(nil),                // 9: cosmos.staking.v1beta1.DVVTriplet
	(*DVVTriplets)(nil),               // 10: cosmos.staking.v1beta1.DVVTriplets
	(*Delegation)(nil),                // 11: cosmos.staking.v1beta1.Delegation
	(*UnbondingDelegation)(nil),       // 12: cosmos.staking.v1beta1.UnbondingDelegation
	(*UnbondingDelegationEntry)(nil),  // 13: cosmos.staking.v1beta1.UnbondingDelegationEntry
	(*RedelegationEntry)(nil),         // 14: cosmos.staking.v1beta1.RedelegationEntry
	(*Redelegation)(nil),              // 15: cosmos.staking.v1beta1.Redelegation
	(*Params)(nil),                    // 16: cosmos.staking.v1beta1.Params
	(*DelegationResponse)(nil),        // 17: cosmos.staking.v1beta1.DelegationResponse
	(*RedelegationEntryResponse)(nil), // 18: cosmos.staking.v1beta1.RedelegationEntryResponse
	(*RedelegationResponse)(nil),      // 19: cosmos.staking.v1beta1.RedelegationResponse
	(*Pool)(nil),                      // 20: cosmos.staking.v1beta1.Pool
	(*types.Header)(nil),              // 21: tendermint.types.Header
	(*timestamppb.Timestamp)(nil),     // 22: google.protobuf.Timestamp
	(*anypb.Any)(nil),                 // 23: google.protobuf.Any
	(*durationpb.Duration)(nil),       // 24: google.protobuf.Duration
	(*v1beta1.Coin)(nil),              // 25: cosmos.base.v1beta1.Coin
}
var file_cosmos_staking_v1beta1_staking_proto_depIdxs = []int32{
	21, // 0: cosmos.staking.v1beta1.HistoricalInfo.header:type_name -> tendermint.types.Header
	5,  // 1: cosmos.staking.v1beta1.HistoricalInfo.valset:type_name -> cosmos.staking.v1beta1.Validator
	2,  // 2: cosmos.staking.v1beta1.Commission.commission_rates:type_name -> cosmos.staking.v1beta1.CommissionRates
	22, // 3: cosmos.staking.v1beta1.Commission.update_time:type_name -> google.protobuf.Timestamp
	23, // 4: cosmos.staking.v1beta1.Validator.consensus_pubkey:type_name -> google.protobuf.Any
	0,  // 5: cosmos.staking.v1beta1.Validator.status:type_name -> cosmos.staking.v1beta1.BondStatus
	4,  // 6: cosmos.staking.v1beta1.Validator.description:type_name -> cosmos.staking.v1beta1.Description
	22, // 7: cosmos.staking.v1beta1.Validator.unbonding_time:type_name -> google.protobuf.Timestamp
	3,  // 8: cosmos.staking.v1beta1.Validator.commission:type_name -> cosmos.staking.v1beta1.Commission
	7,  // 9: cosmos.staking.v1beta1.DVPairs.pairs:type_name -> cosmos.staking.v1beta1.DVPair
	9,  // 10: cosmos.staking.v1beta1.DVVTriplets.triplets:type_name -> cosmos.staking.v1beta1.DVVTriplet
	13, // 11: cosmos.staking.v1beta1.UnbondingDelegation.entries:type_name -> cosmos.staking.v1beta1.UnbondingDelegationEntry
	22, // 12: cosmos.staking.v1beta1.UnbondingDelegationEntry.completion_time:type_name -> google.protobuf.Timestamp
	22, // 13: cosmos.staking.v1beta1.RedelegationEntry.completion_time:type_name -> google.protobuf.Timestamp
	14, // 14: cosmos.staking.v1beta1.Redelegation.entries:type_name -> cosmos.staking.v1beta1.RedelegationEntry
	24, // 15: cosmos.staking.v1beta1.Params.unbonding_time:type_name -> google.protobuf.Duration
	11, // 16: cosmos.staking.v1beta1.DelegationResponse.delegation:type_name -> cosmos.staking.v1beta1.Delegation
	25, // 17: cosmos.staking.v1beta1.DelegationResponse.balance:type_name -> cosmos.base.v1beta1.Coin
	14, // 18: cosmos.staking.v1beta1.RedelegationEntryResponse.redelegation_entry:type_name -> cosmos.staking.v1beta1.RedelegationEntry
	15, // 19: cosmos.staking.v1beta1.RedelegationResponse.redelegation:type_name -> cosmos.staking.v1beta1.Redelegation
	18, // 20: cosmos.staking.v1beta1.RedelegationResponse.entries:type_name -> cosmos.staking.v1beta1.RedelegationEntryResponse
	21, // [21:21] is the sub-list for method output_type
	21, // [21:21] is the sub-list for method input_type
	21, // [21:21] is the sub-list for extension type_name
	21, // [21:21] is the sub-list for extension extendee
	0,  // [0:21] is the sub-list for field type_name
}

func init() { file_cosmos_staking_v1beta1_staking_proto_init() }
func file_cosmos_staking_v1beta1_staking_proto_init() {
	if File_cosmos_staking_v1beta1_staking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_staking_v1beta1_staking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoricalInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_staking_v1beta1_staking_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommissionRates); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_staking_v1beta1_staking_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Commission); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_staking_v1beta1_staking_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Description); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_staking_v1beta1_staking_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Validator); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_staking_v1beta1_staking_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValAddresses); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_staking_v1beta1_staking_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DVPair); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_staking_v1beta1_staking_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DVPairs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_staking_v1beta1_staking_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DVVTriplet); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_staking_v1beta1_staking_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DVVTriplets); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_staking_v1beta1_staking_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Delegation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_staking_v1beta1_staking_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnbondingDelegation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_staking_v1beta1_staking_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnbondingDelegationEntry); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_staking_v1beta1_staking_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedelegationEntry); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_staking_v1beta1_staking_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Redelegation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_staking_v1beta1_staking_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Params); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_staking_v1beta1_staking_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelegationResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_staking_v1beta1_staking_proto_msgTypes[17].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedelegationEntryResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_staking_v1beta1_staking_proto_msgTypes[18].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedelegationResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_staking_v1beta1_staking_proto_msgTypes[19].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pool); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cosmos_staking_v1beta1_staking_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   20,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_staking_v1beta1_staking_proto_goTypes,
		DependencyIndexes: file_cosmos_staking_v1beta1_staking_proto_depIdxs,
		EnumInfos:         file_cosmos_staking_v1beta1_staking_proto_enumTypes,
		MessageInfos:      file_cosmos_staking_v1beta1_staking_proto_msgTypes,
	}.Build()
	File_cosmos_staking_v1beta1_staking_proto = out.File
	file_cosmos_staking_v1beta1_staking_proto_rawDesc = nil
	file_cosmos_staking_v1beta1_staking_proto_goTypes = nil
	file_cosmos_staking_v1beta1_staking_proto_depIdxs = nil
}
