package ihttp

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

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

func UrlExtractor(r *http.Request, formList []string) (queryListMap map[string][]string, err error) {
	queryListMap = make(map[string][]string)
	for i, v := range r.URL.Query() {
		if len(v) > 0 {
			queryListMap[i] = v
		}
	}
	if len(formList) <= len(queryListMap) {
		err = nil
		return
	} else {
		var notPresent string
		for _, v := range formList {
			if len(r.URL.Query()[v]) == 0 {
				notPresent = v
			}
			err = errors.New("One postForm Value not present: " + notPresent)
		}
		return
	}
}

func ParseHeader(r *http.Request, header, split string, entry int, base bool) (string, error) {
	authheader := r.Header.Get(header)
	fmt.Println(authheader)
	idToken := strings.Split(authheader, " ")[entry]
	fmt.Println(idToken)
	if base {
		dIdToken, _ := base64.RawStdEncoding.DecodeString(idToken)
		fmt.Println(string(dIdToken))
		return string(dIdToken), nil
	} else {
		return idToken, nil
	}

}
