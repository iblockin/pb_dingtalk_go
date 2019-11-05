//MIT License
//
//Copyright (c) 2019 BlockIN Inc.
//
//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all
//copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//SOFTWARE.
package pb_dingtalk_go

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const dingTalkOAPI = "oapi.dingtalk.com"
const httpTimoutSecond = time.Duration(30) * time.Second

var dingTalkUrl url.URL = url.URL{
	Scheme: "https",
	Host:   dingTalkOAPI,
	Path:   "robot/send",
}

func Send(message DTMessage, token string) error {
	// append token to url
	value := url.Values{}
	value.Set("access_token", token)
	dtu := dingTalkUrl
	dtu.RawQuery = value.Encode()

	// marshal message to byte
	msgByte, err := message.ToByte()
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", dtu.String(), bytes.NewReader(msgByte))
	if err != nil {
		return err
	}
	req.Header.Add("Accept-Charset", "utf8")
	req.Header.Add("Content-Type", "application/json")

	client := new(http.Client)
	client.Timeout = httpTimoutSecond
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	resultByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	res := Result{}
	err = json.Unmarshal(resultByte, &res)
	if err != nil {
		return errors.New(fmt.Sprintf("unmarshal http response body from json error = %v", err))
	}

	if res.Errcode != 0 {
		return errors.New(fmt.Sprintf("send message to dingtalk error = %s", res.Errmsg))
	}

	return nil
}

func SendWithSign(newMsg DTMessage, token, secret string) error {
	// required a millisecond timestamp
	timeStampMillion := time.Now().UnixNano() / int64(time.Millisecond)

	// string to hash  = timestamp + '\n' + secret
	strToHash := fmt.Sprintf("%d\n%s", timeStampMillion, secret)

	// use `secret` as key to do hash with `strToHash`
	hmac256 := hmac.New(sha256.New, []byte(secret))
	hmac256.Write([]byte(strToHash))
	signedByte := hmac256.Sum(nil)

	// do base64 encode for signedByte which comes from hmac256
	signedBase64 := base64.StdEncoding.EncodeToString(signedByte)

	// last thing is do url encode for signedBase64
	value := url.Values{}
	value.Set("access_token", token)
	value.Set("timestamp", fmt.Sprintf("%d", timeStampMillion))
	value.Set("sign", signedBase64)

	dtu := dingTalkUrl
	dtu.RawQuery = value.Encode()

	// marshal message content
	msgByte, err := newMsg.ToByte()
	if err != nil {
		return err
	}

	// prepare request info
	req, err := http.NewRequest("POST", dtu.String(), bytes.NewReader(msgByte))
	if err != nil {
		return err
	}
	req.Header.Add("Accept-Charset", "utf8")
	req.Header.Add("Content-Type", "application/json")

	client := new(http.Client)
	client.Timeout = httpTimoutSecond
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	resultByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	res := Result{}
	err = json.Unmarshal(resultByte, &res)
	if err != nil {
		return errors.New(fmt.Sprintf("unmarshal http response body from json error = %v", err))
	}

	if res.Errcode != 0 {
		return errors.New(fmt.Sprintf("send message to dingtalk error = %s", res.Errmsg))
	}

	return nil
}
