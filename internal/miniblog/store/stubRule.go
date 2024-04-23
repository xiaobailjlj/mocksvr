package store

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"miniblog/internal/pkg/model"
)

// StubRuleStore 定义了 StubRule 模块在 store 层所实现的方法.
type StubRuleStore interface {
	Create(ctx context.Context, stubRule *model.StubRuleM) error
	Read(ctx context.Context, url string) (*model.StubRuleM, error)
	Update(ctx context.Context, stubRule *model.StubRuleM) error
	Delete(ctx context.Context, url string) error
}

// StubRuleStore 接口的实现.
type stubRules struct {
	db *gorm.DB
}

// 确保 stubRules 实现了 StubRuleStore 接口.
var _ StubRuleStore = (*stubRules)(nil)

func newStubRules(db *gorm.DB) *stubRules {
	return &stubRules{db}
}

// Create 插入一条记录.
func (u *stubRules) Create(ctx context.Context, stubRule *model.StubRuleM) error {
	return u.db.Create(&stubRule).Error
}

// Read 读取一条记录
func (u *stubRules) Read(ctx context.Context, url string) (*model.StubRuleM, error) {
	var stubRule model.StubRuleM
	if err := u.db.Where("url = ?", url).First(&url).Error; err != nil {
		return nil, err
	}

	return &stubRule, nil
}

// Update 更新一条记录.
func (u *stubRules) Update(ctx context.Context, stubRule *model.StubRuleM) error {
	return u.db.Save(stubRule).Error
}

// Delete 删除一条记录
func (u *stubRules) Delete(ctx context.Context, url string) error {
	err := u.db.Where("url = ?", url).Delete(&model.StubRuleM{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	return nil
}
