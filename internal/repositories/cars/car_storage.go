package cars

import (
	"database/sql"
	"log"

	"github.com/M0rdovorot/effective_mobile/configs"
	"github.com/M0rdovorot/effective_mobile/internal/model"
	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/dbscan"
)

const (
	carsOnPage = 10
	CarTable = "public.car"
)

func SelectAllCarsSQL(filter map[string]any, limit int, offset int) squirrel.SelectBuilder {
	builder := squirrel.Select("*").
		From(CarTable + " c")
	where := squirrel.And{}

	for key, value := range filter {
		where = append(where, squirrel.Eq{key: value})
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

func InsertCarSQL(car model.Car) squirrel.InsertBuilder {
	columns := []string{"regNum", "mark", "model", "name", "surname"}
	values := []any{car.RegNum, car.Mark, car.Model, car.Owner.Name, car.Owner.Surname}
	if car.Patronymic != "" {
		columns = append(columns, "patronymic")
		values = append(values, car.Patronymic)
	}
	if car.Year != 0 {
		columns = append(columns, "year")
		values = append(values, car.Year)
	}
	return squirrel.Insert(CarTable).
		Columns(columns...).
		Values(values...).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(squirrel.Dollar)
}

func UpdateCarSQL(id int, patchMap map[string]any) squirrel.UpdateBuilder {
	return squirrel.Update(CarTable).
		SetMap(patchMap).
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar)
}

func DeleteCarSQL(carId int) squirrel.DeleteBuilder {
	return squirrel.Delete(CarTable).
		Where(squirrel.Eq{"id": carId}).
		PlaceholderFormat(squirrel.Dollar)
}

type CarStorage struct {
	db *sql.DB
	config *configs.Config
}

func CreateCarStorage(db *sql.DB, config *configs.Config) CarRepository {
	return &CarStorage{
		db: db,
		config: config,
	}
}

func (storage *CarStorage) GetCars(filter map[string]any, page int) ([]model.Car, error) {
	rows, err := SelectAllCarsSQL(filter, carsOnPage, (page - 1) * carsOnPage).RunWith(storage.db).Query()
	if err != nil {
		return []model.Car{}, err
	}
	defer rows.Close()

	var cars []model.Car
	if err = dbscan.ScanAll(&cars, rows); err != nil {
		return []model.Car{}, err
	}

	return cars, nil
}

func (storage *CarStorage) CreateCars(cars []model.Car) ([]int, error) {
	//begin transaction
	tx, err := storage.db.Begin()
	if err != nil {
		return []int{}, err
	}
	defer func() {
		if err != nil {
			err = tx.Rollback()
		} else {
			err = tx.Commit()
		}
		if err != nil {
			log.Println(err)
		}
	}()
	carIds := []int{}
	for _, car := range cars {
		var carId int

		query, args, err := InsertCarSQL(car).RunWith(storage.db).ToSql()
		if err != nil {
			return []int{}, err
		}
		err = tx.QueryRow(query, args...).Scan(&carId)
		if err != nil {
			return []int{}, err
		}
		carIds = append(carIds, carId)
	}

	return carIds, nil
}

func (storage *CarStorage) PatchCar(id int, patchMap map[string]any) (error) {
	res, err := UpdateCarSQL(id, patchMap).RunWith(storage.db).Exec()
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

func (storage *CarStorage) DeleteCar(carId int) (error) {
	res, err := DeleteCarSQL(carId).RunWith(storage.db).Exec()
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