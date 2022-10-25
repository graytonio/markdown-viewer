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
    environment:
      - VAULT=git
      - GIT_URL=https://github.com/graytonio/my-notes.git
    restart: unless-stopped
```

## Parameters

| Parameter | Function | Default |
| --------- | -------- | ------- |
| `MD_ROOT` | Root directory for markdown files | `/markdown` |
| `PORT` | Port for web interface | `9000` |
| `VAULT`| Type of vault file sync ('git', 'local') | `local` |
| `GIT_URL` | URL to git repo to pull from (required if using git vault) |  |
| `GIT_UPDATE` | Cron string for how often to pull from git | `30 * * * *` |

## Roadmap

- [ ] CI/CD Pipelines for automatic release of new versions when tagged
- [ ] Authentication for private git repos
- [ ] Multiple themes
- [ ] Sidebar tree view of notes
- [ ] Search