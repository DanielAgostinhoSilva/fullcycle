# API GO

## 04 dockerfile

### como buildar uma imagem
```shell
docker build -t my-site .
```

### para listar as imagem my-site
```shell
docker image ls | grep my-site
```

### para rodar a imagem
```shell
docker run --rm --name my-site1 -p 8080:80 my-site
```

### como buildar uma imagem com uma tag
```shell
docker build -t my-site:v1 .
```

### para rodar a imagem com uma tag
```shell
docker run --rm --name my-site1 -p 8080:80 my-site:v1
```

### para remover uma imagem
```shell
docker rmi my-site
```