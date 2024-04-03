package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/charmbracelet/log"
	"go.abhg.dev/gs/internal/git"
	"go.abhg.dev/gs/internal/state"
)

type branchBottomCmd struct{}

func (*branchBottomCmd) Run(ctx context.Context, log *log.Logger) error {
	repo, err := git.Open(ctx, ".", git.OpenOptions{
		Log: log,
	})
	if err != nil {
		return fmt.Errorf("open repository: %w", err)
	}

	// TODO: prompt for init if not initialized
	store, err := state.OpenStore(ctx, repo)
	if err != nil {
		return fmt.Errorf("open storage: %w", err)
	}

	// TODO: ensure no uncommitted changes

	currentBranch, err := repo.CurrentBranch(ctx)
	if err != nil {
		return fmt.Errorf("get current branch: %w", err)
	}

	if currentBranch == store.Trunk() {
		return errors.New("already on trunk")
	}

	var root string
	for {
		b, err := store.GetBranch(ctx, currentBranch)
		if err != nil {
			return fmt.Errorf("get branch %q: %w", currentBranch, err)
		}

		if b.Base == store.Trunk() {
			root = currentBranch
			break
		}

		currentBranch = b.Base
	}

	if err := repo.Checkout(ctx, root); err != nil {
		return fmt.Errorf("checkout %q: %w", root, err)
	}

	return nil
}
