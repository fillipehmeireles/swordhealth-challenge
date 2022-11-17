package repository

import (
	"SwordHealth/pkg/api/models"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	DB         *gorm.DB
	mock       sqlmock.Sqlmock
	repository *ManagerRepository
	manager    *models.Manager
}

func (s *Suite) SetupSuite() {
	var db *sql.DB
	var err error
	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("mysql", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.repository = NewManagerRepository()
}

func (s *Suite) Test_GetAll() {
	s.repository.ReadAll()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
