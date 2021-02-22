package main

import (
	//  "fmt"
	"github.com/webability-go/xamboo/cms/context"
	"github.com/webability-go/xcore/v2"

	"github.com/webability-go/xamboo-env/example/app/bridge"
)

/* This function is MANDATORY and is the point of call from the xamboo
   The enginecontext contains all what you need to link with the system
*/
func Run(ctx *context.Context, template *xcore.XTemplate, language *xcore.XLanguage, e interface{}) interface{} {

	// Let's call out external app library (you should create a wrapper to your app so it's much easier to access funcions instead or writing all this code here)
	myappdata, ok := ctx.Plugins["app"]
	if !ok {
		return "ERROR: THE APPLICATION LIBRARY IS NOT AVAILABLE"
	}

	bridge.Start(myappdata)
	strdata := bridge.GetPageData(ctx, template, language, e)

	// Let's inject some vars into the template
	data := &xcore.XDataset{}
	data.Set("data1", "Data 1 for the template")
	data.Set("dataapp", strdata)

	return template.Execute(data)
}
