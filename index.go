package main

import (
	"context"
	"ecs-ami-checker/internal/aws"
	"ecs-ami-checker/internal/slack"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context) {
	sess := aws.MakeSession()
	latestAmiId := aws.GetParameter(sess)
	functionArn := aws.GetFunctionArn(ctx)
	amiId := aws.GetAmiIdFromTag(functionArn, sess)
	fmt.Println(amiId)
	if amiId != "" && amiId != latestAmiId {
		fmt.Println("change")
		aws.UpdateTag(latestAmiId, sess, ctx)
		slack.NotifySlack(latestAmiId)
	}
	fmt.Println(latestAmiId)
}

func main() {
	lambda.Start(handler)
}
