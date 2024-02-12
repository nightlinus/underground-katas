package cqrs_booking

import "time"

type ReadRegistry struct {
	Rooms []RoomName
}

type RoomName string

type QueryService struct {
	registry ReadRegistry
}

func NewQueryService(registry ReadRegistry) *QueryService {
	return &QueryService{
		registry: registry,
	}
}

func (q QueryService) FreeRooms(from time.Time, to time.Time) []RoomName {
	return q.registry.Rooms
}
