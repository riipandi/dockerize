# Minio on Fly.io

```sh
fly apps create --org domainaja --name domja-minio
fly volumes create minio_data --region sin --size 1 -a domja-minio
fly secrets set $(cat .env | xargs -I %s echo %s)
fly certs create s3.domainaja.id
fly deploy
```
