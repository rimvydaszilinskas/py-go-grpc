// Code generated by protoc-gen-go. DO NOT EDIT.
// source: countries.proto

/*
Package countries is a generated protocol buffer package.

It is generated from these files:
	countries.proto

It has these top-level messages:
	CountryRequest
	Currencies
	CountryResponse
*/
package countries

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

type CountryRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *CountryRequest) Reset()                    { *m = CountryRequest{} }
func (m *CountryRequest) String() string            { return proto.CompactTextString(m) }
func (*CountryRequest) ProtoMessage()               {}
func (*CountryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CountryRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Currencies struct {
	Code   string `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Symbol string `protobuf:"bytes,3,opt,name=symbol" json:"symbol,omitempty"`
}

func (m *Currencies) Reset()                    { *m = Currencies{} }
func (m *Currencies) String() string            { return proto.CompactTextString(m) }
func (*Currencies) ProtoMessage()               {}
func (*Currencies) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Currencies) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Currencies) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Currencies) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

type CountryResponse struct {
	Name       string        `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Alpha2Code string        `protobuf:"bytes,2,opt,name=alpha2Code" json:"alpha2Code,omitempty"`
	Capital    string        `protobuf:"bytes,3,opt,name=capital" json:"capital,omitempty"`
	Subregion  string        `protobuf:"bytes,4,opt,name=subregion" json:"subregion,omitempty"`
	Population int32         `protobuf:"varint,5,opt,name=population" json:"population,omitempty"`
	NativeName string        `protobuf:"bytes,6,opt,name=nativeName" json:"nativeName,omitempty"`
	Currencies []*Currencies `protobuf:"bytes,7,rep,name=currencies" json:"currencies,omitempty"`
}

func (m *CountryResponse) Reset()                    { *m = CountryResponse{} }
func (m *CountryResponse) String() string            { return proto.CompactTextString(m) }
func (*CountryResponse) ProtoMessage()               {}
func (*CountryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CountryResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CountryResponse) GetAlpha2Code() string {
	if m != nil {
		return m.Alpha2Code
	}
	return ""
}

func (m *CountryResponse) GetCapital() string {
	if m != nil {
		return m.Capital
	}
	return ""
}

func (m *CountryResponse) GetSubregion() string {
	if m != nil {
		return m.Subregion
	}
	return ""
}

func (m *CountryResponse) GetPopulation() int32 {
	if m != nil {
		return m.Population
	}
	return 0
}

func (m *CountryResponse) GetNativeName() string {
	if m != nil {
		return m.NativeName
	}
	return ""
}

func (m *CountryResponse) GetCurrencies() []*Currencies {
	if m != nil {
		return m.Currencies
	}
	return nil
}

func init() {
	proto.RegisterType((*CountryRequest)(nil), "countries.CountryRequest")
	proto.RegisterType((*Currencies)(nil), "countries.Currencies")
	proto.RegisterType((*CountryResponse)(nil), "countries.CountryResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Country service

type CountryClient interface {
	Search(ctx context.Context, in *CountryRequest, opts ...grpc.CallOption) (*CountryResponse, error)
}

type countryClient struct {
	cc *grpc.ClientConn
}

func NewCountryClient(cc *grpc.ClientConn) CountryClient {
	return &countryClient{cc}
}

func (c *countryClient) Search(ctx context.Context, in *CountryRequest, opts ...grpc.CallOption) (*CountryResponse, error) {
	out := new(CountryResponse)
	err := grpc.Invoke(ctx, "/countries.Country/Search", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Country service

type CountryServer interface {
	Search(context.Context, *CountryRequest) (*CountryResponse, error)
}

func RegisterCountryServer(s *grpc.Server, srv CountryServer) {
	s.RegisterService(&_Country_serviceDesc, srv)
}

func _Country_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountryServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/countries.Country/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountryServer).Search(ctx, req.(*CountryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Country_serviceDesc = grpc.ServiceDesc{
	ServiceName: "countries.Country",
	HandlerType: (*CountryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _Country_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "countries.proto",
}

func init() { proto.RegisterFile("countries.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x25, 0xb6, 0x4d, 0xc8, 0x08, 0x16, 0x06, 0x94, 0x58, 0x44, 0x42, 0xf0, 0x90, 0x53, 0x0f,
	0x11, 0xcf, 0x1e, 0x72, 0x13, 0xf1, 0x10, 0xbf, 0x60, 0xb3, 0x0e, 0x36, 0x90, 0xee, 0xae, 0xbb,
	0x1b, 0xa1, 0xff, 0xed, 0x07, 0xc8, 0x26, 0xed, 0x6e, 0x0a, 0xbd, 0xcd, 0xbc, 0x37, 0x6f, 0xde,
	0xee, 0x1b, 0x58, 0x73, 0x39, 0x08, 0xab, 0x3b, 0x32, 0x5b, 0xa5, 0xa5, 0x95, 0x98, 0x7a, 0xa0,
	0x78, 0x82, 0x9b, 0x7a, 0x6c, 0x0e, 0x0d, 0xfd, 0x0c, 0x64, 0x2c, 0x22, 0x2c, 0x05, 0xdb, 0x53,
	0x16, 0xe5, 0x51, 0x99, 0x36, 0x63, 0x5d, 0xbc, 0x03, 0xd4, 0x83, 0xd6, 0x24, 0x78, 0x47, 0xc6,
	0x4d, 0x70, 0xf9, 0xe5, 0x27, 0x5c, 0xed, 0x55, 0x57, 0x41, 0x85, 0x77, 0x10, 0x9b, 0xc3, 0xbe,
	0x95, 0x7d, 0xb6, 0x18, 0xd1, 0x63, 0x57, 0xfc, 0x45, 0xb0, 0xf6, 0xa6, 0x46, 0x49, 0x61, 0xe8,
	0x92, 0x2b, 0x3e, 0x02, 0xb0, 0x5e, 0xed, 0x58, 0x55, 0x3b, 0xb7, 0x69, 0xf3, 0x0c, 0xc1, 0x0c,
	0x12, 0xce, 0x54, 0x67, 0xd9, 0xc9, 0xe0, 0xd4, 0xe2, 0x03, 0xa4, 0x66, 0x68, 0x35, 0x7d, 0x77,
	0x52, 0x64, 0xcb, 0x91, 0x0b, 0x80, 0xdb, 0xab, 0xa4, 0x1a, 0x7a, 0x66, 0x1d, 0xbd, 0xca, 0xa3,
	0x72, 0xd5, 0xcc, 0x10, 0xc7, 0x0b, 0x66, 0xbb, 0x5f, 0xfa, 0x70, 0x2f, 0x8a, 0x27, 0xdf, 0x80,
	0xe0, 0x0b, 0x00, 0xf7, 0x69, 0x64, 0x49, 0xbe, 0x28, 0xaf, 0xab, 0xdb, 0x6d, 0x08, 0x39, 0x44,
	0xd5, 0xcc, 0x06, 0xab, 0x37, 0x48, 0x8e, 0xbf, 0xc6, 0x57, 0x88, 0x3f, 0x89, 0x69, 0xbe, 0xc3,
	0xfb, 0xb9, 0xee, 0xec, 0x10, 0x9b, 0xcd, 0x25, 0x6a, 0x8a, 0xab, 0x8d, 0xc7, 0x43, 0x3e, 0xff,
	0x07, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x5b, 0x0a, 0xd6, 0xdb, 0x01, 0x00, 0x00,
}
