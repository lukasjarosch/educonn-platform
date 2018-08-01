package service

import (
	"context"
	pb "github.com/lukasjarosch/educonn-platform/user/proto"
)

type UserApiService struct {
}

func NewUserApiService() *UserApiService {
	return &UserApiService{}
}

func (u *UserApiService) Create(ctx context.Context, req *pb.UserDetails, res *pb.UserResponse) error {

	res.User = &pb.UserDetails{
		FirstName: "Hans",
	}
	return nil
}
