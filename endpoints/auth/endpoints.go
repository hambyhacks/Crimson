package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	authreq "github.com/hambyhacks/CrimsonIMS/app/interface/auth/requests"
	authresp "github.com/hambyhacks/CrimsonIMS/app/interface/auth/responses"
	authsrv "github.com/hambyhacks/CrimsonIMS/service/auth"
)

func MakeAddUserEndpoint(svc authsrv.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(authreq.AddUserRequest)
		msg, err := svc.AddUser(ctx, req.User)
		if err != nil {
			return authresp.AddUserResponse{Msg: "unable to process request", Err: err}, err
		}
		return authresp.AddUserResponse{Msg: msg, Err: nil}, nil
	}
}

func MakeGetUserByEmailEndpoint(svc authsrv.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(authreq.GetUserByEmailRequest)
		userDetails, err := svc.GetByEmail(ctx, req.Email)
		if err != nil {
			return authresp.GetUserByEmailResponse{User: nil, Err: err}, err
		}
		return authresp.GetUserByEmailResponse{User: userDetails, Err: nil}, nil
	}
}

func MakeUpdateUserEndpoint(svc authsrv.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(authreq.UpdateUserRequest)
		msg, err := svc.UpdateUser(ctx, req.User)
		if err != nil {
			return authresp.UpdateUserResponse{Msg: "unable to process request", Err: err}, err
		}
		return authresp.UpdateUserResponse{Msg: msg, Err: nil}, nil
	}
}
