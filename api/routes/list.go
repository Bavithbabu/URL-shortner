package routes

import (
	"os"
	"strconv"
	"time"

	"github.com/bavith/Url_shortern/database"
	"github.com/gofiber/fiber/v2"
)

type URLInfo struct {
	OriginalURL string `json:"original_url"`
	ShortCode   string `json:"short_code"`
	ShortURL    string `json:"short_url"`
	CreatedAt   string `json:"created_at"`
	ExpiresIn   string `json:"expires_in"`
}

func ListUserURLs(c *fiber.Ctx) error {
	r := database.CreateClient(0)

	defer r.Close()

	// Get all shortcodes created by this IP
	userKey := "user:" + c.IP() + ":urls"
	shortcodes, err := r.SMembers(database.Ctx, userKey).Result()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to fetch URLs",
		})
	}

	if len(shortcodes) == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "No URLs found",
			"urls":    []URLInfo{},
		})
	}

	// Fetch details for each shortcode
	var urls []URLInfo

	for _, shortcode := range shortcodes {
		urlKey := "url:" + shortcode

		// Get URL data from hash
		urlData, err := r.HGetAll(database.Ctx, urlKey).Result()

		if err != nil || len(urlData) == 0 {
			continue // Skip if URL data not found
		}

		// Get TTL (time to live)
		ttl, _ := r.TTL(database.Ctx, urlKey).Result()
		expiresIn := ttl / time.Second

		// Parse created timestamp
		createdUnix, _ := strconv.ParseInt(urlData["created"], 10, 64)
		createdAt := time.Unix(createdUnix, 0).Format("2006-01-02 15:04:05")

		urlInfo := URLInfo{
			OriginalURL: urlData["url"],
			ShortCode:   shortcode,
			ShortURL:    os.Getenv("DOMAIN") + "/" + shortcode,
			CreatedAt:   createdAt,
			ExpiresIn:   strconv.FormatInt(int64(expiresIn), 10) + " seconds",
		}

		urls = append(urls, urlInfo)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"count": len(urls),
		"urls":  urls,
	})
}
