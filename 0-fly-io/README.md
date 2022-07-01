# Feelantera Infrastructures

Read how to install `flyctl` at: https://fly.io/docs/flyctl/installing

## Quick Start

First thing first, you need to login to Fly.io: `flyctl auth login`

## Shared Database on Fly.io

```sh
flyctl postgres create \
  --organization feelantera \
  --name feeldbcluster \
  --initial-cluster-size 1 \
  --password $(openssl rand -hex 8) \
  --region sin \
  --vm-size shared-cpu-1x \
  --volume-size 3
```

To destroy the database:

```sh
flyctl apps destroy --yes feeldbcluster
```

## Shared Redis on Fly.io

```sh
(cd redis && fly create --org feelantera --name feelredis)
(cd redis && fly volumes create feelredis_data --region sin --size 1 -a feelredis)
(cd redis && flyctl secrets set REDIS_PASSWORD=secretPass -a feelredis)
(cd redis && flyctl deploy)
```

## Useful Fly.io Commands

```sh
# Detach PostgreSQL from the app
flyctl postgres detach --postgres-app feeldbcluster -a {app_name}

# Check all secrets
flyctl secrets list -a {app_name}

# Viewing the Deployed App
flyctl status -a {app_name}

# View application logs
flyctl logs -a {app_name}

# Destroy an application
flyctl apps destroy --yes {app_name}
```
