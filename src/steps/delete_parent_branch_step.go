package steps

import "github.com/Originate/git-town/src/tools/gittools"

// DeleteParentBranchStep removes the parent branch entry in the Git Town configuration.
type DeleteParentBranchStep struct {
	NoOpStep
	BranchName string
}

// CreateUndoStepBeforeRun returns the undo step for this step before it is run.
func (step *DeleteParentBranchStep) CreateUndoStepBeforeRun() Step {
	parent := gittools.GetParentBranch(step.BranchName)
	if parent == "" {
		return &NoOpStep{}
	}
	return &SetParentBranchStep{BranchName: step.BranchName, ParentBranchName: parent}
}

// Run executes this step.
func (step *DeleteParentBranchStep) Run() error {
	gittools.DeleteParentBranch(step.BranchName)
	return nil
}
