package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/entity-toolkit/ntt-dploy/configs"
	"github.com/entity-toolkit/ntt-dploy/utils"
	"github.com/haykh/tuigo"
)

func main() {
	DependencyMapping := map[int](map[string]interface{}){
		10: configs.KokkosConfigs,
		11: configs.MPIConfigs,
		12: configs.ADIOS2Configs,
		13: configs.HDF5Configs,
	}

	var GetArchs = func(archstr string) (string, string) {
		archs := []string{}
		for _, arch := range strings.Split(archstr, ",") {
			arch = strings.TrimSpace(arch)
			if arch != "" {
				archs = append(archs, arch)
			}
		}
		if len(archs) == 1 {
			archs = append(archs, "")
		}
		return archs[0], archs[1]
	}
	var is_debug = false
	var has_mpi = false
	var has_output = false
	var cpu_arch, gpu_arch, archs string

	b := tuigo.Backend{
		States: []tuigo.AppState{"specs", "depends", "review"},
		Constructors: map[tuigo.AppState]tuigo.Constructor{
			"specs": func(tuigo.Window) tuigo.Window {
				cpusel := tuigo.SelectorWithID(4, configs.CpuArchs, false, nil)
				gpusel := tuigo.SelectorWithID(5, configs.GpuArchs, false, nil)
				return tuigo.Container(
					true,
					tuigo.VerticalContainer,
					tuigo.Text("1. Pick Entity configuration to install", tuigo.LabelText),
					tuigo.Text("options:", tuigo.NormalText),
					tuigo.RadioWithID(1, "use mpi", nil),
					tuigo.RadioWithID(2, "enable output", nil),
					tuigo.RadioWithID(3, "enable debug mode", nil),
					tuigo.Container(true, tuigo.HorizontalContainerTop,
						tuigo.Container(true, tuigo.VerticalContainer,
							tuigo.Text("CPU architecture:", tuigo.NormalText),
							cpusel,
						),
						tuigo.Container(true, tuigo.VerticalContainer,
							tuigo.Text("GPU architecture:", tuigo.NormalText),
							gpusel,
						),
					),
				)
			},
			"depends": func(prev tuigo.Window) tuigo.Window {
				has_mpi = prev.GetElementByID(1).(tuigo.RadioElement).Data().(bool)
				has_output = prev.GetElementByID(2).(tuigo.RadioElement).Data().(bool)
				is_debug = prev.GetElementByID(3).(tuigo.RadioElement).Data().(bool)
				cpu := prev.GetElementByID(4).(tuigo.SelectorElement).Data()
				gpu := prev.GetElementByID(5).(tuigo.SelectorElement).Data()
				var cpu_str = "NATIVE"
				if cpu != nil {
					cpu_str = cpu.(string)
				}
				archs += cpu_str
				if gpu != nil {
					archs += ", " + gpu.(string)
				}
				cpu_arch, gpu_arch = GetArchs(archs)

				cfg := configs.EntityConfigs
				entsfx := cfg["Suffix"].(configs.EntitySuffixFunc)(is_debug, has_mpi, cpu_arch, gpu_arch)

				elements := tuigo.Components{}
				elements = append(elements,
					tuigo.InputWithID(6, cfg["Name"].(string)+" module parent directory", configs.DefaultModPath, "", tuigo.PathInput, nil),
				)
				elements = append(elements,
					tuigo.Container(false, tuigo.HorizontalContainer,
						tuigo.TextWithID(60, "  "+cfg["Name"].(string)+" module path: ", tuigo.NormalText),
						tuigo.TextWithID(61, ".../"+entsfx, tuigo.NormalText),
					),
				)
				elements = append(elements,
					tuigo.TextWithID(500, "  architectures: "+archs, tuigo.NormalText),
				)
				elements = append(elements,
					tuigo.Text("Dependencies:", tuigo.NormalText),
				)
				elements = append(elements,
					tuigo.InputWithID(7, "  CXX module", configs.DefaultCXX, "", tuigo.TextInput, nil),
				)
				if gpu != nil {
					elements = append(elements,
						tuigo.InputWithID(8, "  CUDA module", configs.DefaultCUDA, "", tuigo.TextInput, nil),
					)
				}
				inputs := tuigo.Components{}
				paths_modules := tuigo.Components{
					tuigo.TextWithID(900, "Install & module paths:", tuigo.NormalText),
					tuigo.InputWithID(9, "  parent directory", configs.DefaultOptPath, "", tuigo.PathInput, nil),
				}
				for _, id := range []int{10, 11, 12, 13} {
					conf := DependencyMapping[id]
					if conf["Condition"].(configs.ConditionFunc)(has_mpi, has_output) {
						name := conf["Name"].(string)
						inputs = append(inputs,
							tuigo.InputWithID(
								id,
								"  "+name+" src/module",
								conf["DefaultSrcPath"].(string),
								"",
								tuigo.PathInput,
								conf["UpdateMsg"].(tea.Msg),
							),
						)
						sfx := conf["Suffix"].(configs.SuffixFunc)(has_mpi, cpu_arch, gpu_arch)
						paths_modules = append(paths_modules,
							tuigo.TextWithID(id*10+1, "  "+name+" install path: .../"+sfx, tuigo.NormalText),
							tuigo.TextWithID(id*10+2, "  "+name+" module: .../.modules/"+sfx, tuigo.NormalText),
						)
					}
				}
				elements = append(elements, inputs...)
				elements = append(elements, paths_modules...)
				elements = append(
					tuigo.Components{tuigo.Text("2. Specify paths, compilers & modules to use", tuigo.LabelText)},
					elements...,
				)
				elements = append(
					elements,
					tuigo.Text("[?] to use preinstalled dependencies using modules, use `module:<MODULENAME>`", tuigo.DimmedText),
				)
				return tuigo.Container(
					true,
					tuigo.VerticalContainer,
					elements...,
				)
			},
			"review": func(prev tuigo.Window) tuigo.Window {
				nttpath := prev.GetElementByID(61).(tuigo.TextElement).Data().(string)
				var check = func(cond bool) string {
					if cond {
						return " [✔]"
					} else {
						return " [ ]"
					}
				}
				external_modules := tuigo.Components{}
				compiled_modules := tuigo.Components{}
				new_modules := tuigo.Components{}
				labels := map[int]string{
					7: "CXX", 8: "CUDA", 10: "Kokkos", 11: "MPI", 12: "ADIOS2", 13: "HDF5",
				}
				for _, id := range []int{7, 8, 10, 11, 12, 13} {
					if el := prev.GetElementByID(id); el != nil {
						el_str := el.(tuigo.InputElement).Data().(string)
						if len(el_str) > 7 && el_str[:7] == "module:" {
							modname := el_str[7:]
							external_modules = append(external_modules, tuigo.Text("... "+labels[id]+": "+modname, tuigo.NormalText))
						} else if id < 10 {
							panic("invalid input: module not defined for `" + labels[id] + "`")
						} else {
							srcpath := el_str
							conf := DependencyMapping[id]
							name := conf["Name"].(string)
							sfx := conf["Suffix"].(configs.SuffixFunc)(has_mpi, cpu_arch, gpu_arch)
							optpath := prev.GetElementByID(9).(tuigo.InputElement).Data().(string)
							if optpath[len(optpath)-1] != '/' {
								optpath += "/"
							}
							compiled_modules = append(compiled_modules, tuigo.Text("... "+name+": "+srcpath+" → "+optpath+sfx, tuigo.NormalText))
							new_modules = append(new_modules, tuigo.Text("... "+name+": "+optpath+".modules/"+sfx, tuigo.NormalText))
						}
					}
				}
				if len(external_modules) > 0 {
					external_modules = append(tuigo.Components{tuigo.Text("External modules:", tuigo.NormalText)}, external_modules...)
				}
				if len(compiled_modules) > 0 {
					compiled_modules = append(tuigo.Components{tuigo.Text("Compiled dependencies:", tuigo.NormalText)}, compiled_modules...)
					new_modules = append(tuigo.Components{tuigo.Text("New modules:", tuigo.NormalText)}, new_modules...)
				}
				return tuigo.Container(
					false, tuigo.VerticalContainer,
					append(
						append(
							append(
								tuigo.Components{
									tuigo.Text("3. Review the configuration", tuigo.LabelText),
									tuigo.Text("Entity: ", tuigo.NormalText),
									tuigo.Text("... archs: "+archs, tuigo.NormalText),
									tuigo.Text("... output enabled: "+check(has_output), tuigo.NormalText),
									tuigo.Text("... MPI enabled: "+check(has_mpi), tuigo.NormalText),
									tuigo.Text("... debug mode: "+check(is_debug), tuigo.NormalText),
									tuigo.Text("... modulefile: "+nttpath, tuigo.NormalText),
								},
								external_modules...,
							),
							compiled_modules...,
						),
						new_modules...,
					)...,
				)
			},
		},
		Updaters: map[tuigo.AppState]tuigo.Updater{
			"depends": func(window tuigo.Window, msg tea.Msg) (tuigo.Window, tea.Cmd) {
				cmds := []tea.Cmd{}
				var input_exists_and_is_srcpath = func(id int) (bool, bool) {
					if el := window.GetElementByID(id); el != nil {
						el_str := el.(tuigo.InputElement).Data().(string)
						return true, (len(el_str) < 7) || (el_str[:7] != "module:")
					}
					return true, false
				}
				var hide_unhide_src = func(id_inp, id_src int) {
					if exists, is_src := input_exists_and_is_srcpath(id_inp); exists {
						cmds = append(cmds, tuigo.TgtCmd(
							id_src,
							func(cont tuigo.Wrapper, mod tuigo.Accessor) (tuigo.Wrapper, tuigo.Accessor) {
								if is_src && cont.Hidden() {
									return cont.Unhide().(tuigo.Wrapper), mod
								} else if !is_src && !cont.Hidden() {
									return cont.Hide().(tuigo.Wrapper), mod
								}
								return cont, mod
							},
						))
					}
				}
				var hide_unhide_all = func() {
					all_hidden := true
					for _, id := range []int{10, 11, 12, 13} {
						if exists, is_src := input_exists_and_is_srcpath(id); exists {
							all_hidden = all_hidden && !is_src
						}
					}
					cmds = append(cmds, tuigo.TgtCmd(
						9,
						func(cont tuigo.Wrapper, mod tuigo.Accessor) (tuigo.Wrapper, tuigo.Accessor) {
							if all_hidden {
								return cont.Hide().(tuigo.Wrapper), mod
							} else {
								return cont.Unhide().(tuigo.Wrapper), mod
							}
						},
					))
					cmds = append(cmds, tuigo.TgtCmd(
						900,
						func(cont tuigo.Wrapper, mod tuigo.Accessor) (tuigo.Wrapper, tuigo.Accessor) {
							if all_hidden {
								return cont.Hide().(tuigo.Wrapper), mod
							} else {
								return cont.Unhide().(tuigo.Wrapper), mod
							}
						},
					))
				}
				var change_module = func(id_inp, id_mod int) {
					if exists, is_src := input_exists_and_is_srcpath(id_inp); exists {
						conf := DependencyMapping[id_inp]
						if !is_src {
							prefix := "  " + conf["Name"].(string)
							el := window.GetElementByID(id_inp)
							el_str := el.(tuigo.InputElement).Data().(string)
							modname := el_str[7:]
							if modname == "" {
								modname = "N/A"
							}
							newmod := prefix + " module: " + modname
							cmds = append(cmds, tuigo.TgtCmd(
								id_mod,
								func(cont tuigo.Wrapper, mod tuigo.Accessor) (tuigo.Wrapper, tuigo.Accessor) {
									return cont, mod.(tuigo.TextElement).Set(newmod)
								},
							))
						} else {
							name := conf["Name"].(string)
							sfx := conf["Suffix"].(func(bool, string, string) string)(has_mpi, cpu_arch, gpu_arch)
							cmds = append(cmds, tuigo.TgtCmd(
								id_mod,
								func(cont tuigo.Wrapper, mod tuigo.Accessor) (tuigo.Wrapper, tuigo.Accessor) {
									return cont, mod.(tuigo.TextElement).Set("  " + name + " module: .../.modules/" + sfx)
								},
							))
						}

					}
				}
				switch msg.(type) {
				case configs.OptKokkosUpd:
					hide_unhide_src(10, 101)
					change_module(10, 102)
					hide_unhide_all()
				case configs.OptMPIUpd:
					hide_unhide_src(11, 111)
					change_module(11, 112)
					hide_unhide_all()
				case configs.OptADIOS2Upd:
					hide_unhide_src(12, 121)
					change_module(12, 122)
					hide_unhide_src(12, 131)
					hide_unhide_src(12, 13)
					hide_unhide_all()
				case configs.OptHDF5Upd:
					hide_unhide_src(13, 131)
					change_module(13, 132)
					hide_unhide_all()
				}
				return window, tea.Batch(cmds...)
			},
		},
		Finalizer: func(cs map[tuigo.AppState]tuigo.Window) tuigo.Window {
			depend_window := cs["depends"]
			labels := map[int]string{
				7: "CXX", 8: "CUDA", 10: "Kokkos", 11: "MPI", 12: "ADIOS2", 13: "HDF5",
			}
			optpath := depend_window.GetElementByID(9).(tuigo.InputElement).Data().(string)
			if optpath[len(optpath)-1] != '/' {
				optpath += "/"
			}
			modules := []string{}
			modules_map := map[string]string{}
			tobuild := map[string]([]string){}
			for _, id := range []int{7, 8, 10, 11, 12, 13} {
				if el := depend_window.GetElementByID(id); el != nil {
					el_str := el.(tuigo.InputElement).Data().(string)
					if len(el_str) > 7 && el_str[:7] == "module:" {
						modname := el_str[7:]
						modules = append(modules, modname)
						modules_map[labels[id]] = modname
					} else if id < 10 {
						panic("invalid input: module not defined for `" + labels[id] + "`")
					} else {
						conf := DependencyMapping[id]
						sfx := conf["Suffix"].(configs.SuffixFunc)(has_mpi, cpu_arch, gpu_arch)
						modname := optpath + ".modules/" + sfx
						modules = append(modules, modname)
						modules_map[labels[id]] = modname
						tobuild[labels[id]] = []string{el_str, optpath + sfx}
					}
				}
			}
			entsfx := configs.EntityConfigs["Suffix"].(configs.EntitySuffixFunc)(is_debug, has_mpi, cpu_arch, gpu_arch)
			modulefile := configs.EntityConfigs["ModuleTemplate"].(configs.EntityModTemplate)(
				entsfx,
				is_debug,
				has_mpi,
				cpu_arch,
				gpu_arch,
				modules,
			)
			utils.WriteToFile("temp/module-entity", modulefile)
			for _, id := range []int{10, 11, 12, 13} {
				if paths, ok := tobuild[labels[id]]; ok {
					conf := DependencyMapping[id]
					sfx := conf["Suffix"].(configs.SuffixFunc)(has_mpi, cpu_arch, gpu_arch)
					script := conf["BuildScript"].(configs.BuildScript)
					src := paths[0]
					inst := paths[1]
					prebuild, configure, build, postbuild := script(src, inst, is_debug, has_mpi, cpu_arch, gpu_arch, modules_map)
					lines := []string{
						"## " + conf["Name"].(string),
						"# prebuild",
						strings.Join(prebuild, "\n"),
						"# configure",
						strings.Join(configure, "\n"),
						"# build",
						strings.Join(build, "\n"),
						"# postbuild",
						strings.Join(postbuild, "\n"),
					}
					module := conf["ModuleTemplate"].(configs.ModuleTemplate)(
						sfx,
						inst,
						is_debug,
						has_mpi,
						cpu_arch,
						gpu_arch,
						modules_map,
					)
					utils.WriteToFile("temp/build-"+labels[id]+".sh", lines)
					utils.WriteToFile("temp/module-"+labels[id], module)
				}
			}

			return tuigo.Container(
				false,
				tuigo.VerticalContainer,
				tuigo.Text("All Done!", tuigo.NormalText),
			)
		},
	}
	p := tea.NewProgram(tuigo.App(b, false))
	if _, err := p.Run(); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
