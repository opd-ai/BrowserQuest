BrowserQuest
============

BrowserQuest is a HTML5/JavaScript multiplayer game experiment.

Documentation
-------------

Documentation is located in client and server directories.

Go client bootstrap
-------------------

This repository now also contains the milestone-0 bootstrap for a Go/Ebitengine client that keeps the original BrowserQuest server unchanged.

- Desktop entrypoint: `/home/runner/work/BrowserQuest/BrowserQuest/opd-ai/BrowserQuest/cmd/client-desktop`
- WASM entrypoint: `/home/runner/work/BrowserQuest/BrowserQuest/opd-ai/BrowserQuest/cmd/client-wasm`
- Map conversion tool: `/home/runner/work/BrowserQuest/BrowserQuest/opd-ai/BrowserQuest/cmd/tools-mapconv`
- Embedded asset bundle: `/home/runner/work/BrowserQuest/BrowserQuest/opd-ai/BrowserQuest/assets`

Example commands:

- `go run ./cmd/client-desktop`
- `GOOS=js GOARCH=wasm go build -o client-wasm.wasm ./cmd/client-wasm`
- `go run ./cmd/tools-mapconv -in assets/maps/world_client.json -out /tmp/world_client.json`

License
-------

Code is licensed under MPL 2.0. Content is licensed under CC-BY-SA 3.0.
See the LICENSE file for details.

Credits
-------

Created by [Little Workshop](http://www.littleworkshop.fr):

* Franck Lecollinet - [@whatthefranck](http://twitter.com/whatthefranck)
* Guillaume Lecollinet - [@glecollinet](http://twitter.com/glecollinet)
