package pages

import (
	sel "github.com/entity-toolkit/ntt-dploy/selectors"
	"github.com/haykh/tuigo"
	"os"
	"slices"
	"strings"
)

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func readDepDir(dep, path string, w tuigo.Window) string {
	DEP := strings.ToUpper(dep)
	PATH := strings.ToUpper(path)
	sel.Selectors[DEP+"_"+PATH+"_DIR"].Read(w)
	return sel.Selectors[DEP+"_"+PATH+"_DIR"].Value().(string)
}

func getChoice(w tuigo.Window, label string) interface{} {
	sel.Selectors[label].Read(w)
	return sel.Selectors[label].Value()
}

func getCompilerModule(comp string, w tuigo.Window) string {
	comp_mod := getChoice(w, comp+"_MOD").(string)
	if comp_mod != "" {
		if strings.HasPrefix(comp_mod, "module:") {
			comp_mod = comp_mod[7:]
		}
	}
	return comp_mod
}

func Review(prev tuigo.Window) tuigo.Window {
	install_mode := getChoice(prev, "MODE").([]string)
	if len(install_mode) == 0 {
		panic("No install mode selected")
	}

	enable_modules := getChoice(prev, "ENV_MODULES").(bool)

	var buildscripts tuigo.ComplexContainerElement
	mode_entity := slices.Contains(install_mode, "Entity")
	mode_kokkos := slices.Contains(install_mode, "Kokkos")
	mode_adios2 := slices.Contains(install_mode, "ADIOS2")
	mode_mpi := slices.Contains(install_mode, "MPI")
	mode_hdf5 := slices.Contains(install_mode, "HDF5")
	deps := map[string]bool{
		"Kokkos": mode_kokkos,
		"ADIOS2": mode_adios2,
		"MPI":    mode_mpi,
		"HDF5":   mode_hdf5,
	}
	if !mode_entity {
		Configs["Entity"] = nil
	}
	if !mode_kokkos {
		Configs["Kokkos"] = nil
	}
	if !mode_adios2 {
		Configs["ADIOS2"] = nil
	}
	if !mode_mpi {
		Configs["MPI"] = nil
	}
	if !mode_hdf5 {
		Configs["HDF5"] = nil
	}

	dependencies := tuigo.Components{}

	with_debug := getChoice(prev, "DEBUG").(bool)
	with_cuda := getChoice(prev, "CUDA").(bool)
	with_mpi := getChoice(prev, "MPI").(bool)
	cpuarch := ""
	switch c := getChoice(prev, "CPUARCH").(type) {
	case string:
		cpuarch = c
	default:
		cpuarch = "NATIVE"
	}

	gpuarch := ""
	if with_cuda {
		switch g := getChoice(prev, "GPUARCH").(type) {
		case string:
			gpuarch = g
		default:
			gpuarch = "NATIVE"
		}
	}

	cxx_mod := getCompilerModule("CXX", prev)
	cuda_mod := ""
	if with_cuda {
		cuda_mod = getCompilerModule("CUDA", prev)
	}

	// sanity check
	if mode_mpi && !with_mpi {
		panic("MPI selected but not enabled")
	}

	if mode_kokkos || mode_adios2 || mode_mpi || mode_hdf5 {
		for dep, mode := range deps {
			if mode {
				Configs[dep].AddSetting("debug", with_debug)
				Configs[dep].AddSetting("mpi", with_mpi)
				Configs[dep].AddSetting("cpu", cpuarch)
				Configs[dep].AddSetting("gpu", gpuarch)
				Configs[dep].AddSetting("src_path", "")
				Configs[dep].AddSetting("install_path", "")
				Configs[dep].AddSetting("modules", []string{})
				Configs[dep].AddSetting("flags", []string{})
				Configs[dep].AddSetting("cflags", []string{})

				srcdir := readDepDir(dep, "src", prev)
				if srcdir == "" {
					srcdir = "$HOME/.modules/src/" + strings.ToLower(dep)
				}
				Configs[dep].AddSetting("src_path", srcdir)

				if Configs[dep].Check() {
					Configs[dep].GenSuffix()

					installdir := readDepDir(dep, "install", prev)
					def := "$HOME/.modules/"
					if installdir == "" || installdir == def+strings.ToLower(dep) {
						installdir = def + strings.ToLower(dep) + "/" + strings.ReplaceAll(Configs[dep].Suffix(), "/", ".")
					}
					installdir = strings.ReplaceAll(installdir, "$HOME", os.Getenv("HOME"))
					Configs[dep].AddSetting("install_path", installdir)
					dependencies = append(
						dependencies,
						tuigo.Text(dep+": "+srcdir+" -> "+installdir, tuigo.NormalText),
					)
					Configs[dep].InitFlags()
				} else {
					panic("Configuration for " + dep + " is incomplete")
				}
			}
		}
		// fill out the modules
		if enable_modules {
			// !TODO: add `export` as env variables if mod not specified
			if mode_kokkos {
				if cxx_mod != "" {
					Configs["Kokkos"].AddModule(cxx_mod)
				}
				if with_cuda {
					if cuda_mod != "" {
						Configs["Kokkos"].AddModule(cuda_mod)
					}
				}
			}
			if mode_adios2 {
				if cxx_mod != "" {
					Configs["ADIOS2"].AddModule(cxx_mod)
				}
				if mode_mpi {
					Configs["ADIOS2"].AddModule(Configs["MPI"].Suffix())
				} else if with_mpi {
					mpi_path := readDepDir("MPI", "install", prev)
					if strings.HasPrefix(mpi_path, "module:") {
						mpi_path = mpi_path[7:]
						Configs["ADIOS2"].AddModule(mpi_path)
					}
				}
				if mode_hdf5 {
					Configs["ADIOS2"].AddModule(Configs["HDF5"].Suffix())
				} else {
					hdf5_path := readDepDir("HDF5", "install", prev)
					if strings.HasPrefix(hdf5_path, "module:") {
						hdf5_path = hdf5_path[7:]
						Configs["ADIOS2"].AddModule(hdf5_path)
					}
				}
			}
			if mode_mpi {
				if cxx_mod != "" {
					Configs["MPI"].AddModule(cxx_mod)
				}
				if with_cuda {
					if cuda_mod != "" {
						Configs["MPI"].AddModule(cuda_mod)
					}
				}
			}
			if mode_hdf5 {
				if cxx_mod != "" {
					Configs["HDF5"].AddModule(cxx_mod)
				}
				if with_mpi {
					if mode_mpi {
						Configs["HDF5"].AddModule(Configs["MPI"].Suffix())
					} else {
						mpi_path := readDepDir("MPI", "install", prev)
						if strings.HasPrefix(mpi_path, "module:") {
							mpi_path = mpi_path[7:]
							Configs["HDF5"].AddModule(mpi_path)
						}
					}
				}
			}
		}
		buildscripts = tuigo.Container(
			tuigo.NonFocusable,
			tuigo.HorizontalContainer,
			tuigo.Text("Build scripts: ", tuigo.LabelText),
			tuigo.Container(
				tuigo.NonFocusable,
				tuigo.VerticalContainer,
				dependencies...,
			),
		)
	} else {
		buildscripts = tuigo.Container(
			tuigo.NonFocusable,
			tuigo.HorizontalContainer,
			tuigo.Text("No buildable libraries selected", tuigo.DimmedText),
		)
	}

	if mode_entity {
		Configs["Entity"].AddSetting("debug", with_debug)
		Configs["Entity"].AddSetting("mpi", with_mpi)

		Configs["Entity"].AddSetting("cpu", cpuarch)
		Configs["Entity"].AddSetting("gpu", gpuarch)

		if cxx_mod != "" {
			Configs["Entity"].AddModule(cxx_mod)
		}
		if cuda_mod != "" {
			Configs["Entity"].AddModule(cuda_mod)
		}

		if mode_kokkos {
			Configs["Entity"].AddModule(Configs["Kokkos"].Suffix())
		} else {
			kokkos_path := readDepDir("Kokkos", "install", prev)
			if strings.HasPrefix(kokkos_path, "module:") {
				kokkos_path = kokkos_path[7:]
				Configs["Entity"].AddModule(kokkos_path)
			}
		}

		if mode_adios2 {
			Configs["Entity"].AddModule(Configs["ADIOS2"].Suffix())
		} else {
			adios2_path := readDepDir("ADIOS2", "install", prev)
			if strings.HasPrefix(adios2_path, "module:") {
				adios2_path = adios2_path[7:]
				Configs["Entity"].AddModule(adios2_path)
			}
		}
		if mode_hdf5 {
			Configs["Entity"].AddModule(Configs["HDF5"].Suffix())
		} else {
			hdf5_path := readDepDir("HDF5", "install", prev)
			if strings.HasPrefix(hdf5_path, "module:") {
				hdf5_path = hdf5_path[7:]
				Configs["Entity"].AddModule(hdf5_path)
			}
		}

		if with_mpi {
			if mode_mpi {
				Configs["Entity"].AddModule(Configs["MPI"].Suffix())
			} else {
				mpi_path := readDepDir("MPI", "install", prev)
				if strings.HasPrefix(mpi_path, "module:") {
					mpi_path = mpi_path[7:]
					Configs["Entity"].AddModule(mpi_path)
				}
			}
		}
	}

	var modulefiles tuigo.ComplexContainerElement

	if enable_modules {
		modpath := getChoice(prev, "MODFILE_PATH").(string)
		if modpath == "" {
			modpath = "$HOME/.modules/"
		}
		ModulePath = modpath
		modules := []tuigo.Component{}
		for dep, mode := range deps {
			if mode {
				modules = append(modules, tuigo.Text(Configs[dep].Suffix(), tuigo.NormalText))
			}
		}
		modulefiles = tuigo.Container(
			tuigo.NonFocusable,
			tuigo.HorizontalContainer,
			tuigo.Container(
				tuigo.NonFocusable,
				tuigo.VerticalContainer,
				tuigo.Text("Module files: ", tuigo.LabelText),
				tuigo.Text("path: "+modpath, tuigo.NormalText),
			),
			tuigo.Container(
				tuigo.NonFocusable,
				tuigo.VerticalContainer,
				modules...,
			),
		)
	} else {
		modulefiles = tuigo.Container(
			tuigo.NonFocusable,
			tuigo.HorizontalContainer,
			tuigo.Text("Modulefiles deselected", tuigo.DimmedText),
		)
	}

	return tuigo.Container(
		tuigo.NonFocusable,
		tuigo.VerticalContainer,
		tuigo.Text("Review", tuigo.LabelText),
		tuigo.Text(Labels["review"], tuigo.NormalText),
		modulefiles,
		buildscripts,
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
