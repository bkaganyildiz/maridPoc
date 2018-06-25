package util

import (
	"os"
	"fmt"
	"strconv"
	"bytes"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetStrEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf(key))
	}
	return val
}

func GetBoolEnv(key string) bool {
	val := GetStrEnv(key)
	ret, err := strconv.ParseBool(val)
	Check(err)
	return ret
}

func MapToQueryParamConverter(config map[string]string)  string {
	var buffer bytes.Buffer
	for k, v := range config {
		buffer.WriteString(strings.Join([]string{k, "=", v, "&"}, ""))
	}
	return buffer.String()
}