# Kubernetes


### como instalar o helm
[Documentação do helm](https://helm.sh/docs/intro/install/)
```shell
brew install helm
```

### repositorio de artefatos do helm
[repositorio do helm](https://artifacthub.io/)


### como instalar o manifesto do mysql
#### 1- rode o comando
```shell
helm install my-server oci://registry-1.docker.io/bitnamicharts/mysql
```

#### 2- execute o comando para getar as crendenciais de administrador
```shell
MYSQL_ROOT_PASSWORD=$(kubectl get secret --namespace default my-server-mysql -o jsonpath="{.data.mysql-root-password}" | base64 -d)
```

#### 4- para saber qual crendential foi criada
```shell
echo $MYSQL_ROOT_PASSWORD
```

#### 5- execute um pod com um client do mysql
```shell
kubectl run my-server-mysql-client --rm --tty -i --restart='Never' --image  docker.io/bitnami/mysql:8.0.35-debian-11-r2 --namespace default --env MYSQL_ROOT_PASSWORD=$MYSQL_ROOT_PASSWORD --command -- bash
```

#### 6- como se conectar ao mysql
```shell
mysql -h my-server-mysql.default.svc.cluster.local -uroot -p"$MYSQL_ROOT_PASSWORD"
```
