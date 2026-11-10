# Agent surface

Prefix: AGT. The skill and the plugin manifests under `.agents/` describe this library to other agents. They are public.

- **AGT1:** An API or behavior change updates `.agents/skills/goenv/SKILL.md` and `.agents/skills/goenv/references/setup.md` in the same change. Both restate the API table and the behavior table from `goenv.go`; take them from the code, never from another document.
- **AGT2:** The Codex manifest `version` in `.agents/.codex-plugin/plugin.json` moves with the release tag, in the same release commit. The Claude manifest in `.agents/.claude-plugin/plugin.json` carries no version.
- **AGT3:** The `description` in both manifests stays identical.
- **AGT4:** The skill keeps its `Security & safety` section accurate: it reads only `ENV`, writes nothing and makes no network calls at runtime. Update it if that ever stops being true.
- **AGT5:** The pipeline publishes the skill to ClawHub on tag pushes with `skill_names: '["goenv"]'`. A new or renamed skill directory updates that list.
