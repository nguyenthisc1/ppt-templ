# PPT Templ Property Starter

Property-focused Go-native starter for server-rendered pages with [`templ`](https://templ.guide/) and [`htmx`](https://htmx.org/). It includes a reusable layout shell, a multi-level property navigation, curated listing and project pages, a market news page, a hero-triggered contact modal, and BEM-oriented CSS.

## Stack

- Go HTTP server with `net/http`
- `templ` components for layouts, shared UI, and pages
- plain Go data layer for listing, project, and news mock content
- Plain CSS with tokens and BEM blocks
- local image assets served from `public/assets/img`

## Quick Start

1. Install Go and `templ`.
2. Install optional frontend lint dependencies if you want JS linting:

   ```bash
   npm install
   ```
3. Generate templates:

   ```bash
   templ generate
   ```

4. Sync Go modules:

   ```bash
   go mod tidy
   ```

5. Run the app:

   ```bash
   go run ./cmd/web
   ```

6. Open `http://localhost:8080`.

## Project Structure

- `cmd/web`: app entrypoint
- `internal/content`: property listings, projects, and news mock data
- `internal/http`: route registration and rendering helpers
- `internal/views/layout`: shared page shell and head contract
- `internal/views/components`: reusable UI blocks like header, cards, and contact modal
- `internal/views/pages`: full page templates
- `public/assets`: plain CSS and JavaScript served directly by Go
- `docs/conventions.md`: project rules for layout, tokens, BEM, and performance
- `docs/build-plan.md`: next-phase implementation roadmap
- `docs/documentation-plan.md`: documentation scope and maintenance guidance
- `docs/setup-project-example.md`: baseline setup example for new projects

## Included Pages

- `/`: property-focused homepage with hero, featured listings, projects, news, and contact modal CTA
- `/listings`: searchable listings page with sale/rent filters
- `/listings/:slug`: listing detail page with image slider, property content, and sticky inquiry form
- `/projects`: searchable projects page with stage filters
- `/news`: market news and editorial page

## Performance Defaults

- Inline only the shell-critical CSS required for first paint
- Preload the primary stylesheet before loading it normally
- Defer JavaScript so HTML and CSS stay first
- Keep interactive behavior lightweight and progressively enhance with small JavaScript

## Images

- Put page imagery in `public/assets/img`
- Reference images with `/assets/img/...`
- Current mock data uses `hero.avif`, `listing_*.avif`, and `project_*.avif`

## Conventions

The repo rules live in [docs/conventions.md](./docs/conventions.md). Use that document as the default for future pages and components.

## Setup Files

- `.editorconfig`: shared editor and whitespace defaults
- `.gitignore`: ignored local/build/tooling files
- `eslint.config.js`: ESLint flat config for project JavaScript
- `package.json`: optional frontend lint scripts and dependencies

## Linting

Run frontend linting with:

```bash
npm run lint:js
```
