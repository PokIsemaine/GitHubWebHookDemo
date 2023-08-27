#!/bin/bash

# 设置基准测试文件的路径
benchmark_file="./fib_test.go"

# 设置结果输出文件的路径
output_file="./output.json"

# 编译基准测试文件
go build -o benchmark "$benchmark_file"

# 执行基准测试并将结果保存为JSON格式
./benchmark -bench=. -json > "$output_file"