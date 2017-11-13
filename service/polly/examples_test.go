// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package polly_test

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/polly"
)

var _ time.Duration
var _ strings.Reader
var _ aws.Config

func parseTime(layout, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}

// To delete a lexicon
//
// Deletes a specified pronunciation lexicon stored in an AWS Region.
func ExamplePolly_DeleteLexiconRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := polly.New(cfg)
	input := &polly.DeleteLexiconInput{
		Name: aws.String("example"),
	}

	req := svc.DeleteLexiconRequest(input)
	result, err := req.Send()
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case polly.ErrCodeLexiconNotFoundException:
				fmt.Println(polly.ErrCodeLexiconNotFoundException, aerr.Error())
			case polly.ErrCodeServiceFailureException:
				fmt.Println(polly.ErrCodeServiceFailureException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To describe available voices
//
// Returns the list of voices that are available for use when requesting speech synthesis.
// Displayed languages are those within the specified language code. If no language
// code is specified, voices for all available languages are displayed.
func ExamplePolly_DescribeVoicesRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := polly.New(cfg)
	input := &polly.DescribeVoicesInput{
		LanguageCode: polly.LanguageCodeEnGb,
	}

	req := svc.DescribeVoicesRequest(input)
	result, err := req.Send()
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case polly.ErrCodeInvalidNextTokenException:
				fmt.Println(polly.ErrCodeInvalidNextTokenException, aerr.Error())
			case polly.ErrCodeServiceFailureException:
				fmt.Println(polly.ErrCodeServiceFailureException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To retrieve a lexicon
//
// Returns the content of the specified pronunciation lexicon stored in an AWS Region.
func ExamplePolly_GetLexiconRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := polly.New(cfg)
	input := &polly.GetLexiconInput{
		Name: aws.String(""),
	}

	req := svc.GetLexiconRequest(input)
	result, err := req.Send()
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case polly.ErrCodeLexiconNotFoundException:
				fmt.Println(polly.ErrCodeLexiconNotFoundException, aerr.Error())
			case polly.ErrCodeServiceFailureException:
				fmt.Println(polly.ErrCodeServiceFailureException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To list all lexicons in a region
//
// Returns a list of pronunciation lexicons stored in an AWS Region.
func ExamplePolly_ListLexiconsRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := polly.New(cfg)
	input := &polly.ListLexiconsInput{}

	req := svc.ListLexiconsRequest(input)
	result, err := req.Send()
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case polly.ErrCodeInvalidNextTokenException:
				fmt.Println(polly.ErrCodeInvalidNextTokenException, aerr.Error())
			case polly.ErrCodeServiceFailureException:
				fmt.Println(polly.ErrCodeServiceFailureException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To save a lexicon
//
// Stores a pronunciation lexicon in an AWS Region.
func ExamplePolly_PutLexiconRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := polly.New(cfg)
	input := &polly.PutLexiconInput{
		Content: aws.String("file://example.pls"),
		Name:    aws.String("W3C"),
	}

	req := svc.PutLexiconRequest(input)
	result, err := req.Send()
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case polly.ErrCodeInvalidLexiconException:
				fmt.Println(polly.ErrCodeInvalidLexiconException, aerr.Error())
			case polly.ErrCodeUnsupportedPlsAlphabetException:
				fmt.Println(polly.ErrCodeUnsupportedPlsAlphabetException, aerr.Error())
			case polly.ErrCodeUnsupportedPlsLanguageException:
				fmt.Println(polly.ErrCodeUnsupportedPlsLanguageException, aerr.Error())
			case polly.ErrCodeLexiconSizeExceededException:
				fmt.Println(polly.ErrCodeLexiconSizeExceededException, aerr.Error())
			case polly.ErrCodeMaxLexemeLengthExceededException:
				fmt.Println(polly.ErrCodeMaxLexemeLengthExceededException, aerr.Error())
			case polly.ErrCodeMaxLexiconsNumberExceededException:
				fmt.Println(polly.ErrCodeMaxLexiconsNumberExceededException, aerr.Error())
			case polly.ErrCodeServiceFailureException:
				fmt.Println(polly.ErrCodeServiceFailureException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To synthesize speech
//
// Synthesizes plain text or SSML into a file of human-like speech.
func ExamplePolly_SynthesizeSpeechRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := polly.New(cfg)
	input := &polly.SynthesizeSpeechInput{
		LexiconNames: []*string{
			aws.String("example"),
		},
		OutputFormat: polly.OutputFormatMp3,
		SampleRate:   aws.String("8000"),
		Text:         aws.String("All Gaul is divided into three parts"),
		TextType:     polly.TextTypeText,
		VoiceId:      polly.VoiceIdJoanna,
	}

	req := svc.SynthesizeSpeechRequest(input)
	result, err := req.Send()
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case polly.ErrCodeTextLengthExceededException:
				fmt.Println(polly.ErrCodeTextLengthExceededException, aerr.Error())
			case polly.ErrCodeInvalidSampleRateException:
				fmt.Println(polly.ErrCodeInvalidSampleRateException, aerr.Error())
			case polly.ErrCodeInvalidSsmlException:
				fmt.Println(polly.ErrCodeInvalidSsmlException, aerr.Error())
			case polly.ErrCodeLexiconNotFoundException:
				fmt.Println(polly.ErrCodeLexiconNotFoundException, aerr.Error())
			case polly.ErrCodeServiceFailureException:
				fmt.Println(polly.ErrCodeServiceFailureException, aerr.Error())
			case polly.ErrCodeMarksNotSupportedForFormatException:
				fmt.Println(polly.ErrCodeMarksNotSupportedForFormatException, aerr.Error())
			case polly.ErrCodeSsmlMarksNotSupportedForTextTypeException:
				fmt.Println(polly.ErrCodeSsmlMarksNotSupportedForTextTypeException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}
