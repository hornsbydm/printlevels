package main

import "testing"

func TestParseoid(t *testing.T) {
	for _, tc := range parseOIDCases {
		if isOid, OID := parseOID(tc.in); isOid != tc.isOID || OID != tc.oid {
			t.Fatalf("FAIL: %s want %s got %s", tc.description, tc.oid, OID)
		}
		t.Log("PASS:", tc.description)
	}
	t.Log("Tested", len(parseOIDCases), "cases.")
}
