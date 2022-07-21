package util

import (
	"os"

	"github.com/go-basic/uuid"
)

//get env
func GetEnv(key, defVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defVal
}

//uuid new
func NewUUID() string {
	return uuid.New()
}
