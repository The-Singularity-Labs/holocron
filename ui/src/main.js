import "./assets/css/ash.min.css";
import Alpine from 'alpinejs';
import EncryptForm from './components/encrypt_form.js';
import DecryptForm from './components/decrypt_form.js';
import ConnectionCheck from './components/connection_check.js';
import './wasm_exec'

import wasmURL from "url:./assets/wasm/golib.wasm";

const go = new global.Go();

// register global stores
Alpine.store('images', {
    svgs: {
      copy: new URL('assets/images/copy.svg',import.meta.url),
      check: new URL('assets/images/check.svg',import.meta.url)
    }
})

Alpine.store('global_funcs', {
    go: {},
    setGoFunc(funcName, func) {
        this.go[funcName] =  func
    }
})

Alpine.store('app', {
  isOnline: null
})

// init go functions
export const wasmBrowserInstantiate = async (wasmModuleUrl, importObject) => {
    let response = undefined;
  
    // Check if the browser supports streaming instantiation
    if (WebAssembly.instantiateStreaming) {
      // Fetch the module, and instantiate it as it is downloading
      response = await WebAssembly.instantiateStreaming(
        fetch(wasmModuleUrl),
        importObject
      );
    } else {
      // Fallback to using fetch to download the entire module
      // And then instantiate the module
      const fetchAndInstantiateTask = async () => {
        const wasmArrayBuffer = await fetch(wasmModuleUrl).then(response =>
          response.arrayBuffer()
        );
        return WebAssembly.instantiate(wasmArrayBuffer, importObject);
      };
  
      response = await fetchAndInstantiateTask();
    }
  
    return response;
};

const addWasmFunctions = async () => {
    // Get the importObject from the go instance.
    const importObject = go.importObject;

    // Instantiate our wasm module
    const wasmModule = await wasmBrowserInstantiate(wasmURL, importObject);

    // Allow the wasm_exec go instance, bootstrap and execute our wasm module
    go.run(wasmModule.instance);

    // Set the add function into the wasm store
    Alpine.store('global_funcs').setGoFunc("encrypt", global.encrypt)
    Alpine.store('global_funcs').setGoFunc("decrypt", global.decrypt)
};
addWasmFunctions();


window.Alpine = Alpine;


 

// Register Components
document.addEventListener('alpine:init', () => {
    Alpine.data('encryptForm', EncryptForm);
});

document.addEventListener('alpine:init', () => {
  Alpine.data('decryptForm', DecryptForm);
});


document.addEventListener('alpine:init', () => {
  Alpine.data('connectionCheck', ConnectionCheck);
});

Alpine.start();

