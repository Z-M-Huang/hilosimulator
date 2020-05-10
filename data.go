package hilosimulator

//Configuration simulation configuration
type Configuration struct {
	TotalStack float64
	Odds       float64
	//WinChance needs to be < 100 and > 0
	WinChance        float64
	BaseBet          float64
	RollAmount       uint64
	OnWin            *ConditionalChangeConfiguration
	OnLoss           *ConditionalChangeConfiguration
	RandomClientSeed bool
	AlternateHiLo    bool
}

//ConditionalChangeConfiguration conditional change on win or loss
type ConditionalChangeConfiguration struct {
	ReturnToBaseBet bool
	IncreaseBet     bool
	IncreaseBetBy   float64
	ChangeOdds      bool
	ChangeOddsTo    float64
	//NewWinChance needs to be < 100 and > 0
	NewWinChance float64
}

//SimulationResult simulation result in array format
type SimulationResult struct {
	ServerSeed string
	ClientSeed string
	Nonce      uint64
	Roll       float64

	Won bool
	//Bet `hi` or `lo` indicating bet high or low
	Bet       string
	Odds      float64
	WinChance float64

	Profit float64
	Stack  float64
}
