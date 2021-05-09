//+build windows

package openal

// #include "local.h"
import "C"

import "unsafe"

func (self *Device) cHandle() *C.struct_ALCdevice {
	return (*C.struct_ALCdevice)(unsafe.Pointer(self.handle))
}

func (self *Context) cHandle() *C.struct_ALCcontext {
	return (*C.struct_ALCcontext)(unsafe.Pointer(self.handle))
}
