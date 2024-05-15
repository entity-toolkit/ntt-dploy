package pages

import (
	"github.com/entity-toolkit/ntt-dploy/configs"
	"github.com/entity-toolkit/ntt-dploy/global"
	"github.com/haykh/tuigo"
	"slices"
)

func SpecsSelector(prev tuigo.Window) tuigo.Window {
	// arch_choices := tuigo.Container(configs.FOCUSABLE, tuigo.HorizontalContainerTop,
	// 	tuigo.Container(configs.FOCUSABLE, tuigo.VerticalContainer,
	// 		tuigo.Text(global.Selectors["CPUARCH"].Text()+":", tuigo.NormalText),
	// 		tuigo.SelectorWithID(
	// 			global.Selectors["CPUARCH"].Id(),
	// 			global.Selectors["CPUARCH"].Options().([]string),
	// 			configs.SELECT_ONE,
	// 			configs.NO_CALLBACK,
	// 		),
	// 	),
	// 	tuigo.Container(configs.FOCUSABLE, tuigo.VerticalContainer,
	// 		tuigo.Text(global.Selectors["GPUARCH"].Text()+":", tuigo.NormalText),
	// 		tuigo.SelectorWithID(
	// 			global.Selectors["GPUARCH"].Id(),
	// 			global.Selectors["GPUARCH"].Options().([]string),
	// 			configs.SELECT_ONE,
	// 			configs.NO_CALLBACK,
	// 		),
	// 	),
	// )
	// debug_choice := tuigo.RadioWithID(
	// 	global.Selectors["DEBUG"].Id(),
	// 	global.Selectors["DEBUG"].Text(),
	// 	configs.NO_CALLBACK,
	// )
	// cuda_choice := tuigo.RadioWithID(
	// 	global.Selectors["CUDA"].Id(),
	// 	global.Selectors["CUDA"].Text(),
	// 	configs.NO_CALLBACK,
	// )
	// mpi_choice := tuigo.RadioWithID(
	// 	global.Selectors["MPI"].Id(),
	// 	global.Selectors["MPI"].Text(),
	// 	configs.NO_CALLBACK,
	// )
	// output_choice := tuigo.RadioWithID(
	// 	global.Selectors["OUTPUT"].Id(),
	// 	global.Selectors["OUTPUT"].Text(),
	// 	configs.NO_CALLBACK,
	// )
	if global.Selectors["MODE"].Read(prev) != nil {
		panic("MODE selector not read")
	}

	install_mode := global.Selectors["MODE"].Value().([]string)
	// return tuigo.Container(
	// 	configs.FOCUSABLE,
	// 	tuigo.VerticalContainer,
	// 	tuigo.Text("3. No additional configuration necessary (skip this step)", tuigo.LabelText),
	// )
	if slices.Contains(install_mode, "Entity") {
	}
	if slices.Contains(install_mode, "Kokkos") {
	}
	if slices.Contains(install_mode, "ADIOS2") {
	}
	// if install_mode == "Entity" {
	// 	return tuigo.Container(
	// 		configs.FOCUSABLE,
	// 		tuigo.VerticalContainer,
	// 		tuigo.Text(global.LABELS["pg2:title_entity"], tuigo.LabelText),
	// 		tuigo.Text("options:", tuigo.NormalText),
	// 		debug_choice,
	// 		cuda_choice,
	// 		mpi_choice,
	// 		output_choice,
	// 		arch_choices,
	// 	)
	// } else if install_mode == "Kokkos" {
	// 	return tuigo.Container(
	// 		configs.FOCUSABLE,
	// 		tuigo.VerticalContainer,
	// 		tuigo.Text(global.LABELS["pg2:title_kokkos"], tuigo.LabelText),
	// 		tuigo.InputWithID(
	// 			global.Selectors["KOKKOS_SRC_DIR"].Id(),
	// 			global.Selectors["KOKKOS_SRC_DIR"].Text(),
	// 			"", "",
	// 			tuigo.PathInput,
	// 			configs.NO_CALLBACK,
	// 		),
	// 		tuigo.InputWithID(
	// 			global.Selectors["KOKKOS_INSTALL_DIR"].Id(),
	// 			global.Selectors["KOKKOS_INSTALL_DIR"].Text(),
	// 			global.Selectors["KOKKOS_INSTALL_DIR"].Def().(string), "",
	// 			tuigo.PathInput,
	// 			configs.NO_CALLBACK,
	// 		),
	// 		debug_choice,
	// 		cuda_choice,
	// 		arch_choices,
	// 	)
	// } else if install_mode == "ADIOS2" {
	// 	return tuigo.Container(
	// 		configs.FOCUSABLE,
	// 		tuigo.VerticalContainer,
	// 		tuigo.Text(global.LABELS["pg2:title_adios2"], tuigo.LabelText),
	// 		tuigo.InputWithID(
	// 			global.Selectors["ADIOS2_SRC_DIR"].Id(),
	// 			global.Selectors["ADIOS2_SRC_DIR"].Text(),
	// 			"", "",
	// 			tuigo.PathInput,
	// 			configs.NO_CALLBACK,
	// 		),
	// 		tuigo.InputWithID(
	// 			global.Selectors["ADIOS2_INSTALL_DIR"].Id(),
	// 			global.Selectors["ADIOS2_INSTALL_DIR"].Text(),
	// 			global.Selectors["ADIOS2_INSTALL_DIR"].Def().(string), "",
	// 			tuigo.PathInput,
	// 			configs.NO_CALLBACK,
	// 		),
	// 		mpi_choice,
	// 		cuda_choice,
	// 		tuigo.Text(global.LABELS["pg2:comment_adios2"], tuigo.DimmedText),
	// 	)
	// } else {
	// 	panic("invalid mode selected")
	// }
}
