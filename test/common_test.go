package test

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"golang.org/x/exp/slices"
)

type PersonTest struct {
	Name     string
	Birthday time.Time
}

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

// test MD5
func TestMD5(t *testing.T) {
	// calculate md5: 15985huawei2023-01-01 00:00:002023-01-02 00:00:00
	str := "15985huawei2023-01-01 00:00:002023-01-02 00:00:00"
	fmt.Println(MD5(str))
	fmt.Println(SignWithMD5(str))
}

// MD5 string
func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func SignWithMD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	ret := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return ret
}

func TestSliceSort(t *testing.T) {
	var s1 []PersonTest = []PersonTest{
		{"zhangsan", time.Now()},
		{"lisi", time.Now().Add(time.Hour * 24)},
	}

	slices.SortStableFunc(s1, func(x, y PersonTest) bool {
		return x.Birthday.After(y.Birthday)
	})

	fmt.Println(s1)
}

func TestJsonStr(t *testing.T) {
	str := "{\"debug_mode\":false,\"enable_auto_performance_tracking\":true,\"enable_oom_tracking\":true,\"enable_app_hang_tracking\":true,\"app_hang_timeout_seconds\":5}"
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(m)
}

func TestSwitch(t *testing.T) {
	var a int = 1
	switch a {
	case 1:
		fmt.Println("case 1")
		fallthrough // fallthrough will execute next case
	case 2:
		fmt.Println("case 2")
	}
}

type Cat struct {
	Name string
	Age  int
}

// test modify map when ranging map
func TestModifyMapWhenRangingMap(t *testing.T) {
	m := make(map[string]Cat)
	m["cat1"] = Cat{"cat1", 1}
	m["cat2"] = Cat{"cat2", 2}
	m["cat3"] = Cat{"cat3", 3}

	for k, v := range m {
		v.Age = 10
		m[k] = v
	}
	fmt.Println(m)

	// or use pointer
	m2 := make(map[string]*Cat)
	m2["cat1"] = &Cat{"cat1", 1}
	m2["cat2"] = &Cat{"cat2", 2}
	m2["cat3"] = &Cat{"cat3", 3}

	for _, v := range m2 {
		v.Age = 10
	}
	fmt.Println(m2)
}

func TestCovertTime(t *testing.T) {
	// convert date string `2023-01-01 00:00:00` to time and get month from time
	dateStr := "Tue, 16 May 2023 09:08:53 GMT"

	nt, err := time.Parse(time.RFC1123, dateStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(nt.Unix())
}
