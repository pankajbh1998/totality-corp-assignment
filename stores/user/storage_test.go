package user

import (
	"testing"
	"totality-corp-assignment/models"
)

func TestUserStore_GetByID(t *testing.T) {
	store := New()
	res, err := store.GetByID(1)
	output := &models.User{
		ID:      1,
		FName:   "Pankaj",
		City:    "Greater Noida",
		Phone:   9650536254,
		Height:  5.9,
		Married: false,
	}
	if err != nil {
		if res != output {
			t.Errorf("Failed at %v\nExpected Output : %v\nActual Output   : %v\n", 1, output, res)
		}
	}
	res, err = store.GetByID(6)
	output = nil
	if err != nil {
		if res != output {
			t.Errorf("Failed at %v\nExpected Output : %v\nActual Output   : %v\n", 1, output, res)
		}
	}

}
