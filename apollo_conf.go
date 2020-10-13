package autoconf

import (
	"github.com/eden-framework/apollo"
	"os"
)

func FromApollo(apolloConfig *apollo.ApolloBaseConfig, conf []interface{}) {
	if apolloConfig == nil || os.Getenv("GOENV") == "LOCAL" {
		return
	}

	apollo.AssignConfWithDefault(*apolloConfig, conf...)
}
