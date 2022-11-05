# ARGEDOR

# Kafka
We send the data produced in the producer to kafka and consume the data by running 5 consumers simultaneously by the consumer.

- Kafka(docker) run

We will use docker to run Kafka, docker needs to be installed as well and then all you have to do is use "docker-compose up -d" command.

![Consumer (2)](https://user-images.githubusercontent.com/92402372/200122917-caa15945-743b-4016-8e5b-2a0f98e9557d.png)

<img width="679" alt="Ekran Resmi 2022-11-05 17 02 47" src="https://user-images.githubusercontent.com/92402372/200123691-55a5a799-8dea-41c2-a030-219662a348b6.png">

<img width="1249" alt="Ekran Resmi 2022-11-05 17 01 25" src="https://user-images.githubusercontent.com/92402372/200123687-dbddea02-a120-433c-a759-675e0c75905d.png">


# Rabbitmq
We send the data produced in the producer to rabbitmq and consume the data by running 5 consumers at the same time by the consumer.

- Rabbitmq(docker) run

We will use docker to run rabbitmq, docker needs to be installed as well and then all you have to do is use "docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.10-management" command.

![Consumer (3)](https://user-images.githubusercontent.com/92402372/200122975-91fc046e-c9e5-4d4e-a6bd-6261d6fbb6a9.png)

<img width="581" alt="Ekran Resmi 2022-11-05 16 54 45" src="https://user-images.githubusercontent.com/92402372/200123319-8b524215-d0b8-400f-ab75-afbb0fdff55f.png">

<img width="1252" alt="Ekran Resmi 2022-11-05 16 57 27" src="https://user-images.githubusercontent.com/92402372/200123463-c19a8f14-1771-44f8-b96b-ec127374a5ce.png">
