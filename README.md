> **Status (2026-09-04):** this repo is being folded into [`consolelabs/console-apps`](https://github.com/consolelabs/console-apps) as `apps/log`, served from Cloudflare Workers. Until that cut it still builds and serves https://log.console.so/ from GitHub Pages. The build advances the `vault` submodule to `consolelabs/content` `main` at build time and runs hourly and on demand (Actions, Run workflow); the old submodule-bump dispatch is gone. After the cut this repo is archived.

# log.console.so

[![Built with Devbox](https://jetpack.io/img/devbox/shield_galaxy.svg)](https://jetpack.io/devbox/docs/contributor-quickstart/)

[![Open in DevPod!](https://devpod.sh/assets/open-in-devpod.svg)](https://devpod.sh/open#https://github.com/consolelabs/log.console.so)

This is the repository where we keep our internal notes from everything related to our engineering practices, experiments, as well as bounty notices. We use this repo to share our knowledge, insights, and experiences with each other before we make it generally available to our community.

## Getting Started

We welcome contributions from anyone who is interested in our topics. You can contribute by creating a new note, editing an existing note, or commenting on a note. To do so, you need to follow these steps:

1. Fork this repo and clone it to your local machine.
2. Install [Devbox](^9^), a command-line tool that lets you easily create isolated shells and containers for development. Devbox will help you set up a consistent and reproducible environment for this repo, with all the necessary tools and dependencies installed.
3. Run `devbox shell` in the root directory of the repo to enter an isolated shell. Alternatively, you can use [VSCode](https://code.visualstudio.com/) with the [Remote - Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) extension to open the repo as a devcontainer.

To run the server, you can run our Makefile command:
```sh
make watch-run
```

or through devbox:
```sh
devbox run watch-run
```

## Code of conduct

We expect all contributors to adhere to our [code of conduct](^15^), which is based on the [Contributor Covenant](https://www.contributor-covenant.org/). By participating in this project, you agree to abide by its terms. Please report any unacceptable behavior to [han@console.so](mailto:han@console.so).
