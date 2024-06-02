package utils

import (
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
		modfiles := []string{}
		resolved_modpath := strings.ReplaceAll(global.ModulePath, "$HOME", os.Getenv("HOME"))
		resolved_modpath = strings.ReplaceAll(resolved_modpath, "~", os.Getenv("HOME"))
		resolved_modpath = strings.ReplaceAll(resolved_modpath, "$USER", os.Getenv("USER"))
		resolved_modpath = strings.ReplaceAll(resolved_modpath, "$PWD", pwd)
		for _, template := range global.Configs {
			if template != nil {
				modfile := filepath.Join(resolved_modpath, template.Suffix())
				modfiles = append(modfiles, template.Suffix())
				WriteToFile(
					modfile,
					template.Module(),
				)
			}
		}
		messages = append(
			messages,
			tuigo.Text("The following modules have been written to `"+resolved_modpath+"`:", tuigo.LabelText),
		)
		for _, modfile := range modfiles {
			messages = append(
				messages,
				tuigo.Text(modfile, tuigo.NormalText),
			)
		}
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
			msg := global.HowTo[mod]
			if mod == "HDF5" {
				msg = "1. Create a directory `" + global.Configs[mod].Settings["src_path"].(string) + "`\n" + msg
			} else {
				msg = msg + " to `" + global.Configs[mod].Settings["src_path"].(string) + "`"
			}
			messages = append(
				messages,
				tuigo.Container(
					tuigo.NonFocusable,
					tuigo.HorizontalContainerTop,
					tuigo.Text(mod+":", tuigo.LabelText),
					tuigo.Text(msg, tuigo.NormalText),
				),
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
				`[!] Make sure to follow the instructions above first to download the source codes

Run the following commands to build and install the libraries

% cd `+scriptpath+`
% source ./`+strings.Join(buildfiles, " && source ./")+"\n",
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
