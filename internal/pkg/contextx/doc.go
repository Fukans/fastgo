// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/fastgo. The professional
// version of this repository is https://github.com/onexstack/onex.

/*
Package contextx 提供了对上下文（context）的扩展功能，允许在 context 中存储和提取用户相关的信息，如用户ID、用户名和访问令牌。

使用后缀 x 表示扩展或变体，使得包名简洁且易于记忆。本包中的函数方便了在上下文中传递和管理用户信息，适用于需要上下文传递数据的场景。

典型用法：
在处理 HTTP 请求的中间件或服务函数中，可以使用这些方法将用户信息存储到上下文中，以便在整个请求生命周期内安全地共享，避免使用全局变量和参数传参。

示例：

	// 创建新的上下文
	ctx := context.Background()

	// 将用户ID和用户名存放到上下文中
	ctx = contextx.WithUserID(ctx, "user-xxxx")
	ctx = contextx.WithUsername(ctx, "sampleUser")

	// 从上下文中提取用户信息
	userID := contextx.UserID(ctx)
	username := contextx.Username(ctx)
*/
package contextx // import "github.com/onexstack/fastgo/internal/pkg/contextx"

/*
// 使用空结构体作为上下文键（context key）有以下几个好处：
// 1. 内存效率
// 空结构体作为上下文键可以实现零内存占用，是一种高效的实现方式
1. **内存效率**
```go
// 空结构体不占用内存空间
type requestIDKey struct{}
```
- 空结构体的大小为 0 字节
- 所有空结构体实例共享同一个内存地址
- 相比使用 string 或 int 类型的键更节省内存

2. **类型安全**
```go
// 使用独特的类型作为键
ctx = context.WithValue(ctx, requestIDKey{}, requestID)
```
- 避免了使用字符串作为键可能产生的命名冲突
- 编译器可以检查类型错误
- 提供了更好的类型安全保证

3. **封装性**
```go
// 类型定义在包内部，对外不可见
type requestIDKey struct{}

// 通过函数暴露访问方法
func WithRequestID(ctx context.Context, requestID string) context.Context
func RequestID(ctx context.Context) string
```
- 键的定义对包外是私有的
- 只能通过包提供的函数访问上下文值
- 防止外部直接操作上下文键

4. **唯一性保证**
- 每个空结构体类型都是唯一的
- 即使结构体定义完全相同，不同的类型也是不同的键
- 避免了跨包键冲突的问题

5. **代码清晰度**
- 明确表明这个类型只用作键
- 提高代码可读性
- 体现了 Go 语言的设计哲学：显式优于隐式

这种模式是 Go 语言中处理 context 键的最佳实践，既保证了类型安全，又实现了高效的内存使用。
*/
