#include "../libs.h"

// export a SetFuncHandler to global and golang size set a callback
typedef char* (*FuncHandler)(const char* msg);
static FuncHandler handler = NULL;

void SetFuncHandler(FuncHandler h) {
    handler = h;
}

static int g_invoke(lua_State *L) {
    if (handler != NULL) {
        const char* msg    = NULL;
        const char* result = NULL;
        msg = luaL_checkstring(L, 1);
        printf("%s, %p\n", msg, msg);
        result = handler(msg);
        lua_pushstring(L, result);
        return 1;
    } else {
        printf("handler not set.\n");
    }
    return 0;
}

static int g_print(lua_State *L) {
    printf("call example print func.\n");
    return 0;
}

static const luaL_Reg libs[] = {
    {"print", g_print},
    {"invoke", g_invoke},
    {NULL, NULL}
};

int luaopen_example (lua_State *L) {
    luaL_newlib(L, libs);
    return 1;
}