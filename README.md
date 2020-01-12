# Go WASM Example
As of version `1.13.5` of Go, I believe that targeting WASM is useful. You can expose Go functions to Javascript and synchronously get a value directly back in Javascript in addition to making Asynchronous functions.

To compile this, you first need to get `wasm_exec.js` which is provided with Go. There is a `misc/wasm` directory provided that contains:
```
wasm_exec.html - An example HTML file for using wasm_exec.js + a Go WASM binary
wasm_exec.js - Is a helper script provided to make Go WASM binaries work.
```

Copy `wasm_exec.js` into the `www` directory of this project.


After the `wasm_exec.js` script is copied, compile `main.go` and run the example server:
```
$ GOOS=js GOARCH=wasm go build -o www/test.wasm main.go
$ go run server.go
```

Now you can navigate to `localhost:8000/` in your browser, pop open the console, and see where `func speak(...)` was called from Javascript.

## Use Case
With this example you can begin making something like cryptographic functions in Go and supply APIs to be used by Javascript.

You could also provide compiled closed-source APIs for consumption by Javascript developers.

## Additional Notes
If the `example.html` stops working, reference `wasm_exec.html` to figure out why.

It is possible to also load in WASM binaries in NodeJS, I believe this was fixed in a prior `wasm_exec.js`, but if you do run into problems, you just need to patch that file. 
