package pages

import (
	templates "github.com/entity-toolkit/ntt-dploy/templates"
)

var Labels = map[string]string{
	"title": "Pick & configure the libraries to install",
	"review": `The code will install the selected modules
& generate compilation scripts for the selected
dependencies (if any). You will need to run the
generated scripts separately.`,
}

var Configs = map[string]*templates.LibraryConfiguration{
	"Entity": &templates.EntityTemplate,
	"Kokkos": &templates.KokkosTemplate,
	"ADIOS2": &templates.ADIOS2Template,
	"MPI":    &templates.MPITemplate,
	"HDF5":   &templates.HDF5Template,
}

var ModulePath = ""
