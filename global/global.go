package global

import (
	rand "math/rand"
	time "time"
)

const NMAX = 100000

type Selector struct {
	Id      int
	Text    string
	Value   interface{}
	Options interface{}
}

var MODE = Selector{
	Id:   rand.Intn(NMAX),
	Text: "1. Choose whether you want to install a separate library or the modules for the Entity itself",
}

var ENTITY_CONFIG = Selector{
	Text: "2. Pick Entity configuration to install",
}

var KOKKOS_CONFIG = Selector{
	Text: "2. Pick Kokkos configuration to install",
}

var ADIOS2_CONFIG = Selector{
	Text: "2. Pick ADIOS2 configuration to install",
}

var CPUARCH = Selector{
	Id:   rand.Intn(NMAX),
	Text: "CPU architecture:",
	Options: []string{
		"NATIVE", "A64FX", "AMDAVX", "ARMV80",
		"ARMV81", "ARMV8_THUNDERX", "ARMV8_THUNDERX2", "BDW",
		"BGQ", "HSW", "KNC", "KNL", "POWER7",
		"POWER8", "POWER9", "SKX", "SNB", "SPR",
		"WSM", "ZEN", "ZEN2", "ZEN3",
	},
	Value: nil,
}

var GPUARCH = Selector{
	Id:   rand.Intn(NMAX),
	Text: "GPU architecture:",
	Options: []string{
		"ADA89", "AMD_GFX906", "AMD_GFX908", "AMD_GFX90A",
		"AMD_GFX1030", "AMD_GFX1100", "AMPERE80", "AMPERE86",
		"HOPPER90", "INTEL_GEN", "INTEL_DG1", "INTEL_GEN9",
		"INTEL_GEN11", "INTEL_GEN12LP", "INTEL_XEHP", "INTEL_PVC",
		"KEPLER30", "KEPLER32", "KEPLER35", "KEPLER37", "MAXWELL50",
		"MAXWELL52", "MAXWELL53", "NAVI1030", "PASCAL60", "PASCAL61",
		"TURING75", "VEGA900", "VEGA906", "VEGA908",
		"VEGA90A", "VOLTA70", "VOLTA72",
	},
	Value: nil,
}

var MPI = Selector{
	Id:   rand.Intn(NMAX),
	Text: "use MPI",
}

var OUTPUT = Selector{
	Id:   rand.Intn(NMAX),
	Text: "enable output",
}

var DEBUG = Selector{
	Id:   rand.Intn(NMAX),
	Text: "enable debug mode",
}

var CUDA = Selector{
	Id:   rand.Intn(NMAX),
	Text: "enable CUDA support",
}

var KOKKOS_SRC_DIR = Selector{
	Id:   rand.Intn(NMAX),
	Text: "Kokkos source path",
}

var KOKKOS_INSTALL_DIR = Selector{
	Id:   rand.Intn(NMAX),
	Text: "Kokkos install path",
}

var ADIOS2_SRC_DIR = Selector{
	Id:   rand.Intn(NMAX),
	Text: "ADIOS2 source path",
}

var ADIOS2_INSTALL_DIR = Selector{
	Id:   rand.Intn(NMAX),
	Text: "ADIOS2 install path",
}

var COMMENTS = map[string]string{
	"modules":        "[*] Installation of Entity is done via environment modules",
	"adios2 depends": "[*] Kokkos and HDF5 installation is required",
}

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
