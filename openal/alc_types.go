//+build !windows

package openal

// #include "local.h"
import "C"

import "unsafe"

func (self *Device) cHandle() *C.struct_ALCdevice_struct {
	return (*C.struct_ALCdevice_struct)(unsafe.Pointer(self.handle))
}

func (self *Context) cHandle() *C.struct_ALCcontext_struct {
	return (*C.struct_ALCcontext_struct)(unsafe.Pointer(self.handle))
}
