package memberUsecase

import (
	apperrors "republic-backend/config/errors"

	"republic-backend/models"
	memberRepo "republic-backend/pkgs/members/repository"
	zoneRepo "republic-backend/pkgs/zones/repository"

	"github.com/pkg/errors"
)

type MemberUsecaseI interface {
	CreateMember(b models.MemberModel) (models.MemberModel, error)
	GetMemberById(ID string) (models.MemberModel, error)
	FetchMembers(filter interface{}) ([] models.MemberModel, error)
	UpdateMemberDetails(ID string, update models.MemberModel) (models.MemberModel, error)
	DeleteMember(ID string) error
}

type MemberUsecase struct {
	MemberRepo memberRepo.MemberRepositoryI
	ZoneRepo zoneRepo.ZoneRepositoryI
}

func NewMemberUsecase(memRepo memberRepo.MemberRepositoryI, zRepo zoneRepo.ZoneRepositoryI) MemberUsecase {
	return MemberUsecase{memRepo, zRepo}
}

func (b MemberUsecase) CreateMember(z models.MemberModel) (models.MemberModel, error) {
	z.SetMemberId()

	member, e := b.MemberRepo.CreateMember(z)

	if e != nil {
		return z, errors.Wrap(e, apperrors.NotCreated{Resource: "member"}.Error())
	}

	return member, nil
}

func (b MemberUsecase) GetMemberById(ID string) (models.MemberModel, error) {
	member, e := b.MemberRepo.GetMemberById(ID)
	if e != nil {
		return member, errors.Wrap(e, apperrors.ErrorGetting{Resource: "member"}.Error())
	}
	return member, nil
}

func (b MemberUsecase) FetchMembers(filter interface{}) ([]models.MemberModel, error) {
	members, e := b.MemberRepo.FetchMembers(filter)
	if e != nil {
		return members, errors.Wrap(e, apperrors.ErrorGetting{Resource: "members"}.Error())
	}
	return members, nil
}

func (b MemberUsecase) UpdateMemberDetails(ID string, update models.MemberModel) (models.MemberModel, error) {
	member, e := b.MemberRepo.UpdateMemberDetails(ID, update)
	if e != nil {
		return member, errors.Wrap(e, apperrors.ErrorUpdating{Resource: "member"}.Error())
	}
	return member, nil
}

func (b MemberUsecase) DeleteMember(ID string) error {
	e := b.MemberRepo.DeleteMember(ID)
	if e != nil {
		return errors.Wrap(e, apperrors.ErrorDeleting{Resource: "member"}.Error())
	}
	return nil
}