package lib

import (
	"debug/elf"
	"github.com/pkg/errors"
	"github.com/polyverse/ropoly/lib/architectures/amd64"
	"github.com/polyverse/ropoly/lib/gadgets"
	"github.com/polyverse/ropoly/lib/types"
)

func GadgetsFromExecutable(path string, maxLength int) (types.GadgetInstances, error, []error) {
	b, err := openBinary(path)
	if err != nil {
		return nil, errors.Wrapf(err, "Error opening ELF file %s", path), nil
	}
	defer b.close()

	allGadgets := []*types.GadgetInstance{}

	softerrs := []error{}
	sectionExists, addr, progData, err := b.nextSectionData()
	for sectionExists {
		if err != nil {
			return nil, err, nil
		}
		gadgetinstances, harderr, segment_softerrs := gadgets.Find(progData, amd64.GadgetSpecs, amd64.GadgetDecoder, addr, maxLength)
		softerrs = append(softerrs, segment_softerrs...)
		if harderr != nil {
			return nil, errors.Wrapf(err, "Unable to find gadgets from Program segment in the ELF file."), softerrs
		}
		allGadgets = append(allGadgets, gadgetinstances...)
		sectionExists, addr, progData, err = b.nextSectionData()
	}

	return allGadgets, nil, softerrs
}

type binary interface {
	close() error
	nextSectionData() (bool, types.Addr, []byte, error)
}

func openBinary(path string) (binary, error) {
	file, err := elf.Open(path)
	if err != nil {
		return nil, errors.Wrapf(err, "Error opening ELF file %s", path)
	}
	return elfBinary {
		binary: file,
		sectionIndex: new(int),
	}, nil
}

type elfBinary struct {
	binary *elf.File
	sectionIndex *int
}

func (b elfBinary) close() error {
	return b.binary.Close()
}

func (b elfBinary) nextSectionData() (bool, types.Addr, []byte, error) {
	if *b.sectionIndex == len(b.binary.Sections) {
		return false, 0, nil, nil
	}

	section := b.binary.Sections[*b.sectionIndex]
	*b.sectionIndex++

	if section.Type == elf.SHT_PROGBITS {
		progData, err := section.Data()
		return true, types.Addr(section.Addr), progData, err
	} else {
		return b.nextSectionData()
	}
}