package routes

import (
	"errors"
	"fmt"
	"gofiber/database"
	"gofiber/models"

	"github.com/gofiber/fiber/v2"
)

type UserSerializer struct {
	// this is for serializers
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUser(userModel models.User) UserSerializer {
	return UserSerializer{
		ID:        userModel.ID,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
	}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Create(&user)
	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func GetUsersList(c *fiber.Ctx) error {
	users := []models.User{}

	database.DB.Find(&users)
	responseUsers := []UserSerializer{}
	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.Status(200).JSON(responseUsers)
}

func findUser(id int, user *models.User) error {
	database.DB.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("User doesnt exist")
	}
	return nil
}

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	user := models.User{}
	if err != nil {
		return c.Status(400).JSON("Please ensure id is integer")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	// find the user first based on id
	user := models.User{}
	if err != nil {
		return c.Status(400).JSON("Please ensure id is integer")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// struct for receiving update from json body
	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	updateData := UpdateUser{}

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName

	database.DB.Save(&user)

	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	user := models.User{}
	if err != nil {
		return c.Status(400).JSON("Please ensure id is integer")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}

	message := fmt.Sprintf("Successfully deleted user at id %d", id)
	return c.Status(200).JSON(message)
}
