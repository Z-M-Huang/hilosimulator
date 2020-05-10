# hilosimulator
Online bet game hi/low simulator

# Example Usage
## Simulate
```
  config := &hilowsimulator.Configuration{
		TotalStack: 100000000,
		Odds:       2,
		WinChance:  47.5,
		BaseBet:    100,
		RollAmount: 100,
		OnWin: &hilowsimulator.ConditionalChangeConfiguration{
			ReturnToBaseBet: true,
		},
		OnLoss: &hilowsimulator.ConditionalChangeConfiguration{
			ReturnToBaseBet: false,
			IncreaseBet:     true,
			IncreaseBetBy:   1,
		},
		RandomClientSeed: true,
		AlternateHiLo:    true,
	}

	result, err := Simulate(config)
  if err != nil {
    panic(err)
  }
```

## Verify
```
  valid, err := hilowsimulator.Verify(result[0].ClientSeed, result[0].ServerSeed, result[0].Nonce, result[0].Roll)
```