package go_toolkit_config

import "github.com/gookit/validate"

func Config() {
	// change global opts
	validate.Config(func(opt *validate.GlobalOption) {
		opt.StopOnError = false
	})
}
