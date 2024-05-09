package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/M0rdovorot/effective_mobile/internal/model"
	bannerrep "github.com/M0rdovorot/effective_mobile/internal/repositories/banners"
	"github.com/M0rdovorot/effective_mobile/internal/usecases/authorization"
	ctxusecase "github.com/M0rdovorot/effective_mobile/internal/usecases/context"
	"github.com/gomodule/redigo/redis"
)

type BannerHandler struct {
	Banners bannerrep.BannerRepository
	Cash bannerrep.CashRepository
}

func CreateBannerHandler(banners bannerrep.BannerRepository, cash bannerrep.CashRepository) *BannerHandler {
	return &BannerHandler{
		Banners: banners,
		Cash: cash,
	}
}

func (handler *BannerHandler) GetUserBanner(ctx context.Context, form EmptyForm) (map[string]interface{}, int, error) {
	isAdmin, err := authorization.AuthorizeUserCtx(ctx)
	if err != nil {
		return nil, 0, err
	}

	vars := ctxusecase.GetQueryVars(ctx)
	if vars == nil {
		return nil, 0, ErrNoVars
	}

	tagId, err := strconv.Atoi(vars["tag_id"])
	if err != nil {
		return nil, 0, ErrBadTagID
	}
	featureId, err := strconv.Atoi(vars["feature_id"])
	if err != nil {
		return nil, 0, ErrBadFeatureID
	}

	if vars["use_last_revision"] == "false" || vars["use_last_revision"] == "" {
		log.Println("go to cash")
		JSONContent, isActive, err := handler.Cash.GetUserBanner(tagId, featureId)
		if err != nil && err != redis.ErrNil {
			return nil, 0, err
		}
		if err != redis.ErrNil {
			log.Println("returning from cash")
			var content map[string]interface{}
			err = json.Unmarshal([]byte(JSONContent), &content)
			if err != nil {
				return nil, 0, err
			}
			if !isAdmin && !isActive {
				return nil, 0, bannerrep.ErrForbidden
			}
			return content, http.StatusOK, nil
		}
	}

	log.Println("go to database")
	banner, err := handler.Banners.GetUserBanner(tagId, featureId)
	if err != nil {
		return nil, 0, err
	}
	err = json.Unmarshal([]byte(banner.JSONContent), &banner.Content)
	if err != nil {
		return nil, 0, err
	}

	if vars["use_last_revision"] == "false" || vars["use_last_revision"] == "" {
		log.Println("save to cash")
		err = handler.Cash.CreateBanner(banner, tagId)
		if err != nil {
			return nil, 0, err
		}
	}
	
	if !isAdmin && !banner.IsActive {
		return nil, 0, bannerrep.ErrForbidden
	}

	return banner.Content, http.StatusOK, nil
}

func (handler *BannerHandler) GetAllBanners(ctx context.Context, form EmptyForm) ([]model.Banner, int, error) {
	err := authorization.AuthorizeAdminCtx(ctx)
	if err != nil {
		return []model.Banner{}, 0, err
	}

	vars := ctxusecase.GetQueryVars(ctx)
	if vars == nil {
		return []model.Banner{}, 0, ErrNoVars
	}

	tagId, err := strconv.Atoi(vars["tag_id"])
	if err != nil {
		tagId = 0
	}
	featureId, err := strconv.Atoi(vars["feature_id"])
	if err != nil {
		featureId = 0
	}
	limit, err := strconv.Atoi(vars["limit"])
	if err != nil {
		limit = 0
	}
	offset, err := strconv.Atoi(vars["offset"])
	if err != nil {
		offset = 0
	}


	banners, err := handler.Banners.GetAllBanners(featureId, tagId, limit, offset)
	if err!=nil {
		return []model.Banner{}, 0, err
	}

	for  i := range banners {
		err = json.Unmarshal([]byte(banners[i].JSONContent), &banners[i].Content)
		if err != nil {
			return []model.Banner{}, 0, err
		}
	}

	return banners, http.StatusOK ,nil
}

func (handler *BannerHandler) CreateBanner(ctx context.Context, banner model.Banner) (int, int, error) {
	err := authorization.AuthorizeAdminCtx(ctx)
	if err != nil {
		return 0, 0, err
	}

	content, err := json.Marshal(banner.Content)
	if err != nil {
		return 0, 0, err
	}
	banner.JSONContent = string(content)
	
	id, err := handler.Banners.CreateBanner(banner)
	if err != nil {
		return 0, 0, err
	}

	return id, http.StatusCreated ,nil
}

func (handler *BannerHandler) PatchBanner(ctx context.Context, banner model.Banner) (any, int, error) {
	err := authorization.AuthorizeAdminCtx(ctx)
	if err != nil {
		return 0, 0, err
	}

	vars := ctxusecase.GetVars(ctx)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return 0, 0, ErrBadID
	}
	banner.Id = id

	content, err := json.Marshal(banner.Content)
	if err != nil {
		return 0, 0, err
	}
	banner.JSONContent = string(content)

	return 0, http.StatusOK, handler.Banners.PatchBanner(banner) 
}

func (handler *BannerHandler) DeleteBanner(ctx context.Context, form EmptyForm) (any, int, error) {
	err := authorization.AuthorizeAdminCtx(ctx)
	if err != nil {
		return 0, 0, err
	}

	vars := ctxusecase.GetVars(ctx)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return 0, 0, ErrBadID
	}

	return 0, http.StatusNoContent, handler.Banners.DeleteBanner(id)
}
