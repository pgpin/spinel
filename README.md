# Spinel - Secure access to internal websites

## Summary
Spinel (spi·nell) is a small, fast, and secure authentication handler for Nginx that allows your organization to securely serve internal web applications over the public internet without requiring a VPN (Virtual Private Network) connection or other secure tunneling strategies.

It is very **fast** (typically adding less than 200 microseconds to a request), **simple** to setup with existing Nginx reverse proxies (requiring no special modules or compilation), and **secure** with tamper-proof sha256 checksums. Authentication rules can be setup to allow whitelisted ip ranges (specified as CIDR block) or require user login to an Active Directory domain.

## Production Readiness
Beta. This software is not yet rated for production use.

## Installation
Download the single binary file from the release page <link>.

    ./spinel -config=/path/to/spinel.yaml

## Configuration
|name| description|
|---|---|
|debug|enable debug logging|
|socket|unix socket file to bind|
|expire|Number of hours to consider a token valid|
|secret|a server secret to use in the bearer token checksum|
|ad.host|the hostname of the ldap server|
|ad.port|the port of the ldap server|
|ad.dn|the BaseDN for the ldap quthenticaiton query|
|cidrs|a list of CIDR bocks to whitelist|
|html.logintitle|content to display on login page|
