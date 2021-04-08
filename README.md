# goblock

A basic prototype of a Go based blockchain.

From a project structure perspective this is a monorepo, having all the services included into it.

<br/>

## High Level Architecture

The following actors (systems with specific roles) are being considered:

- `node` - it holds the blockchain, receiving transactions (`Txn`s) from `Client`s, and generating `Block`s to be added to the chain
- `delegate` - a support service that generated the hash of a `Block` before being added to the chain
- `validator` - is another support service that validates the generated hash
- `explorer` - provides the UI for getting a visual look of the whole deployment

```
  .-----------.                                             .---------------.
  |  Client   |                                             | Explorer (UI) |
  '-----------'                                             '---------------'
        |                                                           |
        |                                                           |
        |                                                           |
        | submit Txn    .-----------.                               |
        '-------------->|    Node   |<------------------------------'
                        '-----------'                               |
                              |                                     |
                              | 1/2 get Block Hash                  |
                              |                                     |
                              |      .-----------.                  |
                              |----->| Delegate  |<-----------------|
                              |      '-----------'                  |
                              |                                     |
                              | 2/2 validate Block Hash             |
                              |                                     |
                              |      .-----------.                  |
                              '----->| Validator |<-----------------'
                                     '-----------'

```

<br/>
