package handlers

import (
	"dozn/account-server/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func ListHandler(c *fiber.Ctx) error {
	fmt.Println("ListHandler started!!")
	return c.SendStatus(fiber.StatusOK)
}

func AccountHandler(c *fiber.Ctx) error {
	fmt.Println("AccountHandler started!!")

	a := new(models.Account)

	fmt.Println(c.Body())

	if err := c.BodyParser(a); err != nil {
		return err
	}

	fmt.Println(a.AccoutSeq)
	fmt.Println(a.AccountName)

	return c.SendStatus(fiber.StatusOK)
}
