package products

import (
	"github.com/gofiber/fiber/v2"
	"myfiber-api/pkg/common/models"
)

func (h handler) GetProducts(c *fiber.Ctx) error {
	var products []models.Product

	if result := h.DB.Find(&products); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&products)
}
