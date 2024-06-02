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

var HowTo = map[string]string{
	"Kokkos": "Download the latest source code:\n`https://github.com/kokkos/kokkos/releases`",
	"ADIOS2": "Download the latest source code:\n`https://github.com/ornladios/ADIOS2/releases`",
	"MPI":    "",
	"HDF5": "2. Download the latest source codes as tar.xz:\n" +
		"   main source: `https://github.com/HDFGroup/hdf5/releases/`\n" +
		"   plugins: `https://github.com/HDFGroup/hdf5_plugins/releases`\n" +
		"   ZLIB: `https://github.com/madler/zlib/releases`\n" +
		"   LIBAEC: `https://github.com/MathisRosenhauer/libaec/releases`\n" +
		"3. Unpack only the main source code archive (leave plugins/zlib/libaec tar.xz files as is)\n" +
		"4. Copy `CTestScript.cmake`, `HDF5config.cmake`, and `HDF5options.cmake`\n" +
		"   from the source code directory `./config/cmake/scripts/` to the working directory",
}

var Configs = map[string]*templates.LibraryConfiguration{
	"Entity": &templates.EntityTemplate,
	"Kokkos": &templates.KokkosTemplate,
	"ADIOS2": &templates.ADIOS2Template,
	"MPI":    &templates.MPITemplate,
	"HDF5":   &templates.HDF5Template,
}

var ModulePath = ""
