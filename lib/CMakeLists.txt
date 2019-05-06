cmake_minimum_required(VERSION 2.8)
project (lua)  # project here actually means solution in premake

if(WIN32)
    add_definitions( -D_CRT_SECURE_NO_WARNINGS )
endif()

if ( WIN32 AND NOT CYGWIN AND NOT ( CMAKE_SYSTEM_NAME STREQUAL "WindowsStore" ) )
    set(CMAKE_C_FLAGS_RELEASE "${CMAKE_C_FLAGS_RELEASE} /MT" CACHE STRING "")
    set(CMAKE_C_FLAGS_DEBUG "${CMAKE_C_FLAGS_DEBUG} /MTd" CACHE STRING "")
    set(CMAKE_CXX_FLAGS_RELEASE "${CMAKE_CXX_FLAGS_RELEASE} /MT" CACHE STRING "")
    set(CMAKE_CXX_FLAGS_DEBUG "${CMAKE_CXX_FLAGS_DEBUG} /MTd" CACHE STRING "")
endif ()

macro(source_group_by_dir proj_dir source_files)
    if(MSVC OR APPLE)
        get_filename_component(sgbd_cur_dir ${proj_dir} ABSOLUTE)
        foreach(sgbd_file ${${source_files}})
            get_filename_component(sgbd_abs_file ${sgbd_file} ABSOLUTE)
            file(RELATIVE_PATH sgbd_fpath ${sgbd_cur_dir} ${sgbd_abs_file})
            string(REGEX REPLACE "\(.*\)/.*" \\1 sgbd_group_name ${sgbd_fpath})
            string(COMPARE EQUAL ${sgbd_fpath} ${sgbd_group_name} sgbd_nogroup)
            string(REPLACE "/" "\\" sgbd_group_name ${sgbd_group_name})
            if(sgbd_nogroup)
                set(sgbd_group_name "\\")
            endif(sgbd_nogroup)
            source_group(${sgbd_group_name} FILES ${sgbd_file})
        endforeach(sgbd_file)
    endif(MSVC OR APPLE)
endmacro(source_group_by_dir)

macro(print_all_variables)
    message(STATUS "print_all_variables------------------------------------------{")
    get_cmake_property(_variableNames VARIABLES)
    foreach (_variableName ${_variableNames})
        message(STATUS "${_variableName}=${${_variableName}}")
    endforeach()
    message(STATUS "print_all_variables------------------------------------------}")
endmacro()

# 1. lua static library
include_directories(
        ${CMAKE_SOURCE_DIR}
        lua-5.3.5/src
        ${CMAKE_CURRENT_BINARY_DIR}
        ${THIRDPART_INC}
    )

set ( LUA_CORE lua-5.3.5/src/lapi.c lua-5.3.5/src/lcode.c lua-5.3.5/src/lctype.c lua-5.3.5/src/ldebug.c lua-5.3.5/src/ldo.c lua-5.3.5/src/ldump.c
      lua-5.3.5/src/lfunc.c lua-5.3.5/src/lgc.c lua-5.3.5/src/llex.c lua-5.3.5/src/lmem.c lua-5.3.5/src/lobject.c lua-5.3.5/src/lopcodes.c lua-5.3.5/src/lparser.c
      lua-5.3.5/src/lstate.c lua-5.3.5/src/lstring.c lua-5.3.5/src/ltable.c lua-5.3.5/src/ltm.c lua-5.3.5/src/lundump.c lua-5.3.5/src/lvm.c lua-5.3.5/src/lzio.c )
set ( LUA_LIB lua-5.3.5/src/lauxlib.c lua-5.3.5/src/lbaselib.c lua-5.3.5/src/lbitlib.c lua-5.3.5/src/lcorolib.c lua-5.3.5/src/ldblib.c
  lua-5.3.5/src/liolib.c lua-5.3.5/src/lmathlib.c lua-5.3.5/src/loslib.c lua-5.3.5/src/lstrlib.c lua-5.3.5/src/ltablib.c lua-5.3.5/src/linit.c
  lua-5.3.5/src/lutf8lib.c lua-5.3.5/src/loadlib.c )

set ( LUA_LIBBOOT boot/boot.c boot/libs.c )
set ( LUA_LIBBOOT_MODS "")
set ( MOD_GLUA boot/mod/glua.c )
set ( MOD_EXAMPLE boot/mod/example.c )
set (
     LUA_LIBBOOT_MODS
     ${LUA_LIBBOOT_MODS}
     ${MOD_GLUA}
     ${MOD_EXAMPLE}
)

set ( LUA_EXT_LIB_DIR  "" )

source_group_by_dir(${CMAKE_CURRENT_SOURCE_DIR} LUA_CORE)
source_group_by_dir(${CMAKE_CURRENT_SOURCE_DIR} LUA_LIB)

add_library (lualib STATIC
            ${LUA_CORE}
            ${LUA_LIB}
            )
# rename the library output name
set_target_properties(lualib PROPERTIES OUTPUT_NAME "lua")

# 1. liblua shared
add_library (luasharedlib SHARED
            ${LUA_CORE}
            ${LUA_LIB}
            ${LUA_LIBBOOT}
            ${LUA_LIBBOOT_MODS}
            )
# rename the library output name
set_target_properties(luasharedlib PROPERTIES OUTPUT_NAME "lua")
# 2. lua
add_executable (lua lua-5.3.5/src/lua.c)
target_link_libraries (lua lualib)
if(UNIX)
        target_link_libraries( lua m )
endif()

# 3. luac
link_directories (${LUA_BINARY_DIR})
add_executable (luac lua-5.3.5/src/luac.c)
target_link_libraries (luac lualib)
if(UNIX)
        target_link_libraries( luac m )
endif()

#print_all_variables()