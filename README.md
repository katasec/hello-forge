# Hello Forge

The smallest forge example: use Forge's xAI provider, build an agent with `forge.Config`, ask one question, and print the latest assistant text.

## Run

```bash
export XAI_API_KEY=xai-...
go run .
```

## What this shows

- `forge.Config` as the agent setup point
- using `provider/xai` against xAI's Responses API
- `agent.Ask(ctx, prompt)` for the common path
- `resp.LastText()` for the latest assistant answer
