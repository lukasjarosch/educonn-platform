package mongodb

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/lukasjarosch/educonn-platform/user/internal/platform/config"
	"time"
	"github.com/lukasjarosch/educonn-platform/user/internal/platform/errors"
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
func (u *UserRepository) CreateUser(user *User) (*User, error) {
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
func (u *UserRepository) FindByEmail(email string) (*User, error) {
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
func (u *UserRepository) FindById (id string) (*User, error) {
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
func (u *UserRepository) GetAll() ([]*User, error) {
	session := u.session.Clone()
	defer session.Close()

	var users []*User
	err := session.DB(config.DbName).C(config.UserCollection).Find(bson.M{}).All(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserRepository) DeleteUser(id string) error {
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
