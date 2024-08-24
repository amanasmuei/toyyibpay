package handlers

import (
	"encoding/json"

	"github.com/amanasmuei/toyyibpay.git/client"
	"github.com/amanasmuei/toyyibpay.git/models"
	"github.com/gofiber/fiber/v2"
)

// CreateBillHandler handles the creation of a new bill via ToyyibPay.
func CreateBillHandler(c *client.Client) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req models.PaymentRequest
		if err := ctx.BodyParser(&req); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request payload",
			})
		}

		resp, err := c.CreateBill(req)
		if err != nil {
			if apiErr, ok := err.(*client.APIError); ok {
				return ctx.Status(apiErr.StatusCode).JSON(fiber.Map{
					"error": apiErr.Message,
				})
			}
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		defer resp.Body.Close()

		var paymentResp models.PaymentResponse
		if err := json.NewDecoder(resp.Body).Decode(&paymentResp); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to parse response",
			})
		}

		return ctx.JSON(paymentResp)
	}
}
