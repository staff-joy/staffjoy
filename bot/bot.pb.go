// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bot.proto

package bot

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	math "math"
	company "v2.staffjoy.com/company"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type OnboardWorkerRequest struct {
	CompanyUuid          string   `protobuf:"bytes,1,opt,name=company_uuid,json=companyUuid,proto3" json:"company_uuid,omitempty"`
	UserUuid             string   `protobuf:"bytes,2,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OnboardWorkerRequest) Reset()         { *m = OnboardWorkerRequest{} }
func (m *OnboardWorkerRequest) String() string { return proto.CompactTextString(m) }
func (*OnboardWorkerRequest) ProtoMessage()    {}
func (*OnboardWorkerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_51d7d70385167023, []int{0}
}
func (m *OnboardWorkerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OnboardWorkerRequest.Unmarshal(m, b)
}
func (m *OnboardWorkerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OnboardWorkerRequest.Marshal(b, m, deterministic)
}
func (m *OnboardWorkerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OnboardWorkerRequest.Merge(m, src)
}
func (m *OnboardWorkerRequest) XXX_Size() int {
	return xxx_messageInfo_OnboardWorkerRequest.Size(m)
}
func (m *OnboardWorkerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OnboardWorkerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OnboardWorkerRequest proto.InternalMessageInfo

func (m *OnboardWorkerRequest) GetCompanyUuid() string {
	if m != nil {
		return m.CompanyUuid
	}
	return ""
}

func (m *OnboardWorkerRequest) GetUserUuid() string {
	if m != nil {
		return m.UserUuid
	}
	return ""
}

type AlertNewShiftRequest struct {
	UserUuid             string         `protobuf:"bytes,1,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	NewShift             *company.Shift `protobuf:"bytes,2,opt,name=new_shift,json=newShift,proto3" json:"new_shift,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AlertNewShiftRequest) Reset()         { *m = AlertNewShiftRequest{} }
func (m *AlertNewShiftRequest) String() string { return proto.CompactTextString(m) }
func (*AlertNewShiftRequest) ProtoMessage()    {}
func (*AlertNewShiftRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_51d7d70385167023, []int{1}
}
func (m *AlertNewShiftRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlertNewShiftRequest.Unmarshal(m, b)
}
func (m *AlertNewShiftRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlertNewShiftRequest.Marshal(b, m, deterministic)
}
func (m *AlertNewShiftRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertNewShiftRequest.Merge(m, src)
}
func (m *AlertNewShiftRequest) XXX_Size() int {
	return xxx_messageInfo_AlertNewShiftRequest.Size(m)
}
func (m *AlertNewShiftRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertNewShiftRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AlertNewShiftRequest proto.InternalMessageInfo

func (m *AlertNewShiftRequest) GetUserUuid() string {
	if m != nil {
		return m.UserUuid
	}
	return ""
}

func (m *AlertNewShiftRequest) GetNewShift() *company.Shift {
	if m != nil {
		return m.NewShift
	}
	return nil
}

type AlertNewShiftsRequest struct {
	UserUuid             string           `protobuf:"bytes,1,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	NewShifts            []*company.Shift `protobuf:"bytes,2,rep,name=new_shifts,json=newShifts,proto3" json:"new_shifts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *AlertNewShiftsRequest) Reset()         { *m = AlertNewShiftsRequest{} }
func (m *AlertNewShiftsRequest) String() string { return proto.CompactTextString(m) }
func (*AlertNewShiftsRequest) ProtoMessage()    {}
func (*AlertNewShiftsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_51d7d70385167023, []int{2}
}
func (m *AlertNewShiftsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlertNewShiftsRequest.Unmarshal(m, b)
}
func (m *AlertNewShiftsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlertNewShiftsRequest.Marshal(b, m, deterministic)
}
func (m *AlertNewShiftsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertNewShiftsRequest.Merge(m, src)
}
func (m *AlertNewShiftsRequest) XXX_Size() int {
	return xxx_messageInfo_AlertNewShiftsRequest.Size(m)
}
func (m *AlertNewShiftsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertNewShiftsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AlertNewShiftsRequest proto.InternalMessageInfo

func (m *AlertNewShiftsRequest) GetUserUuid() string {
	if m != nil {
		return m.UserUuid
	}
	return ""
}

func (m *AlertNewShiftsRequest) GetNewShifts() []*company.Shift {
	if m != nil {
		return m.NewShifts
	}
	return nil
}

type AlertRemovedShiftRequest struct {
	UserUuid             string         `protobuf:"bytes,1,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	OldShift             *company.Shift `protobuf:"bytes,2,opt,name=old_shift,json=oldShift,proto3" json:"old_shift,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AlertRemovedShiftRequest) Reset()         { *m = AlertRemovedShiftRequest{} }
func (m *AlertRemovedShiftRequest) String() string { return proto.CompactTextString(m) }
func (*AlertRemovedShiftRequest) ProtoMessage()    {}
func (*AlertRemovedShiftRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_51d7d70385167023, []int{3}
}
func (m *AlertRemovedShiftRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlertRemovedShiftRequest.Unmarshal(m, b)
}
func (m *AlertRemovedShiftRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlertRemovedShiftRequest.Marshal(b, m, deterministic)
}
func (m *AlertRemovedShiftRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertRemovedShiftRequest.Merge(m, src)
}
func (m *AlertRemovedShiftRequest) XXX_Size() int {
	return xxx_messageInfo_AlertRemovedShiftRequest.Size(m)
}
func (m *AlertRemovedShiftRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertRemovedShiftRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AlertRemovedShiftRequest proto.InternalMessageInfo

func (m *AlertRemovedShiftRequest) GetUserUuid() string {
	if m != nil {
		return m.UserUuid
	}
	return ""
}

func (m *AlertRemovedShiftRequest) GetOldShift() *company.Shift {
	if m != nil {
		return m.OldShift
	}
	return nil
}

type AlertRemovedShiftsRequest struct {
	UserUuid             string           `protobuf:"bytes,1,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	OldShifts            []*company.Shift `protobuf:"bytes,2,rep,name=old_shifts,json=oldShifts,proto3" json:"old_shifts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *AlertRemovedShiftsRequest) Reset()         { *m = AlertRemovedShiftsRequest{} }
func (m *AlertRemovedShiftsRequest) String() string { return proto.CompactTextString(m) }
func (*AlertRemovedShiftsRequest) ProtoMessage()    {}
func (*AlertRemovedShiftsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_51d7d70385167023, []int{4}
}
func (m *AlertRemovedShiftsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlertRemovedShiftsRequest.Unmarshal(m, b)
}
func (m *AlertRemovedShiftsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlertRemovedShiftsRequest.Marshal(b, m, deterministic)
}
func (m *AlertRemovedShiftsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertRemovedShiftsRequest.Merge(m, src)
}
func (m *AlertRemovedShiftsRequest) XXX_Size() int {
	return xxx_messageInfo_AlertRemovedShiftsRequest.Size(m)
}
func (m *AlertRemovedShiftsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertRemovedShiftsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AlertRemovedShiftsRequest proto.InternalMessageInfo

func (m *AlertRemovedShiftsRequest) GetUserUuid() string {
	if m != nil {
		return m.UserUuid
	}
	return ""
}

func (m *AlertRemovedShiftsRequest) GetOldShifts() []*company.Shift {
	if m != nil {
		return m.OldShifts
	}
	return nil
}

type AlertChangedShiftRequest struct {
	UserUuid             string         `protobuf:"bytes,1,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	OldShift             *company.Shift `protobuf:"bytes,2,opt,name=old_shift,json=oldShift,proto3" json:"old_shift,omitempty"`
	NewShift             *company.Shift `protobuf:"bytes,3,opt,name=new_shift,json=newShift,proto3" json:"new_shift,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AlertChangedShiftRequest) Reset()         { *m = AlertChangedShiftRequest{} }
func (m *AlertChangedShiftRequest) String() string { return proto.CompactTextString(m) }
func (*AlertChangedShiftRequest) ProtoMessage()    {}
func (*AlertChangedShiftRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_51d7d70385167023, []int{5}
}
func (m *AlertChangedShiftRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlertChangedShiftRequest.Unmarshal(m, b)
}
func (m *AlertChangedShiftRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlertChangedShiftRequest.Marshal(b, m, deterministic)
}
func (m *AlertChangedShiftRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertChangedShiftRequest.Merge(m, src)
}
func (m *AlertChangedShiftRequest) XXX_Size() int {
	return xxx_messageInfo_AlertChangedShiftRequest.Size(m)
}
func (m *AlertChangedShiftRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertChangedShiftRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AlertChangedShiftRequest proto.InternalMessageInfo

func (m *AlertChangedShiftRequest) GetUserUuid() string {
	if m != nil {
		return m.UserUuid
	}
	return ""
}

func (m *AlertChangedShiftRequest) GetOldShift() *company.Shift {
	if m != nil {
		return m.OldShift
	}
	return nil
}

func (m *AlertChangedShiftRequest) GetNewShift() *company.Shift {
	if m != nil {
		return m.NewShift
	}
	return nil
}

func init() {
	proto.RegisterType((*OnboardWorkerRequest)(nil), "staffjoy.bot.OnboardWorkerRequest")
	proto.RegisterType((*AlertNewShiftRequest)(nil), "staffjoy.bot.AlertNewShiftRequest")
	proto.RegisterType((*AlertNewShiftsRequest)(nil), "staffjoy.bot.AlertNewShiftsRequest")
	proto.RegisterType((*AlertRemovedShiftRequest)(nil), "staffjoy.bot.AlertRemovedShiftRequest")
	proto.RegisterType((*AlertRemovedShiftsRequest)(nil), "staffjoy.bot.AlertRemovedShiftsRequest")
	proto.RegisterType((*AlertChangedShiftRequest)(nil), "staffjoy.bot.AlertChangedShiftRequest")
}

func init() { proto.RegisterFile("bot.proto", fileDescriptor_51d7d70385167023) }

var fileDescriptor_51d7d70385167023 = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0x4f, 0x6f, 0xda, 0x40,
	0x10, 0xc5, 0x6b, 0xa8, 0x2a, 0x3c, 0x40, 0xa5, 0x6e, 0xa1, 0xa5, 0x70, 0xa1, 0x1b, 0x29, 0xc9,
	0x69, 0x91, 0x48, 0x94, 0x7b, 0x88, 0x72, 0x4c, 0x90, 0x40, 0x04, 0x29, 0x17, 0x64, 0xe3, 0x35,
	0x38, 0xb1, 0xbd, 0x8e, 0x77, 0x0d, 0xe2, 0x9a, 0xaf, 0x92, 0x2f, 0x1a, 0xf9, 0x1f, 0x62, 0x83,
	0x63, 0xcc, 0x21, 0x47, 0x7b, 0x66, 0x7e, 0x6f, 0x9e, 0xf6, 0x0d, 0xa8, 0x3a, 0x13, 0xc4, 0xf3,
	0x99, 0x60, 0xa8, 0xc6, 0x85, 0x66, 0x9a, 0x4f, 0x6c, 0x43, 0x74, 0x26, 0xda, 0x9d, 0x05, 0x63,
	0x0b, 0x9b, 0xf6, 0xa2, 0x9a, 0x1e, 0x98, 0x3d, 0xea, 0x78, 0x62, 0x13, 0xb7, 0xb6, 0xeb, 0x73,
	0xe6, 0x78, 0x9a, 0x9b, 0x7c, 0xe2, 0x07, 0x68, 0x0c, 0x5d, 0x9d, 0x69, 0xbe, 0x31, 0x65, 0xfe,
	0x33, 0xf5, 0x47, 0xf4, 0x25, 0xa0, 0x5c, 0xa0, 0xff, 0x50, 0x4b, 0x1a, 0x67, 0x41, 0x60, 0x19,
	0x2d, 0xa5, 0xab, 0x9c, 0xab, 0xa3, 0x6a, 0xf2, 0x6f, 0x12, 0x58, 0x06, 0xea, 0x80, 0x1a, 0x70,
	0xea, 0xc7, 0xf5, 0x52, 0x54, 0xaf, 0x84, 0x3f, 0xc2, 0x22, 0xb6, 0xa0, 0x71, 0x6d, 0x53, 0x5f,
	0xdc, 0xd3, 0xf5, 0x78, 0x69, 0x99, 0x22, 0xe5, 0x4a, 0x43, 0x8a, 0x3c, 0x84, 0x2e, 0x41, 0x75,
	0xe9, 0x7a, 0xc6, 0xc3, 0x81, 0x88, 0x58, 0xed, 0xff, 0x25, 0x5b, 0x6b, 0xe9, 0xe2, 0x31, 0xaf,
	0xe2, 0x26, 0x64, 0x6c, 0x43, 0x53, 0x92, 0xe2, 0x85, 0xb4, 0xae, 0x00, 0xb6, 0x5a, 0xbc, 0x55,
	0xea, 0x96, 0xf3, 0xc4, 0xd4, 0x54, 0x8c, 0x63, 0x07, 0x5a, 0x91, 0xda, 0x88, 0x3a, 0x6c, 0x45,
	0x8d, 0xa3, 0xcc, 0x31, 0xdb, 0x28, 0x68, 0x8e, 0xd9, 0x31, 0x19, 0x7b, 0xf0, 0x6f, 0x4f, 0xae,
	0xb0, 0xc1, 0xad, 0xde, 0x61, 0x83, 0xa9, 0x20, 0xc7, 0x6f, 0x4a, 0xe2, 0xf0, 0x66, 0xa9, 0xb9,
	0x8b, 0x2f, 0x77, 0x28, 0x3f, 0x7a, 0xb9, 0xe0, 0xa3, 0xf7, 0x5f, 0xbf, 0x03, 0x0c, 0x98, 0x18,
	0x53, 0x7f, 0x65, 0xcd, 0x29, 0xba, 0x83, 0xba, 0x14, 0x63, 0x84, 0xc9, 0xee, 0x49, 0x90, 0xac,
	0x8c, 0xb7, 0xff, 0x90, 0xf8, 0x50, 0x48, 0x7a, 0x28, 0xe4, 0x36, 0x3c, 0x14, 0xfc, 0x2d, 0xc4,
	0x49, 0x91, 0xfa, 0x88, 0xcb, 0x8a, 0x76, 0x0e, 0x6e, 0x08, 0x3f, 0xe5, 0x84, 0xa2, 0x93, 0x1c,
	0x1e, 0x3f, 0x0c, 0x9c, 0xc0, 0xaf, 0xbd, 0x54, 0xa0, 0xd3, 0x0c, 0x66, 0x46, 0x4a, 0x73, 0xb0,
	0x53, 0x40, 0xfb, 0x61, 0x43, 0x67, 0x07, 0xb8, 0x47, 0xec, 0xbb, 0x1b, 0xa9, 0xcc, 0x7d, 0x33,
	0x32, 0xf7, 0x39, 0x76, 0xd0, 0x7c, 0xfc, 0xbd, 0xea, 0x4b, 0x59, 0xe9, 0xe9, 0x4c, 0xe8, 0x3f,
	0xa2, 0xc6, 0x8b, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x81, 0x0b, 0x24, 0xa3, 0x21, 0x05, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BotServiceClient is the client API for BotService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BotServiceClient interface {
	OnboardWorker(ctx context.Context, in *OnboardWorkerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AlertNewShift(ctx context.Context, in *AlertNewShiftRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AlertNewShifts(ctx context.Context, in *AlertNewShiftsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AlertRemovedShift(ctx context.Context, in *AlertRemovedShiftRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AlertRemovedShifts(ctx context.Context, in *AlertRemovedShiftsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AlertChangedShift(ctx context.Context, in *AlertChangedShiftRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type botServiceClient struct {
	cc *grpc.ClientConn
}

func NewBotServiceClient(cc *grpc.ClientConn) BotServiceClient {
	return &botServiceClient{cc}
}

func (c *botServiceClient) OnboardWorker(ctx context.Context, in *OnboardWorkerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/staffjoy.bot.BotService/OnboardWorker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServiceClient) AlertNewShift(ctx context.Context, in *AlertNewShiftRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/staffjoy.bot.BotService/AlertNewShift", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServiceClient) AlertNewShifts(ctx context.Context, in *AlertNewShiftsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/staffjoy.bot.BotService/AlertNewShifts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServiceClient) AlertRemovedShift(ctx context.Context, in *AlertRemovedShiftRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/staffjoy.bot.BotService/AlertRemovedShift", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServiceClient) AlertRemovedShifts(ctx context.Context, in *AlertRemovedShiftsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/staffjoy.bot.BotService/AlertRemovedShifts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServiceClient) AlertChangedShift(ctx context.Context, in *AlertChangedShiftRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/staffjoy.bot.BotService/AlertChangedShift", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BotServiceServer is the server API for BotService service.
type BotServiceServer interface {
	OnboardWorker(context.Context, *OnboardWorkerRequest) (*emptypb.Empty, error)
	AlertNewShift(context.Context, *AlertNewShiftRequest) (*emptypb.Empty, error)
	AlertNewShifts(context.Context, *AlertNewShiftsRequest) (*emptypb.Empty, error)
	AlertRemovedShift(context.Context, *AlertRemovedShiftRequest) (*emptypb.Empty, error)
	AlertRemovedShifts(context.Context, *AlertRemovedShiftsRequest) (*emptypb.Empty, error)
	AlertChangedShift(context.Context, *AlertChangedShiftRequest) (*emptypb.Empty, error)
}

// UnimplementedBotServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBotServiceServer struct {
}

func (*UnimplementedBotServiceServer) OnboardWorker(ctx context.Context, req *OnboardWorkerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnboardWorker not implemented")
}
func (*UnimplementedBotServiceServer) AlertNewShift(ctx context.Context, req *AlertNewShiftRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlertNewShift not implemented")
}
func (*UnimplementedBotServiceServer) AlertNewShifts(ctx context.Context, req *AlertNewShiftsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlertNewShifts not implemented")
}
func (*UnimplementedBotServiceServer) AlertRemovedShift(ctx context.Context, req *AlertRemovedShiftRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlertRemovedShift not implemented")
}
func (*UnimplementedBotServiceServer) AlertRemovedShifts(ctx context.Context, req *AlertRemovedShiftsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlertRemovedShifts not implemented")
}
func (*UnimplementedBotServiceServer) AlertChangedShift(ctx context.Context, req *AlertChangedShiftRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlertChangedShift not implemented")
}

func RegisterBotServiceServer(s *grpc.Server, srv BotServiceServer) {
	s.RegisterService(&_BotService_serviceDesc, srv)
}

func _BotService_OnboardWorker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnboardWorkerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServiceServer).OnboardWorker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/staffjoy.bot.BotService/OnboardWorker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServiceServer).OnboardWorker(ctx, req.(*OnboardWorkerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotService_AlertNewShift_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertNewShiftRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServiceServer).AlertNewShift(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/staffjoy.bot.BotService/AlertNewShift",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServiceServer).AlertNewShift(ctx, req.(*AlertNewShiftRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotService_AlertNewShifts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertNewShiftsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServiceServer).AlertNewShifts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/staffjoy.bot.BotService/AlertNewShifts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServiceServer).AlertNewShifts(ctx, req.(*AlertNewShiftsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotService_AlertRemovedShift_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertRemovedShiftRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServiceServer).AlertRemovedShift(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/staffjoy.bot.BotService/AlertRemovedShift",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServiceServer).AlertRemovedShift(ctx, req.(*AlertRemovedShiftRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotService_AlertRemovedShifts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertRemovedShiftsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServiceServer).AlertRemovedShifts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/staffjoy.bot.BotService/AlertRemovedShifts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServiceServer).AlertRemovedShifts(ctx, req.(*AlertRemovedShiftsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotService_AlertChangedShift_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertChangedShiftRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServiceServer).AlertChangedShift(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/staffjoy.bot.BotService/AlertChangedShift",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServiceServer).AlertChangedShift(ctx, req.(*AlertChangedShiftRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BotService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "staffjoy.bot.BotService",
	HandlerType: (*BotServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OnboardWorker",
			Handler:    _BotService_OnboardWorker_Handler,
		},
		{
			MethodName: "AlertNewShift",
			Handler:    _BotService_AlertNewShift_Handler,
		},
		{
			MethodName: "AlertNewShifts",
			Handler:    _BotService_AlertNewShifts_Handler,
		},
		{
			MethodName: "AlertRemovedShift",
			Handler:    _BotService_AlertRemovedShift_Handler,
		},
		{
			MethodName: "AlertRemovedShifts",
			Handler:    _BotService_AlertRemovedShifts_Handler,
		},
		{
			MethodName: "AlertChangedShift",
			Handler:    _BotService_AlertChangedShift_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bot.proto",
}