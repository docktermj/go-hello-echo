package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// Hello world.

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})

	// Streaming.

	e.GET("/streaming", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c.Response().WriteHeader(http.StatusOK)

		// Stream the response body.  Note: loop will not stop.

		ticker := time.NewTicker(time.Millisecond * 500)
		for aTime := range ticker.C {
			if err := json.NewEncoder(c.Response()).Encode(aTime.String()); err != nil {
				return err
			}
			c.Response().Flush()
		}
		return nil
	})

	e.Logger.Fatal(e.Start(":1323"))
}
