# gomdview

An easy way to display and share markdown notes in a web browser.  This project is specifically optimized for [Obsidian](https://obsidian.md) Vaults but can be used with any set of markdown files.

## Usage

### docker-compose

```yaml
version: "3.8"
services:
  mdview:
    image: ghcr.io/graytonio/mdview:latest
    container_name: mdview
    ports:
      - "9000:9000"
    volumes:
      - /path/to/my/vault:/markdown
    restart: unless-stopped
```

## Parameters

| Parameter | Function |
| --------- | -------- |
| `MD_ROOT` | Root directory for markdown files |
| `PORT` | Port for web interface |
