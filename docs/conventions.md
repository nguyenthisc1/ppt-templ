# Frontend Conventions

## Setup Rules

- Commit `.editorconfig`, `.gitignore`, and the base lint config at project start.
- Keep formatting rules simple and shared across editors.
- Limit lint tooling to code we own; do not lint vendored assets.
- Treat Node tooling as optional support for frontend quality checks unless a real asset pipeline is introduced.
- Document setup commands in `README.md` and keep an example setup flow in `docs/setup-project-example.md`.

## Layout Rules

- All pages should render through `layout.Base`.
- Keep `Header`, `Nav`, `Main`, and `Footer` shared unless the shell changes globally.
- Put full document rendering in `internal/views/pages`.
- Put HTMX-only fragments in `internal/views/partials`.
- Prefer one clear HTMX target per interaction.

## CSS Rules

- Use design tokens from `:root` before introducing hard-coded values.
- Use BEM naming for components: `.block`, `.block__element`, `.block--modifier`.
- Keep one block responsible for one concept.
- Avoid deep selector nesting and selector chains that depend on markup trivia.
- Use page-specific classes only when shared blocks are insufficient.
- Keep global element selectors limited to resets and typography baseline behavior.

## Token Rules

Required token families for new work:

- brand colors: primary, secondary, strong, soft
- text and surface colors
- border and shadow values
- spacing scale
- radius scale
- type scale
- container width

## Performance Rules

- Preserve the critical CSS slice in `layout.Base`.
- Keep the primary stylesheet loaded through preload + stylesheet link tags.
- Keep scripts deferred by default.
- Add JavaScript only when the interaction cannot be expressed well with HTML and HTMX alone.
- Prefer server-rendered fallbacks for HTMX endpoints when direct access is possible.

## Component Rules

- Add shared interface pieces to `internal/views/components`.
- If a component needs a modifier, prefer `.block--modifier` over creating a sibling block with duplicated styles.
- When adding a new page, start by composing existing blocks before inventing a new layout pattern.
- Keep copy and semantic structure inside templates, and keep styling decisions inside the CSS token and block system.
