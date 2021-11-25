// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.13.0
// source: entities_http.proto

package entities_http

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit     int32          `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset    int32          `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	DateStart string         `protobuf:"bytes,3,opt,name=dateStart,proto3" json:"dateStart,omitempty"`
	DateEnd   string         `protobuf:"bytes,4,opt,name=dateEnd,proto3" json:"dateEnd,omitempty"`
	Params    []*FilterParam `protobuf:"bytes,5,rep,name=params,proto3" json:"params,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_http_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_entities_http_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_entities_http_proto_rawDescGZIP(), []int{0}
}

func (x *Filter) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Filter) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *Filter) GetDateStart() string {
	if x != nil {
		return x.DateStart
	}
	return ""
}

func (x *Filter) GetDateEnd() string {
	if x != nil {
		return x.DateEnd
	}
	return ""
}

func (x *Filter) GetParams() []*FilterParam {
	if x != nil {
		return x.Params
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//Количество записей
	Count int32          `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Items []*empty.Empty `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_http_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_entities_http_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_entities_http_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Response) GetItems() []*empty.Empty {
	if x != nil {
		return x.Items
	}
	return nil
}

type FilterParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value []string `protobuf:"bytes,2,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *FilterParam) Reset() {
	*x = FilterParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_http_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterParam) ProtoMessage() {}

func (x *FilterParam) ProtoReflect() protoreflect.Message {
	mi := &file_entities_http_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterParam.ProtoReflect.Descriptor instead.
func (*FilterParam) Descriptor() ([]byte, []int) {
	return file_entities_http_proto_rawDescGZIP(), []int{2}
}

func (x *FilterParam) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FilterParam) GetValue() []string {
	if x != nil {
		return x.Value
	}
	return nil
}

type InfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *InfoResponse) Reset() {
	*x = InfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_http_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoResponse) ProtoMessage() {}

func (x *InfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_entities_http_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoResponse.ProtoReflect.Descriptor instead.
func (*InfoResponse) Descriptor() ([]byte, []int) {
	return file_entities_http_proto_rawDescGZIP(), []int{3}
}

func (x *InfoResponse) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	CreatedDate string `protobuf:"bytes,3,opt,name=createdDate,proto3" json:"createdDate,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_http_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_entities_http_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_entities_http_proto_rawDescGZIP(), []int{4}
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Item) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Item) GetCreatedDate() string {
	if x != nil {
		return x.CreatedDate
	}
	return ""
}

type MetaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Phone string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *MetaRequest) Reset() {
	*x = MetaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_http_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaRequest) ProtoMessage() {}

func (x *MetaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entities_http_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaRequest.ProtoReflect.Descriptor instead.
func (*MetaRequest) Descriptor() ([]byte, []int) {
	return file_entities_http_proto_rawDescGZIP(), []int{5}
}

func (x *MetaRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MetaRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *MetaRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type MetaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta []*Meta `protobuf:"bytes,1,rep,name=meta,proto3" json:"meta,omitempty"`
}

func (x *MetaResponse) Reset() {
	*x = MetaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_http_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaResponse) ProtoMessage() {}

func (x *MetaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_entities_http_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaResponse.ProtoReflect.Descriptor instead.
func (*MetaResponse) Descriptor() ([]byte, []int) {
	return file_entities_http_proto_rawDescGZIP(), []int{6}
}

func (x *MetaResponse) GetMeta() []*Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Count int32  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_http_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_entities_http_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_entities_http_proto_rawDescGZIP(), []int{7}
}

func (x *Meta) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Meta) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type EventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//Количество записей
	Count int32          `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Items []*empty.Empty `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *EventsResponse) Reset() {
	*x = EventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_http_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventsResponse) ProtoMessage() {}

func (x *EventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_entities_http_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventsResponse.ProtoReflect.Descriptor instead.
func (*EventsResponse) Descriptor() ([]byte, []int) {
	return file_entities_http_proto_rawDescGZIP(), []int{8}
}

func (x *EventsResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *EventsResponse) GetItems() []*empty.Empty {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_entities_http_proto protoreflect.FileDescriptor

var file_entities_http_proto_rawDesc = []byte{
	0x0a, 0x13, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5f,
	0x68, 0x74, 0x74, 0x70, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x62, 0x6f,
	0x64, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xae, 0x01, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0xe0, 0x41, 0x02, 0xfa, 0x42,
	0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x12, 0x32, 0x0a, 0x06,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x22, 0x73, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x3a, 0x23, 0x92, 0x41, 0x20, 0x32, 0x1e, 0x7b, 0x22, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3a,
	0x20, 0x31, 0x32, 0x2c, 0x20, 0x22, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x3a, 0x5b, 0x7b, 0x7d,
	0x2c, 0x7b, 0x7d, 0x5d, 0x7d, 0x22, 0x37, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x39,
	0x0a, 0x0c, 0x69, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x69, 0x74,
	0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x5e, 0x0a, 0x04, 0x69, 0x74, 0x65,
	0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0x82, 0x01, 0x0a, 0x0b, 0x6d, 0x65,
	0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xe0, 0x41, 0x02, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xe0, 0x41, 0x02, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x60, 0x01, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x31, 0x0a, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0xe0, 0x41, 0x02, 0xfa,
	0x42, 0x15, 0x72, 0x13, 0x28, 0x80, 0x02, 0x32, 0x0e, 0x5e, 0x5c, 0x2b, 0x3f, 0x37, 0x28, 0x39,
	0x5c, 0x64, 0x7b, 0x39, 0x7d, 0x29, 0x24, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x37,
	0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27,
	0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x30, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x79, 0x0a, 0x0e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x2c, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x3a,
	0x23, 0x92, 0x41, 0x20, 0x32, 0x1e, 0x7b, 0x22, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3a, 0x20,
	0x31, 0x32, 0x2c, 0x20, 0x22, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x3a, 0x5b, 0x7b, 0x7d, 0x2c,
	0x7b, 0x7d, 0x5d, 0x7d, 0x32, 0xc5, 0x02, 0x0a, 0x08, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x12, 0x99, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x12, 0x15, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x68, 0x74,
	0x74, 0x70, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x22,
	0x5c, 0x92, 0x41, 0x49, 0x4a, 0x47, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x40, 0x0a, 0x1b, 0xd0,
	0xa3, 0xd1, 0x81, 0xd0, 0xbf, 0xd0, 0xb5, 0xd1, 0x88, 0xd0, 0xbd, 0xd1, 0x8b, 0xd0, 0xb9, 0x20,
	0xd0, 0xbe, 0xd1, 0x82, 0xd0, 0xb2, 0xd0, 0xb5, 0xd1, 0x82, 0x12, 0x21, 0x0a, 0x1f, 0x1a, 0x1d,
	0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0a, 0x22, 0x05, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x4a, 0x0a,
	0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x69, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0d, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x07, 0x12, 0x05, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x51, 0x0a, 0x04, 0x4d, 0x65, 0x74,
	0x61, 0x12, 0x1a, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6d, 0x65,
	0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0a, 0x22, 0x05, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x3a, 0x01, 0x2a, 0x42, 0x72, 0x5a, 0x10,
	0x2e, 0x2f, 0x3b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x68, 0x74, 0x74, 0x70,
	0x92, 0x41, 0x5d, 0x12, 0x33, 0x0a, 0x2a, 0xd0, 0xa1, 0xd0, 0xb5, 0xd1, 0x80, 0xd0, 0xb2, 0xd0,
	0xb8, 0xd1, 0x81, 0x20, 0xd0, 0xbd, 0xd0, 0xb5, 0xd0, 0xba, 0xd0, 0xb8, 0xd1, 0x85, 0x20, 0xd1,
	0x81, 0xd1, 0x83, 0xd1, 0x89, 0xd0, 0xbd, 0xd0, 0xbe, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xb5, 0xd0,
	0xb9, 0x32, 0x05, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entities_http_proto_rawDescOnce sync.Once
	file_entities_http_proto_rawDescData = file_entities_http_proto_rawDesc
)

func file_entities_http_proto_rawDescGZIP() []byte {
	file_entities_http_proto_rawDescOnce.Do(func() {
		file_entities_http_proto_rawDescData = protoimpl.X.CompressGZIP(file_entities_http_proto_rawDescData)
	})
	return file_entities_http_proto_rawDescData
}

var file_entities_http_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_entities_http_proto_goTypes = []interface{}{
	(*Filter)(nil),            // 0: entities_http.filter
	(*Response)(nil),          // 1: entities_http.response
	(*FilterParam)(nil),       // 2: entities_http.filterParam
	(*InfoResponse)(nil),      // 3: entities_http.infoResponse
	(*Item)(nil),              // 4: entities_http.item
	(*MetaRequest)(nil),       // 5: entities_http.metaRequest
	(*MetaResponse)(nil),      // 6: entities_http.metaResponse
	(*Meta)(nil),              // 7: entities_http.meta
	(*EventsResponse)(nil),    // 8: entities_http.eventsResponse
	(*empty.Empty)(nil),       // 9: google.protobuf.Empty
	(*httpbody.HttpBody)(nil), // 10: google.api.HttpBody
}
var file_entities_http_proto_depIdxs = []int32{
	2,  // 0: entities_http.filter.params:type_name -> entities_http.filterParam
	9,  // 1: entities_http.response.items:type_name -> google.protobuf.Empty
	4,  // 2: entities_http.infoResponse.items:type_name -> entities_http.item
	7,  // 3: entities_http.metaResponse.meta:type_name -> entities_http.meta
	9,  // 4: entities_http.eventsResponse.items:type_name -> google.protobuf.Empty
	0,  // 5: entities_http.Entities.ListByFilter:input_type -> entities_http.filter
	9,  // 6: entities_http.Entities.Info:input_type -> google.protobuf.Empty
	5,  // 7: entities_http.Entities.Meta:input_type -> entities_http.metaRequest
	10, // 8: entities_http.Entities.ListByFilter:output_type -> google.api.HttpBody
	3,  // 9: entities_http.Entities.Info:output_type -> entities_http.infoResponse
	6,  // 10: entities_http.Entities.Meta:output_type -> entities_http.metaResponse
	8,  // [8:11] is the sub-list for method output_type
	5,  // [5:8] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_entities_http_proto_init() }
func file_entities_http_proto_init() {
	if File_entities_http_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_entities_http_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
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
		file_entities_http_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_entities_http_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterParam); i {
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
		file_entities_http_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoResponse); i {
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
		file_entities_http_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_entities_http_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetaRequest); i {
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
		file_entities_http_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetaResponse); i {
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
		file_entities_http_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta); i {
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
		file_entities_http_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventsResponse); i {
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
			RawDescriptor: file_entities_http_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_entities_http_proto_goTypes,
		DependencyIndexes: file_entities_http_proto_depIdxs,
		MessageInfos:      file_entities_http_proto_msgTypes,
	}.Build()
	File_entities_http_proto = out.File
	file_entities_http_proto_rawDesc = nil
	file_entities_http_proto_goTypes = nil
	file_entities_http_proto_depIdxs = nil
}
