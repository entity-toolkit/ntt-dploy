package pages

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/entity-toolkit/ntt-dploy/configs"
	"github.com/haykh/tuigo"
)

func DependenciesSelector(prev tuigo.Window) tuigo.Window {
	// install_mode := configs.Global["has_mpi"].(string)
	install_mode := global.MODE.Value.(string)
	if install_mode == "Entity" {
		has_mpi := prev.GetElementByID(
			global.MPI.Id,
		).(tuigo.RadioElement).Data().(bool)
		has_output := prev.GetElementByID(
			global.OUTPUT.Id,
		).(tuigo.RadioElement).Data().(bool)
		is_debug = prev.GetElementByID(
			global.DEBUG.Id,
		).(tuigo.RadioElement).Data().(bool)
		cpu := prev.GetElementByID(
			global.CPUARCH.Id,
		).(tuigo.SelectorElement).Data()
		gpu := prev.GetElementByID(
			global.GPUARCH.Id,
		).(tuigo.SelectorElement).Data()
		var cpu_str = "NATIVE"
		if cpu != nil {
			cpu_str = cpu.(string)
		}
		archs += cpu_str
		if gpu != nil {
			archs += ", " + gpu.(string)
		}
		cpu_arch, gpu_arch = configs.GetArchs(archs)

		cfg := configs.EntityConfigs
		entsfx := cfg["Suffix"].(configs.EntitySuffixFunc)(is_debug, has_mpi, cpu_arch, gpu_arch)

		elements := tuigo.Components{}
		elements = append(elements,
			tuigo.InputWithID(
				configs.HashID["MODULE_DIR"],
				cfg["Name"].(string)+" module parent directory",
				configs.DefaultModPath, "",
				tuigo.PathInput,
				configs.NO_CALLBACK,
			),
		)
		elements = append(elements,
			tuigo.Container(false, tuigo.HorizontalContainer,
				tuigo.TextWithID(
					configs.HashID["MODULE_PATH"],
					"  "+cfg["Name"].(string)+" module path: ",
					tuigo.NormalText,
				),
				tuigo.TextWithID(
					configs.HashID["NTT_PATH"],
					".../"+entsfx,
					tuigo.NormalText,
				),
			),
		)
		elements = append(elements,
			tuigo.TextWithID(
				configs.HashID["ARCHITECTURES"],
				"  architectures: "+archs,
				tuigo.NormalText,
			),
		)
		elements = append(elements,
			tuigo.Text("Dependencies:", tuigo.NormalText),
		)
		elements = append(elements,
			tuigo.InputWithID(
				configs.HashID["CXX_MODULE"],
				"  CXX module",
				configs.DefaultCXX, "",
				tuigo.TextInput,
				configs.NO_CALLBACK,
			),
		)
		if gpu != nil {
			elements = append(elements,
				tuigo.InputWithID(
					configs.HashID["CUDA_MODULE"],
					"  CUDA module",
					configs.DefaultCUDA, "",
					tuigo.TextInput,
					configs.NO_CALLBACK,
				),
			)
		}
		inputs := tuigo.Components{}
		paths_modules := tuigo.Components{
			tuigo.TextWithID(
				configs.HashID["INSTALL_MODULE_PATH"],
				"Install & module paths:",
				tuigo.NormalText,
			),
			tuigo.InputWithID(
				configs.HashID["PARENT_DIR"],
				"  parent directory",
				configs.DefaultOptPath, "",
				tuigo.PathInput,
				configs.NO_CALLBACK,
			),
		}
		for id := range configs.DependencyMapping {
			conf := configs.DependencyMapping[id]
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
					tuigo.TextWithID(id+50, "  "+name+" install path: .../"+sfx, tuigo.NormalText),
					tuigo.TextWithID(id+150, "  "+name+" module: .../.modules/"+sfx, tuigo.NormalText),
				)
			}
		}
		elements = append(elements, inputs...)
		elements = append(elements, paths_modules...)
		elements = append(
			tuigo.Components{tuigo.Text("3. Specify paths, compilers & modules to use", tuigo.LabelText)},
			elements...,
		)
		elements = append(
			elements,
			tuigo.Text("[?] to use preinstalled dependencies using modules, use `module:<MODULENAME>`", tuigo.DimmedText),
		)
		return tuigo.Container(
			configs.FOCUSABLE,
			tuigo.VerticalContainer,
			elements...,
		)
	} else if install_mode == "Kokkos" {
		is_debug = prev.GetElementByID(configs.HashID["ENABLE_DEBUG"]).(tuigo.RadioElement).Data().(bool)
		cpu := prev.GetElementByID(configs.HashID["CPU_SELECTOR"]).(tuigo.SelectorElement).Data()
		gpu := prev.GetElementByID(configs.HashID["GPU_SELECTOR"]).(tuigo.SelectorElement).Data()
		var cpu_str = "NATIVE"
		if cpu != nil {
			cpu_str = cpu.(string)
		}
		archs += cpu_str
		if gpu != nil {
			archs += ", " + gpu.(string)
		}
		cpu_arch, gpu_arch = configs.GetArchs(archs)
		kokkos_src = prev.GetElementByID(configs.HashID["KOKKOS_SRC"]).(tuigo.InputElement).Data().(string)
		kokkos_install_path = prev.GetElementByID(configs.HashID["KOKKOS_INSTALL"]).(tuigo.InputElement).Data().(string)
		return tuigo.Container(
			true,
			tuigo.VerticalContainer,
			tuigo.Text("3. No additional configuration necessary (skip this step)", tuigo.LabelText),
		)
	} else if install_mode == "ADIOS2" {
		has_mpi = prev.GetElementByID(configs.HashID["USE_MPI"]).(tuigo.RadioElement).Data().(bool)
		kokkos_path = prev.GetElementByID(configs.HashID["KOKKOS_PATH"]).(tuigo.InputElement).Data().(string)
		use_cuda = prev.GetElementByID(configs.HashID["USE_CUDA"]).(tuigo.RadioElement).Data().(bool)
		adios2_src = prev.GetElementByID(configs.HashID["ADIOS2_SRC"]).(tuigo.InputElement).Data().(string)
		adios2_install_path = prev.GetElementByID(configs.HashID["ADIOS2_INSTALL"]).(tuigo.InputElement).Data().(string)
		return tuigo.Container(
			true,
			tuigo.VerticalContainer,
			tuigo.Text("3. No additional configuration necessary (skip this step)", tuigo.LabelText),
		)
	} else {
		panic("invalid mode selected")
	}
}
