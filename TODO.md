# Demo build board — Prompt → Loop → Tool → Respond (ReAct)

Weekend goal: working demo complete, then presentation. Hand-code for recall; notes become slide ore.

**Definition of done:** Prompt in → ReAct loop (Thought → Action → Observation on the transcript) → one tool can fire → final respond when the model returns no tool calls (or budget stop). Narrate cold.

**Stop rule:** no tool calls ⇒ that assistant text is the answer. No structured `is_final`. No Submit/Prompt control tools for v1.

**Codebase scan (2026-08-03):** Tasks 5 and 6 are live-smoke verified; Tasks 7 and 8 are implementation/test verified; next is the agent tool turn.

| Area | State |
|------|--------|
| `internal/infra/config` | Done — `Load()` / `AI_BASE_URL` / `AI_API_KEY` |
| `cmd/main` | Thin — config → OpenAI provider → agent → `Run` → answer |
| `docs/architecture.md` | Done — Agent ≈ model + harness; package split (`model` / `agent` / `skill` / `tools`) |
| `internal/model` | Tool-aware message types + OpenAI-compatible Responses provider; text/tool conversion tested |
| `internal/agent` | Done — ReAct loop, empty-tool stop, iteration budget, fake-model tests |
| `internal/tools` | Done — tool metadata/schema/handler, registry list/execute, room-temperature tool |
| `internal/skill` | Does not exist |
| OpenAI provider | Done for text — owns client, maps messages, and returned a live OpenRouter response through `agent.Run` |

Current implementation checkpoint: Tasks 5–8 done; next wire the registry into the agent loop.

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

- **Description:** `NewOpenAIProvider` owns the client; `[]Message` → Responses input mapping and `model.Model` assertion are implemented. Proven through a live OpenRouter response via `agent.Run`.
- **Priority:** P0
- **Done:** [x]

### Task 6 — Thin `main` through agent

- **Description:** `main` loads config → builds provider → `NewAgent` → `Run`; the inline SDK call is gone. Live one-shot smoke returned text through the complete no-tool path.
- **Priority:** P0
- **Done:** [x]

---

## Phase B — One tool + ReAct loop

### Task 7 — `Message` / `Response` carry tools

- **Description:** Types can represent assistant Thought + Action (tool call id/name/args) and Observation (tool result tied to call id). Provider maps to/from Responses API tool items.
- **Priority:** P0
- **Done:** [x]

### Task 8 — `tools` thin slice

- **Description:** Added `internal/tools`: provider-neutral tool metadata/schema/handler, registry list/execute with unknown-tool errors, and a deterministic room-temperature tool factory.
- **Priority:** P0
- **Done:** [x]

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

---

## Extras backlog (explicitly deferred)

- **Provider HTTP contract test:** use `httptest.Server` to verify the real SDK request shape (`POST /responses`, model, messages) and decode a synthetic assistant `output_text`. Valuable for learning and regression coverage, but not required before the live Task 6 smoke.

### From code review (2026-08-02)

- **`RoleTool` cannot reach the provider — blocks Task 7.** The loop appends a `RoleTool` message (`internal/agent/agent.go:60-63`), but `convertRoleToOpenAIRole` has no `RoleTool` case and returns an error (`internal/model/openai.go:20-31`). The first real tool call therefore dies on the *next* `Prompt` with `role "tool" cannot be converted to OpenAI`. Invisible today only because the provider never populates `ToolCalls`. Not optional — Task 7 has to close this seam, and the Responses API wants a tool result item tied to a call id, not an `EasyInputMessage` role.
- **`log.Fatal` in `NewAgent`.** `internal/agent/agent.go:78` calls `os.Exit` from library code on a nil model. Violates the project's own "return errors explicitly" rule. A nil model is a programmer error: panic or return an error, but don't kill the caller's process from inside `internal/agent`.
- **`Role = string` is a type alias, not a distinct type.** `internal/model/model.go:20` — the `=` means every string in the program is a valid `Role`, so there is zero compile-time safety on roles. Dropping the `=` costs one character. Check first whether the alias is still load-bearing for the SDK conversion.
- **`NewAgent` returns unexported `*agent`.** Exported constructor handing back an unexported type with exported fields (`internal/agent/agent.go:15,70`). Callers outside the package can hold the value but cannot name it in a signature.
- **`config.Load()` does not validate `APIKey`.** `internal/infra/config/environment.go:32-37` accepts an empty key, so a missing `AI_API_KEY` surfaces as a confusing 401 from OpenRouter instead of a clear startup failure. Trust boundary, cheap fix.
- **Iteration budget uses `==` instead of `>=`.** `internal/agent/agent.go:54` is correct today only because the counter increments by exactly one per lap. Brittle the moment anything else touches it.

Deliberately *not* on this list (reviewed and accepted as-is): `ToolCalls []any` as an honest stub that Task 7 replaces, `config.Env` as an unguarded package global under a single-threaded `main`, and the hardcoded model string in `cmd/main/main.go:18`.
