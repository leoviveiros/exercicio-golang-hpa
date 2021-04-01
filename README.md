# exercicio-golang-hpa

### Comandos de criação do pod "loader"

```
kubectl run -it loader --image=busybox /bin/sh
while true; do wget -q -O- http://go-hpa.default.svc.cluster.local; done;
```