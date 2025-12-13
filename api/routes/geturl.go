package routes

import (
	"os"
	"strconv"
	"time"

	"github.com/bavith/Url_shortern/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func GetURLDetails(c *fiber.Ctx) error {
	shortcode := c.Params("shortcode")

	r := database.CreateClient(0)
	defer r.Close()

	urlKey := "url:" + shortcode

	// Get URL data from Redis Hash
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

	// Get TTL (time remaining until expiry)
	ttl, _ := r.TTL(database.Ctx, urlKey).Result()
	expiresIn := ttl / time.Second

	// Parse created timestamp
	createdUnix, _ := strconv.ParseInt(urlData["created"], 10, 64)
	createdAt := time.Unix(createdUnix, 0).Format("2006-01-02 15:04:05")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"original_url": urlData["url"],
		"short_code":   shortcode,
		"short_url":    os.Getenv("DOMAIN") + "/" + shortcode,
		"created_by":   urlData["ip"],
		"created_at":   createdAt,
		"expires_in":   strconv.FormatInt(int64(expiresIn), 10) + " seconds",
	})
}
