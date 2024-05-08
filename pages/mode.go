package pages

import (
	"github.com/entity-toolkit/ntt-dploy/configs"
	"github.com/entity-toolkit/ntt-dploy/global"
	"github.com/haykh/tuigo"
)

func ModeSelector(tuigo.Window) tuigo.Window {
	return tuigo.Container(
		configs.FOCUSABLE,
		tuigo.VerticalContainer,
		tuigo.Text(
			global.MODE.Text,
			tuigo.LabelText,
		),
		tuigo.SelectorWithID(
			global.MODE.Id,
			[]string{"Entity", "Kokkos", "ADIOS2"},
			configs.SELECT_ONE,
			configs.NO_CALLBACK,
		),
		tuigo.Text(
			global.COMMENTS["modules"],
			tuigo.DimmedText,
		),
	)
}
