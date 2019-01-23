package types

type ProcessingRequest struct {
	CommitID   int64
	Req        ReqType
	Key        string
	Value      string
	Next       *ProcessingRequest
	CreateTime int64
}
