// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: wishlist.proto

package pb

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

type WishlistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *WishlistRequest) Reset() {
	*x = WishlistRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wishlist_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WishlistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WishlistRequest) ProtoMessage() {}

func (x *WishlistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wishlist_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WishlistRequest.ProtoReflect.Descriptor instead.
func (*WishlistRequest) Descriptor() ([]byte, []int) {
	return file_wishlist_proto_rawDescGZIP(), []int{0}
}

func (x *WishlistRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type WishlistResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WishlistId uint32 `protobuf:"varint,1,opt,name=wishlist_id,json=wishlistId,proto3" json:"wishlist_id,omitempty"`
	UserId     uint32 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	IsEmpty    bool   `protobuf:"varint,3,opt,name=is_empty,json=isEmpty,proto3" json:"is_empty,omitempty"`
}

func (x *WishlistResponce) Reset() {
	*x = WishlistResponce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wishlist_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WishlistResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WishlistResponce) ProtoMessage() {}

func (x *WishlistResponce) ProtoReflect() protoreflect.Message {
	mi := &file_wishlist_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WishlistResponce.ProtoReflect.Descriptor instead.
func (*WishlistResponce) Descriptor() ([]byte, []int) {
	return file_wishlist_proto_rawDescGZIP(), []int{1}
}

func (x *WishlistResponce) GetWishlistId() uint32 {
	if x != nil {
		return x.WishlistId
	}
	return 0
}

func (x *WishlistResponce) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *WishlistResponce) GetIsEmpty() bool {
	if x != nil {
		return x.IsEmpty
	}
	return false
}

type AddtoWishlistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProductId uint32 `protobuf:"varint,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *AddtoWishlistRequest) Reset() {
	*x = AddtoWishlistRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wishlist_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddtoWishlistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddtoWishlistRequest) ProtoMessage() {}

func (x *AddtoWishlistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wishlist_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddtoWishlistRequest.ProtoReflect.Descriptor instead.
func (*AddtoWishlistRequest) Descriptor() ([]byte, []int) {
	return file_wishlist_proto_rawDescGZIP(), []int{2}
}

func (x *AddtoWishlistRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AddtoWishlistRequest) GetProductId() uint32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

var File_wishlist_proto protoreflect.FileDescriptor

var file_wishlist_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x77, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x77, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x1a, 0x0d, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x0f, 0x57, 0x69, 0x73,
	0x68, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x67, 0x0a, 0x10, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x69, 0x73,
	0x68, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x77, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x4e,
	0x0a, 0x14, 0x41, 0x64, 0x64, 0x74, 0x6f, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x32, 0x93,
	0x03, 0x0a, 0x0f, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x47, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x69, 0x73, 0x68,
	0x6c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x77, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e,
	0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x77, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x57, 0x69, 0x73, 0x68, 0x6c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x77, 0x69, 0x73,
	0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x63, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4c, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x74, 0x6f, 0x57,
	0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x77, 0x69, 0x73, 0x68, 0x6c, 0x69,
	0x73, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x74, 0x6f, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x69,
	0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1e, 0x2e, 0x77, 0x69, 0x73,
	0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x74, 0x6f, 0x57, 0x69, 0x73, 0x68, 0x6c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x30, 0x01, 0x12, 0x49, 0x0a, 0x10, 0x54, 0x72, 0x75,
	0x6e, 0x63, 0x61, 0x74, 0x65, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x2e,
	0x77, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x77, 0x69, 0x73, 0x68, 0x6c,
	0x69, 0x73, 0x74, 0x2e, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x63, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wishlist_proto_rawDescOnce sync.Once
	file_wishlist_proto_rawDescData = file_wishlist_proto_rawDesc
)

func file_wishlist_proto_rawDescGZIP() []byte {
	file_wishlist_proto_rawDescOnce.Do(func() {
		file_wishlist_proto_rawDescData = protoimpl.X.CompressGZIP(file_wishlist_proto_rawDescData)
	})
	return file_wishlist_proto_rawDescData
}

var file_wishlist_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_wishlist_proto_goTypes = []interface{}{
	(*WishlistRequest)(nil),      // 0: wishlist.WishlistRequest
	(*WishlistResponce)(nil),     // 1: wishlist.WishlistResponce
	(*AddtoWishlistRequest)(nil), // 2: wishlist.AddtoWishlistRequest
	(*AddProductResponce)(nil),   // 3: product.AddProductResponce
}
var file_wishlist_proto_depIdxs = []int32{
	0, // 0: wishlist.WishlistService.CreateWishlist:input_type -> wishlist.WishlistRequest
	0, // 1: wishlist.WishlistService.GetWishlist:input_type -> wishlist.WishlistRequest
	2, // 2: wishlist.WishlistService.AddtoWishlist:input_type -> wishlist.AddtoWishlistRequest
	2, // 3: wishlist.WishlistService.DeleteWishlistItem:input_type -> wishlist.AddtoWishlistRequest
	0, // 4: wishlist.WishlistService.TruncateWishlist:input_type -> wishlist.WishlistRequest
	1, // 5: wishlist.WishlistService.CreateWishlist:output_type -> wishlist.WishlistResponce
	3, // 6: wishlist.WishlistService.GetWishlist:output_type -> product.AddProductResponce
	3, // 7: wishlist.WishlistService.AddtoWishlist:output_type -> product.AddProductResponce
	3, // 8: wishlist.WishlistService.DeleteWishlistItem:output_type -> product.AddProductResponce
	1, // 9: wishlist.WishlistService.TruncateWishlist:output_type -> wishlist.WishlistResponce
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wishlist_proto_init() }
func file_wishlist_proto_init() {
	if File_wishlist_proto != nil {
		return
	}
	file_product_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wishlist_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WishlistRequest); i {
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
		file_wishlist_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WishlistResponce); i {
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
		file_wishlist_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddtoWishlistRequest); i {
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
			RawDescriptor: file_wishlist_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wishlist_proto_goTypes,
		DependencyIndexes: file_wishlist_proto_depIdxs,
		MessageInfos:      file_wishlist_proto_msgTypes,
	}.Build()
	File_wishlist_proto = out.File
	file_wishlist_proto_rawDesc = nil
	file_wishlist_proto_goTypes = nil
	file_wishlist_proto_depIdxs = nil
}
