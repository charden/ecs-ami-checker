package aws

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

const TagName string = "LatestAmiId"

func GetFunctionArn(ctx context.Context) string {
	lc, _ := lambdacontext.FromContext(ctx)
	fmt.Println(lc.InvokedFunctionArn)
	return lc.InvokedFunctionArn
}

func GetAmiIdFromTag(functionArn string, session *session.Session) string {
	svc := lambda.New(session)
	result, err := svc.ListTags(&lambda.ListTagsInput{
		Resource: aws.String(functionArn),
	})

	if err != nil {
		fmt.Println("タグを取得できませんでした:", err)
		return ""
	}

	latestAmiId := ""
	for key, value := range result.Tags {
		if key == TagName {
			latestAmiId = *value
			break
		}
	}
	return latestAmiId
}

func UpdateTag(amiId string, session *session.Session, ctx context.Context) {
	svc := lambda.New(session)
	lc, _ := lambdacontext.FromContext(ctx)

	newTags := map[string]*string{
		TagName: aws.String(amiId),
	}

	functionArn := lc.InvokedFunctionArn
	_, err := svc.TagResource(&lambda.TagResourceInput{
		Resource: aws.String(functionArn),
		Tags:     newTags,
	})
	if err != nil {
		fmt.Println("タグを更新できませんでした:", err)
		return
	}
}
