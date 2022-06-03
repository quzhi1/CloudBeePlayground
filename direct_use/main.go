package main

import (
	"fmt"
	"os"

	"github.com/rollout/rox-go/v5/core/context"
	"github.com/rollout/rox-go/v5/server"
)

// Create Roxflags in the Flags container class
type Flags struct {
	EnableTutorial server.RoxFlag
	TitleColors    server.RoxString
	Page           server.RoxInt
	Percentage     server.RoxDouble
}

func main() {
	// Create flag
	flags := Flags{
		// Define the feature flags
		EnableTutorial: server.NewRoxFlag(false),
		// TitleColors:    server.NewRoxString("Green", []string{"White", "Blue", "Green", "Yellow"}),
		// Page:           server.NewRoxInt(1, []int{1, 2, 3}),
		// Percentage:     server.NewRoxDouble(99.9, []float64{10.5, 50.0, 99.9}),
	}

	// Register the flags container
	rox := server.NewRox()
	rox.RegisterWithEmptyNamespace(flags)
	rox.SetCustomComputedStringProperty("flag_attribute_a", func(context context.Context) string {
		value := context.Get("flag_attribute_a")
		if strVal, ok := value.(string); ok {
			return strVal
		} else {
			return ""
		}
	})

	// Setup the environment key
	<-rox.Setup(getApiKey(), server.NewRoxOptions(server.RoxOptionsBuilder{}))
	defer rox.Shutdown()

	// Boolean flag example
	ctx := context.NewContext(map[string]interface{}{
		"flag_attribute_a": "hello",
	})
	flagEnabled(flags, ctx)
	ctx = context.NewContext(map[string]interface{}{
		"flag_attribute_a": "world",
	})
	flagEnabled(flags, ctx)

	// String flag example
	// fmt.Println("TitleColors is " + flags.TitleColors.GetValue(nil))

	// Int flag example
	// fmt.Printf("Page is %v", flags.Page.GetValue(nil))

	// Double flag example
	// fmt.Printf("Percentage is %v", flags.Percentage.GetValue(nil))

}

func getApiKey() string {
	dat, err := os.ReadFile("api_key")
	if err != nil {
		panic(err)
	}
	return string(dat)
}

func flagEnabled(flags Flags, ctx context.Context) {
	if flags.EnableTutorial.IsEnabled(ctx) {
		fmt.Printf("EnableTutorial flag is true for context flag_attribute_a == %s\n", ctx.Get("flag_attribute_a"))
	} else {
		fmt.Printf("EnableTutorial flag is false for context flag_attribute_a == %s\n", ctx.Get("flag_attribute_a"))
	}
}
