# Spinel - Secure gateway to internal websites 

## Summary
Spinel (spi·nell) is a small, fast, and secure authentication handler for Nginx that allows your organization to securely serve internal web applications over the public internet without requiring a VPN (Virtual Private Network) connection or other secure tunneling strategies.

It is very **fast** (typically adding less than 200 microseconds to a request), **simple** to setup with existing Nginx reverse proxies (requiring no special modules or compilation), and **secure** with tamper-proof sha256 checksums. Authentication rules can be setup to allow whitelisted ip ranges (specified as CIDR block) or require user login to an Active Directory domain.

## Production Readiness
Beta. This software is not yet rated for production use.

## Installation
Download and install the RPM or DEB pakage from the release page at https://github.com/pgpin/spinel/releases/

Edit the configuration in /etc/spinel.yaml as needed.

Start the service with 

    service spinel start

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
|cidrs|a list of CIDR blocks to whitelist|
|html.logintitle|content to display on login page|

### Example spinel.yaml

    secret: foobar
    expires: 24
    ad:
      host: ad.host 
      port: 389
      dn: OU=FOO,dc=bar,dc=bim,dc=bam
    cidrs:
      - 10.0.0.0/8
    html:
      logintitle: Login with your ActiveDirectory account 

### Example nginx configuration

    upstream myinsecure_backed{
      server myinsecureservice.internal:80
    }

    server{
      server_name mysecureservice.com;
      listen 443;
      # [...] ssl options 
      location /_spinel_auth {
        proxy_pass              http://unix:/tmp/spinel.sock:/_spinel_auth;
      }
      location = /_spinel_auth_check {
        internal;
        proxy_pass              http://unix:/tmp/spinel.sock:/_spinel_auth_check;
        proxy_pass_request_body off;
        proxy_set_header        Content-Length "";
        proxy_set_header        X-Original-URI $request_uri;
        proxy_set_header        X-Original-IP $remote_addr;
      }
      error_page 401 /_spinel_login?url=$request_uri;
      location = /_spinel_login {
        proxy_pass              http://unix:/tmp/spinel.sock:/_spinel_login;
      }
      location / {
        auth_request     /_spinel_auth_check;
        proxy_pass http://myinsecure_backend;
      }
    }

