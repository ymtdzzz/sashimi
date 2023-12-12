# sashimi
sashimi is a command-line tool that provides an abstracted interface for batch job splitting.

If you're a batch implementer, implement the following proto and place it in an environment (like a container) where it can be executed as a command to start the gRPC server (This command uses [go-plugin](https://github.com/hashicorp/go-plugin) to communicate sashimi with jobs).

```proto
message SplitJobResponse {
  repeated string commands = 1;
}

service Job {
  rpc SplitJob(google.protobuf.Empty) returns (SplitJobResponse) {}
}
```

For example, for the job executed in the form `job [start] [end]`, an array of individual commands chunked according to the number of jobs should be returned.

```sh
# split command implemented the proto of Job service
$ job_split
## it returns: ["job 1 3000", "job 3001 4000", "job 4001 4238"]

# main job command
$ job [start] [end]
```

By executing `sashimi` (actually through a service such as AWS EventBridge), the commands for splitting can be invoked transparently via gRPC and the resulting split commands can be retrieved as an array. By passing these commands to a scalable container-based batch execution infrastructure such as AWS Batch, it is possible to standardise the processing involved in splitting jobs.

```sh
$ CMDS=$(sashimi job_split)
# returns: "job 1 3000,job 3001 4000,job 4001 4238"
# split it by ',' and execute each command on distributed containers.
```
