package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"os"
)

func MakeSession() *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("DEFAULT_AWS_REGION")), // 使用するAWSリージョンを指定してください
	})
	if err != nil {
		fmt.Println("セッションを作成できませんでした:", err)
		return nil
	}
	return sess
}
