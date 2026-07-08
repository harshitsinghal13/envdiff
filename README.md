# envdiff

> Catch `.env` config drift before it breaks your team.

`envdiff` is a fast, zero-dependency CLI tool that compares `.env` files and catches missing keys, extra keys, and changed values — before they cause confusing bugs in production.

---

## The problem

Every team has hit this:

- New developer joins → copies `.env.example` → misses three keys → spends half a day debugging
- Someone adds `STRIPE_WEBHOOK_SECRET` to `.env.example` → forgets to tell the team → production deploy fails at 2am
- Staging and local environments drift silently over months → nobody knows why staging works but local doesn't

`envdiff` fixes this in one command.

---

## Install

**If you have Go installed:**
```bash
go install github.com/harshitsinghal13/envdiff@latest
```

**Download binary (no Go needed):**

Go to [Releases](https://github.com/harshitsinghal13/envdiff/releases/latest) and download the binary for your OS:

| OS | File |
|----|------|
| Windows | `envdiff_windows_amd64.zip` |
| Mac (M1/M2/M3) | `envdiff_darwin_arm64.tar.gz` |
| Mac (Intel) | `envdiff_darwin_amd64.tar.gz` |
| Linux | `envdiff_linux_amd64.tar.gz` |

**Mac/Linux — make it global:**
```bash
chmod +x envdiff
sudo mv envdiff /usr/local/bin/
```

**Windows — make it global:**

Move `envdiff.exe` to a folder like `C:\tools` and add it to your PATH.

---

## Usage

### Compare two `.env` files

```bash
envdiff compare .env.example .env
```

**Output:**
```
Comparing .env.example → .env

  - STRIPE_KEY        (in .env.example, missing in .env)
  + OLD_FLAG          (in .env, not in .env.example)
  ~ API_TIMEOUT       (different value)
```

### No differences found

```bash
envdiff compare .env.example .env
✓ No differences found
```

### Help

```bash
envdiff --help
envdiff compare --help
```

---

## What each symbol means

| Symbol | Meaning |
|--------|---------|
| `-` | Key exists in first file but **missing** in second |
| `+` | Key exists in second file but **not in first** (possibly stale) |
| `~` | Key exists in both files but has a **different value** |

---

## Roadmap

- [x] **v0.1.0** — Compare two `.env` files (missing, extra, changed keys)
- [ ] **v0.2.0** — Lint a single `.env` file (duplicates, empty values, malformed lines)
- [ ] **v0.3.0** — Validate `.env` against `.env.example` with strict rules
- [ ] **v0.4.0** — GitHub Action for CI/CD integration
- [ ] **v0.5.0** — Support Docker Compose, Kubernetes ConfigMaps
- [ ] **v0.6.0** — Interactive TUI mode

---

## Why envdiff?

There are tools for secret management (Vault, AWS Secrets Manager) and tools for syncing secrets (Doppler, dotenv-vault). But there's no simple, fast, offline CLI that just **tells you what's different** between two config files without sending your secrets anywhere.

`envdiff` is:
- **Offline** — never sends your data anywhere
- **Fast** — instant results, single binary
- **Zero config** — no setup, no account, just run it
- **Open source** — MIT license, free forever

---

## Contributing

Contributions are welcome. Please open an issue first to discuss what you'd like to change.

```bash
git clone https://github.com/harshitsinghal13/envdiff.git
cd envdiff
go run main.go compare .env.example .env
```

---

## License

MIT © [Harshit Singhal](https://github.com/harshitsinghal13)
