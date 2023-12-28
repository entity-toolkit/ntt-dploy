package configs

import "strings"

var ADIOS2Configs = map[string]interface{}{
	"Name":           "ADIOS2",
	"DefaultSrcPath": "$HOME/opt/src/adios2",
	"Suffix": func(mpi bool, cpu string, gpu string) string {
		suffix := "adios2"
		if mpi {
			suffix += "/mpi"
		} else {
			suffix += "/serial"
		}
		if cpu != "" {
			suffix += "/" + strings.ToLower(cpu)
		}
		if gpu != "" {
			suffix += "/" + strings.ToLower(gpu)
		}
		return suffix
	},
	"Condition": func(has_mpi, has_output bool) bool {
		return has_output
	},
	"UpdateMsg": OptADIOS2Upd{},
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
			if cxx, ok := opts["CXX"]; ok {
				modules = append(modules, cxx)
			} else {
				panic("CXX module is not specified")
			}
		}
		if gpu_arch != "" {
			if cuda, ok := opts["CUDA"]; ok {
				flags = append(flags, "ADIOS2_USE_CUDA ON")
				modules = append(modules, cuda)
			} else {
				panic("GPU enabled but CUDA module is not specified")
			}
		}
		if kokkos, ok := opts["Kokkos"]; ok {
			modules = append(modules, kokkos)
		} else {
			panic("Kokkos module is not specified")
		}
		if has_mpi {
			if mpi, ok := opts["MPI"]; ok {
				flags = append(flags, "ADIOS2_USE_MPI ON")
				modules = append(modules, mpi)
			} else {
				panic("MPI enabled but MPI module is not specified")
			}
		} else {
			flags = append(flags, "ADIOS2_USE_MPI OFF")
			flags = append(flags, "ADIOS2_HAVE_HDF5_VOL OFF")
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
		modulefile = append(modulefile, "conflict           adios2")
		modulefile = append(modulefile, "")
		modulefile = append(modulefile, "prereqs						 "+strings.Join(modules, " "))
		modulefile = append(modulefile, "")
		modulefile = append(modulefile, "set                basedir      "+install_path)
		modulefile = append(modulefile, "append-path        PATH         $basedir/bin")
		modulefile = append(modulefile, "setenv             adios2_DIR   $basedir")
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
		if cpu_arch != "" {
			if cxx, ok := opts["CXX"]; ok {
				modules = append(modules, cxx)
			} else {
				panic("CXX module is not specified")
			}
		}
		if gpu_arch != "" {
			if cuda, ok := opts["CUDA"]; ok {
				modules = append(modules, cuda)
			} else {
				panic("GPU enabled but CUDA module is not specified")
			}
		}
		if kokkos, ok := opts["Kokkos"]; ok {
			modules = append(modules, kokkos)
		} else {
			panic("Kokkos module is not specified")
		}
		flags := []string{
			"CMAKE_CXX_STANDARD=17",
			"CMAKE_CXX_EXTENSIONS=OFF",
			"CMAKE_POSITION_INDEPENDENT_CODE=TRUE",
			"BUILD_SHARED_LIBS=ON",
			"ADIOS2_USE_HDF5=ON",
			"ADIOS2_USE_Kokkos=ON",
			"ADIOS2_USE_Python=OFF",
			"ADIOS2_USE_Fortran=OFF",
			"ADIOS2_USE_ZeroMQ=OFF",
			"BUILD_TESTING=OFF",
			"ADIOS2_BUILD_EXAMPLES=OFF",
			"CMAKE_INSTALL_PREFIX=" + install_path,
		}
		if gpu_arch != "" {
			flags = append(flags, "ADIOS2_USE_CUDA=ON")
		}
		if has_mpi {
			flags = append(flags, "ADIOS2_USE_MPI=ON")
		} else {
			flags = append(flags, "ADIOS2_USE_MPI=OFF")
			flags = append(flags, "ADIOS2_HAVE_HDF5_VOL=OFF")
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
