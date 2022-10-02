package http

import (
	"encoding/json"
	"io"

	"github.com/rs/zerolog/log"
)

func ParseBody() {
	var keys map[string]string
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error().Err(err).Msg("Unable to read body")
	}
	err = json.Unmarshal(body, &keys)
	if err != nil {
		log.Error().Err(err).Msg("Unable to Unmarshal file")
	}
}
