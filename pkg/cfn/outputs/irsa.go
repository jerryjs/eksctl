package outputs

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
)

// ServiceAccount/Role
type IRSA struct {
	IAMRole        string
	ServiceAccount string
}

var (
	IRSAOutput chan IRSA
)

func init() {
	IRSAOutput = make(chan IRSA)
}

func Get(stack types.Stack, key string) *string {
	for _, x := range stack.Outputs {
		if *x.OutputKey == key {
			return x.OutputValue
		}
	}
	return nil
}
