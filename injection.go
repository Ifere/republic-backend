package main

import (
	"republic-backend/config/db"
	memberDelivery "republic-backend/pkgs/members/delivery/http"
	memberRepo2 "republic-backend/pkgs/members/repository"
	memberUsecase "republic-backend/pkgs/members/usecase"
	userDelivery "republic-backend/pkgs/users/delivery/http"
	userRepo "republic-backend/pkgs/users/repository"
	userUsecase "republic-backend/pkgs/users/usecase"
	zoneDelivery "republic-backend/pkgs/zones/delivery/http"
	zoneRepo "republic-backend/pkgs/zones/repository"
	zoneUsecase "republic-backend/pkgs/zones/usecase"
)

func InitializeUser(conn db.MongoDB) userDelivery.UserDelivery {
	zRepo := zoneRepo.NewZoneRepo(conn)
	uRepo := userRepo.NewUserRepo(conn)
	usecase := userUsecase.NewUserUsecase(uRepo, zRepo)
	delivery := userDelivery.NewUsersDelivery(usecase)
	return delivery
}

func InitializeZone(conn db.MongoDB) zoneDelivery.ZoneDelivery {
	zRepo := zoneRepo.NewZoneRepo(conn)
	mRepo := memberRepo2.NewMemberRepo(conn)
	usecase := zoneUsecase.NewZoneUsecase(zRepo, mRepo)
	delivery := zoneDelivery.NewZoneDelivery(usecase)
	return delivery
}

func InitializeMember(conn db.MongoDB) memberDelivery.MemberDelivery {
	mRepo := memberRepo2.NewMemberRepo(conn)
	zRepo := zoneRepo.NewZoneRepo(conn)
	usecase := memberUsecase.NewMemberUsecase(mRepo, zRepo)
	delivery := memberDelivery.NewMembersDelivery(usecase)
	return delivery
}
