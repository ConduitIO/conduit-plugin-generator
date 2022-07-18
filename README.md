# Conduit Connector Generator

### General

The generator connector is one of [Conduit](https://github.com/ConduitIO/conduit) builtin plugins. It generates sample
records using its source connector. It has no destination and trying to use that will result in an error.

### How to build it

Run `make`.

### Testing

Run `make test` to run all the unit tests.

### How it works

The data is generated in JSON format. The JSON objects themselves are generated using a field specification, which is
explained in more details in the [Configuration section](#Configuration) below.

The connector is great for getting started with Conduit but also for certain types of performance tests.

It's possible to simulate a 'read' time for records. It's also possible to simulate bursts through "sleep and generate"
cycles, where the connector is sleeping for some time (not generating any records), then generating records for the 
specified time, and then repeating the same cycle. The connector always start with the sleeping phase.

### Configuration

| name               | description                                                                                                          | required | Default             |
|--------------------|----------------------------------------------------------------------------------------------------------------------|----------|---------------------|
| recordCount        | Number of records to be generated. -1 for no limit.                                                                  | false    | "-1"                |
| readTime           | The time it takes to 'read' a record.                                                                                | false    | "0s"                |
| burst.sleepTime    | The time the generator 'sleeps'. Must be non-negative.                                                               | false    | "0s"                |
| burst.generateTime | The amount of time the generator is generating records. Must be positive.<br/>Has effect only if `sleepTime` is set. | false    | max. duration in Go |
| fields             | A comma-separated list of name:type tokens, where type can be: int, string, time, bool.                              | true     | ""                  |
| format             | Format of the generated payload data: raw, structured.                                                               | false    | "raw"               |
