# smtp2http

A simple SMTP server that forwards incoming emails to an HTTP endpoint in
the json format compatible with `cloudflare-email` worker.

Docker image: `ghcr.io/robinbraemer/cloudflare-email/smtp2http:latest`

## Usage

```shell
Usage: smtp2http [--addr ADDR] --user USER --pass PASS --post-url POST-URL --post-token POST-TOKEN

Options:
  --addr ADDR [default: :1587, env: ADDR]
  --user USER [env: USER]
  --pass PASS [env: PASS]
  --post-url POST-URL [env: POST_URL]
  --post-token POST-TOKEN [env: POST_TOKEN]
  --help, -h             display this help and exit
```
