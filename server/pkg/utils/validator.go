package utils

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// 获取验证错误信息，支持通过 msg tag 自定义错误文案
func GetValidateErrorMessage(err error, obj any) string {
	if err == nil {
		return ""
	}

	// 处理切片校验错误（如 []Req），指出第几条数据出错
	var sliceErr binding.SliceValidationError
	if errors.As(err, &sliceErr) && len(sliceErr) > 0 {
		if index, msg := getFirstSliceItemMessage(obj); msg != "" {
			return fmt.Sprintf("第%d条数据，%s", index, msg)
		}
		return sliceErr[0].Error()
	}

	// 提取字段级校验错误
	var vErrs validator.ValidationErrors
	if errors.As(err, &vErrs) {
		if msg := getFieldErrorMessage(vErrs, obj); msg != "" {
			return msg
		}
	}

	return err.Error()
}

// 获取结构体字段的 msg tag 自定义错误
func getFieldErrorMessage(vErrs validator.ValidationErrors, obj any) string {
	t := reflect.TypeOf(obj)
	if t == nil {
		return ""
	}
	// 解指针拿到结构体类型
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return ""
	}

	for _, e := range vErrs {
		if f, ok := t.FieldByName(e.Field()); ok {
			if msg := f.Tag.Get("msg"); msg != "" {
				return msg
			}
			return e.Error()
		}
	}
	return ""
}

// 对切片对象逐条校验，返回首条失败数据的真实序号和错误信息
func getFirstSliceItemMessage(obj any) (int, string) {
	v := reflect.ValueOf(obj)
	if !v.IsValid() {
		return 0, ""
	}
	for v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return 0, ""
		}
		v = v.Elem()
	}
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return 0, ""
	}

	validate, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok || validate == nil {
		return 0, ""
	}

	for i := 0; i < v.Len(); i++ {
		item := v.Index(i)
		for item.Kind() == reflect.Pointer {
			if item.IsNil() {
				break
			}
			item = item.Elem()
		}
		if item.Kind() != reflect.Struct {
			continue
		}
		if err := validate.Struct(item.Interface()); err != nil {
			var vErrs validator.ValidationErrors
			if !errors.As(err, &vErrs) {
				return i + 1, err.Error()
			}
			msg := getFieldErrorMessage(vErrs, item.Interface())
			if msg == "" {
				msg = err.Error()
			}
			return i + 1, msg
		}
	}
	return 0, ""
}
