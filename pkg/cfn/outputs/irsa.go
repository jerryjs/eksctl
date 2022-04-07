package outputs

import (
	cfn "github.com/aws/aws-sdk-go/service/cloudformation"
)

// ServiceAccount/Role
type SARole struct {
	ServiceAccount string
	Role           string
}

var (
	SARoleOutput chan SARole
)

func init() {
	SARoleOutput = make(chan SARole)
}

func Get(stack cfn.Stack, key string) *string {
	for _, x := range stack.Outputs {
		if *x.OutputKey == key {
			return x.OutputValue
		}
	}
	return nil
}
