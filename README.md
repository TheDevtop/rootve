# RootVE

Management framework for chroot environments on NetBSD.

### How to build

1. Clone the repository.
2. Execute the build.sh script.
3. Execute the install.sh script, as root.
4. Start the service.

### How to use

You can manage Virtual Environments with rootctl.

### Example config
```toml
[site]
Root = "/root/site/"
Autoboot = false
Directory = "/root"
Uid = 0
Gid = 0
Environment = ["TERM=xterm", "HOME=/root"]
Interface = "tap0"
CommandPath = "/bin/ksh"
CommandArgs = ["-l"]

```
