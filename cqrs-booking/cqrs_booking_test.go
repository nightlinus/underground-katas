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
