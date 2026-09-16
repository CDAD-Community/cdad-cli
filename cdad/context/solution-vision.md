# Solution Vision

## Problem

Adopting CDAD into a project today requires manual work: cloning
`cdad-bootstrap`, copying directories by hand, and understanding CDAD's
internal context-layer structure and how the initial agent prompt is built.
That friction falls on every developer who wants a CDAD-governed project,
and it costs them time and correctness — a manual copy is exactly the kind
of step where drift and inconsistency start before governance has even
begun.

## Users

Developers who want to bring an existing (or new) software project under
CDAD governance, working from a terminal — VS Code's integrated terminal is
the first-class case, but any shell (PowerShell, Windows Terminal, Linux,
macOS, CI/CD runners) is supported. The primary AI agent they pair with is
Claude Code; the adapter architecture anticipates Kiro, GitHub Copilot, and
Codex later, but only Claude Code is a v3 target.

## What success looks like

- Running `cdad init` immediately yields a CDAD-governed project: the
  developer does not need to know where `cdad-bootstrap` lives, which files
  to copy, or how the context layers are structured.
- `cdad status`, `cdad inspect`, `cdad validate`, and `cdad audit` work
  offline and produce deterministic output — no LLM required.
- `cdad validate --ci` functions as a pipeline gate with stable exit codes
  and machine-readable output.
- Freeze is mechanical and irreversible: the CLI enforces freeze
  prerequisites, but never decides whether the architecture itself is
  correct, and never offers an `unfreeze`.
- A developer can stop work, come back later (`cdad resume`), and quickly
  understand where the project was left without maintaining a handoff file
  by hand.

## Non-goals

The CLI deliberately does not become:

- an autonomous software architect;
- the entity that decides whether a proposal is accepted;
- a replacement for the human decision point;
- an automatic resolver of architectural conflicts;
- the source of truth for architecture (that remains `cdad/context/`);
- a replacement for Claude Code, Kiro, or GitHub Copilot;
- an ADE in its own right;
- a way to bypass CDAD governance or silently modify frozen architecture;
- a tool that offers `unfreeze` under any circumstance.

This section exists to stop scope creep from being rationalized later as a
natural extension of "just inspecting" or "just validating."

---
Governance: L0. Read-only for AI agents.
