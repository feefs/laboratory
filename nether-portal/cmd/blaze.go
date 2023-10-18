package cmd

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/spf13/cobra"
)

var blazeCmd = &cobra.Command{
	Use:           "blaze",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE:          blazeRun,
}

func blazeRun(cmd *cobra.Command, args []string) error {
	r, err := git.PlainOpenWithOptions(".", &git.PlainOpenOptions{DetectDotGit: true})
	if err != nil {
		return err
	}
	head, err := r.Head()
	if err != nil {
		return err
	}

	until := time.Now()
	since := until.Add(-1 * (7 * 24 * time.Hour)) // TODO: Support cli flag to set time range
	cIter, err := r.Log(&git.LogOptions{From: head.Hash(), Since: &since, Until: &until})
	if err != nil {
		return err
	}

	var commits []*object.Commit
	_ = cIter.ForEach(func(c *object.Commit) error {
		commits = append(commits, c)
		return nil
	})
	if len(commits) == 0 {
		fmt.Println(color.YellowString("No commits found!"))
		return nil
	}
	latestCommit := commits[0]
	earliestCommit := commits[len(commits)-1]
	var baseCommit *object.Commit
	if earliestCommit.NumParents() > 0 {
		parent, err := earliestCommit.Parent(0)
		if err != nil {
			return err
		}
		baseCommit = parent
	} else {
		baseCommit = earliestCommit
	}

	markdownChanges, err := getMarkdownChanges(baseCommit, latestCommit)
	if err != nil {
		return err
	}
	if len(markdownChanges) == 0 {
		fmt.Println(color.YellowString("No markdown files found!"))
		return nil
	}

	patch, err := markdownChanges.Patch()
	if err != nil {
		return err
	}

	fmt.Println(patch)

	return nil
}

func getMarkdownChanges(baseCommit, latestCommit *object.Commit) (object.Changes, error) {
	baseTree, err := baseCommit.Tree()
	if err != nil {
		return nil, err
	}
	latestTree, err := latestCommit.Tree()
	if err != nil {
		return nil, err
	}

	changes, err := baseTree.Diff(latestTree)
	if err != nil {
		return nil, err
	}

	var markdownChanges object.Changes
	for _, change := range changes {
		extension := filepath.Ext(change.To.Name)
		if extension == ".md" {
			markdownChanges = append(markdownChanges, change)
		}
	}

	return markdownChanges, nil
}
