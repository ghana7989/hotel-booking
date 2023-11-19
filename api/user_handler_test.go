package api_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/ghana7989/hotel-booking/api"
	"github.com/ghana7989/hotel-booking/db"
	"github.com/ghana7989/hotel-booking/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type testDB struct {
	db.UserStore
}

func setup(t *testing.T) *testDB {
	const testDBUri = "mongodb://localhost:27017"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testDBUri))
	if err != nil {
		t.Fatal(err)
	}
	return &testDB{
		UserStore: db.NewMongoUserStore(client, db.TEST_DB_NAME),
	}
}

func (testDB testDB) tearDown(t *testing.T) {
	if err := testDB.UserStore.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func TestCreateUser(t *testing.T) {
	testDb := setup(t)
	defer testDb.tearDown(t)

	app := fiber.New()

	// app.Listen(":3000")

	userHandler := api.NewUserHandler(testDb.UserStore)
	app.Post("/", userHandler.HandleCreateUser)
	params := types.CreateUserParams{
		Email:     "test@test.com",
		Password:  "test123fdjfjhfhjdfbh",
		FirstName: "test",
		LastName:  "test",
	}
	b, _ := json.Marshal(params)

	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Fatal("expected 200")
	}
	responseBody := new(types.User)
	if err := json.NewDecoder(resp.Body).Decode(responseBody); err != nil {
		t.Fatal(err)
	}
	if responseBody.Email != params.Email {
		t.Fatal("expected email to be equal")
	}
	if responseBody.FirstName != params.FirstName {
		t.Fatal("expected firstName to be equal")
	}
	if responseBody.LastName != params.LastName {
		t.Fatal("expected lastName to be equal")
	}

}
