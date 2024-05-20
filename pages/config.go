package pages

import (
	"slices"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	sel "github.com/entity-toolkit/ntt-dploy/selectors"
	"github.com/haykh/tuigo"
)

func MainSelector(prev tuigo.Window) tuigo.Window {
	main := tuigo.Container(
		tuigo.Focusable,
		tuigo.VerticalContainer,
		tuigo.Text(
			Labels["title"],
			tuigo.LabelText,
		),
		tuigo.SelectorWithID(
			sel.Selectors["MODE"].Id(),
			sel.Selectors["MODE"].Options().([]string),
			tuigo.MultiSelect,
			tuigo.NoViewLimit,
			sel.Selectors["MODE"].Callback(),
		),
	)
	cpuarch_choice_label := tuigo.TextWithID(
		sel.Selectors["CPUARCH"].Id()+1,
		sel.Selectors["CPUARCH"].Text()+":",
		tuigo.NormalText,
	)
	cpuarch_choice_elements := tuigo.SelectorWithID(
		sel.Selectors["CPUARCH"].Id(),
		sel.Selectors["CPUARCH"].Options().([]string),
		tuigo.SelectOne,
		5,
		sel.Selectors["CPUARCH"].Callback(),
	)
	gpuarch_choice_label := tuigo.TextWithID(
		sel.Selectors["GPUARCH"].Id()+1,
		sel.Selectors["GPUARCH"].Text()+":",
		tuigo.NormalText,
	)
	gpuarch_choice_elements := tuigo.SelectorWithID(
		sel.Selectors["GPUARCH"].Id(),
		sel.Selectors["GPUARCH"].Options().([]string),
		tuigo.SelectOne,
		5,
		sel.Selectors["GPUARCH"].Callback(),
	)
	envmod_choice := tuigo.RadioWithID(
		sel.Selectors["ENV_MODULES"].Id(),
		sel.Selectors["ENV_MODULES"].Text(),
		sel.Selectors["ENV_MODULES"].Callback(),
	)
	if sel.Selectors["ENV_MODULES"].Def().(bool) {
		envmod_choice = envmod_choice.Set().(tuigo.SimpleContainerElement)
	}
	modfile_dir := tuigo.InputWithID(
		sel.Selectors["MODFILE_PATH"].Id(),
		sel.Selectors["MODFILE_PATH"].Text(),
		sel.Selectors["MODFILE_PATH"].Def().(string), "",
		tuigo.PathInput,
		tuigo.NoCallback,
	)
	debug_choice := tuigo.RadioWithID(
		sel.Selectors["DEBUG"].Id(),
		sel.Selectors["DEBUG"].Text(),
		tuigo.NoCallback,
	)
	cxx_path := tuigo.InputWithID(
		sel.Selectors["CXX_MOD"].Id(),
		sel.Selectors["CXX_MOD"].Text(),
		"",
		sel.Selectors["CXX_MOD"].Placeholder(),
		tuigo.TextInput,
		tuigo.NoCallback,
	)
	cuda_choice := tuigo.RadioWithID(
		sel.Selectors["CUDA"].Id(),
		sel.Selectors["CUDA"].Text(),
		sel.Selectors["CUDA"].Callback(),
	)
	cuda_path := tuigo.InputWithID(
		sel.Selectors["CUDA_MOD"].Id(),
		sel.Selectors["CUDA_MOD"].Text(),
		"",
		sel.Selectors["CUDA_MOD"].Placeholder(),
		tuigo.TextInput,
		tuigo.NoCallback,
	)
	if sel.Selectors["CUDA"].Def().(bool) {
		cuda_choice = cuda_choice.Set().(tuigo.SimpleContainerElement)
	}
	mpi_choice := tuigo.RadioWithID(
		sel.Selectors["MPI"].Id(),
		sel.Selectors["MPI"].Text(),
		sel.Selectors["MPI"].Callback(),
	)

	lib_names := []string{"KOKKOS", "ADIOS2", "MPI", "HDF5"}
	lib_elements := []tuigo.Component{}
	for _, name := range lib_names {
		src_dir_label := tuigo.TextWithID(
			sel.Selectors[name+"_SRC_DIR"].Id()+1,
			"source path:",
			tuigo.NormalText,
		)
		src_dir := tuigo.InputWithID(
			sel.Selectors[name+"_SRC_DIR"].Id(),
			"",
			sel.Selectors[name+"_SRC_DIR"].Def().(string),
			sel.Selectors[name+"_SRC_DIR"].Placeholder(),
			tuigo.PathInput,
			sel.Selectors[name+"_SRC_DIR"].Callback(),
		)
		install_dir_label := tuigo.TextWithID(
			sel.Selectors[name+"_INSTALL_DIR"].Id()+1,
			"install path:",
			tuigo.NormalText,
		)
		install_dir := tuigo.InputWithID(
			sel.Selectors[name+"_INSTALL_DIR"].Id(),
			"",
			sel.Selectors[name+"_INSTALL_DIR"].Def().(string),
			sel.Selectors[name+"_INSTALL_DIR"].Placeholder(),
			tuigo.TextInput,
			tuigo.NoCallback,
		)
		text_label := tuigo.TextWithID(
			sel.Selectors[name+"_INSTALL_DIR"].Id()+2,
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
		sel.Selectors["MODE_TEXT"].Id(),
		sel.Selectors["MODE_TEXT"].Def().(string),
		tuigo.NormalText,
	)
	elements = append(elements, packages)

	for _, el := range []*tuigo.SimpleContainerElement{
		&cuda_choice, &mpi_choice, &envmod_choice, &modfile_dir,
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
		debug_choice,
	)
	elements = append(elements, envmod_choice)
	elements = append(elements, modfile_dir)

	elements = append(elements, arch_choices)

	elements = append(elements, lib_elements...)

	comments := []tuigo.Component{}
	for label, sel := range sel.Selectors {
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
			sel.Selectors["HELP_KEYS"].Text(),
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
	case sel.StateChange: // one of configs has been switched
		sel.Selectors["MODE"].Read(window)
		install_mode := sel.Selectors["MODE"].Value().([]string)
		if len(install_mode) == 0 {
			set_text(
				sel.Selectors["MODE_TEXT"].Id(),
				sel.Selectors["MODE_TEXT"].Def().(string),
			)
			for label, sel := range sel.Selectors {
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

			strict_conds := map[string]bool{
				"ENV_MODULES": mode_entity,
				"MPI":         mode_mpi,
			}
			for elem, cond := range strict_conds {
				if cond {
					action(
						sel.Selectors[elem].Id(),
						func(cont tuigo.Wrapper, mod tuigo.Accessor) (tuigo.Wrapper, tuigo.Accessor) {
							if !mod.(tuigo.RadioElement).Data().(bool) {
								mod = mod.(tuigo.RadioElement).Toggle()
							}
							return cont.Disable().(tuigo.Wrapper), mod
						},
					)
				} else {
					action(
						sel.Selectors[elem].Id(),
						func(cont tuigo.Wrapper, mod tuigo.Accessor) (tuigo.Wrapper, tuigo.Accessor) {
							return cont.Enable().(tuigo.Wrapper), mod
						},
					)
				}
			}

			// if mode_entity {
			// 	action(
			// 		sel.Selectors["ENV_MODULES"].Id(),
			// 		func(cont tuigo.Wrapper, mod tuigo.Accessor) (tuigo.Wrapper, tuigo.Accessor) {
			// 			if !mod.(tuigo.RadioElement).Data().(bool) {
			// 				mod = mod.(tuigo.RadioElement).Toggle()
			// 			}
			// 			return cont.Disable().(tuigo.Wrapper), mod
			// 		},
			// 	)
			// }
			// if mode_adios2 {
			// 	action(
			// 		sel.Selectors["OUTPUT"].Id(),
			// 		func(cont tuigo.Wrapper, mod tuigo.Accessor) (tuigo.Wrapper, tuigo.Accessor) {
			// 			if !mod.(tuigo.RadioElement).Data().(bool) {
			// 				mod = mod.(tuigo.RadioElement).Toggle()
			// 			}
			// 			return cont.Disable().(tuigo.Wrapper), mod
			// 		},
			// 	)
			// }
			// if mode_mpi {
			// 	action(
			// 		sel.Selectors["MPI"].Id(),
			// 		func(cont tuigo.Wrapper, mod tuigo.Accessor) (tuigo.Wrapper, tuigo.Accessor) {
			// 			if !mod.(tuigo.RadioElement).Data().(bool) {
			// 				mod = mod.(tuigo.RadioElement).Toggle()
			// 			}
			// 			return cont.Disable().(tuigo.Wrapper), mod
			// 		},
			// 	)
			// }

			visible(sel.Selectors["ENV_MODULES"].Id(), mode_any)
			sel.Selectors["ENV_MODULES"].Read(window)
			with_envmod := sel.Selectors["ENV_MODULES"].Value().(bool)
			visible(sel.Selectors["MODFILE_PATH"].Id(), with_envmod)

			visible(sel.Selectors["DEBUG"].Id(), mode_entity || mode_kokkos)
			visible(sel.Selectors["CUDA"].Id(), mode_entity || mode_kokkos || mode_mpi)

			visible(sel.Selectors["CXX_MOD"].Id(), mode_any)

			// cpuarch
			visible(sel.Selectors["CPUARCH"].Id()+1, mode_entity || mode_kokkos)
			visible(sel.Selectors["CPUARCH"].Id(), mode_entity || mode_kokkos)

			// cuda
			sel.Selectors["CUDA"].Read(window)
			with_cuda := sel.Selectors["CUDA"].Value().(bool)
			visible(
				sel.Selectors["CUDA_MOD"].Id(),
				(mode_entity || mode_kokkos || mode_mpi) && with_cuda,
			)

			// gpuarch
			visible(
				sel.Selectors["GPUARCH"].Id()+1,
				(mode_entity || mode_kokkos) && with_cuda,
			)
			visible(
				sel.Selectors["GPUARCH"].Id(),
				(mode_entity || mode_kokkos) && with_cuda,
			)

			// kokkos
			visible(
				sel.Selectors["KOKKOS_INSTALL_DIR"].Id(),
				mode_entity || mode_kokkos || mode_adios2,
			)
			visible(sel.Selectors["KOKKOS_SRC_DIR"].Id(), mode_kokkos)
			visible(sel.Selectors["KOKKOS_SRC_DIR"].Id()+1, mode_kokkos)
			visible(
				sel.Selectors["KOKKOS_INSTALL_DIR"].Id()+1,
				mode_entity || mode_kokkos || mode_adios2,
			)
			visible(
				sel.Selectors["KOKKOS_INSTALL_DIR"].Id()+2,
				mode_entity || mode_kokkos || mode_adios2,
			)

			// adios2
			visible(
				sel.Selectors["ADIOS2_INSTALL_DIR"].Id()+2,
				mode_entity || mode_adios2,
			)
			visible(
				sel.Selectors["ADIOS2_INSTALL_DIR"].Id(),
				mode_entity || mode_adios2,
			)
			visible(
				sel.Selectors["ADIOS2_INSTALL_DIR"].Id()+1,
				mode_entity || mode_adios2,
			)
			visible(sel.Selectors["ADIOS2_SRC_DIR"].Id(), mode_adios2)
			visible(sel.Selectors["ADIOS2_SRC_DIR"].Id()+1, mode_adios2)

			// mpi
			visible(
				sel.Selectors["MPI"].Id(),
				mode_entity || mode_adios2 || mode_hdf5,
			)
			sel.Selectors["MPI"].Read(window)
			with_mpi := sel.Selectors["MPI"].Value().(bool)
			visible(
				sel.Selectors["MPI_INSTALL_DIR"].Id()+2,
				((mode_entity || mode_adios2 || mode_hdf5) && with_mpi) || mode_mpi,
			)
			visible(
				sel.Selectors["MPI_INSTALL_DIR"].Id(),
				((mode_entity || mode_adios2 || mode_hdf5) && with_mpi) || mode_mpi,
			)
			visible(
				sel.Selectors["MPI_INSTALL_DIR"].Id()+1,
				((mode_entity || mode_adios2 || mode_hdf5) && with_mpi) || mode_mpi,
			)
			visible(sel.Selectors["MPI_SRC_DIR"].Id(), mode_mpi)
			visible(sel.Selectors["MPI_SRC_DIR"].Id()+1, mode_mpi)

			// hdf5
			visible(
				sel.Selectors["HDF5_INSTALL_DIR"].Id()+2,
				(mode_entity || mode_adios2) || mode_hdf5,
			)
			visible(
				sel.Selectors["HDF5_INSTALL_DIR"].Id(),
				(mode_entity || mode_adios2) || mode_hdf5,
			)
			visible(
				sel.Selectors["HDF5_INSTALL_DIR"].Id()+1,
				(mode_entity || mode_adios2) || mode_hdf5,
			)
			visible(sel.Selectors["HDF5_SRC_DIR"].Id(), mode_hdf5)
			visible(sel.Selectors["HDF5_SRC_DIR"].Id()+1, mode_hdf5)

			set_text(
				sel.Selectors["MODE_TEXT"].Id(),
				sel.Selectors["MODE_TEXT"].Text(),
			)
		}
	}
	return window, tea.Batch(cmds...)
}
