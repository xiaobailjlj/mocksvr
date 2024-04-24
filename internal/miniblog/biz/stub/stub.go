package stub

//go:generate mockgen -destination mock_stub.go -package stub miniblog/internal/miniblog/biz/stub StubBiz

import (
	"context"
	"miniblog/internal/miniblog/store"
	"miniblog/internal/pkg/model"
	v1 "miniblog/pkg/api/miniblog/v1"
)

// StubBiz defines functions used to handle stub request.
type StubBiz interface {
	CreateStubInterface(ctx context.Context, r *v1.CreateStubInterfaceRequest) error
	CreateStubRule(ctx context.Context, r *v1.CreateStubRuleRequest) error
}

// The implementation of StubBiz interface.
type stubBiz struct {
	ds store.IStore
}

// Make sure that stubBiz implements the StubBiz interface.
// We can find this problem in the compile stage with the following assignment statement.
var _ StubBiz = (*stubBiz)(nil)

func New(ds store.IStore) *stubBiz {
	return &stubBiz{ds: ds}
}

// Create is the implementation of the `Create` method in StubBiz interface.
func (b *stubBiz) CreateStubInterface(ctx context.Context, r *v1.CreateStubInterfaceRequest) error {
	stubInterfaceM := &model.StubInterfaceM{
		DefRespCode:   r.DefRespCode,
		DefRespHeader: r.DefRespHeader,
		DefRespBody:   r.DefRespBody,
		Owner:         r.Owner,
		Desc:          r.Desc,
		Status:        "AVAILABLE",
	}

	if err := b.ds.StubInterface().Create(ctx, stubInterfaceM); err != nil {
		return err
	}

	return nil
}

// Create is the implementation of the `Create` method in StubBiz interface.
func (b *stubBiz) CreateStubRule(ctx context.Context, r *v1.CreateStubRuleRequest) error {
	stubRuleM := &model.StubRuleM{
		InterfaceID: r.InterfaceId,
		MatchRule:   r.MatchRule,
		MatchType:   r.MatchType,
		RespCode:    r.RespCode,
		RespHeader:  r.RespHeader,
		RespBody:    r.RespBody,
		DelayTime:   r.DelayTime,
		Desc:        r.Desc,
		Status:      "AVAILABLE",
	}

	if err := b.ds.StubRule().Create(ctx, stubRuleM); err != nil {
		return err
	}

	return nil
}
