# 0006 - Store Brag File in SQLite

Date: 2025-04-05

## Status

Accepted

## Context

The "brag file" is a running log of accomplishments or achievements. Initially considered as plain text, but storing it in SQLite offers better integration with the rest of the app and allows for richer querying and management.

## Decision

Store brag entries in a SQLite table called `brags`. Each entry includes:
- ID
- Title (optional)
- Content (text)
- Created timestamp
- Updated timestamp

This allows users to manage brag entries via CLI commands like `brag add`, `brag list`, `brag edit`.

## Consequences

- ✅ Centralized storage with other project/todo data
- ✅ Easier to build search/export/edit functionality
- ✅ Versioned along with the rest of the app
- ❌ Slightly more complex than a flat file
- ✅ Enables tagging or categorization later
