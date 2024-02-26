We want to make a booking solution for one hotel.

The first 2 user stories are :

As a user I want to see all free rooms.
As a user I want to book a room.
They want to use the CQRS pattern. To do that we will have :

One Command Service with a function bookARoom(Booking)
that calls the WriteRegistry
that notifies the ReadRegistry called by the Query Service
One Query Service with function Room[] freeRooms(arrival: Date, departure: Date)
The Booking struct contains

client id
room name
arrival date
departure date
And the Room struct contain only

room name


TODO:
 - [x] Тест на вторую границу (должны ли включать или исключать границу при поиске?)
 - [ ] Резервирование в прошлом - полная хуйня (Миша)
 - [x] refactoring from, to -> value object
 - [x] refactoring FreeRooms
 - [] refactoring Registry
 - [ ] Закончить рефакторинг FreeRooms
 - [ ] Продумать тест для нового типа RoomOccupancy
 - [ ] Протестировать что последней день не забронирован
 - [ ] И отрефакторить Test_booked_period ( ассерты в цикле написать )