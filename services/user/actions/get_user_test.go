package actions_test

import (
	"context"

	"connectrpc.com/connect"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/mock"

	proto "github.com/sweetloveinyourheart/planning-pocker/proto/code/userserver/go"
	"github.com/sweetloveinyourheart/planning-pocker/services/user/actions"
	"github.com/sweetloveinyourheart/planning-pocker/services/user/models"
)

func (as *ActionsSuite) Test_GetUserByID_NoUserWasFound() {
	as.setupEnvironment()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	as.mockUserRepository.On("GetUserByID", mock.Anything, mock.Anything).Return(models.User{}, false, nil)

	userID := uuid.Must(uuid.NewV7())

	actions := actions.NewActions(ctx, "test")
	resp, err := actions.GetUser(
		ctx,
		connect.NewRequest(
			&proto.GetUserRequest{
				UserId: userID.String(),
			},
		),
	)

	as.NoError(err)
	as.Nil(resp)

	as.mockUserRepository.AssertExpectations(as.T())
}