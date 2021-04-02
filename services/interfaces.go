package services

import "totality-corp-assignment/models"

type Repository interface {
	GetByID(id int) (*models.User, error)
}
