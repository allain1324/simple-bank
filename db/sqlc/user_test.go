package db

import (
	"context"
	"testing"
	"time"

	"github.com/allain1324/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) (user User) {
	arg := CreateUserParams{
		Username:    util.RandomOwner(),
		HashedPassword: "secret",
		FullName:  util.RandomOwner(),
		Email: util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	account1 := createRandomUser(t)
	account2, err := testQueries.GetUser(context.Background(), account1.Username)

	require.NoError(t, err)
	require.NotEmpty(t, account1)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.Username, account2.Username)
	require.Equal(t, account1.HashedPassword, account2.HashedPassword)
	require.Equal(t, account1.FullName, account2.FullName)
	require.Equal(t, account1.Email, account2.Email)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}