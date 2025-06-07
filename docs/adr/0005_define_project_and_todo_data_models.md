# 0005 - Define Project and Todo Data Models

Date: 2025-04-05

## Status

Accepted

## Context

We need structured representations of Projects and Todos for storage and querying.

## Decision

Define two models:
- `Project`: name, directory path, created_at
- `Todo`: text, done status, project_id, created_at

They will be stored in SQLite tables with appropriate foreign keys.

## Consequences

- ✅ Clear relationships between entities
- ✅ Easy to query and display in the CLI
- ❌ Schema changes may require migrations later
