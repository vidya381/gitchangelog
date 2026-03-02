# Contributing

GitChangelog is a small project. Contributions go directly into something people use.

## Pick an issue

Leave a comment on the issue you want to work on before starting, so we don't end up with two people doing the same thing. Issues labeled `good first issue` are scoped small and have clear instructions.

## Setup

You need Go 1.23+ and Node.js 20+.

```bash
# Backend
cd backend && go run ./cmd/server

# Frontend
cd frontend && npm install && npm run dev
```

Add `frontend/.env.local`:
```
NEXT_PUBLIC_API_URL=http://localhost:8080
```

## Making changes

Branch off `main`:

```bash
git checkout main
git checkout -b feature/issue-3-commit-parser
```

Use `fix/` for bug fixes, `docs/` for documentation changes.

Follow [conventional commits](https://www.conventionalcommits.org/) for your commit messages — `feat:`, `fix:`, `docs:`, `chore:`, etc.

## Opening a PR

- Open your PR against `main`
- Reference the issue it closes: `Closes #3`
- Fill in the PR template
- Wait for CI to pass before requesting a review

## Before you commit

```bash
# Backend
go fmt ./... && go vet ./...

# Frontend
npm run lint
```

## Scope
Check the open issues before proposing large changes. 
Big features need a discussion issue before a PR.
