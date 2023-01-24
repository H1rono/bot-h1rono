package util

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/traPtitech/go-traq"
)

type Stamps []traq.Stamp

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
