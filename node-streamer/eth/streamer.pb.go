// Code generated by protoc-gen-go. DO NOT EDIT.
// source: streamer.proto

/*
Package eth is a generated protocol buffer package.

It is generated from these files:
	streamer.proto

It has these top-level messages:
	ETHTransaction
	BlockHeight
	MempoolToDelete
	WatchAddress
	MempoolRecord
	Empty
	RawTx
	AddressToResync
	UsersData
	AddressExtended
	ReplyInfo
*/
package eth

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ETHTransaction struct {
	UserID       string `protobuf:"bytes,1,opt,name=UserID" json:"UserID,omitempty"`
	WalletIndex  int32  `protobuf:"varint,2,opt,name=WalletIndex" json:"WalletIndex,omitempty"`
	AddressIndex int32  `protobuf:"varint,3,opt,name=AddressIndex" json:"AddressIndex,omitempty"`
	Hash         string `protobuf:"bytes,4,opt,name=Hash" json:"Hash,omitempty"`
	From         string `protobuf:"bytes,5,opt,name=From" json:"From,omitempty"`
	To           string `protobuf:"bytes,6,opt,name=To" json:"To,omitempty"`
	Amount       int64  `protobuf:"varint,7,opt,name=Amount" json:"Amount,omitempty"`
	GasPrice     int64  `protobuf:"varint,8,opt,name=GasPrice" json:"GasPrice,omitempty"`
	GasLimit     int64  `protobuf:"varint,9,opt,name=GasLimit" json:"GasLimit,omitempty"`
	Nonce        int32  `protobuf:"varint,10,opt,name=Nonce" json:"Nonce,omitempty"`
	Status       int32  `protobuf:"varint,11,opt,name=Status" json:"Status,omitempty"`
	BlockTime    int64  `protobuf:"varint,12,opt,name=BlockTime" json:"BlockTime,omitempty"`
	TxpoolTime   int64  `protobuf:"varint,13,opt,name=TxpoolTime" json:"TxpoolTime,omitempty"`
	BlockHeight  int64  `protobuf:"varint,14,opt,name=BlockHeight" json:"BlockHeight,omitempty"`
	Resync       bool   `protobuf:"varint,15,opt,name=Resync" json:"Resync,omitempty"`
}

func (m *ETHTransaction) Reset()                    { *m = ETHTransaction{} }
func (m *ETHTransaction) String() string            { return proto.CompactTextString(m) }
func (*ETHTransaction) ProtoMessage()               {}
func (*ETHTransaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ETHTransaction) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *ETHTransaction) GetWalletIndex() int32 {
	if m != nil {
		return m.WalletIndex
	}
	return 0
}

func (m *ETHTransaction) GetAddressIndex() int32 {
	if m != nil {
		return m.AddressIndex
	}
	return 0
}

func (m *ETHTransaction) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *ETHTransaction) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *ETHTransaction) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *ETHTransaction) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *ETHTransaction) GetGasPrice() int64 {
	if m != nil {
		return m.GasPrice
	}
	return 0
}

func (m *ETHTransaction) GetGasLimit() int64 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *ETHTransaction) GetNonce() int32 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *ETHTransaction) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ETHTransaction) GetBlockTime() int64 {
	if m != nil {
		return m.BlockTime
	}
	return 0
}

func (m *ETHTransaction) GetTxpoolTime() int64 {
	if m != nil {
		return m.TxpoolTime
	}
	return 0
}

func (m *ETHTransaction) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *ETHTransaction) GetResync() bool {
	if m != nil {
		return m.Resync
	}
	return false
}

type BlockHeight struct {
	Height int64 `protobuf:"varint,1,opt,name=height" json:"height,omitempty"`
}

func (m *BlockHeight) Reset()                    { *m = BlockHeight{} }
func (m *BlockHeight) String() string            { return proto.CompactTextString(m) }
func (*BlockHeight) ProtoMessage()               {}
func (*BlockHeight) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BlockHeight) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type MempoolToDelete struct {
	Hash string `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
}

func (m *MempoolToDelete) Reset()                    { *m = MempoolToDelete{} }
func (m *MempoolToDelete) String() string            { return proto.CompactTextString(m) }
func (*MempoolToDelete) ProtoMessage()               {}
func (*MempoolToDelete) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MempoolToDelete) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type WatchAddress struct {
	Address      string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	UserID       string `protobuf:"bytes,2,opt,name=userID" json:"userID,omitempty"`
	WalletIndex  int32  `protobuf:"varint,3,opt,name=WalletIndex" json:"WalletIndex,omitempty"`
	AddressIndex int32  `protobuf:"varint,4,opt,name=AddressIndex" json:"AddressIndex,omitempty"`
}

func (m *WatchAddress) Reset()                    { *m = WatchAddress{} }
func (m *WatchAddress) String() string            { return proto.CompactTextString(m) }
func (*WatchAddress) ProtoMessage()               {}
func (*WatchAddress) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *WatchAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *WatchAddress) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *WatchAddress) GetWalletIndex() int32 {
	if m != nil {
		return m.WalletIndex
	}
	return 0
}

func (m *WatchAddress) GetAddressIndex() int32 {
	if m != nil {
		return m.AddressIndex
	}
	return 0
}

type MempoolRecord struct {
	Category int32  `protobuf:"varint,1,opt,name=category" json:"category,omitempty"`
	HashTX   string `protobuf:"bytes,2,opt,name=hashTX" json:"hashTX,omitempty"`
}

func (m *MempoolRecord) Reset()                    { *m = MempoolRecord{} }
func (m *MempoolRecord) String() string            { return proto.CompactTextString(m) }
func (*MempoolRecord) ProtoMessage()               {}
func (*MempoolRecord) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *MempoolRecord) GetCategory() int32 {
	if m != nil {
		return m.Category
	}
	return 0
}

func (m *MempoolRecord) GetHashTX() string {
	if m != nil {
		return m.HashTX
	}
	return ""
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type RawTx struct {
	Transaction string `protobuf:"bytes,1,opt,name=transaction" json:"transaction,omitempty"`
}

func (m *RawTx) Reset()                    { *m = RawTx{} }
func (m *RawTx) String() string            { return proto.CompactTextString(m) }
func (*RawTx) ProtoMessage()               {}
func (*RawTx) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RawTx) GetTransaction() string {
	if m != nil {
		return m.Transaction
	}
	return ""
}

type AddressToResync struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
}

func (m *AddressToResync) Reset()                    { *m = AddressToResync{} }
func (m *AddressToResync) String() string            { return proto.CompactTextString(m) }
func (*AddressToResync) ProtoMessage()               {}
func (*AddressToResync) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *AddressToResync) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type UsersData struct {
	Map map[string]*AddressExtended `protobuf:"bytes,1,rep,name=map" json:"map,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *UsersData) Reset()                    { *m = UsersData{} }
func (m *UsersData) String() string            { return proto.CompactTextString(m) }
func (*UsersData) ProtoMessage()               {}
func (*UsersData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *UsersData) GetMap() map[string]*AddressExtended {
	if m != nil {
		return m.Map
	}
	return nil
}

type AddressExtended struct {
	UserID       string `protobuf:"bytes,1,opt,name=UserID" json:"UserID,omitempty"`
	WalletIndex  int32  `protobuf:"varint,2,opt,name=WalletIndex" json:"WalletIndex,omitempty"`
	AddressIndex int32  `protobuf:"varint,3,opt,name=AddressIndex" json:"AddressIndex,omitempty"`
}

func (m *AddressExtended) Reset()                    { *m = AddressExtended{} }
func (m *AddressExtended) String() string            { return proto.CompactTextString(m) }
func (*AddressExtended) ProtoMessage()               {}
func (*AddressExtended) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *AddressExtended) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *AddressExtended) GetWalletIndex() int32 {
	if m != nil {
		return m.WalletIndex
	}
	return 0
}

func (m *AddressExtended) GetAddressIndex() int32 {
	if m != nil {
		return m.AddressIndex
	}
	return 0
}

type ReplyInfo struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *ReplyInfo) Reset()                    { *m = ReplyInfo{} }
func (m *ReplyInfo) String() string            { return proto.CompactTextString(m) }
func (*ReplyInfo) ProtoMessage()               {}
func (*ReplyInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ReplyInfo) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*ETHTransaction)(nil), "eth.ETHTransaction")
	proto.RegisterType((*BlockHeight)(nil), "eth.BlockHeight")
	proto.RegisterType((*MempoolToDelete)(nil), "eth.MempoolToDelete")
	proto.RegisterType((*WatchAddress)(nil), "eth.WatchAddress")
	proto.RegisterType((*MempoolRecord)(nil), "eth.MempoolRecord")
	proto.RegisterType((*Empty)(nil), "eth.Empty")
	proto.RegisterType((*RawTx)(nil), "eth.RawTx")
	proto.RegisterType((*AddressToResync)(nil), "eth.AddressToResync")
	proto.RegisterType((*UsersData)(nil), "eth.UsersData")
	proto.RegisterType((*AddressExtended)(nil), "eth.AddressExtended")
	proto.RegisterType((*ReplyInfo)(nil), "eth.ReplyInfo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NodeCommuunications service

type NodeCommuunicationsClient interface {
	EventInitialAdd(ctx context.Context, in *UsersData, opts ...grpc.CallOption) (*ReplyInfo, error)
	EventAddNewAddress(ctx context.Context, in *WatchAddress, opts ...grpc.CallOption) (*ReplyInfo, error)
	EventGetBlockHeight(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BlockHeight, error)
	EventGetAllMempool(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeCommuunications_EventGetAllMempoolClient, error)
	EventAddMempoolRecord(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeCommuunications_EventAddMempoolRecordClient, error)
	EventDeleteMempool(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeCommuunications_EventDeleteMempoolClient, error)
	EventResyncAddress(ctx context.Context, in *AddressToResync, opts ...grpc.CallOption) (*ReplyInfo, error)
	EventSendRawTx(ctx context.Context, in *RawTx, opts ...grpc.CallOption) (*ReplyInfo, error)
	NewTx(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeCommuunications_NewTxClient, error)
}

type nodeCommuunicationsClient struct {
	cc *grpc.ClientConn
}

func NewNodeCommuunicationsClient(cc *grpc.ClientConn) NodeCommuunicationsClient {
	return &nodeCommuunicationsClient{cc}
}

func (c *nodeCommuunicationsClient) EventInitialAdd(ctx context.Context, in *UsersData, opts ...grpc.CallOption) (*ReplyInfo, error) {
	out := new(ReplyInfo)
	err := grpc.Invoke(ctx, "/eth.NodeCommuunications/EventInitialAdd", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeCommuunicationsClient) EventAddNewAddress(ctx context.Context, in *WatchAddress, opts ...grpc.CallOption) (*ReplyInfo, error) {
	out := new(ReplyInfo)
	err := grpc.Invoke(ctx, "/eth.NodeCommuunications/EventAddNewAddress", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeCommuunicationsClient) EventGetBlockHeight(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BlockHeight, error) {
	out := new(BlockHeight)
	err := grpc.Invoke(ctx, "/eth.NodeCommuunications/EventGetBlockHeight", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeCommuunicationsClient) EventGetAllMempool(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeCommuunications_EventGetAllMempoolClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_NodeCommuunications_serviceDesc.Streams[0], c.cc, "/eth.NodeCommuunications/EventGetAllMempool", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeCommuunicationsEventGetAllMempoolClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeCommuunications_EventGetAllMempoolClient interface {
	Recv() (*MempoolRecord, error)
	grpc.ClientStream
}

type nodeCommuunicationsEventGetAllMempoolClient struct {
	grpc.ClientStream
}

func (x *nodeCommuunicationsEventGetAllMempoolClient) Recv() (*MempoolRecord, error) {
	m := new(MempoolRecord)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeCommuunicationsClient) EventAddMempoolRecord(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeCommuunications_EventAddMempoolRecordClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_NodeCommuunications_serviceDesc.Streams[1], c.cc, "/eth.NodeCommuunications/EventAddMempoolRecord", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeCommuunicationsEventAddMempoolRecordClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeCommuunications_EventAddMempoolRecordClient interface {
	Recv() (*MempoolRecord, error)
	grpc.ClientStream
}

type nodeCommuunicationsEventAddMempoolRecordClient struct {
	grpc.ClientStream
}

func (x *nodeCommuunicationsEventAddMempoolRecordClient) Recv() (*MempoolRecord, error) {
	m := new(MempoolRecord)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeCommuunicationsClient) EventDeleteMempool(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeCommuunications_EventDeleteMempoolClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_NodeCommuunications_serviceDesc.Streams[2], c.cc, "/eth.NodeCommuunications/EventDeleteMempool", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeCommuunicationsEventDeleteMempoolClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeCommuunications_EventDeleteMempoolClient interface {
	Recv() (*MempoolToDelete, error)
	grpc.ClientStream
}

type nodeCommuunicationsEventDeleteMempoolClient struct {
	grpc.ClientStream
}

func (x *nodeCommuunicationsEventDeleteMempoolClient) Recv() (*MempoolToDelete, error) {
	m := new(MempoolToDelete)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeCommuunicationsClient) EventResyncAddress(ctx context.Context, in *AddressToResync, opts ...grpc.CallOption) (*ReplyInfo, error) {
	out := new(ReplyInfo)
	err := grpc.Invoke(ctx, "/eth.NodeCommuunications/EventResyncAddress", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeCommuunicationsClient) EventSendRawTx(ctx context.Context, in *RawTx, opts ...grpc.CallOption) (*ReplyInfo, error) {
	out := new(ReplyInfo)
	err := grpc.Invoke(ctx, "/eth.NodeCommuunications/EventSendRawTx", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeCommuunicationsClient) NewTx(ctx context.Context, in *Empty, opts ...grpc.CallOption) (NodeCommuunications_NewTxClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_NodeCommuunications_serviceDesc.Streams[3], c.cc, "/eth.NodeCommuunications/NewTx", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeCommuunicationsNewTxClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeCommuunications_NewTxClient interface {
	Recv() (*ETHTransaction, error)
	grpc.ClientStream
}

type nodeCommuunicationsNewTxClient struct {
	grpc.ClientStream
}

func (x *nodeCommuunicationsNewTxClient) Recv() (*ETHTransaction, error) {
	m := new(ETHTransaction)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for NodeCommuunications service

type NodeCommuunicationsServer interface {
	EventInitialAdd(context.Context, *UsersData) (*ReplyInfo, error)
	EventAddNewAddress(context.Context, *WatchAddress) (*ReplyInfo, error)
	EventGetBlockHeight(context.Context, *Empty) (*BlockHeight, error)
	EventGetAllMempool(*Empty, NodeCommuunications_EventGetAllMempoolServer) error
	EventAddMempoolRecord(*Empty, NodeCommuunications_EventAddMempoolRecordServer) error
	EventDeleteMempool(*Empty, NodeCommuunications_EventDeleteMempoolServer) error
	EventResyncAddress(context.Context, *AddressToResync) (*ReplyInfo, error)
	EventSendRawTx(context.Context, *RawTx) (*ReplyInfo, error)
	NewTx(*Empty, NodeCommuunications_NewTxServer) error
}

func RegisterNodeCommuunicationsServer(s *grpc.Server, srv NodeCommuunicationsServer) {
	s.RegisterService(&_NodeCommuunications_serviceDesc, srv)
}

func _NodeCommuunications_EventInitialAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeCommuunicationsServer).EventInitialAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eth.NodeCommuunications/EventInitialAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeCommuunicationsServer).EventInitialAdd(ctx, req.(*UsersData))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeCommuunications_EventAddNewAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WatchAddress)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeCommuunicationsServer).EventAddNewAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eth.NodeCommuunications/EventAddNewAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeCommuunicationsServer).EventAddNewAddress(ctx, req.(*WatchAddress))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeCommuunications_EventGetBlockHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeCommuunicationsServer).EventGetBlockHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eth.NodeCommuunications/EventGetBlockHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeCommuunicationsServer).EventGetBlockHeight(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeCommuunications_EventGetAllMempool_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeCommuunicationsServer).EventGetAllMempool(m, &nodeCommuunicationsEventGetAllMempoolServer{stream})
}

type NodeCommuunications_EventGetAllMempoolServer interface {
	Send(*MempoolRecord) error
	grpc.ServerStream
}

type nodeCommuunicationsEventGetAllMempoolServer struct {
	grpc.ServerStream
}

func (x *nodeCommuunicationsEventGetAllMempoolServer) Send(m *MempoolRecord) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeCommuunications_EventAddMempoolRecord_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeCommuunicationsServer).EventAddMempoolRecord(m, &nodeCommuunicationsEventAddMempoolRecordServer{stream})
}

type NodeCommuunications_EventAddMempoolRecordServer interface {
	Send(*MempoolRecord) error
	grpc.ServerStream
}

type nodeCommuunicationsEventAddMempoolRecordServer struct {
	grpc.ServerStream
}

func (x *nodeCommuunicationsEventAddMempoolRecordServer) Send(m *MempoolRecord) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeCommuunications_EventDeleteMempool_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeCommuunicationsServer).EventDeleteMempool(m, &nodeCommuunicationsEventDeleteMempoolServer{stream})
}

type NodeCommuunications_EventDeleteMempoolServer interface {
	Send(*MempoolToDelete) error
	grpc.ServerStream
}

type nodeCommuunicationsEventDeleteMempoolServer struct {
	grpc.ServerStream
}

func (x *nodeCommuunicationsEventDeleteMempoolServer) Send(m *MempoolToDelete) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeCommuunications_EventResyncAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressToResync)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeCommuunicationsServer).EventResyncAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eth.NodeCommuunications/EventResyncAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeCommuunicationsServer).EventResyncAddress(ctx, req.(*AddressToResync))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeCommuunications_EventSendRawTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawTx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeCommuunicationsServer).EventSendRawTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eth.NodeCommuunications/EventSendRawTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeCommuunicationsServer).EventSendRawTx(ctx, req.(*RawTx))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeCommuunications_NewTx_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeCommuunicationsServer).NewTx(m, &nodeCommuunicationsNewTxServer{stream})
}

type NodeCommuunications_NewTxServer interface {
	Send(*ETHTransaction) error
	grpc.ServerStream
}

type nodeCommuunicationsNewTxServer struct {
	grpc.ServerStream
}

func (x *nodeCommuunicationsNewTxServer) Send(m *ETHTransaction) error {
	return x.ServerStream.SendMsg(m)
}

var _NodeCommuunications_serviceDesc = grpc.ServiceDesc{
	ServiceName: "eth.NodeCommuunications",
	HandlerType: (*NodeCommuunicationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EventInitialAdd",
			Handler:    _NodeCommuunications_EventInitialAdd_Handler,
		},
		{
			MethodName: "EventAddNewAddress",
			Handler:    _NodeCommuunications_EventAddNewAddress_Handler,
		},
		{
			MethodName: "EventGetBlockHeight",
			Handler:    _NodeCommuunications_EventGetBlockHeight_Handler,
		},
		{
			MethodName: "EventResyncAddress",
			Handler:    _NodeCommuunications_EventResyncAddress_Handler,
		},
		{
			MethodName: "EventSendRawTx",
			Handler:    _NodeCommuunications_EventSendRawTx_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EventGetAllMempool",
			Handler:       _NodeCommuunications_EventGetAllMempool_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "EventAddMempoolRecord",
			Handler:       _NodeCommuunications_EventAddMempoolRecord_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "EventDeleteMempool",
			Handler:       _NodeCommuunications_EventDeleteMempool_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "NewTx",
			Handler:       _NodeCommuunications_NewTx_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "streamer.proto",
}

func init() { proto.RegisterFile("streamer.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 725 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x5d, 0x4f, 0xdb, 0x48,
	0x14, 0xc5, 0x71, 0x0c, 0xc9, 0x0d, 0x24, 0xec, 0xc0, 0xee, 0x5a, 0xd1, 0x6a, 0x15, 0x8d, 0x84,
	0x14, 0x76, 0x57, 0xd1, 0x0a, 0xb4, 0x12, 0x4b, 0xfb, 0x92, 0x42, 0x0a, 0x91, 0x20, 0xaa, 0x8c,
	0x2b, 0xfa, 0x3a, 0xb5, 0x6f, 0xb1, 0x85, 0xed, 0x89, 0xec, 0x09, 0x24, 0xef, 0x55, 0x5f, 0xfa,
	0x17, 0xfa, 0x63, 0xab, 0xf9, 0x70, 0x70, 0xda, 0x54, 0xed, 0x4b, 0xdf, 0xee, 0xb9, 0x1f, 0x73,
	0x8f, 0xcf, 0x9c, 0x91, 0xa1, 0x5d, 0x88, 0x1c, 0x59, 0x8a, 0xf9, 0x60, 0x9a, 0x73, 0xc1, 0x89,
	0x8d, 0x22, 0xa2, 0x9f, 0x6c, 0x68, 0x8f, 0xfc, 0x4b, 0x3f, 0x67, 0x59, 0xc1, 0x02, 0x11, 0xf3,
	0x8c, 0xfc, 0x06, 0x9b, 0xaf, 0x0b, 0xcc, 0xc7, 0xe7, 0xae, 0xd5, 0xb3, 0xfa, 0x4d, 0xcf, 0x20,
	0xd2, 0x83, 0xd6, 0x2d, 0x4b, 0x12, 0x14, 0xe3, 0x2c, 0xc4, 0xb9, 0x5b, 0xeb, 0x59, 0x7d, 0xc7,
	0xab, 0xa6, 0x08, 0x85, 0xed, 0x61, 0x18, 0xe6, 0x58, 0x14, 0xba, 0xc5, 0x56, 0x2d, 0x2b, 0x39,
	0x42, 0xa0, 0x7e, 0xc9, 0x8a, 0xc8, 0xad, 0xab, 0xb3, 0x55, 0x2c, 0x73, 0x2f, 0x73, 0x9e, 0xba,
	0x8e, 0xce, 0xc9, 0x98, 0xb4, 0xa1, 0xe6, 0x73, 0x77, 0x53, 0x65, 0x6a, 0x3e, 0x97, 0xac, 0x86,
	0x29, 0x9f, 0x65, 0xc2, 0xdd, 0xea, 0x59, 0x7d, 0xdb, 0x33, 0x88, 0x74, 0xa1, 0x71, 0xc1, 0x8a,
	0x57, 0x79, 0x1c, 0xa0, 0xdb, 0x50, 0x95, 0x25, 0x36, 0xb5, 0xab, 0x38, 0x8d, 0x85, 0xdb, 0x5c,
	0xd6, 0x14, 0x26, 0xfb, 0xe0, 0x4c, 0x78, 0x16, 0xa0, 0x0b, 0x8a, 0xa4, 0x06, 0x72, 0xcb, 0x8d,
	0x60, 0x62, 0x56, 0xb8, 0x2d, 0x95, 0x36, 0x88, 0xfc, 0x01, 0xcd, 0x17, 0x09, 0x0f, 0xee, 0xfd,
	0x38, 0x45, 0x77, 0x5b, 0x1d, 0xf5, 0x94, 0x20, 0x7f, 0x02, 0xf8, 0xf3, 0x29, 0xe7, 0x89, 0x2a,
	0xef, 0xa8, 0x72, 0x25, 0x23, 0x95, 0x53, 0xcd, 0x97, 0x18, 0xdf, 0x45, 0xc2, 0x6d, 0xab, 0x86,
	0x6a, 0x4a, 0xee, 0xf5, 0xb0, 0x58, 0x64, 0x81, 0xdb, 0xe9, 0x59, 0xfd, 0x86, 0x67, 0x10, 0x3d,
	0x80, 0x2f, 0xdb, 0x22, 0x7d, 0x86, 0xa5, 0x45, 0xd0, 0x88, 0x1e, 0x40, 0xe7, 0x1a, 0x53, 0xb5,
	0x8f, 0x9f, 0x63, 0x82, 0x02, 0xa5, 0xa6, 0x91, 0xd4, 0x59, 0xdf, 0xa1, 0x8a, 0xe9, 0x07, 0x0b,
	0xb6, 0x6f, 0x99, 0x08, 0x22, 0x73, 0x23, 0xc4, 0x85, 0x2d, 0xa6, 0x43, 0xd3, 0x57, 0x42, 0xb9,
	0x69, 0xa6, 0x4d, 0x50, 0xd3, 0x26, 0x98, 0xad, 0x35, 0x81, 0xfd, 0x7d, 0x13, 0xd4, 0xbf, 0x36,
	0x01, 0x3d, 0x83, 0x1d, 0xc3, 0xd7, 0xc3, 0x80, 0xe7, 0xa1, 0xbc, 0xa9, 0x80, 0x09, 0xbc, 0xe3,
	0xf9, 0x42, 0x31, 0x71, 0xbc, 0x25, 0x56, 0x1f, 0xcd, 0x8a, 0xc8, 0x7f, 0x53, 0x52, 0xd1, 0x88,
	0x6e, 0x81, 0x33, 0x4a, 0xa7, 0x62, 0x41, 0x0f, 0xc1, 0xf1, 0xd8, 0xa3, 0x3f, 0x97, 0xe4, 0xc4,
	0x93, 0x91, 0xcd, 0x27, 0x55, 0x53, 0xf4, 0x6f, 0xe8, 0x18, 0x22, 0x3e, 0xd7, 0x12, 0x7f, 0x5b,
	0x03, 0xfa, 0xde, 0x82, 0xa6, 0xf4, 0x7e, 0x71, 0xce, 0x04, 0x23, 0x87, 0x60, 0xa7, 0x6c, 0xea,
	0x5a, 0x3d, 0xbb, 0xdf, 0x3a, 0xfa, 0x7d, 0x80, 0x22, 0x1a, 0x2c, 0x8b, 0x83, 0x6b, 0x36, 0x1d,
	0x65, 0x22, 0x5f, 0x78, 0xb2, 0xa7, 0x7b, 0x05, 0x8d, 0x32, 0x41, 0x76, 0xc1, 0xbe, 0xc7, 0x85,
	0x39, 0x5a, 0x86, 0xe4, 0x2f, 0x70, 0x1e, 0x58, 0x32, 0x43, 0xf5, 0x39, 0xad, 0xa3, 0x7d, 0x75,
	0x94, 0x61, 0x35, 0x9a, 0x0b, 0xcc, 0x42, 0x0c, 0x3d, 0xdd, 0x72, 0x5a, 0x3b, 0xb1, 0x28, 0x5f,
	0x72, 0x2e, 0xab, 0x3f, 0xf7, 0x89, 0xd2, 0x03, 0x68, 0x7a, 0x38, 0x4d, 0x16, 0xe3, 0xec, 0x1d,
	0x97, 0xf2, 0xa4, 0x58, 0x14, 0xec, 0x0e, 0x4b, 0x79, 0x0c, 0x3c, 0xfa, 0x58, 0x87, 0xbd, 0x09,
	0x0f, 0xf1, 0x8c, 0xa7, 0xe9, 0x6c, 0x96, 0xc5, 0x01, 0x93, 0x12, 0x17, 0xe4, 0x18, 0x3a, 0xa3,
	0x07, 0xcc, 0xc4, 0x38, 0x8b, 0x45, 0xcc, 0x92, 0x61, 0x18, 0x92, 0xf6, 0xaa, 0x5c, 0x5d, 0x8d,
	0x97, 0x4b, 0xe8, 0x06, 0xf9, 0x1f, 0x88, 0x1a, 0x1a, 0x86, 0xe1, 0x04, 0x1f, 0x4b, 0x7f, 0xfe,
	0xa2, 0xfa, 0xaa, 0x96, 0x5d, 0x33, 0xfa, 0x1f, 0xec, 0xa9, 0xd1, 0x0b, 0x14, 0xd5, 0xb7, 0x02,
	0xaa, 0x51, 0x39, 0xa4, 0xbb, 0xab, 0xe2, 0x4a, 0x95, 0x6e, 0x90, 0x13, 0xb3, 0xf1, 0x02, 0xc5,
	0x30, 0x49, 0x8c, 0x1d, 0x57, 0xa6, 0x88, 0x8a, 0x57, 0x8c, 0x4a, 0x37, 0xfe, 0xb5, 0xc8, 0x33,
	0xf8, 0xb5, 0xe4, 0xba, 0xea, 0xe2, 0x1f, 0x19, 0x3e, 0x35, 0x6b, 0xf5, 0x33, 0x5d, 0xb7, 0x76,
	0xbf, 0x3a, 0x59, 0xbe, 0x67, 0x35, 0xfb, 0xdc, 0xcc, 0x6a, 0xe7, 0x96, 0x22, 0xad, 0x18, 0xa8,
	0xb4, 0xf5, 0x1a, 0x9d, 0x06, 0xd0, 0x56, 0xd3, 0x37, 0x98, 0x85, 0xfa, 0xbd, 0xe8, 0xad, 0x2a,
	0x5e, 0xd3, 0xff, 0x0f, 0x38, 0x13, 0x7c, 0x6a, 0xd3, 0xe4, 0xf6, 0x74, 0xbc, 0xf2, 0xc7, 0x90,
	0xdc, 0xde, 0x6e, 0xaa, 0x9f, 0xca, 0xf1, 0xe7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x4f, 0xa8,
	0xae, 0x66, 0x06, 0x00, 0x00,
}
