# Principles

Design principles in force. A principle earns its place only if it rules
something out. If a principle would never cause you to reject a pull request,
delete it.

## Format

Each principle: the rule, then the trade-off it accepts.

- **<Principle>** — <what it rules out>. Accepts: <cost of holding this line>.

## Principles in force

- **The CLI never performs architectural reasoning** — rules out adding
  heuristics or LLM calls inside the CLI to judge whether a design is
  correct. Accepts: some questions ("is this the right architecture?")
  stay permanently unanswerable by the CLI itself, even when automating them
  would be convenient.

- **Structural validation is deterministic and LLM-free** — rules out
  `validate`, `inspect`, `status`, and `audit` depending on network calls or
  an LLM for any check they perform. Accepts: these commands can confirm
  structural invariants (a marker exists, a reference resolves, a gap is
  unresolved) but never semantic or architectural correctness.

- **No `unfreeze`, ever** — rules out any command, flag, or code path that
  reverses a freeze. Accepts: every post-freeze change, however small, goes
  through change-request -> proposal -> human decision -> new freeze; there
  is no shortcut for urgent fixes.

- **ADE adapters stay isolated from CDAD core governance** — rules out
  `internal/lifecycle`, `internal/validator`, `internal/session`, and the
  rest of the governance core importing Claude-Code-specific (or any other
  ADE-specific) types or logic. Accepts: some duplication across adapters
  instead of one adapter-shaped core; this is what keeps CDAD artifacts
  portable across ADEs.

- **Session handoff is orientation, never evidence** — rules out session
  state influencing the outcome of `freeze` or `validate`, or overriding L0,
  ADRs, constraints, or governed context in any way. Accepts: a developer
  must re-derive or re-confirm real state at freeze time rather than trust
  the handoff blindly.

- **Offline by default** — rules out requiring network access for `status`,
  `inspect`, `validate`, or `audit`. Accepts: bootstrap acquisition, version
  discovery, and updates are the only commands allowed to need a network.

## Anti-examples

"Write clean code", "prefer simplicity", "follow best practices". These rule
nothing out and cost context tokens to carry.

---
Governance: L0. Read-only for AI agents.
