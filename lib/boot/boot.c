#include <stdio.h>
#include "libs.h"

void NewLuaState(const char* bootCode) {
    lua_State *L = luaL_newstate();
    luaL_openlibs(L);
    OpenLibs(L);
    int ok = luaL_dostring(L, bootCode);
    if (ok != LUA_OK) {
        printf("Error: %s\n", lua_tostring(L,-1));
        fflush(stdout);
    }
}