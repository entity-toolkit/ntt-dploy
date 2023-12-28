package configs

type SuffixFunc = (func(bool, string, string) string)
type ConditionFunc = (func(bool, bool) bool)
type BuildScript = (func(string, string, bool, bool, string, string, map[string]string) ([]string, []string, []string, []string))
type ModuleTemplate = (func(string, string, bool, bool, string, string, map[string]string) []string)

var (
	DefaultOptPath    = "$HOME/opt/"
	DefaultOptModPath = "$HOME/opt/.modules/"
	DefaultModPath    = "$HOME/.modules/"
	DefaultCUDA       = "module:cuda/12.0"
	DefaultCXX        = "module:gcc"
	CpuArchs          = []string{"NATIVE", "A64FX", "AMDAVX", "ARMV80", "ARMV81", "ARMV8_THUNDERX", "ARMV8_THUNDERX2", "BDW", "BGQ", "HSW", "KNC", "KNL", "POWER7", "POWER8", "POWER9", "SKX", "SNB", "SPR", "WSM", "ZEN", "ZEN2", "ZEN3"}
	GpuArchs          = []string{"ADA89", "AMD_GFX906", "AMD_GFX908", "AMD_GFX90A", "AMD_GFX1030", "AMD_GFX1100", "AMPERE80", "AMPERE86", "HOPPER90", "INTEL_GEN", "INTEL_DG1", "INTEL_GEN9", "INTEL_GEN11", "INTEL_GEN12LP", "INTEL_XEHP", "INTEL_PVC", "KEPLER30", "KEPLER32", "KEPLER35", "KEPLER37", "MAXWELL50", "MAXWELL52", "MAXWELL53", "NAVI1030", "PASCAL60", "PASCAL61", "TURING75", "VEGA900", "VEGA906", "VEGA908", "VEGA90A", "VOLTA70", "VOLTA72"}
)

type OptKokkosUpd struct{}
type OptMPIUpd struct{}
type OptADIOS2Upd struct{}
type OptHDF5Upd struct{}
