package cmd

import (
	"github.com/Originate/git-town/src/tools/gittools"
	"github.com/spf13/cobra"
)

var setParentBranchCommand = &cobra.Command{
	Use:   "set-parent-branch <child_branch> <parent_branch>",
	Short: "Updates a branch's parent",
	Long: `Updates a branch's parent

Updates the parent branch of a feature branch in Git Town's configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		setParentBranch(args[0], args[1])
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return errortools.FirstError(
			validateMaxArgsFunc(args, 2),
			gittools.ValidateIsRepository,
		)
	},
}

func setParentBranch(childBranch, parentBranch string) {
	workflows.EnsureHasBranch(childBranch)
	workflows.EnsureHasBranch(parentBranch)
	gittools.SetParentBranch(childBranch, parentBranch)
	gittools.DeleteAllAncestorBranches()
}

func init() {
	RootCmd.AddCommand(setParentBranchCommand)
}
