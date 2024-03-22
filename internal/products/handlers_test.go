package products_test

import (
	"fmt"
	"net/http/httptest"
	"os"
	"testing"

	"co.bastriguez/inventory/internal/databases"
	"co.bastriguez/inventory/internal/products"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
    setup()
    code := m.Run()
    shutdown()
    os.Exit(code)
}

func setup(){
    mongo, err := databases.NewMongo()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    storage := products.NewMongoStorage(mongo)
    products.InitHandlers(storage)    
}

func shutdown(){

}

func TestHandler(t *testing.T) {
    tests := []struct {
        name string
        expectedCode int
        expectedBody string
    }{
        {
            name: "given no products in storage, body is empty and status code is 200",
            expectedCode: 200,
            expectedBody: "{}",
        },
    }

    app := fiber.New()
    app.Get("/", products.HandleGetProducts)

    for _, test := range tests {
        req := httptest.NewRequest("GET", "/", nil)

        resp, _ := app.Test(req, 1)

        assert.Equal(t, test.expectedCode, resp.StatusCode)
    }
}
