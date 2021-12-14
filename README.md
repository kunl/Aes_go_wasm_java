# aes_go_wasm_java
aes for go and java; build go fo wasm and use wasm parse java response.


vscode setting config

settings.json
``` json
{
    "go.toolsEnvVars": {
        "GOOS": "js",
        "GOARCH": "wasm"
    }
}
```

## coding in vscode 
    When using Go modules related to WebAssembly, namely syscall/js, the default settings in VS Code will trigger error reports like "Build constraints exclude all Go files" in the editor. In this lesson, we will learn how to fix this and improve our coding experience.
https://egghead.io/lessons/go-configure-go-build-constraints-in-vs-code-to-work-with-webassembly


## coding in goLand
    WebAssembly (Wasm) is a binary code that you can run in a browser. GoLand supports generation of WASM files from GO files. 
 https://www.jetbrains.com/help/go/webassembly-project.html
