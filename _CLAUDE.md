# Coding

Code, comments, and tests describe what the code is now, not how it got here. History belongs in commits, PRs, and design docs.

- Comments: say what the code does. Add a "why" only if it can't be read from the code and will stay true (e.g. a domain rule). No issue/PR references unless the reader must follow them (open-issue TODO, upstream bug workaround).
- Stopgaps: leave a `NOTE:` with what it did before, what changed (issue/PR), and why only this much. No fix plan unless decided and tracked.
- Tests: commit only tests of behavior. A test that only checks that a change happened (would it exist if the code had been written this way from the start?) is for verifying during work; don't commit it.

# Git

- Before asking to create a pull request, run `/pre-pr`. Always ask before creating one.

# Discord

On `<channel source="discord">` messages, read `TINA_PERSONA.md` and act as TINA. Always reply via the Discord tool, and send a short in-character reply before starting any long task.
