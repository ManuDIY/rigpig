[![CircleCI](https://circleci.com/gh/serainville/rigpig/tree/master.svg?style=svg)](https://circleci.com/gh/serainville/rigpig/tree/master) [![Go Report Card](https://goreportcard.com/badge/github.com/serainville/rigpig)](https://goreportcard.com/report/github.com/serainville/rigpig) [![Codacy Badge](https://api.codacy.com/project/badge/Grade/77f9cb5f1a5d4f089c8f4edd32216372)](https://www.codacy.com/project/serainville/rigpig/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=serainville/rigpig&amp;utm_campaign=Badge_Grade_Dashboard)

# RIGPIG
Rigpig is a single pool / multi-pool cryptocurrency miner. It also operates
as an centralized remote rig manager.

Rigpig is written in Go for the purpose of providing a multi-platform solution.
The following platforms are supported.

- Windows 64-bit
- OSX
- Linux 64-bit
- Linux 32-bit

## Server
A RigPig server is a dual-purpose role - it runs as a standalone rig and
provides centralized management for remote rigs.
 

## Remote Agent
A remote RigPig agent is a mimimalistic service that runs on a remote
host. It connects to and is controlled by a server, reporting all status and
operations to the server.


## Feature Implementation Status
### In Progress
- Standalone miner -- *in progress*
- Central miner manager -- *in progress*

### Not Started
- Remote Miner
- Download latest miners
- Add custom miner
- Add custom pools
- GPU Benchmarking
- CPU Benchmarking
- Desktop GUI
- Cloud Service for remote management


## Version History
1.0.0
- Initial release
