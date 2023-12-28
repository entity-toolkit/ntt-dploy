package configs

import "strings"

var KokkosConfigs = map[string]interface{}{
	"Name":           "Kokkos",
	"DefaultSrcPath": "$HOME/opt/src/kokkos",
	"Suffix": func(mpi bool, cpu string, gpu string) string {
		suffix := "kokkos"
		if cpu != "" {
			suffix += "/" + strings.ToLower(cpu)
		}
		if gpu != "" {
			suffix += "/" + strings.ToLower(gpu)
		}
		return suffix
	},
	"Condition": func(has_mpi, has_output bool) bool {
		return true
	},
	"UpdateMsg": OptKokkosUpd{},
	"ModuleTemplate": func(
		sfx string,
		install_path string,
		is_debug bool,
		has_mpi bool,
		cpu_arch string,
		gpu_arch string,
		opts map[string]string) []string {
		modules := []string{}
		flags := []string{}
		if cpu_arch != "" {
			flags = append(flags, "Kokkos_ARCH_"+strings.ToUpper(cpu_arch)+" ON")
			if cxx, ok := opts["CXX"]; ok {
				modules = append(modules, cxx)
			} else {
				panic("CXX module is not specified")
			}
		}
		if gpu_arch != "" {
			flags = append(flags, "Kokkos_ARCH_"+strings.ToUpper(gpu_arch)+" ON")
			flags = append(flags, "Kokkos_ENABLE_CUDA ON")
			if cuda, ok := opts["CUDA"]; ok {
				modules = append(modules, cuda)
			} else {
				panic("GPU enabled but CUDA module is not specified")
			}
		}
		if !has_mpi {
			flags = append(flags, "Kokkos_ENABLE_OPENMP ON")
		} else {
			if mpi, ok := opts["MPI"]; ok {
				modules = append(modules, mpi)
			} else {
				panic("MPI enabled but MPI module is not specified")
			}
		}

		modulefile := []string{}
		modulefile = append(modulefile, "#%Module1.0######################################################################")
		modulefile = append(modulefile, "##")
		conf := strings.ToUpper(strings.ReplaceAll(sfx, "/", " @ "))
		modulefile = append(modulefile, "## "+conf)
		modulefile = append(modulefile, "##")
		modulefile = append(modulefile, "################################################################################")
		modulefile = append(modulefile, "proc ModulesHelp { } {")
		modulefile = append(modulefile, "	 puts stderr \"\\t"+conf+"\\n\"")
		modulefile = append(modulefile, "}")
		modulefile = append(modulefile, "")
		modulefile = append(modulefile, "module-whatis      \"Sets up "+conf+"\"")
		modulefile = append(modulefile, "")
		modulefile = append(modulefile, "conflict           kokkos")
		modulefile = append(modulefile, "")
		modulefile = append(modulefile, "prereqs						 "+strings.Join(modules, " "))
		modulefile = append(modulefile, "")
		modulefile = append(modulefile, "set                basedir      "+install_path)
		modulefile = append(modulefile, "append-path        PATH         $basedir/bin")
		modulefile = append(modulefile, "setenv             Kokkos_DIR   $basedir")
		modulefile = append(modulefile, "")
		for _, flag := range flags {
			modulefile = append(modulefile, "setenv             "+flag)
		}
		return modulefile
	},
	"BuildScript": func(
		src_path, install_path string,
		is_debug, has_mpi bool,
		cpu_arch, gpu_arch string,
		opts map[string]string,
	) ([]string, []string, []string, []string) {
		modules := []string{}
		flags := []string{
			"CMAKE_CXX_STANDARD=17",
			"CMAKE_CXX_EXTENSIONS=OFF",
			"CMAKE_POSITION_INDEPENDENT_CODE=TRUE",
			"CMAKE_INSTALL_PREFIX=" + install_path,
		}
		if cpu_arch != "" {
			flags = append(flags, "Kokkos_ARCH_"+strings.ToUpper(cpu_arch)+"=ON")
			if cxx, ok := opts["CXX"]; ok {
				modules = append(modules, cxx)
			} else {
				panic("CXX module is not specified")
			}
		}
		if gpu_arch != "" {
			flags = append(flags, "Kokkos_ARCH_"+strings.ToUpper(gpu_arch)+"=ON")
			flags = append(flags, "Kokkos_ENABLE_CUDA=ON")
			if cuda, ok := opts["CUDA"]; ok {
				modules = append(modules, cuda)
			} else {
				panic("GPU enabled but CUDA module is not specified")
			}
		}
		if !has_mpi {
			flags = append(flags, "Kokkos_ENABLE_OPENMP=ON")
		} else {
			if mpi, ok := opts["MPI"]; ok {
				modules = append(modules, mpi)
			} else {
				panic("MPI enabled but MPI module is not specified")
			}
		}
		prebuild := []string{
			"module purge",
			"module load " + strings.Join(modules, " "),
		}
		configure := []string{
			"cd " + src_path,
			"rm -rf build",
			"cmake -B build -D " + strings.Join(flags, " -D "),
		}
		build := []string{
			"cmake --build build -j",
			"cmake --install build",
		}
		postbuild := []string{
			"cd " + src_path,
			"rm -rf build",
		}
		return prebuild, configure, build, postbuild
	},
}
