// file: middleware/basicauth.go
// 未使用

package middleware

import "github.com/kataras/iris/v12/middleware/basicauth"

// BasicAuth middleware sample.
var BasicAuth = basicauth.New(basicauth.Config{
	Users: map[string]string{
		"admin": "admin123",
	},
})
