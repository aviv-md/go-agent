# Session memory index

One-line pointers for future sessions. Newest first.

| Date | Milestone |
|------|-----------|
| 2026-08-03 | Speaker materials packaged: indexed English/Hebrew Markdown and standalone HTML cheatsheets |
| 2026-08-03 | Tasks 13–15 done: demo frozen; 15-slide predictive workshop drafted and timing-checked |
| 2026-08-03 | Tasks 11–12 done: no-tool smoke passed; exact tool lifecycle passed twice untouched |
| 2026-08-03 | Task 10 logging cleanup deferred: transcript boundary, centralized styling, rendering tests |
| 2026-08-03 | Task 10 done: colored visible Thought → Action → Observation → Respond lifecycle |
| 2026-08-03 | All remaining Task 9 concerns moved to post-demo polish; none block the demo |
| 2026-08-03 | Task 9 accepted for demo: live tool turn passed; focused regression coverage deferred |
| 2026-08-03 | Task 9 live smoke passed end-to-end; focused transcript regression test remains |
| 2026-08-03 | Task 9 boundary: tool parameters travel with each prompt; metadata stays beside handlers |
| 2026-08-03 | Task 8 done: provider-neutral Tool + registry + deterministic room-temperature tool; next = Task 9 agent tool turn |
| 2026-08-02 | Task 7 done: Message interface hierarchy + constructors, provider fan-out both directions, `RoleTool` scar retired; next = Task 8 tools slice |
| 2026-08-02 | Review fixes landed (panic on nil model, `>=` budget, config validation); Task 7 deferred fresh; `CLAUDE.md` now imports the pairing contract |
| 2026-08-02 | Tasks 5 + 6 done: thin main returned a live OpenRouter response through provider → agent |
| 2026-08-02 | Provider implemented + conversion-tested; live Task 6 smoke next, `httptest.Server` contract test deferred to extras |
| 2026-08-01 | TODO reconciled: Task 4 done, current `model.Model` signature/name captured; next = Task 5 |
| 2026-08-01 | Session wrap: Task 4 done; next convo → Task 5 OpenAI provider |
| 2026-08-01 | Task 4 done (loop + fake tests); next = Task 5 provider |
| 2026-08-01 | Task 4 loop shipped; next = fake Model tests (`*_test.go`) |
| 2026-08-01 | Task 4: Agent as `*agent`, NewAgent(model, maxIters); next = fake + Run loop |
| 2026-08-01 | Task 4 design: fake Model, Run(ctx), stop=empty ToolCalls, Thought≠reasoning knob |
| 2026-08-01 | Session wrap: Task 3 committed; next convo → Task 4 agent loop |
| 2026-08-01 | Task 3 types: flat `Response`, `Model.Prompt`, agent maps to history `Message` |
| 2026-08-01 | Rename `internal/ai` → `internal/model`; living docs updated |
| 2026-08-01 | Reset to full round-trip; `TODO.md` rescanned (model/agent empty stubs; ReAct still the plan) |
| 2026-07-29 | First LLM round-trip works (OpenRouter + Responses via `cmd/main`) |
| 2026-07-28 | Package split documented (then `ai`; now `model`) / agent / skill / tools |
| 2026-07-28 | Pair-programming contract locked; Go-only, coach-not-code, session memory |
