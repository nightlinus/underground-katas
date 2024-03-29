package cqrs_booking_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	cqrs_booking "underground-katas/cqrs-booking"
)

func Test_no_free_room_available(t *testing.T) {
	query := cqrs_booking.NewQueryService(&cqrs_booking.ReadRegistry{})

	rooms := query.FreeRooms(period(time.Now().Add(-time.Hour*24), time.Now()))

	assert.Empty(t, rooms)
}

func Test_free_room_available(t *testing.T) {
	registry := cqrs_booking.ReadRegistry{Rooms: []cqrs_booking.RoomName{"room1"}}
	query := cqrs_booking.NewQueryService(&registry)
	roomsFree := query.FreeRooms(period(time.Now().Add(-time.Hour*24), time.Now()))

	assert.Equal(t, registry.Rooms, roomsFree)
}

func Test_free_room_check_with_date(t *testing.T) {
	registry := cqrs_booking.ReadRegistry{
		Rooms: []cqrs_booking.RoomName{"room1", "room2"},
		BookedRooms: []cqrs_booking.BookedRoom{
			{Name: "room2", BookedAt: day(2024, 2, 12)},
		},
	}
	arrival := day(2024, 2, 12)
	departure := day(2024, 2, 13)
	query := cqrs_booking.NewQueryService(&registry)

	roomsFree := query.FreeRooms(period(arrival, departure))

	assert.Equal(t, []cqrs_booking.RoomName{"room1"}, roomsFree)
}

func Test_no_free_room_when_room_booked_in_middle(t *testing.T) {
	registry := cqrs_booking.ReadRegistry{
		Rooms: []cqrs_booking.RoomName{"room1"},
		BookedRooms: []cqrs_booking.BookedRoom{
			{Name: "room1", BookedAt: day(2024, 2, 12)},
		},
	}
	arrival := day(2024, 2, 11)
	departure := day(2024, 2, 13)
	query := cqrs_booking.NewQueryService(&registry)

	roomsFree := query.FreeRooms(period(arrival, departure))

	assert.Empty(t, roomsFree)
}

func Test_free_room_when_booked_at_departure_date(t *testing.T) {
	registry := cqrs_booking.ReadRegistry{
		Rooms: []cqrs_booking.RoomName{"room1"},
		BookedRooms: []cqrs_booking.BookedRoom{
			{Name: "room1", BookedAt: day(2024, 2, 12)},
		},
	}
	arrival := day(2024, 2, 11)
	departure := day(2024, 2, 12)
	query := cqrs_booking.NewQueryService(&registry)

	roomsFree := query.FreeRooms(period(arrival, departure))

	assert.Equal(t, []cqrs_booking.RoomName{"room1"}, roomsFree)
}

func Test_room_is_booked_when_arrival_date_booked(t *testing.T) {
	registry := cqrs_booking.ReadRegistry{
		Rooms: []cqrs_booking.RoomName{"room1"},
		BookedRooms: []cqrs_booking.BookedRoom{
			{Name: "room1", BookedAt: day(2024, 2, 12)},
		},
	}
	arrival := day(2024, 2, 12)
	departure := day(2024, 2, 13)
	query := cqrs_booking.NewQueryService(&registry)

	roomsFree := query.FreeRooms(period(arrival, departure))

	assert.Empty(t, roomsFree)
}

func Test_invalid_booking_period_arrival_after_departure(t *testing.T) {
	arrival := day(2024, 2, 12)
	departure := day(2024, 2, 11)

	_, err := cqrs_booking.NewPeriod(arrival, departure)
	assert.Error(t, err)
}

func Test_invalid_booking_period_arrival_same_as_departure(t *testing.T) {
	arrival := day(2024, 2, 12)
	departure := day(2024, 2, 12)

	_, err := cqrs_booking.NewPeriod(arrival, departure)
	assert.Error(t, err)
}

func Test_booked_one_day(t *testing.T) {
	registry := cqrs_booking.ReadRegistry{
		Rooms: []cqrs_booking.RoomName{"room1"},
	}
	query := cqrs_booking.NewQueryService(&registry)
	cmd := cqrs_booking.BookCommand{
		ClientID: 1,
		Room:     cqrs_booking.RoomName("room1"),
		Period:   period(day(2024, 2, 12), day(2024, 2, 13)),
	}

	cqrs_booking.Book(cmd, &registry)

	assert.Empty(t, query.FreeRooms(cmd.Period))
}

func Test_booked_period(t *testing.T) {
	registry := cqrs_booking.ReadRegistry{
		Rooms: []cqrs_booking.RoomName{"room1"},
	}
	query := cqrs_booking.NewQueryService(&registry)
	cmd := cqrs_booking.BookCommand{
		ClientID: 1,
		Room:     cqrs_booking.RoomName("room1"),
		Period:   period(day(2024, 2, 12), day(2024, 2, 15)),
	}

	cqrs_booking.Book(cmd, &registry)

	assertRangeIsBooked(t, query, period(day(2024, 2, 12), day(2024, 2, 15)))
}

func Test_the_room_is_free_for_departure_date(t *testing.T) {
	registry := cqrs_booking.ReadRegistry{
		Rooms: []cqrs_booking.RoomName{"room1"},
	}
	query := cqrs_booking.NewQueryService(&registry)
	cmd := cqrs_booking.BookCommand{
		ClientID: 1,
		Room:     cqrs_booking.RoomName("room1"),
		Period:   period(day(2024, 2, 12), day(2024, 2, 15)),
	}

	cqrs_booking.Book(cmd, &registry)

	assert.NotEmpty(t, query.FreeRooms(period(day(2024, 2, 15), day(2024, 2, 16))))
}

func Test_WriteRegistryExists(t *testing.T) {
	registry := cqrs_booking.WriteRegistryMock{}
	cmd := cqrs_booking.BookCommand{
		ClientID: 1,
		Room:     cqrs_booking.RoomName("room1"),
		Period:   period(day(2024, 2, 12), day(2024, 2, 15)),
	}
	cqrs_booking.BookIntoWriteRegistry(cmd, &registry)

	assert.Equal(t, 1, registry.Called)
}

func day(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
}

func period(from time.Time, to time.Time) cqrs_booking.BookingPeriod {
	p, err := cqrs_booking.NewPeriod(from, to)
	if err != nil {
		panic(err)
	}

	return p
}

func assertRangeIsBooked(t *testing.T, query *cqrs_booking.QueryService, bookedPeriod cqrs_booking.BookingPeriod) {
	t.Helper()

	for _, d := range bookedPeriod.AsRange() {
		assert.Empty(t, query.FreeRooms(period(d, d.Add(time.Hour*24))))
	}
}
