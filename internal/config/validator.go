package config

import (
	"errors"
	"fmt"
)

func validate(cfg *Config) error {
	var errs []error

	// TODO: add new configuration validators here
	if cfg.Dummy {
		errs = append(errs, fmt.Errorf("dummy not implemented yet"))
	}

	return errors.Join(errs...)
}
