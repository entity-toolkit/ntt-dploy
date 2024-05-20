package templates

import (
	"strings"
	"text/template"
)

var func_map = template.FuncMap{
	"lower": strings.ToLower,
	"upper": strings.ToUpper,
	"settings": func(sfx string) string {
		return strings.ToUpper(strings.ReplaceAll(sfx, "/", " @ "))
	},
}

var CreateTemplate = func(name, t string) *template.Template {
	return template.Must(template.New(name).Funcs(func_map).Parse(t))
}

var DEBUG_suffix = "{{if .debug}}/debug{{end}}"
var MPI_suffix = "/{{if .mpi}}mpi{{else}}nompi{{end}}"
var GPU_suffix = "/{{if .gpu -}}gpu-{{upper .gpu}}{{else}}nogpu{{end}}"
var CPU_suffix = "/cpu-{{upper .cpu}}"

type LibraryConfiguration struct {
	Settings            map[string]interface{}
	CheckTemplate       func(LibraryConfiguration) bool
	InitFlagsTemplate   func(*LibraryConfiguration)
	SuffixTemplate      *template.Template
	ModuleTemplate      *template.Template
	BuildScriptTemplate *template.Template
}

func (lc *LibraryConfiguration) AddModule(mod string) {
	if _, ok := lc.Settings["modules"]; !ok {
		lc.Settings["modules"] = []string{}
	}
	lc.Settings["modules"] = append(lc.Settings["modules"].([]string), mod)
}

func (lc *LibraryConfiguration) AddFlag(flag string) {
	if _, ok := lc.Settings["flags"]; !ok {
		lc.Settings["flags"] = []string{}
	}
	lc.Settings["flags"] = append(lc.Settings["flags"].([]string), flag)
}

func (lc *LibraryConfiguration) AddCFlag(cflag string) {
	if _, ok := lc.Settings["cflags"]; !ok {
		lc.Settings["cflags"] = []string{}
	}
	lc.Settings["cflags"] = append(lc.Settings["cflags"].([]string), cflag)
}

func (lc *LibraryConfiguration) AddSetting(key string, value interface{}) {
	if lc.Settings == nil {
		lc.Settings = make(map[string]interface{})
	}
	lc.Settings[key] = value
}

func (lc *LibraryConfiguration) GenSuffix() {
	if ok := lc.Check(); !ok {
		panic("LibraryConfiguration is not valid")
	}
	var suffix strings.Builder
	lc.SuffixTemplate.Execute(&suffix, lc.Settings)
	lc.AddSetting("suffix", suffix.String())
}

func (lc *LibraryConfiguration) InitFlags() {
	lc.InitFlagsTemplate(lc)
}

func (lc LibraryConfiguration) Check() bool {
	return lc.CheckTemplate(lc)
}

func (lc LibraryConfiguration) Suffix() string {
	if suffix, ok := lc.Settings["suffix"].(string); ok {
		return suffix
	} else {
		lc.GenSuffix()
		return lc.Settings["suffix"].(string)
	}
}

func (lc LibraryConfiguration) Module() string {
	if ok := lc.Check(); !ok {
		panic("LibraryConfiguration is not valid")
	}
	var module strings.Builder
	lc.ModuleTemplate.Execute(&module, lc.Settings)
	return module.String()
}

func (lc LibraryConfiguration) BuildScript() string {
	if ok := lc.Check(); !ok {
		panic("LibraryConfiguration is not valid")
	}
	var script strings.Builder
	lc.BuildScriptTemplate.Execute(&script, lc.Settings)
	return script.String()
}
