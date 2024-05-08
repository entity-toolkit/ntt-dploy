package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/entity-toolkit/ntt-dploy/configs"
	"github.com/entity-toolkit/ntt-dploy/pages"
	"github.com/entity-toolkit/ntt-dploy/utils"
	"github.com/haykh/tuigo"
)

func mkdir(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}

func main() {

	// DependencyMapping := map[int](map[string]interface{}){
	// 	10: configs.KokkosConfigs,
	// 	11: configs.MPIConfigs,
	// 	12: configs.ADIOS2Configs,
	// 	13: configs.HDF5Configs,
	// }

	// var is_debug = false
	// var has_mpi = false
	// var has_output = false
	// var use_cuda = false
	// var cpu_arch, gpu_arch, archs string
	// var install_mode string
	// var kokkos_path string
	// var kokkos_src string
	// var adios2_src string
	// var kokkos_install_path string
	// var adios2_install_path string

	backend := tuigo.Backend{
		States: []tuigo.AppState{"mode", "specs", "depends", "review"},
		Constructors: map[tuigo.AppState]tuigo.Constructor{
			"mode":    pages.ModeSelector,
			"specs":   pages.SpecsSelector,
			"depends": pages.DependenciesSelector,
			"review":  pages.Review},
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
					for _, id := range []int{1050, 1051, 1052, 1053} {
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
					has_mpi := configs.Global["has_mpi"].(bool)
					cpu_arch := configs.Global["cpu_arch"].(string)
					gpu_arch := configs.Global["gpu_arch"].(string)
					if exists, is_src := input_exists_and_is_srcpath(id_inp); exists {
						conf := configs.DependencyMapping[id_inp]
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
					hide_unhide_src(1000, 1050)
					change_module(1000, 1150)
					hide_unhide_all()
				case configs.OptMPIUpd:
					hide_unhide_src(1001, 1051)
					change_module(1001, 1151)
					hide_unhide_all()
				case configs.OptADIOS2Upd:
					hide_unhide_src(1002, 1052)
					change_module(1002, 1152)
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
			install_mode := configs.Global["install_mode"].(string)
			is_debug := configs.Global["is_debug"].(bool)
			has_mpi := configs.Global["has_mpi"].(bool)
			use_cuda := configs.Global["use_cuda"].(bool)
			cpu_arch := configs.Global["cpu_arch"].(string)
			gpu_arch := configs.Global["gpu_arch"].(string)
			kokkos_src := configs.Global["kokkos_src"].(string)
			kokkos_path := configs.Global["kokkos_path"].(string)
			kokkos_install_path := configs.Global["kokkos_install_path"].(string)
			adios2_src := configs.Global["adios2_src"].(string)
			adios2_install_path := configs.Global["adios2_install_path"].(string)
			if install_mode == "Entity" {
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
							conf := configs.DependencyMapping[id]
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
						conf := configs.DependencyMapping[id]
						sfx := conf["Suffix"].(configs.SuffixFunc)(has_mpi, cpu_arch, gpu_arch)
						script := conf["BuildScript"].(configs.BuildScript)
						src := paths[0]
						inst := paths[1]
						prebuild, configure, build, postbuild := script(src, inst, is_debug, has_mpi, cpu_arch, gpu_arch, true, modules_map)
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
						mkdir("temp")
						utils.WriteToFile("temp/build-"+labels[id]+".sh", lines)
						utils.WriteToFile("temp/module-"+labels[id], module)
					}
				}

				return tuigo.Container(
					false,
					tuigo.VerticalContainer,
					tuigo.Text("All Done!", tuigo.NormalText),
				)
			} else if install_mode == "Kokkos" {
				_, configure, build, postbuild := configs.DependencyMapping[10]["BuildScript"].(configs.BuildScript)(
					kokkos_src,
					kokkos_install_path,
					is_debug,
					false,
					cpu_arch,
					gpu_arch,
					false,
					map[string]string{},
				)
				lines := []string{
					"## Kokkos installation script",
					"# configure",
					strings.Join(configure, "\n"),
					"# build",
					strings.Join(build, "\n"),
					"# postbuild",
					strings.Join(postbuild, "\n"),
				}
				mkdir("temp")
				utils.WriteToFile("temp/build-kokkos.sh", lines)
				return tuigo.Container(
					false,
					tuigo.VerticalContainer,
					tuigo.Text("All Done! Script has been saved to `temp/build-kokkos.sh`", tuigo.NormalText),
				)
			} else if install_mode == "ADIOS2" {
				var cuda_str string
				if use_cuda {
					cuda_str = "CUDA"
				} else {
					cuda_str = ""
				}
				_, configure, build, postbuild := configs.DependencyMapping[12]["BuildScript"].(configs.BuildScript)(
					adios2_src,
					adios2_install_path,
					false,
					has_mpi,
					"",
					cuda_str,
					false,
					map[string]string{},
				)
				configure[2] = configure[2] + " -D Kokkos_ROOT=" + kokkos_path
				lines := []string{
					"## ADIOS2 installation script",
					"# configure",
					strings.Join(configure, "\n"),
					"# build",
					strings.Join(build, "\n"),
					"# postbuild",
					strings.Join(postbuild, "\n"),
				}
				mkdir("temp")
				utils.WriteToFile("temp/build-adios2.sh", lines)
				return tuigo.Container(
					false,
					tuigo.VerticalContainer,
					tuigo.Text("All Done! Script has been saved to `temp/build-adios2.sh`", tuigo.NormalText),
				)
			} else {
				panic("invalid mode selected")
			}
		},
	}
	program := tea.NewProgram(tuigo.App(backend, false))
	if _, err := program.Run(); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
