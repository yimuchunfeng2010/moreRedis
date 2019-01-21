package types

import "time"

type ProcessingRequest struct {
	CommitID int64
	Req ReqType
	Key string
	Value string
	Next *ProcessingRequest
	CreateTime time.Time

}
