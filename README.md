# Infoblox CLI tool written in go for managing various infoblox things

To get started, download app from the releases and configure with:
```bash
# use your infoblox ip address for host
> infoblox config set
Infoblox Host: 10.5.5.5
Infoblox Username: exampleuser
Infoblox Password: '****************'
Config saved successfully
```

Once you have it configured with your credentials see help for further info on whats possible:
```bash
❯ infoblox -h
Infoblox cli tool for managing things and stuff

Usage:
  infoblox [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  config      config management
  dns         dns record management
  help        Help about any command

Flags:
  -h, --help   help for infoblox

Use "infoblox [command] --help" for more information about a command.
```

### Adding a DNS record
```bash
❯ infoblox dns records add myserver1.mycompany.com 10.10.10.42
```

### Removing a DNS record
```bash
❯ infoblox dns records remove myserver1.mycompany.com
```
