# go-kit

## Endpoint

Go kit primarily deals in the RPC messaging pattern. We use an abstraction called an endpoint to model individual RPCs. An endpoint can be implemented by a server, and called by a client. It's the fundamental building block of many Go kit components.