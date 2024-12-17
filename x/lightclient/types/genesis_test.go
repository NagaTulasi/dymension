package types_test

import (
	"testing"

	"github.com/dymensionxyz/dymension/v3/x/lightclient/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisValidate(t *testing.T) {
	tests := []struct {
		name  string
		g     types.GenesisState
		valid bool
	}{
		{
			name: "valid",
			g: types.GenesisState{
				CanonicalClients: []types.CanonicalClient{
					{RollappId: "rollapp-1", IbcClientId: "client-1"},
					{RollappId: "rollapp-2", IbcClientId: "client-2"},
				},
			},
			valid: true,
		},
		{
			name: "invalid rollapp id",
			g: types.GenesisState{
				CanonicalClients: []types.CanonicalClient{
					{RollappId: "", IbcClientId: "client-1"},
				},
			},
			valid: false,
		},
		{
			name: "invalid ibc client id",
			g: types.GenesisState{
				CanonicalClients: []types.CanonicalClient{
					{RollappId: "rollapp-1", IbcClientId: ""},
				},
			},
			valid: false,
		},

		{
			name:  "empty",
			g:     types.GenesisState{},
			valid: true,
		},
		{
			name:  "default",
			g:     types.DefaultGenesisState(),
			valid: true,
		},
		{
			name: "duplicate canonical client",
			g: types.GenesisState{
				CanonicalClients: []types.CanonicalClient{
					{RollappId: "rollapp-1", IbcClientId: "client-1"},
					{RollappId: "rollapp-1", IbcClientId: "client-1"},
				},
			},
			valid: false,
		},
		{
			name: "duplicate header signer",
			g: types.GenesisState{
				HeaderSigners: []types.HeaderSignerEntry{
					{SequencerAddress: "sequencer1", ClientId: "client1", Height: 100},
					{SequencerAddress: "sequencer1", ClientId: "client1", Height: 100},
				},
			},
			valid: false,
		},
		{
			name: "empty sequencer address in header signer",
			g: types.GenesisState{
				HeaderSigners: []types.HeaderSignerEntry{
					{SequencerAddress: "", ClientId: "client1", Height: 100},
				},
			},
			valid: false,
		},
		{
			name: "empty client id in header signer",
			g: types.GenesisState{
				HeaderSigners: []types.HeaderSignerEntry{
					{SequencerAddress: "sequencer1", ClientId: "", Height: 100},
				},
			},
			valid: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.valid {
				require.NoError(t, tt.g.Validate())
			} else {
				require.Error(t, tt.g.Validate())
			}
		})
	}
}
