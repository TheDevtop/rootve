# RootVE

Management framework for chroot environments on NetBSD.

### How to build

1. Clone the repository.
2. Execute the build.sh script.
3. Execute the install.sh script, as **root**.
4. Start the service.

### How to use

You can manage Virtual Environments with rootctl:

- Start VE `rootctl start [name]`
- Stop VE `rootctl stop [name]`
- Spawn shell `rootctl shell [name]`
- Remove VE `rootctl rm [name]`
- Pause VE `rootctl pause [name]`
- Resume VE `rooctl resume [name]`
- List all VE `rootctl ls`
- List active VE `rootctl ps`

### Example config
```toml
[site]
Root = "/root/site/"
Autoboot = false
Directory = "/root"
Uid = 0
Gid = 0
Environment = ["TERM=xterm", "HOME=/root"]
CommandPath = "/bin/ksh"
CommandArgs = ["-l"]
Networking = true
Interface = "tap0"
Bridge = "bridge0"
AddressV4 = "192.168.1.10/24"
```
