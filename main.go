package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get(
		"https://list.mgtv.com/1/a1-a1--------c1-25---.html?channelId=1")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}

	//e := determineEncoding(resp.Body)
	//utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	printChannelList(all)
	//
	//fmt.Printf("%s\n", all)

}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
}

func printChannelList(contents []byte) {
	re := regexp.MustCompile(`<li><a[^>]*href="(/-------------\.html\?channelId=\d+)">([^<]+)</a></li>`)
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		fmt.Printf("Channel: %s,  URL: https://list.mgtv.com%s \n", m[2], m[1])
	}

	fmt.Printf("Matches found: %d\n", len(matches))

}
