# blueplanet-screencapture

Executes screen captures of a Blue Planet UI

Usage:

```bash
go install github.com/jgroom33/blueplanet-screencapture
~/go/bin/blueplanet-screencapture -path=https://10.182.18.132/orchestrate/#/list/resource-types -element=.main-body -file=foo.png
```

## Options

-element - https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelector

-file - filename to save
