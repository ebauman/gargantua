package apiserver

import "github.com/gofiber/fiber"

func RunServer(port int, stop chan interface{}) {
	app := fiber.New()

	app.Use(func (c *fiber.Ctx) {
		// error handler
		if c.Error() == nil {
			c.Next() // don't process errors if one does not exist
		}

		// an error exists
		// if type is HFError, extract status and so on

	})


	// app registrations go here



}

