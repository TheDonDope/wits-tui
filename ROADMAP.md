# 🥦 Wits Roadmap 🏗️

Instead of overly complicating the development process with GitHub Issues,
we will keep it simple and list features and planned changes in this file.

## 🚀 Tracking Features and Changes

Each feature or change is tracked with:
- **Status**: Planned | In Progress | Implemented (since `<commitSha or tag>`)
- **Relevant Commits**: `<commitSha1>, <commitSha2>, ...`
- **Description**: A brief explanation of the feature or change.

## Detailed Changelog

A detailed changelog with all commits can be found in the [CHANGELOG.md](./CHANGELOG.md).

---

## 📌 Planned Features

### 🔹Reading/Writing of Configuration & Settings

- **Status**: In Progress
- **Description**: Implement reading and writing of the app configuration and settings to a `.wits/settings.yml`
- **Tasks**:
  [ ] tbd
- **Relevant Commits**: tbd

### 🔹 Persistent Local Storage for Strains

- **Status**: In Progress
- **Description**: Store strains in a `.wits/strains.yml` file for persistence.
- **Tasks**:
  - [x] Check/Manage `.wits` folder creation
  - [x] Implement serializing/deserialzing to YAML
  - [ ] Tests for storage / YAML
- **Relevant Commits**: [47a13ae](https://github.com/TheDonDope/wits-tui/commit/47a13aef4390fdae6fcebb13c57ef01207bfecf1)

### 🔹 Terminal Form for Adding Strains

- **Status**: In Progress
- **Description**: Users can add new strains using a Bubble Tea-powered form.
- **Tasks**:
  - [ ] tbd
- **Relevant Commits**: tbd

---

## 🔧 Improvements & Refactoring

### 🔹 Debug Logging to Logfile

- **Status**: In Progress
- **Description**: see <https://github.com/charmbracelet/bubbletea/blob/main/README.md#logging-stuff>
- **Tasks**:
  - [ ] tbd
- **Relevant Commits**: tbd

### 🔹 Improve Error Handling in StrainForm Submission

- **Status**: In Progress
- **Description**: Ensure safe parsing of form inputs and prevent crashes from invalid data.
- **Tasks**:
  - [ ] tbd
- **Relevant Commits**: tbd

### 🔹 Refactor `strain_store.go`

- **Status**: Planned
- **Description**: The different `StrainStore` implementations (in-memory vs. yaml persistance) need to be handled better, e.g. in `AddStrain` there is a check `if sstr.Persistance == YMLFile` that does yml specific logic, but this if looks ugly and will get even worse once more implementations (database storage) will be added.
- **Tasks**:
  - [ ] (somehow) Remove the if check for Persistance
  - [ ] Make specific implementation behaviour less tedious to work with as a user
- **Relevant Commits**: tbd

---

## 📜 Notes

- Changes will be updated as development progresses.
- Feature status will be marked as **Implemented** once merged into `main`.
- Commits are referenced for traceability and rollback if needed.
