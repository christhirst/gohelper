package ijson

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

type Teststruct struct  {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
	Field3 string `json:"field3"`
}


func TestRenderJSON(t *testing.T) {
	got := Teststruct{Field1: "test1",Field2: "test2", Field3: "test3"}
	wr := httptest.NewRecorder()
	RenderJSON(wr,got,200)
	fmt.Println(wr.Body)
	
	t.Error()
}
func TestJsonFromStruct(t *testing.T) {
	got := Teststruct{Field1: "test1",Field2: "test2", Field3: "test3"}
	//suits := []string{"Diamonds", "Hearts", "Clubs", "Spades"}
	wr := httptest.NewRecorder()
	JsonFromStruct(wr,got,200)
	fmt.Println(wr.Body)
	
	t.Error()
}
