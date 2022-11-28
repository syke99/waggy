# Waggy

The dead simple, easy-to-use library for writing WAGI (Web Assembly Gateway Interface) API handlers in Go

What problem does Waggy solve?
=====
With WAGI (Web Assembly Gateway Interface), HTTP requests are routed to specific handlers as defined in a modules.toml file
(more information can be found [here](https://github.com/deislabs/wagi)) to specific functions in a WASM module or WAT file.
It accomplishes this by passing in the HTTP request information via os.Stdin and os.Args, and returning HTTP responses via
os.Stdout. To remove considerable amounts of boilerplate code and to provide a familiar API for handling these WAGI HTTP
requests, Waggy was created. It aims to mimic the net/http package's libraries as much as possible (provided the functionality
is available in WAGI) to provide a seamless experience whenever switching between writing regular HTTP request handlers and
WAGI HTTP request handlers.

How do I use go-c2dmc?
====

### Installing
To install Waggy, simply run

```bash
$ go get github.com/syke99/waggy
```

Then you can import the package in any go file you'd like

```go
import "github.com/syke99/waggy"
```

### Basic usage

Then simply define your handler

```go
// Waggy handler signatures are slightly different from normal HTTP handlers,
// but function very similarly
func helloWorldHandler() {
	// instead of having a http.ResponseWriter and a http.Request passed
	// in at the function signature, with Waggy, you much initialize
	// them at the beginning of your handler
	w, r := waggy.Init()
	
	// then, it's as simple as using the regular net/http library from
	// the stdlib
	name := r.URL.Query.Get("name")
	
	responseBody := fmt.Sprintf("Hello world, my name is %s.", name)
	
	if r.URL.Query.Has("goodbye") {
        body = fmt.Sprintf("%s\n\n Goodbye for now!", responseBody)
	}
	
	_, _ := w.Write([]byte(responseBody))
}
```

**!!NOTE!!**

To learn more about configuring, routing, compiling, and deploying WAGI routes, please consult 
the [WAGI](https://github.com/deislabs/wagi/tree/main/docs) docs and the [TinyGo](https://tinygo.org/docs/guides/webassembly/) WASM docs

Who?
====

This library was developed by Quinn Millican ([@syke99](https://github.com/syke99))


## License

This repo is under the BSD 3 license, see [LICENSE](LICENSE) for details.