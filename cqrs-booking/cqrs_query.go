package cqrs_booking

import "time"

type ReadRegistry struct {
	BookedRooms []BookedRoom
	Rooms       []RoomName
}

type BookedRoom struct {
	Name     RoomName
	BookedAt time.Time
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
	freeRooms := make(map[RoomName]bool)
	for _, name := range q.registry.Rooms {
		freeRooms[name] = true
	}

	for day := from; to.After(day); day = day.AddDate(0, 0, 1) {
		for _, room := range q.registry.BookedRooms {
			if day == room.BookedAt {
				freeRooms[room.Name] = false
			}
		}
	}

	result := make([]RoomName, 0)
	for room, available := range freeRooms {
		if available {
			result = append(result, room)
		}
	}

	return result
}
