package main

import (
	"fmt"
	"os"

	"github.com/rollout/rox-go/v5/core/context"
	"github.com/rollout/rox-go/v5/server"
)

func main() {
	// Init rox
	rox := server.NewRox()
	<-rox.Setup(getApiKey(), server.NewRoxOptions(server.RoxOptionsBuilder{}))
	defer rox.Shutdown()

	// Directly check flags without defining it
	ctx := context.NewContext(map[string]interface{}{
		"flag_attribute_a": "hello",
	})
	if rox.DynamicAPI().IsEnabled("EnableTutorial", false, ctx) {
		fmt.Printf("EnableTutorial flag is true for context flag_attribute_a == %s\n", ctx.Get("flag_attribute_a"))
	} else {
		fmt.Printf("EnableTutorial flag is false for context flag_attribute_a == %s\n", ctx.Get("flag_attribute_a"))
	}

	ctx = context.NewContext(map[string]interface{}{
		"flag_attribute_a": "world",
	})
	if rox.DynamicAPI().IsEnabled("EnableTutorial", false, ctx) {
		fmt.Printf("EnableTutorial flag is true for context flag_attribute_a == %s\n", ctx.Get("flag_attribute_a"))
	} else {
		fmt.Printf("EnableTutorial flag is false for context flag_attribute_a == %s\n", ctx.Get("flag_attribute_a"))
	}
}

func getApiKey() string {
	dat, err := os.ReadFile("api_key")
	if err != nil {
		panic(err)
	}
	return string(dat)
}
