# 0002 - Choose Cobra as the CLI Framework

Date: 2025-04-05

## Status

Accepted

## Context

The application needs a robust command-line interface with support for nested commands, flags, and help documentation.

## Decision

Use [Cobra](https://github.com/spf13/cobra)  to structure the CLI. It offers excellent support for complex command trees and integrates well with Go idioms.

## Consequences

- ✅ Clear command hierarchy
- ✅ Built-in help and usage generation
- ❌ Larger dependency footprint than minimal alternatives
- ✅ Supports auto-generated shell completions
