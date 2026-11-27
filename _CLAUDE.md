# Coding

## Code Comments

- Write what the code does. A one-line "what" is usually enough.
- Do not write a "why" that depends on temporary context (issue numbers, how a bug was found) or on the implementation elsewhere (e.g. "because the later `groupby` raises KeyError"). It goes stale when the code changes and becomes debt.
- Write a "why" only when it cannot be read from the code and will stay true for a long time (e.g. a domain rule).
- Do not add issue or PR references (e.g. "(Issue #123)") to comments or docstrings by default. `git blame` already links the code to its history, and a reference that only says where the change came from adds nothing for the reader.
  - Add one only when the reader needs to follow it to understand or change the code: a TODO or known limitation tracked in an open issue, a stopgap with a `NOTE:` (see below), a workaround for an upstream bug (link the upstream issue), or a decision whose reasoning is too long for a comment and is recorded there.
- Put the history and reasoning of a change in the PR description, commit message, or design docs instead.
- When the code is a stopgap that will likely need a proper fix later, leave a `NOTE:` with its history: what the code did before, what changed and in which issue or PR, and why only this much was done. This is an exception to the rules above on history and references, because the reader needs it to judge whether the code can be changed. A closed issue is fine to reference.
  - Do not write how it should be fixed, and do not make it a `TODO:`, unless the fix is decided and tracked. An undecided plan goes stale; the history does not.

## Tests

- Keep tests that check the behavior spec: what the code returns for a given input.
- Do not keep tests that only check the act of changing the code. A good check: would this test have been written if the code had been written this way from the start? If not, it is temporary. Examples: a test that the SQL now uses parameter binding after rewriting an f-string, or a test that a renamed function still exists.
  - Use such tests to verify the change while working, then do not commit them.
- Exception, depending on the situation: in a large refactoring, it can be worth committing temporary tests (e.g. tests that pin the old output) and deleting them later. In that case, say in the PR that they are temporary and when they will be removed.

# Regarding Responses via Discord Channel

When a message is received via the `<channel source="discord">` tag, read `TINA_PERSONA.md` and embody the persona defined there. The Discord user is talking directly to TINA.

**Important**: When a message requires work that takes time (coding, research, etc.), you MUST send a short in-character reply to the Discord channel BEFORE making any tool calls. Never begin work silently.

**Important**: ALWAYS reply to Discord messages through the Discord channel tool. Never generate a response as internal text only — if you have something to say, it must be sent via the tool so the user actually sees it.
