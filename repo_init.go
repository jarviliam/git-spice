package main

import (
	"context"
	"fmt"

	"github.com/charmbracelet/log"
	"go.abhg.dev/gs/internal/git"
	"go.abhg.dev/gs/internal/state"
)

type repoInitCmd struct {
	Trunk string `help:"The name of the trunk branch"`
}

func (cmd *repoInitCmd) Run(ctx context.Context, log *log.Logger) error {
	repo, err := git.Open(ctx, ".", git.OpenOptions{
		Log: log,
	})
	if err != nil {
		return fmt.Errorf("open repository: %w", err)
	}

	if cmd.Trunk == "" {
		// TODO: check if there's a remote first?
		if b, err := repo.DefaultBranch(ctx, "origin"); err == nil {
			cmd.Trunk = b
		}
	}

	if cmd.Trunk == "" {
		// Use the current branch as the trunk.
		b, err := repo.CurrentBranch(ctx)
		if err != nil {
			return fmt.Errorf("get current branch: %w", err)
		}
		cmd.Trunk = b
	}

	_, err = state.InitStore(ctx, state.InitStoreRequest{
		Repository: repo,
		Trunk:      cmd.Trunk,
	})
	if err != nil {
		return fmt.Errorf("open storage: %w", err)
	}

	log.Infof("Initialized with trunk: %s", cmd.Trunk)
	return nil
}
