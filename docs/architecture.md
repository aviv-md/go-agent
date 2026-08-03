# Architecture

## Mental model

**Agent ≈ model + harness**

| Piece | Role |
|--------|------|
| **Model** | Judgment / planning / language. Next turn: text and/or tool calls. |
| **Harness** | Everything that turns the model into a working system: loop, tools, skills, context, stop rules. |

### Harness pieces

| Piece | Meaning in this project |
|--------|-------------------------|
| **Loop** | Drive turns: assemble context → call model → maybe run tools → feed results → repeat → stop. |
| **Tools** | Callable actions advertised with parameter definitions (represented as JSON Schema) and executed by Go handlers. |
| **Skills** | Agent Skills: folders with `SKILL.md` (and optional assets). Playbooks the harness can load into context. Skills teach *how*; they do not execute side effects themselves. |
| **Also harness** | System prompt / persona, context assembly, budgets and stop conditions. |

---

## Package separation

These are target ownership boundaries, not a claim that every package already exists. `TODO.md` is the source of truth for implementation status.

| Package | Owns | Does not own |
|---------|------|----------------|
| `internal/model` | Model I/O (messages and available-tool definitions in → assistant turn out) | Loop, skills, tool implementations |
| `internal/agent` | Loop, stop rules, what goes into each turn; wires the harness | Tool implementations, skill file parsing |
| `internal/skill` | Discover / load / select `SKILL.md` (and related assets) | Calling the model, running tools |
| `internal/tools` | Tool metadata and parameters; register + execute handlers | Prompt assembly, skill matching |

### Target dependency direction

```
agent → model, skill, tools
skill ↛ tools   (skills may name tools in markdown; Go code should not import tools)
tools ↛ skill
model ↛ agent / skill / tools
```

If `skill` starts importing `tools`, playbook text and runtime registry are probably mixed.

### Planned thin slices (per package)

- **`tools`**: one tool, invoke it, get a result.
- **`skill`**: later, after the demo scope freeze, load one `SKILL.md` into agent context.
- **`agent`**: drive model turns now; add one real tool turn next. Skills remain deferred.
- **`model`**: OpenAI-compatible Responses API boundary; text is implemented first, tool-call mapping follows.

---

## Tool-definition boundary

Tool names, descriptions, and parameters live beside their handlers in `internal/tools`. The agent selects available tools from the registry for each prompt; `internal/model` exposes the provider-neutral prompt boundary and converts those definitions for the provider.
