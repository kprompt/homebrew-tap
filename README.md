# homebrew-tap

Homebrew formulas for [kprompt](https://kprompt.ai).

## Install

```bash
brew install kprompt/tap/kprompt
```

Or:

```bash
brew tap kprompt/tap
brew install kprompt
```

## Releases

Formulas install official binaries from [kprompt/kprompt releases](https://github.com/kprompt/kprompt/releases).
GoReleaser on the CLI repo bumps `Formula/kprompt.rb` on each tagged release when `HOMEBREW_TAP_GITHUB_TOKEN` is set on [`kprompt/kprompt`](https://github.com/kprompt/kprompt).

### Maintainer: wire the tap token

1. Create a fine-grained PAT (or classic `repo`) with **Contents: Read and write** on `kprompt/homebrew-tap` only.
2. Add it as a repository secret on the CLI repo:

```bash
cd kprompt
gh secret set HOMEBREW_TAP_GITHUB_TOKEN --repo kprompt/kprompt
# paste the PAT when prompted
```

If the secret is missing, GoReleaser still publishes GitHub Release assets; the Formula must be bumped by hand (as for v0.11.0).

