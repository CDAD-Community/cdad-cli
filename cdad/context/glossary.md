# Glossary

Domain terms whose meaning in this project differs from the everyday meaning, or
that are ambiguous across teams. Do not define general industry terms.

| Term | Meaning in this project | Not to be confused with |
|---|---|---|
| Freeze | The mechanical CDAD operation that locks governed architecture context after an explicit human decision. There is no `unfreeze`; a mistaken freeze is corrected only through change-request → proposal → new freeze. | A temporary release/code freeze that gets lifted later. |
| ADE | The environment actually executing the AI agent (Claude Code, Kiro, GitHub Copilot, Codex) — detected from the execution environment, never from the underlying model. | The underlying LLM/model itself (a Claude model is not Claude Code). |
| Session Handoff | The CLI's machine-derived operational continuity artifact (`cdad resume`). Explicitly not evidence and never overrides governed context. | Claude Code's own memory/instruction mechanisms, or a source of architectural truth. |
| OPEN / BLOCKING | Governance states of an unresolved decision: OPEN does not block freeze, BLOCKING does. | Generic issue-tracker severity labels. |
| Proposal | A marked, not-yet-decided architectural suggestion produced by an agent. Must never be silently converted into a ratified decision. | An accepted/ratified ADR. |
| Governance Boundary | The line between what the CLI can determine structurally (e.g. "does this provenance marker reference a real file/line?") and what only a human can judge (e.g. "is this architecture correct?"). | A permissions/security boundary. |

---
Governance: L0. Read-only for AI agents.
