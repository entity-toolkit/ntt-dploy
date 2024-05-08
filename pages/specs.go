package pages

import (
	"github.com/entity-toolkit/ntt-dploy/configs"
	"github.com/entity-toolkit/ntt-dploy/global"
	"github.com/haykh/tuigo"
)

func SpecsSelector(prev tuigo.Window) tuigo.Window {
	global.MODE.Value = prev.GetElementByID(
		global.MODE.Id,
	).(tuigo.SelectorElement).Data().(string)
	install_mode := global.MODE.Value.(string)
	arch_choices := tuigo.Container(configs.FOCUSABLE, tuigo.HorizontalContainerTop,
		tuigo.Container(configs.FOCUSABLE, tuigo.VerticalContainer,
			tuigo.Text(global.CPUARCH.Text, tuigo.NormalText),
			tuigo.SelectorWithID(
				global.CPUARCH.Id,
				global.CPUARCH.Options.([]string),
				configs.SELECT_ONE,
				configs.NO_CALLBACK,
			),
		),
		tuigo.Container(configs.FOCUSABLE, tuigo.VerticalContainer,
			tuigo.Text(global.GPUARCH.Text, tuigo.NormalText),
			tuigo.SelectorWithID(
				global.GPUARCH.Id,
				global.GPUARCH.Options.([]string),
				configs.SELECT_ONE,
				configs.NO_CALLBACK,
			),
		),
	)
	debug_choice := tuigo.RadioWithID(
		global.DEBUG.Id,
		global.DEBUG.Text,
		configs.NO_CALLBACK,
	)
	cuda_choice := tuigo.RadioWithID(
		global.CUDA.Id,
		global.CUDA.Text,
		configs.NO_CALLBACK,
	)
	mpi_choice := tuigo.RadioWithID(
		global.MPI.Id,
		global.MPI.Text,
		configs.NO_CALLBACK,
	)
	output_choice := tuigo.RadioWithID(
		global.OUTPUT.Id,
		global.OUTPUT.Text,
		configs.NO_CALLBACK,
	)
	kokkos_install_dir_choice := tuigo.InputWithID(
		global.KOKKOS_INSTALL_DIR.Id,
		global.KOKKOS_INSTALL_DIR.Text,
		"$HOME/opt/Kokkos/", "",
		tuigo.PathInput,
		configs.NO_CALLBACK,
	)

	if install_mode == "Entity" {
		return tuigo.Container(
			configs.FOCUSABLE,
			tuigo.VerticalContainer,
			tuigo.Text(global.ENTITY_CONFIG.Text, tuigo.LabelText),
			tuigo.Text("options:", tuigo.NormalText),
			mpi_choice,
			output_choice,
			debug_choice,
			arch_choices,
		)
	} else if install_mode == "Kokkos" {
		return tuigo.Container(
			configs.FOCUSABLE,
			tuigo.VerticalContainer,
			tuigo.Text(global.KOKKOS_CONFIG.Text, tuigo.LabelText),
			tuigo.InputWithID(
				global.KOKKOS_SRC_DIR.Id,
				global.KOKKOS_SRC_DIR.Text,
				"", "",
				tuigo.PathInput,
				configs.NO_CALLBACK,
			),
			kokkos_install_dir_choice,
			debug_choice,
			arch_choices,
		)
	} else if install_mode == "ADIOS2" {
		return tuigo.Container(
			configs.FOCUSABLE,
			tuigo.VerticalContainer,
			tuigo.Text(global.ADIOS2_CONFIG.Text, tuigo.LabelText),
			tuigo.InputWithID(
				global.ADIOS2_SRC_DIR.Id,
				global.ADIOS2_SRC_DIR.Text,
				"", "",
				tuigo.PathInput,
				configs.NO_CALLBACK,
			),
			tuigo.InputWithID(
				global.ADIOS2_INSTALL_DIR.Id,
				global.ADIOS2_INSTALL_DIR.Text,
				"$HOME/opt/ADIOS2/", "",
				tuigo.PathInput,
				configs.NO_CALLBACK,
			),
			mpi_choice,
			cuda_choice,
			kokkos_install_dir_choice,
			tuigo.Text(global.COMMENTS["adios2 depends"], tuigo.DimmedText),
		)
	} else {
		panic("invalid mode selected")
	}
}
