package main

//db.open
// db.ping
// db.Exec
// db.QuryRow
//db.Close
// db.
//

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	dataname = "mydatabase"
	username = "myuser"
	password = "mypassword"
)

var db *sql.DB

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, dataname)

	dbs, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	db = dbs
	defer db.Close()
	defer fmt.Println("finished")
	db.Ping()

	app := fiber.New()
	app.Get("/product/:id", getProductHandler)
	app.Post("/product", createProductsHandler)
	app.Put("/product/:id", updateProductHandler)
	app.Delete("/product/:id", deleteProductHandler)
	app.Get("/product", getallProductHandler)

	app.Listen(":8080")
}

func getProductHandler(c *fiber.Ctx) error {

	productid, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(" id not found ")
	}
	p, err := getProduct(productid)

	if err != nil {

		return c.Status(fiber.StatusBadRequest).SendString("not in stock")
	}

	return c.JSON(p)
}

func createProductsHandler(c *fiber.Ctx) error {
	product := new(Product)

	err := c.BodyParser(product)
	if err != nil {

		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	err = createProducts(product)

	if err != nil {

		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.JSON(product)
}

func updateProductHandler(c *fiber.Ctx) error {

	productid, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	p := new(Product)
	err = c.BodyParser(p)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	product, err := UpdateProducts(productid, p)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(product)
}

func deleteProductHandler(c *fiber.Ctx) error {
	productid, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	err = deleteProduct(productid)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return err
}

func getallProductHandler(c *fiber.Ctx) error {

	product, err := getProducts()
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(product)
}
