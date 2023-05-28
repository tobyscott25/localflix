## Static file server

This serves static files from the `assets` directory over HTTP for the client to consume.

Install dependancies

```bash
go mod download
```

Set your library location

```bash
export LIBRARY_LOCATION=/path/to/your/library
```

Run the server

```bash
go run main.go
```

Visit `localhost:8080/assets/video.mp4` to see the static file server in action.
