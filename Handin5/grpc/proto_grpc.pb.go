// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: grpc/proto.proto

package proto

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

const (
	Auction_Bid_FullMethodName        = "/handin5.auction/bid"
	Auction_Result_FullMethodName     = "/handin5.auction/result"
	Auction_Election_FullMethodName   = "/handin5.auction/election"
	Auction_Victory_FullMethodName    = "/handin5.auction/victory"
	Auction_DoElection_FullMethodName = "/handin5.auction/doElection"
)

// AuctionClient is the client API for Auction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuctionClient interface {
	Bid(ctx context.Context, in *BidRequest, opts ...grpc.CallOption) (*BidAck, error)
	Result(ctx context.Context, in *ResultRequest, opts ...grpc.CallOption) (*OutcomeResponse, error)
	Election(ctx context.Context, in *ElectionRequest, opts ...grpc.CallOption) (*Alive, error)
	Victory(ctx context.Context, in *VictoryMessage, opts ...grpc.CallOption) (*Alive, error)
	DoElection(ctx context.Context, in *ElectionWarning, opts ...grpc.CallOption) (*Empty, error)
}

type auctionClient struct {
	cc grpc.ClientConnInterface
}

func NewAuctionClient(cc grpc.ClientConnInterface) AuctionClient {
	return &auctionClient{cc}
}

func (c *auctionClient) Bid(ctx context.Context, in *BidRequest, opts ...grpc.CallOption) (*BidAck, error) {
	out := new(BidAck)
	err := c.cc.Invoke(ctx, Auction_Bid_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionClient) Result(ctx context.Context, in *ResultRequest, opts ...grpc.CallOption) (*OutcomeResponse, error) {
	out := new(OutcomeResponse)
	err := c.cc.Invoke(ctx, Auction_Result_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionClient) Election(ctx context.Context, in *ElectionRequest, opts ...grpc.CallOption) (*Alive, error) {
	out := new(Alive)
	err := c.cc.Invoke(ctx, Auction_Election_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionClient) Victory(ctx context.Context, in *VictoryMessage, opts ...grpc.CallOption) (*Alive, error) {
	out := new(Alive)
	err := c.cc.Invoke(ctx, Auction_Victory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionClient) DoElection(ctx context.Context, in *ElectionWarning, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Auction_DoElection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuctionServer is the server API for Auction service.
// All implementations must embed UnimplementedAuctionServer
// for forward compatibility
type AuctionServer interface {
	Bid(context.Context, *BidRequest) (*BidAck, error)
	Result(context.Context, *ResultRequest) (*OutcomeResponse, error)
	Election(context.Context, *ElectionRequest) (*Alive, error)
	Victory(context.Context, *VictoryMessage) (*Alive, error)
	DoElection(context.Context, *ElectionWarning) (*Empty, error)
	mustEmbedUnimplementedAuctionServer()
}

// UnimplementedAuctionServer must be embedded to have forward compatible implementations.
type UnimplementedAuctionServer struct {
}

func (UnimplementedAuctionServer) Bid(context.Context, *BidRequest) (*BidAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bid not implemented")
}
func (UnimplementedAuctionServer) Result(context.Context, *ResultRequest) (*OutcomeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Result not implemented")
}
func (UnimplementedAuctionServer) Election(context.Context, *ElectionRequest) (*Alive, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Election not implemented")
}
func (UnimplementedAuctionServer) Victory(context.Context, *VictoryMessage) (*Alive, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Victory not implemented")
}
func (UnimplementedAuctionServer) DoElection(context.Context, *ElectionWarning) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoElection not implemented")
}
func (UnimplementedAuctionServer) mustEmbedUnimplementedAuctionServer() {}

// UnsafeAuctionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuctionServer will
// result in compilation errors.
type UnsafeAuctionServer interface {
	mustEmbedUnimplementedAuctionServer()
}

func RegisterAuctionServer(s grpc.ServiceRegistrar, srv AuctionServer) {
	s.RegisterService(&Auction_ServiceDesc, srv)
}

func _Auction_Bid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServer).Bid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auction_Bid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServer).Bid(ctx, req.(*BidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auction_Result_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServer).Result(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auction_Result_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServer).Result(ctx, req.(*ResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auction_Election_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ElectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServer).Election(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auction_Election_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServer).Election(ctx, req.(*ElectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auction_Victory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VictoryMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServer).Victory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auction_Victory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServer).Victory(ctx, req.(*VictoryMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auction_DoElection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ElectionWarning)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServer).DoElection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auction_DoElection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServer).DoElection(ctx, req.(*ElectionWarning))
	}
	return interceptor(ctx, in, info, handler)
}

// Auction_ServiceDesc is the grpc.ServiceDesc for Auction service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auction_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "handin5.auction",
	HandlerType: (*AuctionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "bid",
			Handler:    _Auction_Bid_Handler,
		},
		{
			MethodName: "result",
			Handler:    _Auction_Result_Handler,
		},
		{
			MethodName: "election",
			Handler:    _Auction_Election_Handler,
		},
		{
			MethodName: "victory",
			Handler:    _Auction_Victory_Handler,
		},
		{
			MethodName: "doElection",
			Handler:    _Auction_DoElection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/proto.proto",
}