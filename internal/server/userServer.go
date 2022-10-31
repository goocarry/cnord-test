package server

import (
	"context"
	"log"

	"github.com/goocarry/cnord-test/internal/store"
	"github.com/goocarry/cnord-test/proto/user"
)

// UserServer ...
type UserServer struct {
	log *log.Logger
	store *store.Store
	user.UnimplementedUserServiceServer
}

// NewUserServer ...
func NewUserServer(l *log.Logger, store *store.Store) *UserServer {
	return &UserServer{
		log: l,
		store: store,
	}
}

// SaveUser ...
func (us *UserServer) SaveUser(ctx context.Context, u *user.SaveUserRequest) (*user.SaveUserResponse, error) {
	us.log.Printf("Handle SaveUser, Firstname: %s, Lastname: %s", u.FirstName, u.LastName)

	id, err := us.store.User().Create(u.FirstName, u.LastName)
	if err != nil {
		return nil, err
	}

	resp := &user.SaveUserResponse{
		ID: id,
	}

	return resp, nil
}

// GetUserByID ...
func (us *UserServer) GetUserByID(ctx context.Context, u *user.UserRequest) (*user.UserResponse, error) {
	us.log.Printf("Handle GetUserByID, ID: %s", u.ID)

	getuser, err := us.store.User().GetByID(u.ID)
	if err != nil {
		return nil, err
	}

	resp := &user.UserResponse{
		FirstName: getuser.FirstName,
		LastName: getuser.LastName,
	}

	return resp, nil
}
