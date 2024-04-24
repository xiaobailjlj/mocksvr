package stub

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"

	"miniblog/internal/pkg/core"
	"miniblog/internal/pkg/errno"
	"miniblog/internal/pkg/log"
	v1 "miniblog/pkg/api/miniblog/v1"
)

// CreateStubInterface 创建一个接口.
func (ctrl *StubController) CreateStubInterface(c *gin.Context) {
	log.C(c).Infow("Create post function called")

	var r v1.CreateStubInterfaceRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)

		return
	}

	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)

		return
	}

	err := ctrl.b.Stubs().CreateStubInterface(c, &r)
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}

// CreateStubRule 创建一条规则.
func (ctrl *StubController) CreateStubRule(c *gin.Context) {
	log.C(c).Infow("Create post function called")

	var r v1.CreateStubRuleRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)

		return
	}

	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)

		return
	}

	err := ctrl.b.Stubs().CreateStubRule(c, &r)
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
