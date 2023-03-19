package main

import (
  "syscall/js"

  "github.com/the-singularity-labs/holocron/pkg/forge"
)

// Declare a main function, this is the entrypoint into our go module
// That will be run. In our example, we won't need this
func main() {
	js.Global().Set("encrypt", encryptWrapper())
  js.Global().Set("decrypt", decryptWrapper())
  <-make(chan bool)
}

func encryptWrapper()  js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 5 {
			return wrap("", "Not enough arguments")
		}
		h := forge.NewHolocron(args[0].String(), args[1].String(), args[2].String(), args[3].String(), args[4].String())

    encoded, err := h.ToCompressedString()
		if err != nil {
			return wrap(map[string]interface{}{}, err.Error())
		}

    base64QrCode, err := h.ToCompressedQrCode() 
		if err != nil {
			return wrap(map[string]interface{}{}, err.Error())
		}
  
    return wrap(map[string]interface{}{"encrypted": encoded, "base64_qr_code": base64QrCode}, "")
	})
}

func decryptWrapper()  js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 2 {
			return wrap("", "Not enough arguments")
		}
		
    // Decrypt(passphrase, ciphertext string)
		decoded, err  := forge.Decrypt(args[0].String(), args[1].String())
		if err != nil {
			return wrap(map[string]interface{}{}, err.Error())
		}
  
    return wrap(map[string]interface{}{"decoded": decoded}, "")
	})
}

func wrap(result interface{}, errString string) map[string]interface{} {
	return map[string]interface{}{
		"error":   errString,
		"data": result,
	}
}
