# Spinel provides secure access to your organisations internal web resources

## Summary
Spinel (spi·nelle) is a small, fast, and secure authentication handler for Nginx that allows your organization to securely serve internal web applications over the public internet without requiring a VPN (Virtual Private Network) connection or other secure tunneling strategies. It is very *fast* (typically adding less than 30 microseconds to a request), *simple* to setup with existing Nginx reverse proxies (requiring no special modules or compilation), and *secure* with tamper-proof sha256 checksums. Authentication rules can be setup to allow whitelisted ip ranges (specified as CIDR block) or require user login to an Active Directory domain.

## Installation
Download the single binary file from the release page <link>. It is recommended to run this daemon on the same server as your Nginx daemon to ensure that network health does not add latency to the authentication however it can be run anywhere that is accessible to your Nginx daemon.

    ./spinel -config=/path/to/spinel.yaml

## Configuration
|name| default | description|
|---|---|---|
|secret|NA|a server secret to use in the bearer token checksum|
|ad.host|NA|the hostname of the ldap server|
|ad.port|NA|the port of the ldap server|
|ad.dsn|NA|the base DSN for the ldap quthenticaiton query|
|cidrs|NA|a list of CIDR bocks to whitelist|
