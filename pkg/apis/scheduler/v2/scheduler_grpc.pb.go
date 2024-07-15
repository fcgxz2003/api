// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: pkg/apis/scheduler/v2/scheduler.proto

package scheduler

import (
	context "context"
	v2 "github.com/fcgxz2003/api/v2/pkg/apis/common/v2"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SchedulerClient is the client API for Scheduler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SchedulerClient interface {
	// AnnouncePeer announces peer to scheduler.
	AnnouncePeer(ctx context.Context, opts ...grpc.CallOption) (Scheduler_AnnouncePeerClient, error)
	// Checks information of peer.
	StatPeer(ctx context.Context, in *StatPeerRequest, opts ...grpc.CallOption) (*v2.Peer, error)
	// DeletePeer releases peer in scheduler.
	DeletePeer(ctx context.Context, in *DeletePeerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Checks information of task.
	StatTask(ctx context.Context, in *StatTaskRequest, opts ...grpc.CallOption) (*v2.Task, error)
	// DeleteTask releases task in scheduler.
	DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// AnnounceHost announces host to scheduler.
	AnnounceHost(ctx context.Context, in *AnnounceHostRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// DeleteHost releases host in scheduler.
	DeleteHost(ctx context.Context, in *DeleteHostRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// SyncProbes sync probes of the host.
	SyncProbes(ctx context.Context, opts ...grpc.CallOption) (Scheduler_SyncProbesClient, error)
	// AnnounceCachePeer announces cache peer to scheduler.
	AnnounceCachePeer(ctx context.Context, opts ...grpc.CallOption) (Scheduler_AnnounceCachePeerClient, error)
	// Checks information of cache peer.
	StatCachePeer(ctx context.Context, in *StatCachePeerRequest, opts ...grpc.CallOption) (*v2.CachePeer, error)
	// DeleteCachePeer releases cache peer in scheduler.
	DeleteCachePeer(ctx context.Context, in *DeleteCachePeerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// UploadCacheTaskStarted uploads cache task started to scheduler.
	UploadCacheTaskStarted(ctx context.Context, in *UploadCacheTaskStartedRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// UploadCacheTaskFinished uploads cache task finished to scheduler.
	UploadCacheTaskFinished(ctx context.Context, in *UploadCacheTaskFinishedRequest, opts ...grpc.CallOption) (*v2.CacheTask, error)
	// UploadCacheTaskFailed uploads cache task failed to scheduler.
	UploadCacheTaskFailed(ctx context.Context, in *UploadCacheTaskFailedRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Checks information of cache task.
	StatCacheTask(ctx context.Context, in *StatCacheTaskRequest, opts ...grpc.CallOption) (*v2.CacheTask, error)
	// DeleteCacheTask releases cache task in scheduler.
	DeleteCacheTask(ctx context.Context, in *DeleteCacheTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type schedulerClient struct {
	cc grpc.ClientConnInterface
}

func NewSchedulerClient(cc grpc.ClientConnInterface) SchedulerClient {
	return &schedulerClient{cc}
}

func (c *schedulerClient) AnnouncePeer(ctx context.Context, opts ...grpc.CallOption) (Scheduler_AnnouncePeerClient, error) {
	stream, err := c.cc.NewStream(ctx, &Scheduler_ServiceDesc.Streams[0], "/scheduler.v2.Scheduler/AnnouncePeer", opts...)
	if err != nil {
		return nil, err
	}
	x := &schedulerAnnouncePeerClient{stream}
	return x, nil
}

type Scheduler_AnnouncePeerClient interface {
	Send(*AnnouncePeerRequest) error
	Recv() (*AnnouncePeerResponse, error)
	grpc.ClientStream
}

type schedulerAnnouncePeerClient struct {
	grpc.ClientStream
}

func (x *schedulerAnnouncePeerClient) Send(m *AnnouncePeerRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *schedulerAnnouncePeerClient) Recv() (*AnnouncePeerResponse, error) {
	m := new(AnnouncePeerResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *schedulerClient) StatPeer(ctx context.Context, in *StatPeerRequest, opts ...grpc.CallOption) (*v2.Peer, error) {
	out := new(v2.Peer)
	err := c.cc.Invoke(ctx, "/scheduler.v2.Scheduler/StatPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) DeletePeer(ctx context.Context, in *DeletePeerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/scheduler.v2.Scheduler/DeletePeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) StatTask(ctx context.Context, in *StatTaskRequest, opts ...grpc.CallOption) (*v2.Task, error) {
	out := new(v2.Task)
	err := c.cc.Invoke(ctx, "/scheduler.v2.Scheduler/StatTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/scheduler.v2.Scheduler/DeleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) AnnounceHost(ctx context.Context, in *AnnounceHostRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/scheduler.v2.Scheduler/AnnounceHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) DeleteHost(ctx context.Context, in *DeleteHostRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/scheduler.v2.Scheduler/DeleteHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) SyncProbes(ctx context.Context, opts ...grpc.CallOption) (Scheduler_SyncProbesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Scheduler_ServiceDesc.Streams[1], "/scheduler.v2.Scheduler/SyncProbes", opts...)
	if err != nil {
		return nil, err
	}
	x := &schedulerSyncProbesClient{stream}
	return x, nil
}

type Scheduler_SyncProbesClient interface {
	Send(*SyncProbesRequest) error
	Recv() (*SyncProbesResponse, error)
	grpc.ClientStream
}

type schedulerSyncProbesClient struct {
	grpc.ClientStream
}

func (x *schedulerSyncProbesClient) Send(m *SyncProbesRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *schedulerSyncProbesClient) Recv() (*SyncProbesResponse, error) {
	m := new(SyncProbesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *schedulerClient) AnnounceCachePeer(ctx context.Context, opts ...grpc.CallOption) (Scheduler_AnnounceCachePeerClient, error) {
	stream, err := c.cc.NewStream(ctx, &Scheduler_ServiceDesc.Streams[2], "/scheduler.v2.Scheduler/AnnounceCachePeer", opts...)
	if err != nil {
		return nil, err
	}
	x := &schedulerAnnounceCachePeerClient{stream}
	return x, nil
}

type Scheduler_AnnounceCachePeerClient interface {
	Send(*AnnounceCachePeerRequest) error
	Recv() (*AnnounceCachePeerResponse, error)
	grpc.ClientStream
}

type schedulerAnnounceCachePeerClient struct {
	grpc.ClientStream
}

func (x *schedulerAnnounceCachePeerClient) Send(m *AnnounceCachePeerRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *schedulerAnnounceCachePeerClient) Recv() (*AnnounceCachePeerResponse, error) {
	m := new(AnnounceCachePeerResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *schedulerClient) StatCachePeer(ctx context.Context, in *StatCachePeerRequest, opts ...grpc.CallOption) (*v2.CachePeer, error) {
	out := new(v2.CachePeer)
	err := c.cc.Invoke(ctx, "/scheduler.v2.Scheduler/StatCachePeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) DeleteCachePeer(ctx context.Context, in *DeleteCachePeerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/scheduler.v2.Scheduler/DeleteCachePeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) UploadCacheTaskStarted(ctx context.Context, in *UploadCacheTaskStartedRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/scheduler.v2.Scheduler/UploadCacheTaskStarted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) UploadCacheTaskFinished(ctx context.Context, in *UploadCacheTaskFinishedRequest, opts ...grpc.CallOption) (*v2.CacheTask, error) {
	out := new(v2.CacheTask)
	err := c.cc.Invoke(ctx, "/scheduler.v2.Scheduler/UploadCacheTaskFinished", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) UploadCacheTaskFailed(ctx context.Context, in *UploadCacheTaskFailedRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/scheduler.v2.Scheduler/UploadCacheTaskFailed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) StatCacheTask(ctx context.Context, in *StatCacheTaskRequest, opts ...grpc.CallOption) (*v2.CacheTask, error) {
	out := new(v2.CacheTask)
	err := c.cc.Invoke(ctx, "/scheduler.v2.Scheduler/StatCacheTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) DeleteCacheTask(ctx context.Context, in *DeleteCacheTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/scheduler.v2.Scheduler/DeleteCacheTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchedulerServer is the server API for Scheduler service.
// All implementations should embed UnimplementedSchedulerServer
// for forward compatibility
type SchedulerServer interface {
	// AnnouncePeer announces peer to scheduler.
	AnnouncePeer(Scheduler_AnnouncePeerServer) error
	// Checks information of peer.
	StatPeer(context.Context, *StatPeerRequest) (*v2.Peer, error)
	// DeletePeer releases peer in scheduler.
	DeletePeer(context.Context, *DeletePeerRequest) (*emptypb.Empty, error)
	// Checks information of task.
	StatTask(context.Context, *StatTaskRequest) (*v2.Task, error)
	// DeleteTask releases task in scheduler.
	DeleteTask(context.Context, *DeleteTaskRequest) (*emptypb.Empty, error)
	// AnnounceHost announces host to scheduler.
	AnnounceHost(context.Context, *AnnounceHostRequest) (*emptypb.Empty, error)
	// DeleteHost releases host in scheduler.
	DeleteHost(context.Context, *DeleteHostRequest) (*emptypb.Empty, error)
	// SyncProbes sync probes of the host.
	SyncProbes(Scheduler_SyncProbesServer) error
	// AnnounceCachePeer announces cache peer to scheduler.
	AnnounceCachePeer(Scheduler_AnnounceCachePeerServer) error
	// Checks information of cache peer.
	StatCachePeer(context.Context, *StatCachePeerRequest) (*v2.CachePeer, error)
	// DeleteCachePeer releases cache peer in scheduler.
	DeleteCachePeer(context.Context, *DeleteCachePeerRequest) (*emptypb.Empty, error)
	// UploadCacheTaskStarted uploads cache task started to scheduler.
	UploadCacheTaskStarted(context.Context, *UploadCacheTaskStartedRequest) (*emptypb.Empty, error)
	// UploadCacheTaskFinished uploads cache task finished to scheduler.
	UploadCacheTaskFinished(context.Context, *UploadCacheTaskFinishedRequest) (*v2.CacheTask, error)
	// UploadCacheTaskFailed uploads cache task failed to scheduler.
	UploadCacheTaskFailed(context.Context, *UploadCacheTaskFailedRequest) (*emptypb.Empty, error)
	// Checks information of cache task.
	StatCacheTask(context.Context, *StatCacheTaskRequest) (*v2.CacheTask, error)
	// DeleteCacheTask releases cache task in scheduler.
	DeleteCacheTask(context.Context, *DeleteCacheTaskRequest) (*emptypb.Empty, error)
}

// UnimplementedSchedulerServer should be embedded to have forward compatible implementations.
type UnimplementedSchedulerServer struct {
}

func (UnimplementedSchedulerServer) AnnouncePeer(Scheduler_AnnouncePeerServer) error {
	return status.Errorf(codes.Unimplemented, "method AnnouncePeer not implemented")
}
func (UnimplementedSchedulerServer) StatPeer(context.Context, *StatPeerRequest) (*v2.Peer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatPeer not implemented")
}
func (UnimplementedSchedulerServer) DeletePeer(context.Context, *DeletePeerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePeer not implemented")
}
func (UnimplementedSchedulerServer) StatTask(context.Context, *StatTaskRequest) (*v2.Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatTask not implemented")
}
func (UnimplementedSchedulerServer) DeleteTask(context.Context, *DeleteTaskRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
}
func (UnimplementedSchedulerServer) AnnounceHost(context.Context, *AnnounceHostRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnnounceHost not implemented")
}
func (UnimplementedSchedulerServer) DeleteHost(context.Context, *DeleteHostRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHost not implemented")
}
func (UnimplementedSchedulerServer) SyncProbes(Scheduler_SyncProbesServer) error {
	return status.Errorf(codes.Unimplemented, "method SyncProbes not implemented")
}
func (UnimplementedSchedulerServer) AnnounceCachePeer(Scheduler_AnnounceCachePeerServer) error {
	return status.Errorf(codes.Unimplemented, "method AnnounceCachePeer not implemented")
}
func (UnimplementedSchedulerServer) StatCachePeer(context.Context, *StatCachePeerRequest) (*v2.CachePeer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatCachePeer not implemented")
}
func (UnimplementedSchedulerServer) DeleteCachePeer(context.Context, *DeleteCachePeerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCachePeer not implemented")
}
func (UnimplementedSchedulerServer) UploadCacheTaskStarted(context.Context, *UploadCacheTaskStartedRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadCacheTaskStarted not implemented")
}
func (UnimplementedSchedulerServer) UploadCacheTaskFinished(context.Context, *UploadCacheTaskFinishedRequest) (*v2.CacheTask, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadCacheTaskFinished not implemented")
}
func (UnimplementedSchedulerServer) UploadCacheTaskFailed(context.Context, *UploadCacheTaskFailedRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadCacheTaskFailed not implemented")
}
func (UnimplementedSchedulerServer) StatCacheTask(context.Context, *StatCacheTaskRequest) (*v2.CacheTask, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatCacheTask not implemented")
}
func (UnimplementedSchedulerServer) DeleteCacheTask(context.Context, *DeleteCacheTaskRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCacheTask not implemented")
}

// UnsafeSchedulerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SchedulerServer will
// result in compilation errors.
type UnsafeSchedulerServer interface {
	mustEmbedUnimplementedSchedulerServer()
}

func RegisterSchedulerServer(s grpc.ServiceRegistrar, srv SchedulerServer) {
	s.RegisterService(&Scheduler_ServiceDesc, srv)
}

func _Scheduler_AnnouncePeer_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SchedulerServer).AnnouncePeer(&schedulerAnnouncePeerServer{stream})
}

type Scheduler_AnnouncePeerServer interface {
	Send(*AnnouncePeerResponse) error
	Recv() (*AnnouncePeerRequest, error)
	grpc.ServerStream
}

type schedulerAnnouncePeerServer struct {
	grpc.ServerStream
}

func (x *schedulerAnnouncePeerServer) Send(m *AnnouncePeerResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *schedulerAnnouncePeerServer) Recv() (*AnnouncePeerRequest, error) {
	m := new(AnnouncePeerRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Scheduler_StatPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatPeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).StatPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.v2.Scheduler/StatPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).StatPeer(ctx, req.(*StatPeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_DeletePeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).DeletePeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.v2.Scheduler/DeletePeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).DeletePeer(ctx, req.(*DeletePeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_StatTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).StatTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.v2.Scheduler/StatTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).StatTask(ctx, req.(*StatTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.v2.Scheduler/DeleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).DeleteTask(ctx, req.(*DeleteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_AnnounceHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnnounceHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).AnnounceHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.v2.Scheduler/AnnounceHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).AnnounceHost(ctx, req.(*AnnounceHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_DeleteHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).DeleteHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.v2.Scheduler/DeleteHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).DeleteHost(ctx, req.(*DeleteHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_SyncProbes_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SchedulerServer).SyncProbes(&schedulerSyncProbesServer{stream})
}

type Scheduler_SyncProbesServer interface {
	Send(*SyncProbesResponse) error
	Recv() (*SyncProbesRequest, error)
	grpc.ServerStream
}

type schedulerSyncProbesServer struct {
	grpc.ServerStream
}

func (x *schedulerSyncProbesServer) Send(m *SyncProbesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *schedulerSyncProbesServer) Recv() (*SyncProbesRequest, error) {
	m := new(SyncProbesRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Scheduler_AnnounceCachePeer_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SchedulerServer).AnnounceCachePeer(&schedulerAnnounceCachePeerServer{stream})
}

type Scheduler_AnnounceCachePeerServer interface {
	Send(*AnnounceCachePeerResponse) error
	Recv() (*AnnounceCachePeerRequest, error)
	grpc.ServerStream
}

type schedulerAnnounceCachePeerServer struct {
	grpc.ServerStream
}

func (x *schedulerAnnounceCachePeerServer) Send(m *AnnounceCachePeerResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *schedulerAnnounceCachePeerServer) Recv() (*AnnounceCachePeerRequest, error) {
	m := new(AnnounceCachePeerRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Scheduler_StatCachePeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatCachePeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).StatCachePeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.v2.Scheduler/StatCachePeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).StatCachePeer(ctx, req.(*StatCachePeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_DeleteCachePeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCachePeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).DeleteCachePeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.v2.Scheduler/DeleteCachePeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).DeleteCachePeer(ctx, req.(*DeleteCachePeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_UploadCacheTaskStarted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadCacheTaskStartedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).UploadCacheTaskStarted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.v2.Scheduler/UploadCacheTaskStarted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).UploadCacheTaskStarted(ctx, req.(*UploadCacheTaskStartedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_UploadCacheTaskFinished_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadCacheTaskFinishedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).UploadCacheTaskFinished(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.v2.Scheduler/UploadCacheTaskFinished",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).UploadCacheTaskFinished(ctx, req.(*UploadCacheTaskFinishedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_UploadCacheTaskFailed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadCacheTaskFailedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).UploadCacheTaskFailed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.v2.Scheduler/UploadCacheTaskFailed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).UploadCacheTaskFailed(ctx, req.(*UploadCacheTaskFailedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_StatCacheTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatCacheTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).StatCacheTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.v2.Scheduler/StatCacheTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).StatCacheTask(ctx, req.(*StatCacheTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_DeleteCacheTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCacheTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).DeleteCacheTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.v2.Scheduler/DeleteCacheTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).DeleteCacheTask(ctx, req.(*DeleteCacheTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Scheduler_ServiceDesc is the grpc.ServiceDesc for Scheduler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Scheduler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "scheduler.v2.Scheduler",
	HandlerType: (*SchedulerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StatPeer",
			Handler:    _Scheduler_StatPeer_Handler,
		},
		{
			MethodName: "DeletePeer",
			Handler:    _Scheduler_DeletePeer_Handler,
		},
		{
			MethodName: "StatTask",
			Handler:    _Scheduler_StatTask_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _Scheduler_DeleteTask_Handler,
		},
		{
			MethodName: "AnnounceHost",
			Handler:    _Scheduler_AnnounceHost_Handler,
		},
		{
			MethodName: "DeleteHost",
			Handler:    _Scheduler_DeleteHost_Handler,
		},
		{
			MethodName: "StatCachePeer",
			Handler:    _Scheduler_StatCachePeer_Handler,
		},
		{
			MethodName: "DeleteCachePeer",
			Handler:    _Scheduler_DeleteCachePeer_Handler,
		},
		{
			MethodName: "UploadCacheTaskStarted",
			Handler:    _Scheduler_UploadCacheTaskStarted_Handler,
		},
		{
			MethodName: "UploadCacheTaskFinished",
			Handler:    _Scheduler_UploadCacheTaskFinished_Handler,
		},
		{
			MethodName: "UploadCacheTaskFailed",
			Handler:    _Scheduler_UploadCacheTaskFailed_Handler,
		},
		{
			MethodName: "StatCacheTask",
			Handler:    _Scheduler_StatCacheTask_Handler,
		},
		{
			MethodName: "DeleteCacheTask",
			Handler:    _Scheduler_DeleteCacheTask_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AnnouncePeer",
			Handler:       _Scheduler_AnnouncePeer_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SyncProbes",
			Handler:       _Scheduler_SyncProbes_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "AnnounceCachePeer",
			Handler:       _Scheduler_AnnounceCachePeer_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pkg/apis/scheduler/v2/scheduler.proto",
}
