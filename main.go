package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")

	// hot reload templates so you don't have to constantly restart the application in design mode
	engine.ShouldReload = true

	// create a new fiber application
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	// setup our routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Ticker": "",
		})
	})

	// search can be a partial, full, or JSON api
	app.Get("/search", func(c *fiber.Ctx) error {
		ticker := c.Query("ticker")
		results := SearchTicker(ticker)

		switch c.Accepts(fiber.MIMETextHTML, fiber.MIMEApplicationJSON) {
		case fiber.MIMETextHTML:
			// if not an htmx request, render full page
			if c.Get("HX-Request") == "" {
				return c.Render("index", fiber.Map{
					"Ticker":  ticker,
					"Results": results,
				})
			}

			// render partial view
			return c.Render("partials/results", fiber.Map{
				"Ticker":  ticker,
				"Results": results,
			})

		case fiber.MIMEApplicationJSON:
			return c.JSON(fiber.Map{
				"ticker":  ticker,
				"results": results,
			})

		default:
			return c.SendStatus(http.StatusNotAcceptable)

		}
	})

	// detail of values
	app.Get("/values/:ticker", func(c *fiber.Ctx) error {
		ticker := c.Params("ticker")
		values := GetDailyValues(ticker)

		return c.Render("partials/values", fiber.Map{
			"Ticker": ticker,
			"Values": values,
		})
	})

	_ = app.Listen(":8000")
}

const PoligonPath = "https://api.polygon.io"
const TickerPath = PoligonPath + "/v3/reference/tickers"
const DailyValuesPath = PoligonPath + "/v1/open-close"

var (
	ApiKey = "apiKey=" + os.Getenv("POLYGON_API_KEY")
)

type Stock struct {
	Ticker          string    `json:"ticker"`
	Name            string    `json:"name"`
	Market          string    `json:"market"`
	Locale          string    `json:"locale"`
	PrimaryExchange string    `json:"primary_exchange"`
	Type            string    `json:"type"`
	Active          bool      `json:"active"`
	Currency        string    `json:"currency_name"`
	LastUpdated     time.Time `json:"last_updated_utc"`
}

type SearchResult struct {
	Results []Stock `json:"results"`
}

func SearchTicker(ticker string) []Stock {
	body := Fetch(TickerPath + "?" + ApiKey + "&ticker=" + strings.ToUpper(ticker))
	data := SearchResult{}
	_ = json.Unmarshal(body, &data)
	return data.Results
}

type Values struct {
	Open  float64 `json:"open"`
	Close float64 `json:"close"`
	High  float64 `json:"high"`
	Low   float64 `json:"low"`
}

func GetDailyValues(ticker string) Values {
	yesterday := time.Now().UTC().Add(-(time.Hour * 24)).Format("2006-01-02")
	body := Fetch(DailyValuesPath + "/" + strings.ToUpper(ticker) + "/" + yesterday + "?" + ApiKey)
	data := Values{}
	_ = json.Unmarshal(body, &data)
	return data
}
