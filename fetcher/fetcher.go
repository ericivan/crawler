package fetcher

import (
	"golang.org/x/text/transform"
	"io/ioutil"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"github.com/henrylee2cn/surfer"
)

func Fetch(url string) ([]byte, error) {
	//resp, err := http.Get(url)

	resp, err := surfer.Download(&surfer.Request{
		Url: url,
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(resp.Body)

	e, err := determineEncoding(reader)

	utf8Reader := transform.NewReader(reader, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r *bufio.Reader) (encoding.Encoding, error) {

	bytes, err := bufio.NewReader(r).Peek(1024)

	if err != nil {
		return nil, err
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e, nil
}
