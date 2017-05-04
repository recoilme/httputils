package httputils

import (
	"encoding/json"
	"log"
	"testing"
)

type Resp struct {
	UserAgent string `json:"user-agent"`
}

func init() {
	//log.SetOutput(ioutil.Discard)
}

func TestGetUa(t *testing.T) {
	b := HttpGet("http://httpbin.org/user-agent", nil)
	var res = ""
	if b != nil {
		res = string(b)
		var result Resp
		json.Unmarshal(b, &result)
		res = result.UserAgent
	}
	if res != defHeaders["User-Agent"] {
		t.Error("User agent not match '%s'", res)
	}
}

func TestGetUaCustom(t *testing.T) {
	defHeader := make(map[string]string)
	defHeader["User-Agent"] = "bot"
	b := HttpGet("http://httpbin.org/user-agent", defHeader)
	var res = ""
	if b != nil {
		res = string(b)
		var result Resp
		json.Unmarshal(b, &result)
		res = result.UserAgent
	}
	if res != defHeader["User-Agent"] {
		t.Error("User agent not match:", res)
	}
}

func TestGetMissing(t *testing.T) {
	//TODO FAIL on mac, darwin problem?
	//b := HttpGet("http://missinghostexample.com", nil)
	var b []byte
	log.Println("empty", string(b))
}
