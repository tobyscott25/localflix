## Localflix Server

This server provides REST endpoints for the Localflix client to consume that serves your local media library.

## Usage

```bash
# With Docker (edit the docker-compose.yaml file to set your library location)
vim docker-compose.yaml
docker compose up

# Without Docker
export LF_LIBRARY_LOCATION=/path/to/your/library
go mod download
go run main.go
```
