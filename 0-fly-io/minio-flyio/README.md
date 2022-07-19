# Minio on Fly.io

### Launch and deploy

```sh
fly create --org personal --name minio
fly volumes create minio_data --region lax --size 1 -a minio
fly secrets set $(cat .env | xargs -I %s echo %s)
fly certs create minio.example.com
fly deploy
```
