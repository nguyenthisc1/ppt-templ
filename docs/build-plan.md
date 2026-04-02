# Build Plan

## Phase 1

- Keep the current mock-data approach for listings, projects, and news while validating layout and conversion flow.
- Use the shared contact modal as the default inquiry touchpoint from hero and card-level actions.
- Keep search and filtering query-string based so the site remains easy to test and index.

## Phase 2

- Move listing, project, and news content into structured JSON, CMS, or database-backed sources.
- Add detail pages for individual listings and projects.
- Add lead handling for the contact form with validation, persistence, and notification delivery.

## Phase 3

- Add asset fingerprinting and stronger cache headers for production.
- Add SEO enhancements: sitemap, structured metadata, and Open Graph previews per page.
- Consider HTMX-powered partial filtering once the server-side query behavior is stable.
