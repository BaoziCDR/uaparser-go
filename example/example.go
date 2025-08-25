package main

import (
	"fmt"
	"log"
	"time"

	"github.com/BaoziCDR/uaparser-go"
)

func main() {
	fmt.Println("=== Basic Usage ===")
	Example()

	fmt.Println("\n=== Custom Logger ===")
	ExampleWithCustomLogger()

	fmt.Println("\n=== Performance Optimization Options ===")
	ExamplePerformanceOptions()
}

// Example shows basic usage of the uaparser library
func Example() {
	// Create a parser with default settings (no logging)
	parser, err := uaparser.NewFromSaved()
	if err != nil {
		log.Fatal(err)
	}

	// Parse a user agent string
	ua := parser.Parse("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	fmt.Printf("Family: %s, Version: %s\n", ua.Family, ua.Version)
	// Output: Family: Chrome, Version: 91.0.4472.124
}

// ExampleWithCustomLogger shows how to use a custom logger
func ExampleWithCustomLogger() {
	// Create a custom logger
	customLogger := uaparser.NewDefaultLogger()

	// Create a parser with custom logger and debug mode
	parser, err := uaparser.NewFromSaved(
		uaparser.WithLogger(customLogger),
		uaparser.WithDebugMode(true),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Parse a user agent string
	ua := parser.Parse("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	fmt.Printf("Family: %s, Version: %s\n", ua.Family, ua.Version)
}

// ExamplePerformanceOptions demonstrates the advanced performance tuning options
// available in the uaparser library. These options are designed to optimize
// parsing performance in high-throughput scenarios.
//
// Performance Configuration Options:
//
//  1. WithUseSort(true) - Dynamic Sorting Optimization
//     Enables automatic reordering of regex patterns based on match frequency.
//     Each regex pattern maintains a MatchesCount counter, and sorting is performed
//     in descending order of usage frequency. This significantly improves performance
//     when processing large volumes of similar User-Agent types.
//
//  2. WithMatchIdxNotOk(threshold) - Match Index Threshold
//     Defines what constitutes a "poor match" position (default: 20).
//     When a regex pattern is matched at an index greater than this threshold,
//     the UserAgentMisses counter is incremented. This indicates that frequently
//     used patterns are located too far down in the array and reordering is needed.
//     Recommended: 5-10 for applications primarily handling common browsers.
//
//  3. WithMissesThreshold(count) - Sorting Trigger Threshold
//     Sets the number of "poor matches" that triggers automatic reordering
//     (default: 500,000, minimum: 100,000). When UserAgentMisses >= missesThreshold,
//     the parser executes sort.Sort() to reorder patterns by frequency.
//     Recommended: 100,000-200,000 for high-concurrency applications.
//
// Configuration Recommendations:
//   - Low concurrency: Use default configuration
//   - High concurrency: Enable WithUseSort(true) with smaller missesThreshold
//   - Mobile-focused: Set smaller matchIdxNotOk (5-10)
//   - Diverse UA types: Increase missesThreshold to reduce sorting frequency
func ExamplePerformanceOptions() {
	logger := uaparser.NewDefaultLogger()

	// Create a parser with performance optimization settings
	parser, err := uaparser.NewFromSaved(
		uaparser.WithLogger(logger),
		uaparser.WithDebugMode(true),
		uaparser.WithUseSort(true),           // Enable dynamic sorting optimization
		uaparser.WithMatchIdxNotOk(10),       // Consider matches at index >10 as poor
		uaparser.WithMissesThreshold(100000), // Trigger reordering after 100K poor matches
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("=== Performance Optimization Demo ===")

	// Test with common browsers (should match early in the array)
	testUAs := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.1 Safari/605.1.15",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:91.0) Gecko/20100101 Firefox/91.0",
	}

	start := time.Now()
	for i, ua := range testUAs {
		result := parser.Parse(ua)
		fmt.Printf("Test %d: %s -> %s\n", i+1, result.Family, result.Version)
	}
	elapsed := time.Since(start)
	fmt.Printf("Parsing 3 UAs took: %v\n", elapsed)
}
