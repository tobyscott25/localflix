## Static file server

This serves static files from the `assets` directory over HTTP for the client to consume.

Install dependancies

```bash
go mod download
```

Set your library location as an environment variable

```bash
vim docker-compose.yaml

# or if you're not using Docker
export LIBRARY_LOCATION=/path/to/your/library
```

Run the server

```bash
docker compose up

# or if you're not using Docker
go run main.go
```

Visit `localhost:8080/assets/video.mp4` to see the static file server in action.
