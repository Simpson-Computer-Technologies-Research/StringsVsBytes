# Strings VS Bytes ![Stars](https://img.shields.io/github/stars/Simpson-Computer-Technologies-Research/StringsVsBytes?color=brightgreen) ![Watchers](https://img.shields.io/github/watchers/Simpson-Computer-Technologies-Research/StringsVsBytes?label=Watchers)

![banner](https://user-images.githubusercontent.com/75189508/194193330-1f29dcee-1c3c-490b-b88b-6eb109c53d3c.png)
Appending to a String or Byte Array, Which is faster?


# Benchmarks

```go
// About 0.11 Seconds
>> Bytes Result: (1288890) in 11796900ns

// About 5.17 Seconds
>> Strings Result: (1288890) in 5179652000ns
```

# Functions

```go
// Test Bytes Method
func TestBytes(n int) {
    // Perform 'n' iterations
	for i := 0; i < n; i++ {
		// Append the string to the bytes array
		bytes = append(bytes, []byte(fmt.Sprintf("Bytes_: %d", i))...)
	}
}

// Test Strings Method
func TestStrings(n int) {
	// Perform 'n' iterations
	for i := 0; i < n; i++ {
		// Append the string to the str variable
		str += fmt.Sprintf("String: %d", i)
	}
}
```

# License
MIT License

Copyright (c) 2022 Tristan Simpson

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
