package users

import (
	"context"

	"github.com/jackmcguire1/go-twitch-analytics/configuration"
	"github.com/jackmcguire1/go-twitch-analytics/internal/aws/dynamo"
	"github.com/nicklaw5/helix"
)

// Users user service to interact with a user
type Users struct {
	table *dynamo.Table
}

// Client initialises users service
func Client(tableName string) *Users {
	return &Users{
		table: dynamo.Client(env.UsersTable),
	}
}

// User User object
type User struct {
	UserID          string                       `json:"id"`
	UserAccessToken *helix.UserAccessCredentials `json:"user_access_token"`
}

// Put store user in dynamo
func (u *Users) Put(
	ctx context.Context,
	input *User,
) (
	err error,
) {
	_, err = u.table.Put(ctx, input)
	return
}

// Get get user from dynamo
func (u *Users) Get(
	ctx context.Context,
	userID string,
) (
	user *User,
	err error,
) {
	user = &User{}

	err = u.table.Get(ctx, userID, &user)

	return
}
