// Copyright 2026 The Moov Authors
// Use of this source code is governed by an Apache License
// that can be found in the LICENSE file.

package fed

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMercuryChoiceSupplementalParticipants(t *testing.T) {
	t.Run("FedACH", func(t *testing.T) {
		_, dictionary := loadTestACHFiles(t)
		participant := dictionary.RoutingNumberSearchSingle("091311229")

		require.NotNil(t, participant)
		require.Equal(t, "O", participant.OfficeCode)
		require.Equal(t, "091000080", participant.ServicingFRBNumber)
		require.Equal(t, "1", participant.RecordTypeCode)
		require.Equal(t, "071525", participant.Revised)
		require.Equal(t, "000000000", participant.NewRoutingNumber)
		require.Equal(t, "MERCURY, PARTNERING WITH CHOICE BANK", participant.CustomerName)
		require.Equal(t, "4501 23RD AVE SOUTH", participant.Address)
		require.Equal(t, "FARGO", participant.City)
		require.Equal(t, "ND", participant.State)
		require.Equal(t, "58104", participant.PostalCode)
		require.Equal(t, "0000", participant.PostalCodeExtension)
		require.Equal(t, "8773803623", participant.PhoneNumber)
		require.Equal(t, "1", participant.StatusCode)
		require.Equal(t, "1", participant.ViewCode)

		var matches int
		for _, candidate := range dictionary.ACHParticipants {
			if candidate.RoutingNumber == "091311229" {
				matches++
			}
		}
		require.Equal(t, 1, matches)
	})

	t.Run("Fedwire", func(t *testing.T) {
		_, dictionary := loadTestWireFiles(t)
		participant := dictionary.RoutingNumberSearchSingle("091311229")

		require.NotNil(t, participant)
		require.Equal(t, "CHOICEFIN", participant.TelegraphicName)
		require.Equal(t, "CHOICE FINANCIAL GROUP", participant.CustomerName)
		require.Equal(t, "FARGO", participant.City)
		require.Equal(t, "ND", participant.State)
		require.Equal(t, "Y", participant.FundsTransferStatus)
		require.Equal(t, " ", participant.FundsSettlementOnlyStatus)
		require.Equal(t, "N", participant.BookEntrySecuritiesTransferStatus)
		require.Equal(t, "20211210", participant.Date)

		var matches int
		for _, candidate := range dictionary.WIREParticipants {
			if candidate.RoutingNumber == "091311229" {
				matches++
			}
		}
		require.Equal(t, 1, matches)
	})
}
