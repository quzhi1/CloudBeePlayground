package main

import (
	"fmt"
	"os"

	"github.com/rollout/rox-go/v5/core/context"
	"github.com/rollout/rox-go/v5/server"
)

// type Flags struct{}

func main() {
	// Init rox
	rox := server.NewRox()
	// rox.RegisterWithEmptyNamespace(Flags{})
	<-rox.Setup(getApiKey(), server.NewRoxOptions(server.RoxOptionsBuilder{}))
	defer rox.Shutdown()

	// Directly check flags without defining it
	ctx := context.NewContext(map[string]interface{}{
		"flag_attribute_b": "hello",
	})
	if rox.DynamicAPI().IsEnabled("DynamicApiTest", false, ctx) {
		fmt.Printf("DynamicApiTest flag is true for context flag_attribute_b == %s\n", ctx.Get("flag_attribute_b"))
	} else {
		fmt.Printf("DynamicApiTest flag is false for context flag_attribute_b == %s\n", ctx.Get("flag_attribute_b"))
	}

	ctx = context.NewContext(map[string]interface{}{
		"flag_attribute_b": "world",
	})
	if rox.DynamicAPI().IsEnabled("DynamicApiTest", false, ctx) {
		fmt.Printf("DynamicApiTest flag is true for context flag_attribute_b == %s\n", ctx.Get("flag_attribute_b"))
	} else {
		fmt.Printf("DynamicApiTest flag is false for context flag_attribute_b == %s\n", ctx.Get("flag_attribute_b"))
	}
}

func getApiKey() string {
	dat, err := os.ReadFile("api_key")
	if err != nil {
		panic(err)
	}
	return string(dat)
}
