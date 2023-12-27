// messagingConventions.go

package config

const (
	// MessageDelimiter is the delimiter used to separate messages in the TCP communication
	MessageDelimiter = "\n"
	// ErrorMessage is the message sent when there is an error in the communication
	ErrorMessage = "ERROR"

	// SuccessMessage is the message sent when the communication is successful
	SuccessMessage = "SUCCESS"

	// OtherMessage is used for any other custom message in the communication
	OtherMessage = "OTHER"

	// MessageTimeout is the timeout for the TCP communication
	MessageTimeout = "TIMEOUT"

	// MessageSolveProblem is the message sent when the server wants client to solve a problem
	MessageSolveProblem = "SOLVE"

	// MessageSolvedProblem is the message sent when the client has solved the problem
	MessageSolvedProblem = "SOLVED"
)
