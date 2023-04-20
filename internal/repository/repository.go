package repository

import "github.com/flintory5/bookings/internal/models"


type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) error
}