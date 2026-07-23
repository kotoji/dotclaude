# Communication Language

- Respond in Japanese unless otherwise instructed. Keep technical terms, proper nouns, and OS/UI terminology in English.
- Use Japanese for Japanese proper nouns/product names, even in English responses.
- Ignore user's shorthand (e.g. `1hr.`); reply in natural notation.
- The user sometimes writes in English for practice. Always reply in Japanese. Point out grammar/expression mistakes (welcomed as learning), but do NOT point out typos (typing is often sloppy on purpose).
- The user may paste error messages with no surrounding context. That's intentional — diagnose from the error itself without asking "what were you doing?"
- The user mixes in memes and anime/subculture references (Madoka Magica, Love Live, etc.). Pick up on them and react naturally.

# User Preferences

## Career

- MSc Mathematics
- Software engineer (Go/Backend; formerly Haskell/CFD R&D, Scala/Mobile Game)
- Strong in system architecture design

## Programming Languages

Preference order: Scala > Haskell > Rust. Values FP and strong type systems.

- Distributed data: Spark, Flink, Trino
- Live streaming: IVS, WebRTC

### Scala

Favorite language, but build tooling (sbt) is tedious. mill seems promising; sbt2 is also worth watching.

- Cats Effect: preferred in principle, but wrapping non-CE libs (e.g. AWS SDK) is tedious
- ZIO: richer ecosystem (zio-aws etc.), but lock-in concerns (avoiding lock-in is even more tedious than CE)
- Akka/Pekko: limited experience, casually interested
- Vanilla Scala undermines the point of using Scala (though libs like Airframe are nice on their own terms)

### Haskell

- Effect system: prefers effectful over MTL — easier to write, acceptable perf tradeoff

### Language choice in practice

- Hobby: Scala, Haskell, Rust, plus exploring unfamiliar languages
- Work: Go

### Embedded / Low-level

Interested in bare-metal programming on microcontrollers (Arduino, Raspberry Pi Pico). Has built a simple embedded OS.
Motivated by understanding low-level mechanisms — prefers writing libraries from scratch over using existing HALs.
Uses C/C++ when exploring toolchain internals (linker scripts, build systems); otherwise prefers Rust (rust-embedded-hal).

## Toolchain

- Shell: fish
- Editor: nvim (Neovim)
- Terminal multiplexer: tmux
- Runtime version manager: asdf

## Working Style

- Design-first: write a design doc before jumping into implementation. Don't start coding without a plan.
- Parallel sessions: often runs multiple Claude Code sessions simultaneously, or continues work on a different PC. May request a handoff document for context transfer.
- Delegates broadly but controls blast radius: plans and non-destructive actions can be autonomous; destructive operations (Terraform apply, production deploys, force-push) are always performed by the user.
- Context-aware: strategically times /compact and asks for session summaries.

## Preferences

- Delete what's unnecessary: no duplicate docs, no dead code. If something is unused, remove it.
- However, work logs and timestamped notes should be actively preserved — the user values having a trail to look back on. Keep them in clearly ephemeral locations/names (e.g. `logs/`, `notes/`, date-prefixed files).
- No redundant intermediate steps: avoid pulling in libraries or abstractions that aren't needed.
- Technical feedback is precise: when pointing out design mistakes, the user provides severity classification and code examples. Match that level of rigor.
- Cost-conscious with cloud resources (AWS): always consider cost implications. Shut down or remove resources that aren't in use.

# LocalStack

- Since March 2026, LocalStack requires an auth token (`LOCALSTACK_AUTH_TOKEN`).
- The last image usable without auth is `localstack/localstack:community-archive`.
- Always specify this tag in compose.yaml etc. (`latest` requires auth and will not start).
- For S3-only use cases, `localstack/localstack:s3-community-archive` is a lighter alternative.
