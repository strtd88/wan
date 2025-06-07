# 0003 - Serve Static HTML Using Go’s `net/http`

Date: 2025-04-05

## Status

Accepted

## Context

We want to serve a simple HTML page of developer links locally without requiring a separate web server.

## Decision

Use Go’s built-in `net/http` package to implement a lightweight HTTP server that serves static HTML files from a known directory.

## Consequences

- ✅ Zero external dependencies
- ✅ Easy to embed within the same binary
- ❌ Lacks advanced features like compression or TLS (not needed for local use)
