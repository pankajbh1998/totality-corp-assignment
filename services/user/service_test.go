package user

import (
	"reflect"
	"testing"
	"totality-corp-assignment/errors"
	"totality-corp-assignment/models"
	"totality-corp-assignment/services"

	"github.com/golang/mock/gomock"
)

func TestUserService_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	storeUser := services.NewMockRepository(ctrl)
	serviceUser := New(storeUser)

	testCases := []struct {
		id       int
		exOutput *models.User
		err      error
	}{
		{
			id: 1,
			exOutput: &models.User{
				ID:      1,
				FName:   "Pankaj",
				City:    "Greater Noida",
				Phone:   9650536254,
				Height:  5.9,
				Married: false,
			},
		},
	}

	for i, v := range testCases {
		storeUser.EXPECT().GetByID(v.id).Return(v.exOutput, v.err)
		res, err := serviceUser.GetByID(v.id)

		if v.err != err {
			t.Errorf("Failed at %v\nExpected Error : %v\nActual Error   : %v\n", i+1, v.err, err)
		} else if !reflect.DeepEqual(res, v.exOutput) {
			t.Errorf("Failed at %v\nExpected Output : %v\nActual Output   : %v\n", i+1, v.exOutput, res)
		}
	}
}

func TestUserService_GetByIDs(t *testing.T) {
	ctrl := gomock.NewController(t)
	storeUser := services.NewMockRepository(ctrl)
	serviceUser := New(storeUser)

	testCases := []struct {
		id       []int
		exOutput []models.User
		err      error
	}{
		{
			id: []int{1},
			exOutput: []models.User{
				{
					ID:      1,
					FName:   "Pankaj",
					City:    "Greater Noida",
					Phone:   9650536254,
					Height:  5.9,
					Married: false,
				},
			},
		},
		{
			id:       []int{1},
			exOutput: nil,
			err:      errors.ErrContentNotFound,
		},
	}

	for i, v := range testCases {
		if v.err == nil {
			storeUser.EXPECT().GetByID(v.id[0]).Return(&v.exOutput[0], v.err)
		} else {
			storeUser.EXPECT().GetByID(v.id[0]).Return(nil, v.err)
		}
		res, err := serviceUser.GetByIDs(v.id)

		if v.err != err {
			t.Errorf("Failed at %v\nExpected Error : %v\nActual Error   : %v\n", i+1, v.err, err)
		} else if !reflect.DeepEqual(res, v.exOutput) {
			t.Errorf("Failed at %v\nExpected Output : %v\nActual Output   : %v\n", i+1, v.exOutput, res)
		}
	}
}
