package util

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

func GenCommitID(msg string) (commitID int64, err error) {
	msgByte := Md5Encode(msg)
	msgByte = msgByte[0 : len(msgByte)/2-1]
	commitID , err = ByteToInt(msgByte)
	if err != nil {
		return -1, err
	}
	return commitID, nil
}


func Md5Encode(msg string) []byte {
	Md5Inst := md5.New()
	Md5Inst.Write([]byte(msg))
	result := Md5Inst.Sum([]byte(""))
	return result
}

func ByteToInt(msg []byte) (int64, error) {
	encodedStr := "0x" + hex.EncodeToString(msg)
	data, err := strconv.ParseInt(encodedStr, 0, 64)
	if err != nil {
		return -1, err
	}
	return data, nil
}
