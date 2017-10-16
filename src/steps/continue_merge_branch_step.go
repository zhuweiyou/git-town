package steps

import (
	"github.com/Originate/git-town/src/flows/scriptflows"
	"github.com/Originate/git-town/src/tools/gittools"
)

// ContinueMergeBranchStep finishes an ongoing merge conflict
// assuming all conflicts have been resolved by the user.
type ContinueMergeBranchStep struct {
	NoOpStep
}

// CreateAbortStep returns the abort step for this step.
func (step *ContinueMergeBranchStep) CreateAbortStep() Step {
	return &NoOpStep{}
}

// CreateContinueStep returns the continue step for this step.
func (step *ContinueMergeBranchStep) CreateContinueStep() Step {
	return step
}

// Run executes this step.
func (step *ContinueMergeBranchStep) Run() error {
	if gittools.IsMergeInProgress() {
		return scriptflows.RunCommand("git", "commit", "--no-edit")
	}
	return nil
}
