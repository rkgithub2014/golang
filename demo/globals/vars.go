package globals

// flags, args
const (
	DbtypeFlag        = "dbtype"
	DbtypeFlagUsage   = "-dbtype: (Required), values: moodysdb|olapdb|tenantdb"
	FttypeFlag        = "fttype"
	ValidateFlag      = "validate"
	ValidateFlagUsage = "-validate: (Optional), values:true or false, validates command without commit"
	PkgdirArg         = "pkgdir"
	ConfigfileArg     = "config"
	DbCreateFlagSet   = "dbcreatefset"
	EmptyString       = ""
)

// CmdType CreditLens Command struct
type CmdType struct {
	Cmd        string
	SubCommand string
	CmdArgs    map[string]string
}

// CmdHandler CreditLens Command Handler
type CmdHandler interface {
	InitFlags(args []string)
	Validate() error
	Execute() error
}
