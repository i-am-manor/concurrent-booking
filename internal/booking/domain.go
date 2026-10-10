package booking

import "errors"

var ErrSeatAlreadyBooked = errors.New("seat is already taken")

// represents a confirmed seat reservation.
type Booking struct {
	MovieID string
	SeatID  string
	UserID  string
	Status  string
}

type BookingStore interface {
	Book(b Booking) error
	ListBookings(movieID string) []Booking
}
