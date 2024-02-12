package cqrs_booking_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	cqrs_booking "underground-katas/cqrs-booking"
)

func Test_no_free_room_available(t *testing.T) {
	rooms := cqrs_booking.FreeRooms(time.Now(), time.Now())

	assert.Empty(t, rooms)
}
