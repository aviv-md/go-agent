# Session memory index

One-line pointers for future sessions. Newest first.

| Date | Milestone |
|------|-----------|
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
