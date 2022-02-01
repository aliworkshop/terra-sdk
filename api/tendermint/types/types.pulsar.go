package types

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	crypto "github.com/aliworkshop/terra-sdk/api/tendermint/crypto"
	version "github.com/aliworkshop/terra-sdk/api/tendermint/version"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_PartSetHeader       protoreflect.MessageDescriptor
	fd_PartSetHeader_total protoreflect.FieldDescriptor
	fd_PartSetHeader_hash  protoreflect.FieldDescriptor
)

func init() {
	file_tendermint_types_types_proto_init()
	md_PartSetHeader = File_tendermint_types_types_proto.Messages().ByName("PartSetHeader")
	fd_PartSetHeader_total = md_PartSetHeader.Fields().ByName("total")
	fd_PartSetHeader_hash = md_PartSetHeader.Fields().ByName("hash")
}

var _ protoreflect.Message = (*fastReflection_PartSetHeader)(nil)

type fastReflection_PartSetHeader PartSetHeader

func (x *PartSetHeader) ProtoReflect() protoreflect.Message {
	return (*fastReflection_PartSetHeader)(x)
}

func (x *PartSetHeader) slowProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_PartSetHeader_messageType fastReflection_PartSetHeader_messageType
var _ protoreflect.MessageType = fastReflection_PartSetHeader_messageType{}

type fastReflection_PartSetHeader_messageType struct{}

func (x fastReflection_PartSetHeader_messageType) Zero() protoreflect.Message {
	return (*fastReflection_PartSetHeader)(nil)
}
func (x fastReflection_PartSetHeader_messageType) New() protoreflect.Message {
	return new(fastReflection_PartSetHeader)
}
func (x fastReflection_PartSetHeader_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_PartSetHeader
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_PartSetHeader) Descriptor() protoreflect.MessageDescriptor {
	return md_PartSetHeader
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_PartSetHeader) Type() protoreflect.MessageType {
	return _fastReflection_PartSetHeader_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_PartSetHeader) New() protoreflect.Message {
	return new(fastReflection_PartSetHeader)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_PartSetHeader) Interface() protoreflect.ProtoMessage {
	return (*PartSetHeader)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_PartSetHeader) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Total != uint32(0) {
		value := protoreflect.ValueOfUint32(x.Total)
		if !f(fd_PartSetHeader_total, value) {
			return
		}
	}
	if len(x.Hash) != 0 {
		value := protoreflect.ValueOfBytes(x.Hash)
		if !f(fd_PartSetHeader_hash, value) {
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
func (x *fastReflection_PartSetHeader) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "tendermint.types.PartSetHeader.total":
		return x.Total != uint32(0)
	case "tendermint.types.PartSetHeader.hash":
		return len(x.Hash) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.PartSetHeader"))
		}
		panic(fmt.Errorf("message tendermint.types.PartSetHeader does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PartSetHeader) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "tendermint.types.PartSetHeader.total":
		x.Total = uint32(0)
	case "tendermint.types.PartSetHeader.hash":
		x.Hash = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.PartSetHeader"))
		}
		panic(fmt.Errorf("message tendermint.types.PartSetHeader does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_PartSetHeader) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "tendermint.types.PartSetHeader.total":
		value := x.Total
		return protoreflect.ValueOfUint32(value)
	case "tendermint.types.PartSetHeader.hash":
		value := x.Hash
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.PartSetHeader"))
		}
		panic(fmt.Errorf("message tendermint.types.PartSetHeader does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_PartSetHeader) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "tendermint.types.PartSetHeader.total":
		x.Total = uint32(value.Uint())
	case "tendermint.types.PartSetHeader.hash":
		x.Hash = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.PartSetHeader"))
		}
		panic(fmt.Errorf("message tendermint.types.PartSetHeader does not contain field %s", fd.FullName()))
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
func (x *fastReflection_PartSetHeader) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.PartSetHeader.total":
		panic(fmt.Errorf("field total of message tendermint.types.PartSetHeader is not mutable"))
	case "tendermint.types.PartSetHeader.hash":
		panic(fmt.Errorf("field hash of message tendermint.types.PartSetHeader is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.PartSetHeader"))
		}
		panic(fmt.Errorf("message tendermint.types.PartSetHeader does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_PartSetHeader) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.PartSetHeader.total":
		return protoreflect.ValueOfUint32(uint32(0))
	case "tendermint.types.PartSetHeader.hash":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.PartSetHeader"))
		}
		panic(fmt.Errorf("message tendermint.types.PartSetHeader does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_PartSetHeader) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in tendermint.types.PartSetHeader", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_PartSetHeader) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PartSetHeader) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_PartSetHeader) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_PartSetHeader) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*PartSetHeader)
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
		if x.Total != 0 {
			n += 1 + runtime.Sov(uint64(x.Total))
		}
		l = len(x.Hash)
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
		x := input.Message.Interface().(*PartSetHeader)
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
		if len(x.Hash) > 0 {
			i -= len(x.Hash)
			copy(dAtA[i:], x.Hash)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Hash)))
			i--
			dAtA[i] = 0x12
		}
		if x.Total != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Total))
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
		x := input.Message.Interface().(*PartSetHeader)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: PartSetHeader: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: PartSetHeader: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
				}
				x.Total = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Total |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
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
				x.Hash = append(x.Hash[:0], dAtA[iNdEx:postIndex]...)
				if x.Hash == nil {
					x.Hash = []byte{}
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
	md_Part       protoreflect.MessageDescriptor
	fd_Part_index protoreflect.FieldDescriptor
	fd_Part_bytes protoreflect.FieldDescriptor
	fd_Part_proof protoreflect.FieldDescriptor
)

func init() {
	file_tendermint_types_types_proto_init()
	md_Part = File_tendermint_types_types_proto.Messages().ByName("Part")
	fd_Part_index = md_Part.Fields().ByName("index")
	fd_Part_bytes = md_Part.Fields().ByName("bytes")
	fd_Part_proof = md_Part.Fields().ByName("proof")
}

var _ protoreflect.Message = (*fastReflection_Part)(nil)

type fastReflection_Part Part

func (x *Part) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Part)(x)
}

func (x *Part) slowProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Part_messageType fastReflection_Part_messageType
var _ protoreflect.MessageType = fastReflection_Part_messageType{}

type fastReflection_Part_messageType struct{}

func (x fastReflection_Part_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Part)(nil)
}
func (x fastReflection_Part_messageType) New() protoreflect.Message {
	return new(fastReflection_Part)
}
func (x fastReflection_Part_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Part
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Part) Descriptor() protoreflect.MessageDescriptor {
	return md_Part
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Part) Type() protoreflect.MessageType {
	return _fastReflection_Part_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Part) New() protoreflect.Message {
	return new(fastReflection_Part)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Part) Interface() protoreflect.ProtoMessage {
	return (*Part)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Part) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Index != uint32(0) {
		value := protoreflect.ValueOfUint32(x.Index)
		if !f(fd_Part_index, value) {
			return
		}
	}
	if len(x.Bytes) != 0 {
		value := protoreflect.ValueOfBytes(x.Bytes)
		if !f(fd_Part_bytes, value) {
			return
		}
	}
	if x.Proof != nil {
		value := protoreflect.ValueOfMessage(x.Proof.ProtoReflect())
		if !f(fd_Part_proof, value) {
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
func (x *fastReflection_Part) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "tendermint.types.Part.index":
		return x.Index != uint32(0)
	case "tendermint.types.Part.bytes":
		return len(x.Bytes) != 0
	case "tendermint.types.Part.proof":
		return x.Proof != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Part"))
		}
		panic(fmt.Errorf("message tendermint.types.Part does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Part) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "tendermint.types.Part.index":
		x.Index = uint32(0)
	case "tendermint.types.Part.bytes":
		x.Bytes = nil
	case "tendermint.types.Part.proof":
		x.Proof = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Part"))
		}
		panic(fmt.Errorf("message tendermint.types.Part does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Part) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "tendermint.types.Part.index":
		value := x.Index
		return protoreflect.ValueOfUint32(value)
	case "tendermint.types.Part.bytes":
		value := x.Bytes
		return protoreflect.ValueOfBytes(value)
	case "tendermint.types.Part.proof":
		value := x.Proof
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Part"))
		}
		panic(fmt.Errorf("message tendermint.types.Part does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Part) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "tendermint.types.Part.index":
		x.Index = uint32(value.Uint())
	case "tendermint.types.Part.bytes":
		x.Bytes = value.Bytes()
	case "tendermint.types.Part.proof":
		x.Proof = value.Message().Interface().(*crypto.Proof)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Part"))
		}
		panic(fmt.Errorf("message tendermint.types.Part does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Part) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.Part.proof":
		if x.Proof == nil {
			x.Proof = new(crypto.Proof)
		}
		return protoreflect.ValueOfMessage(x.Proof.ProtoReflect())
	case "tendermint.types.Part.index":
		panic(fmt.Errorf("field index of message tendermint.types.Part is not mutable"))
	case "tendermint.types.Part.bytes":
		panic(fmt.Errorf("field bytes of message tendermint.types.Part is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Part"))
		}
		panic(fmt.Errorf("message tendermint.types.Part does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Part) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.Part.index":
		return protoreflect.ValueOfUint32(uint32(0))
	case "tendermint.types.Part.bytes":
		return protoreflect.ValueOfBytes(nil)
	case "tendermint.types.Part.proof":
		m := new(crypto.Proof)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Part"))
		}
		panic(fmt.Errorf("message tendermint.types.Part does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Part) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in tendermint.types.Part", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Part) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Part) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Part) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Part) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Part)
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
		if x.Index != 0 {
			n += 1 + runtime.Sov(uint64(x.Index))
		}
		l = len(x.Bytes)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Proof != nil {
			l = options.Size(x.Proof)
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
		x := input.Message.Interface().(*Part)
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
		if x.Proof != nil {
			encoded, err := options.Marshal(x.Proof)
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
		if len(x.Bytes) > 0 {
			i -= len(x.Bytes)
			copy(dAtA[i:], x.Bytes)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Bytes)))
			i--
			dAtA[i] = 0x12
		}
		if x.Index != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Index))
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
		x := input.Message.Interface().(*Part)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Part: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Part: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
				}
				x.Index = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Index |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Bytes", wireType)
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
				x.Bytes = append(x.Bytes[:0], dAtA[iNdEx:postIndex]...)
				if x.Bytes == nil {
					x.Bytes = []byte{}
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
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
				if x.Proof == nil {
					x.Proof = &crypto.Proof{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Proof); err != nil {
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
	md_BlockID                 protoreflect.MessageDescriptor
	fd_BlockID_hash            protoreflect.FieldDescriptor
	fd_BlockID_part_set_header protoreflect.FieldDescriptor
)

func init() {
	file_tendermint_types_types_proto_init()
	md_BlockID = File_tendermint_types_types_proto.Messages().ByName("BlockID")
	fd_BlockID_hash = md_BlockID.Fields().ByName("hash")
	fd_BlockID_part_set_header = md_BlockID.Fields().ByName("part_set_header")
}

var _ protoreflect.Message = (*fastReflection_BlockID)(nil)

type fastReflection_BlockID BlockID

func (x *BlockID) ProtoReflect() protoreflect.Message {
	return (*fastReflection_BlockID)(x)
}

func (x *BlockID) slowProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_BlockID_messageType fastReflection_BlockID_messageType
var _ protoreflect.MessageType = fastReflection_BlockID_messageType{}

type fastReflection_BlockID_messageType struct{}

func (x fastReflection_BlockID_messageType) Zero() protoreflect.Message {
	return (*fastReflection_BlockID)(nil)
}
func (x fastReflection_BlockID_messageType) New() protoreflect.Message {
	return new(fastReflection_BlockID)
}
func (x fastReflection_BlockID_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_BlockID
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_BlockID) Descriptor() protoreflect.MessageDescriptor {
	return md_BlockID
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_BlockID) Type() protoreflect.MessageType {
	return _fastReflection_BlockID_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_BlockID) New() protoreflect.Message {
	return new(fastReflection_BlockID)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_BlockID) Interface() protoreflect.ProtoMessage {
	return (*BlockID)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_BlockID) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Hash) != 0 {
		value := protoreflect.ValueOfBytes(x.Hash)
		if !f(fd_BlockID_hash, value) {
			return
		}
	}
	if x.PartSetHeader != nil {
		value := protoreflect.ValueOfMessage(x.PartSetHeader.ProtoReflect())
		if !f(fd_BlockID_part_set_header, value) {
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
func (x *fastReflection_BlockID) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "tendermint.types.BlockID.hash":
		return len(x.Hash) != 0
	case "tendermint.types.BlockID.part_set_header":
		return x.PartSetHeader != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.BlockID"))
		}
		panic(fmt.Errorf("message tendermint.types.BlockID does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BlockID) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "tendermint.types.BlockID.hash":
		x.Hash = nil
	case "tendermint.types.BlockID.part_set_header":
		x.PartSetHeader = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.BlockID"))
		}
		panic(fmt.Errorf("message tendermint.types.BlockID does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_BlockID) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "tendermint.types.BlockID.hash":
		value := x.Hash
		return protoreflect.ValueOfBytes(value)
	case "tendermint.types.BlockID.part_set_header":
		value := x.PartSetHeader
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.BlockID"))
		}
		panic(fmt.Errorf("message tendermint.types.BlockID does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_BlockID) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "tendermint.types.BlockID.hash":
		x.Hash = value.Bytes()
	case "tendermint.types.BlockID.part_set_header":
		x.PartSetHeader = value.Message().Interface().(*PartSetHeader)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.BlockID"))
		}
		panic(fmt.Errorf("message tendermint.types.BlockID does not contain field %s", fd.FullName()))
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
func (x *fastReflection_BlockID) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.BlockID.part_set_header":
		if x.PartSetHeader == nil {
			x.PartSetHeader = new(PartSetHeader)
		}
		return protoreflect.ValueOfMessage(x.PartSetHeader.ProtoReflect())
	case "tendermint.types.BlockID.hash":
		panic(fmt.Errorf("field hash of message tendermint.types.BlockID is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.BlockID"))
		}
		panic(fmt.Errorf("message tendermint.types.BlockID does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_BlockID) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.BlockID.hash":
		return protoreflect.ValueOfBytes(nil)
	case "tendermint.types.BlockID.part_set_header":
		m := new(PartSetHeader)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.BlockID"))
		}
		panic(fmt.Errorf("message tendermint.types.BlockID does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_BlockID) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in tendermint.types.BlockID", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_BlockID) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BlockID) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_BlockID) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_BlockID) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*BlockID)
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
		l = len(x.Hash)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.PartSetHeader != nil {
			l = options.Size(x.PartSetHeader)
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
		x := input.Message.Interface().(*BlockID)
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
		if x.PartSetHeader != nil {
			encoded, err := options.Marshal(x.PartSetHeader)
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
		if len(x.Hash) > 0 {
			i -= len(x.Hash)
			copy(dAtA[i:], x.Hash)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Hash)))
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
		x := input.Message.Interface().(*BlockID)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BlockID: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BlockID: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
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
				x.Hash = append(x.Hash[:0], dAtA[iNdEx:postIndex]...)
				if x.Hash == nil {
					x.Hash = []byte{}
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field PartSetHeader", wireType)
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
				if x.PartSetHeader == nil {
					x.PartSetHeader = &PartSetHeader{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.PartSetHeader); err != nil {
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
	md_Header                      protoreflect.MessageDescriptor
	fd_Header_version              protoreflect.FieldDescriptor
	fd_Header_chain_id             protoreflect.FieldDescriptor
	fd_Header_height               protoreflect.FieldDescriptor
	fd_Header_time                 protoreflect.FieldDescriptor
	fd_Header_last_block_id        protoreflect.FieldDescriptor
	fd_Header_last_commit_hash     protoreflect.FieldDescriptor
	fd_Header_data_hash            protoreflect.FieldDescriptor
	fd_Header_validators_hash      protoreflect.FieldDescriptor
	fd_Header_next_validators_hash protoreflect.FieldDescriptor
	fd_Header_consensus_hash       protoreflect.FieldDescriptor
	fd_Header_app_hash             protoreflect.FieldDescriptor
	fd_Header_last_results_hash    protoreflect.FieldDescriptor
	fd_Header_evidence_hash        protoreflect.FieldDescriptor
	fd_Header_proposer_address     protoreflect.FieldDescriptor
)

func init() {
	file_tendermint_types_types_proto_init()
	md_Header = File_tendermint_types_types_proto.Messages().ByName("Header")
	fd_Header_version = md_Header.Fields().ByName("version")
	fd_Header_chain_id = md_Header.Fields().ByName("chain_id")
	fd_Header_height = md_Header.Fields().ByName("height")
	fd_Header_time = md_Header.Fields().ByName("time")
	fd_Header_last_block_id = md_Header.Fields().ByName("last_block_id")
	fd_Header_last_commit_hash = md_Header.Fields().ByName("last_commit_hash")
	fd_Header_data_hash = md_Header.Fields().ByName("data_hash")
	fd_Header_validators_hash = md_Header.Fields().ByName("validators_hash")
	fd_Header_next_validators_hash = md_Header.Fields().ByName("next_validators_hash")
	fd_Header_consensus_hash = md_Header.Fields().ByName("consensus_hash")
	fd_Header_app_hash = md_Header.Fields().ByName("app_hash")
	fd_Header_last_results_hash = md_Header.Fields().ByName("last_results_hash")
	fd_Header_evidence_hash = md_Header.Fields().ByName("evidence_hash")
	fd_Header_proposer_address = md_Header.Fields().ByName("proposer_address")
}

var _ protoreflect.Message = (*fastReflection_Header)(nil)

type fastReflection_Header Header

func (x *Header) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Header)(x)
}

func (x *Header) slowProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Header_messageType fastReflection_Header_messageType
var _ protoreflect.MessageType = fastReflection_Header_messageType{}

type fastReflection_Header_messageType struct{}

func (x fastReflection_Header_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Header)(nil)
}
func (x fastReflection_Header_messageType) New() protoreflect.Message {
	return new(fastReflection_Header)
}
func (x fastReflection_Header_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Header
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Header) Descriptor() protoreflect.MessageDescriptor {
	return md_Header
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Header) Type() protoreflect.MessageType {
	return _fastReflection_Header_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Header) New() protoreflect.Message {
	return new(fastReflection_Header)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Header) Interface() protoreflect.ProtoMessage {
	return (*Header)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Header) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Version != nil {
		value := protoreflect.ValueOfMessage(x.Version.ProtoReflect())
		if !f(fd_Header_version, value) {
			return
		}
	}
	if x.ChainId != "" {
		value := protoreflect.ValueOfString(x.ChainId)
		if !f(fd_Header_chain_id, value) {
			return
		}
	}
	if x.Height != int64(0) {
		value := protoreflect.ValueOfInt64(x.Height)
		if !f(fd_Header_height, value) {
			return
		}
	}
	if x.Time != nil {
		value := protoreflect.ValueOfMessage(x.Time.ProtoReflect())
		if !f(fd_Header_time, value) {
			return
		}
	}
	if x.LastBlockId != nil {
		value := protoreflect.ValueOfMessage(x.LastBlockId.ProtoReflect())
		if !f(fd_Header_last_block_id, value) {
			return
		}
	}
	if len(x.LastCommitHash) != 0 {
		value := protoreflect.ValueOfBytes(x.LastCommitHash)
		if !f(fd_Header_last_commit_hash, value) {
			return
		}
	}
	if len(x.DataHash) != 0 {
		value := protoreflect.ValueOfBytes(x.DataHash)
		if !f(fd_Header_data_hash, value) {
			return
		}
	}
	if len(x.ValidatorsHash) != 0 {
		value := protoreflect.ValueOfBytes(x.ValidatorsHash)
		if !f(fd_Header_validators_hash, value) {
			return
		}
	}
	if len(x.NextValidatorsHash) != 0 {
		value := protoreflect.ValueOfBytes(x.NextValidatorsHash)
		if !f(fd_Header_next_validators_hash, value) {
			return
		}
	}
	if len(x.ConsensusHash) != 0 {
		value := protoreflect.ValueOfBytes(x.ConsensusHash)
		if !f(fd_Header_consensus_hash, value) {
			return
		}
	}
	if len(x.AppHash) != 0 {
		value := protoreflect.ValueOfBytes(x.AppHash)
		if !f(fd_Header_app_hash, value) {
			return
		}
	}
	if len(x.LastResultsHash) != 0 {
		value := protoreflect.ValueOfBytes(x.LastResultsHash)
		if !f(fd_Header_last_results_hash, value) {
			return
		}
	}
	if len(x.EvidenceHash) != 0 {
		value := protoreflect.ValueOfBytes(x.EvidenceHash)
		if !f(fd_Header_evidence_hash, value) {
			return
		}
	}
	if len(x.ProposerAddress) != 0 {
		value := protoreflect.ValueOfBytes(x.ProposerAddress)
		if !f(fd_Header_proposer_address, value) {
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
func (x *fastReflection_Header) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "tendermint.types.Header.version":
		return x.Version != nil
	case "tendermint.types.Header.chain_id":
		return x.ChainId != ""
	case "tendermint.types.Header.height":
		return x.Height != int64(0)
	case "tendermint.types.Header.time":
		return x.Time != nil
	case "tendermint.types.Header.last_block_id":
		return x.LastBlockId != nil
	case "tendermint.types.Header.last_commit_hash":
		return len(x.LastCommitHash) != 0
	case "tendermint.types.Header.data_hash":
		return len(x.DataHash) != 0
	case "tendermint.types.Header.validators_hash":
		return len(x.ValidatorsHash) != 0
	case "tendermint.types.Header.next_validators_hash":
		return len(x.NextValidatorsHash) != 0
	case "tendermint.types.Header.consensus_hash":
		return len(x.ConsensusHash) != 0
	case "tendermint.types.Header.app_hash":
		return len(x.AppHash) != 0
	case "tendermint.types.Header.last_results_hash":
		return len(x.LastResultsHash) != 0
	case "tendermint.types.Header.evidence_hash":
		return len(x.EvidenceHash) != 0
	case "tendermint.types.Header.proposer_address":
		return len(x.ProposerAddress) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Header"))
		}
		panic(fmt.Errorf("message tendermint.types.Header does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Header) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "tendermint.types.Header.version":
		x.Version = nil
	case "tendermint.types.Header.chain_id":
		x.ChainId = ""
	case "tendermint.types.Header.height":
		x.Height = int64(0)
	case "tendermint.types.Header.time":
		x.Time = nil
	case "tendermint.types.Header.last_block_id":
		x.LastBlockId = nil
	case "tendermint.types.Header.last_commit_hash":
		x.LastCommitHash = nil
	case "tendermint.types.Header.data_hash":
		x.DataHash = nil
	case "tendermint.types.Header.validators_hash":
		x.ValidatorsHash = nil
	case "tendermint.types.Header.next_validators_hash":
		x.NextValidatorsHash = nil
	case "tendermint.types.Header.consensus_hash":
		x.ConsensusHash = nil
	case "tendermint.types.Header.app_hash":
		x.AppHash = nil
	case "tendermint.types.Header.last_results_hash":
		x.LastResultsHash = nil
	case "tendermint.types.Header.evidence_hash":
		x.EvidenceHash = nil
	case "tendermint.types.Header.proposer_address":
		x.ProposerAddress = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Header"))
		}
		panic(fmt.Errorf("message tendermint.types.Header does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Header) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "tendermint.types.Header.version":
		value := x.Version
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "tendermint.types.Header.chain_id":
		value := x.ChainId
		return protoreflect.ValueOfString(value)
	case "tendermint.types.Header.height":
		value := x.Height
		return protoreflect.ValueOfInt64(value)
	case "tendermint.types.Header.time":
		value := x.Time
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "tendermint.types.Header.last_block_id":
		value := x.LastBlockId
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "tendermint.types.Header.last_commit_hash":
		value := x.LastCommitHash
		return protoreflect.ValueOfBytes(value)
	case "tendermint.types.Header.data_hash":
		value := x.DataHash
		return protoreflect.ValueOfBytes(value)
	case "tendermint.types.Header.validators_hash":
		value := x.ValidatorsHash
		return protoreflect.ValueOfBytes(value)
	case "tendermint.types.Header.next_validators_hash":
		value := x.NextValidatorsHash
		return protoreflect.ValueOfBytes(value)
	case "tendermint.types.Header.consensus_hash":
		value := x.ConsensusHash
		return protoreflect.ValueOfBytes(value)
	case "tendermint.types.Header.app_hash":
		value := x.AppHash
		return protoreflect.ValueOfBytes(value)
	case "tendermint.types.Header.last_results_hash":
		value := x.LastResultsHash
		return protoreflect.ValueOfBytes(value)
	case "tendermint.types.Header.evidence_hash":
		value := x.EvidenceHash
		return protoreflect.ValueOfBytes(value)
	case "tendermint.types.Header.proposer_address":
		value := x.ProposerAddress
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Header"))
		}
		panic(fmt.Errorf("message tendermint.types.Header does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Header) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "tendermint.types.Header.version":
		x.Version = value.Message().Interface().(*version.Consensus)
	case "tendermint.types.Header.chain_id":
		x.ChainId = value.Interface().(string)
	case "tendermint.types.Header.height":
		x.Height = value.Int()
	case "tendermint.types.Header.time":
		x.Time = value.Message().Interface().(*timestamppb.Timestamp)
	case "tendermint.types.Header.last_block_id":
		x.LastBlockId = value.Message().Interface().(*BlockID)
	case "tendermint.types.Header.last_commit_hash":
		x.LastCommitHash = value.Bytes()
	case "tendermint.types.Header.data_hash":
		x.DataHash = value.Bytes()
	case "tendermint.types.Header.validators_hash":
		x.ValidatorsHash = value.Bytes()
	case "tendermint.types.Header.next_validators_hash":
		x.NextValidatorsHash = value.Bytes()
	case "tendermint.types.Header.consensus_hash":
		x.ConsensusHash = value.Bytes()
	case "tendermint.types.Header.app_hash":
		x.AppHash = value.Bytes()
	case "tendermint.types.Header.last_results_hash":
		x.LastResultsHash = value.Bytes()
	case "tendermint.types.Header.evidence_hash":
		x.EvidenceHash = value.Bytes()
	case "tendermint.types.Header.proposer_address":
		x.ProposerAddress = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Header"))
		}
		panic(fmt.Errorf("message tendermint.types.Header does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Header) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.Header.version":
		if x.Version == nil {
			x.Version = new(version.Consensus)
		}
		return protoreflect.ValueOfMessage(x.Version.ProtoReflect())
	case "tendermint.types.Header.time":
		if x.Time == nil {
			x.Time = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.Time.ProtoReflect())
	case "tendermint.types.Header.last_block_id":
		if x.LastBlockId == nil {
			x.LastBlockId = new(BlockID)
		}
		return protoreflect.ValueOfMessage(x.LastBlockId.ProtoReflect())
	case "tendermint.types.Header.chain_id":
		panic(fmt.Errorf("field chain_id of message tendermint.types.Header is not mutable"))
	case "tendermint.types.Header.height":
		panic(fmt.Errorf("field height of message tendermint.types.Header is not mutable"))
	case "tendermint.types.Header.last_commit_hash":
		panic(fmt.Errorf("field last_commit_hash of message tendermint.types.Header is not mutable"))
	case "tendermint.types.Header.data_hash":
		panic(fmt.Errorf("field data_hash of message tendermint.types.Header is not mutable"))
	case "tendermint.types.Header.validators_hash":
		panic(fmt.Errorf("field validators_hash of message tendermint.types.Header is not mutable"))
	case "tendermint.types.Header.next_validators_hash":
		panic(fmt.Errorf("field next_validators_hash of message tendermint.types.Header is not mutable"))
	case "tendermint.types.Header.consensus_hash":
		panic(fmt.Errorf("field consensus_hash of message tendermint.types.Header is not mutable"))
	case "tendermint.types.Header.app_hash":
		panic(fmt.Errorf("field app_hash of message tendermint.types.Header is not mutable"))
	case "tendermint.types.Header.last_results_hash":
		panic(fmt.Errorf("field last_results_hash of message tendermint.types.Header is not mutable"))
	case "tendermint.types.Header.evidence_hash":
		panic(fmt.Errorf("field evidence_hash of message tendermint.types.Header is not mutable"))
	case "tendermint.types.Header.proposer_address":
		panic(fmt.Errorf("field proposer_address of message tendermint.types.Header is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Header"))
		}
		panic(fmt.Errorf("message tendermint.types.Header does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Header) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.Header.version":
		m := new(version.Consensus)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "tendermint.types.Header.chain_id":
		return protoreflect.ValueOfString("")
	case "tendermint.types.Header.height":
		return protoreflect.ValueOfInt64(int64(0))
	case "tendermint.types.Header.time":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "tendermint.types.Header.last_block_id":
		m := new(BlockID)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "tendermint.types.Header.last_commit_hash":
		return protoreflect.ValueOfBytes(nil)
	case "tendermint.types.Header.data_hash":
		return protoreflect.ValueOfBytes(nil)
	case "tendermint.types.Header.validators_hash":
		return protoreflect.ValueOfBytes(nil)
	case "tendermint.types.Header.next_validators_hash":
		return protoreflect.ValueOfBytes(nil)
	case "tendermint.types.Header.consensus_hash":
		return protoreflect.ValueOfBytes(nil)
	case "tendermint.types.Header.app_hash":
		return protoreflect.ValueOfBytes(nil)
	case "tendermint.types.Header.last_results_hash":
		return protoreflect.ValueOfBytes(nil)
	case "tendermint.types.Header.evidence_hash":
		return protoreflect.ValueOfBytes(nil)
	case "tendermint.types.Header.proposer_address":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Header"))
		}
		panic(fmt.Errorf("message tendermint.types.Header does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Header) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in tendermint.types.Header", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Header) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Header) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Header) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Header) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Header)
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
		if x.Version != nil {
			l = options.Size(x.Version)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ChainId)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Height != 0 {
			n += 1 + runtime.Sov(uint64(x.Height))
		}
		if x.Time != nil {
			l = options.Size(x.Time)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.LastBlockId != nil {
			l = options.Size(x.LastBlockId)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.LastCommitHash)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.DataHash)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ValidatorsHash)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.NextValidatorsHash)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ConsensusHash)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.AppHash)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.LastResultsHash)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.EvidenceHash)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ProposerAddress)
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
		x := input.Message.Interface().(*Header)
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
		if len(x.ProposerAddress) > 0 {
			i -= len(x.ProposerAddress)
			copy(dAtA[i:], x.ProposerAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ProposerAddress)))
			i--
			dAtA[i] = 0x72
		}
		if len(x.EvidenceHash) > 0 {
			i -= len(x.EvidenceHash)
			copy(dAtA[i:], x.EvidenceHash)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.EvidenceHash)))
			i--
			dAtA[i] = 0x6a
		}
		if len(x.LastResultsHash) > 0 {
			i -= len(x.LastResultsHash)
			copy(dAtA[i:], x.LastResultsHash)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.LastResultsHash)))
			i--
			dAtA[i] = 0x62
		}
		if len(x.AppHash) > 0 {
			i -= len(x.AppHash)
			copy(dAtA[i:], x.AppHash)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.AppHash)))
			i--
			dAtA[i] = 0x5a
		}
		if len(x.ConsensusHash) > 0 {
			i -= len(x.ConsensusHash)
			copy(dAtA[i:], x.ConsensusHash)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ConsensusHash)))
			i--
			dAtA[i] = 0x52
		}
		if len(x.NextValidatorsHash) > 0 {
			i -= len(x.NextValidatorsHash)
			copy(dAtA[i:], x.NextValidatorsHash)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.NextValidatorsHash)))
			i--
			dAtA[i] = 0x4a
		}
		if len(x.ValidatorsHash) > 0 {
			i -= len(x.ValidatorsHash)
			copy(dAtA[i:], x.ValidatorsHash)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ValidatorsHash)))
			i--
			dAtA[i] = 0x42
		}
		if len(x.DataHash) > 0 {
			i -= len(x.DataHash)
			copy(dAtA[i:], x.DataHash)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.DataHash)))
			i--
			dAtA[i] = 0x3a
		}
		if len(x.LastCommitHash) > 0 {
			i -= len(x.LastCommitHash)
			copy(dAtA[i:], x.LastCommitHash)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.LastCommitHash)))
			i--
			dAtA[i] = 0x32
		}
		if x.LastBlockId != nil {
			encoded, err := options.Marshal(x.LastBlockId)
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
			dAtA[i] = 0x2a
		}
		if x.Time != nil {
			encoded, err := options.Marshal(x.Time)
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
		if x.Height != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Height))
			i--
			dAtA[i] = 0x18
		}
		if len(x.ChainId) > 0 {
			i -= len(x.ChainId)
			copy(dAtA[i:], x.ChainId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ChainId)))
			i--
			dAtA[i] = 0x12
		}
		if x.Version != nil {
			encoded, err := options.Marshal(x.Version)
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
		x := input.Message.Interface().(*Header)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Header: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Header: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
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
				if x.Version == nil {
					x.Version = &version.Consensus{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Version); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
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
				x.ChainId = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
				}
				x.Height = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Height |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
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
				if x.Time == nil {
					x.Time = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Time); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field LastBlockId", wireType)
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
				if x.LastBlockId == nil {
					x.LastBlockId = &BlockID{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.LastBlockId); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 6:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field LastCommitHash", wireType)
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
				x.LastCommitHash = append(x.LastCommitHash[:0], dAtA[iNdEx:postIndex]...)
				if x.LastCommitHash == nil {
					x.LastCommitHash = []byte{}
				}
				iNdEx = postIndex
			case 7:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DataHash", wireType)
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
				x.DataHash = append(x.DataHash[:0], dAtA[iNdEx:postIndex]...)
				if x.DataHash == nil {
					x.DataHash = []byte{}
				}
				iNdEx = postIndex
			case 8:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ValidatorsHash", wireType)
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
				x.ValidatorsHash = append(x.ValidatorsHash[:0], dAtA[iNdEx:postIndex]...)
				if x.ValidatorsHash == nil {
					x.ValidatorsHash = []byte{}
				}
				iNdEx = postIndex
			case 9:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field NextValidatorsHash", wireType)
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
				x.NextValidatorsHash = append(x.NextValidatorsHash[:0], dAtA[iNdEx:postIndex]...)
				if x.NextValidatorsHash == nil {
					x.NextValidatorsHash = []byte{}
				}
				iNdEx = postIndex
			case 10:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ConsensusHash", wireType)
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
				x.ConsensusHash = append(x.ConsensusHash[:0], dAtA[iNdEx:postIndex]...)
				if x.ConsensusHash == nil {
					x.ConsensusHash = []byte{}
				}
				iNdEx = postIndex
			case 11:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AppHash", wireType)
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
				x.AppHash = append(x.AppHash[:0], dAtA[iNdEx:postIndex]...)
				if x.AppHash == nil {
					x.AppHash = []byte{}
				}
				iNdEx = postIndex
			case 12:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field LastResultsHash", wireType)
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
				x.LastResultsHash = append(x.LastResultsHash[:0], dAtA[iNdEx:postIndex]...)
				if x.LastResultsHash == nil {
					x.LastResultsHash = []byte{}
				}
				iNdEx = postIndex
			case 13:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field EvidenceHash", wireType)
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
				x.EvidenceHash = append(x.EvidenceHash[:0], dAtA[iNdEx:postIndex]...)
				if x.EvidenceHash == nil {
					x.EvidenceHash = []byte{}
				}
				iNdEx = postIndex
			case 14:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProposerAddress", wireType)
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
				x.ProposerAddress = append(x.ProposerAddress[:0], dAtA[iNdEx:postIndex]...)
				if x.ProposerAddress == nil {
					x.ProposerAddress = []byte{}
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

var _ protoreflect.List = (*_Data_1_list)(nil)

type _Data_1_list struct {
	list *[][]byte
}

func (x *_Data_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_Data_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfBytes((*x.list)[i])
}

func (x *_Data_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Bytes()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_Data_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Bytes()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_Data_1_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message Data at list field Txs as it is not of Message kind"))
}

func (x *_Data_1_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_Data_1_list) NewElement() protoreflect.Value {
	var v []byte
	return protoreflect.ValueOfBytes(v)
}

func (x *_Data_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_Data     protoreflect.MessageDescriptor
	fd_Data_txs protoreflect.FieldDescriptor
)

func init() {
	file_tendermint_types_types_proto_init()
	md_Data = File_tendermint_types_types_proto.Messages().ByName("Data")
	fd_Data_txs = md_Data.Fields().ByName("txs")
}

var _ protoreflect.Message = (*fastReflection_Data)(nil)

type fastReflection_Data Data

func (x *Data) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Data)(x)
}

func (x *Data) slowProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Data_messageType fastReflection_Data_messageType
var _ protoreflect.MessageType = fastReflection_Data_messageType{}

type fastReflection_Data_messageType struct{}

func (x fastReflection_Data_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Data)(nil)
}
func (x fastReflection_Data_messageType) New() protoreflect.Message {
	return new(fastReflection_Data)
}
func (x fastReflection_Data_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Data
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Data) Descriptor() protoreflect.MessageDescriptor {
	return md_Data
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Data) Type() protoreflect.MessageType {
	return _fastReflection_Data_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Data) New() protoreflect.Message {
	return new(fastReflection_Data)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Data) Interface() protoreflect.ProtoMessage {
	return (*Data)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Data) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Txs) != 0 {
		value := protoreflect.ValueOfList(&_Data_1_list{list: &x.Txs})
		if !f(fd_Data_txs, value) {
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
func (x *fastReflection_Data) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "tendermint.types.Data.txs":
		return len(x.Txs) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Data"))
		}
		panic(fmt.Errorf("message tendermint.types.Data does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Data) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "tendermint.types.Data.txs":
		x.Txs = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Data"))
		}
		panic(fmt.Errorf("message tendermint.types.Data does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Data) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "tendermint.types.Data.txs":
		if len(x.Txs) == 0 {
			return protoreflect.ValueOfList(&_Data_1_list{})
		}
		listValue := &_Data_1_list{list: &x.Txs}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Data"))
		}
		panic(fmt.Errorf("message tendermint.types.Data does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Data) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "tendermint.types.Data.txs":
		lv := value.List()
		clv := lv.(*_Data_1_list)
		x.Txs = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Data"))
		}
		panic(fmt.Errorf("message tendermint.types.Data does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Data) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.Data.txs":
		if x.Txs == nil {
			x.Txs = [][]byte{}
		}
		value := &_Data_1_list{list: &x.Txs}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Data"))
		}
		panic(fmt.Errorf("message tendermint.types.Data does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Data) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.Data.txs":
		list := [][]byte{}
		return protoreflect.ValueOfList(&_Data_1_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Data"))
		}
		panic(fmt.Errorf("message tendermint.types.Data does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Data) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in tendermint.types.Data", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Data) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Data) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Data) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Data) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Data)
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
		if len(x.Txs) > 0 {
			for _, b := range x.Txs {
				l = len(b)
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
		x := input.Message.Interface().(*Data)
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
		if len(x.Txs) > 0 {
			for iNdEx := len(x.Txs) - 1; iNdEx >= 0; iNdEx-- {
				i -= len(x.Txs[iNdEx])
				copy(dAtA[i:], x.Txs[iNdEx])
				i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Txs[iNdEx])))
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
		x := input.Message.Interface().(*Data)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Data: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Data: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Txs", wireType)
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
				x.Txs = append(x.Txs, make([]byte, postIndex-iNdEx))
				copy(x.Txs[len(x.Txs)-1], dAtA[iNdEx:postIndex])
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
	md_Vote                   protoreflect.MessageDescriptor
	fd_Vote_type              protoreflect.FieldDescriptor
	fd_Vote_height            protoreflect.FieldDescriptor
	fd_Vote_round             protoreflect.FieldDescriptor
	fd_Vote_block_id          protoreflect.FieldDescriptor
	fd_Vote_timestamp         protoreflect.FieldDescriptor
	fd_Vote_validator_address protoreflect.FieldDescriptor
	fd_Vote_validator_index   protoreflect.FieldDescriptor
	fd_Vote_signature         protoreflect.FieldDescriptor
)

func init() {
	file_tendermint_types_types_proto_init()
	md_Vote = File_tendermint_types_types_proto.Messages().ByName("Vote")
	fd_Vote_type = md_Vote.Fields().ByName("type")
	fd_Vote_height = md_Vote.Fields().ByName("height")
	fd_Vote_round = md_Vote.Fields().ByName("round")
	fd_Vote_block_id = md_Vote.Fields().ByName("block_id")
	fd_Vote_timestamp = md_Vote.Fields().ByName("timestamp")
	fd_Vote_validator_address = md_Vote.Fields().ByName("validator_address")
	fd_Vote_validator_index = md_Vote.Fields().ByName("validator_index")
	fd_Vote_signature = md_Vote.Fields().ByName("signature")
}

var _ protoreflect.Message = (*fastReflection_Vote)(nil)

type fastReflection_Vote Vote

func (x *Vote) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Vote)(x)
}

func (x *Vote) slowProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_types_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Vote_messageType fastReflection_Vote_messageType
var _ protoreflect.MessageType = fastReflection_Vote_messageType{}

type fastReflection_Vote_messageType struct{}

func (x fastReflection_Vote_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Vote)(nil)
}
func (x fastReflection_Vote_messageType) New() protoreflect.Message {
	return new(fastReflection_Vote)
}
func (x fastReflection_Vote_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Vote
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Vote) Descriptor() protoreflect.MessageDescriptor {
	return md_Vote
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Vote) Type() protoreflect.MessageType {
	return _fastReflection_Vote_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Vote) New() protoreflect.Message {
	return new(fastReflection_Vote)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Vote) Interface() protoreflect.ProtoMessage {
	return (*Vote)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Vote) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Type_ != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.Type_))
		if !f(fd_Vote_type, value) {
			return
		}
	}
	if x.Height != int64(0) {
		value := protoreflect.ValueOfInt64(x.Height)
		if !f(fd_Vote_height, value) {
			return
		}
	}
	if x.Round != int32(0) {
		value := protoreflect.ValueOfInt32(x.Round)
		if !f(fd_Vote_round, value) {
			return
		}
	}
	if x.BlockId != nil {
		value := protoreflect.ValueOfMessage(x.BlockId.ProtoReflect())
		if !f(fd_Vote_block_id, value) {
			return
		}
	}
	if x.Timestamp != nil {
		value := protoreflect.ValueOfMessage(x.Timestamp.ProtoReflect())
		if !f(fd_Vote_timestamp, value) {
			return
		}
	}
	if len(x.ValidatorAddress) != 0 {
		value := protoreflect.ValueOfBytes(x.ValidatorAddress)
		if !f(fd_Vote_validator_address, value) {
			return
		}
	}
	if x.ValidatorIndex != int32(0) {
		value := protoreflect.ValueOfInt32(x.ValidatorIndex)
		if !f(fd_Vote_validator_index, value) {
			return
		}
	}
	if len(x.Signature) != 0 {
		value := protoreflect.ValueOfBytes(x.Signature)
		if !f(fd_Vote_signature, value) {
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
func (x *fastReflection_Vote) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "tendermint.types.Vote.type":
		return x.Type_ != 0
	case "tendermint.types.Vote.height":
		return x.Height != int64(0)
	case "tendermint.types.Vote.round":
		return x.Round != int32(0)
	case "tendermint.types.Vote.block_id":
		return x.BlockId != nil
	case "tendermint.types.Vote.timestamp":
		return x.Timestamp != nil
	case "tendermint.types.Vote.validator_address":
		return len(x.ValidatorAddress) != 0
	case "tendermint.types.Vote.validator_index":
		return x.ValidatorIndex != int32(0)
	case "tendermint.types.Vote.signature":
		return len(x.Signature) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Vote"))
		}
		panic(fmt.Errorf("message tendermint.types.Vote does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Vote) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "tendermint.types.Vote.type":
		x.Type_ = 0
	case "tendermint.types.Vote.height":
		x.Height = int64(0)
	case "tendermint.types.Vote.round":
		x.Round = int32(0)
	case "tendermint.types.Vote.block_id":
		x.BlockId = nil
	case "tendermint.types.Vote.timestamp":
		x.Timestamp = nil
	case "tendermint.types.Vote.validator_address":
		x.ValidatorAddress = nil
	case "tendermint.types.Vote.validator_index":
		x.ValidatorIndex = int32(0)
	case "tendermint.types.Vote.signature":
		x.Signature = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Vote"))
		}
		panic(fmt.Errorf("message tendermint.types.Vote does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Vote) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "tendermint.types.Vote.type":
		value := x.Type_
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	case "tendermint.types.Vote.height":
		value := x.Height
		return protoreflect.ValueOfInt64(value)
	case "tendermint.types.Vote.round":
		value := x.Round
		return protoreflect.ValueOfInt32(value)
	case "tendermint.types.Vote.block_id":
		value := x.BlockId
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "tendermint.types.Vote.timestamp":
		value := x.Timestamp
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "tendermint.types.Vote.validator_address":
		value := x.ValidatorAddress
		return protoreflect.ValueOfBytes(value)
	case "tendermint.types.Vote.validator_index":
		value := x.ValidatorIndex
		return protoreflect.ValueOfInt32(value)
	case "tendermint.types.Vote.signature":
		value := x.Signature
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Vote"))
		}
		panic(fmt.Errorf("message tendermint.types.Vote does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Vote) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "tendermint.types.Vote.type":
		x.Type_ = (SignedMsgType)(value.Enum())
	case "tendermint.types.Vote.height":
		x.Height = value.Int()
	case "tendermint.types.Vote.round":
		x.Round = int32(value.Int())
	case "tendermint.types.Vote.block_id":
		x.BlockId = value.Message().Interface().(*BlockID)
	case "tendermint.types.Vote.timestamp":
		x.Timestamp = value.Message().Interface().(*timestamppb.Timestamp)
	case "tendermint.types.Vote.validator_address":
		x.ValidatorAddress = value.Bytes()
	case "tendermint.types.Vote.validator_index":
		x.ValidatorIndex = int32(value.Int())
	case "tendermint.types.Vote.signature":
		x.Signature = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Vote"))
		}
		panic(fmt.Errorf("message tendermint.types.Vote does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Vote) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.Vote.block_id":
		if x.BlockId == nil {
			x.BlockId = new(BlockID)
		}
		return protoreflect.ValueOfMessage(x.BlockId.ProtoReflect())
	case "tendermint.types.Vote.timestamp":
		if x.Timestamp == nil {
			x.Timestamp = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.Timestamp.ProtoReflect())
	case "tendermint.types.Vote.type":
		panic(fmt.Errorf("field type of message tendermint.types.Vote is not mutable"))
	case "tendermint.types.Vote.height":
		panic(fmt.Errorf("field height of message tendermint.types.Vote is not mutable"))
	case "tendermint.types.Vote.round":
		panic(fmt.Errorf("field round of message tendermint.types.Vote is not mutable"))
	case "tendermint.types.Vote.validator_address":
		panic(fmt.Errorf("field validator_address of message tendermint.types.Vote is not mutable"))
	case "tendermint.types.Vote.validator_index":
		panic(fmt.Errorf("field validator_index of message tendermint.types.Vote is not mutable"))
	case "tendermint.types.Vote.signature":
		panic(fmt.Errorf("field signature of message tendermint.types.Vote is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Vote"))
		}
		panic(fmt.Errorf("message tendermint.types.Vote does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Vote) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.Vote.type":
		return protoreflect.ValueOfEnum(0)
	case "tendermint.types.Vote.height":
		return protoreflect.ValueOfInt64(int64(0))
	case "tendermint.types.Vote.round":
		return protoreflect.ValueOfInt32(int32(0))
	case "tendermint.types.Vote.block_id":
		m := new(BlockID)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "tendermint.types.Vote.timestamp":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "tendermint.types.Vote.validator_address":
		return protoreflect.ValueOfBytes(nil)
	case "tendermint.types.Vote.validator_index":
		return protoreflect.ValueOfInt32(int32(0))
	case "tendermint.types.Vote.signature":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Vote"))
		}
		panic(fmt.Errorf("message tendermint.types.Vote does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Vote) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in tendermint.types.Vote", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Vote) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Vote) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Vote) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Vote) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Vote)
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
		if x.Type_ != 0 {
			n += 1 + runtime.Sov(uint64(x.Type_))
		}
		if x.Height != 0 {
			n += 1 + runtime.Sov(uint64(x.Height))
		}
		if x.Round != 0 {
			n += 1 + runtime.Sov(uint64(x.Round))
		}
		if x.BlockId != nil {
			l = options.Size(x.BlockId)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Timestamp != nil {
			l = options.Size(x.Timestamp)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ValidatorAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.ValidatorIndex != 0 {
			n += 1 + runtime.Sov(uint64(x.ValidatorIndex))
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
		x := input.Message.Interface().(*Vote)
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
			dAtA[i] = 0x42
		}
		if x.ValidatorIndex != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.ValidatorIndex))
			i--
			dAtA[i] = 0x38
		}
		if len(x.ValidatorAddress) > 0 {
			i -= len(x.ValidatorAddress)
			copy(dAtA[i:], x.ValidatorAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ValidatorAddress)))
			i--
			dAtA[i] = 0x32
		}
		if x.Timestamp != nil {
			encoded, err := options.Marshal(x.Timestamp)
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
			dAtA[i] = 0x2a
		}
		if x.BlockId != nil {
			encoded, err := options.Marshal(x.BlockId)
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
		if x.Round != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Round))
			i--
			dAtA[i] = 0x18
		}
		if x.Height != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Height))
			i--
			dAtA[i] = 0x10
		}
		if x.Type_ != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Type_))
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
		x := input.Message.Interface().(*Vote)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Vote: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Vote: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Type_", wireType)
				}
				x.Type_ = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Type_ |= SignedMsgType(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
				}
				x.Height = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Height |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Round", wireType)
				}
				x.Round = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Round |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BlockId", wireType)
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
				if x.BlockId == nil {
					x.BlockId = &BlockID{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.BlockId); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
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
				if x.Timestamp == nil {
					x.Timestamp = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Timestamp); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 6:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
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
				x.ValidatorAddress = append(x.ValidatorAddress[:0], dAtA[iNdEx:postIndex]...)
				if x.ValidatorAddress == nil {
					x.ValidatorAddress = []byte{}
				}
				iNdEx = postIndex
			case 7:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ValidatorIndex", wireType)
				}
				x.ValidatorIndex = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.ValidatorIndex |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 8:
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

var _ protoreflect.List = (*_Commit_4_list)(nil)

type _Commit_4_list struct {
	list *[]*CommitSig
}

func (x *_Commit_4_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_Commit_4_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_Commit_4_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*CommitSig)
	(*x.list)[i] = concreteValue
}

func (x *_Commit_4_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*CommitSig)
	*x.list = append(*x.list, concreteValue)
}

func (x *_Commit_4_list) AppendMutable() protoreflect.Value {
	v := new(CommitSig)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Commit_4_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_Commit_4_list) NewElement() protoreflect.Value {
	v := new(CommitSig)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Commit_4_list) IsValid() bool {
	return x.list != nil
}

var (
	md_Commit            protoreflect.MessageDescriptor
	fd_Commit_height     protoreflect.FieldDescriptor
	fd_Commit_round      protoreflect.FieldDescriptor
	fd_Commit_block_id   protoreflect.FieldDescriptor
	fd_Commit_signatures protoreflect.FieldDescriptor
)

func init() {
	file_tendermint_types_types_proto_init()
	md_Commit = File_tendermint_types_types_proto.Messages().ByName("Commit")
	fd_Commit_height = md_Commit.Fields().ByName("height")
	fd_Commit_round = md_Commit.Fields().ByName("round")
	fd_Commit_block_id = md_Commit.Fields().ByName("block_id")
	fd_Commit_signatures = md_Commit.Fields().ByName("signatures")
}

var _ protoreflect.Message = (*fastReflection_Commit)(nil)

type fastReflection_Commit Commit

func (x *Commit) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Commit)(x)
}

func (x *Commit) slowProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_types_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Commit_messageType fastReflection_Commit_messageType
var _ protoreflect.MessageType = fastReflection_Commit_messageType{}

type fastReflection_Commit_messageType struct{}

func (x fastReflection_Commit_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Commit)(nil)
}
func (x fastReflection_Commit_messageType) New() protoreflect.Message {
	return new(fastReflection_Commit)
}
func (x fastReflection_Commit_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Commit
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Commit) Descriptor() protoreflect.MessageDescriptor {
	return md_Commit
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Commit) Type() protoreflect.MessageType {
	return _fastReflection_Commit_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Commit) New() protoreflect.Message {
	return new(fastReflection_Commit)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Commit) Interface() protoreflect.ProtoMessage {
	return (*Commit)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Commit) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Height != int64(0) {
		value := protoreflect.ValueOfInt64(x.Height)
		if !f(fd_Commit_height, value) {
			return
		}
	}
	if x.Round != int32(0) {
		value := protoreflect.ValueOfInt32(x.Round)
		if !f(fd_Commit_round, value) {
			return
		}
	}
	if x.BlockId != nil {
		value := protoreflect.ValueOfMessage(x.BlockId.ProtoReflect())
		if !f(fd_Commit_block_id, value) {
			return
		}
	}
	if len(x.Signatures) != 0 {
		value := protoreflect.ValueOfList(&_Commit_4_list{list: &x.Signatures})
		if !f(fd_Commit_signatures, value) {
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
func (x *fastReflection_Commit) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "tendermint.types.Commit.height":
		return x.Height != int64(0)
	case "tendermint.types.Commit.round":
		return x.Round != int32(0)
	case "tendermint.types.Commit.block_id":
		return x.BlockId != nil
	case "tendermint.types.Commit.signatures":
		return len(x.Signatures) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Commit"))
		}
		panic(fmt.Errorf("message tendermint.types.Commit does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Commit) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "tendermint.types.Commit.height":
		x.Height = int64(0)
	case "tendermint.types.Commit.round":
		x.Round = int32(0)
	case "tendermint.types.Commit.block_id":
		x.BlockId = nil
	case "tendermint.types.Commit.signatures":
		x.Signatures = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Commit"))
		}
		panic(fmt.Errorf("message tendermint.types.Commit does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Commit) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "tendermint.types.Commit.height":
		value := x.Height
		return protoreflect.ValueOfInt64(value)
	case "tendermint.types.Commit.round":
		value := x.Round
		return protoreflect.ValueOfInt32(value)
	case "tendermint.types.Commit.block_id":
		value := x.BlockId
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "tendermint.types.Commit.signatures":
		if len(x.Signatures) == 0 {
			return protoreflect.ValueOfList(&_Commit_4_list{})
		}
		listValue := &_Commit_4_list{list: &x.Signatures}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Commit"))
		}
		panic(fmt.Errorf("message tendermint.types.Commit does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Commit) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "tendermint.types.Commit.height":
		x.Height = value.Int()
	case "tendermint.types.Commit.round":
		x.Round = int32(value.Int())
	case "tendermint.types.Commit.block_id":
		x.BlockId = value.Message().Interface().(*BlockID)
	case "tendermint.types.Commit.signatures":
		lv := value.List()
		clv := lv.(*_Commit_4_list)
		x.Signatures = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Commit"))
		}
		panic(fmt.Errorf("message tendermint.types.Commit does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Commit) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.Commit.block_id":
		if x.BlockId == nil {
			x.BlockId = new(BlockID)
		}
		return protoreflect.ValueOfMessage(x.BlockId.ProtoReflect())
	case "tendermint.types.Commit.signatures":
		if x.Signatures == nil {
			x.Signatures = []*CommitSig{}
		}
		value := &_Commit_4_list{list: &x.Signatures}
		return protoreflect.ValueOfList(value)
	case "tendermint.types.Commit.height":
		panic(fmt.Errorf("field height of message tendermint.types.Commit is not mutable"))
	case "tendermint.types.Commit.round":
		panic(fmt.Errorf("field round of message tendermint.types.Commit is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Commit"))
		}
		panic(fmt.Errorf("message tendermint.types.Commit does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Commit) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.Commit.height":
		return protoreflect.ValueOfInt64(int64(0))
	case "tendermint.types.Commit.round":
		return protoreflect.ValueOfInt32(int32(0))
	case "tendermint.types.Commit.block_id":
		m := new(BlockID)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "tendermint.types.Commit.signatures":
		list := []*CommitSig{}
		return protoreflect.ValueOfList(&_Commit_4_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Commit"))
		}
		panic(fmt.Errorf("message tendermint.types.Commit does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Commit) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in tendermint.types.Commit", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Commit) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Commit) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Commit) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Commit) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Commit)
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
		if x.Height != 0 {
			n += 1 + runtime.Sov(uint64(x.Height))
		}
		if x.Round != 0 {
			n += 1 + runtime.Sov(uint64(x.Round))
		}
		if x.BlockId != nil {
			l = options.Size(x.BlockId)
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
		x := input.Message.Interface().(*Commit)
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
				dAtA[i] = 0x22
			}
		}
		if x.BlockId != nil {
			encoded, err := options.Marshal(x.BlockId)
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
		if x.Round != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Round))
			i--
			dAtA[i] = 0x10
		}
		if x.Height != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Height))
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
		x := input.Message.Interface().(*Commit)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Commit: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Commit: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
				}
				x.Height = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Height |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Round", wireType)
				}
				x.Round = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Round |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BlockId", wireType)
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
				if x.BlockId == nil {
					x.BlockId = &BlockID{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.BlockId); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 4:
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
				x.Signatures = append(x.Signatures, &CommitSig{})
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
	md_CommitSig                   protoreflect.MessageDescriptor
	fd_CommitSig_block_id_flag     protoreflect.FieldDescriptor
	fd_CommitSig_validator_address protoreflect.FieldDescriptor
	fd_CommitSig_timestamp         protoreflect.FieldDescriptor
	fd_CommitSig_signature         protoreflect.FieldDescriptor
)

func init() {
	file_tendermint_types_types_proto_init()
	md_CommitSig = File_tendermint_types_types_proto.Messages().ByName("CommitSig")
	fd_CommitSig_block_id_flag = md_CommitSig.Fields().ByName("block_id_flag")
	fd_CommitSig_validator_address = md_CommitSig.Fields().ByName("validator_address")
	fd_CommitSig_timestamp = md_CommitSig.Fields().ByName("timestamp")
	fd_CommitSig_signature = md_CommitSig.Fields().ByName("signature")
}

var _ protoreflect.Message = (*fastReflection_CommitSig)(nil)

type fastReflection_CommitSig CommitSig

func (x *CommitSig) ProtoReflect() protoreflect.Message {
	return (*fastReflection_CommitSig)(x)
}

func (x *CommitSig) slowProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_types_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_CommitSig_messageType fastReflection_CommitSig_messageType
var _ protoreflect.MessageType = fastReflection_CommitSig_messageType{}

type fastReflection_CommitSig_messageType struct{}

func (x fastReflection_CommitSig_messageType) Zero() protoreflect.Message {
	return (*fastReflection_CommitSig)(nil)
}
func (x fastReflection_CommitSig_messageType) New() protoreflect.Message {
	return new(fastReflection_CommitSig)
}
func (x fastReflection_CommitSig_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_CommitSig
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_CommitSig) Descriptor() protoreflect.MessageDescriptor {
	return md_CommitSig
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_CommitSig) Type() protoreflect.MessageType {
	return _fastReflection_CommitSig_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_CommitSig) New() protoreflect.Message {
	return new(fastReflection_CommitSig)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_CommitSig) Interface() protoreflect.ProtoMessage {
	return (*CommitSig)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_CommitSig) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BlockIdFlag != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.BlockIdFlag))
		if !f(fd_CommitSig_block_id_flag, value) {
			return
		}
	}
	if len(x.ValidatorAddress) != 0 {
		value := protoreflect.ValueOfBytes(x.ValidatorAddress)
		if !f(fd_CommitSig_validator_address, value) {
			return
		}
	}
	if x.Timestamp != nil {
		value := protoreflect.ValueOfMessage(x.Timestamp.ProtoReflect())
		if !f(fd_CommitSig_timestamp, value) {
			return
		}
	}
	if len(x.Signature) != 0 {
		value := protoreflect.ValueOfBytes(x.Signature)
		if !f(fd_CommitSig_signature, value) {
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
func (x *fastReflection_CommitSig) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "tendermint.types.CommitSig.block_id_flag":
		return x.BlockIdFlag != 0
	case "tendermint.types.CommitSig.validator_address":
		return len(x.ValidatorAddress) != 0
	case "tendermint.types.CommitSig.timestamp":
		return x.Timestamp != nil
	case "tendermint.types.CommitSig.signature":
		return len(x.Signature) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.CommitSig"))
		}
		panic(fmt.Errorf("message tendermint.types.CommitSig does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_CommitSig) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "tendermint.types.CommitSig.block_id_flag":
		x.BlockIdFlag = 0
	case "tendermint.types.CommitSig.validator_address":
		x.ValidatorAddress = nil
	case "tendermint.types.CommitSig.timestamp":
		x.Timestamp = nil
	case "tendermint.types.CommitSig.signature":
		x.Signature = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.CommitSig"))
		}
		panic(fmt.Errorf("message tendermint.types.CommitSig does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_CommitSig) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "tendermint.types.CommitSig.block_id_flag":
		value := x.BlockIdFlag
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	case "tendermint.types.CommitSig.validator_address":
		value := x.ValidatorAddress
		return protoreflect.ValueOfBytes(value)
	case "tendermint.types.CommitSig.timestamp":
		value := x.Timestamp
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "tendermint.types.CommitSig.signature":
		value := x.Signature
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.CommitSig"))
		}
		panic(fmt.Errorf("message tendermint.types.CommitSig does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_CommitSig) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "tendermint.types.CommitSig.block_id_flag":
		x.BlockIdFlag = (BlockIDFlag)(value.Enum())
	case "tendermint.types.CommitSig.validator_address":
		x.ValidatorAddress = value.Bytes()
	case "tendermint.types.CommitSig.timestamp":
		x.Timestamp = value.Message().Interface().(*timestamppb.Timestamp)
	case "tendermint.types.CommitSig.signature":
		x.Signature = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.CommitSig"))
		}
		panic(fmt.Errorf("message tendermint.types.CommitSig does not contain field %s", fd.FullName()))
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
func (x *fastReflection_CommitSig) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.CommitSig.timestamp":
		if x.Timestamp == nil {
			x.Timestamp = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.Timestamp.ProtoReflect())
	case "tendermint.types.CommitSig.block_id_flag":
		panic(fmt.Errorf("field block_id_flag of message tendermint.types.CommitSig is not mutable"))
	case "tendermint.types.CommitSig.validator_address":
		panic(fmt.Errorf("field validator_address of message tendermint.types.CommitSig is not mutable"))
	case "tendermint.types.CommitSig.signature":
		panic(fmt.Errorf("field signature of message tendermint.types.CommitSig is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.CommitSig"))
		}
		panic(fmt.Errorf("message tendermint.types.CommitSig does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_CommitSig) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.CommitSig.block_id_flag":
		return protoreflect.ValueOfEnum(0)
	case "tendermint.types.CommitSig.validator_address":
		return protoreflect.ValueOfBytes(nil)
	case "tendermint.types.CommitSig.timestamp":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "tendermint.types.CommitSig.signature":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.CommitSig"))
		}
		panic(fmt.Errorf("message tendermint.types.CommitSig does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_CommitSig) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in tendermint.types.CommitSig", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_CommitSig) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_CommitSig) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_CommitSig) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_CommitSig) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*CommitSig)
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
		if x.BlockIdFlag != 0 {
			n += 1 + runtime.Sov(uint64(x.BlockIdFlag))
		}
		l = len(x.ValidatorAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Timestamp != nil {
			l = options.Size(x.Timestamp)
			n += 1 + l + runtime.Sov(uint64(l))
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
		x := input.Message.Interface().(*CommitSig)
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
			dAtA[i] = 0x22
		}
		if x.Timestamp != nil {
			encoded, err := options.Marshal(x.Timestamp)
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
		if len(x.ValidatorAddress) > 0 {
			i -= len(x.ValidatorAddress)
			copy(dAtA[i:], x.ValidatorAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ValidatorAddress)))
			i--
			dAtA[i] = 0x12
		}
		if x.BlockIdFlag != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.BlockIdFlag))
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
		x := input.Message.Interface().(*CommitSig)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: CommitSig: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: CommitSig: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BlockIdFlag", wireType)
				}
				x.BlockIdFlag = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.BlockIdFlag |= BlockIDFlag(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
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
				x.ValidatorAddress = append(x.ValidatorAddress[:0], dAtA[iNdEx:postIndex]...)
				if x.ValidatorAddress == nil {
					x.ValidatorAddress = []byte{}
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
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
				if x.Timestamp == nil {
					x.Timestamp = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Timestamp); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 4:
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

var (
	md_Proposal           protoreflect.MessageDescriptor
	fd_Proposal_type      protoreflect.FieldDescriptor
	fd_Proposal_height    protoreflect.FieldDescriptor
	fd_Proposal_round     protoreflect.FieldDescriptor
	fd_Proposal_pol_round protoreflect.FieldDescriptor
	fd_Proposal_block_id  protoreflect.FieldDescriptor
	fd_Proposal_timestamp protoreflect.FieldDescriptor
	fd_Proposal_signature protoreflect.FieldDescriptor
)

func init() {
	file_tendermint_types_types_proto_init()
	md_Proposal = File_tendermint_types_types_proto.Messages().ByName("Proposal")
	fd_Proposal_type = md_Proposal.Fields().ByName("type")
	fd_Proposal_height = md_Proposal.Fields().ByName("height")
	fd_Proposal_round = md_Proposal.Fields().ByName("round")
	fd_Proposal_pol_round = md_Proposal.Fields().ByName("pol_round")
	fd_Proposal_block_id = md_Proposal.Fields().ByName("block_id")
	fd_Proposal_timestamp = md_Proposal.Fields().ByName("timestamp")
	fd_Proposal_signature = md_Proposal.Fields().ByName("signature")
}

var _ protoreflect.Message = (*fastReflection_Proposal)(nil)

type fastReflection_Proposal Proposal

func (x *Proposal) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Proposal)(x)
}

func (x *Proposal) slowProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_types_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Proposal_messageType fastReflection_Proposal_messageType
var _ protoreflect.MessageType = fastReflection_Proposal_messageType{}

type fastReflection_Proposal_messageType struct{}

func (x fastReflection_Proposal_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Proposal)(nil)
}
func (x fastReflection_Proposal_messageType) New() protoreflect.Message {
	return new(fastReflection_Proposal)
}
func (x fastReflection_Proposal_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Proposal
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Proposal) Descriptor() protoreflect.MessageDescriptor {
	return md_Proposal
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Proposal) Type() protoreflect.MessageType {
	return _fastReflection_Proposal_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Proposal) New() protoreflect.Message {
	return new(fastReflection_Proposal)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Proposal) Interface() protoreflect.ProtoMessage {
	return (*Proposal)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Proposal) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Type_ != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.Type_))
		if !f(fd_Proposal_type, value) {
			return
		}
	}
	if x.Height != int64(0) {
		value := protoreflect.ValueOfInt64(x.Height)
		if !f(fd_Proposal_height, value) {
			return
		}
	}
	if x.Round != int32(0) {
		value := protoreflect.ValueOfInt32(x.Round)
		if !f(fd_Proposal_round, value) {
			return
		}
	}
	if x.PolRound != int32(0) {
		value := protoreflect.ValueOfInt32(x.PolRound)
		if !f(fd_Proposal_pol_round, value) {
			return
		}
	}
	if x.BlockId != nil {
		value := protoreflect.ValueOfMessage(x.BlockId.ProtoReflect())
		if !f(fd_Proposal_block_id, value) {
			return
		}
	}
	if x.Timestamp != nil {
		value := protoreflect.ValueOfMessage(x.Timestamp.ProtoReflect())
		if !f(fd_Proposal_timestamp, value) {
			return
		}
	}
	if len(x.Signature) != 0 {
		value := protoreflect.ValueOfBytes(x.Signature)
		if !f(fd_Proposal_signature, value) {
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
func (x *fastReflection_Proposal) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "tendermint.types.Proposal.type":
		return x.Type_ != 0
	case "tendermint.types.Proposal.height":
		return x.Height != int64(0)
	case "tendermint.types.Proposal.round":
		return x.Round != int32(0)
	case "tendermint.types.Proposal.pol_round":
		return x.PolRound != int32(0)
	case "tendermint.types.Proposal.block_id":
		return x.BlockId != nil
	case "tendermint.types.Proposal.timestamp":
		return x.Timestamp != nil
	case "tendermint.types.Proposal.signature":
		return len(x.Signature) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Proposal"))
		}
		panic(fmt.Errorf("message tendermint.types.Proposal does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Proposal) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "tendermint.types.Proposal.type":
		x.Type_ = 0
	case "tendermint.types.Proposal.height":
		x.Height = int64(0)
	case "tendermint.types.Proposal.round":
		x.Round = int32(0)
	case "tendermint.types.Proposal.pol_round":
		x.PolRound = int32(0)
	case "tendermint.types.Proposal.block_id":
		x.BlockId = nil
	case "tendermint.types.Proposal.timestamp":
		x.Timestamp = nil
	case "tendermint.types.Proposal.signature":
		x.Signature = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Proposal"))
		}
		panic(fmt.Errorf("message tendermint.types.Proposal does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Proposal) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "tendermint.types.Proposal.type":
		value := x.Type_
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	case "tendermint.types.Proposal.height":
		value := x.Height
		return protoreflect.ValueOfInt64(value)
	case "tendermint.types.Proposal.round":
		value := x.Round
		return protoreflect.ValueOfInt32(value)
	case "tendermint.types.Proposal.pol_round":
		value := x.PolRound
		return protoreflect.ValueOfInt32(value)
	case "tendermint.types.Proposal.block_id":
		value := x.BlockId
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "tendermint.types.Proposal.timestamp":
		value := x.Timestamp
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "tendermint.types.Proposal.signature":
		value := x.Signature
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Proposal"))
		}
		panic(fmt.Errorf("message tendermint.types.Proposal does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Proposal) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "tendermint.types.Proposal.type":
		x.Type_ = (SignedMsgType)(value.Enum())
	case "tendermint.types.Proposal.height":
		x.Height = value.Int()
	case "tendermint.types.Proposal.round":
		x.Round = int32(value.Int())
	case "tendermint.types.Proposal.pol_round":
		x.PolRound = int32(value.Int())
	case "tendermint.types.Proposal.block_id":
		x.BlockId = value.Message().Interface().(*BlockID)
	case "tendermint.types.Proposal.timestamp":
		x.Timestamp = value.Message().Interface().(*timestamppb.Timestamp)
	case "tendermint.types.Proposal.signature":
		x.Signature = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Proposal"))
		}
		panic(fmt.Errorf("message tendermint.types.Proposal does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Proposal) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.Proposal.block_id":
		if x.BlockId == nil {
			x.BlockId = new(BlockID)
		}
		return protoreflect.ValueOfMessage(x.BlockId.ProtoReflect())
	case "tendermint.types.Proposal.timestamp":
		if x.Timestamp == nil {
			x.Timestamp = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.Timestamp.ProtoReflect())
	case "tendermint.types.Proposal.type":
		panic(fmt.Errorf("field type of message tendermint.types.Proposal is not mutable"))
	case "tendermint.types.Proposal.height":
		panic(fmt.Errorf("field height of message tendermint.types.Proposal is not mutable"))
	case "tendermint.types.Proposal.round":
		panic(fmt.Errorf("field round of message tendermint.types.Proposal is not mutable"))
	case "tendermint.types.Proposal.pol_round":
		panic(fmt.Errorf("field pol_round of message tendermint.types.Proposal is not mutable"))
	case "tendermint.types.Proposal.signature":
		panic(fmt.Errorf("field signature of message tendermint.types.Proposal is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Proposal"))
		}
		panic(fmt.Errorf("message tendermint.types.Proposal does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Proposal) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.Proposal.type":
		return protoreflect.ValueOfEnum(0)
	case "tendermint.types.Proposal.height":
		return protoreflect.ValueOfInt64(int64(0))
	case "tendermint.types.Proposal.round":
		return protoreflect.ValueOfInt32(int32(0))
	case "tendermint.types.Proposal.pol_round":
		return protoreflect.ValueOfInt32(int32(0))
	case "tendermint.types.Proposal.block_id":
		m := new(BlockID)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "tendermint.types.Proposal.timestamp":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "tendermint.types.Proposal.signature":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.Proposal"))
		}
		panic(fmt.Errorf("message tendermint.types.Proposal does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Proposal) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in tendermint.types.Proposal", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Proposal) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Proposal) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Proposal) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Proposal) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Proposal)
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
		if x.Type_ != 0 {
			n += 1 + runtime.Sov(uint64(x.Type_))
		}
		if x.Height != 0 {
			n += 1 + runtime.Sov(uint64(x.Height))
		}
		if x.Round != 0 {
			n += 1 + runtime.Sov(uint64(x.Round))
		}
		if x.PolRound != 0 {
			n += 1 + runtime.Sov(uint64(x.PolRound))
		}
		if x.BlockId != nil {
			l = options.Size(x.BlockId)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Timestamp != nil {
			l = options.Size(x.Timestamp)
			n += 1 + l + runtime.Sov(uint64(l))
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
		x := input.Message.Interface().(*Proposal)
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
			dAtA[i] = 0x3a
		}
		if x.Timestamp != nil {
			encoded, err := options.Marshal(x.Timestamp)
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
			dAtA[i] = 0x32
		}
		if x.BlockId != nil {
			encoded, err := options.Marshal(x.BlockId)
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
			dAtA[i] = 0x2a
		}
		if x.PolRound != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.PolRound))
			i--
			dAtA[i] = 0x20
		}
		if x.Round != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Round))
			i--
			dAtA[i] = 0x18
		}
		if x.Height != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Height))
			i--
			dAtA[i] = 0x10
		}
		if x.Type_ != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Type_))
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
		x := input.Message.Interface().(*Proposal)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Proposal: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Proposal: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Type_", wireType)
				}
				x.Type_ = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Type_ |= SignedMsgType(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
				}
				x.Height = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Height |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Round", wireType)
				}
				x.Round = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Round |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field PolRound", wireType)
				}
				x.PolRound = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.PolRound |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BlockId", wireType)
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
				if x.BlockId == nil {
					x.BlockId = &BlockID{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.BlockId); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 6:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
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
				if x.Timestamp == nil {
					x.Timestamp = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Timestamp); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 7:
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

var (
	md_SignedHeader        protoreflect.MessageDescriptor
	fd_SignedHeader_header protoreflect.FieldDescriptor
	fd_SignedHeader_commit protoreflect.FieldDescriptor
)

func init() {
	file_tendermint_types_types_proto_init()
	md_SignedHeader = File_tendermint_types_types_proto.Messages().ByName("SignedHeader")
	fd_SignedHeader_header = md_SignedHeader.Fields().ByName("header")
	fd_SignedHeader_commit = md_SignedHeader.Fields().ByName("commit")
}

var _ protoreflect.Message = (*fastReflection_SignedHeader)(nil)

type fastReflection_SignedHeader SignedHeader

func (x *SignedHeader) ProtoReflect() protoreflect.Message {
	return (*fastReflection_SignedHeader)(x)
}

func (x *SignedHeader) slowProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_types_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_SignedHeader_messageType fastReflection_SignedHeader_messageType
var _ protoreflect.MessageType = fastReflection_SignedHeader_messageType{}

type fastReflection_SignedHeader_messageType struct{}

func (x fastReflection_SignedHeader_messageType) Zero() protoreflect.Message {
	return (*fastReflection_SignedHeader)(nil)
}
func (x fastReflection_SignedHeader_messageType) New() protoreflect.Message {
	return new(fastReflection_SignedHeader)
}
func (x fastReflection_SignedHeader_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_SignedHeader
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_SignedHeader) Descriptor() protoreflect.MessageDescriptor {
	return md_SignedHeader
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_SignedHeader) Type() protoreflect.MessageType {
	return _fastReflection_SignedHeader_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_SignedHeader) New() protoreflect.Message {
	return new(fastReflection_SignedHeader)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_SignedHeader) Interface() protoreflect.ProtoMessage {
	return (*SignedHeader)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_SignedHeader) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Header != nil {
		value := protoreflect.ValueOfMessage(x.Header.ProtoReflect())
		if !f(fd_SignedHeader_header, value) {
			return
		}
	}
	if x.Commit != nil {
		value := protoreflect.ValueOfMessage(x.Commit.ProtoReflect())
		if !f(fd_SignedHeader_commit, value) {
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
func (x *fastReflection_SignedHeader) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "tendermint.types.SignedHeader.header":
		return x.Header != nil
	case "tendermint.types.SignedHeader.commit":
		return x.Commit != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.SignedHeader"))
		}
		panic(fmt.Errorf("message tendermint.types.SignedHeader does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SignedHeader) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "tendermint.types.SignedHeader.header":
		x.Header = nil
	case "tendermint.types.SignedHeader.commit":
		x.Commit = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.SignedHeader"))
		}
		panic(fmt.Errorf("message tendermint.types.SignedHeader does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_SignedHeader) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "tendermint.types.SignedHeader.header":
		value := x.Header
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "tendermint.types.SignedHeader.commit":
		value := x.Commit
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.SignedHeader"))
		}
		panic(fmt.Errorf("message tendermint.types.SignedHeader does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_SignedHeader) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "tendermint.types.SignedHeader.header":
		x.Header = value.Message().Interface().(*Header)
	case "tendermint.types.SignedHeader.commit":
		x.Commit = value.Message().Interface().(*Commit)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.SignedHeader"))
		}
		panic(fmt.Errorf("message tendermint.types.SignedHeader does not contain field %s", fd.FullName()))
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
func (x *fastReflection_SignedHeader) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.SignedHeader.header":
		if x.Header == nil {
			x.Header = new(Header)
		}
		return protoreflect.ValueOfMessage(x.Header.ProtoReflect())
	case "tendermint.types.SignedHeader.commit":
		if x.Commit == nil {
			x.Commit = new(Commit)
		}
		return protoreflect.ValueOfMessage(x.Commit.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.SignedHeader"))
		}
		panic(fmt.Errorf("message tendermint.types.SignedHeader does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_SignedHeader) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.SignedHeader.header":
		m := new(Header)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "tendermint.types.SignedHeader.commit":
		m := new(Commit)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.SignedHeader"))
		}
		panic(fmt.Errorf("message tendermint.types.SignedHeader does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_SignedHeader) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in tendermint.types.SignedHeader", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_SignedHeader) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SignedHeader) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_SignedHeader) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_SignedHeader) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*SignedHeader)
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
		if x.Commit != nil {
			l = options.Size(x.Commit)
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
		x := input.Message.Interface().(*SignedHeader)
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
		if x.Commit != nil {
			encoded, err := options.Marshal(x.Commit)
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
		x := input.Message.Interface().(*SignedHeader)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SignedHeader: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SignedHeader: illegal tag %d (wire type %d)", fieldNum, wire)
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
					x.Header = &Header{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Header); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Commit", wireType)
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
				if x.Commit == nil {
					x.Commit = &Commit{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Commit); err != nil {
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
	md_LightBlock               protoreflect.MessageDescriptor
	fd_LightBlock_signed_header protoreflect.FieldDescriptor
	fd_LightBlock_validator_set protoreflect.FieldDescriptor
)

func init() {
	file_tendermint_types_types_proto_init()
	md_LightBlock = File_tendermint_types_types_proto.Messages().ByName("LightBlock")
	fd_LightBlock_signed_header = md_LightBlock.Fields().ByName("signed_header")
	fd_LightBlock_validator_set = md_LightBlock.Fields().ByName("validator_set")
}

var _ protoreflect.Message = (*fastReflection_LightBlock)(nil)

type fastReflection_LightBlock LightBlock

func (x *LightBlock) ProtoReflect() protoreflect.Message {
	return (*fastReflection_LightBlock)(x)
}

func (x *LightBlock) slowProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_types_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_LightBlock_messageType fastReflection_LightBlock_messageType
var _ protoreflect.MessageType = fastReflection_LightBlock_messageType{}

type fastReflection_LightBlock_messageType struct{}

func (x fastReflection_LightBlock_messageType) Zero() protoreflect.Message {
	return (*fastReflection_LightBlock)(nil)
}
func (x fastReflection_LightBlock_messageType) New() protoreflect.Message {
	return new(fastReflection_LightBlock)
}
func (x fastReflection_LightBlock_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_LightBlock
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_LightBlock) Descriptor() protoreflect.MessageDescriptor {
	return md_LightBlock
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_LightBlock) Type() protoreflect.MessageType {
	return _fastReflection_LightBlock_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_LightBlock) New() protoreflect.Message {
	return new(fastReflection_LightBlock)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_LightBlock) Interface() protoreflect.ProtoMessage {
	return (*LightBlock)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_LightBlock) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.SignedHeader != nil {
		value := protoreflect.ValueOfMessage(x.SignedHeader.ProtoReflect())
		if !f(fd_LightBlock_signed_header, value) {
			return
		}
	}
	if x.ValidatorSet != nil {
		value := protoreflect.ValueOfMessage(x.ValidatorSet.ProtoReflect())
		if !f(fd_LightBlock_validator_set, value) {
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
func (x *fastReflection_LightBlock) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "tendermint.types.LightBlock.signed_header":
		return x.SignedHeader != nil
	case "tendermint.types.LightBlock.validator_set":
		return x.ValidatorSet != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.LightBlock"))
		}
		panic(fmt.Errorf("message tendermint.types.LightBlock does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_LightBlock) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "tendermint.types.LightBlock.signed_header":
		x.SignedHeader = nil
	case "tendermint.types.LightBlock.validator_set":
		x.ValidatorSet = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.LightBlock"))
		}
		panic(fmt.Errorf("message tendermint.types.LightBlock does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_LightBlock) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "tendermint.types.LightBlock.signed_header":
		value := x.SignedHeader
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "tendermint.types.LightBlock.validator_set":
		value := x.ValidatorSet
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.LightBlock"))
		}
		panic(fmt.Errorf("message tendermint.types.LightBlock does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_LightBlock) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "tendermint.types.LightBlock.signed_header":
		x.SignedHeader = value.Message().Interface().(*SignedHeader)
	case "tendermint.types.LightBlock.validator_set":
		x.ValidatorSet = value.Message().Interface().(*ValidatorSet)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.LightBlock"))
		}
		panic(fmt.Errorf("message tendermint.types.LightBlock does not contain field %s", fd.FullName()))
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
func (x *fastReflection_LightBlock) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.LightBlock.signed_header":
		if x.SignedHeader == nil {
			x.SignedHeader = new(SignedHeader)
		}
		return protoreflect.ValueOfMessage(x.SignedHeader.ProtoReflect())
	case "tendermint.types.LightBlock.validator_set":
		if x.ValidatorSet == nil {
			x.ValidatorSet = new(ValidatorSet)
		}
		return protoreflect.ValueOfMessage(x.ValidatorSet.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.LightBlock"))
		}
		panic(fmt.Errorf("message tendermint.types.LightBlock does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_LightBlock) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.LightBlock.signed_header":
		m := new(SignedHeader)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "tendermint.types.LightBlock.validator_set":
		m := new(ValidatorSet)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.LightBlock"))
		}
		panic(fmt.Errorf("message tendermint.types.LightBlock does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_LightBlock) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in tendermint.types.LightBlock", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_LightBlock) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_LightBlock) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_LightBlock) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_LightBlock) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*LightBlock)
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
		if x.SignedHeader != nil {
			l = options.Size(x.SignedHeader)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.ValidatorSet != nil {
			l = options.Size(x.ValidatorSet)
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
		x := input.Message.Interface().(*LightBlock)
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
		if x.ValidatorSet != nil {
			encoded, err := options.Marshal(x.ValidatorSet)
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
		if x.SignedHeader != nil {
			encoded, err := options.Marshal(x.SignedHeader)
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
		x := input.Message.Interface().(*LightBlock)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: LightBlock: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: LightBlock: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SignedHeader", wireType)
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
				if x.SignedHeader == nil {
					x.SignedHeader = &SignedHeader{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.SignedHeader); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ValidatorSet", wireType)
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
				if x.ValidatorSet == nil {
					x.ValidatorSet = &ValidatorSet{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.ValidatorSet); err != nil {
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
	md_BlockMeta            protoreflect.MessageDescriptor
	fd_BlockMeta_block_id   protoreflect.FieldDescriptor
	fd_BlockMeta_block_size protoreflect.FieldDescriptor
	fd_BlockMeta_header     protoreflect.FieldDescriptor
	fd_BlockMeta_num_txs    protoreflect.FieldDescriptor
)

func init() {
	file_tendermint_types_types_proto_init()
	md_BlockMeta = File_tendermint_types_types_proto.Messages().ByName("BlockMeta")
	fd_BlockMeta_block_id = md_BlockMeta.Fields().ByName("block_id")
	fd_BlockMeta_block_size = md_BlockMeta.Fields().ByName("block_size")
	fd_BlockMeta_header = md_BlockMeta.Fields().ByName("header")
	fd_BlockMeta_num_txs = md_BlockMeta.Fields().ByName("num_txs")
}

var _ protoreflect.Message = (*fastReflection_BlockMeta)(nil)

type fastReflection_BlockMeta BlockMeta

func (x *BlockMeta) ProtoReflect() protoreflect.Message {
	return (*fastReflection_BlockMeta)(x)
}

func (x *BlockMeta) slowProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_types_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_BlockMeta_messageType fastReflection_BlockMeta_messageType
var _ protoreflect.MessageType = fastReflection_BlockMeta_messageType{}

type fastReflection_BlockMeta_messageType struct{}

func (x fastReflection_BlockMeta_messageType) Zero() protoreflect.Message {
	return (*fastReflection_BlockMeta)(nil)
}
func (x fastReflection_BlockMeta_messageType) New() protoreflect.Message {
	return new(fastReflection_BlockMeta)
}
func (x fastReflection_BlockMeta_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_BlockMeta
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_BlockMeta) Descriptor() protoreflect.MessageDescriptor {
	return md_BlockMeta
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_BlockMeta) Type() protoreflect.MessageType {
	return _fastReflection_BlockMeta_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_BlockMeta) New() protoreflect.Message {
	return new(fastReflection_BlockMeta)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_BlockMeta) Interface() protoreflect.ProtoMessage {
	return (*BlockMeta)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_BlockMeta) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BlockId != nil {
		value := protoreflect.ValueOfMessage(x.BlockId.ProtoReflect())
		if !f(fd_BlockMeta_block_id, value) {
			return
		}
	}
	if x.BlockSize != int64(0) {
		value := protoreflect.ValueOfInt64(x.BlockSize)
		if !f(fd_BlockMeta_block_size, value) {
			return
		}
	}
	if x.Header != nil {
		value := protoreflect.ValueOfMessage(x.Header.ProtoReflect())
		if !f(fd_BlockMeta_header, value) {
			return
		}
	}
	if x.NumTxs != int64(0) {
		value := protoreflect.ValueOfInt64(x.NumTxs)
		if !f(fd_BlockMeta_num_txs, value) {
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
func (x *fastReflection_BlockMeta) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "tendermint.types.BlockMeta.block_id":
		return x.BlockId != nil
	case "tendermint.types.BlockMeta.block_size":
		return x.BlockSize != int64(0)
	case "tendermint.types.BlockMeta.header":
		return x.Header != nil
	case "tendermint.types.BlockMeta.num_txs":
		return x.NumTxs != int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.BlockMeta"))
		}
		panic(fmt.Errorf("message tendermint.types.BlockMeta does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BlockMeta) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "tendermint.types.BlockMeta.block_id":
		x.BlockId = nil
	case "tendermint.types.BlockMeta.block_size":
		x.BlockSize = int64(0)
	case "tendermint.types.BlockMeta.header":
		x.Header = nil
	case "tendermint.types.BlockMeta.num_txs":
		x.NumTxs = int64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.BlockMeta"))
		}
		panic(fmt.Errorf("message tendermint.types.BlockMeta does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_BlockMeta) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "tendermint.types.BlockMeta.block_id":
		value := x.BlockId
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "tendermint.types.BlockMeta.block_size":
		value := x.BlockSize
		return protoreflect.ValueOfInt64(value)
	case "tendermint.types.BlockMeta.header":
		value := x.Header
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "tendermint.types.BlockMeta.num_txs":
		value := x.NumTxs
		return protoreflect.ValueOfInt64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.BlockMeta"))
		}
		panic(fmt.Errorf("message tendermint.types.BlockMeta does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_BlockMeta) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "tendermint.types.BlockMeta.block_id":
		x.BlockId = value.Message().Interface().(*BlockID)
	case "tendermint.types.BlockMeta.block_size":
		x.BlockSize = value.Int()
	case "tendermint.types.BlockMeta.header":
		x.Header = value.Message().Interface().(*Header)
	case "tendermint.types.BlockMeta.num_txs":
		x.NumTxs = value.Int()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.BlockMeta"))
		}
		panic(fmt.Errorf("message tendermint.types.BlockMeta does not contain field %s", fd.FullName()))
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
func (x *fastReflection_BlockMeta) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.BlockMeta.block_id":
		if x.BlockId == nil {
			x.BlockId = new(BlockID)
		}
		return protoreflect.ValueOfMessage(x.BlockId.ProtoReflect())
	case "tendermint.types.BlockMeta.header":
		if x.Header == nil {
			x.Header = new(Header)
		}
		return protoreflect.ValueOfMessage(x.Header.ProtoReflect())
	case "tendermint.types.BlockMeta.block_size":
		panic(fmt.Errorf("field block_size of message tendermint.types.BlockMeta is not mutable"))
	case "tendermint.types.BlockMeta.num_txs":
		panic(fmt.Errorf("field num_txs of message tendermint.types.BlockMeta is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.BlockMeta"))
		}
		panic(fmt.Errorf("message tendermint.types.BlockMeta does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_BlockMeta) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.BlockMeta.block_id":
		m := new(BlockID)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "tendermint.types.BlockMeta.block_size":
		return protoreflect.ValueOfInt64(int64(0))
	case "tendermint.types.BlockMeta.header":
		m := new(Header)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "tendermint.types.BlockMeta.num_txs":
		return protoreflect.ValueOfInt64(int64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.BlockMeta"))
		}
		panic(fmt.Errorf("message tendermint.types.BlockMeta does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_BlockMeta) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in tendermint.types.BlockMeta", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_BlockMeta) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BlockMeta) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_BlockMeta) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_BlockMeta) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*BlockMeta)
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
		if x.BlockId != nil {
			l = options.Size(x.BlockId)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.BlockSize != 0 {
			n += 1 + runtime.Sov(uint64(x.BlockSize))
		}
		if x.Header != nil {
			l = options.Size(x.Header)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.NumTxs != 0 {
			n += 1 + runtime.Sov(uint64(x.NumTxs))
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
		x := input.Message.Interface().(*BlockMeta)
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
		if x.NumTxs != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.NumTxs))
			i--
			dAtA[i] = 0x20
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
			dAtA[i] = 0x1a
		}
		if x.BlockSize != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.BlockSize))
			i--
			dAtA[i] = 0x10
		}
		if x.BlockId != nil {
			encoded, err := options.Marshal(x.BlockId)
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
		x := input.Message.Interface().(*BlockMeta)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BlockMeta: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BlockMeta: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BlockId", wireType)
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
				if x.BlockId == nil {
					x.BlockId = &BlockID{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.BlockId); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BlockSize", wireType)
				}
				x.BlockSize = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.BlockSize |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
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
					x.Header = &Header{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Header); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field NumTxs", wireType)
				}
				x.NumTxs = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.NumTxs |= int64(b&0x7F) << shift
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
	md_TxProof           protoreflect.MessageDescriptor
	fd_TxProof_root_hash protoreflect.FieldDescriptor
	fd_TxProof_data      protoreflect.FieldDescriptor
	fd_TxProof_proof     protoreflect.FieldDescriptor
)

func init() {
	file_tendermint_types_types_proto_init()
	md_TxProof = File_tendermint_types_types_proto.Messages().ByName("TxProof")
	fd_TxProof_root_hash = md_TxProof.Fields().ByName("root_hash")
	fd_TxProof_data = md_TxProof.Fields().ByName("data")
	fd_TxProof_proof = md_TxProof.Fields().ByName("proof")
}

var _ protoreflect.Message = (*fastReflection_TxProof)(nil)

type fastReflection_TxProof TxProof

func (x *TxProof) ProtoReflect() protoreflect.Message {
	return (*fastReflection_TxProof)(x)
}

func (x *TxProof) slowProtoReflect() protoreflect.Message {
	mi := &file_tendermint_types_types_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_TxProof_messageType fastReflection_TxProof_messageType
var _ protoreflect.MessageType = fastReflection_TxProof_messageType{}

type fastReflection_TxProof_messageType struct{}

func (x fastReflection_TxProof_messageType) Zero() protoreflect.Message {
	return (*fastReflection_TxProof)(nil)
}
func (x fastReflection_TxProof_messageType) New() protoreflect.Message {
	return new(fastReflection_TxProof)
}
func (x fastReflection_TxProof_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_TxProof
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_TxProof) Descriptor() protoreflect.MessageDescriptor {
	return md_TxProof
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_TxProof) Type() protoreflect.MessageType {
	return _fastReflection_TxProof_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_TxProof) New() protoreflect.Message {
	return new(fastReflection_TxProof)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_TxProof) Interface() protoreflect.ProtoMessage {
	return (*TxProof)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_TxProof) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.RootHash) != 0 {
		value := protoreflect.ValueOfBytes(x.RootHash)
		if !f(fd_TxProof_root_hash, value) {
			return
		}
	}
	if len(x.Data) != 0 {
		value := protoreflect.ValueOfBytes(x.Data)
		if !f(fd_TxProof_data, value) {
			return
		}
	}
	if x.Proof != nil {
		value := protoreflect.ValueOfMessage(x.Proof.ProtoReflect())
		if !f(fd_TxProof_proof, value) {
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
func (x *fastReflection_TxProof) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "tendermint.types.TxProof.root_hash":
		return len(x.RootHash) != 0
	case "tendermint.types.TxProof.data":
		return len(x.Data) != 0
	case "tendermint.types.TxProof.proof":
		return x.Proof != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.TxProof"))
		}
		panic(fmt.Errorf("message tendermint.types.TxProof does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_TxProof) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "tendermint.types.TxProof.root_hash":
		x.RootHash = nil
	case "tendermint.types.TxProof.data":
		x.Data = nil
	case "tendermint.types.TxProof.proof":
		x.Proof = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.TxProof"))
		}
		panic(fmt.Errorf("message tendermint.types.TxProof does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_TxProof) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "tendermint.types.TxProof.root_hash":
		value := x.RootHash
		return protoreflect.ValueOfBytes(value)
	case "tendermint.types.TxProof.data":
		value := x.Data
		return protoreflect.ValueOfBytes(value)
	case "tendermint.types.TxProof.proof":
		value := x.Proof
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.TxProof"))
		}
		panic(fmt.Errorf("message tendermint.types.TxProof does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_TxProof) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "tendermint.types.TxProof.root_hash":
		x.RootHash = value.Bytes()
	case "tendermint.types.TxProof.data":
		x.Data = value.Bytes()
	case "tendermint.types.TxProof.proof":
		x.Proof = value.Message().Interface().(*crypto.Proof)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.TxProof"))
		}
		panic(fmt.Errorf("message tendermint.types.TxProof does not contain field %s", fd.FullName()))
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
func (x *fastReflection_TxProof) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.TxProof.proof":
		if x.Proof == nil {
			x.Proof = new(crypto.Proof)
		}
		return protoreflect.ValueOfMessage(x.Proof.ProtoReflect())
	case "tendermint.types.TxProof.root_hash":
		panic(fmt.Errorf("field root_hash of message tendermint.types.TxProof is not mutable"))
	case "tendermint.types.TxProof.data":
		panic(fmt.Errorf("field data of message tendermint.types.TxProof is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.TxProof"))
		}
		panic(fmt.Errorf("message tendermint.types.TxProof does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_TxProof) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "tendermint.types.TxProof.root_hash":
		return protoreflect.ValueOfBytes(nil)
	case "tendermint.types.TxProof.data":
		return protoreflect.ValueOfBytes(nil)
	case "tendermint.types.TxProof.proof":
		m := new(crypto.Proof)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: tendermint.types.TxProof"))
		}
		panic(fmt.Errorf("message tendermint.types.TxProof does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_TxProof) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in tendermint.types.TxProof", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_TxProof) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_TxProof) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_TxProof) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_TxProof) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*TxProof)
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
		l = len(x.RootHash)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Data)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Proof != nil {
			l = options.Size(x.Proof)
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
		x := input.Message.Interface().(*TxProof)
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
		if x.Proof != nil {
			encoded, err := options.Marshal(x.Proof)
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
		if len(x.Data) > 0 {
			i -= len(x.Data)
			copy(dAtA[i:], x.Data)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Data)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.RootHash) > 0 {
			i -= len(x.RootHash)
			copy(dAtA[i:], x.RootHash)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.RootHash)))
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
		x := input.Message.Interface().(*TxProof)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: TxProof: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: TxProof: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RootHash", wireType)
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
				x.RootHash = append(x.RootHash[:0], dAtA[iNdEx:postIndex]...)
				if x.RootHash == nil {
					x.RootHash = []byte{}
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
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
				x.Data = append(x.Data[:0], dAtA[iNdEx:postIndex]...)
				if x.Data == nil {
					x.Data = []byte{}
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
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
				if x.Proof == nil {
					x.Proof = &crypto.Proof{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Proof); err != nil {
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
// source: tendermint/types/types.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// BlockIdFlag indicates which BlcokID the signature is for
type BlockIDFlag int32

const (
	BlockIDFlag_BLOCK_ID_FLAG_UNKNOWN BlockIDFlag = 0
	BlockIDFlag_BLOCK_ID_FLAG_ABSENT  BlockIDFlag = 1
	BlockIDFlag_BLOCK_ID_FLAG_COMMIT  BlockIDFlag = 2
	BlockIDFlag_BLOCK_ID_FLAG_NIL     BlockIDFlag = 3
)

// Enum value maps for BlockIDFlag.
var (
	BlockIDFlag_name = map[int32]string{
		0: "BLOCK_ID_FLAG_UNKNOWN",
		1: "BLOCK_ID_FLAG_ABSENT",
		2: "BLOCK_ID_FLAG_COMMIT",
		3: "BLOCK_ID_FLAG_NIL",
	}
	BlockIDFlag_value = map[string]int32{
		"BLOCK_ID_FLAG_UNKNOWN": 0,
		"BLOCK_ID_FLAG_ABSENT":  1,
		"BLOCK_ID_FLAG_COMMIT":  2,
		"BLOCK_ID_FLAG_NIL":     3,
	}
)

func (x BlockIDFlag) Enum() *BlockIDFlag {
	p := new(BlockIDFlag)
	*p = x
	return p
}

func (x BlockIDFlag) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BlockIDFlag) Descriptor() protoreflect.EnumDescriptor {
	return file_tendermint_types_types_proto_enumTypes[0].Descriptor()
}

func (BlockIDFlag) Type() protoreflect.EnumType {
	return &file_tendermint_types_types_proto_enumTypes[0]
}

func (x BlockIDFlag) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BlockIDFlag.Descriptor instead.
func (BlockIDFlag) EnumDescriptor() ([]byte, []int) {
	return file_tendermint_types_types_proto_rawDescGZIP(), []int{0}
}

// SignedMsgType is a type of signed message in the consensus.
type SignedMsgType int32

const (
	SignedMsgType_SIGNED_MSG_TYPE_UNKNOWN SignedMsgType = 0
	// Votes
	SignedMsgType_SIGNED_MSG_TYPE_PREVOTE   SignedMsgType = 1
	SignedMsgType_SIGNED_MSG_TYPE_PRECOMMIT SignedMsgType = 2
	// Proposals
	SignedMsgType_SIGNED_MSG_TYPE_PROPOSAL SignedMsgType = 32
)

// Enum value maps for SignedMsgType.
var (
	SignedMsgType_name = map[int32]string{
		0:  "SIGNED_MSG_TYPE_UNKNOWN",
		1:  "SIGNED_MSG_TYPE_PREVOTE",
		2:  "SIGNED_MSG_TYPE_PRECOMMIT",
		32: "SIGNED_MSG_TYPE_PROPOSAL",
	}
	SignedMsgType_value = map[string]int32{
		"SIGNED_MSG_TYPE_UNKNOWN":   0,
		"SIGNED_MSG_TYPE_PREVOTE":   1,
		"SIGNED_MSG_TYPE_PRECOMMIT": 2,
		"SIGNED_MSG_TYPE_PROPOSAL":  32,
	}
)

func (x SignedMsgType) Enum() *SignedMsgType {
	p := new(SignedMsgType)
	*p = x
	return p
}

func (x SignedMsgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SignedMsgType) Descriptor() protoreflect.EnumDescriptor {
	return file_tendermint_types_types_proto_enumTypes[1].Descriptor()
}

func (SignedMsgType) Type() protoreflect.EnumType {
	return &file_tendermint_types_types_proto_enumTypes[1]
}

func (x SignedMsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SignedMsgType.Descriptor instead.
func (SignedMsgType) EnumDescriptor() ([]byte, []int) {
	return file_tendermint_types_types_proto_rawDescGZIP(), []int{1}
}

// PartsetHeader
type PartSetHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total uint32 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Hash  []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *PartSetHeader) Reset() {
	*x = PartSetHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartSetHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartSetHeader) ProtoMessage() {}

// Deprecated: Use PartSetHeader.ProtoReflect.Descriptor instead.
func (*PartSetHeader) Descriptor() ([]byte, []int) {
	return file_tendermint_types_types_proto_rawDescGZIP(), []int{0}
}

func (x *PartSetHeader) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PartSetHeader) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

type Part struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index uint32        `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Bytes []byte        `protobuf:"bytes,2,opt,name=bytes,proto3" json:"bytes,omitempty"`
	Proof *crypto.Proof `protobuf:"bytes,3,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (x *Part) Reset() {
	*x = Part{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Part) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Part) ProtoMessage() {}

// Deprecated: Use Part.ProtoReflect.Descriptor instead.
func (*Part) Descriptor() ([]byte, []int) {
	return file_tendermint_types_types_proto_rawDescGZIP(), []int{1}
}

func (x *Part) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Part) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

func (x *Part) GetProof() *crypto.Proof {
	if x != nil {
		return x.Proof
	}
	return nil
}

// BlockID
type BlockID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash          []byte         `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	PartSetHeader *PartSetHeader `protobuf:"bytes,2,opt,name=part_set_header,json=partSetHeader,proto3" json:"part_set_header,omitempty"`
}

func (x *BlockID) Reset() {
	*x = BlockID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockID) ProtoMessage() {}

// Deprecated: Use BlockID.ProtoReflect.Descriptor instead.
func (*BlockID) Descriptor() ([]byte, []int) {
	return file_tendermint_types_types_proto_rawDescGZIP(), []int{2}
}

func (x *BlockID) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *BlockID) GetPartSetHeader() *PartSetHeader {
	if x != nil {
		return x.PartSetHeader
	}
	return nil
}

// Header defines the structure of a Tendermint block header.
type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// basic block info
	Version *version.Consensus     `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	ChainId string                 `protobuf:"bytes,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Height  int64                  `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Time    *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=time,proto3" json:"time,omitempty"`
	// prev block info
	LastBlockId *BlockID `protobuf:"bytes,5,opt,name=last_block_id,json=lastBlockId,proto3" json:"last_block_id,omitempty"`
	// hashes of block data
	LastCommitHash []byte `protobuf:"bytes,6,opt,name=last_commit_hash,json=lastCommitHash,proto3" json:"last_commit_hash,omitempty"` // commit from validators from the last block
	DataHash       []byte `protobuf:"bytes,7,opt,name=data_hash,json=dataHash,proto3" json:"data_hash,omitempty"`                     // transactions
	// hashes from the app output from the prev block
	ValidatorsHash     []byte `protobuf:"bytes,8,opt,name=validators_hash,json=validatorsHash,proto3" json:"validators_hash,omitempty"`               // validators for the current block
	NextValidatorsHash []byte `protobuf:"bytes,9,opt,name=next_validators_hash,json=nextValidatorsHash,proto3" json:"next_validators_hash,omitempty"` // validators for the next block
	ConsensusHash      []byte `protobuf:"bytes,10,opt,name=consensus_hash,json=consensusHash,proto3" json:"consensus_hash,omitempty"`                 // consensus params for current block
	AppHash            []byte `protobuf:"bytes,11,opt,name=app_hash,json=appHash,proto3" json:"app_hash,omitempty"`                                   // state after txs from the previous block
	LastResultsHash    []byte `protobuf:"bytes,12,opt,name=last_results_hash,json=lastResultsHash,proto3" json:"last_results_hash,omitempty"`         // root hash of all results from the txs from the previous block
	// consensus info
	EvidenceHash    []byte `protobuf:"bytes,13,opt,name=evidence_hash,json=evidenceHash,proto3" json:"evidence_hash,omitempty"`          // evidence included in the block
	ProposerAddress []byte `protobuf:"bytes,14,opt,name=proposer_address,json=proposerAddress,proto3" json:"proposer_address,omitempty"` // original proposer of the block
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_tendermint_types_types_proto_rawDescGZIP(), []int{3}
}

func (x *Header) GetVersion() *version.Consensus {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *Header) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

func (x *Header) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Header) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Header) GetLastBlockId() *BlockID {
	if x != nil {
		return x.LastBlockId
	}
	return nil
}

func (x *Header) GetLastCommitHash() []byte {
	if x != nil {
		return x.LastCommitHash
	}
	return nil
}

func (x *Header) GetDataHash() []byte {
	if x != nil {
		return x.DataHash
	}
	return nil
}

func (x *Header) GetValidatorsHash() []byte {
	if x != nil {
		return x.ValidatorsHash
	}
	return nil
}

func (x *Header) GetNextValidatorsHash() []byte {
	if x != nil {
		return x.NextValidatorsHash
	}
	return nil
}

func (x *Header) GetConsensusHash() []byte {
	if x != nil {
		return x.ConsensusHash
	}
	return nil
}

func (x *Header) GetAppHash() []byte {
	if x != nil {
		return x.AppHash
	}
	return nil
}

func (x *Header) GetLastResultsHash() []byte {
	if x != nil {
		return x.LastResultsHash
	}
	return nil
}

func (x *Header) GetEvidenceHash() []byte {
	if x != nil {
		return x.EvidenceHash
	}
	return nil
}

func (x *Header) GetProposerAddress() []byte {
	if x != nil {
		return x.ProposerAddress
	}
	return nil
}

// Data contains the set of transactions included in the block
type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Txs that will be applied by state @ block.Height+1.
	// NOTE: not all txs here are valid.  We're just agreeing on the order first.
	// This means that block.AppHash does not include these txs.
	Txs [][]byte `protobuf:"bytes,1,rep,name=txs,proto3" json:"txs,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_tendermint_types_types_proto_rawDescGZIP(), []int{4}
}

func (x *Data) GetTxs() [][]byte {
	if x != nil {
		return x.Txs
	}
	return nil
}

// Vote represents a prevote, precommit, or commit vote from validators for
// consensus.
type Vote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type_            SignedMsgType          `protobuf:"varint,1,opt,name=type,proto3,enum=tendermint.types.SignedMsgType" json:"type,omitempty"`
	Height           int64                  `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Round            int32                  `protobuf:"varint,3,opt,name=round,proto3" json:"round,omitempty"`
	BlockId          *BlockID               `protobuf:"bytes,4,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"` // zero if vote is nil.
	Timestamp        *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ValidatorAddress []byte                 `protobuf:"bytes,6,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	ValidatorIndex   int32                  `protobuf:"varint,7,opt,name=validator_index,json=validatorIndex,proto3" json:"validator_index,omitempty"`
	Signature        []byte                 `protobuf:"bytes,8,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Vote) Reset() {
	*x = Vote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vote) ProtoMessage() {}

// Deprecated: Use Vote.ProtoReflect.Descriptor instead.
func (*Vote) Descriptor() ([]byte, []int) {
	return file_tendermint_types_types_proto_rawDescGZIP(), []int{5}
}

func (x *Vote) GetType_() SignedMsgType {
	if x != nil {
		return x.Type_
	}
	return SignedMsgType_SIGNED_MSG_TYPE_UNKNOWN
}

func (x *Vote) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Vote) GetRound() int32 {
	if x != nil {
		return x.Round
	}
	return 0
}

func (x *Vote) GetBlockId() *BlockID {
	if x != nil {
		return x.BlockId
	}
	return nil
}

func (x *Vote) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Vote) GetValidatorAddress() []byte {
	if x != nil {
		return x.ValidatorAddress
	}
	return nil
}

func (x *Vote) GetValidatorIndex() int32 {
	if x != nil {
		return x.ValidatorIndex
	}
	return 0
}

func (x *Vote) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

// Commit contains the evidence that a block was committed by a set of validators.
type Commit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height     int64        `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Round      int32        `protobuf:"varint,2,opt,name=round,proto3" json:"round,omitempty"`
	BlockId    *BlockID     `protobuf:"bytes,3,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	Signatures []*CommitSig `protobuf:"bytes,4,rep,name=signatures,proto3" json:"signatures,omitempty"`
}

func (x *Commit) Reset() {
	*x = Commit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_types_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Commit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commit) ProtoMessage() {}

// Deprecated: Use Commit.ProtoReflect.Descriptor instead.
func (*Commit) Descriptor() ([]byte, []int) {
	return file_tendermint_types_types_proto_rawDescGZIP(), []int{6}
}

func (x *Commit) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Commit) GetRound() int32 {
	if x != nil {
		return x.Round
	}
	return 0
}

func (x *Commit) GetBlockId() *BlockID {
	if x != nil {
		return x.BlockId
	}
	return nil
}

func (x *Commit) GetSignatures() []*CommitSig {
	if x != nil {
		return x.Signatures
	}
	return nil
}

// CommitSig is a part of the Vote included in a Commit.
type CommitSig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockIdFlag      BlockIDFlag            `protobuf:"varint,1,opt,name=block_id_flag,json=blockIdFlag,proto3,enum=tendermint.types.BlockIDFlag" json:"block_id_flag,omitempty"`
	ValidatorAddress []byte                 `protobuf:"bytes,2,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	Timestamp        *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Signature        []byte                 `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *CommitSig) Reset() {
	*x = CommitSig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_types_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitSig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitSig) ProtoMessage() {}

// Deprecated: Use CommitSig.ProtoReflect.Descriptor instead.
func (*CommitSig) Descriptor() ([]byte, []int) {
	return file_tendermint_types_types_proto_rawDescGZIP(), []int{7}
}

func (x *CommitSig) GetBlockIdFlag() BlockIDFlag {
	if x != nil {
		return x.BlockIdFlag
	}
	return BlockIDFlag_BLOCK_ID_FLAG_UNKNOWN
}

func (x *CommitSig) GetValidatorAddress() []byte {
	if x != nil {
		return x.ValidatorAddress
	}
	return nil
}

func (x *CommitSig) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *CommitSig) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type Proposal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type_     SignedMsgType          `protobuf:"varint,1,opt,name=type,proto3,enum=tendermint.types.SignedMsgType" json:"type,omitempty"`
	Height    int64                  `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Round     int32                  `protobuf:"varint,3,opt,name=round,proto3" json:"round,omitempty"`
	PolRound  int32                  `protobuf:"varint,4,opt,name=pol_round,json=polRound,proto3" json:"pol_round,omitempty"`
	BlockId   *BlockID               `protobuf:"bytes,5,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Signature []byte                 `protobuf:"bytes,7,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Proposal) Reset() {
	*x = Proposal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_types_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Proposal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Proposal) ProtoMessage() {}

// Deprecated: Use Proposal.ProtoReflect.Descriptor instead.
func (*Proposal) Descriptor() ([]byte, []int) {
	return file_tendermint_types_types_proto_rawDescGZIP(), []int{8}
}

func (x *Proposal) GetType_() SignedMsgType {
	if x != nil {
		return x.Type_
	}
	return SignedMsgType_SIGNED_MSG_TYPE_UNKNOWN
}

func (x *Proposal) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Proposal) GetRound() int32 {
	if x != nil {
		return x.Round
	}
	return 0
}

func (x *Proposal) GetPolRound() int32 {
	if x != nil {
		return x.PolRound
	}
	return 0
}

func (x *Proposal) GetBlockId() *BlockID {
	if x != nil {
		return x.BlockId
	}
	return nil
}

func (x *Proposal) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Proposal) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type SignedHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Commit *Commit `protobuf:"bytes,2,opt,name=commit,proto3" json:"commit,omitempty"`
}

func (x *SignedHeader) Reset() {
	*x = SignedHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_types_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedHeader) ProtoMessage() {}

// Deprecated: Use SignedHeader.ProtoReflect.Descriptor instead.
func (*SignedHeader) Descriptor() ([]byte, []int) {
	return file_tendermint_types_types_proto_rawDescGZIP(), []int{9}
}

func (x *SignedHeader) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *SignedHeader) GetCommit() *Commit {
	if x != nil {
		return x.Commit
	}
	return nil
}

type LightBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SignedHeader *SignedHeader `protobuf:"bytes,1,opt,name=signed_header,json=signedHeader,proto3" json:"signed_header,omitempty"`
	ValidatorSet *ValidatorSet `protobuf:"bytes,2,opt,name=validator_set,json=validatorSet,proto3" json:"validator_set,omitempty"`
}

func (x *LightBlock) Reset() {
	*x = LightBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_types_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LightBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LightBlock) ProtoMessage() {}

// Deprecated: Use LightBlock.ProtoReflect.Descriptor instead.
func (*LightBlock) Descriptor() ([]byte, []int) {
	return file_tendermint_types_types_proto_rawDescGZIP(), []int{10}
}

func (x *LightBlock) GetSignedHeader() *SignedHeader {
	if x != nil {
		return x.SignedHeader
	}
	return nil
}

func (x *LightBlock) GetValidatorSet() *ValidatorSet {
	if x != nil {
		return x.ValidatorSet
	}
	return nil
}

type BlockMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockId   *BlockID `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	BlockSize int64    `protobuf:"varint,2,opt,name=block_size,json=blockSize,proto3" json:"block_size,omitempty"`
	Header    *Header  `protobuf:"bytes,3,opt,name=header,proto3" json:"header,omitempty"`
	NumTxs    int64    `protobuf:"varint,4,opt,name=num_txs,json=numTxs,proto3" json:"num_txs,omitempty"`
}

func (x *BlockMeta) Reset() {
	*x = BlockMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_types_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockMeta) ProtoMessage() {}

// Deprecated: Use BlockMeta.ProtoReflect.Descriptor instead.
func (*BlockMeta) Descriptor() ([]byte, []int) {
	return file_tendermint_types_types_proto_rawDescGZIP(), []int{11}
}

func (x *BlockMeta) GetBlockId() *BlockID {
	if x != nil {
		return x.BlockId
	}
	return nil
}

func (x *BlockMeta) GetBlockSize() int64 {
	if x != nil {
		return x.BlockSize
	}
	return 0
}

func (x *BlockMeta) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *BlockMeta) GetNumTxs() int64 {
	if x != nil {
		return x.NumTxs
	}
	return 0
}

// TxProof represents a Merkle proof of the presence of a transaction in the Merkle tree.
type TxProof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RootHash []byte        `protobuf:"bytes,1,opt,name=root_hash,json=rootHash,proto3" json:"root_hash,omitempty"`
	Data     []byte        `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Proof    *crypto.Proof `protobuf:"bytes,3,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (x *TxProof) Reset() {
	*x = TxProof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tendermint_types_types_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxProof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxProof) ProtoMessage() {}

// Deprecated: Use TxProof.ProtoReflect.Descriptor instead.
func (*TxProof) Descriptor() ([]byte, []int) {
	return file_tendermint_types_types_proto_rawDescGZIP(), []int{12}
}

func (x *TxProof) GetRootHash() []byte {
	if x != nil {
		return x.RootHash
	}
	return nil
}

func (x *TxProof) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *TxProof) GetProof() *crypto.Proof {
	if x != nil {
		return x.Proof
	}
	return nil
}

var File_tendermint_types_types_proto protoreflect.FileDescriptor

var file_tendermint_types_types_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x74, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6f, 0x66,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x74, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x0d, 0x50, 0x61, 0x72, 0x74,
	0x53, 0x65, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x22, 0x68, 0x0a, 0x04, 0x50, 0x61, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x74, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x6f, 0x66,
	0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x22, 0x6c, 0x0a,
	0x07, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x4d, 0x0a, 0x0f,
	0x70, 0x61, 0x72, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x53, 0x65, 0x74,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x0d, 0x70, 0x61,
	0x72, 0x74, 0x53, 0x65, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0xe6, 0x04, 0x0a, 0x06,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e,
	0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xe2, 0xde, 0x1f, 0x07, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x49, 0x44, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x38, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x08, 0xc8, 0xde, 0x1f, 0x00, 0x90, 0xdf, 0x1f, 0x01, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x43, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49,
	0x44, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e,
	0x6c, 0x61, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1b,
	0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x48, 0x61, 0x73, 0x68, 0x12, 0x27, 0x0a, 0x0f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x30, 0x0a, 0x14, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x12, 0x6e, 0x65, 0x78, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x73, 0x48, 0x61, 0x73, 0x68, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x73, 0x75, 0x73, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d,
	0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x48, 0x61, 0x73, 0x68, 0x12, 0x19, 0x0a,
	0x08, 0x61, 0x70, 0x70, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x61, 0x70, 0x70, 0x48, 0x61, 0x73, 0x68, 0x12, 0x2a, 0x0a, 0x11, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x65, 0x76, 0x69,
	0x64, 0x65, 0x6e, 0x63, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x22, 0x18, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03,
	0x74, 0x78, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x03, 0x74, 0x78, 0x73, 0x22, 0xe8,
	0x02, 0x0a, 0x04, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4d,
	0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x45, 0x0a, 0x08, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x42, 0x0f, 0xc8, 0xde, 0x1f, 0x00, 0xe2, 0xde, 0x1f,
	0x07, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49,
	0x64, 0x12, 0x42, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x42, 0x08, 0xc8, 0xde, 0x1f, 0x00, 0x90, 0xdf, 0x1f, 0x01, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2b, 0x0a, 0x11, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x10, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0xc0, 0x01, 0x0a, 0x06, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x12, 0x45, 0x0a, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x42,
	0x0f, 0xc8, 0xde, 0x1f, 0x00, 0xe2, 0xde, 0x1f, 0x07, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44,
	0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x0a, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x53, 0x69, 0x67, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00,
	0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x22, 0xdd, 0x01, 0x0a,
	0x09, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x53, 0x69, 0x67, 0x12, 0x41, 0x0a, 0x0d, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x46, 0x6c, 0x61, 0x67,
	0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x2b, 0x0a,
	0x11, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x42, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xc8, 0xde, 0x1f, 0x00, 0x90,
	0xdf, 0x1f, 0x01, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0xb3, 0x02, 0x0a,
	0x08, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x12, 0x33, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x6f, 0x6c, 0x5f, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x70, 0x6f, 0x6c, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x45, 0x0a, 0x08, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x42, 0x0f, 0xc8, 0xde, 0x1f, 0x00, 0xe2, 0xde, 0x1f, 0x07,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64,
	0x12, 0x42, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x08, 0xc8, 0xde, 0x1f, 0x00, 0x90, 0xdf, 0x1f, 0x01, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x22, 0x72, 0x0a, 0x0c, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x30, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x06,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x22, 0x96, 0x01, 0x0a, 0x0a, 0x4c, 0x69, 0x67, 0x68, 0x74,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x43, 0x0a, 0x0d, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x0c, 0x73, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x0d, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65,
	0x74, 0x52, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x22,
	0xc2, 0x01, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x45, 0x0a,
	0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x42, 0x0f, 0xc8, 0xde, 0x1f, 0x00,
	0xe2, 0xde, 0x1f, 0x07, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x52, 0x07, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x04, 0xc8,
	0xde, 0x1f, 0x00, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x6e,
	0x75, 0x6d, 0x5f, 0x74, 0x78, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75,
	0x6d, 0x54, 0x78, 0x73, 0x22, 0x6a, 0x0a, 0x07, 0x54, 0x78, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12,
	0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x2e, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66,
	0x2a, 0xd7, 0x01, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x46, 0x6c, 0x61, 0x67,
	0x12, 0x31, 0x0a, 0x15, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x49, 0x44, 0x5f, 0x46, 0x4c, 0x41,
	0x47, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x1a, 0x16, 0x8a, 0x9d, 0x20,
	0x12, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x46, 0x6c, 0x61, 0x67, 0x55, 0x6e, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x12, 0x2f, 0x0a, 0x14, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x49, 0x44, 0x5f,
	0x46, 0x4c, 0x41, 0x47, 0x5f, 0x41, 0x42, 0x53, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x1a, 0x15, 0x8a,
	0x9d, 0x20, 0x11, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x46, 0x6c, 0x61, 0x67, 0x41, 0x62,
	0x73, 0x65, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x14, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x49, 0x44,
	0x5f, 0x46, 0x4c, 0x41, 0x47, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x54, 0x10, 0x02, 0x1a, 0x15,
	0x8a, 0x9d, 0x20, 0x11, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x46, 0x6c, 0x61, 0x67, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x29, 0x0a, 0x11, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x49,
	0x44, 0x5f, 0x46, 0x4c, 0x41, 0x47, 0x5f, 0x4e, 0x49, 0x4c, 0x10, 0x03, 0x1a, 0x12, 0x8a, 0x9d,
	0x20, 0x0e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x46, 0x6c, 0x61, 0x67, 0x4e, 0x69, 0x6c,
	0x1a, 0x08, 0x88, 0xa3, 0x1e, 0x00, 0xa8, 0xa4, 0x1e, 0x01, 0x2a, 0xd7, 0x01, 0x0a, 0x0d, 0x53,
	0x69, 0x67, 0x6e, 0x65, 0x64, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x17,
	0x53, 0x49, 0x47, 0x4e, 0x45, 0x44, 0x5f, 0x4d, 0x53, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x1a, 0x0f, 0x8a, 0x9d, 0x20, 0x0b, 0x55,
	0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x17, 0x53, 0x49,
	0x47, 0x4e, 0x45, 0x44, 0x5f, 0x4d, 0x53, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x52,
	0x45, 0x56, 0x4f, 0x54, 0x45, 0x10, 0x01, 0x1a, 0x0f, 0x8a, 0x9d, 0x20, 0x0b, 0x50, 0x72, 0x65,
	0x76, 0x6f, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x19, 0x53, 0x49, 0x47, 0x4e,
	0x45, 0x44, 0x5f, 0x4d, 0x53, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x52, 0x45, 0x43,
	0x4f, 0x4d, 0x4d, 0x49, 0x54, 0x10, 0x02, 0x1a, 0x11, 0x8a, 0x9d, 0x20, 0x0d, 0x50, 0x72, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x18, 0x53, 0x49,
	0x47, 0x4e, 0x45, 0x44, 0x5f, 0x4d, 0x53, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x52,
	0x4f, 0x50, 0x4f, 0x53, 0x41, 0x4c, 0x10, 0x20, 0x1a, 0x10, 0x8a, 0x9d, 0x20, 0x0c, 0x50, 0x72,
	0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x08, 0x88, 0xa3, 0x1e, 0x00,
	0xa8, 0xa4, 0x1e, 0x01, 0x42, 0xb6, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x42, 0x0a, 0x54,
	0x79, 0x70, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0xa2, 0x02,
	0x03, 0x54, 0x54, 0x58, 0xaa, 0x02, 0x10, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0xca, 0x02, 0x10, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x74, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x73, 0xe2, 0x02, 0x1c, 0x54, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x54, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x3a, 0x3a, 0x54, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tendermint_types_types_proto_rawDescOnce sync.Once
	file_tendermint_types_types_proto_rawDescData = file_tendermint_types_types_proto_rawDesc
)

func file_tendermint_types_types_proto_rawDescGZIP() []byte {
	file_tendermint_types_types_proto_rawDescOnce.Do(func() {
		file_tendermint_types_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_tendermint_types_types_proto_rawDescData)
	})
	return file_tendermint_types_types_proto_rawDescData
}

var file_tendermint_types_types_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_tendermint_types_types_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_tendermint_types_types_proto_goTypes = []interface{}{
	(BlockIDFlag)(0),              // 0: tendermint.types.BlockIDFlag
	(SignedMsgType)(0),            // 1: tendermint.types.SignedMsgType
	(*PartSetHeader)(nil),         // 2: tendermint.types.PartSetHeader
	(*Part)(nil),                  // 3: tendermint.types.Part
	(*BlockID)(nil),               // 4: tendermint.types.BlockID
	(*Header)(nil),                // 5: tendermint.types.Header
	(*Data)(nil),                  // 6: tendermint.types.Data
	(*Vote)(nil),                  // 7: tendermint.types.Vote
	(*Commit)(nil),                // 8: tendermint.types.Commit
	(*CommitSig)(nil),             // 9: tendermint.types.CommitSig
	(*Proposal)(nil),              // 10: tendermint.types.Proposal
	(*SignedHeader)(nil),          // 11: tendermint.types.SignedHeader
	(*LightBlock)(nil),            // 12: tendermint.types.LightBlock
	(*BlockMeta)(nil),             // 13: tendermint.types.BlockMeta
	(*TxProof)(nil),               // 14: tendermint.types.TxProof
	(*crypto.Proof)(nil),          // 15: tendermint.crypto.Proof
	(*version.Consensus)(nil),     // 16: tendermint.version.Consensus
	(*timestamppb.Timestamp)(nil), // 17: google.protobuf.Timestamp
	(*ValidatorSet)(nil),          // 18: tendermint.types.ValidatorSet
}
var file_tendermint_types_types_proto_depIdxs = []int32{
	15, // 0: tendermint.types.Part.proof:type_name -> tendermint.crypto.Proof
	2,  // 1: tendermint.types.BlockID.part_set_header:type_name -> tendermint.types.PartSetHeader
	16, // 2: tendermint.types.Header.version:type_name -> tendermint.version.Consensus
	17, // 3: tendermint.types.Header.time:type_name -> google.protobuf.Timestamp
	4,  // 4: tendermint.types.Header.last_block_id:type_name -> tendermint.types.BlockID
	1,  // 5: tendermint.types.Vote.type:type_name -> tendermint.types.SignedMsgType
	4,  // 6: tendermint.types.Vote.block_id:type_name -> tendermint.types.BlockID
	17, // 7: tendermint.types.Vote.timestamp:type_name -> google.protobuf.Timestamp
	4,  // 8: tendermint.types.Commit.block_id:type_name -> tendermint.types.BlockID
	9,  // 9: tendermint.types.Commit.signatures:type_name -> tendermint.types.CommitSig
	0,  // 10: tendermint.types.CommitSig.block_id_flag:type_name -> tendermint.types.BlockIDFlag
	17, // 11: tendermint.types.CommitSig.timestamp:type_name -> google.protobuf.Timestamp
	1,  // 12: tendermint.types.Proposal.type:type_name -> tendermint.types.SignedMsgType
	4,  // 13: tendermint.types.Proposal.block_id:type_name -> tendermint.types.BlockID
	17, // 14: tendermint.types.Proposal.timestamp:type_name -> google.protobuf.Timestamp
	5,  // 15: tendermint.types.SignedHeader.header:type_name -> tendermint.types.Header
	8,  // 16: tendermint.types.SignedHeader.commit:type_name -> tendermint.types.Commit
	11, // 17: tendermint.types.LightBlock.signed_header:type_name -> tendermint.types.SignedHeader
	18, // 18: tendermint.types.LightBlock.validator_set:type_name -> tendermint.types.ValidatorSet
	4,  // 19: tendermint.types.BlockMeta.block_id:type_name -> tendermint.types.BlockID
	5,  // 20: tendermint.types.BlockMeta.header:type_name -> tendermint.types.Header
	15, // 21: tendermint.types.TxProof.proof:type_name -> tendermint.crypto.Proof
	22, // [22:22] is the sub-list for method output_type
	22, // [22:22] is the sub-list for method input_type
	22, // [22:22] is the sub-list for extension type_name
	22, // [22:22] is the sub-list for extension extendee
	0,  // [0:22] is the sub-list for field type_name
}

func init() { file_tendermint_types_types_proto_init() }
func file_tendermint_types_types_proto_init() {
	if File_tendermint_types_types_proto != nil {
		return
	}
	file_tendermint_types_validator_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tendermint_types_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartSetHeader); i {
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
		file_tendermint_types_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Part); i {
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
		file_tendermint_types_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockID); i {
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
		file_tendermint_types_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
		file_tendermint_types_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_tendermint_types_types_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vote); i {
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
		file_tendermint_types_types_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Commit); i {
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
		file_tendermint_types_types_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitSig); i {
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
		file_tendermint_types_types_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Proposal); i {
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
		file_tendermint_types_types_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedHeader); i {
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
		file_tendermint_types_types_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LightBlock); i {
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
		file_tendermint_types_types_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockMeta); i {
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
		file_tendermint_types_types_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxProof); i {
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
			RawDescriptor: file_tendermint_types_types_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tendermint_types_types_proto_goTypes,
		DependencyIndexes: file_tendermint_types_types_proto_depIdxs,
		EnumInfos:         file_tendermint_types_types_proto_enumTypes,
		MessageInfos:      file_tendermint_types_types_proto_msgTypes,
	}.Build()
	File_tendermint_types_types_proto = out.File
	file_tendermint_types_types_proto_rawDesc = nil
	file_tendermint_types_types_proto_goTypes = nil
	file_tendermint_types_types_proto_depIdxs = nil
}
