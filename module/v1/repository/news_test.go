package repositories

import (
	"database/sql/driver"
	"regexp"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

// func TestSave(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	database, _ := gorm.Open("mysql", db)

// 	queryInsert := regexp.QuoteMeta("INSERT INTO `news` (`author`,`body`,`created`) VALUES (?,?,?,?)")

// 	insertModel := models.News{
// 		//ID:      1,
// 		Author:  "yaya",
// 		Body:    "yaya",
// 		Created: time.Now(),
// 	}

// 	//nilModel := models.News{}

// 	dataArgs := []driver.Value{
// 		insertModel.Author,
// 		insertModel.Body,
// 		AnyTime{},
// 	}

// 	mock.ExpectExec(queryInsert).
// 		WithArgs(dataArgs...).
// 		WillReturnResult(sqlmock.NewResult(1, 1))

// 	repository := NewsRepositoryHandler(database)

// 	err = repository.Store(&insertModel)

// 	assert.Nil(t, err)

// 	if err = mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}
// }

func TestFind(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	database, _ := gorm.Open("mysql", db)
	rows := sqlmock.NewRows([]string{"id", "author", "body", "created"}).
		AddRow(1, "yaya", "yaya", time.Now()).
		AddRow(2, "yaya", "yaya", time.Now())

	querySelect := regexp.QuoteMeta("SELECT * FROM `news` WHERE (id = ?)")

	mock.ExpectQuery(querySelect).
		WillReturnRows(rows)

	repository := NewsRepositoryHandler(database)
	_, err = repository.Find(1)

	assert.Nil(t, err)

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
