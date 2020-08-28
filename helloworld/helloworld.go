package helloworld

import "github.com/aws/aws-sdk-go/service/qldbsession/qldbsessioniface"

//MyDriver to handle own driver
type MyDriver struct {
	ledgerName                string
	qldbSession               qldbsessioniface.QLDBSessionAPI
	retryLimit                uint8
	iseRetryLimit             uint16
	maxConcurrentTransactions uint16
}

//New returns a new MyDriver
func New() *MyDriver {
	return &MyDriver{"", nil, 0, 0, 0}
}
