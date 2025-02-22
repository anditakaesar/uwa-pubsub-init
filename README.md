# UWA PubSub INIT

UWA PubSub INIT pubsub topics and subscribers.

# Configuration

These configurations are required in order the initializer to works.

## Environment Variables

PUBSUB_EMULATOR_HOST = set the Google's pubsub emulator host e.g. `localhost:8085`
ProjectID = set the projectID of the emulator e.g. `local-project`
PubSubInitFile = specify the file path of the init file e.g. `init.pubsub/my-init-file.yaml`

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

TBD

# License

MIT