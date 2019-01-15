mkdir -p build_osx && cd build_osx
cmake -GXcode ../
cd ..
cmake --build build_osx --config Release
mkdir -p release/osx/
cp build_osx/Release/lua release/osx/
cp build_osx/Release/luac release/osx/
cp build_osx/Release/liblua.dylib release/osx/
# copy to your working dir
cp build_osx/Release/liblua.dylib ../cmd/lua/
