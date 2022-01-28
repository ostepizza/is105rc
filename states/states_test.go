// Package for testing the represention of the state of the riverworld
package main

import "testing"

import "github.com/ostepizza/is105rc/states"

// Test function implemented in line with the Golang's testing tool
func test(t *testing.T) {
    wanted := "\\_Farmer_Fox_/"
    state := states.BoatSlot
    if state != wanted {
         t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
    }
}
