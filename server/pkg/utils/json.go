package utils

import (
	"github.com/bytedance/sonic"
)

// 将 struct 转换成 JSON 字符串
func StructToJsonString(v any) (string, error) {
	jsonStr, err := sonic.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(jsonStr), nil
}

// 将 JSON 字符串转换为 struct
func JsonStringToStruct(jsonStr string, v any) error {
	return sonic.Unmarshal([]byte(jsonStr), v)
}
