package banners

import (
	"database/sql"

	"github.com/M0rdovorot/effective_mobile/configs"
	"github.com/M0rdovorot/effective_mobile/internal/model"
	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/dbscan"
)

func SelectBannerSQL(tagId int, featureId int) squirrel.SelectBuilder {
	return squirrel.Select("b.id, b.content, b.created_at, b.updated_at, b.is_active, b.feature_id").
		From(configs.BannerTable + " b").
		Join("public.banner_to_tag btt ON btt.banner_id = b.id").
		Where(squirrel.And{
			squirrel.Eq{"tag_id": tagId},
			squirrel.Eq{"feature_id": featureId},
		}).
		PlaceholderFormat(squirrel.Dollar)
}

func SelectAllBannersSQL(featureId int, tagId int, limit int, offset int) squirrel.SelectBuilder {
	builder := squirrel.Select("b.id, b.content, b.created_at, b.updated_at, b.is_active, b.feature_id").
		From(configs.BannerTable + " b")
	where := squirrel.And{}

	if tagId != 0 {
		builder = builder.Join("public.banner_to_tag btt ON btt.banner_id = b.id")
		where = append(where, squirrel.Eq{"tag_id": tagId})
	}
	if featureId != 0 {
			where = append(where, squirrel.Eq{"feature_id": featureId})
	} 
	if len(where) > 0 {
		builder = builder.Where(where)
	}
	if limit != 0 {
		builder = builder.Limit(uint64(limit))
	} 
	if offset != 0 {
		builder = builder.Offset(uint64(offset))
	}
	return builder.PlaceholderFormat(squirrel.Dollar)
}

func InsertBannerSQL(banner model.Banner) squirrel.InsertBuilder {
	return squirrel.Insert(configs.BannerTable).
		Columns("content", "is_active", "feature_id").
		Values(banner.JSONContent, banner.IsActive, banner.FeatureId).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(squirrel.Dollar)
}

func InsertTagToBannerSQL(tagId int, bannerId int) squirrel.InsertBuilder {
	return squirrel.Insert(configs.TagToBannerTable).
		Columns("tag_id", "banner_id").
		Values(tagId, bannerId).
		PlaceholderFormat(squirrel.Dollar)
}

func UpdateBannerSQL(banner model.Banner) squirrel.UpdateBuilder {
	return squirrel.Update(configs.BannerTable).
		SetMap(map[string]interface{}{
			"content": banner.JSONContent,
			"feature_id": banner.FeatureId,
			"is_active": banner.IsActive,
		}).
		Where(squirrel.Eq{"id": banner.Id}).
		PlaceholderFormat(squirrel.Dollar)
}

func DeleteTagsByBannerIdSQL(bannerId int) squirrel.DeleteBuilder {
	return squirrel.Delete(configs.TagToBannerTable).
		Where(squirrel.Eq{"banner_id": bannerId}).
		PlaceholderFormat(squirrel.Dollar)
}

func DeleteBannerSQL(bannerId int) squirrel.DeleteBuilder {
	return squirrel.Delete(configs.BannerTable).
		Where(squirrel.Eq{"id": bannerId}).
		// Suffix("RETURNING count(*)").
		PlaceholderFormat(squirrel.Dollar)
}

type BannerStorage struct {
	db *sql.DB
}

func CreateBannerStorage(db *sql.DB) BannerRepository {
	return &BannerStorage{
		db: db,
	}
}

func (storage *BannerStorage) GetUserBanner(tagId int, featureId int) (model.Banner, error) {
	rows, err := SelectBannerSQL(tagId, featureId).RunWith(storage.db).Query()
	if err != nil {
		return model.Banner{}, err
	}
	defer rows.Close()

	var banner []model.Banner
	if err = dbscan.ScanAll(&banner, rows); err != nil {
		return model.Banner{}, err
	} 

	if len(banner) == 0 {
		return model.Banner{}, ErrNotFound
	} 

	return banner[0], nil
}

func (storage *BannerStorage) GetAllBanners(featureId int, tagId int, limit int, offset int) ([]model.Banner, error) {
	rows, err := SelectAllBannersSQL(featureId, tagId, limit, offset).RunWith(storage.db).Query()
	if err != nil {
		return []model.Banner{}, err
	}
	defer rows.Close()

	var banners []model.Banner
	if err = dbscan.ScanAll(&banners, rows); err != nil {
		return []model.Banner{}, err
	}

	return banners, nil
}

func (storage *BannerStorage) CreateBanner(banner model.Banner) (int, error) {
	var bannerId int
	err := InsertBannerSQL(banner).RunWith(storage.db).QueryRow().Scan(&bannerId)
	if err != nil {
		return 0, err
	}

	for _, tagId := range banner.TagIds {
		rows, err := InsertTagToBannerSQL(tagId, bannerId).RunWith(storage.db).Query()
		if err != nil {
			return 0, err
		}
		defer rows.Close()
	}

	return bannerId, nil
}

func (storage *BannerStorage) PatchBanner(banner model.Banner) (error) {
	res, err := UpdateBannerSQL(banner).RunWith(storage.db).Exec()
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return ErrNotFound
	}

	_, err = DeleteTagsByBannerIdSQL(banner.Id).RunWith(storage.db).Exec()
	if err != nil {
		return err
	}
	
	for _, tagId := range banner.TagIds {
		rows, err := InsertTagToBannerSQL(tagId, banner.Id).RunWith(storage.db).Query()
		if err != nil {
			return err
		}
		defer rows.Close()
	} 
	return nil
}

func (storage *BannerStorage) DeleteBanner(bannerId int) (error) {
	res, err := DeleteBannerSQL(bannerId).RunWith(storage.db).Exec()
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return ErrNotFound
	}
	return nil
}