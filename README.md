# dynatrace-go-metrics
Go binary makes it easy to push custom metrics into Dynatrace.

## API Token
Create an API token with `API v2: Ingest Metrics` permissions

## Custom Metric Format

Custom metrics take the following format: `metric key SPACE value`:

```
batchjobs.myjob.processedrows 765
```

You can also add a dimension which splits the metric in the format: `batchjobs.myjob,dimension splitting key=value SPACE value to push`. You receive a new datapoint for each metric dimension.

For example, if a metric had a `key` of `batchjobs.myjob` and you'd like to pass dimensions with that metric:

```
batchjobs.myjob,results="Status" 1
batchjobs.myjob,results="Time" 123
batchjobs.myjob,results="Rows Processed" 40000
```

## Usage

Push multiple metrics using the [line protocol](https://www.dynatrace.com/support/help/how-to-use-dynatrace/metrics/metric-ingestion/metric-ingestion-protocol/). Split each metric with a delimiter.

By default, the delimiter using this tool is `#`. You can override this using the `DT_DELIMITER=` environment variable (see additional examples below)

Remove all trailing slashes from the `BASEURL`.

### Windows: Use Default Delimiter

Using Default Delimiter (`#`)

```
set DT_BASEURL=https://***.live.dynatrace.com
set DT_APITOKEN=***
set DT_METRICSLIST=documents.processed.success,documenttype="Mortgages" 3#documents.processed.success,documenttype="Loans" 143
.\bin\windows\dt-send-metrics-x64.exe
```

### Windows: Use Custom Delimiter
```
set DT_DELIMITER=#!#!#
set DT_BASEURL=https://***.live.dynatrace.com
set DT_APITOKEN=***
set DT_METRICSLIST=documents.processed.success,documenttype="Mortgages" 3#!#!#documents.processed.success,documenttype="Loans" 143
.\bin\windows\dt-send-metrics-x64.exe
```

### Linux: Use Default Delimiter

Using Default Delimiter (`#`)

```
export DT_BASEURL=https://***.live.dynatrace.com
export DT_APITOKEN=***
export DT_METRICSLIST=documents.processed.success,documenttype="Mortgages" 3#documents.processed.success,documenttype="Loans" 143
./bin/linux/dt-send-metrics-x64
```

### Linux: Use Custom Delimiter
```
export DT_DELIMITER=#!#!#
export DT_BASEURL=https://***.live.dynatrace.com
export DT_APITOKEN=***
export DT_METRICSLIST=documents.processed.success,documenttype="Mortgages" 3#!#!#documents.processed.success,documenttype="Loans" 143
./bin/linux/dt-send-metrics-x64
```

## Building
```
make build_windows    // or
make build_linux      // or
make build_macos      // or

// Build for all OS
make build
```
