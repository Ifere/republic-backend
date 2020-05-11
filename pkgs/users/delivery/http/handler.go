package userDelivery

import (
	"net/http"
	"republic-backend/config/responses"
	"republic-backend/models"
	userusecase "republic-backend/pkgs/users/usecase"
	httplib "republic-backend/utils/http"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

type UserDelivery struct {
	UserUsecase userusecase.UserUsecaseI
}

func (b UserDelivery) CreateUser(rw http.ResponseWriter, r *http.Request) {
	var user models.UserModel
	payload := httplib.C{
		W: rw,
		R: r,
	}
	payload.BindJSON(&user)
	user, err := b.UserUsecase.CreateUser(user)
	if err != nil {
		resp := responses.GeneralResponse{
			Success: false,
			Data:    nil,
			Message: errors.Cause(err).Error(),
		}
		payload.Response(resp)
	} else {
		resp := responses.GeneralResponse{
			Success: true,
			Data:    user,
			Message: "user created",
		}
		payload.Response(resp)

	}
}

func (b UserDelivery) GetUserById(rw http.ResponseWriter, r *http.Request) {
	payload := httplib.C{
		W: rw,
		R: r,
	}
	ID := payload.Params("userID")

	user, err := b.UserUsecase.GetUserById(ID)
	if err != nil {
		resp := responses.GeneralResponse{
			Success: false,
			Data:    nil,
			Message: errors.Cause(err).Error(),
		}
		payload.Response(resp)
	}
	resp := responses.GeneralResponse{
		Success: true,
		Data:    user,
		Message: "user returned",
	}
	payload.Response(resp)
}

func (b UserDelivery) FetchUsers(rw http.ResponseWriter, r *http.Request) {
	filter := bson.D{{}}
	payload := httplib.C{
		W: rw,
		R: r,
	}
	users, err := b.UserUsecase.FetchUsers(filter)
	if err != nil {
		resp := responses.GeneralResponse{
			Success: false,
			Data:    nil,
			Message: errors.Cause(err).Error(),
		}
		payload.Response(resp)
	}
	resp := responses.GeneralResponse{
		Success: true,
		Data:    users,
		Message: "users returned",
	}
	payload.Response(resp)
}

func (b UserDelivery) UpdateUserDetails(rw http.ResponseWriter, r *http.Request) {
	var user models.UserModel
	payload := httplib.C{
		W: rw,
		R: r,
	}
	payload.BindJSON(&user)
	ID := payload.Params("userID")
	user, err := b.UserUsecase.UpdateUserDetails(ID, user)
	if err != nil {
		resp := responses.GeneralResponse{
			Success: false,
			Data:    nil,
			Message: errors.Cause(err).Error(),
		}
		payload.Response(resp)
	}
	resp := responses.GeneralResponse{
		Success: true,
		Data:    user,
		Message: "user details updated",
	}
	payload.Response(resp)
}

func (b UserDelivery) DeleteUser(rw http.ResponseWriter, r *http.Request) {
	payload := httplib.C{
		W: rw,
		R: r,
	}
	ID := payload.Params("userID")
	err := b.UserUsecase.DeleteUser(ID)
	if err != nil {
		resp := responses.GeneralResponse{
			Success: false,
			Data:    nil,
			Message: errors.Cause(err).Error(),
		}
		payload.Response(resp)
	}
	resp := responses.GeneralResponse{
		Success: true,
		Message: "user deleted",
	}
	payload.Response(resp)
}

func NewUsersDelivery(zUsecase userusecase.UserUsecaseI) UserDelivery {
	return UserDelivery{zUsecase}
}
