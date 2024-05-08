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
		// modulefile := []string{}
		conf := strings.ToUpper(strings.ReplaceAll(sfx, "/", " @ "))
		cpu_setenv := []string{}
		gpu_setenv := []string{}
		if cpu_arch != "" {
			cpu_setenv = []string{
				"setenv             Kokkos_ARCH_" + strings.ToUpper(cpu_arch) + "ON",
			}
		}
		if gpu_arch != "" {
			gpu_setenv = []string{
				"setenv             Kokkos_ARCH_" + strings.ToUpper(gpu_arch) + "ON",
				"setenv             Kokkos_ENABLE_CUDA ON",
			}
		}
		modulefile := []string{
			"#%Module1.0######################################################################",
			"##",
			"##" + conf,
			"##",
			"################################################################################",
			"proc ModulesHelp { } {",
			"  puts stderr \"\\t" + conf + "\\n\"",
			"}",
			"",
			"module-whatis      \"Sets up " + conf + "\"",
			"",
			"conflict           entity",
			"",
			cpu_setenv...,
			gpu_setenv...,
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
