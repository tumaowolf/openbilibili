// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: api/http/v1/captcha.proto

/*
Package v1 is a generated blademaster stub package.
This code was generated with go-common/app/tool/bmgen/protoc-gen-bm v0.1.

It is generated from these files:
	api/http/v1/captcha.proto
*/
package v1

import (
	"context"

	bm "go-common/library/net/http/blademaster"
	"go-common/library/net/http/blademaster/binding"
)

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

// =================
// Captcha Interface
// =================

// captcha 相关服务
type Captcha interface {
	// 创建验证码 `midware:"auth"`
	Create(ctx context.Context, req *CreateCaptchaReq) (resp *CreateCaptchaResp, err error)

	// 校验接口 `midware:"auth" method:"POST"`
	Verify(ctx context.Context, req *VerifyReq) (resp *VerifyResp, err error)
}

var v1CaptchaSvc Captcha

// @params CreateCaptchaReq
// @router GET /xlive/web-room/v1/captcha/create
// @response CreateCaptchaResp
func captchaCreate(c *bm.Context) {
	p := new(CreateCaptchaReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1CaptchaSvc.Create(c, p)
	c.JSON(resp, err)
}

// @params VerifyReq
// @router POST /xlive/web-room/v1/captcha/verify
// @response VerifyResp
func captchaVerify(c *bm.Context) {
	p := new(VerifyReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1CaptchaSvc.Verify(c, p)
	c.JSON(resp, err)
}

// RegisterV1CaptchaService Register the blademaster route with middleware map
// midMap is the middleware map, the key is defined in proto
func RegisterV1CaptchaService(e *bm.Engine, svc Captcha, midMap map[string]bm.HandlerFunc) {
	auth := midMap["auth"]
	v1CaptchaSvc = svc
	e.GET("/xlive/web-room/v1/captcha/create", auth, captchaCreate)
	e.POST("/xlive/web-room/v1/captcha/verify", auth, captchaVerify)
}
