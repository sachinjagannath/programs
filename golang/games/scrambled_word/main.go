package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// WordProvider interface defines how to get words
type WordProvider interface {
	GetRandomWord() string
	GetWordCount() int
}

// GameScorer interface defines scoring behavior
type GameScorer interface {
	AddScore(points int)
	GetScore() int
	CalculateBonus(timeRemaining int) int
}

// UIRenderer interface for display logic
type UIRenderer interface {
	ShowWelcome()
	ShowRound(roundNum int, scrambled string, wordLen int)
	ShowCorrect(word string, points, bonus, total, score int)
	ShowIncorrect(word string, score int)
	ShowTimeout(word string)
	ShowFinalScore(rounds, score int)
	ClearScreen()
}

// WordBank implements WordProvider
type WordBank struct {
	words []string
}

func NewWordBank() *WordBank {
	return &WordBank{
		words: []string{
			"computer", "programming", "golang", "developer", "keyboard",
			"algorithm", "function", "variable", "terminal", "software",
			"database", "network", "security", "interface", "compiler",
			"memory", "processor", "internet", "browser", "application",
			"struct", "method", "pointer", "channel", "goroutine",
			"package", "import", "export", "module", "library",
		},
	}
}

func (wb *WordBank) GetRandomWord() string {
	return wb.words[rand.Intn(len(wb.words))]
}

func (wb *WordBank) GetWordCount() int {
	return len(wb.words)
}

// ScoreKeeper implements GameScorer
type ScoreKeeper struct {
	score           int
	basePoints      int
	bonusMultiplier int
}

func NewScoreKeeper(basePoints, bonusMultiplier int) *ScoreKeeper {
	return &ScoreKeeper{
		score:           0,
		basePoints:      basePoints,
		bonusMultiplier: bonusMultiplier,
	}
}

func (sk *ScoreKeeper) AddScore(points int) {
	sk.score += points
}

func (sk *ScoreKeeper) GetScore() int {
	return sk.score
}

func (sk *ScoreKeeper) CalculateBonus(timeRemaining int) int {
	return timeRemaining * sk.bonusMultiplier
}

func (sk *ScoreKeeper) GetBasePoints() int {
	return sk.basePoints
}

// TerminalUI implements UIRenderer
type TerminalUI struct{}

func NewTerminalUI() *TerminalUI {
	return &TerminalUI{}
}

func (ui *TerminalUI) ShowWelcome() {
	fmt.Println("╔═══════════════════════════════════════════╗")
	fmt.Println("║     WORD SCRAMBLE - ANAGRAM GAME         ║")
	fmt.Println("╚═══════════════════════════════════════════╝")
	fmt.Println("\n📖 Rules:")
	fmt.Println("  • Unscramble the word within the time limit")
	fmt.Println("  • Each correct answer = 10 points")
	fmt.Println("  • Time bonus: remaining seconds × 2")
	fmt.Println("  • You have 30 seconds per word")
	fmt.Println("\nGood luck! 🎯")
}

func (ui *TerminalUI) ShowRound(roundNum int, scrambled string, wordLen int) {
	fmt.Printf("\n🎮 Round %d\n", roundNum)
	fmt.Println("═══════════════════════════════════════")
	fmt.Printf("\n🔤 Scrambled Word: %s\n", strings.ToUpper(scrambled))
	fmt.Printf("💡 Hint: %d letters\n\n", wordLen)
}

func (ui *TerminalUI) ShowCorrect(word string, points, bonus, total, score int) {
	fmt.Printf("\n\n✅ CORRECT! The word was: %s\n", strings.ToUpper(word))
	fmt.Printf("🎯 Base points: %d\n", points)
	fmt.Printf("⚡ Time bonus: %d\n", bonus)
	fmt.Printf("📊 Total points earned: %d\n", total)
	fmt.Printf("🏆 Your score: %d\n", score)
}

func (ui *TerminalUI) ShowIncorrect(word string, score int) {
	fmt.Printf("\n\n❌ Wrong! The correct word was: %s\n", strings.ToUpper(word))
	fmt.Printf("🏆 Your score: %d\n", score)
}

func (ui *TerminalUI) ShowTimeout(word string) {
	fmt.Printf("\n\n⏰ Time's up! The word was: %s\n", strings.ToUpper(word))
}

func (ui *TerminalUI) ShowFinalScore(rounds, score int) {
	fmt.Println("\n╔═══════════════════════════════════════════╗")
	fmt.Println("║           GAME OVER - FINAL SCORE        ║")
	fmt.Println("╚═══════════════════════════════════════════╝")
	fmt.Printf("\n📊 Rounds Played: %d\n", rounds)
	fmt.Printf("🏆 Final Score: %d\n", score)

	if rounds > 0 {
		avg := float64(score) / float64(rounds)
		fmt.Printf("📈 Average per Round: %.1f\n", avg)
	}

	fmt.Println("\n🎮 Thanks for playing!")
	fmt.Println("═══════════════════════════════════════")
}

func (ui *TerminalUI) ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

// GameConfig holds game configuration
type GameConfig struct {
	TimeLimit       int
	BasePoints      int
	BonusMultiplier int
}

func DefaultGameConfig() *GameConfig {
	return &GameConfig{
		TimeLimit:       30,
		BasePoints:      10,
		BonusMultiplier: 2,
	}
}

// Game orchestrates the game logic
type Game struct {
	config       *GameConfig
	wordProvider WordProvider
	scorer       *ScoreKeeper
	ui           UIRenderer
	reader       *bufio.Reader
	roundsPlayed int
}

func NewGame(config *GameConfig, wp WordProvider, scorer *ScoreKeeper, ui UIRenderer) *Game {
	return &Game{
		config:       config,
		wordProvider: wp,
		scorer:       scorer,
		ui:           ui,
		reader:       bufio.NewReader(os.Stdin),
		roundsPlayed: 0,
	}
}

func (g *Game) Start() {
	g.ui.ClearScreen()
	g.ui.ShowWelcome()

	for {
		fmt.Print("\nPress ENTER to start a new round (or type 'quit' to exit): ")
		input, _ := g.reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "quit" {
			break
		}

		g.ui.ClearScreen()
		g.playRound()
	}

	g.ui.ShowFinalScore(g.roundsPlayed, g.scorer.GetScore())
}

func (g *Game) playRound() {
	word := g.wordProvider.GetRandomWord()
	scrambled := scrambleWord(word)
	g.roundsPlayed++
	g.ui.ShowRound(g.roundsPlayed, scrambled, len(word))
	fmt.Printf("⏱️  Time limit: %d seconds - Type your answer below:\n\n", g.config.TimeLimit)
	fmt.Print("Your answer: ")

	// Buffered channel to receive the user's answer
	answerChan := make(chan string, 1)

	// Goroutine to get user input
	go func() {
		// ReadString will block until '\n' is entered.
		answer, _ := g.reader.ReadString('\n')
		// Clean and normalize the input before sending it to the channel
		answerChan <- strings.TrimSpace(strings.ToLower(answer))
	}()

	startTime := time.Now()
	timeoutTimer := time.NewTimer(time.Duration(g.config.TimeLimit) * time.Second)
	defer timeoutTimer.Stop()

	select {
	case answer := <-answerChan:
		// **CASE 1: User answered (and won the race against the timer)**

		// Calculate the time taken for scoring purposes
		elapsed := time.Since(startTime).Seconds()
		// Ensure timeRemaining is not negative for scoring calculation
		timeRemaining := int(float64(g.config.TimeLimit) - elapsed)
		if timeRemaining < 0 {
			timeRemaining = 0
		}

		if answer == word {
			// Answer is correct
			basePoints := g.scorer.GetBasePoints()
			bonus := g.scorer.CalculateBonus(timeRemaining)
			totalPoints := basePoints + bonus
			g.scorer.AddScore(totalPoints)
			g.ui.ShowCorrect(word, basePoints, bonus, totalPoints, g.scorer.GetScore())
		} else {
			// Answer is incorrect
			g.ui.ShowIncorrect(word, g.scorer.GetScore())
		}

	case <-timeoutTimer.C:
		// **CASE 2: Time ran out (and won the race against the answer)**

		fmt.Println("\n")
		g.ui.ShowTimeout(word)
		fmt.Print("\n[Press ENTER to continue...]")

		// Wait for the user's pending or future input and discard it.
		// This prevents the user's buffered input from being read in a later round
		// and prevents the goroutine from blocking indefinitely.
		<-answerChan
	}

	// Brief pause to allow the user to read the round summary
	time.Sleep(2 * time.Second)
}

//func (g *Game) playRound() {
//	word := g.wordProvider.GetRandomWord()
//	scrambled := scrambleWord(word)
//	g.roundsPlayed++
//	g.ui.ShowRound(g.roundsPlayed, scrambled, len(word))
//	fmt.Printf("⏱️  Time limit: %d seconds - Type your answer below:\n\n", g.config.TimeLimit)
//	fmt.Print("Your answer: ")
//	answerChan := make(chan string, 1)
//
//	// Goroutine to get user input
//	go func() {
//		answer, _ := g.reader.ReadString('\n')
//		answerChan <- strings.TrimSpace(strings.ToLower(answer))
//	}()
//
//	startTime := time.Now()
//	timeoutTimer := time.NewTimer(time.Duration(g.config.TimeLimit) * time.Second)
//	defer timeoutTimer.Stop()
//
//	select {
//	case answer := <-answerChan:
//		// User answered
//		elapsed := int(time.Since(startTime).Seconds())
//		timeRemaining := g.config.TimeLimit - elapsed
//
//		if timeRemaining <= 0 {
//			// Answer came in, but after timeout
//			fmt.Println("\n")
//			g.ui.ShowTimeout(word)
//		} else if answer == word {
//			basePoints := g.scorer.GetBasePoints()
//			bonus := g.scorer.CalculateBonus(timeRemaining)
//			totalPoints := basePoints + bonus
//			g.scorer.AddScore(totalPoints)
//			g.ui.ShowCorrect(word, basePoints, bonus, totalPoints, g.scorer.GetScore())
//		} else {
//			g.ui.ShowIncorrect(word, g.scorer.GetScore())
//		}
//
//	case <-timeoutTimer.C:
//		// Time ran out
//		fmt.Println("\n")
//		g.ui.ShowTimeout(word)
//		fmt.Print("\n[Press ENTER to continue...]")
//
//		// Wait for the user's input and discard it
//		<-answerChan
//	}
//
//	time.Sleep(2 * time.Second)
//}

// scrambleWord handles word scrambling logic
func scrambleWord(word string) string {
	runes := []rune(word)
	for i := 0; i < 20; i++ {
		a := rand.Intn(len(runes))
		b := rand.Intn(len(runes))
		runes[a], runes[b] = runes[b], runes[a]
	}

	scrambled := string(runes)
	if scrambled == word && len(word) > 1 {
		return scrambleWord(word)
	}
	return scrambled
}

func main() {
	// Note: rand.Seed() is deprecated in Go 1.20+
	// The random number generator is automatically seeded

	// Initialize dependencies
	config := DefaultGameConfig()
	wordBank := NewWordBank()
	scorer := NewScoreKeeper(config.BasePoints, config.BonusMultiplier)
	ui := NewTerminalUI()

	// Create and start game
	game := NewGame(config, wordBank, scorer, ui)
	game.Start()
}
