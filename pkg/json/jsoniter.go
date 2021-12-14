// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build jsoniter
// +build jsoniter

// ywh: 插件化选择 JSON 库，表示 tags 是 jsoniter 的时候编译这个 Go 文件。
// 可根据结构化和非结构化场景选择。

package json

import jsoniter "github.com/json-iterator/go"

// RawMessage is exported by component-base/pkg/json package.
type RawMessage = jsoniter.RawMessage

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
	// Marshal is exported by component-base/pkg/json package.
	Marshal = json.Marshal
	// Unmarshal is exported by component-base/pkg/json package.
	Unmarshal = json.Unmarshal
	// MarshalIndent is exported by component-base/pkgn/json package.
	MarshalIndent = json.MarshalIndent
	// NewDecoder is exported by component-base/pkg/json package.
	NewDecoder = json.NewDecoder
	// NewEncoder is exported by component-base/pkg/json package.
	NewEncoder = json.NewEncoder
)
