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
