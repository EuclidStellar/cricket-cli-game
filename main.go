/*************************** Making evevrything work together ****************************/

/* If you want to know behind the scene of this code, then you can check out the https://x.com/EuclidStellar/status/1713094088969658756?s=20 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

)

func main() {
	fmt.Println("Welcome to Command Line Cricket!")

	var player1Score, player2Score, wickets int
	players := []string{"Player 1", "Player 2"}
	batsman := 0

	for over := 1; over <= 2; over++ {
		fmt.Printf("\n--- Over %d ---\n", over)

		for ball := 1; ball <= 6; ball++ {
			fmt.Printf("\n%s, enter your shot (hit, miss, or out): ", players[batsman])

			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			input := strings.ToLower(scanner.Text())

			switch input {
			case "hit":
				score := 1
				fmt.Printf("Good shot! %s scores %d run.\n", players[batsman], score)
				players[batsman] = strconv.Itoa(player1Score + score)
			case "miss":
				fmt.Printf("A swing and a miss!\n")
			case "out":
				fmt.Printf("Oh no! %s is out!\n", players[batsman])
				wickets++
				if wickets == 2 {
					fmt.Println("All out! Innings over.")
					break
				}
			default:
				fmt.Println("Invalid input. Please enter 'hit', 'miss', or 'out'.")
				ball-- // Reattempt the ball
			}

			batsman = (batsman + 1) % 2
		}

		if wickets == 2 {
			fmt.Println("Innings over. Both players have been dismissed.")
			break
		}
	}

	fmt.Println("\n--- Game Over ---")
	fmt.Printf("Player 1 scored %d runs.\n", player1Score)
	fmt.Printf("Player 2 scored %d runs.\n", player2Score)

	if player1Score > player2Score {
		fmt.Println("Player 1 wins!")
	} else if player1Score < player2Score {
		fmt.Println("Player 2 wins!")
	} else {
		fmt.Println("It's a tie!")
	}
}
