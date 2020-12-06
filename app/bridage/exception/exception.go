package exception

import (
	"ChatoolsAPIs/app/bridage/constant"

	"github.com/York-xia/tools/curd/common"
)

// ExceptionController ...
type ExceptionController struct {
	common.BaseController
}

// URLMapping ...
func (c *ExceptionController) URLMapping() {
	c.Mapping("ExceptToken", c.ExceptToken)
}

// ExceptToken ...
func (c *ExceptionController) ExceptToken() {
	c.Data["json"] = common.RestResult{
		Code:    constant.TOKEN_IS_NIL,
		Message: "Authorization can not be null",
	}
	c.ServeJSON()
}

var exception = ExceptionController{}

// GetInst ...
func GetInst() ExceptionController {
	return exception
}
