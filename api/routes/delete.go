package routes

import (
	"github.com/bavith/Url_shortern/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func DeleteURL(c *fiber.Ctx) error {
	shortcode := c.Params("shortcode")

	r := database.CreateClient(0)
	defer r.Close()

	urlKey := "url:" + shortcode

	// Get URL data to check ownership
	urlData, err := r.HGetAll(database.Ctx, urlKey).Result()

	if err == redis.Nil || len(urlData) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "URL not found",
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to fetch URL details",
		})
	}

	// Check if the requesting IP matches the creator's IP
	if urlData["ip"] != c.IP() {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "You are not authorized to delete this URL",
		})
	}

	// Delete the URL hash
	err = r.Del(database.Ctx, urlKey).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to delete URL",
		})
	}

	// Remove shortcode from user's set
	userKey := "user:" + c.IP() + ":urls"
	r.SRem(database.Ctx, userKey, shortcode)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "URL deleted successfully",
		"deleted": shortcode,
	})
}
