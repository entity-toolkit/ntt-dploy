package configs

import "strings"

var MPIConfigs = map[string]interface{}{
	"Name":           "MPI",
	"DefaultSrcPath": "$HOME/opt/src/ompi",
	"Suffix": func(mpi bool, cpu string, gpu string) string {
		return "mpi"
	},
	"Condition": func(has_mpi, has_output bool) bool {
		return has_mpi
	},
	"UpdateMsg": OptMPIUpd{},
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
		if gpu_arch != "" {
			if cuda, ok := opts["CUDA"]; ok {
				modules = append(modules, cuda)
			} else {
				panic("GPU enabled but CUDA module is not specified")
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
		modulefile = append(modulefile, "conflict           mpi openmpi ompi open-mpi")
		modulefile = append(modulefile, "")
		modulefile = append(modulefile, "prereqs						 "+strings.Join(modules, " "))
		modulefile = append(modulefile, "")
		modulefile = append(modulefile, "set                basedir               $install_path")
		modulefile = append(modulefile, "prepend-path       PATH                  $basedir/bin")
		modulefile = append(modulefile, "prepend-path       LD_LIBRARY_PATH       $basedir/lib")
		modulefile = append(modulefile, "")
		modulefile = append(modulefile, "append-path -d { } LOCAL_LDFLAGS      -L $basedir/lib")
		modulefile = append(modulefile, "append-path -d { } LOCAL_INCLUDE      -I $basedir/include")
		modulefile = append(modulefile, "append-path -d { } LOCAL_CFLAGS       -I $basedir/include")
		modulefile = append(modulefile, "append-path -d { } LOCAL_FCFLAGS      -I $basedir/include")
		modulefile = append(modulefile, "append-path -d { } LOCAL_CXXFLAGS     -I $basedir/include")
		modulefile = append(modulefile, "")
		modulefile = append(modulefile, "setenv             CXX                   $basedir/bin/mpicxx")
		modulefile = append(modulefile, "setenv             CC                    $basedir/bin/mpicc")
		modulefile = append(modulefile, "")
		modulefile = append(modulefile, "setenv             SLURM_MPI_TYPE        pmix_v3")
		modulefile = append(modulefile, "setenv             MPIHOME               $basedir")
		modulefile = append(modulefile, "setenv             MPI_HOME              $basedir")
		modulefile = append(modulefile, "setenv             OPENMPI_HOME          $basedir")
		return modulefile
	},
	"BuildScript": func(
		src_path, install_path string,
		is_debug, has_mpi bool,
		cpu_arch, gpu_arch string,
		opts map[string]string,
	) ([]string, []string, []string, []string) {
		compile_args := []string{}
		modules := []string{}
		compilers := []string{}
		if cpu_arch != "" {
			compilers = append(compilers, "export CC=$(which gcc) CXX=$(which g++)")
			if cxx, ok := opts["CXX"]; ok {
				modules = append(modules, cxx)
			} else {
				panic("CXX module is not specified")
			}
		}
		if gpu_arch != "" {
			compile_args = append(compile_args, "--with-cuda=$CUDA_HOME")
			if cuda, ok := opts["CUDA"]; ok {
				modules = append(modules, cuda)
			} else {
				panic("GPU enabled but CUDA module is not specified")
			}
		}
		compile_args = append(compile_args, "--with-devel-headers")
		compile_args = append(compile_args, "--prefix="+install_path)
		prebuild := []string{
			"module purge",
			"module load " + strings.Join(modules, " "),
		}
		prebuild = append(prebuild, compilers...)
		configure := []string{
			"cd " + src_path,
			"rm -rf build",
			"./autogen.pl",
			"mkdir -p build",
			"cd build",
			"../configure " + strings.Join(compile_args, " "),
		}
		build := []string{
			"make -j",
			"make install",
		}
		postbuild := []string{
			"cd " + src_path,
			"rm -rf build",
		}
		return prebuild, configure, build, postbuild
	},
}
