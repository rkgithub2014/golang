package globals

import "errors"

var (
	//ErrArgsNotFound - Arugument not found error
	ErrArgsNotFound = errors.New("Arguments not found")
	//ErrSubCommandMissing - Subcommand not found error
	ErrSubCommandMissing = errors.New("Subcommand missing")
	//ErrdbtypeFlagMissing - dbtype flag missing
	ErrdbtypeFlagMissing = errors.New("dbtype flag missing")
)
