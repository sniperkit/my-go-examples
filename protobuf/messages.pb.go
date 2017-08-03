// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	messages.proto

It has these top-level messages:
	Token
*/
package main

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Token struct {
	AccessToken  string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	TokenType    string `protobuf:"bytes,2,opt,name=token_type,json=tokenType" json:"token_type,omitempty"`
	RefreshToken string `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken" json:"refresh_token,omitempty"`
	ExpiresAt    int64  `protobuf:"varint,4,opt,name=expires_at,json=expiresAt" json:"expires_at,omitempty"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (m *Token) String() string            { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Token) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *Token) GetTokenType() string {
	if m != nil {
		return m.TokenType
	}
	return ""
}

func (m *Token) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *Token) GetExpiresAt() int64 {
	if m != nil {
		return m.ExpiresAt
	}
	return 0
}

func init() {
	proto.RegisterType((*Token)(nil), "main.Token")
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc,
	0x53, 0xea, 0x65, 0xe4, 0x62, 0x0d, 0xc9, 0xcf, 0x4e, 0xcd, 0x13, 0x52, 0xe4, 0xe2, 0x49, 0x4c,
	0x4e, 0x4e, 0x2d, 0x2e, 0x8e, 0x2f, 0x01, 0xf1, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xb8,
	0x21, 0x62, 0x10, 0x25, 0xb2, 0x5c, 0x5c, 0x60, 0xb9, 0xf8, 0x92, 0xca, 0x82, 0x54, 0x09, 0x26,
	0xb0, 0x02, 0x4e, 0xb0, 0x48, 0x48, 0x65, 0x41, 0xaa, 0x90, 0x32, 0x17, 0x6f, 0x51, 0x6a, 0x5a,
	0x51, 0x6a, 0x71, 0x06, 0xd4, 0x08, 0x66, 0xb0, 0x0a, 0x1e, 0xa8, 0x20, 0xdc, 0x8c, 0xd4, 0x8a,
	0x82, 0xcc, 0xa2, 0xd4, 0xe2, 0xf8, 0xc4, 0x12, 0x09, 0x16, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x4e,
	0xa8, 0x88, 0x63, 0x49, 0x12, 0x1b, 0xd8, 0x71, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xef,
	0x6c, 0x01, 0xd0, 0xae, 0x00, 0x00, 0x00,
}
