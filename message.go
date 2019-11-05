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

import "encoding/json"

type Content struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}
type At struct {
	IsAtAll bool `json:"isAtAll"`
}

type DTMessage struct {
	MsgType string  `json:"msgtype"`
	Content Content `json:"markdown"`
	At      At      `json:"at"`
}

func (dt *DTMessage) SetContent(c Content) {
	dt.Content = c
}

func (dt *DTMessage) SetAtAll(isAtAll bool) {
	dt.At.IsAtAll = isAtAll
}

func (dt DTMessage) ToByte() ([]byte, error) {
	dt.MsgType = "markdown"
	jsonByte, err := json.Marshal(dt)
	return jsonByte, err
}

func (dt DTMessage) String() string {
	jb, err := dt.ToByte()
	if jb == nil || err != nil {
		return ""
	}
	return string(jb)
}

type Result struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}
