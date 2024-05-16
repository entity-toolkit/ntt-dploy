package pages

import (
	"slices"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/entity-toolkit/ntt-dploy/configs"
	"github.com/entity-toolkit/ntt-dploy/global"
	"github.com/haykh/tuigo"
)

func MainSelector(prev tuigo.Window) tuigo.Window {
	main := tuigo.Container(
		tuigo.Focusable,
		tuigo.VerticalContainer,
		tuigo.Text(
			global.Labels["title"],
			tuigo.LabelText,
		),
		tuigo.SelectorWithID(
			global.Selectors["MODE"].Id(),
			global.Selectors["MODE"].Options().([]string),
			tuigo.MultiSelect,
			tuigo.NoViewLimit,
			global.Selectors["MODE"].Callback(),
		),
	)
	cpuarch_choice_label := tuigo.TextWithID(
		global.Selectors["CPUARCH"].Id()+1,
		global.Selectors["CPUARCH"].Text()+":",
		tuigo.NormalText,
	)
	cpuarch_choice_elements := tuigo.SelectorWithID(
		global.Selectors["CPUARCH"].Id(),
		global.Selectors["CPUARCH"].Options().([]string),
		tuigo.SelectOne,
		5,
		global.Selectors["CPUARCH"].Callback(),
	)
	gpuarch_choice_label := tuigo.TextWithID(
		global.Selectors["GPUARCH"].Id()+1,
		global.Selectors["GPUARCH"].Text()+":",
		tuigo.NormalText,
	)
	gpuarch_choice_elements := tuigo.SelectorWithID(
		global.Selectors["GPUARCH"].Id(),
		global.Selectors["GPUARCH"].Options().([]string),
		tuigo.SelectOne,
		5,
		global.Selectors["GPUARCH"].Callback(),
	)
	debug_choice := tuigo.RadioWithID(
		global.Selectors["DEBUG"].Id(),
		global.Selectors["DEBUG"].Text(),
		configs.NO_CALLBACK,
	)
	cxx_path := tuigo.InputWithID(
		global.Selectors["CXX_PATH"].Id(),
		global.Selectors["CXX_PATH"].Text(),
		"", "[default]",
		tuigo.TextInput,
		configs.NO_CALLBACK,
	)
	cuda_choice := tuigo.RadioWithID(
		global.Selectors["CUDA"].Id(),
		global.Selectors["CUDA"].Text(),
		global.Selectors["CUDA"].Callback(),
	)
	cuda_path := tuigo.InputWithID(
		global.Selectors["CUDA_PATH"].Id(),
		global.Selectors["CUDA_PATH"].Text(),
		"", "[default]",
		tuigo.TextInput,
		configs.NO_CALLBACK,
	)
	if global.Selectors["CUDA"].Def().(bool) {
		cuda_choice = cuda_choice.Set().(tuigo.SimpleContainerElement)
	}
	mpi_choice := tuigo.RadioWithID(
		global.Selectors["MPI"].Id(),
		global.Selectors["MPI"].Text(),
		global.Selectors["MPI"].Callback(),
	)
	output_choice := tuigo.RadioWithID(
		global.Selectors["OUTPUT"].Id(),
		global.Selectors["OUTPUT"].Text(),
		global.Selectors["OUTPUT"].Callback(),
	)
	if global.Selectors["OUTPUT"].Def().(bool) {
		output_choice = output_choice.Set().(tuigo.SimpleContainerElement)
	}

	lib_names := []string{"KOKKOS", "ADIOS2", "MPI", "HDF5"}
	lib_elements := []tuigo.Component{}
	for _, name := range lib_names {
		src_dir_label := tuigo.TextWithID(
			global.Selectors[name+"_SRC_DIR"].Id()+1,
			"source path:",
			tuigo.NormalText,
		)
		src_dir := tuigo.InputWithID(
			global.Selectors[name+"_SRC_DIR"].Id(),
			"", "", "",
			tuigo.PathInput,
			global.Selectors[name+"_SRC_DIR"].Callback(),
		)
		install_dir_label := tuigo.TextWithID(
			global.Selectors[name+"_INSTALL_DIR"].Id()+1,
			"install path:",
			tuigo.NormalText,
		)
		install_dir := tuigo.InputWithID(
			global.Selectors[name+"_INSTALL_DIR"].Id(),
			"",
			global.Selectors[name+"_INSTALL_DIR"].Def().(string), "",
			tuigo.TextInput,
			configs.NO_CALLBACK,
		)
		text_label := tuigo.TextWithID(
			global.Selectors[name+"_INSTALL_DIR"].Id()+2,
			name, tuigo.LabelText,
		)
		text_label = text_label.Hide().(tuigo.SimpleContainerElement)
		src_dir = src_dir.Hide().(tuigo.SimpleContainerElement)
		install_dir = install_dir.Hide().(tuigo.SimpleContainerElement)
		src_dir_label = src_dir_label.Hide().(tuigo.SimpleContainerElement)
		install_dir_label = install_dir_label.Hide().(tuigo.SimpleContainerElement)
		lib_elements = append(
			lib_elements,
			tuigo.Container(
				tuigo.Focusable,
				tuigo.HorizontalContainer,
				text_label,
				tuigo.Container(
					tuigo.Focusable,
					tuigo.VerticalContainer,
					src_dir_label,
					src_dir,
				),
				tuigo.Container(
					tuigo.Focusable,
					tuigo.VerticalContainer,
					install_dir_label,
					install_dir,
				),
			),
		)
	}

	elements := tuigo.Components{}

	packages := tuigo.TextWithID(
		global.Selectors["MODE_TEXT"].Id(),
		global.Selectors["MODE_TEXT"].Def().(string),
		tuigo.NormalText,
	)
	elements = append(elements, packages)

	for _, el := range []*tuigo.SimpleContainerElement{
		&cuda_choice, &output_choice, &mpi_choice,
		&debug_choice, &cxx_path, &cpuarch_choice_label, &cpuarch_choice_elements,
		&gpuarch_choice_label, &gpuarch_choice_elements, &cuda_path,
	} {
		*el = el.Hide().(tuigo.SimpleContainerElement)
	}

	arch_choices := tuigo.Container(
		tuigo.Focusable,
		tuigo.HorizontalContainerTop,
		tuigo.Container(
			tuigo.Focusable,
			tuigo.VerticalContainer,
			cpuarch_choice_label,
			cpuarch_choice_elements,
		),
		tuigo.Container(
			tuigo.Focusable,
			tuigo.VerticalContainer,
			gpuarch_choice_label,
			gpuarch_choice_elements,
		),
	)

	elements = append(elements, cxx_path)
	elements = append(elements, cuda_path)
	elements = append(
		elements,
		tuigo.Container(
			tuigo.Focusable,
			tuigo.HorizontalContainerTop,
			cuda_choice,
			mpi_choice,
		),
	)
	elements = append(
		elements,
		tuigo.Container(
			tuigo.Focusable,
			tuigo.HorizontalContainerTop,
			debug_choice,
			output_choice,
		),
	)

	elements = append(elements, arch_choices)

	elements = append(elements, lib_elements...)

	comments := []tuigo.Component{}
	for label, sel := range global.Selectors {
		if (len(label) >= 5 && label[:5] == "HELP_" && !strings.Contains(label, "KEYS")) ||
			(len(label) >= 8 && label[:8] == "COMMENT_") {
			comments = append(
				comments,
				tuigo.Text(sel.Text(), tuigo.DimmedText),
			)
		}
	}
	comments = append(
		comments,
		tuigo.Text(
			global.Selectors["HELP_KEYS"].Text(),
			tuigo.DimmedText,
		),
	)
	return tuigo.Container(
		tuigo.Focusable,
		tuigo.VerticalContainer,
		main,
		tuigo.Container(
			tuigo.Focusable,
			tuigo.HorizontalContainer,
			tuigo.Container(
				tuigo.Focusable,
				tuigo.VerticalContainer,
				elements...,
			),
			tuigo.Container(
				tuigo.NonFocusable,
				tuigo.VerticalContainer,
				comments...,
			),
		),
	)
}

func MainUpdater(window tuigo.Window, msg tea.Msg) (tuigo.Window, tea.Cmd) {
	cmds := []tea.Cmd{}
	var action = func(id int, f func(tuigo.Wrapper, tuigo.Accessor) (tuigo.Wrapper, tuigo.Accessor)) {
		cmds = append(cmds, tuigo.TgtCmd(
			id,
			func(cont tuigo.Wrapper, mod tuigo.Accessor) (tuigo.Wrapper, tuigo.Accessor) {
				return f(cont, mod)
			},
		))
	}
	var hide = func(id int) {
		action(id, func(cont tuigo.Wrapper, mod tuigo.Accessor) (tuigo.Wrapper, tuigo.Accessor) {
			return cont.Hide().(tuigo.Wrapper), mod
		})
	}
	var unhide = func(id int) {
		action(id, func(cont tuigo.Wrapper, mod tuigo.Accessor) (tuigo.Wrapper, tuigo.Accessor) {
			return cont.Unhide().(tuigo.Wrapper), mod
		})
	}
	var visible = func(id int, flag bool) {
		if flag {
			unhide(id)
		} else {
			hide(id)
		}
	}
	var set_text = func(id int, text string) {
		cmds = append(
			cmds,
			tuigo.TgtCmd(
				id,
				func(cont tuigo.Wrapper, mod tuigo.Accessor) (tuigo.Wrapper, tuigo.Accessor) {
					return cont, mod.(tuigo.TextElement).Set(text)
				},
			),
		)
	}
	switch msg.(type) {
	case global.StateChange: // one of configs has been switched
		global.Selectors["MODE"].Read(window)
		install_mode := global.Selectors["MODE"].Value().([]string)
		if len(install_mode) == 0 {
			set_text(
				global.Selectors["MODE_TEXT"].Id(),
				global.Selectors["MODE_TEXT"].Def().(string),
			)
			for label, sel := range global.Selectors {
				if (len(label) < 4 || label[:4] != "MODE") &&
					!(len(label) >= 5 && label[:5] == "HELP_") &&
					!(len(label) >= 8 && label[:8] == "COMMENT_") {
					hide(sel.Id())
					if window.GetElementByID(sel.Id()+1) != nil {
						hide(sel.Id() + 1)
					}
					if window.GetElementByID(sel.Id()+2) != nil {
						hide(sel.Id() + 2)
					}
				}
			}
		} else {
			mode_entity := slices.Contains(install_mode, "Entity")
			mode_kokkos := slices.Contains(install_mode, "Kokkos")
			mode_adios2 := slices.Contains(install_mode, "ADIOS2")
			mode_mpi := slices.Contains(install_mode, "MPI")
			mode_hdf5 := slices.Contains(install_mode, "HDF5")
			mode_any := mode_entity || mode_kokkos || mode_adios2 || mode_mpi || mode_hdf5

			visible(global.Selectors["OUTPUT"].Id(), mode_entity)
			global.Selectors["OUTPUT"].Read(window)
			with_output := global.Selectors["OUTPUT"].Value().(bool)

			visible(global.Selectors["DEBUG"].Id(), mode_entity || mode_kokkos)
			visible(global.Selectors["CUDA"].Id(), mode_entity || mode_kokkos || mode_mpi)

			visible(global.Selectors["CXX_PATH"].Id(), mode_any)

			// cpuarch
			visible(global.Selectors["CPUARCH"].Id()+1, mode_entity || mode_kokkos)
			visible(global.Selectors["CPUARCH"].Id(), mode_entity || mode_kokkos)

			// cuda
			global.Selectors["CUDA"].Read(window)
			with_cuda := global.Selectors["CUDA"].Value().(bool)
			visible(
				global.Selectors["CUDA_PATH"].Id(),
				(mode_entity || mode_kokkos || mode_mpi) && with_cuda,
			)

			// gpuarch
			visible(
				global.Selectors["GPUARCH"].Id()+1,
				(mode_entity || mode_kokkos) && with_cuda,
			)
			visible(
				global.Selectors["GPUARCH"].Id(),
				(mode_entity || mode_kokkos) && with_cuda,
			)

			// kokkos
			visible(
				global.Selectors["KOKKOS_INSTALL_DIR"].Id(),
				mode_entity || mode_kokkos || mode_adios2,
			)
			visible(global.Selectors["KOKKOS_SRC_DIR"].Id(), mode_kokkos)
			visible(global.Selectors["KOKKOS_SRC_DIR"].Id()+1, mode_kokkos)
			visible(
				global.Selectors["KOKKOS_INSTALL_DIR"].Id()+1,
				mode_entity || mode_kokkos || mode_adios2,
			)
			visible(
				global.Selectors["KOKKOS_INSTALL_DIR"].Id()+2,
				mode_entity || mode_kokkos || mode_adios2,
			)

			// adios2
			visible(
				global.Selectors["ADIOS2_INSTALL_DIR"].Id()+2,
				(mode_entity && with_output) || mode_adios2,
			)
			visible(
				global.Selectors["ADIOS2_INSTALL_DIR"].Id(),
				(mode_entity && with_output) || mode_adios2,
			)
			visible(
				global.Selectors["ADIOS2_INSTALL_DIR"].Id()+1,
				(mode_entity && with_output) || mode_adios2,
			)
			visible(global.Selectors["ADIOS2_SRC_DIR"].Id(), mode_adios2)
			visible(global.Selectors["ADIOS2_SRC_DIR"].Id()+1, mode_adios2)

			// mpi
			visible(
				global.Selectors["MPI"].Id(),
				mode_entity || mode_adios2 || mode_hdf5,
			)
			global.Selectors["MPI"].Read(window)
			with_mpi := global.Selectors["MPI"].Value().(bool)
			visible(
				global.Selectors["MPI_INSTALL_DIR"].Id()+2,
				((mode_entity || mode_adios2 || mode_hdf5) && with_mpi) || mode_mpi,
			)
			visible(
				global.Selectors["MPI_INSTALL_DIR"].Id(),
				((mode_entity || mode_adios2 || mode_hdf5) && with_mpi) || mode_mpi,
			)
			visible(
				global.Selectors["MPI_INSTALL_DIR"].Id()+1,
				((mode_entity || mode_adios2 || mode_hdf5) && with_mpi) || mode_mpi,
			)
			visible(global.Selectors["MPI_SRC_DIR"].Id(), mode_mpi)
			visible(global.Selectors["MPI_SRC_DIR"].Id()+1, mode_mpi)

			// hdf5
			visible(
				global.Selectors["HDF5_INSTALL_DIR"].Id()+2,
				((mode_entity || mode_adios2) && with_output) || mode_hdf5,
			)
			visible(
				global.Selectors["HDF5_INSTALL_DIR"].Id(),
				((mode_entity || mode_adios2) && with_output) || mode_hdf5,
			)
			visible(
				global.Selectors["HDF5_INSTALL_DIR"].Id()+1,
				((mode_entity || mode_adios2) && with_output) || mode_hdf5,
			)
			visible(global.Selectors["HDF5_SRC_DIR"].Id(), mode_hdf5)
			visible(global.Selectors["HDF5_SRC_DIR"].Id()+1, mode_hdf5)

			set_text(
				global.Selectors["MODE_TEXT"].Id(),
				global.Selectors["MODE_TEXT"].Text(),
			)
		}
	}
	return window, tea.Batch(cmds...)
}
