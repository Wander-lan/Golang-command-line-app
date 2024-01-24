This is a simple application that shows IPs and server names from a given Url.

Usage:
There are two main commands in this: "ip" and "servers_names", and one flag: "host".
The "ip" command shows a list of the public IPs of a given domain. The "servers_names" command shows a list the servers names from a given domain.
If no domain is given than the app will use the default domain that is "google.com.br".

Example:
./command-line ip --host github.com
./command-line servers_names --host github.com
