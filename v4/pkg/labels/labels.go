package labels

const (
	ProviderLabel               = "hobbyfarm.io/provider"
	EnvironmentLabel            = "hobbyfarm.io/environment"
	ScheduledEventCompleteLabel = "hobbyfarm.io/scheduled-event-complete"
	UsernameLabel               = "hobbyfarm.io/username"
)

// auth-related

const (
	LdapPrincipalKey  = "auth.hobbyfarm.io/ldap-principal"
	LocalPrincipalKey = "auth.hobbyfarm.io/local-principal"
	LocalUsernameKey  = "auth.hobbyfarm.io/local-username"
)

// accesscode related

const (
	AccessCodeLabel                = "hobbyfarm.io/accesscode"
	OneTimeAccessCodeLabel         = "hobbyfarm.io/onetimeaccesscode"
	OneTimeAccessCodeSetLabel      = "hobbyfarm.io/onetimeaccesscodeset"
	OneTimeAccessCodeRedeemedLabel = "hobbyfarm.io/otac-redeemed"
)
