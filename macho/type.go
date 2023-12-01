package macho

import (
	"debug/buildinfo"
	"debug/dwarf"
	mac "debug/macho"
	"runtime/debug"
)

// Macho is a struct representing a Mach-O file.
type macho struct {
	file *mac.File
	path string
}

// Open opens a Mach-O file specified by the given path.
func (m *macho) Open() (err error) {
	m.file, err = mac.Open(m.path)
	return err
}

func (m *macho) BuildInfo() (*debug.BuildInfo, error) {
	return buildinfo.ReadFile(m.path)
}

// Close closes the Mach-O file.
func (m *macho) Close() {

}

func (m *macho) Dwarf() (*dwarf.Data, error) {
	return m.file.DWARF()
}

func NewMacho(a string) *macho {
	return &macho{
		path: a,
	}
}
