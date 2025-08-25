# uaparser-go

[English](README.md) | 简体中文

[![Go Version](https://img.shields.io/badge/go-1.16+-blue.svg)](https://golang.org)
[![MIT License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/BaoziCDR/uaparser-go)](https://goreportcard.com/report/github.com/BaoziCDR/uaparser-go)
[![GoDoc](https://godoc.org/github.com/BaoziCDR/uaparser-go?status.svg)](https://godoc.org/github.com/BaoziCDR/uaparser-go)

> 🚀 快速、可靠的 Go 语言 User Agent 字符串解析器

一个高性能的 User Agent 解析库，支持自定义日志记录、缓存优化和线程安全操作。

**[查看文档](README.md)** · 
**[报告Bug](https://github.com/BaoziCDR/uaparser-go/issues)** · 
**[提出新特性](https://github.com/BaoziCDR/uaparser-go/issues)**

## 目录

- [uaparser-go](#uaparser-go)
  - [目录](#目录)
  - [项目介绍](#项目介绍)
    - [主要改进](#主要改进)
    - [设计目标](#设计目标)
  - [主要特性](#主要特性)
  - [快速开始](#快速开始)
    - [环境要求](#环境要求)
    - [安装步骤](#安装步骤)
    - [基本使用](#基本使用)
    - [版本比较功能](#版本比较功能)
    - [高级版本解析](#高级版本解析)
  - [项目结构](#项目结构)
  - [配置选项](#配置选项)
    - [基本配置](#基本配置)
    - [自定义日志记录](#自定义日志记录)
    - [性能优化选项](#性能优化选项)
  - [性能优化](#性能优化)
    - [缓存机制](#缓存机制)
    - [动态排序](#动态排序)
    - [内存对齐](#内存对齐)
  - [支持的浏览器](#支持的浏览器)
    - [桌面浏览器](#桌面浏览器)
    - [移动浏览器](#移动浏览器)
    - [特殊应用](#特殊应用)
  - [开发指南](#开发指南)
    - [运行测试](#运行测试)
    - [运行示例](#运行示例)
    - [性能测试](#性能测试)
  - [如何贡献](#如何贡献)
    - [贡献步骤](#贡献步骤)
    - [开发规范](#开发规范)
  - [版本历史](#版本历史)
  - [许可证](#许可证)
  - [致谢](#致谢)

## 项目介绍

UA Parser 是一个专为 Go 语言设计的高性能 User Agent 字符串解析库。它可以快速准确地从 User Agent 字符串中提取浏览器信息，包括浏览器类型和版本号。

本项目基于优秀的 [ua-parser/uap-go](https://github.com/ua-parser/uap-go) 项目构建，在版本解析能力和实用工具方面进行了重大增强。

### 主要改进

- **增强版本解析**: 支持任意位数的版本号解析（如 "1.2.3.4.5.6"）
- **版本比较工具**: 内置版本范围匹配和比较实用工具
- **性能优化**: 先进的缓存和动态排序机制
- **自定义日志**: 灵活的日志接口，便于调试和监控

### 设计目标

- **高性能**: 内置 LRU 缓存和动态排序优化
- **易于使用**: 简洁的 API 设计，开箱即用
- **高度可配置**: 支持自定义日志记录和性能调优
- **线程安全**: 支持高并发场景
- **扩展功能**: 版本比较和范围匹配实用工具

## 主要特性

- ✅ **快速解析**: 解析 User Agent 字符串提取浏览器类型和版本
- ✅ **增强版本支持**: 解析任意位数的版本号（如 "1.2.3.4.5.6"）
- ✅ **版本比较**: 内置版本比较和范围匹配实用工具
- ✅ **高性能缓存**: 内置 LRU 缓存机制，避免重复解析
- ✅ **自定义日志**: 支持任意实现 Logger 接口的日志库
- ✅ **动态优化**: 根据使用频率自动优化正则表达式顺序
- ✅ **线程安全**: 完全支持并发访问
- ✅ **广泛兼容**: 支持数百种浏览器和设备识别

## 快速开始

### 环境要求

- Go 1.16 或更高版本

### 安装步骤

```bash
go get github.com/BaoziCDR/uaparser-go
```

### 基本使用

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/BaoziCDR/uaparser-go"
)

func main() {
    // 创建解析器
    parser, err := uaparser.NewFromSaved()
    if err != nil {
        log.Fatal(err)
    }

    // 解析 User Agent 字符串
    ua := parser.Parse("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
    
    fmt.Printf("浏览器: %s\n", ua.Family)    // Chrome
    fmt.Printf("版本: %s\n", ua.Version)     // 91.0.4472.124
    fmt.Printf("完整信息: %s\n", ua.ToString()) // Chrome 91.0.4472.124
}
```

### 版本比较功能

库中包含强大的版本比较实用工具：

```go
package main

import (
    "fmt"
    "github.com/BaoziCDR/uaparser-go"
)

func main() {
    // 创建版本可比较对象
    version1 := uaparser.VersionComparable("91.0.4472.124")
    version2 := uaparser.VersionComparable("91.0.4472.125")
    
    // 比较版本
    result := version1.Compare(version2)
    if result < 0 {
        fmt.Println("version1 比 version2 旧")
    } else if result > 0 {
        fmt.Println("version1 比 version2 新")
    } else {
        fmt.Println("版本相同")
    }
    
    // 范围匹配
    chromeVersion := uaparser.VersionComparable("91.0.4472.124")
    
    // 检查版本是否在 [90.0.0.0, 92.0.0.0) 范围内
    inRange := uaparser.MatchRange("[90.0.0.0,92.0.0.0)", chromeVersion)
    fmt.Printf("Chrome 91.0.4472.124 在范围 [90.0.0.0, 92.0.0.0) 内: %v\n", inRange)
    
    // 检查版本是否大于 90.0.0.0
    isNewer := uaparser.MatchRange("(90.0.0.0,)", chromeVersion)
    fmt.Printf("Chrome 91.0.4472.124 比 90.0.0.0 新: %v\n", isNewer)
}
```

### 高级版本解析

与原始 ua-parser 实现不同，本库支持任意位数的版本号：

```go
package main

import (
    "fmt"
    "log"
    "github.com/BaoziCDR/uaparser-go"
)

func main() {
    parser, err := uaparser.NewFromSaved()
    if err != nil {
        log.Fatal(err)
    }

    // 解析复杂版本号
    testCases := []string{
        "CustomBrowser/1.2.3.4.5.6",
        "MyApp/10.15.7.21.1050.2.1",
        "Enterprise/2023.1.15.build.12345",
    }
    
    for _, ua := range testCases {
        result := parser.Parse(ua)
        fmt.Printf("UA: %s\n", ua)
        fmt.Printf("浏览器: %s, 版本: %s\n\n", result.Family, result.Version)
    }
}
```

## 项目结构

```
uaparser/
├── parser.go              # 核心解析器实现
├── user_agent.go         # User Agent 数据结构
├── logger.go             # 日志接口和实现
├── option.go             # 配置选项
├── cache.go              # 缓存实现
├── comparable.go         # 版本比较工具
├── defualt_yaml.go       # 内置正则表达式规则
├── test/                 # 测试文件
│   └── parser_test.go    # 单元测试
├── example/              # 使用示例
│   └── example.go        # 示例代码
├── README.md             # 英文文档
├── README_Zh.md          # 中文文档（本文件）
└── go.mod                # Go 模块定义
```

## 配置选项

### 基本配置

```go
parser, err := uaparser.NewFromSaved(
    uaparser.WithLogger(myLogger),           // 自定义日志器
    uaparser.WithDebugMode(true),            // 启用调试模式
    uaparser.WithUseSort(true),              // 启用排序优化
    uaparser.WithMissesThreshold(1000000),   // 缓存错失阈值
    uaparser.WithMatchIdxNotOk(25),          // 匹配索引阈值
)
```

### 自定义日志记录

```go
// 使用内置的默认日志器
logger := uaparser.NewDefaultLogger()

// 或使用自定义日志器
type MyLogger struct{}

func (l *MyLogger) Infof(format string, args ...interface{}) {
    // 实现您的日志逻辑
}

parser, _ := uaparser.NewFromSaved(uaparser.WithLogger(&MyLogger{}))
```

### 性能优化选项

| 选项 | 说明 | 默认值 | 推荐场景 |
|-----|------|--------|----------|
| `WithUseSort` | 启用动态排序优化 | `false` | 高并发场景 |
| `WithMissesThreshold` | 触发重排序的错失次数 | `500,000` | 根据并发量调整 |
| `WithMatchIdxNotOk` | 定义"不良匹配"的索引阈值 | `20` | 根据主要用户群体调整 |

## 性能优化

### 缓存机制

- **LRU 缓存**: 自动缓存解析结果，避免重复计算
- **缓存大小**: 默认 1024 条记录
- **命中率**: 在典型应用中可达 95% 以上

### 动态排序

```go
// 启用动态排序，常用的正则表达式会自动移到前面
parser, _ := uaparser.NewFromSaved(
    uaparser.WithUseSort(true),
    uaparser.WithMissesThreshold(100000), // 降低阈值以更频繁地优化
)
```

### 内存对齐

所有关键数据结构都经过内存对齐优化，在 64 位系统上性能最佳。

## 支持的浏览器

### 桌面浏览器
- Chrome、Firefox、Safari、Edge
- Internet Explorer、Opera
- 各种基于 Chromium 的浏览器

### 移动浏览器  
- Chrome Mobile、Firefox Mobile、Safari Mobile
- 微信内置浏览器、QQ 浏览器、UC 浏览器
- 百度浏览器、小米浏览器、华为浏览器

### 特殊应用
- Electron 应用、WebView
- 爬虫和机器人
- 各种 API 客户端

## 开发指南

### 运行测试

```bash
# 运行所有测试
go test ./test/ -v

# 运行基准测试
go test ./test/ -bench=. -benchmem
```

### 运行示例

```bash
# 运行基本示例
go run ./example/example.go
```

### 性能测试

项目包含完整的性能测试示例，展示了各种配置选项的效果。

## 如何贡献

贡献让开源社区成为学习、启发和创造的绝佳场所。您的任何贡献都是**非常感谢**的。

### 贡献步骤

1. **Fork** 本项目
2. 创建您的特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交您的更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启一个 **Pull Request**

### 开发规范

- 请确保代码通过所有测试
- 遵循 Go 语言代码规范
- 为新功能添加相应的测试用例
- 更新相关文档

## 版本历史

- **v1.0.0** - 初始版本发布
  - 基本的 User Agent 解析功能
  - 内置缓存机制
  - 自定义日志支持

查看 [版本发布](https://github.com/BaoziCDR/uaparser-go/releases) 获取详细的版本历史。

## 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 致谢

- **基于 [ua-parser/uap-go](https://github.com/ua-parser/uap-go)**: 本项目建立在官方 ua-parser Go 实现提供的优秀基础之上
- 感谢 [hashicorp/golang-lru](https://github.com/hashicorp/golang-lru) 提供的 LRU 缓存实现
- 感谢 [yaml.v2](https://gopkg.in/yaml.v2) 提供的 YAML 解析支持
- 感谢 [ua-parser 社区](https://github.com/ua-parser) 维护正则表达式定义
- 感谢所有为本项目做出贡献的开发者

---

**[⬆ 回到顶部](#uaparser-go)**

如果这个项目对您有帮助，请给它一个 ⭐️ 吧！
