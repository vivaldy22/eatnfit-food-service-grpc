package packet

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

func (s *Service) Create(ctx context.Context, packet *foodproto.DetailPacket) (*foodproto.DetailPacket, error) {
	panic("implement me")
}

func (s *Service) Update(ctx context.Context, request *foodproto.DetailPacketUpdateRequest) (*foodproto.DetailPacket, error) {
	panic("implement me")
}

func (s *Service) Delete(ctx context.Context, id *foodproto.ID) (*empty.Empty, error) {
	panic("implement me")
}

func NewService(db *sql.DB) foodproto.PacketCRUDServer {
	return &Service{db}
}
