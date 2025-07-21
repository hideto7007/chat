package user

import (
	"chat/application/usecase/users"
	"chat/lib/cast"
	"chat/lib/gin/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
    listUserUseCase   *users.ListUserUseCase
    createUserUseCase *users.CreateUserUseCase
	deleteUserUseCase *users.DeleteUserUseCase
	updateUserUseCase *users.UpdateUserUseCase
	getEmailUseCase   *users.GetEmailUserUseCase
	getIDUseCase      *users.GetIdUserUseCase
	loginUserUseCase  *users.LoginUserUseCase
}

func NewUserController(
    listUserUseCase *users.ListUserUseCase,
    createUserUseCase *users.CreateUserUseCase,
	deleteUserUseCase *users.DeleteUserUseCase,
	updateUserUseCase *users.UpdateUserUseCase,
	getEmailUseCase *users.GetEmailUserUseCase,
	getIDUseCase *users.GetIdUserUseCase,
	loginUserUseCase *users.LoginUserUseCase,
) *UserController {
    return &UserController{
        listUserUseCase:   listUserUseCase,
        createUserUseCase: createUserUseCase,
		deleteUserUseCase: deleteUserUseCase,
		updateUserUseCase: updateUserUseCase,
		getEmailUseCase:   getEmailUseCase,
		getIDUseCase:      getIDUseCase,
        loginUserUseCase:  loginUserUseCase,
    }
}

// @Summary ユーザー一覧取得
// @Description ユーザーの一覧を返します
// @Tags users
// @Produce json
// @Success 200 {object} UsersResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/users [get]
func (c *UserController) List(ctx *gin.Context) {
    users, err := c.listUserUseCase.Execute(ctx)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, response.Error(err))
        return
    }

    var res UsersResponse
    for _, u := range users {
        res.Users = append(res.Users, UserResponse{
            ID:    u.ID,
            Name:  u.Name,
            Email: u.Email,
        })
    }
    fmt.Println(res)
    ctx.JSON(http.StatusOK, res)
}

// @Summary ユーザーIDで取得
// @Description ユーザーIDを指定してユーザー情報を取得します
// @Tags users
// @Param id path string true "User ID"
// @Produce json
// @Success 200 {object} UserResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/users/id/{id} [get]
func (c *UserController) GetIdUser(ctx *gin.Context) {
    userID, err := cast.Unit(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, response.Error(err))
        return
    }
    user, err := c.getIDUseCase.Execute(ctx, userID)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, response.Error(err))
        return
    }

    ctx.JSON(http.StatusOK, UserResponse{
        ID:    user.ID,
        Name:  user.Name,
        Email: user.Email,
    })
}

// @Summary ユーザーEmailで取得
// @Description ユーザーEmailを指定してユーザー情報を取得します
// @Tags users
// @Param email path string true "User Email"
// @Produce json
// @Success 200 {object} UserResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/users/email/{email} [get]
func (c *UserController) GetEmailUser(ctx *gin.Context) {
    userEmail := ctx.Param("email")
    user, err := c.getEmailUseCase.Execute(ctx, userEmail)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, response.Error(err))
        return
    }

    ctx.JSON(http.StatusOK, UserResponse{
        ID:    user.ID,
        Name:  user.Name,
        Email: user.Email,
    })
}

// @Summary ユーザーログイン取得
// @Description ユーザーEmailとパスワードを指定してユーザー情報を取得します
// @Tags users
// @Accept json
// @Produce json
// @Param user body LoginRequest true "ユーザーログインリクエスト"
// @Success 200 {object} UserResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/users/login [post]
func (c *UserController) Login(ctx *gin.Context) {
    var req LoginRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, response.Error(err))
        return
    }
    user, err := c.loginUserUseCase.Execute(ctx, req.Email, req.Password)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, response.Error(err))
        return
    }

    ctx.JSON(http.StatusOK, UserResponse{
        ID:    user.ID,
        Name:  user.Name,
        Email: user.Email,
    })
}



// @Summary ユーザー作成
// @Description 新しいユーザーを作成します
// @Tags users
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "ユーザー作成リクエスト"
// @Success 201 {object} UserResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/users [post]
func (c *UserController) Create(ctx *gin.Context) {
    var req CreateUserRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, response.Error(err))
        return
    }

	newUser := users.NewCreateUserDto(
		req.Name,
		req.Email,
		req.Password,
	)

    user, err := c.createUserUseCase.Execute(ctx, newUser)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, response.Error(err))
        return
    }

    res := UserResponse{
        ID:    user.ID,
        Name:  user.Name,
        Email: user.Email,
    }
    ctx.JSON(http.StatusCreated, res)
}

// @Summary ユーザー更新
// @Description 既存のユーザー情報を更新します
// @Tags users
// @Accept json
// @Produce json
// @Param user body UpdateUserRequest true "ユーザー更新リクエスト"
// @Success 201 {object} UserResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/users [put]
func (c *UserController) Update(ctx *gin.Context) {
    var req UpdateUserRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, response.Error(err))
        return
    }

	newUser := users.NewUpdateUserDto(
		req.ID,
		req.Name,
		req.Email,
		req.Password,
	)

    userDto, err := c.updateUserUseCase.Execute(ctx, newUser)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, response.Error(err))
        return
    }

    res := UserResponse{
        ID:    userDto.ID,
        Name:  userDto.Name,
        Email: userDto.Email,
    }
    ctx.JSON(http.StatusCreated, res)
}

// @Summary ユーザー削除
// @Description 指定したユーザーを削除します
// @Tags users
// @Param id path string true "User ID"
// @Success 204 "No Content"
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/users/{id} [delete]
func (c *UserController) Delete(ctx *gin.Context) {
	userID, err := cast.Unit(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error(err))
		return
	}

	err = c.deleteUserUseCase.Execute(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Error(err))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
