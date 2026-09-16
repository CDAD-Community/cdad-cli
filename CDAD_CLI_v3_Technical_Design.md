# CDAD CLI — Vision, Scope & Technical Design

**Project:** Context-Driven AI Development (CDAD)  
**Component:** `cdad` Command Line Interface  
**Target:** CDAD v3 / v3.x  
**Primary ADE:** Claude Code running inside VS Code  
**Design principle:** CDAD governs; the AI agent reasons and executes within the governed context.

---

# 1. Executive Summary

The CDAD CLI is the operational entry point for CDAD.

The user should not need to clone `cdad-bootstrap`, copy directories manually, or understand the internal implementation structure of CDAD.

The desired experience is:

```bash
cdad init
```

From this point, the CLI prepares the project, obtains the required CDAD implementation/artifacts, detects the development environment and ADE, prepares the corresponding integration, and provides the governed bootstrap instructions to the AI agent.

The CLI is **not an AI agent** and does not conduct architectural reasoning.

Its responsibility is deterministic:

- bootstrap;
- inspect;
- validate;
- manage lifecycle events;
- maintain derived operational state;
- generate session handoff state;
- integrate with CI/CD;
- expose CDAD governance status.

Claude Code performs reasoning and implementation work under CDAD governance.

---

# 2. Vision

## 2.1 User vision

A developer should be able to enter an existing software project and run:

```bash
cdad init
```

and immediately obtain a CDAD-governed project.

The user should not need to know:

- where `cdad-bootstrap` lives;
- which files need to be copied;
- which Claude Code files need to be created;
- how CDAD context layers are structured;
- how the initial agent prompt is constructed.

CDAD CLI hides this operational complexity.

## 2.2 Conceptual vision

```text
                 SOFTWARE PROJECT
                        │
                        ▼
                   cdad init
                        │
          ┌─────────────┼─────────────┐
          │             │             │
          ▼             ▼             ▼
      Bootstrap       ADE          Context
          │          Detection      Setup
          │             │             │
          └─────────────┼─────────────┘
                        ▼
                Governed Project
                        │
                        ▼
                  Claude Code
                        │
                        ▼
                 CDAD-governed work
```

---

# 3. Core Principle

The CLI must preserve the CDAD separation of responsibilities.

```text
CDAD CLI
    │
    ├── prepares
    ├── verifies
    ├── observes
    ├── validates
    └── orchestrates lifecycle
             │
             ▼
        AI Agent / ADE
             │
             ├── reasons
             ├── analyzes
             ├── proposes
             └── implements
```

The CLI does not become another reasoning agent.

The CDAD governance model remains:

```text
Evidence
   ↓
Grounding
   ↓
Dossier
   ↓
Workers / THINK Agents
   ↓
Marked Proposals
   ↓
Human Decision
   ↓
Freeze
```

---

# 4. Relationship with `cdad-bootstrap`

`cdad-bootstrap` is the reference implementation.

The CLI becomes the user-facing distribution mechanism.

The desired model is:

```text
CDAD Methodology
       │
       ▼
Reference Implementation
cdad-bootstrap
       │
       ▼
CDAD CLI
       │
       ▼
Developer Project
```

The developer normally interacts with:

```bash
cdad
```

rather than directly cloning:

```text
cdad-bootstrap
```

The CLI may obtain the appropriate bootstrap version from a release, package, registry, or other official distribution mechanism.

The exact distribution mechanism is an implementation decision.

---

# 5. Goals

## Primary goals

- Provide a single entry point to CDAD.
- Eliminate manual cloning of `cdad-bootstrap`.
- Bootstrap CDAD into existing projects.
- Detect the development environment.
- Detect supported ADE integrations.
- Prepare Claude Code integration.
- Provide deterministic CDAD validation.
- Provide project inspection.
- Support CDAD lifecycle events.
- Generate machine-derived session state.
- Support session continuity across days and sessions.
- Integrate with CI/CD.
- Preserve CDAD governance boundaries.
- Remain portable across AI development environments.

## Secondary goals

- Provide version management.
- Provide migration mechanisms.
- Provide diagnostics.
- Provide audit information.
- Provide automation hooks.
- Provide a stable interface for future ADE integrations.

---

# 6. Non-Goals

The CLI must not:

- become an autonomous software architect;
- make architectural decisions;
- decide whether a proposal should be accepted;
- replace the human decision point;
- resolve architectural conflicts automatically;
- act as the source of truth for architecture;
- replace Claude Code;
- replace Kiro;
- replace Copilot;
- become an ADE;
- allow `unfreeze`;
- bypass CDAD governance;
- silently modify frozen architecture.

The existing CDAD rule remains:

```text
No unfreeze.

CHANGE REQUEST
      ↓
IMPACT ANALYSIS
      ↓
THINK
      ↓
HUMAN DECISION
      ↓
NEW FREEZE
```

---

# 7. Primary Target Environment

## v3 target

The first-class implementation target is:

```text
VS Code
   │
   └── Integrated Terminal
           │
           ▼
        cdad CLI
           │
           ▼
      Claude Code
```

The CLI is not a VS Code extension.

It is a standalone executable that can be invoked from:

- VS Code integrated terminal;
- PowerShell;
- Windows Terminal;
- Linux terminal;
- macOS terminal;
- CI/CD runners.

VS Code is therefore an important development environment, not a runtime dependency.

---

# 8. Claude Code Integration

## 8.1 Principle

CDAD must govern Claude Code without becoming coupled to Claude Code as a methodology.

The CLI prepares the integration layer.

Claude Code remains responsible for AI interaction.

Conceptually:

```text
              CDAD
               │
        Governance Layer
               │
               ▼
         CDAD CLI
               │
        Integration Layer
               │
               ▼
          Claude Code
               │
               ▼
         Project Work
```

## 8.2 Bootstrap

Example:

```bash
cdad init --agent claude
```

The CLI:

1. Detects the project.
2. Detects whether Claude Code configuration exists.
3. Obtains the appropriate CDAD bootstrap artifacts.
4. Creates the required CDAD structure.
5. Creates or updates the CDAD/Claude integration artifacts.
6. Generates the CDAD bootstrap prompt/contract.
7. Provides that prompt to Claude Code.
8. Claude Code analyzes the project and populates the appropriate CDAD artifacts.
9. CLI validates the resulting structure.

---

# 9. Agent Bootstrap Prompt

The CLI may generate a canonical bootstrap instruction for Claude Code.

Conceptually:

```text
You are operating in a CDAD-governed project.

Follow the CDAD governance artifacts available in this repository.

Analyze the existing project and populate the required CDAD artifacts.

Do not invent facts.

When evidence is unavailable:
    mark the appropriate gap.

When evidence conflicts:
    expose the conflict.

When proposing new architectural content:
    mark it as a proposal.

Do not convert proposals into decisions.

Respect frozen CDAD context.

Use the appropriate CDAD proposal/change-request path for
post-freeze changes.
```

The exact prompt is maintained as part of the CDAD implementation and versioned independently from individual agent sessions.

The CLI should preferably provide the contract to Claude Code through the normal Claude Code integration mechanism rather than attempting to emulate Claude Code.

---

# 10. CLI Command Model

The initial command surface should remain small.

## Core commands

```bash
cdad init
cdad status
cdad inspect
cdad validate
cdad freeze
cdad resume
cdad doctor
cdad audit
```

## Lifecycle commands

```bash
cdad event <event>
cdad change
cdad proposal
```

## Information

```bash
cdad version
cdad help
```

Future commands may be added only when they represent a stable CDAD capability.

---

# 11. Command Responsibilities

## `cdad init`

Bootstrap CDAD.

Responsibilities:

- detect project;
- detect existing CDAD installation;
- detect ADE;
- select CDAD mode;
- obtain bootstrap artifacts;
- create CDAD structure;
- prepare integration;
- generate bootstrap instruction;
- validate initialization.

Example:

```bash
cdad init
```

Interactive mode may ask:

```text
Detected project: my-api

Detected ADE:
  Claude Code

CDAD mode:
  1. Prompt-only
  2. Existing design/documentation
  3. Professional methodology

Select:
```

---

# 12. `cdad status`

Show current CDAD state.

Example:

```text
CDAD Status
────────────────────────────
Project: my-api
CDAD: v3
Mode: Professional

Freeze: ACTIVE
Last Freeze: 2026-09-16

OPEN: 3
BLOCKING: 0
Proposals: 2
Change Requests: 1

Session Handoff: CURRENT

Validation: OK
```

Status is deterministic.

It should be derived from project artifacts rather than AI-generated prose.

---

# 13. `cdad inspect`

Provide a deeper inspection.

Example:

```bash
cdad inspect
```

Possible output:

```text
CDAD Project Inspection

[OK] CDAD structure
[OK] Frozen context
[OK] AGENTS.md
[OK] ADR references
[WARN] 2 OPEN decisions
[WARN] 1 pending proposal
[OK] No blocking gaps
[OK] Session state current
```

---

# 14. `cdad validate`

Validate CDAD invariants and structure.

Potential checks:

- required artifacts;
- directory structure;
- valid references;
- provenance markers;
- unresolved conflicts;
- unresolved BLOCKING gaps;
- malformed proposals;
- invalid frozen state;
- invalid ADR references;
- stale session state;
- configuration consistency.

Example:

```bash
cdad validate
```

Exit codes should be deterministic.

Example:

```text
0 = valid
1 = validation failure
2 = configuration/error
```

Exact exit-code semantics should be formally specified.

---

# 15. `cdad validate --ci`

CI/CD-oriented validation.

Example:

```bash
cdad validate --ci
```

The command must be:

- non-interactive;
- deterministic;
- machine-readable when requested;
- suitable for pipeline gates.

Possible:

```bash
cdad validate --format json
```

---

# 16. `cdad freeze`

The CLI performs the mechanical freeze operation and validates the freeze conditions.

The CLI does not decide whether the architecture is correct.

The human decision remains outside the CLI.

Conceptually:

```text
Human Decision
      ↓
cdad freeze
      ↓
Deterministic checks
      ↓
Freeze artifact/state
```

The CLI must reject freeze when governance conditions fail.

Examples:

```text
BLOCKING gap exists       → FAIL
Unresolved conflict       → FAIL
Unresolved proposal       → FAIL
OPEN ratified             → ALLOWED
Required provenance       → REQUIRED
```

---

# 17. `cdad resume`

Prepare the environment for continuing work.

Example:

```bash
cdad resume
```

It should surface:

- current project state;
- last freeze;
- active change request;
- pending proposals;
- OPEN items;
- latest session handoff;
- next known work.

The output is orientation, not authority.

The handoff is never evidence.

---

# 18. Session Handoff

CDAD should maintain a machine-generated session continuity mechanism.

Terminology:

```text
Session State
Session Handoff
Context Resumption
Session Continuity
```

The preferred artifact concept is:

```text
Session Handoff Brief
```

## Important rule

The handoff must not become a second source of truth.

It must never override:

```text
L0
ADR
constraints
governed context
freeze state
```

It is operational history/orientation.

---

# 19. Automatic Session State

The session state should be generated by CDAD rather than requiring the human to maintain it manually.

The service can periodically derive it from:

- git state;
- git log;
- current branch;
- last freeze;
- proposals;
- change requests;
- ADRs;
- CDAD status;
- active work;
- project artifacts.

Example:

```text
SESSION STATE
────────────────────────

Last Freeze:
  2026-09-16 12:42

Current Branch:
  feature/api-consumer

Active Change Request:
  CR-004

Pending Proposals:
  2

OPEN:
  Cache strategy

Last Activity:
  API consumer implementation

Next Known Step:
  Validate external API contract
```

This allows:

```text
Day 1
Developer stops
       ↓
CDAD generates state
       ↓
Day 2
cdad resume
       ↓
Developer continues
```

---

# 20. Why CDAD Session State Does Not Replace Claude Code Memory

Claude Code may have its own mechanisms for instructions, rules, and continuity.

CDAD should not attempt to replace them.

The distinction is:

```text
Claude Code mechanisms
    ↓
Agent-specific behavior/context

CDAD Session State
    ↓
CDAD project lifecycle state
```

The CDAD handoff should therefore be:

- portable;
- ADE-independent;
- derived;
- non-authoritative;
- excluded from grounding.

---

# 21. `cdad doctor`

Diagnostic command.

Example:

```bash
cdad doctor
```

Checks:

```text
CDAD CLI
[OK] Version
[OK] Project detected
[OK] Git detected
[OK] CDAD structure
[OK] Claude Code integration
[WARN] Session state stale
[OK] Bootstrap version
```

It should explain remediation where possible.

---

# 22. `cdad audit`

Generate a governance/traceability report.

Potential contents:

```text
Project
CDAD version
Initial bootstrap
Freeze history
ADR history
Change Requests
Proposals
OPEN decisions
BLOCKING decisions
Validation history
Session continuity
```

Possible formats:

```bash
cdad audit
cdad audit --format json
cdad audit --format markdown
```

---

# 23. Lifecycle Events

The CLI should expose a small event model.

Potential events:

```text
project.init
think.started
think.completed
proposal.created
change.requested
freeze.started
freeze.completed
work.started
work.completed
session.updated
project.completed
```

The exact event taxonomy should be formalized before implementation.

The purpose is automation.

Example:

```text
EVENT
  freeze.completed
       │
       ├── update session state
       ├── generate audit entry
       └── trigger configured hooks
```

---

# 24. Event Hooks

CDAD may allow configured actions associated with lifecycle events.

Example:

```yaml
events:
  freeze.completed:
    - update-session
    - validate
    - audit

  work.completed:
    - validate
    - update-session
```

The hook system must not bypass governance.

Hooks are automation, not authorization.

---

# 25. Architecture

## High-level

```text
                    ┌──────────────────────┐
                    │      CDAD CLI        │
                    │       Go             │
                    └──────────┬───────────┘
                               │
             ┌─────────────────┼─────────────────┐
             │                 │                 │
             ▼                 ▼                 ▼
       Lifecycle          Inspector          Validator
             │                 │                 │
             └─────────────────┼─────────────────┘
                               │
                               ▼
                     CDAD Project State
                               │
             ┌─────────────────┼─────────────────┐
             ▼                 ▼                 ▼
          Context           ADRs             Proposals
             │                 │                 │
             └─────────────────┼─────────────────┘
                               ▼
                         Claude Code
                               │
                               ▼
                          Project Code
```

---

# 26. Internal CLI Architecture

Recommended modules:

```text
cmd/
  cdad/

internal/
  bootstrap/
  lifecycle/
  inspector/
  validator/
  session/
  events/
  integration/
  config/
  git/
  artifacts/
  audit/
  output/
```

Potential architecture:

```text
CLI Layer
   │
   ▼
Application Services
   │
   ├── BootstrapService
   ├── ValidationService
   ├── InspectionService
   ├── SessionService
   ├── LifecycleService
   └── AuditService
   │
   ▼
Domain
   │
   ├── CDADProject
   ├── FreezeState
   ├── Proposal
   ├── ChangeRequest
   ├── SessionState
   └── Provenance
   │
   ▼
Infrastructure
   │
   ├── Filesystem
   ├── Git
   ├── Network
   └── ADE adapters
```

---

# 27. Technology Stack

## Primary language

**Go**

Reasons:

- native binaries;
- Windows/Linux/macOS;
- fast startup;
- simple distribution;
- no runtime dependency;
- strong CLI ecosystem;
- good filesystem/process/network support;
- suitable for CI/CD;
- easy static builds.

Target binaries:

```text
Windows → cdad.exe
Linux   → cdad
macOS   → cdad
```

---

# 28. Recommended Go Stack

Potential libraries:

```text
CLI:
Cobra

Configuration:
Viper or simple typed configuration

Terminal UX:
Charm ecosystem / Lip Gloss where appropriate

Validation:
Custom deterministic validation layer

Git:
go-git or Git subprocess integration

HTTP:
Go standard library

Serialization:
encoding/json
YAML library where required

Testing:
Go testing
```

Library selection should remain deliberately minimal.

CDAD should avoid unnecessary dependencies.

---

# 29. Distribution

Desired developer experience:

```bash
install cdad
```

Then:

```bash
cdad init
```

Distribution options:

```text
GitHub Releases
Homebrew
Scoop
winget
package managers
installer scripts
```

The exact initial distribution channel is an implementation decision.

The binary should be versioned.

Example:

```bash
cdad version
```

Output:

```text
CDAD CLI
Version: 3.0.0
CDAD Schema: 3
```

---

# 30. Bootstrap Versioning

The CLI must distinguish:

```text
CLI version
CDAD methodology version
Bootstrap version
Project schema version
```

Example:

```text
CLI:       3.1.0
Methodology: 3
Bootstrap:  3.0.4
Schema:      3
```

This prevents the CLI from silently changing a project structure.

---

# 31. ADE Detection

The first implementation targets Claude Code.

Architecture should allow future adapters:

```text
ADE Adapter
    │
    ├── Claude Code
    ├── Kiro
    ├── GitHub Copilot
    ├── Codex
    └── Other
```

Detection must be conservative.

If no ADE is detected:

```text
No supported ADE detected.

CDAD can still initialize the project.
```

The methodology remains independent of any specific ADE.

---

# 32. Claude Code Adapter

Potential responsibilities:

```text
detect Claude Code
detect project Claude configuration
prepare integration
provide CDAD instructions
provide bootstrap prompt
validate expected integration artifacts
```

The adapter must not contain CDAD governance logic that belongs to the core.

```text
Core CDAD
   │
   ├── governance
   ├── validation
   ├── lifecycle
   └── state
        │
        ▼
ADE Adapter
   │
   └── Claude Code integration
```

---

# 33. Security Model

The CLI must assume that project files may contain sensitive information.

Principles:

- do not upload project source code by default;
- do not send source code to external services unless explicitly configured;
- do not expose credentials;
- do not copy secrets into generated prompts;
- avoid logging sensitive values;
- validate downloaded bootstrap artifacts;
- use HTTPS;
- verify release integrity where practical.

The CLI should not require network access for ordinary validation commands.

---

# 34. Offline Mode

Commands such as:

```bash
cdad status
cdad inspect
cdad validate
cdad audit
```

should work offline whenever possible.

Network access should primarily be required for:

```text
bootstrap acquisition
version discovery
updates
```

This improves reliability and CI compatibility.

---

# 35. Machine-Readable Output

Commands should support:

```bash
--format json
```

Example:

```bash
cdad status --format json
```

This enables:

```text
CI/CD
scripts
IDE integrations
dashboards
automation
```

The human-readable output remains the default.

---

# 36. Exit Codes

The CLI should provide stable exit semantics.

Conceptual model:

```text
0 → successful operation
1 → governance/validation failure
2 → CLI/configuration error
3 → environment/integration error
```

The exact specification should be finalized before implementation.

---

# 37. Project Structure

The CLI should bootstrap the CDAD structure defined by the current CDAD specification.

Conceptually:

```text
project/
│
├── cdad/
│   ├── context/
│   ├── adr/
│   ├── proposals/
│   ├── docs/
│   └── ...
│
├── AGENTS.md
├── ...
└── application source
```

The exact v3 artifact structure remains governed by the CDAD specification and must not be duplicated independently inside the CLI.

The CLI should consume a versioned bootstrap definition.

---

# 38. Governance Boundary

The CLI validates structural consequences of governance.

It does not judge architectural correctness.

Example:

```text
Question:
"Is Redis the correct architecture?"

CLI:
    Cannot decide.

Question:
"Is there an unresolved BLOCKING gap?"

CLI:
    Can determine structurally.

Question:
"Does this provenance marker reference a real file/line?"

CLI:
    Can determine structurally.

Question:
"Is this architecture a good design?"

CLI:
    Human decision.
```

This follows the CDAD distinction between deterministic verification and human architectural judgment.

---

# 39. Provenance Validation

The validator should check structural provenance.

Example:

```text
[FUENTE: file:line]
[VACÍO]
[CONFLICTO]
[PROPUESTA]
```

Potential deterministic checks:

```text
[ ] marker exists
[ ] referenced file exists
[ ] referenced line exists
[ ] required conflict citations exist
[ ] unresolved proposals are not treated as frozen decisions
```

Semantic correctness remains outside deterministic validation.

---

# 40. Freeze Validation

Before freeze:

```text
Validate
   │
   ├── BLOCKING gaps?
   ├── unresolved conflicts?
   ├── unresolved proposals?
   ├── invalid provenance?
   └── invalid structure?
```

Then:

```text
FAIL
```

or:

```text
READY FOR FREEZE
```

The CLI must never fabricate information to make the freeze pass.

---

# 41. Session State Architecture

Recommended distinction:

```text
                CDAD STATE
                    │
       ┌────────────┼────────────┐
       ▼            ▼            ▼
   Governed      Derived      Historical
    Context       State        Handoff
       │            │            │
       │            │            │
      L0          CLI          Session
   authority    generated     continuity
```

The session handoff is not evidence.

It is not part of the grounding corpus.

It cannot override governed context.

---

# 42. Automatic Session Generation

The CLI/service should periodically update session state.

Possible triggers:

```text
git activity
lifecycle event
explicit cdad resume
explicit cdad status
timer/background service
IDE integration
CI event
```

The initial v3 implementation should prefer deterministic event/command triggers.

A background daemon may be considered later.

---

# 43. Service vs CLI

The CLI itself is the first implementation.

Later CDAD may introduce:

```text
CDAD CLI
   +
CDAD Local Service
```

The service could provide:

- scheduled session state;
- lifecycle event processing;
- local project inspection;
- background validation;
- integrations.

However, the service should not be required for basic CDAD operation.

---

# 44. Recommended Evolution

```text
v3.0
  CLI bootstrap
  status
  validate
  inspect
  Claude Code integration

v3.1
  session state
  session handoff
  lifecycle events
  audit
  improved ADE abstraction

v3.2+
  local CDAD service
  richer integrations
  additional ADEs
  advanced automation
```

This is a proposal, not a committed release plan.

---

# 45. Main Use Cases

## UC-01 — Initialize project

```text
Actor: Developer

Given:
  existing software project

When:
  cdad init

Then:
  CDAD is initialized
  ADE is detected
  required artifacts are prepared
  bootstrap instructions are available
```

---

## UC-02 — Bootstrap Claude Code

```text
Actor: Developer

When:
  cdad init

And:
  Claude Code is detected

Then:
  Claude Code integration is prepared
  CDAD bootstrap instructions are provided
```

---

## UC-03 — Inspect project

```text
cdad inspect
```

Expected:

```text
Current CDAD state
Governance status
Open items
Proposals
Change Requests
Freeze state
Session state
```

---

## UC-04 — Validate project

```text
cdad validate
```

Expected:

```text
Deterministic governance validation
```

---

## UC-05 — CI governance

```text
CI
 │
 ▼
cdad validate --ci
 │
 ├── PASS
 └── FAIL
```

---

## UC-06 — Resume next day

```text
Developer
    │
    ▼
cdad resume
    │
    ▼
Session Handoff
    │
    ▼
Current CDAD state
    │
    ▼
Claude Code
```

---

## UC-07 — Freeze

```text
Human decision
      ↓
cdad freeze
      ↓
validation
      ↓
freeze
```

---

## UC-08 — Change after freeze

```text
Developer discovers change
          ↓
Change Request
          ↓
THINK
          ↓
Proposal
          ↓
Human Decision
          ↓
cdad freeze
```

The CLI does not provide `unfreeze`.

---

# 46. Epic Stories

## EPIC-01 — CDAD CLI Bootstrap

### Story 01.1

**As a developer, I want to install the CDAD CLI so that I can use CDAD without cloning the bootstrap repository.**

Acceptance criteria:

- CLI can be installed as a standalone executable.
- Windows/Linux/macOS are supported.
- `cdad version` works.

### Story 01.2

**As a developer, I want to run `cdad init` so that CDAD is initialized in my project.**

Acceptance criteria:

- Existing project detected.
- CDAD structure created.
- Version recorded.
- Existing files are not silently overwritten.

---

# 47. EPIC-02 — ADE Integration

### Story 02.1

**As a developer, I want CDAD to detect Claude Code so that the appropriate integration is prepared automatically.**

Acceptance criteria:

- Claude Code detection is deterministic.
- Integration is isolated in an adapter.
- CDAD core does not depend on Claude Code internals.

### Story 02.2

**As a developer, I want CDAD to generate the bootstrap instructions for Claude Code so that the agent can populate CDAD artifacts.**

Acceptance criteria:

- Canonical instruction is versioned.
- Instructions respect CDAD governance.
- No architectural decision is made by the CLI.

---

# 48. EPIC-03 — Deterministic Validation

### Story 03.1

**As a developer, I want `cdad validate` to identify governance inconsistencies.**

Acceptance criteria:

- Validation is deterministic.
- No LLM is required.
- Exit code indicates result.

### Story 03.2

**As a DevOps engineer, I want `cdad validate --ci` so that CDAD can operate as a pipeline gate.**

Acceptance criteria:

- Non-interactive.
- Machine-readable output.
- Stable exit codes.

---

# 49. EPIC-04 — Inspection

### Story 04.1

**As a developer, I want `cdad status` to show the current CDAD state.**

Acceptance criteria:

- Freeze status shown.
- OPEN/BLOCKING shown.
- Proposals shown.
- Change Requests shown.
- Session state shown.

### Story 04.2

**As a developer, I want `cdad inspect` to provide deeper diagnostics.**

Acceptance criteria:

- Structural checks shown.
- Warnings distinguished from failures.
- Remediation guidance provided where deterministic.

---

# 50. EPIC-05 — Session Continuity

### Story 05.1

**As a developer, I want CDAD to automatically maintain session state so that I can continue work later without manually maintaining a handoff file.**

Acceptance criteria:

- State derives from real project artifacts.
- Agent does not need to write the authoritative state.
- Handoff is not evidence.
- Handoff does not override governed context.

### Story 05.2

**As a developer, I want `cdad resume` so that I can quickly understand where the project was left.**

Acceptance criteria:

- Last known state shown.
- Active change shown.
- Pending proposals shown.
- OPEN items shown.
- Latest handoff shown.

---

# 51. EPIC-06 — Lifecycle Automation

### Story 06.1

**As a CDAD user, I want lifecycle events so that CDAD can automate actions at important project moments.**

Examples:

```text
init
think
proposal
change-request
freeze
work
complete
session-update
```

### Story 06.2

**As a project owner, I want configurable lifecycle hooks so that CDAD can execute deterministic automation.**

Acceptance criteria:

- Hooks are explicit.
- Hooks cannot bypass governance.
- Hook failures are observable.

---

# 52. EPIC-07 — Audit

### Story 07.1

**As a project owner, I want a CDAD audit report so that I can understand the project's governance history.**

Acceptance criteria:

- Freeze history included.
- ADR/change history included.
- Proposals included.
- Validation results included.
- Machine-readable output supported.

---

# 53. EPIC-08 — Versioning and Migration

### Story 08.1

**As a developer, I want the CLI to identify the CDAD version of my project.**

### Story 08.2

**As a developer, I want controlled migration between CDAD schema versions.**

Acceptance criteria:

- Migration is explicit.
- Existing artifacts are backed up or recoverable.
- Migration does not silently alter frozen decisions.

---

# 54. EPIC-09 — ADE Portability

### Story 09.1

**As a developer, I want CDAD to remain independent of Claude Code so that I can use another ADE later.**

Acceptance criteria:

- Core governance does not depend on Claude Code.
- ADE integration is adapter-based.
- CDAD artifacts remain portable.

Future adapters:

```text
Claude Code
Kiro
Copilot
Codex
other ADEs
```

---

# 55. EPIC-10 — Security

### Story 10.1

**As a developer, I want CDAD CLI to avoid exposing secrets so that project bootstrap and validation are safe.**

Acceptance criteria:

- Credentials are not logged.
- Source is not uploaded by default.
- Network operations are explicit.
- Downloaded artifacts can be integrity-checked.

---

# 56. Example End-to-End Experience

Developer opens an existing project in VS Code.

Terminal:

```bash
cdad init
```

CLI:

```text
CDAD CLI 3.0

Project detected:
  payments-api

ADE detected:
  Claude Code

CDAD mode:
  Existing project

Initializing CDAD...
✓ Bootstrap obtained
✓ CDAD structure created
✓ Claude Code integration prepared
✓ Governance artifacts created
✓ Bootstrap instruction prepared

Project ready for CDAD initialization.

Next:
  Start Claude Code and execute the CDAD bootstrap instruction.
```

Claude Code:

```text
Analyzing project under CDAD governance...
```

It populates the appropriate artifacts.

Developer then runs:

```bash
cdad validate
```

CLI:

```text
CDAD Validation

✓ Structure
✓ Provenance
✓ ADR references
✓ Proposals
✓ Freeze prerequisites

Result: VALID
```

---

# 57. Example Daily Workflow

```text
Morning
   │
   ▼
VS Code
   │
   ▼
cdad resume
   │
   ▼
Current CDAD state
   │
   ▼
Claude Code
   │
   ▼
Work
   │
   ▼
cdad validate
   │
   ▼
Continue / Change Request / Freeze
```

---

# 58. Example CI/CD Workflow

```text
Developer
    │
    ▼
Git Push
    │
    ▼
CI Pipeline
    │
    ▼
cdad validate --ci
    │
 ┌──┴──┐
 ▼     ▼
PASS   FAIL
 │      │
 ▼      ▼
Build  Stop
```

CDAD therefore becomes part of the development governance pipeline without becoming the CI/CD system itself.

---

# 59. Design Invariants

The following should be treated as architectural invariants for the CLI.

## I-CLI-01

**The CLI does not make architectural decisions.**

## I-CLI-02

**The CLI does not replace the human decision point.**

## I-CLI-03

**The CLI cannot bypass freeze governance.**

## I-CLI-04

**The CLI must not introduce an `unfreeze` operation.**

## I-CLI-05

**Session handoff is not evidence.**

## I-CLI-06

**Session state must be derived from project state wherever possible.**

## I-CLI-07

**ADE integrations must remain separated from CDAD core governance.**

## I-CLI-08

**Validation must be deterministic whenever the property being checked is structural.**

## I-CLI-09

**The CLI should work without an LLM for inspection and validation.**

## I-CLI-10

**CDAD artifacts remain portable across ADEs.**

---

# 60. Open Design Questions

These should be resolved before implementation is considered stable.

- How exactly is `cdad-bootstrap` distributed?
- GitHub Releases, package registry, embedded templates, or hybrid?
- Does `cdad init` modify existing ADE files or create dedicated CDAD integration files?
- How exactly should Claude Code be detected?
- What is the canonical Claude Code integration mechanism for CDAD?
- Should the CLI invoke Claude Code automatically or only prepare instructions?
- What is the canonical event schema?
- Should session state be stored in `.cdad-session/` or another structure?
- Which session artifacts are derived versus optional agent notes?
- How frequently should automatic session state be refreshed?
- Is a local background CDAD service needed in v3.1?
- What is the canonical bootstrap manifest?
- How are bootstrap versions pinned?
- What is the exact CDAD v3 schema?
- Which ADEs are officially supported in v3?
- Which commands belong in v3.0 versus v3.1?

---

# 61. Proposed v3 Minimum Viable CLI

The first implementation should be intentionally small.

```text
cdad init
cdad status
cdad inspect
cdad validate
cdad freeze
cdad resume
cdad version
```

Everything else can build on these primitives.

The most important capability is not the number of commands.

It is the transition:

```text
Clone bootstrap manually
        ↓
Configure ADE manually
        ↓
Copy CDAD files
        ↓
Understand setup
```

to:

```bash
cdad init
```

followed by:

```text
CDAD prepares
Claude Code reasons
Human decides
CDAD validates
```

---

# 62. Final Architectural Vision

```text
                         DEVELOPER
                             │
                             ▼
                    ┌────────────────┐
                    │   VS CODE      │
                    │                │
                    │  Integrated    │
                    │   Terminal     │
                    └───────┬────────┘
                            │
                            ▼
                       ┌─────────┐
                       │   cdad  │
                       │   CLI   │
                       └────┬────┘
                            │
          ┌─────────────────┼─────────────────┐
          │                 │                 │
          ▼                 ▼                 ▼
      Bootstrap         Lifecycle        Validation
      & ADE setup        & Events         & Inspect
          │                 │                 │
          └─────────────────┼─────────────────┘
                            │
                            ▼
                     CDAD Project State
                            │
                            ▼
                     ┌──────────────┐
                     │ Claude Code  │
                     └──────┬───────┘
                            │
                       THINK / WORK
                            │
                            ▼
                       Proposals
                            │
                            ▼
                         HUMAN
                            │
                            ▼
                         FREEZE
                            │
                            ▼
                    Governed Software
```

The central idea is:

> **CDAD CLI is the operational control surface of CDAD, not the intelligence layer.**

It makes CDAD easy to install, initialize, inspect, validate, resume and automate while keeping reasoning with the AI agents and authority with the human.

---

# 63. Relationship to CDAD v3 D11/D12

This CLI design directly supports the current v3 analysis:

- `cdad init` selects the CDAD entry mode.
- `cdad status` provides deterministic project state.
- `cdad validate` implements structural governance checks.
- `cdad freeze` executes the mechanical freeze boundary.
- `cdad resume` exposes session continuity without making the handoff authoritative.
- lifecycle events provide automation around CDAD rather than replacing CDAD governance.
- ADE adapters keep Claude Code integration separate from the CDAD core.
- the CLI never becomes a reasoning worker.

This preserves the central CDAD boundary:

```text
GROUNDING
    ↓
DOSSIER
    ↓
WORKERS / THINK
    ↓
PROPOSALS
    ↓
HUMAN
    ↓
FREEZE
```

while adding an operational layer around it:

```text
                 CDAD CLI
                    │
       ┌────────────┼────────────┐
       ▼            ▼            ▼
    Bootstrap    Lifecycle    Validation
       │            │            │
       └────────────┼────────────┘
                    ▼
              CDAD Governance
                    │
                    ▼
              AI Development
```

**Status:** Proposed technical design for discussion; not yet a ratified CDAD specification.
