# 概要
blue green デプロイテストのために異なるDockerImageをビルドしている。

blue
```
docker build -t mygoapp:blue --build-arg APP_VERSION=Blue .
```

green
```
docker build -t mygoapp:blue --build-arg APP_VERSION=Green .
```

# 環境構築
## clusterの構築
```
kind create cluster --config kind-config.yaml 
```

## Istioのインストール
```
istioctl install --set profile=demo -f nodeport.yml -y
```

## Istioのインストール確認
```
$ kubectl get svc -n istio-system
NAME                   TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)                                                                      AGE
istio-egressgateway    ClusterIP   10.96.129.31    <none>        80/TCP,443/TCP                                                               64s
istio-ingressgateway   NodePort    10.96.109.201   <none>        15021:32107/TCP,80:30080/TCP,443:30236/TCP,31400:31459/TCP,15443:32178/TCP   64s
istiod                 ClusterIP   10.96.105.170   <none>        15010/TCP,15012/TCP,443/TCP,15014/TCP                                        80s
```