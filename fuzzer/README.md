# Build

To build the `Fuzzy Wuzzy` executable, simply run `go build`. To build the module images, run `bash build_images.sh`.

# Usage 

Once built, `Fuzzy Wuzzy` is invoked like any other executable. 
By default, `Go` places the executable in the same folder from which
`go build` was invoked, with the executable taking the name of the the containing folder.


`Fuzzy Wuzzy` only takes one argument, a config file. The config file itself has two sections:

1. Runner config
2. Module config

The runner config contains configuration variables that will be evaluated and passed to each individual module.
The module config is only evaluated and passed to the relevant module.

tl;dr run `./fuzzer <your config file name>` in the folder where you built `Fuzzy Wuzzy`.

## Runner Config

### control

The control field defines a script that will be run before and after each module run.
This script is used to control the starting and stopping of the services being fuzzed.
The script must adhere to several properties:

1. The script takes 1 argument, a string that will be `start` or `stop`. 
`start` will start the services.
`stop` will stop the services.

2. The script is idempotent. 
That means after the script exists after `start`, all services must be up and running and ready to be fuzzed.
After the scripts exists after `stop`, all services must be stopped, and the state must be that of before `stop` was called.

In theory, this is to maintain a clean state for each module run, allowing for reproducable runs. 
In practice, the fuzzer doesn't actually care about the contents of the script. A
All that matters to the fuzzer is that all of the services are available after `start`, 
and after `stop` is ran `start` can be called again without any errors.

### api

The API block contains various fields that defines the API being fuzzed. 
In theory, most of these can be fetched from the schema itself

#### schema

The schema file of the service to be fuzzed.

#### http_scheme

HTTP vs HTTPS. TODO: Obtain from schema

#### host, port

host and port of service. TODO: Obtain from schema

#### security

This block defines various files that are relevant to service security

##### clients_file

A file that contains a list of clients that the fuzzer can use
to make auth requests.
This provides a list of clients, their ids, secrets and scopes.

## Module Config

For the config needed for each module, please see the README for the relevant module.

# Misc.

While the runner itself doesn't define any particular schema types, 
currently only OpenAPI schemas are supported.
