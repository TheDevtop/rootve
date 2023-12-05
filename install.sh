#!/bin/sh

# Run this script as root!

cp rootexec/rootexec /usr/local/bin/rootexec && echo 'Installed: rootexec'
chmod ug+s /usr/local/bin/rootexec

cp rootd/rootd /usr/local/bin/rootd && echo 'Installed: rootd'
chmod ug+s /usr/local/bin/rootd

cp rootctl/rootctl /usr/local/bin/rootctl && echo 'Installed: rootctl'

cp rc.d/rootd /etc/rc.d/rootd && echo 'Installed: Service file'
chmod a+x /etc/rc.d/rootd

exit 0
