#!/bin/bash

# ./new-adr.sh "adr title"


TITLE=$(echo "$@" | tr ' ' '-' | tr '[A-Z]' '[a-z]')
NEXT_NUM=$(printf "%04d" $(ls docs/adr/*.md | wc -l | awk '{print $1 + 1}'))

cat <<EOT > docs/adr/${NEXT_NUM}_${TITLE}.md
# ${NEXT_NUM} - $@

Date: $(date +%Y-%m-%d)

## Status

Proposed

## Context

Why was this decision made?

## Decision

What was decided?

## Consequences

What are the positive/negative consequences?
EOT

echo "ADR created: docs/adr/${NEXT_NUM}_${TITLE}.md"
