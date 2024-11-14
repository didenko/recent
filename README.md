The `recent` utility collects a number of most recently modified files from the specified directory tree. Current directory is used by default.

Usage:

    recent [flags] [folder]

Flags:

  `-d`, `--date`        include date in file timestamps output

  `-h`, `--help`        help for recent

  `-n`, `--limit <int>`   a number of recent files to list (default 10)

  `-m`, `--mill`        include milliseconts in file timestamps output - only if time included as well

  `-t`, `--time`        include time in file timestamps output

