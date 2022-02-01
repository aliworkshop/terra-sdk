package types

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	crypto "github.com/aliworkshop/terra-sdk/api/tendermint/crypto"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var _ protoreflect.List = (*_ValidatorSet_1_list)(nil)

type _ValidatorSet_1_list struct {
	list *[]*Validator
}

func (x *_ValidatorSet_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_ValidatorSet_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_ValidatorSet_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Validator)
	(*x.list)[i] = concreteValue
}

func (x *_ValidatorSet_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Validator)
	*x.list = append(*x.list, concreteValue)
}

func (x *_ValidatorSet_1_list) AppendMutable() protoreflect.Value {
	v := new(Validator)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_ValidatorSet_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_ValidatorSet_1_list) NewElement() protoreflect.Value {
	v := new(Validator)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_ValidatorSet_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_ValidatorSet                    protoreflect.MessageDescriptor
	fd_ValidatorSet_validators         protoreflect.FieldDescriptor
	fd_ValidatorSet_proposer           protoreflect.FieldDescriptor
	fd_ValidatorSet_total_voting_power protoreflect.FieldDescriptor
)

func init() {
	file_tendermint_types_validator_proto_init()
	md_ValidatorSet = File_tendermint_types_validator_proto.Messages().ByName("ValidatorSet")
	fd_ValidatorSet_validators = md_ValidatorSet.Fields().ByName("validators")
	fd_ValidatorSet_proposer = md_ValidatorSet.Fields().ByName("proposer")
	fd_ValidatorSet_total_voting_power = md_ValidatorSet.Fields().ByName("total_voting_power")
}

var _ protoreflect.Message = (*fastReflection_ValidatorSet)(nil)

type fastReflection_ValidatorSet ValidatorSet

func (x *ValidatorSet) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ValidatorSet)(x)
}

func (x *ValidatorSet) slowProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_validator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ValidatorSet_messageType fastReflection_ValidatorSet_messageType
var _ protoreflect.MessageType = fastReflection_ValidatorSet_messageType{}

type fastReflection_ValidatorSet_messageType struct{}

func (x fastReflection_ValidatorSet_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ValidatorSet)(nil)
}
func (x fastReflection_ValidatorSet_messageType) New() protoreflect.Message {
	return new(fastReflection_ValidatorSet)
}
func (x fastReflection_ValidatorSet_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ValidatorSet
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ValidatorSet) Descriptor() protoreflect.MessageDescriptor {
	return md_ValidatorSet
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ValidatorSet) Type() protoreflect.MessageType {
	return _fastReflection_ValidatorSet_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ValidatorSet) New() protoreflect.Message {
	return new(fastReflection_ValidatorSet)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ValidatorSet) Interface() protoreflect.ProtoMessage {
	return (*ValidatorSet)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ValidatorSet) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Validators) != 0 {
		value := protoreflect.ValueOfList(&_ValidatorSet_1_list{list: &x.Validators})
		if !f(fd_ValidatorSet_validators, value) {
			return
		}
	}
	if x.Proposer != nil {
		value := protoreflect.ValueOfMessage(x.Proposer.ProtoReflect())
		if !f(fd_ValidatorSet_proposer, value) {
			return
		}
	}
	if x.TotalVotingPower != int64(0) {
		value := protoreflect.ValueOfInt64(x.TotalVotingPower)
		if !f(fd_ValidatorSet_total_voting_power, value) {
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
func (x *fastReflection_ValidatorSet) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "tendermint.types.ValidatorSet.validators":
		return len(x.Validators) != 0
	case "tendermint.types.ValidatorSet.proposer":
		return x.Proposer != nil
	case "tendermint.types.ValidatorSet.total_voting_power":
		return x.TotalVotingPower != int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.ValidatorSet"))
		}
		panic(fmt.Errorf("message tendermint.types.ValidatorSet does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorSet) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "tendermint.types.ValidatorSet.validators":
		x.Validators = nil
	case "tendermint.types.ValidatorSet.proposer":
		x.Proposer = nil
	case "tendermint.types.ValidatorSet.total_voting_power":
		x.TotalVotingPower = int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.ValidatorSet"))
		}
		panic(fmt.Errorf("message tendermint.types.ValidatorSet does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ValidatorSet) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "tendermint.types.ValidatorSet.validators":
		if len(x.Validators) == 0 {
			return protoreflect.ValueOfList(&_ValidatorSet_1_list{})
		}
		listValue := &_ValidatorSet_1_list{list: &x.Validators}
		return protoreflect.ValueOfList(listValue)
	case "tendermint.types.ValidatorSet.proposer":
		value := x.Proposer
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "tendermint.types.ValidatorSet.total_voting_power":
		value := x.TotalVotingPower
		return protoreflect.ValueOfInt64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.ValidatorSet"))
		}
		panic(fmt.Errorf("message tendermint.types.ValidatorSet does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_ValidatorSet) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "tendermint.types.ValidatorSet.validators":
		lv := value.List()
		clv := lv.(*_ValidatorSet_1_list)
		x.Validators = *clv.list
	case "tendermint.types.ValidatorSet.proposer":
		x.Proposer = value.Message().Interface().(*Validator)
	case "tendermint.types.ValidatorSet.total_voting_power":
		x.TotalVotingPower = value.Int()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.ValidatorSet"))
		}
		panic(fmt.Errorf("message tendermint.types.ValidatorSet does not contain field %s", fd.FullName()))
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
func (x *fastReflection_ValidatorSet) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.ValidatorSet.validators":
		if x.Validators == nil {
			x.Validators = []*Validator{}
		}
		value := &_ValidatorSet_1_list{list: &x.Validators}
		return protoreflect.ValueOfList(value)
	case "tendermint.types.ValidatorSet.proposer":
		if x.Proposer == nil {
			x.Proposer = new(Validator)
		}
		return protoreflect.ValueOfMessage(x.Proposer.ProtoReflect())
	case "tendermint.types.ValidatorSet.total_voting_power":
		panic(fmt.Errorf("field total_voting_power of message tendermint.types.ValidatorSet is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.ValidatorSet"))
		}
		panic(fmt.Errorf("message tendermint.types.ValidatorSet does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ValidatorSet) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.ValidatorSet.validators":
		list := []*Validator{}
		return protoreflect.ValueOfList(&_ValidatorSet_1_list{list: &list})
	case "tendermint.types.ValidatorSet.proposer":
		m := new(Validator)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "tendermint.types.ValidatorSet.total_voting_power":
		return protoreflect.ValueOfInt64(int64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.ValidatorSet"))
		}
		panic(fmt.Errorf("message tendermint.types.ValidatorSet does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ValidatorSet) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in tendermint.types.ValidatorSet", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ValidatorSet) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorSet) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_ValidatorSet) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ValidatorSet) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ValidatorSet)
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
		if len(x.Validators) > 0 {
			for _, e := range x.Validators {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.Proposer != nil {
			l = options.Size(x.Proposer)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.TotalVotingPower != 0 {
			n += 1 + runtime.Sov(uint64(x.TotalVotingPower))
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
		x := input.Message.Interface().(*ValidatorSet)
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
		if x.TotalVotingPower != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.TotalVotingPower))
			i--
			dAtA[i] = 0x18
		}
		if x.Proposer != nil {
			encoded, err := options.Marshal(x.Proposer)
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
		if len(x.Validators) > 0 {
			for iNdEx := len(x.Validators) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Validators[iNdEx])
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
		x := input.Message.Interface().(*ValidatorSet)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ValidatorSet: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ValidatorSet: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
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
				x.Validators = append(x.Validators, &Validator{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Validators[len(x.Validators)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Proposer", wireType)
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
				if x.Proposer == nil {
					x.Proposer = &Validator{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Proposer); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TotalVotingPower", wireType)
				}
				x.TotalVotingPower = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.TotalVotingPower |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
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
	md_Validator                   protoreflect.MessageDescriptor
	fd_Validator_address           protoreflect.FieldDescriptor
	fd_Validator_pub_key           protoreflect.FieldDescriptor
	fd_Validator_voting_power      protoreflect.FieldDescriptor
	fd_Validator_proposer_priority protoreflect.FieldDescriptor
)

func init() {
	file_tendermint_types_validator_proto_init()
	md_Validator = File_tendermint_types_validator_proto.Messages().ByName("Validator")
	fd_Validator_address = md_Validator.Fields().ByName("address")
	fd_Validator_pub_key = md_Validator.Fields().ByName("pub_key")
	fd_Validator_voting_power = md_Validator.Fields().ByName("voting_power")
	fd_Validator_proposer_priority = md_Validator.Fields().ByName("proposer_priority")
}

var _ protoreflect.Message = (*fastReflection_Validator)(nil)

type fastReflection_Validator Validator

func (x *Validator) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Validator)(x)
}

func (x *Validator) slowProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_validator_proto_msgTypes[1]
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
	if len(x.Address) != 0 {
		value := protoreflect.ValueOfBytes(x.Address)
		if !f(fd_Validator_address, value) {
			return
		}
	}
	if x.PubKey != nil {
		value := protoreflect.ValueOfMessage(x.PubKey.ProtoReflect())
		if !f(fd_Validator_pub_key, value) {
			return
		}
	}
	if x.VotingPower != int64(0) {
		value := protoreflect.ValueOfInt64(x.VotingPower)
		if !f(fd_Validator_voting_power, value) {
			return
		}
	}
	if x.ProposerPriority != int64(0) {
		value := protoreflect.ValueOfInt64(x.ProposerPriority)
		if !f(fd_Validator_proposer_priority, value) {
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
	case "tendermint.types.Validator.address":
		return len(x.Address) != 0
	case "tendermint.types.Validator.pub_key":
		return x.PubKey != nil
	case "tendermint.types.Validator.voting_power":
		return x.VotingPower != int64(0)
	case "tendermint.types.Validator.proposer_priority":
		return x.ProposerPriority != int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Validator"))
		}
		panic(fmt.Errorf("message tendermint.types.Validator does not contain field %s", fd.FullName()))
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
	case "tendermint.types.Validator.address":
		x.Address = nil
	case "tendermint.types.Validator.pub_key":
		x.PubKey = nil
	case "tendermint.types.Validator.voting_power":
		x.VotingPower = int64(0)
	case "tendermint.types.Validator.proposer_priority":
		x.ProposerPriority = int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Validator"))
		}
		panic(fmt.Errorf("message tendermint.types.Validator does not contain field %s", fd.FullName()))
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
	case "tendermint.types.Validator.address":
		value := x.Address
		return protoreflect.ValueOfBytes(value)
	case "tendermint.types.Validator.pub_key":
		value := x.PubKey
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "tendermint.types.Validator.voting_power":
		value := x.VotingPower
		return protoreflect.ValueOfInt64(value)
	case "tendermint.types.Validator.proposer_priority":
		value := x.ProposerPriority
		return protoreflect.ValueOfInt64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Validator"))
		}
		panic(fmt.Errorf("message tendermint.types.Validator does not contain field %s", descriptor.FullName()))
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
	case "tendermint.types.Validator.address":
		x.Address = value.Bytes()
	case "tendermint.types.Validator.pub_key":
		x.PubKey = value.Message().Interface().(*crypto.PublicKey)
	case "tendermint.types.Validator.voting_power":
		x.VotingPower = value.Int()
	case "tendermint.types.Validator.proposer_priority":
		x.ProposerPriority = value.Int()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Validator"))
		}
		panic(fmt.Errorf("message tendermint.types.Validator does not contain field %s", fd.FullName()))
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
	case "tendermint.types.Validator.pub_key":
		if x.PubKey == nil {
			x.PubKey = new(crypto.PublicKey)
		}
		return protoreflect.ValueOfMessage(x.PubKey.ProtoReflect())
	case "tendermint.types.Validator.address":
		panic(fmt.Errorf("field address of message tendermint.types.Validator is not mutable"))
	case "tendermint.types.Validator.voting_power":
		panic(fmt.Errorf("field voting_power of message tendermint.types.Validator is not mutable"))
	case "tendermint.types.Validator.proposer_priority":
		panic(fmt.Errorf("field proposer_priority of message tendermint.types.Validator is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Validator"))
		}
		panic(fmt.Errorf("message tendermint.types.Validator does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Validator) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.Validator.address":
		return protoreflect.ValueOfBytes(nil)
	case "tendermint.types.Validator.pub_key":
		m := new(crypto.PublicKey)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "tendermint.types.Validator.voting_power":
		return protoreflect.ValueOfInt64(int64(0))
	case "tendermint.types.Validator.proposer_priority":
		return protoreflect.ValueOfInt64(int64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Validator"))
		}
		panic(fmt.Errorf("message tendermint.types.Validator does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Validator) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in tendermint.types.Validator", d.FullName()))
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
		l = len(x.Address)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.PubKey != nil {
			l = options.Size(x.PubKey)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.VotingPower != 0 {
			n += 1 + runtime.Sov(uint64(x.VotingPower))
		}
		if x.ProposerPriority != 0 {
			n += 1 + runtime.Sov(uint64(x.ProposerPriority))
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
		if x.ProposerPriority != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.ProposerPriority))
			i--
			dAtA[i] = 0x20
		}
		if x.VotingPower != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.VotingPower))
			i--
			dAtA[i] = 0x18
		}
		if x.PubKey != nil {
			encoded, err := options.Marshal(x.PubKey)
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
		if len(x.Address) > 0 {
			i -= len(x.Address)
			copy(dAtA[i:], x.Address)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Address)))
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Address = append(x.Address[:0], dAtA[iNdEx:postIndex]...)
				if x.Address == nil {
					x.Address = []byte{}
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
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
				if x.PubKey == nil {
					x.PubKey = &crypto.PublicKey{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.PubKey); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field VotingPower", wireType)
				}
				x.VotingPower = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.VotingPower |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProposerPriority", wireType)
				}
				x.ProposerPriority = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.ProposerPriority |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
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
	md_SimpleValidator              protoreflect.MessageDescriptor
	fd_SimpleValidator_pub_key      protoreflect.FieldDescriptor
	fd_SimpleValidator_voting_power protoreflect.FieldDescriptor
)

func init() {
	file_tendermint_types_validator_proto_init()
	md_SimpleValidator = File_tendermint_types_validator_proto.Messages().ByName("SimpleValidator")
	fd_SimpleValidator_pub_key = md_SimpleValidator.Fields().ByName("pub_key")
	fd_SimpleValidator_voting_power = md_SimpleValidator.Fields().ByName("voting_power")
}

var _ protoreflect.Message = (*fastReflection_SimpleValidator)(nil)

type fastReflection_SimpleValidator SimpleValidator

func (x *SimpleValidator) ProtoReflect() protoreflect.Message {
	return (*fastReflection_SimpleValidator)(x)
}

func (x *SimpleValidator) slowProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_validator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_SimpleValidator_messageType fastReflection_SimpleValidator_messageType
var _ protoreflect.MessageType = fastReflection_SimpleValidator_messageType{}

type fastReflection_SimpleValidator_messageType struct{}

func (x fastReflection_SimpleValidator_messageType) Zero() protoreflect.Message {
	return (*fastReflection_SimpleValidator)(nil)
}
func (x fastReflection_SimpleValidator_messageType) New() protoreflect.Message {
	return new(fastReflection_SimpleValidator)
}
func (x fastReflection_SimpleValidator_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_SimpleValidator
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_SimpleValidator) Descriptor() protoreflect.MessageDescriptor {
	return md_SimpleValidator
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_SimpleValidator) Type() protoreflect.MessageType {
	return _fastReflection_SimpleValidator_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_SimpleValidator) New() protoreflect.Message {
	return new(fastReflection_SimpleValidator)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_SimpleValidator) Interface() protoreflect.ProtoMessage {
	return (*SimpleValidator)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_SimpleValidator) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.PubKey != nil {
		value := protoreflect.ValueOfMessage(x.PubKey.ProtoReflect())
		if !f(fd_SimpleValidator_pub_key, value) {
			return
		}
	}
	if x.VotingPower != int64(0) {
		value := protoreflect.ValueOfInt64(x.VotingPower)
		if !f(fd_SimpleValidator_voting_power, value) {
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
func (x *fastReflection_SimpleValidator) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "tendermint.types.SimpleValidator.pub_key":
		return x.PubKey != nil
	case "tendermint.types.SimpleValidator.voting_power":
		return x.VotingPower != int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.SimpleValidator"))
		}
		panic(fmt.Errorf("message tendermint.types.SimpleValidator does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SimpleValidator) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "tendermint.types.SimpleValidator.pub_key":
		x.PubKey = nil
	case "tendermint.types.SimpleValidator.voting_power":
		x.VotingPower = int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.SimpleValidator"))
		}
		panic(fmt.Errorf("message tendermint.types.SimpleValidator does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_SimpleValidator) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "tendermint.types.SimpleValidator.pub_key":
		value := x.PubKey
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "tendermint.types.SimpleValidator.voting_power":
		value := x.VotingPower
		return protoreflect.ValueOfInt64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.SimpleValidator"))
		}
		panic(fmt.Errorf("message tendermint.types.SimpleValidator does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_SimpleValidator) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "tendermint.types.SimpleValidator.pub_key":
		x.PubKey = value.Message().Interface().(*crypto.PublicKey)
	case "tendermint.types.SimpleValidator.voting_power":
		x.VotingPower = value.Int()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.SimpleValidator"))
		}
		panic(fmt.Errorf("message tendermint.types.SimpleValidator does not contain field %s", fd.FullName()))
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
func (x *fastReflection_SimpleValidator) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.SimpleValidator.pub_key":
		if x.PubKey == nil {
			x.PubKey = new(crypto.PublicKey)
		}
		return protoreflect.ValueOfMessage(x.PubKey.ProtoReflect())
	case "tendermint.types.SimpleValidator.voting_power":
		panic(fmt.Errorf("field voting_power of message tendermint.types.SimpleValidator is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.SimpleValidator"))
		}
		panic(fmt.Errorf("message tendermint.types.SimpleValidator does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_SimpleValidator) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.SimpleValidator.pub_key":
		m := new(crypto.PublicKey)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "tendermint.types.SimpleValidator.voting_power":
		return protoreflect.ValueOfInt64(int64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.SimpleValidator"))
		}
		panic(fmt.Errorf("message tendermint.types.SimpleValidator does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_SimpleValidator) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in tendermint.types.SimpleValidator", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_SimpleValidator) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SimpleValidator) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_SimpleValidator) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_SimpleValidator) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*SimpleValidator)
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
		if x.PubKey != nil {
			l = options.Size(x.PubKey)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.VotingPower != 0 {
			n += 1 + runtime.Sov(uint64(x.VotingPower))
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
		x := input.Message.Interface().(*SimpleValidator)
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
		if x.VotingPower != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.VotingPower))
			i--
			dAtA[i] = 0x10
		}
		if x.PubKey != nil {
			encoded, err := options.Marshal(x.PubKey)
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
		x := input.Message.Interface().(*SimpleValidator)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SimpleValidator: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SimpleValidator: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
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
				if x.PubKey == nil {
					x.PubKey = &crypto.PublicKey{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.PubKey); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field VotingPower", wireType)
				}
				x.VotingPower = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.VotingPower |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
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
// source: tendermint/types/validator.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ValidatorSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Validators       []*Validator `protobuf:"bytes,1,rep,name=validators,proto3" json:"validators,omitempty"`
	Proposer         *Validator   `protobuf:"bytes,2,opt,name=proposer,proto3" json:"proposer,omitempty"`
	TotalVotingPower int64        `protobuf:"varint,3,opt,name=total_voting_power,json=totalVotingPower,proto3" json:"total_voting_power,omitempty"`
}

func (x *ValidatorSet) Reset() {
	*x = ValidatorSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_validator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatorSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatorSet) ProtoMessage() {}

// Deprecated: Use ValidatorSet.ProtoReflect.Descriptor instead.
func (*ValidatorSet) Descriptor() ([]byte, []int) {
	return file_tendermint_types_validator_proto_rawDescGZIP(), []int{0}
}

func (x *ValidatorSet) GetValidators() []*Validator {
	if x != nil {
		return x.Validators
	}
	return nil
}

func (x *ValidatorSet) GetProposer() *Validator {
	if x != nil {
		return x.Proposer
	}
	return nil
}

func (x *ValidatorSet) GetTotalVotingPower() int64 {
	if x != nil {
		return x.TotalVotingPower
	}
	return 0
}

type Validator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address          []byte            `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	PubKey           *crypto.PublicKey `protobuf:"bytes,2,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	VotingPower      int64             `protobuf:"varint,3,opt,name=voting_power,json=votingPower,proto3" json:"voting_power,omitempty"`
	ProposerPriority int64             `protobuf:"varint,4,opt,name=proposer_priority,json=proposerPriority,proto3" json:"proposer_priority,omitempty"`
}

func (x *Validator) Reset() {
	*x = Validator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_validator_proto_msgTypes[1]
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
	return file_tendermint_types_validator_proto_rawDescGZIP(), []int{1}
}

func (x *Validator) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Validator) GetPubKey() *crypto.PublicKey {
	if x != nil {
		return x.PubKey
	}
	return nil
}

func (x *Validator) GetVotingPower() int64 {
	if x != nil {
		return x.VotingPower
	}
	return 0
}

func (x *Validator) GetProposerPriority() int64 {
	if x != nil {
		return x.ProposerPriority
	}
	return 0
}

type SimpleValidator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PubKey      *crypto.PublicKey `protobuf:"bytes,1,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	VotingPower int64             `protobuf:"varint,2,opt,name=voting_power,json=votingPower,proto3" json:"voting_power,omitempty"`
}

func (x *SimpleValidator) Reset() {
	*x = SimpleValidator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_validator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleValidator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleValidator) ProtoMessage() {}

// Deprecated: Use SimpleValidator.ProtoReflect.Descriptor instead.
func (*SimpleValidator) Descriptor() ([]byte, []int) {
	return file_tendermint_types_validator_proto_rawDescGZIP(), []int{2}
}

func (x *SimpleValidator) GetPubKey() *crypto.PublicKey {
	if x != nil {
		return x.PubKey
	}
	return nil
}

func (x *SimpleValidator) GetVotingPower() int64 {
	if x != nil {
		return x.VotingPower
	}
	return 0
}

var File_tendermint_types_validator_proto protoreflect.FileDescriptor

var file_tendermint_types_validator_proto_rawDesc = []byte{
	0x0a, 0x20, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x74, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x6b, 0x65,
	0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x01, 0x0a, 0x0c, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x12, 0x3b, 0x0a, 0x0a, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x37, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x12,
	0x2c, 0x0a, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x70, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x56, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x22, 0xb2, 0x01,
	0x0a, 0x09, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x3b, 0x0a, 0x07, 0x70, 0x75, 0x62, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x74, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x06, 0x70, 0x75, 0x62, 0x4b,
	0x65, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x77,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67,
	0x50, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65,
	0x72, 0x5f, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x10, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x22, 0x6b, 0x0a, 0x0f, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x35, 0x0a, 0x07, 0x70, 0x75, 0x62, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x74, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x52, 0x06, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x0c,
	0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x42,
	0xba, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x42, 0x0e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0xa2, 0x02, 0x03,
	0x54, 0x54, 0x58, 0xaa, 0x02, 0x10, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0xca, 0x02, 0x10, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x74, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x73, 0xe2, 0x02, 0x1c, 0x54, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x54, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x74, 0x3a, 0x3a, 0x54, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tendermint_types_validator_proto_rawDescOnce sync.Once
	file_tendermint_types_validator_proto_rawDescData = file_tendermint_types_validator_proto_rawDesc
)

func file_tendermint_types_validator_proto_rawDescGZIP() []byte {
	file_tendermint_types_validator_proto_rawDescOnce.Do(func() {
		file_tendermint_types_validator_proto_rawDescData = protoimpl.X.CompressGZIP(file_tendermint_types_validator_proto_rawDescData)
	})
	return file_tendermint_types_validator_proto_rawDescData
}

var file_tendermint_types_validator_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tendermint_types_validator_proto_goTypes = []interface{}{
	(*ValidatorSet)(nil),     // 0: tendermint.types.ValidatorSet
	(*Validator)(nil),        // 1: tendermint.types.Validator
	(*SimpleValidator)(nil),  // 2: tendermint.types.SimpleValidator
	(*crypto.PublicKey)(nil), // 3: tendermint.crypto.PublicKey
}
var file_tendermint_types_validator_proto_depIdxs = []int32{
	1, // 0: tendermint.types.ValidatorSet.validators:type_name -> tendermint.types.Validator
	1, // 1: tendermint.types.ValidatorSet.proposer:type_name -> tendermint.types.Validator
	3, // 2: tendermint.types.Validator.pub_key:type_name -> tendermint.crypto.PublicKey
	3, // 3: tendermint.types.SimpleValidator.pub_key:type_name -> tendermint.crypto.PublicKey
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tendermint_types_validator_proto_init() }
func file_tendermint_types_validator_proto_init() {
	if File_tendermint_types_validator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tendermint_types_validator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatorSet); i {
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
		file_tendermint_types_validator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_tendermint_types_validator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleValidator); i {
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
			RawDescriptor: file_tendermint_types_validator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tendermint_types_validator_proto_goTypes,
		DependencyIndexes: file_tendermint_types_validator_proto_depIdxs,
		MessageInfos:      file_tendermint_types_validator_proto_msgTypes,
	}.Build()
	File_tendermint_types_validator_proto = out.File
	file_tendermint_types_validator_proto_rawDesc = nil
	file_tendermint_types_validator_proto_goTypes = nil
	file_tendermint_types_validator_proto_depIdxs = nil
}
