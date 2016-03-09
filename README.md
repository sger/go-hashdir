[![Build Status](https://travis-ci.org/sger/go-hashdir.svg?branch=master)](https://travis-ci.org/sger/go-hashdir)

# hashdir

Simple package for hashing directory strings 

```
package main

import (
	"fmt"
	hashdir "github.com/sger/go-hashdir"
)

func main() {

	hash, err := hashdir.Create("test/directory", "md5")
	fmt.Println(hash) // returns d41d8cd98f00b204e9800998ecf8427e
}
```

Testing
-----

```ruby
$ go test -v
```

Author
-----

__Spiros Gerokostas__ 

- [![](https://img.shields.io/badge/twitter-sger-brightgreen.svg)](https://twitter.com/sger) 
- :email: spiros.gerokostas@gmail.com

License
-----

go-hashdir is available under the MIT license. See the LICENSE file for more info.