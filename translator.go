package main

import (
	"context"
	"fmt"
	// "log"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

// https://github.com/GoogleCloudPlatform/golang-samples/blob/master/translate/translate_quickstart/main.go

// https://cloud.google.com/translate/docs/languages

func translateText(targetLanguage, text, model string) (string, error) {
	// targetLanguage := "ja"
	// text := "The Go Gopher is cute"
	// model := "nmt"

	ctx := context.Background()

	lang, err := language.Parse(targetLanguage)
	if err != nil {
			return "", fmt.Errorf("language.Parse: %v", err)
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
			return "", fmt.Errorf("translate.NewClient: %v", err)
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, lang, &translate.Options{
			Model: model, // Either "mnt" or "base".
	})
	if err != nil {
			return "", fmt.Errorf("Translate: %v", err)
	}
	if len(resp) == 0 {
			return "", nil
	}
	// fmt.Println(resp[0].Text)
	return resp[0].Text, nil
}

