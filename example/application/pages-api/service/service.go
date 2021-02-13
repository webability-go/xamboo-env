package main

import (
	"net/http"

	"github.com/webability-go/xamboo/cms/context"
	"github.com/webability-go/xcore/v2"

	"github.com/webability-go/xamboo-env/example/app/bridge"
)

/* This function is MANDATORY and is the point of call from the xamboo
   The enginecontext contains all what you need to link with the system
*/
func Run(ctx *context.Context, xtemplate *xcore.XTemplate, xlanguage *xcore.XLanguage, cms interface{}) interface{} {

	if !bridge.Setup(ctx) {
		return "" // error already managed
	}

	switch ctx.Request.Method {
	case "GET":
		return rGet(ctx, xtemplate, xlanguage, cms)
	case "POST":
		return rPost(ctx, xtemplate, xlanguage, cms)
	case "PUT":
		return rPut(ctx, xtemplate, xlanguage, cms)
	case "DELETE":
		return rDelete(ctx, xtemplate, xlanguage, cms)
	case "OPTIONS": // The Component Origin takes care of this
		return ""
	default:
	}
	// 501 not supported
	http.Error(ctx.Writer, "Method not supported", http.StatusNotImplemented)
	return "Method not supported"
}

func rGet(ctx *context.Context, xtemplate *xcore.XTemplate, xlanguage *xcore.XLanguage, cms interface{}) interface{} {
	return "{\"method:\":\"GET\"}"
}

func rPost(ctx *context.Context, xtemplate *xcore.XTemplate, xlanguage *xcore.XLanguage, cms interface{}) interface{} {
	return "{\"method:\":\"POST\"}"
}

func rPut(ctx *context.Context, xtemplate *xcore.XTemplate, xlanguage *xcore.XLanguage, cms interface{}) interface{} {
	return "{\"method:\":\"PUT\"}"
}

func rDelete(ctx *context.Context, xtemplate *xcore.XTemplate, xlanguage *xcore.XLanguage, cms interface{}) interface{} {
	return "{\"method:\":\"DELETE\"}"
}
