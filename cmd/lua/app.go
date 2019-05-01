package main
/*
#cgo CFLAGS: -Iinclude
#cgo LDFLAGS: -L../../prebuild -llua
#ifndef __glua_h__
#define __glua_h__
extern void NewLuaState(const char* code);
// example 2
#include <stdio.h>
#include <stdlib.h> // C.free
extern void SetFuncHandler(void*);
extern void GOexample2();
static const char* HelloGo(const char* msg) {
    printf("go lua string:%s\n", msg);
    fflush(stdout);
    GOexample2();
    return "ABCDEFG";
}
static void InitHandler() {
    SetFuncHandler(HelloGo);
}
#endif
*/
import "C"

import (
    "log"
    "unsafe"
)

func example1() {
    // example 1, invoke a mod
    code := C.CString("print(1234); exp.print()")
    C.NewLuaState(code)
    C.free(unsafe.Pointer(code))
}

func example2() {
    // example 2, lua size invoke golang function
    C.InitHandler()
    code := C.CString("local rs=exp.invoke('hello world');print(rs)")
    C.NewLuaState(code)
    C.free(unsafe.Pointer(code))
}

//export GOexample2
func GOexample2() {
    log.Println("run in go runtime.")
}

func main()  {
    example1();
    example2();
}


