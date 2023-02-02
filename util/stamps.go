package util

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/traPtitech/go-traq"
)

type Stamps []traq.Stamp

var STAMP_REGEXP = regexp.MustCompile(`:[a-zA-Z0-9_+*?-]+(\.[a-zA-Z0-9_-]+)*:`)

func FetchStamps(client *traq.APIClient, auth context.Context) Stamps {
	s, r, err := client.StampApi.
		GetStamps(auth).
		IncludeUnicode(true).
		Execute()
	if err != nil {
		log.Error(err)
	}
	LogResponse(r)
	return s
}

func ExtractStampPatterns(msg string) []string {
	return STAMP_REGEXP.FindAllString(msg, -1)
}

func Pattern2RegexAndEffect(pattern string) (re *regexp.Regexp, effect string) {
	// :oisu-*.party.parrot: -> oisu-* | .party.parrot
	i := strings.Index(pattern, ".")
	l := len(pattern)
	if i < 0 {
		i = l - 1
	}
	bs := []byte(pattern)
	// :oisu-*.party.parrot: の oisu-*
	body := string(bs[1:i])
	// :oisu-*.party.parrot: の .party.parrot なければ空
	effect = string(bs[i : l-1])
	// パターン->正規表現
	re_src := strings.ReplaceAll(body, `*`, `[a-zA-Z0-9_-]*`)
	re_src = strings.ReplaceAll(re_src, `?`, `[a-zA-Z0-9_-]`)
	re_src = strings.ReplaceAll(re_src, `+`, `[a-zA-Z0-9_-]+`)
	re = regexp.MustCompile(fmt.Sprintf(`(?i)^%s$`, re_src))
	return
}

func FindAllStamps(pattern string, stamps []traq.Stamp) []string {
	re, effect := Pattern2RegexAndEffect(pattern)
	// 後でこの配列をJoinする
	result := make([]string, 0, len(stamps))
	for _, stamp := range stamps {
		if re.Match([]byte(stamp.Name)) {
			res := fmt.Sprintf(":%s%s:", stamp.Name, effect)
			result = append(result, res)
		}
	}
	if len(result) == 0 {
		result = append(result, pattern)
	}
	return result
}
