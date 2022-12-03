package main

import (
	"syscall/js"
	"github.com/the-singularity-labs/holocron"
)



func main() {
	js.Global().Set("ForgeHolocron", encodeWrapper())
	js.Global().Set("ForgeLiteHolocron", encodeLiteWrapper())
	js.Global().Set("DecodeHolocron", decodeWrapper())
	<-make(chan bool)
}

func encodeWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 5 {
			return wrap("", "Not enough arguments")
		}
		// h := holocron.NewHolocron(name, prompt, ascertainment, treasure, salt)
		h := holocron.NewHolocron(args[0].String(), args[1].String(), args[2].String(), args[3].String(), args[4].String())
		encoded, err := h.ForgeForWeb()
		var errString string
		if err != nil {
			errString = err.Error()
		}
		return wrap(encoded, errString)
	})
}

func encodeLiteWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 5 {
			return wrap("", "Not enough arguments")
		}
		// h := holocron.NewHolocron(name, prompt, ascertainment, treasure, salt)
		h := holocron.NewHolocron(args[0].String(), args[1].String(), args[2].String(), args[3].String(), args[4].String())
		encoded, err := h.ForgeLightForWeb()
		var errString string
		if err != nil {
			errString = err.Error()
		}
		return wrap(encoded, errString)
	})
}

func decodeWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 2 {
			return wrap("", "Not enough arguments")
		}
		// Decrypt(passphrase, ciphertext string)
		decoded := holocron.Decrypt(args[0].String(), args[1].String())
		return wrap(decoded, "")
	})
}

func wrap(result string, err string) map[string]interface{} {
	return map[string]interface{}{
		"error":   err,
		"data": result,
	}
}