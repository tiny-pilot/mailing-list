# mailing-list

[![CircleCI](https://circleci.com/gh/tiny-pilot/mailing-list.svg?style=svg&circle-token=88c772285740e28012a545a2537d777601cdcfd0)](https://circleci.com/gh/tiny-pilot/mailing-list)

## Testing

```bash
. .dev.env

TESTER="$(mktemp -d)"
go build -o /tmp/sheets-cli ./dev-scripts/cli.go
```
