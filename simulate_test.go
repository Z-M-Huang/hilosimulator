package hilowsimulator

import "testing"

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
