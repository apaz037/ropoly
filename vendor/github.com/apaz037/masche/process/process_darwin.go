package process

// #cgo CFLAGS: -std=c99
// #include <libproc.h>
// #include <errno.h>
// #include <stdlib.h>
import "C"

import (
	"fmt"
	"path/filepath"
	"reflect"
	"unsafe"
)

type darwinProcess int

func getProcess(pid int) darwinProcess {
	return darwinProcess(pid)
}

func (p process) Name() (name string, harderror error, softerrors []error) {
	cname := C.malloc(C.PROC_PIDPATHINFO_MAXSIZE)
	defer C.free(cname)

	_, err := C.proc_pidpath(C.int(p.pid), cname, C.PROC_PIDPATHINFO_MAXSIZE)
	if err != nil {
		harderr := fmt.Errorf("Error while reading name of process %d: %v", p.pid, err)
		return "", harderr, nil
	}

	name, harderror = filepath.EvalSymlinks(C.GoString((*C.char)(cname)))
	return
}

func (p darwinProcess) Close() (harderror error, softerrors []error) {
	return nil, nil
}

func (p darwinProcess) Handle() uintptr {
	return uintptr(p)
}

func getAllPids() (pids []int, harderror error, softerrors []error) {
	var pid C.pid_t
	pidSize := unsafe.Sizeof(pid)
	cpidsSize := pidSize * 1024 * 2
	cpids := C.malloc(C.size_t(cpidsSize))
	defer C.free(cpids)

	bytesUsed, err := C.proc_listpids(C.PROC_ALL_PIDS, 0, cpids, C.int(cpidsSize))
	if err != nil {
		return nil, err, nil
	}

	numberOfPids := uintptr(bytesUsed) / pidSize
	pids = make([]int, 0, numberOfPids)
	cpidsSlice := *(*[]C.pid_t)(unsafe.Pointer(
		&reflect.SliceHeader{
			Data: uintptr(unsafe.Pointer(cpids)),
			Len:  int(numberOfPids),
			Cap:  int(numberOfPids)}))

	for i, _ := range cpidsSlice {
		if cpidsSlice[i] == 0 {
			continue
		}

		pids = append(pids, int(cpidsSlice[i]))
	}

	return pids, nil, nil
}
