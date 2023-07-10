package user_test

import (
	"database/sql"
	"errors"
	"fmt"
	"testing"

	"github.com/hiroki-Fukumoto/farm2/database"
	"github.com/hiroki-Fukumoto/farm2/database/model"
	"github.com/hiroki-Fukumoto/farm2/domain/user"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println(err)
	}
	m.Run()
}

func TestCreateUserRepository(t *testing.T) {
	db := database.NewTestDB()
	defer func() {
		err := deleteTestData(db)
		if err != nil {
			t.Fatalf("failed delete test data: %v", err)
		}
	}()

	repo := user.NewUserRepository(db)

	tests := []struct {
		name   string
		user   *model.User
		expect *model.User
	}{
		{
			name: "Should confirm success.",
			user: &model.User{
				AccountID: "accountID1",
				Name:      "user1",
			},
			expect: &model.User{
				AccountID: "accountID1",
				Name:      "user1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.CreateUser(tt.user)
			if err != nil {
				t.Fatal(err)
			}
			var u model.User
			query := "SELECT account_id, name FROM users WHERE account_id = ?"
			err = db.Get(&u, query, tt.user.AccountID)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.expect.AccountID, u.AccountID)
			assert.Equal(t, tt.expect.Name, u.Name)
		})
	}
}

func TestFindUserByAccountIDRepository(t *testing.T) {
	db := database.NewTestDB()
	err := createTestData(db)
	if err != nil {
		t.Fatalf("failed create test data: %v", err)
	}
	defer func() {
		err := deleteTestData(db)
		if err != nil {
			t.Fatalf("failed delete test data: %v", err)
		}
	}()

	repo := user.NewUserRepository(db)

	tests := []struct {
		name         string
		accountID    string
		expectName   string
		expectNoRows bool
	}{
		{
			name:       "Should confirm get user for accountID=1.",
			accountID:  "accountID1",
			expectName: "user1",
		},
		{
			name:         "Should confirm not found user.",
			accountID:    "accountID",
			expectNoRows: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := repo.FindUserByAccountID(tt.accountID)
			if err != nil {
				if tt.expectNoRows {
					assert.True(t, errors.Is(err, sql.ErrNoRows))
					return
				}
				t.Fatal(err)
			}
			assert.Equal(t, tt.expectName, user.Name)
		})
	}
}
