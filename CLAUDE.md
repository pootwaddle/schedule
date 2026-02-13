# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

**schedule** is a Go-based job scheduler that replaces/supplements Windows Task Scheduler for automating batch jobs on the MOE server. It runs as a long-running service, executing batch files at configured times.

**Module**: `github.com/pootwaddle/schedule`
**Type**: Executable (service)

## Dependencies

- `github.com/carlescere/scheduler` - Job scheduling library
- `github.com/pootwaddle/slogger` - Structured logging

## Scheduled Jobs

### Daily Jobs
| Time | Job | Batch File | Purpose |
|------|-----|-----------|---------|
| 00:00:01 | RotateLog | (internal) | Log file rotation |
| 00:02:02 | DelOldLogs | `DELAGE.BAT` | Delete old log files |
| 00:03:03 | DelAge | `delage.bat` | File retention cleanup |
| 00:05:05 | BirdBuddy | `birdbuddy.bat` | Bird Buddy data processing |
| 00:05:15 | DailyTemplates | `daytmpl.bat` | Generate daily markdown files |
| 00:23:23 | Backup | `backup_c_drive_to_tech1.bat` | C: drive backup |
| 01:05:10 | WebbyStats | `webby.bat` | Web statistics processing |
| 03:33:33 | GetFit | `getfit.bat` | Fitness data processing |
| 03:33:35 | Fortune | `FORTUN.BAT` | Fortune/joke generation |
| 11:59:55 | LogSumm | `logsumm.bat` | Email log summary (noon) |
| 23:59:55 | LogSumm | `logsumm.bat` | Email log summary (midnight) |

### Weekly Jobs
| Schedule | Job | Batch File | Purpose |
|----------|-----|-----------|---------|
| Monday 06:20 | Reserves | `RESERVES.BAT` | Fire department reserves |
| Friday 03:33 | Gem | `gem.bat` | Gem processing |

### Recurring Jobs
| Interval | Job | Batch File | Purpose |
|----------|-----|-----------|---------|
| 2 min | CleanGrey | `grey.bat` | Greylist cleanup |
| 5 min | LogParse | `logparse.bat` | SMTP log parsing |
| 7 min | SpamParse | `spamparse.bat` | Spam analysis |
| 9 min | Heartbeat | (internal) | Liveness log message |

## Architecture

Single-file application (`schedule.go`). Each job is a closure that calls `exec.Command("CMD", "/C", "C:\\AUTOJOB\\<script>.bat")`. The scheduler library handles cron-like timing. The main goroutine blocks with `select{}`.

### Log Rotation
Daily log rotation via `rotateLogging()` function that detects date changes and calls `slogger.ReloadLogger()`.

## Build Commands

```powershell
go build
```

## Important Notes

- Stop `schedule.exe` on MOE before deploying new version.
- All batch files are at `C:\AUTOJOB\` on the MOE server.
- Runs indefinitely - blocks with `select{}` after scheduling all jobs.
