A tool to interact with OpenAI

```bash
go run main.go -key {your OpenAI key} tell me 1 + 1 =
2


uptime | go run main.go -key {your OpenAI key}  convert this to json
{
    "up": "71 days, 3:04",
    "users": "2 users",
    "load_averages": ["1.90", "2.14", "2.14"]
}
```