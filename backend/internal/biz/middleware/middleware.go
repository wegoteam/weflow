package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
	"github.com/wegoteam/weflow/internal/base"
	"github.com/wegoteam/weflow/internal/biz/entity/bo"
	"net/http"
	"time"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	IdentityKey   = "identity"
)

func InitJwt() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:                 "test zone",
		Key:                   []byte("secret key"),
		Timeout:               time.Hour,
		MaxRefresh:            time.Hour,
		TokenLookup:           "header:Authorization, cookie:token, query:token, form:token, param:token",
		TokenHeadName:         "Bearer", //default Bearer Authorization: Bearer token
		LoginResponse:         LoginResponse,
		Authenticator:         Authenticator,
		IdentityKey:           IdentityKey,
		IdentityHandler:       IdentityHandler,
		PayloadFunc:           PayloadFunc,
		HTTPStatusMessageFunc: HTTPStatusMessageFunc,
		Unauthorized:          Unauthorized,
	})
	if err != nil {
		panic(err)
	}
}

func LoginResponse(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
	c.JSON(http.StatusOK, utils.H{
		"code":    code,
		"token":   token,
		"expire":  expire.Format(time.RFC3339),
		"message": "success",
	})
}

func Authenticator(ctx context.Context, c *app.RequestContext) (interface{}, error) {
	var loginStruct struct {
		Username string `form:"username" json:"username" query:"username" vd:"(len($) > 0 && len($) < 30); msg:'Illegal format'"`
		Password string `form:"password" json:"password" query:"password" vd:"(len($) > 0 && len($) < 30); msg:'Illegal format'"`
	}
	if err := c.BindAndValidate(&loginStruct); err != nil {
		return nil, err
	}
	userInfo := bo.UserInfoResult{
		UserID:   "547",
		UserName: "xuch01",
		Password: "xuch01",
	}

	return userInfo, nil
}

func IdentityHandler(ctx context.Context, c *app.RequestContext) interface{} {
	token := jwt.GetToken(ctx, c)
	fmt.Printf("token = %+v\n", token)
	claims := jwt.ExtractClaims(ctx, c)
	key := claims[IdentityKey]
	fmt.Printf("key = %+v\n", key)
	return &bo.UserInfoResult{
		UserName: "",
	}
}

func PayloadFunc(data interface{}) jwt.MapClaims {
	//if v, ok := data.(*bo.UserInfoResult); ok {
	//	return jwt.MapClaims{
	//		IdentityKey: v.UserName,
	//	}
	//}
	fmt.Printf("data = %+v\n", data)
	return jwt.MapClaims{}
}

func HTTPStatusMessageFunc(e error, ctx context.Context, c *app.RequestContext) string {
	hlog.CtxErrorf(ctx, "jwt biz err = %+v", e.Error())

	if errors.Is(e, jwt.ErrEmptyParamToken) {
		return "Token不存在或者已过期"
	}
	return e.Error()
}

func Unauthorized(ctx context.Context, c *app.RequestContext, code int, message string) {
	c.JSON(http.StatusOK, base.Fail(code, message))
}
