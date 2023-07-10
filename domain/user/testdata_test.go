package user_test

import (
	"fmt"

	"github.com/hiroki-Fukumoto/farm2/database/model"
	"github.com/jmoiron/sqlx"
)

func createTestData(db *sqlx.DB) error {
	sql := "INSERT INTO users (account_id, name) VALUES (:account_id, :name);"
	users := []model.User{}

	for i := 1; i < 10; i++ {
		u := model.User{
			AccountID: fmt.Sprintf("accountID%d", i),
			Name:      fmt.Sprintf("user%d", i),
		}
		users = append(users, u)
	}

	_, err := db.NamedExec(sql, users)
	if err != nil {
		return err
	}
	return nil
}

func deleteTestData(db *sqlx.DB) error {
	sql := "DELETE from users;"
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}
