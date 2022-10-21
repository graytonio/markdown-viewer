# gomdview

An easy way to display and share markdown notes in a web browser.  This project is specifically optimized for [Obsidian](https://obsidian.md) Vaults but can be used with any set of markdown files.

## Usage

### docker-compose

```yaml
version: "3.8"
services:
  markdown-viewer:
    restart: unless-stopped
    container_name: web-service
    build:
      context: .
    volumes:
      - ~/Library/CloudStorage/GoogleDrive-graytonw@jfrog.com/My Drive/Jfrog Notes:/markdown
  nginx:
    image: nginx:alpine
    ports:
      - "80:80"
    volumes:
      - ./nginx_server.conf:/etc/nginx/conf.d/default.conf
```

## Parameters

| Parameter | Function |
| -- | -- |
| `MD_ROOT` | Root directory for markdown files |
| `PORT` | Port for web interface |