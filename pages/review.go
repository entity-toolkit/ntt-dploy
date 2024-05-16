package pages

import "github.com/haykh/tuigo"

func Review(prev tuigo.Window) tuigo.Window {
	return tuigo.Container(
		tuigo.NonFocusable,
		tuigo.VerticalContainer,
		tuigo.Text("Review", tuigo.LabelText),
	)
}

// func Review(prev tuigo.Window) tuigo.Window {
// 	var check = func(cond bool) string {
// 		if cond {
// 			return " [✔]"
// 		} else {
// 			return " [ ]"
// 		}
// 	}
// 	install_mode := global.MODE.Value.(string)
// 	has_mpi := global.MPI.Value.(bool)
// 	has_output := global.OUTPUT.Value.(bool)
// 	use_cuda, _ := global.CUDA.Value.(bool)
// 	is_debug := global.DEBUG.Value.(bool)
// 	archs := global.ARCHS.Value.(string)
// 	kokkos_path, _ := global.KOKKOS_INSTALL_DIR.Value.(string)
// 	kokkos_src, _ := global.KOKKOS_SRC_DIR.Value.(string)
// 	adios2_src, _ := global.ADIOS2_SRC_DIR.Value.(string)
// 	cpu_arch := global.CPUARCH.Value.(string)
// 	gpu_arch := global.GPUARCH.Value.(string)
// 	// install_mode := configs.Global["install_mode"].(string)
// 	// has_mpi := configs.Global["has_mpi"].(bool)
// 	// has_output := configs.Global["has_output"].(bool)
// 	// use_cuda := configs.Global["use_cuda"].(bool)
// 	// is_debug := configs.Global["is_debug"].(bool)
// 	// archs := configs.Global["archs"].(string)
// 	// kokkos_path := configs.Global["kokkos_path"].(string)
// 	// kokkos_src := configs.Global["kokkos_src"].(string)
// 	// adios2_src := configs.Global["adios2_src"].(string)
// 	// kokkos_install_path := configs.Global["kokkos_install_path"].(string)
// 	// adios2_install_path := configs.Global["adios2_install_path"].(string)
// 	// cpu_arch := configs.Global["cpu_arch"].(string)
// 	// gpu_arch := configs.Global["gpu_arch"].(string)

// 	if install_mode == "Entity" {
// 		nttpath, _ := prev.GetElementByID(global.ENTITY_PATH.Id).(tuigo.TextElement).Data().(string)
// 		external_modules := tuigo.Components{}
// 		compiled_modules := tuigo.Components{}
// 		new_modules := tuigo.Components{}
// 		labels := map[int]string{
// 			global.CXX_MODULE.Id:  "CXX",
// 			global.CUDA_MODULE.Id: "CUDA",
// 			1050:                  "Kokkos",
// 			1051:                  "MPI",
// 			1052:                  "ADIOS2",
// 			1053:                  "HDF5",
// 		}
// 		for id := range labels {
// 			if el := prev.GetElementByID(id); el != nil {
// 				el_str := el.(tuigo.InputElement).Data().(string)
// 				if len(el_str) > 7 && el_str[:7] == "module:" {
// 					modname := el_str[7:]
// 					external_modules = append(external_modules, tuigo.Text("... "+labels[id]+": "+modname, tuigo.NormalText))
// 				} else if id < 10 {
// 					panic("invalid input: module not defined for `" + labels[id] + "`")
// 				} else {
// 					srcpath := el_str
// 					conf := configs.DependencyMapping[id]
// 					name := conf["Name"].(string)
// 					sfx := conf["Suffix"].(configs.SuffixFunc)(has_mpi, cpu_arch, gpu_arch)
// 					optpath := prev.GetElementByID(9).(tuigo.InputElement).Data().(string)
// 					if optpath[len(optpath)-1] != '/' {
// 						optpath += "/"
// 					}
// 					compiled_modules = append(compiled_modules, tuigo.Text("... "+name+": "+srcpath+" → "+optpath+sfx, tuigo.NormalText))
// 					new_modules = append(new_modules, tuigo.Text("... "+name+": "+optpath+".modules/"+sfx, tuigo.NormalText))
// 				}
// 			}
// 		}
// 		if len(external_modules) > 0 {
// 			external_modules = append(tuigo.Components{tuigo.Text("External modules:", tuigo.NormalText)}, external_modules...)
// 		}
// 		if len(compiled_modules) > 0 {
// 			compiled_modules = append(tuigo.Components{tuigo.Text("Compiled dependencies:", tuigo.NormalText)}, compiled_modules...)
// 			new_modules = append(tuigo.Components{tuigo.Text("New modules:", tuigo.NormalText)}, new_modules...)
// 		}
// 		return tuigo.Container(
// 			false, tuigo.VerticalContainer,
// 			append(
// 				append(
// 					append(
// 						tuigo.Components{
// 							tuigo.Text("4. Review the configuration", tuigo.LabelText),
// 							tuigo.Text("Entity: ", tuigo.NormalText),
// 							tuigo.Text("... archs: "+archs, tuigo.NormalText),
// 							tuigo.Text("... output enabled: "+check(has_output), tuigo.NormalText),
// 							tuigo.Text("... MPI enabled: "+check(has_mpi), tuigo.NormalText),
// 							tuigo.Text("... debug mode: "+check(is_debug), tuigo.NormalText),
// 							tuigo.Text("... modulefile: "+nttpath, tuigo.NormalText),
// 						},
// 						external_modules...,
// 					),
// 					compiled_modules...,
// 				),
// 				new_modules...,
// 			)...,
// 		)
// 	} else if install_mode == "Kokkos" {
// 		return tuigo.Container(
// 			true,
// 			tuigo.VerticalContainer,
// 			tuigo.Text("4. Review the configuration for Kokkos", tuigo.LabelText),
// 			tuigo.Text("... src path: "+kokkos_src, tuigo.NormalText),
// 			tuigo.Text("... install path: "+kokkos_path, tuigo.NormalText),
// 			tuigo.Text("... archs: "+archs, tuigo.NormalText),
// 			tuigo.Text("... debug mode: "+check(is_debug), tuigo.NormalText),
// 		)
// 	} else if install_mode == "ADIOS2" {
// 		adios2_path := global.ADIOS2_INSTALL_DIR.Value.(string)
// 		return tuigo.Container(
// 			true,
// 			tuigo.VerticalContainer,
// 			tuigo.Text("4. Review the configuration for ADIOS2", tuigo.LabelText),
// 			tuigo.Text("... src path: "+adios2_src, tuigo.NormalText),
// 			tuigo.Text("... install path: "+adios2_path, tuigo.NormalText),
// 			tuigo.Text("... use MPI: "+check(has_mpi), tuigo.NormalText),
// 			tuigo.Text("... use CUDA: "+check(use_cuda), tuigo.NormalText),
// 			tuigo.Text("... Kokkos path: "+kokkos_path, tuigo.NormalText),
// 		)
// 	} else {
// 		panic("invalid mode selected")
// 	}
// }
