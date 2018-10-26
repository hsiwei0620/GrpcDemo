// Code generated by protoc-gen-go. DO NOT EDIT.
// source: APIMTLSTio.proto

package MqGrpcProject

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type APIMTLST_Request struct {
	Mtrlcate             string   `protobuf:"bytes,1,opt,name=mtrlcate,proto3" json:"mtrlcate,omitempty"`
	Mtrlproductid        string   `protobuf:"bytes,2,opt,name=mtrlproductid,proto3" json:"mtrlproductid,omitempty"`
	Mtrllotid            string   `protobuf:"bytes,3,opt,name=mtrllotid,proto3" json:"mtrllotid,omitempty"`
	Mtrlstat             string   `protobuf:"bytes,4,opt,name=mtrlstat,proto3" json:"mtrlstat,omitempty"`
	Mtrlpositionid       string   `protobuf:"bytes,5,opt,name=mtrlpositionid,proto3" json:"mtrlpositionid,omitempty"`
	Coneqptid            string   `protobuf:"bytes,6,opt,name=coneqptid,proto3" json:"coneqptid,omitempty"`
	Salesorder           string   `protobuf:"bytes,7,opt,name=salesorder,proto3" json:"salesorder,omitempty"`
	Lineid               string   `protobuf:"bytes,8,opt,name=lineid,proto3" json:"lineid,omitempty"`
	Createdate           string   `protobuf:"bytes,9,opt,name=createdate,proto3" json:"createdate,omitempty"`
	Lotid                string   `protobuf:"bytes,10,opt,name=lotid,proto3" json:"lotid,omitempty"`
	Nxopeno              string   `protobuf:"bytes,11,opt,name=nxopeno,proto3" json:"nxopeno,omitempty"`
	Querysubmtrl         string   `protobuf:"bytes,12,opt,name=querysubmtrl,proto3" json:"querysubmtrl,omitempty"`
	Serverip             string   `protobuf:"bytes,13,opt,name=serverip,proto3" json:"serverip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *APIMTLST_Request) Reset()         { *m = APIMTLST_Request{} }
func (m *APIMTLST_Request) String() string { return proto.CompactTextString(m) }
func (*APIMTLST_Request) ProtoMessage()    {}
func (*APIMTLST_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_65dfdf7d9b509dd2, []int{0}
}

func (m *APIMTLST_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_APIMTLST_Request.Unmarshal(m, b)
}
func (m *APIMTLST_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_APIMTLST_Request.Marshal(b, m, deterministic)
}
func (m *APIMTLST_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_APIMTLST_Request.Merge(m, src)
}
func (m *APIMTLST_Request) XXX_Size() int {
	return xxx_messageInfo_APIMTLST_Request.Size(m)
}
func (m *APIMTLST_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_APIMTLST_Request.DiscardUnknown(m)
}

var xxx_messageInfo_APIMTLST_Request proto.InternalMessageInfo

func (m *APIMTLST_Request) GetMtrlcate() string {
	if m != nil {
		return m.Mtrlcate
	}
	return ""
}

func (m *APIMTLST_Request) GetMtrlproductid() string {
	if m != nil {
		return m.Mtrlproductid
	}
	return ""
}

func (m *APIMTLST_Request) GetMtrllotid() string {
	if m != nil {
		return m.Mtrllotid
	}
	return ""
}

func (m *APIMTLST_Request) GetMtrlstat() string {
	if m != nil {
		return m.Mtrlstat
	}
	return ""
}

func (m *APIMTLST_Request) GetMtrlpositionid() string {
	if m != nil {
		return m.Mtrlpositionid
	}
	return ""
}

func (m *APIMTLST_Request) GetConeqptid() string {
	if m != nil {
		return m.Coneqptid
	}
	return ""
}

func (m *APIMTLST_Request) GetSalesorder() string {
	if m != nil {
		return m.Salesorder
	}
	return ""
}

func (m *APIMTLST_Request) GetLineid() string {
	if m != nil {
		return m.Lineid
	}
	return ""
}

func (m *APIMTLST_Request) GetCreatedate() string {
	if m != nil {
		return m.Createdate
	}
	return ""
}

func (m *APIMTLST_Request) GetLotid() string {
	if m != nil {
		return m.Lotid
	}
	return ""
}

func (m *APIMTLST_Request) GetNxopeno() string {
	if m != nil {
		return m.Nxopeno
	}
	return ""
}

func (m *APIMTLST_Request) GetQuerysubmtrl() string {
	if m != nil {
		return m.Querysubmtrl
	}
	return ""
}

func (m *APIMTLST_Request) GetServerip() string {
	if m != nil {
		return m.Serverip
	}
	return ""
}

type APIMTLST_Reply struct {
	Trxid                string         `protobuf:"bytes,1,opt,name=trxid,proto3" json:"trxid,omitempty"`
	Typeid               string         `protobuf:"bytes,2,opt,name=typeid,proto3" json:"typeid,omitempty"`
	Rtncode              string         `protobuf:"bytes,3,opt,name=rtncode,proto3" json:"rtncode,omitempty"`
	Rtnmesg              string         `protobuf:"bytes,4,opt,name=rtnmesg,proto3" json:"rtnmesg,omitempty"`
	Mtrlcnt              string         `protobuf:"bytes,5,opt,name=mtrlcnt,proto3" json:"mtrlcnt,omitempty"`
	Oary                 []*APIMTLSToA  `protobuf:"bytes,6,rep,name=oary,proto3" json:"oary,omitempty"`
	Mascnt               string         `protobuf:"bytes,7,opt,name=mascnt,proto3" json:"mascnt,omitempty"`
	Oary2                []*APIMTLSToA2 `protobuf:"bytes,8,rep,name=oary2,proto3" json:"oary2,omitempty"`
	Errmsg               string         `protobuf:"bytes,9,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *APIMTLST_Reply) Reset()         { *m = APIMTLST_Reply{} }
func (m *APIMTLST_Reply) String() string { return proto.CompactTextString(m) }
func (*APIMTLST_Reply) ProtoMessage()    {}
func (*APIMTLST_Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_65dfdf7d9b509dd2, []int{1}
}

func (m *APIMTLST_Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_APIMTLST_Reply.Unmarshal(m, b)
}
func (m *APIMTLST_Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_APIMTLST_Reply.Marshal(b, m, deterministic)
}
func (m *APIMTLST_Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_APIMTLST_Reply.Merge(m, src)
}
func (m *APIMTLST_Reply) XXX_Size() int {
	return xxx_messageInfo_APIMTLST_Reply.Size(m)
}
func (m *APIMTLST_Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_APIMTLST_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_APIMTLST_Reply proto.InternalMessageInfo

func (m *APIMTLST_Reply) GetTrxid() string {
	if m != nil {
		return m.Trxid
	}
	return ""
}

func (m *APIMTLST_Reply) GetTypeid() string {
	if m != nil {
		return m.Typeid
	}
	return ""
}

func (m *APIMTLST_Reply) GetRtncode() string {
	if m != nil {
		return m.Rtncode
	}
	return ""
}

func (m *APIMTLST_Reply) GetRtnmesg() string {
	if m != nil {
		return m.Rtnmesg
	}
	return ""
}

func (m *APIMTLST_Reply) GetMtrlcnt() string {
	if m != nil {
		return m.Mtrlcnt
	}
	return ""
}

func (m *APIMTLST_Reply) GetOary() []*APIMTLSToA {
	if m != nil {
		return m.Oary
	}
	return nil
}

func (m *APIMTLST_Reply) GetMascnt() string {
	if m != nil {
		return m.Mascnt
	}
	return ""
}

func (m *APIMTLST_Reply) GetOary2() []*APIMTLSToA2 {
	if m != nil {
		return m.Oary2
	}
	return nil
}

func (m *APIMTLST_Reply) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

type APIMTLSToA struct {
	Mtrlproductid        string   `protobuf:"bytes,1,opt,name=mtrlproductid,proto3" json:"mtrlproductid,omitempty"`
	Mtrllotid            string   `protobuf:"bytes,2,opt,name=mtrllotid,proto3" json:"mtrllotid,omitempty"`
	Seqno                string   `protobuf:"bytes,3,opt,name=seqno,proto3" json:"seqno,omitempty"`
	Containerid          string   `protobuf:"bytes,4,opt,name=containerid,proto3" json:"containerid,omitempty"`
	Mtrlcate             string   `protobuf:"bytes,5,opt,name=mtrlcate,proto3" json:"mtrlcate,omitempty"`
	Mtrlstat             string   `protobuf:"bytes,6,opt,name=mtrlstat,proto3" json:"mtrlstat,omitempty"`
	Mtrlqty              string   `protobuf:"bytes,7,opt,name=mtrlqty,proto3" json:"mtrlqty,omitempty"`
	Chgdate              string   `protobuf:"bytes,8,opt,name=chgdate,proto3" json:"chgdate,omitempty"`
	Chgtime              string   `protobuf:"bytes,9,opt,name=chgtime,proto3" json:"chgtime,omitempty"`
	Chguser              string   `protobuf:"bytes,10,opt,name=chguser,proto3" json:"chguser,omitempty"`
	Tareqptid            string   `protobuf:"bytes,11,opt,name=tareqptid,proto3" json:"tareqptid,omitempty"`
	Eqptid               string   `protobuf:"bytes,12,opt,name=eqptid,proto3" json:"eqptid,omitempty"`
	Eqptmtg              string   `protobuf:"bytes,13,opt,name=eqptmtg,proto3" json:"eqptmtg,omitempty"`
	Comment              string   `protobuf:"bytes,14,opt,name=comment,proto3" json:"comment,omitempty"`
	Mskpositionid        string   `protobuf:"bytes,15,opt,name=mskpositionid,proto3" json:"mskpositionid,omitempty"`
	Datarepdatetime      string   `protobuf:"bytes,16,opt,name=datarepdatetime,proto3" json:"datarepdatetime,omitempty"`
	Oosflg               string   `protobuf:"bytes,17,opt,name=oosflg,proto3" json:"oosflg,omitempty"`
	Createdate           string   `protobuf:"bytes,18,opt,name=createdate,proto3" json:"createdate,omitempty"`
	Expiredate           string   `protobuf:"bytes,19,opt,name=expiredate,proto3" json:"expiredate,omitempty"`
	Validperiod          string   `protobuf:"bytes,20,opt,name=validperiod,proto3" json:"validperiod,omitempty"`
	Minuseqty            string   `protobuf:"bytes,21,opt,name=minuseqty,proto3" json:"minuseqty,omitempty"`
	Salesorder           string   `protobuf:"bytes,22,opt,name=salesorder,proto3" json:"salesorder,omitempty"`
	Lineid               string   `protobuf:"bytes,23,opt,name=lineid,proto3" json:"lineid,omitempty"`
	Enoughflg            string   `protobuf:"bytes,24,opt,name=enoughflg,proto3" json:"enoughflg,omitempty"`
	Lotid                string   `protobuf:"bytes,25,opt,name=lotid,proto3" json:"lotid,omitempty"`
	Sheets               string   `protobuf:"bytes,26,opt,name=sheets,proto3" json:"sheets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *APIMTLSToA) Reset()         { *m = APIMTLSToA{} }
func (m *APIMTLSToA) String() string { return proto.CompactTextString(m) }
func (*APIMTLSToA) ProtoMessage()    {}
func (*APIMTLSToA) Descriptor() ([]byte, []int) {
	return fileDescriptor_65dfdf7d9b509dd2, []int{2}
}

func (m *APIMTLSToA) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_APIMTLSToA.Unmarshal(m, b)
}
func (m *APIMTLSToA) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_APIMTLSToA.Marshal(b, m, deterministic)
}
func (m *APIMTLSToA) XXX_Merge(src proto.Message) {
	xxx_messageInfo_APIMTLSToA.Merge(m, src)
}
func (m *APIMTLSToA) XXX_Size() int {
	return xxx_messageInfo_APIMTLSToA.Size(m)
}
func (m *APIMTLSToA) XXX_DiscardUnknown() {
	xxx_messageInfo_APIMTLSToA.DiscardUnknown(m)
}

var xxx_messageInfo_APIMTLSToA proto.InternalMessageInfo

func (m *APIMTLSToA) GetMtrlproductid() string {
	if m != nil {
		return m.Mtrlproductid
	}
	return ""
}

func (m *APIMTLSToA) GetMtrllotid() string {
	if m != nil {
		return m.Mtrllotid
	}
	return ""
}

func (m *APIMTLSToA) GetSeqno() string {
	if m != nil {
		return m.Seqno
	}
	return ""
}

func (m *APIMTLSToA) GetContainerid() string {
	if m != nil {
		return m.Containerid
	}
	return ""
}

func (m *APIMTLSToA) GetMtrlcate() string {
	if m != nil {
		return m.Mtrlcate
	}
	return ""
}

func (m *APIMTLSToA) GetMtrlstat() string {
	if m != nil {
		return m.Mtrlstat
	}
	return ""
}

func (m *APIMTLSToA) GetMtrlqty() string {
	if m != nil {
		return m.Mtrlqty
	}
	return ""
}

func (m *APIMTLSToA) GetChgdate() string {
	if m != nil {
		return m.Chgdate
	}
	return ""
}

func (m *APIMTLSToA) GetChgtime() string {
	if m != nil {
		return m.Chgtime
	}
	return ""
}

func (m *APIMTLSToA) GetChguser() string {
	if m != nil {
		return m.Chguser
	}
	return ""
}

func (m *APIMTLSToA) GetTareqptid() string {
	if m != nil {
		return m.Tareqptid
	}
	return ""
}

func (m *APIMTLSToA) GetEqptid() string {
	if m != nil {
		return m.Eqptid
	}
	return ""
}

func (m *APIMTLSToA) GetEqptmtg() string {
	if m != nil {
		return m.Eqptmtg
	}
	return ""
}

func (m *APIMTLSToA) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *APIMTLSToA) GetMskpositionid() string {
	if m != nil {
		return m.Mskpositionid
	}
	return ""
}

func (m *APIMTLSToA) GetDatarepdatetime() string {
	if m != nil {
		return m.Datarepdatetime
	}
	return ""
}

func (m *APIMTLSToA) GetOosflg() string {
	if m != nil {
		return m.Oosflg
	}
	return ""
}

func (m *APIMTLSToA) GetCreatedate() string {
	if m != nil {
		return m.Createdate
	}
	return ""
}

func (m *APIMTLSToA) GetExpiredate() string {
	if m != nil {
		return m.Expiredate
	}
	return ""
}

func (m *APIMTLSToA) GetValidperiod() string {
	if m != nil {
		return m.Validperiod
	}
	return ""
}

func (m *APIMTLSToA) GetMinuseqty() string {
	if m != nil {
		return m.Minuseqty
	}
	return ""
}

func (m *APIMTLSToA) GetSalesorder() string {
	if m != nil {
		return m.Salesorder
	}
	return ""
}

func (m *APIMTLSToA) GetLineid() string {
	if m != nil {
		return m.Lineid
	}
	return ""
}

func (m *APIMTLSToA) GetEnoughflg() string {
	if m != nil {
		return m.Enoughflg
	}
	return ""
}

func (m *APIMTLSToA) GetLotid() string {
	if m != nil {
		return m.Lotid
	}
	return ""
}

func (m *APIMTLSToA) GetSheets() string {
	if m != nil {
		return m.Sheets
	}
	return ""
}

type APIMTLSToA2 struct {
	Mtrlprodid           string   `protobuf:"bytes,1,opt,name=mtrlprodid,proto3" json:"mtrlprodid,omitempty"`
	Submtrlprodid        string   `protobuf:"bytes,2,opt,name=submtrlprodid,proto3" json:"submtrlprodid,omitempty"`
	Proportion           string   `protobuf:"bytes,3,opt,name=proportion,proto3" json:"proportion,omitempty"`
	Userid               string   `protobuf:"bytes,4,opt,name=userid,proto3" json:"userid,omitempty"`
	Chgdate              string   `protobuf:"bytes,5,opt,name=chgdate,proto3" json:"chgdate,omitempty"`
	Chgtime              string   `protobuf:"bytes,6,opt,name=chgtime,proto3" json:"chgtime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *APIMTLSToA2) Reset()         { *m = APIMTLSToA2{} }
func (m *APIMTLSToA2) String() string { return proto.CompactTextString(m) }
func (*APIMTLSToA2) ProtoMessage()    {}
func (*APIMTLSToA2) Descriptor() ([]byte, []int) {
	return fileDescriptor_65dfdf7d9b509dd2, []int{3}
}

func (m *APIMTLSToA2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_APIMTLSToA2.Unmarshal(m, b)
}
func (m *APIMTLSToA2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_APIMTLSToA2.Marshal(b, m, deterministic)
}
func (m *APIMTLSToA2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_APIMTLSToA2.Merge(m, src)
}
func (m *APIMTLSToA2) XXX_Size() int {
	return xxx_messageInfo_APIMTLSToA2.Size(m)
}
func (m *APIMTLSToA2) XXX_DiscardUnknown() {
	xxx_messageInfo_APIMTLSToA2.DiscardUnknown(m)
}

var xxx_messageInfo_APIMTLSToA2 proto.InternalMessageInfo

func (m *APIMTLSToA2) GetMtrlprodid() string {
	if m != nil {
		return m.Mtrlprodid
	}
	return ""
}

func (m *APIMTLSToA2) GetSubmtrlprodid() string {
	if m != nil {
		return m.Submtrlprodid
	}
	return ""
}

func (m *APIMTLSToA2) GetProportion() string {
	if m != nil {
		return m.Proportion
	}
	return ""
}

func (m *APIMTLSToA2) GetUserid() string {
	if m != nil {
		return m.Userid
	}
	return ""
}

func (m *APIMTLSToA2) GetChgdate() string {
	if m != nil {
		return m.Chgdate
	}
	return ""
}

func (m *APIMTLSToA2) GetChgtime() string {
	if m != nil {
		return m.Chgtime
	}
	return ""
}

func init() {
	proto.RegisterType((*APIMTLST_Request)(nil), "MqGrpcProject.APIMTLST_Request")
	proto.RegisterType((*APIMTLST_Reply)(nil), "MqGrpcProject.APIMTLST_Reply")
	proto.RegisterType((*APIMTLSToA)(nil), "MqGrpcProject.APIMTLSTo_a")
	proto.RegisterType((*APIMTLSToA2)(nil), "MqGrpcProject.APIMTLSTo_a2")
}

func init() { proto.RegisterFile("APIMTLSTio.proto", fileDescriptor_65dfdf7d9b509dd2) }

var fileDescriptor_65dfdf7d9b509dd2 = []byte{
	// 722 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0xdd, 0x4e, 0x1b, 0x3d,
	0x10, 0x55, 0x12, 0x12, 0x88, 0x13, 0x7e, 0x3e, 0x7f, 0x40, 0x5d, 0x8a, 0x10, 0x8a, 0xaa, 0x8a,
	0xab, 0x48, 0x85, 0x27, 0x28, 0x37, 0x6d, 0x25, 0x90, 0xa2, 0x14, 0xa9, 0x97, 0x68, 0xd9, 0x9d,
	0x06, 0xb7, 0xd9, 0xb5, 0x63, 0x3b, 0x88, 0xbc, 0x4b, 0xa5, 0x4a, 0x7d, 0x92, 0xbe, 0x59, 0xab,
	0xb1, 0xc7, 0xd9, 0x1f, 0x44, 0x7b, 0x97, 0x73, 0x66, 0x32, 0x1e, 0xcf, 0x39, 0xe3, 0x65, 0x7b,
	0xef, 0x26, 0x1f, 0xaf, 0x6f, 0xae, 0x3e, 0xdd, 0x48, 0x35, 0xd6, 0x46, 0x39, 0xc5, 0xb7, 0xaf,
	0x17, 0xef, 0x8d, 0x4e, 0x27, 0x46, 0x7d, 0x85, 0xd4, 0x8d, 0xbe, 0x77, 0xca, 0x9c, 0xdb, 0x29,
	0x2c, 0x96, 0x60, 0x1d, 0x3f, 0x62, 0x5b, 0xb9, 0x33, 0xf3, 0x34, 0x71, 0x20, 0x5a, 0xa7, 0xad,
	0xb3, 0xfe, 0x74, 0x8d, 0xf9, 0x6b, 0xb6, 0x8d, 0xbf, 0xb5, 0x51, 0xd9, 0x32, 0x75, 0x32, 0x13,
	0x6d, 0x9f, 0x50, 0x27, 0xf9, 0x31, 0xeb, 0x23, 0x31, 0x57, 0x98, 0xd1, 0xf1, 0x19, 0x25, 0x11,
	0xeb, 0x5b, 0x97, 0x38, 0xb1, 0x51, 0xd6, 0x47, 0xcc, 0xdf, 0xb0, 0x1d, 0x5f, 0x4a, 0x59, 0xe9,
	0xa4, 0x2a, 0x64, 0x26, 0xba, 0x3e, 0xa3, 0xc1, 0xe2, 0x09, 0xa9, 0x2a, 0x60, 0xa1, 0xf1, 0x84,
	0x5e, 0x38, 0x61, 0x4d, 0xf0, 0x13, 0xc6, 0x6c, 0x32, 0x07, 0xab, 0x4c, 0x06, 0x46, 0x6c, 0xfa,
	0x70, 0x85, 0xe1, 0x87, 0xac, 0x37, 0x97, 0x05, 0xc8, 0x4c, 0x6c, 0xf9, 0x18, 0x21, 0xfc, 0x5f,
	0x6a, 0x20, 0x71, 0x90, 0xe1, 0xdd, 0xfb, 0xe1, 0x7f, 0x25, 0xc3, 0xf7, 0x59, 0x37, 0xdc, 0x89,
	0xf9, 0x50, 0x00, 0x5c, 0xb0, 0xcd, 0xe2, 0x51, 0x69, 0x28, 0x94, 0x18, 0x78, 0x3e, 0x42, 0x3e,
	0x62, 0xc3, 0xc5, 0x12, 0xcc, 0xca, 0x2e, 0xef, 0xb0, 0x7f, 0x31, 0xf4, 0xe1, 0x1a, 0x87, 0xd3,
	0xb0, 0x60, 0x1e, 0xc0, 0x48, 0x2d, 0xb6, 0xc3, 0x34, 0x22, 0x1e, 0xfd, 0x68, 0xb3, 0x9d, 0x8a,
	0x3c, 0x7a, 0xbe, 0xc2, 0x16, 0x9c, 0x79, 0x94, 0x19, 0x29, 0x13, 0x00, 0x5e, 0xc8, 0xad, 0x34,
	0xac, 0xf5, 0x20, 0x84, 0xad, 0x19, 0x57, 0xa4, 0x2a, 0x03, 0x92, 0x21, 0x42, 0x8a, 0xe4, 0x60,
	0x67, 0xa4, 0x41, 0x84, 0x18, 0xf1, 0x72, 0x17, 0x8e, 0x66, 0x1f, 0x21, 0x1f, 0xb3, 0x0d, 0x95,
	0x98, 0x95, 0xe8, 0x9d, 0x76, 0xce, 0x06, 0xe7, 0x47, 0xe3, 0x9a, 0x97, 0xc6, 0xb1, 0x51, 0x75,
	0x9b, 0x4c, 0x7d, 0x1e, 0x76, 0x95, 0x27, 0x16, 0x0b, 0x05, 0x09, 0x08, 0xf1, 0xb7, 0xac, 0x8b,
	0xf1, 0x73, 0xb1, 0xe5, 0x0b, 0xbd, 0x7a, 0xbe, 0xd0, 0xf9, 0x34, 0x64, 0x62, 0x29, 0x30, 0x26,
	0xb7, 0x33, 0x52, 0x85, 0xd0, 0xe8, 0x77, 0x97, 0x0d, 0x2a, 0xf9, 0x4f, 0xfd, 0xd9, 0xfa, 0xa7,
	0x3f, 0xdb, 0x4d, 0x7f, 0xee, 0xb3, 0xae, 0x85, 0x45, 0xa1, 0x68, 0x64, 0x01, 0xf0, 0x53, 0x36,
	0x48, 0x55, 0xe1, 0x12, 0x59, 0x80, 0x91, 0x19, 0x0d, 0xad, 0x4a, 0xd5, 0xf6, 0xa6, 0xdb, 0xd8,
	0x9b, 0xaa, 0xe7, 0x7b, 0x0d, 0xcf, 0xd3, 0xc0, 0x17, 0x6e, 0x45, 0x73, 0x8a, 0x10, 0x23, 0xe9,
	0xfd, 0xcc, 0x9b, 0x31, 0x18, 0x35, 0x42, 0x8a, 0x38, 0x99, 0x47, 0x9b, 0x46, 0x48, 0x91, 0xa5,
	0x05, 0x43, 0x2e, 0x8d, 0x10, 0x6f, 0xed, 0x12, 0x43, 0x3b, 0x13, 0x9c, 0x5a, 0x12, 0x7e, 0xc2,
	0x21, 0x34, 0xa4, 0x09, 0x07, 0x5e, 0xb0, 0x4d, 0xfc, 0x95, 0xbb, 0x19, 0xd9, 0x33, 0x42, 0x7f,
	0x92, 0xca, 0x73, 0x28, 0x9c, 0xd8, 0xa1, 0x93, 0x02, 0xf4, 0x2a, 0xd8, 0x6f, 0x95, 0x25, 0xde,
	0x25, 0x15, 0xaa, 0x24, 0x3f, 0x63, 0xbb, 0x59, 0x82, 0x0d, 0x68, 0xbc, 0x92, 0xbf, 0xcb, 0x9e,
	0xcf, 0x6b, 0xd2, 0xd8, 0x9b, 0x52, 0xf6, 0xcb, 0x7c, 0x26, 0xfe, 0x0b, 0xbd, 0x05, 0xd4, 0xd8,
	0x57, 0xfe, 0x64, 0x5f, 0x4f, 0x18, 0x83, 0x47, 0x2d, 0x4d, 0x88, 0xff, 0x1f, 0xe2, 0x25, 0x83,
	0x9a, 0x3e, 0x24, 0x73, 0x99, 0x69, 0x30, 0x52, 0x65, 0x62, 0x3f, 0x68, 0x5a, 0xa1, 0xbc, 0x53,
	0x64, 0xb1, 0xb4, 0x80, 0xea, 0x1c, 0x90, 0x53, 0x22, 0xd1, 0x78, 0x67, 0x0e, 0xff, 0xf2, 0xce,
	0xbc, 0xa8, 0xbd, 0x33, 0xc7, 0xac, 0x0f, 0x85, 0x5a, 0xce, 0xee, 0xf1, 0x4a, 0x22, 0x54, 0x5d,
	0x13, 0xe5, 0x2b, 0xf3, 0xb2, 0xfa, 0xca, 0x1c, 0xb2, 0x9e, 0xbd, 0x07, 0x70, 0x56, 0x1c, 0x85,
	0x5a, 0x01, 0x8d, 0x7e, 0xb5, 0xd8, 0xb0, 0xba, 0x31, 0xd8, 0x54, 0x74, 0xfb, 0xda, 0xff, 0x15,
	0x06, 0xc5, 0xa1, 0xb7, 0x87, 0x52, 0xe8, 0x09, 0xaf, 0x91, 0x58, 0x45, 0x1b, 0xa5, 0x95, 0x41,
	0xb1, 0x68, 0x13, 0x2a, 0x0c, 0xb6, 0x83, 0xa6, 0x5a, 0x6f, 0x02, 0xa1, 0xaa, 0x65, 0xbb, 0xcf,
	0x5a, 0xb6, 0x57, 0xb3, 0xec, 0xe5, 0x05, 0x3b, 0x90, 0x6a, 0x3c, 0x33, 0x3a, 0xad, 0xbf, 0x04,
	0x97, 0xbc, 0x06, 0x27, 0xf8, 0x05, 0x9b, 0xb4, 0x7e, 0xb6, 0x3b, 0x1f, 0xae, 0x3e, 0xdf, 0xf5,
	0xfc, 0x07, 0xed, 0xe2, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x23, 0x9a, 0xe9, 0xe4, 0x06,
	0x00, 0x00,
}
