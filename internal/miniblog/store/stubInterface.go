package store

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"miniblog/internal/pkg/model"
)

// StubInterfaceStore 定义了 StubInterface 模块在 store 层所实现的方法.
type StubInterfaceStore interface {
	Create(ctx context.Context, stubInterface *model.StubInterfaceM) error
	Read(ctx context.Context, url string) (*model.StubInterfaceM, error)
	Update(ctx context.Context, stubInterface *model.StubInterfaceM) error
	Delete(ctx context.Context, url string) error
}

// StubInterfaceStore 接口的实现.
type stubInterfaces struct {
	db *gorm.DB
}

// 确保 stubInterfaces 实现了 StubInterfaceStore 接口.
var _ StubInterfaceStore = (*stubInterfaces)(nil)

func newStubInterfaces(db *gorm.DB) *stubInterfaces {
	return &stubInterfaces{db}
}

// Create 插入一条记录.
func (u *stubInterfaces) Create(ctx context.Context, stubInterface *model.StubInterfaceM) error {
	return u.db.Create(&stubInterface).Error
}

// Read 读取一条记录
func (u *stubInterfaces) Read(ctx context.Context, url string) (*model.StubInterfaceM, error) {
	var stubInterface model.StubInterfaceM
	if err := u.db.Where("url = ?", url).First(&url).Error; err != nil {
		return nil, err
	}

	return &stubInterface, nil
}

// Update 更新一条记录.
func (u *stubInterfaces) Update(ctx context.Context, stubInterface *model.StubInterfaceM) error {
	return u.db.Save(stubInterface).Error
}

// Delete 删除一条记录
func (u *stubInterfaces) Delete(ctx context.Context, url string) error {
	err := u.db.Where("url = ?", url).Delete(&model.StubInterfaceM{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	return nil
}
