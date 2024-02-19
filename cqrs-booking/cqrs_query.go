package cqrs_booking

import (
	"errors"
	"time"
)

type BookingPeriod struct {
	from time.Time
	to   time.Time
}

func NewPeriod(from time.Time, to time.Time) (BookingPeriod, error) {
	if from.After(to) {
		return BookingPeriod{}, errors.New(`invalid period`)
	}

	if from.Truncate(time.Hour*24) == to.Truncate(time.Hour*24) {
		return BookingPeriod{}, errors.New(`invalid period`)
	}

	return BookingPeriod{
		from: from,
		to:   to,
	}, nil
}

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

func (q QueryService) FreeRooms(period BookingPeriod) []RoomName {
	freeRooms := make(map[RoomName]bool)
	for _, name := range q.registry.Rooms {
		freeRooms[name] = true
	}

	for day := period.from; period.to.After(day); day = day.AddDate(0, 0, 1) {
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
