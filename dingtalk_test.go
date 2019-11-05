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
	"fmt"
	"testing"
	"time"
)

var s string = fmt.Sprintf("{\"k1\":\"dingtalk_test\",\"times\":%d}", time.Now().Unix())

func TestSend(t *testing.T) {
	context := Content{}
	context.Title = fmt.Sprintf("测试 发送未签名消息%d", time.Now().Unix())
	// test for markdown
	context.Text = "" +
		"# 测试 \n\n" +
		"# 一级标题 \n\n" +
		"-------------------\n\n" +
		"## 二级标题 \n\n" +
		"-------------------\n\n" +
		"### 三级标题 \n\n" +
		"-------------------\n\n" +
		"代码段落 json格式 \n\n" +
		"```json\n" +
		s + "\n" +
		"```\n"
	var msg DTMessage
	msg.SetContent(context)
	msg.SetAtAll(true)
	err := Send(msg, "5ffb53bc964fc3cfa1f5c610ae6ad82c26674f07512806d2c8ab4a63f25fc08f")
	if err != nil {
		t.Error(err)
	}
}

func TestSendWithSign(t *testing.T) {
	token := "5ffb53bc964fc3cfa1f5c610ae6ad82c26674f07512806d2c8ab4a63f25fc08f"
	secret := "SEC0d4f388445fcdad7c7f679a401224018fea774be9b88c76cb6f00a566435dd34"

	context := Content{}
	context.Title = fmt.Sprintf("测试 发送签名消息%d", time.Now().Unix())
	// test for markdown
	context.Text = "" +
		"# 测试 \n\n" +
		"# 一级标题 \n\n" +
		"-------------------\n\n" +
		"## 二级标题 \n\n" +
		"-------------------\n\n" +
		"### 三级标题 \n\n" +
		"-------------------\n\n" +
		"代码段落 json格式 \n\n" +
		"```json\n" +
		s + "\n" +
		"```\n"

	var msg DTMessage
	msg.SetContent(context)

	err := SendWithSign(msg, token, secret)
	if err != nil {
		t.Error(err)
	}
}
