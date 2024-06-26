// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0
// source: api/shop/v1/grocery.proto

package v1

import (
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

// ----------Request----------
type CreateGroceryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                   // 如果id=0，代表需要新建一个商品
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`               // 商品名称
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"` // 商品描述
	Price       int32  `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`            // 商品价格
	Image       string `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`             // 商品图片
	Category    string `protobuf:"bytes,6,opt,name=category,proto3" json:"category,omitempty"`       // 商品分类
}

func (x *CreateGroceryRequest) Reset() {
	*x = CreateGroceryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shop_v1_grocery_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroceryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroceryRequest) ProtoMessage() {}

func (x *CreateGroceryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_v1_grocery_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroceryRequest.ProtoReflect.Descriptor instead.
func (*CreateGroceryRequest) Descriptor() ([]byte, []int) {
	return file_api_shop_v1_grocery_proto_rawDescGZIP(), []int{0}
}

func (x *CreateGroceryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateGroceryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateGroceryRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateGroceryRequest) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateGroceryRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *CreateGroceryRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

type UpdateGroceryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`               // 商品名称
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"` // 商品描述
	Price       int32  `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`            // 商品价格
	Image       string `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`             // 商品图片
	Category    string `protobuf:"bytes,6,opt,name=category,proto3" json:"category,omitempty"`       // 商品分类
}

func (x *UpdateGroceryRequest) Reset() {
	*x = UpdateGroceryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shop_v1_grocery_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGroceryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGroceryRequest) ProtoMessage() {}

func (x *UpdateGroceryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_v1_grocery_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGroceryRequest.ProtoReflect.Descriptor instead.
func (*UpdateGroceryRequest) Descriptor() ([]byte, []int) {
	return file_api_shop_v1_grocery_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateGroceryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateGroceryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateGroceryRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateGroceryRequest) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *UpdateGroceryRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *UpdateGroceryRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

type DeleteGroceryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteGroceryRequest) Reset() {
	*x = DeleteGroceryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shop_v1_grocery_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGroceryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGroceryRequest) ProtoMessage() {}

func (x *DeleteGroceryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_v1_grocery_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGroceryRequest.ProtoReflect.Descriptor instead.
func (*DeleteGroceryRequest) Descriptor() ([]byte, []int) {
	return file_api_shop_v1_grocery_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteGroceryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetGroceryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetGroceryRequest) Reset() {
	*x = GetGroceryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shop_v1_grocery_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroceryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroceryRequest) ProtoMessage() {}

func (x *GetGroceryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_v1_grocery_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroceryRequest.ProtoReflect.Descriptor instead.
func (*GetGroceryRequest) Descriptor() ([]byte, []int) {
	return file_api_shop_v1_grocery_proto_rawDescGZIP(), []int{3}
}

func (x *GetGroceryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListGroceryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []string `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
}

func (x *ListGroceryRequest) Reset() {
	*x = ListGroceryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shop_v1_grocery_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGroceryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGroceryRequest) ProtoMessage() {}

func (x *ListGroceryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_v1_grocery_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGroceryRequest.ProtoReflect.Descriptor instead.
func (*ListGroceryRequest) Descriptor() ([]byte, []int) {
	return file_api_shop_v1_grocery_proto_rawDescGZIP(), []int{4}
}

func (x *ListGroceryRequest) GetId() []string {
	if x != nil {
		return x.Id
	}
	return nil
}

// ----------Reply----------
type CreateGroceryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg       string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Timestamp int64  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Id        string `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateGroceryReply) Reset() {
	*x = CreateGroceryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shop_v1_grocery_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroceryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroceryReply) ProtoMessage() {}

func (x *CreateGroceryReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_v1_grocery_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroceryReply.ProtoReflect.Descriptor instead.
func (*CreateGroceryReply) Descriptor() ([]byte, []int) {
	return file_api_shop_v1_grocery_proto_rawDescGZIP(), []int{5}
}

func (x *CreateGroceryReply) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CreateGroceryReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *CreateGroceryReply) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *CreateGroceryReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateGroceryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg       string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Timestamp int64  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Id        string `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateGroceryReply) Reset() {
	*x = UpdateGroceryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shop_v1_grocery_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGroceryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGroceryReply) ProtoMessage() {}

func (x *UpdateGroceryReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_v1_grocery_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGroceryReply.ProtoReflect.Descriptor instead.
func (*UpdateGroceryReply) Descriptor() ([]byte, []int) {
	return file_api_shop_v1_grocery_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateGroceryReply) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *UpdateGroceryReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *UpdateGroceryReply) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *UpdateGroceryReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteGroceryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg       string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Timestamp int64  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Id        string `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteGroceryReply) Reset() {
	*x = DeleteGroceryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shop_v1_grocery_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGroceryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGroceryReply) ProtoMessage() {}

func (x *DeleteGroceryReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_v1_grocery_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGroceryReply.ProtoReflect.Descriptor instead.
func (*DeleteGroceryReply) Descriptor() ([]byte, []int) {
	return file_api_shop_v1_grocery_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteGroceryReply) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *DeleteGroceryReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *DeleteGroceryReply) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *DeleteGroceryReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetGroceryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    string         `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg       string         `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Timestamp int64          `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Grocery   *GroceryEntity `protobuf:"bytes,4,opt,name=grocery,proto3" json:"grocery,omitempty"`
}

func (x *GetGroceryReply) Reset() {
	*x = GetGroceryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shop_v1_grocery_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroceryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroceryReply) ProtoMessage() {}

func (x *GetGroceryReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_v1_grocery_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroceryReply.ProtoReflect.Descriptor instead.
func (*GetGroceryReply) Descriptor() ([]byte, []int) {
	return file_api_shop_v1_grocery_proto_rawDescGZIP(), []int{8}
}

func (x *GetGroceryReply) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetGroceryReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetGroceryReply) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *GetGroceryReply) GetGrocery() *GroceryEntity {
	if x != nil {
		return x.Grocery
	}
	return nil
}

type ListGroceryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    string         `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg       string         `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Timestamp int64          `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Grocery   *GroceryEntity `protobuf:"bytes,4,opt,name=grocery,proto3" json:"grocery,omitempty"`
}

func (x *ListGroceryReply) Reset() {
	*x = ListGroceryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shop_v1_grocery_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGroceryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGroceryReply) ProtoMessage() {}

func (x *ListGroceryReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_v1_grocery_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGroceryReply.ProtoReflect.Descriptor instead.
func (*ListGroceryReply) Descriptor() ([]byte, []int) {
	return file_api_shop_v1_grocery_proto_rawDescGZIP(), []int{9}
}

func (x *ListGroceryReply) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ListGroceryReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ListGroceryReply) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *ListGroceryReply) GetGrocery() *GroceryEntity {
	if x != nil {
		return x.Grocery
	}
	return nil
}

// ----------商品字段----------
type GroceryEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`               // 商品名称
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"` // 商品描述
	Price       int32  `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`            // 商品价格
	Image       string `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`             // 商品图片
	Category    string `protobuf:"bytes,6,opt,name=category,proto3" json:"category,omitempty"`       // 商品分类
}

func (x *GroceryEntity) Reset() {
	*x = GroceryEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shop_v1_grocery_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroceryEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroceryEntity) ProtoMessage() {}

func (x *GroceryEntity) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_v1_grocery_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroceryEntity.ProtoReflect.Descriptor instead.
func (*GroceryEntity) Descriptor() ([]byte, []int) {
	return file_api_shop_v1_grocery_proto_rawDescGZIP(), []int{10}
}

func (x *GroceryEntity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GroceryEntity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GroceryEntity) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GroceryEntity) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *GroceryEntity) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *GroceryEntity) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

var File_api_shop_v1_grocery_proto protoreflect.FileDescriptor

var file_api_shop_v1_grocery_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72,
	0x6f, 0x63, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x22, 0xa4, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22,
	0xa4, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x23,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x63, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6c, 0x0a, 0x12, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6c, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6c, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47,
	0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x8f, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x63, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x34, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x07, 0x67, 0x72,
	0x6f, 0x63, 0x65, 0x72, 0x79, 0x22, 0x90, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72,
	0x6f, 0x63, 0x65, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x34, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52,
	0x07, 0x67, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x22, 0x9d, 0x01, 0x0a, 0x0d, 0x47, 0x72, 0x6f,
	0x63, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x32, 0xa3, 0x03, 0x0a, 0x07, 0x47, 0x72, 0x6f,
	0x63, 0x65, 0x72, 0x79, 0x12, 0x53, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72,
	0x6f, 0x63, 0x65, 0x72, 0x79, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x70,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f,
	0x63, 0x65, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x53, 0x0a, 0x0d, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47,
	0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x53,
	0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x12,
	0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x4a, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72,
	0x79, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x4d, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x12, 0x1f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x47, 0x72, 0x6f, 0x63, 0x65, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x2b,
	0x0a, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a,
	0x1a, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_shop_v1_grocery_proto_rawDescOnce sync.Once
	file_api_shop_v1_grocery_proto_rawDescData = file_api_shop_v1_grocery_proto_rawDesc
)

func file_api_shop_v1_grocery_proto_rawDescGZIP() []byte {
	file_api_shop_v1_grocery_proto_rawDescOnce.Do(func() {
		file_api_shop_v1_grocery_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_shop_v1_grocery_proto_rawDescData)
	})
	return file_api_shop_v1_grocery_proto_rawDescData
}

var file_api_shop_v1_grocery_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_shop_v1_grocery_proto_goTypes = []interface{}{
	(*CreateGroceryRequest)(nil), // 0: api.shop.v1.CreateGroceryRequest
	(*UpdateGroceryRequest)(nil), // 1: api.shop.v1.UpdateGroceryRequest
	(*DeleteGroceryRequest)(nil), // 2: api.shop.v1.DeleteGroceryRequest
	(*GetGroceryRequest)(nil),    // 3: api.shop.v1.GetGroceryRequest
	(*ListGroceryRequest)(nil),   // 4: api.shop.v1.ListGroceryRequest
	(*CreateGroceryReply)(nil),   // 5: api.shop.v1.CreateGroceryReply
	(*UpdateGroceryReply)(nil),   // 6: api.shop.v1.UpdateGroceryReply
	(*DeleteGroceryReply)(nil),   // 7: api.shop.v1.DeleteGroceryReply
	(*GetGroceryReply)(nil),      // 8: api.shop.v1.GetGroceryReply
	(*ListGroceryReply)(nil),     // 9: api.shop.v1.ListGroceryReply
	(*GroceryEntity)(nil),        // 10: api.shop.v1.GroceryEntity
}
var file_api_shop_v1_grocery_proto_depIdxs = []int32{
	10, // 0: api.shop.v1.GetGroceryReply.grocery:type_name -> api.shop.v1.GroceryEntity
	10, // 1: api.shop.v1.ListGroceryReply.grocery:type_name -> api.shop.v1.GroceryEntity
	0,  // 2: api.shop.v1.Grocery.CreateGrocery:input_type -> api.shop.v1.CreateGroceryRequest
	1,  // 3: api.shop.v1.Grocery.UpdateGrocery:input_type -> api.shop.v1.UpdateGroceryRequest
	2,  // 4: api.shop.v1.Grocery.DeleteGrocery:input_type -> api.shop.v1.DeleteGroceryRequest
	3,  // 5: api.shop.v1.Grocery.GetGrocery:input_type -> api.shop.v1.GetGroceryRequest
	4,  // 6: api.shop.v1.Grocery.ListGrocery:input_type -> api.shop.v1.ListGroceryRequest
	5,  // 7: api.shop.v1.Grocery.CreateGrocery:output_type -> api.shop.v1.CreateGroceryReply
	6,  // 8: api.shop.v1.Grocery.UpdateGrocery:output_type -> api.shop.v1.UpdateGroceryReply
	7,  // 9: api.shop.v1.Grocery.DeleteGrocery:output_type -> api.shop.v1.DeleteGroceryReply
	8,  // 10: api.shop.v1.Grocery.GetGrocery:output_type -> api.shop.v1.GetGroceryReply
	9,  // 11: api.shop.v1.Grocery.ListGrocery:output_type -> api.shop.v1.ListGroceryReply
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_api_shop_v1_grocery_proto_init() }
func file_api_shop_v1_grocery_proto_init() {
	if File_api_shop_v1_grocery_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_shop_v1_grocery_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroceryRequest); i {
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
		file_api_shop_v1_grocery_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGroceryRequest); i {
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
		file_api_shop_v1_grocery_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGroceryRequest); i {
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
		file_api_shop_v1_grocery_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroceryRequest); i {
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
		file_api_shop_v1_grocery_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGroceryRequest); i {
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
		file_api_shop_v1_grocery_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroceryReply); i {
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
		file_api_shop_v1_grocery_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGroceryReply); i {
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
		file_api_shop_v1_grocery_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGroceryReply); i {
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
		file_api_shop_v1_grocery_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroceryReply); i {
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
		file_api_shop_v1_grocery_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGroceryReply); i {
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
		file_api_shop_v1_grocery_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroceryEntity); i {
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
			RawDescriptor: file_api_shop_v1_grocery_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_shop_v1_grocery_proto_goTypes,
		DependencyIndexes: file_api_shop_v1_grocery_proto_depIdxs,
		MessageInfos:      file_api_shop_v1_grocery_proto_msgTypes,
	}.Build()
	File_api_shop_v1_grocery_proto = out.File
	file_api_shop_v1_grocery_proto_rawDesc = nil
	file_api_shop_v1_grocery_proto_goTypes = nil
	file_api_shop_v1_grocery_proto_depIdxs = nil
}
