# BDP.go [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/git/git-scm.com/blob/main/MIT-LICENSE.txt)
A Demo Parser for Portal 2 Demos built to improve the experience of verifiers on both the [Challenge Mode LeaderBoards](https://board.portal2.sr/) and the [SRC Leaderboards](https://www.speedrun.com/portal_2).

## Credits
- [@pektezol](https://github.com/pektezol): For both [BitReader](https://github.com/pektezol/BitReader) & [sdp.go](https://github.com/pektezol/sdp.go) 
- [@NeKzor](https://github.com/NeKzor): For [dem.nekz.me](https://dem.nekz.me)
- [@UncraftedName](https://github.com/UncraftedName): For [UntitledParser](https://github.com/UncraftedName/UntitledParser)

## How to Use
Drag and Drop demo files or a directory containing demos onto the executable to generate a json file containing useful information from the demos. The output will look moderately different depending on the mode the parser is set to. 

## TODO
- Multi-file & Folder input
- Multi-threaded Demo Parsing (When multiple demos are present)
- Demo Folder Support
- Output to json
- `cmd_whitelist.txt` support
- `cvar_whitelist.txt` support
- `sar_whitelist.txt` support
- `filesum_whitelist.txt` support
- Parse SAR Data
- More complete parsing of demos
- Regex filename for challenge mode demos
- Support for 4 output modes (Dump/CM/FullGame/MDP)
