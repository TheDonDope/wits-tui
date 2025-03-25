<a name="unreleased"></a>
## [Unreleased]


<a name="v0.6.0"></a>
## [v0.6.0] - 2025-03-25
### Chore
- update rendered tapes
- set font options for tapes

### Docs
- render env example source in readme
- add dependecy logos to readme
- add initial dev diary

### Feat
- make the main menu more beautiful
- add tape for dev-diary
- add dev-diary script

### Fix
- use correct font family in vhs tapes
- remove unnecessary variable interpolation


<a name="v0.5.0"></a>
## [v0.5.0] - 2025-03-20
### Chore
- **ci:** pin 3rd party github actions to specific commit
- **vhs:** update rendered tapes for docs

### Docs
- update changelog for v0.5.0
- add used tech to readme

### Feat
- add charmbracelt/vhs to gifs for documentation


<a name="v0.4.1"></a>
## [v0.4.1] - 2025-03-17
### Docs
- update changelog for v0.4.1

### Fix
- repair release target in makefile


<a name="v0.4.0"></a>
## [v0.4.0] - 2025-03-17
### Chore
- add release target to makefile
- clean up makefile
- **ci:** upload test coverage results to codacy
- **ci:** fix bug report template formatting
- **ci:** upload test coverage results on github build
- **deps:** bump github.com/charmbracelet/lipgloss from 1.0.0 to 1.1.0

### Docs
- add changelog for v0.4.0
- fix formatting
- update roadmap
- add codacy badge
- update roadmap
- add changelog for v0.3.0

### Feat
- integrate cobra commands
- configure debug logging
- enable configuration through environment variables
- **cmd-wits:** add logging to main cmd
- **pkg:** add stringer methods for strain and store
- **pkg-service:** add logging to strain service
- **pkg-storage:** add logging to strain store
- **pkg-tui:** format strain list item more nicely

### Fix
- **pkg:** correctly initialize strain store and load from yml file
- **pkg-tui:** initialize strain editor properly
- **pkg-tui:** wire up home view correctly
- **pkg-tui:** trigger list loading from menus
- **pkg-tui:** properly update strain list

### Test
- **pkg-service:** add tests for strain service
- **pkg-storage:** add tests for strain store
- **pkg-tui:** add menu model tests
- **pkg-tui:** add statistics home model tests
- **pkg-tui:** add settings home model tests
- **pkg-tui:** add devices home model tests
- **pkg-tui:** add home model tests


<a name="v0.3.0"></a>
## [v0.3.0] - 2025-03-13
### Chore
- **build:** add windows build target to makefile

### Docs
- add roadmap and update readme
- update application run instructions
- add changelog for v0.2.0

### Feat
- **cmd-wits:** run wits in fullscreen
- **pkg-tui:** wire up strain add action
- **pkg-tui:** separate side effects into tea.Cmds
- **pkg-tui:** add mnemonics for appliance actions
- **pkg-tui:** render appliance titles
- **pkg-tui:** render mnemonics with marked text on menu items
- **pkg-tui:** add appliances
- **pkg-tui:** add home view model
- **pkg-tui:** add home view builder
- **pkg-tui:** sort options for strain editor selects alphabetically

### Fix
- **cmd-wits:** remove wrong ignore and re-add wits command
- **pgk-tui:** update documentations
- **pkg-tui:** initialize appliances properly
- **pkg-tui:** handle ctrl+c program exit
- **pkg-tui:** use non-emoji cursor
- **pkg-tui:** use correct cursor emojis

### Refac
- **cmd-wits:** rename command from tui to wits
- **pkg-storage:** extract wits directory
- **pkg-tui:** clean up strains home model
- **pkg-tui:** use tea.Cmd messaging for side effects
- **pkg-tui:** drop the term appliances and instead use home model
- **pkg-tui:** rename HomeView to HomeModel to closer align to bubbletea terminology
- **pkg-tui:** privatize menu model properties


<a name="v0.2.0"></a>
## [v0.2.0] - 2025-03-11
### Chore
- **deps:** bump go version to v1.24.1

### Docs
- add changelog for v0.1.0

### Feat
- add file persistance


<a name="v0.1.0"></a>
## v0.1.0 - 2025-03-11
### Feat
- initial commit


[Unreleased]: https://github.com/TheDonDope/wits-tui/compare/v0.6.0...HEAD
[v0.6.0]: https://github.com/TheDonDope/wits-tui/compare/v0.5.0...v0.6.0
[v0.5.0]: https://github.com/TheDonDope/wits-tui/compare/v0.4.1...v0.5.0
[v0.4.1]: https://github.com/TheDonDope/wits-tui/compare/v0.4.0...v0.4.1
[v0.4.0]: https://github.com/TheDonDope/wits-tui/compare/v0.3.0...v0.4.0
[v0.3.0]: https://github.com/TheDonDope/wits-tui/compare/v0.2.0...v0.3.0
[v0.2.0]: https://github.com/TheDonDope/wits-tui/compare/v0.1.0...v0.2.0
