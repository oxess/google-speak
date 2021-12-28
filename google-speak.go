// Command quickstart generates an audio file with the content "Hello, World!".
package main

import (
	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"context"
	"flag"
	"fmt"
	"google.golang.org/api/option"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
	"io/ioutil"
	"log"
)

func createClient(credentialsFilePath string) (*texttospeech.Client, context.Context) {
	// Instantiates a client.
	ctx := context.Background()

	var err error
	var client *texttospeech.Client

	if credentialsFilePath == "cloud-proxy" {
		client, err = texttospeech.NewClient(ctx)
	} else {
		client, err = texttospeech.NewClient(ctx, option.WithCredentialsFile(credentialsFilePath))
	}

	if err != nil {
		log.Fatal(err)
	}

	return client, ctx
}

func textToFile(client *texttospeech.Client, ctx context.Context, languageCode string, outputFilePath string, text string) {
	req := texttospeechpb.SynthesizeSpeechRequest{
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: text},
		},
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: languageCode,
			SsmlGender:   texttospeechpb.SsmlVoiceGender_FEMALE,
		},
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		},
	}

	resp, err := client.SynthesizeSpeech(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}

	// The resp's AudioContent is binary.
	err = ioutil.WriteFile(outputFilePath, resp.AudioContent, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Audio content written to file: %v\n", outputFilePath)
}

func main() {
	var languageCode string

	credentialsFilePath := flag.String("credentials-path", "cloud-proxy", "Path to credential file (JSON)")
	outputFilePath := flag.String("output-path", "output.mp3", "Output path (mp3 ext required)")
	languageCodePl := flag.Bool("pl", false, "Change context to polish language")
	flag.Parse()

	client, ctx := createClient(*credentialsFilePath)
	defer client.Close()

	if *languageCodePl {
		languageCode = "pl-PL"
	} else {
		languageCode = "en-US"
	}

	textToFile(client, ctx, languageCode, *outputFilePath, flag.Arg(0))
}
