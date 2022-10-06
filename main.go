package main

// Import Packages
import (
	"fmt"
	"time"
)

// Main Function
func main() {
	// Collect Results
	var (
		BytesResult, br   = TestBytes(100000)
		StringsResult, sr = TestStrings(100000)
	)
	// Print the bytes result
	fmt.Printf(
		"\n>> Bytes Result: (%d) in %dns\n", len(br), BytesResult,
	)
	// Print the strings result
	fmt.Printf(
		"\n>> Strings Result: (%d) in %dns\n\n", len(sr), StringsResult,
	)
}

// Function to test appending a string to a bytes
// array. This method is MUCH FASTER than the
// appending to a string method.
func TestBytes(n int) (int64, []byte) {
	// Establish Variables
	var (
		// Start Time for tracking speed
		startTime int64 = time.Now().UnixNano()
		// Bytes array
		bytes []byte
	)
	// Perform 'n' iterations
	for i := 0; i < n; i++ {
		// Append the string to the bytes array
		bytes = append(bytes, []byte(fmt.Sprintf("Bytes_: %d", i))...)
	}
	// Return the speed and the resulting bytes array
	return time.Now().UnixNano() - startTime, bytes
}

// Function to test appending a string to another string.
// This method is MUCH SLOWER than using a bytes array.
func TestStrings(n int) (int64, string) {
	// Establish Variables
	var (
		// Start Time for tracking speed
		startTime int64 = time.Now().UnixNano()
		// String variable
		str string
	)
	// Perform 'n' iterations
	for i := 0; i < n; i++ {
		// Append the string to the str variable
		str += fmt.Sprintf("String: %d", i)
	}
	// Return the speed and the resulting string
	return time.Now().UnixNano() - startTime, str
}
