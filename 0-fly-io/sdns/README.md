# SDNS at fly.io

Read how to install `flyctl` at: https://fly.io/docs/flyctl/installing

## Quick Start

First thing first, you need to login to Fly.io: `flyctl auth login`

## Up and Running

```sh
flyctl apps create --org personal --name sdns --region lax --no-deploy
# flyctl launch --org personal --name sdns --region lax --no-deploy
# flyctl secrets set $(cat .env.production | xargs -I %s echo %s)
flyctl deploy

# Custom domain
flyctl certs create bcr.web.id
flyctl certs show bcr.web.id
```
