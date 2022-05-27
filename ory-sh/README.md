# Ory Self Hosted

## Kratos - Identity Management

```sh
git clone https://github.com/ory/kratos.git ; cd kratos
docker-compose -f quickstart.yml -f quickstart-standalone.yml up -d --build --force-recreate
```

> Open `http://localhost:4455`

## Hydra - OAuth2 & OpenID Connect

```sh
git clone https://github.com/ory/hydra.git ; cd hydra
docker-compose -f quickstart.yml -f quickstart-postgres.yml up --build
```

## Keto - Permission & Access Control

```sh
git clone https://github.com/ory/keto.git ; cd keto
docker-compose -f contrib/cat-videos-example/docker-compose.yml up
```
