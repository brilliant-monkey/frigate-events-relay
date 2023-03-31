# Kafka Relay

[![Main](https://github.com/brilliant-monkey/frigate-kafka-relay/actions/workflows/main.yml/badge.svg?branch=main&event=push)](https://github.com/brilliant-monkey/frigate-kafka-relay/actions/workflows/main.yml)

Relays Frigate events from MQTT to Kafka.

## Requirements

- MQTT
- Frigate

## Description

In order to process events more efficiently by other tools (like [Frigate clips](https://github.com/brilliant-monkey/frigate-clips)), a broker should be created to forward data to a one-to-many queue like Kafka. This enables the ability to spread a task across multiple consumers. In the case of Frigate clips, we can have multiple instances of the Clips app running to consume Frigate events in parallel. Mosquitto MQTT does not allow this one-to-many partitioning so tasks must complete in sequence.

## Usage

Use the `docker-compose.yml` file along with the `docker compose up -d` command to create a relay for forwarding Frigate events from MQTT to Kafka.

Before using `docker compose up -d`, make sure a `secret.yml` file is in located in the root of the project folder. An example file `secret.example.yml` is created for you as a template. Rename the file to `secret.yml` and edit it's contents to suit your environment.
