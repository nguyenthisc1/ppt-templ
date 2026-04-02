# PPT Templ Starter

Minimal Go-native starter for server-rendered pages with [`templ`](https://templ.guide/) and [`htmx`](https://htmx.org/). It includes a reusable layout shell, BEM-oriented CSS, design tokens, a sample HTMX fragment endpoint, and default performance rules.

## Stack

- Go HTTP server with `net/http`
- `templ` components for layouts, shared UI, and pages
- `htmx` for targeted partial updates
- Plain CSS with tokens and BEM blocks

## Quick Start

1. Install Go and `templ`.
2. Generate templates:

   ```bash
   templ generate
   ```

3. Sync Go modules:

   ```bash
   go mod tidy
   ```

4. Run the app:

   ```bash
   go run ./cmd/web
   ```

5. Open `http://localhost:8080`.

## Project Structure

- `cmd/web`: app entrypoint
- `internal/http`: route registration and rendering helpers
- `internal/views/layout`: shared page shell and head contract
- `internal/views/components`: reusable UI blocks like header and footer
- `internal/views/pages`: full page templates
- `internal/views/partials`: HTMX fragment templates
- `public/assets`: plain CSS and JavaScript served directly by Go
- `docs/conventions.md`: project rules for layout, tokens, BEM, and performance

## Included Pages

- `/`: home page with HTMX fragment demo
- `/about`: structure and conventions overview
- `/contact`: checklist-style example content page
- `/fragments/highlight?audience=founders|teams|marketers`: fragment endpoint with full-page fallback

## Performance Defaults

- Inline only the shell-critical CSS required for first paint
- Preload the primary stylesheet before loading it normally
- Defer JavaScript so HTML and CSS stay first
- Keep HTMX usage scoped to focused content regions

## Conventions

The repo rules live in [docs/conventions.md](./docs/conventions.md). Use that document as the default for future pages and components.
# ppt-templ
