package abciv1beta1

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	abci "github.com/aliworkshop/terra-sdk/api/tendermint/abci"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var _ protoreflect.List = (*_TxResponse_7_list)(nil)

type _TxResponse_7_list struct {
	list *[]*ABCIMessageLog
}

func (x *_TxResponse_7_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_TxResponse_7_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_TxResponse_7_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*ABCIMessageLog)
	(*x.list)[i] = concreteValue
}

func (x *_TxResponse_7_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*ABCIMessageLog)
	*x.list = append(*x.list, concreteValue)
}

func (x *_TxResponse_7_list) AppendMutable() protoreflect.Value {
	v := new(ABCIMessageLog)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_TxResponse_7_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_TxResponse_7_list) NewElement() protoreflect.Value {
	v := new(ABCIMessageLog)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_TxResponse_7_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_TxResponse_13_list)(nil)

type _TxResponse_13_list struct {
	list *[]*abci.Event
}

func (x *_TxResponse_13_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_TxResponse_13_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_TxResponse_13_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*abci.Event)
	(*x.list)[i] = concreteValue
}

func (x *_TxResponse_13_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*abci.Event)
	*x.list = append(*x.list, concreteValue)
}

func (x *_TxResponse_13_list) AppendMutable() protoreflect.Value {
	v := new(abci.Event)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_TxResponse_13_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_TxResponse_13_list) NewElement() protoreflect.Value {
	v := new(abci.Event)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_TxResponse_13_list) IsValid() bool {
	return x.list != nil
}

var (
	md_TxResponse            protoreflect.MessageDescriptor
	fd_TxResponse_height     protoreflect.FieldDescriptor
	fd_TxResponse_txhash     protoreflect.FieldDescriptor
	fd_TxResponse_codespace  protoreflect.FieldDescriptor
	fd_TxResponse_code       protoreflect.FieldDescriptor
	fd_TxResponse_data       protoreflect.FieldDescriptor
	fd_TxResponse_raw_log    protoreflect.FieldDescriptor
	fd_TxResponse_logs       protoreflect.FieldDescriptor
	fd_TxResponse_info       protoreflect.FieldDescriptor
	fd_TxResponse_gas_wanted protoreflect.FieldDescriptor
	fd_TxResponse_gas_used   protoreflect.FieldDescriptor
	fd_TxResponse_tx         protoreflect.FieldDescriptor
	fd_TxResponse_timestamp  protoreflect.FieldDescriptor
	fd_TxResponse_events     protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_abci_v1beta1_abci_proto_init()
	md_TxResponse = File_cosmos_base_abci_v1beta1_abci_proto.Messages().ByName("TxResponse")
	fd_TxResponse_height = md_TxResponse.Fields().ByName("height")
	fd_TxResponse_txhash = md_TxResponse.Fields().ByName("txhash")
	fd_TxResponse_codespace = md_TxResponse.Fields().ByName("codespace")
	fd_TxResponse_code = md_TxResponse.Fields().ByName("code")
	fd_TxResponse_data = md_TxResponse.Fields().ByName("data")
	fd_TxResponse_raw_log = md_TxResponse.Fields().ByName("raw_log")
	fd_TxResponse_logs = md_TxResponse.Fields().ByName("logs")
	fd_TxResponse_info = md_TxResponse.Fields().ByName("info")
	fd_TxResponse_gas_wanted = md_TxResponse.Fields().ByName("gas_wanted")
	fd_TxResponse_gas_used = md_TxResponse.Fields().ByName("gas_used")
	fd_TxResponse_tx = md_TxResponse.Fields().ByName("tx")
	fd_TxResponse_timestamp = md_TxResponse.Fields().ByName("timestamp")
	fd_TxResponse_events = md_TxResponse.Fields().ByName("events")
}

var _ protoreflect.Message = (*fastReflection_TxResponse)(nil)

type fastReflection_TxResponse TxResponse

func (x *TxResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_TxResponse)(x)
}

func (x *TxResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_abci_v1beta1_abci_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_TxResponse_messageType fastReflection_TxResponse_messageType
var _ protoreflect.MessageType = fastReflection_TxResponse_messageType{}

type fastReflection_TxResponse_messageType struct{}

func (x fastReflection_TxResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_TxResponse)(nil)
}
func (x fastReflection_TxResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_TxResponse)
}
func (x fastReflection_TxResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_TxResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_TxResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_TxResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_TxResponse) Type() protoreflect.MessageType {
	return _fastReflection_TxResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_TxResponse) New() protoreflect.Message {
	return new(fastReflection_TxResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_TxResponse) Interface() protoreflect.ProtoMessage {
	return (*TxResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_TxResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Height != int64(0) {
		value := protoreflect.ValueOfInt64(x.Height)
		if !f(fd_TxResponse_height, value) {
			return
		}
	}
	if x.Txhash != "" {
		value := protoreflect.ValueOfString(x.Txhash)
		if !f(fd_TxResponse_txhash, value) {
			return
		}
	}
	if x.Codespace != "" {
		value := protoreflect.ValueOfString(x.Codespace)
		if !f(fd_TxResponse_codespace, value) {
			return
		}
	}
	if x.Code != uint32(0) {
		value := protoreflect.ValueOfUint32(x.Code)
		if !f(fd_TxResponse_code, value) {
			return
		}
	}
	if x.Data != "" {
		value := protoreflect.ValueOfString(x.Data)
		if !f(fd_TxResponse_data, value) {
			return
		}
	}
	if x.RawLog != "" {
		value := protoreflect.ValueOfString(x.RawLog)
		if !f(fd_TxResponse_raw_log, value) {
			return
		}
	}
	if len(x.Logs) != 0 {
		value := protoreflect.ValueOfList(&_TxResponse_7_list{list: &x.Logs})
		if !f(fd_TxResponse_logs, value) {
			return
		}
	}
	if x.Info != "" {
		value := protoreflect.ValueOfString(x.Info)
		if !f(fd_TxResponse_info, value) {
			return
		}
	}
	if x.GasWanted != int64(0) {
		value := protoreflect.ValueOfInt64(x.GasWanted)
		if !f(fd_TxResponse_gas_wanted, value) {
			return
		}
	}
	if x.GasUsed != int64(0) {
		value := protoreflect.ValueOfInt64(x.GasUsed)
		if !f(fd_TxResponse_gas_used, value) {
			return
		}
	}
	if x.Tx != nil {
		value := protoreflect.ValueOfMessage(x.Tx.ProtoReflect())
		if !f(fd_TxResponse_tx, value) {
			return
		}
	}
	if x.Timestamp != "" {
		value := protoreflect.ValueOfString(x.Timestamp)
		if !f(fd_TxResponse_timestamp, value) {
			return
		}
	}
	if len(x.Events) != 0 {
		value := protoreflect.ValueOfList(&_TxResponse_13_list{list: &x.Events})
		if !f(fd_TxResponse_events, value) {
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
func (x *fastReflection_TxResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.TxResponse.height":
		return x.Height != int64(0)
	case "cosmos.base.abci.v1beta1.TxResponse.txhash":
		return x.Txhash != ""
	case "cosmos.base.abci.v1beta1.TxResponse.codespace":
		return x.Codespace != ""
	case "cosmos.base.abci.v1beta1.TxResponse.code":
		return x.Code != uint32(0)
	case "cosmos.base.abci.v1beta1.TxResponse.data":
		return x.Data != ""
	case "cosmos.base.abci.v1beta1.TxResponse.raw_log":
		return x.RawLog != ""
	case "cosmos.base.abci.v1beta1.TxResponse.logs":
		return len(x.Logs) != 0
	case "cosmos.base.abci.v1beta1.TxResponse.info":
		return x.Info != ""
	case "cosmos.base.abci.v1beta1.TxResponse.gas_wanted":
		return x.GasWanted != int64(0)
	case "cosmos.base.abci.v1beta1.TxResponse.gas_used":
		return x.GasUsed != int64(0)
	case "cosmos.base.abci.v1beta1.TxResponse.tx":
		return x.Tx != nil
	case "cosmos.base.abci.v1beta1.TxResponse.timestamp":
		return x.Timestamp != ""
	case "cosmos.base.abci.v1beta1.TxResponse.events":
		return len(x.Events) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.TxResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.TxResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_TxResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.TxResponse.height":
		x.Height = int64(0)
	case "cosmos.base.abci.v1beta1.TxResponse.txhash":
		x.Txhash = ""
	case "cosmos.base.abci.v1beta1.TxResponse.codespace":
		x.Codespace = ""
	case "cosmos.base.abci.v1beta1.TxResponse.code":
		x.Code = uint32(0)
	case "cosmos.base.abci.v1beta1.TxResponse.data":
		x.Data = ""
	case "cosmos.base.abci.v1beta1.TxResponse.raw_log":
		x.RawLog = ""
	case "cosmos.base.abci.v1beta1.TxResponse.logs":
		x.Logs = nil
	case "cosmos.base.abci.v1beta1.TxResponse.info":
		x.Info = ""
	case "cosmos.base.abci.v1beta1.TxResponse.gas_wanted":
		x.GasWanted = int64(0)
	case "cosmos.base.abci.v1beta1.TxResponse.gas_used":
		x.GasUsed = int64(0)
	case "cosmos.base.abci.v1beta1.TxResponse.tx":
		x.Tx = nil
	case "cosmos.base.abci.v1beta1.TxResponse.timestamp":
		x.Timestamp = ""
	case "cosmos.base.abci.v1beta1.TxResponse.events":
		x.Events = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.TxResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.TxResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_TxResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.abci.v1beta1.TxResponse.height":
		value := x.Height
		return protoreflect.ValueOfInt64(value)
	case "cosmos.base.abci.v1beta1.TxResponse.txhash":
		value := x.Txhash
		return protoreflect.ValueOfString(value)
	case "cosmos.base.abci.v1beta1.TxResponse.codespace":
		value := x.Codespace
		return protoreflect.ValueOfString(value)
	case "cosmos.base.abci.v1beta1.TxResponse.code":
		value := x.Code
		return protoreflect.ValueOfUint32(value)
	case "cosmos.base.abci.v1beta1.TxResponse.data":
		value := x.Data
		return protoreflect.ValueOfString(value)
	case "cosmos.base.abci.v1beta1.TxResponse.raw_log":
		value := x.RawLog
		return protoreflect.ValueOfString(value)
	case "cosmos.base.abci.v1beta1.TxResponse.logs":
		if len(x.Logs) == 0 {
			return protoreflect.ValueOfList(&_TxResponse_7_list{})
		}
		listValue := &_TxResponse_7_list{list: &x.Logs}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.base.abci.v1beta1.TxResponse.info":
		value := x.Info
		return protoreflect.ValueOfString(value)
	case "cosmos.base.abci.v1beta1.TxResponse.gas_wanted":
		value := x.GasWanted
		return protoreflect.ValueOfInt64(value)
	case "cosmos.base.abci.v1beta1.TxResponse.gas_used":
		value := x.GasUsed
		return protoreflect.ValueOfInt64(value)
	case "cosmos.base.abci.v1beta1.TxResponse.tx":
		value := x.Tx
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.base.abci.v1beta1.TxResponse.timestamp":
		value := x.Timestamp
		return protoreflect.ValueOfString(value)
	case "cosmos.base.abci.v1beta1.TxResponse.events":
		if len(x.Events) == 0 {
			return protoreflect.ValueOfList(&_TxResponse_13_list{})
		}
		listValue := &_TxResponse_13_list{list: &x.Events}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.TxResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.TxResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_TxResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.TxResponse.height":
		x.Height = value.Int()
	case "cosmos.base.abci.v1beta1.TxResponse.txhash":
		x.Txhash = value.Interface().(string)
	case "cosmos.base.abci.v1beta1.TxResponse.codespace":
		x.Codespace = value.Interface().(string)
	case "cosmos.base.abci.v1beta1.TxResponse.code":
		x.Code = uint32(value.Uint())
	case "cosmos.base.abci.v1beta1.TxResponse.data":
		x.Data = value.Interface().(string)
	case "cosmos.base.abci.v1beta1.TxResponse.raw_log":
		x.RawLog = value.Interface().(string)
	case "cosmos.base.abci.v1beta1.TxResponse.logs":
		lv := value.List()
		clv := lv.(*_TxResponse_7_list)
		x.Logs = *clv.list
	case "cosmos.base.abci.v1beta1.TxResponse.info":
		x.Info = value.Interface().(string)
	case "cosmos.base.abci.v1beta1.TxResponse.gas_wanted":
		x.GasWanted = value.Int()
	case "cosmos.base.abci.v1beta1.TxResponse.gas_used":
		x.GasUsed = value.Int()
	case "cosmos.base.abci.v1beta1.TxResponse.tx":
		x.Tx = value.Message().Interface().(*anypb.Any)
	case "cosmos.base.abci.v1beta1.TxResponse.timestamp":
		x.Timestamp = value.Interface().(string)
	case "cosmos.base.abci.v1beta1.TxResponse.events":
		lv := value.List()
		clv := lv.(*_TxResponse_13_list)
		x.Events = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.TxResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.TxResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_TxResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.TxResponse.logs":
		if x.Logs == nil {
			x.Logs = []*ABCIMessageLog{}
		}
		value := &_TxResponse_7_list{list: &x.Logs}
		return protoreflect.ValueOfList(value)
	case "cosmos.base.abci.v1beta1.TxResponse.tx":
		if x.Tx == nil {
			x.Tx = new(anypb.Any)
		}
		return protoreflect.ValueOfMessage(x.Tx.ProtoReflect())
	case "cosmos.base.abci.v1beta1.TxResponse.events":
		if x.Events == nil {
			x.Events = []*abci.Event{}
		}
		value := &_TxResponse_13_list{list: &x.Events}
		return protoreflect.ValueOfList(value)
	case "cosmos.base.abci.v1beta1.TxResponse.height":
		panic(fmt.Errorf("field height of message cosmos.base.abci.v1beta1.TxResponse is not mutable"))
	case "cosmos.base.abci.v1beta1.TxResponse.txhash":
		panic(fmt.Errorf("field txhash of message cosmos.base.abci.v1beta1.TxResponse is not mutable"))
	case "cosmos.base.abci.v1beta1.TxResponse.codespace":
		panic(fmt.Errorf("field codespace of message cosmos.base.abci.v1beta1.TxResponse is not mutable"))
	case "cosmos.base.abci.v1beta1.TxResponse.code":
		panic(fmt.Errorf("field code of message cosmos.base.abci.v1beta1.TxResponse is not mutable"))
	case "cosmos.base.abci.v1beta1.TxResponse.data":
		panic(fmt.Errorf("field data of message cosmos.base.abci.v1beta1.TxResponse is not mutable"))
	case "cosmos.base.abci.v1beta1.TxResponse.raw_log":
		panic(fmt.Errorf("field raw_log of message cosmos.base.abci.v1beta1.TxResponse is not mutable"))
	case "cosmos.base.abci.v1beta1.TxResponse.info":
		panic(fmt.Errorf("field info of message cosmos.base.abci.v1beta1.TxResponse is not mutable"))
	case "cosmos.base.abci.v1beta1.TxResponse.gas_wanted":
		panic(fmt.Errorf("field gas_wanted of message cosmos.base.abci.v1beta1.TxResponse is not mutable"))
	case "cosmos.base.abci.v1beta1.TxResponse.gas_used":
		panic(fmt.Errorf("field gas_used of message cosmos.base.abci.v1beta1.TxResponse is not mutable"))
	case "cosmos.base.abci.v1beta1.TxResponse.timestamp":
		panic(fmt.Errorf("field timestamp of message cosmos.base.abci.v1beta1.TxResponse is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.TxResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.TxResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_TxResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.TxResponse.height":
		return protoreflect.ValueOfInt64(int64(0))
	case "cosmos.base.abci.v1beta1.TxResponse.txhash":
		return protoreflect.ValueOfString("")
	case "cosmos.base.abci.v1beta1.TxResponse.codespace":
		return protoreflect.ValueOfString("")
	case "cosmos.base.abci.v1beta1.TxResponse.code":
		return protoreflect.ValueOfUint32(uint32(0))
	case "cosmos.base.abci.v1beta1.TxResponse.data":
		return protoreflect.ValueOfString("")
	case "cosmos.base.abci.v1beta1.TxResponse.raw_log":
		return protoreflect.ValueOfString("")
	case "cosmos.base.abci.v1beta1.TxResponse.logs":
		list := []*ABCIMessageLog{}
		return protoreflect.ValueOfList(&_TxResponse_7_list{list: &list})
	case "cosmos.base.abci.v1beta1.TxResponse.info":
		return protoreflect.ValueOfString("")
	case "cosmos.base.abci.v1beta1.TxResponse.gas_wanted":
		return protoreflect.ValueOfInt64(int64(0))
	case "cosmos.base.abci.v1beta1.TxResponse.gas_used":
		return protoreflect.ValueOfInt64(int64(0))
	case "cosmos.base.abci.v1beta1.TxResponse.tx":
		m := new(anypb.Any)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.base.abci.v1beta1.TxResponse.timestamp":
		return protoreflect.ValueOfString("")
	case "cosmos.base.abci.v1beta1.TxResponse.events":
		list := []*abci.Event{}
		return protoreflect.ValueOfList(&_TxResponse_13_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.TxResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.TxResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_TxResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.abci.v1beta1.TxResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_TxResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_TxResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_TxResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_TxResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*TxResponse)
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
		l = len(x.Txhash)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Codespace)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Code != 0 {
			n += 1 + runtime.Sov(uint64(x.Code))
		}
		l = len(x.Data)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.RawLog)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Logs) > 0 {
			for _, e := range x.Logs {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		l = len(x.Info)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.GasWanted != 0 {
			n += 1 + runtime.Sov(uint64(x.GasWanted))
		}
		if x.GasUsed != 0 {
			n += 1 + runtime.Sov(uint64(x.GasUsed))
		}
		if x.Tx != nil {
			l = options.Size(x.Tx)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Timestamp)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Events) > 0 {
			for _, e := range x.Events {
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
		x := input.Message.Interface().(*TxResponse)
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
		if len(x.Events) > 0 {
			for iNdEx := len(x.Events) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Events[iNdEx])
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
				dAtA[i] = 0x6a
			}
		}
		if len(x.Timestamp) > 0 {
			i -= len(x.Timestamp)
			copy(dAtA[i:], x.Timestamp)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Timestamp)))
			i--
			dAtA[i] = 0x62
		}
		if x.Tx != nil {
			encoded, err := options.Marshal(x.Tx)
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
			dAtA[i] = 0x5a
		}
		if x.GasUsed != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.GasUsed))
			i--
			dAtA[i] = 0x50
		}
		if x.GasWanted != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.GasWanted))
			i--
			dAtA[i] = 0x48
		}
		if len(x.Info) > 0 {
			i -= len(x.Info)
			copy(dAtA[i:], x.Info)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Info)))
			i--
			dAtA[i] = 0x42
		}
		if len(x.Logs) > 0 {
			for iNdEx := len(x.Logs) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Logs[iNdEx])
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
		}
		if len(x.RawLog) > 0 {
			i -= len(x.RawLog)
			copy(dAtA[i:], x.RawLog)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.RawLog)))
			i--
			dAtA[i] = 0x32
		}
		if len(x.Data) > 0 {
			i -= len(x.Data)
			copy(dAtA[i:], x.Data)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Data)))
			i--
			dAtA[i] = 0x2a
		}
		if x.Code != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Code))
			i--
			dAtA[i] = 0x20
		}
		if len(x.Codespace) > 0 {
			i -= len(x.Codespace)
			copy(dAtA[i:], x.Codespace)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Codespace)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Txhash) > 0 {
			i -= len(x.Txhash)
			copy(dAtA[i:], x.Txhash)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Txhash)))
			i--
			dAtA[i] = 0x12
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
		x := input.Message.Interface().(*TxResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: TxResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: TxResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Txhash", wireType)
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
				x.Txhash = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Codespace", wireType)
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
				x.Codespace = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
				}
				x.Code = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Code |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
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
				x.Data = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 6:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RawLog", wireType)
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
				x.RawLog = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 7:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Logs", wireType)
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
				x.Logs = append(x.Logs, &ABCIMessageLog{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Logs[len(x.Logs)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 8:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
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
				x.Info = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 9:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field GasWanted", wireType)
				}
				x.GasWanted = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.GasWanted |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 10:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field GasUsed", wireType)
				}
				x.GasUsed = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.GasUsed |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 11:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
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
				if x.Tx == nil {
					x.Tx = &anypb.Any{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Tx); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 12:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
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
				x.Timestamp = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 13:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
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
				x.Events = append(x.Events, &abci.Event{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Events[len(x.Events)-1]); err != nil {
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

var _ protoreflect.List = (*_ABCIMessageLog_3_list)(nil)

type _ABCIMessageLog_3_list struct {
	list *[]*StringEvent
}

func (x *_ABCIMessageLog_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_ABCIMessageLog_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_ABCIMessageLog_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*StringEvent)
	(*x.list)[i] = concreteValue
}

func (x *_ABCIMessageLog_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*StringEvent)
	*x.list = append(*x.list, concreteValue)
}

func (x *_ABCIMessageLog_3_list) AppendMutable() protoreflect.Value {
	v := new(StringEvent)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_ABCIMessageLog_3_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_ABCIMessageLog_3_list) NewElement() protoreflect.Value {
	v := new(StringEvent)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_ABCIMessageLog_3_list) IsValid() bool {
	return x.list != nil
}

var (
	md_ABCIMessageLog           protoreflect.MessageDescriptor
	fd_ABCIMessageLog_msg_index protoreflect.FieldDescriptor
	fd_ABCIMessageLog_log       protoreflect.FieldDescriptor
	fd_ABCIMessageLog_events    protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_abci_v1beta1_abci_proto_init()
	md_ABCIMessageLog = File_cosmos_base_abci_v1beta1_abci_proto.Messages().ByName("ABCIMessageLog")
	fd_ABCIMessageLog_msg_index = md_ABCIMessageLog.Fields().ByName("msg_index")
	fd_ABCIMessageLog_log = md_ABCIMessageLog.Fields().ByName("log")
	fd_ABCIMessageLog_events = md_ABCIMessageLog.Fields().ByName("events")
}

var _ protoreflect.Message = (*fastReflection_ABCIMessageLog)(nil)

type fastReflection_ABCIMessageLog ABCIMessageLog

func (x *ABCIMessageLog) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ABCIMessageLog)(x)
}

func (x *ABCIMessageLog) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_abci_v1beta1_abci_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ABCIMessageLog_messageType fastReflection_ABCIMessageLog_messageType
var _ protoreflect.MessageType = fastReflection_ABCIMessageLog_messageType{}

type fastReflection_ABCIMessageLog_messageType struct{}

func (x fastReflection_ABCIMessageLog_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ABCIMessageLog)(nil)
}
func (x fastReflection_ABCIMessageLog_messageType) New() protoreflect.Message {
	return new(fastReflection_ABCIMessageLog)
}
func (x fastReflection_ABCIMessageLog_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ABCIMessageLog
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ABCIMessageLog) Descriptor() protoreflect.MessageDescriptor {
	return md_ABCIMessageLog
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ABCIMessageLog) Type() protoreflect.MessageType {
	return _fastReflection_ABCIMessageLog_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ABCIMessageLog) New() protoreflect.Message {
	return new(fastReflection_ABCIMessageLog)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ABCIMessageLog) Interface() protoreflect.ProtoMessage {
	return (*ABCIMessageLog)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ABCIMessageLog) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.MsgIndex != uint32(0) {
		value := protoreflect.ValueOfUint32(x.MsgIndex)
		if !f(fd_ABCIMessageLog_msg_index, value) {
			return
		}
	}
	if x.Log != "" {
		value := protoreflect.ValueOfString(x.Log)
		if !f(fd_ABCIMessageLog_log, value) {
			return
		}
	}
	if len(x.Events) != 0 {
		value := protoreflect.ValueOfList(&_ABCIMessageLog_3_list{list: &x.Events})
		if !f(fd_ABCIMessageLog_events, value) {
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
func (x *fastReflection_ABCIMessageLog) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.ABCIMessageLog.msg_index":
		return x.MsgIndex != uint32(0)
	case "cosmos.base.abci.v1beta1.ABCIMessageLog.log":
		return x.Log != ""
	case "cosmos.base.abci.v1beta1.ABCIMessageLog.events":
		return len(x.Events) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.ABCIMessageLog"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.ABCIMessageLog does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ABCIMessageLog) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.ABCIMessageLog.msg_index":
		x.MsgIndex = uint32(0)
	case "cosmos.base.abci.v1beta1.ABCIMessageLog.log":
		x.Log = ""
	case "cosmos.base.abci.v1beta1.ABCIMessageLog.events":
		x.Events = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.ABCIMessageLog"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.ABCIMessageLog does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ABCIMessageLog) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.abci.v1beta1.ABCIMessageLog.msg_index":
		value := x.MsgIndex
		return protoreflect.ValueOfUint32(value)
	case "cosmos.base.abci.v1beta1.ABCIMessageLog.log":
		value := x.Log
		return protoreflect.ValueOfString(value)
	case "cosmos.base.abci.v1beta1.ABCIMessageLog.events":
		if len(x.Events) == 0 {
			return protoreflect.ValueOfList(&_ABCIMessageLog_3_list{})
		}
		listValue := &_ABCIMessageLog_3_list{list: &x.Events}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.ABCIMessageLog"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.ABCIMessageLog does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_ABCIMessageLog) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.ABCIMessageLog.msg_index":
		x.MsgIndex = uint32(value.Uint())
	case "cosmos.base.abci.v1beta1.ABCIMessageLog.log":
		x.Log = value.Interface().(string)
	case "cosmos.base.abci.v1beta1.ABCIMessageLog.events":
		lv := value.List()
		clv := lv.(*_ABCIMessageLog_3_list)
		x.Events = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.ABCIMessageLog"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.ABCIMessageLog does not contain field %s", fd.FullName()))
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
func (x *fastReflection_ABCIMessageLog) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.ABCIMessageLog.events":
		if x.Events == nil {
			x.Events = []*StringEvent{}
		}
		value := &_ABCIMessageLog_3_list{list: &x.Events}
		return protoreflect.ValueOfList(value)
	case "cosmos.base.abci.v1beta1.ABCIMessageLog.msg_index":
		panic(fmt.Errorf("field msg_index of message cosmos.base.abci.v1beta1.ABCIMessageLog is not mutable"))
	case "cosmos.base.abci.v1beta1.ABCIMessageLog.log":
		panic(fmt.Errorf("field log of message cosmos.base.abci.v1beta1.ABCIMessageLog is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.ABCIMessageLog"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.ABCIMessageLog does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ABCIMessageLog) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.ABCIMessageLog.msg_index":
		return protoreflect.ValueOfUint32(uint32(0))
	case "cosmos.base.abci.v1beta1.ABCIMessageLog.log":
		return protoreflect.ValueOfString("")
	case "cosmos.base.abci.v1beta1.ABCIMessageLog.events":
		list := []*StringEvent{}
		return protoreflect.ValueOfList(&_ABCIMessageLog_3_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.ABCIMessageLog"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.ABCIMessageLog does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ABCIMessageLog) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.abci.v1beta1.ABCIMessageLog", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ABCIMessageLog) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ABCIMessageLog) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_ABCIMessageLog) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ABCIMessageLog) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ABCIMessageLog)
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
		if x.MsgIndex != 0 {
			n += 1 + runtime.Sov(uint64(x.MsgIndex))
		}
		l = len(x.Log)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Events) > 0 {
			for _, e := range x.Events {
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
		x := input.Message.Interface().(*ABCIMessageLog)
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
		if len(x.Events) > 0 {
			for iNdEx := len(x.Events) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Events[iNdEx])
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
		if len(x.Log) > 0 {
			i -= len(x.Log)
			copy(dAtA[i:], x.Log)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Log)))
			i--
			dAtA[i] = 0x12
		}
		if x.MsgIndex != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.MsgIndex))
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
		x := input.Message.Interface().(*ABCIMessageLog)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ABCIMessageLog: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ABCIMessageLog: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MsgIndex", wireType)
				}
				x.MsgIndex = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.MsgIndex |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Log", wireType)
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
				x.Log = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
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
				x.Events = append(x.Events, &StringEvent{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Events[len(x.Events)-1]); err != nil {
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

var _ protoreflect.List = (*_StringEvent_2_list)(nil)

type _StringEvent_2_list struct {
	list *[]*Attribute
}

func (x *_StringEvent_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_StringEvent_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_StringEvent_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Attribute)
	(*x.list)[i] = concreteValue
}

func (x *_StringEvent_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Attribute)
	*x.list = append(*x.list, concreteValue)
}

func (x *_StringEvent_2_list) AppendMutable() protoreflect.Value {
	v := new(Attribute)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_StringEvent_2_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_StringEvent_2_list) NewElement() protoreflect.Value {
	v := new(Attribute)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_StringEvent_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_StringEvent            protoreflect.MessageDescriptor
	fd_StringEvent_type       protoreflect.FieldDescriptor
	fd_StringEvent_attributes protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_abci_v1beta1_abci_proto_init()
	md_StringEvent = File_cosmos_base_abci_v1beta1_abci_proto.Messages().ByName("StringEvent")
	fd_StringEvent_type = md_StringEvent.Fields().ByName("type")
	fd_StringEvent_attributes = md_StringEvent.Fields().ByName("attributes")
}

var _ protoreflect.Message = (*fastReflection_StringEvent)(nil)

type fastReflection_StringEvent StringEvent

func (x *StringEvent) ProtoReflect() protoreflect.Message {
	return (*fastReflection_StringEvent)(x)
}

func (x *StringEvent) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_abci_v1beta1_abci_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_StringEvent_messageType fastReflection_StringEvent_messageType
var _ protoreflect.MessageType = fastReflection_StringEvent_messageType{}

type fastReflection_StringEvent_messageType struct{}

func (x fastReflection_StringEvent_messageType) Zero() protoreflect.Message {
	return (*fastReflection_StringEvent)(nil)
}
func (x fastReflection_StringEvent_messageType) New() protoreflect.Message {
	return new(fastReflection_StringEvent)
}
func (x fastReflection_StringEvent_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_StringEvent
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_StringEvent) Descriptor() protoreflect.MessageDescriptor {
	return md_StringEvent
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_StringEvent) Type() protoreflect.MessageType {
	return _fastReflection_StringEvent_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_StringEvent) New() protoreflect.Message {
	return new(fastReflection_StringEvent)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_StringEvent) Interface() protoreflect.ProtoMessage {
	return (*StringEvent)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_StringEvent) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Type_ != "" {
		value := protoreflect.ValueOfString(x.Type_)
		if !f(fd_StringEvent_type, value) {
			return
		}
	}
	if len(x.Attributes) != 0 {
		value := protoreflect.ValueOfList(&_StringEvent_2_list{list: &x.Attributes})
		if !f(fd_StringEvent_attributes, value) {
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
func (x *fastReflection_StringEvent) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.StringEvent.type":
		return x.Type_ != ""
	case "cosmos.base.abci.v1beta1.StringEvent.attributes":
		return len(x.Attributes) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.StringEvent"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.StringEvent does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_StringEvent) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.StringEvent.type":
		x.Type_ = ""
	case "cosmos.base.abci.v1beta1.StringEvent.attributes":
		x.Attributes = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.StringEvent"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.StringEvent does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_StringEvent) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.abci.v1beta1.StringEvent.type":
		value := x.Type_
		return protoreflect.ValueOfString(value)
	case "cosmos.base.abci.v1beta1.StringEvent.attributes":
		if len(x.Attributes) == 0 {
			return protoreflect.ValueOfList(&_StringEvent_2_list{})
		}
		listValue := &_StringEvent_2_list{list: &x.Attributes}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.StringEvent"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.StringEvent does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_StringEvent) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.StringEvent.type":
		x.Type_ = value.Interface().(string)
	case "cosmos.base.abci.v1beta1.StringEvent.attributes":
		lv := value.List()
		clv := lv.(*_StringEvent_2_list)
		x.Attributes = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.StringEvent"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.StringEvent does not contain field %s", fd.FullName()))
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
func (x *fastReflection_StringEvent) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.StringEvent.attributes":
		if x.Attributes == nil {
			x.Attributes = []*Attribute{}
		}
		value := &_StringEvent_2_list{list: &x.Attributes}
		return protoreflect.ValueOfList(value)
	case "cosmos.base.abci.v1beta1.StringEvent.type":
		panic(fmt.Errorf("field type of message cosmos.base.abci.v1beta1.StringEvent is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.StringEvent"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.StringEvent does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_StringEvent) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.StringEvent.type":
		return protoreflect.ValueOfString("")
	case "cosmos.base.abci.v1beta1.StringEvent.attributes":
		list := []*Attribute{}
		return protoreflect.ValueOfList(&_StringEvent_2_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.StringEvent"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.StringEvent does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_StringEvent) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.abci.v1beta1.StringEvent", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_StringEvent) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_StringEvent) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_StringEvent) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_StringEvent) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*StringEvent)
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
		l = len(x.Type_)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Attributes) > 0 {
			for _, e := range x.Attributes {
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
		x := input.Message.Interface().(*StringEvent)
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
		if len(x.Attributes) > 0 {
			for iNdEx := len(x.Attributes) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Attributes[iNdEx])
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
		if len(x.Type_) > 0 {
			i -= len(x.Type_)
			copy(dAtA[i:], x.Type_)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Type_)))
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
		x := input.Message.Interface().(*StringEvent)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: StringEvent: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: StringEvent: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Type_", wireType)
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
				x.Type_ = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Attributes", wireType)
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
				x.Attributes = append(x.Attributes, &Attribute{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Attributes[len(x.Attributes)-1]); err != nil {
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
	md_Attribute       protoreflect.MessageDescriptor
	fd_Attribute_key   protoreflect.FieldDescriptor
	fd_Attribute_value protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_abci_v1beta1_abci_proto_init()
	md_Attribute = File_cosmos_base_abci_v1beta1_abci_proto.Messages().ByName("Attribute")
	fd_Attribute_key = md_Attribute.Fields().ByName("key")
	fd_Attribute_value = md_Attribute.Fields().ByName("value")
}

var _ protoreflect.Message = (*fastReflection_Attribute)(nil)

type fastReflection_Attribute Attribute

func (x *Attribute) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Attribute)(x)
}

func (x *Attribute) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_abci_v1beta1_abci_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Attribute_messageType fastReflection_Attribute_messageType
var _ protoreflect.MessageType = fastReflection_Attribute_messageType{}

type fastReflection_Attribute_messageType struct{}

func (x fastReflection_Attribute_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Attribute)(nil)
}
func (x fastReflection_Attribute_messageType) New() protoreflect.Message {
	return new(fastReflection_Attribute)
}
func (x fastReflection_Attribute_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Attribute
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Attribute) Descriptor() protoreflect.MessageDescriptor {
	return md_Attribute
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Attribute) Type() protoreflect.MessageType {
	return _fastReflection_Attribute_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Attribute) New() protoreflect.Message {
	return new(fastReflection_Attribute)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Attribute) Interface() protoreflect.ProtoMessage {
	return (*Attribute)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Attribute) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Key != "" {
		value := protoreflect.ValueOfString(x.Key)
		if !f(fd_Attribute_key, value) {
			return
		}
	}
	if x.Value != "" {
		value := protoreflect.ValueOfString(x.Value)
		if !f(fd_Attribute_value, value) {
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
func (x *fastReflection_Attribute) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.Attribute.key":
		return x.Key != ""
	case "cosmos.base.abci.v1beta1.Attribute.value":
		return x.Value != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.Attribute"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.Attribute does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Attribute) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.Attribute.key":
		x.Key = ""
	case "cosmos.base.abci.v1beta1.Attribute.value":
		x.Value = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.Attribute"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.Attribute does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Attribute) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.abci.v1beta1.Attribute.key":
		value := x.Key
		return protoreflect.ValueOfString(value)
	case "cosmos.base.abci.v1beta1.Attribute.value":
		value := x.Value
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.Attribute"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.Attribute does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Attribute) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.Attribute.key":
		x.Key = value.Interface().(string)
	case "cosmos.base.abci.v1beta1.Attribute.value":
		x.Value = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.Attribute"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.Attribute does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Attribute) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.Attribute.key":
		panic(fmt.Errorf("field key of message cosmos.base.abci.v1beta1.Attribute is not mutable"))
	case "cosmos.base.abci.v1beta1.Attribute.value":
		panic(fmt.Errorf("field value of message cosmos.base.abci.v1beta1.Attribute is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.Attribute"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.Attribute does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Attribute) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.Attribute.key":
		return protoreflect.ValueOfString("")
	case "cosmos.base.abci.v1beta1.Attribute.value":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.Attribute"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.Attribute does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Attribute) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.abci.v1beta1.Attribute", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Attribute) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Attribute) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Attribute) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Attribute) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Attribute)
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
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Attribute)
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
		x := input.Message.Interface().(*Attribute)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Attribute: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Attribute: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
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
				x.Key = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
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
				x.Value = string(dAtA[iNdEx:postIndex])
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
	md_GasInfo            protoreflect.MessageDescriptor
	fd_GasInfo_gas_wanted protoreflect.FieldDescriptor
	fd_GasInfo_gas_used   protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_abci_v1beta1_abci_proto_init()
	md_GasInfo = File_cosmos_base_abci_v1beta1_abci_proto.Messages().ByName("GasInfo")
	fd_GasInfo_gas_wanted = md_GasInfo.Fields().ByName("gas_wanted")
	fd_GasInfo_gas_used = md_GasInfo.Fields().ByName("gas_used")
}

var _ protoreflect.Message = (*fastReflection_GasInfo)(nil)

type fastReflection_GasInfo GasInfo

func (x *GasInfo) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GasInfo)(x)
}

func (x *GasInfo) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_abci_v1beta1_abci_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GasInfo_messageType fastReflection_GasInfo_messageType
var _ protoreflect.MessageType = fastReflection_GasInfo_messageType{}

type fastReflection_GasInfo_messageType struct{}

func (x fastReflection_GasInfo_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GasInfo)(nil)
}
func (x fastReflection_GasInfo_messageType) New() protoreflect.Message {
	return new(fastReflection_GasInfo)
}
func (x fastReflection_GasInfo_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GasInfo
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GasInfo) Descriptor() protoreflect.MessageDescriptor {
	return md_GasInfo
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GasInfo) Type() protoreflect.MessageType {
	return _fastReflection_GasInfo_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GasInfo) New() protoreflect.Message {
	return new(fastReflection_GasInfo)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GasInfo) Interface() protoreflect.ProtoMessage {
	return (*GasInfo)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GasInfo) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.GasWanted != uint64(0) {
		value := protoreflect.ValueOfUint64(x.GasWanted)
		if !f(fd_GasInfo_gas_wanted, value) {
			return
		}
	}
	if x.GasUsed != uint64(0) {
		value := protoreflect.ValueOfUint64(x.GasUsed)
		if !f(fd_GasInfo_gas_used, value) {
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
func (x *fastReflection_GasInfo) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.GasInfo.gas_wanted":
		return x.GasWanted != uint64(0)
	case "cosmos.base.abci.v1beta1.GasInfo.gas_used":
		return x.GasUsed != uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.GasInfo"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.GasInfo does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GasInfo) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.GasInfo.gas_wanted":
		x.GasWanted = uint64(0)
	case "cosmos.base.abci.v1beta1.GasInfo.gas_used":
		x.GasUsed = uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.GasInfo"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.GasInfo does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GasInfo) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.abci.v1beta1.GasInfo.gas_wanted":
		value := x.GasWanted
		return protoreflect.ValueOfUint64(value)
	case "cosmos.base.abci.v1beta1.GasInfo.gas_used":
		value := x.GasUsed
		return protoreflect.ValueOfUint64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.GasInfo"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.GasInfo does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_GasInfo) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.GasInfo.gas_wanted":
		x.GasWanted = value.Uint()
	case "cosmos.base.abci.v1beta1.GasInfo.gas_used":
		x.GasUsed = value.Uint()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.GasInfo"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.GasInfo does not contain field %s", fd.FullName()))
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
func (x *fastReflection_GasInfo) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.GasInfo.gas_wanted":
		panic(fmt.Errorf("field gas_wanted of message cosmos.base.abci.v1beta1.GasInfo is not mutable"))
	case "cosmos.base.abci.v1beta1.GasInfo.gas_used":
		panic(fmt.Errorf("field gas_used of message cosmos.base.abci.v1beta1.GasInfo is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.GasInfo"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.GasInfo does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GasInfo) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.GasInfo.gas_wanted":
		return protoreflect.ValueOfUint64(uint64(0))
	case "cosmos.base.abci.v1beta1.GasInfo.gas_used":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.GasInfo"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.GasInfo does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GasInfo) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.abci.v1beta1.GasInfo", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GasInfo) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GasInfo) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_GasInfo) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GasInfo) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GasInfo)
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
		if x.GasWanted != 0 {
			n += 1 + runtime.Sov(uint64(x.GasWanted))
		}
		if x.GasUsed != 0 {
			n += 1 + runtime.Sov(uint64(x.GasUsed))
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
		x := input.Message.Interface().(*GasInfo)
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
		if x.GasUsed != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.GasUsed))
			i--
			dAtA[i] = 0x10
		}
		if x.GasWanted != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.GasWanted))
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
		x := input.Message.Interface().(*GasInfo)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GasInfo: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GasInfo: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field GasWanted", wireType)
				}
				x.GasWanted = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.GasWanted |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field GasUsed", wireType)
				}
				x.GasUsed = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.GasUsed |= uint64(b&0x7F) << shift
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

var _ protoreflect.List = (*_Result_3_list)(nil)

type _Result_3_list struct {
	list *[]*abci.Event
}

func (x *_Result_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_Result_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_Result_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*abci.Event)
	(*x.list)[i] = concreteValue
}

func (x *_Result_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*abci.Event)
	*x.list = append(*x.list, concreteValue)
}

func (x *_Result_3_list) AppendMutable() protoreflect.Value {
	v := new(abci.Event)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Result_3_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_Result_3_list) NewElement() protoreflect.Value {
	v := new(abci.Event)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Result_3_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_Result_4_list)(nil)

type _Result_4_list struct {
	list *[]*anypb.Any
}

func (x *_Result_4_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_Result_4_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_Result_4_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*anypb.Any)
	(*x.list)[i] = concreteValue
}

func (x *_Result_4_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*anypb.Any)
	*x.list = append(*x.list, concreteValue)
}

func (x *_Result_4_list) AppendMutable() protoreflect.Value {
	v := new(anypb.Any)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Result_4_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_Result_4_list) NewElement() protoreflect.Value {
	v := new(anypb.Any)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Result_4_list) IsValid() bool {
	return x.list != nil
}

var (
	md_Result               protoreflect.MessageDescriptor
	fd_Result_data          protoreflect.FieldDescriptor
	fd_Result_log           protoreflect.FieldDescriptor
	fd_Result_events        protoreflect.FieldDescriptor
	fd_Result_msg_responses protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_abci_v1beta1_abci_proto_init()
	md_Result = File_cosmos_base_abci_v1beta1_abci_proto.Messages().ByName("Result")
	fd_Result_data = md_Result.Fields().ByName("data")
	fd_Result_log = md_Result.Fields().ByName("log")
	fd_Result_events = md_Result.Fields().ByName("events")
	fd_Result_msg_responses = md_Result.Fields().ByName("msg_responses")
}

var _ protoreflect.Message = (*fastReflection_Result)(nil)

type fastReflection_Result Result

func (x *Result) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Result)(x)
}

func (x *Result) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_abci_v1beta1_abci_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Result_messageType fastReflection_Result_messageType
var _ protoreflect.MessageType = fastReflection_Result_messageType{}

type fastReflection_Result_messageType struct{}

func (x fastReflection_Result_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Result)(nil)
}
func (x fastReflection_Result_messageType) New() protoreflect.Message {
	return new(fastReflection_Result)
}
func (x fastReflection_Result_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Result
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Result) Descriptor() protoreflect.MessageDescriptor {
	return md_Result
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Result) Type() protoreflect.MessageType {
	return _fastReflection_Result_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Result) New() protoreflect.Message {
	return new(fastReflection_Result)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Result) Interface() protoreflect.ProtoMessage {
	return (*Result)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Result) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Data) != 0 {
		value := protoreflect.ValueOfBytes(x.Data)
		if !f(fd_Result_data, value) {
			return
		}
	}
	if x.Log != "" {
		value := protoreflect.ValueOfString(x.Log)
		if !f(fd_Result_log, value) {
			return
		}
	}
	if len(x.Events) != 0 {
		value := protoreflect.ValueOfList(&_Result_3_list{list: &x.Events})
		if !f(fd_Result_events, value) {
			return
		}
	}
	if len(x.MsgResponses) != 0 {
		value := protoreflect.ValueOfList(&_Result_4_list{list: &x.MsgResponses})
		if !f(fd_Result_msg_responses, value) {
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
func (x *fastReflection_Result) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.Result.data":
		return len(x.Data) != 0
	case "cosmos.base.abci.v1beta1.Result.log":
		return x.Log != ""
	case "cosmos.base.abci.v1beta1.Result.events":
		return len(x.Events) != 0
	case "cosmos.base.abci.v1beta1.Result.msg_responses":
		return len(x.MsgResponses) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.Result"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.Result does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Result) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.Result.data":
		x.Data = nil
	case "cosmos.base.abci.v1beta1.Result.log":
		x.Log = ""
	case "cosmos.base.abci.v1beta1.Result.events":
		x.Events = nil
	case "cosmos.base.abci.v1beta1.Result.msg_responses":
		x.MsgResponses = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.Result"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.Result does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Result) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.abci.v1beta1.Result.data":
		value := x.Data
		return protoreflect.ValueOfBytes(value)
	case "cosmos.base.abci.v1beta1.Result.log":
		value := x.Log
		return protoreflect.ValueOfString(value)
	case "cosmos.base.abci.v1beta1.Result.events":
		if len(x.Events) == 0 {
			return protoreflect.ValueOfList(&_Result_3_list{})
		}
		listValue := &_Result_3_list{list: &x.Events}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.base.abci.v1beta1.Result.msg_responses":
		if len(x.MsgResponses) == 0 {
			return protoreflect.ValueOfList(&_Result_4_list{})
		}
		listValue := &_Result_4_list{list: &x.MsgResponses}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.Result"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.Result does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Result) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.Result.data":
		x.Data = value.Bytes()
	case "cosmos.base.abci.v1beta1.Result.log":
		x.Log = value.Interface().(string)
	case "cosmos.base.abci.v1beta1.Result.events":
		lv := value.List()
		clv := lv.(*_Result_3_list)
		x.Events = *clv.list
	case "cosmos.base.abci.v1beta1.Result.msg_responses":
		lv := value.List()
		clv := lv.(*_Result_4_list)
		x.MsgResponses = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.Result"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.Result does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Result) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.Result.events":
		if x.Events == nil {
			x.Events = []*abci.Event{}
		}
		value := &_Result_3_list{list: &x.Events}
		return protoreflect.ValueOfList(value)
	case "cosmos.base.abci.v1beta1.Result.msg_responses":
		if x.MsgResponses == nil {
			x.MsgResponses = []*anypb.Any{}
		}
		value := &_Result_4_list{list: &x.MsgResponses}
		return protoreflect.ValueOfList(value)
	case "cosmos.base.abci.v1beta1.Result.data":
		panic(fmt.Errorf("field data of message cosmos.base.abci.v1beta1.Result is not mutable"))
	case "cosmos.base.abci.v1beta1.Result.log":
		panic(fmt.Errorf("field log of message cosmos.base.abci.v1beta1.Result is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.Result"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.Result does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Result) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.Result.data":
		return protoreflect.ValueOfBytes(nil)
	case "cosmos.base.abci.v1beta1.Result.log":
		return protoreflect.ValueOfString("")
	case "cosmos.base.abci.v1beta1.Result.events":
		list := []*abci.Event{}
		return protoreflect.ValueOfList(&_Result_3_list{list: &list})
	case "cosmos.base.abci.v1beta1.Result.msg_responses":
		list := []*anypb.Any{}
		return protoreflect.ValueOfList(&_Result_4_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.Result"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.Result does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Result) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.abci.v1beta1.Result", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Result) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Result) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Result) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Result) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Result)
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
		l = len(x.Data)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Log)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Events) > 0 {
			for _, e := range x.Events {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if len(x.MsgResponses) > 0 {
			for _, e := range x.MsgResponses {
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
		x := input.Message.Interface().(*Result)
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
		if len(x.MsgResponses) > 0 {
			for iNdEx := len(x.MsgResponses) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.MsgResponses[iNdEx])
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
		if len(x.Events) > 0 {
			for iNdEx := len(x.Events) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Events[iNdEx])
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
		if len(x.Log) > 0 {
			i -= len(x.Log)
			copy(dAtA[i:], x.Log)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Log)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Data) > 0 {
			i -= len(x.Data)
			copy(dAtA[i:], x.Data)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Data)))
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
		x := input.Message.Interface().(*Result)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Result: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Result: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
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
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Log", wireType)
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
				x.Log = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
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
				x.Events = append(x.Events, &abci.Event{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Events[len(x.Events)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MsgResponses", wireType)
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
				x.MsgResponses = append(x.MsgResponses, &anypb.Any{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.MsgResponses[len(x.MsgResponses)-1]); err != nil {
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
	md_SimulationResponse          protoreflect.MessageDescriptor
	fd_SimulationResponse_gas_info protoreflect.FieldDescriptor
	fd_SimulationResponse_result   protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_abci_v1beta1_abci_proto_init()
	md_SimulationResponse = File_cosmos_base_abci_v1beta1_abci_proto.Messages().ByName("SimulationResponse")
	fd_SimulationResponse_gas_info = md_SimulationResponse.Fields().ByName("gas_info")
	fd_SimulationResponse_result = md_SimulationResponse.Fields().ByName("result")
}

var _ protoreflect.Message = (*fastReflection_SimulationResponse)(nil)

type fastReflection_SimulationResponse SimulationResponse

func (x *SimulationResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_SimulationResponse)(x)
}

func (x *SimulationResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_abci_v1beta1_abci_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_SimulationResponse_messageType fastReflection_SimulationResponse_messageType
var _ protoreflect.MessageType = fastReflection_SimulationResponse_messageType{}

type fastReflection_SimulationResponse_messageType struct{}

func (x fastReflection_SimulationResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_SimulationResponse)(nil)
}
func (x fastReflection_SimulationResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_SimulationResponse)
}
func (x fastReflection_SimulationResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_SimulationResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_SimulationResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_SimulationResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_SimulationResponse) Type() protoreflect.MessageType {
	return _fastReflection_SimulationResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_SimulationResponse) New() protoreflect.Message {
	return new(fastReflection_SimulationResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_SimulationResponse) Interface() protoreflect.ProtoMessage {
	return (*SimulationResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_SimulationResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.GasInfo != nil {
		value := protoreflect.ValueOfMessage(x.GasInfo.ProtoReflect())
		if !f(fd_SimulationResponse_gas_info, value) {
			return
		}
	}
	if x.Result != nil {
		value := protoreflect.ValueOfMessage(x.Result.ProtoReflect())
		if !f(fd_SimulationResponse_result, value) {
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
func (x *fastReflection_SimulationResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.SimulationResponse.gas_info":
		return x.GasInfo != nil
	case "cosmos.base.abci.v1beta1.SimulationResponse.result":
		return x.Result != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.SimulationResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.SimulationResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SimulationResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.SimulationResponse.gas_info":
		x.GasInfo = nil
	case "cosmos.base.abci.v1beta1.SimulationResponse.result":
		x.Result = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.SimulationResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.SimulationResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_SimulationResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.abci.v1beta1.SimulationResponse.gas_info":
		value := x.GasInfo
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.base.abci.v1beta1.SimulationResponse.result":
		value := x.Result
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.SimulationResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.SimulationResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_SimulationResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.SimulationResponse.gas_info":
		x.GasInfo = value.Message().Interface().(*GasInfo)
	case "cosmos.base.abci.v1beta1.SimulationResponse.result":
		x.Result = value.Message().Interface().(*Result)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.SimulationResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.SimulationResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_SimulationResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.SimulationResponse.gas_info":
		if x.GasInfo == nil {
			x.GasInfo = new(GasInfo)
		}
		return protoreflect.ValueOfMessage(x.GasInfo.ProtoReflect())
	case "cosmos.base.abci.v1beta1.SimulationResponse.result":
		if x.Result == nil {
			x.Result = new(Result)
		}
		return protoreflect.ValueOfMessage(x.Result.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.SimulationResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.SimulationResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_SimulationResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.SimulationResponse.gas_info":
		m := new(GasInfo)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.base.abci.v1beta1.SimulationResponse.result":
		m := new(Result)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.SimulationResponse"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.SimulationResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_SimulationResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.abci.v1beta1.SimulationResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_SimulationResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SimulationResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_SimulationResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_SimulationResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*SimulationResponse)
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
		if x.GasInfo != nil {
			l = options.Size(x.GasInfo)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Result != nil {
			l = options.Size(x.Result)
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
		x := input.Message.Interface().(*SimulationResponse)
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
		if x.Result != nil {
			encoded, err := options.Marshal(x.Result)
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
		if x.GasInfo != nil {
			encoded, err := options.Marshal(x.GasInfo)
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
		x := input.Message.Interface().(*SimulationResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SimulationResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SimulationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field GasInfo", wireType)
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
				if x.GasInfo == nil {
					x.GasInfo = &GasInfo{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.GasInfo); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
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
				if x.Result == nil {
					x.Result = &Result{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Result); err != nil {
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
	md_MsgData          protoreflect.MessageDescriptor
	fd_MsgData_msg_type protoreflect.FieldDescriptor
	fd_MsgData_data     protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_abci_v1beta1_abci_proto_init()
	md_MsgData = File_cosmos_base_abci_v1beta1_abci_proto.Messages().ByName("MsgData")
	fd_MsgData_msg_type = md_MsgData.Fields().ByName("msg_type")
	fd_MsgData_data = md_MsgData.Fields().ByName("data")
}

var _ protoreflect.Message = (*fastReflection_MsgData)(nil)

type fastReflection_MsgData MsgData

func (x *MsgData) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgData)(x)
}

func (x *MsgData) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_abci_v1beta1_abci_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgData_messageType fastReflection_MsgData_messageType
var _ protoreflect.MessageType = fastReflection_MsgData_messageType{}

type fastReflection_MsgData_messageType struct{}

func (x fastReflection_MsgData_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgData)(nil)
}
func (x fastReflection_MsgData_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgData)
}
func (x fastReflection_MsgData_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgData
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgData) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgData
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgData) Type() protoreflect.MessageType {
	return _fastReflection_MsgData_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgData) New() protoreflect.Message {
	return new(fastReflection_MsgData)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgData) Interface() protoreflect.ProtoMessage {
	return (*MsgData)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgData) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.MsgType != "" {
		value := protoreflect.ValueOfString(x.MsgType)
		if !f(fd_MsgData_msg_type, value) {
			return
		}
	}
	if len(x.Data) != 0 {
		value := protoreflect.ValueOfBytes(x.Data)
		if !f(fd_MsgData_data, value) {
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
func (x *fastReflection_MsgData) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.MsgData.msg_type":
		return x.MsgType != ""
	case "cosmos.base.abci.v1beta1.MsgData.data":
		return len(x.Data) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.MsgData"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.MsgData does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgData) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.MsgData.msg_type":
		x.MsgType = ""
	case "cosmos.base.abci.v1beta1.MsgData.data":
		x.Data = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.MsgData"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.MsgData does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgData) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.abci.v1beta1.MsgData.msg_type":
		value := x.MsgType
		return protoreflect.ValueOfString(value)
	case "cosmos.base.abci.v1beta1.MsgData.data":
		value := x.Data
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.MsgData"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.MsgData does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgData) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.MsgData.msg_type":
		x.MsgType = value.Interface().(string)
	case "cosmos.base.abci.v1beta1.MsgData.data":
		x.Data = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.MsgData"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.MsgData does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgData) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.MsgData.msg_type":
		panic(fmt.Errorf("field msg_type of message cosmos.base.abci.v1beta1.MsgData is not mutable"))
	case "cosmos.base.abci.v1beta1.MsgData.data":
		panic(fmt.Errorf("field data of message cosmos.base.abci.v1beta1.MsgData is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.MsgData"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.MsgData does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgData) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.MsgData.msg_type":
		return protoreflect.ValueOfString("")
	case "cosmos.base.abci.v1beta1.MsgData.data":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.MsgData"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.MsgData does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgData) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.abci.v1beta1.MsgData", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgData) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgData) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgData) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgData) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgData)
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
		l = len(x.MsgType)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Data)
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
		x := input.Message.Interface().(*MsgData)
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
		if len(x.Data) > 0 {
			i -= len(x.Data)
			copy(dAtA[i:], x.Data)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Data)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.MsgType) > 0 {
			i -= len(x.MsgType)
			copy(dAtA[i:], x.MsgType)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.MsgType)))
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
		x := input.Message.Interface().(*MsgData)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgData: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgData: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MsgType", wireType)
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
				x.MsgType = string(dAtA[iNdEx:postIndex])
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

var _ protoreflect.List = (*_TxMsgData_1_list)(nil)

type _TxMsgData_1_list struct {
	list *[]*MsgData
}

func (x *_TxMsgData_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_TxMsgData_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_TxMsgData_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*MsgData)
	(*x.list)[i] = concreteValue
}

func (x *_TxMsgData_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*MsgData)
	*x.list = append(*x.list, concreteValue)
}

func (x *_TxMsgData_1_list) AppendMutable() protoreflect.Value {
	v := new(MsgData)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_TxMsgData_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_TxMsgData_1_list) NewElement() protoreflect.Value {
	v := new(MsgData)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_TxMsgData_1_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_TxMsgData_2_list)(nil)

type _TxMsgData_2_list struct {
	list *[]*anypb.Any
}

func (x *_TxMsgData_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_TxMsgData_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_TxMsgData_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*anypb.Any)
	(*x.list)[i] = concreteValue
}

func (x *_TxMsgData_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*anypb.Any)
	*x.list = append(*x.list, concreteValue)
}

func (x *_TxMsgData_2_list) AppendMutable() protoreflect.Value {
	v := new(anypb.Any)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_TxMsgData_2_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_TxMsgData_2_list) NewElement() protoreflect.Value {
	v := new(anypb.Any)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_TxMsgData_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_TxMsgData               protoreflect.MessageDescriptor
	fd_TxMsgData_data          protoreflect.FieldDescriptor
	fd_TxMsgData_msg_responses protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_abci_v1beta1_abci_proto_init()
	md_TxMsgData = File_cosmos_base_abci_v1beta1_abci_proto.Messages().ByName("TxMsgData")
	fd_TxMsgData_data = md_TxMsgData.Fields().ByName("data")
	fd_TxMsgData_msg_responses = md_TxMsgData.Fields().ByName("msg_responses")
}

var _ protoreflect.Message = (*fastReflection_TxMsgData)(nil)

type fastReflection_TxMsgData TxMsgData

func (x *TxMsgData) ProtoReflect() protoreflect.Message {
	return (*fastReflection_TxMsgData)(x)
}

func (x *TxMsgData) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_abci_v1beta1_abci_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_TxMsgData_messageType fastReflection_TxMsgData_messageType
var _ protoreflect.MessageType = fastReflection_TxMsgData_messageType{}

type fastReflection_TxMsgData_messageType struct{}

func (x fastReflection_TxMsgData_messageType) Zero() protoreflect.Message {
	return (*fastReflection_TxMsgData)(nil)
}
func (x fastReflection_TxMsgData_messageType) New() protoreflect.Message {
	return new(fastReflection_TxMsgData)
}
func (x fastReflection_TxMsgData_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_TxMsgData
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_TxMsgData) Descriptor() protoreflect.MessageDescriptor {
	return md_TxMsgData
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_TxMsgData) Type() protoreflect.MessageType {
	return _fastReflection_TxMsgData_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_TxMsgData) New() protoreflect.Message {
	return new(fastReflection_TxMsgData)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_TxMsgData) Interface() protoreflect.ProtoMessage {
	return (*TxMsgData)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_TxMsgData) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Data) != 0 {
		value := protoreflect.ValueOfList(&_TxMsgData_1_list{list: &x.Data})
		if !f(fd_TxMsgData_data, value) {
			return
		}
	}
	if len(x.MsgResponses) != 0 {
		value := protoreflect.ValueOfList(&_TxMsgData_2_list{list: &x.MsgResponses})
		if !f(fd_TxMsgData_msg_responses, value) {
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
func (x *fastReflection_TxMsgData) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.TxMsgData.data":
		return len(x.Data) != 0
	case "cosmos.base.abci.v1beta1.TxMsgData.msg_responses":
		return len(x.MsgResponses) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.TxMsgData"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.TxMsgData does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_TxMsgData) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.TxMsgData.data":
		x.Data = nil
	case "cosmos.base.abci.v1beta1.TxMsgData.msg_responses":
		x.MsgResponses = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.TxMsgData"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.TxMsgData does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_TxMsgData) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.abci.v1beta1.TxMsgData.data":
		if len(x.Data) == 0 {
			return protoreflect.ValueOfList(&_TxMsgData_1_list{})
		}
		listValue := &_TxMsgData_1_list{list: &x.Data}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.base.abci.v1beta1.TxMsgData.msg_responses":
		if len(x.MsgResponses) == 0 {
			return protoreflect.ValueOfList(&_TxMsgData_2_list{})
		}
		listValue := &_TxMsgData_2_list{list: &x.MsgResponses}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.TxMsgData"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.TxMsgData does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_TxMsgData) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.TxMsgData.data":
		lv := value.List()
		clv := lv.(*_TxMsgData_1_list)
		x.Data = *clv.list
	case "cosmos.base.abci.v1beta1.TxMsgData.msg_responses":
		lv := value.List()
		clv := lv.(*_TxMsgData_2_list)
		x.MsgResponses = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.TxMsgData"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.TxMsgData does not contain field %s", fd.FullName()))
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
func (x *fastReflection_TxMsgData) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.TxMsgData.data":
		if x.Data == nil {
			x.Data = []*MsgData{}
		}
		value := &_TxMsgData_1_list{list: &x.Data}
		return protoreflect.ValueOfList(value)
	case "cosmos.base.abci.v1beta1.TxMsgData.msg_responses":
		if x.MsgResponses == nil {
			x.MsgResponses = []*anypb.Any{}
		}
		value := &_TxMsgData_2_list{list: &x.MsgResponses}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.TxMsgData"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.TxMsgData does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_TxMsgData) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.TxMsgData.data":
		list := []*MsgData{}
		return protoreflect.ValueOfList(&_TxMsgData_1_list{list: &list})
	case "cosmos.base.abci.v1beta1.TxMsgData.msg_responses":
		list := []*anypb.Any{}
		return protoreflect.ValueOfList(&_TxMsgData_2_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.TxMsgData"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.TxMsgData does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_TxMsgData) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.abci.v1beta1.TxMsgData", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_TxMsgData) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_TxMsgData) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_TxMsgData) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_TxMsgData) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*TxMsgData)
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
		if len(x.Data) > 0 {
			for _, e := range x.Data {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if len(x.MsgResponses) > 0 {
			for _, e := range x.MsgResponses {
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
		x := input.Message.Interface().(*TxMsgData)
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
		if len(x.MsgResponses) > 0 {
			for iNdEx := len(x.MsgResponses) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.MsgResponses[iNdEx])
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
		if len(x.Data) > 0 {
			for iNdEx := len(x.Data) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Data[iNdEx])
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
		x := input.Message.Interface().(*TxMsgData)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: TxMsgData: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: TxMsgData: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
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
				x.Data = append(x.Data, &MsgData{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Data[len(x.Data)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MsgResponses", wireType)
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
				x.MsgResponses = append(x.MsgResponses, &anypb.Any{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.MsgResponses[len(x.MsgResponses)-1]); err != nil {
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

var _ protoreflect.List = (*_SearchTxsResult_6_list)(nil)

type _SearchTxsResult_6_list struct {
	list *[]*TxResponse
}

func (x *_SearchTxsResult_6_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_SearchTxsResult_6_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_SearchTxsResult_6_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*TxResponse)
	(*x.list)[i] = concreteValue
}

func (x *_SearchTxsResult_6_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*TxResponse)
	*x.list = append(*x.list, concreteValue)
}

func (x *_SearchTxsResult_6_list) AppendMutable() protoreflect.Value {
	v := new(TxResponse)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_SearchTxsResult_6_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_SearchTxsResult_6_list) NewElement() protoreflect.Value {
	v := new(TxResponse)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_SearchTxsResult_6_list) IsValid() bool {
	return x.list != nil
}

var (
	md_SearchTxsResult             protoreflect.MessageDescriptor
	fd_SearchTxsResult_total_count protoreflect.FieldDescriptor
	fd_SearchTxsResult_count       protoreflect.FieldDescriptor
	fd_SearchTxsResult_page_number protoreflect.FieldDescriptor
	fd_SearchTxsResult_page_total  protoreflect.FieldDescriptor
	fd_SearchTxsResult_limit       protoreflect.FieldDescriptor
	fd_SearchTxsResult_txs         protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_base_abci_v1beta1_abci_proto_init()
	md_SearchTxsResult = File_cosmos_base_abci_v1beta1_abci_proto.Messages().ByName("SearchTxsResult")
	fd_SearchTxsResult_total_count = md_SearchTxsResult.Fields().ByName("total_count")
	fd_SearchTxsResult_count = md_SearchTxsResult.Fields().ByName("count")
	fd_SearchTxsResult_page_number = md_SearchTxsResult.Fields().ByName("page_number")
	fd_SearchTxsResult_page_total = md_SearchTxsResult.Fields().ByName("page_total")
	fd_SearchTxsResult_limit = md_SearchTxsResult.Fields().ByName("limit")
	fd_SearchTxsResult_txs = md_SearchTxsResult.Fields().ByName("txs")
}

var _ protoreflect.Message = (*fastReflection_SearchTxsResult)(nil)

type fastReflection_SearchTxsResult SearchTxsResult

func (x *SearchTxsResult) ProtoReflect() protoreflect.Message {
	return (*fastReflection_SearchTxsResult)(x)
}

func (x *SearchTxsResult) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_base_abci_v1beta1_abci_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_SearchTxsResult_messageType fastReflection_SearchTxsResult_messageType
var _ protoreflect.MessageType = fastReflection_SearchTxsResult_messageType{}

type fastReflection_SearchTxsResult_messageType struct{}

func (x fastReflection_SearchTxsResult_messageType) Zero() protoreflect.Message {
	return (*fastReflection_SearchTxsResult)(nil)
}
func (x fastReflection_SearchTxsResult_messageType) New() protoreflect.Message {
	return new(fastReflection_SearchTxsResult)
}
func (x fastReflection_SearchTxsResult_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_SearchTxsResult
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_SearchTxsResult) Descriptor() protoreflect.MessageDescriptor {
	return md_SearchTxsResult
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_SearchTxsResult) Type() protoreflect.MessageType {
	return _fastReflection_SearchTxsResult_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_SearchTxsResult) New() protoreflect.Message {
	return new(fastReflection_SearchTxsResult)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_SearchTxsResult) Interface() protoreflect.ProtoMessage {
	return (*SearchTxsResult)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_SearchTxsResult) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.TotalCount != uint64(0) {
		value := protoreflect.ValueOfUint64(x.TotalCount)
		if !f(fd_SearchTxsResult_total_count, value) {
			return
		}
	}
	if x.Count != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Count)
		if !f(fd_SearchTxsResult_count, value) {
			return
		}
	}
	if x.PageNumber != uint64(0) {
		value := protoreflect.ValueOfUint64(x.PageNumber)
		if !f(fd_SearchTxsResult_page_number, value) {
			return
		}
	}
	if x.PageTotal != uint64(0) {
		value := protoreflect.ValueOfUint64(x.PageTotal)
		if !f(fd_SearchTxsResult_page_total, value) {
			return
		}
	}
	if x.Limit != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Limit)
		if !f(fd_SearchTxsResult_limit, value) {
			return
		}
	}
	if len(x.Txs) != 0 {
		value := protoreflect.ValueOfList(&_SearchTxsResult_6_list{list: &x.Txs})
		if !f(fd_SearchTxsResult_txs, value) {
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
func (x *fastReflection_SearchTxsResult) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.SearchTxsResult.total_count":
		return x.TotalCount != uint64(0)
	case "cosmos.base.abci.v1beta1.SearchTxsResult.count":
		return x.Count != uint64(0)
	case "cosmos.base.abci.v1beta1.SearchTxsResult.page_number":
		return x.PageNumber != uint64(0)
	case "cosmos.base.abci.v1beta1.SearchTxsResult.page_total":
		return x.PageTotal != uint64(0)
	case "cosmos.base.abci.v1beta1.SearchTxsResult.limit":
		return x.Limit != uint64(0)
	case "cosmos.base.abci.v1beta1.SearchTxsResult.txs":
		return len(x.Txs) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.SearchTxsResult"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.SearchTxsResult does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SearchTxsResult) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.SearchTxsResult.total_count":
		x.TotalCount = uint64(0)
	case "cosmos.base.abci.v1beta1.SearchTxsResult.count":
		x.Count = uint64(0)
	case "cosmos.base.abci.v1beta1.SearchTxsResult.page_number":
		x.PageNumber = uint64(0)
	case "cosmos.base.abci.v1beta1.SearchTxsResult.page_total":
		x.PageTotal = uint64(0)
	case "cosmos.base.abci.v1beta1.SearchTxsResult.limit":
		x.Limit = uint64(0)
	case "cosmos.base.abci.v1beta1.SearchTxsResult.txs":
		x.Txs = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.SearchTxsResult"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.SearchTxsResult does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_SearchTxsResult) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.base.abci.v1beta1.SearchTxsResult.total_count":
		value := x.TotalCount
		return protoreflect.ValueOfUint64(value)
	case "cosmos.base.abci.v1beta1.SearchTxsResult.count":
		value := x.Count
		return protoreflect.ValueOfUint64(value)
	case "cosmos.base.abci.v1beta1.SearchTxsResult.page_number":
		value := x.PageNumber
		return protoreflect.ValueOfUint64(value)
	case "cosmos.base.abci.v1beta1.SearchTxsResult.page_total":
		value := x.PageTotal
		return protoreflect.ValueOfUint64(value)
	case "cosmos.base.abci.v1beta1.SearchTxsResult.limit":
		value := x.Limit
		return protoreflect.ValueOfUint64(value)
	case "cosmos.base.abci.v1beta1.SearchTxsResult.txs":
		if len(x.Txs) == 0 {
			return protoreflect.ValueOfList(&_SearchTxsResult_6_list{})
		}
		listValue := &_SearchTxsResult_6_list{list: &x.Txs}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.SearchTxsResult"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.SearchTxsResult does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_SearchTxsResult) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.SearchTxsResult.total_count":
		x.TotalCount = value.Uint()
	case "cosmos.base.abci.v1beta1.SearchTxsResult.count":
		x.Count = value.Uint()
	case "cosmos.base.abci.v1beta1.SearchTxsResult.page_number":
		x.PageNumber = value.Uint()
	case "cosmos.base.abci.v1beta1.SearchTxsResult.page_total":
		x.PageTotal = value.Uint()
	case "cosmos.base.abci.v1beta1.SearchTxsResult.limit":
		x.Limit = value.Uint()
	case "cosmos.base.abci.v1beta1.SearchTxsResult.txs":
		lv := value.List()
		clv := lv.(*_SearchTxsResult_6_list)
		x.Txs = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.SearchTxsResult"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.SearchTxsResult does not contain field %s", fd.FullName()))
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
func (x *fastReflection_SearchTxsResult) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.SearchTxsResult.txs":
		if x.Txs == nil {
			x.Txs = []*TxResponse{}
		}
		value := &_SearchTxsResult_6_list{list: &x.Txs}
		return protoreflect.ValueOfList(value)
	case "cosmos.base.abci.v1beta1.SearchTxsResult.total_count":
		panic(fmt.Errorf("field total_count of message cosmos.base.abci.v1beta1.SearchTxsResult is not mutable"))
	case "cosmos.base.abci.v1beta1.SearchTxsResult.count":
		panic(fmt.Errorf("field count of message cosmos.base.abci.v1beta1.SearchTxsResult is not mutable"))
	case "cosmos.base.abci.v1beta1.SearchTxsResult.page_number":
		panic(fmt.Errorf("field page_number of message cosmos.base.abci.v1beta1.SearchTxsResult is not mutable"))
	case "cosmos.base.abci.v1beta1.SearchTxsResult.page_total":
		panic(fmt.Errorf("field page_total of message cosmos.base.abci.v1beta1.SearchTxsResult is not mutable"))
	case "cosmos.base.abci.v1beta1.SearchTxsResult.limit":
		panic(fmt.Errorf("field limit of message cosmos.base.abci.v1beta1.SearchTxsResult is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.SearchTxsResult"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.SearchTxsResult does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_SearchTxsResult) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.base.abci.v1beta1.SearchTxsResult.total_count":
		return protoreflect.ValueOfUint64(uint64(0))
	case "cosmos.base.abci.v1beta1.SearchTxsResult.count":
		return protoreflect.ValueOfUint64(uint64(0))
	case "cosmos.base.abci.v1beta1.SearchTxsResult.page_number":
		return protoreflect.ValueOfUint64(uint64(0))
	case "cosmos.base.abci.v1beta1.SearchTxsResult.page_total":
		return protoreflect.ValueOfUint64(uint64(0))
	case "cosmos.base.abci.v1beta1.SearchTxsResult.limit":
		return protoreflect.ValueOfUint64(uint64(0))
	case "cosmos.base.abci.v1beta1.SearchTxsResult.txs":
		list := []*TxResponse{}
		return protoreflect.ValueOfList(&_SearchTxsResult_6_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.base.abci.v1beta1.SearchTxsResult"))
		}
		panic(fmt.Errorf("message cosmos.base.abci.v1beta1.SearchTxsResult does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_SearchTxsResult) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.base.abci.v1beta1.SearchTxsResult", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_SearchTxsResult) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SearchTxsResult) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_SearchTxsResult) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_SearchTxsResult) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*SearchTxsResult)
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
		if x.TotalCount != 0 {
			n += 1 + runtime.Sov(uint64(x.TotalCount))
		}
		if x.Count != 0 {
			n += 1 + runtime.Sov(uint64(x.Count))
		}
		if x.PageNumber != 0 {
			n += 1 + runtime.Sov(uint64(x.PageNumber))
		}
		if x.PageTotal != 0 {
			n += 1 + runtime.Sov(uint64(x.PageTotal))
		}
		if x.Limit != 0 {
			n += 1 + runtime.Sov(uint64(x.Limit))
		}
		if len(x.Txs) > 0 {
			for _, e := range x.Txs {
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
		x := input.Message.Interface().(*SearchTxsResult)
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
				encoded, err := options.Marshal(x.Txs[iNdEx])
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
		}
		if x.Limit != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Limit))
			i--
			dAtA[i] = 0x28
		}
		if x.PageTotal != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.PageTotal))
			i--
			dAtA[i] = 0x20
		}
		if x.PageNumber != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.PageNumber))
			i--
			dAtA[i] = 0x18
		}
		if x.Count != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Count))
			i--
			dAtA[i] = 0x10
		}
		if x.TotalCount != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.TotalCount))
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
		x := input.Message.Interface().(*SearchTxsResult)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SearchTxsResult: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SearchTxsResult: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TotalCount", wireType)
				}
				x.TotalCount = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.TotalCount |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
				}
				x.Count = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Count |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field PageNumber", wireType)
				}
				x.PageNumber = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.PageNumber |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field PageTotal", wireType)
				}
				x.PageTotal = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.PageTotal |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 5:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
				}
				x.Limit = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Limit |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 6:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Txs", wireType)
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
				x.Txs = append(x.Txs, &TxResponse{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Txs[len(x.Txs)-1]); err != nil {
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
// source: cosmos/base/abci/v1beta1/abci.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// TxResponse defines a structure containing relevant tx data and metadata. The
// tags are stringified and the log is JSON decoded.
type TxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The block height
	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	// The transaction hash.
	Txhash string `protobuf:"bytes,2,opt,name=txhash,proto3" json:"txhash,omitempty"`
	// Namespace for the Code
	Codespace string `protobuf:"bytes,3,opt,name=codespace,proto3" json:"codespace,omitempty"`
	// Response code.
	Code uint32 `protobuf:"varint,4,opt,name=code,proto3" json:"code,omitempty"`
	// Result bytes, if any.
	Data string `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	// The output of the application's logger (raw string). May be
	// non-deterministic.
	RawLog string `protobuf:"bytes,6,opt,name=raw_log,json=rawLog,proto3" json:"raw_log,omitempty"`
	// The output of the application's logger (typed). May be non-deterministic.
	Logs []*ABCIMessageLog `protobuf:"bytes,7,rep,name=logs,proto3" json:"logs,omitempty"`
	// Additional information. May be non-deterministic.
	Info string `protobuf:"bytes,8,opt,name=info,proto3" json:"info,omitempty"`
	// Amount of gas requested for transaction.
	GasWanted int64 `protobuf:"varint,9,opt,name=gas_wanted,json=gasWanted,proto3" json:"gas_wanted,omitempty"`
	// Amount of gas consumed by transaction.
	GasUsed int64 `protobuf:"varint,10,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
	// The request transaction bytes.
	Tx *anypb.Any `protobuf:"bytes,11,opt,name=tx,proto3" json:"tx,omitempty"`
	// Time of the previous block. For heights > 1, it's the weighted median of
	// the timestamps of the valid votes in the block.LastCommit. For height == 1,
	// it's genesis time.
	Timestamp string `protobuf:"bytes,12,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Events defines all the events emitted by processing a transaction. Note,
	// these events include those emitted by processing all the messages and those
	// emitted from the ante handler. Whereas Logs contains the events, with
	// additional metadata, emitted only by processing the messages.
	//
	// Since: cosmos-sdk 0.42.11, 0.44.5, 0.45
	Events []*abci.Event `protobuf:"bytes,13,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *TxResponse) Reset() {
	*x = TxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_abci_v1beta1_abci_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxResponse) ProtoMessage() {}

// Deprecated: Use TxResponse.ProtoReflect.Descriptor instead.
func (*TxResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_base_abci_v1beta1_abci_proto_rawDescGZIP(), []int{0}
}

func (x *TxResponse) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *TxResponse) GetTxhash() string {
	if x != nil {
		return x.Txhash
	}
	return ""
}

func (x *TxResponse) GetCodespace() string {
	if x != nil {
		return x.Codespace
	}
	return ""
}

func (x *TxResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *TxResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *TxResponse) GetRawLog() string {
	if x != nil {
		return x.RawLog
	}
	return ""
}

func (x *TxResponse) GetLogs() []*ABCIMessageLog {
	if x != nil {
		return x.Logs
	}
	return nil
}

func (x *TxResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

func (x *TxResponse) GetGasWanted() int64 {
	if x != nil {
		return x.GasWanted
	}
	return 0
}

func (x *TxResponse) GetGasUsed() int64 {
	if x != nil {
		return x.GasUsed
	}
	return 0
}

func (x *TxResponse) GetTx() *anypb.Any {
	if x != nil {
		return x.Tx
	}
	return nil
}

func (x *TxResponse) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *TxResponse) GetEvents() []*abci.Event {
	if x != nil {
		return x.Events
	}
	return nil
}

// ABCIMessageLog defines a structure containing an indexed tx ABCI message log.
type ABCIMessageLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgIndex uint32 `protobuf:"varint,1,opt,name=msg_index,json=msgIndex,proto3" json:"msg_index,omitempty"`
	Log      string `protobuf:"bytes,2,opt,name=log,proto3" json:"log,omitempty"`
	// Events contains a slice of Event objects that were emitted during some
	// execution.
	Events []*StringEvent `protobuf:"bytes,3,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *ABCIMessageLog) Reset() {
	*x = ABCIMessageLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_abci_v1beta1_abci_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ABCIMessageLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ABCIMessageLog) ProtoMessage() {}

// Deprecated: Use ABCIMessageLog.ProtoReflect.Descriptor instead.
func (*ABCIMessageLog) Descriptor() ([]byte, []int) {
	return file_cosmos_base_abci_v1beta1_abci_proto_rawDescGZIP(), []int{1}
}

func (x *ABCIMessageLog) GetMsgIndex() uint32 {
	if x != nil {
		return x.MsgIndex
	}
	return 0
}

func (x *ABCIMessageLog) GetLog() string {
	if x != nil {
		return x.Log
	}
	return ""
}

func (x *ABCIMessageLog) GetEvents() []*StringEvent {
	if x != nil {
		return x.Events
	}
	return nil
}

// StringEvent defines en Event object wrapper where all the attributes
// contain key/value pairs that are strings instead of raw bytes.
type StringEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type_      string       `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Attributes []*Attribute `protobuf:"bytes,2,rep,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *StringEvent) Reset() {
	*x = StringEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_abci_v1beta1_abci_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringEvent) ProtoMessage() {}

// Deprecated: Use StringEvent.ProtoReflect.Descriptor instead.
func (*StringEvent) Descriptor() ([]byte, []int) {
	return file_cosmos_base_abci_v1beta1_abci_proto_rawDescGZIP(), []int{2}
}

func (x *StringEvent) GetType_() string {
	if x != nil {
		return x.Type_
	}
	return ""
}

func (x *StringEvent) GetAttributes() []*Attribute {
	if x != nil {
		return x.Attributes
	}
	return nil
}

// Attribute defines an attribute wrapper where the key and value are
// strings instead of raw bytes.
type Attribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Attribute) Reset() {
	*x = Attribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_abci_v1beta1_abci_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attribute) ProtoMessage() {}

// Deprecated: Use Attribute.ProtoReflect.Descriptor instead.
func (*Attribute) Descriptor() ([]byte, []int) {
	return file_cosmos_base_abci_v1beta1_abci_proto_rawDescGZIP(), []int{3}
}

func (x *Attribute) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Attribute) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// GasInfo defines tx execution gas context.
type GasInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// GasWanted is the maximum units of work we allow this tx to perform.
	GasWanted uint64 `protobuf:"varint,1,opt,name=gas_wanted,json=gasWanted,proto3" json:"gas_wanted,omitempty"`
	// GasUsed is the amount of gas actually consumed.
	GasUsed uint64 `protobuf:"varint,2,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
}

func (x *GasInfo) Reset() {
	*x = GasInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_abci_v1beta1_abci_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GasInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GasInfo) ProtoMessage() {}

// Deprecated: Use GasInfo.ProtoReflect.Descriptor instead.
func (*GasInfo) Descriptor() ([]byte, []int) {
	return file_cosmos_base_abci_v1beta1_abci_proto_rawDescGZIP(), []int{4}
}

func (x *GasInfo) GetGasWanted() uint64 {
	if x != nil {
		return x.GasWanted
	}
	return 0
}

func (x *GasInfo) GetGasUsed() uint64 {
	if x != nil {
		return x.GasUsed
	}
	return 0
}

// Result is the union of ResponseFormat and ResponseCheckTx.
type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Data is any data returned from message or handler execution. It MUST be
	// length prefixed in order to separate data from multiple message executions.
	// Deprecated. This field is still populated, but prefer msg_response instead
	// because it also contains the Msg response typeURL.
	//
	// Deprecated: Do not use.
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// Log contains the log information from message or handler execution.
	Log string `protobuf:"bytes,2,opt,name=log,proto3" json:"log,omitempty"`
	// Events contains a slice of Event objects that were emitted during message
	// or handler execution.
	Events []*abci.Event `protobuf:"bytes,3,rep,name=events,proto3" json:"events,omitempty"`
	// msg_responses contains the Msg handler responses type packed in Anys.
	//
	// Since: cosmos-sdk 0.45
	MsgResponses []*anypb.Any `protobuf:"bytes,4,rep,name=msg_responses,json=msgResponses,proto3" json:"msg_responses,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_abci_v1beta1_abci_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_cosmos_base_abci_v1beta1_abci_proto_rawDescGZIP(), []int{5}
}

// Deprecated: Do not use.
func (x *Result) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Result) GetLog() string {
	if x != nil {
		return x.Log
	}
	return ""
}

func (x *Result) GetEvents() []*abci.Event {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *Result) GetMsgResponses() []*anypb.Any {
	if x != nil {
		return x.MsgResponses
	}
	return nil
}

// SimulationResponse defines the response generated when a transaction is
// successfully simulated.
type SimulationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GasInfo *GasInfo `protobuf:"bytes,1,opt,name=gas_info,json=gasInfo,proto3" json:"gas_info,omitempty"`
	Result  *Result  `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *SimulationResponse) Reset() {
	*x = SimulationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_abci_v1beta1_abci_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimulationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimulationResponse) ProtoMessage() {}

// Deprecated: Use SimulationResponse.ProtoReflect.Descriptor instead.
func (*SimulationResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_base_abci_v1beta1_abci_proto_rawDescGZIP(), []int{6}
}

func (x *SimulationResponse) GetGasInfo() *GasInfo {
	if x != nil {
		return x.GasInfo
	}
	return nil
}

func (x *SimulationResponse) GetResult() *Result {
	if x != nil {
		return x.Result
	}
	return nil
}

// MsgData defines the data returned in a Result object during message
// execution.
//
// Deprecated: Do not use.
type MsgData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgType string `protobuf:"bytes,1,opt,name=msg_type,json=msgType,proto3" json:"msg_type,omitempty"`
	Data    []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *MsgData) Reset() {
	*x = MsgData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_abci_v1beta1_abci_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgData) ProtoMessage() {}

// Deprecated: Use MsgData.ProtoReflect.Descriptor instead.
func (*MsgData) Descriptor() ([]byte, []int) {
	return file_cosmos_base_abci_v1beta1_abci_proto_rawDescGZIP(), []int{7}
}

func (x *MsgData) GetMsgType() string {
	if x != nil {
		return x.MsgType
	}
	return ""
}

func (x *MsgData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// TxMsgData defines a list of MsgData. A transaction will have a MsgData object
// for each message.
type TxMsgData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// data field is deprecated and not populated.
	//
	// Deprecated: Do not use.
	Data []*MsgData `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	// msg_responses contains the Msg handler responses packed into Anys.
	//
	// Since: cosmos-sdk 0.45
	MsgResponses []*anypb.Any `protobuf:"bytes,2,rep,name=msg_responses,json=msgResponses,proto3" json:"msg_responses,omitempty"`
}

func (x *TxMsgData) Reset() {
	*x = TxMsgData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_abci_v1beta1_abci_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxMsgData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxMsgData) ProtoMessage() {}

// Deprecated: Use TxMsgData.ProtoReflect.Descriptor instead.
func (*TxMsgData) Descriptor() ([]byte, []int) {
	return file_cosmos_base_abci_v1beta1_abci_proto_rawDescGZIP(), []int{8}
}

// Deprecated: Do not use.
func (x *TxMsgData) GetData() []*MsgData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *TxMsgData) GetMsgResponses() []*anypb.Any {
	if x != nil {
		return x.MsgResponses
	}
	return nil
}

// SearchTxsResult defines a structure for querying txs pageable
type SearchTxsResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Count of all txs
	TotalCount uint64 `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	// Count of txs in current page
	Count uint64 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	// Index of current page, start from 1
	PageNumber uint64 `protobuf:"varint,3,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	// Count of total pages
	PageTotal uint64 `protobuf:"varint,4,opt,name=page_total,json=pageTotal,proto3" json:"page_total,omitempty"`
	// Max count txs per page
	Limit uint64 `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	// List of txs in current page
	Txs []*TxResponse `protobuf:"bytes,6,rep,name=txs,proto3" json:"txs,omitempty"`
}

func (x *SearchTxsResult) Reset() {
	*x = SearchTxsResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_base_abci_v1beta1_abci_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchTxsResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchTxsResult) ProtoMessage() {}

// Deprecated: Use SearchTxsResult.ProtoReflect.Descriptor instead.
func (*SearchTxsResult) Descriptor() ([]byte, []int) {
	return file_cosmos_base_abci_v1beta1_abci_proto_rawDescGZIP(), []int{9}
}

func (x *SearchTxsResult) GetTotalCount() uint64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *SearchTxsResult) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *SearchTxsResult) GetPageNumber() uint64 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *SearchTxsResult) GetPageTotal() uint64 {
	if x != nil {
		return x.PageTotal
	}
	return 0
}

func (x *SearchTxsResult) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *SearchTxsResult) GetTxs() []*TxResponse {
	if x != nil {
		return x.Txs
	}
	return nil
}

var File_cosmos_base_abci_v1beta1_abci_proto protoreflect.FileDescriptor

var file_cosmos_base_abci_v1beta1_abci_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x61, 0x62,
	0x63, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x62, 0x63, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x61, 0x62, 0x63, 0x69, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a,
	0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x74, 0x2f, 0x61, 0x62, 0x63, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x03,
	0x0a, 0x0a, 0x54, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x22, 0x0a, 0x06, 0x74, 0x78, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xe2, 0xde, 0x1f, 0x06, 0x54, 0x78, 0x48, 0x61, 0x73, 0x68,
	0x52, 0x06, 0x74, 0x78, 0x68, 0x61, 0x73, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x64, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x64,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x17,
	0x0a, 0x07, 0x72, 0x61, 0x77, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x61, 0x77, 0x4c, 0x6f, 0x67, 0x12, 0x55, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x61, 0x62, 0x63, 0x69, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x41, 0x42, 0x43, 0x49, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x42,
	0x17, 0xc8, 0xde, 0x1f, 0x00, 0xaa, 0xdf, 0x1f, 0x0f, 0x41, 0x42, 0x43, 0x49, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x61, 0x73, 0x5f, 0x77, 0x61, 0x6e, 0x74, 0x65, 0x64,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x67, 0x61, 0x73, 0x57, 0x61, 0x6e, 0x74, 0x65,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x61, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x61, 0x73, 0x55, 0x73, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x02,
	0x74, 0x78, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x02,
	0x74, 0x78, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x34, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x61, 0x62,
	0x63, 0x69, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x06,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x04, 0x88, 0xa0, 0x1f, 0x00, 0x22, 0x9a, 0x01, 0x0a,
	0x0e, 0x41, 0x42, 0x43, 0x49, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x12,
	0x1b, 0x0a, 0x09, 0x6d, 0x73, 0x67, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x6d, 0x73, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x10, 0x0a, 0x03,
	0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x12, 0x53,
	0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x61, 0x62, 0x63,
	0x69, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x14, 0xc8, 0xde, 0x1f, 0x00, 0xaa, 0xdf, 0x1f, 0x0c, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x06, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x3a, 0x04, 0x80, 0xdc, 0x20, 0x01, 0x22, 0x72, 0x0a, 0x0b, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x49, 0x0a, 0x0a,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x61,
	0x62, 0x63, 0x69, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x0a, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x3a, 0x04, 0x80, 0xdc, 0x20, 0x01, 0x22, 0x33, 0x0a,
	0x09, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x43, 0x0a, 0x07, 0x47, 0x61, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a,
	0x0a, 0x67, 0x61, 0x73, 0x5f, 0x77, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x09, 0x67, 0x61, 0x73, 0x57, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x67, 0x61, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x67, 0x61, 0x73, 0x55, 0x73, 0x65, 0x64, 0x22, 0xa9, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x16, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x42, 0x02, 0x18, 0x01, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x12, 0x34, 0x0a, 0x06,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x61, 0x62, 0x63, 0x69, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x39, 0x0a, 0x0d, 0x6d, 0x73, 0x67, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52,
	0x0c, 0x6d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x3a, 0x04, 0x88,
	0xa0, 0x1f, 0x00, 0x22, 0x96, 0x01, 0x0a, 0x12, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x08, 0x67, 0x61,
	0x73, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x61, 0x62, 0x63, 0x69, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x61, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x42,
	0x08, 0xc8, 0xde, 0x1f, 0x00, 0xd0, 0xde, 0x1f, 0x01, 0x52, 0x07, 0x67, 0x61, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x38, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x61, 0x62, 0x63, 0x69, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x40, 0x0a, 0x07,
	0x4d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x73, 0x67, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x3a, 0x06, 0x18, 0x01, 0x80, 0xdc, 0x20, 0x01, 0x22, 0x87,
	0x01, 0x0a, 0x09, 0x54, 0x78, 0x4d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x39, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x61, 0x62, 0x63, 0x69, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x42, 0x02, 0x18,
	0x01, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x39, 0x0a, 0x0d, 0x6d, 0x73, 0x67, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0c, 0x6d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x73, 0x3a, 0x04, 0x80, 0xdc, 0x20, 0x01, 0x22, 0xdc, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x54, 0x78, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x36, 0x0a, 0x03, 0x74, 0x78, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x61, 0x62, 0x63, 0x69, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x54, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x03, 0x74, 0x78,
	0x73, 0x3a, 0x04, 0x80, 0xdc, 0x20, 0x01, 0x42, 0xf7, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x61, 0x62, 0x63, 0x69,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x09, 0x41, 0x62, 0x63, 0x69, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x2f, 0x61, 0x62, 0x63, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x3b, 0x61, 0x62, 0x63, 0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43,
	0x42, 0x41, 0xaa, 0x02, 0x18, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x42, 0x61, 0x73, 0x65,
	0x2e, 0x41, 0x62, 0x63, 0x69, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x18,
	0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x42, 0x61, 0x73, 0x65, 0x5c, 0x41, 0x62, 0x63, 0x69,
	0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x24, 0x43, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x5c, 0x42, 0x61, 0x73, 0x65, 0x5c, 0x41, 0x62, 0x63, 0x69, 0x5c, 0x56, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x1b, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x42, 0x61, 0x73, 0x65, 0x3a, 0x3a,
	0x41, 0x62, 0x63, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xd8, 0xe1, 0x1e,
	0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_base_abci_v1beta1_abci_proto_rawDescOnce sync.Once
	file_cosmos_base_abci_v1beta1_abci_proto_rawDescData = file_cosmos_base_abci_v1beta1_abci_proto_rawDesc
)

func file_cosmos_base_abci_v1beta1_abci_proto_rawDescGZIP() []byte {
	file_cosmos_base_abci_v1beta1_abci_proto_rawDescOnce.Do(func() {
		file_cosmos_base_abci_v1beta1_abci_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_base_abci_v1beta1_abci_proto_rawDescData)
	})
	return file_cosmos_base_abci_v1beta1_abci_proto_rawDescData
}

var file_cosmos_base_abci_v1beta1_abci_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_cosmos_base_abci_v1beta1_abci_proto_goTypes = []interface{}{
	(*TxResponse)(nil),         // 0: cosmos.base.abci.v1beta1.TxResponse
	(*ABCIMessageLog)(nil),     // 1: cosmos.base.abci.v1beta1.ABCIMessageLog
	(*StringEvent)(nil),        // 2: cosmos.base.abci.v1beta1.StringEvent
	(*Attribute)(nil),          // 3: cosmos.base.abci.v1beta1.Attribute
	(*GasInfo)(nil),            // 4: cosmos.base.abci.v1beta1.GasInfo
	(*Result)(nil),             // 5: cosmos.base.abci.v1beta1.Result
	(*SimulationResponse)(nil), // 6: cosmos.base.abci.v1beta1.SimulationResponse
	(*MsgData)(nil),            // 7: cosmos.base.abci.v1beta1.MsgData
	(*TxMsgData)(nil),          // 8: cosmos.base.abci.v1beta1.TxMsgData
	(*SearchTxsResult)(nil),    // 9: cosmos.base.abci.v1beta1.SearchTxsResult
	(*anypb.Any)(nil),          // 10: google.protobuf.Any
	(*abci.Event)(nil),         // 11: tendermint.abci.Event
}
var file_cosmos_base_abci_v1beta1_abci_proto_depIdxs = []int32{
	1,  // 0: cosmos.base.abci.v1beta1.TxResponse.logs:type_name -> cosmos.base.abci.v1beta1.ABCIMessageLog
	10, // 1: cosmos.base.abci.v1beta1.TxResponse.tx:type_name -> google.protobuf.Any
	11, // 2: cosmos.base.abci.v1beta1.TxResponse.events:type_name -> tendermint.abci.Event
	2,  // 3: cosmos.base.abci.v1beta1.ABCIMessageLog.events:type_name -> cosmos.base.abci.v1beta1.StringEvent
	3,  // 4: cosmos.base.abci.v1beta1.StringEvent.attributes:type_name -> cosmos.base.abci.v1beta1.Attribute
	11, // 5: cosmos.base.abci.v1beta1.Result.events:type_name -> tendermint.abci.Event
	10, // 6: cosmos.base.abci.v1beta1.Result.msg_responses:type_name -> google.protobuf.Any
	4,  // 7: cosmos.base.abci.v1beta1.SimulationResponse.gas_info:type_name -> cosmos.base.abci.v1beta1.GasInfo
	5,  // 8: cosmos.base.abci.v1beta1.SimulationResponse.result:type_name -> cosmos.base.abci.v1beta1.Result
	7,  // 9: cosmos.base.abci.v1beta1.TxMsgData.data:type_name -> cosmos.base.abci.v1beta1.MsgData
	10, // 10: cosmos.base.abci.v1beta1.TxMsgData.msg_responses:type_name -> google.protobuf.Any
	0,  // 11: cosmos.base.abci.v1beta1.SearchTxsResult.txs:type_name -> cosmos.base.abci.v1beta1.TxResponse
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_cosmos_base_abci_v1beta1_abci_proto_init() }
func file_cosmos_base_abci_v1beta1_abci_proto_init() {
	if File_cosmos_base_abci_v1beta1_abci_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_base_abci_v1beta1_abci_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxResponse); i {
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
		file_cosmos_base_abci_v1beta1_abci_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ABCIMessageLog); i {
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
		file_cosmos_base_abci_v1beta1_abci_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringEvent); i {
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
		file_cosmos_base_abci_v1beta1_abci_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attribute); i {
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
		file_cosmos_base_abci_v1beta1_abci_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GasInfo); i {
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
		file_cosmos_base_abci_v1beta1_abci_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_cosmos_base_abci_v1beta1_abci_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimulationResponse); i {
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
		file_cosmos_base_abci_v1beta1_abci_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgData); i {
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
		file_cosmos_base_abci_v1beta1_abci_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxMsgData); i {
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
		file_cosmos_base_abci_v1beta1_abci_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchTxsResult); i {
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
			RawDescriptor: file_cosmos_base_abci_v1beta1_abci_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_base_abci_v1beta1_abci_proto_goTypes,
		DependencyIndexes: file_cosmos_base_abci_v1beta1_abci_proto_depIdxs,
		MessageInfos:      file_cosmos_base_abci_v1beta1_abci_proto_msgTypes,
	}.Build()
	File_cosmos_base_abci_v1beta1_abci_proto = out.File
	file_cosmos_base_abci_v1beta1_abci_proto_rawDesc = nil
	file_cosmos_base_abci_v1beta1_abci_proto_goTypes = nil
	file_cosmos_base_abci_v1beta1_abci_proto_depIdxs = nil
}
