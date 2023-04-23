package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCommon(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	}))
	defer server.Close()

	httpGet, err := server.Client().Get("https://wwww.baidu.com")
	if err != nil {
		fmt.Println("something goes wrong")
		t.Error(err)
		return
	}
	defer httpGet.Body.Close()
}
