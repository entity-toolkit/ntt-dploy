package configs

import (
	"fmt"
	"strings"
)

var HDF5Configs = map[string]interface{}{
	"Name":           "HDF5",
	"DefaultSrcPath": "$HOME/opt/src/hdf5",
	"Suffix": func(mpi bool, cpu string, gpu string) string {
		suffix := "hdf5"
		if mpi {
			suffix += "/mpi"
		} else {
			suffix += "/serial"
		}
		return suffix
	},
	"Condition": func(has_mpi, has_output bool) bool {
		return has_output
	},
	"UpdateMsg": OptHDF5Upd{},
	"ModuleTemplate": func(
		sfx string,
		install_path string,
		is_debug bool,
		has_mpi bool,
		cpu_arch string,
		gpu_arch string,
		opts map[string]string) []string {
		modules := []string{}
		if cpu_arch != "" {
			if cxx, ok := opts["CXX"]; ok {
				modules = append(modules, cxx)
			} else {
				panic("CXX module is not specified")
			}
		}
		if has_mpi {
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
		modulefile = append(modulefile, "conflict           hdf5")
		modulefile = append(modulefile, "")
		modulefile = append(modulefile, "prereqs						 "+strings.Join(modules, " "))
		modulefile = append(modulefile, "")
		modulefile = append(modulefile, "set                basedir      "+install_path)
		modulefile = append(modulefile, "prepend-path       PATH                   $basedir/bin")
		modulefile = append(modulefile, "prepend-path       LD_LIBRARY_PATH        $basedir/lib")
		modulefile = append(modulefile, "prepend-path       LIBRARY_PATH           $basedir/lib")
		modulefile = append(modulefile, "prepend-path       MANPATH                $basedir/man")
		modulefile = append(modulefile, "prepend-path       HDF5_ROOT              $basedir")
		modulefile = append(modulefile, "prepend-path       HDF5DIR                $basedir")
		modulefile = append(modulefile, "append-path        -d { } LDFLAGS         -L$basedir/lib")
		modulefile = append(modulefile, "append-path        -d { } INCLUDE         -I$basedir/include")
		modulefile = append(modulefile, "append-path        CPATH                  $basedir/include")
		modulefile = append(modulefile, "append-path        -d { } FFLAGS          -I$basedir/include")
		modulefile = append(modulefile, "append-path        -d { } FCFLAGS         -I$basedir/include")
		modulefile = append(modulefile, "append-path        -d { } LOCAL_LDFLAGS   -L$basedir/lib")
		modulefile = append(modulefile, "append-path        -d { } LOCAL_INCLUDE   -I$basedir/include")
		modulefile = append(modulefile, "append-path        -d { } LOCAL_CFLAGS    -I$basedir/include")
		modulefile = append(modulefile, "append-path        -d { } LOCAL_FFLAGS    -I$basedir/include")
		modulefile = append(modulefile, "append-path        -d { } LOCAL_FCFLAGS   -I$basedir/include")
		modulefile = append(modulefile, "append-path        -d { } LOCAL_CXXFLAGS  -I$basedir/include")
		return modulefile
	},
	"BuildScript": func(
		src_path, install_path string,
		is_debug, has_mpi bool,
		cpu_arch, gpu_arch string,
		use_modules bool,
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
		if has_mpi {
			if mpi, ok := opts["MPI"]; ok {
				modules = append(modules, mpi)
			} else {
				panic("MPI enabled but MPI module is not specified")
			}
		}
		flags := fmt.Sprintf("-S HDF5config.cmake,HPC=sbatch,MPI=%v,BUILD_GENERATOR=Unix,INSTALLDIR=%s", has_mpi, install_path)
		flags += " -C Release -V -O hdf5.log"

		prebuild := []string{
			"module purge",
			"module load " + strings.Join(modules, " "),
		}
		configure := []string{
			"cd " + src_path,
			"rm -rf build",
		}
		build := []string{
			"ctest " + flags,
			"cd build",
			"make install",
			"cd HDF5_ZLIB-prefix/src/HDF5_ZLIB-build",
			"make install",
			"cd ../../../SZIP-prefix/src/SZIP-build",
			"make install",
		}
		postbuild := []string{
			"cd " + src_path,
			"rm -rf build",
		}
		return prebuild, configure, build, postbuild
	},
}
