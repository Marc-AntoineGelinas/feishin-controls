# Feishin-Controls
feishin-controls is a CLI project written in go to interact with [feishin](https://github.com/jeffvli/feishin) using its [remote-types](https://github.com/jeffvli/feishin/blob/904f05ff61c5ce90e8edede1b39068d3fb6a3c83/src/shared/types/remote-types.ts) via websockets.

# How to install
## Using release
TODO
## Build manually
feishin-controls is a go project, which means you can easily build the executable for your operating system.

If you do not have go installed, look at the [official install doc](https://go.dev/doc/install)

Clone the project and cd into the new folder
```bash
git clone git@github.com:Marc-AntoineGelinas/feishin-controls.git
cd feishin-controls/
```

Then, simply build the project
```
go build .
```

You will then have an executable that you can use for your os.

# How to use
## Setup
feishin-controls uses Feishin's remote control server. You first need to enable it in Feishin itself.
At the time of writing, you need to go into Settings->General and toggle "Enable remote control server"

As you'll see underneath the control, there is also a port, username and password field. The default values respectively are 4333, feishin and a random unsecure password.

You can modify these values if needed, but if you do not care about authentication you can also leave the username and password fields blank. Only the port is required.

If you've cloned the project, you'll find a config.yml file at the root. Otherwise, you'll need to create it alongside the executable.
Fill in the values as you've configured in Feishin and feishin-controls will read these values at runtime to authenticate to the remote control server.

## Usage
You can then get a list of the different commands using
```bash
./feishin-controls -h
```

To have more info on a specific command, also use the -h flag

```bash
./feishin-controls previous -h
```
