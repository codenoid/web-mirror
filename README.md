# Web Mirror

![screenshot](screenshot.png)

Based on [reverseproxy](https://github.com/cssivision/reverseproxy)

Solution for:

1. Website that only set 'X-Frame-Options' to 'sameorigin'.
2. Hide website real url
3. Content interception & manipulation [reverseproxy.go#L311](https://github.com/codenoid/web-mirror/blob/master/reverseproxy.go#L311)

## Setup

```sh
# Installation
curl -sf https://gobinaries.com/codenoid/web-mirror@latest | sh

# Running UP
export BIND="0.0.0.0:7000"
./web-mirror
```

## Usage

1. Put requested url to ?mirror= args
2. the payload must be url encoded

Open [http://localhost:7000/?mirror=https%3A%2F%2Fwww.tiktok.com%2F%40khaby.lame](http://localhost:7000/?mirror=https%3A%2F%2Fwww.tiktok.com%2F%40khaby.lame)
