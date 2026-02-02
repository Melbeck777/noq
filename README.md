# noq

`noq` is a CLI tool for managing local Markdown files as the single source of truth, and synchronizing theme to external services:

- **memo** -> Notion database pages
- **article** -> Qiita articles (planned)

The goal is to write everything locally in Markdwon and safely push/pull changes.

## Status

This project is in start development.

## Planned features

- Local-first memo and article management in Markdown
- The information of connection written in destinations

# How to use

Initial new file
`noq memo new`

```
---
id: "yyyymmdd-{num}" # Automatically input
kind: "memo" # Automatically input
title: "teset memo" # Input your self
created_at: "2026-01-08T08:50:00+09:00" # Automatically input
updated_at: "2026-01-08T08:50:00+09:00" # Automatically input

destinations:
  - type: notion
    db: "notion-db-id" input your selef
    properties:
    page_id: null
    url: null
---
# Title

sentence…
![](assets/xxx.png)
```

After push `page_id` and `url` filled automaticaly.
`noq memo push`

```
---
id: "yyyymmdd-{num}"
kind: "memo"
title: "teset memo"
created_at: "2026-01-08T08:50:00+09:00"

destinations:
  - type: notion
    db: "notion-db-id"
    properties:
        status: "Progress"
        category: "企画"
    page_id: null
    url: null
---
# Title

sentence…
![](assets/xxx.png)
```

## Planned tech stack

- Language: Go
- CLI: `spf13/cobra`
- Config: 'spf13/viper' (for `.noq/config.yaml`)
- YAML: `gopkg.in/yaml.v3`
- Colored output (optional): `github.com/faith/color`

# ToDo

- [ ] Setup config with command
- [ ] Make new memo file in local `memo.root`
- [ ] Consider to about design article

config

```
memo_root: "/home/user/documents/noq/memo"
article_root: "/home/user/documents/noq/article"

notion:
    default_memo_db: "xxxxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxx"
    databases:
        memo: "xxxxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxx"
        article: "yyyyyyyyyy-yyyy-yyyy-yyyy-yyyyyyyyyy"

qiita:
    tokne: "YOUR_TOKEN"
```
