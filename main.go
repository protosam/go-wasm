package main

import(
	"fmt"
	"syscall/js"
)

var wasm_loop = make(chan bool)

// This function will be exposed to javascript for use, it just spits back a value.
func speak(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		return args[0]
	}
	return "No value provided"
}

// This function is called when the WASM binary is started up in Javascript.
func main(){
	// Print hello to the console.
	fmt.Println("Hello")

	// Expose the speak function to javascript as "speak"
	js.Global().Set("speak", js.FuncOf(speak))

	// Keep the WASM binary running indefinitely.
	<-wasm_loop
}
