# TODO

Implementation tracking for all `pim` commands.

## Top-level Commands

- [ ] **pim init** — Initialize a fresh workspace manifest and select a Python version.
- [ ] **pim add** — Add a new package dependency to the active project context.
- [ ] **pim remove** — Remove a package dependency from the project configuration.
- [ ] **pim sync** — Align the physical cache and environment state with the manifest file.
- [ ] **pim tidy** — Scan project source files to prune or add dependencies automatically.
- [ ] **pim clear** — Clear the local workspace configuration states and target environment links.
- [ ] **pim run** — Execute a Python script using dynamic, environmentless injection.
- [ ] **pim tool** — Execute an installed package command-line utility directly.
- [ ] **pim cache** — Audit and manage the storage states of global package caches.

## `py` Subcommands

- [ ] **pim py** — Manage standalone Python interpreter installations.
- [ ] **pim py search** — Search for available remote CPython runtime versions.
- [ ] **pim py versions** — List all Python interpreter versions downloaded locally.
- [ ] **pim py use** — Switch the current project workspace to a specific Python version.
- [ ] **pim py download** — Download and cache a portable standalone CPython runtime version.
- [ ] **pim py status** — Display full configuration diagnostics for the active runtime.

## `pkg` Subcommands

- [ ] **pim pkg** — Inspect and query PyPI package metadata globally.
- [ ] **pim pkg search** — Perform a rapid local fuzzy search against all PyPI package names.
- [ ] **pim pkg info** — Display comprehensive metadata and caching status for a package.
