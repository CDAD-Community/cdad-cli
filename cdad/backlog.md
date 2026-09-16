# Project Backlog

> The project's living development line — Epics, Stories, and the work
> currently expected to be built. This is a development-planning artifact,
> not architecture: see *Precedence* below. It answers "what exists, what's
> next, what's blocked" — not "how may it be built" or "what decisions are
> authoritative."

**Last verified:** `2026-09-16` — run the `cdad-audit` skill to reconcile against defined requirements.

---

## Development Line

Initial backlog for CDAD CLI v3, populated from the source design
`CDAD_CLI_v3_Technical_Design.md` via
`cdad/proposals/PROPOSAL-initial-backlog-population.md` (approved by the
Solution Designer). No implementation exists yet — every Story below is
pre-work, not in progress.

## Precedence

```text
Governed Context / L0
        v
ADR / governed decisions
        v
This backlog
        v
Implementation work
```

If a Story appears to contradict governed context or an accepted ADR, that
is a finding, not a resolution. Raise it through `cdad/CHANGE-REQUEST.md` —
a Story never silently overrides architecture.

---

## Epics

### EPIC-01 — CDAD CLI Bootstrap

**Status:** Proposed
**Goal:** Let a developer install the CDAD CLI and initialize CDAD in a
project without cloning `cdad-bootstrap` or copying files manually.

#### Stories

##### STORY-01.1 — Install the CDAD CLI

- **Status:** Proposed
- **Priority:** Not specified in source
- **Description:** As a developer, I want to install the CDAD CLI so that I
  can use CDAD without cloning the bootstrap repository.
- **Acceptance Criteria:**
  - CLI can be installed as a standalone executable.
  - Windows/Linux/macOS are supported.
  - `cdad version` works.
- **Dependencies:** None
- **Notes:** Source §46 (EPIC-01, Story 01.1).

##### STORY-01.2 — Run `cdad init`

- **Status:** Proposed
- **Priority:** Not specified in source
- **Description:** As a developer, I want to run `cdad init` so that CDAD is
  initialized in my project.
- **Acceptance Criteria:**
  - Existing project detected.
  - CDAD structure created.
  - Version recorded.
  - Existing files are not silently overwritten.
- **Dependencies:** STORY-01.1
- **Notes:** Source §46 (EPIC-01, Story 01.2).

---

### EPIC-02 — ADE Integration

**Status:** Proposed
**Goal:** Detect the ADE (Claude Code first) and prepare the corresponding
integration automatically, isolated behind an adapter.

#### Stories

##### STORY-02.1 — Detect Claude Code

- **Status:** Proposed
- **Priority:** Not specified in source
- **Description:** As a developer, I want CDAD to detect Claude Code so that
  the appropriate integration is prepared automatically.
- **Acceptance Criteria:**
  - Claude Code detection is deterministic.
  - Integration is isolated in an adapter.
  - CDAD core does not depend on Claude Code internals.
- **Dependencies:** STORY-01.2
- **Notes:** Source §47 (EPIC-02, Story 02.1).

##### STORY-02.2 — Generate the Claude Code bootstrap instructions

- **Status:** Proposed
- **Priority:** Not specified in source
- **Description:** As a developer, I want CDAD to generate the bootstrap
  instructions for Claude Code so that the agent can populate CDAD
  artifacts.
- **Acceptance Criteria:**
  - Canonical instruction is versioned.
  - Instructions respect CDAD governance.
  - No architectural decision is made by the CLI.
- **Dependencies:** STORY-02.1
- **Notes:** Source §47 (EPIC-02, Story 02.2).

---

### EPIC-03 — Deterministic Validation

**Status:** Proposed
**Goal:** Provide governance validation that is deterministic, LLM-free, and
usable as a CI gate.

#### Stories

##### STORY-03.1 — `cdad validate` identifies governance inconsistencies

- **Status:** Proposed
- **Priority:** Not specified in source
- **Description:** As a developer, I want `cdad validate` to identify
  governance inconsistencies.
- **Acceptance Criteria:**
  - Validation is deterministic.
  - No LLM is required.
  - Exit code indicates result.
- **Dependencies:** STORY-01.2
- **Notes:** Source §48 (EPIC-03, Story 03.1).

##### STORY-03.2 — `cdad validate --ci` as a pipeline gate

- **Status:** Proposed
- **Priority:** Not specified in source
- **Description:** As a DevOps engineer, I want `cdad validate --ci` so that
  CDAD can operate as a pipeline gate.
- **Acceptance Criteria:**
  - Non-interactive.
  - Machine-readable output.
  - Stable exit codes.
- **Dependencies:** STORY-03.1
- **Notes:** Source §48 (EPIC-03, Story 03.2).

---

### EPIC-04 — Inspection

**Status:** Proposed
**Goal:** Give the developer deterministic visibility into current CDAD
state and structural diagnostics.

#### Stories

##### STORY-04.1 — `cdad status` shows current CDAD state

- **Status:** Proposed
- **Priority:** Not specified in source
- **Description:** As a developer, I want `cdad status` to show the current
  CDAD state.
- **Acceptance Criteria:**
  - Freeze status shown.
  - OPEN/BLOCKING shown.
  - Proposals shown.
  - Change Requests shown.
  - Session state shown.
- **Dependencies:** STORY-01.2
- **Notes:** Source §49 (EPIC-04, Story 04.1).

##### STORY-04.2 — `cdad inspect` provides deeper diagnostics

- **Status:** Proposed
- **Priority:** Not specified in source
- **Description:** As a developer, I want `cdad inspect` to provide deeper
  diagnostics.
- **Acceptance Criteria:**
  - Structural checks shown.
  - Warnings distinguished from failures.
  - Remediation guidance provided where deterministic.
- **Dependencies:** STORY-04.1
- **Notes:** Source §49 (EPIC-04, Story 04.2).

---

### EPIC-05 — Session Continuity

**Status:** Proposed
**Goal:** Maintain machine-generated session state so a developer can resume
work without hand-maintaining a handoff file, without that state ever
becoming a second source of truth.

#### Stories

##### STORY-05.1 — Automatic session state

- **Status:** Proposed
- **Priority:** Not specified in source
- **Description:** As a developer, I want CDAD to automatically maintain
  session state so that I can continue work later without manually
  maintaining a handoff file.
- **Acceptance Criteria:**
  - State derives from real project artifacts.
  - Agent does not need to write the authoritative state.
  - Handoff is not evidence.
  - Handoff does not override governed context.
- **Dependencies:** STORY-01.2
- **Notes:** Source §50 (EPIC-05, Story 05.1).

##### STORY-05.2 — `cdad resume`

- **Status:** Proposed
- **Priority:** Not specified in source
- **Description:** As a developer, I want `cdad resume` so that I can
  quickly understand where the project was left.
- **Acceptance Criteria:**
  - Last known state shown.
  - Active change shown.
  - Pending proposals shown.
  - OPEN items shown.
  - Latest handoff shown.
- **Dependencies:** STORY-05.1
- **Notes:** Source §50 (EPIC-05, Story 05.2).

---

### EPIC-06 — Lifecycle Automation

**Status:** Proposed
**Goal:** Expose a small lifecycle event model with configurable hooks that
automate around CDAD without bypassing governance.

#### Stories

##### STORY-06.1 — Lifecycle events

- **Status:** Proposed
- **Priority:** Not specified in source
- **Description:** As a CDAD user, I want lifecycle events so that CDAD can
  automate actions at important project moments (init, think, proposal,
  change-request, freeze, work, complete, session-update).
- **Acceptance Criteria:** Not specified beyond the event taxonomy example
  in the source document (§23, §51) — the exact canonical event schema is
  an open design question (source §60).
- **Dependencies:** STORY-01.2
- **Notes:** Source §51 (EPIC-06, Story 06.1).

##### STORY-06.2 — Configurable lifecycle hooks

- **Status:** Proposed
- **Priority:** Not specified in source
- **Description:** As a project owner, I want configurable lifecycle hooks
  so that CDAD can execute deterministic automation.
- **Acceptance Criteria:**
  - Hooks are explicit.
  - Hooks cannot bypass governance.
  - Hook failures are observable.
- **Dependencies:** STORY-06.1
- **Notes:** Source §51 (EPIC-06, Story 06.2).

---

### EPIC-07 — Audit

**Status:** Proposed
**Goal:** Produce a governance/traceability report of the project's CDAD
history.

#### Stories

##### STORY-07.1 — CDAD audit report

- **Status:** Proposed
- **Priority:** Not specified in source
- **Description:** As a project owner, I want a CDAD audit report so that I
  can understand the project's governance history.
- **Acceptance Criteria:**
  - Freeze history included.
  - ADR/change history included.
  - Proposals included.
  - Validation results included.
  - Machine-readable output supported.
- **Dependencies:** STORY-03.1
- **Notes:** Source §52 (EPIC-07, Story 07.1).

---

### EPIC-08 — Versioning and Migration

**Status:** Proposed
**Goal:** Let the CLI identify its own and the project's CDAD schema
version, and migrate between schema versions in a controlled way.

#### Stories

##### STORY-08.1 — Identify project CDAD version

- **Status:** Proposed
- **Priority:** Not specified in source
- **Description:** As a developer, I want the CLI to identify the CDAD
  version of my project.
- **Acceptance Criteria:** Not specified beyond the story statement itself
  in the source document.
- **Dependencies:** STORY-01.2
- **Notes:** Source §53 (EPIC-08, Story 08.1).

##### STORY-08.2 — Controlled schema migration

- **Status:** Proposed
- **Priority:** Not specified in source
- **Description:** As a developer, I want controlled migration between CDAD
  schema versions.
- **Acceptance Criteria:**
  - Migration is explicit.
  - Existing artifacts are backed up or recoverable.
  - Migration does not silently alter frozen decisions.
- **Dependencies:** STORY-08.1
- **Notes:** Source §53 (EPIC-08, Story 08.2).

---

### EPIC-09 — ADE Portability

**Status:** Proposed
**Goal:** Keep CDAD core independent of Claude Code so other ADEs (Kiro,
Copilot, Codex) can be added later without touching governance.

#### Stories

##### STORY-09.1 — CDAD independent of Claude Code

- **Status:** Proposed
- **Priority:** Not specified in source
- **Description:** As a developer, I want CDAD to remain independent of
  Claude Code so that I can use another ADE later.
- **Acceptance Criteria:**
  - Core governance does not depend on Claude Code.
  - ADE integration is adapter-based.
  - CDAD artifacts remain portable.
- **Dependencies:** STORY-02.1
- **Notes:** Source §54 (EPIC-09, Story 09.1). Future adapters named: Claude
  Code, Kiro, GitHub Copilot, Codex, other ADEs.

---

### EPIC-10 — Security

**Status:** Proposed
**Goal:** Ensure the CLI never exposes secrets and treats network operations
as explicit and safe by default.

#### Stories

##### STORY-10.1 — Avoid exposing secrets

- **Status:** Proposed
- **Priority:** Not specified in source
- **Description:** As a developer, I want CDAD CLI to avoid exposing secrets
  so that project bootstrap and validation are safe.
- **Acceptance Criteria:**
  - Credentials are not logged.
  - Source is not uploaded by default.
  - Network operations are explicit.
  - Downloaded artifacts can be integrity-checked.
- **Dependencies:** STORY-01.2
- **Notes:** Source §55 (EPIC-10, Story 10.1).

---

## General Development Work

Work tracked here that is not tied to a specific Story.

- None

## Current Focus

What is actively being worked on right now.

- STORY-01.1 (Install the CDAD CLI) — Go module and Cobra CLI skeleton
  created (`go.mod`, `cmd/cdad/main.go`, `cdad version` working).
- STORY-01.2 (`cdad init`) — mechanical part implemented: detects the
  target directory, creates `cdad/{adr,context,docs,proposals,scripts}/`
  and a starter `cdad/backlog.md`, never overwrites existing paths, and
  records each run in `cdad/CDAD-COMPLETION.md`. Deliberately does **not**
  create the portable-core files (`AGENTS.md`, `README-CDAD.md`/`.es.md`)
  or the generic `cdad/` docs (`INDEX.md`, `CHANGE-REQUEST.md`,
  `INSTALLATION*`, `USAGE*`, `ADR-TEMPLATE.md`) — sourcing that canonical
  content is `internal/artifacts` (bootstrap artifact acquisition), which
  `cdad/context/stack.md` still marks pending/undecided. Reported as
  "Pending" by `cdad init` rather than guessed at.
- STORY-03.1 (`cdad validate`) — backlog structural-integrity check
  (`internal/validator`) ported natively from `cdad/scripts/cdad-check-backlog.sh`:
  duplicate Epic/Story ID detection, status-vocabulary enforcement, empty-Epic
  warnings; exit 0/1/2 matches the shell script's contract. Unit-tested
  (`go test ./internal/validator/...`). While porting, found and fixed a real
  bug in `cdad-check-backlog.sh` itself: its Story-ID regex truncated at the
  first `.`, so `STORY-01.1`/`STORY-01.2` collapsed to the same key and were
  flagged as false-positive duplicates against this project's own backlog —
  fixed in the script too. Other `cdad validate` checks (adapter matrix,
  stack-map freshness) are not yet ported: adapter matching needs ADE
  detection (EPIC-02), and stack-map freshness needs git integration
  (`internal/git`, not yet built).

## Next Work

What comes after Current Focus.

- None declared yet. Per §61 of the source design, EPIC-01 (bootstrap) is
  the natural starting point since every other Epic's Stories depend on
  `cdad init` (STORY-01.2) existing first — informational, not a decided
  sequencing.

## Blocked

- None

---
Governance: adding or removing an Epic/Story, or materially changing its
scope or acceptance criteria, goes through `cdad/CHANGE-REQUEST.md` ->
`cdad/proposals/` -> Solution Designer decision — the same funnel as an
architecture change. Updating a Story's status, or the *Current Focus* /
*Next Work* / *Blocked* lists, as part of already-approved implementation
work does not need a change request. See `AGENTS.md` → *Backlog governance*
for the full rule and precedence relative to L0/ADRs.
