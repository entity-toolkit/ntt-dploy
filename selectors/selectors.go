package selectors

import (
	"errors"
	"math/rand"

	"github.com/haykh/tuigo"
)

type Selector struct {
	id          int
	text        string
	value       interface{}
	options     interface{}
	def         interface{}
	placeholder string
	callback    interface{}
}

const NMAX = 1000000

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
	"ENV_MODULES": {
		id:       rand.Intn(NMAX),
		text:     "use environment modules",
		callback: StateChange{},
		def:      true,
	},
	"MODFILE_PATH": {
		id:       rand.Intn(NMAX),
		text:     "path to write module files",
		callback: StateChange{},
		def:      "$HOME/.modules/modfiles",
	},
	"CXX_MOD": {
		id:          rand.Intn(NMAX),
		text:        "Host compiler module",
		placeholder: "leave blank to use default",
	},
	"CUDA": {
		id:       rand.Intn(NMAX),
		text:     "CUDA support",
		callback: StateChange{},
		def:      true,
	},
	"CUDA_MOD": {
		id:          rand.Intn(NMAX),
		text:        "CUDA compiler module",
		placeholder: "leave blank to use default",
	},
	"MPI": {
		id:       rand.Intn(NMAX),
		text:     "MPI support",
		callback: StateChange{},
	},
	"MPI_SRC_DIR": {
		id:          rand.Intn(NMAX),
		text:        "path to MPI source code",
		def:         "$HOME/src/mpi",
		placeholder: "default",
	},
	"MPI_INSTALL_DIR": {
		id:          rand.Intn(NMAX),
		text:        "MPI install path",
		def:         "$HOME/.modules/mpi",
		placeholder: "default",
	},
	"HDF5_SRC_DIR": {
		id:          rand.Intn(NMAX),
		text:        "path to HDF5 source code",
		def:         "$HOME/src/hdf5",
		placeholder: "default",
	},
	"HDF5_INSTALL_DIR": {
		id:          rand.Intn(NMAX),
		text:        "HDF5 install path",
		def:         "$HOME/.modules/hdf5",
		placeholder: "default",
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
		id:          rand.Intn(NMAX),
		text:        "path to Kokkos source code",
		callback:    StateChange{},
		def:         "$HOME/src/kokkos",
		placeholder: "default",
	},
	"KOKKOS_INSTALL_DIR": {
		id:          rand.Intn(NMAX),
		text:        "Kokkos install path",
		def:         "$HOME/.modules/kokkos",
		placeholder: "default",
	},
	"ADIOS2_SRC_DIR": {
		id:          rand.Intn(NMAX),
		text:        "path to ADIOS2 source code",
		callback:    StateChange{},
		def:         "$HOME/src/adios2",
		placeholder: "default",
	},
	"ADIOS2_INSTALL_DIR": {
		id:          rand.Intn(NMAX),
		text:        "ADIOS2 install path",
		def:         "$HOME/.modules/adios2",
		placeholder: "default",
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
		id: rand.Intn(NMAX),
		text: `[*] To use environment module as a
dependency, specify module:<MOD>
instead of a path.
Example:
MPI install path: module:ompi/4.0.5`,
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

func (s Selector) Placeholder() string {
	return s.placeholder
}

func (s Selector) Callback() interface{} {
	return s.callback
}
