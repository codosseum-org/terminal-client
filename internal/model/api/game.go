package modelapi

type Submission struct {
	Language string // The language of the solution attempt, as an identifier supported by Piston
	Code     string // The code to test, which is assumed to be in the language specified
}

// Information about a game round
type Round struct {
	Challenge Challenge           // Challenge for that round
	Over      bool                // True if the round is over, false otherwise
	Results   []PlayerRoundResult // (Partial) results of the round, depending on its progress. Ordered from best to worst.
}

type PlayerRoundResult struct {
	Name      string // Name of the player
	Language  string // Language they used
	Code      string // Code submitted by the player. Optional.
	ByteCount int    // Number of bytes in the solution (relevant for golfing modes)
	TimeLeft  int    // Time left on the clock (in seconds) when they pressed submit
	Score     struct {
		TestsRun       int     // Number of tests run
		TestsPassed    int     // Number of tests passed
		AverageRuntime float64 // Average runtime in milliseconds
	}
}

// The final result of a player, after all rounds (and therefore the game) are over.
// Not defined in the OpenAPI spec (yet).
type PlayerGameResult struct {
}

// Information about a game
type Game struct {
	AllowedLanguages []string  // Which programming languages are allowed to be used for submissions. If omitted, all languages are allowed.
	GameMode         GameMode  // Game mode of this game
	MaxPlayers       int       // Maximum player count for the game
	TimeLimit        int       // Time limit for the game, in seconds.
	MaxWarmupTime    int       // Maximum warmup time for a game in seconds - that is, a timer that starts once 2 or more players have joined the game that delays the starting of the game to allow more players to join.
	ID               string    // Game ID
	State            GameState // Current game state
	TimeLeft         int       // Time left on the clock for the current game state. If state is "gameOver", this value is absent
	Players          Players
	Results          []PlayerGameResult // Game results. This is different from *round results* in that it combines the results from all rounds thus far and interprets them according to the game mode being played. Progressive/partial results are given throughout the game, even if it is not over yet. Ordered from best (so far) to worst.
	Round            int                // Number of the currently active round ("inProgress") or the most recent round ("roundOver", "gameOver"). Absent if state = "warmup"
}

// Information about the players in a game
type Players struct {
	Players []Player // The names of the players currently in the game
	Admin   string   // The name of the player who created the game
}

// The available game states
type GameState string

// A game mode supported by codosseum
type GameMode string

// Settings for a Codosseum game
type GameSettings struct {
	AllowedLanguages []string // Which programming languages are allowed to be used for submissions. If omitted, all languages are allowed.
	GameMode         GameMode // Game mode of this game
	MaxPlayers       int      // Maximum player count for the game
	TimeLimit        int      // Time limit for the game, in seconds.
	MaxWarmupTime    int      // Maximum warmup time for a game in seconds - that is, a timer that starts once 2 or more players have joined the game that delays the starting of the game to allow more players to join.
}

// Not defined in the OpenAPI spec (yet).
type Challenge struct {
}
