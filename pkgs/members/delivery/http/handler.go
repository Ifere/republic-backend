package memberDelivery

import (
	"net/http"
	"republic-backend/config/responses"
	"republic-backend/models"
	memberusecase "republic-backend/pkgs/members/usecase"
	httplib "republic-backend/utils/http"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

type MemberDelivery struct {
	MemberUsecase memberusecase.MemberUsecaseI
}

func (b MemberDelivery) CreateMember(rw http.ResponseWriter, r *http.Request) {
	var member models.MemberModel
	payload := httplib.C{
		W: rw,
		R: r,
	}
	payload.BindJSON(&member)
	member, err := b.MemberUsecase.CreateMember(member)
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
			Data:    member,
			Message: "member created",
		}
		payload.Response(resp)

	}
}

func (b MemberDelivery) GetMemberById(rw http.ResponseWriter, r *http.Request) {
	payload := httplib.C{
		W: rw,
		R: r,
	}
	ID := payload.Params("memberID")

	member, err := b.MemberUsecase.GetMemberById(ID)
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
		Data:    member,
		Message: "member returned",
	}
	payload.Response(resp)
}

func (b MemberDelivery) FetchMembers(rw http.ResponseWriter, r *http.Request) {
	filter := bson.D{{}}
	payload := httplib.C{
		W: rw,
		R: r,
	}
	members, err := b.MemberUsecase.FetchMembers(filter)
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
		Data:    members,
		Message: "members returned",
	}
	payload.Response(resp)
}

func (b MemberDelivery) UpdateMemberDetails(rw http.ResponseWriter, r *http.Request) {
	var member models.MemberModel
	payload := httplib.C{
		W: rw,
		R: r,
	}
	payload.BindJSON(&member)
	ID := payload.Params("memberID")
	member, err := b.MemberUsecase.UpdateMemberDetails(ID, member)
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
		Data:    member,
		Message: "member details updated",
	}
	payload.Response(resp)
}

func (b MemberDelivery) DeleteMember(rw http.ResponseWriter, r *http.Request) {
	payload := httplib.C{
		W: rw,
		R: r,
	}
	ID := payload.Params("memberID")
	err := b.MemberUsecase.DeleteMember(ID)
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
		Message: "member deleted",
	}
	payload.Response(resp)
}

func NewMembersDelivery(zUsecase memberusecase.MemberUsecaseI) MemberDelivery {
	return MemberDelivery{zUsecase}
}