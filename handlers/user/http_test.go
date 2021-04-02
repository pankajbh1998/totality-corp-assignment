package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"totality-corp-assignment/errors"
	"totality-corp-assignment/handlers"
	"totality-corp-assignment/models"

	"github.com/gorilla/mux"

	"github.com/golang/mock/gomock"
)

func TestHandler_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	serviceUser := handlers.NewMockUserService(ctrl)
	handlerUser := New(serviceUser)

	testCases := []struct {
		id         string
		exOutput   interface{}
		statusCode int
	}{
		{
			id: "1",
			exOutput: &models.User{
				ID:      1,
				FName:   "Pankaj",
				City:    "Greater Noida",
				Phone:   9650536254,
				Height:  5.9,
				Married: false,
			},
			statusCode: http.StatusOK,
		},
		{
			id:         "abc",
			exOutput:   errors.ErrIdCantBeZeroOrNegative,
			statusCode: http.StatusBadRequest,
		},
		{
			id:         "1",
			exOutput:   errors.ErrContentNotFound,
			statusCode: http.StatusInternalServerError,
		},
	}

	for i, v := range testCases {
		if i == 0 {
			serviceUser.EXPECT().GetByID(v.exOutput.(*models.User).ID).Return(v.exOutput.(*models.User), nil)
		}
		if i == 2 {
			serviceUser.EXPECT().GetByID(1).Return(nil, v.exOutput.(error))
		}
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/users"+v.id, nil)
		r = mux.SetURLVars(r, map[string]string{"id": v.id})
		handlerUser.GetByID(w, r)
		result := w.Result()
		res, err := ioutil.ReadAll(result.Body)
		expectedOutput, _ := json.Marshal(v.exOutput)
		if err != nil {
			t.Fatalf(err.Error())
		} else if v.statusCode != result.StatusCode {
			t.Errorf("Failed at %v Wrong Status Code\n", i+1)
		} else if !reflect.DeepEqual(res, expectedOutput) {
			t.Errorf("Failed at %v\nExpected Output : %v\nActual Output   : %v\n", i+1, string(expectedOutput), string(res))
		}
	}
}

func TestHandler_GetByIDs(t *testing.T) {
	ctrl := gomock.NewController(t)
	serviceUser := handlers.NewMockUserService(ctrl)
	handlerUser := New(serviceUser)

	testCases := []struct {
		id         string
		exOutput   interface{}
		statusCode int
	}{
		{
			id: "[1]",
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
			statusCode: http.StatusOK,
		},
		{
			id:         "[abc]",
			exOutput:   errors.ErrIdCantBeZeroOrNegative,
			statusCode: http.StatusBadRequest,
		},
		{
			id:         "[1]",
			exOutput:   errors.ErrContentNotFound,
			statusCode: http.StatusInternalServerError,
		},
	}

	for i, v := range testCases {
		arr := []int{1}
		if i == 0 {
			serviceUser.EXPECT().GetByIDs(arr).Return(v.exOutput, nil)
		}
		if i == 2 {
			serviceUser.EXPECT().GetByIDs(arr).Return(nil, v.exOutput.(error))
		}
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/users"+v.id, nil)
		r = mux.SetURLVars(r, map[string]string{"ids": v.id})
		handlerUser.GetByIDs(w, r)
		result := w.Result()
		res, err := ioutil.ReadAll(result.Body)
		expectedOutput, _ := json.Marshal(v.exOutput)
		if err != nil {
			t.Fatalf(err.Error())
		} else if v.statusCode != result.StatusCode {
			t.Errorf("Failed at %v Wrong Status Code\n", i+1)
		} else if !reflect.DeepEqual(res, expectedOutput) {
			t.Errorf("Failed at %v\nExpected Output : %v\nActual Output   : %v\n", i+1, string(expectedOutput), string(res))
		}
	}
}
