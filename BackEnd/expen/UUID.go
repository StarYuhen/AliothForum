package expen

import (
	"github.com/google/uuid"
)

// CreatUid 生成uid
func CreatUid() string {
	uid, _ := uuid.NewUUID()
	return uid.String()
}
