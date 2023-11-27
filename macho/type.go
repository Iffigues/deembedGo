package macho

import "fmt"

// Macho is a struct representing a Mach-O file.
type Macho struct {
}

// Open opens a Mach-O file specified by the given path.
func (m *Macho) Open(path string) {
	fmt.Println("path = ", path)
}

// Close closes the Mach-O file.
func (m *Macho) Close() {

}
