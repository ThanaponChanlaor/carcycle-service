# CarCycleService

Go REST API backend for the CarCycle app. Implements the contract `CarCycleUI` already
expects (see that project's README), backed by SQLite (pure-Go driver, no CGO / C compiler
needed).

## Stack

- Go standard library `net/http` (Go 1.22+ method/pattern routing — `GET /api/cars/{id}` etc.),
  no web framework.
- `modernc.org/sqlite` — pure-Go SQLite driver, single file database (`carcycle.db`), no
  separate database server to install or run.
- No ORM — plain `database/sql` queries in `cars.go`.

## Endpoints

| Method | Path                 | Purpose                                    |
|--------|----------------------|---------------------------------------------|
| GET    | `/api/cars`          | List all cars                              |
| GET    | `/api/cars/{id}`      | Get one car                                |
| GET    | `/api/cars/{id}/history` | Get a car's change history               |
| PATCH  | `/api/cars/{id}`      | Update a car (any subset of editable fields) |
| GET    | `/`                   | Health check                               |

`PATCH` diffs the incoming fields against the current row and — only for fields that actually
changed — appends a `car_history` row (field, old value, new value, timestamp). `changedBy` is
currently hardcoded to `"CarCycle Staff"`; there's no real user identity yet since the FE's admin
login is just a client-side gate (see [[project-car-detail-fe-scaffold]] memory / FE README).

## Getting started

```bash
go mod tidy   # fetches modernc.org/sqlite and writes go.sum (needed once)
cp .env.example .env
go run .
```

The server listens on `PORT` (default `8080`) and stores data in `DB_PATH` (default
`carcycle.db`, created automatically with the same demo fleet the FE mock uses, so switching
the FE from mock data to this API shouldn't change what you see).

## Wiring up the frontend

In `CarCycleUI/.env`:

```
VITE_USE_MOCK_DATA=false
VITE_API_BASE_URL=http://localhost:8080/api
```

If you're testing from a phone on the same network (QR-code flow), `localhost` won't resolve on
the phone — use this PC's LAN IP instead (same one Vite's `Network:` URL prints), e.g.
`http://192.168.x.x:8080/api`. With `ALLOWED_ORIGINS` empty, CORS is wide open
(`Access-Control-Allow-Origin: *`) so this works regardless of which origin the FE is served from.
Set `ALLOWED_ORIGINS` (comma-separated) to restrict it — production sets it to the Netlify URL.

## Branching & deployment

- Git flow: `feature/*` → PR into `dev` → PR `dev` into `main`. The `Branch policy` workflow
  (`source-branch` check) fails PRs into `dev` that aren't from `feature/*` and PRs into `main`
  that aren't from `dev`. Both branches are protected: no direct pushes, CI must pass, and the
  repo owner merges.
- `CI` (`go vet`, `go build`, `go test`) runs on every PR to `dev`/`main` and on pushes to `dev`.
- `Deploy to Render` runs on push to `main`: CI steps, then POSTs the Render deploy hook
  (repo secret `RENDER_DEPLOY_HOOK_URL`). The Render service is defined in `render.yaml`
  (`autoDeploy: false`, so only this workflow deploys).
- Production data lives in Firestore (the FE talks to it directly). The SQLite DB on Render is
  demo/seed data only and resets on every deploy.

## Not implemented yet

- No authentication — every request can read and write. The FE's admin login is FE-only right
  now; this backend should eventually check the caller's role on `PATCH` itself once real auth
  exists, rather than trusting the FE to gate it.
- No endpoint to create a new car (FE doesn't have that UI yet either) — only the seeded fleet
  exists until one is added.
