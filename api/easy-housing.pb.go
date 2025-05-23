// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: api/easy-housing.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// AppointmentType represents the type of appointment to schedule.
type AppointmentType int32

const (
	AppointmentType_APPRAISAL  AppointmentType = 0
	AppointmentType_INSPECTION AppointmentType = 1
)

// Enum value maps for AppointmentType.
var (
	AppointmentType_name = map[int32]string{
		0: "APPRAISAL",
		1: "INSPECTION",
	}
	AppointmentType_value = map[string]int32{
		"APPRAISAL":  0,
		"INSPECTION": 1,
	}
)

func (x AppointmentType) Enum() *AppointmentType {
	p := new(AppointmentType)
	*p = x
	return p
}

func (x AppointmentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppointmentType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_easy_housing_proto_enumTypes[0].Descriptor()
}

func (AppointmentType) Type() protoreflect.EnumType {
	return &file_api_easy_housing_proto_enumTypes[0]
}

func (x AppointmentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AppointmentType.Descriptor instead.
func (AppointmentType) EnumDescriptor() ([]byte, []int) {
	return file_api_easy_housing_proto_rawDescGZIP(), []int{0}
}

type EchoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       *string                `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EchoRequest) Reset() {
	*x = EchoRequest{}
	mi := &file_api_easy_housing_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EchoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoRequest) ProtoMessage() {}

func (x *EchoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_easy_housing_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoRequest.ProtoReflect.Descriptor instead.
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return file_api_easy_housing_proto_rawDescGZIP(), []int{0}
}

func (x *EchoRequest) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

type EchoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       *string                `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EchoResponse) Reset() {
	*x = EchoResponse{}
	mi := &file_api_easy_housing_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EchoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoResponse) ProtoMessage() {}

func (x *EchoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_easy_housing_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoResponse.ProtoReflect.Descriptor instead.
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return file_api_easy_housing_proto_rawDescGZIP(), []int{1}
}

func (x *EchoResponse) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

// RegisterUserRequest represents a request to register a user.
type RegisterUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterUserRequest) Reset() {
	*x = RegisterUserRequest{}
	mi := &file_api_easy_housing_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUserRequest) ProtoMessage() {}

func (x *RegisterUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_easy_housing_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUserRequest.ProtoReflect.Descriptor instead.
func (*RegisterUserRequest) Descriptor() ([]byte, []int) {
	return file_api_easy_housing_proto_rawDescGZIP(), []int{2}
}

// RegisterUserResponse represents a response to a request to register a user.
type RegisterUserResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// message is a message to display to the user.
	Message *string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	// userId is the id of the user
	UserId        *string `protobuf:"bytes,2,opt,name=userId" json:"userId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterUserResponse) Reset() {
	*x = RegisterUserResponse{}
	mi := &file_api_easy_housing_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUserResponse) ProtoMessage() {}

func (x *RegisterUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_easy_housing_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUserResponse.ProtoReflect.Descriptor instead.
func (*RegisterUserResponse) Descriptor() ([]byte, []int) {
	return file_api_easy_housing_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterUserResponse) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

func (x *RegisterUserResponse) GetUserId() string {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return ""
}

// AuthenticateUserRequest represents a request to authenticate a user.
type AuthenticateUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthenticateUserRequest) Reset() {
	*x = AuthenticateUserRequest{}
	mi := &file_api_easy_housing_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthenticateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateUserRequest) ProtoMessage() {}

func (x *AuthenticateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_easy_housing_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateUserRequest.ProtoReflect.Descriptor instead.
func (*AuthenticateUserRequest) Descriptor() ([]byte, []int) {
	return file_api_easy_housing_proto_rawDescGZIP(), []int{4}
}

// AuthenticateUserResponse represents a response to a request to authenticate a user.
type AuthenticateUserResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// message is a message to display to the user.
	Message       *string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthenticateUserResponse) Reset() {
	*x = AuthenticateUserResponse{}
	mi := &file_api_easy_housing_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthenticateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateUserResponse) ProtoMessage() {}

func (x *AuthenticateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_easy_housing_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateUserResponse.ProtoReflect.Descriptor instead.
func (*AuthenticateUserResponse) Descriptor() ([]byte, []int) {
	return file_api_easy_housing_proto_rawDescGZIP(), []int{5}
}

func (x *AuthenticateUserResponse) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

// RegisterHomeRequest represents a request to register a home.
type RegisterHomeRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// address is the address of the home
	Address *string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	// ownerId is the id of the owner of the home
	OwnerId       *string `protobuf:"bytes,2,opt,name=ownerId" json:"ownerId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterHomeRequest) Reset() {
	*x = RegisterHomeRequest{}
	mi := &file_api_easy_housing_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterHomeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterHomeRequest) ProtoMessage() {}

func (x *RegisterHomeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_easy_housing_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterHomeRequest.ProtoReflect.Descriptor instead.
func (*RegisterHomeRequest) Descriptor() ([]byte, []int) {
	return file_api_easy_housing_proto_rawDescGZIP(), []int{6}
}

func (x *RegisterHomeRequest) GetAddress() string {
	if x != nil && x.Address != nil {
		return *x.Address
	}
	return ""
}

func (x *RegisterHomeRequest) GetOwnerId() string {
	if x != nil && x.OwnerId != nil {
		return *x.OwnerId
	}
	return ""
}

// RegisterHomeResponse represents a response to a request to register a home.
type RegisterHomeResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// message is a message to display to the user.
	Message       *string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterHomeResponse) Reset() {
	*x = RegisterHomeResponse{}
	mi := &file_api_easy_housing_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterHomeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterHomeResponse) ProtoMessage() {}

func (x *RegisterHomeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_easy_housing_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterHomeResponse.ProtoReflect.Descriptor instead.
func (*RegisterHomeResponse) Descriptor() ([]byte, []int) {
	return file_api_easy_housing_proto_rawDescGZIP(), []int{7}
}

func (x *RegisterHomeResponse) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

// AddDocRequest represents a request to add docs to a home.
type AddDocRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// homeId is the id of the home to add docs to
	HomeId *string `protobuf:"bytes,1,opt,name=homeId" json:"homeId,omitempty"`
	// docLocation is the object storage location of the provided doc
	DocLocation   *string `protobuf:"bytes,2,opt,name=docLocation" json:"docLocation,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddDocRequest) Reset() {
	*x = AddDocRequest{}
	mi := &file_api_easy_housing_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddDocRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDocRequest) ProtoMessage() {}

func (x *AddDocRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_easy_housing_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDocRequest.ProtoReflect.Descriptor instead.
func (*AddDocRequest) Descriptor() ([]byte, []int) {
	return file_api_easy_housing_proto_rawDescGZIP(), []int{8}
}

func (x *AddDocRequest) GetHomeId() string {
	if x != nil && x.HomeId != nil {
		return *x.HomeId
	}
	return ""
}

func (x *AddDocRequest) GetDocLocation() string {
	if x != nil && x.DocLocation != nil {
		return *x.DocLocation
	}
	return ""
}

// AddDocResponse represents a response to a request to add docs to a home.
type AddDocResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// message is a message to display to the user.
	Message       *string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddDocResponse) Reset() {
	*x = AddDocResponse{}
	mi := &file_api_easy_housing_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddDocResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDocResponse) ProtoMessage() {}

func (x *AddDocResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_easy_housing_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDocResponse.ProtoReflect.Descriptor instead.
func (*AddDocResponse) Descriptor() ([]byte, []int) {
	return file_api_easy_housing_proto_rawDescGZIP(), []int{9}
}

func (x *AddDocResponse) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

// ScheduleAppointmentRequest represents a request to schedule an appointment for a home.
type ScheduleAppointmentRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// homeId is the id of the home to schedule an appointment for
	HomeId *string `protobuf:"bytes,1,opt,name=homeId" json:"homeId,omitempty"`
	// appointmentTime is in seconds since epoch
	AppointmentTime *int64 `protobuf:"varint,2,opt,name=appointmentTime" json:"appointmentTime,omitempty"`
	// appointmentType is the type of appointment to schedule
	AppointmentType *AppointmentType `protobuf:"varint,3,opt,name=appointmentType,enum=easy_housing.AppointmentType" json:"appointmentType,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ScheduleAppointmentRequest) Reset() {
	*x = ScheduleAppointmentRequest{}
	mi := &file_api_easy_housing_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScheduleAppointmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleAppointmentRequest) ProtoMessage() {}

func (x *ScheduleAppointmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_easy_housing_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleAppointmentRequest.ProtoReflect.Descriptor instead.
func (*ScheduleAppointmentRequest) Descriptor() ([]byte, []int) {
	return file_api_easy_housing_proto_rawDescGZIP(), []int{10}
}

func (x *ScheduleAppointmentRequest) GetHomeId() string {
	if x != nil && x.HomeId != nil {
		return *x.HomeId
	}
	return ""
}

func (x *ScheduleAppointmentRequest) GetAppointmentTime() int64 {
	if x != nil && x.AppointmentTime != nil {
		return *x.AppointmentTime
	}
	return 0
}

func (x *ScheduleAppointmentRequest) GetAppointmentType() AppointmentType {
	if x != nil && x.AppointmentType != nil {
		return *x.AppointmentType
	}
	return AppointmentType_APPRAISAL
}

// ScheduleAppointmentResponse represents a response to a request to schedule an appointment for a home.
type ScheduleAppointmentResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// message is a message to display to the user.
	Message       *string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ScheduleAppointmentResponse) Reset() {
	*x = ScheduleAppointmentResponse{}
	mi := &file_api_easy_housing_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScheduleAppointmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleAppointmentResponse) ProtoMessage() {}

func (x *ScheduleAppointmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_easy_housing_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleAppointmentResponse.ProtoReflect.Descriptor instead.
func (*ScheduleAppointmentResponse) Descriptor() ([]byte, []int) {
	return file_api_easy_housing_proto_rawDescGZIP(), []int{11}
}

func (x *ScheduleAppointmentResponse) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

// CompleteListingRequest represents a request to complete a listing for a home.
type CompleteListingRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// homeId is the id of the home to complete a listing for
	HomeId        *string `protobuf:"bytes,1,opt,name=homeId" json:"homeId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CompleteListingRequest) Reset() {
	*x = CompleteListingRequest{}
	mi := &file_api_easy_housing_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompleteListingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteListingRequest) ProtoMessage() {}

func (x *CompleteListingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_easy_housing_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteListingRequest.ProtoReflect.Descriptor instead.
func (*CompleteListingRequest) Descriptor() ([]byte, []int) {
	return file_api_easy_housing_proto_rawDescGZIP(), []int{12}
}

func (x *CompleteListingRequest) GetHomeId() string {
	if x != nil && x.HomeId != nil {
		return *x.HomeId
	}
	return ""
}

// CompleteListingResponse represents a response to a request to complete a listing for a home.
type CompleteListingResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// message is a message to display to the user.
	Message       *string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CompleteListingResponse) Reset() {
	*x = CompleteListingResponse{}
	mi := &file_api_easy_housing_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompleteListingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteListingResponse) ProtoMessage() {}

func (x *CompleteListingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_easy_housing_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteListingResponse.ProtoReflect.Descriptor instead.
func (*CompleteListingResponse) Descriptor() ([]byte, []int) {
	return file_api_easy_housing_proto_rawDescGZIP(), []int{13}
}

func (x *CompleteListingResponse) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

// GetListingStatusRequest represents a request to get the status of a home.
type GetListingStatusRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// homeId is the id of the home to get the status of
	HomeId        *string `protobuf:"bytes,1,opt,name=homeId" json:"homeId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetListingStatusRequest) Reset() {
	*x = GetListingStatusRequest{}
	mi := &file_api_easy_housing_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetListingStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListingStatusRequest) ProtoMessage() {}

func (x *GetListingStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_easy_housing_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListingStatusRequest.ProtoReflect.Descriptor instead.
func (*GetListingStatusRequest) Descriptor() ([]byte, []int) {
	return file_api_easy_housing_proto_rawDescGZIP(), []int{14}
}

func (x *GetListingStatusRequest) GetHomeId() string {
	if x != nil && x.HomeId != nil {
		return *x.HomeId
	}
	return ""
}

// GetListingStatusResponse represents a response to a request to get the status of a home.
type GetListingStatusResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// message is a message to display to the user.
	Message *string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	// ownershipVerified is true if the home has been verified as owned by the user
	OwnershipVerified *bool `protobuf:"varint,3,opt,name=ownershipVerified" json:"ownershipVerified,omitempty"`
	// inspectionCompleted is true if the home has been inspected
	InspectionCompleted *bool `protobuf:"varint,4,opt,name=inspectionCompleted" json:"inspectionCompleted,omitempty"`
	// appraisalCompleted is true if the home has been appraised
	AppraisalCompleted *bool `protobuf:"varint,5,opt,name=appraisalCompleted" json:"appraisalCompleted,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *GetListingStatusResponse) Reset() {
	*x = GetListingStatusResponse{}
	mi := &file_api_easy_housing_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetListingStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListingStatusResponse) ProtoMessage() {}

func (x *GetListingStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_easy_housing_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListingStatusResponse.ProtoReflect.Descriptor instead.
func (*GetListingStatusResponse) Descriptor() ([]byte, []int) {
	return file_api_easy_housing_proto_rawDescGZIP(), []int{15}
}

func (x *GetListingStatusResponse) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

func (x *GetListingStatusResponse) GetOwnershipVerified() bool {
	if x != nil && x.OwnershipVerified != nil {
		return *x.OwnershipVerified
	}
	return false
}

func (x *GetListingStatusResponse) GetInspectionCompleted() bool {
	if x != nil && x.InspectionCompleted != nil {
		return *x.InspectionCompleted
	}
	return false
}

func (x *GetListingStatusResponse) GetAppraisalCompleted() bool {
	if x != nil && x.AppraisalCompleted != nil {
		return *x.AppraisalCompleted
	}
	return false
}

var File_api_easy_housing_proto protoreflect.FileDescriptor

const file_api_easy_housing_proto_rawDesc = "" +
	"\n" +
	"\x16api/easy-housing.proto\x12\feasy_housing\"'\n" +
	"\vEchoRequest\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"(\n" +
	"\fEchoResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"\x15\n" +
	"\x13RegisterUserRequest\"H\n" +
	"\x14RegisterUserResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\x12\x16\n" +
	"\x06userId\x18\x02 \x01(\tR\x06userId\"\x19\n" +
	"\x17AuthenticateUserRequest\"4\n" +
	"\x18AuthenticateUserResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"I\n" +
	"\x13RegisterHomeRequest\x12\x18\n" +
	"\aaddress\x18\x01 \x01(\tR\aaddress\x12\x18\n" +
	"\aownerId\x18\x02 \x01(\tR\aownerId\"0\n" +
	"\x14RegisterHomeResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"I\n" +
	"\rAddDocRequest\x12\x16\n" +
	"\x06homeId\x18\x01 \x01(\tR\x06homeId\x12 \n" +
	"\vdocLocation\x18\x02 \x01(\tR\vdocLocation\"*\n" +
	"\x0eAddDocResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"\xa7\x01\n" +
	"\x1aScheduleAppointmentRequest\x12\x16\n" +
	"\x06homeId\x18\x01 \x01(\tR\x06homeId\x12(\n" +
	"\x0fappointmentTime\x18\x02 \x01(\x03R\x0fappointmentTime\x12G\n" +
	"\x0fappointmentType\x18\x03 \x01(\x0e2\x1d.easy_housing.AppointmentTypeR\x0fappointmentType\"7\n" +
	"\x1bScheduleAppointmentResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"0\n" +
	"\x16CompleteListingRequest\x12\x16\n" +
	"\x06homeId\x18\x01 \x01(\tR\x06homeId\"3\n" +
	"\x17CompleteListingResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"1\n" +
	"\x17GetListingStatusRequest\x12\x16\n" +
	"\x06homeId\x18\x01 \x01(\tR\x06homeId\"\xc4\x01\n" +
	"\x18GetListingStatusResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\x12,\n" +
	"\x11ownershipVerified\x18\x03 \x01(\bR\x11ownershipVerified\x120\n" +
	"\x13inspectionCompleted\x18\x04 \x01(\bR\x13inspectionCompleted\x12.\n" +
	"\x12appraisalCompleted\x18\x05 \x01(\bR\x12appraisalCompleted*0\n" +
	"\x0fAppointmentType\x12\r\n" +
	"\tAPPRAISAL\x10\x00\x12\x0e\n" +
	"\n" +
	"INSPECTION\x10\x012\xe1\x05\n" +
	"\vEasyHousing\x12?\n" +
	"\x04Echo\x12\x19.easy_housing.EchoRequest\x1a\x1a.easy_housing.EchoResponse\"\x00\x12W\n" +
	"\fRegisterUser\x12!.easy_housing.RegisterUserRequest\x1a\".easy_housing.RegisterUserResponse\"\x00\x12c\n" +
	"\x10AuthenticateUser\x12%.easy_housing.AuthenticateUserRequest\x1a&.easy_housing.AuthenticateUserResponse\"\x00\x12W\n" +
	"\fRegisterHome\x12!.easy_housing.RegisterHomeRequest\x1a\".easy_housing.RegisterHomeResponse\"\x00\x12E\n" +
	"\x06AddDoc\x12\x1b.easy_housing.AddDocRequest\x1a\x1c.easy_housing.AddDocResponse\"\x00\x12l\n" +
	"\x13ScheduleAppointment\x12(.easy_housing.ScheduleAppointmentRequest\x1a).easy_housing.ScheduleAppointmentResponse\"\x00\x12`\n" +
	"\x0fCompleteListing\x12$.easy_housing.CompleteListingRequest\x1a%.easy_housing.CompleteListingResponse\"\x00\x12c\n" +
	"\x10GetListingStatus\x12%.easy_housing.GetListingStatusRequest\x1a&.easy_housing.GetListingStatusResponse\"\x00B\x12Z\x10easy-housing/apib\beditionsp\xe8\a"

var (
	file_api_easy_housing_proto_rawDescOnce sync.Once
	file_api_easy_housing_proto_rawDescData []byte
)

func file_api_easy_housing_proto_rawDescGZIP() []byte {
	file_api_easy_housing_proto_rawDescOnce.Do(func() {
		file_api_easy_housing_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_easy_housing_proto_rawDesc), len(file_api_easy_housing_proto_rawDesc)))
	})
	return file_api_easy_housing_proto_rawDescData
}

var file_api_easy_housing_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_easy_housing_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_api_easy_housing_proto_goTypes = []any{
	(AppointmentType)(0),                // 0: easy_housing.AppointmentType
	(*EchoRequest)(nil),                 // 1: easy_housing.EchoRequest
	(*EchoResponse)(nil),                // 2: easy_housing.EchoResponse
	(*RegisterUserRequest)(nil),         // 3: easy_housing.RegisterUserRequest
	(*RegisterUserResponse)(nil),        // 4: easy_housing.RegisterUserResponse
	(*AuthenticateUserRequest)(nil),     // 5: easy_housing.AuthenticateUserRequest
	(*AuthenticateUserResponse)(nil),    // 6: easy_housing.AuthenticateUserResponse
	(*RegisterHomeRequest)(nil),         // 7: easy_housing.RegisterHomeRequest
	(*RegisterHomeResponse)(nil),        // 8: easy_housing.RegisterHomeResponse
	(*AddDocRequest)(nil),               // 9: easy_housing.AddDocRequest
	(*AddDocResponse)(nil),              // 10: easy_housing.AddDocResponse
	(*ScheduleAppointmentRequest)(nil),  // 11: easy_housing.ScheduleAppointmentRequest
	(*ScheduleAppointmentResponse)(nil), // 12: easy_housing.ScheduleAppointmentResponse
	(*CompleteListingRequest)(nil),      // 13: easy_housing.CompleteListingRequest
	(*CompleteListingResponse)(nil),     // 14: easy_housing.CompleteListingResponse
	(*GetListingStatusRequest)(nil),     // 15: easy_housing.GetListingStatusRequest
	(*GetListingStatusResponse)(nil),    // 16: easy_housing.GetListingStatusResponse
}
var file_api_easy_housing_proto_depIdxs = []int32{
	0,  // 0: easy_housing.ScheduleAppointmentRequest.appointmentType:type_name -> easy_housing.AppointmentType
	1,  // 1: easy_housing.EasyHousing.Echo:input_type -> easy_housing.EchoRequest
	3,  // 2: easy_housing.EasyHousing.RegisterUser:input_type -> easy_housing.RegisterUserRequest
	5,  // 3: easy_housing.EasyHousing.AuthenticateUser:input_type -> easy_housing.AuthenticateUserRequest
	7,  // 4: easy_housing.EasyHousing.RegisterHome:input_type -> easy_housing.RegisterHomeRequest
	9,  // 5: easy_housing.EasyHousing.AddDoc:input_type -> easy_housing.AddDocRequest
	11, // 6: easy_housing.EasyHousing.ScheduleAppointment:input_type -> easy_housing.ScheduleAppointmentRequest
	13, // 7: easy_housing.EasyHousing.CompleteListing:input_type -> easy_housing.CompleteListingRequest
	15, // 8: easy_housing.EasyHousing.GetListingStatus:input_type -> easy_housing.GetListingStatusRequest
	2,  // 9: easy_housing.EasyHousing.Echo:output_type -> easy_housing.EchoResponse
	4,  // 10: easy_housing.EasyHousing.RegisterUser:output_type -> easy_housing.RegisterUserResponse
	6,  // 11: easy_housing.EasyHousing.AuthenticateUser:output_type -> easy_housing.AuthenticateUserResponse
	8,  // 12: easy_housing.EasyHousing.RegisterHome:output_type -> easy_housing.RegisterHomeResponse
	10, // 13: easy_housing.EasyHousing.AddDoc:output_type -> easy_housing.AddDocResponse
	12, // 14: easy_housing.EasyHousing.ScheduleAppointment:output_type -> easy_housing.ScheduleAppointmentResponse
	14, // 15: easy_housing.EasyHousing.CompleteListing:output_type -> easy_housing.CompleteListingResponse
	16, // 16: easy_housing.EasyHousing.GetListingStatus:output_type -> easy_housing.GetListingStatusResponse
	9,  // [9:17] is the sub-list for method output_type
	1,  // [1:9] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_api_easy_housing_proto_init() }
func file_api_easy_housing_proto_init() {
	if File_api_easy_housing_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_easy_housing_proto_rawDesc), len(file_api_easy_housing_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_easy_housing_proto_goTypes,
		DependencyIndexes: file_api_easy_housing_proto_depIdxs,
		EnumInfos:         file_api_easy_housing_proto_enumTypes,
		MessageInfos:      file_api_easy_housing_proto_msgTypes,
	}.Build()
	File_api_easy_housing_proto = out.File
	file_api_easy_housing_proto_goTypes = nil
	file_api_easy_housing_proto_depIdxs = nil
}
