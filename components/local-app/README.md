# local-app

## nxpod-cli

All of the accessible commands can be listed with `nxpod --help` .

### Installing

1. Download the CLI for your platform and make it executable:

```bash
wget -O nxpod https://nxpod.khulnasoft.com/static/bin/nxpod-cli-darwin-arm64
chmod u+x nxpod
```

2. Optionally, make it available globally. On macOS:

```bash
sudo mv nxpod /usr/local/bin/
```

### Usage

Start by logging in with `nxpod login`, which will also create a default context in the configuration file (`~/.nxpod/config.yaml`).

### Development

To develop the CLI with Nxpod, you can run it just like locally, but in Nxpod workspaces, a browser and a keyring are not available. To log in despite these limitations, provide a PAT via the `NXPOD_TOKEN` environment variable, or use the `--token` flag with the login command.

#### In a Nxpod workspace

[![Open in Nxpod](https://www.nxpod.khulnasoft.com/svg/open-in-nxpod.svg)](https://nxpod.khulnasoft.com/#https://github.com/nxpkg/nxpod)

You will have nxpod-cli ready as `nxpod` on any Workspace based on `https://github.com/nxpkg/nxpod`.

```
# Reinstall `nxpod`
leeway run components/local-app:install-cli

# Reinstall completion
leeway run components/local-app:cli-completion
```

### Versioning and Release Management

The CLI is versioned independently of other Nxpod artifacts due to its auto-updating behaviour.
To create a new version that existing clients will consume increment the number in `version.txt`. Make sure to use semantic versioning. The minor version can be greater than 10, e.g. `0.342` is a valid version.

## local-app

**Beware**: this is very much work in progress and will likely break things.

### How to install

```
docker run --rm -it -v /tmp/dest:/out ghcr.io/nxpkg/build/local-app:<version>
```

### How to run

```
./local-app
```

### How to run in Nxpod against a dev-staging environment

```
cd components/local-app
BROWSER= NXPOD_HOST=<URL-of-your-preview-env> go run main.go --mock-keyring run
```
