package configs

import "strings"

type EntitySuffixFunc = (func(bool, bool, string, string) string)
type EntityModTemplate = (func(string, bool, bool, string, string, []string) []string)

var EntityConfigs = map[string]interface{}{
	"Name": "Entity",
	"Suffix": func(is_debug bool, has_mpi bool, cpu_arch string, gpu_arch string) string {
		suffix := "entity"
		if is_debug {
			suffix += "/debug"
		}
		if has_mpi {
			suffix += "/mpi"
		} else {
			suffix += "/serial"
		}
		suffix += "/" + strings.ToLower(cpu_arch)
		if gpu_arch != "" {
			suffix += "/" + strings.ToLower(gpu_arch)
		}
		return suffix
	},
	"ModuleTemplate": func(sfx string, is_debug bool, has_mpi bool, cpu_arch string, gpu_arch string, modules []string) []string {
		modulefile := []string{}
		modulefile = append(modulefile, "#%Module1.0######################################################################")
		modulefile = append(modulefile, "##")
		conf := strings.ToUpper(strings.ReplaceAll(sfx, "/", " @ "))
		modulefile = append(modulefile, "## "+conf+"")
		modulefile = append(modulefile, "##")
		modulefile = append(modulefile, "################################################################################")
		modulefile = append(modulefile, "proc ModulesHelp { } {")
		modulefile = append(modulefile, "	 puts stderr \"\\t"+conf+"\\n\"")
		modulefile = append(modulefile, "}")
		modulefile = append(modulefile, "")
		modulefile = append(modulefile, "module-whatis      \"Sets up "+conf+"\"")
		modulefile = append(modulefile, "")
		modulefile = append(modulefile, "conflict           entity")
		modulefile = append(modulefile, "")
		if cpu_arch != "" {
			modulefile = append(modulefile, "setenv             Kokkos_ARCH_"+strings.ToUpper(cpu_arch)+" ON")
		}
		if gpu_arch != "" {
			modulefile = append(modulefile, "setenv             Kokkos_ARCH_"+strings.ToUpper(gpu_arch)+" ON")
			modulefile = append(modulefile, "setenv             Kokkos_ENABLE_CUDA ON")
		}
		if !has_mpi {
			modulefile = append(modulefile, "setenv             Kokkos_ENABLE_OPENMP ON")
			modulefile = append(modulefile, "")
			modulefile = append(modulefile, "setenv             OMP_PROC_BIND spread")
			modulefile = append(modulefile, "setenv             OMP_PLACES threads")
			modulefile = append(modulefile, "setenv             OMP_NUM_THREADS [exec nproc]")
		}
		modulefile = append(modulefile, "")
		if is_debug {
			modulefile = append(modulefile, "setenv             Entity_ENABLE_DEBUG ON")
		}
		if has_mpi {
			modulefile = append(modulefile, "setenv             Entity_ENABLE_MPI ON")
		}
		for _, module := range modules {
			modulefile = append(modulefile, "module load        "+module+"")
		}
		return modulefile
	},
}
