mkdir -p build64_mingw && pushd build64_mingw
cmake -G "MinGW Makefiles" .. -DWITH_FEATURE_GP2P=ON
mingw32-make VERBOSE=1
popd
mkdir -p release/windows/
cp build64_mingw/lua.exe release/windows/
cp build64_mingw/luac.exe release/windows/
cp build64_mingw/liblua.dll release/windows/
# copy to your working dir
#cp build64_mingw/liblua.dll ../cmd/lua/
cp build64_mingw/liblua.dll ../prebuild/