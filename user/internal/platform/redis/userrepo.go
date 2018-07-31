package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/lukasjarosch/educonn-platform/user/proto"
)

type UserRepository struct {
	redisDialString string
}

// NewUserRepository creates a new user redis repo
func NewUserRepository(redisDialString string) *UserRepository {
	return &UserRepository{redisDialString: redisDialString}
}

func (u *UserRepository) GetUserDetails(id string) (user *educonn_user.UserDetails, err error) {
	c, err := redis.Dial("tcp", u.redisDialString)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	userKey := fmt.Sprintf("user:%s", id)
	res, err := redis.Values(c.Do("HGETALL", userKey))
	if err != nil {
		return nil, err
	}

	var userDetails redisUserDetails
	err = redis.ScanStruct(res, &userDetails)
	if err != nil {
		return nil, err
	}

	user = &educonn_user.UserDetails{
		Id:        userDetails.ID,
		FirstName: userDetails.FirstName,
		LastName:  userDetails.LastName,
		Email:     userDetails.Email,
		Password:  userDetails.Password,
	}

	return user, nil
}

type redisUserDetails struct {
	ID        string `redis:"id"`
	FirstName string `redis:"first_name"`
	LastName  string `redis:"last_name"`
	Email     string `redis:"email"`
	Password  string `redis:"password"`
}
