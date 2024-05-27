#!/usr/bin/env bash

RUN_NAME="ProjGPT"

# 创建输出目录
mkdir -p build/bin

# 定义一个函数来编译特定平台的二进制文件
compile_for_platform() {
    local os=$1
    local arch=$2
    local output_name="${RUN_NAME}_${os}_${arch}"

    # 编译二进制文件
    GOOS=$os GOARCH=$arch go build -ldflags="-s -w" -o build/bin/${output_name} ./src/main.go && cmd/upx -9 build/bin/${output_name}
    echo "${output_name} build success"
}

# 编译Linux的64位和32位版本
compile_for_platform linux amd64
compile_for_platform linux 386

# 编译Windows的64位和32位版本
compile_for_platform windows amd64
compile_for_platform windows 386

# 编译macOS的64位版本
compile_for_platform darwin amd64

# 编译完成后的提示
echo "All platforms build success"