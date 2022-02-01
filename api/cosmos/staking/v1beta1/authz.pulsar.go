package stakingv1beta1

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	v1beta1 "github.com/aliworkshop/terra-sdk/api/cosmos/base/v1beta1"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_StakeAuthorization                    protoreflect.MessageDescriptor
	fd_StakeAuthorization_max_tokens         protoreflect.FieldDescriptor
	fd_StakeAuthorization_allow_list         protoreflect.FieldDescriptor
	fd_StakeAuthorization_deny_list          protoreflect.FieldDescriptor
	fd_StakeAuthorization_authorization_type protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_authz_proto_init()
	md_StakeAuthorization = File_cosmos_staking_v1beta1_authz_proto.Messages().ByName("StakeAuthorization")
	fd_StakeAuthorization_max_tokens = md_StakeAuthorization.Fields().ByName("max_tokens")
	fd_StakeAuthorization_allow_list = md_StakeAuthorization.Fields().ByName("allow_list")
	fd_StakeAuthorization_deny_list = md_StakeAuthorization.Fields().ByName("deny_list")
	fd_StakeAuthorization_authorization_type = md_StakeAuthorization.Fields().ByName("authorization_type")
}

var _ protoreflect.Message = (*fastReflection_StakeAuthorization)(nil)

type fastReflection_StakeAuthorization StakeAuthorization

func (x *StakeAuthorization) ProtoReflect() protoreflect.Message {
	return (*fastReflection_StakeAuthorization)(x)
}

func (x *StakeAuthorization) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_authz_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_StakeAuthorization_messageType fastReflection_StakeAuthorization_messageType
var _ protoreflect.MessageType = fastReflection_StakeAuthorization_messageType{}

type fastReflection_StakeAuthorization_messageType struct{}

func (x fastReflection_StakeAuthorization_messageType) Zero() protoreflect.Message {
	return (*fastReflection_StakeAuthorization)(nil)
}
func (x fastReflection_StakeAuthorization_messageType) New() protoreflect.Message {
	return new(fastReflection_StakeAuthorization)
}
func (x fastReflection_StakeAuthorization_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_StakeAuthorization
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_StakeAuthorization) Descriptor() protoreflect.MessageDescriptor {
	return md_StakeAuthorization
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_StakeAuthorization) Type() protoreflect.MessageType {
	return _fastReflection_StakeAuthorization_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_StakeAuthorization) New() protoreflect.Message {
	return new(fastReflection_StakeAuthorization)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_StakeAuthorization) Interface() protoreflect.ProtoMessage {
	return (*StakeAuthorization)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_StakeAuthorization) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.MaxTokens != nil {
		value := protoreflect.ValueOfMessage(x.MaxTokens.ProtoReflect())
		if !f(fd_StakeAuthorization_max_tokens, value) {
			return
		}
	}
	if x.Validators != nil {
		switch o := x.Validators.(type) {
		case *StakeAuthorization_AllowList:
			v := o.AllowList
			value := protoreflect.ValueOfMessage(v.ProtoReflect())
			if !f(fd_StakeAuthorization_allow_list, value) {
				return
			}
		case *StakeAuthorization_DenyList:
			v := o.DenyList
			value := protoreflect.ValueOfMessage(v.ProtoReflect())
			if !f(fd_StakeAuthorization_deny_list, value) {
				return
			}
		}
	}
	if x.AuthorizationType != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.AuthorizationType))
		if !f(fd_StakeAuthorization_authorization_type, value) {
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
func (x *fastReflection_StakeAuthorization) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.StakeAuthorization.max_tokens":
		return x.MaxTokens != nil
	case "cosmos.staking.v1beta1.StakeAuthorization.allow_list":
		if x.Validators == nil {
			return false
		} else if _, ok := x.Validators.(*StakeAuthorization_AllowList); ok {
			return true
		} else {
			return false
		}
	case "cosmos.staking.v1beta1.StakeAuthorization.deny_list":
		if x.Validators == nil {
			return false
		} else if _, ok := x.Validators.(*StakeAuthorization_DenyList); ok {
			return true
		} else {
			return false
		}
	case "cosmos.staking.v1beta1.StakeAuthorization.authorization_type":
		return x.AuthorizationType != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.StakeAuthorization"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.StakeAuthorization does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_StakeAuthorization) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.StakeAuthorization.max_tokens":
		x.MaxTokens = nil
	case "cosmos.staking.v1beta1.StakeAuthorization.allow_list":
		x.Validators = nil
	case "cosmos.staking.v1beta1.StakeAuthorization.deny_list":
		x.Validators = nil
	case "cosmos.staking.v1beta1.StakeAuthorization.authorization_type":
		x.AuthorizationType = 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.StakeAuthorization"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.StakeAuthorization does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_StakeAuthorization) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.StakeAuthorization.max_tokens":
		value := x.MaxTokens
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.staking.v1beta1.StakeAuthorization.allow_list":
		if x.Validators == nil {
			return protoreflect.ValueOfMessage((*StakeAuthorization_Validators)(nil).ProtoReflect())
		} else if v, ok := x.Validators.(*StakeAuthorization_AllowList); ok {
			return protoreflect.ValueOfMessage(v.AllowList.ProtoReflect())
		} else {
			return protoreflect.ValueOfMessage((*StakeAuthorization_Validators)(nil).ProtoReflect())
		}
	case "cosmos.staking.v1beta1.StakeAuthorization.deny_list":
		if x.Validators == nil {
			return protoreflect.ValueOfMessage((*StakeAuthorization_Validators)(nil).ProtoReflect())
		} else if v, ok := x.Validators.(*StakeAuthorization_DenyList); ok {
			return protoreflect.ValueOfMessage(v.DenyList.ProtoReflect())
		} else {
			return protoreflect.ValueOfMessage((*StakeAuthorization_Validators)(nil).ProtoReflect())
		}
	case "cosmos.staking.v1beta1.StakeAuthorization.authorization_type":
		value := x.AuthorizationType
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.StakeAuthorization"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.StakeAuthorization does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_StakeAuthorization) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.StakeAuthorization.max_tokens":
		x.MaxTokens = value.Message().Interface().(*v1beta1.Coin)
	case "cosmos.staking.v1beta1.StakeAuthorization.allow_list":
		cv := value.Message().Interface().(*StakeAuthorization_Validators)
		x.Validators = &StakeAuthorization_AllowList{AllowList: cv}
	case "cosmos.staking.v1beta1.StakeAuthorization.deny_list":
		cv := value.Message().Interface().(*StakeAuthorization_Validators)
		x.Validators = &StakeAuthorization_DenyList{DenyList: cv}
	case "cosmos.staking.v1beta1.StakeAuthorization.authorization_type":
		x.AuthorizationType = (AuthorizationType)(value.Enum())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.StakeAuthorization"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.StakeAuthorization does not contain field %s", fd.FullName()))
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
func (x *fastReflection_StakeAuthorization) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.StakeAuthorization.max_tokens":
		if x.MaxTokens == nil {
			x.MaxTokens = new(v1beta1.Coin)
		}
		return protoreflect.ValueOfMessage(x.MaxTokens.ProtoReflect())
	case "cosmos.staking.v1beta1.StakeAuthorization.allow_list":
		if x.Validators == nil {
			value := &StakeAuthorization_Validators{}
			oneofValue := &StakeAuthorization_AllowList{AllowList: value}
			x.Validators = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
		switch m := x.Validators.(type) {
		case *StakeAuthorization_AllowList:
			return protoreflect.ValueOfMessage(m.AllowList.ProtoReflect())
		default:
			value := &StakeAuthorization_Validators{}
			oneofValue := &StakeAuthorization_AllowList{AllowList: value}
			x.Validators = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
	case "cosmos.staking.v1beta1.StakeAuthorization.deny_list":
		if x.Validators == nil {
			value := &StakeAuthorization_Validators{}
			oneofValue := &StakeAuthorization_DenyList{DenyList: value}
			x.Validators = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
		switch m := x.Validators.(type) {
		case *StakeAuthorization_DenyList:
			return protoreflect.ValueOfMessage(m.DenyList.ProtoReflect())
		default:
			value := &StakeAuthorization_Validators{}
			oneofValue := &StakeAuthorization_DenyList{DenyList: value}
			x.Validators = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
	case "cosmos.staking.v1beta1.StakeAuthorization.authorization_type":
		panic(fmt.Errorf("field authorization_type of message cosmos.staking.v1beta1.StakeAuthorization is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.StakeAuthorization"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.StakeAuthorization does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_StakeAuthorization) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.StakeAuthorization.max_tokens":
		m := new(v1beta1.Coin)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.staking.v1beta1.StakeAuthorization.allow_list":
		value := &StakeAuthorization_Validators{}
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.staking.v1beta1.StakeAuthorization.deny_list":
		value := &StakeAuthorization_Validators{}
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.staking.v1beta1.StakeAuthorization.authorization_type":
		return protoreflect.ValueOfEnum(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.StakeAuthorization"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.StakeAuthorization does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_StakeAuthorization) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	case "cosmos.staking.v1beta1.StakeAuthorization.validators":
		if x.Validators == nil {
			return nil
		}
		switch x.Validators.(type) {
		case *StakeAuthorization_AllowList:
			return x.Descriptor().Fields().ByName("allow_list")
		case *StakeAuthorization_DenyList:
			return x.Descriptor().Fields().ByName("deny_list")
		}
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.StakeAuthorization", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_StakeAuthorization) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_StakeAuthorization) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_StakeAuthorization) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_StakeAuthorization) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*StakeAuthorization)
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
		if x.MaxTokens != nil {
			l = options.Size(x.MaxTokens)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		switch x := x.Validators.(type) {
		case *StakeAuthorization_AllowList:
			if x == nil {
				break
			}
			l = options.Size(x.AllowList)
			n += 1 + l + runtime.Sov(uint64(l))
		case *StakeAuthorization_DenyList:
			if x == nil {
				break
			}
			l = options.Size(x.DenyList)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.AuthorizationType != 0 {
			n += 1 + runtime.Sov(uint64(x.AuthorizationType))
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
		x := input.Message.Interface().(*StakeAuthorization)
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
		switch x := x.Validators.(type) {
		case *StakeAuthorization_AllowList:
			encoded, err := options.Marshal(x.AllowList)
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
		case *StakeAuthorization_DenyList:
			encoded, err := options.Marshal(x.DenyList)
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
		if x.AuthorizationType != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.AuthorizationType))
			i--
			dAtA[i] = 0x20
		}
		if x.MaxTokens != nil {
			encoded, err := options.Marshal(x.MaxTokens)
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
		x := input.Message.Interface().(*StakeAuthorization)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: StakeAuthorization: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: StakeAuthorization: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MaxTokens", wireType)
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
				if x.MaxTokens == nil {
					x.MaxTokens = &v1beta1.Coin{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.MaxTokens); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AllowList", wireType)
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
				v := &StakeAuthorization_Validators{}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], v); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				x.Validators = &StakeAuthorization_AllowList{v}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DenyList", wireType)
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
				v := &StakeAuthorization_Validators{}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], v); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				x.Validators = &StakeAuthorization_DenyList{v}
				iNdEx = postIndex
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AuthorizationType", wireType)
				}
				x.AuthorizationType = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.AuthorizationType |= AuthorizationType(b&0x7F) << shift
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

var _ protoreflect.List = (*_StakeAuthorization_Validators_1_list)(nil)

type _StakeAuthorization_Validators_1_list struct {
	list *[]string
}

func (x *_StakeAuthorization_Validators_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_StakeAuthorization_Validators_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfString((*x.list)[i])
}

func (x *_StakeAuthorization_Validators_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_StakeAuthorization_Validators_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_StakeAuthorization_Validators_1_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message StakeAuthorization_Validators at list field Address as it is not of Message kind"))
}

func (x *_StakeAuthorization_Validators_1_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_StakeAuthorization_Validators_1_list) NewElement() protoreflect.Value {
	v := ""
	return protoreflect.ValueOfString(v)
}

func (x *_StakeAuthorization_Validators_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_StakeAuthorization_Validators         protoreflect.MessageDescriptor
	fd_StakeAuthorization_Validators_address protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_staking_v1beta1_authz_proto_init()
	md_StakeAuthorization_Validators = File_cosmos_staking_v1beta1_authz_proto.Messages().ByName("StakeAuthorization").Messages().ByName("Validators")
	fd_StakeAuthorization_Validators_address = md_StakeAuthorization_Validators.Fields().ByName("address")
}

var _ protoreflect.Message = (*fastReflection_StakeAuthorization_Validators)(nil)

type fastReflection_StakeAuthorization_Validators StakeAuthorization_Validators

func (x *StakeAuthorization_Validators) ProtoReflect() protoreflect.Message {
	return (*fastReflection_StakeAuthorization_Validators)(x)
}

func (x *StakeAuthorization_Validators) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_staking_v1beta1_authz_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_StakeAuthorization_Validators_messageType fastReflection_StakeAuthorization_Validators_messageType
var _ protoreflect.MessageType = fastReflection_StakeAuthorization_Validators_messageType{}

type fastReflection_StakeAuthorization_Validators_messageType struct{}

func (x fastReflection_StakeAuthorization_Validators_messageType) Zero() protoreflect.Message {
	return (*fastReflection_StakeAuthorization_Validators)(nil)
}
func (x fastReflection_StakeAuthorization_Validators_messageType) New() protoreflect.Message {
	return new(fastReflection_StakeAuthorization_Validators)
}
func (x fastReflection_StakeAuthorization_Validators_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_StakeAuthorization_Validators
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_StakeAuthorization_Validators) Descriptor() protoreflect.MessageDescriptor {
	return md_StakeAuthorization_Validators
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_StakeAuthorization_Validators) Type() protoreflect.MessageType {
	return _fastReflection_StakeAuthorization_Validators_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_StakeAuthorization_Validators) New() protoreflect.Message {
	return new(fastReflection_StakeAuthorization_Validators)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_StakeAuthorization_Validators) Interface() protoreflect.ProtoMessage {
	return (*StakeAuthorization_Validators)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_StakeAuthorization_Validators) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Address) != 0 {
		value := protoreflect.ValueOfList(&_StakeAuthorization_Validators_1_list{list: &x.Address})
		if !f(fd_StakeAuthorization_Validators_address, value) {
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
func (x *fastReflection_StakeAuthorization_Validators) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.StakeAuthorization.Validators.address":
		return len(x.Address) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.StakeAuthorization.Validators"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.StakeAuthorization.Validators does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_StakeAuthorization_Validators) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.StakeAuthorization.Validators.address":
		x.Address = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.StakeAuthorization.Validators"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.StakeAuthorization.Validators does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_StakeAuthorization_Validators) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.staking.v1beta1.StakeAuthorization.Validators.address":
		if len(x.Address) == 0 {
			return protoreflect.ValueOfList(&_StakeAuthorization_Validators_1_list{})
		}
		listValue := &_StakeAuthorization_Validators_1_list{list: &x.Address}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.StakeAuthorization.Validators"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.StakeAuthorization.Validators does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_StakeAuthorization_Validators) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.StakeAuthorization.Validators.address":
		lv := value.List()
		clv := lv.(*_StakeAuthorization_Validators_1_list)
		x.Address = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.StakeAuthorization.Validators"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.StakeAuthorization.Validators does not contain field %s", fd.FullName()))
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
func (x *fastReflection_StakeAuthorization_Validators) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.StakeAuthorization.Validators.address":
		if x.Address == nil {
			x.Address = []string{}
		}
		value := &_StakeAuthorization_Validators_1_list{list: &x.Address}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.StakeAuthorization.Validators"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.StakeAuthorization.Validators does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_StakeAuthorization_Validators) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.staking.v1beta1.StakeAuthorization.Validators.address":
		list := []string{}
		return protoreflect.ValueOfList(&_StakeAuthorization_Validators_1_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.staking.v1beta1.StakeAuthorization.Validators"))
		}
		panic(fmt.Errorf("message cosmos.staking.v1beta1.StakeAuthorization.Validators does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_StakeAuthorization_Validators) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.staking.v1beta1.StakeAuthorization.Validators", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_StakeAuthorization_Validators) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_StakeAuthorization_Validators) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_StakeAuthorization_Validators) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_StakeAuthorization_Validators) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*StakeAuthorization_Validators)
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
		if len(x.Address) > 0 {
			for _, s := range x.Address {
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
		x := input.Message.Interface().(*StakeAuthorization_Validators)
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
		if len(x.Address) > 0 {
			for iNdEx := len(x.Address) - 1; iNdEx >= 0; iNdEx-- {
				i -= len(x.Address[iNdEx])
				copy(dAtA[i:], x.Address[iNdEx])
				i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Address[iNdEx])))
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
		x := input.Message.Interface().(*StakeAuthorization_Validators)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: StakeAuthorization_Validators: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: StakeAuthorization_Validators: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
				x.Address = append(x.Address, string(dAtA[iNdEx:postIndex]))
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
// source: cosmos/staking/v1beta1/authz.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// AuthorizationType defines the type of staking module authorization type
//
// Since: cosmos-sdk 0.43
type AuthorizationType int32

const (
	// AUTHORIZATION_TYPE_UNSPECIFIED specifies an unknown authorization type
	AuthorizationType_AUTHORIZATION_TYPE_UNSPECIFIED AuthorizationType = 0
	// AUTHORIZATION_TYPE_DELEGATE defines an authorization type for Msg/Delegate
	AuthorizationType_AUTHORIZATION_TYPE_DELEGATE AuthorizationType = 1
	// AUTHORIZATION_TYPE_UNDELEGATE defines an authorization type for Msg/Undelegate
	AuthorizationType_AUTHORIZATION_TYPE_UNDELEGATE AuthorizationType = 2
	// AUTHORIZATION_TYPE_REDELEGATE defines an authorization type for Msg/BeginRedelegate
	AuthorizationType_AUTHORIZATION_TYPE_REDELEGATE AuthorizationType = 3
)

// Enum value maps for AuthorizationType.
var (
	AuthorizationType_name = map[int32]string{
		0: "AUTHORIZATION_TYPE_UNSPECIFIED",
		1: "AUTHORIZATION_TYPE_DELEGATE",
		2: "AUTHORIZATION_TYPE_UNDELEGATE",
		3: "AUTHORIZATION_TYPE_REDELEGATE",
	}
	AuthorizationType_value = map[string]int32{
		"AUTHORIZATION_TYPE_UNSPECIFIED": 0,
		"AUTHORIZATION_TYPE_DELEGATE":    1,
		"AUTHORIZATION_TYPE_UNDELEGATE":  2,
		"AUTHORIZATION_TYPE_REDELEGATE":  3,
	}
)

func (x AuthorizationType) Enum() *AuthorizationType {
	p := new(AuthorizationType)
	*p = x
	return p
}

func (x AuthorizationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthorizationType) Descriptor() protoreflect.EnumDescriptor {
	return file_cosmos_staking_v1beta1_authz_proto_enumTypes[0].Descriptor()
}

func (AuthorizationType) Type() protoreflect.EnumType {
	return &file_cosmos_staking_v1beta1_authz_proto_enumTypes[0]
}

func (x AuthorizationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthorizationType.Descriptor instead.
func (AuthorizationType) EnumDescriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_authz_proto_rawDescGZIP(), []int{0}
}

// StakeAuthorization defines authorization for delegate/undelegate/redelegate.
//
// Since: cosmos-sdk 0.43
type StakeAuthorization struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// max_tokens specifies the maximum amount of tokens can be delegate to a validator. If it is
	// empty, there is no spend limit and any amount of coins can be delegated.
	MaxTokens *v1beta1.Coin `protobuf:"bytes,1,opt,name=max_tokens,json=maxTokens,proto3" json:"max_tokens,omitempty"`
	// validators is the oneof that represents either allow_list or deny_list
	//
	// Types that are assignable to Validators:
	//	*StakeAuthorization_AllowList
	//	*StakeAuthorization_DenyList
	Validators isStakeAuthorization_Validators `protobuf_oneof:"validators"`
	// authorization_type defines one of AuthorizationType.
	AuthorizationType AuthorizationType `protobuf:"varint,4,opt,name=authorization_type,json=authorizationType,proto3,enum=cosmos.staking.v1beta1.AuthorizationType" json:"authorization_type,omitempty"`
}

func (x *StakeAuthorization) Reset() {
	*x = StakeAuthorization{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_authz_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StakeAuthorization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StakeAuthorization) ProtoMessage() {}

// Deprecated: Use StakeAuthorization.ProtoReflect.Descriptor instead.
func (*StakeAuthorization) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_authz_proto_rawDescGZIP(), []int{0}
}

func (x *StakeAuthorization) GetMaxTokens() *v1beta1.Coin {
	if x != nil {
		return x.MaxTokens
	}
	return nil
}

func (x *StakeAuthorization) GetValidators() isStakeAuthorization_Validators {
	if x != nil {
		return x.Validators
	}
	return nil
}

func (x *StakeAuthorization) GetAllowList() *StakeAuthorization_Validators {
	if x, ok := x.GetValidators().(*StakeAuthorization_AllowList); ok {
		return x.AllowList
	}
	return nil
}

func (x *StakeAuthorization) GetDenyList() *StakeAuthorization_Validators {
	if x, ok := x.GetValidators().(*StakeAuthorization_DenyList); ok {
		return x.DenyList
	}
	return nil
}

func (x *StakeAuthorization) GetAuthorizationType() AuthorizationType {
	if x != nil {
		return x.AuthorizationType
	}
	return AuthorizationType_AUTHORIZATION_TYPE_UNSPECIFIED
}

type isStakeAuthorization_Validators interface {
	isStakeAuthorization_Validators()
}

type StakeAuthorization_AllowList struct {
	// allow_list specifies list of validator addresses to whom grantee can delegate tokens on behalf of granter's
	// account.
	AllowList *StakeAuthorization_Validators `protobuf:"bytes,2,opt,name=allow_list,json=allowList,proto3,oneof"`
}

type StakeAuthorization_DenyList struct {
	// deny_list specifies list of validator addresses to whom grantee can not delegate tokens.
	DenyList *StakeAuthorization_Validators `protobuf:"bytes,3,opt,name=deny_list,json=denyList,proto3,oneof"`
}

func (*StakeAuthorization_AllowList) isStakeAuthorization_Validators() {}

func (*StakeAuthorization_DenyList) isStakeAuthorization_Validators() {}

// Validators defines list of validator addresses.
type StakeAuthorization_Validators struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address []string `protobuf:"bytes,1,rep,name=address,proto3" json:"address,omitempty"`
}

func (x *StakeAuthorization_Validators) Reset() {
	*x = StakeAuthorization_Validators{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_staking_v1beta1_authz_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StakeAuthorization_Validators) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StakeAuthorization_Validators) ProtoMessage() {}

// Deprecated: Use StakeAuthorization_Validators.ProtoReflect.Descriptor instead.
func (*StakeAuthorization_Validators) Descriptor() ([]byte, []int) {
	return file_cosmos_staking_v1beta1_authz_proto_rawDescGZIP(), []int{0, 0}
}

func (x *StakeAuthorization_Validators) GetAddress() []string {
	if x != nil {
		return x.Address
	}
	return nil
}

var File_cosmos_staking_v1beta1_authz_proto protoreflect.FileDescriptor

var file_cosmos_staking_v1beta1_authz_proto_rawDesc = []byte{
	0x0a, 0x22, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x14, 0x67, 0x6f,
	0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x03,
	0x0a, 0x12, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x65, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43,
	0x6f, 0x69, 0x6e, 0x42, 0x2b, 0xaa, 0xdf, 0x1f, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x69, 0x6e,
	0x52, 0x09, 0x6d, 0x61, 0x78, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x56, 0x0a, 0x0a, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x35, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x48, 0x00, 0x52, 0x09, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x54, 0x0a, 0x09, 0x64, 0x65, 0x6e, 0x79, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x53, 0x74, 0x61, 0x6b, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x48, 0x00, 0x52,
	0x08, 0x64, 0x65, 0x6e, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x58, 0x0a, 0x12, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73,
	0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x11, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x1a, 0x40, 0x0a, 0x0a, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x12, 0x32, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x3a, 0x11, 0xca, 0xb4, 0x2d, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0c, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2a, 0x9e, 0x01, 0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x1e,
	0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x1f, 0x0a, 0x1b, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x47, 0x41, 0x54, 0x45, 0x10,
	0x01, 0x12, 0x21, 0x0a, 0x1d, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x4c, 0x45, 0x47, 0x41,
	0x54, 0x45, 0x10, 0x02, 0x12, 0x21, 0x0a, 0x1d, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x44, 0x45, 0x4c,
	0x45, 0x47, 0x41, 0x54, 0x45, 0x10, 0x03, 0x42, 0xea, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73,
	0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x73, 0x74,
	0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x73, 0x74,
	0x61, 0x6b, 0x69, 0x6e, 0x67, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43,
	0x53, 0x58, 0xaa, 0x02, 0x16, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x16, 0x43, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x22, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x53, 0x74,
	0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x18, 0x43, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x3a, 0x3a, 0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_staking_v1beta1_authz_proto_rawDescOnce sync.Once
	file_cosmos_staking_v1beta1_authz_proto_rawDescData = file_cosmos_staking_v1beta1_authz_proto_rawDesc
)

func file_cosmos_staking_v1beta1_authz_proto_rawDescGZIP() []byte {
	file_cosmos_staking_v1beta1_authz_proto_rawDescOnce.Do(func() {
		file_cosmos_staking_v1beta1_authz_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_staking_v1beta1_authz_proto_rawDescData)
	})
	return file_cosmos_staking_v1beta1_authz_proto_rawDescData
}

var file_cosmos_staking_v1beta1_authz_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cosmos_staking_v1beta1_authz_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cosmos_staking_v1beta1_authz_proto_goTypes = []interface{}{
	(AuthorizationType)(0),                // 0: cosmos.staking.v1beta1.AuthorizationType
	(*StakeAuthorization)(nil),            // 1: cosmos.staking.v1beta1.StakeAuthorization
	(*StakeAuthorization_Validators)(nil), // 2: cosmos.staking.v1beta1.StakeAuthorization.Validators
	(*v1beta1.Coin)(nil),                  // 3: cosmos.base.v1beta1.Coin
}
var file_cosmos_staking_v1beta1_authz_proto_depIdxs = []int32{
	3, // 0: cosmos.staking.v1beta1.StakeAuthorization.max_tokens:type_name -> cosmos.base.v1beta1.Coin
	2, // 1: cosmos.staking.v1beta1.StakeAuthorization.allow_list:type_name -> cosmos.staking.v1beta1.StakeAuthorization.Validators
	2, // 2: cosmos.staking.v1beta1.StakeAuthorization.deny_list:type_name -> cosmos.staking.v1beta1.StakeAuthorization.Validators
	0, // 3: cosmos.staking.v1beta1.StakeAuthorization.authorization_type:type_name -> cosmos.staking.v1beta1.AuthorizationType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_cosmos_staking_v1beta1_authz_proto_init() }
func file_cosmos_staking_v1beta1_authz_proto_init() {
	if File_cosmos_staking_v1beta1_authz_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_staking_v1beta1_authz_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StakeAuthorization); i {
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
		file_cosmos_staking_v1beta1_authz_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StakeAuthorization_Validators); i {
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
	file_cosmos_staking_v1beta1_authz_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*StakeAuthorization_AllowList)(nil),
		(*StakeAuthorization_DenyList)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cosmos_staking_v1beta1_authz_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_staking_v1beta1_authz_proto_goTypes,
		DependencyIndexes: file_cosmos_staking_v1beta1_authz_proto_depIdxs,
		EnumInfos:         file_cosmos_staking_v1beta1_authz_proto_enumTypes,
		MessageInfos:      file_cosmos_staking_v1beta1_authz_proto_msgTypes,
	}.Build()
	File_cosmos_staking_v1beta1_authz_proto = out.File
	file_cosmos_staking_v1beta1_authz_proto_rawDesc = nil
	file_cosmos_staking_v1beta1_authz_proto_goTypes = nil
	file_cosmos_staking_v1beta1_authz_proto_depIdxs = nil
}
