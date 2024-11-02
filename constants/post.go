package constants

type Role int

const (
	Admin Role = iota
	Mod
	User
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
