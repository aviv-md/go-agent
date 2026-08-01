# Demo build board — Prompt → Loop → Tool → Respond (ReAct)

Weekend goal: working demo complete, then presentation. Hand-code for recall; notes become slide ore.

**Definition of done:** Prompt in → ReAct loop (Thought → Action → Observation on the transcript) → one tool can fire → final respond when the model returns no tool calls (or budget stop). Narrate cold.

**Stop rule:** no tool calls ⇒ that assistant text is the answer. No structured `is_final`. No Submit/Prompt control tools for v1.

**Codebase scan (2026-08-01, updated):** foundations + `ai` types exist (`IsFinal` still on `Response` — to drop/replace). Agent loop still branches on `IsFinal`. OpenAI provider is a stub. `cmd/main` still one-shots Responses inline. No `internal/tools` / `internal/skill` packages yet.

---

## Phase 0 — Foundations

### Task 1 — Config + first LLM round-trip

- **Description:** `internal/infra/config` loads `AI_BASE_URL` / `AI_API_KEY`; `cmd/main` can call OpenRouter Responses and print a reply (“Anybody home?” path).
- **Priority:** P0
- **Done:** [x]

### Task 2 — Architecture mental model

- **Description:** Document Agent ≈ model + harness; package split `ai` / `agent` / `skill` / `tools` in `docs/architecture.md`.
- **Priority:** P1
- **Done:** [x]

### Task 3 — `ai` types + Provider interface

- **Description:** `Provider.Prompt` + `Message` / `Response` / roles exist. **Revisit:** grow shape for tool calls + tool results; drop `IsFinal` in favor of “has tool calls?”.
- **Priority:** P0
- **Done:** [~] (skeleton; not ReAct-ready)

### Task 4 — Agent loop (ReAct harness)

- **Description:** `agent.Run`: seed user message → loop `Prompt` → if tool calls: append assistant (text + tool call) + tool result(s) → continue; if no tool calls: final respond; budget backstop. Done only when you own it and can teach Thought / Action / Observation.
- **Priority:** P0
- **Done:** [ ]

---

## Phase A — Provider + thin main

### Task 5 — Finish OpenAI provider (text path)

- **Description:** Wire client into `NewOpenAIProvider`; map `[]Message` → Responses input; `Prompt` returns a real text `Response` (compile + run through `ai.Provider`, not inline in `main`).
- **Priority:** P0
- **Done:** [ ]

### Task 6 — Thin `main` through agent

- **Description:** `main` loads config → builds provider → `NewAgent` → `Run`; delete/bypass the inline one-shot Responses call. Early smoke: one-shot ask → answer (no tools) still proves the pipe.
- **Priority:** P0
- **Done:** [ ]

---

## Phase B — One tool + ReAct loop

### Task 7 — `Message` / `Response` carry tools

- **Description:** Types can represent assistant Thought + Action (tool call id/name/args) and Observation (tool result tied to call id). Provider maps to/from Responses API tool items. Remove `IsFinal`.
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
