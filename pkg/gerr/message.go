// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package gerr

import "fmt"

type ErrorMessage struct {
	Name string
	en   string
	zhCN string
}

func (em ErrorMessage) Message(locale string, err error, a ...interface{}) string {
	format := ""
	switch locale {
	case En:
		format = em.en
	case ZhCN:
		if len(em.zhCN) > 0 {
			format = em.zhCN
		} else {
			format = em.en
		}
	}
	if err != nil {
		return fmt.Sprintf("%s: %s", fmt.Sprintf(format, a...), err.Error())
	} else {
		return fmt.Sprintf(format, a...)
	}
}

var (
	ErrorMissingParameter = ErrorMessage{
		Name: "missing_parameter",
		en:   "missing parameter [%s]",
		zhCN: "缺少参数[%s]",
	}
	ErrorGetServiceConfigFailed = ErrorMessage{
		Name: "get_service_config_failed",
		en:   "get_service_config_failed",
		zhCN: "获取通知服务参数配置失败",
	}

	ErrorIllegalEmailFormat = ErrorMessage{
		Name: "illegal_email_format",
		en:   "illegal Email format [%s]",
		zhCN: "非法的Email格式[%s]",
	}
	ErrorIllegalPort = ErrorMessage{
		Name: "illegal_Port",
		en:   "illegal Port [%s]",
		zhCN: "错误的端口号[%s]",
	}
)