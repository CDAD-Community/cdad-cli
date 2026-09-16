# Constraints

Hard limits this solution must respect. This file is loaded into every AI
session, so keep it short and concrete: only constraints that would change a
decision. Delete every line you have not actually committed to.

## Platform

- Cloud provider: not applicable — the CLI is a standalone local/CI tool, not a hosted service.
- Compute model: standalone native binary, invoked as a local process (developer terminal or CI runner).
- IaC tooling: none — nothing is provisioned; the CLI is distributed as a binary, not deployed infrastructure.

## Technical

- Runtime and language version: Go — exact minimum version not pinned by the source design; pending.
- Datastore: none — governed and derived state live as files under the project's `cdad/` directory, never a database.
- Communication style: local filesystem/process I/O for all offline commands (`status`, `inspect`, `validate`, `audit`); HTTPS only, and only when explicitly needed, for bootstrap artifact acquisition, version discovery, and updates.

## Regulatory and organizational

- Data residency: not applicable — the CLI does not transmit project data by default.
- Compliance regime: none specified by the source design.
- Budget or quota ceilings that constrain design: none specified by the source design.

## Explicitly out of scope

- Becoming an autonomous software architect or making architectural decisions.
- Deciding whether a proposal should be accepted, or replacing the human decision point.
- Resolving architectural conflicts automatically.
- Acting as the source of truth for architecture (that is `cdad/context/`).
- Replacing Claude Code, Kiro, or GitHub Copilot, or becoming an ADE itself.
- Any `unfreeze` operation.
- Bypassing CDAD governance or silently modifying frozen architecture.
- Uploading project source code to external services by default.
- Sending source code to external services unless explicitly configured.
- Exposing credentials, or copying secrets into generated prompts.
- Requiring network access for ordinary validation/inspection commands.

---
Governance: L0. Read-only for AI agents. Changes require Solution Designer
approval via the `cdad-propose-change` skill.
