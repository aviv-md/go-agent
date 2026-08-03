# Agents Under the Hype — Speaker Cheatsheet

Keep this beside the terminal. It is a rescue map, not a script.

## North star

> An agent is a model inside a loop, with tools and a stop rule.

If stuck, return to:

> The model chooses the next turn. The harness constrains, executes, observes, and stops.

## Time checkpoints

| By | Be at |
|---:|---|
| 03:00 | Vocabulary pile finished |
| 11:00 | Model + harness + tools finished |
| 17:00 | ReAct explained |
| 23:00 | Prediction rounds finished |
| 27:00 | Live demo finished |
| 37:00 | Code dissection finished |
| 40:00 | Scars finished |
| 45:00 | Takeaway + Q&A finished |

## Slide-by-slide speaker notes

Each entry combines the intended explanation, rescue line, interaction, and transition.

### 1 — Agents Under the Hype · 1:00

**Say:** Agent language sounds mystical: autonomy, reasoning, planning, memory, orchestration. Production systems can be complex, but their irreducible core is small enough to understand today.

**If stuck:** “The irreducible core is smaller than the vocabulary around it.”

**Ask:** “What made something feel like an agent rather than a chatbot?” Take two answers.

**Move:** “Let’s remove the labels and find the mechanism.”

### 2 — The vocabulary pile · 2:00

**Say:** The market bundles many capabilities under “agent.” Memory, streaming, planning, retrieval, and multiple agents are useful layers, but none is required for our first working loop.

**If stuck:** “Memory, planning, and orchestration are layers—not the core loop.”

**Move:** “First, separate the thing that talks from the thing that acts.”

### 3 — A model is not yet an agent · 2:00

**Say:** The model receives context and proposes the next turn. The harness owns execution, observations, budgets, and stopping. The model cannot measure a room; it can only request a capability.

**If stuck:** “The model proposes. The harness disposes.”

**Move:** “Put those responsibilities together and we get our working definition.”

### 4 — The whole architecture · 3:00

**Say:** Agent is approximately model plus harness. The model contributes judgment and language. The harness turns those into constrained behavior.

**If stuck:** “Model chooses; harness constrains.”

**Ask:** “Who should own API keys and actually execute side effects?” Land on harness/tool code.

**Move:** “So what exactly does the model receive when a tool is available?”

### 5 — Tool contract + handler · 3:00

**Say:** The model sees name, description, and parameters. The harness owns the deterministic handler, validation, and result. Our temperature tool always returns 72 because the loop—not meteorology—is today’s lesson.

**If stuck:** “The model sees a contract; ordinary Go code owns the handler.”

**Move:** “One action is useful. A loop turns that action into behavior.”

### 6 — The stop rule · 3:00

**Say:** Prompt the model. If it requests a tool, execute it, append the observation, and repeat. If it requests no tools, its text is the answer. The iteration budget is only a runaway backstop.

**If stuck:** “No tool calls means the assistant text is the answer.”

**Ask:** “Helpful text and zero tool calls—should we prompt again?” Land on no.

**Move:** “Now let’s give each observable part of a lap a name.”

### 7 — ReAct · 3:00

**Say:** Thought is a short user-visible status, not private chain-of-thought. Action is the structured tool call. Observation is the tool result. Respond is final text with no tool call.

**If stuck:** “ReAct gives names to the observable rhythm.”

**Ask:** Tell the audience they now own prediction: they must call the next lifecycle step.

**Move:** “We’ll start with a prompt that should not act.”

### 8 — Predict: no tool · 2:00

**Say:** “How is it going?” needs no external fact. The correct behavior is one Respond and an immediate stop.

**If stuck:** “Available tools do not make tool use mandatory.”

**Ask:** “Which lifecycle step comes next?” Land on Respond.

**Move:** “Now let’s ask something the model cannot know by itself.”

### 9 — Predict: tool path · 4:00

**Say:** Reveal one line at a time. The model announces that it will check, requests `room_temperature`, receives 72 degrees from Go code, then answers without another tool call.

**If stuck:** “Model speaks, harness acts, model sees the result, harness stops.”

**Ask in order:** Can it answer truthfully? Which tool? Who produces Observation? Continue or stop?

**Move:** “You just predicted the complete run. Let’s see whether the machine agrees.”

### 10 — Live demo · 4:00

**Say:** Watch four things: tool selection, room argument, observation, and stop. Do not narrate ahead of the terminal—let the audience call each expected line.

**If stuck:** “We are testing the mental model, not trying to surprise ourselves.”

**Do:** Run the exact bedroom prompt. Run no-tool only if comfortably on time.

**Move:** “That output is not decoration; it is an architectural trace.”

### 11 — Dissect the transcript · 3:00

**Say:** Ownership alternates. Model produces Thought and Action. Harness executes and supplies Observation. Model produces Respond. Harness recognizes zero tool calls and stops.

**If stuck:** “The transcript alternates ownership between model and harness.”

**Ask:** Point at each line and have the room call “model” or “harness.”

**Move:** “Now let’s map that trace to the smallest code boundaries.”

### 12 — Code: the boundary · 3:00

**Say:** Messages and available tools go into `Prompt`; one assistant turn comes out. The provider only adapts neutral domain types to SDK request types.

**If stuck:** “Messages and tools go in; one assistant turn comes out.”

**Point out:** Agent code does not know OpenAI SDK structures; tool handlers do not know the provider.

**Move:** “Here is the loop consuming that tiny boundary.”

### 13 — Code: the entire trick · 4:00

**Say:** Prompt, append assistant turn, stop if there are no tool calls, otherwise execute, bind the observation to the call ID, append, and repeat.

**If stuck:** “Prompt, append, stop or execute, append observation, repeat.”

**Ask:** “Where is autonomy?” Land on the model choosing the next turn inside harness constraints.

**Move:** “The mechanism is small. That does not make it bug-free.”

### 14 — The scars · 3:00

**Say:** The failures were ordinary boundary bugs: missing descriptions, reversed call ID/content, empty registries, weak error assertions, and presentation leaking into the loop. Skills, memory, streaming, and multi-agent systems were deliberately excluded.

**If stuck:** “Most bugs were boundary bugs, not intelligence bugs.”

**Move:** “Those advanced features still matter—now we know what they are layers around.”

### 15 — Takeaway + Q&A · 5:00

**Say:** An agent is a familiar program that delegates next-turn judgment to a model. Model chooses. Harness constrains. Tools act. Observations ground. Stop rule ends.

**If stuck:** “Agents are pretty simple underneath the hype.”

**Ask:** “Which production layer would you add first—and what concrete requirement justifies it?”

## Audience questions and expected landing

### What makes an agent different from a chatbot?

Accept a couple of answers. Land on:

> It can choose actions inside a harnessed loop.

### Who owns the API key and executes tools?

> The harness/tool implementation—not the model.

### Helpful text, zero tool calls: prompt again?

> No. That text is the final response.

### Who produces an Observation?

> The harness executes ordinary code; the model only requested the Action.

### Where is autonomy in the loop?

> The model chooses the next turn inside boundaries enforced by the harness.

## Demo card

### Command

```bash
go run ./cmd/main/main.go
```

### Exact tool prompt

```text
Do you know the temprature in my bedroom?
```

### Expected lifecycle

```text
[Thought] I'll check the temperature in your bedroom for you.
[Action] room_temperature map[room:bedroom]
[Observation] The temperature in bedroom is 72 degrees
[Respond] The temperature in your bedroom is currently 72 degrees Fahrenheit.
```

### No-tool contrast

```text
How is it going?
```

Expected: one `Respond`; no Thought, Action, or Observation.

## If the live demo fails

Say:

> “The provider or network failed; the architecture did not become mysterious. Let’s inspect the last verified transcript.”

Then:

1. Show the static transcript.
2. Continue to Slide 11.
3. Do not debug live.

## If you lose your place

Use one of these bridges:

- “Which component owns this line: model or harness?”
- “What should happen next?”
- “What makes the loop stop?”
- “This looks intelligent, but which ordinary boundary made it possible?”
- “Let’s map that back to the four-step lifecycle.”

## If running late

Cut in this order:

1. Skip the second no-tool live run.
2. Show only the `Model` interface on Slide 12.
3. Take one audience prediction instead of all four.
4. Compress the scar list to: description, call ID, wrong error.

Never cut:

- The stop rule
- The live tool run or fallback transcript
- The final thesis

## Rabbit holes to park

> “Good production question; it is deliberately outside today’s smallest loop.”

Park:

- Memory
- Skills
- Streaming
- Parallel tools
- Multi-agent orchestration
- Production logging architecture

## Closing

> “An agent is not a new kind of computer program. It is a familiar program that delegates next-turn judgment to a model.”

Final question:

> “Which production layer would you add first—and what concrete requirement justifies it?”
