# Hello Forge

The smallest forge-core example: use a Forge provider, build an agent with `forge.Config`, ask one question, and print the latest assistant text.

## Run

```bash
export XAI_API_KEY=xai-...
export OPENAI_API_KEY=sk-...
go run .
```

## What this shows

- `forge.Config` as the agent setup point
- using `provider/openai` or `provider/xai`
- `agent.Ask(ctx, prompt)` for the common path
- `resp.LastText()` for the latest assistant answer
