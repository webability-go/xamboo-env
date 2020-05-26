package main

import (
	"fmt"

	"github.com/webability-go/xamboo/assets"
	"github.com/webability-go/xcore/v2"
)

const VERSION = "1.0.0"

func init() {
	fmt.Println("External APP Main SO library initialized, VERSION =", VERSION)

	// Load CONFIG ofr DatasourceSet
}

var Application = LocalApp{}

type LocalApp struct{}

func (la *LocalApp) StartHost(h assets.Host) {
	fmt.Println("External APP Main SO library started, HOST =", h.Name, "VERSION =", VERSION)
}

func (la *LocalApp) StartContext(ctx *assets.Context) {
	fmt.Println("External APP Context Start, URL=", ctx.Request.URL)
}

func (la *LocalApp) GetDatasourcesConfigFile() string {
	return ""
}

func (la *LocalApp) GetDatasourceSet() assets.DatasourceSet {
	return nil
}

func (la *LocalApp) GetCompiledModules() assets.ModuleSet {
	return nil
}

// Local Function to call from the page
func GetPageData(ctx *assets.Context, template *xcore.XTemplate, language *xcore.XLanguage, e interface{}) string {
	fmt.Println("Distributes a page data called by a page library from app.go")

	return "This is the code of the external application after build all what you need into it. This is a shared library compiled as a plugin."
}
