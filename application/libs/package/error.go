/**
 * Created by 宅职社 -- mr.wang
 * User: wang
 * Date: 2021-4-8
 * Time: 下午 15:57
 */
package _package

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
)

var errs = map[string]string{
	"required": "必填",
}

type Errors struct {
	err error
}

func StrErrs(err validator.FieldError, ref reflect.Type, Tag string) string {
	s, _ := ref.FieldByName(err.Field())
	return fmt.Sprintf("%s%s", s.Tag.Get(Tag), errs[err.Tag()])
}
