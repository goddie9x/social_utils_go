package constants

type Role int

const (
	// Unknown is the zero value and must stay the most restrictive role.
	// A Role unmarshalled from a missing JSON field (or a zero-valued
	// UserAuth) must never escalate to Admin.
	Unknown Role = iota
	User
	Mod
	Admin
)

type PrivacyType int

const (
	Private PrivacyType = iota
	Public
	Friend
)

type PostType int

const (
	PersonalPost PostType = iota
	GroupPost
)
