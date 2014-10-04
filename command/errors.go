package command

import (
	"errors"
)

var NoSpecPathError = errors.New("Please provide a spec")
var NoTargetError = errors.New("Please provide a target")
