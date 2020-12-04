package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	model "golang-clean-architecture/domain/model/user"
	"golang-clean-architecture/testutil"
	"gorm.io/gorm"
	"regexp"
	"testing"
	"time"
)

func setup(t *testing.T) (u UserRepository, mock sqlmock.Sqlmock) {
	db, mock, err := testutil.NewDBMock(t)
	require.NoError(t, err)
	u = NewUserRepository(db)
	return u, mock
}

func TestUserRepository_Register(t *testing.T) {
	u, mock := setup(t)

	user := model.User{
		ID:        0,
		Name:      "asd",
		Username:  "qwe",
		Password:  nil,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: gorm.DeletedAt{},
	}

	insertQueryRgx := `INSERT INTO "users" ("name","username","password","created_at","updated_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`
	mock.ExpectQuery(regexp.QuoteMeta(insertQueryRgx)).WithArgs(user.Name,user.Username, user.Password,user.CreatedAt,user.UpdatedAt,user.DeletedAt).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	_, err := u.Register(user)
	require.NoError(t, err)
	err = mock.ExpectationsWereMet()
	require.NoError(t, err)

}