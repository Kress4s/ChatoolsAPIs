package constant

import "errors"

var (
	// ErrNilAuthorization token权限验证
	ErrNilAuthorization = errors.New("Header Authorization Cant be null")
)
