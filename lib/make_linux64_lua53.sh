mkdir -p build_linux64 && cd build_linux64
cmake ../
cd ..
cmake --build build_linux64 --config Release
mkdir -p release/linux/
cp build_linux64/lua release/linux/
cp build_linux64/luac release/linux/
cp build_linux64/liblua.so release/linux/
# copy to your working dir
cp build_linux64/liblua.so ../cmd/lua/
