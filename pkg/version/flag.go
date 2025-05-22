// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.
//

// Package verflag defines utility functions to handle command line flags
// related to version of Kubernetes.
package version

import (
	"fmt"
	"os"
	"strconv"

	flag "github.com/spf13/pflag"
)

// versionValue 类型：实现了 flag.Value 接口
type versionValue int

// 通过实现 flag.Value 接口，将 -version 参数与 versionValue 类型绑定，支持以下输入：
// -version       → 等价于 -version=true（输出简化的版本信息）
// -version=true  → 输出简化版本信息
// -version=false → 不输出版本信息
// -version=raw   → 输出原始版本信息（如 raw 或具体版本字符串）
const (
	// 未设置版本.
	VersionNotSet versionValue = 0
	// 启用版本.
	VersionEnabled versionValue = 1
	// 原始版本.
	VersionRaw versionValue = 2
)

const strRawVersion string = "raw"

func (v *versionValue) IsBoolFlag() bool {
	return true
}

func (v *versionValue) Get() any {
	return *v
}

// Set 方法​​：解析输入字符串，支持 true/false 和 raw。
func (v *versionValue) Set(s string) error {
	if s == strRawVersion {
		*v = VersionRaw
		return nil
	}
	boolVal, err := strconv.ParseBool(s)
	if boolVal {
		*v = VersionEnabled
	} else {
		*v = VersionNotSet
	}
	return err
}

// ​​String 方法​​：返回当前状态的字符串表示（如 true、false 或 raw）。
func (v *versionValue) String() string {
	if *v == VersionRaw {
		return strRawVersion
	}
	return fmt.Sprintf("%v", bool(*v == VersionEnabled))
}

// The type of the flag as required by the pflag.Value interface.
func (v *versionValue) Type() string {
	return "version"
}

func VersionVar(p *versionValue, name string, value versionValue, usage string) {
	*p = value
	flag.Var(p, name, usage)
	// "--version" will be treated as "--version=true"
	flag.Lookup(name).NoOptDefVal = "true"
}

// Version 函数：创建一个新的版本标志，并将其添加到全局 FlagSet 中。
// 是 VersionVar 的便捷包装
func Version(name string, value versionValue, usage string) *versionValue {
	p := new(versionValue)
	VersionVar(p, name, value, usage)
	return p
}

const versionFlagName = "version"

var versionFlag = Version(versionFlagName, VersionNotSet, "Print version information and quit")

// AddFlags registers this package's flags on arbitrary FlagSets, such that they point to the
// same value as the global flags.
func AddFlags(fs *flag.FlagSet) {
	fs.AddFlag(flag.Lookup(versionFlagName))
}

// PrintAndExitIfRequested will check if the -version flag was passed
// and, if so, print the version and exit.
func PrintAndExitIfRequested() {
	// 检查版本标志的值并打印相应的信息
	if *versionFlag == VersionRaw {
		fmt.Printf("%s\n", Get().Text())
		os.Exit(0)
	} else if *versionFlag == VersionEnabled {
		fmt.Printf("%s\n", Get().String())
		os.Exit(0)
	}
}
