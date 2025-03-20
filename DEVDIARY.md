# Development Diary

## 2025-03-18

### CI Hardening

Pinned third-party GitHub Actions to specific commit SHAs to ensure deterministic builds and mitigate supply chain risks. Maintained compatibility with existing workflows while improving security posture.

- `c519bfd` chore(ci): pin 3rd party github actions to specific commit

### Notes for 2025-03-18

Nope

## 2025-03-17

### Testing Blitz & Release Automation

Added comprehensive test coverage for TUI components and strain management logic. Implemented Codacy integration for test reporting and quality metrics. Fixed release automation in Makefile, added versioned changelog entries through v0.4.1, and improved CI pipeline reliability. Updated documentation with badges and roadmap progress.

- `87a6465` fix(pkg-tui): wire up home view correctly
- `8075465` fix(pkg-tui): initialize strain editor properly
- `73db715` chore(deps): bump lipgloss to 1.1.0
- `65f5b81` test(pkg-storage): strain store tests
- `a9ae565` test(pkg-service): strain service tests
- `da5a8e3` test(pkg-tui): devices home model
- `56a6a66` test(pkg-tui): home model tests
- `53df9ac` test(pkg-tui): settings home model
- `2c83a30` test(pkg-tui): statistics model
- `7f84de4` test(pkg-tui): menu model tests
- `2c86440` chore(ci): upload coverage to Codacy
- `52c4ee1` docs: add codacy badge
- `aabff4c` chore(ci): fix bug report template
- `8db798a` docs: update roadmap
- `9621bef` chore(ci): upload coverage on build
- `d149c32` docs: fix formatting
- `29e0b52` chore: add release target
- `01e728d` docs: changelog v0.4.0
- `f456e72` fix: repair release target
- `2727ccc` docs: changelog v0.4.1

### Notes for 2025-03-17

Nope

## 2025-03-16

### Core Functionality & Observability

Implemented proper initialization for strain store persistence and TUI data binding. Added structured logging throughout service and storage layers. Integrated Cobra CLI framework and refined Makefile targets. Improved list rendering and menu navigation in TUI.

- `68f653b` fix(pkg-tui): trigger list loading
- `0562b72` feat: integrate cobra commands
- `eed3770` feat(pkg-tui): format strain list
- `5400544` fix(pkg): init strain store
- `63b3bf9` feat(pkg-storage): strain store logging
- `dc369f0` feat(pkg-service): strain service logging
- `05bb7d8` feat(cmd-wits): main cmd logginga

### Notes for 2025-03-16

Nope

## 2025-03-13

### Configuration & TUI Architecture

Added environment variable configuration support and debug logging initialization. Refactored TUI components to use proper Bubble Tea command patterns. Published v0.3.0 changelog and updated roadmap documentation.

- `43314a7` docs: roadmap update
- `4a5afbf` docs: changelog v0.3.0
- `fb098e4` feat: enable env config
- `5137f20` fix(pkg-tui): strain list update
- `ff8ddc3` feat(pkg-tui): tea.Cmd messaging

### Notes for 2025-03-13

Nope

## 2025-03-12

### TUI Foundation

Built core TUI components including home view model, menu navigation, and strain editor. Implemented fullscreen mode and proper application exit handling. Added Windows build target and improved documentation for application usage.

- `9278519` feat(pkg-tui): home view builder
- `432bc97` feat(pkg-tui): home view model
- `1b8ecff` feat(pkg-tui): add appliances
- `99da745` fix(pkg-tui): handle ctrl+c exit
- `0016bb1` feat(cmd-wits): fullscreen mode
- `47a13ae` fix(cmd-wits): command rename

### Notes for 2025-03-12

Nope

## 2025-03-11

### Initial Implementation

Established core architecture with file persistence and initial TUI rendering. Configured logging infrastructure and added basic strain management capabilities. Set up Makefile build system and published initial v0.1.0 changelog.

- `8128889` docs: changelog v0.1.0
- `62d5175` feat: file persistence
- `1b4bc06` fix(pkg-tui): non-emoji cursor
- `6f440f4` refac(cmd): rename tui to wits
- `30b4d4a` refac(pkg-storage): extract dir
- `0016bb1` feat(cmd): fullscreen mode

### Notes for 2025-03-11

Nope
