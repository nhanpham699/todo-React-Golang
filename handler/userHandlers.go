package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/nhanpham699/demo/dto"
	"github.com/nhanpham699/demo/service"
)

type UserHandlers struct {
	Service service.UserService
}

func (th *UserHandlers) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var request dto.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	fmt.Println(request)
	res, error := th.Service.Register(request)
	if error != nil {
		log.Fatal(err)
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Fatal(err)
	}
}

func (th *UserHandlers) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var request dto.AuthRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	fmt.Println(request)
	res, error := th.Service.Login(request)
	if error != nil {
		log.Fatal(err)
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Fatal(err)
	}
}
