// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"github.com/marmotedu/log"
)

// ErrResponse defines the return messages when an error occurred.
// Reference will be omitted if it does not exist.
// swagger:model
type ErrResponse struct {
	// Code defines the business error code.
	Code int `json:"code"`

	// Message contains the detail of this message.
	// This message is suitable to be exposed to external
	Message string `json:"message"`

	// Reference returns the reference document which maybe useful to solve this error.
	Reference string `json:"reference,omitempty"`
}

// WriteResponse write an error or the response data into http response body.
// It use errors.ParseCoder to parse any error into errors.Coder
// errors.Coder contains error code, user-safe error message and http status code.
// ywh: 统一的返回结构，API 通过同一个函数返回。
func WriteResponse(c *gin.Context, err error, data interface{}) {
	if err != nil {
		// 将 err 解析为 github.com/marmotedu/errors 包中定义的 Coder 类型的错误。
		log.Errorf("%#+v", err)
		coder := errors.ParseCoder(err)
		// 调用 Coder 接口提供的  Code() 、String() 、Reference() 方法，获取该错误的业务码。
		c.JSON(coder.HTTPStatus(), ErrResponse{
			Code:      coder.Code(),
			Message:   coder.String(),
			Reference: coder.Reference(),
		})

		return
	}
	c.JSON(http.StatusOK, data)
}
