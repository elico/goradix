# GoRadix [![Build Status](https://travis-ci.org/falmar/goradix.svg?branch=master)](https://travis-ci.org/falmar/goradix)

Radix Tree implementation written in Golang. **Still under development**.

Thread Safe (branches):
- master - Will be concurrent safe by default
- ts - Concurrent Safe
- nts - Not Concurrent Safe

#### Radix Tree
> In computer science, a radix tree (also radix trie or compact prefix tree) is a data structure that represents a space-optimized trie in which each node that is the only child is merged with its parent. - [Wikipedia](https://en.wikipedia.org/wiki/Radix_tree)

![](https://upload.wikimedia.org/wikipedia/commons/thumb/a/ae/Patricia_trie.svg/400px-Patricia_trie.svg.png)

## Content
 - [Usage](#usage)
 - [Benchmarks](#benchmarks)
 - [License](#license)

## TODO:


| Code | Implementation | Test | Benchmark | Usage Examples |
|---|:---:|:---:|:---:|:---:|
| Insert | x | x | x | x |
| LookUp | x | x | x | x |
| AutoComplete | x | x |  | x |
| Remove |  |  |  | | |


## Usage:

Download: `go get github.com/falmar/goradix`

Import `import "github.com/falmar/goradix"`

### Insert

```go
// string required, value (optional)
func (r *Radix) Insert(s string, value ...interface{}){...}
```
```go
package main

import "github.com/falmar/goradix"

func main() {
	radix := goradix.New()

  // Simple string insert
	radix.Insert("romanus")
	radix.Insert("romane")
	radix.Insert("romulus")
}
```

### LookUp

```go
// string required
func (r *Radix) LookUp(s string) (interface{}, error) {...} {...}
```
```go
package main

import (
	"fmt"

	"github.com/falmar/goradix"
)

func main() {
	radix := goradix.New()
	radix.Insert("romanus", 1)
	radix.Insert("romane", 100)
	radix.Insert("romulus", 1000)

	value, err := radix.LookUp("romane")

	if err != nil { // No Match Found
		return
	}

	// Output: Found node, Value: 100
	fmt.Println("Found node, Value: ", value)
}
```

### AutoComplete
```go
// string, bool required
func (r Radix) AutoComplete(s string, wholeWord bool) ([]string, error) {...}
```
```go
package main

import (
	"fmt"

	"github.com/falmar/goradix"
)

func main() {
	radix := goradix.New()
	radix.Insert("romanus", 1)
	radix.Insert("romane", 100)
	radix.Insert("romulus", 1000)

	// Return remaining text
	words, err := radix.AutoComplete("ro", false)

	if err != nil { // No Match Found
		return
	}

	// Output: AutoComplete 'rom'; Words: [manus mane mulus]
	fmt.Printf("AutoComplete: '%s'; Words: %v\n", "ro", words)

	// Return whole words
	words, _ = radix.AutoComplete("ro", true)

	// Output: AutoComplete 'rom'; Words: [romanus romane romulus]
	fmt.Printf("AutoComplete: '%s'; Words: %v\n", "ro", words)
}
```

### Match
Under development

### Remove
Under development

## Benchmarks

### Insert

**Go 1.6**

```text
BenchmarkInsertString-2      	 2000000	       809 ns/op	     238 B/op	       7 allocs/op
BenchmarkInsertBytes-2       	 2000000	       730 ns/op	     227 B/op	       6 allocs/op

```

**Go 1.7**

```text
BenchmarkInsertString-2   	 2000000	       618 ns/op	     238 B/op	       7 allocs/op
BenchmarkInsertBytes-2    	 3000000	       546 ns/op	     227 B/op	       6 allocs/op
```

### LookUp

**Go 1.6**
```text
BenchmarkLookUpStringSingle-2	10000000	       138 ns/op	       0 B/op	       0 allocs/op
BenchmarkLookUpStringRandom-2	 3000000	       389 ns/op	       0 B/op	       0 allocs/op
BenchmarkLookUpBytesSingle-2 	10000000	       103 ns/op	       0 B/op	       0 allocs/op
BenchmarkLookUpBytesRandom-2 	 5000000	       422 ns/op	       0 B/op	       0 allocs/op
```

**Go 1.7**

```text
BenchmarkLookUpStringSingle-2   	20000000	       175 ns/op	       0 B/op	       0 allocs/op
BenchmarkLookUpStringRandom-2   	 5000000	       386 ns/op	       0 B/op	       0 allocs/op
BenchmarkLookUpBytesSingle-2    	20000000	       174 ns/op	       0 B/op	       0 allocs/op
BenchmarkLookUpBytesRandom-2    	 5000000	       364 ns/op	       0 B/op	       0 allocs/op

```

**Benchmark machine**

```text
Architecture:          x86_64
CPU op-mode(s):        32-bit, 64-bit
Byte Order:            Little Endian
CPU(s):                2
On-line CPU(s) list:   0,1
Thread(s) per core:    1
Core(s) per socket:    1
Socket(s):             2
NUMA node(s):          1
Vendor ID:             GenuineIntel
CPU family:            6
Model:                 62
Model name:            Intel(R) Xeon(R) CPU E5-2630L v2 @ 2.40GHz
Stepping:              4
CPU MHz:               2399.998
BogoMIPS:              4799.99
Virtualization:        VT-x
Hypervisor vendor:     KVM
Virtualization type:   full
L1d cache:             32K
L1i cache:             32K
L2 cache:              256K
L3 cache:              15360K
NUMA node0 CPU(s):     0,1
```

## License

Copyright 2016 David Lavieri. All rights reserved.
Use of this source code is governed by a MIT License that can be found in the LICENSE file.
