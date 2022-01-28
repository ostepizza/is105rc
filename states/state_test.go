// Package for testing the represention of the state of the riverworld
package state

import "testing"

// Test function implemented in line with the Golang's testing tool
func test(t *testing.T) {
    wanted := "\\_Farmer_"*"_/" || "\\_"*"_Farmer_/"
    state := boat
    if state != wanted {
         t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
    }
}
