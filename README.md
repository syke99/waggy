# Waggy
[![Go Reference](https://pkg.go.dev/badge/github.com/syke99/waggy.svg)](https://pkg.go.dev/github.com/syke99/waggy)
[![Go Reportcard](https://goreportcard.com/badge/github.com/syke99/waggy)](https://goreportcard.com/report/github.com/syke99/waggy)
[![Codecov](https://codecov.io/gh/syke99/waggy/branch/main/graph/badge.svg?token=KDYH3JO1QI)](https://codecov.io/gh/syke99/waggy)
[![LICENSE](https://img.shields.io/github/license/syke99/waggy)](https://pkg.go.dev/github.com/syke99/waggy/blob/master/LICENSE)

The dead simple, easy-to-use library for writing HTTP handlers and routers in Go that can be used in standard HTTP server environments or in WAGI (Web Assembly Gateway Interface) environments

What problems do Waggy solve?
=====
With WAGI (Web Assembly Gateway Interface), HTTP requests are routed to specific handlers as defined in a modules.toml file
(more information can be found [here](https://github.com/deislabs/wagi)) to specific functions in a WASM module or WAT file.
It accomplishes this by passing in the HTTP request information via os.Stdin and os.Args, and returning HTTP responses via
os.Stdout. To remove considerable amounts of boilerplate code and to provide a familiar API for handling these WAGI HTTP
requests, Waggy was created. It provides path parameter access functionality that will feel very reminiscent to those who 
have used [gorilla/mux](https://github.com/gorilla/mux). Additionally, you can also map multiple handlers to a specific route
based on the specific HTTP method that was used in the incoming request. Waggy also allows users to compile an entire server's 
worth of routes into a single WASM module and bypass setting up their routes via a modules.toml file if they so choose by 
handling mapping the route to the correct entry point (handler). But don't worry, you can also compile individual routes into 
their own WASM modules, too, so you can use the conventional modules.toml file for routing. 

</br>

#### v0.5.0 improvements:
* **Restricting Methods**: With v0.5.0, you have the ability to restrict HTTP methods on your handlers to make sure your endpoints
are more secure 

* **No Matched Routes**: Another added feature in v0.5.0 is handling cases where no routes are matches to the route of the incoming
request with generic 405 responses, as well as the ability to set custom handlers for the same edge case

How do I use Waggy?
====

### Installing
To install Waggy in a repo, simply run

```bash
$ go get github.com/syke99/waggy
```

Then you can import the package in any go file you'd like

```go
import "github.com/syke99/waggy"
```

### Basic usage

Examples of using both Routers and Handlers for compiling WASM modules for WAGI can be found in the [examples](https://github.com/syke99/waggy-examples) repo.



**!!NOTE!!**

To learn more about configuring, routing, compiling, and deploying WAGI routes, as well as
the limitations of WAGI routes, please consult the [WAGI](https://github.com/deislabs/wagi/tree/main/docs) docs
and the [TinyGo](https://tinygo.org/docs/guides/webassembly/) WASM docs

Who?
====

This library was developed by Quinn Millican ([@syke99](https://github.com/syke99))


## License

This repo is under the BSD 3 license, see [LICENSE](../LICENSE) for details.
