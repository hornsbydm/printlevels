package main

var parseOIDCases = []struct {
	description string
	in          string
	isOID       bool
	oid         string
}{
	{
		"genuine oid",
		"1.3.6.1.2.1.1.1.0",
		true,
		"1.3.6.1.2.1.1.1.0",
	},
	{
		"manual value 3 digits",
		"M100",
		false,
		"100",
	},
	{
		"manual value 1 digit",
		"M1",
		false,
		"1"},
	{
		"manual no value",
		"M",
		false,
		"0"},
}
