package products

import (
	"github.com/gofiber/fiber/v2"
	"myfiber-api/pkg/common/models"
)

func (h handler) GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	var product models.Product

	if result := h.DB.First(&product, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&product)
}
