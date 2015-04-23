package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
	"time"
)

//go run ghb.go -n 3 -url http://www.baidu.com

func main() {
	number := flag.Int("n", 10, "loop number")
	url := flag.String("url", "http://www.baidu.com", "loop number")
	flag.Parse()
	num := *number
	ch := make(chan string, num)
	for i := 0; i < num; i++ {
		go GetHTTP(*url, ch)
	}
	fmt.Println("loop is ready for number ", num)
	for i := 0; i < num; i++ {
		v := <-ch
		fmt.Println("HTTP GET : index=", i, v)
	}
}
func BenchmarkHTTP(b *testing.B) {

	for i := 0; i < b.N; i++ {
		fmt.Sprintf("Hello")
	}

}

func GetHTTP(url string, ch chan string) (ret string, flag bool) {
	ret = "url=" + url
	start := time.Now().Unix()
	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		// return err.Error(), false
		ret = ret + ";err=" + err.Error()
	}
	if response.StatusCode != 200 {
		// return string(response.StatusCode), false
		ret = ret + ";code=" + string(response.StatusCode)
	}
	ret = ret + ";code=200"
	body, _ := ioutil.ReadAll(response.Body)
	end := time.Now().Unix()
	//strconv.Itoa
	ret = ret + ";usetime=" + strconv.FormatInt(end-start, 10)
	ch <- ret
	return string(body), true
}
