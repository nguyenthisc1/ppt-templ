const js = require("@eslint/js");

module.exports = [
  {
    ignores: [
      "node_modules/**",
      "dist/**",
      "coverage/**",
      "public/assets/js/htmx.min.js",
    ],
  },
  js.configs.recommended,
  {
    files: ["public/assets/js/**/*.js"],
    languageOptions: {
      ecmaVersion: 2022,
      sourceType: "script",
      globals: {
        document: "readonly",
        window: "readonly",
      },
    },
    rules: {
      "no-unused-vars": ["error", { args: "none" }],
    },
  },
];
