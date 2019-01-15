package main
/*
#cgo CFLAGS: -Iinclude
#cgo LDFLAGS: -L./ -llua
#ifndef __glua_h__
#define __glua_h__
extern void NewLuaState(const char* BootCode);
// example 2
#include <stdio.h>
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
import "log"

func example1() {
    // example 1, invoke a mod
    C.NewLuaState(C.CString("print(1234); exp.print()"))
}

func example2() {
    // example 2, lua size invoke golang function
    C.InitHandler()
    C.NewLuaState(C.CString("local rs=exp.invoke('hello world');print(rs)"))
}

//export GOexample2
func GOexample2() {
    log.Println("run in go runtime.")
}

func main()  {
    example1();
    example2();
}

