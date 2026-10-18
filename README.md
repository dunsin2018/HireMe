<div align="center">

<picture>
  <source media="(prefers-color-scheme: dark)" srcset="docs/assets/logo-dark.svg">
  <img src="docs/assets/logo-light.svg" alt="HireMe" width="84" height="84">
</picture>

# HireMe

### Every IT job, straight from the source.

Postings pulled directly from company career pages — no recruiters, no reposts, no dead links. Fully open source.

[Features](docs/features.md) · [API](docs/API.md) · [Contributing](CONTRIBUTING.md)

[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
![Go version](https://img.shields.io/github/go-mod/go-version/dunsin2018/HireMe)

</div>

## Why HireMe?

- **Straight from the source.** Every listing is crawled directly from a company's
  own ATS — Workday, Greenhouse, Lever, Ashby, iCIMS and a long tail of others — and
  links to the original posting. No recruiter reposts, no aggregator middlemen, no
  dead links.
- **One schema, deduplicated.** The same role posted to three boards collapses into
  one entry: every posting is normalized into a single shape and deduplicated on a
  stable key.
- **Search that understands jobs.** Faceted full-text search over region, work mode,
  seniority, skills and salary — derived from curated dictionaries, never guessed.
- **Actually open.** MIT-licensed and self-hostable, pipeline and data both in the
  open. Adding a company is one line of YAML.
- **Yours to build on.** A clean HTTP API, saved searches with digests, and a whole
  workspace on top of the catalogue — see [Beyond the catalogue](#beyond-the-catalogue).

## Beyond the catalogue

Finding the posting is half the problem. The other half — writing the CV,
sending it, and knowing what happened to it — is built on the same data, in the
same repository, under the same licence.

| | |
| --- | --- |
| **Find** | Faceted search, curated collections, saved searches with email digests, market analytics, the ghost-job signal |
| **Apply** | CV builder with ATS-safe PDF templates, deterministic CV↔vacancy scoring, CV tailoring that invents nothing, referrals |
| **Track** | An application board with stages, a mail inbox that links recruiter replies to the application they answer, an append-only event ledger, reminders |

**[Full feature reference → `docs/features.md`](docs/features.md)**

## Run it

```bash
git clone https://github.com/dunsin2018/HireMe
cd HireMe
cp .env.example .env   # fill in DATABASE_URL and JWT_SECRET
make up                # app + Postgres via Docker Compose
```

The web app is served same-origin by the Go server. See [`.env.example`](.env.example)
for every knob, and [`docs/`](docs/) for architecture, API, and operations.

## Stack

**Go** (Fiber) · **PostgreSQL** · **Meilisearch** · **SvelteKit** · **Docker Compose**

## Origin

HireMe is a fork of [freehire](https://github.com/strelov1/freehire) by Ilya Strelov,
rebranded and independently operated. It is MIT-licensed; the original copyright
notice is retained in [LICENSE](LICENSE).

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md). Issues and pull requests are welcome —
an opened issue counts as much as merged code.
