# uaparser-go

English | [简体中文](README_Zh.md)

[![Go Version](https://img.shields.io/badge/go-1.16+-blue.svg)](https://golang.org)
[![MIT License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/BaoziCDR/uaparser-go)](https://goreportcard.com/report/github.com/BaoziCDR/uaparser-go)
[![GoDoc](https://godoc.org/github.com/BaoziCDR/uaparser-go?status.svg)](https://godoc.org/github.com/BaoziCDR/uaparser-go)

> 🚀 A fast, reliable User Agent string parser written in Go

A high-performance User Agent parsing library with support for custom logging, caching optimization, and thread-safe operations.

**[View Documentation](README.md)** · 
**[Report Bug](https://github.com/BaoziCDR/uaparser-go/issues)** · 
**[Request Feature](https://github.com/BaoziCDR/uaparser-go/issues)**

## Table of Contents

- [uaparser-go](#uaparser-go)
  - [Table of Contents](#table-of-contents)
  - [About the Project](#about-the-project)
    - [Key Improvements](#key-improvements)
    - [Design Goals](#design-goals)
  - [Key Features](#key-features)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
    - [Quick Start](#quick-start)
    - [Version Comparison](#version-comparison)
    - [Advanced Version Parsing](#advanced-version-parsing)
  - [Project Structure](#project-structure)
  - [Configuration Options](#configuration-options)
    - [Basic Configuration](#basic-configuration)
    - [Custom Logging](#custom-logging)
    - [Performance Tuning Options](#performance-tuning-options)
    - [Available Options](#available-options)
    - [Built-in Loggers](#built-in-loggers)
      - [DefaultLogger](#defaultlogger)
      - [NoOpLogger](#nooplogger)
  - [Performance Optimization](#performance-optimization)
    - [Caching Mechanism](#caching-mechanism)
    - [Dynamic Sorting](#dynamic-sorting)
    - [Memory Alignment](#memory-alignment)
    - [Performance Features](#performance-features)
  - [Supported Browsers](#supported-browsers)
    - [Desktop Browsers](#desktop-browsers)
    - [Mobile Browsers](#mobile-browsers)
    - [Special Applications](#special-applications)
  - [Development Guide](#development-guide)
    - [Running Tests](#running-tests)
    - [Running Examples](#running-examples)
    - [Performance Testing](#performance-testing)
  - [Contributing](#contributing)
    - [How to Contribute](#how-to-contribute)
    - [Development Guidelines](#development-guidelines)
  - [Changelog](#changelog)
  - [License](#license)
  - [Acknowledgments](#acknowledgments)

## About the Project

UA Parser is a high-performance User Agent string parsing library designed specifically for Go. It can quickly and accurately extract browser information from User Agent strings, including browser type and version numbers.

This project is inspired by and builds upon the excellent work of [ua-parser/uap-go](https://github.com/ua-parser/uap-go), with significant enhancements in version parsing capabilities and additional utility features.

### Key Improvements

- **Enhanced Version Parsing**: Support for arbitrary-length version numbers (e.g., "1.2.3.4.5.6")
- **Version Comparison Tools**: Built-in utilities for version range matching and comparison
- **Performance Optimizations**: Advanced caching and dynamic sorting mechanisms
- **Custom Logging**: Flexible logging interface for debugging and monitoring

### Design Goals

- **High Performance**: Built-in LRU caching and dynamic sorting optimization
- **Easy to Use**: Clean API design, ready to use out of the box
- **Highly Configurable**: Support for custom logging and performance tuning
- **Thread Safe**: Support for high-concurrency scenarios
- **Extended Functionality**: Version comparison and range matching utilities

## Key Features

- ✅ **Fast Parsing**: Extract browser family and version from User Agent strings
- ✅ **Enhanced Version Support**: Parse arbitrary-length version numbers (e.g., "1.2.3.4.5.6")
- ✅ **Version Comparison**: Built-in version comparison and range matching utilities
- ✅ **High-Performance Caching**: Built-in LRU cache mechanism to avoid repeated parsing
- ✅ **Custom Logging**: Support for any logger implementing the Logger interface
- ✅ **Dynamic Optimization**: Automatically optimize regex order based on usage frequency
- ✅ **Thread Safe**: Full support for concurrent access
- ✅ **Wide Compatibility**: Support for hundreds of browsers and device identification

## Getting Started

### Prerequisites

- Go 1.16 or higher

### Installation

```bash
go get github.com/BaoziCDR/uaparser-go
```

### Quick Start

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/BaoziCDR/uaparser-go"
)

func main() {
    // Create a new parser
    parser, err := uaparser.NewFromSaved()
    if err != nil {
        log.Fatal(err)
    }

    // Parse a user agent string
    ua := parser.Parse("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
    
    fmt.Printf("Family: %s\n", ua.Family)     // Chrome
    fmt.Printf("Version: %s\n", ua.Version)   // 91.0.4472.124
    fmt.Printf("String: %s\n", ua.ToString()) // Chrome 91.0.4472.124
}
```

### Version Comparison

The library includes powerful version comparison utilities:

```go
package main

import (
    "fmt"
    "github.com/BaoziCDR/uaparser-go"
)

func main() {
    // Create version comparables
    version1 := uaparser.VersionComparable("91.0.4472.124")
    version2 := uaparser.VersionComparable("91.0.4472.125")
    
    // Compare versions
    result := version1.Compare(version2)
    if result < 0 {
        fmt.Println("version1 is older than version2")
    } else if result > 0 {
        fmt.Println("version1 is newer than version2")
    } else {
        fmt.Println("versions are equal")
    }
    
    // Range matching
    chromeVersion := uaparser.VersionComparable("91.0.4472.124")
    
    // Check if version is in range [90.0.0.0, 92.0.0.0)
    inRange := uaparser.MatchRange("[90.0.0.0,92.0.0.0)", chromeVersion)
    fmt.Printf("Chrome 91.0.4472.124 is in range [90.0.0.0, 92.0.0.0): %v\n", inRange)
    
    // Check if version is greater than 90.0.0.0
    isNewer := uaparser.MatchRange("(90.0.0.0,)", chromeVersion)
    fmt.Printf("Chrome 91.0.4472.124 is newer than 90.0.0.0: %v\n", isNewer)
}
```

### Advanced Version Parsing

Unlike the original ua-parser implementation, this library supports arbitrary-length version numbers:

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

    // Parse complex version numbers
    testCases := []string{
        "CustomBrowser/1.2.3.4.5.6",
        "MyApp/10.15.7.21.1050.2.1",
        "Enterprise/2023.1.15.build.12345",
    }
    
    for _, ua := range testCases {
        result := parser.Parse(ua)
        fmt.Printf("UA: %s\n", ua)
        fmt.Printf("Family: %s, Version: %s\n\n", result.Family, result.Version)
    }
}
```

## Project Structure

```
uaparser/
├── parser.go            # Core parser implementation
├── user_agent.go        # User agent structures
├── logger.go            # Logger interface and implementations
├── option.go            # Configuration options  
├── cache.go             # Caching implementation
├── comparable.go        # Version comparison utilities
├── defualt_yaml.go      # Built-in regex definitions
├── test/                # Test files
│   └── parser_test.go   # Unit tests
├── example/             # Usage examples
│   └── example.go       # Example implementations
├── README.md            # English Documentation
├── README_Zh.md         # Chinese Documentation
└── go.mod               # Go Module Definition
```

## Configuration Options

### Basic Configuration

```go
parser, err := uaparser.NewFromSaved(
    uaparser.WithLogger(myLogger),           // Custom logger
    uaparser.WithDebugMode(true),            // Enable debug logging
    uaparser.WithUseSort(true),              // Enable sorting optimization
    uaparser.WithMissesThreshold(1000000),   // Cache miss threshold
    uaparser.WithMatchIdxNotOk(25),          // Match index threshold
)
```

### Custom Logging

```go
// Use built-in default logger
logger := uaparser.NewDefaultLogger()

// Or use a custom logger
type MyLogger struct{}

func (l *MyLogger) Infof(format string, args ...interface{}) {
    // Implement your logging logic
}

parser, _ := uaparser.NewFromSaved(uaparser.WithLogger(&MyLogger{}))
```

### Performance Tuning Options

| Option | Description | Default | Recommended Scenario |
|--------|-------------|---------|----------------------|
| `WithUseSort` | Enable dynamic sorting optimization | `false` | High-concurrency scenarios |
| `WithMissesThreshold` | Number of misses to trigger reordering | `500,000` | Adjust based on concurrency level |
| `WithMatchIdxNotOk` | Index threshold for "poor matches" | `20` | Adjust based on main user base |

### Available Options

- `WithLogger(logger Logger)` - Set a custom logger
- `WithDebugMode(bool)` - Enable/disable debug logging
- `WithUseSort(bool)` - Enable/disable automatic sorting of regex patterns by usage
- `WithMissesThreshold(uint64)` - Set the threshold for triggering pattern sorting
- `WithMatchIdxNotOk(int)` - Set the index threshold for counting cache misses

### Built-in Loggers

#### DefaultLogger
Logs to stdout with timestamps:
```go
logger := uaparser.NewDefaultLogger()
```

#### NoOpLogger
Disables all logging (default):
```go
logger := uaparser.NewNoOpLogger()
```

## Performance Optimization

### Caching Mechanism

- **LRU Cache**: Automatically cache parsing results to avoid repeated calculations
- **Cache Size**: Default 1024 records
- **Hit Rate**: Can reach 95%+ in typical applications

### Dynamic Sorting

```go
// Enable dynamic sorting, frequently used regex patterns will automatically move to the front
parser, _ := uaparser.NewFromSaved(
    uaparser.WithUseSort(true),
    uaparser.WithMissesThreshold(100000), // Lower threshold for more frequent optimization
)
```

### Memory Alignment

All critical data structures are optimized for memory alignment, delivering best performance on 64-bit systems.

### Performance Features

- **Caching**: Parsed results are cached using LRU cache
- **Pattern Sorting**: Frequently matched patterns are moved to the front
- **Atomic Operations**: Thread-safe counters for statistics
- **Memory Alignment**: Optimized struct layout for better performance

## Supported Browsers

### Desktop Browsers
- Chrome, Firefox, Safari, Edge
- Internet Explorer, Opera
- Various Chromium-based browsers

### Mobile Browsers  
- Chrome Mobile, Firefox Mobile, Safari Mobile
- WeChat Browser, QQ Browser, UC Browser
- Baidu Browser, Mi Browser, Huawei Browser

### Special Applications
- Electron applications, WebView
- Crawlers and bots
- Various API clients

## Development Guide

### Running Tests

```bash
# Run all tests
go test ./test/ -v

# Run benchmark tests
go test ./test/ -bench=. -benchmem
```

### Running Examples

```bash
# Run basic example
go run ./example/example.go
```

### Performance Testing

The project includes comprehensive performance testing examples showcasing the effects of various configuration options.

## Contributing

Contributions make the open source community an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

### How to Contribute

1. **Fork** the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a **Pull Request**

### Development Guidelines

- Ensure code passes all tests
- Follow Go language coding standards
- Add appropriate test cases for new features
- Update relevant documentation

## Changelog

- **v1.0.0** - Initial release
  - Basic User Agent parsing functionality
  - Built-in caching mechanism
  - Custom logging support

See [Releases](https://github.com/BaoziCDR/uaparser-go/releases) for detailed changelog.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- **Based on [ua-parser/uap-go](https://github.com/ua-parser/uap-go)**: This project builds upon the excellent foundation provided by the official ua-parser Go implementation
- Thanks to [hashicorp/golang-lru](https://github.com/hashicorp/golang-lru) for LRU cache implementation
- Thanks to [yaml.v2](https://gopkg.in/yaml.v2) for YAML parsing support
- Thanks to the [ua-parser community](https://github.com/ua-parser) for maintaining the regex definitions
- Thanks to all developers who contributed to this project

---

**[⬆ Back to top](#uaparser-go)**

If this project helps you, please give it a ⭐️!
