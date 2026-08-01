# Demo build board — Prompt → Loop → Tool → Respond (ReAct)

Weekend goal: working demo complete, then presentation. Hand-code for recall; notes become slide ore.

**Definition of done:** Prompt in → ReAct loop (Thought → Action → Observation on the transcript) → one tool can fire → final respond when the model returns no tool calls (or budget stop). Narrate cold.

**Stop rule:** no tool calls ⇒ that assistant text is the answer. No structured `is_final`. No Submit/Prompt control tools for v1.

**Codebase scan (2026-08-01):** Task 4 complete; Task 5 is next.

| Area | State |
|------|--------|
| `internal/infra/config` | Done — `Load()` / `AI_BASE_URL` / `AI_API_KEY` |
| `cmd/main` | Done — inline OpenRouter Responses one-shot (“Anybody home?”) |
| `docs/architecture.md` | Done — Agent ≈ model + harness; package split (`model` / `agent` / `skill` / `tools`) |
| `internal/model` | Types in: `Model.Prompt`, `Message`, flat `Response`, roles; `ToolCalls` still `[]any`; no OpenAI impl yet |
| `internal/agent` | Done — ReAct loop, empty-tool stop, iteration budget, fake-model tests |
| `internal/tools` / `internal/skill` | Do not exist |
| OpenAI provider | Does not exist — client lives in `main` |

Current implementation checkpoint: `cd948d8` (agent loop tested with a fake model).

---

## Phase 0 — Foundations

### Task 1 — Config + first LLM round-trip

- **Description:** `internal/infra/config` loads `AI_BASE_URL` / `AI_API_KEY`; `cmd/main` can call OpenRouter Responses and print a reply (“Anybody home?” path).
- **Priority:** P0
- **Done:** [x]

### Task 2 — Architecture mental model

- **Description:** Document Agent ≈ model + harness; package split `model` / `agent` / `skill` / `tools` in `docs/architecture.md`.
- **Priority:** P1
- **Done:** [x]

### Task 3 — `model` types + `Model` interface

- **Description:** In `internal/model`: `Model.Prompt(context.Context, []Message) (Response, error)` with flat `Response` (`Content` + `ToolCalls` stub), roles (user/assistant/system/tool). Agent maps response → history `Message`. Plan for ReAct: “has tool calls?” not `IsFinal`.
- **Priority:** P0
- **Done:** [x]

### Task 4 — Agent loop (ReAct harness)

- **Description:** `agent.Run`: seed user message → loop `Prompt` → if tool calls: append assistant text + stub tool result(s) → continue; if no tool calls: final respond; budget backstop. Proven with fake-model tests; real tool execution remains Task 9.
- **Priority:** P0
- **Done:** [x]

---

## Phase A — Provider + thin main

### Task 5 — Finish OpenAI provider (text path)

- **Description:** `NewOpenAIProvider` owns the client; map `[]Message` → Responses input; `Prompt` returns a real text `Response`. Prove it satisfies `model.Model`, not through an inline call in `main`.
- **Priority:** P0
- **Done:** [ ]

### Task 6 — Thin `main` through agent

- **Description:** `main` loads config → builds provider → `NewAgent` → `Run`; delete/bypass the inline one-shot Responses call. Early smoke: one-shot ask → answer (no tools) still proves the pipe.
- **Priority:** P0
- **Done:** [ ]

---

## Phase B — One tool + ReAct loop

### Task 7 — `Message` / `Response` carry tools

- **Description:** Types can represent assistant Thought + Action (tool call id/name/args) and Observation (tool result tied to call id). Provider maps to/from Responses API tool items.
- **Priority:** P0
- **Done:** [ ]

### Task 8 — `tools` thin slice

- **Description:** Add `internal/tools`: register + execute one dumb tool (keep it boring; domain still open).
- **Priority:** P0
- **Done:** [ ]

### Task 9 — Agent tool turn (verbose transcript)

- **Description:** On tool calls: append assistant blabber + tool call, run tool, append result, `Prompt` again. On no tool calls: stop with that text as the answer. Prefer **one Action per Thought** for the teaching demo (parallel tools = later).
- **Priority:** P0
- **Done:** [ ]

### Task 10 — ReAct system nudge

- **Description:** System/developer prompt: light Thought → Action → Observation guidance without over-constraining. Decide if mid-loop Thought is shown in the demo or log-only.
- **Priority:** P1
- **Done:** [ ]

### Task 11 — Smoke: tool vs no-tool

- **Description:** Prompt that must use the tool (multi-lap ReAct); prompt that must not (single respond); budget still kills runaway. Checkpoint: Prompt → Loop → Tool → Respond is real.
- **Priority:** P0
- **Done:** [ ]

---

## Phase C — Demo harden + freeze

### Task 12 — Happy-path script

- **Description:** Exact prompt(s) for Monday; run twice without touching code.
- **Priority:** P1
- **Done:** [ ]

### Task 13 — Failure / build notes

- **Description:** Capture scars while coding (bad tool shape, missing assistant tool-call on transcript, stop bugs, parallel-tool surprises) — presentation ore.
- **Priority:** P1
- **Done:** [ ]

### Task 14 — Scope freeze

- **Description:** No skills, no second tool, no streaming, no Submit/end_turn control tool, no platform pitch in the binary.
- **Priority:** P0
- **Done:** [ ]

---

## Phase D — Presentation (only after freeze)

### Task 15 — Synthesize notes → deck beats

- **Description:** Reverse-engineer slides from demystify frame + working ReAct demo (“What even are agents?”).
- **Priority:** P1
- **Done:** [ ]

### Task 16 — Rehearse once

- **Description:** Run happy path cold; narrate Thought → Action → Observation → Respond without debugging live.
- **Priority:** P1
- **Done:** [ ]

---

## Out of scope (until freeze)

- Skills package
- Second tool / parallel tool calls
- Structured `is_final` / Submit / Prompt control tools
- Streaming / pubsub
- Cyvore agent platform pitch
- Article (future, if talk lands)
