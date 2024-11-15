blue green デプロイテストのために異なるDockerImageをビルドしている。

blue
```
docker build -t mygoapp:blue --build-arg APP_VERSION=Blue .
```

green
```
docker build -t mygoapp:blue --build-arg APP_VERSION=Green .
```