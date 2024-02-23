package config

import (
	"github.com/m-mizutani/swarm/pkg/infra/policy"
	"github.com/urfave/cli/v2"
)

type Policy struct {
	dir cli.StringSlice
}

func (x *Policy) Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringSliceFlag{
			Name:        "policy-dir",
			Aliases:     []string{"d"},
			Usage:       "Directory path of policy files",
			EnvVars:     []string{"SWARM_POLICY_DIR"},
			Destination: &x.dir,
			Required:    true,
		},
	}
}

func (x *Policy) Configure() (*policy.Client, error) {
	var options []policy.Option
	for _, dir := range x.dir.Value() {
		options = append(options, policy.WithDir(dir))
	}

	return policy.New(options...)
}
