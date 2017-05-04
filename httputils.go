package httputils

import (
	"bytes"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

var defHeaders = make(map[string]string)

func init() {
	http.DefaultClient.Transport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 1 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 1 * time.Second,
	}
	http.DefaultClient = &http.Client{
		Timeout: time.Second * 10,
	}
	defHeaders["User-Agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.12; rv:52.0) Gecko/20100101 Firefox/52.0"
	defHeaders["Accept-Language"] = "ru-RU,ru;q=0.8,en-US;q=0.5,en;q=0.3"
	defHeaders["Accept-Encoding"] = "gzip, deflate, br"
	defHeaders["Referer"] = "https://ya.ru/"
	defHeaders["Cookie"] = ""
}

// HttpGet create request with default headers + custom headers
func HttpGet(url string, headers map[string]string) []byte {
	log.Println("httpGet", url)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	//req.Host = "yandex.ru"
	for k, v := range defHeaders {
		req.Header.Set(k, v)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	} else {
		return body
	}

	return nil
}

// HttpPut create request with default headers + custom headers
func HttpPut(url string, headers map[string]string, b []byte) (result bool) {
	log.Println("httpPut", url)
	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(b))
	if err != nil {
		log.Fatalln(err)
	}
	//req.Host = "yandex.ru"
	for k, v := range defHeaders {
		req.Header.Set(k, v)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	} else {
		result = true
	}
	defer resp.Body.Close()
	return
}
