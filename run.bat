@echo off
:: 获取输入的前缀和路径
set name=module
set path=api/%name%/%name%_v1/
set logPath=internal/logic/l_%name%/
set package=%name%_v1
set logicPackage=l_%name%

:: 检查指定路径是否存在，如果不存在则创建
if not exist "%path%" (
    echo The specified path does not exist. Creating directory...
    mkdir "%path%"
)

:: 创建 name_create.go 文件
echo package %package% > "%path%\%name%_create.go"

:: 创建 name_del.go 文件
echo package %package% > "%path%\%name%_del.go"

:: 创建name_modify.go文件
echo package %package% > "%path%\%name%_modify.go"

:: 创建name_list.go文件
echo package %package% > "%path%\%name%_list.go"

echo Files created successfully!


:: 检查指定路径是否存在，如果不存在则创建
if not exist "%logPath%" (
    echo The specified path does not exist. Creating directory...
    mkdir "%logPath%"
)

echo package %logicPackage% > "%logPath%\%name%.go"
:: 创建 name_create.go 文件
echo package %logicPackage% > "%logPath%\%name%_create.go"

:: 创建 name_del.go 文件
echo package %logicPackage% > "%logPath%\%name%_del.go"

:: 创建name_modify.go文件
echo package %logicPackage% > "%logPath%\%name%_modify.go"

:: 创建name_list.go文件
echo package %logicPackage% > "%logPath%\%name%_list.go"

echo Files created successfully!
