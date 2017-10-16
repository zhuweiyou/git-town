package cmd

import (
	"github.com/Originate/git-town/src/drivers"
	"github.com/Originate/git-town/src/flows/gitflows"
	"github.com/Originate/git-town/src/lib/gitlib"
	"github.com/Originate/git-town/src/steps"
	"github.com/spf13/cobra"
)

type newPullRequestConfig struct {
	InitialBranch  string
	BranchesToSync []string
}

var newPullRequestCommand = &cobra.Command{
	Use:   "new-pull-request",
	Short: "Creates a new pull request",
	Long: `Creates a new pull request

Syncs the current branch
and opens a browser window to the new pull request page of your repository.

The form is pre-populated for the current branch
so that the pull request only shows the changes made
against the immediate parent branch.

Supported only for repositories hosted on GitHub, GitLab, and Bitbucket.
When using hosted versions of GitHub, GitLab, or Bitbucket,
make sure that your SSH identity contains the phrase "github", "gitlab" or
"bitbucket", so that Git Town can derive which hosting service you use.

Example: your SSH identity should be something like
         "git@github-as-account1:Originate/git town.git"`,
	Run: func(cmd *cobra.Command, args []string) {
		steps.Run(steps.RunOptions{
			CanSkip:              func() bool { return false },
			Command:              "new-pull-request",
			IsAbort:              abortFlag,
			IsContinue:           continueFlag,
			IsSkip:               false,
			IsUndo:               false,
			SkipMessageGenerator: func() string { return "" },
			StepListGenerator: func() steps.StepList {
				config := getNewPullRequestConfig()
				return getNewPullRequestStepList(config)
			},
		})
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return errortools.FirstError(
			validateMaxArgsFunc(args, 0),
			gittools.ValidateIsRepository,
			validateIsConfigured,
			gittools.ValidateIsOnline,
			drivers.ValidateHasDriver,
		)
	},
}

func getNewPullRequestConfig() (result newPullRequestConfig) {
	if gittools.HasRemote("origin") {
		scriptlib.Fetch()
	}
	result.InitialBranch = gitlib.GetCurrentBranchName()
	gitflows.EnsureKnowsParentBranches([]string{result.InitialBranch})
	result.BranchesToSync = append(gittools.GetAncestorBranches(result.InitialBranch), result.InitialBranch)
	return
}

func getNewPullRequestStepList(config newPullRequestConfig) (result steps.StepList) {
	for _, branchName := range config.BranchesToSync {
		result.AppendList(steps.GetSyncBranchSteps(branchName))
	}
	result.Wrap(steps.WrapOptions{RunInGitRoot: true, StashOpenChanges: true})
	result.Append(&steps.CreatePullRequestStep{BranchName: config.InitialBranch})
	return
}

func init() {
	newPullRequestCommand.Flags().BoolVar(&abortFlag, "abort", false, abortFlagDescription)
	newPullRequestCommand.Flags().BoolVar(&continueFlag, "continue", false, continueFlagDescription)
	RootCmd.AddCommand(newPullRequestCommand)
}
