# URL Shortener API

A powerful and efficient URL shortening service built with Go, Fiber, and Redis. This API allows users to create, manage, and track shortened URLs with rate limiting and IP-based user management.

##  Features

- ✅ **Shorten URLs** - Create short, custom URLs
- ✅ **Custom Short Codes** - Use your own custom identifiers
- ✅ **List All URLs** - View all URLs created by your IP
- ✅ **Get URL Details** - Fetch detailed information about specific shortened URLs
- ✅ **Delete URLs** - Remove URLs you created
- ✅ **Auto Expiry** - URLs automatically expire after set time (default 24 hours)
- ✅ **Rate Limiting** - Prevent API abuse with IP-based rate limiting
- ✅ **URL Validation** - Ensures only valid URLs are shortened
- ✅ **Docker Support** - Easy deployment with Docker Compose

## 🛠️ Tech Stack

- **Backend:** Go (Golang)
- **Framework:** Fiber v2
- **Database:** Redis
- **Containerization:** Docker & Docker Compose
- **Validation:** govalidator

##  Prerequisites

- Docker & Docker Compose installed
- Go 1.23+ (for local development)
- Redis (handled by Docker)

## 🔧 Installation & Setup

### 1. Clone the repository

```bash
git clone https://github.com/Bavithbabu/URL-shortner.git
cd URL-shortner
```

### 2. Create environment file

Create a `.env` file in the `api` directory:

```env
DB_ADDR=db:6379
DB_PASS=
APP_PORT=:3000
DOMAIN=http://localhost:3000
API_QUOTA=10
```

### 3. Run with Docker Compose

```bash
docker-compose up -d --build
```

The API will be available at `http://localhost:3000`

### 4. Check logs

```bash
docker-compose logs -f api
```

##  API Endpoints

### 1. Shorten URL

**POST** `/api/v1`

Create a shortened URL.

**Request Body:**
```json
{
    "url": "https://www.example.com",
    "short": "custom",
    "expiry": 24
}
```

**Parameters:**
- `url` (required): The original URL to shorten
- `short` (optional): Custom short code (auto-generated if not provided)
- `expiry` (optional): Expiry time in hours (default: 24)

**Response:**
```json
{
    "url": "https://www.example.com",
    "short": "http://localhost:3000/custom",
    "expiry": 24,
    "rate_limit": 9,
    "rate_limit_reset": 30
}
```

---

### 2. List All URLs

**GET** `/api/v1/urls`

Get all URLs created by your IP address.

**Response:**
```json
{
    "count": 2,
    "urls": [
        {
            "original_url": "https://www.example.com",
            "short_code": "custom",
            "short_url": "http://localhost:3000/custom",
            "created_at": "2025-12-13 15:30:45",
            "expires_in": "86400 seconds"
        }
    ]
}
```

---

### 3. Get URL Details

**GET** `/api/v1/url/:shortcode`

Get detailed information about a specific shortened URL.

**Example:** `GET /api/v1/url/custom`

**Response:**
```json
{
    "original_url": "https://www.example.com",
    "short_code": "custom",
    "short_url": "http://localhost:3000/custom",
    "created_by": "172.18.0.1",
    "created_at": "2025-12-13 15:30:45",
    "expires_in": "86400 seconds"
}
```

---

### 4. Delete URL

**DELETE** `/api/v1/url/:shortcode`

Delete a shortened URL (only if you created it).

**Example:** `DELETE /api/v1/url/custom`

**Response:**
```json
{
    "message": "URL deleted successfully",
    "deleted": "custom"
}
```

**Error (unauthorized):**
```json
{
    "error": "You are not authorized to delete this URL"
}
```

---

### 5. Redirect to Original URL

**GET** `/:shortcode`

Redirects to the original URL.

**Example:** `GET /custom` → Redirects to `https://www.example.com`

##  Rate Limiting

- **Default Quota:** 10 requests per IP
- **Reset Time:** 30 minutes
- Rate limit info included in response headers

##  Project Structure

```
URL-shortner/
├── api/
│   ├── database/
│   │   └── database.go          # Redis connection
│   ├── helpers/
│   │   └── helpers.go           # URL validation helpers
│   ├── routes/
│   │   ├── shorten.go          # Shorten URL endpoint
│   │   ├── resolve.go          # Redirect endpoint
│   │   ├── list.go             # List URLs endpoint
│   │   ├── geturl.go           # Get URL details endpoint
│   │   └── delete.go           # Delete URL endpoint
│   ├── main.go                 # Application entry point
│   ├── Dockerfile              # Docker config for API
│   ├── go.mod                  # Go dependencies
│   └── .env                    # Environment variables
├── db/
│   └── Dockerfile              # Docker config for Redis
├── docker-compose.yml          # Docker Compose configuration
├── .gitignore
└── README.md
```
 

##  Features to Add (Future)

- [ ] User authentication (JWT)
- [ ] Click/visit analytics
- [ ] URL preview before redirect
- [ ] QR code generation
- [ ] Bulk URL shortening
- [ ] Custom domains
- [ ] API key authentication

## 👨‍💻 Author

**Bavithbabu**
- GitHub: [@Bavithbabu](https://github.com/Bavithbabu)

## 📄 License

This project is open source and available under the MIT License.

## 🤝Contributing

Contributions, issues, and feature requests are welcome!

1. Fork the project
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request
 
---

⭐ **Star this repo if you found it helpful!**
