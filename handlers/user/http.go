package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"totality-corp-assignment/errors"
	services "totality-corp-assignment/handlers"

	"github.com/gorilla/mux"
)

type handler struct {
	service services.UserService
}

func New(userService services.UserService) *handler {
	return &handler{service: userService}
}

func validateId(id string) (int, error) {
	numId, _ := strconv.Atoi(id)
	if numId <= 0 {
		return 0, errors.ErrIdCantBeZeroOrNegative
	}
	return numId, nil
}

func printHeader(w http.ResponseWriter, code int, result interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)
	post, _ := json.Marshal(result)
	fmt.Fprintf(w, "%v", string(post))
}

func (user *handler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	numId, err := validateId(id)

	if err != nil {
		printHeader(w, http.StatusBadRequest, err)
		return
	}

	userDetail, err := user.service.GetByID(numId)

	if err != nil {
		printHeader(w, http.StatusInternalServerError, err)
		return
	}

	printHeader(w, 200, userDetail)
}

func (user *handler) GetByIDs(w http.ResponseWriter, r *http.Request) {
	str := mux.Vars(r)["ids"]
	ids := strings.Split(str[1:len(str)-1], ",")
	var numIDs []int
	for _, id := range ids {
		numId, err := validateId(id)

		if err != nil {
			printHeader(w, http.StatusBadRequest, err)
			return
		}
		numIDs = append(numIDs, numId)
	}

	userDetail, err := user.service.GetByIDs(numIDs)

	if err != nil {
		printHeader(w, http.StatusInternalServerError, err)
		return
	}

	printHeader(w, http.StatusOK, userDetail)
}
