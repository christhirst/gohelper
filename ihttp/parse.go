package ihttp

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

func ParseBody[t any](r *http.Request, jsonTarget t) (t, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error().Err(err).Msg("Unable to read body")
	}
	err = json.Unmarshal(body, &jsonTarget)
	if err != nil {
		log.Error().Err(err).Msg("Unable to Unmarshal file")
	}

	return jsonTarget, err
}
