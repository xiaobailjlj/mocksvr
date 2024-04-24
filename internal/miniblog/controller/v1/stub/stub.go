// Copyright 2022 Innkeeper Belm(孔令飞) <nosbelm@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://miniblog.

package stub

import (
	"miniblog/internal/miniblog/biz"
	"miniblog/internal/miniblog/store"
)

// StubController 是 stub 模块在 Controller 层的实现，用来处理博客模块的请求.
type StubController struct {
	b biz.IBiz
}

// New 创建一个 stub controller.
func New(ds store.IStore) *StubController {
	return &StubController{b: biz.NewBiz(ds)}
}
