package command

import (
	"errors"
)

var NoSpecPathError = errors.New("Please provide a spec")
var NoTargetError = errors.New("Please provide a target")
var NoServiceNameError = errors.New("Please provide a service name")
