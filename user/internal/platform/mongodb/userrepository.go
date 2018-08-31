package mongodb

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/lukasjarosch/educonn-platform/user/internal/platform/config"
	"time"
	"github.com/lukasjarosch/educonn-platform/user/internal/platform/errors"
	"context"
	"github.com/opentracing/opentracing-go"
)

type UserRepository struct {
	session *mgo.Session
}

// NewUserRepository creates a new user repository
func NewUserRepository(host string, port string, user string, pass string, dbName string) (*UserRepository, error) {
	connString := fmt.Sprintf("%s:%s/%s", host, port, dbName)
	session, err := mgo.Dial(connString)
	if err != nil {
		return nil, err
	}
	return &UserRepository{
		session: session,
	}, nil
}

// CreateUser creates a new user entry
func (u *UserRepository) CreateUser(ctx context.Context, user *User) (*User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserRepository.CreateUser")
	defer span.Finish()

	session := u.session.Clone()
	defer session.Close()

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	user.ID = bson.NewObjectId()
	err := session.DB(config.DbName).C(config.UserCollection).Insert(user)
	if err != nil {
	    return nil, err
	}

	return user, nil
}

// FindByEmail retrieves a user by email
func (u *UserRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserRepository.FindByEmail")
	defer span.Finish()

	session := u.session.Clone()
	defer session.Close()

	var user User
	err := session.DB(config.DbName).C(config.UserCollection).Find(bson.M{"email": email}).One(&user)
	if err != nil {
		return nil, errors.UserNotFound
	}

	return &user, nil
}

// FindById searches  for an user by ID
func (u *UserRepository) FindById (ctx context.Context, id string) (*User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserRepository.FindByEmail")
	defer span.Finish()

	session := u.session.Clone()
	defer session.Close()

	var user User
	err := session.DB(config.DbName).C(config.UserCollection).FindId(bson.ObjectIdHex(id)).One(&user)
	if err != nil {
		return nil, errors.UserNotFound
	}

	return &user, nil
}

// GetAll returns all users
func (u *UserRepository) GetAll(ctx context.Context) ([]*User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserRepository.FindByEmail")
	defer span.Finish()

	session := u.session.Clone()
	defer session.Close()

	var users []*User
	err := session.DB(config.DbName).C(config.UserCollection).Find(bson.M{}).All(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserRepository) DeleteUser(ctx context.Context, id string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserRepository.FindByEmail")
	defer span.Finish()

	session := u.session.Clone()
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		return errors.MalformedUserId
	}

	err := session.DB(config.DbName).C(config.UserCollection).RemoveId(bson.ObjectIdHex(id))
	if err != nil {
		return err
	}

	return nil
}
