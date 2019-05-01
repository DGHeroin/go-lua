mkdir -p build_osx && cd build_osx
cmake -GXcode ../
cd ..
cmake --build build_osx --config Release

install_name_tool -id "@executable_path/liblua.dylib" build_osx/Release/liblua.dylib

mkdir -p release/osx/
cp build_osx/Release/lua release/osx/
cp build_osx/Release/luac release/osx/
cp build_osx/Release/liblua.dylib release/osx/

# copy to your working dir
#cp build_osx/Release/liblua.dylib ../cmd/lua/
cp build_osx/Release/liblua.dylib ../prebuild

