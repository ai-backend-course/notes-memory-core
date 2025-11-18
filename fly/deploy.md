# notes-memory-core — Deployment Notes (Fly.io)

## App Name
`notes-memory-core-api`

---

## Initial Setup

```bash
flyctl launch
```

**Prompts:**

- **App name:** `notes-memory-core-api`
- **Postgres:** Yes
- **Deploy now:** Yes

This command:

- generated `fly.toml`  
- provisioned a Fly Postgres instance  
- attached Postgres to the app  
- built the Docker image  
- deployed the app  
- assigned a public URL  

---

## Connecting Fly Postgres to the App

After `flyctl launch`, Fly provides a connection string.

Store it as an environment variable:

```bash
flyctl secrets set DATABASE_URL="postgres://<user>:<pass>@<host>:<port>/<db>?sslmode=disable"
```

(Use the connection string Fly shows you after provisioning the DB.)

---

## Restart After Setting Secrets

```bash
flyctl apps restart notes-memory-core-api
```

---

## Health Check

```bash
curl https://notes-memory-core-api.fly.dev/health
```

**Expected Response:**

```json
{"status":"ok"}
```

---

## Get Notes Test

```bash
curl https://notes-memory-core-api.fly.dev/notes
```

---

## Logs

```bash
flyctl logs -a notes-memory-core-api
```

---

## Live URL

```
https://notes-memory-core-api.fly.dev
```

(Custom domain `api.jeffellis.dev` will be added on Day 9.)

---

## Deployment Complete

The Notes-Memory-Core API is fully deployed, running with Fly Postgres, tested, and documented.
