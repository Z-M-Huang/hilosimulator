package hilosimulator

import (
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/Z-M-Huang/provablyfair"
)

var client *provablyfair.Client

func init() {
	serverSeed, err := provablyfair.GenerateNewSeed(32)
	if err != nil {
		panic(err)
	}
	client = &provablyfair.Client{
		ServerSeed: serverSeed,
	}
}

//Simulate simulate game outcome
func Simulate(config *Configuration) ([]*SimulationResult, error) {
	valid, err := validateConfig(config)
	if !valid {
		return nil, err
	}

	totalStack := config.TotalStack
	odds := config.Odds
	baseBet := config.BaseBet
	winChance := config.WinChance
	betHi := false
	var ret []*SimulationResult
	clientSeed, err := provablyfair.GenerateNewSeed(32)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate client seed. %s", err.Error())
	}

	for i := uint64(0); i < config.RollAmount; i++ {
		if totalStack < baseBet {
			break
		}
		simResult := &SimulationResult{}
		profit := float64(0)
		won := false
		if config.RandomClientSeed {
			clientSeed, err = provablyfair.GenerateNewSeed(32)
			if err != nil {
				return nil, fmt.Errorf("Failed to generate client seed. %s", err.Error())
			}
		}
		simResult.ClientSeed = hex.EncodeToString(clientSeed)

		if config.AlternateHiLo {
			betHi = !betHi
		}

		rollNum, serverSeed, nonce, err := client.Generate(clientSeed)
		if err != nil {
			return nil, err
		}
		simResult.Roll = rollNum
		simResult.Nonce = nonce
		simResult.ServerSeed = hex.EncodeToString(serverSeed)

		if betHi {
			simResult.Bet = "hi"
			target := 100 - winChance
			won = rollNum >= target
		} else {
			simResult.Bet = "lo"
			won = rollNum <= winChance
		}
		simResult.Won = won

		if won {
			profit = (odds - 1) * baseBet
		} else {
			profit = -1 * baseBet
		}
		simResult.Profit = profit

		if won && config.OnWin != nil {
			if config.OnWin.ReturnToBaseBet {
				baseBet = config.BaseBet
			} else if config.OnWin.IncreaseBet {
				baseBet = baseBet * (1 + config.OnWin.IncreaseBetBy)
			}

			if config.OnWin.ChangeOdds {
				odds = config.OnWin.ChangeOddsTo
				winChance = config.OnWin.NewWinChance
			}
		} else if !won && config.OnLoss != nil {
			if config.OnLoss.ReturnToBaseBet {
				baseBet = config.BaseBet
			} else if config.OnLoss.IncreaseBet {
				baseBet = baseBet * (1 + config.OnLoss.IncreaseBetBy)
			}

			if config.OnLoss.ChangeOdds {
				odds = config.OnLoss.ChangeOddsTo
				winChance = config.OnLoss.NewWinChance
			}
		}

		totalStack += profit
		simResult.Stack = totalStack

		ret = append(ret, simResult)
	}
	return ret, nil
}

//Verify verify num
func Verify(clientSeed, serverSeed string, nonce uint64, roll float64) (bool, error) {
	return provablyfair.VerifyFromString(clientSeed, serverSeed, nonce, roll)
}

func validateConfig(config *Configuration) (bool, error) {
	//TotalStac
	if !(config.TotalStack > 0) {
		return false, errors.New("Total stack is too low to run simulation")
	} else if config.TotalStack < config.BaseBet {
		return false, fmt.Errorf("Total stack is too low to run simulation with base bet of %f", config.BaseBet)
	}

	//WinChance
	if !(config.WinChance > 0 && config.WinChance < 100) {
		return false, errors.New("Win chance needs to be between 0 and 100")
	}

	//Odds
	if !(config.Odds > 0) {
		return false, errors.New("Odds cannot be less than 0")
	}

	//BaseBet
	if !(config.BaseBet > 0) {
		return false, errors.New("Base bet needs to be greater than 0")
	}

	//RollAmount
	if !(config.RollAmount > 0) {
		return false, errors.New("Roll amount needs to be greater than 0")
	}

	//OnWin
	if config.OnWin != nil {
		valid, err := validateConditionalChangeConfiguration(config.OnWin)
		if !valid {
			return false, err
		}
	}

	//OnLoss
	if config.OnLoss != nil {
		valid, err := validateConditionalChangeConfiguration(config.OnLoss)
		if !valid {
			return false, err
		}
	}

	return true, nil
}

func validateConditionalChangeConfiguration(config *ConditionalChangeConfiguration) (bool, error) {
	//ReturnToBaseBet
	if config.ReturnToBaseBet && config.IncreaseBet {
		return false, errors.New("Return to base bet and Increase bet cannot be both true")
	}

	if config.ChangeOdds && !(config.NewWinChance > 0 && config.NewWinChance < 100) {
		return false, errors.New("New win chance can only be between 0 and 100")
	}
	return true, nil
}
