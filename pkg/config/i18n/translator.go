package i18n

import (
	"embed"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhuangpeng/rabbit-go/pkg/utils"
	"github.com/zhuangpeng/rabbit-go/pkg/xerr"
	"github.com/zhuangpeng/rabbit-go/pkg/xerr/errcode"
	"golang.org/x/text/language"
	"google.golang.org/grpc/status"
)

//go:embed locale/*.json
var LocaleFS embed.FS

type Translator struct {
	bundle       *i18n.Bundle
	localizer    map[language.Tag]*i18n.Localizer
	supportLangs []language.Tag
}

func (l *Translator) NewBundle(file embed.FS) {
	bundle := i18n.NewBundle(language.Chinese)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	_, err := bundle.LoadMessageFileFS(file, "locale/zh.json")
	logx.Must(err)
	_, err = bundle.LoadMessageFileFS(file, "locale/en.json")
	logx.Must(err)

	l.bundle = bundle
}

func (l *Translator) NewTranslator() {
	l.supportLangs = append(l.supportLangs, language.Chinese)
	l.supportLangs = append(l.supportLangs, language.English)
	l.localizer = make(map[language.Tag]*i18n.Localizer)
	l.localizer[language.Chinese] = i18n.NewLocalizer(l.bundle, language.Chinese.String())
	l.localizer[language.English] = i18n.NewLocalizer(l.bundle, language.English.String())
}

func (l *Translator) Trans(lang string, msgId string) string {
	message, err := l.MatchLocalizer(lang).LocalizeMessage(&i18n.Message{ID: msgId})
	if err != nil {
		return msgId
	}

	if message == "" {
		return msgId
	}

	return message
}

func (l *Translator) TransError(lang string, err error) error {
	if errcode.IsGrpcError(err) {
		message, e := l.MatchLocalizer(lang).LocalizeMessage(&i18n.Message{ID: strings.Split(err.Error(), "desc = ")[1]})
		if e != nil || message == "" {
			message = err.Error()
		}
		return status.Error(status.Code(err), message)
	} else if codeErr, ok := err.(*xerr.CodeError); ok {
		message, e := l.MatchLocalizer(lang).LocalizeMessage(&i18n.Message{ID: codeErr.Error()})
		if e != nil || message == "" {
			message = codeErr.Error()
		}
		return xerr.NewCodeError(codeErr.Code, message)
	} else if apiErr, ok := err.(*xerr.ApiError); ok {
		message, e := l.MatchLocalizer(lang).LocalizeMessage(&i18n.Message{ID: apiErr.Error()})
		if e != nil {
			message = apiErr.Error()
		}
		return xerr.NewApiError(apiErr.Code, message)
	} else {
		return xerr.NewApiError(http.StatusInternalServerError, "failed to translate error message")
	}
}

func (l *Translator) MatchLocalizer(lang string) *i18n.Localizer {
	tags := utils.ParseTags(lang)
	for _, v := range tags {
		if val, ok := l.localizer[v]; ok {
			return val
		}
	}

	return l.localizer[language.Chinese]
}
