package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/fredericobenevides/posgoexpert/9-APIS/internal/dto"
	"github.com/fredericobenevides/posgoexpert/9-APIS/internal/entity"
	"github.com/fredericobenevides/posgoexpert/9-APIS/internal/infra/database"
	"github.com/go-chi/jwtauth"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(db database.UserInterface) *UserHandler {
	return &UserHandler{UserDB: db}
}

// GetJWT godoc
// @Summary Get a user JWT
// @Description Get a user JWT
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.GetJWTInput true "user credentials"
// @Success 200 {object} dto.GetJWTOutput
// @Failure 404 {object} Error
// @Failure 500 {object} Error
// @Router /users/generate_token [post]
func (h *UserHandler) GetJwt(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("jwtExpiresIn").(int)

	var user dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})

	accessToken := dto.GetJWTOutput{
		AccessToken: tokenString,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// Create user godoc
// @Summary Create user
// @Description Create a new user with the input payload
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.CreateUserInput true "user request"
// @Success 201
// @Failure 400
// @Failure 500 {object} Error
// @Router /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
