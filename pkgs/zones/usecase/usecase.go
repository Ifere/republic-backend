package zoneUsecase

import (
	apperrors "republic-backend/config/errors"
	"republic-backend/models"
	memberRepo "republic-backend/pkgs/members/repository"
	zoneRepo "republic-backend/pkgs/zones/repository"
	"github.com/pkg/errors"
)

type ZoneUsecaseI interface {
	CreateZone(b models.ZoneModel) (models.ZoneModel, error)
	GetZoneById(ID string) (models.ZoneModel, error)
	FetchZones(filter interface{}) ([] models.ZoneModel, error)
	UpdateZoneDetails(ID string, update models.ZoneModel) (models.ZoneModel, error)
	DeleteZone(ID string) error
}

type ZoneUsecase struct {
	ZoneRepo zoneRepo.ZoneRepositoryI
	memberRepo.MemberRepositoryI
}

func NewZoneUsecase(zRepo zoneRepo.ZoneRepositoryI, memRepo memberRepo.MemberRepositoryI) ZoneUsecase {
	return ZoneUsecase{zRepo, memRepo}
}

func (b ZoneUsecase) CreateZone(z models.ZoneModel) (models.ZoneModel, error) {
	z.SetZoneId()

	zone, e := b.ZoneRepo.CreateZone(z)

	if e != nil {
		return z, errors.Wrap(e, apperrors.NotCreated{Resource: "zone"}.Error())
	}

	return zone, nil
}

func (b ZoneUsecase) GetZoneById(ID string) (models.ZoneModel, error) {
	zone, e := b.ZoneRepo.GetZoneById(ID)
	if e != nil {
		return zone, errors.Wrap(e, apperrors.ErrorGetting{Resource: "zone"}.Error())
	}
	return zone, nil
}

func (b ZoneUsecase) FetchZones(filter interface{}) ([]models.ZoneModel, error) {
	zones, e := b.ZoneRepo.FetchZones(filter)
	if e != nil {
		return zones, errors.Wrap(e, apperrors.ErrorGetting{Resource: "zone"}.Error())
	}
	return zones, nil
}

func (b ZoneUsecase) UpdateZoneDetails(ID string, update models.ZoneModel) (models.ZoneModel, error) {
	zones, e := b.ZoneRepo.UpdateZoneDetails(ID, update)
	if e != nil {
		return zones, errors.Wrap(e, apperrors.ErrorUpdating{Resource: "zone"}.Error())
	}
	return zones, nil
}

func (b ZoneUsecase) DeleteZone(ID string) error {
	e := b.ZoneRepo.DeleteZone(ID)
	if e != nil {
		return errors.Wrap(e, apperrors.ErrorDeleting{Resource: "zone"}.Error())
	}
	return nil
}
