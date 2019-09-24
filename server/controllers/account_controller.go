package controllers

import (
	"Chess/server/middleware"
	"Chess/server/models"
	"Chess/server/repositories/accountRepo"
	"Chess/server/services/accountService"
	"encoding/json"
	"github.com/harlow/authtoken"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/go-chi/render"
)

// AccountController include some handler for account
type AccountController interface {
	CreateNewUser(w http.ResponseWriter, r *http.Request)
	FilterPaging(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	LoginWithToken(w http.ResponseWriter, r *http.Request)
	ChangePassword(w http.ResponseWriter, r *http.Request)
	RemoveUser(w http.ResponseWriter, r *http.Request)
	GetUserById(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

type accountController struct {
	accountService accountService.Service
}

// Create new Account godoc
// @tags account-manager-apis
// @Summary Create new Account
// @Description Create new account with role default is customer
// @Id 00001
// @Accept json
// @Produce json
// @Param UserInformation body controllers.UserM true "User information"
// @Success 200 {object} controllers.AccountResponse
// @Router  /account/create [post]
func (ac *accountController) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data UserM
	err := decoder.Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	user := model.User{
		Name:     data.Name,
		Password: data.Password,
	}
	newUser, err := ac.accountService.CreateNewUser(&user)
	var res *AccountResponse
	if err != nil {
		res = &AccountResponse{
			Data:    nil,
			Message: "not ok",
			Success: false,
		}
	} else {
		res = &AccountResponse{
			Data:    newUser,
			Message: "Ok",
			Success: true,
		}
	}

	render.JSON(w, r, res)

}

// Filter users and paging godoc
// @tags account-manager-apis
// @Summary Filter users and paging
// @Description filter list user and paging filtered
// @Id 00002
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param name query string false "name for user"
// @Param page query integer false "page number for user"
// @Param pageSize query integer false "page size each page"
// @Success 200 {object} model.Response
// @Router /account/accounts [get]
func (ac *accountController) FilterPaging(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()

	var page int
	var pageSize int
	var name *string
	var role *string
	if queryValues.Get("page") == "" {
		page = 1
	} else {
		page, _ = strconv.Atoi(queryValues.Get("page"))
	}

	if queryValues.Get("pageSize") == "" {
		pageSize = 10
	} else {
		pageSize, _ = strconv.Atoi(queryValues.Get("pageSize"))
	}

	nameStr := queryValues.Get("name")
	if nameStr == "" {
		name = nil
	} else {
		name = &nameStr
	}

	roleStr := queryValues.Get("role")
	if roleStr == "" {
		role = nil
	} else {
		role = &roleStr
	}

	listUser, total, err := ac.accountService.GetFilterListUser(name, role, page, pageSize)
	var listUserResponse []model.User
	for _, index := range listUser{
		listUserResponse = append(listUserResponse, model.User{
			ID : 		index.ID,
			Name : 		index.Name,
			CreatedAt :	index.CreatedAt,
		})
	}

	var res *model.Response
	if err != nil {
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			MetaData: model.MetaData{
				Page:page,
				PageSize:pageSize,
				Total:total,
			},
			Data:    listUserResponse,
			Message: "ok",
			Success: true,
		}
	}

	render.JSON(w, r, res)
}

// Update user godoc
// @tags account-manager-apis
// @Summary Update user
// @Description update user by field:name, avatar, status, role
// @Id 00003
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param accountId path integer true "id of user account"
// @Param update_model body controllers.UpdateAccountResponse true "inlucde old user and new update user"
// @Success 200 {object} controllers.AccountResponse
// @Router  /account/{accountId} [put]
func (ac *accountController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data UpdateAccountResponse
	err := decoder.Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		return
	}

	accountIDStr := chi.URLParam(r, "accountId")
	accountID, err := strconv.Atoi(accountIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		return
	}

	infoUpdate := model.User{
		Name: data.Name,
	}

	err = ac.accountService.UpdateUser(accountID, &infoUpdate)

	var res *AccountResponse
	if err != nil {
		res = &AccountResponse{
			Data:    nil,
			Message: "update user failed. " + err.Error(),
			Success: false,
		}
	} else {
		res = &AccountResponse{
			Data:    true,
			Message: "Update user success",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

// Update password godoc
// @tags account-manager-apis
// @Summary Update password
// @Description update password for exists user
// @Id 00005
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Update_Password body controllers.UpdatePassResponse true "cuple value id and reset password"
// @Success 200 {object} controllers.AccountResponse
// @Router  /account/password [put]
func (ac *accountController) ChangePassword(w http.ResponseWriter, r *http.Request) {
	var res *AccountResponse

	decoder := json.NewDecoder(r.Body)
	var data UpdatePassResponse
	if err := decoder.Decode(&data); err != nil {
		res = &AccountResponse{
			Data:    nil,
			Message: "don't meet require",
			Success: false,
		}
	}
	if err := ac.accountService.UpdatePassword(data.UserId, data.OldPass, data.NewPass); err != nil {
		res = &AccountResponse{
			Data:    nil,
			Message: "ERROR: " + err.Error(),
			Success: false,
		}
	} else {
		res = &AccountResponse{
			Data:    nil,
			Message: "chane password success",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

// Remove user godoc
// @tags account-manager-apis
// @Summary Remove User
// @Description Soft Delete user by user id
// @Id 00006
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param userId path string true "user id is wanted remove"
// @Success 200 {object} controllers.AccountResponse
// @Router  /account/{userId} [delete]
func (ac *accountController) RemoveUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userId")
	err := ac.accountService.DeleteUser(userID)

	var res *AccountResponse
	if err != nil {
		res = &AccountResponse{
			Data:    nil,
			Message: "not ok",
			Success: false,
		}
	} else {
		res = &AccountResponse{
			Data:    true,
			Message: "ok",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

// Get user by id godoc
// @tags account-manager-apis
// @Summary Get user by user id
// @Description find user by user id
// @Id 00007
// @Accept json
// @Produce json
// @Param userId path string true "user id is wanted find" default(1)
// @Security ApiKeyAuth
// @Success 200 {object} controllers.AccountResponse
// @Router  /account/{userId} [get]
func (ac *accountController) GetUserById(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userId")

	user, err := ac.accountService.GetUserByID(userID)

	var res *AccountResponse
	if err != nil {
		res = &AccountResponse{
			Data:    nil,
			Message: "not ok",
			Success: false,
		}
	} else {
		res = &AccountResponse{
			Data:    user,
			Message: "ok",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

// Login godoc
// @tags access-apis
// @Summary Login system
// @Description login with username, password. return token string jwt
// @Accept json
// @Produce json
// @Param accountInfo body controllers.LoginRequest true "username and password"
// @Success 200 {object} controllers.LoginResponse
// @Router  /access/login [post]
func (ac *accountController) Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data LoginRequest
	if err := decoder.Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		return
	}

	userData, tokenString, err := ac.accountService.LoginRequest(data.Name, data.Password)
	var res *LoginResponse
	if err != nil {
		res = &LoginResponse{
			Data: err,
			Message: "ERROR: " + err.Error(),
			Success: false,
		}
	} else {
		tempData := model.UserResponse{
			User: userData,
		}

		if tokenString == "" {
			res = &LoginResponse{
				Message: "password incorect",
				Success: false,
			}
		} else {
			res = &LoginResponse{
				Token: tokenString,
				Data:    tempData,
				Message: "login success",
				Success: true,
			}
		}
	}
	
	render.JSON(w, r, res)
}

// LoginWithToken godoc
// @tags access-apis
// @Summary Login system by token string
// @Description login with token string. return new token string jwt
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} controllers.LoginResponse
// Failure 401 {string} string
// @Router  /access/login/token [post]
func (ac *accountController) LoginWithToken(w http.ResponseWriter, r *http.Request) {
	token, err := authtoken.FromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, http.StatusText(401), 401)
		return
	}

	var res *LoginResponse
	userData, tokenString, success, err := ac.accountService.LoginWithToken(token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	if success != true {
		res = &LoginResponse{
			Message: tokenString,
			Success: success,
		}
	} else {
		res = &LoginResponse {
			Data: userData,
			Token: tokenString,
			Message: "login ok",
			Success: success,
		}
	}

	render.JSON(w, r, res)
}

func (ac *accountController) Logout(w http.ResponseWriter, r *http.Request) {

	cookie := http.Cookie{
		Name: "jwt",
		MaxAge: -1,
	}
	http.SetCookie(w, &cookie)

	var res *AccountResponse
	res = &AccountResponse{
		Data:    nil,
		Message: "Logout success",
		Success: true,
	}

	render.JSON(w, r, res)
}

// NewAccountController export account handler inteface
func NewAccountController() (AccountController, error) {
	userRepo := accountRepo.NewUserRepository()
	authentication, err := middleAccess.NewAuthentication()
	accountService := accountService.NewAccountService(userRepo, authentication)
	if err != nil {
		log.Printf("Has problem at new account controller, because set authenication faile: %v\n", err)
		return nil, err
	}
	return &accountController{
		accountService: accountService,
	}, nil
}

// UserM payload for create api
type UserM struct {
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Password string `json:"password"`
}

// UpdateAccountResponse update account response
type UpdateAccountResponse struct {
	Name string `json:"name"`
	Avatar string `json:"avatar"`
	Password string `json:"password"`
	Role string `json:"role"`
	Status string `json:"status"`
}

// AccountResponse general response
type AccountResponse struct {
	Data    interface{} `json:"data"`
	Message string      `jon:"message"`
	Success bool        `json:"success"`
}

// LoginResponse login response
type LoginResponse struct {
	Token string `json:"token"`
	Data interface{} `json:"data"`
	Message string `json:"message"`
	Success bool `json:"success"`
}

// UpdatePassResponse response for reset password api
type UpdatePassResponse struct {
	UserId  int    `json:"userId"`
	OldPass string `json:"oldPass"`
	NewPass string `json:"newPass"`
}

// LoginRequest login payload
type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// LoginTokenPayload payload for login token api
type LoginTokenPayload struct {
	TokenString string `json:"tokenString"`
}