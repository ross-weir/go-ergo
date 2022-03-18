package ergo

/*
#include "ergo.h"
*/
import "C"
import (
	"runtime"
	"unsafe"
)

type Tree interface {
	// Base16 converts the Tree to a base16 encoded string.
	Base16() (*string, error)

	// Address converts the Tree to an Address.
	Address() (Address, error)
}

type tree struct {
	p C.ErgoTreePtr
}

// NewTree creates a new ergo tree from the supplied base16 string.
func NewTree(s string) (Tree, error) {
	treeStr := C.CString(s)
	defer C.free(unsafe.Pointer(treeStr))

	var p C.ErgoTreePtr

	errPtr := C.ergo_lib_ergo_tree_from_base16_bytes(treeStr, &p)
	err := newError(errPtr)

	if err.isError() {
		return nil, err.error()
	}

	t := &tree{p}

	runtime.SetFinalizer(t, finalizeTree)

	return t, nil
}

func (t *tree) Base16() (*string, error) {
	var outStr *C.char
	defer C.free(unsafe.Pointer(outStr))

	errPtr := C.ergo_lib_ergo_tree_to_base16_bytes(t.p, &outStr)
	err := newError(errPtr)

	if err.isError() {
		return nil, err.error()
	}

	result := C.GoString(outStr)

	return &result, nil
}

func (t *tree) Address() (Address, error) {
	var p C.AddressPtr

	errPtr := C.ergo_lib_address_from_ergo_tree(t.p, &p)
	err := newError(errPtr)

	if err.isError() {
		return nil, err.error()
	}

	a := &address{p}

	return newAddress(a), nil
}

func finalizeTree(t *tree) {
	C.ergo_lib_ergo_tree_delete(t.p)
}
