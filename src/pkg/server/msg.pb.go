// Code generated by protoc-gen-go from "msg.proto"
// DO NOT EDIT!

package server

import proto "goprotobuf.googlecode.com/hg/proto"
import "math"
import "os"

// Reference proto, math & os imports to suppress error if they are not otherwise used.
var _ = proto.GetString
var _ = math.Inf
var _ os.Error


type request_Verb int32

const (
	request_GET    = 1
	request_SET    = 2
	request_DEL    = 3
	request_ESET   = 4
	request_REV    = 5
	request_WAIT   = 6
	request_NOP    = 7
	request_WATCH  = 8
	request_WALK   = 9
	request_CANCEL = 10
	request_GETDIR = 14
	request_STAT   = 16
)

var request_Verb_name = map[int32]string{
	1:  "GET",
	2:  "SET",
	3:  "DEL",
	4:  "ESET",
	5:  "REV",
	6:  "WAIT",
	7:  "NOP",
	8:  "WATCH",
	9:  "WALK",
	10: "CANCEL",
	14: "GETDIR",
	16: "STAT",
}
var request_Verb_value = map[string]int32{
	"GET":    1,
	"SET":    2,
	"DEL":    3,
	"ESET":   4,
	"REV":    5,
	"WAIT":   6,
	"NOP":    7,
	"WATCH":  8,
	"WALK":   9,
	"CANCEL": 10,
	"GETDIR": 14,
	"STAT":   16,
}

func newRequest_Verb(x int32) *request_Verb {
	e := request_Verb(x)
	return &e
}

type response_Err int32

const (
	response_OTHER        = 127
	response_TAG_IN_USE   = 1
	response_UNKNOWN_VERB = 2
	response_REDIRECT     = 3
	response_TOO_LATE     = 4
	response_REV_MISMATCH = 5
	response_BAD_PATH     = 6
	response_MISSING_ARG  = 7
	response_RANGE        = 8
	response_NOTDIR       = 20
	response_ISDIR        = 21
	response_NOENT        = 22
)

var response_Err_name = map[int32]string{
	127: "OTHER",
	1:   "TAG_IN_USE",
	2:   "UNKNOWN_VERB",
	3:   "REDIRECT",
	4:   "TOO_LATE",
	5:   "REV_MISMATCH",
	6:   "BAD_PATH",
	7:   "MISSING_ARG",
	8:   "RANGE",
	20:  "NOTDIR",
	21:  "ISDIR",
	22:  "NOENT",
}
var response_Err_value = map[string]int32{
	"OTHER":        127,
	"TAG_IN_USE":   1,
	"UNKNOWN_VERB": 2,
	"REDIRECT":     3,
	"TOO_LATE":     4,
	"REV_MISMATCH": 5,
	"BAD_PATH":     6,
	"MISSING_ARG":  7,
	"RANGE":        8,
	"NOTDIR":       20,
	"ISDIR":        21,
	"NOENT":        22,
}

func newResponse_Err(x int32) *response_Err {
	e := response_Err(x)
	return &e
}

type request struct {
	Tag              *int32        "PB(varint,1,opt,name=tag)"
	Verb             *request_Verb "PB(varint,2,opt,name=verb,enum=server.request_Verb)"
	Path             *string       "PB(bytes,4,opt,name=path)"
	Value            []byte        "PB(bytes,5,opt,name=value)"
	OtherTag         *int32        "PB(varint,6,opt,name=other_tag)"
	Offset           *int32        "PB(varint,7,opt,name=offset)"
	Rev              *int64        "PB(varint,9,opt,name=rev)"
	XXX_unrecognized []byte
}

func (this *request) Reset() {
	*this = request{}
}

type response struct {
	Tag              *int32        "PB(varint,1,opt,name=tag)"
	Flags            *int32        "PB(varint,2,opt,name=flags)"
	Rev              *int64        "PB(varint,3,opt,name=rev)"
	Path             *string       "PB(bytes,5,opt,name=path)"
	Value            []byte        "PB(bytes,6,opt,name=value)"
	Len              *int32        "PB(varint,8,opt,name=len)"
	ErrCode          *response_Err "PB(varint,100,opt,name=err_code,enum=server.response_Err)"
	ErrDetail        *string       "PB(bytes,101,opt,name=err_detail)"
	XXX_unrecognized []byte
}

func (this *response) Reset() {
	*this = response{}
}

func init() {
	proto.RegisterEnum("server.request_Verb", request_Verb_name, request_Verb_value)
	proto.RegisterEnum("server.response_Err", response_Err_name, response_Err_value)
}
