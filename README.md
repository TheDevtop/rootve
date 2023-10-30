# RootVE

Management framework for chroot environments on NetBSD.

### Example config
```toml
[site]
Root = "/root/site/"
Directory = "/root"
Uid = 0
Gid = 0
Environment = ["TERM=xterm", "HOME=/root"]
CommandPath = "/bin/ksh"
CommandArgs = ["-l"]

```
