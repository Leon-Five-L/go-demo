package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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

// test get month from a time
func TestGetMonth(t *testing.T) {
	// convert date string `2023-01-01 00:00:00` to time and get month from time
	beginTime, _ := time.Parse(time.DateTime, "2023-01-01 00:00:00")
	month := beginTime.Month()
	fmt.Println(int(month))
}
