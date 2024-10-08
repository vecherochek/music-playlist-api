// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: player.proto

package player_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PlayerV1Client is the client API for PlayerV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlayerV1Client interface {
	AddPlayer(ctx context.Context, in *AddPlayerRequest, opts ...grpc.CallOption) (*AddPlayerResponse, error)
	DeletePlayer(ctx context.Context, in *DeletePlayerRequest, opts ...grpc.CallOption) (*DeletePlayerResponse, error)
	AddSong(ctx context.Context, in *AddSongRequest, opts ...grpc.CallOption) (*AddSongResponse, error)
	DeleteSong(ctx context.Context, in *DeleteSongFromListRequest, opts ...grpc.CallOption) (*DeleteSongFromListResponse, error)
	Play(ctx context.Context, in *PlayRequest, opts ...grpc.CallOption) (PlayerV1_PlayClient, error)
	Pause(ctx context.Context, in *PauseRequest, opts ...grpc.CallOption) (*PauseResponse, error)
	Next(ctx context.Context, in *NextRequest, opts ...grpc.CallOption) (PlayerV1_NextClient, error)
	Prev(ctx context.Context, in *PrevRequest, opts ...grpc.CallOption) (PlayerV1_PrevClient, error)
}

type playerV1Client struct {
	cc grpc.ClientConnInterface
}

func NewPlayerV1Client(cc grpc.ClientConnInterface) PlayerV1Client {
	return &playerV1Client{cc}
}

func (c *playerV1Client) AddPlayer(ctx context.Context, in *AddPlayerRequest, opts ...grpc.CallOption) (*AddPlayerResponse, error) {
	out := new(AddPlayerResponse)
	err := c.cc.Invoke(ctx, "/music_player_v1.Player_v1/AddPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerV1Client) DeletePlayer(ctx context.Context, in *DeletePlayerRequest, opts ...grpc.CallOption) (*DeletePlayerResponse, error) {
	out := new(DeletePlayerResponse)
	err := c.cc.Invoke(ctx, "/music_player_v1.Player_v1/DeletePlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerV1Client) AddSong(ctx context.Context, in *AddSongRequest, opts ...grpc.CallOption) (*AddSongResponse, error) {
	out := new(AddSongResponse)
	err := c.cc.Invoke(ctx, "/music_player_v1.Player_v1/AddSong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerV1Client) DeleteSong(ctx context.Context, in *DeleteSongFromListRequest, opts ...grpc.CallOption) (*DeleteSongFromListResponse, error) {
	out := new(DeleteSongFromListResponse)
	err := c.cc.Invoke(ctx, "/music_player_v1.Player_v1/DeleteSong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerV1Client) Play(ctx context.Context, in *PlayRequest, opts ...grpc.CallOption) (PlayerV1_PlayClient, error) {
	stream, err := c.cc.NewStream(ctx, &PlayerV1_ServiceDesc.Streams[0], "/music_player_v1.Player_v1/Play", opts...)
	if err != nil {
		return nil, err
	}
	x := &playerV1PlayClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PlayerV1_PlayClient interface {
	Recv() (*StreamPlayResponse, error)
	grpc.ClientStream
}

type playerV1PlayClient struct {
	grpc.ClientStream
}

func (x *playerV1PlayClient) Recv() (*StreamPlayResponse, error) {
	m := new(StreamPlayResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *playerV1Client) Pause(ctx context.Context, in *PauseRequest, opts ...grpc.CallOption) (*PauseResponse, error) {
	out := new(PauseResponse)
	err := c.cc.Invoke(ctx, "/music_player_v1.Player_v1/Pause", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerV1Client) Next(ctx context.Context, in *NextRequest, opts ...grpc.CallOption) (PlayerV1_NextClient, error) {
	stream, err := c.cc.NewStream(ctx, &PlayerV1_ServiceDesc.Streams[1], "/music_player_v1.Player_v1/Next", opts...)
	if err != nil {
		return nil, err
	}
	x := &playerV1NextClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PlayerV1_NextClient interface {
	Recv() (*StreamNextResponse, error)
	grpc.ClientStream
}

type playerV1NextClient struct {
	grpc.ClientStream
}

func (x *playerV1NextClient) Recv() (*StreamNextResponse, error) {
	m := new(StreamNextResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *playerV1Client) Prev(ctx context.Context, in *PrevRequest, opts ...grpc.CallOption) (PlayerV1_PrevClient, error) {
	stream, err := c.cc.NewStream(ctx, &PlayerV1_ServiceDesc.Streams[2], "/music_player_v1.Player_v1/Prev", opts...)
	if err != nil {
		return nil, err
	}
	x := &playerV1PrevClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PlayerV1_PrevClient interface {
	Recv() (*StreamPrevResponse, error)
	grpc.ClientStream
}

type playerV1PrevClient struct {
	grpc.ClientStream
}

func (x *playerV1PrevClient) Recv() (*StreamPrevResponse, error) {
	m := new(StreamPrevResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PlayerV1Server is the server API for PlayerV1 service.
// All implementations must embed UnimplementedPlayerV1Server
// for forward compatibility
type PlayerV1Server interface {
	AddPlayer(context.Context, *AddPlayerRequest) (*AddPlayerResponse, error)
	DeletePlayer(context.Context, *DeletePlayerRequest) (*DeletePlayerResponse, error)
	AddSong(context.Context, *AddSongRequest) (*AddSongResponse, error)
	DeleteSong(context.Context, *DeleteSongFromListRequest) (*DeleteSongFromListResponse, error)
	Play(*PlayRequest, PlayerV1_PlayServer) error
	Pause(context.Context, *PauseRequest) (*PauseResponse, error)
	Next(*NextRequest, PlayerV1_NextServer) error
	Prev(*PrevRequest, PlayerV1_PrevServer) error
	mustEmbedUnimplementedPlayerV1Server()
}

// UnimplementedPlayerV1Server must be embedded to have forward compatible implementations.
type UnimplementedPlayerV1Server struct {
}

func (UnimplementedPlayerV1Server) AddPlayer(context.Context, *AddPlayerRequest) (*AddPlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPlayer not implemented")
}
func (UnimplementedPlayerV1Server) DeletePlayer(context.Context, *DeletePlayerRequest) (*DeletePlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePlayer not implemented")
}
func (UnimplementedPlayerV1Server) AddSong(context.Context, *AddSongRequest) (*AddSongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSong not implemented")
}
func (UnimplementedPlayerV1Server) DeleteSong(context.Context, *DeleteSongFromListRequest) (*DeleteSongFromListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSong not implemented")
}
func (UnimplementedPlayerV1Server) Play(*PlayRequest, PlayerV1_PlayServer) error {
	return status.Errorf(codes.Unimplemented, "method Play not implemented")
}
func (UnimplementedPlayerV1Server) Pause(context.Context, *PauseRequest) (*PauseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pause not implemented")
}
func (UnimplementedPlayerV1Server) Next(*NextRequest, PlayerV1_NextServer) error {
	return status.Errorf(codes.Unimplemented, "method Next not implemented")
}
func (UnimplementedPlayerV1Server) Prev(*PrevRequest, PlayerV1_PrevServer) error {
	return status.Errorf(codes.Unimplemented, "method Prev not implemented")
}
func (UnimplementedPlayerV1Server) mustEmbedUnimplementedPlayerV1Server() {}

// UnsafePlayerV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlayerV1Server will
// result in compilation errors.
type UnsafePlayerV1Server interface {
	mustEmbedUnimplementedPlayerV1Server()
}

func RegisterPlayerV1Server(s grpc.ServiceRegistrar, srv PlayerV1Server) {
	s.RegisterService(&PlayerV1_ServiceDesc, srv)
}

func _PlayerV1_AddPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerV1Server).AddPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/music_player_v1.Player_v1/AddPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerV1Server).AddPlayer(ctx, req.(*AddPlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerV1_DeletePlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerV1Server).DeletePlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/music_player_v1.Player_v1/DeletePlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerV1Server).DeletePlayer(ctx, req.(*DeletePlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerV1_AddSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerV1Server).AddSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/music_player_v1.Player_v1/AddSong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerV1Server).AddSong(ctx, req.(*AddSongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerV1_DeleteSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSongFromListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerV1Server).DeleteSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/music_player_v1.Player_v1/DeleteSong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerV1Server).DeleteSong(ctx, req.(*DeleteSongFromListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerV1_Play_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PlayRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PlayerV1Server).Play(m, &playerV1PlayServer{stream})
}

type PlayerV1_PlayServer interface {
	Send(*StreamPlayResponse) error
	grpc.ServerStream
}

type playerV1PlayServer struct {
	grpc.ServerStream
}

func (x *playerV1PlayServer) Send(m *StreamPlayResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _PlayerV1_Pause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PauseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerV1Server).Pause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/music_player_v1.Player_v1/Pause",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerV1Server).Pause(ctx, req.(*PauseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerV1_Next_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NextRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PlayerV1Server).Next(m, &playerV1NextServer{stream})
}

type PlayerV1_NextServer interface {
	Send(*StreamNextResponse) error
	grpc.ServerStream
}

type playerV1NextServer struct {
	grpc.ServerStream
}

func (x *playerV1NextServer) Send(m *StreamNextResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _PlayerV1_Prev_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrevRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PlayerV1Server).Prev(m, &playerV1PrevServer{stream})
}

type PlayerV1_PrevServer interface {
	Send(*StreamPrevResponse) error
	grpc.ServerStream
}

type playerV1PrevServer struct {
	grpc.ServerStream
}

func (x *playerV1PrevServer) Send(m *StreamPrevResponse) error {
	return x.ServerStream.SendMsg(m)
}

// PlayerV1_ServiceDesc is the grpc.ServiceDesc for PlayerV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PlayerV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "music_player_v1.Player_v1",
	HandlerType: (*PlayerV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPlayer",
			Handler:    _PlayerV1_AddPlayer_Handler,
		},
		{
			MethodName: "DeletePlayer",
			Handler:    _PlayerV1_DeletePlayer_Handler,
		},
		{
			MethodName: "AddSong",
			Handler:    _PlayerV1_AddSong_Handler,
		},
		{
			MethodName: "DeleteSong",
			Handler:    _PlayerV1_DeleteSong_Handler,
		},
		{
			MethodName: "Pause",
			Handler:    _PlayerV1_Pause_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Play",
			Handler:       _PlayerV1_Play_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Next",
			Handler:       _PlayerV1_Next_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Prev",
			Handler:       _PlayerV1_Prev_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "player.proto",
}
