// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// If you're modifying this file, please follow the protobuf style guide:
//   https://protobuf.dev/programming-guides/style/
// and also the Google API design guide
//   https://cloud.google.com/apis/design/
// also see the comments in the http grpc source file:
//   https://github.com/googleapis/googleapis/blob/master/google/api/http.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: api.proto

package v2

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	OpenMatchService_CreateTicket_FullMethodName               = "/open_match.v2.OpenMatchService/CreateTicket"
	OpenMatchService_DeactivateTickets_FullMethodName          = "/open_match.v2.OpenMatchService/DeactivateTickets"
	OpenMatchService_ActivateTickets_FullMethodName            = "/open_match.v2.OpenMatchService/ActivateTickets"
	OpenMatchService_InvokeMatchmakingFunctions_FullMethodName = "/open_match.v2.OpenMatchService/InvokeMatchmakingFunctions"
	OpenMatchService_CreateAssignments_FullMethodName          = "/open_match.v2.OpenMatchService/CreateAssignments"
	OpenMatchService_WatchAssignments_FullMethodName           = "/open_match.v2.OpenMatchService/WatchAssignments"
)

// OpenMatchServiceClient is the client API for OpenMatchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OpenMatchServiceClient interface {
	// CreateTicket puts an immutable ticket into state storage, and returns it's Ticket Id.
	// Tickets will be actively expired after the configured OM_TICKET_TTL_SECS has passed.
	// Tickets are placed in the 'inactive' state when created (they will not show up in
	//
	//	pools sent to your matchmaking functions).  Use the ActivateTickets() RPC to move
	//	them to the 'active' state.
	CreateTicket(ctx context.Context, in *CreateTicketRequest, opts ...grpc.CallOption) (*CreateTicketResponse, error)
	// Deactivate tickets takes a list of ticket ids which it will move to the 'inactive'
	// state and returns the (estimated) completion time, after which the tickets will no
	// longer appear in pools sent to new matchmaking function invocations.
	DeactivateTickets(ctx context.Context, in *DeactivateTicketsRequest, opts ...grpc.CallOption) (*DeactivateTicketsResponse, error)
	// Activate tickets takes a list of ticket ids which it will move to the 'active'
	// state and returns the (estimated) completion time, after which the ticket will
	// appear in pools sent to new matchmaking function invocations.
	ActivateTickets(ctx context.Context, in *ActivateTicketsRequest, opts ...grpc.CallOption) (*ActivateTicketsResponse, error)
	// InvokeMatchmakingFunctions is the core of open match. As input, it receives:
	// - A Match Profile, consisting of:
	//   - A list of empty ticket pools, with filters defining how to find all
	//     players that belong to that pool.
	//   - A list of empty rosters, representing groups of tickets to assign to the resulting
	//     match. This is an optional construct but often used to represent teams (see
	//     the example matchmaking functions and the docs for more details)
	//
	// - A list of matchmaking function (aka MMF) endpoints
	// The RPC first looks at the pools in the request profile, and evaluates all the filters
	// in each to fill the pools with eligible tickets. Once all pools in the profile are
	// filled with all eligible tickets, a copy of the profile is sent to each matchmaking
	// function endpoint specified in the request. MMF calls are defined by a separate gRPC
	// service defined in proto/v2/mmf.proto, please reference it for details as to the input
	// and output of matchmaking functions.
	// All matches returned from all MMFs before the context deadline are then sent back
	// to OM's InvokeMatchmakingFunctions handler, which does the following:
	//   - Reads all rosters of all matches returned, and moves every ticket it finds in those
	//     rosters to the 'inactive' state.
	//   - Records statistics/telemetry for the results.
	//   - Sends those results out the output stream as a StreamedMmfResponse
	//
	// Note: cancelling the context (by, for example, closing the connection) does NOT
	//
	//	cancel the invoked MMFs, by design. They will run, make matches, and exit
	//	when they finish.
	InvokeMatchmakingFunctions(ctx context.Context, in *MmfRequest, opts ...grpc.CallOption) (OpenMatchService_InvokeMatchmakingFunctionsClient, error)
	// CreateAssignments creates an assignment for each ticket in the request's AssignmentRoster
	// field. Assignments are only guaranteed to exist until the ticket expires, although they MAY
	// continue to exist afterwords. This RPC is considered deprecated and should not be used in production.
	CreateAssignments(ctx context.Context, in *CreateAssignmentsRequest, opts ...grpc.CallOption) (*CreateAssignmentsResponse, error)
	// WatchAssignments streams back one assignment of each ticketID requested, if it exists before the
	// timeout is reached. This RPC is considered deprecated and should not be used in production.
	WatchAssignments(ctx context.Context, in *WatchAssignmentsRequest, opts ...grpc.CallOption) (OpenMatchService_WatchAssignmentsClient, error)
}

type openMatchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOpenMatchServiceClient(cc grpc.ClientConnInterface) OpenMatchServiceClient {
	return &openMatchServiceClient{cc}
}

func (c *openMatchServiceClient) CreateTicket(ctx context.Context, in *CreateTicketRequest, opts ...grpc.CallOption) (*CreateTicketResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTicketResponse)
	err := c.cc.Invoke(ctx, OpenMatchService_CreateTicket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openMatchServiceClient) DeactivateTickets(ctx context.Context, in *DeactivateTicketsRequest, opts ...grpc.CallOption) (*DeactivateTicketsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeactivateTicketsResponse)
	err := c.cc.Invoke(ctx, OpenMatchService_DeactivateTickets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openMatchServiceClient) ActivateTickets(ctx context.Context, in *ActivateTicketsRequest, opts ...grpc.CallOption) (*ActivateTicketsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ActivateTicketsResponse)
	err := c.cc.Invoke(ctx, OpenMatchService_ActivateTickets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openMatchServiceClient) InvokeMatchmakingFunctions(ctx context.Context, in *MmfRequest, opts ...grpc.CallOption) (OpenMatchService_InvokeMatchmakingFunctionsClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &OpenMatchService_ServiceDesc.Streams[0], OpenMatchService_InvokeMatchmakingFunctions_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &openMatchServiceInvokeMatchmakingFunctionsClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OpenMatchService_InvokeMatchmakingFunctionsClient interface {
	Recv() (*StreamedMmfResponse, error)
	grpc.ClientStream
}

type openMatchServiceInvokeMatchmakingFunctionsClient struct {
	grpc.ClientStream
}

func (x *openMatchServiceInvokeMatchmakingFunctionsClient) Recv() (*StreamedMmfResponse, error) {
	m := new(StreamedMmfResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *openMatchServiceClient) CreateAssignments(ctx context.Context, in *CreateAssignmentsRequest, opts ...grpc.CallOption) (*CreateAssignmentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateAssignmentsResponse)
	err := c.cc.Invoke(ctx, OpenMatchService_CreateAssignments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openMatchServiceClient) WatchAssignments(ctx context.Context, in *WatchAssignmentsRequest, opts ...grpc.CallOption) (OpenMatchService_WatchAssignmentsClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &OpenMatchService_ServiceDesc.Streams[1], OpenMatchService_WatchAssignments_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &openMatchServiceWatchAssignmentsClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OpenMatchService_WatchAssignmentsClient interface {
	Recv() (*StreamedWatchAssignmentsResponse, error)
	grpc.ClientStream
}

type openMatchServiceWatchAssignmentsClient struct {
	grpc.ClientStream
}

func (x *openMatchServiceWatchAssignmentsClient) Recv() (*StreamedWatchAssignmentsResponse, error) {
	m := new(StreamedWatchAssignmentsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OpenMatchServiceServer is the server API for OpenMatchService service.
// All implementations must embed UnimplementedOpenMatchServiceServer
// for forward compatibility
type OpenMatchServiceServer interface {
	// CreateTicket puts an immutable ticket into state storage, and returns it's Ticket Id.
	// Tickets will be actively expired after the configured OM_TICKET_TTL_SECS has passed.
	// Tickets are placed in the 'inactive' state when created (they will not show up in
	//
	//	pools sent to your matchmaking functions).  Use the ActivateTickets() RPC to move
	//	them to the 'active' state.
	CreateTicket(context.Context, *CreateTicketRequest) (*CreateTicketResponse, error)
	// Deactivate tickets takes a list of ticket ids which it will move to the 'inactive'
	// state and returns the (estimated) completion time, after which the tickets will no
	// longer appear in pools sent to new matchmaking function invocations.
	DeactivateTickets(context.Context, *DeactivateTicketsRequest) (*DeactivateTicketsResponse, error)
	// Activate tickets takes a list of ticket ids which it will move to the 'active'
	// state and returns the (estimated) completion time, after which the ticket will
	// appear in pools sent to new matchmaking function invocations.
	ActivateTickets(context.Context, *ActivateTicketsRequest) (*ActivateTicketsResponse, error)
	// InvokeMatchmakingFunctions is the core of open match. As input, it receives:
	// - A Match Profile, consisting of:
	//   - A list of empty ticket pools, with filters defining how to find all
	//     players that belong to that pool.
	//   - A list of empty rosters, representing groups of tickets to assign to the resulting
	//     match. This is an optional construct but often used to represent teams (see
	//     the example matchmaking functions and the docs for more details)
	//
	// - A list of matchmaking function (aka MMF) endpoints
	// The RPC first looks at the pools in the request profile, and evaluates all the filters
	// in each to fill the pools with eligible tickets. Once all pools in the profile are
	// filled with all eligible tickets, a copy of the profile is sent to each matchmaking
	// function endpoint specified in the request. MMF calls are defined by a separate gRPC
	// service defined in proto/v2/mmf.proto, please reference it for details as to the input
	// and output of matchmaking functions.
	// All matches returned from all MMFs before the context deadline are then sent back
	// to OM's InvokeMatchmakingFunctions handler, which does the following:
	//   - Reads all rosters of all matches returned, and moves every ticket it finds in those
	//     rosters to the 'inactive' state.
	//   - Records statistics/telemetry for the results.
	//   - Sends those results out the output stream as a StreamedMmfResponse
	//
	// Note: cancelling the context (by, for example, closing the connection) does NOT
	//
	//	cancel the invoked MMFs, by design. They will run, make matches, and exit
	//	when they finish.
	InvokeMatchmakingFunctions(*MmfRequest, OpenMatchService_InvokeMatchmakingFunctionsServer) error
	// CreateAssignments creates an assignment for each ticket in the request's AssignmentRoster
	// field. Assignments are only guaranteed to exist until the ticket expires, although they MAY
	// continue to exist afterwords. This RPC is considered deprecated and should not be used in production.
	CreateAssignments(context.Context, *CreateAssignmentsRequest) (*CreateAssignmentsResponse, error)
	// WatchAssignments streams back one assignment of each ticketID requested, if it exists before the
	// timeout is reached. This RPC is considered deprecated and should not be used in production.
	WatchAssignments(*WatchAssignmentsRequest, OpenMatchService_WatchAssignmentsServer) error
	mustEmbedUnimplementedOpenMatchServiceServer()
}

// UnimplementedOpenMatchServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOpenMatchServiceServer struct {
}

func (UnimplementedOpenMatchServiceServer) CreateTicket(context.Context, *CreateTicketRequest) (*CreateTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTicket not implemented")
}
func (UnimplementedOpenMatchServiceServer) DeactivateTickets(context.Context, *DeactivateTicketsRequest) (*DeactivateTicketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeactivateTickets not implemented")
}
func (UnimplementedOpenMatchServiceServer) ActivateTickets(context.Context, *ActivateTicketsRequest) (*ActivateTicketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateTickets not implemented")
}
func (UnimplementedOpenMatchServiceServer) InvokeMatchmakingFunctions(*MmfRequest, OpenMatchService_InvokeMatchmakingFunctionsServer) error {
	return status.Errorf(codes.Unimplemented, "method InvokeMatchmakingFunctions not implemented")
}
func (UnimplementedOpenMatchServiceServer) CreateAssignments(context.Context, *CreateAssignmentsRequest) (*CreateAssignmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAssignments not implemented")
}
func (UnimplementedOpenMatchServiceServer) WatchAssignments(*WatchAssignmentsRequest, OpenMatchService_WatchAssignmentsServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchAssignments not implemented")
}
func (UnimplementedOpenMatchServiceServer) mustEmbedUnimplementedOpenMatchServiceServer() {}

// UnsafeOpenMatchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OpenMatchServiceServer will
// result in compilation errors.
type UnsafeOpenMatchServiceServer interface {
	mustEmbedUnimplementedOpenMatchServiceServer()
}

func RegisterOpenMatchServiceServer(s grpc.ServiceRegistrar, srv OpenMatchServiceServer) {
	s.RegisterService(&OpenMatchService_ServiceDesc, srv)
}

func _OpenMatchService_CreateTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenMatchServiceServer).CreateTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenMatchService_CreateTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenMatchServiceServer).CreateTicket(ctx, req.(*CreateTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpenMatchService_DeactivateTickets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeactivateTicketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenMatchServiceServer).DeactivateTickets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenMatchService_DeactivateTickets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenMatchServiceServer).DeactivateTickets(ctx, req.(*DeactivateTicketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpenMatchService_ActivateTickets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateTicketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenMatchServiceServer).ActivateTickets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenMatchService_ActivateTickets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenMatchServiceServer).ActivateTickets(ctx, req.(*ActivateTicketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpenMatchService_InvokeMatchmakingFunctions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MmfRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OpenMatchServiceServer).InvokeMatchmakingFunctions(m, &openMatchServiceInvokeMatchmakingFunctionsServer{ServerStream: stream})
}

type OpenMatchService_InvokeMatchmakingFunctionsServer interface {
	Send(*StreamedMmfResponse) error
	grpc.ServerStream
}

type openMatchServiceInvokeMatchmakingFunctionsServer struct {
	grpc.ServerStream
}

func (x *openMatchServiceInvokeMatchmakingFunctionsServer) Send(m *StreamedMmfResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _OpenMatchService_CreateAssignments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAssignmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenMatchServiceServer).CreateAssignments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OpenMatchService_CreateAssignments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenMatchServiceServer).CreateAssignments(ctx, req.(*CreateAssignmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpenMatchService_WatchAssignments_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchAssignmentsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OpenMatchServiceServer).WatchAssignments(m, &openMatchServiceWatchAssignmentsServer{ServerStream: stream})
}

type OpenMatchService_WatchAssignmentsServer interface {
	Send(*StreamedWatchAssignmentsResponse) error
	grpc.ServerStream
}

type openMatchServiceWatchAssignmentsServer struct {
	grpc.ServerStream
}

func (x *openMatchServiceWatchAssignmentsServer) Send(m *StreamedWatchAssignmentsResponse) error {
	return x.ServerStream.SendMsg(m)
}

// OpenMatchService_ServiceDesc is the grpc.ServiceDesc for OpenMatchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OpenMatchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "open_match.v2.OpenMatchService",
	HandlerType: (*OpenMatchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTicket",
			Handler:    _OpenMatchService_CreateTicket_Handler,
		},
		{
			MethodName: "DeactivateTickets",
			Handler:    _OpenMatchService_DeactivateTickets_Handler,
		},
		{
			MethodName: "ActivateTickets",
			Handler:    _OpenMatchService_ActivateTickets_Handler,
		},
		{
			MethodName: "CreateAssignments",
			Handler:    _OpenMatchService_CreateAssignments_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "InvokeMatchmakingFunctions",
			Handler:       _OpenMatchService_InvokeMatchmakingFunctions_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "WatchAssignments",
			Handler:       _OpenMatchService_WatchAssignments_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api.proto",
}
