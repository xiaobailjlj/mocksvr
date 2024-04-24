// Copyright 2022 Innkeeper Belm(孔令飞) <nosbelm@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://miniblog.

package biz

//go:generate mockgen -destination mock_biz.go -package biz miniblog/internal/miniblog/biz IBiz

import (
	"miniblog/internal/miniblog/biz/post"
	"miniblog/internal/miniblog/biz/stub"
	"miniblog/internal/miniblog/biz/user"
	"miniblog/internal/miniblog/store"
)

// IBiz 定义了 Biz 层需要实现的方法.
type IBiz interface {
	Users() user.UserBiz
	Posts() post.PostBiz
	Stubs() stub.StubBiz
}

// 确保 biz 实现了 IBiz 接口.
var _ IBiz = (*biz)(nil)

// biz 是 IBiz 的一个具体实现.
type biz struct {
	ds store.IStore
}

// 确保 biz 实现了 IBiz 接口.
var _ IBiz = (*biz)(nil)

// NewBiz 创建一个 IBiz 类型的实例.
func NewBiz(ds store.IStore) *biz {
	return &biz{ds: ds}
}

// Users 返回一个实现了 UserBiz 接口的实例.
func (b *biz) Users() user.UserBiz {
	return user.New(b.ds)
}

// Posts 返回一个实现了 PostBiz 接口的实例.
func (b *biz) Posts() post.PostBiz {
	return post.New(b.ds)
}

// Stubs 返回一个实现了 StubBiz 接口的实例.
func (b *biz) Stubs() stub.StubBiz {
	return stub.New(b.ds)
}
