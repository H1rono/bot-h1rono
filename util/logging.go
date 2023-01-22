package util

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/traPtitech/go-traq"
)

func SetupLogging() {
	// TRACE, DEBUG, INFO, WARNING, ERROR, FATAL, PANIC
	l := os.Getenv("LOG_LEVEL")
	switch l {
	case "TRACE":
		log.SetLevel(log.TraceLevel)
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "INFO", "":
		// l = "" の場合に後の出力が残念になるため
		l = "INFO"
		log.SetLevel(log.InfoLevel)
	case "WARNING":
		log.SetLevel(log.WarnLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	case "FATAL":
		log.SetLevel(log.FatalLevel)
	case "PANIC":
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
		log.Debug(r.Request.Body)
		log.Debug("レスポンス:")
		log.Debug(r.Header)
		log.Debug(r.Body)
	}
}

func LogSentMessage(m *traq.Message) {
	log.Debug("メッセージを投稿しました。")
	log.Debugf("ID: %s", m.Id)
	log.Debugf("チャンネルID: %s", m.ChannelId)
}
