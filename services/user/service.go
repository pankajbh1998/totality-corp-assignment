package user

import (
	"totality-corp-assignment/models"
	"totality-corp-assignment/services"
)

type userService struct {
	repository services.Repository
}

func New(userRepo services.Repository) *userService {
	return &userService{repository: userRepo}
}

// GetByID method takes id as parameter and return details of user on the basis of id with corresponding error
func (user *userService) GetByID(id int) (*models.User, error) {
	return user.repository.GetByID(id)
}

// GetByIDs method takes ids as parameter and return details of users on the basis of ids with corresponding error
func (user *userService) GetByIDs(ids []int) ([]models.User, error) {
	var (
		usersDetails []models.User
	)
	for _, id := range ids {
		userDetails, err := user.repository.GetByID(id)
		if err != nil {
			return nil, err
		}
		usersDetails = append(usersDetails, *userDetails)
	}
	return usersDetails, nil
}
