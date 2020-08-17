# tfgen

Build wasm binary

```
GOARCH=wasm GOOS=js go build -o ./server/web/app.wasm
```

Build server binary

```
go build -o ./server/server ./server/server.go
```

Run server

```
./server/server
```