package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type Game struct {
	word           string
	guessedLetters map[rune]bool
	attemptsLeft   int
	maxAttempts    int
}

func NewGame(word string, maxAttempts int) *Game {
	return &Game{
		word:           strings.ToUpper(word),
		guessedLetters: make(map[rune]bool),
		attemptsLeft:   maxAttempts,
		maxAttempts:    maxAttempts,
	}
}

func (g *Game) GetDisplayWord() string {
	var display strings.Builder
	for _, char := range g.word {
		if g.guessedLetters[char] {
			display.WriteRune(char)
		} else {
			display.WriteRune('_')
		}
		display.WriteRune(' ')
	}
	return strings.TrimSpace(display.String())
}
func (g *Game) GuessLetter(letter rune) (correct bool, alreadyGuessed bool) {
	letter = rune(strings.ToUpper(string(letter))[0])

	if g.guessedLetters[letter] {
		return false, true
	}

	g.guessedLetters[letter] = true

	if strings.ContainsRune(g.word, letter) {
		return true, false
	}

	g.attemptsLeft--
	return false, false
}

func (g *Game) isWon() bool {
	for _, char := range g.word {
		if !g.guessedLetters[char] {
			return false
		}
	}
	return true
}
func (g *Game) isLost() bool {
	return g.attemptsLeft <= 0
}
func (g *Game) GetGuessedLetters() string {
	var letters []rune
	for letter := range g.guessedLetters {
		letters = append(letters, letter)
	}
	for i := 0; i < len(letters); i++ {
		for j := i + 1; j < len(letters); j++ {
			if letters[i] > letters[j] {
				letters[i], letters[j] = letters[j], letters[i]
			}
		}
	}
	return string(letters)
}

func DrawHangman(attemptsLeft, maxAttempt int) {
	stages := []string{
		`
  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
=========`,
		`
  +---+
  |   |
  O   |
 /|\  |
 /    |
      |
=========`,
		`
  +---+
  |   |
  O   |
 /|\  |
      |
      |
=========`,
		`
  +---+
  |   |
  O   |
 /|   |
      |
      |
=========`,
		`
  +---+
  |   |
  O   |
  |   |
      |
      |
=========`,
		`
  +---+
  |   |
  O   |
      |
      |
      |
=========`,
		`
  +---+
  |   |
      |
      |
      |
      |
=========`,
	}
	index := attemptsLeft
	if index >= len(stages) {
		index = len(stages) - 1
	}
	fmt.Println(stages[index])
}

var WordBank = []string{
	"PROGRAMMING", "COMPUTER", "ALGORITHM", "FUNCTION",
	"VARIABLE", "DATABASE", "NETWORK", "SECURITY",
	"INTERFACE", "DEVELOPER", "SOFTWARE", "HARDWARE",
	"TERMINAL", "KEYBOARD", "MONITOR", "POINTER",
}

func GetRandomWord() string {
	return WordBank[rand.Intn(len(WordBank))]
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func DisplayGameState(game *Game) {
	ClearScreen()
	fmt.Println("=== HANGMAN GAME ===\n")
	DrawHangman(game.attemptsLeft, game.maxAttempts)
	fmt.Printf("\nWord: %s\n", game.GetDisplayWord())
	fmt.Printf("Attempts left: %d\n", game.attemptsLeft)
	fmt.Printf("Guessed letters: %s\n", game.GetGuessedLetters())
	fmt.Println()
}

func GetPlayerInput(reader *bufio.Reader) (rune, error) {
	fmt.Print("Enter a letter: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	if len(input) != 1 {
		return 0, fmt.Errorf("please enter exactly one letter")
	}

	letter := rune(input[0])
	if !((letter >= 'a' && letter <= 'z') || (letter >= 'A' && letter <= 'Z')) {
		return 0, fmt.Errorf("please enter a valid letter")
	}

	return letter, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to Hangman!")
	fmt.Println("Press Enter to start...")

	reader.ReadString('\n')

	word := GetRandomWord()
	game := NewGame(word, 6)

	for !game.isWon() && !game.isLost() {
		DisplayGameState(game)

		letter, err := GetPlayerInput(reader)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			fmt.Println("Press Enter to continue...")
			reader.ReadString('\n')
			continue
		}

		correct, alreadyGuessed := game.GuessLetter(letter)

		if alreadyGuessed {
			fmt.Println("You have already guessed that letter!")
		} else if correct {
			fmt.Println("Correct! Good guess!")
		} else {
			fmt.Println("Wrong! That letter is not in the word.")
		}

		fmt.Println("Press Enter to continue...")
		reader.ReadString('\n')
	}
	DisplayGameState(game)
	if game.isWon() {
		fmt.Println("🎉 Congratulations! You won!")
	} else {
		fmt.Printf("😢 Game Over! The word was: %s\n", game.word)
	}
}
