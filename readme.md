# Environment Path Appender

Environment Path Appender is a middleware plugin for [Traefik](https://github.com/traefik/traefik). With this middleware plugin you can add a value from the environment to the request path.

A use-case where this functionality will be useful is for example when an API key needs to be added to the path from a secret,


## Configuration

### Static

```yaml
pilot:
  token: "xxxxx"

experimental:
  plugins:
    environmentpathappender:
      moduleName: github.com/portofrotterdam/environmentpathappender
      version: v0.0.1
```

### Dynamic

```yaml
http:
  middlewares:
    environmentpathappender-foo:
      environmentpathappender:
        env: "MY_ENV_VARIABLE"
```
