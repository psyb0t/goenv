# Docs and releases

Prefix: DOC.

- **DOC1:** `README.md` is the front door: pitch, install, a usage example, a short summary of each topic and the Documentation section. Reference material (API tables, behavior table, agent install routes, development and CI) lives in `docs/`. Keep each fact in one place and link to it.
- **DOC2:** The README links to the docs site with full `https://psyb0t.github.io/goenv/...` URLs and to repository files with full `https://github.com/psyb0t/goenv/blob/main/...` URLs, because the README is also shown on pkg.go.dev and the mirrors.
- **DOC3:** Every page in `docs/` has front matter with a unique `title` and a `description` of 50 to 170 characters, and one `#` heading. Page-to-page links use the relative `.md` path. A link to a file outside `docs/` uses the full GitHub URL.
- **DOC4:** After a docs change, run `make docs` and check the output in `.testing/docs-site`: every internal link and anchor resolves, every GitHub link points at a file that exists on `main`, and the build printed no warnings.
- **DOC5:** Prose keeps the README's voice: blunt and profane where it already is, first person singular when the owner speaks, no em dashes, "directory" never "folder".
- **DOC6:** Releases are annotated `vX.Y.Z` tags on a `main` commit that passed CI, with the version's CHANGELOG section as the tag message. A tag push does not rerun the checks: the pipeline waits for that commit's `main` run, then cuts the GitHub release and publishes the skill. There is no version file; the tag is the version.
