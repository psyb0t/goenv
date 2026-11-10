---
title: "Agent integrations"
description: "Install the goenv agent skill in Claude Code or Codex from the psyb0t marketplace, in OpenClaw from ClawHub, or let an agent read .agents/skills directly."
nav_order: 3
---

# Agent integrations

goenv ships an agent skill that teaches a coding agent the API, the exact dev/prod matching and when not to use the library. It lives in [`.agents/skills/goenv/`](https://github.com/psyb0t/goenv/tree/main/.agents/skills/goenv): [`SKILL.md`](https://github.com/psyb0t/goenv/blob/main/.agents/skills/goenv/SKILL.md) plus [`references/setup.md`](https://github.com/psyb0t/goenv/blob/main/.agents/skills/goenv/references/setup.md).

The skill works in any agent that reads `.agents/skills/` (such as [peen](https://github.com/psyb0t/peen)), and installs natively in the clients below.

## Claude Code

```bash
claude plugin marketplace add psyb0t/agents
claude plugin install goenv@psyb0t
```

The plugin manifest is [`.agents/.claude-plugin/plugin.json`](https://github.com/psyb0t/goenv/blob/main/.agents/.claude-plugin/plugin.json).

## Codex

```bash
codex plugin marketplace add psyb0t/agents
codex plugin add goenv@psyb0t
```

Installed via the marketplace, the skill invokes as `$goenv:goenv`. Codex also picks the skill up automatically, with no install, in any repo containing `.agents/skills/`, and there it invokes as plain `$goenv`. The plugin manifest is [`.agents/.codex-plugin/plugin.json`](https://github.com/psyb0t/goenv/blob/main/.agents/.codex-plugin/plugin.json).

## OpenClaw

The skill is published to ClawHub on every release:

```bash
openclaw skills install @psyb0t/goenv
```

The `publish-to-clawhub` job in the [pipeline](development.md#ci) pushes it after the GitHub release of each tag.

## What the skill can touch

The library reads `ENV` and nothing else, writes nothing and makes no network calls at runtime. The skill only guides an agent through adding the import and a few calls to your own Go code, plus `go get` to fetch the module.
