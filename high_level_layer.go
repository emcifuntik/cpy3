/*
Unless explicitly stated otherwise all files in this repository are licensed
under the MIT License.
This product includes software developed at Datadog (https://www.datadoghq.com/).
Copyright 2018 Datadog, Inc.
*/

package python3

/*
#cgo pkg-config: python3
#include "Python.h"
#include "macro.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// Py_Main : https://docs.python.org/3/c-api/veryhigh.html?highlight=pycompilerflags#c.Py_Main
// "error" will be set if we fail to call "Py_DecodeLocale" on every "args".
func Py_Main(args []string) (int, error) {
	argc := C.int(len(args))
	argv := make([]*C.wchar_t, argc, argc)
	for i, arg := range args {
		carg := C.CString(arg)
		defer C.free(unsafe.Pointer(carg))

		warg := C.Py_DecodeLocale(carg, nil)
		if warg == nil {
			return -1, fmt.Errorf("fail to call Py_DecodeLocale on '%s'", arg)
		}
		// Py_DecodeLocale requires a call to PyMem_RawFree to free the memory
		defer C.PyMem_RawFree(unsafe.Pointer(warg))
		argv[i] = warg
	}

	return int(C.Py_Main(argc, (**C.wchar_t)(unsafe.Pointer(&argv[0])))), nil
}

// PyRun_AnyFile : https://docs.python.org/3/c-api/veryhigh.html?highlight=pycompilerflags#c.PyRun_AnyFile
// "error" will be set if we fail to open "filename".
func PyRun_AnyFile(filename string) (int, error) {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))

	sstring := C.CString("s")
	defer C.free(unsafe.Pointer(sstring))

	pyobj := C._go_Py_BuildValue(sstring, cfilename)
	defer C.PyGC_Collect()

	mode := C.CString("r+")
	defer C.free(unsafe.Pointer(mode))

	cfile := C._Py_fopen_obj(pyobj, mode)

	// C.PyRun_AnyFile is a macro, using C.PyRun_AnyFileFlags instead
	return int(C.PyRun_AnyFileFlags(cfile, cfilename, nil)), nil
}

// PyRun_SimpleString : https://docs.python.org/3/c-api/veryhigh.html?highlight=pycompilerflags#c.PyRun_SimpleString
func PyRun_SimpleString(command string) int {
	ccommand := C.CString(command)
	defer C.free(unsafe.Pointer(ccommand))

	// C.PyRun_SimpleString is a macro, using C.PyRun_SimpleStringFlags instead
	return int(C.PyRun_SimpleStringFlags(ccommand, nil))
}

// PyRun_File : https://docs.python.org/3/c-api/veryhigh.html?highlight=pycompilerflags#c.PyRun_File
func PyRun_File(filename string, globals *PyObject, locals *PyObject) *PyObject {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))

	sstring := C.CString("s")
	defer C.free(unsafe.Pointer(sstring))

	pyobj := C._go_Py_BuildValue(sstring, cfilename)
	defer C.PyGC_Collect()

	mode := C.CString("r+")
	defer C.free(unsafe.Pointer(mode))

	cfile := C._Py_fopen_obj(pyobj, mode)

	closeit := 0

	return togo(C.PyRun_FileExFlags(cfile, cfilename, C.Py_file_input, toc(globals), toc(locals), C.int(closeit), nil))
}

// PyRun_SimpleFile
func PyRun_SimpleFile(filename string) int {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))

	sstring := C.CString("s")
	defer C.free(unsafe.Pointer(sstring))

	pyobj := C._go_Py_BuildValue(sstring, cfilename)
	defer C.PyGC_Collect()

	mode := C.CString("r+")

	cfile := C._Py_fopen_obj(pyobj, mode)
	defer C.free(unsafe.Pointer(mode))

	closeit := 0

	return int(C.PyRun_SimpleFileExFlags(cfile, cfilename, C.int(closeit), nil))
}
