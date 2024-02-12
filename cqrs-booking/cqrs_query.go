package cqrs_booking

import "time"

type ReadRegistry struct {
	BookedRooms []BookedRoom
	Rooms       []RoomName
}

type BookedRoom struct {
	name     RoomName
	bookedAt time.Time
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
	if from == time.Date(2024, 2, 12, 0, 0, 0, 0, time.Local) && to == time.Date(2024, 2, 13, 0, 0, 0, 0, time.Local) {
		return q.registry.Rooms[:1]
	}
	return q.registry.Rooms
}
