package utils

import (
	"fmt"
	global "github.com/entity-toolkit/ntt-dploy/pages"
	"github.com/haykh/tuigo"
	"os"
	"path/filepath"
	"strings"
)

func Finalizer(cs map[tuigo.AppState]tuigo.Window) tuigo.Window {
	messages := []tuigo.Component{}
	pwd, _ := os.Getwd()
	if global.ModulePath != "" {
		temporary_modpath := filepath.Join(pwd, "temp/modules")
		for _, template := range global.Configs {
			if template != nil {
				WriteToFile(
					filepath.Join(temporary_modpath, template.Suffix()),
					template.Module(),
				)
			}
		}
		messages = append(
			messages,
			tuigo.Text("Modules have been written to `"+temporary_modpath+"`", tuigo.LabelText),
		)
		messages = append(messages,
			tuigo.Text(
				fmt.Sprintf(
					"Run `mv %s/* %s` to move the modules to the correct location",
					temporary_modpath,
					global.ModulePath,
				),
				tuigo.NormalText,
			),
		)
	}
	scriptpath := filepath.Join(pwd, "temp")
	buildfiles := []string{}
	for _, mod := range []string{"Kokkos", "MPI", "HDF5", "ADIOS2"} {
		template := global.Configs[mod]
		if template != nil {
			buildfile := "build_" + strings.ToLower(mod) + ".sh"
			buildfiles = append(buildfiles, buildfile)
			WriteToFile(
				filepath.Join(scriptpath, buildfile),
				template.BuildScript(),
			)
		}
	}
	if len(buildfiles) > 0 {
		messages = append(
			messages,
			tuigo.Text("Build scripts have been written to `"+scriptpath+"`", tuigo.LabelText),
		)
		messages = append(
			messages,
			tuigo.Text(
				`Run the following commands to build and install the libraries
% cd `+scriptpath+`
% chmod +x `+strings.Join(buildfiles, " ")+`
% ./`+strings.Join(buildfiles, " && ./"),
				tuigo.NormalText,
			),
		)

	}
	return tuigo.Container(
		tuigo.NonFocusable,
		tuigo.VerticalContainer,
		messages...,
	)
}
