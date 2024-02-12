package cqrs_booking_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	cqrs_booking "underground-katas/cqrs-booking"
)

func Test_no_free_room_available(t *testing.T) {
	query := cqrs_booking.NewQueryService(cqrs_booking.ReadRegistry{})
	rooms := query.FreeRooms(time.Now(), time.Now())

	assert.Empty(t, rooms)
}

func Test_free_room_available(t *testing.T) {
	registry := cqrs_booking.ReadRegistry{Rooms: []cqrs_booking.RoomName{"room1"}}
	query := cqrs_booking.NewQueryService(registry)
	roomsFree := query.FreeRooms(time.Now(), time.Now())

	assert.Equal(t, registry.Rooms, roomsFree)
}

func Test_free_room_check_with_date(t *testing.T) {
	registry := cqrs_booking.ReadRegistry{
		Rooms: []cqrs_booking.RoomName{"room1", "room2", "room3", "room4"},
	}
	arrival, _ := time.Parse(time.DateOnly, "2024-02-12")
	departure, _ := time.Parse(time.DateOnly, "2024-02-13")
	query := cqrs_booking.NewQueryService(registry)

	roomsFree := query.FreeRooms(arrival, departure)

	assert.Equal(t, []cqrs_booking.RoomName{"room1"}, roomsFree)
}
