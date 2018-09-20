# EXEC_SERVER

Simple HTTP server that runs commands.

Written in Go.

## USE

Launch with a single argument to a config file (see `cfg-example.json`). Once running, call the run endpoint with the desired command. For example, if you're running on localhost:80, and the cfg has a command named "print", then hit the endpoint

```http://localhost:80/run/print```

To run the file configured in the `print` command.