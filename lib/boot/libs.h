#ifndef __libs_h__
#define __libs_h__
#include "lua.h"
#include "lualib.h"
#include "lauxlib.h"
void OpenLibs(lua_State *L);

// Example
#define LUA_LIBNAME_EXAMPLE "exp"
extern int luaopen_example(lua_State *L);

//GLUA
#define LUA_LIBNAME_GLUA "glua"
extern int luaopen_glua(lua_State *L);
#endif