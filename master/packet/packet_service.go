package packet

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

func (s *Service) GetAll(ctx context.Context, pagination *foodproto.Pagination) (*foodproto.PacketList, error) {
	var packets = new(foodproto.PacketList)
	page, _ := strconv.Atoi(pagination.Page)
	limit, _ := strconv.Atoi(pagination.Limit)
	offset := (page * limit) - limit
	query := fmt.Sprintf(queries.GET_ALL_PACKET, offset, pagination.Limit)
	rows, err := s.db.Query(query, "%"+pagination.Keyword+"%")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var each = new(foodproto.Packet)
		if err := rows.Scan(&each.PacketId, &each.PacketName, &each.PacketPrice, &each.PacketDesc, &each.PacketStatus); err != nil {
			return nil, err
		}
		packets.List = append(packets.List, each)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return packets, nil
}

func (s *Service) GetTotal(ctx context.Context, e *empty.Empty) (*foodproto.Total, error) {
	var total int
	row := s.db.QueryRow(queries.GET_TOTAL_PACKET)
	err := row.Scan(&total)
	if err != nil {
		return nil, err
	}
	return &foodproto.Total{TotalData: strconv.Itoa(total)}, nil
}

func (s *Service) GetByID(ctx context.Context, id *foodproto.ID) (*foodproto.DetailPacket, error) {
	var packet = new(foodproto.Packet)
	var foods = new(foodproto.FoodList)
	row := s.db.QueryRow(queries.GET_PACKET_BY_ID, id.Id)

	err := row.Scan(&packet.PacketId, &packet.PacketName, &packet.PacketPrice, &packet.PacketDesc, &packet.PacketStatus)
	if err != nil {
		return nil, err
	}

	rows, err := s.db.Query(queries.GET_FOODS_BY_PACKET_ID, id.Id)
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

	return &foodproto.DetailPacket{
		Packet:   packet,
		ListFood: foods.List,
	}, nil
}

func (s *Service) Create(ctx context.Context, packet *foodproto.DetailPacketInsert) (*foodproto.DetailPacketInsert, error) {
	tx, err := s.db.Begin()

	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare(queries.CREATE_PACKET)

	if err != nil {
		return nil, err
	}

	idPacket := uuid.New().String()
	_, err = stmt.Exec(idPacket, packet.Packet.PacketName, packet.Packet.PacketPrice, packet.Packet.PacketDesc)
	if err != nil {
		return nil, tx.Rollback()
	}

	packet.Packet.PacketId = idPacket

	for _, food := range packet.ListFood {
		stmt, err = tx.Prepare(queries.CREATE_DETAIL_PACKET)
		if err != nil {
			return nil, tx.Rollback()
		}

		idDetail := uuid.New().String()
		_, err = stmt.Exec(idDetail, idPacket, food.FoodId)
		if err != nil {
			return nil, tx.Rollback()
		}
	}
	stmt.Close()

	return &foodproto.DetailPacketInsert{
		Packet:   packet.Packet,
		ListFood: packet.ListFood,
	}, tx.Commit()
}

func (s *Service) Update(ctx context.Context, request *foodproto.DetailPacketUpdateRequest) (*foodproto.DetailPacketInsert, error) {
	tx, err := s.db.Begin()

	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare(queries.UPDATE_PACKET)

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(request.Packet.Packet.PacketName, request.Packet.Packet.PacketPrice,
		request.Packet.Packet.PacketDesc, request.Id.Id)
	if err != nil {
		return nil, tx.Rollback()
	}

	for _, food := range request.Packet.ListFood {
		stmt, err = tx.Prepare(queries.UPDATE_DETAIL_PACKET)
		if err != nil {
			return nil, tx.Rollback()
		}

		_, err = stmt.Exec(food.FoodId, request.Id.Id)
		if err != nil {
			return nil, tx.Rollback()
		}
	}
	stmt.Close()

	return &foodproto.DetailPacketInsert{
		Packet:   request.Packet.Packet,
		ListFood: request.Packet.ListFood,
	}, tx.Commit()
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

func NewService(db *sql.DB) foodproto.PacketCRUDServer {
	return &Service{db}
}
