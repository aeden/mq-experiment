## Message Queue Experiment

This repo contains a simple demonstration of setting up a message queue, posting to it from a web API and then consuming those messages and processing them from a backend.

The goal is to demonstrate how an HTTP API application could be designed using asynchronous messaging at its core.

### Running

You'll need Go installed. Instructions here:

[https://golang.org/doc/install](https://golang.org/doc/install)

Next, install the gnatsd service and forego and run the server:

`./install.sh && cd server && forego start`

In a new console:

`cd webapp && make && forego start`

In a new console:

`cd worker && make && forego start`

## Fire and Forget

Once you have the three services up and running you can hit the web API with a simple curl call:

`curl -i http://localhost:8080/ping`

You should receive a 200 response from the webapp and you should see a message on the worker console. This demonstrates the concept of fire-and-forget, where a message is sent and the caller is not concerned with the result.
