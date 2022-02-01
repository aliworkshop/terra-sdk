package signingv1beta1

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	v1beta1 "github.com/aliworkshop/terra-sdk/api/cosmos/crypto/multisig/v1beta1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var _ protoreflect.List = (*_SignatureDescriptors_1_list)(nil)

type _SignatureDescriptors_1_list struct {
	list *[]*SignatureDescriptor
}

func (x *_SignatureDescriptors_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_SignatureDescriptors_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_SignatureDescriptors_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*SignatureDescriptor)
	(*x.list)[i] = concreteValue
}

func (x *_SignatureDescriptors_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*SignatureDescriptor)
	*x.list = append(*x.list, concreteValue)
}

func (x *_SignatureDescriptors_1_list) AppendMutable() protoreflect.Value {
	v := new(SignatureDescriptor)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_SignatureDescriptors_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_SignatureDescriptors_1_list) NewElement() protoreflect.Value {
	v := new(SignatureDescriptor)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_SignatureDescriptors_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_SignatureDescriptors            protoreflect.MessageDescriptor
	fd_SignatureDescriptors_signatures protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_tx_signing_v1beta1_signing_proto_init()
	md_SignatureDescriptors = File_cosmos_tx_signing_v1beta1_signing_proto.Messages().ByName("SignatureDescriptors")
	fd_SignatureDescriptors_signatures = md_SignatureDescriptors.Fields().ByName("signatures")
}

var _ protoreflect.Message = (*fastReflection_SignatureDescriptors)(nil)

type fastReflection_SignatureDescriptors SignatureDescriptors

func (x *SignatureDescriptors) ProtoReflect() protoreflect.Message {
	return (*fastReflection_SignatureDescriptors)(x)
}

func (x *SignatureDescriptors) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_tx_signing_v1beta1_signing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_SignatureDescriptors_messageType fastReflection_SignatureDescriptors_messageType
var _ protoreflect.MessageType = fastReflection_SignatureDescriptors_messageType{}

type fastReflection_SignatureDescriptors_messageType struct{}

func (x fastReflection_SignatureDescriptors_messageType) Zero() protoreflect.Message {
	return (*fastReflection_SignatureDescriptors)(nil)
}
func (x fastReflection_SignatureDescriptors_messageType) New() protoreflect.Message {
	return new(fastReflection_SignatureDescriptors)
}
func (x fastReflection_SignatureDescriptors_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_SignatureDescriptors
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_SignatureDescriptors) Descriptor() protoreflect.MessageDescriptor {
	return md_SignatureDescriptors
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_SignatureDescriptors) Type() protoreflect.MessageType {
	return _fastReflection_SignatureDescriptors_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_SignatureDescriptors) New() protoreflect.Message {
	return new(fastReflection_SignatureDescriptors)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_SignatureDescriptors) Interface() protoreflect.ProtoMessage {
	return (*SignatureDescriptors)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_SignatureDescriptors) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Signatures) != 0 {
		value := protoreflect.ValueOfList(&_SignatureDescriptors_1_list{list: &x.Signatures})
		if !f(fd_SignatureDescriptors_signatures, value) {
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
func (x *fastReflection_SignatureDescriptors) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.tx.signing.v1beta1.SignatureDescriptors.signatures":
		return len(x.Signatures) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.signing.v1beta1.SignatureDescriptors"))
		}
		panic(fmt.Errorf("message cosmos.tx.signing.v1beta1.SignatureDescriptors does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SignatureDescriptors) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.tx.signing.v1beta1.SignatureDescriptors.signatures":
		x.Signatures = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.signing.v1beta1.SignatureDescriptors"))
		}
		panic(fmt.Errorf("message cosmos.tx.signing.v1beta1.SignatureDescriptors does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_SignatureDescriptors) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.tx.signing.v1beta1.SignatureDescriptors.signatures":
		if len(x.Signatures) == 0 {
			return protoreflect.ValueOfList(&_SignatureDescriptors_1_list{})
		}
		listValue := &_SignatureDescriptors_1_list{list: &x.Signatures}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.signing.v1beta1.SignatureDescriptors"))
		}
		panic(fmt.Errorf("message cosmos.tx.signing.v1beta1.SignatureDescriptors does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_SignatureDescriptors) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.tx.signing.v1beta1.SignatureDescriptors.signatures":
		lv := value.List()
		clv := lv.(*_SignatureDescriptors_1_list)
		x.Signatures = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.signing.v1beta1.SignatureDescriptors"))
		}
		panic(fmt.Errorf("message cosmos.tx.signing.v1beta1.SignatureDescriptors does not contain field %s", fd.FullName()))
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
func (x *fastReflection_SignatureDescriptors) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.signing.v1beta1.SignatureDescriptors.signatures":
		if x.Signatures == nil {
			x.Signatures = []*SignatureDescriptor{}
		}
		value := &_SignatureDescriptors_1_list{list: &x.Signatures}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.signing.v1beta1.SignatureDescriptors"))
		}
		panic(fmt.Errorf("message cosmos.tx.signing.v1beta1.SignatureDescriptors does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_SignatureDescriptors) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.signing.v1beta1.SignatureDescriptors.signatures":
		list := []*SignatureDescriptor{}
		return protoreflect.ValueOfList(&_SignatureDescriptors_1_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.signing.v1beta1.SignatureDescriptors"))
		}
		panic(fmt.Errorf("message cosmos.tx.signing.v1beta1.SignatureDescriptors does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_SignatureDescriptors) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.tx.signing.v1beta1.SignatureDescriptors", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_SignatureDescriptors) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SignatureDescriptors) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_SignatureDescriptors) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_SignatureDescriptors) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*SignatureDescriptors)
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
		if len(x.Signatures) > 0 {
			for _, e := range x.Signatures {
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
		x := input.Message.Interface().(*SignatureDescriptors)
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
		if len(x.Signatures) > 0 {
			for iNdEx := len(x.Signatures) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Signatures[iNdEx])
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
		x := input.Message.Interface().(*SignatureDescriptors)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SignatureDescriptors: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SignatureDescriptors: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Signatures", wireType)
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
				x.Signatures = append(x.Signatures, &SignatureDescriptor{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Signatures[len(x.Signatures)-1]); err != nil {
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
	md_SignatureDescriptor            protoreflect.MessageDescriptor
	fd_SignatureDescriptor_public_key protoreflect.FieldDescriptor
	fd_SignatureDescriptor_data       protoreflect.FieldDescriptor
	fd_SignatureDescriptor_sequence   protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_tx_signing_v1beta1_signing_proto_init()
	md_SignatureDescriptor = File_cosmos_tx_signing_v1beta1_signing_proto.Messages().ByName("SignatureDescriptor")
	fd_SignatureDescriptor_public_key = md_SignatureDescriptor.Fields().ByName("public_key")
	fd_SignatureDescriptor_data = md_SignatureDescriptor.Fields().ByName("data")
	fd_SignatureDescriptor_sequence = md_SignatureDescriptor.Fields().ByName("sequence")
}

var _ protoreflect.Message = (*fastReflection_SignatureDescriptor)(nil)

type fastReflection_SignatureDescriptor SignatureDescriptor

func (x *SignatureDescriptor) ProtoReflect() protoreflect.Message {
	return (*fastReflection_SignatureDescriptor)(x)
}

func (x *SignatureDescriptor) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_tx_signing_v1beta1_signing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_SignatureDescriptor_messageType fastReflection_SignatureDescriptor_messageType
var _ protoreflect.MessageType = fastReflection_SignatureDescriptor_messageType{}

type fastReflection_SignatureDescriptor_messageType struct{}

func (x fastReflection_SignatureDescriptor_messageType) Zero() protoreflect.Message {
	return (*fastReflection_SignatureDescriptor)(nil)
}
func (x fastReflection_SignatureDescriptor_messageType) New() protoreflect.Message {
	return new(fastReflection_SignatureDescriptor)
}
func (x fastReflection_SignatureDescriptor_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_SignatureDescriptor
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_SignatureDescriptor) Descriptor() protoreflect.MessageDescriptor {
	return md_SignatureDescriptor
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_SignatureDescriptor) Type() protoreflect.MessageType {
	return _fastReflection_SignatureDescriptor_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_SignatureDescriptor) New() protoreflect.Message {
	return new(fastReflection_SignatureDescriptor)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_SignatureDescriptor) Interface() protoreflect.ProtoMessage {
	return (*SignatureDescriptor)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_SignatureDescriptor) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.PublicKey != nil {
		value := protoreflect.ValueOfMessage(x.PublicKey.ProtoReflect())
		if !f(fd_SignatureDescriptor_public_key, value) {
			return
		}
	}
	if x.Data != nil {
		value := protoreflect.ValueOfMessage(x.Data.ProtoReflect())
		if !f(fd_SignatureDescriptor_data, value) {
			return
		}
	}
	if x.Sequence != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Sequence)
		if !f(fd_SignatureDescriptor_sequence, value) {
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
func (x *fastReflection_SignatureDescriptor) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.public_key":
		return x.PublicKey != nil
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.data":
		return x.Data != nil
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.sequence":
		return x.Sequence != uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.signing.v1beta1.SignatureDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.tx.signing.v1beta1.SignatureDescriptor does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SignatureDescriptor) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.public_key":
		x.PublicKey = nil
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.data":
		x.Data = nil
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.sequence":
		x.Sequence = uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.signing.v1beta1.SignatureDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.tx.signing.v1beta1.SignatureDescriptor does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_SignatureDescriptor) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.public_key":
		value := x.PublicKey
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.data":
		value := x.Data
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.sequence":
		value := x.Sequence
		return protoreflect.ValueOfUint64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.signing.v1beta1.SignatureDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.tx.signing.v1beta1.SignatureDescriptor does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_SignatureDescriptor) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.public_key":
		x.PublicKey = value.Message().Interface().(*anypb.Any)
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.data":
		x.Data = value.Message().Interface().(*SignatureDescriptor_Data)
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.sequence":
		x.Sequence = value.Uint()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.signing.v1beta1.SignatureDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.tx.signing.v1beta1.SignatureDescriptor does not contain field %s", fd.FullName()))
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
func (x *fastReflection_SignatureDescriptor) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.public_key":
		if x.PublicKey == nil {
			x.PublicKey = new(anypb.Any)
		}
		return protoreflect.ValueOfMessage(x.PublicKey.ProtoReflect())
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.data":
		if x.Data == nil {
			x.Data = new(SignatureDescriptor_Data)
		}
		return protoreflect.ValueOfMessage(x.Data.ProtoReflect())
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.sequence":
		panic(fmt.Errorf("field sequence of message cosmos.tx.signing.v1beta1.SignatureDescriptor is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.signing.v1beta1.SignatureDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.tx.signing.v1beta1.SignatureDescriptor does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_SignatureDescriptor) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.public_key":
		m := new(anypb.Any)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.data":
		m := new(SignatureDescriptor_Data)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.sequence":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.signing.v1beta1.SignatureDescriptor"))
		}
		panic(fmt.Errorf("message cosmos.tx.signing.v1beta1.SignatureDescriptor does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_SignatureDescriptor) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.tx.signing.v1beta1.SignatureDescriptor", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_SignatureDescriptor) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SignatureDescriptor) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_SignatureDescriptor) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_SignatureDescriptor) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*SignatureDescriptor)
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
		if x.PublicKey != nil {
			l = options.Size(x.PublicKey)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Data != nil {
			l = options.Size(x.Data)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Sequence != 0 {
			n += 1 + runtime.Sov(uint64(x.Sequence))
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
		x := input.Message.Interface().(*SignatureDescriptor)
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
		if x.Sequence != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Sequence))
			i--
			dAtA[i] = 0x18
		}
		if x.Data != nil {
			encoded, err := options.Marshal(x.Data)
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
		if x.PublicKey != nil {
			encoded, err := options.Marshal(x.PublicKey)
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
		x := input.Message.Interface().(*SignatureDescriptor)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SignatureDescriptor: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SignatureDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
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
				if x.PublicKey == nil {
					x.PublicKey = &anypb.Any{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.PublicKey); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
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
				if x.Data == nil {
					x.Data = &SignatureDescriptor_Data{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Data); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
				}
				x.Sequence = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Sequence |= uint64(b&0x7F) << shift
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
	md_SignatureDescriptor_Data        protoreflect.MessageDescriptor
	fd_SignatureDescriptor_Data_single protoreflect.FieldDescriptor
	fd_SignatureDescriptor_Data_multi  protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_tx_signing_v1beta1_signing_proto_init()
	md_SignatureDescriptor_Data = File_cosmos_tx_signing_v1beta1_signing_proto.Messages().ByName("SignatureDescriptor").Messages().ByName("Data")
	fd_SignatureDescriptor_Data_single = md_SignatureDescriptor_Data.Fields().ByName("single")
	fd_SignatureDescriptor_Data_multi = md_SignatureDescriptor_Data.Fields().ByName("multi")
}

var _ protoreflect.Message = (*fastReflection_SignatureDescriptor_Data)(nil)

type fastReflection_SignatureDescriptor_Data SignatureDescriptor_Data

func (x *SignatureDescriptor_Data) ProtoReflect() protoreflect.Message {
	return (*fastReflection_SignatureDescriptor_Data)(x)
}

func (x *SignatureDescriptor_Data) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_tx_signing_v1beta1_signing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_SignatureDescriptor_Data_messageType fastReflection_SignatureDescriptor_Data_messageType
var _ protoreflect.MessageType = fastReflection_SignatureDescriptor_Data_messageType{}

type fastReflection_SignatureDescriptor_Data_messageType struct{}

func (x fastReflection_SignatureDescriptor_Data_messageType) Zero() protoreflect.Message {
	return (*fastReflection_SignatureDescriptor_Data)(nil)
}
func (x fastReflection_SignatureDescriptor_Data_messageType) New() protoreflect.Message {
	return new(fastReflection_SignatureDescriptor_Data)
}
func (x fastReflection_SignatureDescriptor_Data_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_SignatureDescriptor_Data
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_SignatureDescriptor_Data) Descriptor() protoreflect.MessageDescriptor {
	return md_SignatureDescriptor_Data
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_SignatureDescriptor_Data) Type() protoreflect.MessageType {
	return _fastReflection_SignatureDescriptor_Data_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_SignatureDescriptor_Data) New() protoreflect.Message {
	return new(fastReflection_SignatureDescriptor_Data)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_SignatureDescriptor_Data) Interface() protoreflect.ProtoMessage {
	return (*SignatureDescriptor_Data)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_SignatureDescriptor_Data) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Sum != nil {
		switch o := x.Sum.(type) {
		case *SignatureDescriptor_Data_Single_:
			v := o.Single
			value := protoreflect.ValueOfMessage(v.ProtoReflect())
			if !f(fd_SignatureDescriptor_Data_single, value) {
				return
			}
		case *SignatureDescriptor_Data_Multi_:
			v := o.Multi
			value := protoreflect.ValueOfMessage(v.ProtoReflect())
			if !f(fd_SignatureDescriptor_Data_multi, value) {
				return
			}
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
func (x *fastReflection_SignatureDescriptor_Data) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.single":
		if x.Sum == nil {
			return false
		} else if _, ok := x.Sum.(*SignatureDescriptor_Data_Single_); ok {
			return true
		} else {
			return false
		}
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.multi":
		if x.Sum == nil {
			return false
		} else if _, ok := x.Sum.(*SignatureDescriptor_Data_Multi_); ok {
			return true
		} else {
			return false
		}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.signing.v1beta1.SignatureDescriptor.Data"))
		}
		panic(fmt.Errorf("message cosmos.tx.signing.v1beta1.SignatureDescriptor.Data does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SignatureDescriptor_Data) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.single":
		x.Sum = nil
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.multi":
		x.Sum = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.signing.v1beta1.SignatureDescriptor.Data"))
		}
		panic(fmt.Errorf("message cosmos.tx.signing.v1beta1.SignatureDescriptor.Data does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_SignatureDescriptor_Data) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.single":
		if x.Sum == nil {
			return protoreflect.ValueOfMessage((*SignatureDescriptor_Data_Single)(nil).ProtoReflect())
		} else if v, ok := x.Sum.(*SignatureDescriptor_Data_Single_); ok {
			return protoreflect.ValueOfMessage(v.Single.ProtoReflect())
		} else {
			return protoreflect.ValueOfMessage((*SignatureDescriptor_Data_Single)(nil).ProtoReflect())
		}
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.multi":
		if x.Sum == nil {
			return protoreflect.ValueOfMessage((*SignatureDescriptor_Data_Multi)(nil).ProtoReflect())
		} else if v, ok := x.Sum.(*SignatureDescriptor_Data_Multi_); ok {
			return protoreflect.ValueOfMessage(v.Multi.ProtoReflect())
		} else {
			return protoreflect.ValueOfMessage((*SignatureDescriptor_Data_Multi)(nil).ProtoReflect())
		}
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.signing.v1beta1.SignatureDescriptor.Data"))
		}
		panic(fmt.Errorf("message cosmos.tx.signing.v1beta1.SignatureDescriptor.Data does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_SignatureDescriptor_Data) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.single":
		cv := value.Message().Interface().(*SignatureDescriptor_Data_Single)
		x.Sum = &SignatureDescriptor_Data_Single_{Single: cv}
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.multi":
		cv := value.Message().Interface().(*SignatureDescriptor_Data_Multi)
		x.Sum = &SignatureDescriptor_Data_Multi_{Multi: cv}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.signing.v1beta1.SignatureDescriptor.Data"))
		}
		panic(fmt.Errorf("message cosmos.tx.signing.v1beta1.SignatureDescriptor.Data does not contain field %s", fd.FullName()))
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
func (x *fastReflection_SignatureDescriptor_Data) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.single":
		if x.Sum == nil {
			value := &SignatureDescriptor_Data_Single{}
			oneofValue := &SignatureDescriptor_Data_Single_{Single: value}
			x.Sum = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
		switch m := x.Sum.(type) {
		case *SignatureDescriptor_Data_Single_:
			return protoreflect.ValueOfMessage(m.Single.ProtoReflect())
		default:
			value := &SignatureDescriptor_Data_Single{}
			oneofValue := &SignatureDescriptor_Data_Single_{Single: value}
			x.Sum = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.multi":
		if x.Sum == nil {
			value := &SignatureDescriptor_Data_Multi{}
			oneofValue := &SignatureDescriptor_Data_Multi_{Multi: value}
			x.Sum = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
		switch m := x.Sum.(type) {
		case *SignatureDescriptor_Data_Multi_:
			return protoreflect.ValueOfMessage(m.Multi.ProtoReflect())
		default:
			value := &SignatureDescriptor_Data_Multi{}
			oneofValue := &SignatureDescriptor_Data_Multi_{Multi: value}
			x.Sum = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.signing.v1beta1.SignatureDescriptor.Data"))
		}
		panic(fmt.Errorf("message cosmos.tx.signing.v1beta1.SignatureDescriptor.Data does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_SignatureDescriptor_Data) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.single":
		value := &SignatureDescriptor_Data_Single{}
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.multi":
		value := &SignatureDescriptor_Data_Multi{}
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.signing.v1beta1.SignatureDescriptor.Data"))
		}
		panic(fmt.Errorf("message cosmos.tx.signing.v1beta1.SignatureDescriptor.Data does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_SignatureDescriptor_Data) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.sum":
		if x.Sum == nil {
			return nil
		}
		switch x.Sum.(type) {
		case *SignatureDescriptor_Data_Single_:
			return x.Descriptor().Fields().ByName("single")
		case *SignatureDescriptor_Data_Multi_:
			return x.Descriptor().Fields().ByName("multi")
		}
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.tx.signing.v1beta1.SignatureDescriptor.Data", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_SignatureDescriptor_Data) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SignatureDescriptor_Data) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_SignatureDescriptor_Data) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_SignatureDescriptor_Data) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*SignatureDescriptor_Data)
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
		switch x := x.Sum.(type) {
		case *SignatureDescriptor_Data_Single_:
			if x == nil {
				break
			}
			l = options.Size(x.Single)
			n += 1 + l + runtime.Sov(uint64(l))
		case *SignatureDescriptor_Data_Multi_:
			if x == nil {
				break
			}
			l = options.Size(x.Multi)
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
		x := input.Message.Interface().(*SignatureDescriptor_Data)
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
		switch x := x.Sum.(type) {
		case *SignatureDescriptor_Data_Single_:
			encoded, err := options.Marshal(x.Single)
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
		case *SignatureDescriptor_Data_Multi_:
			encoded, err := options.Marshal(x.Multi)
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
		x := input.Message.Interface().(*SignatureDescriptor_Data)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SignatureDescriptor_Data: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SignatureDescriptor_Data: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Single", wireType)
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
				v := &SignatureDescriptor_Data_Single{}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], v); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				x.Sum = &SignatureDescriptor_Data_Single_{v}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Multi", wireType)
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
				v := &SignatureDescriptor_Data_Multi{}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], v); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				x.Sum = &SignatureDescriptor_Data_Multi_{v}
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
	md_SignatureDescriptor_Data_Single           protoreflect.MessageDescriptor
	fd_SignatureDescriptor_Data_Single_mode      protoreflect.FieldDescriptor
	fd_SignatureDescriptor_Data_Single_signature protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_tx_signing_v1beta1_signing_proto_init()
	md_SignatureDescriptor_Data_Single = File_cosmos_tx_signing_v1beta1_signing_proto.Messages().ByName("SignatureDescriptor").Messages().ByName("Data").Messages().ByName("Single")
	fd_SignatureDescriptor_Data_Single_mode = md_SignatureDescriptor_Data_Single.Fields().ByName("mode")
	fd_SignatureDescriptor_Data_Single_signature = md_SignatureDescriptor_Data_Single.Fields().ByName("signature")
}

var _ protoreflect.Message = (*fastReflection_SignatureDescriptor_Data_Single)(nil)

type fastReflection_SignatureDescriptor_Data_Single SignatureDescriptor_Data_Single

func (x *SignatureDescriptor_Data_Single) ProtoReflect() protoreflect.Message {
	return (*fastReflection_SignatureDescriptor_Data_Single)(x)
}

func (x *SignatureDescriptor_Data_Single) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_tx_signing_v1beta1_signing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_SignatureDescriptor_Data_Single_messageType fastReflection_SignatureDescriptor_Data_Single_messageType
var _ protoreflect.MessageType = fastReflection_SignatureDescriptor_Data_Single_messageType{}

type fastReflection_SignatureDescriptor_Data_Single_messageType struct{}

func (x fastReflection_SignatureDescriptor_Data_Single_messageType) Zero() protoreflect.Message {
	return (*fastReflection_SignatureDescriptor_Data_Single)(nil)
}
func (x fastReflection_SignatureDescriptor_Data_Single_messageType) New() protoreflect.Message {
	return new(fastReflection_SignatureDescriptor_Data_Single)
}
func (x fastReflection_SignatureDescriptor_Data_Single_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_SignatureDescriptor_Data_Single
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_SignatureDescriptor_Data_Single) Descriptor() protoreflect.MessageDescriptor {
	return md_SignatureDescriptor_Data_Single
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_SignatureDescriptor_Data_Single) Type() protoreflect.MessageType {
	return _fastReflection_SignatureDescriptor_Data_Single_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_SignatureDescriptor_Data_Single) New() protoreflect.Message {
	return new(fastReflection_SignatureDescriptor_Data_Single)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_SignatureDescriptor_Data_Single) Interface() protoreflect.ProtoMessage {
	return (*SignatureDescriptor_Data_Single)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_SignatureDescriptor_Data_Single) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Mode != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.Mode))
		if !f(fd_SignatureDescriptor_Data_Single_mode, value) {
			return
		}
	}
	if len(x.Signature) != 0 {
		value := protoreflect.ValueOfBytes(x.Signature)
		if !f(fd_SignatureDescriptor_Data_Single_signature, value) {
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
func (x *fastReflection_SignatureDescriptor_Data_Single) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Single.mode":
		return x.Mode != 0
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Single.signature":
		return len(x.Signature) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Single"))
		}
		panic(fmt.Errorf("message cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Single does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SignatureDescriptor_Data_Single) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Single.mode":
		x.Mode = 0
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Single.signature":
		x.Signature = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Single"))
		}
		panic(fmt.Errorf("message cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Single does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_SignatureDescriptor_Data_Single) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Single.mode":
		value := x.Mode
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Single.signature":
		value := x.Signature
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Single"))
		}
		panic(fmt.Errorf("message cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Single does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_SignatureDescriptor_Data_Single) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Single.mode":
		x.Mode = (SignMode)(value.Enum())
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Single.signature":
		x.Signature = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Single"))
		}
		panic(fmt.Errorf("message cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Single does not contain field %s", fd.FullName()))
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
func (x *fastReflection_SignatureDescriptor_Data_Single) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Single.mode":
		panic(fmt.Errorf("field mode of message cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Single is not mutable"))
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Single.signature":
		panic(fmt.Errorf("field signature of message cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Single is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Single"))
		}
		panic(fmt.Errorf("message cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Single does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_SignatureDescriptor_Data_Single) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Single.mode":
		return protoreflect.ValueOfEnum(0)
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Single.signature":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Single"))
		}
		panic(fmt.Errorf("message cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Single does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_SignatureDescriptor_Data_Single) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Single", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_SignatureDescriptor_Data_Single) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SignatureDescriptor_Data_Single) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_SignatureDescriptor_Data_Single) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_SignatureDescriptor_Data_Single) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*SignatureDescriptor_Data_Single)
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
		if x.Mode != 0 {
			n += 1 + runtime.Sov(uint64(x.Mode))
		}
		l = len(x.Signature)
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
		x := input.Message.Interface().(*SignatureDescriptor_Data_Single)
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
		if len(x.Signature) > 0 {
			i -= len(x.Signature)
			copy(dAtA[i:], x.Signature)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Signature)))
			i--
			dAtA[i] = 0x12
		}
		if x.Mode != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Mode))
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
		x := input.Message.Interface().(*SignatureDescriptor_Data_Single)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SignatureDescriptor_Data_Single: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SignatureDescriptor_Data_Single: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Mode", wireType)
				}
				x.Mode = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Mode |= SignMode(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
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
				x.Signature = append(x.Signature[:0], dAtA[iNdEx:postIndex]...)
				if x.Signature == nil {
					x.Signature = []byte{}
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

var _ protoreflect.List = (*_SignatureDescriptor_Data_Multi_2_list)(nil)

type _SignatureDescriptor_Data_Multi_2_list struct {
	list *[]*SignatureDescriptor_Data
}

func (x *_SignatureDescriptor_Data_Multi_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_SignatureDescriptor_Data_Multi_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_SignatureDescriptor_Data_Multi_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*SignatureDescriptor_Data)
	(*x.list)[i] = concreteValue
}

func (x *_SignatureDescriptor_Data_Multi_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*SignatureDescriptor_Data)
	*x.list = append(*x.list, concreteValue)
}

func (x *_SignatureDescriptor_Data_Multi_2_list) AppendMutable() protoreflect.Value {
	v := new(SignatureDescriptor_Data)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_SignatureDescriptor_Data_Multi_2_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_SignatureDescriptor_Data_Multi_2_list) NewElement() protoreflect.Value {
	v := new(SignatureDescriptor_Data)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_SignatureDescriptor_Data_Multi_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_SignatureDescriptor_Data_Multi            protoreflect.MessageDescriptor
	fd_SignatureDescriptor_Data_Multi_bitarray   protoreflect.FieldDescriptor
	fd_SignatureDescriptor_Data_Multi_signatures protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_tx_signing_v1beta1_signing_proto_init()
	md_SignatureDescriptor_Data_Multi = File_cosmos_tx_signing_v1beta1_signing_proto.Messages().ByName("SignatureDescriptor").Messages().ByName("Data").Messages().ByName("Multi")
	fd_SignatureDescriptor_Data_Multi_bitarray = md_SignatureDescriptor_Data_Multi.Fields().ByName("bitarray")
	fd_SignatureDescriptor_Data_Multi_signatures = md_SignatureDescriptor_Data_Multi.Fields().ByName("signatures")
}

var _ protoreflect.Message = (*fastReflection_SignatureDescriptor_Data_Multi)(nil)

type fastReflection_SignatureDescriptor_Data_Multi SignatureDescriptor_Data_Multi

func (x *SignatureDescriptor_Data_Multi) ProtoReflect() protoreflect.Message {
	return (*fastReflection_SignatureDescriptor_Data_Multi)(x)
}

func (x *SignatureDescriptor_Data_Multi) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_tx_signing_v1beta1_signing_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_SignatureDescriptor_Data_Multi_messageType fastReflection_SignatureDescriptor_Data_Multi_messageType
var _ protoreflect.MessageType = fastReflection_SignatureDescriptor_Data_Multi_messageType{}

type fastReflection_SignatureDescriptor_Data_Multi_messageType struct{}

func (x fastReflection_SignatureDescriptor_Data_Multi_messageType) Zero() protoreflect.Message {
	return (*fastReflection_SignatureDescriptor_Data_Multi)(nil)
}
func (x fastReflection_SignatureDescriptor_Data_Multi_messageType) New() protoreflect.Message {
	return new(fastReflection_SignatureDescriptor_Data_Multi)
}
func (x fastReflection_SignatureDescriptor_Data_Multi_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_SignatureDescriptor_Data_Multi
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_SignatureDescriptor_Data_Multi) Descriptor() protoreflect.MessageDescriptor {
	return md_SignatureDescriptor_Data_Multi
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_SignatureDescriptor_Data_Multi) Type() protoreflect.MessageType {
	return _fastReflection_SignatureDescriptor_Data_Multi_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_SignatureDescriptor_Data_Multi) New() protoreflect.Message {
	return new(fastReflection_SignatureDescriptor_Data_Multi)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_SignatureDescriptor_Data_Multi) Interface() protoreflect.ProtoMessage {
	return (*SignatureDescriptor_Data_Multi)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_SignatureDescriptor_Data_Multi) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Bitarray != nil {
		value := protoreflect.ValueOfMessage(x.Bitarray.ProtoReflect())
		if !f(fd_SignatureDescriptor_Data_Multi_bitarray, value) {
			return
		}
	}
	if len(x.Signatures) != 0 {
		value := protoreflect.ValueOfList(&_SignatureDescriptor_Data_Multi_2_list{list: &x.Signatures})
		if !f(fd_SignatureDescriptor_Data_Multi_signatures, value) {
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
func (x *fastReflection_SignatureDescriptor_Data_Multi) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Multi.bitarray":
		return x.Bitarray != nil
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Multi.signatures":
		return len(x.Signatures) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Multi"))
		}
		panic(fmt.Errorf("message cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Multi does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SignatureDescriptor_Data_Multi) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Multi.bitarray":
		x.Bitarray = nil
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Multi.signatures":
		x.Signatures = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Multi"))
		}
		panic(fmt.Errorf("message cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Multi does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_SignatureDescriptor_Data_Multi) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Multi.bitarray":
		value := x.Bitarray
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Multi.signatures":
		if len(x.Signatures) == 0 {
			return protoreflect.ValueOfList(&_SignatureDescriptor_Data_Multi_2_list{})
		}
		listValue := &_SignatureDescriptor_Data_Multi_2_list{list: &x.Signatures}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Multi"))
		}
		panic(fmt.Errorf("message cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Multi does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_SignatureDescriptor_Data_Multi) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Multi.bitarray":
		x.Bitarray = value.Message().Interface().(*v1beta1.CompactBitArray)
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Multi.signatures":
		lv := value.List()
		clv := lv.(*_SignatureDescriptor_Data_Multi_2_list)
		x.Signatures = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Multi"))
		}
		panic(fmt.Errorf("message cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Multi does not contain field %s", fd.FullName()))
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
func (x *fastReflection_SignatureDescriptor_Data_Multi) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Multi.bitarray":
		if x.Bitarray == nil {
			x.Bitarray = new(v1beta1.CompactBitArray)
		}
		return protoreflect.ValueOfMessage(x.Bitarray.ProtoReflect())
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Multi.signatures":
		if x.Signatures == nil {
			x.Signatures = []*SignatureDescriptor_Data{}
		}
		value := &_SignatureDescriptor_Data_Multi_2_list{list: &x.Signatures}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Multi"))
		}
		panic(fmt.Errorf("message cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Multi does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_SignatureDescriptor_Data_Multi) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Multi.bitarray":
		m := new(v1beta1.CompactBitArray)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Multi.signatures":
		list := []*SignatureDescriptor_Data{}
		return protoreflect.ValueOfList(&_SignatureDescriptor_Data_Multi_2_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Multi"))
		}
		panic(fmt.Errorf("message cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Multi does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_SignatureDescriptor_Data_Multi) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Multi", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_SignatureDescriptor_Data_Multi) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SignatureDescriptor_Data_Multi) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_SignatureDescriptor_Data_Multi) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_SignatureDescriptor_Data_Multi) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*SignatureDescriptor_Data_Multi)
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
		if x.Bitarray != nil {
			l = options.Size(x.Bitarray)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Signatures) > 0 {
			for _, e := range x.Signatures {
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
		x := input.Message.Interface().(*SignatureDescriptor_Data_Multi)
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
		if len(x.Signatures) > 0 {
			for iNdEx := len(x.Signatures) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Signatures[iNdEx])
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
		if x.Bitarray != nil {
			encoded, err := options.Marshal(x.Bitarray)
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
		x := input.Message.Interface().(*SignatureDescriptor_Data_Multi)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SignatureDescriptor_Data_Multi: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SignatureDescriptor_Data_Multi: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Bitarray", wireType)
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
				if x.Bitarray == nil {
					x.Bitarray = &v1beta1.CompactBitArray{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Bitarray); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Signatures", wireType)
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
				x.Signatures = append(x.Signatures, &SignatureDescriptor_Data{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Signatures[len(x.Signatures)-1]); err != nil {
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: cosmos/tx/signing/v1beta1/signing.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// SignMode represents a signing mode with its own security guarantees.
type SignMode int32

const (
	// SIGN_MODE_UNSPECIFIED specifies an unknown signing mode and will be
	// rejected.
	SignMode_SIGN_MODE_UNSPECIFIED SignMode = 0
	// SIGN_MODE_DIRECT specifies a signing mode which uses SignDoc and is
	// verified with raw bytes from Tx.
	SignMode_SIGN_MODE_DIRECT SignMode = 1
	// SIGN_MODE_TEXTUAL is a future signing mode that will verify some
	// human-readable textual representation on top of the binary representation
	// from SIGN_MODE_DIRECT. It is currently not supported.
	SignMode_SIGN_MODE_TEXTUAL SignMode = 2
	// SIGN_MODE_DIRECT_AUX specifies a signing mode which uses
	// SignDocDirectAux. As opposed to SIGN_MODE_DIRECT, this sign mode does not
	// require signers signing over other signers' `signer_info`. It also allows
	// for adding Tips in transactions.
	SignMode_SIGN_MODE_DIRECT_AUX SignMode = 3
	// SIGN_MODE_LEGACY_AMINO_JSON is a backwards compatibility mode which uses
	// Amino JSON and will be removed in the future.
	SignMode_SIGN_MODE_LEGACY_AMINO_JSON SignMode = 127
)

// Enum value maps for SignMode.
var (
	SignMode_name = map[int32]string{
		0:   "SIGN_MODE_UNSPECIFIED",
		1:   "SIGN_MODE_DIRECT",
		2:   "SIGN_MODE_TEXTUAL",
		3:   "SIGN_MODE_DIRECT_AUX",
		127: "SIGN_MODE_LEGACY_AMINO_JSON",
	}
	SignMode_value = map[string]int32{
		"SIGN_MODE_UNSPECIFIED":       0,
		"SIGN_MODE_DIRECT":            1,
		"SIGN_MODE_TEXTUAL":           2,
		"SIGN_MODE_DIRECT_AUX":        3,
		"SIGN_MODE_LEGACY_AMINO_JSON": 127,
	}
)

func (x SignMode) Enum() *SignMode {
	p := new(SignMode)
	*p = x
	return p
}

func (x SignMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SignMode) Descriptor() protoreflect.EnumDescriptor {
	return file_cosmos_tx_signing_v1beta1_signing_proto_enumTypes[0].Descriptor()
}

func (SignMode) Type() protoreflect.EnumType {
	return &file_cosmos_tx_signing_v1beta1_signing_proto_enumTypes[0]
}

func (x SignMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SignMode.Descriptor instead.
func (SignMode) EnumDescriptor() ([]byte, []int) {
	return file_cosmos_tx_signing_v1beta1_signing_proto_rawDescGZIP(), []int{0}
}

// SignatureDescriptors wraps multiple SignatureDescriptor's.
type SignatureDescriptors struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// signatures are the signature descriptors
	Signatures []*SignatureDescriptor `protobuf:"bytes,1,rep,name=signatures,proto3" json:"signatures,omitempty"`
}

func (x *SignatureDescriptors) Reset() {
	*x = SignatureDescriptors{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_tx_signing_v1beta1_signing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignatureDescriptors) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignatureDescriptors) ProtoMessage() {}

// Deprecated: Use SignatureDescriptors.ProtoReflect.Descriptor instead.
func (*SignatureDescriptors) Descriptor() ([]byte, []int) {
	return file_cosmos_tx_signing_v1beta1_signing_proto_rawDescGZIP(), []int{0}
}

func (x *SignatureDescriptors) GetSignatures() []*SignatureDescriptor {
	if x != nil {
		return x.Signatures
	}
	return nil
}

// SignatureDescriptor is a convenience type which represents the full data for
// a signature including the public key of the signer, signing modes and the
// signature itself. It is primarily used for coordinating signatures between
// clients.
type SignatureDescriptor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// public_key is the public key of the signer
	PublicKey *anypb.Any                `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Data      *SignatureDescriptor_Data `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// sequence is the sequence of the account, which describes the
	// number of committed transactions signed by a given address. It is used to prevent
	// replay attacks.
	Sequence uint64 `protobuf:"varint,3,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (x *SignatureDescriptor) Reset() {
	*x = SignatureDescriptor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_tx_signing_v1beta1_signing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignatureDescriptor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignatureDescriptor) ProtoMessage() {}

// Deprecated: Use SignatureDescriptor.ProtoReflect.Descriptor instead.
func (*SignatureDescriptor) Descriptor() ([]byte, []int) {
	return file_cosmos_tx_signing_v1beta1_signing_proto_rawDescGZIP(), []int{1}
}

func (x *SignatureDescriptor) GetPublicKey() *anypb.Any {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *SignatureDescriptor) GetData() *SignatureDescriptor_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SignatureDescriptor) GetSequence() uint64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

// Data represents signature data
type SignatureDescriptor_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// sum is the oneof that specifies whether this represents single or multi-signature data
	//
	// Types that are assignable to Sum:
	//	*SignatureDescriptor_Data_Single_
	//	*SignatureDescriptor_Data_Multi_
	Sum isSignatureDescriptor_Data_Sum `protobuf_oneof:"sum"`
}

func (x *SignatureDescriptor_Data) Reset() {
	*x = SignatureDescriptor_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_tx_signing_v1beta1_signing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignatureDescriptor_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignatureDescriptor_Data) ProtoMessage() {}

// Deprecated: Use SignatureDescriptor_Data.ProtoReflect.Descriptor instead.
func (*SignatureDescriptor_Data) Descriptor() ([]byte, []int) {
	return file_cosmos_tx_signing_v1beta1_signing_proto_rawDescGZIP(), []int{1, 0}
}

func (x *SignatureDescriptor_Data) GetSum() isSignatureDescriptor_Data_Sum {
	if x != nil {
		return x.Sum
	}
	return nil
}

func (x *SignatureDescriptor_Data) GetSingle() *SignatureDescriptor_Data_Single {
	if x, ok := x.GetSum().(*SignatureDescriptor_Data_Single_); ok {
		return x.Single
	}
	return nil
}

func (x *SignatureDescriptor_Data) GetMulti() *SignatureDescriptor_Data_Multi {
	if x, ok := x.GetSum().(*SignatureDescriptor_Data_Multi_); ok {
		return x.Multi
	}
	return nil
}

type isSignatureDescriptor_Data_Sum interface {
	isSignatureDescriptor_Data_Sum()
}

type SignatureDescriptor_Data_Single_ struct {
	// single represents a single signer
	Single *SignatureDescriptor_Data_Single `protobuf:"bytes,1,opt,name=single,proto3,oneof"`
}

type SignatureDescriptor_Data_Multi_ struct {
	// multi represents a multisig signer
	Multi *SignatureDescriptor_Data_Multi `protobuf:"bytes,2,opt,name=multi,proto3,oneof"`
}

func (*SignatureDescriptor_Data_Single_) isSignatureDescriptor_Data_Sum() {}

func (*SignatureDescriptor_Data_Multi_) isSignatureDescriptor_Data_Sum() {}

// Single is the signature data for a single signer
type SignatureDescriptor_Data_Single struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// mode is the signing mode of the single signer
	Mode SignMode `protobuf:"varint,1,opt,name=mode,proto3,enum=cosmos.tx.signing.v1beta1.SignMode" json:"mode,omitempty"`
	// signature is the raw signature bytes
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *SignatureDescriptor_Data_Single) Reset() {
	*x = SignatureDescriptor_Data_Single{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_tx_signing_v1beta1_signing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignatureDescriptor_Data_Single) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignatureDescriptor_Data_Single) ProtoMessage() {}

// Deprecated: Use SignatureDescriptor_Data_Single.ProtoReflect.Descriptor instead.
func (*SignatureDescriptor_Data_Single) Descriptor() ([]byte, []int) {
	return file_cosmos_tx_signing_v1beta1_signing_proto_rawDescGZIP(), []int{1, 0, 0}
}

func (x *SignatureDescriptor_Data_Single) GetMode() SignMode {
	if x != nil {
		return x.Mode
	}
	return SignMode_SIGN_MODE_UNSPECIFIED
}

func (x *SignatureDescriptor_Data_Single) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

// Multi is the signature data for a multisig public key
type SignatureDescriptor_Data_Multi struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// bitarray specifies which keys within the multisig are signing
	Bitarray *v1beta1.CompactBitArray `protobuf:"bytes,1,opt,name=bitarray,proto3" json:"bitarray,omitempty"`
	// signatures is the signatures of the multi-signature
	Signatures []*SignatureDescriptor_Data `protobuf:"bytes,2,rep,name=signatures,proto3" json:"signatures,omitempty"`
}

func (x *SignatureDescriptor_Data_Multi) Reset() {
	*x = SignatureDescriptor_Data_Multi{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_tx_signing_v1beta1_signing_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignatureDescriptor_Data_Multi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignatureDescriptor_Data_Multi) ProtoMessage() {}

// Deprecated: Use SignatureDescriptor_Data_Multi.ProtoReflect.Descriptor instead.
func (*SignatureDescriptor_Data_Multi) Descriptor() ([]byte, []int) {
	return file_cosmos_tx_signing_v1beta1_signing_proto_rawDescGZIP(), []int{1, 0, 1}
}

func (x *SignatureDescriptor_Data_Multi) GetBitarray() *v1beta1.CompactBitArray {
	if x != nil {
		return x.Bitarray
	}
	return nil
}

func (x *SignatureDescriptor_Data_Multi) GetSignatures() []*SignatureDescriptor_Data {
	if x != nil {
		return x.Signatures
	}
	return nil
}

var File_cosmos_tx_signing_v1beta1_signing_proto protoreflect.FileDescriptor

var file_cosmos_tx_signing_v1beta1_signing_proto_rawDesc = []byte{
	0x0a, 0x27, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x74, 0x78, 0x2f, 0x73, 0x69, 0x67, 0x6e,
	0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x73, 0x69, 0x67, 0x6e,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x74, 0x78, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x1a, 0x2d, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66,
	0x0a, 0x14, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x4e, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x74, 0x78, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x22, 0xf5, 0x04, 0x0a, 0x13, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x33,
	0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x4b, 0x65, 0x79, 0x12, 0x47, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x33, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x74, 0x78, 0x2e, 0x73, 0x69,
	0x67, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08,
	0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0xc3, 0x03, 0x0a, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x54, 0x0a, 0x06, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x3a, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x74, 0x78, 0x2e, 0x73, 0x69,
	0x67, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x48, 0x00, 0x52,
	0x06, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x51, 0x0a, 0x05, 0x6d, 0x75, 0x6c, 0x74, 0x69,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x74, 0x78, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x48, 0x00, 0x52, 0x05, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x1a, 0x5f, 0x0a, 0x06, 0x53, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x74, 0x78, 0x2e, 0x73,
	0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53,
	0x69, 0x67, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x1a, 0xa9, 0x01, 0x0a, 0x05,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x12, 0x4b, 0x0a, 0x08, 0x62, 0x69, 0x74, 0x61, 0x72, 0x72, 0x61,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x69, 0x67,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74,
	0x42, 0x69, 0x74, 0x41, 0x72, 0x72, 0x61, 0x79, 0x52, 0x08, 0x62, 0x69, 0x74, 0x61, 0x72, 0x72,
	0x61, 0x79, 0x12, 0x53, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x74, 0x78, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x42, 0x05, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x2a, 0x8d,
	0x01, 0x0a, 0x08, 0x53, 0x69, 0x67, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x53,
	0x49, 0x47, 0x4e, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x49, 0x47, 0x4e, 0x5f, 0x4d,
	0x4f, 0x44, 0x45, 0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11,
	0x53, 0x49, 0x47, 0x4e, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x54, 0x45, 0x58, 0x54, 0x55, 0x41,
	0x4c, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x49, 0x47, 0x4e, 0x5f, 0x4d, 0x4f, 0x44, 0x45,
	0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x5f, 0x41, 0x55, 0x58, 0x10, 0x03, 0x12, 0x1f, 0x0a,
	0x1b, 0x53, 0x49, 0x47, 0x4e, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x4c, 0x45, 0x47, 0x41, 0x43,
	0x59, 0x5f, 0x41, 0x4d, 0x49, 0x4e, 0x4f, 0x5f, 0x4a, 0x53, 0x4f, 0x4e, 0x10, 0x7f, 0x42, 0xff,
	0x01, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x74, 0x78,
	0x2e, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x42, 0x0c, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x74, 0x78, 0x2f, 0x73, 0x69, 0x67,
	0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x73, 0x69, 0x67,
	0x6e, 0x69, 0x6e, 0x67, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x54,
	0x53, 0xaa, 0x02, 0x19, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x54, 0x78, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x19,
	0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x54, 0x78, 0x5c, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e,
	0x67, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x25, 0x43, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x5c, 0x54, 0x78, 0x5c, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x1c, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x54, 0x78, 0x3a, 0x3a,
	0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_tx_signing_v1beta1_signing_proto_rawDescOnce sync.Once
	file_cosmos_tx_signing_v1beta1_signing_proto_rawDescData = file_cosmos_tx_signing_v1beta1_signing_proto_rawDesc
)

func file_cosmos_tx_signing_v1beta1_signing_proto_rawDescGZIP() []byte {
	file_cosmos_tx_signing_v1beta1_signing_proto_rawDescOnce.Do(func() {
		file_cosmos_tx_signing_v1beta1_signing_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_tx_signing_v1beta1_signing_proto_rawDescData)
	})
	return file_cosmos_tx_signing_v1beta1_signing_proto_rawDescData
}

var file_cosmos_tx_signing_v1beta1_signing_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cosmos_tx_signing_v1beta1_signing_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cosmos_tx_signing_v1beta1_signing_proto_goTypes = []interface{}{
	(SignMode)(0),                           // 0: cosmos.tx.signing.v1beta1.SignMode
	(*SignatureDescriptors)(nil),            // 1: cosmos.tx.signing.v1beta1.SignatureDescriptors
	(*SignatureDescriptor)(nil),             // 2: cosmos.tx.signing.v1beta1.SignatureDescriptor
	(*SignatureDescriptor_Data)(nil),        // 3: cosmos.tx.signing.v1beta1.SignatureDescriptor.Data
	(*SignatureDescriptor_Data_Single)(nil), // 4: cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Single
	(*SignatureDescriptor_Data_Multi)(nil),  // 5: cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Multi
	(*anypb.Any)(nil),                       // 6: google.protobuf.Any
	(*v1beta1.CompactBitArray)(nil),         // 7: cosmos.crypto.multisig.v1beta1.CompactBitArray
}
var file_cosmos_tx_signing_v1beta1_signing_proto_depIdxs = []int32{
	2, // 0: cosmos.tx.signing.v1beta1.SignatureDescriptors.signatures:type_name -> cosmos.tx.signing.v1beta1.SignatureDescriptor
	6, // 1: cosmos.tx.signing.v1beta1.SignatureDescriptor.public_key:type_name -> google.protobuf.Any
	3, // 2: cosmos.tx.signing.v1beta1.SignatureDescriptor.data:type_name -> cosmos.tx.signing.v1beta1.SignatureDescriptor.Data
	4, // 3: cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.single:type_name -> cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Single
	5, // 4: cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.multi:type_name -> cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Multi
	0, // 5: cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Single.mode:type_name -> cosmos.tx.signing.v1beta1.SignMode
	7, // 6: cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Multi.bitarray:type_name -> cosmos.crypto.multisig.v1beta1.CompactBitArray
	3, // 7: cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Multi.signatures:type_name -> cosmos.tx.signing.v1beta1.SignatureDescriptor.Data
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_cosmos_tx_signing_v1beta1_signing_proto_init() }
func file_cosmos_tx_signing_v1beta1_signing_proto_init() {
	if File_cosmos_tx_signing_v1beta1_signing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_tx_signing_v1beta1_signing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignatureDescriptors); i {
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
		file_cosmos_tx_signing_v1beta1_signing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignatureDescriptor); i {
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
		file_cosmos_tx_signing_v1beta1_signing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignatureDescriptor_Data); i {
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
		file_cosmos_tx_signing_v1beta1_signing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignatureDescriptor_Data_Single); i {
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
		file_cosmos_tx_signing_v1beta1_signing_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignatureDescriptor_Data_Multi); i {
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
	file_cosmos_tx_signing_v1beta1_signing_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*SignatureDescriptor_Data_Single_)(nil),
		(*SignatureDescriptor_Data_Multi_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cosmos_tx_signing_v1beta1_signing_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_tx_signing_v1beta1_signing_proto_goTypes,
		DependencyIndexes: file_cosmos_tx_signing_v1beta1_signing_proto_depIdxs,
		EnumInfos:         file_cosmos_tx_signing_v1beta1_signing_proto_enumTypes,
		MessageInfos:      file_cosmos_tx_signing_v1beta1_signing_proto_msgTypes,
	}.Build()
	File_cosmos_tx_signing_v1beta1_signing_proto = out.File
	file_cosmos_tx_signing_v1beta1_signing_proto_rawDesc = nil
	file_cosmos_tx_signing_v1beta1_signing_proto_goTypes = nil
	file_cosmos_tx_signing_v1beta1_signing_proto_depIdxs = nil
}
