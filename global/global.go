package global

import (
	"errors"
	"github.com/haykh/tuigo"
	"math/rand"
)

type Selector struct {
	id      int
	text    string
	value   interface{}
	options interface{}
	def     interface{}
}

var LABELS = map[string]string{
	"pg1:title":           "1. Choose whether you want to install a separate library or the modules for the Entity itself",
	"pg1:comment_modules": "[*] Installation of Entity is done via environment modules",
	"pg2:title_entity":    "2. Pick Entity configuration to install",
	"pg2:title_kokkos":    "2. Pick Kokkos configuration to install",
	"pg2:title_adios2":    "2. Pick ADIOS2 configuration to install",
	"pg2:comment_adios2":  "[*] Kokkos and HDF5 installation is required",
}

const NMAX = 100000

var Selectors = map[string]*Selector{
	"MODE": {
		id:      rand.Intn(NMAX),
		options: []string{"Entity", "Kokkos", "ADIOS2"},
	},
	"DEBUG": {
		id:   rand.Intn(NMAX),
		text: "enable debug mode",
	},
	"CUDA": {
		id:   rand.Intn(NMAX),
		text: "enable CUDA support",
	},
	"MPI": {
		id:   rand.Intn(NMAX),
		text: "use MPI",
	},
	"OUTPUT": {
		id:   rand.Intn(NMAX),
		text: "enable output",
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
		id:   rand.Intn(NMAX),
		text: "path to Kokkos source code",
	},
	"KOKKOS_INSTALL_DIR": {
		id:   rand.Intn(NMAX),
		text: "Kokkos install path",
		def:  "$HOME/opt/Kokkos/",
	},
	"ADIOS2_SRC_DIR": {
		id:   rand.Intn(NMAX),
		text: "path to ADIOS2 source code",
	},
	"ADIOS2_INSTALL_DIR": {
		id:   rand.Intn(NMAX),
		text: "ADIOS2 install path",
		def:  "$HOME/opt/ADIOS2/",
	},
}

func (s *Selector) Read(w tuigo.Window) error {
	val := w.GetElementByID(s.id).(tuigo.SelectorElement).Data()
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
