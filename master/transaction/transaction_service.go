package transaction

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/golang/protobuf/ptypes/empty"
	foodproto "github.com/vivaldy22/eatnfit-food-service/proto"
	"github.com/vivaldy22/eatnfit-food-service/tools/queries"
)

type Service struct {
	db *sql.DB
}

func (s *Service) GetAll(ctx context.Context, pagination *foodproto.Pagination) (*foodproto.TransactionList, error) {
	var transactions = new(foodproto.TransactionList)
	page, _ := strconv.Atoi(pagination.Page)
	limit, _ := strconv.Atoi(pagination.Limit)
	offset := (page * limit) - limit
	query := fmt.Sprintf(queries.GET_ALL_TRANSACTION, offset, pagination.Limit)
	rows, err := s.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var each = new(foodproto.Transaction)
		if err := rows.Scan(&each.TransId, &each.TransDate, &each.UserId, &each.PacketId, &each.Portion, &each.TotalPrice,
			&each.StartDate, &each.StartTime, &each.Address, &each.PaymentId, &each.TransStatus); err != nil {
			return nil, err
		}
		transactions.List = append(transactions.List, each)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return transactions, nil
}

func (s *Service) GetTotal(ctx context.Context, empty *empty.Empty) (*foodproto.Total, error) {
	var total int
	row := s.db.QueryRow(queries.GET_TOTAL_TRANSACTION)
	err := row.Scan(&total)
	if err != nil {
		return nil, err
	}
	return &foodproto.Total{TotalData: strconv.Itoa(total)}, nil
}

func (s *Service) GetByTransID(ctx context.Context, id *foodproto.ID) (*foodproto.Transaction, error) {
	var trans = new(foodproto.Transaction)
	row := s.db.QueryRow(queries.GET_TRANS_BY_ID_TRANSACTION, id.Id)

	err := row.Scan(&trans.TransId, &trans.TransDate, &trans.UserId, &trans.PacketId, &trans.Portion, &trans.TotalPrice,
		&trans.StartDate, &trans.StartTime, &trans.Address, &trans.PaymentId, &trans.TransStatus)
	if err != nil {
		return nil, err
	}
	return trans, nil
}

func (s *Service) GetByUserID(ctx context.Context, id *foodproto.ID) (*foodproto.Transaction, error) {
	var trans = new(foodproto.Transaction)
	row := s.db.QueryRow(queries.GET_TRANS_BY_ID_USER, id.Id)

	err := row.Scan(&trans.TransId, &trans.TransDate, &trans.UserId, &trans.PacketId, &trans.Portion, &trans.TotalPrice,
		&trans.StartDate, &trans.StartTime, &trans.Address, &trans.PaymentId, &trans.TransStatus)
	if err != nil {
		return nil, err
	}
	return trans, nil
}

func (s *Service) Create(ctx context.Context, transaction *foodproto.Transaction) (*foodproto.Transaction, error) {
	panic("implement me")
}

func (s *Service) Update(ctx context.Context, request *foodproto.TransactionUpdateRequest) (*foodproto.Transaction, error) {
	panic("implement me")
}

func (s *Service) Delete(ctx context.Context, id *foodproto.ID) (*empty.Empty, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return new(empty.Empty), err
	}

	stmt, err := tx.Prepare(queries.DELETE_PACKET)
	if err != nil {
		return new(empty.Empty), err
	}

	_, err = stmt.Exec(id.Id)
	if err != nil {
		return new(empty.Empty), tx.Rollback()
	}

	stmt, err = tx.Prepare(queries.DELETE_DETAIL_PACKET)
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

func NewService(db *sql.DB) foodproto.TransactionCRUDServer {
	return &Service{db}
}
