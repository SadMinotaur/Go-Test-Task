package types

type aToken struct {
	GUID string
	Token string
	Created int64
	Expires int64
}

type rToken struct {
	GUID          string
	Token         string
	Created       int64
	AccessCreated int64
	Expires       int64
	AccessExpires int64
}

