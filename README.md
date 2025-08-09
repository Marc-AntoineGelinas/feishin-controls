# Feishin-Controls
feishin-controls is a CLI project written in go to interact with [feishin](https://github.com/jeffvli/feishin) using its [remote-types](https://github.com/jeffvli/feishin/blob/904f05ff61c5ce90e8edede1b39068d3fb6a3c83/src/shared/types/remote-types.ts) via websockets.

# How to install
## Using release
Download the release for your operating system.

You may need to set permissions on the file to be able to use it.

On Linux :
```bash
chmod +x feishin-controls
```

Download config.yml from this repo, or create it manually with the same name. Then refer to the [setup section](#setup)
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

You'll use these values to initialize feishin-controls' config file.
You can run
```bash
./feishin-controls init
```
for an interactive prompt where you'll be asked to fill in the parameters.
Or you can use
```bash
./feishin-controls init [url] [username] [password]
```
to directly pass the parameters and create the config file
## Usage
You can then get a list of the different commands using
```bash
./feishin-controls -h
```

To have more info on a specific command, also use the -h flag

```bash
./feishin-controls previous -h
```
