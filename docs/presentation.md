# Agents Under the Hype

45-minute predictive workshop for a general technical audience.

Core thesis: **agents are surprisingly simple underneath the hype: a model inside a loop, given tools and a stop rule.**

Timing target: 40 minutes of material + 5 minutes of Q&A.

---

## Slide 1 — Agents Under the Hype

**Time:** 1 minute

### On slide

> What even is an agent?

**A model inside a loop, with tools and a stop rule.**

### Visual

A large question mark dissolving into four plain boxes:

`Model → Loop → Tools → Stop`

### Speaker notes

Agent language often sounds mystical: autonomy, reasoning, planning, memory, orchestration. Today we are going to build a smaller and more useful mental model.

The claim is not that production agents are trivial. The claim is that their irreducible core is simple enough to understand in one session.

### Audience cue

Ask: “Before today, what made something feel like an agent rather than a chatbot?”

Take two quick answers without correcting them yet.

---

## Slide 2 — The vocabulary pile

**Time:** 2 minutes

### On slide

- LLM
- Assistant
- Agent
- Copilot
- Tool calling
- Workflow
- Multi-agent system

> Different products pile features onto the same small mechanism.

### Visual

A messy pile of labels above one small loop.

### Speaker notes

The market bundles many capabilities under “agent.” That makes the category feel more complicated than the mechanism.

Memory, planning, streaming, retrieval, permissions, and multiple agents can matter. None of them is required to explain the first working agent loop.

We will strip the system down until every moving part is visible.

---

## Slide 3 — A model is not yet an agent

**Time:** 2 minutes

### On slide

**Model**

- Receives context
- Produces the next turn
- May request a tool

**Harness**

- Owns the loop
- Executes side effects
- Feeds observations back
- Decides when to stop

### Visual

Two columns joined by arrows:

`Harness → Model → Harness`

### Speaker notes

The model proposes. The harness disposes.

The model cannot measure the bedroom temperature. It can only emit a structured request asking something else to do that work.

This distinction is the foundation for debugging. If the model chose the wrong tool, inspect instructions and tool descriptions. If the tool ran incorrectly, inspect ordinary application code. If the loop never stopped, inspect the harness.

---

## Slide 4 — The whole architecture

**Time:** 3 minutes

### On slide

> **Agent ≈ model + harness**

The harness contains:

- Context
- Tools
- Loop
- Budget
- Stop rule

### Visual

```text
User
  ↓
Harness ⇄ Model
  ↓
Tools
```

### Speaker notes

This is our working definition for the rest of the workshop.

The model provides judgment and language. The harness turns those capabilities into a system that can act.

The approximation sign matters. Real products add security, persistence, observability, evaluation, and user experience. Those are engineering layers around the same core.

### Audience cue

Ask: “Which box should own the API key and actually execute a tool?”

Expected answer: the harness or tool implementation, never the model.

---

## Slide 5 — A tool is a contract plus a handler

**Time:** 3 minutes

### On slide

**What the model sees**

- Name
- Description
- Parameters

**What the harness owns**

- Handler
- Validation
- Side effects
- Result

### Visual

```text
room_temperature
├── description: Get the temperature of a room
├── parameters: { room: string }
└── handler: Go function
```

### Speaker notes

The tool definition is an affordance for the model. Clear names and descriptions improve tool selection.

The handler is ordinary deterministic code. In this demo it always returns 72 degrees, which is perfect: we are teaching the loop, not building a weather station.

The model never receives the Go function. It receives metadata describing when and how to request it.

---

## Slide 6 — The loop needs a stop rule

**Time:** 3 minutes

### On slide

```text
Prompt model
     ↓
Tool calls?
  yes → execute → append observation → repeat
  no  → return assistant text
```

**Backstop:** iteration budget

### Visual

A circular flow ending in a clearly marked exit:

`No tool calls → Respond`

### Speaker notes

Many agent diagrams emphasize how the loop continues. The more important question is how it stops.

Our v1 rule is deliberately boring:

> If the assistant returns no tool calls, its text is the final answer.

There is no special `is_final` field and no Submit tool. An iteration budget kills runaway loops.

Simplicity is doing real work here. Every extra control mechanism creates another state the model and harness can disagree about.

### Audience cue

Ask: “If the model returns helpful text and zero tool calls, should we prompt it again?”

Expected answer: no. That text is the answer.

---

## Slide 7 — ReAct makes the loop visible

**Time:** 3 minutes

### On slide

1. **Thought** — brief visible status
2. **Action** — structured tool call
3. **Observation** — tool result
4. **Respond** — final answer

> ReAct is not magic reasoning. It is an observable interaction pattern.

### Visual

```text
Thought → Action → Observation
   ↑                    ↓
   └──── repeat? ───────┘
              ↓
           Respond
```

### Speaker notes

For this demo, “Thought” means a short user-visible status such as “I’ll check the bedroom.” It does not mean exposing private chain-of-thought.

Action is the model’s structured tool request. Observation is data returned by the harness. Respond is assistant text with no tool call.

This vocabulary gives us handles for prediction and debugging.

### Audience cue

Tell the room: “From here on, I will pause the system. You tell me which lifecycle step should happen next.”

---

## Slide 8 — Prediction round: no tool

**Time:** 2 minutes

### On slide

**Prompt**

> How is it going?

What should happen next?

### Visual

Initially show only the prompt. Reveal the answer after the audience commits.

### Speaker notes

The prompt does not require external information or an action.

Expected prediction: Respond directly. No Thought, Action, or Observation is needed.

### Reveal

```text
[Respond] It's going great! I'm here and ready to help you manage your home.
```

### Teaching point

An agent does not need to use a tool merely because tools are available.

---

## Slide 9 — Prediction round: use the tool

**Time:** 4 minutes

### On slide

**Prompt**

> Do you know the temprature in my bedroom?

What happens next?

### Visual

Reveal one line at a time.

### Audience cues and reveals

1. Ask: “Can the model answer truthfully from its own context?”
   - Reveal: `[Thought] I'll check the temperature in your bedroom for you.`
2. Ask: “What exact capability should it request?”
   - Reveal: `[Action] room_temperature map[room:bedroom]`
3. Ask: “Who produces the next line?”
   - Reveal: `[Observation] The temperature in bedroom is 72 degrees`
4. Ask: “Does the loop continue or stop?”
   - Reveal: `[Respond] The temperature in your bedroom is currently 72 degrees Fahrenheit.`

### Speaker notes

The misspelling is intentional: this is the exact prompt used in the proven demo. The model still maps the intent to the correct tool parameter.

The critical handoff is between Action and Observation. The model emits the Action; ordinary Go code executes it and supplies the Observation.

---

## Slide 10 — Live demo

**Time:** 4 minutes

### On slide

```bash
go run ./cmd/main/main.go
```

Watch for:

- Did it select the tool?
- Did the call contain the room?
- Did the observation return to the loop?
- Did it stop?

### Speaker notes

Run the exact bedroom prompt without touching code.

Do not narrate ahead of the terminal. Let the audience call each expected step, then let the program confirm or contradict them.

After the run, execute the no-tool prompt if time permits. The contrast proves tool use is conditional.

### Static fallback transcript

If the provider or network fails, reveal this verified run:

```text
2026/08/03 02:35:36 [Thought] I'll check the temperature in your bedroom for you.
2026/08/03 02:35:36 [Action] room_temperature map[room:bedroom]
2026/08/03 02:35:36 [Observation] The temperature in bedroom is 72 degrees
2026/08/03 02:35:36 [Respond] The temperature in your bedroom is currently **72 degrees Fahrenheit**.
```

### Recovery line

“The network failed, but the architecture did not become mysterious. Let’s dissect the last verified transcript.”

---

## Slide 11 — Dissect the transcript

**Time:** 3 minutes

### On slide

| Line | Produced by | Meaning |
|---|---|---|
| Thought | Model | User-visible status |
| Action | Model | Structured request |
| Observation | Tool handler | External fact |
| Respond | Model | No tool call, so stop |

### Visual

Place the transcript on the left and highlight the responsible component on a model/harness diagram on the right.

### Speaker notes

The transcript alternates ownership:

- Model speaks.
- Harness acts.
- Model sees the result.
- Harness recognizes the stop condition.

This alternating ownership is the agent. There is no hidden daemon deciding what to do next.

### Audience cue

Point to each transcript line and have the audience call “model” or “harness.”

---

## Slide 12 — Code dissection: the boundary

**Time:** 3 minutes

### On slide

```go
type Model interface {
    Prompt(
        ctx context.Context,
        input []Message,
        tools Tools,
    ) (AssistantMessage, error)
}
```

And the provider request receives both:

```go
response, err := o.client.Responses.New(ctx, responses.ResponseNewParams{
    Input: responses.ResponseNewParamsInputUnion{
        OfInputItemList: convertedInputs,
    },
    Tools: convertedTools,
    Model: o.model,
})
```

### Speaker notes

The interface is almost disappointingly small: messages and available tools go in; one assistant turn comes out.

The OpenAI-compatible provider is an adapter. It translates our provider-neutral messages and tool definitions into SDK request types.

The agent loop does not know about OpenAI SDK structures. The tool handler does not know about the provider.

### Teaching point

Boundaries make the loop understandable. They do not make it intelligent.

---

## Slide 13 — Code dissection: the entire trick

**Time:** 4 minutes

### On slide

```go
assistantMessage, err := a.Model.Prompt(ctx, messages, modelTools)
messages = append(messages, assistantMessage)

if len(assistantMessage.ToolCalls()) == 0 {
    break
}
```

Then:

```go
content, err := a.Registry.Execute(t.Name(), t.Args())
observation := model.NewToolMessage(content, t.ID())
messages = append(messages, observation)
```

### Speaker notes

This is the heart of the implementation:

1. Ask for the next turn.
2. Preserve the assistant turn in history.
3. Stop when there are no tool calls.
4. Otherwise execute the requested tool.
5. Tie the observation to the original call ID.
6. Repeat.

The call ID matters because the provider must know which Action an Observation answers.

### Audience cue

Ask: “Where is the agent’s autonomy in this code?”

Answer: in the model choosing the next turn inside boundaries enforced by the harness.

---

## Slide 14 — The scars are the lesson

**Time:** 3 minutes

### On slide

What broke while building:

- Tool description silently omitted
- Observation content and call ID reversed
- Empty registry advertised no capabilities
- Tests accepted the wrong error
- Logging concerns leaked into the loop

Deliberately excluded:

- Skills
- Memory
- Streaming
- Parallel tools
- Multi-agent orchestration

### Visual

Two columns: **Real bugs** and **Tempting complexity**.

### Speaker notes

Most failures were boundary failures, not intelligence failures.

The model needed a clear tool description. The provider needed the correct call ID. The test needed to assert the right error. These are ordinary software-engineering problems.

The excluded features are not fake or useless. They are layers we did not need to prove the core mechanism.

This is the practical payoff of demystification: when the system misbehaves, you can name the broken seam.

---

## Slide 15 — The takeaway

**Time:** 5 minutes including Q&A

### On slide

> **Agents are pretty simple under the hype.**

```text
Model
  + loop
  + tools
  + observations
  + stop rule
  = useful agent
```

Build the smallest loop that can prove the behavior.

### Speaker notes

An agent is not a new kind of computer program. It is a familiar program that delegates next-turn judgment to a model.

The model can choose. The harness constrains. Tools act. Observations ground. The stop rule ends.

Once that is clear, advanced agent features become additions you can evaluate rather than magic you must accept.

### Final audience question

“Which production layer would you add first—and what concrete requirement justifies it?”

---

# Demo card

Keep this section open separately during the presentation.

## Tool path

Prompt:

```text
Do you know the temprature in my bedroom?
```

Expected lifecycle:

```text
[Thought] I'll check the temperature in your bedroom for you.
[Action] room_temperature map[room:bedroom]
[Observation] The temperature in bedroom is 72 degrees
[Respond] The temperature in your bedroom is currently 72 degrees Fahrenheit.
```

## No-tool path

Prompt:

```text
How is it going?
```

Expected lifecycle:

```text
[Respond] It's going great! I'm here and ready to help you manage your home.
```

## Failure fallback

If the live call fails:

1. Name the failure plainly: provider/network, not agent architecture.
2. Show the verified static transcript from Slide 10.
3. Continue directly into the dissection.
4. Do not debug live.

---

# Tabletop rehearsal notes

## Timing pass

| Beat | Slides | Budget |
|---|---:|---:|
| Demystify | 1–2 | 3 min |
| Model + harness + tools | 3–5 | 8 min |
| Loop + ReAct | 6–7 | 6 min |
| Audience prediction | 8–9 | 6 min |
| Live demo | 10 | 4 min |
| Dissection | 11–13 | 10 min |
| Scars | 14 | 3 min |
| Takeaway + Q&A | 15 | 5 min |
| **Total** |  | **45 min** |

## Tabletop rehearsal result

- Slide 5 complete at 11:00 — the primitives are established.
- Slide 7 complete at 17:00 — the audience has the ReAct vocabulary.
- Slide 10 complete at 27:00 — the live proof has landed.
- Slide 13 complete at 37:00 — the code dissection is finished.
- Slide 15 complete at 45:00 — takeaway and Q&A included.

The highest overrun risk is audience discussion on Slides 4 and 9. Take two answers, summarize, and move on. The exact tool prompt has already passed twice consecutively; the no-tool prompt has also passed.

## Transition check

- Slide 2 → 3: “Let’s separate the thing that talks from the thing that acts.”
- Slide 5 → 6: “A tool gives us one action. A loop turns that action into behavior.”
- Slide 6 → 7: “Now let’s put names on what we can observe in each lap.”
- Slide 9 → 10: “You have predicted the whole run. Let’s see whether the machine agrees.”
- Slide 10 → 11: “The output is not decoration; it is an architectural trace.”
- Slide 13 → 14: “The loop is small, but small does not mean bug-free.”
- Slide 14 → 15: “Complex production layers remain—but now we know what they are layers around.”

## Rehearsal cuts if running long

1. Skip the second live no-tool run; Slide 8 already establishes it.
2. Shorten Slide 12 to the `Model` interface only.
3. Take one audience prediction per lifecycle rather than all four.
4. Never cut the stop rule, the live tool run, or the final takeaway.
