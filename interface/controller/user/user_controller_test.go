package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/go-querystring/query"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	model "golang-clean-architecture/domain/model/user"
	mock_interactor "golang-clean-architecture/usecase/mock_interactor/user"
	"net/http/httptest"
	"testing"
)

func TestUserController_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	uimock := mock_interactor.NewMockUserInteractor(ctrl)
	uc := NewUserController(uimock)

	body := model.RegisterRequest{
		Name:     "asd",
		Username: "qwe",
		Password: "zxc",
	}

	rec := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, r := gin.CreateTestContext(rec)
	r.Use(func(context *gin.Context) {
		params, err := query.Values(body)
		require.NoError(t, err)
		context.Request.PostForm = params
	})
	r.POST("/register", uc.Register)

	c.Request = httptest.NewRequest("POST","/register",nil)

	uimock.EXPECT().Register(body).Times(1)

	r.ServeHTTP(rec, c.Request)

	assert.Equal(t, 200, rec.Code)
}


func TestUserController_RegisterValidationFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	uimock := mock_interactor.NewMockUserInteractor(ctrl)
	uc := NewUserController(uimock)

	body := model.RegisterRequest{
		Name:     "",
		Username: "",
		Password: "",
	}

	rec := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, r := gin.CreateTestContext(rec)
	r.Use(func(context *gin.Context) {
		params, err := query.Values(body)
		require.NoError(t, err)
		context.Request.PostForm = params
	})
	r.POST("/register", uc.Register)

	c.Request = httptest.NewRequest("POST","/register",nil)

	r.ServeHTTP(rec, c.Request)

	assert.Equal(t, 422, rec.Code)
}
