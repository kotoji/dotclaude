---
name: dewaste
description: Remove change-history noise and waste from the current diff before committing or opening a PR — history-dependent comments, unneeded issue/PR references, tests that only check that a change happened, dead code. Use when asked to clean up, tidy, or de-noise a diff.
---

Review the current diff (staged and unstaged; if empty, the branch against its base) and remove waste. The principle: code, comments, and tests describe what the code is now, not how it got here. History belongs in commits, PRs, and design docs.

Only touch lines the diff adds or changes, unless asked to go wider.

## What to remove or rewrite

### Comments

- A "why" that depends on temporary context: an issue number, how a bug was found, "fixed X", "now uses Y", "previously Z".
- A "why" that depends on implementation elsewhere, e.g. "because the later `groupby` raises KeyError". It goes stale when that code changes.
- A comment that restates the code line by line. Keep a one-line "what" when it helps.
- Issue or PR references like "(Issue #123)" that only say where the change came from. `git blame` already links code to history.

Keep:

- A "why" that cannot be read from the code and will stay true for a long time (e.g. a domain rule).
- A reference the reader must follow to understand or change the code: a TODO or known limitation tracked in an open issue, a workaround for an upstream bug (link the upstream issue), a decision whose reasoning is too long for a comment.

### Stopgaps

If the code is a stopgap that will likely need a proper fix, make sure it has a `NOTE:` with its history: what the code did before, what changed and in which issue or PR (a closed one is fine), and why only this much was done. Remove any undecided fix plan from it, and do not turn it into a `TODO:` unless the fix is decided and tracked.

### Tests

- Remove tests that only check the act of changing the code. Ask: would this test have been written if the code had been written this way from the start? If not, it is temporary. Examples: a test that SQL now uses parameter binding after rewriting an f-string; a test that a renamed function still exists; a test asserting an old bug's exact symptom is gone rather than the correct behavior.
- If such a test covers behavior that matters, rewrite it as a behavior test (input → expected output) instead of deleting it.
- Exception: in a large refactoring, temporary tests that pin old output may be worth committing. Keep them only if the PR says they are temporary and when they will be removed; otherwise ask the user.

### Code

- Dead code the change left behind: unused functions, imports, variables, parameters, branches.
- Leftover debug output and commented-out code.
- Compatibility shims, aliases, or re-exports for something this same change renamed or removed, when nothing outside the change uses them.

## How to work

1. List each finding with `file:line`, what it is, and the proposed change.
2. Apply the clear ones. For judgment calls (is this "why" durable? is this test temporary?), ask before changing.
3. If tests exist, run the related ones after editing and report the result.
4. Do not commit. Suggest that any removed history go into the commit message or PR description instead.
