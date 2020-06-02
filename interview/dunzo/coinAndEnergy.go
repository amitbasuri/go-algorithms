package dunzo

//Alex will visit a number of houses that are arranged in a line.
//Each house has a number of coins and an amount of energy available.
//The journey must begin at the first house, and each house along the journey must be visited.
//None can be skipped over, but Alex can end the journey at any point.
//It costs 1 unit of energy to move from one house to the next.
//Alex can collect either the energy or the coins when visiting a house.
//Determine the maximum number of coins Alex can collect while never having a negative energy amount.
//Example n = 3 initialEnergy = 0 energy = [2, 1, 1] coins = [11, 5, 7]
//There are n houses in a line. The ith house has energy[i] energy and coins[i] coins.
//Alex starts the journey at the first house with initialEnergy energy.
//The best approach is to collect 2 units of energy at the first house,
//then 5 + 7 = 12 coins at the second and third houses.
//Alex's energy level is 0 after moving to the second and third houses.

/*
 * Complete the 'getRich' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. LONG_INTEGER initialEnergy
 *  2. INTEGER_ARRAY energy
 *  3. INTEGER_ARRAY coins
 */

func GetRichCoinAndEnergy(initialEnergy int64, energy []int32, coins []int32) int32 {

	return solveGetRich(initialEnergy, energy, coins, 0)

}

func solveGetRich(initialEnergy int64, energy []int32, coins []int32, index int) int32 {

	if index >= len(energy) || initialEnergy < 0 {
		return 0
	}
	energyChoice := energy[index]
	coinChoice := coins[index]

	coinStolenChoice := solveGetRich(initialEnergy-1, energy, coins, index+1) + coinChoice
	coinNotStolenChoice := solveGetRich(initialEnergy-1+int64(energyChoice), energy, coins, index+1)

	if coinStolenChoice > coinNotStolenChoice {
		return coinStolenChoice
	}
	return coinNotStolenChoice

}
