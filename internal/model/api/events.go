package modelapi

type GameUpdateEvent struct {
	GameSettings GameSettings
}

type PlayerJoinEvent struct {
	Player Player
}

type PlayerLeaveEvent struct {
	Player Player
}

// Fired when a new round starts
type RoundStartEvent struct {
	Challenge Challenge // Challenge for this round
	Number    int       // Round number
}

// Fired when a player clicks submit (but the results are not in yet)
type PlayerSubmitEvent struct {
	Name string // Player name
	Lang string
	Code string
}

// Fired when the submitted code of a player becomes available to the consumer of the event stream. Players see opponents' code as soon as they have submitted their code themselves. Everyone else sees the players' solutions after the round has ended.
type PlayerCodeRevealEvent struct {
	Name string // Player name
	Lang string // Language used
	Code string // Code they submitted
}

// Fired after a player's solution has gone through all submission tests and the final result for the player has been determined. This event contains all values that were sent in the PlayerSubmitEvent before as well as the score of the player.
type PlayerFinishEvent struct {
	PlayerRoundResult PlayerRoundResult
}

// Sent (asynchronously) in response to a request asking to run a test for a solution. This event is sent for every individual test that fails or succeeds.
type TestResultEvent struct {
	TestID string // The ID returned when the test runs were requested. This can be used to determine the state of the solution when the test request was made.
	Num    int    // Test number that was run
	// Test result. input, expectedOutput and actualOutput are only sent if the test failed.
	Result struct {
		Pass           bool   // Whether the test passed or not
		Input          string // Test input
		ExpectedOutput string // Correct output
		ActualOutput   string // Output produced by the player's code
	}
}

// Sent to synchronise clock and game state between server and clients.
type SyncEvent struct {
	State    GameState // Current game state
	TimeLeft int       // Time left in seconds until "next thing" (warmup countdown, round timer, next rount countdown)
}

// Sent to a player if they have been eliminated from the game. This is not relevant for all game modes.
type EliminatedEvent struct {
	Message string // Message detailing elimination.
}

// Fired after the last round of a game is over. Contains final results of the game (may vary depending on game mode)
type GameOverEvent struct {
	FinalResults []PlayerGameResult
}
