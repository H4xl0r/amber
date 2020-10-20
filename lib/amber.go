package amber

import (
	"github.com/EgeBalci/debug/pe"

	"github.com/EgeBalci/keystone-go"
)

// VERSION number
const VERSION = "3.0.0"

// IMAGE_FILE_MACHINE_I386 Intel 386 or later processors and compatible processors
const IMAGE_FILE_MACHINE_I386 = 0x14c

// IMAGE_FILE_MACHINE_AMD64 x64
const IMAGE_FILE_MACHINE_AMD64 = 0x8664
const IMAGE_FILE_DLL = 0x2000

// Blueprint structure contains PE specs, tool parameters and
// OS spesific info
type Blueprint struct {
	// Parameters...
	FileName        string
	FileSize        int
	Reflective      bool
	IAT             bool
	Resource        bool
	ScrapePEHeaders bool
	IgnoreIntegrity bool
	EncodeCount     int

	// PE specs...
	Architecture      int
	SizeOfImage       uint32
	ImageBase         uint64
	AddressOfEntry    uint32
	Subsystem         uint16
	ImportTable       uint64
	ExportTable       uint64
	RelocTable        uint64
	ImportAdressTable uint64
	HasBoundedImports bool
	HasDelayedImports bool
	HasTLSCallbacks   bool
	HasRelocData      bool
	IsCLR             bool
	IsDLL             bool

	// PE File
	file *pe.File
}

// Assemble assembles the given instruction assembly code
func (bp *Blueprint) Assemble(asm string) ([]byte, bool) {
	var mode keystone.Mode
	switch bp.Architecture {
	case 32:
		mode = keystone.MODE_32
	case 64:
		mode = keystone.MODE_64
	default:
		return nil, false
	}

	ks, err := keystone.New(keystone.ARCH_X86, mode)
	if err != nil {
		return nil, false
	}
	defer ks.Close()

	//err = ks.Option(keystone.OPT_SYNTAX, keystone.OPT_SYNTAX_INTEL)
	err = ks.Option(keystone.OPT_SYNTAX, keystone.OPT_SYNTAX_NASM)
	if err != nil {
		return nil, false
	}
	//log.Println(asm)
	bin, _, ok := ks.Assemble(asm, 0)
	return bin, ok

}