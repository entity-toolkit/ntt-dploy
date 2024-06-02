package configs

import (
	tea "github.com/charmbracelet/bubbletea"
)

type SuffixFunc = (func(bool, string, string) string)
type ConditionFunc = (func(bool, bool) bool)
type BuildScript = (func(string, string, bool, bool, string, string, bool, map[string]string) ([]string, []string, []string, []string))
type ModuleTemplate = (func(string, string, bool, bool, string, string, map[string]string) []string)

var (
	DefaultOptPath    = "$HOME/opt/"
	DefaultOptModPath = "$HOME/opt/.modules/"
	DefaultModPath    = "$HOME/.modules/"
	DefaultCUDA       = "module:cuda/12.0"
	DefaultCXX        = "module:gcc"
)

// type OptKokkosUpd struct{}
// type OptMPIUpd struct{}
// type OptADIOS2Upd struct{}
// type OptHDF5Upd struct{}

// const NONFOCUSABLE = true
// const FOCUSABLE = true
// const MULTI_SELECT = true
// const SELECT_ONE = false

var NO_CALLBACK tea.Msg = nil

// var Global = map[string]interface{}{}

// var GetArchs = func(archstr string) (string, string) {
// 	archs := []string{}
// 	for _, arch := range strings.Split(archstr, ",") {
// 		arch = strings.TrimSpace(arch)
// 		if arch != "" {
// 			archs = append(archs, arch)
// 		}
// 	}
// 	if len(archs) == 1 {
// 		archs = append(archs, "")
// 	}
// 	return archs[0], archs[1]
// }

// var DependencyMapping = map[int](map[string]interface{}){
// 	1000: KokkosConfigs,
// 	1001: MPIConfigs,
// 	1002: ADIOS2Configs,
// 	1003: HDF5Configs,
// }
