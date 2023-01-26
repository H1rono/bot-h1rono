package util

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/traPtitech/go-traq"
)

const (
	LOG_LEVEL_TRACE   = "TRACE"
	LOG_LEVEL_DEBUG   = "DEBUG"
	LOG_LEVEL_INFO    = "INFO"
	LOG_LEVEL_WARNING = "WARNING"
	LOG_LEVEL_ERROR   = "ERROR"
	LOG_LEVEL_FATAL   = "FATAL"
	LOG_LEVEL_PANIC   = "PANIC"
)

func SetupLogging(l string) {
	// l: TRACE, DEBUG, INFO, WARNING, ERROR, FATAL, PANIC
	switch l {
	case LOG_LEVEL_TRACE:
		log.SetLevel(log.TraceLevel)
	case LOG_LEVEL_DEBUG:
		log.SetLevel(log.DebugLevel)
	case LOG_LEVEL_INFO, "":
		// l = "" の場合に後の出力が残念になるため
		l = "INFO"
		log.SetLevel(log.InfoLevel)
	case LOG_LEVEL_WARNING:
		log.SetLevel(log.WarnLevel)
	case LOG_LEVEL_ERROR:
		log.SetLevel(log.ErrorLevel)
	case LOG_LEVEL_FATAL:
		log.SetLevel(log.FatalLevel)
	case LOG_LEVEL_PANIC:
		log.SetLevel(log.PanicLevel)
	default:
		log.Fatalf("Unexpected environment variable LOG_LEVEL=%s", l)
	}
	log.Infof("log level is set at %s", l)
}

func LogResponse(r *http.Response) {
	if r.StatusCode >= 400 {
		log.Errorf("エラーレスポンス%sを受け取りました", r.Status)
		log.Error("リクエスト:")
		log.Error(r.Request.Header)
		log.Error(r.Request.Body)
		log.Error("レスポンス:")
		log.Error(r.Header)
		log.Error(r.Body)
	} else {
		log.Debug("リクエスト:")
		log.Debug(r.Request.Header)
		log.Debug("レスポンス:")
		log.Debug(r.Header)
	}
}

func LogSentMessage(m *traq.Message) {
	log.Debug("メッセージを投稿しました。")
	log.Debugf("ID: %s", m.Id)
	log.Debugf("チャンネルID: %s", m.ChannelId)
}
