package main

// ----------------------------- types imported from other libraries -----------------------------

// GroupToml holds the data of the group.toml file.
type GroupToml struct {
	Servers []*ServerToml `toml:"servers"`
}

// ServerToml is one entry in the group.toml file describing one server to use for
// the cothority.
type ServerToml struct {
	Address     Address
	Suite       string
	Public      string
	Description string
	URL         string `toml:"URL,omitempty"`
}

type Address string
