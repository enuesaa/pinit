//go:build dev

package web

import (
	"os/exec"
	"log"
	"fmt"
)

func init() {
	cmd := exec.Command("pnpm", "build")
	cmd.Dir = "./web"
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("%s\n", output)
}

// func (ctl *Controller) ServeStatic(c *fiber.Ctx) error {
// 	fmt.Printf("a\n")

// 	return c.SendString("a")
// }
