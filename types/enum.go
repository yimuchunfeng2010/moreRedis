package types

type ReqType int8
const (
	ReqType_UNKOWN  ReqType = iota
	ReqType_GET
	ReqType_SET
)