package routes

import (
	"errors"
	"fmt"
	"gofiber/database"
	"gofiber/models"

	"github.com/gofiber/fiber/v2"
)

type ProductSerializer struct {
	// this is for serializer
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func CreateResponseProduct(productModel models.Product) ProductSerializer {
	return ProductSerializer{
		ID:           productModel.ID,
		Name:         productModel.Name,
		SerialNumber: productModel.SerialNumber,
	}
}

func CreateProduct(c *fiber.Ctx) error {
	product := models.Product{}

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Create(&product)
	response := CreateResponseProduct(product)

	return c.Status(200).JSON(response)
}

func GetProductsList(c *fiber.Ctx) error {
	products := []models.Product{}

	database.DB.Find(&products)
	response := []ProductSerializer{}
	for _, product := range products {
		responseProduct := CreateResponseProduct(product)
		response = append(response, responseProduct)
	}

	return c.Status(200).JSON(response)
}

func findProduct(id int, product *models.Product) error {
	database.DB.Find(&product, "id = ?", id)
	if product.ID == 0 {
		return errors.New("Product doesnt exist")
	}
	return nil
}

func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	product := models.Product{}
	if err != nil {
		return c.Status(400).JSON("Please ensure id is integer")
	}

	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	response := CreateResponseProduct(product)
	return c.Status(200).JSON(response)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	// find the product first based on id
	product := models.Product{}
	if err != nil {
		return c.Status(400).JSON("Please ensure id is integer")
	}

	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// struct for receiving update from json body
	type UpdateProduct struct {
		Name         string `json:"name"`
		SerialNumber string `json:"serial_number"`
	}

	updateData := UpdateProduct{}

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	product.Name = updateData.Name
	product.SerialNumber = updateData.SerialNumber

	database.DB.Save(&product)

	response := CreateResponseProduct(product)
	return c.Status(200).JSON(response)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	product := models.Product{}
	if err != nil {
		return c.Status(400).JSON("Please ensure id is integer")
	}

	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.DB.Delete(&product).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}

	message := fmt.Sprintf("Successfully deleted product at id %d", id)
	return c.Status(200).JSON(message)
}
