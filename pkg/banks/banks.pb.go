// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: banks.proto

package banks

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Bank struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ExternalId       int64                  `protobuf:"varint,2,opt,name=external_id,proto3" json:"external_id,omitempty"`
	ExternalLegacyId int64                  `protobuf:"varint,3,opt,name=external_legacy_id,proto3" json:"external_legacy_id,omitempty"`
	Name             string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Logo             string                 `protobuf:"bytes,5,opt,name=logo,proto3" json:"logo,omitempty"`
	Url              string                 `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
	CreatedAt        *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt        *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	PeriodFrom       int64                  `protobuf:"varint,9,opt,name=period_from,proto3" json:"period_from,omitempty"`
	PeriodTo         int64                  `protobuf:"varint,10,opt,name=period_to,proto3" json:"period_to,omitempty"`
	AmountFrom       int64                  `protobuf:"varint,11,opt,name=amount_from,proto3" json:"amount_from,omitempty"`
	AmountTo         int64                  `protobuf:"varint,12,opt,name=amount_to,proto3" json:"amount_to,omitempty"`
	RateFrom         float64                `protobuf:"fixed64,13,opt,name=rate_from,proto3" json:"rate_from,omitempty"`
	RateTo           float64                `protobuf:"fixed64,14,opt,name=rate_to,proto3" json:"rate_to,omitempty"`
	ReviewTime       int64                  `protobuf:"varint,15,opt,name=review_time,proto3" json:"review_time,omitempty"`
	Registration     []string               `protobuf:"bytes,16,rep,name=registration,proto3" json:"registration,omitempty"`
	CreditPurpose    []string               `protobuf:"bytes,17,rep,name=credit_purpose,proto3" json:"credit_purpose,omitempty"`
	Documents        []string               `protobuf:"bytes,18,rep,name=documents,proto3" json:"documents,omitempty"`
	ObtainMethod     []string               `protobuf:"bytes,19,rep,name=obtain_method,proto3" json:"obtain_method,omitempty"`
}

func (x *Bank) Reset() {
	*x = Bank{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banks_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bank) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bank) ProtoMessage() {}

func (x *Bank) ProtoReflect() protoreflect.Message {
	mi := &file_banks_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bank.ProtoReflect.Descriptor instead.
func (*Bank) Descriptor() ([]byte, []int) {
	return file_banks_proto_rawDescGZIP(), []int{0}
}

func (x *Bank) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Bank) GetExternalId() int64 {
	if x != nil {
		return x.ExternalId
	}
	return 0
}

func (x *Bank) GetExternalLegacyId() int64 {
	if x != nil {
		return x.ExternalLegacyId
	}
	return 0
}

func (x *Bank) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Bank) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *Bank) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Bank) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Bank) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Bank) GetPeriodFrom() int64 {
	if x != nil {
		return x.PeriodFrom
	}
	return 0
}

func (x *Bank) GetPeriodTo() int64 {
	if x != nil {
		return x.PeriodTo
	}
	return 0
}

func (x *Bank) GetAmountFrom() int64 {
	if x != nil {
		return x.AmountFrom
	}
	return 0
}

func (x *Bank) GetAmountTo() int64 {
	if x != nil {
		return x.AmountTo
	}
	return 0
}

func (x *Bank) GetRateFrom() float64 {
	if x != nil {
		return x.RateFrom
	}
	return 0
}

func (x *Bank) GetRateTo() float64 {
	if x != nil {
		return x.RateTo
	}
	return 0
}

func (x *Bank) GetReviewTime() int64 {
	if x != nil {
		return x.ReviewTime
	}
	return 0
}

func (x *Bank) GetRegistration() []string {
	if x != nil {
		return x.Registration
	}
	return nil
}

func (x *Bank) GetCreditPurpose() []string {
	if x != nil {
		return x.CreditPurpose
	}
	return nil
}

func (x *Bank) GetDocuments() []string {
	if x != nil {
		return x.Documents
	}
	return nil
}

func (x *Bank) GetObtainMethod() []string {
	if x != nil {
		return x.ObtainMethod
	}
	return nil
}

type GetBanksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  uint64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset uint64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *GetBanksRequest) Reset() {
	*x = GetBanksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banks_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBanksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBanksRequest) ProtoMessage() {}

func (x *GetBanksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_banks_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBanksRequest.ProtoReflect.Descriptor instead.
func (*GetBanksRequest) Descriptor() ([]byte, []int) {
	return file_banks_proto_rawDescGZIP(), []int{1}
}

func (x *GetBanksRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetBanksRequest) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetBanksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Banks []*Bank `protobuf:"bytes,1,rep,name=banks,proto3" json:"banks,omitempty"`
}

func (x *GetBanksResponse) Reset() {
	*x = GetBanksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banks_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBanksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBanksResponse) ProtoMessage() {}

func (x *GetBanksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_banks_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBanksResponse.ProtoReflect.Descriptor instead.
func (*GetBanksResponse) Descriptor() ([]byte, []int) {
	return file_banks_proto_rawDescGZIP(), []int{2}
}

func (x *GetBanksResponse) GetBanks() []*Bank {
	if x != nil {
		return x.Banks
	}
	return nil
}

type UpdateBankRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bank *Bank `protobuf:"bytes,1,opt,name=bank,proto3" json:"bank,omitempty"`
}

func (x *UpdateBankRequest) Reset() {
	*x = UpdateBankRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banks_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBankRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBankRequest) ProtoMessage() {}

func (x *UpdateBankRequest) ProtoReflect() protoreflect.Message {
	mi := &file_banks_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBankRequest.ProtoReflect.Descriptor instead.
func (*UpdateBankRequest) Descriptor() ([]byte, []int) {
	return file_banks_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateBankRequest) GetBank() *Bank {
	if x != nil {
		return x.Bank
	}
	return nil
}

type UpdateBankResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateBankResponse) Reset() {
	*x = UpdateBankResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banks_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBankResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBankResponse) ProtoMessage() {}

func (x *UpdateBankResponse) ProtoReflect() protoreflect.Message {
	mi := &file_banks_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBankResponse.ProtoReflect.Descriptor instead.
func (*UpdateBankResponse) Descriptor() ([]byte, []int) {
	return file_banks_proto_rawDescGZIP(), []int{4}
}

var File_banks_proto protoreflect.FileDescriptor

var file_banks_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x61, 0x6e, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x77,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x05, 0x0a, 0x04, 0x42, 0x61, 0x6e, 0x6b, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x12,
	0x2e, 0x0a, 0x12, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6c, 0x65, 0x67, 0x61,
	0x63, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x5f, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x66, 0x72, 0x6f,
	0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x74, 0x6f, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x74, 0x6f, 0x12,
	0x20, 0x0a, 0x0b, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x66, 0x72, 0x6f,
	0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x6f, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x6f, 0x12,
	0x1c, 0x0a, 0x09, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x09, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x6f, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07,
	0x72, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x10, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a,
	0x0e, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x18,
	0x11, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x70, 0x75,
	0x72, 0x70, 0x6f, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x12, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x6f, 0x62, 0x74, 0x61, 0x69, 0x6e, 0x5f, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x13, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x6f, 0x62, 0x74, 0x61,
	0x69, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x3f, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x42, 0x61, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x36, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x42, 0x61, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22,
	0x0a, 0x05, 0x62, 0x61, 0x6e, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x05, 0x62, 0x61, 0x6e,
	0x6b, 0x73, 0x22, 0x35, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x04, 0x62, 0x61, 0x6e, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x42,
	0x61, 0x6e, 0x6b, 0x52, 0x04, 0x62, 0x61, 0x6e, 0x6b, 0x22, 0x14, 0x0a, 0x12, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xc2, 0x01, 0x0a, 0x05, 0x42, 0x61, 0x6e, 0x6b, 0x73, 0x12, 0x57, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x42, 0x61, 0x6e, 0x6b, 0x73, 0x12, 0x17, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x61, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6e, 0x6b, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12,
	0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x73, 0x2f, 0x67,
	0x65, 0x74, 0x12, 0x60, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6b,
	0x12, 0x19, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x77, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a,
	0x01, 0x2a, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x73, 0x2f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x42, 0x15, 0x5a, 0x13, 0x62, 0x61, 0x6e, 0x6b, 0x73, 0x2d, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_banks_proto_rawDescOnce sync.Once
	file_banks_proto_rawDescData = file_banks_proto_rawDesc
)

func file_banks_proto_rawDescGZIP() []byte {
	file_banks_proto_rawDescOnce.Do(func() {
		file_banks_proto_rawDescData = protoimpl.X.CompressGZIP(file_banks_proto_rawDescData)
	})
	return file_banks_proto_rawDescData
}

var file_banks_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_banks_proto_goTypes = []any{
	(*Bank)(nil),                  // 0: worker.Bank
	(*GetBanksRequest)(nil),       // 1: worker.GetBanksRequest
	(*GetBanksResponse)(nil),      // 2: worker.GetBanksResponse
	(*UpdateBankRequest)(nil),     // 3: worker.UpdateBankRequest
	(*UpdateBankResponse)(nil),    // 4: worker.UpdateBankResponse
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_banks_proto_depIdxs = []int32{
	5, // 0: worker.Bank.created_at:type_name -> google.protobuf.Timestamp
	5, // 1: worker.Bank.updated_at:type_name -> google.protobuf.Timestamp
	0, // 2: worker.GetBanksResponse.banks:type_name -> worker.Bank
	0, // 3: worker.UpdateBankRequest.bank:type_name -> worker.Bank
	1, // 4: worker.Banks.GetBanks:input_type -> worker.GetBanksRequest
	3, // 5: worker.Banks.UpdateBank:input_type -> worker.UpdateBankRequest
	2, // 6: worker.Banks.GetBanks:output_type -> worker.GetBanksResponse
	4, // 7: worker.Banks.UpdateBank:output_type -> worker.UpdateBankResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_banks_proto_init() }
func file_banks_proto_init() {
	if File_banks_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_banks_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Bank); i {
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
		file_banks_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetBanksRequest); i {
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
		file_banks_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetBanksResponse); i {
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
		file_banks_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateBankRequest); i {
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
		file_banks_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateBankResponse); i {
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
			RawDescriptor: file_banks_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_banks_proto_goTypes,
		DependencyIndexes: file_banks_proto_depIdxs,
		MessageInfos:      file_banks_proto_msgTypes,
	}.Build()
	File_banks_proto = out.File
	file_banks_proto_rawDesc = nil
	file_banks_proto_goTypes = nil
	file_banks_proto_depIdxs = nil
}
