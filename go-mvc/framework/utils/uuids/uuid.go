package uuids

import (
	"github.com/satori/go.uuid"
	"strconv"
	"strings"
)

// 生成uuid 32位
func UUid32() string {
	var err error
	return strings.ReplaceAll(uuid.Must(uuid.NewV1(), err).String(), "-", "")
}

// 根据雪花算法生成id
func SnowId() (string, error) {
	worker, err := NewWorker(1)
	return strconv.FormatInt(worker.GetId(), 10), err
}
