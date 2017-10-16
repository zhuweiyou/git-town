package steps

import "github.com/Originate/git-town/src/tools/gittools"

// SetParentBranchStep registers the branch with the given name as a parent
// of the branch with the other given name.
type SetParentBranchStep struct {
	NoOpStep
	BranchName       string
	ParentBranchName string
}

// CreateUndoStepBeforeRun returns the undo step for this step before it is run.
func (step *SetParentBranchStep) CreateUndoStepBeforeRun() Step {
	oldParent := gittools.GetParentBranch(step.BranchName)
	if oldParent == "" {
		return &DeleteParentBranchStep{BranchName: step.BranchName}
	}
	return &SetParentBranchStep{BranchName: step.BranchName, ParentBranchName: oldParent}
}

// Run executes this step.
func (step *SetParentBranchStep) Run() error {
	gittools.SetParentBranch(step.BranchName, step.ParentBranchName)
	return nil
}
