package producthandler

import "github.com/gofiber/fiber/v2"

func (h *ProductHandler) GetAll(c *fiber.Ctx) error {
	products, err := h.ProductService.GetAll()
	if err != nil {
		return h.Response.InternalServer(c, "Failed to fetch products")
	}

	return h.Response.Success(c, "", products)
}
