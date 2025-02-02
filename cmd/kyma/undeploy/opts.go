package undeploy

import (
	"github.com/kyma-project/cli/internal/cli"
	"time"
)

//Options defines available options for the command
type Options struct {
	*cli.Options
	KeepCRDs bool
	Timeout  time.Duration
}

//NewOptions creates options with default values
func NewOptions(o *cli.Options) *Options {
	return &Options{Options: o}
}
