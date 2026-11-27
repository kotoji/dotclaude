---
name: pre-pr
description: Pre-PR checklist — run /dewaste, /simplify, and /code-review in order on the current branch, then report and ask whether to create the pull request. Use before creating a PR or when asked to prepare a branch for a PR.
---

Prepare the current branch for a pull request by running these skills in order. Run each with the Skill tool and finish it before starting the next, so each one sees the result of the previous.

1. `/dewaste` — remove change-history noise and waste.
2. `/simplify` — simplify the changed code.
3. `/code-review` — review the final code for bugs. Apply only fixes the user approves.

If a step finds nothing, say so and move on. If a step needs a decision from the user, ask and wait before continuing.

After all three steps:

1. If tests exist, run the related ones and report the result.
2. Summarize what each step changed or found, and anything left unresolved.
3. Ask whether to create the pull request. Do not create it, commit, or push without the user's answer.
