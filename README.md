# Building a C2 Framework in Go

Course materials for the AntiSyphon workshop at Wild West Hackin' Fest (WWHF) Denver.

**Date:** February 10-11, 2025

## Overview

This repository contains the complete lesson solutions for a 16-hour hands-on workshop on building a Command and Control (C2) framework from scratch using Go.

## Structure

Each lesson builds cumulatively on the previous one:

| Lesson | Topic |
|--------|-------|
| 01-02 | Project setup and basic server |
| 03-04 | Agent foundations |
| 05-06 | HTTPS communication |
| 07-08 | DNS covert channel |
| 09-10 | Protocol switching |
| 11-12 | Cryptographic authentication |
| 13-15 | Command API and queue system |
| 16-17 | Agent execution framework |
| 18-20 | Shellcode execution (reflective DLL loading) |
| 21 | Server results handling |
| 22 | Download command |
| 23 | Persistence mechanisms |

## Requirements

- Go 1.21+
- Windows VM for testing agent capabilities

## Usage

Each `lesson_XX_end` directory contains the complete working code for that lesson:

```bash
cd lesson_23_end
go build ./cmd/server
go build ./cmd/agent
```

## Disclaimer

This material is for educational purposes in authorized security training contexts only.

## Author

Faan Ross - AntiSyphon Training
