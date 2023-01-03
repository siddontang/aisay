package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	gogpt "github.com/sashabaranov/go-gpt3"
)

var (
	key     = flag.String("key", "", "Your OpenAI API key")
	verbose = flag.Bool("verbose", false, "output full prompt")
	mode    = flag.String("mode", "text-davinci-003", "OpenAI model, e.g. text-davinci-003, code-davinci-002, etc.")
)

func panicErr(err error) {
	if err == nil {
		return
	}

	panic(err.Error())
}

func readFromPipe() string {
	stat, _ := os.Stdin.Stat()

	// No data in pipe
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		return ""
	}

	var buf []byte
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		buf = append(buf, scanner.Bytes()...)
	}

	panicErr(scanner.Err())
	return string(buf)
}

func buildRequest(prompt string) gogpt.CompletionRequest {
	req := gogpt.CompletionRequest{
		Model:            *mode,
		MaxTokens:        1024,
		Prompt:           prompt,
		Temperature:      0,
		Stop:             nil,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
		BestOf:           1,
	}
	return req
}

func main() {
	flag.Parse()

	c := gogpt.NewClient(*key)
	ctx := context.Background()

	prompt := readFromPipe()

	prompt = fmt.Sprintf("%s\n\n%s", prompt, strings.Join(flag.Args(), " "))

	prompt = strings.TrimSpace(prompt)
	if prompt == "" {
		return
	}

	if *verbose {
		println(prompt)
	}

	req := buildRequest(prompt)
	resp, err := c.CreateCompletion(ctx, req)
	panicErr(err)

	output := strings.TrimSpace(resp.Choices[0].Text)
	println(output)
}
