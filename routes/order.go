package routes

import (
	"errors"
	"fmt"
	"gofiber/database"
	"gofiber/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type OrderSerializer struct {
	ID        uint      `json:"id"`
	User      string    `json:"username"`
	Product   string    `json:"productname"`
	CreatedAt time.Time `json:"order_date"`
}

func CreateResponseOrder(orderModel models.Order, user UserSerializer, product ProductSerializer) OrderSerializer {
	return OrderSerializer{
		ID:        orderModel.ID,
		User:      fmt.Sprintf("%s-%s", user.FirstName, user.LastName),
		Product:   product.Name,
		CreatedAt: orderModel.CreatedAt,
	}
}

func CreateOrder(c *fiber.Ctx) error {
	order := models.Order{}
	product := models.Product{}
	user := models.User{}

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := findUser(order.UserRefer, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := findProduct(order.ProductRefer, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Create(&order)
	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)
	response := CreateResponseOrder(order, responseUser, responseProduct)

	return c.Status(200).JSON(response)
}

func GetOrdersList(c *fiber.Ctx) error {
	orders := []models.Order{}
	database.DB.Find(&orders)
	response := []OrderSerializer{}

	for _, order := range orders {
		user := models.User{}
		product := models.Product{}

		database.DB.Find(&user, "id = ?", order.UserRefer)
		database.DB.Find(&product, "id = ?", order.ProductRefer)

		responseOrder := CreateResponseOrder(order, CreateResponseUser(user), CreateResponseProduct(product))
		response = append(response, responseOrder)
	}

	return c.Status(200).JSON(response)
}

func findOrder(id int, order *models.Order) error {
	database.DB.Find(&order, "id = ?", id)
	if order.ID == 0 {
		return errors.New("Product doesnt exist")
	}
	return nil
}

func GetOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	order := models.Order{}
	user := models.User{}
	product := models.Product{}

	if err != nil {
		return c.Status(400).JSON("Please ensure id is integer")
	}

	if err := findOrder(id, &order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Find(&user, "id = ?", order.UserRefer)
	database.DB.Find(&product, "id = ?", order.ProductRefer)

	response := CreateResponseOrder(order, CreateResponseUser(user), CreateResponseProduct(product))
	return c.Status(200).JSON(response)
}
