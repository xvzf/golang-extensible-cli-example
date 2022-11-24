package main

import (
	"log"
	"plugin"

	"github.com/xvzf/golang-plugin-test/pkg/ext"
)

// this can be dynamic
var ensurePlugins = []string{
	// dynamic compiled object; they can have their own resources and just have to implement an interface
	"./fiz.so",
	"./buz.so",
}

func main() {
	// Load plugins

	for _, ensure := range ensurePlugins {
		plug, err := plugin.Open(ensure)
		if err != nil {
			log.Panic(err)
		}
		log.Println("Loaded plugin from " + ensure)

		// Retrieve public Plugin var from shared object
		rawExt, err := plug.Lookup("Plugin")
		if err != nil {
			log.Panic(err)
		}

		// Cast to interface
		if ext, ok := rawExt.(ext.Plugin); ok {
			rootCmd.AddCommand(ext.Command())
		} else {
			log.Fatalln("Failed to cast loaded extension")
		}
	}

	// Execute
	Execute()
}
