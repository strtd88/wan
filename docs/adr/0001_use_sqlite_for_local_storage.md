# 0001 - Use SQLite for Local Data Storage

Date: 2025-04-05

## Status

Accepted

## Context

We need a lightweight, file-based database that requires no setup for local development use. The app must store data about projects, todos, and brag entries without relying on external services.

## Decision

Use SQLite as the primary local database engine. It is embedded, fast, and supports relational queries via SQL.

## Consequences

- ✅ No external dependencies
- ✅ Easy to version control or back up the database file
- ❌ Limited scalability beyond single-user/local usage
- ❌ Requires careful handling of schema migrations
