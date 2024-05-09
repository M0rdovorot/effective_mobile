package banners

import (
	model "github.com/M0rdovorot/effective_mobile/internal/model"
)

type BannerRepository interface{
	GetUserBanner(tagId int, featureId int) (model.Banner, error) 
	GetAllBanners(featureId int, tagId int, limit int, offset int) ([]model.Banner, error)
	CreateBanner(banner model.Banner) (int, error)
	PatchBanner(banner model.Banner) (error)
	DeleteBanner(bannerId int) (error) 
}