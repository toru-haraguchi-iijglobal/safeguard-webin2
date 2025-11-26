## Purpose

This repository is a small Go tool invoked by OISG-RemoteApp-Launcher. It locates an `Asset` from YAML or JSONL definitions and automates browser actions (login flows) using chromedp.

Key points for an AI coding assistant working here:

- CLI: `webin2` accepts `-jsonl`, `-yaml`, `-asset`, `-account`, `-password`, and the new `-lang` (browser language) flag. `-lang` overrides any `lang` value in the definition file.
- Config formats: YAML supports multiple documents (---) — see `sample_webin2.yaml`. JSONL uses one JSON object per line — see `sample_webin2.jsonl`.
- Actions flow: `Definition` → list of `Action` objects → `run()` converts each action to chromedp calls. Common actions: `navigate`, `click`, `account`, `password`, `sleep`.
- Logging: `logger.go` centralizes Debug/Info/Warn/Error/Fatal; `LogArgs` now logs CLI `lang` and passwords are redacted.

When adding features:
- Update `webin2structs.go` if you add new fields or action parameters.
- Add new action types in `webin2chromedp.go` and wire up logging (`LogActionStart` / `LogActionComplete`).
- Keep platform assumptions in mind — this code expects Windows paths for Edge by default.

Quick example (local run):
```powershell
.\webin2.exe -yaml sample_webin2.yaml -asset PaloAltoNetworksHUB -account user -password pass -lang ja
```

If anything here is missing, tell me which area to expand (examples, tests, or local debugging tips).
