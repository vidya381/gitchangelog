# GitChangelog

Paste in a GitHub repo URL, get a formatted changelog back. That's it.

It reads the commit history through the GitHub API and sorts commits into sections — Features, Bug Fixes, Breaking Changes, etc. Works best with repos that use conventional commit messages like `feat:`, `fix:`, `docs:`. Commits that don't follow that format land in an "Other" section.

## Stack

- **Backend** — Go (Gin), deployed on Oracle Cloud via Docker
- **Frontend** — Next.js + Tailwind, deployed on Vercel

## Running locally

You need Go 1.23+ and Node.js 20+.

```bash
# Backend (runs on :8080)
cd backend && go run ./cmd/server

# Frontend (runs on :3000)
cd frontend && npm install && npm run dev
```

Create `frontend/.env.local`:
```
NEXT_PUBLIC_API_URL=http://localhost:8080
```

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md). Issues marked [`good first issue`](https://github.com/vidya381/gitchangelog/issues?q=label%3A%22good+first+issue%22) are a good place to start.
