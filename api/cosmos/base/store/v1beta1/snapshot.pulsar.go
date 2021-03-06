package storev1beta1

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_SnapshotItem       protoreflect.MessageDescriptor
	fd_SnapshotItem_store protoreflect.FieldDescriptor
	fd_SnapshotItem_iavl  protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_store_v1beta1_snapshot_proto_init()
	md_SnapshotItem = File_cosmos_base_store_v1beta1_snapshot_proto.Messages().ByName("SnapshotItem")
	fd_SnapshotItem_store = md_SnapshotItem.Fields().ByName("store")
	fd_SnapshotItem_iavl = md_SnapshotItem.Fields().ByName("iavl")
}

var _ protoreflect.Message = (*fastReflection_SnapshotItem)(nil)

type fastReflection_SnapshotItem SnapshotItem

func (x *SnapshotItem) ProtoReflect() protoreflect.Message {
	return (*fastReflection_SnapshotItem)(x)
}

func (x *SnapshotItem) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_store_v1beta1_snapshot_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_SnapshotItem_messageType fastReflection_SnapshotItem_messageType
var _ protoreflect.MessageType = fastReflection_SnapshotItem_messageType{}

type fastReflection_SnapshotItem_messageType struct{}

func (x fastReflection_SnapshotItem_messageType) Zero() protoreflect.Message {
	return (*fastReflection_SnapshotItem)(nil)
}
func (x fastReflection_SnapshotItem_messageType) New() protoreflect.Message {
	return new(fastReflection_SnapshotItem)
}
func (x fastReflection_SnapshotItem_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_SnapshotItem
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_SnapshotItem) Descriptor() protoreflect.MessageDescriptor {
	return md_SnapshotItem
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_SnapshotItem) Type() protoreflect.MessageType {
	return _fastReflection_SnapshotItem_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_SnapshotItem) New() protoreflect.Message {
	return new(fastReflection_SnapshotItem)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_SnapshotItem) Interface() protoreflect.ProtoMessage {
	return (*SnapshotItem)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_SnapshotItem) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Item != nil {
		switch o := x.Item.(type) {
		case *SnapshotItem_Store:
			v := o.Store
			value := protoreflect.ValueOfMessage(v.ProtoReflect())
			if !f(fd_SnapshotItem_store, value) {
				return
			}
		case *SnapshotItem_Iavl:
			v := o.Iavl
			value := protoreflect.ValueOfMessage(v.ProtoReflect())
			if !f(fd_SnapshotItem_iavl, value) {
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
func (x *fastReflection_SnapshotItem) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.store.v1beta1.SnapshotItem.store":
		if x.Item == nil {
			return false
		} else if _, ok := x.Item.(*SnapshotItem_Store); ok {
			return true
		} else {
			return false
		}
	case "cosmos.base.store.v1beta1.SnapshotItem.iavl":
		if x.Item == nil {
			return false
		} else if _, ok := x.Item.(*SnapshotItem_Iavl); ok {
			return true
		} else {
			return false
		}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.store.v1beta1.SnapshotItem"))
		}
		panic(fmt.Errorf("message cosmos.base.store.v1beta1.SnapshotItem does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SnapshotItem) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.store.v1beta1.SnapshotItem.store":
		x.Item = nil
	case "cosmos.base.store.v1beta1.SnapshotItem.iavl":
		x.Item = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.store.v1beta1.SnapshotItem"))
		}
		panic(fmt.Errorf("message cosmos.base.store.v1beta1.SnapshotItem does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_SnapshotItem) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.store.v1beta1.SnapshotItem.store":
		if x.Item == nil {
			return protoreflect.ValueOfMessage((*SnapshotStoreItem)(nil).ProtoReflect())
		} else if v, ok := x.Item.(*SnapshotItem_Store); ok {
			return protoreflect.ValueOfMessage(v.Store.ProtoReflect())
		} else {
			return protoreflect.ValueOfMessage((*SnapshotStoreItem)(nil).ProtoReflect())
		}
	case "cosmos.base.store.v1beta1.SnapshotItem.iavl":
		if x.Item == nil {
			return protoreflect.ValueOfMessage((*SnapshotIAVLItem)(nil).ProtoReflect())
		} else if v, ok := x.Item.(*SnapshotItem_Iavl); ok {
			return protoreflect.ValueOfMessage(v.Iavl.ProtoReflect())
		} else {
			return protoreflect.ValueOfMessage((*SnapshotIAVLItem)(nil).ProtoReflect())
		}
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.store.v1beta1.SnapshotItem"))
		}
		panic(fmt.Errorf("message cosmos.base.store.v1beta1.SnapshotItem does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_SnapshotItem) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.store.v1beta1.SnapshotItem.store":
		cv := value.Message().Interface().(*SnapshotStoreItem)
		x.Item = &SnapshotItem_Store{Store: cv}
	case "cosmos.base.store.v1beta1.SnapshotItem.iavl":
		cv := value.Message().Interface().(*SnapshotIAVLItem)
		x.Item = &SnapshotItem_Iavl{Iavl: cv}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.store.v1beta1.SnapshotItem"))
		}
		panic(fmt.Errorf("message cosmos.base.store.v1beta1.SnapshotItem does not contain field %s", fd.FullName()))
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
func (x *fastReflection_SnapshotItem) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.store.v1beta1.SnapshotItem.store":
		if x.Item == nil {
			value := &SnapshotStoreItem{}
			oneofValue := &SnapshotItem_Store{Store: value}
			x.Item = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
		switch m := x.Item.(type) {
		case *SnapshotItem_Store:
			return protoreflect.ValueOfMessage(m.Store.ProtoReflect())
		default:
			value := &SnapshotStoreItem{}
			oneofValue := &SnapshotItem_Store{Store: value}
			x.Item = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
	case "cosmos.base.store.v1beta1.SnapshotItem.iavl":
		if x.Item == nil {
			value := &SnapshotIAVLItem{}
			oneofValue := &SnapshotItem_Iavl{Iavl: value}
			x.Item = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
		switch m := x.Item.(type) {
		case *SnapshotItem_Iavl:
			return protoreflect.ValueOfMessage(m.Iavl.ProtoReflect())
		default:
			value := &SnapshotIAVLItem{}
			oneofValue := &SnapshotItem_Iavl{Iavl: value}
			x.Item = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.store.v1beta1.SnapshotItem"))
		}
		panic(fmt.Errorf("message cosmos.base.store.v1beta1.SnapshotItem does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_SnapshotItem) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.store.v1beta1.SnapshotItem.store":
		value := &SnapshotStoreItem{}
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.base.store.v1beta1.SnapshotItem.iavl":
		value := &SnapshotIAVLItem{}
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.store.v1beta1.SnapshotItem"))
		}
		panic(fmt.Errorf("message cosmos.base.store.v1beta1.SnapshotItem does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_SnapshotItem) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	case "cosmos.base.store.v1beta1.SnapshotItem.item":
		if x.Item == nil {
			return nil
		}
		switch x.Item.(type) {
		case *SnapshotItem_Store:
			return x.Descriptor().Fields().ByName("store")
		case *SnapshotItem_Iavl:
			return x.Descriptor().Fields().ByName("iavl")
		}
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.store.v1beta1.SnapshotItem", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_SnapshotItem) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SnapshotItem) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_SnapshotItem) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_SnapshotItem) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*SnapshotItem)
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
		switch x := x.Item.(type) {
		case *SnapshotItem_Store:
			if x == nil {
				break
			}
			l = options.Size(x.Store)
			n += 1 + l + runtime.Sov(uint64(l))
		case *SnapshotItem_Iavl:
			if x == nil {
				break
			}
			l = options.Size(x.Iavl)
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
		x := input.Message.Interface().(*SnapshotItem)
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
		switch x := x.Item.(type) {
		case *SnapshotItem_Store:
			encoded, err := options.Marshal(x.Store)
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
		case *SnapshotItem_Iavl:
			encoded, err := options.Marshal(x.Iavl)
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
		x := input.Message.Interface().(*SnapshotItem)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SnapshotItem: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SnapshotItem: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Store", wireType)
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
				v := &SnapshotStoreItem{}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], v); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				x.Item = &SnapshotItem_Store{v}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Iavl", wireType)
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
				v := &SnapshotIAVLItem{}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], v); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				x.Item = &SnapshotItem_Iavl{v}
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
	md_SnapshotStoreItem      protoreflect.MessageDescriptor
	fd_SnapshotStoreItem_name protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_store_v1beta1_snapshot_proto_init()
	md_SnapshotStoreItem = File_cosmos_base_store_v1beta1_snapshot_proto.Messages().ByName("SnapshotStoreItem")
	fd_SnapshotStoreItem_name = md_SnapshotStoreItem.Fields().ByName("name")
}

var _ protoreflect.Message = (*fastReflection_SnapshotStoreItem)(nil)

type fastReflection_SnapshotStoreItem SnapshotStoreItem

func (x *SnapshotStoreItem) ProtoReflect() protoreflect.Message {
	return (*fastReflection_SnapshotStoreItem)(x)
}

func (x *SnapshotStoreItem) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_store_v1beta1_snapshot_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_SnapshotStoreItem_messageType fastReflection_SnapshotStoreItem_messageType
var _ protoreflect.MessageType = fastReflection_SnapshotStoreItem_messageType{}

type fastReflection_SnapshotStoreItem_messageType struct{}

func (x fastReflection_SnapshotStoreItem_messageType) Zero() protoreflect.Message {
	return (*fastReflection_SnapshotStoreItem)(nil)
}
func (x fastReflection_SnapshotStoreItem_messageType) New() protoreflect.Message {
	return new(fastReflection_SnapshotStoreItem)
}
func (x fastReflection_SnapshotStoreItem_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_SnapshotStoreItem
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_SnapshotStoreItem) Descriptor() protoreflect.MessageDescriptor {
	return md_SnapshotStoreItem
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_SnapshotStoreItem) Type() protoreflect.MessageType {
	return _fastReflection_SnapshotStoreItem_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_SnapshotStoreItem) New() protoreflect.Message {
	return new(fastReflection_SnapshotStoreItem)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_SnapshotStoreItem) Interface() protoreflect.ProtoMessage {
	return (*SnapshotStoreItem)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_SnapshotStoreItem) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Name != "" {
		value := protoreflect.ValueOfString(x.Name)
		if !f(fd_SnapshotStoreItem_name, value) {
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
func (x *fastReflection_SnapshotStoreItem) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.store.v1beta1.SnapshotStoreItem.name":
		return x.Name != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.store.v1beta1.SnapshotStoreItem"))
		}
		panic(fmt.Errorf("message cosmos.base.store.v1beta1.SnapshotStoreItem does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SnapshotStoreItem) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.store.v1beta1.SnapshotStoreItem.name":
		x.Name = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.store.v1beta1.SnapshotStoreItem"))
		}
		panic(fmt.Errorf("message cosmos.base.store.v1beta1.SnapshotStoreItem does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_SnapshotStoreItem) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.store.v1beta1.SnapshotStoreItem.name":
		value := x.Name
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.store.v1beta1.SnapshotStoreItem"))
		}
		panic(fmt.Errorf("message cosmos.base.store.v1beta1.SnapshotStoreItem does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_SnapshotStoreItem) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.store.v1beta1.SnapshotStoreItem.name":
		x.Name = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.store.v1beta1.SnapshotStoreItem"))
		}
		panic(fmt.Errorf("message cosmos.base.store.v1beta1.SnapshotStoreItem does not contain field %s", fd.FullName()))
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
func (x *fastReflection_SnapshotStoreItem) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.store.v1beta1.SnapshotStoreItem.name":
		panic(fmt.Errorf("field name of message cosmos.base.store.v1beta1.SnapshotStoreItem is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.store.v1beta1.SnapshotStoreItem"))
		}
		panic(fmt.Errorf("message cosmos.base.store.v1beta1.SnapshotStoreItem does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_SnapshotStoreItem) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.store.v1beta1.SnapshotStoreItem.name":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.store.v1beta1.SnapshotStoreItem"))
		}
		panic(fmt.Errorf("message cosmos.base.store.v1beta1.SnapshotStoreItem does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_SnapshotStoreItem) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.store.v1beta1.SnapshotStoreItem", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_SnapshotStoreItem) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SnapshotStoreItem) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_SnapshotStoreItem) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_SnapshotStoreItem) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*SnapshotStoreItem)
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
		l = len(x.Name)
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
		x := input.Message.Interface().(*SnapshotStoreItem)
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
		if len(x.Name) > 0 {
			i -= len(x.Name)
			copy(dAtA[i:], x.Name)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Name)))
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
		x := input.Message.Interface().(*SnapshotStoreItem)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SnapshotStoreItem: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SnapshotStoreItem: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
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
				x.Name = string(dAtA[iNdEx:postIndex])
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
	md_SnapshotIAVLItem         protoreflect.MessageDescriptor
	fd_SnapshotIAVLItem_key     protoreflect.FieldDescriptor
	fd_SnapshotIAVLItem_value   protoreflect.FieldDescriptor
	fd_SnapshotIAVLItem_version protoreflect.FieldDescriptor
	fd_SnapshotIAVLItem_height  protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_store_v1beta1_snapshot_proto_init()
	md_SnapshotIAVLItem = File_cosmos_base_store_v1beta1_snapshot_proto.Messages().ByName("SnapshotIAVLItem")
	fd_SnapshotIAVLItem_key = md_SnapshotIAVLItem.Fields().ByName("key")
	fd_SnapshotIAVLItem_value = md_SnapshotIAVLItem.Fields().ByName("value")
	fd_SnapshotIAVLItem_version = md_SnapshotIAVLItem.Fields().ByName("version")
	fd_SnapshotIAVLItem_height = md_SnapshotIAVLItem.Fields().ByName("height")
}

var _ protoreflect.Message = (*fastReflection_SnapshotIAVLItem)(nil)

type fastReflection_SnapshotIAVLItem SnapshotIAVLItem

func (x *SnapshotIAVLItem) ProtoReflect() protoreflect.Message {
	return (*fastReflection_SnapshotIAVLItem)(x)
}

func (x *SnapshotIAVLItem) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_store_v1beta1_snapshot_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_SnapshotIAVLItem_messageType fastReflection_SnapshotIAVLItem_messageType
var _ protoreflect.MessageType = fastReflection_SnapshotIAVLItem_messageType{}

type fastReflection_SnapshotIAVLItem_messageType struct{}

func (x fastReflection_SnapshotIAVLItem_messageType) Zero() protoreflect.Message {
	return (*fastReflection_SnapshotIAVLItem)(nil)
}
func (x fastReflection_SnapshotIAVLItem_messageType) New() protoreflect.Message {
	return new(fastReflection_SnapshotIAVLItem)
}
func (x fastReflection_SnapshotIAVLItem_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_SnapshotIAVLItem
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_SnapshotIAVLItem) Descriptor() protoreflect.MessageDescriptor {
	return md_SnapshotIAVLItem
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_SnapshotIAVLItem) Type() protoreflect.MessageType {
	return _fastReflection_SnapshotIAVLItem_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_SnapshotIAVLItem) New() protoreflect.Message {
	return new(fastReflection_SnapshotIAVLItem)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_SnapshotIAVLItem) Interface() protoreflect.ProtoMessage {
	return (*SnapshotIAVLItem)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_SnapshotIAVLItem) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Key) != 0 {
		value := protoreflect.ValueOfBytes(x.Key)
		if !f(fd_SnapshotIAVLItem_key, value) {
			return
		}
	}
	if len(x.Value) != 0 {
		value := protoreflect.ValueOfBytes(x.Value)
		if !f(fd_SnapshotIAVLItem_value, value) {
			return
		}
	}
	if x.Version != int64(0) {
		value := protoreflect.ValueOfInt64(x.Version)
		if !f(fd_SnapshotIAVLItem_version, value) {
			return
		}
	}
	if x.Height != int32(0) {
		value := protoreflect.ValueOfInt32(x.Height)
		if !f(fd_SnapshotIAVLItem_height, value) {
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
func (x *fastReflection_SnapshotIAVLItem) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.store.v1beta1.SnapshotIAVLItem.key":
		return len(x.Key) != 0
	case "cosmos.base.store.v1beta1.SnapshotIAVLItem.value":
		return len(x.Value) != 0
	case "cosmos.base.store.v1beta1.SnapshotIAVLItem.version":
		return x.Version != int64(0)
	case "cosmos.base.store.v1beta1.SnapshotIAVLItem.height":
		return x.Height != int32(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.store.v1beta1.SnapshotIAVLItem"))
		}
		panic(fmt.Errorf("message cosmos.base.store.v1beta1.SnapshotIAVLItem does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SnapshotIAVLItem) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.store.v1beta1.SnapshotIAVLItem.key":
		x.Key = nil
	case "cosmos.base.store.v1beta1.SnapshotIAVLItem.value":
		x.Value = nil
	case "cosmos.base.store.v1beta1.SnapshotIAVLItem.version":
		x.Version = int64(0)
	case "cosmos.base.store.v1beta1.SnapshotIAVLItem.height":
		x.Height = int32(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.store.v1beta1.SnapshotIAVLItem"))
		}
		panic(fmt.Errorf("message cosmos.base.store.v1beta1.SnapshotIAVLItem does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_SnapshotIAVLItem) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.store.v1beta1.SnapshotIAVLItem.key":
		value := x.Key
		return protoreflect.ValueOfBytes(value)
	case "cosmos.base.store.v1beta1.SnapshotIAVLItem.value":
		value := x.Value
		return protoreflect.ValueOfBytes(value)
	case "cosmos.base.store.v1beta1.SnapshotIAVLItem.version":
		value := x.Version
		return protoreflect.ValueOfInt64(value)
	case "cosmos.base.store.v1beta1.SnapshotIAVLItem.height":
		value := x.Height
		return protoreflect.ValueOfInt32(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.store.v1beta1.SnapshotIAVLItem"))
		}
		panic(fmt.Errorf("message cosmos.base.store.v1beta1.SnapshotIAVLItem does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_SnapshotIAVLItem) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.store.v1beta1.SnapshotIAVLItem.key":
		x.Key = value.Bytes()
	case "cosmos.base.store.v1beta1.SnapshotIAVLItem.value":
		x.Value = value.Bytes()
	case "cosmos.base.store.v1beta1.SnapshotIAVLItem.version":
		x.Version = value.Int()
	case "cosmos.base.store.v1beta1.SnapshotIAVLItem.height":
		x.Height = int32(value.Int())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.store.v1beta1.SnapshotIAVLItem"))
		}
		panic(fmt.Errorf("message cosmos.base.store.v1beta1.SnapshotIAVLItem does not contain field %s", fd.FullName()))
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
func (x *fastReflection_SnapshotIAVLItem) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.store.v1beta1.SnapshotIAVLItem.key":
		panic(fmt.Errorf("field key of message cosmos.base.store.v1beta1.SnapshotIAVLItem is not mutable"))
	case "cosmos.base.store.v1beta1.SnapshotIAVLItem.value":
		panic(fmt.Errorf("field value of message cosmos.base.store.v1beta1.SnapshotIAVLItem is not mutable"))
	case "cosmos.base.store.v1beta1.SnapshotIAVLItem.version":
		panic(fmt.Errorf("field version of message cosmos.base.store.v1beta1.SnapshotIAVLItem is not mutable"))
	case "cosmos.base.store.v1beta1.SnapshotIAVLItem.height":
		panic(fmt.Errorf("field height of message cosmos.base.store.v1beta1.SnapshotIAVLItem is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.store.v1beta1.SnapshotIAVLItem"))
		}
		panic(fmt.Errorf("message cosmos.base.store.v1beta1.SnapshotIAVLItem does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_SnapshotIAVLItem) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.store.v1beta1.SnapshotIAVLItem.key":
		return protoreflect.ValueOfBytes(nil)
	case "cosmos.base.store.v1beta1.SnapshotIAVLItem.value":
		return protoreflect.ValueOfBytes(nil)
	case "cosmos.base.store.v1beta1.SnapshotIAVLItem.version":
		return protoreflect.ValueOfInt64(int64(0))
	case "cosmos.base.store.v1beta1.SnapshotIAVLItem.height":
		return protoreflect.ValueOfInt32(int32(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.store.v1beta1.SnapshotIAVLItem"))
		}
		panic(fmt.Errorf("message cosmos.base.store.v1beta1.SnapshotIAVLItem does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_SnapshotIAVLItem) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.store.v1beta1.SnapshotIAVLItem", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_SnapshotIAVLItem) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SnapshotIAVLItem) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_SnapshotIAVLItem) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_SnapshotIAVLItem) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*SnapshotIAVLItem)
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
		l = len(x.Key)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Value)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Version != 0 {
			n += 1 + runtime.Sov(uint64(x.Version))
		}
		if x.Height != 0 {
			n += 1 + runtime.Sov(uint64(x.Height))
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
		x := input.Message.Interface().(*SnapshotIAVLItem)
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
		if x.Height != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Height))
			i--
			dAtA[i] = 0x20
		}
		if x.Version != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Version))
			i--
			dAtA[i] = 0x18
		}
		if len(x.Value) > 0 {
			i -= len(x.Value)
			copy(dAtA[i:], x.Value)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Value)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Key) > 0 {
			i -= len(x.Key)
			copy(dAtA[i:], x.Key)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Key)))
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
		x := input.Message.Interface().(*SnapshotIAVLItem)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SnapshotIAVLItem: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SnapshotIAVLItem: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
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
				x.Key = append(x.Key[:0], dAtA[iNdEx:postIndex]...)
				if x.Key == nil {
					x.Key = []byte{}
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
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
				x.Value = append(x.Value[:0], dAtA[iNdEx:postIndex]...)
				if x.Value == nil {
					x.Value = []byte{}
				}
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
				}
				x.Version = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Version |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
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
					x.Height |= int32(b&0x7F) << shift
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
// source: cosmos/base/store/v1beta1/snapshot.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// SnapshotItem is an item contained in a rootmulti.Store snapshot.
type SnapshotItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// item is the specific type of snapshot item.
	//
	// Types that are assignable to Item:
	//	*SnapshotItem_Store
	//	*SnapshotItem_Iavl
	Item isSnapshotItem_Item `protobuf_oneof:"item"`
}

func (x *SnapshotItem) Reset() {
	*x = SnapshotItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_store_v1beta1_snapshot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotItem) ProtoMessage() {}

// Deprecated: Use SnapshotItem.ProtoReflect.Descriptor instead.
func (*SnapshotItem) Descriptor() ([]byte, []int) {
	return file_cosmos_base_store_v1beta1_snapshot_proto_rawDescGZIP(), []int{0}
}

func (x *SnapshotItem) GetItem() isSnapshotItem_Item {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *SnapshotItem) GetStore() *SnapshotStoreItem {
	if x, ok := x.GetItem().(*SnapshotItem_Store); ok {
		return x.Store
	}
	return nil
}

func (x *SnapshotItem) GetIavl() *SnapshotIAVLItem {
	if x, ok := x.GetItem().(*SnapshotItem_Iavl); ok {
		return x.Iavl
	}
	return nil
}

type isSnapshotItem_Item interface {
	isSnapshotItem_Item()
}

type SnapshotItem_Store struct {
	Store *SnapshotStoreItem `protobuf:"bytes,1,opt,name=store,proto3,oneof"`
}

type SnapshotItem_Iavl struct {
	Iavl *SnapshotIAVLItem `protobuf:"bytes,2,opt,name=iavl,proto3,oneof"`
}

func (*SnapshotItem_Store) isSnapshotItem_Item() {}

func (*SnapshotItem_Iavl) isSnapshotItem_Item() {}

// SnapshotStoreItem contains metadata about a snapshotted store.
type SnapshotStoreItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SnapshotStoreItem) Reset() {
	*x = SnapshotStoreItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_store_v1beta1_snapshot_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotStoreItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotStoreItem) ProtoMessage() {}

// Deprecated: Use SnapshotStoreItem.ProtoReflect.Descriptor instead.
func (*SnapshotStoreItem) Descriptor() ([]byte, []int) {
	return file_cosmos_base_store_v1beta1_snapshot_proto_rawDescGZIP(), []int{1}
}

func (x *SnapshotStoreItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// SnapshotIAVLItem is an exported IAVL node.
type SnapshotIAVLItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key     []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value   []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Version int64  `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	Height  int32  `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *SnapshotIAVLItem) Reset() {
	*x = SnapshotIAVLItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_store_v1beta1_snapshot_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotIAVLItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotIAVLItem) ProtoMessage() {}

// Deprecated: Use SnapshotIAVLItem.ProtoReflect.Descriptor instead.
func (*SnapshotIAVLItem) Descriptor() ([]byte, []int) {
	return file_cosmos_base_store_v1beta1_snapshot_proto_rawDescGZIP(), []int{2}
}

func (x *SnapshotIAVLItem) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *SnapshotIAVLItem) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *SnapshotIAVLItem) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *SnapshotIAVLItem) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

var File_cosmos_base_store_v1beta1_snapshot_proto protoreflect.FileDescriptor

var file_cosmos_base_store_v1beta1_snapshot_proto_rawDesc = []byte{
	0x0a, 0x28, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x73, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x01, 0x0a, 0x0c,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x44, 0x0a, 0x05,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x48, 0x00, 0x52, 0x05, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x12, 0x4b, 0x0a, 0x04, 0x69, 0x61, 0x76, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x49, 0x41, 0x56, 0x4c, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x08, 0xe2,
	0xde, 0x1f, 0x04, 0x49, 0x41, 0x56, 0x4c, 0x48, 0x00, 0x52, 0x04, 0x69, 0x61, 0x76, 0x6c, 0x42,
	0x06, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x27, 0x0a, 0x11, 0x53, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x6c, 0x0a, 0x10, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x49, 0x41, 0x56, 0x4c,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x42, 0xfe,
	0x01, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x42, 0x0d, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x42, 0x53,
	0xaa, 0x02, 0x19, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x2e, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x19, 0x43,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x42, 0x61, 0x73, 0x65, 0x5c, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x25, 0x43, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x5c, 0x42, 0x61, 0x73, 0x65, 0x5c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x56, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x1c, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x42, 0x61, 0x73, 0x65, 0x3a,
	0x3a, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_base_store_v1beta1_snapshot_proto_rawDescOnce sync.Once
	file_cosmos_base_store_v1beta1_snapshot_proto_rawDescData = file_cosmos_base_store_v1beta1_snapshot_proto_rawDesc
)

func file_cosmos_base_store_v1beta1_snapshot_proto_rawDescGZIP() []byte {
	file_cosmos_base_store_v1beta1_snapshot_proto_rawDescOnce.Do(func() {
		file_cosmos_base_store_v1beta1_snapshot_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_base_store_v1beta1_snapshot_proto_rawDescData)
	})
	return file_cosmos_base_store_v1beta1_snapshot_proto_rawDescData
}

var file_cosmos_base_store_v1beta1_snapshot_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cosmos_base_store_v1beta1_snapshot_proto_goTypes = []interface{}{
	(*SnapshotItem)(nil),      // 0: cosmos.base.store.v1beta1.SnapshotItem
	(*SnapshotStoreItem)(nil), // 1: cosmos.base.store.v1beta1.SnapshotStoreItem
	(*SnapshotIAVLItem)(nil),  // 2: cosmos.base.store.v1beta1.SnapshotIAVLItem
}
var file_cosmos_base_store_v1beta1_snapshot_proto_depIdxs = []int32{
	1, // 0: cosmos.base.store.v1beta1.SnapshotItem.store:type_name -> cosmos.base.store.v1beta1.SnapshotStoreItem
	2, // 1: cosmos.base.store.v1beta1.SnapshotItem.iavl:type_name -> cosmos.base.store.v1beta1.SnapshotIAVLItem
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cosmos_base_store_v1beta1_snapshot_proto_init() }
func file_cosmos_base_store_v1beta1_snapshot_proto_init() {
	if File_cosmos_base_store_v1beta1_snapshot_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_base_store_v1beta1_snapshot_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotItem); i {
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
		file_cosmos_base_store_v1beta1_snapshot_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotStoreItem); i {
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
		file_cosmos_base_store_v1beta1_snapshot_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotIAVLItem); i {
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
	file_cosmos_base_store_v1beta1_snapshot_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SnapshotItem_Store)(nil),
		(*SnapshotItem_Iavl)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cosmos_base_store_v1beta1_snapshot_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_base_store_v1beta1_snapshot_proto_goTypes,
		DependencyIndexes: file_cosmos_base_store_v1beta1_snapshot_proto_depIdxs,
		MessageInfos:      file_cosmos_base_store_v1beta1_snapshot_proto_msgTypes,
	}.Build()
	File_cosmos_base_store_v1beta1_snapshot_proto = out.File
	file_cosmos_base_store_v1beta1_snapshot_proto_rawDesc = nil
	file_cosmos_base_store_v1beta1_snapshot_proto_goTypes = nil
	file_cosmos_base_store_v1beta1_snapshot_proto_depIdxs = nil
}
