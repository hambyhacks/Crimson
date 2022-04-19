package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	userreq "github.com/hambyhacks/CrimsonIMS/app/interface/users/requests"
	userresp "github.com/hambyhacks/CrimsonIMS/app/interface/users/responses"
	usersrv "github.com/hambyhacks/CrimsonIMS/service/users"
)

func MakeAddUserEndpoint(svc usersrv.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(userreq.AddUserRequest)
		msg, err := svc.AddUser(ctx, req.User)
		if err != nil {
			return userresp.AddUserResponse{Msg: "unable to process request", Err: err}, err
		}
		return userresp.AddUserResponse{Msg: msg, Err: nil}, nil
	}
}

func MakeGetUserByEmailEndpoint(svc usersrv.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(userreq.GetUserByEmailRequest)
		userDetails, err := svc.GetByEmail(ctx, req.Email)
		if err != nil {
			return userresp.GetUserByEmailResponse{User: nil, Err: err}, err
		}
		return userresp.GetUserByEmailResponse{User: userDetails, Err: nil}, nil
	}
}

func MakeUpdateUserEndpoint(svc usersrv.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(userreq.UpdateUserRequest)
		msg, err := svc.UpdateUser(ctx, req.User)
		if err != nil {
			return userresp.UpdateUserResponse{Msg: "unable to process request", Err: err}, err
		}
		return userresp.UpdateUserResponse{Msg: msg, Err: nil}, nil
	}
}
