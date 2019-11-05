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

import "testing"

func TestDTMessage_SetAtAll(t *testing.T) {
	var dt DTMessage
	dt.SetAtAll(true)
	if !dt.At.IsAtAll {
		t.Error("dt message set at all = true failed")
	}
	dt.SetAtAll(false)
	if dt.At.IsAtAll {
		t.Error("dt message set at all = false failed")
	}
}

func TestDTMessage_SetContent(t *testing.T) {
	title := "test_title"
	text := "test_text"
	ct := Content{
		Title: title,
		Text:  text,
	}
	var dt DTMessage
	dt.SetContent(ct)
	if dt.Content.Text != text || dt.Content.Title != title {
		t.Error("dt message set content failed")
	}
}
