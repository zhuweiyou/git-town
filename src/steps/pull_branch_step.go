package steps

import "github.com/Originate/git-town/src/flows/scriptflows"

// PullBranchStep pulls the branch with the given name from the origin remote
type PullBranchStep struct {
	NoOpStep
	BranchName string
}

// Run executes this step.
func (step *PullBranchStep) Run() error {
	return scriptflows.RunCommand("git", "pull")
}
