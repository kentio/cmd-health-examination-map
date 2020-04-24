package fetcher

import (
	"bufio"
	"fmt"
	browser "github.com/EDDYCJY/fake-useragent"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var UserAgent = browser.IPad()

func Get(url string) (resp *http.Response, err error) {
	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("User-Agent", UserAgent)
	return client.Do(request)
}

var rateLimiter = time.Tick(100 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	resp, err := Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error")
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(
	r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8 // default
	}
	e, _, _ := charset.DetermineEncoding( // 识别页面的编码
		bytes, "")
	return e
}
