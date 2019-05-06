#include "../libs.h"

typedef void (*SetFuncHandler) (int, const char* msg);
typedef const char* (*GetFuncHandler)();
static SetFuncHandler setHandler = NULL;
static GetFuncHandler getHandler = NULL;
void SetGLUASetFuncHandler(SetFuncHandler h) {
    setHandler = h;
}
void SetGLUAGetFuncHandler(GetFuncHandler h) {
    getHandler = h;
}

static int g_set(lua_State *L) {
    if (setHandler != NULL) {
        int len = 0;
        const char* msg    = NULL;
        len = luaL_checkinteger(L, 1);
        msg = luaL_checkstring(L, 2);
        setHandler(len, msg);
        return 0;
    }
    return 0;
}
static int g_get(lua_State *L) {
    if (setHandler != NULL) {
        const char* result    = NULL;
        result = getHandler();
        lua_pushstring(L, result);
        return 1;
    }
    return 0;
}

static const luaL_Reg libs[] = {
    {"set", g_set},
    {"get", g_get},
    {NULL, NULL}
};

int luaopen_glua (lua_State *L) {
    luaL_newlib(L, libs);
    return 1;
}