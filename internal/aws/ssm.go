package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"os"
)

func GetParameter(session *session.Session) string {

	ssmClient := ssm.New(session)

	result, err := ssmClient.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(os.Getenv("AMI_PARAMETER_PATH")),
		WithDecryption: aws.Bool(false),
	})
	if err != nil {
		fmt.Println("パラメータを取得できませんでした:", err)
		return ""
	}
	fmt.Println(*result.Parameter.Value)
	return *result.Parameter.Value
}
