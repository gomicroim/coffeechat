#!/bin/sh

# 生成go pb文件
# --descriptor_set_out=pb.desc 生成描述文件，charles抓包需要
protoc --go_out=. $(find proto -iname "*.proto")

# 生成java pb文件
protoc --java_out=. $(find proto -iname "*.proto")