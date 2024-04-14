/**
 * @Author: chentong
 * @Date: 2023/10/30 17:32
 */

package http

import (
	"errors"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh_Hans"
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/zeromicro/go-zero/core/logx"
)

type Validator struct {
	Validator *validator.Validate
	Uni       *ut.UniversalTranslator
	Trans     map[string]ut.Translator
}

func NewValidator() *Validator {
	validate := Validator{
		Validator: validator.New(),
		Uni:       ut.New(zh_Hans.New(), en.New()),
		Trans:     make(map[string]ut.Translator),
	}

	enTrans, _ := validate.Uni.GetTranslator("en")
	zhTrans, _ := validate.Uni.GetTranslator("zh")
	validate.Trans["en"] = enTrans
	validate.Trans["zh"] = zhTrans

	validate.Validator.RegisterTagNameFunc(func(field reflect.StructField) string {
		// 默认取字段名称
		name := field.Name
		if labelTag := strings.SplitN(field.Tag.Get("label"), ",", 2)[0]; labelTag != "" {
			// 尝试读取 label tag
			name = labelTag
		} else if jsonTag := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]; jsonTag != "" {
			// 尝试读取 json tag
			name = jsonTag
		}

		return name
	})

	err := enTranslations.RegisterDefaultTranslations(validate.Validator, enTrans)
	if err != nil {
		logx.Error(logx.Field("Detail", err.Error()))
		return nil
	}
	err = zhTranslations.RegisterDefaultTranslations(validate.Validator, zhTrans)
	if err != nil {
		logx.Error(logx.Field("Detail", err.Error()))
		return nil
	}

	return &validate
}

func (v *Validator) Validate(r *http.Request, data any) error {
	err := v.Validator.Struct(data)
	if err == nil {
		return nil
	}

	var errs validator.ValidationErrors
	if errors.As(err, &errs) {
		// 存在多字段错误, 需要遍历把所有错误都取出来
		s := strings.Builder{}
		for _, e := range errs.Translate(v.Trans["zh"]) {
			s.WriteString(e)
			s.WriteString(";")
		}
		return errors.New(s.String())
	}

	var invalidValidationError *validator.InvalidValidationError
	if errors.As(err, &invalidValidationError) {
		logx.Error(err)
		return err
	}

	return nil
}
