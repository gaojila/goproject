package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	// "golang.org/x/text"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

func DetermineEncoding(r io.Reader) encoding.Encoding {
	byte, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(byte, "")
	return e
}

func printCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`)
	matchs := re.FindAllSubmatch(contents, -1)
	for _, m := range matchs {
		fmt.Printf("City: %s Url %s \n", m[2], m[1])
	}
	fmt.Printf("Matches found: %d\n", len(matchs))
}

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error StatusCode:", resp.StatusCode)
		return
	}

	e := DetermineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, e.NewEncoder())

	s, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	printCityList(s)
}
