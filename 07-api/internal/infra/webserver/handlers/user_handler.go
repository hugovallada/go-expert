package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/hugovallada/go-expert/api/internal/dto"
	"github.com/hugovallada/go-expert/api/internal/entity"
	"github.com/hugovallada/go-expert/api/internal/infra/database"
)

type UserHandler struct {
	UserDB database.UserModel
}

type Error struct {
	Message string `json:"message"`
}

func NewUserHandler(db database.UserModel) *UserHandler {
	return &UserHandler{UserDB: db}
}

// Login godoc
// @Sumary Login
// @Description Login with user
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.LoginUser true "user credentials"
// @Success 200 {object} dto.BearerToken
// @Failure 400 {object} Error
// @Failure 500 {object} Error
// @Router /users/auth [post]
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("jwtExpiresIn").(int)
	var userDto dto.LoginUser
	err := json.NewDecoder(r.Body).Decode(&userDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: "Dado(s) inválido(s)."})
		return
	}
	user, err := h.UserDB.FindByEmail(userDto.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: "Usuário ou senha inválido(s)."})
		return
	}
	if !user.ValidatePassword(userDto.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(Error{Message: "Usuário ou senha inválido(s)."})
		return
	}
	_, token, err := jwt.Encode(map[string]any{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(Error{Message: "Erro interno"})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	accessToken := dto.BearerToken{AccessToken: token}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)

}

// Create user godoc
// @Sumary Create user
// @Description Create user
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.CreateUserInput true "user request"
// @Success 201
// @Failure 400 {object} Error
// @Failure 500 {object} Error
// @Router /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userDto dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&userDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := entity.NewUser(userDto.Name, userDto.Email, userDto.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: "Os dados enviados são inválidos"})
		return
	}
	err = h.UserDB.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{Message: "Erro interno"})
		return
	}
	w.WriteHeader(http.StatusCreated)
}
