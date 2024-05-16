package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/entity-toolkit/ntt-dploy/pages"
	"github.com/haykh/tuigo"
)

// func mkdir(path string) {
// 	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
// 		err := os.Mkdir(path, os.ModePerm)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	}
// }

func main() {
	backend := tuigo.Backend{
		States: []tuigo.AppState{"main", "review"},
		Constructors: map[tuigo.AppState]tuigo.Constructor{
			"main":   pages.MainSelector,
			"review": pages.Review},
		Updaters: map[tuigo.AppState]tuigo.Updater{
			"main": pages.MainUpdater,
		},
		Finalizer: func(cs map[tuigo.AppState]tuigo.Window) tuigo.Window {
			return tuigo.Container(
				tuigo.NonFocusable,
				tuigo.VerticalContainer,
				tuigo.Text("All Done!", tuigo.NormalText),
			)
		},
	}

	program := tea.NewProgram(tuigo.App(backend, false))
	if _, err := program.Run(); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}

// install_mode := global.MODE.Value.(string)
// is_debug := global.DEBUG.Value.(bool)
// has_mpi := global.MPI.Value.(bool)
// use_cuda := global.CUDA.Value.(bool)
// cpu_arch := global.CPUARCH.Value.(string)
// gpu_arch := global.GPUARCH.Value.(string)
// kokkos_src := global.KOKKOS_SRC_DIR.Value.(string)
// kokkos_path := global.KOKKOS_INSTALL_DIR.Value.(string)
// adios2_src := global.ADIOS2_SRC_DIR.Value.(string)
// adios2_install_path := global.ADIOS2_INSTALL_DIR.Value.(string)
// if install_mode == "Entity" {
// 	depend_window := cs["depends"]
// 	labels := map[int]string{
// 		7: "CXX", 8: "CUDA", 10: "Kokkos", 11: "MPI", 12: "ADIOS2", 13: "HDF5",
// 	}
// 	optpath := depend_window.GetElementByID(9).(tuigo.InputElement).Data().(string)
// 	if optpath[len(optpath)-1] != '/' {
// 		optpath += "/"
// 	}
// 	modules := []string{}
// 	modules_map := map[string]string{}
// 	tobuild := map[string]([]string){}
// 	for _, id := range []int{7, 8, 10, 11, 12, 13} {
// 		if el := depend_window.GetElementByID(id); el != nil {
// 			el_str := el.(tuigo.InputElement).Data().(string)
// 			if len(el_str) > 7 && el_str[:7] == "module:" {
// 				modname := el_str[7:]
// 				modules = append(modules, modname)
// 				modules_map[labels[id]] = modname
// 			} else if id < 10 {
// 				panic("invalid input: module not defined for `" + labels[id] + "`")
// 			} else {
// 				conf := configs.DependencyMapping[id]
// 				sfx := conf["Suffix"].(configs.SuffixFunc)(has_mpi, cpu_arch, gpu_arch)
// 				modname := optpath + ".modules/" + sfx
// 				modules = append(modules, modname)
// 				modules_map[labels[id]] = modname
// 				tobuild[labels[id]] = []string{el_str, optpath + sfx}
// 			}
// 		}
// 	}
// 	entsfx := configs.EntityConfigs["Suffix"].(configs.EntitySuffixFunc)(is_debug, has_mpi, cpu_arch, gpu_arch)
// 	modulefile := configs.EntityConfigs["ModuleTemplate"].(configs.EntityModTemplate)(
// 		entsfx,
// 		is_debug,
// 		has_mpi,
// 		cpu_arch,
// 		gpu_arch,
// 		modules,
// 	)
// 	utils.WriteToFile("temp/module-entity", modulefile)
// 	for _, id := range []int{10, 11, 12, 13} {
// 		if paths, ok := tobuild[labels[id]]; ok {
// 			conf := configs.DependencyMapping[id]
// 			sfx := conf["Suffix"].(configs.SuffixFunc)(has_mpi, cpu_arch, gpu_arch)
// 			script := conf["BuildScript"].(configs.BuildScript)
// 			src := paths[0]
// 			inst := paths[1]
// 			prebuild, configure, build, postbuild := script(src, inst, is_debug, has_mpi, cpu_arch, gpu_arch, true, modules_map)
// 			lines := []string{
// 				"## " + conf["Name"].(string),
// 				"# prebuild",
// 				strings.Join(prebuild, "\n"),
// 				"# configure",
// 				strings.Join(configure, "\n"),
// 				"# build",
// 				strings.Join(build, "\n"),
// 				"# postbuild",
// 				strings.Join(postbuild, "\n"),
// 			}
// 			module := conf["ModuleTemplate"].(configs.ModuleTemplate)(
// 				sfx,
// 				inst,
// 				is_debug,
// 				has_mpi,
// 				cpu_arch,
// 				gpu_arch,
// 				modules_map,
// 			)
// 			mkdir("temp")
// 			utils.WriteToFile("temp/build-"+labels[id]+".sh", lines)
// 			utils.WriteToFile("temp/module-"+labels[id], module)
// 		}
// 	}

// 	return tuigo.Container(
// 		false,
// 		tuigo.VerticalContainer,
// 		tuigo.Text("All Done!", tuigo.NormalText),
// 	)
// } else if install_mode == "Kokkos" {
// 	_, configure, build, postbuild := configs.DependencyMapping[10]["BuildScript"].(configs.BuildScript)(
// 		kokkos_src,
// 		kokkos_path,
// 		is_debug,
// 		false,
// 		cpu_arch,
// 		gpu_arch,
// 		false,
// 		map[string]string{},
// 	)
// 	lines := []string{
// 		"## Kokkos installation script",
// 		"# configure",
// 		strings.Join(configure, "\n"),
// 		"# build",
// 		strings.Join(build, "\n"),
// 		"# postbuild",
// 		strings.Join(postbuild, "\n"),
// 	}
// 	mkdir("temp")
// 	utils.WriteToFile("temp/build-kokkos.sh", lines)
// 	return tuigo.Container(
// 		false,
// 		tuigo.VerticalContainer,
// 		tuigo.Text("All Done! Script has been saved to `temp/build-kokkos.sh`", tuigo.NormalText),
// 	)
// } else if install_mode == "ADIOS2" {
// 	var cuda_str string
// 	if use_cuda {
// 		cuda_str = "CUDA"
// 	} else {
// 		cuda_str = ""
// 	}
// 	_, configure, build, postbuild := configs.DependencyMapping[12]["BuildScript"].(configs.BuildScript)(
// 		adios2_src,
// 		adios2_install_path,
// 		false,
// 		has_mpi,
// 		"",
// 		cuda_str,
// 		false,
// 		map[string]string{},
// 	)
// 	configure[2] = configure[2] + " -D Kokkos_ROOT=" + kokkos_path
// 	lines := []string{
// 		"## ADIOS2 installation script",
// 		"# configure",
// 		strings.Join(configure, "\n"),
// 		"# build",
// 		strings.Join(build, "\n"),
// 		"# postbuild",
// 		strings.Join(postbuild, "\n"),
// 	}
// 	mkdir("temp")
// 	utils.WriteToFile("temp/build-adios2.sh", lines)
// 	return tuigo.Container(
// 		false,
// 		tuigo.VerticalContainer,
// 		tuigo.Text("All Done! Script has been saved to `temp/build-adios2.sh`", tuigo.NormalText),
// 	)
// } else {
// 	panic("invalid mode selected")
// }
// },
