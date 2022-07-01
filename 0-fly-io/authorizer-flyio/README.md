# Authorizer on Fly.io

### Launch and deploy

```sh
# Create new application naemspace
fly create --org feelantera --name feelsso

# Load secrets from dotenv file then initialize deployment
flyctl secrets set $(cat .env | xargs -I %s echo %s)
flyctl secrets set REDIS_URL="redis://default:secretPass@feelredis.internal:6379"
flyctl secrets list

# Attach PostgreSQL to the app
flyctl postgres attach --postgres-app feeldbcluster

# Deploy the app
flyctl deploy
```

### Setup custom domain

Point DNS A Record to the assigned IP address. Or, if using subdomain you
can point `feelsso.fly.dev` CNAME record.

```sh
flyctl ips list
flyctl certs create auth.feel.co.id
```
