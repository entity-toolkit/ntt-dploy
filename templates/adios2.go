package templates

var ADIOS2Template = LibraryConfiguration{
	InitFlagsTemplate: func(self *LibraryConfiguration) {
		self.AddCFlag("CMAKE_CXX_STANDARD=17")
		self.AddCFlag("CMAKE_CXX_EXTENSIONS=OFF")
		self.AddCFlag("CMAKE_POSITION_INDEPENDENT_CODE=TRUE")
		self.AddCFlag("BUILD_SHARED_LIBS=ON")
		self.AddCFlag("ADIOS2_USE_HDF5=ON")
		self.AddCFlag("ADIOS2_USE_Kokkos=ON")
		self.AddCFlag("ADIOS2_USE_Python=OFF")
		self.AddCFlag("ADIOS2_USE_Fortran=OFF")
		self.AddCFlag("ADIOS2_USE_ZeroMQ=OFF")
		self.AddCFlag("BUILD_TESTING=OFF")
		self.AddCFlag("ADIOS2_BUILD_EXAMPLES=OFF")
		// if self.Settings["gpu"].(string) != "" {
		// 	self.AddCFlag("ADIOS2_USE_CUDA=ON")
		// 	self.AddFlag("ADIOS2_USE_CUDA ON")
		// }
		if self.Settings["mpi"].(bool) {
			self.AddCFlag("ADIOS2_USE_MPI=ON")
			self.AddFlag("ADIOS2_USE_MPI ON")
		} else {
			self.AddCFlag("ADIOS2_USE_MPI=OFF")
			self.AddFlag("ADIOS2_USE_MPI OFF")
			self.AddCFlag("ADIOS2_HAVE_HDF5_VOL=OFF")
			self.AddFlag("ADIOS2_HAVE_HDF5_VOL OFF")
		}
		self.AddCFlag("CMAKE_INSTALL_PREFIX=" + self.Settings["install_path"].(string))
	},
	CheckTemplate: func(self LibraryConfiguration) bool {
		required_settings := []string{
			"mpi",
			"cpu",
			"gpu",
			"modules",
			"install_path",
			"flags",
			"cflags",
			"src_path",
		}
		for _, setting := range required_settings {
			if _, ok := self.Settings[setting]; !ok {
				return false
			}
		}
		if _, ok := self.Settings["mpi"].(bool); !ok {
			return false
		}
		if _, ok := self.Settings["cpu"].(string); !ok {
			return false
		}
		if _, ok := self.Settings["gpu"].(string); !ok {
			return false
		}
		if _, ok := self.Settings["modules"].([]string); !ok {
			return false
		}
		if _, ok := self.Settings["install_path"].(string); !ok {
			return false
		}
		if _, ok := self.Settings["flags"].([]string); !ok {
			return false
		}
		if _, ok := self.Settings["cflags"].([]string); !ok {
			return false
		}
		if _, ok := self.Settings["src_path"].(string); !ok {
			return false
		}
		return true
	},
	SuffixTemplate: CreateTemplate("suffix", "adios2"+MPI_suffix+CPU_suffix),
	ModuleTemplate: CreateTemplate(
		"module",
		`#%Module1.0######################################################################
##
## {{settings .suffix}}
##
#################################################################################
proc ModulesHelp { } {
  puts stderr "\t{{settings .suffix}}\n"
}

module-whatis      "Sets up {{settings .suffix}}"

conflict           adios2
{{if not .modules}}{{else}}prereqs           {{range .modules}} {{.}}{{end}}{{end}}

set                basedir      {{.install_path}}
append-path        PATH         $basedir/bin
setenv             adios2_DIR   $basedir

{{range .flags}}
setenv             {{.}}{{end}}
`,
	),
	BuildScriptTemplate: CreateTemplate(
		"build",
		`module purge{{if not .modules}}{{else}}{{range .modules}}
module load {{.}}{{end}}{{end}}
cd {{.src_path}}
rm -rf build
cmake -B build{{range .cflags}} -D {{.}}{{end}}
cmake --build build -j
cmake --install build
rm -rf build`,
	),
}

// var ADIOS2Configs = map[string]interface{}{
// 	"Name":           "ADIOS2",
// 	"DefaultSrcPath": "$HOME/opt/src/adios2",
// 	"Suffix": func(mpi bool, cpu string, gpu string) string {
// 		suffix := "adios2"
// 		if mpi {
// 			suffix += "/mpi"
// 		} else {
// 			suffix += "/serial"
// 		}
// 		if cpu != "" {
// 			suffix += "/" + strings.ToLower(cpu)
// 		}
// 		if gpu != "" {
// 			suffix += "/" + strings.ToLower(gpu)
// 		}
// 		return suffix
// 	},
// 	"Condition": func(has_mpi, has_output bool) bool {
// 		return has_output
// 	},
// 	"UpdateMsg": OptADIOS2Upd{},
// 	"ModuleTemplate": func(
// 		sfx string,
// 		install_path string,
// 		is_debug bool,
// 		has_mpi bool,
// 		cpu_arch string,
// 		gpu_arch string,
// 		opts map[string]string) []string {
// 		modules := []string{}
// 		flags := []string{}
// 		if cpu_arch != "" {
// 			if cxx, ok := opts["CXX"]; ok {
// 				modules = append(modules, cxx)
// 			} else {
// 				panic("CXX module is not specified")
// 			}
// 		}
// 		if gpu_arch != "" {
// 			if cuda, ok := opts["CUDA"]; ok {
// 				flags = append(flags, "ADIOS2_USE_CUDA ON")
// 				modules = append(modules, cuda)
// 			} else {
// 				panic("GPU enabled but CUDA module is not specified")
// 			}
// 		}
// 		if kokkos, ok := opts["Kokkos"]; ok {
// 			modules = append(modules, kokkos)
// 		} else {
// 			panic("Kokkos module is not specified")
// 		}
// 		if has_mpi {
// 			if mpi, ok := opts["MPI"]; ok {
// 				flags = append(flags, "ADIOS2_USE_MPI ON")
// 				modules = append(modules, mpi)
// 			} else {
// 				panic("MPI enabled but MPI module is not specified")
// 			}
// 		} else {
// 			flags = append(flags, "ADIOS2_USE_MPI OFF")
// 			flags = append(flags, "ADIOS2_HAVE_HDF5_VOL OFF")
// 		}
//
// 		modulefile := []string{}
// 		modulefile = append(modulefile, "#%Module1.0######################################################################")
// 		modulefile = append(modulefile, "##")
// 		conf := strings.ToUpper(strings.ReplaceAll(sfx, "/", " @ "))
// 		modulefile = append(modulefile, "## "+conf)
// 		modulefile = append(modulefile, "##")
// 		modulefile = append(modulefile, "################################################################################")
// 		modulefile = append(modulefile, "proc ModulesHelp { } {")
// 		modulefile = append(modulefile, "	 puts stderr \"\\t"+conf+"\\n\"")
// 		modulefile = append(modulefile, "}")
// 		modulefile = append(modulefile, "")
// 		modulefile = append(modulefile, "module-whatis      \"Sets up "+conf+"\"")
// 		modulefile = append(modulefile, "")
// 		modulefile = append(modulefile, "conflict           adios2")
// 		modulefile = append(modulefile, "")
// 		modulefile = append(modulefile, "prereqs						 "+strings.Join(modules, " "))
// 		modulefile = append(modulefile, "")
// 		modulefile = append(modulefile, "set                basedir      "+install_path)
// 		modulefile = append(modulefile, "append-path        PATH         $basedir/bin")
// 		modulefile = append(modulefile, "setenv             adios2_DIR   $basedir")
// 		modulefile = append(modulefile, "")
// 		for _, flag := range flags {
// 			modulefile = append(modulefile, "setenv             "+flag)
// 		}
// 		return modulefile
// 	},
// 	"BuildScript": func(
// 		src_path, install_path string,
// 		is_debug, has_mpi bool,
// 		cpu_arch, gpu_arch string,
// 		use_modules bool,
// 		opts map[string]string,
// 	) ([]string, []string, []string, []string) {
// 		modules := []string{}
// 		if use_modules {
// 			if cpu_arch != "" {
// 				if cxx, ok := opts["CXX"]; ok {
// 					modules = append(modules, cxx)
// 				} else {
// 					panic("CXX module is not specified")
// 				}
// 			}
// 			if gpu_arch != "" {
// 				if cuda, ok := opts["CUDA"]; ok {
// 					modules = append(modules, cuda)
// 				} else {
// 					panic("GPU enabled but CUDA module is not specified")
// 				}
// 			}
// 			if kokkos, ok := opts["Kokkos"]; ok {
// 				modules = append(modules, kokkos)
// 			} else {
// 				panic("Kokkos module is not specified")
// 			}
// 		}
// 		flags := []string{
// 			"CMAKE_CXX_STANDARD=17",
// 			"CMAKE_CXX_EXTENSIONS=OFF",
// 			"CMAKE_POSITION_INDEPENDENT_CODE=TRUE",
// 			"BUILD_SHARED_LIBS=ON",
// 			"ADIOS2_USE_HDF5=ON",
// 			"ADIOS2_USE_Kokkos=ON",
// 			"ADIOS2_USE_Python=OFF",
// 			"ADIOS2_USE_Fortran=OFF",
// 			"ADIOS2_USE_ZeroMQ=OFF",
// 			"BUILD_TESTING=OFF",
// 			"ADIOS2_BUILD_EXAMPLES=OFF",
// 			"CMAKE_INSTALL_PREFIX=" + install_path,
// 		}
// 		if gpu_arch != "" {
// 			flags = append(flags, "ADIOS2_USE_CUDA=ON")
// 		}
// 		if has_mpi {
// 			flags = append(flags, "ADIOS2_USE_MPI=ON")
// 		} else {
// 			flags = append(flags, "ADIOS2_USE_MPI=OFF")
// 			flags = append(flags, "ADIOS2_HAVE_HDF5_VOL=OFF")
// 		}
// 		var prebuild []string
// 		if use_modules {
// 			prebuild = []string{
// 				"module purge",
// 				"module load " + strings.Join(modules, " "),
// 			}
// 		} else {
// 			prebuild = []string{}
// 		}
// 		configure := []string{
// 			"cd " + src_path,
// 			"rm -rf build",
// 			"cmake -B build -D " + strings.Join(flags, " -D "),
// 		}
// 		build := []string{
// 			"cmake --build build -j",
// 			"cmake --install build",
// 		}
// 		postbuild := []string{
// 			"cd " + src_path,
// 			"rm -rf build",
// 		}
// 		return prebuild, configure, build, postbuild
// 	},
// }
