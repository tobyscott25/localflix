## Static file server

This serves static files from the `assets` directory over HTTP for the client to consume.

Install dependancies

```bash
go mod download
```

Run the server

```bash
go run main.go
```

Visit `localhost:8080/assets/screenshot.png` to see the static file server in action.
