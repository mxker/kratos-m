// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.1
// - protoc             v3.6.1
// source: api/mykratos/v1/my.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationMyCreateMy = "/api.mykratos.v1.My/CreateMy"
const OperationMyDeleteMy = "/api.mykratos.v1.My/DeleteMy"
const OperationMyGetMy = "/api.mykratos.v1.My/GetMy"
const OperationMyListMy = "/api.mykratos.v1.My/ListMy"
const OperationMyUpdateMy = "/api.mykratos.v1.My/UpdateMy"

type MyHTTPServer interface {
	CreateMy(context.Context, *CreateMyRequest) (*CreateMyReply, error)
	DeleteMy(context.Context, *DeleteMyRequest) (*DeleteMyReply, error)
	GetMy(context.Context, *GetMyRequest) (*GetMyReply, error)
	ListMy(context.Context, *ListMyRequest) (*ListMyReply, error)
	UpdateMy(context.Context, *UpdateMyRequest) (*UpdateMyReply, error)
}

func RegisterMyHTTPServer(s *http.Server, srv MyHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/my", _My_CreateMy0_HTTP_Handler(srv))
	r.PUT("/v1/my/{id}", _My_UpdateMy0_HTTP_Handler(srv))
	r.DELETE("/v1/my/{id}", _My_DeleteMy0_HTTP_Handler(srv))
	r.GET("/v1/my/{id}", _My_GetMy0_HTTP_Handler(srv))
	r.GET("/v1/my", _My_ListMy0_HTTP_Handler(srv))
}

func _My_CreateMy0_HTTP_Handler(srv MyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateMyRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMyCreateMy)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateMy(ctx, req.(*CreateMyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateMyReply)
		return ctx.Result(200, reply)
	}
}

func _My_UpdateMy0_HTTP_Handler(srv MyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateMyRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMyUpdateMy)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateMy(ctx, req.(*UpdateMyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateMyReply)
		return ctx.Result(200, reply)
	}
}

func _My_DeleteMy0_HTTP_Handler(srv MyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteMyRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMyDeleteMy)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteMy(ctx, req.(*DeleteMyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteMyReply)
		return ctx.Result(200, reply)
	}
}

func _My_GetMy0_HTTP_Handler(srv MyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetMyRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMyGetMy)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetMy(ctx, req.(*GetMyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetMyReply)
		return ctx.Result(200, reply)
	}
}

func _My_ListMy0_HTTP_Handler(srv MyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListMyRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMyListMy)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListMy(ctx, req.(*ListMyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListMyReply)
		return ctx.Result(200, reply)
	}
}

type MyHTTPClient interface {
	CreateMy(ctx context.Context, req *CreateMyRequest, opts ...http.CallOption) (rsp *CreateMyReply, err error)
	DeleteMy(ctx context.Context, req *DeleteMyRequest, opts ...http.CallOption) (rsp *DeleteMyReply, err error)
	GetMy(ctx context.Context, req *GetMyRequest, opts ...http.CallOption) (rsp *GetMyReply, err error)
	ListMy(ctx context.Context, req *ListMyRequest, opts ...http.CallOption) (rsp *ListMyReply, err error)
	UpdateMy(ctx context.Context, req *UpdateMyRequest, opts ...http.CallOption) (rsp *UpdateMyReply, err error)
}

type MyHTTPClientImpl struct {
	cc *http.Client
}

func NewMyHTTPClient(client *http.Client) MyHTTPClient {
	return &MyHTTPClientImpl{client}
}

func (c *MyHTTPClientImpl) CreateMy(ctx context.Context, in *CreateMyRequest, opts ...http.CallOption) (*CreateMyReply, error) {
	var out CreateMyReply
	pattern := "/v1/my"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMyCreateMy))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MyHTTPClientImpl) DeleteMy(ctx context.Context, in *DeleteMyRequest, opts ...http.CallOption) (*DeleteMyReply, error) {
	var out DeleteMyReply
	pattern := "/v1/my/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMyDeleteMy))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MyHTTPClientImpl) GetMy(ctx context.Context, in *GetMyRequest, opts ...http.CallOption) (*GetMyReply, error) {
	var out GetMyReply
	pattern := "/v1/my/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMyGetMy))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MyHTTPClientImpl) ListMy(ctx context.Context, in *ListMyRequest, opts ...http.CallOption) (*ListMyReply, error) {
	var out ListMyReply
	pattern := "/v1/my"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMyListMy))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MyHTTPClientImpl) UpdateMy(ctx context.Context, in *UpdateMyRequest, opts ...http.CallOption) (*UpdateMyReply, error) {
	var out UpdateMyReply
	pattern := "/v1/my/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMyUpdateMy))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}