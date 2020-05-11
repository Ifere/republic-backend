package zoneDelivery

import (
	"net/http"
	"republic-backend/config/responses"
	"republic-backend/models"
	zoneusecase "republic-backend/pkgs/zones/usecase"
	httplib "republic-backend/utils/http"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

type ZoneDelivery struct {
	ZoneUsecase zoneusecase.ZoneUsecaseI
}

func (b ZoneDelivery) CreateZone(rw http.ResponseWriter, r *http.Request) {
	var zone models.ZoneModel
	payload := httplib.C{
		W: rw,
		R: r,
	}
	payload.BindJSON(&zone)
	zone, err := b.ZoneUsecase.CreateZone(zone)
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
			Data:    zone,
			Message: "zone created",
		}
		payload.Response(resp)

	}
}

func (b ZoneDelivery) GetZoneById(rw http.ResponseWriter, r *http.Request) {
	payload := httplib.C{
		W: rw,
		R: r,
	}
	ID := payload.Params("zoneID")

	zone, err := b.ZoneUsecase.GetZoneById(ID)
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
		Data:    zone,
		Message: "zone returned",
	}
	payload.Response(resp)
}

func (b ZoneDelivery) FetchZones(rw http.ResponseWriter, r *http.Request) {
	filter := bson.D{{}}
	payload := httplib.C{
		W: rw,
		R: r,
	}
	zones, err := b.ZoneUsecase.FetchZones(filter)
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
		Data:    zones,
		Message: "zones returned",
	}
	payload.Response(resp)
}

func (b ZoneDelivery) UpdateZoneDetails(rw http.ResponseWriter, r *http.Request) {
	var zone models.ZoneModel
	payload := httplib.C{
		W: rw,
		R: r,
	}
	payload.BindJSON(&zone)
	ID := payload.Params("zoneID")
	zone, err := b.ZoneUsecase.UpdateZoneDetails(ID, zone)
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
		Data:    zone,
		Message: "zone details updated",
	}
	payload.Response(resp)
}

func (b ZoneDelivery) DeleteZone(rw http.ResponseWriter, r *http.Request) {
	payload := httplib.C{
		W: rw,
		R: r,
	}
	ID := payload.Params("zoneID")
	err := b.ZoneUsecase.DeleteZone(ID)
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
		Message: "zone deleted",
	}
	payload.Response(resp)
}

func NewZoneDelivery(zUsecase zoneusecase.ZoneUsecaseI) ZoneDelivery {
	return ZoneDelivery{zUsecase}
}
