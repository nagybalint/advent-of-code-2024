package tasks

type Task interface {
	CalculateAnswer(input string) (string, error)
}
