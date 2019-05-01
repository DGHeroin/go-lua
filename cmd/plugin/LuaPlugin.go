package main
/*
#cgo CFLAGS: -Iinclude
#cgo LDFLAGS: -L../../prebuild -llua
#ifndef __glua_h__
#define __glua_h__
#include <stdlib.h> // C.free
#include <stdio.h>
extern void NewLuaState(const char* code);
extern void SetFuncHandler(void*);
#endif
*/
import "C"

import (
    "unsafe"
)

func NewLuaState(luaCode string) error {
    code := C.CString(luaCode)
    C.NewLuaState(code)
    C.free(unsafe.Pointer(code))
    return nil
}