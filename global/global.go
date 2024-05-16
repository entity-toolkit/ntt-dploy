package global

import (
	"errors"
	"math/rand"

	"github.com/haykh/tuigo"
)

type Selector struct {
	id       int
	text     string
	value    interface{}
	options  interface{}
	def      interface{}
	callback interface{}
}

var Labels = map[string]string{
	"title": "Pick & configure the libraries to install",
}

const NMAX = 1000000

// type MODE_CALLBACK struct{}
// type CUDA_CALLBACK struct{}
// type MPI_CALLBACK struct{}
// type OUTPUT_CALLBACK struct{}
// type KOKKOS_SRC_CALLBACK struct{}
// type ADIOS2_SRC_CALLBACK struct{}

type StateChange struct{}

var Selectors = map[string]*Selector{
	"MODE": {
		id:       rand.Intn(NMAX),
		options:  []string{"Entity", "Kokkos", "ADIOS2", "MPI", "HDF5"},
		callback: StateChange{},
	},
	"MODE_TEXT": {
		id:   rand.Intn(NMAX),
		text: "Configure selected libraries below",
		def:  "No libraries selected",
	},
	"DEBUG": {
		id:   rand.Intn(NMAX),
		text: "debug build",
	},
	"CXX_PATH": {
		id:   rand.Intn(NMAX),
		text: "Host compiler",
	},
	"CUDA": {
		id:       rand.Intn(NMAX),
		text:     "CUDA support",
		callback: StateChange{},
		def:      true,
	},
	"CUDA_PATH": {
		id:   rand.Intn(NMAX),
		text: "CUDA compiler",
	},
	"MPI": {
		id:       rand.Intn(NMAX),
		text:     "MPI support",
		callback: StateChange{},
	},
	"MPI_SRC_DIR": {
		id:   rand.Intn(NMAX),
		text: "path to MPI source code",
	},
	"MPI_INSTALL_DIR": {
		id:   rand.Intn(NMAX),
		text: "MPI install path",
		def:  "$MPI_HOME",
	},
	"OUTPUT": {
		id:       rand.Intn(NMAX),
		text:     "enable output",
		callback: StateChange{},
		def:      true,
	},
	"HDF5_SRC_DIR": {
		id:   rand.Intn(NMAX),
		text: "path to HDF5 source code",
	},
	"HDF5_INSTALL_DIR": {
		id:   rand.Intn(NMAX),
		text: "HDF5 install path",
		def:  "$HDF5_DIR",
	},
	"CPUARCH": {
		id:   rand.Intn(NMAX),
		text: "CPU architecture",
		options: []string{
			"NATIVE", "A64FX", "AMDAVX", "ARMV80",
			"ARMV81", "ARMV8_THUNDERX", "ARMV8_THUNDERX2", "BDW",
			"BGQ", "HSW", "KNC", "KNL", "POWER7",
			"POWER8", "POWER9", "SKX", "SNB", "SPR",
			"WSM", "ZEN", "ZEN2", "ZEN3",
		},
	},
	"GPUARCH": {
		id:   rand.Intn(NMAX),
		text: "GPU architecture",
		options: []string{
			"ADA89", "AMD_GFX906", "AMD_GFX908", "AMD_GFX90A",
			"AMD_GFX1030", "AMD_GFX1100", "AMPERE80", "AMPERE86",
			"HOPPER90", "INTEL_GEN", "INTEL_DG1", "INTEL_GEN9",
			"INTEL_GEN11", "INTEL_GEN12LP", "INTEL_XEHP", "INTEL_PVC",
			"KEPLER30", "KEPLER32", "KEPLER35", "KEPLER37", "MAXWELL50",
			"MAXWELL52", "MAXWELL53", "NAVI1030", "PASCAL60", "PASCAL61",
			"TURING75", "VEGA900", "VEGA906", "VEGA908",
			"VEGA90A", "VOLTA70", "VOLTA72",
		},
	},
	"KOKKOS_SRC_DIR": {
		id:       rand.Intn(NMAX),
		text:     "path to Kokkos source code",
		callback: StateChange{},
	},
	"KOKKOS_INSTALL_DIR": {
		id:   rand.Intn(NMAX),
		text: "Kokkos install path",
		def:  "$HOME/opt/Kokkos/",
	},
	"ADIOS2_SRC_DIR": {
		id:       rand.Intn(NMAX),
		text:     "path to ADIOS2 source code",
		callback: StateChange{},
	},
	"ADIOS2_INSTALL_DIR": {
		id:   rand.Intn(NMAX),
		text: "ADIOS2 install path",
		def:  "$HOME/opt/ADIOS2/",
	},
	"COMMENT_MODULES": {
		id:   rand.Intn(NMAX),
		text: "[*] To configure the Entity,\n    environment modules are required",
	},
	"COMMENT_ADIOS2": {
		id:   rand.Intn(NMAX),
		text: "[*] ADIOS2 depends on\n    Kokkos and HDF5 installations",
	},
	"HELP_ENVMOD": {
		id:   rand.Intn(NMAX),
		text: "[*] To use environment module as a\n    dependency, specify `module:<MOD>`\n    instead of a path\n    Example: `MPI install path: module:ompi/4.0.5`",
	},
	"HELP_KEYS": {
		id:   rand.Intn(NMAX),
		text: "[Tab/Shift-Tab] Next/previous field\n[Space] Select/Unselect\n[Up/Down] Navigate options\n[Esc] Quit",
	},
}

func (s *Selector) Read(w tuigo.Window) error {
	val := w.GetElementByID(s.id).Data()
	if val != nil {
		s.value = val
	} else {
		return errors.New("Selector value is nil")
	}
	return nil
}

func (s Selector) Id() int {
	return s.id
}

func (s Selector) Text() string {
	return s.text
}

func (s Selector) Value() interface{} {
	return s.value
}

func (s Selector) Options() interface{} {
	return s.options
}

func (s Selector) Def() interface{} {
	return s.def
}

func (s Selector) Callback() interface{} {
	return s.callback
}

// var ENTITY_PATH = Selector{
// 	Id: rand.Intn(NMAX),
// }

// var CPUARCH = Selector{
// 	Id:   rand.Intn(NMAX),
// 	Text: "CPU architecture:",
// 	Options: []string{
// 		"NATIVE", "A64FX", "AMDAVX", "ARMV80",
// 		"ARMV81", "ARMV8_THUNDERX", "ARMV8_THUNDERX2", "BDW",
// 		"BGQ", "HSW", "KNC", "KNL", "POWER7",
// 		"POWER8", "POWER9", "SKX", "SNB", "SPR",
// 		"WSM", "ZEN", "ZEN2", "ZEN3",
// 	},
// 	Value: nil,
// }

// var GPUARCH = Selector{
// 	Id:   rand.Intn(NMAX),
// 	Text: "GPU architecture:",
// 	Options: []string{
// 		"ADA89", "AMD_GFX906", "AMD_GFX908", "AMD_GFX90A",
// 		"AMD_GFX1030", "AMD_GFX1100", "AMPERE80", "AMPERE86",
// 		"HOPPER90", "INTEL_GEN", "INTEL_DG1", "INTEL_GEN9",
// 		"INTEL_GEN11", "INTEL_GEN12LP", "INTEL_XEHP", "INTEL_PVC",
// 		"KEPLER30", "KEPLER32", "KEPLER35", "KEPLER37", "MAXWELL50",
// 		"MAXWELL52", "MAXWELL53", "NAVI1030", "PASCAL60", "PASCAL61",
// 		"TURING75", "VEGA900", "VEGA906", "VEGA908",
// 		"VEGA90A", "VOLTA70", "VOLTA72",
// 	},
// 	Value: nil,
// }

// var ARCHS = Selector{
// 	Id:   rand.Intn(NMAX),
// 	Text: "architectures",
// }

// var MPI = Selector{
// 	Id:   rand.Intn(NMAX),
// 	Text: "use MPI",
// }

// var OUTPUT = Selector{
// 	Id:   rand.Intn(NMAX),
// 	Text: "enable output",
// }

// var DEBUG = Selector{
// 	Id:   rand.Intn(NMAX),
// 	Text: "enable debug mode",
// }

// var CUDA = Selector{
// 	Id:   rand.Intn(NMAX),
// 	Text: "enable CUDA support",
// }

// var KOKKOS_SRC_DIR = Selector{
// 	Id:   rand.Intn(NMAX),
// 	Text: "Kokkos source path",
// }

// var KOKKOS_INSTALL_DIR = Selector{
// 	Id:   rand.Intn(NMAX),
// 	Text: "Kokkos install path",
// }

// var ADIOS2_SRC_DIR = Selector{
// 	Id:   rand.Intn(NMAX),
// 	Text: "ADIOS2 source path",
// }

// var ADIOS2_INSTALL_DIR = Selector{
// 	Id:   rand.Intn(NMAX),
// 	Text: "ADIOS2 install path",
// }

// var MODULES_DIR = Selector{
// 	Id:   rand.Intn(NMAX),
// 	Text: "module parent directory",
// }

// var MODULE = Selector{
// 	Id:   rand.Intn(NMAX),
// 	Text: "module path",
// }

// var CXX_MODULE = Selector{
// 	Id:   rand.Intn(NMAX),
// 	Text: "CXX module",
// }

// var CUDA_MODULE = Selector{
// 	Id:   rand.Intn(NMAX),
// 	Text: "CUDA module",
// }

// var INSTALL_MODULE_PATHS = Selector{
// 	Id:   rand.Intn(NMAX),
// 	Text: "Install & module paths",
// }

// var PARENT_DIR = Selector{
// 	Id:   rand.Intn(NMAX),
// 	Text: "parent directory",
// }

// const SELECTOR_CPU = rand.Intn(NMAX)
// const GPU_SELECTOR = rand.Intn(NMAX)
// const USE_MPI = rand.Intn(NMAX)
// const ENABLE_OUTPUT = 203
// const ENABLE_DEBUG = 204
// const KOKKOS_SRC = 205
// const KOKKOS_INSTALL = 206
// const ADIOS2_SRC = 207
// const ADIOS2_INSTALL = 208
// const USE_CUDA = 209
// const KOKKOS_PATH = 210
// const MODULE_DIR = 301
// const MODULE_PATH = 302
// const NTT_PATH = 303
// const ARCHITECTURES = 304
// const INSTALL_MODULE_PATH = 305
// const PARENT_DIR = 306
