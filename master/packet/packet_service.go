package packet

//
//import (
//	"context"
//	"database/sql"
//
//	"github.com/golang/protobuf/ptypes/empty"
//
//	foodproto "github.com/vivaldy22/eatnfit-food-service/proto"
//)
//
//type Service struct {
//	db *sql.DB
//}
//
//func (s *Service) GetAll(ctx context.Context, empty *empty.Empty) (*foodproto.DetailPacketList, error) {
//	var packets = new(foodproto.DetailPacketList)
//	rows, err := s.db.Query(queries.GET_ALL_LEVEL)
//
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	for rows.Next() {
//		var each = new(authproto.Level)
//		if err := rows.Scan(&each.LevelId, &each.LevelName, &each.LevelStatus); err != nil {
//			return nil, err
//		}
//		packets.List = append(packets.List, each)
//	}
//
//	if err = rows.Err(); err != nil {
//		return nil, err
//	}
//	return packets, nil
//}
//
//func (s *Service) GetByID(ctx context.Context, id *foodproto.ID) (*foodproto.DetailPacket, error) {
//	panic("implement me")
//}
//
//func (s *Service) Create(ctx context.Context, packet *foodproto.DetailPacket) (*foodproto.DetailPacket, error) {
//	panic("implement me")
//}
//
//func (s *Service) Update(ctx context.Context, request *foodproto.DetailPacketUpdateRequest) (*foodproto.DetailPacket, error) {
//	panic("implement me")
//}
//
//func (s *Service) Delete(ctx context.Context, id *foodproto.ID) (*empty.Empty, error) {
//	panic("implement me")
//}
//
//func NewService(db *sql.DB) foodproto.PacketCRUDServer {
//	return &Service{db}
//}
