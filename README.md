# Bofin

A command line tool that can be used by to make Weblink development more productive.

## Install

To install bofin, you should run the [install script](https://github.com/gbannerman/bofin/blob/main/install.sh). To do that, you may either download and run the script manually, or use the following cURL command:

```bash
curl https://raw.githubusercontent.com/gbannerman/bofin/main/install.sh | bash
```

## Setup

During install, a `.bofin-conf.yaml` file will be created in your home directory. You should update this with the required information.

You should also create a `.env.template` file in your `boot_dir`. This template will be used to generate your `.env` file when boot-related command is run. The following variables should be added to your `.env.example` and will be replaced when generating your `.env`:

- {{.ApiAccessToken}}
- {{.WeblinkStagingDomain}}

## Building from source

You can create a build of bofin by installing `go` and running the following command:

```bash
GOOS=linux GOARCH=amd64 go build .
```

The supported values for `GOOS` and `GOARCH` are [here](https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63).
