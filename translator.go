package main

import (
	"context"
	"fmt"
	// "log"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

// func translateText(text string) {
// 	ctx := context.Background()

// 	// Creates a client.
// 	client, err := translate.NewClient(ctx)
// 	if err != nil {
// 		log.Fatalf("Failed to create client: %v", err)
// 	}

// 	// Sets the target language.
// 	target, err := language.Parse("es")
// 	if err != nil {
// 		log.Fatalf("Failed to parse target language: %v", err)
// 	}

// 	// Translates the text.
// 	translations, err := client.Translate(ctx, []string{text}, target, nil)
// 	if err != nil {
// 		log.Fatalf("Failed to translate text: %v", err)
// 	}

// 	// fmt.Printf("Text: %v\n", text)
// 	fmt.Printf("Translation: %v\n", translations[0].Text)
// 	return translations[0].Text
// }

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

