package main

import (
	"fmt"

	"github.com/linielson/aws-sns-sqs/common"
)

func main() {
	sess := common.BuildSession()
	fmt.Printf("%v", sess)
}
