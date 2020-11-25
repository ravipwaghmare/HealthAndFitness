package interfaces

// Goals interface
type Goals interface {
	SetWeightGoal(id int, weight int) error
	SetStepGoals(id int, steps int) error
}
