package pmysql

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecodePacket(t *testing.T) {

	// test case 通过wireshark 抓包获取
	var testCase = []struct {
		raw       []byte
		expect    packetType
		shouldErr bool
	}{
		{
			raw: []byte{
				0x00, 0x00, 0x0c, 0x07, 0xac, 0xd2, 0x98, 0x01, 0xa7, 0xa5, 0x76, 0x3f,
				0x08, 0x00, 0x45, 0x00, 0x00, 0xf2, 0x00, 0x00, 0x40, 0x00, 0x40, 0x06,
				0xe3, 0x69, 0xac, 0x18, 0x18, 0x1d, 0x0a, 0x5f, 0x88, 0x08, 0xd3, 0x07,
				0x0c, 0xea, 0xd4, 0x37, 0xc6, 0x4f, 0x18, 0x2f, 0xa8, 0xd5, 0x80, 0x18,
				0x10, 0x17, 0x25, 0xc6, 0x00, 0x00, 0x01, 0x01, 0x08, 0x0a, 0x1c, 0xc8,
				0xf1, 0xab, 0xf9, 0x2e, 0xc2, 0x4a, 0xba, 0x00, 0x00, 0x01, 0x85, 0xa6,
				0xff, 0x01, 0x00, 0x00, 0x00, 0x01, 0x21, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x72, 0x6f, 0x6f, 0x74, 0x00, 0x14,
				0xbd, 0xd9, 0x4a, 0x79, 0x53, 0x60, 0x76, 0x93, 0x1a, 0xb7, 0xc8, 0xeb,
				0x97, 0x30, 0xaf, 0x41, 0x1d, 0x3d, 0xa2, 0x59, 0x6d, 0x79, 0x73, 0x71,
				0x6c, 0x5f, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x61, 0x73,
				0x73, 0x77, 0x6f, 0x72, 0x64, 0x00, 0x69, 0x03, 0x5f, 0x6f, 0x73, 0x08,
				0x6f, 0x73, 0x78, 0x31, 0x30, 0x2e, 0x31, 0x31, 0x0c, 0x5f, 0x63, 0x6c,
				0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x08, 0x6c, 0x69,
				0x62, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x04, 0x5f, 0x70, 0x69, 0x64, 0x05,
				0x33, 0x31, 0x38, 0x30, 0x33, 0x0f, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
				0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x06, 0x35, 0x2e,
				0x37, 0x2e, 0x31, 0x33, 0x09, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
				0x72, 0x6d, 0x06, 0x78, 0x38, 0x36, 0x5f, 0x36, 0x34, 0x0c, 0x70, 0x72,
				0x6f, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x05, 0x6d,
				0x79, 0x73, 0x71, 0x6c,
			},
			expect:    _clientAuthPacket,
			shouldErr: false,
		},
		{
			raw: []byte{
				0x98, 0x01, 0xa7, 0xa5, 0x76, 0x3f, 0x00, 0xf2, 0x8b, 0xee, 0x4a, 0x31,
				0x08, 0x00, 0x45, 0x00, 0x02, 0x6f, 0xf6, 0x4e, 0x40, 0x00, 0x38, 0x06,
				0x2a, 0xc9, 0x0a, 0x60, 0x50, 0x86, 0xac, 0x18, 0x18, 0x73, 0x0c, 0xea,
				0xf8, 0xd6, 0xd8, 0xdf, 0xa5, 0xbc, 0x54, 0x61, 0x3e, 0x44, 0x80, 0x18,
				0x00, 0xeb, 0x6c, 0x39, 0x00, 0x00, 0x01, 0x01, 0x08, 0x0a, 0x74, 0x2b,
				0x2f, 0x95, 0x45, 0x5d, 0x84, 0x63, 0x01, 0x00, 0x00, 0x01, 0x06, 0x46,
				0x00, 0x00, 0x02, 0x03, 0x64, 0x65, 0x66, 0x12, 0x69, 0x6e, 0x66, 0x6f,
				0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x63, 0x68, 0x65,
				0x6d, 0x61, 0x07, 0x43, 0x4f, 0x4c, 0x55, 0x4d, 0x4e, 0x53, 0x07, 0x43,
				0x4f, 0x4c, 0x55, 0x4d, 0x4e, 0x53, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64,
				0x0b, 0x43, 0x4f, 0x4c, 0x55, 0x4d, 0x4e, 0x5f, 0x4e, 0x41, 0x4d, 0x45,
				0x0c, 0x21, 0x00, 0xc0, 0x00, 0x00, 0x00, 0xfd, 0x01, 0x00, 0x00, 0x00,
				0x00, 0x45, 0x00, 0x00, 0x03, 0x03, 0x64, 0x65, 0x66, 0x12, 0x69, 0x6e,
				0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x63,
				0x68, 0x65, 0x6d, 0x61, 0x07, 0x43, 0x4f, 0x4c, 0x55, 0x4d, 0x4e, 0x53,
				0x07, 0x43, 0x4f, 0x4c, 0x55, 0x4d, 0x4e, 0x53, 0x04, 0x54, 0x79, 0x70,
				0x65, 0x0b, 0x43, 0x4f, 0x4c, 0x55, 0x4d, 0x4e, 0x5f, 0x54, 0x59, 0x50,
				0x45, 0x0c, 0x21, 0x00, 0xf7, 0xff, 0x08, 0x00, 0xfc, 0x11, 0x00, 0x00,
				0x00, 0x00, 0x45, 0x00, 0x00, 0x04, 0x03, 0x64, 0x65, 0x66, 0x12, 0x69,
				0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
				0x63, 0x68, 0x65, 0x6d, 0x61, 0x07, 0x43, 0x4f, 0x4c, 0x55, 0x4d, 0x4e,
				0x53, 0x07, 0x43, 0x4f, 0x4c, 0x55, 0x4d, 0x4e, 0x53, 0x04, 0x4e, 0x75,
				0x6c, 0x6c, 0x0b, 0x49, 0x53, 0x5f, 0x4e, 0x55, 0x4c, 0x4c, 0x41, 0x42,
				0x4c, 0x45, 0x0c, 0x21, 0x00, 0x09, 0x00, 0x00, 0x00, 0xfd, 0x01, 0x00,
				0x00, 0x00, 0x00, 0x43, 0x00, 0x00, 0x05, 0x03, 0x64, 0x65, 0x66, 0x12,
				0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
				0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x07, 0x43, 0x4f, 0x4c, 0x55, 0x4d,
				0x4e, 0x53, 0x07, 0x43, 0x4f, 0x4c, 0x55, 0x4d, 0x4e, 0x53, 0x03, 0x4b,
				0x65, 0x79, 0x0a, 0x43, 0x4f, 0x4c, 0x55, 0x4d, 0x4e, 0x5f, 0x4b, 0x45,
				0x59, 0x0c, 0x21, 0x00, 0x09, 0x00, 0x00, 0x00, 0xfd, 0x01, 0x00, 0x00,
				0x00, 0x00, 0x4b, 0x00, 0x00, 0x06, 0x03, 0x64, 0x65, 0x66, 0x12, 0x69,
				0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
				0x63, 0x68, 0x65, 0x6d, 0x61, 0x07, 0x43, 0x4f, 0x4c, 0x55, 0x4d, 0x4e,
				0x53, 0x07, 0x43, 0x4f, 0x4c, 0x55, 0x4d, 0x4e, 0x53, 0x07, 0x44, 0x65,
				0x66, 0x61, 0x75, 0x6c, 0x74, 0x0e, 0x43, 0x4f, 0x4c, 0x55, 0x4d, 0x4e,
				0x5f, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x0c, 0x21, 0x00, 0xf7,
				0xff, 0x08, 0x00, 0xfc, 0x10, 0x00, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00,
				0x07, 0x03, 0x64, 0x65, 0x66, 0x12, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
				0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
				0x07, 0x43, 0x4f, 0x4c, 0x55, 0x4d, 0x4e, 0x53, 0x07, 0x43, 0x4f, 0x4c,
				0x55, 0x4d, 0x4e, 0x53, 0x05, 0x45, 0x78, 0x74, 0x72, 0x61, 0x05, 0x45,
				0x58, 0x54, 0x52, 0x41, 0x0c, 0x21, 0x00, 0x5a, 0x00, 0x00, 0x00, 0xfd,
				0x01, 0x00, 0x00, 0x00, 0x00, 0x22, 0x00, 0x00, 0x08, 0x02, 0x69, 0x64,
				0x07, 0x69, 0x6e, 0x74, 0x28, 0x31, 0x30, 0x29, 0x02, 0x4e, 0x4f, 0x03,
				0x50, 0x52, 0x49, 0xfb, 0x0e, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x69, 0x6e,
				0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x17, 0x00, 0x00, 0x09, 0x04,
				0x6e, 0x61, 0x6d, 0x65, 0x0b, 0x76, 0x61, 0x72, 0x63, 0x68, 0x61, 0x72,
				0x28, 0x32, 0x30, 0x29, 0x02, 0x4e, 0x4f, 0x00, 0xfb, 0x00, 0x12, 0x00,
				0x00, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x07, 0x69, 0x6e, 0x74, 0x28, 0x31,
				0x30, 0x29, 0x02, 0x4e, 0x4f, 0x00, 0xfb, 0x00, 0x1a, 0x00, 0x00, 0x0b,
				0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x0c, 0x64, 0x65, 0x63, 0x69,
				0x6d, 0x61, 0x6c, 0x28, 0x38, 0x2c, 0x32, 0x29, 0x02, 0x4e, 0x4f, 0x00,
				0xfb, 0x00, 0x07, 0x00, 0x00, 0x0c, 0xfe, 0x00, 0x00, 0x22, 0x00, 0x00,
				0x00,
			},
			expect:    _resultSetPacket,
			shouldErr: false,
		},
		{
			raw: []byte{
				0x98, 0x01, 0xa7, 0xa5, 0x76, 0x3f, 0x00, 0xf2, 0x8b, 0xee, 0x49, 0xb5,
				0x08, 0x00, 0x45, 0x00, 0x01, 0x53, 0x72, 0x18, 0x40, 0x00, 0x38, 0x06,
				0xb0, 0x11, 0x0a, 0x60, 0x50, 0x86, 0xac, 0x18, 0x18, 0x7d, 0x0c, 0xea,
				0xfa, 0xa4, 0x93, 0x4a, 0xc7, 0x94, 0x85, 0x26, 0x47, 0x98, 0x80, 0x18,
				0x00, 0xe3, 0xeb, 0xb2, 0x00, 0x00, 0x01, 0x01, 0x08, 0x0a, 0x7d, 0x83,
				0xbf, 0xe8, 0x4b, 0x95, 0xec, 0x08, 0x01, 0x00, 0x00, 0x01, 0x04, 0x32,
				0x00, 0x00, 0x02, 0x03, 0x64, 0x65, 0x66, 0x04, 0x74, 0x65, 0x73, 0x74,
				0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x0a,
				0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x02, 0x69,
				0x64, 0x02, 0x69, 0x64, 0x0c, 0x3f, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x03,
				0x03, 0x42, 0x00, 0x00, 0x00, 0x36, 0x00, 0x00, 0x03, 0x03, 0x64, 0x65,
				0x66, 0x04, 0x74, 0x65, 0x73, 0x74, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72,
				0x74, 0x6d, 0x65, 0x6e, 0x74, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74,
				0x6d, 0x65, 0x6e, 0x74, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x04, 0x6e, 0x61,
				0x6d, 0x65, 0x0c, 0x21, 0x00, 0x3c, 0x00, 0x00, 0x00, 0xfd, 0x01, 0x10,
				0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x04, 0x03, 0x64, 0x65, 0x66, 0x04,
				0x74, 0x65, 0x73, 0x74, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
				0x65, 0x6e, 0x74, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
				0x6e, 0x74, 0x03, 0x61, 0x67, 0x65, 0x03, 0x61, 0x67, 0x65, 0x0c, 0x3f,
				0x00, 0x0a, 0x00, 0x00, 0x00, 0x03, 0x01, 0x10, 0x00, 0x00, 0x00, 0x3a,
				0x00, 0x00, 0x05, 0x03, 0x64, 0x65, 0x66, 0x04, 0x74, 0x65, 0x73, 0x74,
				0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x0a,
				0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x06, 0x73,
				0x61, 0x6c, 0x61, 0x72, 0x79, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79,
				0x0c, 0x3f, 0x00, 0x0a, 0x00, 0x00, 0x00, 0xf6, 0x01, 0x10, 0x02, 0x00,
				0x00, 0x05, 0x00, 0x00, 0x06, 0xfe, 0x00, 0x00, 0x22, 0x00, 0x1e, 0x00,
				0x00, 0x07, 0x00, 0x00, 0x14, 0x00, 0x00, 0x00, 0x09, 0x72, 0x6f, 0x6e,
				0x61, 0x6c, 0x64, 0x6f, 0x31, 0x32, 0x23, 0x00, 0x00, 0x00, 0x09, 0x35,
				0x30, 0x30, 0x30, 0x30, 0x30, 0x2e, 0x38, 0x30, 0x05, 0x00, 0x00, 0x08,
				0xfe, 0x00, 0x00, 0x22, 0x00,
			},
			expect:    _stmtExecuteResponse,
			shouldErr: false,
		},
		{
			raw: []byte{
				0x98, 0x01, 0xa7, 0xa5, 0x76, 0x3f, 0x00, 0xf2, 0x8b, 0xee, 0x49, 0xb5,
				0x08, 0x00, 0x45, 0x00, 0x01, 0x8d, 0x72, 0x17, 0x40, 0x00, 0x38, 0x06,
				0xaf, 0xd8, 0x0a, 0x60, 0x50, 0x86, 0xac, 0x18, 0x18, 0x7d, 0x0c, 0xea,
				0xfa, 0xa4, 0x93, 0x4a, 0xc6, 0x3b, 0x85, 0x26, 0x47, 0x68, 0x80, 0x18,
				0x00, 0xe3, 0xc3, 0x12, 0x00, 0x00, 0x01, 0x01, 0x08, 0x0a, 0x7d, 0x83,
				0xbf, 0xc3, 0x4b, 0x95, 0xeb, 0xe1, 0x0c, 0x00, 0x00, 0x01, 0x00, 0x01,
				0x00, 0x00, 0x00, 0x04, 0x00, 0x03, 0x00, 0x00, 0x00, 0x00, 0x17, 0x00,
				0x00, 0x02, 0x03, 0x64, 0x65, 0x66, 0x00, 0x00, 0x00, 0x01, 0x3f, 0x00,
				0x0c, 0x3f, 0x00, 0x00, 0x00, 0x00, 0x00, 0xfd, 0x80, 0x00, 0x00, 0x00,
				0x00, 0x17, 0x00, 0x00, 0x03, 0x03, 0x64, 0x65, 0x66, 0x00, 0x00, 0x00,
				0x01, 0x3f, 0x00, 0x0c, 0x3f, 0x00, 0x00, 0x00, 0x00, 0x00, 0xfd, 0x80,
				0x00, 0x00, 0x00, 0x00, 0x17, 0x00, 0x00, 0x04, 0x03, 0x64, 0x65, 0x66,
				0x00, 0x00, 0x00, 0x01, 0x3f, 0x00, 0x0c, 0x3f, 0x00, 0x00, 0x00, 0x00,
				0x00, 0xfd, 0x80, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x05, 0xfe,
				0x00, 0x00, 0x02, 0x00, 0x32, 0x00, 0x00, 0x06, 0x03, 0x64, 0x65, 0x66,
				0x04, 0x74, 0x65, 0x73, 0x74, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74,
				0x6d, 0x65, 0x6e, 0x74, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
				0x65, 0x6e, 0x74, 0x02, 0x69, 0x64, 0x02, 0x69, 0x64, 0x0c, 0x3f, 0x00,
				0x0a, 0x00, 0x00, 0x00, 0x03, 0x03, 0x42, 0x00, 0x00, 0x00, 0x36, 0x00,
				0x00, 0x07, 0x03, 0x64, 0x65, 0x66, 0x04, 0x74, 0x65, 0x73, 0x74, 0x0a,
				0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x0a, 0x64,
				0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x04, 0x6e, 0x61,
				0x6d, 0x65, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x0c, 0x21, 0x00, 0x3c, 0x00,
				0x00, 0x00, 0xfd, 0x01, 0x10, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x08,
				0x03, 0x64, 0x65, 0x66, 0x04, 0x74, 0x65, 0x73, 0x74, 0x0a, 0x64, 0x65,
				0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x0a, 0x64, 0x65, 0x70,
				0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x03, 0x61, 0x67, 0x65, 0x03,
				0x61, 0x67, 0x65, 0x0c, 0x3f, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x03, 0x01,
				0x10, 0x00, 0x00, 0x00, 0x3a, 0x00, 0x00, 0x09, 0x03, 0x64, 0x65, 0x66,
				0x04, 0x74, 0x65, 0x73, 0x74, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74,
				0x6d, 0x65, 0x6e, 0x74, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
				0x65, 0x6e, 0x74, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x06, 0x73,
				0x61, 0x6c, 0x61, 0x72, 0x79, 0x0c, 0x3f, 0x00, 0x0a, 0x00, 0x00, 0x00,
				0xf6, 0x01, 0x10, 0x02, 0x00, 0x00, 0x05, 0x00, 0x00, 0x0a, 0xfe, 0x00,
				0x00, 0x02, 0x00,
			},
			expect:    _prepareResponse,
			shouldErr: false,
		},
		{
			raw: []byte{
				0x98, 0x01, 0xa7, 0xa5, 0x76, 0x3f, 0x00, 0xf2, 0x8b, 0xee, 0x4a, 0x31,
				0x08, 0x00, 0x45, 0x00, 0x00, 0x82, 0xf6, 0x47, 0x40, 0x00, 0x38, 0x06,
				0x2c, 0xbd, 0x0a, 0x60, 0x50, 0x86, 0xac, 0x18, 0x18, 0x73, 0x0c, 0xea,
				0xf8, 0xd6, 0xd8, 0xdf, 0xa4, 0x1e, 0x54, 0x61, 0x3d, 0x13, 0x80, 0x18,
				0x00, 0xe3, 0x60, 0x98, 0x00, 0x00, 0x01, 0x01, 0x08, 0x0a, 0x74, 0x1e,
				0x52, 0x7d, 0x45, 0x50, 0xb7, 0xb4, 0x4a, 0x00, 0x00, 0x00, 0x0a, 0x35,
				0x2e, 0x37, 0x2e, 0x32, 0x30, 0x00, 0x16, 0x00, 0x00, 0x00, 0x01, 0x28,
				0x04, 0x0d, 0x16, 0x1e, 0x62, 0x17, 0x00, 0xff, 0xf7, 0x08, 0x02, 0x00,
				0xff, 0x81, 0x15, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x27, 0x41, 0x63, 0x5d, 0x65, 0x7a, 0x19, 0x50, 0x61, 0x62, 0x5f,
				0x21, 0x00, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x5f, 0x6e, 0x61, 0x74, 0x69,
				0x76, 0x65, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x00,
			},
			expect:    _serverGreetingPacket,
			shouldErr: false,
		},
		{
			raw: []byte{
				0x98, 0x01, 0xa7, 0xa5, 0x76, 0x3f, 0x00, 0xf2, 0x8b, 0xee, 0x4a, 0x31,
				0x08, 0x00, 0x45, 0x00, 0x00, 0x48, 0xf6, 0x49, 0x40, 0x00, 0x38, 0x06,
				0x2c, 0xf5, 0x0a, 0x60, 0x50, 0x86, 0xac, 0x18, 0x18, 0x73, 0x0c, 0xea,
				0xf8, 0xd6, 0xd8, 0xdf, 0xa4, 0x6c, 0x54, 0x61, 0x3d, 0xd6, 0x80, 0x18,
				0x00, 0xeb, 0x38, 0xfe, 0x00, 0x00, 0x01, 0x01, 0x08, 0x0a, 0x74, 0x1e,
				0x52, 0xad, 0x45, 0x50, 0xb8, 0x02, 0x10, 0x00, 0x00, 0x02, 0x00, 0x00,
				0x00, 0x02, 0x40, 0x00, 0x00, 0x00, 0x07, 0x01, 0x05, 0x04, 0x74, 0x65,
				0x73, 0x74,
			},
			expect:    _okPacket,
			shouldErr: false,
		},
		{
			raw: []byte{
				0x00, 0x00, 0x0c, 0x07, 0xac, 0xd2, 0x98, 0x01, 0xa7, 0xa5, 0x76, 0x3f,
				0x08, 0x00, 0x45, 0x00, 0x00, 0x59, 0x00, 0x00, 0x40, 0x00, 0x40, 0x06,
				0x1b, 0x2e, 0xac, 0x18, 0x18, 0x73, 0x0a, 0x60, 0x50, 0x86, 0xf8, 0xd6,
				0x0c, 0xea, 0x54, 0x61, 0x3d, 0xd6, 0xd8, 0xdf, 0xa4, 0x80, 0x80, 0x18,
				0x10, 0x16, 0xdf, 0xba, 0x00, 0x00, 0x01, 0x01, 0x08, 0x0a, 0x45, 0x50,
				0xb8, 0x2e, 0x74, 0x1e, 0x52, 0xad, 0x21, 0x00, 0x00, 0x00, 0x03, 0x73,
				0x65, 0x6c, 0x65, 0x63, 0x74, 0x20, 0x40, 0x40, 0x76, 0x65, 0x72, 0x73,
				0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x20,
				0x6c, 0x69, 0x6d, 0x69, 0x74, 0x20, 0x31,
			},
			expect:    _queryPacket,
			shouldErr: false,
		},
		{
			raw: []byte{
				0x98, 0x01, 0xa7, 0xa5, 0x76, 0x3f, 0x00, 0xf2, 0x8b, 0xee, 0x4a, 0x31,
				0x08, 0x00, 0x45, 0x00, 0x00, 0x90, 0xf6, 0x4a, 0x40, 0x00, 0x38, 0x06,
				0x2c, 0xac, 0x0a, 0x60, 0x50, 0x86, 0xac, 0x18, 0x18, 0x73, 0x0c, 0xea,
				0xf8, 0xd6, 0xd8, 0xdf, 0xa4, 0x80, 0x54, 0x61, 0x3d, 0xfb, 0x80, 0x18,
				0x00, 0xeb, 0xfb, 0xc6, 0x00, 0x00, 0x01, 0x01, 0x08, 0x0a, 0x74, 0x1e,
				0x52, 0xd5, 0x45, 0x50, 0xb8, 0x2e, 0x01, 0x00, 0x00, 0x01, 0x01, 0x27,
				0x00, 0x00, 0x02, 0x03, 0x64, 0x65, 0x66, 0x00, 0x00, 0x00, 0x11, 0x40,
				0x40, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6d,
				0x6d, 0x65, 0x6e, 0x74, 0x00, 0x0c, 0x21, 0x00, 0x54, 0x00, 0x00, 0x00,
				0xfd, 0x00, 0x00, 0x1f, 0x00, 0x00, 0x1d, 0x00, 0x00, 0x03, 0x1c, 0x4d,
				0x79, 0x53, 0x51, 0x4c, 0x20, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
				0x74, 0x79, 0x20, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x28, 0x47,
				0x50, 0x4c, 0x29, 0x07, 0x00, 0x00, 0x04, 0xfe, 0x00, 0x00, 0x02, 0x00,
				0x00, 0x00,
			},
			expect:    _resultSetPacket,
			shouldErr: false,
		},
		{
			raw: []byte{
				0x00, 0x00, 0x0c, 0x07, 0xac, 0xd2, 0x98, 0x01, 0xa7, 0xa5, 0x76, 0x3f,
				0x08, 0x00, 0x45, 0x00, 0x00, 0x74, 0x00, 0x00, 0x40, 0x00, 0x40, 0x06,
				0x1b, 0x09, 0xac, 0x18, 0x18, 0x7d, 0x0a, 0x60, 0x50, 0x86, 0xfa, 0xa4,
				0x0c, 0xea, 0x85, 0x26, 0x47, 0x28, 0x93, 0x4a, 0xc6, 0x3b, 0x80, 0x18,
				0x10, 0x16, 0x60, 0x07, 0x00, 0x00, 0x01, 0x01, 0x08, 0x0a, 0x4b, 0x95,
				0xeb, 0xe1, 0x7d, 0x83, 0xb8, 0x13, 0x3c, 0x00, 0x00, 0x00, 0x16, 0x53,
				0x45, 0x4c, 0x45, 0x43, 0x54, 0x20, 0x2a, 0x20, 0x46, 0x52, 0x4f, 0x4d,
				0x20, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x20,
				0x57, 0x48, 0x45, 0x52, 0x45, 0x20, 0x28, 0x6e, 0x61, 0x6d, 0x65, 0x3d,
				0x3f, 0x20, 0x41, 0x4e, 0x44, 0x20, 0x61, 0x67, 0x65, 0x3e, 0x3f, 0x20,
				0x41, 0x4e, 0x44, 0x20, 0x61, 0x67, 0x65, 0x3c, 0x3f, 0x29,
			},
			expect:    _prepareQuery,
			shouldErr: false,
		},
		{
			raw: []byte{
				0x00, 0x00, 0x0c, 0x07, 0xac, 0xd2, 0x98, 0x01, 0xa7, 0xa5, 0x76, 0x3f,
				0x08, 0x00, 0x45, 0x00, 0x00, 0x64, 0x00, 0x00, 0x40, 0x00, 0x40, 0x06,
				0x1b, 0x19, 0xac, 0x18, 0x18, 0x7d, 0x0a, 0x60, 0x50, 0x86, 0xfa, 0xa4,
				0x0c, 0xea, 0x85, 0x26, 0x47, 0x68, 0x93, 0x4a, 0xc7, 0x94, 0x80, 0x18,
				0x10, 0x0c, 0x98, 0x29, 0x00, 0x00, 0x01, 0x01, 0x08, 0x0a, 0x4b, 0x95,
				0xec, 0x08, 0x7d, 0x83, 0xbf, 0xc3, 0x2c, 0x00, 0x00, 0x00, 0x17, 0x01,
				0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x01, 0xfe, 0x00,
				0x08, 0x00, 0x08, 0x00, 0x09, 0x72, 0x6f, 0x6e, 0x61, 0x6c, 0x64, 0x6f,
				0x31, 0x32, 0x17, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x32, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
			expect:    _stmtExecute,
			shouldErr: false,
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		actual := DecodePacket(tc.raw)
		if tc.shouldErr {
			should.Nil(actual, "case #%d", idx)
			continue
		}
		should.NotNil(actual, "case #%d", idx)
		should.Equal(tc.expect, actual["protocol_type"], "case #%d", idx)
	}
}

func TestNoPanic(t *testing.T) {

	// test case 通过wireshark 抓包获取
	var testCase = []struct {
		raw []byte
	}{
		{
			raw: []byte{
				8, 0, 0, 0, 2, 99, 97, 114, 112, 111, 111, 108,
			},
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		should.NotPanics(func() {
			actual := DecodePacketWithoutHeader(tc.raw)
			should.Nil(actual, "case #%d fail", idx)
		}, "case #%d fail", idx)
	}
}
