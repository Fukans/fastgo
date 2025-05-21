// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/fastgo. The professional
// version of this repository is https://github.com/onexstack/onex.

package apiserver

import (
	"fmt"

	genericoptions "github.com/onexstack/fastgo/pkg/options"
)

// Config 配置结构体，用于存储应用相关的配置.
// 不用 viper.Get，是因为这种方式能更加清晰的知道应用提供了哪些配置项.
type Config struct {
	MySQLOptions *genericoptions.MySQLOptions
}

// Server 定义一个服务器结构体类型.
type Server struct {
	cfg *Config
}

// NewServer 根据配置创建服务器.
func (cfg *Config) NewServer() (*Server, error) {
	return &Server{cfg: cfg}, nil
}

// Run 运行应用.
func (s *Server) Run() error {
	fmt.Printf("Read MySQL host from config: %s\n", s.cfg.MySQLOptions.Addr)

	select {} // 调用 select 语句，阻塞防止进程退出
}

// 为什么要定义这个server？

// 定义 Server 结构体的重要原因:
//
// 1. 关注点分离
//    - 将配置（Config）和服务器实例（Server）分开
//    - 配置负责存储和管理配置项
//    - 服务器负责具体的业务逻辑实现
//
// 2. 依赖注入
//    - Server 通过持有 cfg *Config 字段获取所需的配置
//    - 这种设计使得配置可以在服务器创建时注入
//    - 便于测试时注入mock配置
//
// 3. 生命周期管理
//    - Server 结构体封装了服务器的生命周期
//    - NewServer 负责服务器的创建和初始化
//    - Run 方法负责启动和运行服务器
//
// 4. 扩展性考虑
//    - 当需要添加新的服务器功能时，可以直接在 Server 结构体中添加新的方法
//    - 需要新的配置项时，可以在 Config 中添加，不影响现有代码
//
// 这种设计模式是 Go 语言中常见的最佳实践，它提供了清晰的代码结构和良好的可维护性。

// 为什么server类不直接调用serverOptions类，而是调用Config类创建配置信息？

// 这种设计采用了中间层模式，主要有以下几个优点：
//
// 1. **职责分离**
//    - `ServerOptions` 负责配置的验证和解析，处理命令行参数和配置文件
//    - `Config` 负责配置的存储和管理，提供配置的访问接口
//    - `Server` 专注于业务逻辑的实现，不需要关心配置来源
//
// 2. **配置转换和清理**
//    - `ServerOptions` 到 `Config` 的转换过程可以：
//      - 清理和规范化配置数据
//      - 设置默认值
//      - 进行必要的配置转换
//    - 确保 `Server` 获得的是经过处理的、规范的配置
//
// 3. **解耦和灵活性**
//    - `Server` 不直接依赖 `ServerOptions`，降低了组件间的耦合
//    - 配置可以来自不同源（命令行、环境变量、配置文件）
//    - 方便未来扩展新的配置源或更改配置处理逻辑
//
// 4. **测试友好**
//    - 可以直接构造 `Config` 对象进行单元测试
//    - 不需要模拟完整的命令行环境
//    - 配置验证和服务逻辑可以分别测试
//
// 5. **维护性**
//    - 配置结构的变更只需要修改 `ServerOptions` 和 `Config`
//    - `Server` 的业务逻辑不会受到配置变更的影响
//    - 配置处理逻辑集中在一处，易于维护
//
// 这种设计模式在大型项目中特别有价值，它提供了更好的可维护性和扩展性，同时保持了代码的清晰结构。
