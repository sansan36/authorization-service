package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"
	userv1 "github.com/sansan36/authorization-service/gen/authorization/v1"
	"github.com/sansan36/authorization-service/gen/authorization/v1/authorizationv1connect"
	commonv1 "github.com/sansan36/authorization-service/gen/common/v1"
	"github.com/sansan36/authorization-service/libs/auth"
	"github.com/sansan36/authorization-service/repository"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type userServiceHandler struct {
	authorizationv1connect.UnimplementedAuthorizationServiceHandler
	Repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) *userServiceHandler {
	return &userServiceHandler{
		Repo: repo,
	}
}

func (h *userServiceHandler) GetUserList(ctx context.Context, req *connect.Request[userv1.GetUserListRequest]) (res *connect.Response[userv1.GetUserListResponse], err error) {
	users, pagination, err := h.Repo.GetUserList(ctx, &userv1.User{}, req.Msg)
	if err != nil {
		return
	}

	res = connect.NewResponse(&userv1.GetUserListResponse{
		Users:      users,
		Pagination: pagination,
		HttpStatus: &commonv1.StandardResponse{
			Status: "success",
			Code:   http.StatusOK,
		},
	})

	return
}

func (h *userServiceHandler) GetUser(ctx context.Context, req *connect.Request[userv1.GetUserRequest]) (res *connect.Response[userv1.GetUserResponse], err error) {
	payload := req.Msg
	user, err := h.Repo.GetUser(ctx, &userv1.User{
		Id: payload.Id,
	})
	if err != nil {
		return
	}

	res = connect.NewResponse(&userv1.GetUserResponse{
		User: user,
		HttpStatus: &commonv1.StandardResponse{
			Status: "success",
			Code:   http.StatusOK,
		},
	})

	return
}

func (h *userServiceHandler) AddUser(ctx context.Context, req *connect.Request[userv1.AddUserRequest]) (res *connect.Response[userv1.AddUserResponse], err error) {
	payload := req.Msg

	v, err := protovalidate.New()
	if err != nil {
		return
	}
	err = v.Validate(payload)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	hashedPass, err := auth.HashPassword(payload.Password)
	if err != nil {
		return
	}

	// check if userName or email allready exist
	if err = checkUserDuplication(ctx, h, &userv1.User{UserName: payload.UserName, Email: payload.Email}); err != nil {
		return
	}

	user, err := h.Repo.AddUser(ctx, &userv1.User{
		UserName:       payload.UserName,
		Email:          payload.Email,
		Name:           payload.Name,
		HashedPassword: hashedPass,
		CreatedAt:      timestamppb.Now(),
	})
	if err != nil {
		return
	}

	res = connect.NewResponse(&userv1.AddUserResponse{
		User: user,
		HttpStatus: &commonv1.StandardResponse{
			Status: "success",
			Code:   http.StatusOK,
		},
	})

	return
}

func (h *userServiceHandler) EditUser(ctx context.Context, req *connect.Request[userv1.EditUserRequest]) (res *connect.Response[userv1.EditUserResponse], err error) {
	payload := req.Msg
	user, err := h.Repo.GetUser(ctx, &userv1.User{
		Id: payload.Id,
	})
	if err != nil {
		return
	}

	// check if userName or email allready exist
	if err = checkUserDuplication(ctx, h, &userv1.User{UserName: payload.UserName, Email: payload.Email}); err != nil {
		return
	}

	user, err = h.Repo.EditUser(ctx, &userv1.User{
		Id:             payload.Id,
		UserName:       payload.UserName,
		Email:          payload.Email,
		Name:           payload.Name,
		HashedPassword: user.HashedPassword,
		CreatedAt:      user.CreatedAt,
	})
	if err != nil {
		return
	}

	res = connect.NewResponse(&userv1.EditUserResponse{
		User: user,
		HttpStatus: &commonv1.StandardResponse{
			Status: "success",
			Code:   http.StatusOK,
		},
	})

	return
}

func (h *userServiceHandler) RemoveUser(ctx context.Context, req *connect.Request[userv1.RemoveUserRequest]) (res *connect.Response[userv1.RemoveUserResponse], err error) {
	payload := req.Msg
	user, err := h.Repo.GetUser(ctx, &userv1.User{
		Id: payload.Id,
	})
	if err != nil {
		return
	}

	_, err = h.Repo.RemoveUser(ctx, &userv1.User{
		Id: user.Id,
	})
	if err != nil {
		return
	}

	res = connect.NewResponse(&userv1.RemoveUserResponse{
		Message: fmt.Sprintf(`User with ID:%v has been removed`, payload.Id),
		HttpStatus: &commonv1.StandardResponse{
			Status: "success",
			Code:   http.StatusOK,
		},
	})

	return
}

// function to check userName or email allready exist
func checkUserDuplication(ctx context.Context, h *userServiceHandler, user *userv1.User) error {
	// check if userName already exist
	if _, err := h.Repo.GetUser(ctx, &userv1.User{
		UserName: user.UserName,
	}); err == nil {
		return connect.NewError(connect.CodeAlreadyExists, errors.New("username already exist"))
	}

	// check if email already exist
	if _, err := h.Repo.GetUser(ctx, &userv1.User{
		Email: user.Email,
	}); err == nil {
		return connect.NewError(connect.CodeAlreadyExists, errors.New("email already exist"))
	}

	return nil
}
