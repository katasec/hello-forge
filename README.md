# Hello Kiln

The smallest [kiln](https://github.com/katasec/kiln) example: build an agent for
each provider, ask a question, and call a tool.

## Run

```bash
export OPENAI_API_KEY=sk-...
export XAI_API_KEY=xai-...
export ANTHROPIC_API_KEY=sk-ant-...
go run .
```

All three keys are required — the example exercises every provider in one run.

## What this shows

- `kiln.Config` as the agent setup point, one agent per provider
- `provider/openai`, `provider/xai` and `provider/anthropic`
- `xai.WithWebSearch()` plus `provider.LastCitations()` for sourced answers
- `tool.Func[In, Out]` registering a typed Go function the model can call
- `agent.Ask(ctx, prompt)` for the common path, and `resp.LastText()` for the answer
- `resp.Usage` for per-call token counts

The `add` tool prints when it runs, so a real tool call is visible rather than
inferred from the answer:

```
[Anthropic Tools]
  -> tool add(21, 21) invoked
21 + 21 = 42
```
