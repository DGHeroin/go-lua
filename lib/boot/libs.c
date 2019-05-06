#include <stdio.h>
#include "libs.h"

static const luaL_Reg loadedlibs[] = {
  {LUA_LIBNAME_EXAMPLE, luaopen_example},
  {LUA_LIBNAME_GLUA, luaopen_glua},
  {NULL, NULL}
};

void OpenLibs(lua_State *L) {
    const luaL_Reg *lib;
    for (lib = loadedlibs; lib->func; lib++) {
        luaL_requiref(L, lib->name, lib->func, 1);
        lua_pop(L, 1);  /* remove lib */
    }
}