# UWA PubSub INIT

UWA PubSub INIT pubsub topics and subscribers.

# Configuration

These configurations are required in order the initializer to works.

## Environment Variables

| Name          | Description   | Example       |
| ------------- | ------------- | ------------- |
| PUBSUB_EMULATOR_HOST | set the Google's pubsub emulator host | `localhost:8085` |
| ProjectID | set the projectID of the emulator | `local-project` |
| PubSubInitFile | specify the file path of the init file | `init.pubsub/my-init-file.yaml` |

## Init File

The init file is essential to initialize topic and subscriber. The configuration is arranged like this:

```yaml
topics:
    my.topic.name:
        - my.topic.name.subscriber1
        - my.topic.name.subscriber2
    my.topic.name.without.subs:
```

# Running application

## Single run

Typically you would want to run it in a containerized application. But generally, prepare the configurations above and run as a go application.

## Docker Container

Minimum docker container configuration that also run the pubsub emulator:

```yaml
version: '3.8'
services:
  pubsub-emulator:
    image: google/cloud-sdk:emulators
    restart: always
    command: gcloud beta emulators pubsub start --host-port=0.0.0.0:8085
    ports:
      - "8085:8085"
    environment:
      - PUBSUB_PROJECT_ID=local-project

  pubsub-init:
    image: ghcr.io/anditakaesar/uwa-pubsub-init:latest
    environment:
      - PUBSUB_EMULATOR_HOST=pubsub-emulator:8085
      - ProjectID=local-project
      - PubSubInitFile=my-init-file.yaml
    volumes:
      - ./my-init-file.yaml:/root/my-init-file.yaml
    depends_on:
      - pubsub-emulator
```

Run with docker compose

```
docker compose up -d
```

# License

MIT