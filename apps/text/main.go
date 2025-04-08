package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Dice faces represented as strings
var diceFaces = [][]string{
	{
		"+-----+",
		"|     |",
		"|  *  |",
		"|     |",
		"+-----+",
	},
	{
		"+-----+",
		"| *   |",
		"|     |",
		"|   * |",
		"+-----+",
	},
	{
		"+-----+",
		"| *   |",
		"|  *  |",
		"|   * |",
		"+-----+",
	},
	{
		"+-----+",
		"| * * |",
		"|     |",
		"| * * |",
		"+-----+",
	},
	{
		"+-----+",
		"| * * |",
		"|  *  |",
		"| * * |",
		"+-----+",
	},
	{
		"+-----+",
		"| * * |",
		"| * * |",
		"| * * |",
		"+-----+",
	},
}

func main() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano())) // Create a new random generator
	balance := 100                                         // Starting balance

	fmt.Println("Welcome to my establishment.")
	fmt.Println("How about a nice game of craps?")
	fmt.Println()
	fmt.Printf("You start with $%d.\n", balance)

	roll := 0
	for {
		if balance <= 0 {
			fmt.Println("You have no money left. Game over!")
			break
		}

		fmt.Printf("\nRoll %d: You have $%d. Enter your bet: ", roll, balance)
		var bet int
		fmt.Scan(&bet)

		// Validate the bet
		if bet <= 0 || bet > balance {
			fmt.Println(strings.Join([]string{
				"",
				"Look at the sign on the wall:",
				"",
				"+--------------------+",
				"| This establishment |",
				"|   Does not accept  |",
				"|       credit.      |",
				"+--------------------+",
			}, "\n"))
			continue
		}

		// Generate random numbers for each die (1 to 6)
		die1 := rng.Intn(6)
		die2 := rng.Intn(6)

		fmt.Printf("\nRoll %d:\n", roll)
		for i := 0; i < len(diceFaces[die1]); i++ {
			// Print corresponding lines of the two dice side by side
			fmt.Printf("%s   %s\n", diceFaces[die1][i], diceFaces[die2][i])
		}

		// Calculate the sum of the dice
		sum := die1 + 1 + die2 + 1 // Add 1 because diceFaces are 0-indexed
		fmt.Printf("\nYou rolled a %d.\n", sum)

		// Determine win or loss
		if sum == 7 || sum == 11 {
			winnings := bet * 2
			balance += winnings
			fmt.Printf("You win $%d!\n", bet*2)
		} else {
			balance -= bet
			fmt.Printf("You lose $%d.\n", bet)
		}

		time.Sleep(1 * time.Second) // Pause for 1 second between rolls
		roll++
	}

	fmt.Printf("\nGame over! You leave with $%d after %d rolls.\n", balance, roll)
}
