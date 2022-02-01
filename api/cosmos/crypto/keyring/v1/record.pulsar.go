package keyringv1

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	v1 "github.com/aliworkshop/terra-sdk/api/cosmos/crypto/hd/v1"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_Record         protoreflect.MessageDescriptor
	fd_Record_name    protoreflect.FieldDescriptor
	fd_Record_pub_key protoreflect.FieldDescriptor
	fd_Record_local   protoreflect.FieldDescriptor
	fd_Record_ledger  protoreflect.FieldDescriptor
	fd_Record_multi   protoreflect.FieldDescriptor
	fd_Record_offline protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_crypto_keyring_v1_record_proto_init()
	md_Record = File_cosmos_crypto_keyring_v1_record_proto.Messages().ByName("Record")
	fd_Record_name = md_Record.Fields().ByName("name")
	fd_Record_pub_key = md_Record.Fields().ByName("pub_key")
	fd_Record_local = md_Record.Fields().ByName("local")
	fd_Record_ledger = md_Record.Fields().ByName("ledger")
	fd_Record_multi = md_Record.Fields().ByName("multi")
	fd_Record_offline = md_Record.Fields().ByName("offline")
}

var _ protoreflect.Message = (*fastReflection_Record)(nil)

type fastReflection_Record Record

func (x *Record) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Record)(x)
}

func (x *Record) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_crypto_keyring_v1_record_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Record_messageType fastReflection_Record_messageType
var _ protoreflect.MessageType = fastReflection_Record_messageType{}

type fastReflection_Record_messageType struct{}

func (x fastReflection_Record_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Record)(nil)
}
func (x fastReflection_Record_messageType) New() protoreflect.Message {
	return new(fastReflection_Record)
}
func (x fastReflection_Record_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Record
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Record) Descriptor() protoreflect.MessageDescriptor {
	return md_Record
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Record) Type() protoreflect.MessageType {
	return _fastReflection_Record_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Record) New() protoreflect.Message {
	return new(fastReflection_Record)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Record) Interface() protoreflect.ProtoMessage {
	return (*Record)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Record) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Name != "" {
		value := protoreflect.ValueOfString(x.Name)
		if !f(fd_Record_name, value) {
			return
		}
	}
	if x.PubKey != nil {
		value := protoreflect.ValueOfMessage(x.PubKey.ProtoReflect())
		if !f(fd_Record_pub_key, value) {
			return
		}
	}
	if x.Item != nil {
		switch o := x.Item.(type) {
		case *Record_Local_:
			v := o.Local
			value := protoreflect.ValueOfMessage(v.ProtoReflect())
			if !f(fd_Record_local, value) {
				return
			}
		case *Record_Ledger_:
			v := o.Ledger
			value := protoreflect.ValueOfMessage(v.ProtoReflect())
			if !f(fd_Record_ledger, value) {
				return
			}
		case *Record_Multi_:
			v := o.Multi
			value := protoreflect.ValueOfMessage(v.ProtoReflect())
			if !f(fd_Record_multi, value) {
				return
			}
		case *Record_Offline_:
			v := o.Offline
			value := protoreflect.ValueOfMessage(v.ProtoReflect())
			if !f(fd_Record_offline, value) {
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
func (x *fastReflection_Record) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.crypto.keyring.v1.Record.name":
		return x.Name != ""
	case "cosmos.crypto.keyring.v1.Record.pub_key":
		return x.PubKey != nil
	case "cosmos.crypto.keyring.v1.Record.local":
		if x.Item == nil {
			return false
		} else if _, ok := x.Item.(*Record_Local_); ok {
			return true
		} else {
			return false
		}
	case "cosmos.crypto.keyring.v1.Record.ledger":
		if x.Item == nil {
			return false
		} else if _, ok := x.Item.(*Record_Ledger_); ok {
			return true
		} else {
			return false
		}
	case "cosmos.crypto.keyring.v1.Record.multi":
		if x.Item == nil {
			return false
		} else if _, ok := x.Item.(*Record_Multi_); ok {
			return true
		} else {
			return false
		}
	case "cosmos.crypto.keyring.v1.Record.offline":
		if x.Item == nil {
			return false
		} else if _, ok := x.Item.(*Record_Offline_); ok {
			return true
		} else {
			return false
		}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.keyring.v1.Record"))
		}
		panic(fmt.Errorf("message cosmos.crypto.keyring.v1.Record does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Record) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.crypto.keyring.v1.Record.name":
		x.Name = ""
	case "cosmos.crypto.keyring.v1.Record.pub_key":
		x.PubKey = nil
	case "cosmos.crypto.keyring.v1.Record.local":
		x.Item = nil
	case "cosmos.crypto.keyring.v1.Record.ledger":
		x.Item = nil
	case "cosmos.crypto.keyring.v1.Record.multi":
		x.Item = nil
	case "cosmos.crypto.keyring.v1.Record.offline":
		x.Item = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.keyring.v1.Record"))
		}
		panic(fmt.Errorf("message cosmos.crypto.keyring.v1.Record does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Record) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.crypto.keyring.v1.Record.name":
		value := x.Name
		return protoreflect.ValueOfString(value)
	case "cosmos.crypto.keyring.v1.Record.pub_key":
		value := x.PubKey
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.crypto.keyring.v1.Record.local":
		if x.Item == nil {
			return protoreflect.ValueOfMessage((*Record_Local)(nil).ProtoReflect())
		} else if v, ok := x.Item.(*Record_Local_); ok {
			return protoreflect.ValueOfMessage(v.Local.ProtoReflect())
		} else {
			return protoreflect.ValueOfMessage((*Record_Local)(nil).ProtoReflect())
		}
	case "cosmos.crypto.keyring.v1.Record.ledger":
		if x.Item == nil {
			return protoreflect.ValueOfMessage((*Record_Ledger)(nil).ProtoReflect())
		} else if v, ok := x.Item.(*Record_Ledger_); ok {
			return protoreflect.ValueOfMessage(v.Ledger.ProtoReflect())
		} else {
			return protoreflect.ValueOfMessage((*Record_Ledger)(nil).ProtoReflect())
		}
	case "cosmos.crypto.keyring.v1.Record.multi":
		if x.Item == nil {
			return protoreflect.ValueOfMessage((*Record_Multi)(nil).ProtoReflect())
		} else if v, ok := x.Item.(*Record_Multi_); ok {
			return protoreflect.ValueOfMessage(v.Multi.ProtoReflect())
		} else {
			return protoreflect.ValueOfMessage((*Record_Multi)(nil).ProtoReflect())
		}
	case "cosmos.crypto.keyring.v1.Record.offline":
		if x.Item == nil {
			return protoreflect.ValueOfMessage((*Record_Offline)(nil).ProtoReflect())
		} else if v, ok := x.Item.(*Record_Offline_); ok {
			return protoreflect.ValueOfMessage(v.Offline.ProtoReflect())
		} else {
			return protoreflect.ValueOfMessage((*Record_Offline)(nil).ProtoReflect())
		}
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.keyring.v1.Record"))
		}
		panic(fmt.Errorf("message cosmos.crypto.keyring.v1.Record does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Record) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.crypto.keyring.v1.Record.name":
		x.Name = value.Interface().(string)
	case "cosmos.crypto.keyring.v1.Record.pub_key":
		x.PubKey = value.Message().Interface().(*anypb.Any)
	case "cosmos.crypto.keyring.v1.Record.local":
		cv := value.Message().Interface().(*Record_Local)
		x.Item = &Record_Local_{Local: cv}
	case "cosmos.crypto.keyring.v1.Record.ledger":
		cv := value.Message().Interface().(*Record_Ledger)
		x.Item = &Record_Ledger_{Ledger: cv}
	case "cosmos.crypto.keyring.v1.Record.multi":
		cv := value.Message().Interface().(*Record_Multi)
		x.Item = &Record_Multi_{Multi: cv}
	case "cosmos.crypto.keyring.v1.Record.offline":
		cv := value.Message().Interface().(*Record_Offline)
		x.Item = &Record_Offline_{Offline: cv}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.keyring.v1.Record"))
		}
		panic(fmt.Errorf("message cosmos.crypto.keyring.v1.Record does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Record) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.crypto.keyring.v1.Record.pub_key":
		if x.PubKey == nil {
			x.PubKey = new(anypb.Any)
		}
		return protoreflect.ValueOfMessage(x.PubKey.ProtoReflect())
	case "cosmos.crypto.keyring.v1.Record.local":
		if x.Item == nil {
			value := &Record_Local{}
			oneofValue := &Record_Local_{Local: value}
			x.Item = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
		switch m := x.Item.(type) {
		case *Record_Local_:
			return protoreflect.ValueOfMessage(m.Local.ProtoReflect())
		default:
			value := &Record_Local{}
			oneofValue := &Record_Local_{Local: value}
			x.Item = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
	case "cosmos.crypto.keyring.v1.Record.ledger":
		if x.Item == nil {
			value := &Record_Ledger{}
			oneofValue := &Record_Ledger_{Ledger: value}
			x.Item = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
		switch m := x.Item.(type) {
		case *Record_Ledger_:
			return protoreflect.ValueOfMessage(m.Ledger.ProtoReflect())
		default:
			value := &Record_Ledger{}
			oneofValue := &Record_Ledger_{Ledger: value}
			x.Item = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
	case "cosmos.crypto.keyring.v1.Record.multi":
		if x.Item == nil {
			value := &Record_Multi{}
			oneofValue := &Record_Multi_{Multi: value}
			x.Item = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
		switch m := x.Item.(type) {
		case *Record_Multi_:
			return protoreflect.ValueOfMessage(m.Multi.ProtoReflect())
		default:
			value := &Record_Multi{}
			oneofValue := &Record_Multi_{Multi: value}
			x.Item = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
	case "cosmos.crypto.keyring.v1.Record.offline":
		if x.Item == nil {
			value := &Record_Offline{}
			oneofValue := &Record_Offline_{Offline: value}
			x.Item = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
		switch m := x.Item.(type) {
		case *Record_Offline_:
			return protoreflect.ValueOfMessage(m.Offline.ProtoReflect())
		default:
			value := &Record_Offline{}
			oneofValue := &Record_Offline_{Offline: value}
			x.Item = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
	case "cosmos.crypto.keyring.v1.Record.name":
		panic(fmt.Errorf("field name of message cosmos.crypto.keyring.v1.Record is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.keyring.v1.Record"))
		}
		panic(fmt.Errorf("message cosmos.crypto.keyring.v1.Record does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Record) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.crypto.keyring.v1.Record.name":
		return protoreflect.ValueOfString("")
	case "cosmos.crypto.keyring.v1.Record.pub_key":
		m := new(anypb.Any)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.crypto.keyring.v1.Record.local":
		value := &Record_Local{}
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.crypto.keyring.v1.Record.ledger":
		value := &Record_Ledger{}
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.crypto.keyring.v1.Record.multi":
		value := &Record_Multi{}
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.crypto.keyring.v1.Record.offline":
		value := &Record_Offline{}
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.keyring.v1.Record"))
		}
		panic(fmt.Errorf("message cosmos.crypto.keyring.v1.Record does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Record) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	case "cosmos.crypto.keyring.v1.Record.item":
		if x.Item == nil {
			return nil
		}
		switch x.Item.(type) {
		case *Record_Local_:
			return x.Descriptor().Fields().ByName("local")
		case *Record_Ledger_:
			return x.Descriptor().Fields().ByName("ledger")
		case *Record_Multi_:
			return x.Descriptor().Fields().ByName("multi")
		case *Record_Offline_:
			return x.Descriptor().Fields().ByName("offline")
		}
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.crypto.keyring.v1.Record", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Record) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Record) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Record) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Record) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Record)
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
		if x.PubKey != nil {
			l = options.Size(x.PubKey)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		switch x := x.Item.(type) {
		case *Record_Local_:
			if x == nil {
				break
			}
			l = options.Size(x.Local)
			n += 1 + l + runtime.Sov(uint64(l))
		case *Record_Ledger_:
			if x == nil {
				break
			}
			l = options.Size(x.Ledger)
			n += 1 + l + runtime.Sov(uint64(l))
		case *Record_Multi_:
			if x == nil {
				break
			}
			l = options.Size(x.Multi)
			n += 1 + l + runtime.Sov(uint64(l))
		case *Record_Offline_:
			if x == nil {
				break
			}
			l = options.Size(x.Offline)
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
		x := input.Message.Interface().(*Record)
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
		case *Record_Local_:
			encoded, err := options.Marshal(x.Local)
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
		case *Record_Ledger_:
			encoded, err := options.Marshal(x.Ledger)
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
		case *Record_Multi_:
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
			dAtA[i] = 0x2a
		case *Record_Offline_:
			encoded, err := options.Marshal(x.Offline)
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
		x := input.Message.Interface().(*Record)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Record: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Record: illegal tag %d (wire type %d)", fieldNum, wire)
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
					x.PubKey = &anypb.Any{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.PubKey); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Local", wireType)
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
				v := &Record_Local{}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], v); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				x.Item = &Record_Local_{v}
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Ledger", wireType)
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
				v := &Record_Ledger{}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], v); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				x.Item = &Record_Ledger_{v}
				iNdEx = postIndex
			case 5:
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
				v := &Record_Multi{}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], v); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				x.Item = &Record_Multi_{v}
				iNdEx = postIndex
			case 6:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Offline", wireType)
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
				v := &Record_Offline{}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], v); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				x.Item = &Record_Offline_{v}
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
	md_Record_Local               protoreflect.MessageDescriptor
	fd_Record_Local_priv_key      protoreflect.FieldDescriptor
	fd_Record_Local_priv_key_type protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_crypto_keyring_v1_record_proto_init()
	md_Record_Local = File_cosmos_crypto_keyring_v1_record_proto.Messages().ByName("Record").Messages().ByName("Local")
	fd_Record_Local_priv_key = md_Record_Local.Fields().ByName("priv_key")
	fd_Record_Local_priv_key_type = md_Record_Local.Fields().ByName("priv_key_type")
}

var _ protoreflect.Message = (*fastReflection_Record_Local)(nil)

type fastReflection_Record_Local Record_Local

func (x *Record_Local) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Record_Local)(x)
}

func (x *Record_Local) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_crypto_keyring_v1_record_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Record_Local_messageType fastReflection_Record_Local_messageType
var _ protoreflect.MessageType = fastReflection_Record_Local_messageType{}

type fastReflection_Record_Local_messageType struct{}

func (x fastReflection_Record_Local_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Record_Local)(nil)
}
func (x fastReflection_Record_Local_messageType) New() protoreflect.Message {
	return new(fastReflection_Record_Local)
}
func (x fastReflection_Record_Local_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Record_Local
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Record_Local) Descriptor() protoreflect.MessageDescriptor {
	return md_Record_Local
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Record_Local) Type() protoreflect.MessageType {
	return _fastReflection_Record_Local_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Record_Local) New() protoreflect.Message {
	return new(fastReflection_Record_Local)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Record_Local) Interface() protoreflect.ProtoMessage {
	return (*Record_Local)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Record_Local) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.PrivKey != nil {
		value := protoreflect.ValueOfMessage(x.PrivKey.ProtoReflect())
		if !f(fd_Record_Local_priv_key, value) {
			return
		}
	}
	if x.PrivKeyType != "" {
		value := protoreflect.ValueOfString(x.PrivKeyType)
		if !f(fd_Record_Local_priv_key_type, value) {
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
func (x *fastReflection_Record_Local) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.crypto.keyring.v1.Record.Local.priv_key":
		return x.PrivKey != nil
	case "cosmos.crypto.keyring.v1.Record.Local.priv_key_type":
		return x.PrivKeyType != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.keyring.v1.Record.Local"))
		}
		panic(fmt.Errorf("message cosmos.crypto.keyring.v1.Record.Local does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Record_Local) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.crypto.keyring.v1.Record.Local.priv_key":
		x.PrivKey = nil
	case "cosmos.crypto.keyring.v1.Record.Local.priv_key_type":
		x.PrivKeyType = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.keyring.v1.Record.Local"))
		}
		panic(fmt.Errorf("message cosmos.crypto.keyring.v1.Record.Local does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Record_Local) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.crypto.keyring.v1.Record.Local.priv_key":
		value := x.PrivKey
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.crypto.keyring.v1.Record.Local.priv_key_type":
		value := x.PrivKeyType
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.keyring.v1.Record.Local"))
		}
		panic(fmt.Errorf("message cosmos.crypto.keyring.v1.Record.Local does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Record_Local) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.crypto.keyring.v1.Record.Local.priv_key":
		x.PrivKey = value.Message().Interface().(*anypb.Any)
	case "cosmos.crypto.keyring.v1.Record.Local.priv_key_type":
		x.PrivKeyType = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.keyring.v1.Record.Local"))
		}
		panic(fmt.Errorf("message cosmos.crypto.keyring.v1.Record.Local does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Record_Local) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.crypto.keyring.v1.Record.Local.priv_key":
		if x.PrivKey == nil {
			x.PrivKey = new(anypb.Any)
		}
		return protoreflect.ValueOfMessage(x.PrivKey.ProtoReflect())
	case "cosmos.crypto.keyring.v1.Record.Local.priv_key_type":
		panic(fmt.Errorf("field priv_key_type of message cosmos.crypto.keyring.v1.Record.Local is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.keyring.v1.Record.Local"))
		}
		panic(fmt.Errorf("message cosmos.crypto.keyring.v1.Record.Local does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Record_Local) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.crypto.keyring.v1.Record.Local.priv_key":
		m := new(anypb.Any)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.crypto.keyring.v1.Record.Local.priv_key_type":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.keyring.v1.Record.Local"))
		}
		panic(fmt.Errorf("message cosmos.crypto.keyring.v1.Record.Local does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Record_Local) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.crypto.keyring.v1.Record.Local", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Record_Local) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Record_Local) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Record_Local) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Record_Local) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Record_Local)
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
		if x.PrivKey != nil {
			l = options.Size(x.PrivKey)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.PrivKeyType)
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
		x := input.Message.Interface().(*Record_Local)
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
		if len(x.PrivKeyType) > 0 {
			i -= len(x.PrivKeyType)
			copy(dAtA[i:], x.PrivKeyType)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.PrivKeyType)))
			i--
			dAtA[i] = 0x12
		}
		if x.PrivKey != nil {
			encoded, err := options.Marshal(x.PrivKey)
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
		x := input.Message.Interface().(*Record_Local)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Record_Local: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Record_Local: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field PrivKey", wireType)
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
				if x.PrivKey == nil {
					x.PrivKey = &anypb.Any{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.PrivKey); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field PrivKeyType", wireType)
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
				x.PrivKeyType = string(dAtA[iNdEx:postIndex])
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
	md_Record_Ledger      protoreflect.MessageDescriptor
	fd_Record_Ledger_path protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_crypto_keyring_v1_record_proto_init()
	md_Record_Ledger = File_cosmos_crypto_keyring_v1_record_proto.Messages().ByName("Record").Messages().ByName("Ledger")
	fd_Record_Ledger_path = md_Record_Ledger.Fields().ByName("path")
}

var _ protoreflect.Message = (*fastReflection_Record_Ledger)(nil)

type fastReflection_Record_Ledger Record_Ledger

func (x *Record_Ledger) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Record_Ledger)(x)
}

func (x *Record_Ledger) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_crypto_keyring_v1_record_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Record_Ledger_messageType fastReflection_Record_Ledger_messageType
var _ protoreflect.MessageType = fastReflection_Record_Ledger_messageType{}

type fastReflection_Record_Ledger_messageType struct{}

func (x fastReflection_Record_Ledger_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Record_Ledger)(nil)
}
func (x fastReflection_Record_Ledger_messageType) New() protoreflect.Message {
	return new(fastReflection_Record_Ledger)
}
func (x fastReflection_Record_Ledger_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Record_Ledger
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Record_Ledger) Descriptor() protoreflect.MessageDescriptor {
	return md_Record_Ledger
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Record_Ledger) Type() protoreflect.MessageType {
	return _fastReflection_Record_Ledger_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Record_Ledger) New() protoreflect.Message {
	return new(fastReflection_Record_Ledger)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Record_Ledger) Interface() protoreflect.ProtoMessage {
	return (*Record_Ledger)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Record_Ledger) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Path != nil {
		value := protoreflect.ValueOfMessage(x.Path.ProtoReflect())
		if !f(fd_Record_Ledger_path, value) {
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
func (x *fastReflection_Record_Ledger) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.crypto.keyring.v1.Record.Ledger.path":
		return x.Path != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.keyring.v1.Record.Ledger"))
		}
		panic(fmt.Errorf("message cosmos.crypto.keyring.v1.Record.Ledger does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Record_Ledger) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.crypto.keyring.v1.Record.Ledger.path":
		x.Path = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.keyring.v1.Record.Ledger"))
		}
		panic(fmt.Errorf("message cosmos.crypto.keyring.v1.Record.Ledger does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Record_Ledger) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.crypto.keyring.v1.Record.Ledger.path":
		value := x.Path
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.keyring.v1.Record.Ledger"))
		}
		panic(fmt.Errorf("message cosmos.crypto.keyring.v1.Record.Ledger does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Record_Ledger) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.crypto.keyring.v1.Record.Ledger.path":
		x.Path = value.Message().Interface().(*v1.BIP44Params)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.keyring.v1.Record.Ledger"))
		}
		panic(fmt.Errorf("message cosmos.crypto.keyring.v1.Record.Ledger does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Record_Ledger) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.crypto.keyring.v1.Record.Ledger.path":
		if x.Path == nil {
			x.Path = new(v1.BIP44Params)
		}
		return protoreflect.ValueOfMessage(x.Path.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.keyring.v1.Record.Ledger"))
		}
		panic(fmt.Errorf("message cosmos.crypto.keyring.v1.Record.Ledger does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Record_Ledger) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.crypto.keyring.v1.Record.Ledger.path":
		m := new(v1.BIP44Params)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.keyring.v1.Record.Ledger"))
		}
		panic(fmt.Errorf("message cosmos.crypto.keyring.v1.Record.Ledger does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Record_Ledger) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.crypto.keyring.v1.Record.Ledger", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Record_Ledger) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Record_Ledger) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Record_Ledger) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Record_Ledger) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Record_Ledger)
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
		if x.Path != nil {
			l = options.Size(x.Path)
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
		x := input.Message.Interface().(*Record_Ledger)
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
		if x.Path != nil {
			encoded, err := options.Marshal(x.Path)
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
		x := input.Message.Interface().(*Record_Ledger)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Record_Ledger: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Record_Ledger: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
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
				if x.Path == nil {
					x.Path = &v1.BIP44Params{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Path); err != nil {
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
	md_Record_Multi protoreflect.MessageDescriptor
)

func init() {
	file_cosmos_crypto_keyring_v1_record_proto_init()
	md_Record_Multi = File_cosmos_crypto_keyring_v1_record_proto.Messages().ByName("Record").Messages().ByName("Multi")
}

var _ protoreflect.Message = (*fastReflection_Record_Multi)(nil)

type fastReflection_Record_Multi Record_Multi

func (x *Record_Multi) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Record_Multi)(x)
}

func (x *Record_Multi) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_crypto_keyring_v1_record_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Record_Multi_messageType fastReflection_Record_Multi_messageType
var _ protoreflect.MessageType = fastReflection_Record_Multi_messageType{}

type fastReflection_Record_Multi_messageType struct{}

func (x fastReflection_Record_Multi_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Record_Multi)(nil)
}
func (x fastReflection_Record_Multi_messageType) New() protoreflect.Message {
	return new(fastReflection_Record_Multi)
}
func (x fastReflection_Record_Multi_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Record_Multi
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Record_Multi) Descriptor() protoreflect.MessageDescriptor {
	return md_Record_Multi
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Record_Multi) Type() protoreflect.MessageType {
	return _fastReflection_Record_Multi_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Record_Multi) New() protoreflect.Message {
	return new(fastReflection_Record_Multi)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Record_Multi) Interface() protoreflect.ProtoMessage {
	return (*Record_Multi)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Record_Multi) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_Record_Multi) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.keyring.v1.Record.Multi"))
		}
		panic(fmt.Errorf("message cosmos.crypto.keyring.v1.Record.Multi does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Record_Multi) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.keyring.v1.Record.Multi"))
		}
		panic(fmt.Errorf("message cosmos.crypto.keyring.v1.Record.Multi does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Record_Multi) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.keyring.v1.Record.Multi"))
		}
		panic(fmt.Errorf("message cosmos.crypto.keyring.v1.Record.Multi does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Record_Multi) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.keyring.v1.Record.Multi"))
		}
		panic(fmt.Errorf("message cosmos.crypto.keyring.v1.Record.Multi does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Record_Multi) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.keyring.v1.Record.Multi"))
		}
		panic(fmt.Errorf("message cosmos.crypto.keyring.v1.Record.Multi does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Record_Multi) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.keyring.v1.Record.Multi"))
		}
		panic(fmt.Errorf("message cosmos.crypto.keyring.v1.Record.Multi does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Record_Multi) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.crypto.keyring.v1.Record.Multi", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Record_Multi) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Record_Multi) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Record_Multi) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Record_Multi) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Record_Multi)
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
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Record_Multi)
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
		x := input.Message.Interface().(*Record_Multi)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Record_Multi: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Record_Multi: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
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
	md_Record_Offline protoreflect.MessageDescriptor
)

func init() {
	file_cosmos_crypto_keyring_v1_record_proto_init()
	md_Record_Offline = File_cosmos_crypto_keyring_v1_record_proto.Messages().ByName("Record").Messages().ByName("Offline")
}

var _ protoreflect.Message = (*fastReflection_Record_Offline)(nil)

type fastReflection_Record_Offline Record_Offline

func (x *Record_Offline) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Record_Offline)(x)
}

func (x *Record_Offline) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_crypto_keyring_v1_record_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Record_Offline_messageType fastReflection_Record_Offline_messageType
var _ protoreflect.MessageType = fastReflection_Record_Offline_messageType{}

type fastReflection_Record_Offline_messageType struct{}

func (x fastReflection_Record_Offline_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Record_Offline)(nil)
}
func (x fastReflection_Record_Offline_messageType) New() protoreflect.Message {
	return new(fastReflection_Record_Offline)
}
func (x fastReflection_Record_Offline_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Record_Offline
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Record_Offline) Descriptor() protoreflect.MessageDescriptor {
	return md_Record_Offline
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Record_Offline) Type() protoreflect.MessageType {
	return _fastReflection_Record_Offline_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Record_Offline) New() protoreflect.Message {
	return new(fastReflection_Record_Offline)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Record_Offline) Interface() protoreflect.ProtoMessage {
	return (*Record_Offline)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Record_Offline) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_Record_Offline) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.keyring.v1.Record.Offline"))
		}
		panic(fmt.Errorf("message cosmos.crypto.keyring.v1.Record.Offline does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Record_Offline) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.keyring.v1.Record.Offline"))
		}
		panic(fmt.Errorf("message cosmos.crypto.keyring.v1.Record.Offline does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Record_Offline) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.keyring.v1.Record.Offline"))
		}
		panic(fmt.Errorf("message cosmos.crypto.keyring.v1.Record.Offline does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Record_Offline) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.keyring.v1.Record.Offline"))
		}
		panic(fmt.Errorf("message cosmos.crypto.keyring.v1.Record.Offline does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Record_Offline) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.keyring.v1.Record.Offline"))
		}
		panic(fmt.Errorf("message cosmos.crypto.keyring.v1.Record.Offline does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Record_Offline) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.crypto.keyring.v1.Record.Offline"))
		}
		panic(fmt.Errorf("message cosmos.crypto.keyring.v1.Record.Offline does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Record_Offline) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.crypto.keyring.v1.Record.Offline", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Record_Offline) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Record_Offline) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Record_Offline) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Record_Offline) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Record_Offline)
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
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Record_Offline)
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
		x := input.Message.Interface().(*Record_Offline)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Record_Offline: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Record_Offline: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
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
// source: cosmos/crypto/keyring/v1/record.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Record is used for representing a key in the keyring.
type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name represents a name of Record
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// pub_key represents a public key in any format
	PubKey *anypb.Any `protobuf:"bytes,2,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	// Record contains one of the following items
	//
	// Types that are assignable to Item:
	//	*Record_Local_
	//	*Record_Ledger_
	//	*Record_Multi_
	//	*Record_Offline_
	Item isRecord_Item `protobuf_oneof:"item"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_crypto_keyring_v1_record_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_cosmos_crypto_keyring_v1_record_proto_rawDescGZIP(), []int{0}
}

func (x *Record) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Record) GetPubKey() *anypb.Any {
	if x != nil {
		return x.PubKey
	}
	return nil
}

func (x *Record) GetItem() isRecord_Item {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *Record) GetLocal() *Record_Local {
	if x, ok := x.GetItem().(*Record_Local_); ok {
		return x.Local
	}
	return nil
}

func (x *Record) GetLedger() *Record_Ledger {
	if x, ok := x.GetItem().(*Record_Ledger_); ok {
		return x.Ledger
	}
	return nil
}

func (x *Record) GetMulti() *Record_Multi {
	if x, ok := x.GetItem().(*Record_Multi_); ok {
		return x.Multi
	}
	return nil
}

func (x *Record) GetOffline() *Record_Offline {
	if x, ok := x.GetItem().(*Record_Offline_); ok {
		return x.Offline
	}
	return nil
}

type isRecord_Item interface {
	isRecord_Item()
}

type Record_Local_ struct {
	// local stores the public information about a locally stored key
	Local *Record_Local `protobuf:"bytes,3,opt,name=local,proto3,oneof"`
}

type Record_Ledger_ struct {
	// ledger stores the public information about a Ledger key
	Ledger *Record_Ledger `protobuf:"bytes,4,opt,name=ledger,proto3,oneof"`
}

type Record_Multi_ struct {
	// Multi does not store any information.
	Multi *Record_Multi `protobuf:"bytes,5,opt,name=multi,proto3,oneof"`
}

type Record_Offline_ struct {
	// Offline does not store any information.
	Offline *Record_Offline `protobuf:"bytes,6,opt,name=offline,proto3,oneof"`
}

func (*Record_Local_) isRecord_Item() {}

func (*Record_Ledger_) isRecord_Item() {}

func (*Record_Multi_) isRecord_Item() {}

func (*Record_Offline_) isRecord_Item() {}

// Item is a keyring item stored in a keyring backend.
// Local item
type Record_Local struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrivKey     *anypb.Any `protobuf:"bytes,1,opt,name=priv_key,json=privKey,proto3" json:"priv_key,omitempty"`
	PrivKeyType string     `protobuf:"bytes,2,opt,name=priv_key_type,json=privKeyType,proto3" json:"priv_key_type,omitempty"`
}

func (x *Record_Local) Reset() {
	*x = Record_Local{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_crypto_keyring_v1_record_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record_Local) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record_Local) ProtoMessage() {}

// Deprecated: Use Record_Local.ProtoReflect.Descriptor instead.
func (*Record_Local) Descriptor() ([]byte, []int) {
	return file_cosmos_crypto_keyring_v1_record_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Record_Local) GetPrivKey() *anypb.Any {
	if x != nil {
		return x.PrivKey
	}
	return nil
}

func (x *Record_Local) GetPrivKeyType() string {
	if x != nil {
		return x.PrivKeyType
	}
	return ""
}

// Ledger item
type Record_Ledger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path *v1.BIP44Params `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *Record_Ledger) Reset() {
	*x = Record_Ledger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_crypto_keyring_v1_record_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record_Ledger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record_Ledger) ProtoMessage() {}

// Deprecated: Use Record_Ledger.ProtoReflect.Descriptor instead.
func (*Record_Ledger) Descriptor() ([]byte, []int) {
	return file_cosmos_crypto_keyring_v1_record_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Record_Ledger) GetPath() *v1.BIP44Params {
	if x != nil {
		return x.Path
	}
	return nil
}

// Multi item
type Record_Multi struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Record_Multi) Reset() {
	*x = Record_Multi{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_crypto_keyring_v1_record_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record_Multi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record_Multi) ProtoMessage() {}

// Deprecated: Use Record_Multi.ProtoReflect.Descriptor instead.
func (*Record_Multi) Descriptor() ([]byte, []int) {
	return file_cosmos_crypto_keyring_v1_record_proto_rawDescGZIP(), []int{0, 2}
}

// Offline item
type Record_Offline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Record_Offline) Reset() {
	*x = Record_Offline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_crypto_keyring_v1_record_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record_Offline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record_Offline) ProtoMessage() {}

// Deprecated: Use Record_Offline.ProtoReflect.Descriptor instead.
func (*Record_Offline) Descriptor() ([]byte, []int) {
	return file_cosmos_crypto_keyring_v1_record_proto_rawDescGZIP(), []int{0, 3}
}

var File_cosmos_crypto_keyring_v1_record_proto protoreflect.FileDescriptor

var file_cosmos_crypto_keyring_v1_record_proto_rawDesc = []byte{
	0x0a, 0x25, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f,
	0x6b, 0x65, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x6b, 0x65, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x2f, 0x68, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8e, 0x04, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x2d, 0x0a, 0x07, 0x70, 0x75, 0x62, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x3e,
	0x0a, 0x05, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x6b, 0x65,
	0x79, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e,
	0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x41,
	0x0a, 0x06, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x6b,
	0x65, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x2e, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x12, 0x3e, 0x0a, 0x05, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x2e, 0x6b, 0x65, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x48, 0x00, 0x52, 0x05, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x12, 0x44, 0x0a, 0x07, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x2e, 0x6b, 0x65, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x2e, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x48, 0x00, 0x52, 0x07,
	0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x1a, 0x5c, 0x0a, 0x05, 0x4c, 0x6f, 0x63, 0x61, 0x6c,
	0x12, 0x2f, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x76, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x70, 0x72, 0x69, 0x76, 0x4b, 0x65,
	0x79, 0x12, 0x22, 0x0a, 0x0d, 0x70, 0x72, 0x69, 0x76, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x4b, 0x65,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x3e, 0x0a, 0x06, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x12,
	0x34, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x68, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x49, 0x50, 0x34, 0x34, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x1a, 0x07, 0x0a, 0x05, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x1a, 0x09,
	0x0a, 0x07, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x69, 0x74, 0x65,
	0x6d, 0x42, 0xf7, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x6b, 0x65, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x42, 0x0b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x2f, 0x6b, 0x65, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x6b, 0x65, 0x79,
	0x72, 0x69, 0x6e, 0x67, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x43, 0x4b, 0xaa, 0x02, 0x18, 0x43,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79,
	0x72, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x18, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x5c, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5c, 0x4b, 0x65, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x24, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x43, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x5c, 0x4b, 0x65, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1b, 0x43, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x3a, 0x3a, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x3a, 0x3a, 0x4b, 0x65, 0x79, 0x72,
	0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0xc8, 0xe1, 0x1e, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_cosmos_crypto_keyring_v1_record_proto_rawDescOnce sync.Once
	file_cosmos_crypto_keyring_v1_record_proto_rawDescData = file_cosmos_crypto_keyring_v1_record_proto_rawDesc
)

func file_cosmos_crypto_keyring_v1_record_proto_rawDescGZIP() []byte {
	file_cosmos_crypto_keyring_v1_record_proto_rawDescOnce.Do(func() {
		file_cosmos_crypto_keyring_v1_record_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_crypto_keyring_v1_record_proto_rawDescData)
	})
	return file_cosmos_crypto_keyring_v1_record_proto_rawDescData
}

var file_cosmos_crypto_keyring_v1_record_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cosmos_crypto_keyring_v1_record_proto_goTypes = []interface{}{
	(*Record)(nil),         // 0: cosmos.crypto.keyring.v1.Record
	(*Record_Local)(nil),   // 1: cosmos.crypto.keyring.v1.Record.Local
	(*Record_Ledger)(nil),  // 2: cosmos.crypto.keyring.v1.Record.Ledger
	(*Record_Multi)(nil),   // 3: cosmos.crypto.keyring.v1.Record.Multi
	(*Record_Offline)(nil), // 4: cosmos.crypto.keyring.v1.Record.Offline
	(*anypb.Any)(nil),      // 5: google.protobuf.Any
	(*v1.BIP44Params)(nil), // 6: cosmos.crypto.hd.v1.BIP44Params
}
var file_cosmos_crypto_keyring_v1_record_proto_depIdxs = []int32{
	5, // 0: cosmos.crypto.keyring.v1.Record.pub_key:type_name -> google.protobuf.Any
	1, // 1: cosmos.crypto.keyring.v1.Record.local:type_name -> cosmos.crypto.keyring.v1.Record.Local
	2, // 2: cosmos.crypto.keyring.v1.Record.ledger:type_name -> cosmos.crypto.keyring.v1.Record.Ledger
	3, // 3: cosmos.crypto.keyring.v1.Record.multi:type_name -> cosmos.crypto.keyring.v1.Record.Multi
	4, // 4: cosmos.crypto.keyring.v1.Record.offline:type_name -> cosmos.crypto.keyring.v1.Record.Offline
	5, // 5: cosmos.crypto.keyring.v1.Record.Local.priv_key:type_name -> google.protobuf.Any
	6, // 6: cosmos.crypto.keyring.v1.Record.Ledger.path:type_name -> cosmos.crypto.hd.v1.BIP44Params
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_cosmos_crypto_keyring_v1_record_proto_init() }
func file_cosmos_crypto_keyring_v1_record_proto_init() {
	if File_cosmos_crypto_keyring_v1_record_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_crypto_keyring_v1_record_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record); i {
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
		file_cosmos_crypto_keyring_v1_record_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record_Local); i {
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
		file_cosmos_crypto_keyring_v1_record_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record_Ledger); i {
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
		file_cosmos_crypto_keyring_v1_record_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record_Multi); i {
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
		file_cosmos_crypto_keyring_v1_record_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record_Offline); i {
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
	file_cosmos_crypto_keyring_v1_record_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Record_Local_)(nil),
		(*Record_Ledger_)(nil),
		(*Record_Multi_)(nil),
		(*Record_Offline_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cosmos_crypto_keyring_v1_record_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_crypto_keyring_v1_record_proto_goTypes,
		DependencyIndexes: file_cosmos_crypto_keyring_v1_record_proto_depIdxs,
		MessageInfos:      file_cosmos_crypto_keyring_v1_record_proto_msgTypes,
	}.Build()
	File_cosmos_crypto_keyring_v1_record_proto = out.File
	file_cosmos_crypto_keyring_v1_record_proto_rawDesc = nil
	file_cosmos_crypto_keyring_v1_record_proto_goTypes = nil
	file_cosmos_crypto_keyring_v1_record_proto_depIdxs = nil
}
