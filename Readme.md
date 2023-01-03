A tool to interact with OpenAI

## Build

```bash
go build .
```

### Examples 

```
export AISAY_KEY={your OpenAI key}
```

```bash
aisay tell me 1 + 1 =
2


uptime | aisay convert this to json
{
    "up": "71 days, 3:04",
    "users": "2 users",
    "load_averages": ["1.90", "2.14", "2.14"]
}
```