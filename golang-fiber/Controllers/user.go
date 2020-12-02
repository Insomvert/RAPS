package Controllers

import (
	"golang-fiber/Models"

	"github.com/gofiber/fiber"
)

func GetListUser(c *fiber.Ctx) {
	var user []Models.User
	err := Models.GetListUser(&user)
	if err != nil {
		c.JSON(err)
	} else {
		c.JSON(user)
	}
}

func CreateUser(c *fiber.Ctx) {
	var user Models.User
	err := Models.CreateUser(&user)
	if err != nil {
		c.JSON(err)
	} else {
		c.JSON(user)
	}
}

func UpdateUser(c *fiber.Ctx) {
	var id = c.Params("id")
	var user Models.User
	err := Models.UpdateUser(&user, id)
	if err != nil {
		c.JSON(err)
	} else {
		c.JSON(user)
	}
}

func DeleteUser(c *fiber.Ctx) {
	var id = c.Params("id")
	var user Models.User
	err := Models.DeleteUser(&user, id)
	if err != nil {
		c.JSON(err)
	} else {
		c.JSON(map[string]string{"status": "user deleted"})
	}
}
