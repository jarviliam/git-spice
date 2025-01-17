---
icon: material/frequently-asked-questions
description: >-
  Frequently asked questions about git-spice,
  a tool to manage stacked branches in Git.
---

# Frequently Asked Questions

## What is stacking?

Stacking refers to the practice of creating interdependent branches
on top of each other.
Each branch in the stack builds on top of the previous one.
For example, you might have a branch `feature-a` that adds a new feature,
and a branch `feature-b` that builds on top of `feature-a`.

Stacking your changes lets you:

- **Unblock yourself**:
  While one branch is under review, you can start working on the next one.
- **Be kinder to your team**:
  By keeping changes small and focused, you make it easier for your team
  to review, test, and understand your work.
  A 100 line PR gets a more meaningful review than a 1000 line PR.

git-spice helps you manage your stack of branches,
keeping them up-to-date and in sync with each other.

**Related**:

- [The stacking workflow](https://www.stacking.dev/)

## Where is the GitHub authentication token stored?

git-spice stores the GitHub authentication in a system-specific secure storage.
See [Secret storage](guide/internals.md#secret-storage) for more details.

## Why doesn't git-spice create one PR per commit?

With tooling like this, there are two options:
each commit is an atomic unit of work, or each branch is.
While the former might be more in line with Git's original philosophy,
the latter is more practical for most teams with GitHub-based workflows.

With a PR per commit, when a PR gets review feedback,
you must amend that commit with fixes and force-push.
This is inconvenient for PR reviewers as there's no distinction
between the original changes and those addressing feedback.

However, with a PR per branch, you can keep the original changes separate
from follow-up fixes, even if the branch is force-pushed.
This makes it easier for PR reviewers to work through the changes.

And with GitHub squash-merges, you can still get a clean history
consisting of atomic, revertible commits on the trunk branch.
