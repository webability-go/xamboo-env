package bridge

import (
	"errors"
	"fmt"
	"net/http"
	"plugin"

	"github.com/webability-go/xamboo/cms/context"
	"github.com/webability-go/xcore/v2"
)

/* This package declare all the available functions of the app to be able to call them. */
/* Include this package when you want to call the app */

var linked bool = false

var GetPageData func(*context.Context, *xcore.XTemplate, *xcore.XLanguage, interface{}) string

func Setup(ctx *context.Context) bool {

	// Ask for the plugins we need
	app, ok := ctx.Plugins["app"]
	if !ok {
		// 500 internal error
		http.Error(ctx.Writer, "Library app not available", http.StatusInternalServerError)
		return false
	}

	// Initialize the plugin (just in case)
	err := Start(app)
	if err != nil {
		// 500 internal error
		http.Error(ctx.Writer, "Library not linked", http.StatusInternalServerError)
		return false
	}

	// You may add here login process, verify user language, device, IPs, session, etc.

	return true
}

func Start(lib *plugin.Plugin) error {
	if linked {
		return nil
	}

	fmt.Println("Linking the app plugin to mapped functions")

	fct, err := lib.Lookup("GetPageData")
	if err != nil {
		fmt.Println(err)
		return errors.New("ERROR: THE APPLICATION LIBRARY DOES NOT CONTAIN GETPAGEDATA FUNCTION")
	}
	GetPageData = fct.(func(*context.Context, *xcore.XTemplate, *xcore.XLanguage, interface{}) string)

	linked = true
	return nil
}
