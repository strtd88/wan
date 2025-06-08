# 0007 - Use Sqitch for Database Migrations

Date: 2025-04-06

## Status

Accepted

## Context

We are building a local-first CLI tool that uses SQLite to store structured data about projects, todos, and brag entries. As the schema evolves, we need a lightweight, version-controlled way to manage migrations.

Since this project is personal but built with professional standards, we want a solution consistent with real-world practices — ideally one already used or familiar from work.

## Decision

Use [Sqitch](https://sqitch.org/)  as the database migration tool for managing SQLite schema changes.

Sqitch provides:
- Plain SQL-based migrations
- Git-friendly workflow
- Engine support for SQLite
- Deploy, revert, and verify capabilities
- No runtime dependencies

It aligns well with our goals of being local-first, simple, and maintainable.

## Consequences

### ✅ Positive
