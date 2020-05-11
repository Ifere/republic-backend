package userUsecase


import (
	apperrors "republic-backend/config/errors"
	"republic-backend/models"
	userRepo "republic-backend/pkgs/users/repository"
	zoneRepo "republic-backend/pkgs/zones/repository"
	"github.com/pkg/errors"
)

type UserUsecaseI interface {
	CreateUser(b models.UserModel) (models.UserModel, error)
	GetUserById(ID string) (models.UserModel, error)
	FetchUsers(filter interface{}) ([] models.UserModel, error)
	UpdateUserDetails(ID string, update models.UserModel) (models.UserModel, error)
	DeleteUser(ID string) error
}

type UserUsecase struct {
	UserRepo userRepo.UserRepositoryI
	ZoneRepo zoneRepo.ZoneRepositoryI
}

func NewUserUsecase(uRepo userRepo.UserRepositoryI, zRepo zoneRepo.ZoneRepositoryI) UserUsecase {
	return UserUsecase{uRepo, zRepo}
}

func (b UserUsecase) CreateUser(z models.UserModel) (models.UserModel, error) {
	z.SetUserID()

	user, e := b.UserRepo.CreateUser(z)

	if e != nil {
		return z, errors.Wrap(e, apperrors.NotCreated{Resource: "user"}.Error())
	}

	return user, nil
}

func (b UserUsecase) GetUserById(ID string) (models.UserModel, error) {
	user, e := b.UserRepo.GetUserById(ID)
	if e != nil {
		return user, errors.Wrap(e, apperrors.ErrorGetting{Resource: "user"}.Error())
	}
	return user, nil
}

func (b UserUsecase) FetchUsers(filter interface{}) ([]models.UserModel, error) {
	users, e := b.UserRepo.FetchUsers(filter)
	if e != nil {
		return users, errors.Wrap(e, apperrors.ErrorGetting{Resource: "users"}.Error())
	}
	return users, nil
}

func (b UserUsecase) UpdateUserDetails(ID string, update models.UserModel) (models.UserModel, error) {
	user, e := b.UserRepo.UpdateUserDetails(ID, update)
	if e != nil {
		return user, errors.Wrap(e, apperrors.ErrorUpdating{Resource: "user"}.Error())
	}
	return user, nil
}

func (b UserUsecase) DeleteUser(ID string) error {
	e := b.UserRepo.DeleteUser(ID)
	if e != nil {
		return errors.Wrap(e, apperrors.ErrorDeleting{Resource: "user"}.Error())
	}
	return nil
}