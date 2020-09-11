package food

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/google/uuid"

	"github.com/golang/protobuf/ptypes/empty"
	foodproto "github.com/vivaldy22/eatnfit-food-service/proto"
	"github.com/vivaldy22/eatnfit-food-service/tools/queries"
)

type Service struct {
	db *sql.DB
}

func (s *Service) GetAll(ctx context.Context, pagination *foodproto.Pagination) (*foodproto.FoodList, error) {
	var foods = new(foodproto.FoodList)
	page, _ := strconv.Atoi(pagination.Page)
	limit, _ := strconv.Atoi(pagination.Limit)
	offset := (page * limit) - limit
	query := fmt.Sprintf(queries.GET_ALL_FOOD, offset, pagination.Limit)
	rows, err := s.db.Query(query, "%"+pagination.Keyword+"%")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var each = new(foodproto.Food)
		if err := rows.Scan(&each.FoodId, &each.FoodPortion, &each.FoodName, &each.FoodCalories, &each.FoodFat,
			&each.FoodCarbs, &each.FoodProtein, &each.FoodPrice, &each.FoodDesc, &each.FoodStatus); err != nil {
			return nil, err
		}
		foods.List = append(foods.List, each)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return foods, nil
}

func (s *Service) GetTotal(ctx context.Context, e *empty.Empty) (*foodproto.Total, error) {
	var total int
	row := s.db.QueryRow(queries.GET_TOTAL_FOOD)
	err := row.Scan(&total)
	if err != nil {
		return nil, err
	}
	return &foodproto.Total{TotalData: strconv.Itoa(total)}, nil
}

func (s *Service) GetByID(ctx context.Context, id *foodproto.ID) (*foodproto.Food, error) {
	var food = new(foodproto.Food)
	row := s.db.QueryRow(queries.GET_BY_ID_FOOD, id.Id)

	err := row.Scan(&food.FoodId, &food.FoodPortion, &food.FoodName, &food.FoodCalories, &food.FoodFat,
		&food.FoodCarbs, &food.FoodProtein, &food.FoodPrice, &food.FoodDesc, &food.FoodStatus)
	if err != nil {
		return nil, err
	}
	return food, nil
}

func (s *Service) Create(ctx context.Context, food *foodproto.Food) (*foodproto.Food, error) {
	tx, err := s.db.Begin()

	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare(queries.CREATE_FOOD)

	if err != nil {
		return nil, err
	}

	id := uuid.New().String()
	_, err = stmt.Exec(id, food.FoodPortion, food.FoodName, food.FoodCalories, food.FoodFat,
		food.FoodCarbs, food.FoodProtein, food.FoodPrice, food.FoodDesc)
	if err != nil {
		return nil, tx.Rollback()
	}

	food.FoodId = id
	stmt.Close()
	return food, tx.Commit()
}

func (s *Service) Update(ctx context.Context, request *foodproto.FoodUpdateRequest) (*foodproto.Food, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare(queries.UPDATE_FOOD)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(request.Food.FoodPortion, request.Food.FoodName, request.Food.FoodCalories, request.Food.FoodFat,
		request.Food.FoodCarbs, request.Food.FoodProtein, request.Food.FoodPrice, request.Food.FoodDesc, request.Id.Id)
	if err != nil {
		return nil, tx.Rollback()
	}

	stmt.Close()
	request.Food.FoodId = request.Id.Id
	return request.Food, tx.Commit()
}

func (s *Service) Delete(ctx context.Context, id *foodproto.ID) (*empty.Empty, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return new(empty.Empty), err
	}

	stmt, err := tx.Prepare(queries.DELETE_FOOD)
	if err != nil {
		return new(empty.Empty), err
	}

	_, err = stmt.Exec(id.Id)
	if err != nil {
		return new(empty.Empty), tx.Rollback()
	}

	stmt.Close()
	return new(empty.Empty), tx.Commit()
}

func NewService(db *sql.DB) foodproto.FoodCRUDServer {
	return &Service{db}
}
