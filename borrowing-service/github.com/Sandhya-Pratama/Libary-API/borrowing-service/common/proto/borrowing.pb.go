// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: common/proto/borrowing.proto

package borrowing

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

type BorrowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	BookId int32 `protobuf:"varint,2,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
}

func (x *BorrowRequest) Reset() {
	*x = BorrowRequest{}
	mi := &file_common_proto_borrowing_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BorrowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BorrowRequest) ProtoMessage() {}

func (x *BorrowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_borrowing_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BorrowRequest.ProtoReflect.Descriptor instead.
func (*BorrowRequest) Descriptor() ([]byte, []int) {
	return file_common_proto_borrowing_proto_rawDescGZIP(), []int{0}
}

func (x *BorrowRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *BorrowRequest) GetBookId() int32 {
	if x != nil {
		return x.BookId
	}
	return 0
}

type ReturnRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BorrowingId int32 `protobuf:"varint,1,opt,name=borrowing_id,json=borrowingId,proto3" json:"borrowing_id,omitempty"`
}

func (x *ReturnRequest) Reset() {
	*x = ReturnRequest{}
	mi := &file_common_proto_borrowing_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReturnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReturnRequest) ProtoMessage() {}

func (x *ReturnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_borrowing_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReturnRequest.ProtoReflect.Descriptor instead.
func (*ReturnRequest) Descriptor() ([]byte, []int) {
	return file_common_proto_borrowing_proto_rawDescGZIP(), []int{1}
}

func (x *ReturnRequest) GetBorrowingId() int32 {
	if x != nil {
		return x.BorrowingId
	}
	return 0
}

type BorrowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BorrowingId int32  `protobuf:"varint,1,opt,name=borrowing_id,json=borrowingId,proto3" json:"borrowing_id,omitempty"`
	Message     string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *BorrowResponse) Reset() {
	*x = BorrowResponse{}
	mi := &file_common_proto_borrowing_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BorrowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BorrowResponse) ProtoMessage() {}

func (x *BorrowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_borrowing_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BorrowResponse.ProtoReflect.Descriptor instead.
func (*BorrowResponse) Descriptor() ([]byte, []int) {
	return file_common_proto_borrowing_proto_rawDescGZIP(), []int{2}
}

func (x *BorrowResponse) GetBorrowingId() int32 {
	if x != nil {
		return x.BorrowingId
	}
	return 0
}

func (x *BorrowResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ReturnResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BorrowingId int32  `protobuf:"varint,1,opt,name=borrowing_id,json=borrowingId,proto3" json:"borrowing_id,omitempty"`
	Message     string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ReturnResponse) Reset() {
	*x = ReturnResponse{}
	mi := &file_common_proto_borrowing_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReturnResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReturnResponse) ProtoMessage() {}

func (x *ReturnResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_borrowing_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReturnResponse.ProtoReflect.Descriptor instead.
func (*ReturnResponse) Descriptor() ([]byte, []int) {
	return file_common_proto_borrowing_proto_rawDescGZIP(), []int{3}
}

func (x *ReturnResponse) GetBorrowingId() int32 {
	if x != nil {
		return x.BorrowingId
	}
	return 0
}

func (x *ReturnResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UserBorrowingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserBorrowingsRequest) Reset() {
	*x = UserBorrowingsRequest{}
	mi := &file_common_proto_borrowing_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserBorrowingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBorrowingsRequest) ProtoMessage() {}

func (x *UserBorrowingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_borrowing_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBorrowingsRequest.ProtoReflect.Descriptor instead.
func (*UserBorrowingsRequest) Descriptor() ([]byte, []int) {
	return file_common_proto_borrowing_proto_rawDescGZIP(), []int{4}
}

func (x *UserBorrowingsRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type BorrowingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BorrowingId int32  `protobuf:"varint,1,opt,name=borrowing_id,json=borrowingId,proto3" json:"borrowing_id,omitempty"`
	UserId      int32  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	BookId      int32  `protobuf:"varint,3,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	BorrowDate  string `protobuf:"bytes,4,opt,name=borrow_date,json=borrowDate,proto3" json:"borrow_date,omitempty"`
	ReturnDate  string `protobuf:"bytes,5,opt,name=return_date,json=returnDate,proto3" json:"return_date,omitempty"`
}

func (x *BorrowingResponse) Reset() {
	*x = BorrowingResponse{}
	mi := &file_common_proto_borrowing_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BorrowingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BorrowingResponse) ProtoMessage() {}

func (x *BorrowingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_borrowing_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BorrowingResponse.ProtoReflect.Descriptor instead.
func (*BorrowingResponse) Descriptor() ([]byte, []int) {
	return file_common_proto_borrowing_proto_rawDescGZIP(), []int{5}
}

func (x *BorrowingResponse) GetBorrowingId() int32 {
	if x != nil {
		return x.BorrowingId
	}
	return 0
}

func (x *BorrowingResponse) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *BorrowingResponse) GetBookId() int32 {
	if x != nil {
		return x.BookId
	}
	return 0
}

func (x *BorrowingResponse) GetBorrowDate() string {
	if x != nil {
		return x.BorrowDate
	}
	return ""
}

func (x *BorrowingResponse) GetReturnDate() string {
	if x != nil {
		return x.ReturnDate
	}
	return ""
}

type BorrowingListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Borrowings []*BorrowingResponse `protobuf:"bytes,1,rep,name=borrowings,proto3" json:"borrowings,omitempty"`
}

func (x *BorrowingListResponse) Reset() {
	*x = BorrowingListResponse{}
	mi := &file_common_proto_borrowing_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BorrowingListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BorrowingListResponse) ProtoMessage() {}

func (x *BorrowingListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_borrowing_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BorrowingListResponse.ProtoReflect.Descriptor instead.
func (*BorrowingListResponse) Descriptor() ([]byte, []int) {
	return file_common_proto_borrowing_proto_rawDescGZIP(), []int{6}
}

func (x *BorrowingListResponse) GetBorrowings() []*BorrowingResponse {
	if x != nil {
		return x.Borrowings
	}
	return nil
}

var File_common_proto_borrowing_proto protoreflect.FileDescriptor

var file_common_proto_borrowing_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62,
	0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x22, 0x41, 0x0a, 0x0d, 0x42, 0x6f, 0x72,
	0x72, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x0d,
	0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x49, 0x64,
	0x22, 0x4d, 0x0a, 0x0e, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77,
	0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x4d, 0x0a, 0x0e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69,
	0x6e, 0x67, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x30,
	0x0a, 0x15, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0xaa, 0x01, 0x0a, 0x11, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77,
	0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x62, 0x6f,
	0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x62,
	0x6f, 0x72, 0x72, 0x6f, 0x77, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x22, 0x55, 0x0a,
	0x15, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0a, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77,
	0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x6f, 0x72,
	0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77,
	0x69, 0x6e, 0x67, 0x73, 0x32, 0xf3, 0x01, 0x0a, 0x10, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x42, 0x6f, 0x72,
	0x72, 0x6f, 0x77, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x18, 0x2e, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77,
	0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f,
	0x72, 0x72, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a,
	0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x18, 0x2e, 0x62, 0x6f, 0x72,
	0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67,
	0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x59, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x73,
	0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69,
	0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x62, 0x6f, 0x72, 0x72, 0x6f,
	0x77, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x50, 0x5a, 0x4e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x61, 0x6e, 0x64, 0x68, 0x79, 0x61,
	0x2d, 0x50, 0x72, 0x61, 0x74, 0x61, 0x6d, 0x61, 0x2f, 0x4c, 0x69, 0x62, 0x61, 0x72, 0x79, 0x2d,
	0x41, 0x50, 0x49, 0x2f, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x3b, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_borrowing_proto_rawDescOnce sync.Once
	file_common_proto_borrowing_proto_rawDescData = file_common_proto_borrowing_proto_rawDesc
)

func file_common_proto_borrowing_proto_rawDescGZIP() []byte {
	file_common_proto_borrowing_proto_rawDescOnce.Do(func() {
		file_common_proto_borrowing_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_borrowing_proto_rawDescData)
	})
	return file_common_proto_borrowing_proto_rawDescData
}

var file_common_proto_borrowing_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_common_proto_borrowing_proto_goTypes = []any{
	(*BorrowRequest)(nil),         // 0: borrowing.BorrowRequest
	(*ReturnRequest)(nil),         // 1: borrowing.ReturnRequest
	(*BorrowResponse)(nil),        // 2: borrowing.BorrowResponse
	(*ReturnResponse)(nil),        // 3: borrowing.ReturnResponse
	(*UserBorrowingsRequest)(nil), // 4: borrowing.UserBorrowingsRequest
	(*BorrowingResponse)(nil),     // 5: borrowing.BorrowingResponse
	(*BorrowingListResponse)(nil), // 6: borrowing.BorrowingListResponse
}
var file_common_proto_borrowing_proto_depIdxs = []int32{
	5, // 0: borrowing.BorrowingListResponse.borrowings:type_name -> borrowing.BorrowingResponse
	0, // 1: borrowing.BorrowingService.BorrowBook:input_type -> borrowing.BorrowRequest
	1, // 2: borrowing.BorrowingService.ReturnBook:input_type -> borrowing.ReturnRequest
	4, // 3: borrowing.BorrowingService.GetBorrowingsByUser:input_type -> borrowing.UserBorrowingsRequest
	2, // 4: borrowing.BorrowingService.BorrowBook:output_type -> borrowing.BorrowResponse
	3, // 5: borrowing.BorrowingService.ReturnBook:output_type -> borrowing.ReturnResponse
	6, // 6: borrowing.BorrowingService.GetBorrowingsByUser:output_type -> borrowing.BorrowingListResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_common_proto_borrowing_proto_init() }
func file_common_proto_borrowing_proto_init() {
	if File_common_proto_borrowing_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_proto_borrowing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_common_proto_borrowing_proto_goTypes,
		DependencyIndexes: file_common_proto_borrowing_proto_depIdxs,
		MessageInfos:      file_common_proto_borrowing_proto_msgTypes,
	}.Build()
	File_common_proto_borrowing_proto = out.File
	file_common_proto_borrowing_proto_rawDesc = nil
	file_common_proto_borrowing_proto_goTypes = nil
	file_common_proto_borrowing_proto_depIdxs = nil
}
