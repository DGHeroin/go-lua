package main
/*
#cgo CFLAGS: -Iinclude
#cgo CFLAGS: -Wno-error=implicit-function-declaration
#cgo LDFLAGS: -L../../prebuild -llua -lm
#ifndef __glua_h__
#define __glua_h__
extern void NewLuaState(const char* code);
#include <stdio.h>
#include <stdlib.h> // C.free

extern void SetGLUASetFuncHandler(void*);
extern void SetGLUAGetFuncHandler(void*);

void Go_Set(int, void*);
void* Go_Get(int*);

static void Set(int len, const char* msg) {
    Go_Set(len, (void*)msg);
}
static const char* Get() {
    //return Go_Get();
    int len = 0;
    void* msg = Go_Get(&len);
    //printf("init ===> %s %d\n", msg, len);
    //printf("ptr->%p\n", msg);
    free(msg);
    return msg;
}
static void InitHandler() {
    SetGLUASetFuncHandler(Set);
    SetGLUAGetFuncHandler(Get);
}
#endif
*/
import "C"

import (
    "fmt"
    "io/ioutil"
    "os"
    "unsafe"
)

//export Go_Set
func Go_Set(sz C.int, msg unsafe.Pointer) {
    data := C.GoBytes(msg, sz)
    fmt.Println("Go_Set", sz, string(data))
}
//export Go_Get
func Go_Get(sz *C.int) unsafe.Pointer {
    content := "hello world!!!"
    *sz = C.int(len(content))
    ptr := unsafe.Pointer(C.CString(content))
    //fmt.Println("ptr:", ptr)
    return ptr
}

func run (luaCode string) error {
    // 注册函数
    C.InitHandler()
    code := C.CString(luaCode)
    C.NewLuaState(code)
    C.free(unsafe.Pointer(code))
    return nil
}

func main() {
    if len(os.Args) == 2 {
        filepath := os.Args[1]
        if _, err := os.Stat(filepath); os.IsNotExist(err) {
            // path/to/whatever does not exist
            run(filepath) // args as lua code
            return
        }
        data, err := ioutil.ReadFile(filepath)
        if err != nil {
            fmt.Println(err)
            return
        }
        run(string(data))
        return
    }
    // todo run as
    return
}


