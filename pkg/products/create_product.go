package products

import (
	"github.com/gofiber/fiber/v2"
	"myfiber-api/pkg/common/models"
)

type AddProductRequest struct {
	Name  string  `json:"name"`
	Stock int32   `json:"stock"`
	Price float64 `json:"price"`
}

func (h handler) CreateProduct(c *fiber.Ctx) error {

	body := AddProductRequest{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var product models.Product

	product.Name = body.Name
	product.Stock = body.Stock
	product.Price = body.Price

	// insert new db entry
	if result := h.DB.Create(&product); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&product)
}
