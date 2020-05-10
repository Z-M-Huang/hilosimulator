package hilowsimulator

import (
	"testing"
)

func TestSimulate(t *testing.T) {
	config := &Configuration{
		TotalStack: 100000000,
		Odds:       2,
		WinChance:  47.5,
		BaseBet:    100,
		RollAmount: 100,
		OnWin: &ConditionalChangeConfiguration{
			ReturnToBaseBet: true,
		},
		OnLoss: &ConditionalChangeConfiguration{
			ReturnToBaseBet: false,
			IncreaseBet:     true,
			IncreaseBetBy:   1,
		},
		RandomClientSeed: true,
		AlternateHiLo:    true,
	}

	ret, err := Simulate(config)

	if err != nil {
		t.Error(err)
	} else if len(ret) != 100 {
		t.Error("Invalid amount of simulations returned")
	}
}

func TestVerify(t *testing.T) {
	serverSeed := "71ef42d840a9bb600f5e57dc22acb07f5675eae0c7b52c18db2f1fa42ac92b0a"
	clientSeed := "e3e154b55fe29671ca2f48b7dd58210b38cfae7f168e6639b350767f8c2f58ad"
	nonce := uint64(0)
	rollNum := float64(24.88)

	valid, err := Verify(clientSeed, serverSeed, nonce, rollNum)

	if err != nil {
		t.Error(err)
	} else if !valid {
		t.Error("Result invalid")
	}
}
