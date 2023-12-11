<img src="https://github.com/kelvinzer0/mac-tty-cleaner/blob/main/example.gif?raw=true" width="100%" style="justify-content: center;align-items: center;"/>


# Mac TTY Cleaner

[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

## Overview

**Mac TTY Cleaner** is a lightweight Go program designed to clean TTY devices on macOS by running a specified command. This utility is particularly useful for system maintenance, providing an easy way to clear TTY consoles.

## Features

- Cleans TTY devices in the /dev/ttys0* directory
- Efficient and lightweight design

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Installation

```bash
git clone https://github.com/kelvinzer0/mac-tty-cleaner.git
cd mac-tty-cleaner
go build -o clears ./cmd/mac-tty-cleaner
```

## Usage

Run the application:

```bash
sudo chmod +x ./clears
sudo mv ./clears /usr/local/bin/clears
```

## Contributing

Contributions are welcome! Open issues or submit pull requests to enhance the functionality.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
