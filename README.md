# golang extensible CLI example

Using the golang plugin package for having an extensible CLI subsystem.

`make build` gets you started and builds two demo plugins and the core CLI!

## How does it work?

Plugins implement this (very simple) interface:
```go
type Plugin interface {
	Command() *cobra.Command
}
```

Golang plugins are compiled into (static!) shared objects, so dependencies can stay within the Plugin as long as the interface does not change.
For it to work, `CGO=1` alongside the `-buildmode=plugin` has to be passed to the build process for plugins.

The resulting artifacts can be loaded using the internal `plugin` package of Golang.

TL;DR: This loads a shared plugin and parses it.
```go
plug, _ := plugin.Open("/path/to/my/plugin.so")
rawExt, _ := plug.Lookup("Plugin") // load exported variable

rootCmd.AddCommand(ext.Command()) // Retrieve the Cobra Command and add it to the root command dynamically
```
