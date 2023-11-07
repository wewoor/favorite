# Favorite 

Preview, create, and edit the document site of [Docusaurus](https://docusaurus.io/) on the cloud.

## Introduction

Favorite is a  document site editor running on the cloud, you can create and change your Docusaurus markdown file, and preview it on the web. 
This is just a toy project, just for learning, not used for production.

## Tech Stack

- Next.js (Web Frontend)
- Gin (Server-Side)
- Docker (Git/Node.js Exec Env)
- Docker Images: Node:16
- MySQL (Database)
- Prisma (ORM)
- Project File Storage (local temporary)

## Schema initial

```bash
> go run -mod=mod entgo.io/ent/cmd/ent init --target schema Project User
```
