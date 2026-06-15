# PLAN.md — BrowserQuest Client → Go / Ebitengine Port

> **Project**: `gobrowserquest-client` (working title)
> **Goal**: Port **only the client** of the deprecated
> [mozilla/BrowserQuest](https://github.com/mozilla/BrowserQuest) to **Go**, using
> **[Ebitengine](https://ebitengine.org/)**, targeting **WASM (browser)**,
> **Desktop** (Windows/macOS/Linux), and **Mobile** (Android/iOS via gomobile).
> **The existing server is kept as-is** — the Go client must speak the original
> BrowserQuest wire protocol over WebSocket without any server changes.
> **License compliance**: Code under MPL 2.0, content (art, maps, audio) under CC-BY-SA 3.0.
> Credit retained to *Little Workshop* (Franck & Guillaume Lecollinet).

---

## 0. Guiding Principles

- **Server is authoritative and unchanged**: the Go client renders & predicts; it must be
  byte-compatible with the original protocol so it can connect to the live/existing server.
- **Shared protocol package**: a single Go package mirrors `shared/js/gametypes.js`
  (message opcodes, entity kinds, item ranks, orientations) for use across the client.
- **Extensibility first**: every client subsystem (rendering layers, UI widgets, input
  schemes, audio backends, protocol message handlers, asset providers) exposes a
  registry + interface so downstream adopters add features *without forking core*.
- **Platform-agnostic core**: no `build`-tag leakage into game logic. Platform specifics
  isolated behind interfaces (storage, audio, transport, asset loading).
- **Asset fidelity**: reuse original sprites/maps/sounds where license allows; convert formats only.

---

## 1. Repository Layout (Target)

```
gobrowserquest-client/
├── PLAN.md
├── LICENSE                # MPL 2.0 + CC-BY-SA 3.0 notices
├── go.mod
├── cmd/
│   ├── client-desktop/    # Ebitengine desktop entrypoint
│   ├── client-wasm/       # WASM build entrypoint + index.html shell
│   └── tools-mapconv/     # Tiled/JSON map → Go map format converter
├── internal/              # non-extensible internals
├── pkg/
│   ├── protocol/          # message codec + opcode registry (EXT) — mirrors gametypes.js
│   ├── types/             # entity kinds, item ranks, orientations (from original Types)
│   ├── gamemap/           # client map model: tiles, collisions, zones, doors, checkpoints
│   └── client/
│       ├── game/          # client game loop, world model, prediction/interpolation
│       ├── render/        # layered renderer (EXT layers)
│       ├── input/         # input mapping + pathfinding (EXT)
│       ├── ui/            # HUD, chat, inventory, achievements (EXT widgets)
│       ├── audio/         # audio interface (EXT)
│       ├── netclient/     # WebSocket client transport (EXT)
│       └── platform/      # build-tagged platform adapters (storage/audio/transport)
├── assets/                # converted sprites, maps, sfx (CC-BY-SA)
└── docs/
```

---

## 2. Extension-Point Catalogue (Cross-Cutting)

Every item below ships a **registry + interface** pattern (`Register(name, factory)`),
discoverable via a central `Registry` so mods self-register in `init()`.

- [ ] **Protocol message handlers** — register client-side handlers for opcodes; add new
      messages the server may send (forward-compat) without editing core.
- [ ] **Entity kinds** — register new entity archetypes with sprite/animation/serialize hooks.
- [ ] **Render layers** — insert custom draw passes (weather, lighting, day/night, debug).
- [ ] **UI widgets** — register HUD panels and modal screens.
- [ ] **Input schemes** — remappable actions; touch/gamepad/keyboard providers.
- [ ] **Audio backends** — platform-specific implementations behind one interface.
- [ ] **Asset providers** — embedded FS, HTTP, or platform bundle.
- [ ] **Storage drivers** — browser localStorage / file / mobile prefs behind one interface.
- [ ] **Transport** — WebSocket default; pluggable (in-process for tests, replay/record).
- [ ] **Client event hooks** — `OnSpawn`, `OnDespawn`, `OnDeath`, `OnChat`, `OnZoneChange`.

---

## 3. Milestones

### Milestone 0 — Project Bootstrap & Asset Recovery
**Outcome**: Buildable empty client skeleton; original assets converted & licensed.

- [ ] Initialize Go module, CI (build matrix: linux/win/mac/wasm/android/ios).
- [ ] Set up MPL 2.0 + CC-BY-SA 3.0 headers and attribution to Little Workshop.
- [ ] Extract original `client/maps`, `client/sprites`, `client/audio`.
- [ ] Write `cmd/tools-mapconv` to convert BrowserQuest JSON maps → Go map model.
- [ ] Define `pkg/gamemap` model (tiles, collisions, zones, checkpoints, doors).
- [ ] Embed assets via `embed.FS` with pluggable `AssetProvider` interface.
- [ ] Stub Ebitengine window that opens & clears screen on all 3 platform builds.

**Exit criteria**: `go build ./...` passes for every target; window opens on desktop + WASM.

---

### Milestone 1 — Protocol Parity with Existing Server
**Outcome**: A Go client that can connect to and talk with the unmodified server.

- [ ] Catalogue server messages (`shared/js/gametypes.js`): HELLO, WELCOME, SPAWN,
      DESPAWN, MOVE, LOOTMOVE, AGGRO, ATTACK, HIT, HURT, HEALTH, CHAT, EQUIP,
      DROP, TELEPORT, DAMAGE, POPULATION, KILL, LIST, WHO, ZONE, DESTROY,
      BLINK, OPEN, CHECK, etc.
- [ ] Implement `protocol.Codec` matching the original JSON-array framing exactly.
- [ ] Build **opcode/message-handler registry** so adopters add handlers without editing core.
- [ ] Port `Types` enums: entity kinds, orientations, item ranks.
- [ ] `netclient` WebSocket transport that performs the handshake (HELLO → WELCOME).
- [ ] Golden-file tests against captured real traffic from the existing server.

**Exit criteria**: Go client completes handshake with the live server and decodes the
initial entity list / population messages correctly.

---

### Milestone 2 — World Model & Rendering
**Outcome**: Ebitengine client renders the world and reflects server state.

- [ ] Client `world` model: entities, players, mobs, items, NPCs keyed by id.
- [ ] Apply server messages to local model (spawn/despawn/move/health/equip/teleport).
- [ ] **Layered renderer**: terrain → entities (depth-sorted) → high tiles → UI → debug.
- [ ] Sprite sheet loader + animation system (port animation specs from original).
- [ ] Tile map renderer with viewport culling; smooth camera follow.
- [ ] Entity movement interpolation/prediction.
- [ ] **Render-layer registry** for downstream visual additions.
- [ ] **Client event hooks** wired (OnSpawn/OnDespawn/OnDeath/OnChat/OnZoneChange).

**Exit criteria**: Client visually matches server state for movement, spawns, and combat
animations in real time against the live server.

---

### Milestone 3 — Input, Pathfinding & Interaction
**Outcome**: Player can move, fight, and loot via the client.

- [ ] **Input provider interface**: keyboard+mouse (desktop), touch (mobile), gamepad.
- [ ] Click/tap-to-move with A* pathfinding on collision grid (port original behavior).
- [ ] Send MOVE/LOOTMOVE/ATTACK messages matching server expectations.
- [ ] Attack-on-click, loot-on-click, target highlighting, door/teleport interaction.
- [ ] Local prediction reconciled with server corrections.
- [ ] **Input scheme registry** for remapping and alternate control sets.

**Exit criteria**: Full move/fight/loot/equip loop works end-to-end against the live server.

---

### Milestone 4 — UI/HUD, Audio & Chat
**Outcome**: Feature-complete, polished client UX.

- [ ] **UI widget registry**: health bar, chat box, inventory, population counter,
      death screen, achievements notifications, name-entry screen.
- [ ] Chat send/receive against server CHAT messages.
- [ ] **Audio interface** + platform backends (Ebitengine `audio` for desktop/WASM,
      gomobile-compatible for mobile); SFX + music with mute toggle.
- [ ] Settings menu (audio, name entry, controls remap).
- [ ] Achievements parity with original client logic.

**Exit criteria**: A full play session is possible and feels equivalent to the original client.

---

### Milestone 5 — Platform Targets Hardening
**Outcome**: First-class builds on all platforms.

- [ ] **WASM**: optimized `wasm_exec.js` shell, gzip/brotli, loading screen, resize handling.
      WebSocket via the browser bridge.
- [ ] **Desktop**: window scaling, fullscreen, DPI, per-OS packaging.
- [ ] **Android**: gomobile build, touch tuned, back-button handling, lifecycle pause/resume.
- [ ] **iOS**: gomobile build, safe-area insets, app lifecycle.
- [ ] Platform `storage` adapter (browser localStorage bridge / file / mobile prefs).
- [ ] Network reconnect logic & offline detection per platform (reconnect to existing server).

**Exit criteria**: Verified manual playthrough on each of the 5 targets against the live server.

---

### Milestone 6 — Extensibility & Modding SDK
**Outcome**: Documented, demonstrable client extension story.

- [ ] Publish stable interfaces for all registries (semver-guarded).
- [ ] **Sample render-layer mod**: e.g., day/night tint overlay.
- [ ] **Sample UI widget mod**: custom HUD panel.
- [ ] **Sample message-handler mod**: handle a new server message gracefully.
- [ ] **Sample input scheme**: alternate control mapping.
- [ ] `docs/extending.md` cookbook for each client extension point.

**Exit criteria**: A third party can add a render layer + UI widget + input scheme without
touching core packages.

---

### Milestone 7 — Polish, Parity Audit & Release
**Outcome**: 1.0 release matching original feel.

- [ ] Side-by-side parity audit vs original client (movement speed, animation timing,
      camera, combat feedback, map rendering).
- [ ] Performance pass (allocations, draw calls, WASM bundle size).
- [ ] Integration tests with a recorded-traffic replay transport.
- [ ] Accessibility & i18n scaffolding (string table is an extension point).
- [ ] Docs: README, architecture overview, build/deploy guide, contribution guide.
- [ ] Tagged `v1.0.0`, CI release artifacts for all platforms.

**Exit criteria**: Tagged release; reproducible builds; parity checklist signed off.

---

## 4. Master Completion Checklist

- [ ] M0 — Bootstrap & assets
- [ ] M1 — Protocol parity with existing server
- [ ] M2 — World model & rendering
- [ ] M3 — Input, pathfinding & interaction
- [ ] M4 — UI / HUD / audio / chat
- [ ] M5 — Platform hardening (WASM/Desktop/Android/iOS)
- [ ] M6 — Extensibility & modding SDK
- [ ] M7 — Polish, parity audit, release

---

## 5. Risks & Mitigations

- **Protocol drift from server** — the server is fixed, so the client must match it exactly;
  golden-file + live-traffic capture tests in M1 prevent regressions.
- **WebSocket on WASM** — must route through the browser's native WS; isolate behind the
  `netclient` transport interface and test the JS bridge early.
- **Asset license scope** — confirm each asset is CC-BY-SA 3.0; retain attribution; document provenance.
- **gomobile audio/input quirks** — isolate behind interfaces early (M0/M4) to swap backends.
- **WASM bundle size** — TinyGo evaluation as optional path; lazy-load assets.
- **Extensibility vs performance** — registries resolved at init; hot paths avoid reflection.

---

## 6. Definition of Done (per task)

A checkbox is complete only when: code merged, unit/integration tests added & green on the
full platform CI matrix, public interfaces documented, and (where user-facing) verified on
at least desktop **and** WASM connecting to the existing server.