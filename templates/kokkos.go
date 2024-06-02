package templates

import "strings"

var KokkosTemplate = LibraryConfiguration{
	InitFlagsTemplate: func(self *LibraryConfiguration) {
		self.AddCFlag("CMAKE_CXX_STANDARD=17")
		self.AddCFlag("CMAKE_CXX_EXTENSIONS=OFF")
		self.AddCFlag("CMAKE_POSITION_INDEPENDENT_CODE=TRUE")
		if self.Settings["cpu"].(string) != "" {
			self.AddCFlag("Kokkos_ARCH_" + strings.ToUpper(self.Settings["cpu"].(string)) + "=ON")
			self.AddFlag("Kokkos_ARCH_" + strings.ToUpper(self.Settings["cpu"].(string)) + " ON")
		}
		if self.Settings["gpu"].(string) != "" {
			self.AddCFlag("Kokkos_ARCH_" + strings.ToUpper(self.Settings["gpu"].(string)) + "=ON")
			self.AddFlag("Kokkos_ARCH_" + strings.ToUpper(self.Settings["gpu"].(string)) + " ON")
			self.AddCFlag("Kokkos_ENABLE_CUDA=ON")
			self.AddFlag("Kokkos_ENABLE_CUDA ON")
		} else {
			if !self.Settings["mpi"].(bool) {
				self.AddCFlag("Kokkos_ENABLE_OPENMP=ON")
				self.AddFlag("Kokkos_ENABLE_OPENMP ON")
				self.AddFlag("OMP_PROC_BIND spread")
				self.AddFlag("OMP_PLACES threads")
				self.AddFlag("OMP_NUM_THREADS [exec nproc]")
			}
		}
		self.AddCFlag("CMAKE_INSTALL_PREFIX=" + self.Settings["install_path"].(string))
	},
	CheckTemplate: func(self LibraryConfiguration) bool {
		required_settings := []string{
			"debug",
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
		if _, ok := self.Settings["debug"].(bool); !ok {
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
	SuffixTemplate: CreateTemplate("suffix", "kokkos"+DEBUG_suffix+GPU_suffix+CPU_suffix),
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

conflict           kokkos
{{if not .modules}}{{else}}prereq            {{range .modules}} {{.}}{{end}}{{end}}

set                basedir      {{.install_path}}
prepend-path       PATH         $basedir/bin
setenv             Kokkos_DIR   $basedir

{{range .flags}}
setenv             {{.}}{{end}}
`,
	),
	BuildScriptTemplate: CreateTemplate(
		"build",
		`CURRENT_DIR=$(pwd) &&\
module purge{{if not .modules}}{{else}}{{range .modules}} &&\
module load {{.}}{{end}}{{end}} &&\
cd {{.src_path}} &&\
rm -rf build &&\
cmake -B build{{range .cflags}} -D {{.}}{{end}} &&\
cmake --build build -j &&\
cmake --install build &&\
rm -rf build &&\
cd $CURRENT_DIR &&\
unset CURRENT_DIR`,
	),
}

//  func(
// 	src_path, install_path string,
// 	is_debug, has_mpi bool,
// 	cpu_arch, gpu_arch string,
// 	use_modules bool,
// 	opts map[string]string,
// ) ([]string, []string, []string, []string) {
// 	modules := []string{}
// 	flags := []string{
// 		"CMAKE_CXX_STANDARD=17",
// 		"CMAKE_CXX_EXTENSIONS=OFF",
// 		"CMAKE_POSITION_INDEPENDENT_CODE=TRUE",
// 		"CMAKE_INSTALL_PREFIX=" + install_path,
// 	}
// 	if cpu_arch != "" {
// 		flags = append(flags, "Kokkos_ARCH_"+strings.ToUpper(cpu_arch)+"=ON")
// 		if use_modules {
// 			if cxx, ok := opts["CXX"]; ok {
// 				modules = append(modules, cxx)
// 			} else {
// 				panic("CXX module is not specified")
// 			}
// 		}
// 	}
// 	if gpu_arch != "" {
// 		flags = append(flags, "Kokkos_ARCH_"+strings.ToUpper(gpu_arch)+"=ON")
// 		flags = append(flags, "Kokkos_ENABLE_CUDA=ON")
// 		if use_modules {
// 			if cuda, ok := opts["CUDA"]; ok {
// 				modules = append(modules, cuda)
// 			} else {
// 				panic("GPU enabled but CUDA module is not specified")
// 			}
// 		}
// 	}
// 	if !has_mpi {
// 		flags = append(flags, "Kokkos_ENABLE_OPENMP=ON")
// 	} else {
// 		if use_modules {
// 			if mpi, ok := opts["MPI"]; ok {
// 				modules = append(modules, mpi)
// 			} else {
// 				panic("MPI enabled but MPI module is not specified")
// 			}
// 		}
// 	}
// 	var prebuild []string
// 	if use_modules {
// 		prebuild = []string{
// 			"module purge",
// 			"module load " + strings.Join(modules, " "),
// 		}
// 	} else {
// 		prebuild = []string{}
// 	}
// 	configure := []string{
// 		"cd " + src_path,
// 		"rm -rf build",
// 		"cmake -B build -D " + strings.Join(flags, " -D "),
// 	}
// 	build := []string{
// 		"cmake --build build -j",
// 		"cmake --install build",
// 	}
// 	postbuild := []string{
// 		"cd " + src_path,
// 		"rm -rf build",
// 	}
// 	return prebuild, configure, build, postbuild
// },
// }
