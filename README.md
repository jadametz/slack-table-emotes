# Slack Table Emotes

[![Build Status](https://travis-ci.org/jadametz/slack-table-emotes.svg?branch=master)](https://travis-ci.org/jadametz/slack-table-emotes)
[![Coverage Status](https://coveralls.io/repos/github/jadametz/slack-table-emotes/badge.svg?branch=master)](https://coveralls.io/github/jadametz/slack-table-emotes?branch=master)

Table flipping and catching because Slack only builds in one emote ¯\\_(ツ)_/¯

## Setup

I use [up](https://github.com/apex/up) to effortlessly host this. There is minimal (optional only) configuration.

### Environment Variables

| ENV         | Default | Description                                                             |
|-------------|---------|-------------------------------------------------------------------------|
| PORT        | `8080`  | The HTTP server will listen on this port                                |
| ATTACHMENTS | `unset` | If set to `"yes"` responses will include the emote and descriptive text |

## Usage

Depending on your situation, the following are possible:

```
/table flip
    (╯°□°)╯︵ ┻━┻

/table catch
    ┬─┬ノ( º _ ºノ)
```

Enjoy!
