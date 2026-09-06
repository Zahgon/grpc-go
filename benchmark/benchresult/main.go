package main

import (
	"os"
	"time"

	"google.golang.org/grpc/benchmark/stats"
)

func createMap(fileName string) map[string]stats.BenchResults {
	_ = "STUB: not implemented"
	return nil
}

func intChange(title string, val1, val2 uint64) string { _ = "STUB: not implemented"; return "" }

func floatChange(title string, val1, val2 float64) string { _ = "STUB: not implemented"; return "" }

func timeChange(title string, val1, val2 time.Duration) string {
	_ = "STUB: not implemented"
	return ""
}

func strDiff(title, val1, val2 string) string { _ = "STUB: not implemented"; return "" }

func compareTwoMap(m1, m2 map[string]stats.BenchResults) { _ = "STUB: not implemented"; return }

func compareBenchmark(file1, file2 string) { _ = "STUB: not implemented"; return }

func printHeader() { _ = "STUB: not implemented"; return }

func printline(benchName string, d stats.RunData) { _ = "STUB: not implemented"; return }

func formatBenchmark(fileName string) { _ = "STUB: not implemented"; return }

func main() {
	if len(os.Args) == 2 {
		formatBenchmark(os.Args[1])
	} else {
		compareBenchmark(os.Args[1], os.Args[2])
	}
}
