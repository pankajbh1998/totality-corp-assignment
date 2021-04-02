package user

import (
	"totality-corp-assignment/errors"
	"totality-corp-assignment/models"
)

type userStore struct{}

func New() *userStore {
	return &userStore{}
}

func (s *userStore) GetByID(id int) (*models.User, error) {
	userDetails := []models.User{
		{
			ID:      1,
			FName:   "Pankaj",
			City:    "Greater Noida",
			Phone:   9650536254,
			Height:  5.9,
			Married: false,
		},
		{
			ID:      2,
			FName:   "Steve",
			City:    "LA",
			Phone:   1234567890,
			Height:  5.8,
			Married: true,
		},
		{
			ID:      3,
			FName:   "Rudra",
			City:    "Noida",
			Phone:   9540754235,
			Height:  5.5,
			Married: false,
		},
		{
			ID:      4,
			FName:   "Shivam",
			City:    "Noida",
			Phone:   9876543210,
			Height:  5.6,
			Married: true,
		},
		{
			ID:      5,
			FName:   "Shubham",
			City:    "Delhi",
			Phone:   9891234560,
			Height:  5.7,
			Married: false,
		},
	}

	for _, v := range userDetails {
		if v.ID == id {
			return &v, nil
		}
	}

	return nil, errors.ErrContentNotFound

}
