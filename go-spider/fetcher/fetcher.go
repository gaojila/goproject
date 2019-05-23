package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	// "regexp"
)

func DetermineEncoding(r io.Reader) encoding.Encoding {
	byte, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(byte, "")
	return e
}

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error StatusCode: %d", resp.StatusCode)
	}

	e := DetermineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, e.NewEncoder())

	// s, err := ioutil.ReadAll(utf8Reader)
	// if err != nil {
	// 	panic(err)
	// }
	return ioutil.ReadAll(utf8Reader)
}
