package transaction

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/vivaldy22/eatnfit-food-service/tools/consts"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	foodproto "github.com/vivaldy22/eatnfit-food-service/proto"
	"github.com/vivaldy22/eatnfit-food-service/tools/queries"
)

type Service struct {
	db *sql.DB
}

func (s *Service) ConfirmTransaction(ctx context.Context, id *foodproto.ID) (*foodproto.Transaction, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare(queries.CONFIRM_TRANSACTION)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(id.Id)
	if err != nil {
		return nil, tx.Rollback()
	}

	stmt.Close()
	return &foodproto.Transaction{
		TransId: id.Id,
	}, tx.Commit()
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
			&each.StartDate, &each.StartTime, &each.Address, &each.PaymentId, &each.OrderStatus, &each.TransStatus); err != nil {
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
		&trans.StartDate, &trans.StartTime, &trans.Address, &trans.PaymentId, &trans.OrderStatus, &trans.TransStatus)
	if err != nil {
		return nil, err
	}
	return trans, nil
}

func (s *Service) GetByUserID(ctx context.Context, id *foodproto.ID) (*foodproto.TransactionList, error) {
	var transactions = new(foodproto.TransactionList)
	rows, err := s.db.Query(queries.GET_TRANS_BY_ID_USER, id.Id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var each = new(foodproto.Transaction)
		if err := rows.Scan(&each.TransId, &each.TransDate, &each.UserId, &each.PacketId, &each.Portion, &each.TotalPrice,
			&each.StartDate, &each.StartTime, &each.Address, &each.PaymentId, &each.OrderStatus, &each.TransStatus); err != nil {
			return nil, err
		}
		transactions.List = append(transactions.List, each)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return transactions, nil
}

func (s *Service) Create(ctx context.Context, trans *foodproto.Transaction) (*foodproto.Transaction, error) {
	tx, err := s.db.Begin()

	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare(queries.CREATE_TRANS_PACKET)

	if err != nil {
		return nil, err
	}

	id := uuid.New().String()
	dateNow := time.Now().Format(consts.DATE_FORMAT)
	_, err = stmt.Exec(id, dateNow, trans.UserId, trans.PacketId, trans.Portion, trans.PacketId,
		trans.Portion, trans.StartDate, trans.StartTime, trans.Address, trans.PaymentId)
	if err != nil {
		return nil, tx.Rollback()
	}

	trans.TransId = id
	trans.TransDate = dateNow
	stmt.Close()
	return trans, tx.Commit()
}

func (s *Service) Delete(ctx context.Context, id *foodproto.ID) (*empty.Empty, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return new(empty.Empty), err
	}

	stmt, err := tx.Prepare(queries.DELETE_TRANS)
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
