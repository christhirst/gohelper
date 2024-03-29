package ijson

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func StructToJson(bodybytes io.Reader) (map[string]interface{}, error) {
	decoder := json.NewDecoder(bodybytes)
	var empJSON map[string]interface{}
	err := decoder.Decode(&empJSON)
	if err != nil {
		return nil, errors.New("Unable to convert struct to json")
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

func JsonFromStruct[K any](w http.ResponseWriter, k K, statusCode int) {
	jsonData, err := json.Marshal(k)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)

}
