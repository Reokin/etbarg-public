package handler

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var app *gin.Engine

// Rate limiter
func RateLimiter() gin.HandlerFunc {
	pwd_limiter := rate.NewLimiter(3, 3)       // Password page limiter
	default_limiter := rate.NewLimiter(1, 100) // Other pages
	return func(c *gin.Context) {
		switch c.Request.URL.String() {
		case "/jS2xXust7iye":
			if pwd_limiter.Allow() {
				c.Next()
			} else {
				c.JSON(http.StatusTooManyRequests, gin.H{
					"status":  "429",
					"message": "Too many requests",
				})
				c.Abort()
			}
		default:
			if default_limiter.Allow() {
				c.Next()
			} else {
				c.JSON(http.StatusTooManyRequests, gin.H{
					"status":  "429",
					"message": "Too many requests",
				})
				c.Abort()
			}
		}
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// Get the current lever stage
func getLevel(c *gin.Context) {
	req, err := http.NewRequest(http.MethodGet, "Database URL to get the level", nil) // DB should send the level number without anything else in the response
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500",
			"message": "Internal server error",
		})
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500",
			"message": "Internal server error",
		})
		return
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500",
			"message": "Internal server error",
		})
		return
	}

	lvlstr := fmt.Sprintf("%s", resBody)

	lvlstr2 := strings.ReplaceAll(lvlstr, "\"", "")

	level, err := strconv.Atoi(lvlstr2)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500",
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"level": level})
}

// Set lever stage
func postLevel(c *gin.Context) {

	var level int = 0
	switch c.Request.URL.String() {
	case "/post_level_1":
		level = 1
	case "/post_level_2":
		level = 2
	case "/post_level_3":
		level = 3
	case "/post_level_4":
		level = 4
	case "/final_page_tracker":
		level = 5
	default:
		break
	}

	req, err := http.NewRequest(http.MethodGet, "Database URL to get the level", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500",
			"message": "Internal server error",
		})
		return
	}

	res1, err2 := http.DefaultClient.Do(req)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500",
			"message": "Internal server error",
		})
		return
	}

	resBody, err3 := io.ReadAll(res1.Body)
	if err3 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500",
			"message": "Internal server error",
		})
		return
	}

	lvlstr := fmt.Sprintf("%s", resBody)

	lvlstr2 := strings.ReplaceAll(lvlstr, "\"", "")

	clevel, err1 := strconv.Atoi(lvlstr2)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500",
			"message": "Internal server error",
		})
		return
	}

	if level <= clevel {
		c.JSON(http.StatusOK, false)
		return
	}

	requestURL := fmt.Sprintf("URL to obs remote control %d", level) // GET request sends the new level number as a parameter (idk why i did it this way)
	_, err4 := http.Get(requestURL)
	if err4 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500",
			"message": "Internal server error",
		})
		return
	}

	requestURL1 := fmt.Sprintf("Database URL to set the level %d", level) // GET request sends the new level number as a parameter (idk why i did it this way)
	res, err5 := http.Get(requestURL1)
	if err5 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500",
			"message": "Internal server error",
		})
		return
	}

	if res.Status != "200 OK" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500",
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, true)
}

// Main function
func Handler(w http.ResponseWriter, r *http.Request) {
	gin.SetMode(gin.ReleaseMode)
	app = gin.Default()
	app.Use(RateLimiter())
	app.Use(CORSMiddleware())

	app.GET("/get_level", getLevel)
	app.POST("/post_level_1", postLevel)
	app.POST("/post_level_2", postLevel)
	app.POST("/post_level_3", postLevel)
	app.POST("/post_level_4", postLevel)
	app.POST("/final_page_tracker", postLevel)

	app.GET("/check_password", func(c *gin.Context) {
		if c.Request.Header.Get("password") == "VUXUWB]UXVSDTX_^" {
			c.JSON(http.StatusOK, gin.H{
				"status":  "200",
				"message": "final_page",
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  "401",
				"message": "Wrong password",
			})
			return
		}
	})

	app.ServeHTTP(w, r)
}
