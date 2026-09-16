# Architecture

The current architecture of this solution. Loaded on demand, not at session
start — so it can be longer than `constraints.md`, but every section should
still be something an agent would act on.

## Architectural style

Layered, with an adapter (ports-and-adapters) boundary for ADE integration.

```text
CLI Layer
   |
   v
Application Services
   |
   +-- BootstrapService
   +-- ValidationService
   +-- InspectionService
   +-- SessionService
   +-- LifecycleService
   +-- AuditService
   |
   v
Domain
   |
   +-- CDADProject
   +-- FreezeState
   +-- Proposal
   +-- ChangeRequest
   +-- SessionState
   +-- Provenance
   |
   v
Infrastructure
   |
   +-- Filesystem
   +-- Git
   +-- Network
   +-- ADE adapters
```

The CLI is not a reasoning agent: the Application Services layer performs
deterministic operations only (bootstrap, validate, inspect, session,
lifecycle, audit). Architectural or content reasoning belongs to the AI
agent (Claude Code), one layer above this CLI, never inside it.

## Modules and boundaries

| Module | Responsibility | May depend on | Must not depend on |
|---|---|---|---|
| `cmd/cdad` | CLI entrypoint, command wiring (Cobra) | Application Services | Infrastructure directly, Domain internals |
| `internal/bootstrap` | Obtain/apply bootstrap artifacts, create CDAD structure | Domain, `internal/artifacts`, `internal/integration` | ADE-specific adapter logic |
| `internal/lifecycle` | Lifecycle event model and hooks | Domain, `internal/events` | `internal/integration` (ADE-specific code) |
| `internal/inspector` | Deeper structural inspection (`cdad inspect`) | Domain, `internal/git` | Network |
| `internal/validator` | Deterministic structural validation (`cdad validate`) | Domain | Network, LLM/agent calls |
| `internal/session` | Derive/maintain session handoff state | Domain, `internal/git` | Governed context write path |
| `internal/events` | Lifecycle event taxonomy and dispatch | Domain | `internal/integration` |
| `internal/integration` | ADE adapters (Claude Code first) | Domain (read-only) | CDAD governance logic (must live in core, not here) |
| `internal/config` | Typed configuration | — | Domain business logic |
| `internal/git` | Git state/log access | — | Domain business logic |
| `internal/artifacts` | Bootstrap artifact acquisition/versioning | `internal/config` | ADE-specific logic |
| `internal/audit` | Governance/traceability report generation | Domain | Network (unless explicitly fetching) |
| `internal/output` | Human-readable and `--format json` rendering | — | Domain business logic |

The boundary that matters most: **ADE adapters (`internal/integration`)
must not contain CDAD governance logic that belongs to the core**, and core
governance (`internal/lifecycle`, `internal/validator`, `internal/session`,
etc.) must never depend on Claude-Code-specific (or any other ADE-specific)
internals. This is what keeps CDAD artifacts portable across ADEs.

## Integration strategy

The CLI prepares the integration layer; the AI agent remains responsible
for AI interaction. Concretely, for `cdad init --agent claude`:

1. Detect the project.
2. Detect whether Claude Code configuration exists.
3. Obtain the appropriate CDAD bootstrap artifacts.
4. Create the required CDAD structure.
5. Create or update the CDAD/Claude Code integration artifacts.
6. Generate the CDAD bootstrap prompt/contract.
7. Provide that prompt to Claude Code through Claude Code's own normal
   integration mechanism — the CLI does not attempt to emulate Claude Code.
8. Claude Code analyzes the project and populates the appropriate CDAD
   artifacts.
9. The CLI validates the resulting structure.

ADE detection must be conservative: if no supported ADE is detected, CDAD
still initializes the project structurally, without an ADE integration.

## Data model ownership

There is no database. Governed and derived state both live as artifacts on
the project filesystem, under `cdad/`:

- **Governed context** (`cdad/context/`, `cdad/adr/`) — owned by the Solution
  Designer; the CLI validates it, never authors its content.
- **Derived state** (session handoff, `cdad/backlog.md` status fields) — CLI
  generates this from git state, lifecycle events, and project artifacts.
- **Historical/audit state** (`cdad/CDAD-COMPLETION.md`, audit reports) —
  append-only records the CLI writes.

The exact on-disk shape for session state (e.g. whether it lives under a
`.cdad-session/` directory or elsewhere) is **not decided by the source
design** — see open question in the source document, §60. Leave as pending
until resolved through the normal change process.

## Deployment topology

Standalone native binaries, no runtime dependency:

```text
Windows -> cdad.exe
Linux   -> cdad
macOS   -> cdad
```

Invoked from: VS Code integrated terminal, PowerShell, Windows Terminal,
Linux/macOS terminals, and CI/CD runners. The CLI is not a VS Code
extension and is not itself a service that gets deployed anywhere — it is
distributed as a binary and run locally or in CI.

## Known deviations

None. This is a pre-implementation design; no code exists yet, so there are
no known deviations between implementation and this architecture to record.

---
Governance: L0. Read-only for AI agents. Changes require an approved ADR.
