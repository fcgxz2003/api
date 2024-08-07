// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: pkg/apis/dfdaemon/v1/dfdaemon.proto

package dfdaemon

import (
	context "context"
	v1 "github.com/fcgxz2003/api/v2/pkg/apis/common/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DaemonClient is the client API for Daemon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DaemonClient interface {
	// Trigger client to download file
	Download(ctx context.Context, in *DownRequest, opts ...grpc.CallOption) (Daemon_DownloadClient, error)
	// Get piece tasks from other peers
	GetPieceTasks(ctx context.Context, in *v1.PieceTaskRequest, opts ...grpc.CallOption) (*v1.PiecePacket, error)
	// Check daemon health
	CheckHealth(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Sync piece tasks with other peers
	SyncPieceTasks(ctx context.Context, opts ...grpc.CallOption) (Daemon_SyncPieceTasksClient, error)
	// Check if given task exists in P2P cache system
	StatTask(ctx context.Context, in *StatTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Import the given file into P2P cache system
	ImportTask(ctx context.Context, in *ImportTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Export or download file from P2P cache system
	ExportTask(ctx context.Context, in *ExportTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Delete file from P2P cache system
	DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// LeaveHost releases host in scheduler.
	LeaveHost(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Exchange peers between daemons
	PeerExchange(ctx context.Context, opts ...grpc.CallOption) (Daemon_PeerExchangeClient, error)
}

type daemonClient struct {
	cc grpc.ClientConnInterface
}

func NewDaemonClient(cc grpc.ClientConnInterface) DaemonClient {
	return &daemonClient{cc}
}

func (c *daemonClient) Download(ctx context.Context, in *DownRequest, opts ...grpc.CallOption) (Daemon_DownloadClient, error) {
	stream, err := c.cc.NewStream(ctx, &Daemon_ServiceDesc.Streams[0], "/dfdaemon.Daemon/Download", opts...)
	if err != nil {
		return nil, err
	}
	x := &daemonDownloadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Daemon_DownloadClient interface {
	Recv() (*DownResult, error)
	grpc.ClientStream
}

type daemonDownloadClient struct {
	grpc.ClientStream
}

func (x *daemonDownloadClient) Recv() (*DownResult, error) {
	m := new(DownResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *daemonClient) GetPieceTasks(ctx context.Context, in *v1.PieceTaskRequest, opts ...grpc.CallOption) (*v1.PiecePacket, error) {
	out := new(v1.PiecePacket)
	err := c.cc.Invoke(ctx, "/dfdaemon.Daemon/GetPieceTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) CheckHealth(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dfdaemon.Daemon/CheckHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) SyncPieceTasks(ctx context.Context, opts ...grpc.CallOption) (Daemon_SyncPieceTasksClient, error) {
	stream, err := c.cc.NewStream(ctx, &Daemon_ServiceDesc.Streams[1], "/dfdaemon.Daemon/SyncPieceTasks", opts...)
	if err != nil {
		return nil, err
	}
	x := &daemonSyncPieceTasksClient{stream}
	return x, nil
}

type Daemon_SyncPieceTasksClient interface {
	Send(*v1.PieceTaskRequest) error
	Recv() (*v1.PiecePacket, error)
	grpc.ClientStream
}

type daemonSyncPieceTasksClient struct {
	grpc.ClientStream
}

func (x *daemonSyncPieceTasksClient) Send(m *v1.PieceTaskRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *daemonSyncPieceTasksClient) Recv() (*v1.PiecePacket, error) {
	m := new(v1.PiecePacket)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *daemonClient) StatTask(ctx context.Context, in *StatTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dfdaemon.Daemon/StatTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) ImportTask(ctx context.Context, in *ImportTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dfdaemon.Daemon/ImportTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) ExportTask(ctx context.Context, in *ExportTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dfdaemon.Daemon/ExportTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dfdaemon.Daemon/DeleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) LeaveHost(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dfdaemon.Daemon/LeaveHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) PeerExchange(ctx context.Context, opts ...grpc.CallOption) (Daemon_PeerExchangeClient, error) {
	stream, err := c.cc.NewStream(ctx, &Daemon_ServiceDesc.Streams[2], "/dfdaemon.Daemon/PeerExchange", opts...)
	if err != nil {
		return nil, err
	}
	x := &daemonPeerExchangeClient{stream}
	return x, nil
}

type Daemon_PeerExchangeClient interface {
	Send(*PeerExchangeData) error
	Recv() (*PeerExchangeData, error)
	grpc.ClientStream
}

type daemonPeerExchangeClient struct {
	grpc.ClientStream
}

func (x *daemonPeerExchangeClient) Send(m *PeerExchangeData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *daemonPeerExchangeClient) Recv() (*PeerExchangeData, error) {
	m := new(PeerExchangeData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DaemonServer is the server API for Daemon service.
// All implementations should embed UnimplementedDaemonServer
// for forward compatibility
type DaemonServer interface {
	// Trigger client to download file
	Download(*DownRequest, Daemon_DownloadServer) error
	// Get piece tasks from other peers
	GetPieceTasks(context.Context, *v1.PieceTaskRequest) (*v1.PiecePacket, error)
	// Check daemon health
	CheckHealth(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	// Sync piece tasks with other peers
	SyncPieceTasks(Daemon_SyncPieceTasksServer) error
	// Check if given task exists in P2P cache system
	StatTask(context.Context, *StatTaskRequest) (*emptypb.Empty, error)
	// Import the given file into P2P cache system
	ImportTask(context.Context, *ImportTaskRequest) (*emptypb.Empty, error)
	// Export or download file from P2P cache system
	ExportTask(context.Context, *ExportTaskRequest) (*emptypb.Empty, error)
	// Delete file from P2P cache system
	DeleteTask(context.Context, *DeleteTaskRequest) (*emptypb.Empty, error)
	// LeaveHost releases host in scheduler.
	LeaveHost(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	// Exchange peers between daemons
	PeerExchange(Daemon_PeerExchangeServer) error
}

// UnimplementedDaemonServer should be embedded to have forward compatible implementations.
type UnimplementedDaemonServer struct {
}

func (UnimplementedDaemonServer) Download(*DownRequest, Daemon_DownloadServer) error {
	return status.Errorf(codes.Unimplemented, "method Download not implemented")
}
func (UnimplementedDaemonServer) GetPieceTasks(context.Context, *v1.PieceTaskRequest) (*v1.PiecePacket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPieceTasks not implemented")
}
func (UnimplementedDaemonServer) CheckHealth(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckHealth not implemented")
}
func (UnimplementedDaemonServer) SyncPieceTasks(Daemon_SyncPieceTasksServer) error {
	return status.Errorf(codes.Unimplemented, "method SyncPieceTasks not implemented")
}
func (UnimplementedDaemonServer) StatTask(context.Context, *StatTaskRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatTask not implemented")
}
func (UnimplementedDaemonServer) ImportTask(context.Context, *ImportTaskRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportTask not implemented")
}
func (UnimplementedDaemonServer) ExportTask(context.Context, *ExportTaskRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportTask not implemented")
}
func (UnimplementedDaemonServer) DeleteTask(context.Context, *DeleteTaskRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
}
func (UnimplementedDaemonServer) LeaveHost(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveHost not implemented")
}
func (UnimplementedDaemonServer) PeerExchange(Daemon_PeerExchangeServer) error {
	return status.Errorf(codes.Unimplemented, "method PeerExchange not implemented")
}

// UnsafeDaemonServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DaemonServer will
// result in compilation errors.
type UnsafeDaemonServer interface {
	mustEmbedUnimplementedDaemonServer()
}

func RegisterDaemonServer(s grpc.ServiceRegistrar, srv DaemonServer) {
	s.RegisterService(&Daemon_ServiceDesc, srv)
}

func _Daemon_Download_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DaemonServer).Download(m, &daemonDownloadServer{stream})
}

type Daemon_DownloadServer interface {
	Send(*DownResult) error
	grpc.ServerStream
}

type daemonDownloadServer struct {
	grpc.ServerStream
}

func (x *daemonDownloadServer) Send(m *DownResult) error {
	return x.ServerStream.SendMsg(m)
}

func _Daemon_GetPieceTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.PieceTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).GetPieceTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfdaemon.Daemon/GetPieceTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).GetPieceTasks(ctx, req.(*v1.PieceTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_CheckHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).CheckHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfdaemon.Daemon/CheckHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).CheckHealth(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_SyncPieceTasks_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DaemonServer).SyncPieceTasks(&daemonSyncPieceTasksServer{stream})
}

type Daemon_SyncPieceTasksServer interface {
	Send(*v1.PiecePacket) error
	Recv() (*v1.PieceTaskRequest, error)
	grpc.ServerStream
}

type daemonSyncPieceTasksServer struct {
	grpc.ServerStream
}

func (x *daemonSyncPieceTasksServer) Send(m *v1.PiecePacket) error {
	return x.ServerStream.SendMsg(m)
}

func (x *daemonSyncPieceTasksServer) Recv() (*v1.PieceTaskRequest, error) {
	m := new(v1.PieceTaskRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Daemon_StatTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).StatTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfdaemon.Daemon/StatTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).StatTask(ctx, req.(*StatTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_ImportTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).ImportTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfdaemon.Daemon/ImportTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).ImportTask(ctx, req.(*ImportTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_ExportTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).ExportTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfdaemon.Daemon/ExportTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).ExportTask(ctx, req.(*ExportTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfdaemon.Daemon/DeleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).DeleteTask(ctx, req.(*DeleteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_LeaveHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).LeaveHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfdaemon.Daemon/LeaveHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).LeaveHost(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_PeerExchange_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DaemonServer).PeerExchange(&daemonPeerExchangeServer{stream})
}

type Daemon_PeerExchangeServer interface {
	Send(*PeerExchangeData) error
	Recv() (*PeerExchangeData, error)
	grpc.ServerStream
}

type daemonPeerExchangeServer struct {
	grpc.ServerStream
}

func (x *daemonPeerExchangeServer) Send(m *PeerExchangeData) error {
	return x.ServerStream.SendMsg(m)
}

func (x *daemonPeerExchangeServer) Recv() (*PeerExchangeData, error) {
	m := new(PeerExchangeData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Daemon_ServiceDesc is the grpc.ServiceDesc for Daemon service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Daemon_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dfdaemon.Daemon",
	HandlerType: (*DaemonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPieceTasks",
			Handler:    _Daemon_GetPieceTasks_Handler,
		},
		{
			MethodName: "CheckHealth",
			Handler:    _Daemon_CheckHealth_Handler,
		},
		{
			MethodName: "StatTask",
			Handler:    _Daemon_StatTask_Handler,
		},
		{
			MethodName: "ImportTask",
			Handler:    _Daemon_ImportTask_Handler,
		},
		{
			MethodName: "ExportTask",
			Handler:    _Daemon_ExportTask_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _Daemon_DeleteTask_Handler,
		},
		{
			MethodName: "LeaveHost",
			Handler:    _Daemon_LeaveHost_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Download",
			Handler:       _Daemon_Download_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SyncPieceTasks",
			Handler:       _Daemon_SyncPieceTasks_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "PeerExchange",
			Handler:       _Daemon_PeerExchange_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pkg/apis/dfdaemon/v1/dfdaemon.proto",
}
