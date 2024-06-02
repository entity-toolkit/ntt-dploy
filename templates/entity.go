package templates

var EntityTemplate = LibraryConfiguration{
	InitFlagsTemplate: func(self *LibraryConfiguration) {
		if self.Settings["debug"].(bool) {
			self.AddFlag("Entity_ENABLE_DEBUG ON")
		}
		if self.Settings["mpi"].(bool) {
			self.AddFlag("Entity_ENABLE_MPI ON")
		} else if self.Settings["gpu"].(string) == "" {
			self.AddFlag("Entity_ENABLE_OPENMP ON")
		}
	},
	CheckTemplate: func(self LibraryConfiguration) bool {
		required_settings := []string{
			"mpi",
			"cpu",
			"gpu",
			"modules",
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
		return true
	},
	SuffixTemplate: CreateTemplate("suffix", "entity"+DEBUG_suffix+MPI_suffix+GPU_suffix+CPU_suffix),
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

conflict           entity
{{if not .modules}}{{else}}{{range .modules}}
module load       {{.}}{{end}}{{end}}
`,
	),
	BuildScriptTemplate: nil,
}

// import (
// 	"github.com/entity-toolkit/ntt-dploy/configs"
// 	"strings"
// 	"text/template"
// )
//
// type EntitySuffixFunc = (func(bool, bool, string, string) string)
// type EntityModTemplate = (func(string, bool, bool, string, string, []string) []string)
//
// var EntityConfigs = map[string]interface{}{
// 	"Name": "Entity",
// 	"Suffix": func(is_debug bool, has_mpi bool, cpu_arch string, gpu_arch string) string {
// 		suffix := "entity"
// 		if is_debug {
// 			suffix += "/debug"
// 		}
// 		if has_mpi {
// 			suffix += "/mpi"
// 		} else {
// 			suffix += "/serial"
// 		}
// 		suffix += "/" + strings.ToLower(cpu_arch)
// 		if gpu_arch != "" {
// 			suffix += "/" + strings.ToLower(gpu_arch)
// 		}
// 		return suffix
// 	},
// 	"ModuleTemplate": func(sfx string, is_debug bool, has_mpi bool, cpu_arch string, gpu_arch string, modules []string) []string {
// 		conf := strings.ToUpper(strings.ReplaceAll(sfx, "/", " @ "))
// 		cpu_setenv := []string{}
// 		if cpu_arch != "" {
// 			cpu_setenv = []string{
// 				"setenv             Kokkos_ARCH_" + strings.ToUpper(cpu_arch) + "ON",
// 			}
// 		}
// 		gpu_setenv := []string{}
// 		if gpu_arch != "" {
// 			gpu_setenv = []string{
// 				"setenv             Kokkos_ARCH_" + strings.ToUpper(gpu_arch) + "ON",
// 				"setenv             Kokkos_ENABLE_CUDA ON",
// 			}
// 		}
// 		omp_setenv := []string{}
// 		if !has_mpi {
// 			omp_setenv = []string{
// 				"setenv             Entity_ENABLE_OPENMP ON",
// 				"setenv             OMP_PROC_BIND spread",
// 				"setenv             OMP_PLACES threads",
// 				"setenv             OMP_NUM_THREADS [exec nproc]",
// 			}
// 		}
// 		debug_setenv := []string{}
// 		if is_debug {
// 			debug_setenv = []string{
// 				"setenv             Entity_ENABLE_DEBUG ON",
// 			}
// 		}
// 		mpi_setenv := []string{}
// 		if has_mpi {
// 			mpi_setenv = []string{
// 				"setenv             Entity_ENABLE_MPI ON",
// 			}
// 		}
// 		module_setenv := []string{}
// 		for _, module := range modules {
// 			module_setenv = append(module_setenv, "module load        "+module+"")
// 		}
// 		return append([]string{
// 			"#%Module1.0######################################################################",
// 			"##",
// 			"##" + conf,
// 			"##",
// 			"################################################################################",
// 			"proc ModulesHelp { } {",
// 			"  puts stderr \"\\t" + conf + "\\n\"",
// 			"}",
// 			"",
// 			"module-whatis      \"Sets up " + conf + "\"",
// 			"",
// 			"conflict           entity",
// 			"",
// 			cpu_setenv[0],
// 			gpu_setenv[0],
// 			gpu_setenv[1],
// 			omp_setenv[0],
// 			"",
// 			omp_setenv[1],
// 			omp_setenv[2],
// 			omp_setenv[3],
// 			debug_setenv[0],
// 			mpi_setenv[0],
// 			"",
// 		},
// 			module_setenv...,
// 		)
// 	},
// }
