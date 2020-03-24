# Twitter Reporter

É projeto que realiza o consumo das APIs do [Twitter](https://developer.twitter.com/). Após iniciar o projeto, é realizado uma request para o endpoint `/reporter-api/v1/reporters` do serviço [twitter-reporter-api](https://github.com/dalmarcogd/twitter-reporter/tree/master/twitter-reporter-api) com a hashtag que deve ser encontrada no Twitter. Após isso API publica uma mensagem em um tópico do RabbitMQ que por sua ver, direciona para uma fila. O serviço [twitter-reporter-processor](https://github.com/dalmarcogd/twitter-reporter/tree/master/twitter-reporter-processor) consome as mensangens das filas, procura no twitter os tweets que contenham a hashtag e os persiste em uma base.

Os dois serviços são monitorados pelo APM Server da Elastic, que por sua vez utiliza o Elasticsearch e apresenta os dados através do Kibana.

```sh
$ cd digital-account
$ docker-compose up -d rabbit
$ docker-compose up -d kibana
$ docker-compose up -d postgres
$ docker-compose up -d twitter-reporter-api twitter-reporter-processor 
```
![alt text](https://github.com/dalmarcogd/twitter-reporter/blob/master/Design.png)

[Download collections](https://github.com/dalmarcogd/twitter-reporter/blob/master/Twitter%20Reporter.postman_collection.json)
