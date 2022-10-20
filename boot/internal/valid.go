package internal

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	znTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
	"zcw-admin-server/global"
)

func Valid() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New() //chinese
		enT := en.New() //english
		uni := ut.New(enT, zhT, enT)

		local := global.CONFIG.App.Language
		var o bool
		global.Trans, o = uni.GetTranslator(local)
		if !o {
			global.LOG.Info("验证器翻译失败")
			return
		}
		var err error
		// 注册翻译器
		switch local {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, global.Trans)
		case "zh":
			err = znTranslations.RegisterDefaultTranslations(v, global.Trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, global.Trans)
		}
		if err != nil {
			global.LOG.Info("验证器翻译失败")
		}
	}
}
