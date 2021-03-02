package types

type AToken struct {
	GUID    string
	Token   string
	Created int64
	Expires int64
}

type RToken struct {
	GUID          string
	Token         string
	Created       int64
	AccessCreated int64
	Expires       int64
	AccessExpires int64
}

