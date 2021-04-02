package handlers

import "totality-corp-assignment/models"

// UserService interface contains methods of user service so that we can consume them in handler
type UserService interface {
	GetByID(id int) (*models.User, error)
	GetByIDs(ids []int) ([]models.User, error)
}
