//go:build dev

package web

import (
	"os/exec"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func init() {
	cmd := exec.Command("pnpm", "dev")
	cmd.Dir = "./web"
	cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

	err := cmd.Run()
    if err != nil {
		log.Fatalf("Error: %s\n", err)
    }

	// defer
}

func Serve(c *fiber.Ctx, path string) error {
	url := "http://localhost:3001" + path
	return proxy.Forward(url)(c)
}
