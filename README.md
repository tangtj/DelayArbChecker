# Delay Arb Checker

This README was written with the assistance of ChatGPT.

## Overview

The Delay Arb Checker is a tool designed to measure the latency of a service by performing repeated tests to an EVM RPC endpoint. It offers a straightforward way to monitor and analyze the response times of your RPC endpoints, which is critical for performance tuning and service health monitoring.

## Features

- Perform latency tests against a specified RPC URL.
- Customize the number of test iterations with ease.
- Simple command-line interface for easy integration into monitoring systems.

## Getting Started

### Prerequisites

Before using the Delay Arb Checker, ensure you have:

- A system with command-line access.
- Network access to the RPC endpoint you wish to test.

### Download

To begin using the Delay Arb Checker, download the latest version directly from the GitHub releases:

1. Visit the [latest release page](https://github.com/tangtj/DelayArbChecker/releases/download/latest) for the Delay Arb Checker project.
2. Download the `checker` binary suitable for your operating system.
3. Make sure to set the downloaded binary as executable. For Unix-like systems, you can use `chmod +x checker`.

```shell
curl -o checker https://github.com/tangtj/DelayArbChecker/releases/download/latest/checker

chmod +x checker
```

### Usage

To use the tool, open a command prompt or terminal.

Execute the application with the following options:

```bash
./checker -url <rpc-url> -t <number-of-tests>
```

- `-url` specifies the RPC URL you want to test; the default is "https://arb1-sequencer.arbitrum.io/rpc".
- `-t` sets the number of latency tests to perform. If not specified, the default is 20.

Example command:

```bash
./checker -url https://arb1-sequencer.arbitrum.io/rpc -t 20
```

This will perform 20 latency tests against the specified Arbitrum RPC endpoint.