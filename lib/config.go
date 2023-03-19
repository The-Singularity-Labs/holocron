package lib

import (
	"github.com/guark/guark/app"
	"github.com/guark/plugins/clipboard"
	"github.com/guark/plugins/dialog"
	"github.com/guark/plugins/notify"
	"github.com/the-singularity-labs/create-guark-alpinejs-app/lib/funcs"
	"github.com/the-singularity-labs/create-guark-alpinejs-app/lib/hooks"
)

// Exposed functions to guark Javascript api.
var Funcs = app.Funcs{
	"hello_world": funcs.HelloWorld,
}

// App hooks.
var Hooks = app.Hooks{
	"created": hooks.Created,
	"mounted": hooks.Mounted,
}

// App plugins.
var Plugins = app.Plugins{
	"dialog":    &dialog.Plugin{},
	"notify":    &notify.Plugin{},
	"clipboard": &clipboard.Plugin{},
}
