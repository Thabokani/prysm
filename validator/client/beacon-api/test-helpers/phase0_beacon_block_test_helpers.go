package test_helpers

import (
	"github.com/prysmaticlabs/prysm/v4/api/server/structs"
	ethpb "github.com/prysmaticlabs/prysm/v4/proto/prysm/v1alpha1"
)

func GenerateProtoPhase0BeaconBlock() *ethpb.BeaconBlock {
	return &ethpb.BeaconBlock{
		Slot:          1,
		ProposerIndex: 2,
		ParentRoot:    FillByteSlice(32, 3),
		StateRoot:     FillByteSlice(32, 4),
		Body: &ethpb.BeaconBlockBody{
			RandaoReveal: FillByteSlice(96, 5),
			Eth1Data: &ethpb.Eth1Data{
				DepositRoot:  FillByteSlice(32, 6),
				DepositCount: 7,
				BlockHash:    FillByteSlice(32, 8),
			},
			Graffiti: FillByteSlice(32, 9),
			ProposerSlashings: []*ethpb.ProposerSlashing{
				{
					Header_1: &ethpb.SignedBeaconBlockHeader{
						Header: &ethpb.BeaconBlockHeader{
							Slot:          10,
							ProposerIndex: 11,
							ParentRoot:    FillByteSlice(32, 12),
							StateRoot:     FillByteSlice(32, 13),
							BodyRoot:      FillByteSlice(32, 14),
						},
						Signature: FillByteSlice(96, 15),
					},
					Header_2: &ethpb.SignedBeaconBlockHeader{
						Header: &ethpb.BeaconBlockHeader{
							Slot:          16,
							ProposerIndex: 17,
							ParentRoot:    FillByteSlice(32, 18),
							StateRoot:     FillByteSlice(32, 19),
							BodyRoot:      FillByteSlice(32, 20),
						},
						Signature: FillByteSlice(96, 21),
					},
				},
				{
					Header_1: &ethpb.SignedBeaconBlockHeader{
						Header: &ethpb.BeaconBlockHeader{
							Slot:          22,
							ProposerIndex: 23,
							ParentRoot:    FillByteSlice(32, 24),
							StateRoot:     FillByteSlice(32, 25),
							BodyRoot:      FillByteSlice(32, 26),
						},
						Signature: FillByteSlice(96, 27),
					},
					Header_2: &ethpb.SignedBeaconBlockHeader{
						Header: &ethpb.BeaconBlockHeader{
							Slot:          28,
							ProposerIndex: 29,
							ParentRoot:    FillByteSlice(32, 30),
							StateRoot:     FillByteSlice(32, 31),
							BodyRoot:      FillByteSlice(32, 32),
						},
						Signature: FillByteSlice(96, 33),
					},
				},
			},
			AttesterSlashings: []*ethpb.AttesterSlashing{
				{
					Attestation_1: &ethpb.IndexedAttestation{
						AttestingIndices: []uint64{34, 35},
						Data: &ethpb.AttestationData{
							Slot:            36,
							CommitteeIndex:  37,
							BeaconBlockRoot: FillByteSlice(32, 38),
							Source: &ethpb.Checkpoint{
								Epoch: 39,
								Root:  FillByteSlice(32, 40),
							},
							Target: &ethpb.Checkpoint{
								Epoch: 41,
								Root:  FillByteSlice(32, 42),
							},
						},
						Signature: FillByteSlice(96, 43),
					},
					Attestation_2: &ethpb.IndexedAttestation{
						AttestingIndices: []uint64{44, 45},
						Data: &ethpb.AttestationData{
							Slot:            46,
							CommitteeIndex:  47,
							BeaconBlockRoot: FillByteSlice(32, 48),
							Source: &ethpb.Checkpoint{
								Epoch: 49,
								Root:  FillByteSlice(32, 50),
							},
							Target: &ethpb.Checkpoint{
								Epoch: 51,
								Root:  FillByteSlice(32, 52),
							},
						},
						Signature: FillByteSlice(96, 53),
					},
				},
				{
					Attestation_1: &ethpb.IndexedAttestation{
						AttestingIndices: []uint64{54, 55},
						Data: &ethpb.AttestationData{
							Slot:            56,
							CommitteeIndex:  57,
							BeaconBlockRoot: FillByteSlice(32, 58),
							Source: &ethpb.Checkpoint{
								Epoch: 59,
								Root:  FillByteSlice(32, 60),
							},
							Target: &ethpb.Checkpoint{
								Epoch: 61,
								Root:  FillByteSlice(32, 62),
							},
						},
						Signature: FillByteSlice(96, 63),
					},
					Attestation_2: &ethpb.IndexedAttestation{
						AttestingIndices: []uint64{64, 65},
						Data: &ethpb.AttestationData{
							Slot:            66,
							CommitteeIndex:  67,
							BeaconBlockRoot: FillByteSlice(32, 68),
							Source: &ethpb.Checkpoint{
								Epoch: 69,
								Root:  FillByteSlice(32, 70),
							},
							Target: &ethpb.Checkpoint{
								Epoch: 71,
								Root:  FillByteSlice(32, 72),
							},
						},
						Signature: FillByteSlice(96, 73),
					},
				},
			},
			Attestations: []*ethpb.Attestation{
				{
					AggregationBits: FillByteSlice(32, 74),
					Data: &ethpb.AttestationData{
						Slot:            75,
						CommitteeIndex:  76,
						BeaconBlockRoot: FillByteSlice(32, 77),
						Source: &ethpb.Checkpoint{
							Epoch: 78,
							Root:  FillByteSlice(32, 79),
						},
						Target: &ethpb.Checkpoint{
							Epoch: 80,
							Root:  FillByteSlice(32, 81),
						},
					},
					Signature: FillByteSlice(96, 82),
				},
				{
					AggregationBits: FillByteSlice(4, 83),
					Data: &ethpb.AttestationData{
						Slot:            84,
						CommitteeIndex:  85,
						BeaconBlockRoot: FillByteSlice(32, 38),
						Source: &ethpb.Checkpoint{
							Epoch: 87,
							Root:  FillByteSlice(32, 88),
						},
						Target: &ethpb.Checkpoint{
							Epoch: 89,
							Root:  FillByteSlice(32, 90),
						},
					},
					Signature: FillByteSlice(96, 91),
				},
			},
			Deposits: []*ethpb.Deposit{
				{
					Proof: FillByteArraySlice(33, FillByteSlice(32, 92)),
					Data: &ethpb.Deposit_Data{
						PublicKey:             FillByteSlice(48, 94),
						WithdrawalCredentials: FillByteSlice(32, 95),
						Amount:                96,
						Signature:             FillByteSlice(96, 97),
					},
				},
				{
					Proof: FillByteArraySlice(33, FillByteSlice(32, 98)),
					Data: &ethpb.Deposit_Data{
						PublicKey:             FillByteSlice(48, 100),
						WithdrawalCredentials: FillByteSlice(32, 101),
						Amount:                102,
						Signature:             FillByteSlice(96, 103),
					},
				},
			},
			VoluntaryExits: []*ethpb.SignedVoluntaryExit{
				{
					Exit: &ethpb.VoluntaryExit{
						Epoch:          104,
						ValidatorIndex: 105,
					},
					Signature: FillByteSlice(96, 106),
				},
				{
					Exit: &ethpb.VoluntaryExit{
						Epoch:          107,
						ValidatorIndex: 108,
					},
					Signature: FillByteSlice(96, 109),
				},
			},
		},
	}
}

func GenerateJsonPhase0BeaconBlock() *structs.BeaconBlock {
	return &structs.BeaconBlock{
		Slot:          "1",
		ProposerIndex: "2",
		ParentRoot:    FillEncodedByteSlice(32, 3),
		StateRoot:     FillEncodedByteSlice(32, 4),
		Body: &structs.BeaconBlockBody{
			RandaoReveal: FillEncodedByteSlice(96, 5),
			Eth1Data: &structs.Eth1Data{
				DepositRoot:  FillEncodedByteSlice(32, 6),
				DepositCount: "7",
				BlockHash:    FillEncodedByteSlice(32, 8),
			},
			Graffiti: FillEncodedByteSlice(32, 9),
			ProposerSlashings: []*structs.ProposerSlashing{
				{
					SignedHeader1: &structs.SignedBeaconBlockHeader{
						Message: &structs.BeaconBlockHeader{
							Slot:          "10",
							ProposerIndex: "11",
							ParentRoot:    FillEncodedByteSlice(32, 12),
							StateRoot:     FillEncodedByteSlice(32, 13),
							BodyRoot:      FillEncodedByteSlice(32, 14),
						},
						Signature: FillEncodedByteSlice(96, 15),
					},
					SignedHeader2: &structs.SignedBeaconBlockHeader{
						Message: &structs.BeaconBlockHeader{
							Slot:          "16",
							ProposerIndex: "17",
							ParentRoot:    FillEncodedByteSlice(32, 18),
							StateRoot:     FillEncodedByteSlice(32, 19),
							BodyRoot:      FillEncodedByteSlice(32, 20),
						},
						Signature: FillEncodedByteSlice(96, 21),
					},
				},
				{
					SignedHeader1: &structs.SignedBeaconBlockHeader{
						Message: &structs.BeaconBlockHeader{
							Slot:          "22",
							ProposerIndex: "23",
							ParentRoot:    FillEncodedByteSlice(32, 24),
							StateRoot:     FillEncodedByteSlice(32, 25),
							BodyRoot:      FillEncodedByteSlice(32, 26),
						},
						Signature: FillEncodedByteSlice(96, 27),
					},
					SignedHeader2: &structs.SignedBeaconBlockHeader{
						Message: &structs.BeaconBlockHeader{
							Slot:          "28",
							ProposerIndex: "29",
							ParentRoot:    FillEncodedByteSlice(32, 30),
							StateRoot:     FillEncodedByteSlice(32, 31),
							BodyRoot:      FillEncodedByteSlice(32, 32),
						},
						Signature: FillEncodedByteSlice(96, 33),
					},
				},
			},
			AttesterSlashings: []*structs.AttesterSlashing{
				{
					Attestation1: &structs.IndexedAttestation{
						AttestingIndices: []string{"34", "35"},
						Data: &structs.AttestationData{
							Slot:            "36",
							CommitteeIndex:  "37",
							BeaconBlockRoot: FillEncodedByteSlice(32, 38),
							Source: &structs.Checkpoint{
								Epoch: "39",
								Root:  FillEncodedByteSlice(32, 40),
							},
							Target: &structs.Checkpoint{
								Epoch: "41",
								Root:  FillEncodedByteSlice(32, 42),
							},
						},
						Signature: FillEncodedByteSlice(96, 43),
					},
					Attestation2: &structs.IndexedAttestation{
						AttestingIndices: []string{"44", "45"},
						Data: &structs.AttestationData{
							Slot:            "46",
							CommitteeIndex:  "47",
							BeaconBlockRoot: FillEncodedByteSlice(32, 48),
							Source: &structs.Checkpoint{
								Epoch: "49",
								Root:  FillEncodedByteSlice(32, 50),
							},
							Target: &structs.Checkpoint{
								Epoch: "51",
								Root:  FillEncodedByteSlice(32, 52),
							},
						},
						Signature: FillEncodedByteSlice(96, 53),
					},
				},
				{
					Attestation1: &structs.IndexedAttestation{
						AttestingIndices: []string{"54", "55"},
						Data: &structs.AttestationData{
							Slot:            "56",
							CommitteeIndex:  "57",
							BeaconBlockRoot: FillEncodedByteSlice(32, 58),
							Source: &structs.Checkpoint{
								Epoch: "59",
								Root:  FillEncodedByteSlice(32, 60),
							},
							Target: &structs.Checkpoint{
								Epoch: "61",
								Root:  FillEncodedByteSlice(32, 62),
							},
						},
						Signature: FillEncodedByteSlice(96, 63),
					},
					Attestation2: &structs.IndexedAttestation{
						AttestingIndices: []string{"64", "65"},
						Data: &structs.AttestationData{
							Slot:            "66",
							CommitteeIndex:  "67",
							BeaconBlockRoot: FillEncodedByteSlice(32, 68),
							Source: &structs.Checkpoint{
								Epoch: "69",
								Root:  FillEncodedByteSlice(32, 70),
							},
							Target: &structs.Checkpoint{
								Epoch: "71",
								Root:  FillEncodedByteSlice(32, 72),
							},
						},
						Signature: FillEncodedByteSlice(96, 73),
					},
				},
			},
			Attestations: []*structs.Attestation{
				{
					AggregationBits: FillEncodedByteSlice(32, 74),
					Data: &structs.AttestationData{
						Slot:            "75",
						CommitteeIndex:  "76",
						BeaconBlockRoot: FillEncodedByteSlice(32, 77),
						Source: &structs.Checkpoint{
							Epoch: "78",
							Root:  FillEncodedByteSlice(32, 79),
						},
						Target: &structs.Checkpoint{
							Epoch: "80",
							Root:  FillEncodedByteSlice(32, 81),
						},
					},
					Signature: FillEncodedByteSlice(96, 82),
				},
				{
					AggregationBits: FillEncodedByteSlice(4, 83),
					Data: &structs.AttestationData{
						Slot:            "84",
						CommitteeIndex:  "85",
						BeaconBlockRoot: FillEncodedByteSlice(32, 38),
						Source: &structs.Checkpoint{
							Epoch: "87",
							Root:  FillEncodedByteSlice(32, 88),
						},
						Target: &structs.Checkpoint{
							Epoch: "89",
							Root:  FillEncodedByteSlice(32, 90),
						},
					},
					Signature: FillEncodedByteSlice(96, 91),
				},
			},
			Deposits: []*structs.Deposit{
				{
					Proof: FillEncodedByteArraySlice(33, FillEncodedByteSlice(32, 92)),
					Data: &structs.DepositData{
						Pubkey:                FillEncodedByteSlice(48, 94),
						WithdrawalCredentials: FillEncodedByteSlice(32, 95),
						Amount:                "96",
						Signature:             FillEncodedByteSlice(96, 97),
					},
				},
				{
					Proof: FillEncodedByteArraySlice(33, FillEncodedByteSlice(32, 98)),
					Data: &structs.DepositData{
						Pubkey:                FillEncodedByteSlice(48, 100),
						WithdrawalCredentials: FillEncodedByteSlice(32, 101),
						Amount:                "102",
						Signature:             FillEncodedByteSlice(96, 103),
					},
				},
			},
			VoluntaryExits: []*structs.SignedVoluntaryExit{
				{
					Message: &structs.VoluntaryExit{
						Epoch:          "104",
						ValidatorIndex: "105",
					},
					Signature: FillEncodedByteSlice(96, 106),
				},
				{
					Message: &structs.VoluntaryExit{
						Epoch:          "107",
						ValidatorIndex: "108",
					},
					Signature: FillEncodedByteSlice(96, 109),
				},
			},
		},
	}
}
