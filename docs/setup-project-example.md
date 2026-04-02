# Setup Project Example

Use this starter setup as the baseline for new projects built on the same stack.

## Required Project Files

- `.editorconfig`: shared editor defaults for whitespace, line endings, and indentation
- `.gitignore`: ignore local/editor/build artifacts before the first commit
- `eslint.config.js`: lint the hand-written frontend JavaScript
- `package.json`: optional Node-based tooling only for linting frontend JS
- `README.md`: quick start, route map, and project structure

## Example Setup Flow

1. Initialize the Go module:

   ```bash
   go mod init your-project-name
   ```

2. Add templ and sync dependencies:

   ```bash
   go get github.com/a-h/templ
   go mod tidy
   ```

3. Install frontend linting dependencies when working on custom JavaScript:

   ```bash
   npm install
   ```

4. Generate templates and run the app:

   ```bash
   templ generate
   go run ./cmd/web
   ```

5. Run frontend linting when JavaScript changes:

   ```bash
   npm run lint:js
   ```

## Rules For New Projects

- Keep Node tooling optional and limited to frontend quality checks unless the project truly needs a bundling pipeline.
- Add ignore rules before adding generated files or local tooling caches.
- Keep editor defaults committed so all contributors format files consistently from day one.
- Scope ESLint to hand-written application JavaScript, not vendored assets like `htmx.min.js`.
- Update the README and conventions docs whenever setup commands or required tools change.
