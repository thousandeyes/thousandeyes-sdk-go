[![GoDoc](https://godoc.org/github.com/thousandeyes/thousandeyes-sdk-go?status.svg)](http://godoc.org/github.com/thousandeyes/thousandeyes-sdk-go) [![Go Report Card](https://goreportcard.com/badge/github.com/thousandeyes/thousandeyes-sdk-go)](https://goreportcard.com/report/github.com/thousandeyes/thousandeyes-sdk-go) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/gojp/goreportcard/blob/master/LICENSE)

# thousandeyes-sdk-go

DEPRECATED TO BE UPDATED LATER


`thousandeyes-sdk-go` is a [go](https://golang.org/) client library for the [Thousandeyes v6 API](https://developer.thousandeyes.com/v6).

## Installation
thousandeyes-sdk-go is compatible with modern Go releases in module mode, with Go installed:

```cli
go get github.com/thousandeyes/thousandeyes-sdk-go/v2
```

will resolve and add the package to the current development module, along with its dependencies.

Alternatively the same can be achieved if you use import in a package:
```go
import "github.com/thousandeyes/thousandeyes-sdk-go/v2"
```

and run go get without parameters.

## Usage
Example code to list ThousandEyes agents:

```go
package main

import (
	"fmt"
	"os"

	"github.com/thousandeyes/thousandeyes-sdk-go/v2"
)

func main() {
	opts := thousandeyes.ClientOptions{
		AuthToken: os.Getenv("TE_TOKEN"),
		AccountID: os.Getenv("TE_AID"),
	}

	client := thousandeyes.NewClient(&opts)
	agents, err := client.GetAgents()
	if err != nil {
		panic(err)
	}
	for _, a := range *agents {
		fmt.Println(*a.AgentName)
	}
}
```

## Contributing
1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request

## License
This library is distributed under the Apache 2.0 license found in the [LICENSE](/LICENSE) file.

## Maintenance and Acknowledgements
This project is maintained by the ThousandEyes engineering team and accepts community contributions.

ThousandEyes would like to extend a thank you to William Fleming, John Dyer, and Joshua Blanchard for their contributions and community maintenance of this project.
