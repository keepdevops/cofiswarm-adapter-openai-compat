# cofiswarm-adapter-openai-compat

Cofiswarm component: `adapter-openai-compat`.

- Layout: [REPO-STANDARD-LAYOUT](https://github.com/keepdevops/cofiswarm-docs/blob/main/REPO-STANDARD-LAYOUT.md)
- Migration: [MIGRATION-SPRINTS](https://github.com/keepdevops/cofiswarm-docs/blob/main/MIGRATION-SPRINTS.md)

## FHS paths

| Path | Purpose |
|------|---------|
| `/etc/cofiswarm/adapter-openai-compat/` | config |
| `/var/lib/cofiswarm/adapter-openai-compat/` | state |
| `/var/log/cofiswarm/adapter-openai-compat/` | logs |

## Test

```bash
./test/scripts/assert-layout.sh adapter-openai-compat
```
