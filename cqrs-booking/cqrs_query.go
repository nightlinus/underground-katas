package cqrs_booking

import (
	"errors"
	"time"
)

type BookingPeriod struct {
	from time.Time
	to   time.Time
}

func (bp BookingPeriod) AsRange() []time.Time {
	result := make([]time.Time, 0)
	for day := bp.from; bp.to.After(day); day = day.AddDate(0, 0, 1) {
		result = append(result, day)
	}

	return result
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

type (
	RoomName      string
	RoomOccupancy map[RoomName]bool
)

func CreateFreeRooms(roomNames []RoomName) RoomOccupancy {
	freeRooms := make(RoomOccupancy)
	for _, name := range roomNames {
		freeRooms[name] = true
	}

	return freeRooms
}

func (r *RoomOccupancy) MarkOccupied(roomName RoomName) {
	(*r)[roomName] = false
}

type QueryService struct {
	registry ReadRegistry
}

func NewQueryService(registry ReadRegistry) *QueryService {
	return &QueryService{
		registry: registry,
	}
}

func (r *RoomOccupancy) GetFreeRooms() []RoomName {
	result := make([]RoomName, 0)
	for room, available := range *r {
		if available {
			result = append(result, room)
		}
	}
	return result
}

func (q QueryService) FreeRooms(period BookingPeriod) []RoomName {
	rooms := CreateFreeRooms(q.registry.Rooms)
	for _, day := range period.AsRange() {
		for _, room := range q.registry.BookedRooms {
			if day == room.BookedAt {
				rooms.MarkOccupied(room.Name)
			}
		}
	}

	return rooms.GetFreeRooms()
}

type BookCommand struct {
	ClientID int
	Room     RoomName
	Period   BookingPeriod
}

func Book(command BookCommand) {

}
