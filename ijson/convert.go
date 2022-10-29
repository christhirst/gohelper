package ijson

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

func StructToJson(bodybytes *bytes.Buffer) (map[string]interface{}, error) {
	decoder := json.NewDecoder(bodybytes)
	var empJSON map[string]interface{}
	err := decoder.Decode(&empJSON)
	if err != nil {
		return nil, errors.New("Unable to convert request body")
	}

	return empJSON, nil
}

func RenderJSON(w http.ResponseWriter, v interface{}, statusCode int) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	_, _ = w.Write(buf.Bytes())
}
