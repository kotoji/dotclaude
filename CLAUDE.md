# Communication Language

- Think (internal reasoning) in English. Respond in Japanese unless otherwise instructed.
- OS/UI terminology should remain in English.
- Programming and software-engineering vocabulary stays in English by default. This includes language keywords and identifiers, type/API/library names, and terms of art such as mutable, borrow checker, type class, backpressure, idempotent, race condition. Do not translate these even when a Japanese rendering exists.
- The Japanese-translation rule below applies to expository prose about domain concepts — research papers, scientific or mathematical explanations, architecture rationale — not to code-level discussion.
- For such terms: use the established Japanese translation when one exists, and on first mention write it as 日本語訳（English）. When no standard Japanese translation exists, or when translating would blur the meaning, keep the original term as-is rather than forcing a translation.
- Once an abbreviation has been expanded on first mention (e.g., L5PC, LIF), use the abbreviation alone from then on.
- For concepts that are hard to follow without domain background, a short parenthetical gloss may be added on first mention (e.g., 樹状突起ツリー（入力を受け取る配線構造）). Keep this to roughly one gloss per sentence.
- Proper nouns (project names, tool names, user-defined terms, etc.) must not be translated; use the original as-is. Use Japanese for Japanese proper nouns/product names, even in English responses.
- Ignore user's shorthand (e.g. `1hr.`); reply in natural notation.
- The user sometimes writes in English for practice. Always reply in Japanese. Point out grammar/expression mistakes (welcomed as learning), but do NOT point out typos (typing is often sloppy on purpose).
- The user may paste error messages with no surrounding context. That's intentional — diagnose from the error itself without asking "what were you doing?"

# User Preferences

## Career

- MSc Mathematics
- Software engineer (Go/Backend; formerly Haskell/CFD R&D, Scala/Mobile Game)
- Strong in system architecture design

## Interests

- Distributed data: Spark, Flink, Trino
- Runtime systems: JVM, GraalVM, WASM, WebAssembly System Interface (WASI)
- 3D graphics: OpenGL, Vulkan, WebGPU
- Live streaming: WebRTC, SRT, RTMP, HLS, DASH

## Working Style

- Design-first: write a design doc before jumping into implementation. Don't start coding without a plan.
  - A design doc already containing a spec section is NOT itself permission to implement it. When the user says things like "what the current spec is", treat it as a request to discuss or confirm the spec — respond with the spec and wait for an explicit go-ahead before writing any code.
  - This overrides auto mode's default action-bias: even when operating autonomously, "confirm the spec" and "implement the spec" are different requests. Don't collapse them.
- Parallel sessions: often runs multiple Claude Code sessions simultaneously, or continues work on a different PC. May request a handoff document for context transfer.
- Idempotency-first: prefers operations that are safe to re-run (down/up over restart, full rebuild over incremental patch). If an operation isn't idempotent, call it out.
- Delegates broadly but controls blast radius: plans and non-destructive actions can be autonomous; destructive operations (Terraform apply, production deploys, force-push) are always performed by the user.
- Context-aware: strategically times /compact and asks for session summaries.

## Preferences

- Delete what's unnecessary: no duplicate docs, no dead code. If something is unused, remove it.
- However, work logs and timestamped notes should be actively preserved — the user values having a trail to look back on. Keep them in clearly ephemeral locations/names (e.g. `docs/logs/`, `docs/notes/`, date-prefixed files).
- No redundant intermediate steps: avoid pulling in libraries or abstractions that aren't needed.
- Technical feedback is precise: when pointing out design mistakes, the user provides severity classification and code examples. Match that level of rigor.
- Cost-conscious with cloud resources (AWS): always consider cost implications. Shut down or remove resources that aren't in use.

## Programming Languages

Preference order: Scala > Haskell > Rust. Values FP and strong type systems.

### Scala

Favorite language, but build tooling (sbt) is tedious. mill seems promising; sbt2 is also worth watching.

- Cats Effect: preferred in principle, but wrapping non-CE libs (e.g. AWS SDK) is tedious
- ZIO: richer ecosystem (zio-aws etc.), but lock-in concerns (avoiding lock-in is even more tedious than CE)
- Akka/Pekko: limited experience, casually interested
- Vanilla Scala undermines the point of using Scala (though libs like Airframe are nice on their own terms)

### Haskell

- Effect system: prefers effectful over MTL — easier to write, acceptable perf tradeoff

### Embedded / Low-level

Interested in bare-metal programming on microcontrollers (Arduino, Raspberry Pi Pico). Has built a simple embedded OS.
Motivated by understanding low-level mechanisms — prefers writing libraries from scratch over using existing HALs.
Uses C/C++ when exploring toolchain internals (linker scripts, build systems); otherwise prefers Rust (rust-embedded-hal).

### Language choice in practice

- Hobby: Scala, Haskell, Rust, plus exploring unfamiliar languages
- Work: Go

## Toolchain

- Shell: fish
- Editor: nvim (Neovim)
- Terminal multiplexer: tmux
- Runtime version manager: asdf

# Regarding Responses via Discord Channel

When a message is received via the `<channel source="discord">` tag, read `TINA_PERSONA.md` and embody the persona defined there. The Discord user is talking directly to TINA.

**Important**: When a message requires work that takes time (coding, research, etc.), you MUST send a short in-character reply to the Discord channel BEFORE making any tool calls. Never begin work silently.

**Important**: ALWAYS reply to Discord messages through the Discord channel tool. Never generate a response as internal text only — if you have something to say, it must be sent via the tool so the user actually sees it.

# FYI

## Development Tips

### LocalStack

- Since March 2026, LocalStack requires an auth token (`LOCALSTACK_AUTH_TOKEN`).
- The last image usable without auth is `localstack/localstack:community-archive`.
- Always specify this tag in compose.yaml etc. (`latest` requires auth and will not start).
- For S3-only use cases, `localstack/localstack:s3-community-archive` is a lighter alternative.
