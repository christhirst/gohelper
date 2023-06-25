package ihttp

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"
)

type TemplateData struct {
	StringMap string
	IntMap    int
}

func TestPrefixFromDB(t *testing.T) {
	w := new(TemplateData)
	ww := new(TemplateData)
	ww.IntMap = 22
	ww.StringMap = "ee"
	ldapBytes, _ := json.Marshal(ww)

	// Create a bytes.Buffer object that implements io.Reader
	buf := bytes.NewBuffer(ldapBytes)
	req := httptest.NewRequest("POST", "/ldap/firstrq", buf)

	s, _ := ParseBody(req, w)
	t.Error(s)
}

func TestParseHeader(t *testing.T) {
	req := httptest.NewRequest("POST", "/ldap/firstrq", nil)
	req.Header.Set("Authorization", "Basic Y2xpZW50bmFtZTQ6dGVzdHB3")
	split := " "
	header := "Authorization"
	s, _ := ParseHeader(req, header, split, 1, true)
	t.Error(s)
}
