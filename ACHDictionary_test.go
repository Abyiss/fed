// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package feddir

import (
	"os"
	"strings"
	"testing"
)

func TestParseParticipant(t *testing.T) {
	var line = "073905527O0710003011012908000000000LINCOLN SAVINGS BANK                P O BOX E                           REINBECK            IA506690159319788644111     "

	f := NewACHDictionary(strings.NewReader(line))
	f.Read()

	// TODO should I consider getting this from a accessor or a keyed dictionary?
	participant := f.ACHParticipants[0]

	if participant.RoutingNumber != "073905527" {
		t.Errorf("CustomerName Expected '073905527' got: %v", participant.RoutingNumber)
	}
	if participant.OfficeCode != "O" {
		t.Errorf("OfficeCode Expected 'O' got: %v", participant.OfficeCode)
	}
	if participant.ServicingFrbNumber != "071000301" {
		t.Errorf("ServicingFrbNumber Expected '071000301' got: %v", participant.ServicingFrbNumber)
	}
	if participant.RecordTypeCode != "1" {
		t.Errorf("RecordTypeCode Expected '1' got: %v", participant.RecordTypeCode)
	}
	if participant.Revised != "012908" {
		t.Errorf("Revised Expected '012908' got: %v", participant.Revised)
	}
	if participant.NewRoutingNumber != "000000000" {
		t.Errorf("NewRoutingNumber Expected '000000000' got: %v", participant.NewRoutingNumber)
	}
	if participant.CustomerName != "LINCOLN SAVINGS BANK" {
		t.Errorf("CustomerName Expected 'LINCOLN SAVINGS BANK' got: %v", participant.CustomerName)
	}
	if participant.Address != "P O BOX E" {
		t.Errorf("Address Expected 'P O BOX E' got: %v", participant.Address)
	}
	if participant.City != "REINBECK" {
		t.Errorf("City Expected 'REINBECK' got: %v", participant.City)
	}
	if participant.State != "IA" {
		t.Errorf("State Expected 'REINBECK' got: %v", participant.State)
	}
	if participant.PostalCode != "50669" {
		t.Errorf("PostalCode Expected '50669' got: %v", participant.PostalCode)
	}
	if participant.PostalCodeExtension != "0159" {
		t.Errorf("PostalCodeExtension Expected '0159' got: %v", participant.PostalCodeExtension)
	}
	if participant.PhoneNumber != "3197886441" {
		t.Errorf("PhoneNumber Expected '3197886441' got: %v", participant.PhoneNumber)
	}
	if participant.StatusCode != "1" {
		t.Errorf("StatusCode Expected '1' got: %v", participant.StatusCode)
	}
	if participant.ViewCode != "1" {
		t.Errorf("ViewCode Expected '1' got: %v", participant.ViewCode)
	}
}

func TestACHDirectoryRead(t *testing.T) {
	f, err := os.Open("./data/FedACHdir.txt")
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	achDir := NewACHDictionary(f)
	err = achDir.Read()
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if len(achDir.ACHParticipants) != 18198 {
		t.Errorf("Expected '19189' got: %v", len(achDir.ACHParticipants))
	}

	if fi, ok := achDir.IndexACHParticipant["073905527"]; ok {
		if !ok {
			t.Errorf("Expected `073905527` got : %v", fi.RoutingNumber)
		}
	}
}

func TestInvalidRecordLength(t *testing.T) {
	var line = "073905527O0710003011012908000000000LINCOLN SAVINGS BANK                P O BOX E"
	f := NewACHDictionary(strings.NewReader(line))
	if err := f.Read(); err != nil {
		if !Has(err, NewRecordWrongLengthErr(80)) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestParticipantLabel(t *testing.T) {
	var line = "073905527O0710003011012908000000000LINCOLN SAVINGS BANK                P O BOX E                           REINBECK            IA506690159319788644111     "

	f := NewACHDictionary(strings.NewReader(line))
	f.Read()

	// TODO should I consider getting this from a accessor or a keyed dictionary?
	participant := f.ACHParticipants[0]

	if participant.CustomerNameLabel() != "Lincoln Savings Bank" {
		t.Errorf("CustomerNameLabel Expected 'Lincoln Savings Bank' got: %v", participant.CustomerNameLabel())
	}

}

// TestRoutingNumberSearch tests that a valid routing number defined in FedACHDir returns the participant data
func TestRoutingNumberSearch(t *testing.T) {
	f, err := os.Open("./data/FedACHdir.txt")
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	achDir := NewACHDictionary(f)
	err = achDir.Read()
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}

	p := achDir.RoutingNumberSearch("325183657")

	if p.CustomerName != "LOWER VALLEY CU" {
		t.Errorf("Expected `LOWER VALLEY CU` got : %v", p.CustomerName)
	}
}

// TestInvalidRoutingNumberSearch tests that an invalid routing number returns nil
func TestInvalidRoutingNumberSearch(t *testing.T) {
	f, err := os.Open("./data/FedACHdir.txt")
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	achDir := NewACHDictionary(f)
	err = achDir.Read()
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}

	p := achDir.RoutingNumberSearch("433")

	if p != nil {
		t.Errorf("%s", "433 should have returned nil")
	}
}
