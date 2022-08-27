package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	cartdto "waysbuck-API/dto/cart"
	dto "waysbuck-API/dto/result"
	"waysbuck-API/models"
	"waysbuck-API/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerCart struct {
	CartRepository repositories.CartRepository
}

// // Create `path_file` Global variable here ...
var path_file_cart = "http://localhost:2500/uploads/"

func HandlerCart(CartRepository repositories.CartRepository) *handlerCart {
	return &handlerCart{CartRepository}
}

func (h *handlerCart) FindCarts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	carts, err := h.CartRepository.FindCarts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: carts}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCart) GetCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// var cart models.Cart
	cart, err := h.CartRepository.GetCart(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: cart}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCart) CreateCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// GET DATA FOR USER TOKEN
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	Transaction_ID := int(userInfo["time"].(float64))

	request := new(cartdto.CartRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	requestCart := models.Cart{
		ProductID:     request.ProductID,
		TransactionID: Transaction_ID,
		SubAmount:     request.SubAmount,
		ToppingID:     request.ToppingID,
	}

	validations := validator.New()
	errrr := validations.Struct(requestCart)
	if errrr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	topping, _ := h.CartRepository.FindToppingsID(request.ToppingID)

	cart := models.Cart{
		ProductID:     request.ProductID,
		TransactionID: Transaction_ID,
		SubAmount:     request.SubAmount,
		Topping:       topping,
	}

	data, err := h.CartRepository.CreateCart(cart)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCart) UpdateCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(cartdto.UpdateCart)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	cart, err := h.CartRepository.GetCart(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	if request.ProductID != 0 {
		cart.ProductID = request.ProductID
	}

	data, err := h.CartRepository.UpdateCart(cart)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCart) DeleteCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	fmt.Println(id)
	cart, err := h.CartRepository.GetCart(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.CartRepository.DeleteCart(cart)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCart) FindCartsByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	idTrans := int(userInfo["time"].(float64))

	cart, err := h.CartRepository.FindCartsTransaction(idTrans)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	for i, p := range cart {
		cart[i].Product.Image = path_file_cart + p.Product.Image
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: cart}
	json.NewEncoder(w).Encode(response)
}

func convertResponseCart(u models.Cart) models.CartResponse {
	return models.CartResponse{
		ID:        u.ID,
		SubAmount: u.SubAmount,
		Topping:   u.Topping,
	}
}
