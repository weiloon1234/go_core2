package rules

import "core2/validator"

func init() {
	// Register all core rules
	validator.RegisterRule("required", Required)
	validator.RegisterRule("email", Email)
	validator.RegisterRule("min", Min)
	validator.RegisterRule("max", Max)
	validator.RegisterRule("alphanumeric", Alphanumeric)
}
