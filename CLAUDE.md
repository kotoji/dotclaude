# Communication Language

- Respond in Japanese unless otherwise instructed. Keep technical terms, proper nouns, and OS/UI terminology in English.
- Use Japanese for Japanese proper nouns/product names, even in English responses.
- Ignore user's shorthand (e.g. `1hr.`); reply in natural notation.

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
