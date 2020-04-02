package fetcher

import (
	"bufio"
	"fmt"
	"git.trac.cn/nv/spider/pkg/setting"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var rateLimiter <-chan time.Time

func SetUp() {
	rateLimiter = time.Tick(setting.ServerSetting.RateLimit)
}

func Fetch(url string) ([]byte, error) {
	<-rateLimiter

	resp, err := http.Get(strings.TrimSpace(url))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	//bodyReader := bufio.NewReader(resp.Body)
	//e := determineEncoding(bodyReader)
	//utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	return ioutil.ReadAll(resp.Body)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
}
