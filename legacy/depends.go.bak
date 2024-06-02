package pages

import (
	"github.com/entity-toolkit/ntt-dploy/configs"
	"github.com/entity-toolkit/ntt-dploy/global"
	"github.com/haykh/tuigo"
)

func DependenciesSelector(prev tuigo.Window) tuigo.Window {
	install_mode := global.Selectors["MODE"].Value().(string)

	if install_mode == "Entity" || install_mode == "Kokkos" {
		global.Selectors["DEBUG"].Read(prev)
	}

	if install_mode == "Entity" {
		global.Selectors["OUTPUT"].Read(prev)
	}

	if install_mode == "Kokkos" {
		global.Selectors["CPUARCH"].Read(prev)
		global.Selectors["GPUARCH"].Read(prev)
	}

	if install_mode == "Entity" || install_mode == "ADIOS2" {
		global.Selectors["MPI"].Read(prev)
	}
	global.Selectors["CUDA"].Read(prev)

	return tuigo.Container(
		configs.FOCUSABLE,
		tuigo.VerticalContainer,
		tuigo.Text("3. No additional configuration necessary (skip this step)", tuigo.LabelText),
	)

	// if install_mode == "Entity" {
	// entsfx := configs.EntityConfigs["Suffix"].(configs.EntitySuffixFunc)(
	// 	global.Selectors["DEBUG"].Value.(bool),
	// 	global.MPI.Value.(bool),
	// 	global.CPUARCH.Value.(string),
	// 	global.GPUARCH.Value.(string),
	// )

	// 		elements := tuigo.Components{}
	// 		elements = append(elements,
	// 			tuigo.InputWithID(
	// 				global.MODULES_DIR.Id,
	// 				cfg["Name"].(string)+" "+global.MODULES_DIR.Text,
	// 				configs.DefaultModPath, "",
	// 				tuigo.PathInput,
	// 				configs.NO_CALLBACK,
	// 			),
	// 		)
	// 		elements = append(elements,
	// 			tuigo.Container(false, tuigo.HorizontalContainer,
	// 				tuigo.TextWithID(
	// 					global.MODULE.Id,
	// 					"  "+cfg["Name"].(string)+" "+global.MODULE.Text+": ",
	// 					tuigo.NormalText,
	// 				),
	// 				tuigo.TextWithID(
	// 					// configs.HashID["NTT_PATH"],
	// 					global.ENTITY_PATH.Id,
	// 					".../"+entsfx,
	// 					tuigo.NormalText,
	// 				),
	// 			),
	// 		)
	// 		elements = append(elements,
	// 			tuigo.TextWithID(
	// 				global.ARCHS.Id,
	// 				"  "+global.ARCHS.Text+": "+global.ARCHS.Value.(string),
	// 				tuigo.NormalText,
	// 			),
	// 		)
	// 		elements = append(elements,
	// 			tuigo.Text("Dependencies:", tuigo.NormalText),
	// 		)
	// 		elements = append(elements,
	// 			tuigo.InputWithID(
	// 				global.CXX_MODULE.Id,
	// 				"  "+global.CXX_MODULE.Text,
	// 				configs.DefaultCXX, "",
	// 				tuigo.TextInput,
	// 				configs.NO_CALLBACK,
	// 			),
	// 		)
	// 		if _, ok := global.GPUARCH.Value.(string); ok {
	// 			elements = append(elements,
	// 				tuigo.InputWithID(
	// 					global.CUDA_MODULE.Id,
	// 					"  "+global.CUDA_MODULE.Text,
	// 					configs.DefaultCUDA, "",
	// 					tuigo.TextInput,
	// 					configs.NO_CALLBACK,
	// 				),
	// 			)
	// 		}
	// 		inputs := tuigo.Components{}
	// 		paths_modules := tuigo.Components{
	// 			tuigo.TextWithID(
	// 				global.INSTALL_MODULE_PATHS.Id,
	// 				global.INSTALL_MODULE_PATHS.Text+":",
	// 				tuigo.NormalText,
	// 			),
	// 			tuigo.InputWithID(
	// 				global.PARENT_DIR.Id,
	// 				"  "+global.PARENT_DIR.Text,
	// 				configs.DefaultOptPath, "",
	// 				tuigo.PathInput,
	// 				configs.NO_CALLBACK,
	// 			),
	// 		}
	// 		for id := range configs.DependencyMapping {
	// 			conf := configs.DependencyMapping[id]
	// 			if conf["Condition"].(configs.ConditionFunc)(
	// 				global.MPI.Value.(bool),
	// 				global.OUTPUT.Value.(bool),
	// 			) {
	// 				name := conf["Name"].(string)
	// 				inputs = append(inputs,
	// 					tuigo.InputWithID(
	// 						id,
	// 						"  "+name+" src/module",
	// 						conf["DefaultSrcPath"].(string),
	// 						"",
	// 						tuigo.PathInput,
	// 						conf["UpdateMsg"].(tea.Msg),
	// 					),
	// 				)
	// 				sfx := conf["Suffix"].(configs.SuffixFunc)(
	// 					global.MPI.Value.(bool),
	// 					global.CPUARCH.Value.(string),
	// 					global.GPUARCH.Value.(string),
	// 				)
	// 				paths_modules = append(paths_modules,
	// 					tuigo.TextWithID(id+50, "  "+name+" install path: .../"+sfx, tuigo.NormalText),
	// 					tuigo.TextWithID(id+150, "  "+name+" module: .../.modules/"+sfx, tuigo.NormalText),
	// 				)
	// 			}
	// 		}
	// 		elements = append(elements, inputs...)
	// 		elements = append(elements, paths_modules...)
	// 		elements = append(
	// 			tuigo.Components{tuigo.Text("3. Specify paths, compilers & modules to use", tuigo.LabelText)},
	// 			elements...,
	// 		)
	// 		elements = append(
	// 			elements,
	// 			tuigo.Text("[?] to use preinstalled dependencies using modules, use `module:<MODULENAME>`", tuigo.DimmedText),
	// 		)
	// 		return tuigo.Container(
	// 			configs.FOCUSABLE,
	// 			tuigo.VerticalContainer,
	// 			elements...,
	// 		)
	// } else if install_mode == "Kokkos" {
	// } else if install_mode == "ADIOS2" {
	// } else {
	// 	panic("invalid mode selected")
	// }
}

// func DependenciesSelector(prev tuigo.Window) tuigo.Window {
// 	install_mode := global.MODE.Value.(string)
// 	global.MPI.Value = prev.GetElementByID(
// 		global.MPI.Id,
// 	).(tuigo.RadioElement).Data().(bool)

// 	if install_mode == "Entity" || install_mode == "Kokkos" {
// 		global.OUTPUT.Value = prev.GetElementByID(
// 			global.OUTPUT.Id,
// 		).(tuigo.RadioElement).Data().(bool)
// 		global.DEBUG.Value = prev.GetElementByID(
// 			global.DEBUG.Id,
// 		).(tuigo.RadioElement).Data().(bool)

// 		cpu := prev.GetElementByID(
// 			global.CPUARCH.Id,
// 		).(tuigo.SelectorElement).Data()
// 		gpu := prev.GetElementByID(
// 			global.GPUARCH.Id,
// 		).(tuigo.SelectorElement).Data()

// 		var cpu_str = "NATIVE"
// 		if cpu != nil {
// 			cpu_str = cpu.(string)
// 		}
// 		archs := ""
// 		archs += cpu_str
// 		if gpu != nil {
// 			archs += ", " + gpu.(string)
// 		}
// 		cpu_arch, gpu_arch := configs.GetArchs(archs)
// 		global.ARCHS.Value = archs
// 		global.CPUARCH.Value = cpu_arch
// 		global.GPUARCH.Value = gpu_arch
// 	} else if install_mode == "ADIOS2" {
// 		global.KOKKOS_INSTALL_DIR.Value = prev.GetElementByID(
// 			global.KOKKOS_INSTALL_DIR.Id,
// 		).(tuigo.InputElement).Data().(string)
// 		global.CUDA.Value = prev.GetElementByID(
// 			global.CUDA.Id,
// 		).(tuigo.RadioElement).Data().(bool)
// 		global.ADIOS2_SRC_DIR.Value = prev.GetElementByID(
// 			global.ADIOS2_SRC_DIR.Id,
// 		).(tuigo.InputElement).Data().(string)
// 		global.ADIOS2_INSTALL_DIR.Value = prev.GetElementByID(
// 			global.ADIOS2_INSTALL_DIR.Id,
// 		).(tuigo.InputElement).Data().(string)
// 	}

// 	if install_mode == "Entity" {
// 	} else if install_mode == "Kokkos" {
// 		global.KOKKOS_SRC_DIR.Value = prev.GetElementByID(global.KOKKOS_SRC_DIR.Id).(tuigo.InputElement).Data().(string)
// 		global.KOKKOS_INSTALL_DIR.Value = prev.GetElementByID(global.KOKKOS_INSTALL_DIR.Id).(tuigo.InputElement).Data().(string)
// 		return tuigo.Container(
// 			true,
// 			tuigo.VerticalContainer,
// 			tuigo.Text("3. No additional configuration necessary (skip this step)", tuigo.LabelText),
// 		)
// 	} else if install_mode == "ADIOS2" {
// 		return tuigo.Container(
// 			true,
// 			tuigo.VerticalContainer,
// 			tuigo.Text("3. No additional configuration necessary (skip this step)", tuigo.LabelText),
// 		)
// 	} else {
// 		panic("invalid mode selected")
// 	}
// }
