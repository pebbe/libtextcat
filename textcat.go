/*
A wrapper for libtextcat

*/
package textcat

/*
#cgo LDFLAGS: -ltextcat
#include <textcat.h>
#include <stdlib.h>
#include <string.h>
void *tc_init(const char *configfile) {
    return textcat_Init(configfile);
}
void tc_done(void *h) {
    textcat_Done(h);
}
char *tc_classify(void *h, const char *buffer) {
    size_t size;
    size = strlen(buffer);
    return textcat_Classify(h, buffer, size);
}
*/
import "C"

import (
	"errors"
	"unsafe"
)

type Textcat struct {
    h      unsafe.Pointer
	isOpen bool
}

func NewTextcat(configfile string) (t *Textcat, e error) {
	t = &Textcat{}
	cs := C.CString(configfile)
	t.h = C.tc_init(cs)
	C.free(unsafe.Pointer(cs))
	if uintptr(t.h) == 0 {
		e = errors.New("Init textcat failed for config file \"" + configfile + "\"")
	} else {
		t.isOpen = true
	}
	return
}

func (t *Textcat) Classify(s string) string {
	if ! t.isOpen {
		panic("Textcat is closed")
	}
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	return C.GoString(C.tc_classify(t.h, cs))
}

func (t *Textcat) Close() {
	if t.isOpen {
		C.tc_done(t.h)
		t.isOpen = false
	}
}

