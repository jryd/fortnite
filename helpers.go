package fortnite

import (
	"fmt"
	"math"
	"strings"
)

func processBRStats(stats RawBRStatsResponse, account User, platform string) FormattedBRStats {
	totalTime := 0.00

	var results FormattedBRStats

	for _, stat := range stats {

		//Solo, Duos, or Squads
		var mode string

		if strings.Contains(stat.Name, "_p2") {
			mode = "solo"
		} else if strings.Contains(stat.Name, "_p10") {
			mode = "duo"
		} else {
			mode = "squad"
		}

		//Which type of result is this element
		if strings.Contains(stat.Name, fmt.Sprintf("placetop1_%v", platform)) {
			if mode == "solo" {
				results.Group.Solo.Wins = stat.Value
			} else if mode == "duo" {
				results.Group.Duo.Wins = stat.Value
			} else {
				results.Group.Squad.Wins = stat.Value
			}
		} else if strings.Contains(stat.Name, fmt.Sprintf("placetop3_%v", platform)) {
			if mode == "solo" {
				results.Group.Solo.Top3 = stat.Value
			} else if mode == "duo" {
				results.Group.Duo.Top3 = stat.Value
			} else {
				results.Group.Squad.Top3 = stat.Value
			}
		} else if strings.Contains(stat.Name, fmt.Sprintf("placetop5_%v", platform)) {
			if mode == "solo" {
				results.Group.Solo.Top5 = stat.Value
			} else if mode == "duo" {
				results.Group.Duo.Top5 = stat.Value
			} else {
				results.Group.Squad.Top5 = stat.Value
			}
		} else if strings.Contains(stat.Name, fmt.Sprintf("placetop6_%v", platform)) {
			if mode == "solo" {
				results.Group.Solo.Top6 = stat.Value
			} else if mode == "duo" {
				results.Group.Duo.Top6 = stat.Value
			} else {
				results.Group.Squad.Top6 = stat.Value
			}
		} else if strings.Contains(stat.Name, fmt.Sprintf("placetop10_%v", platform)) {
			if mode == "solo" {
				results.Group.Solo.Top10 = stat.Value
			} else if mode == "duo" {
				results.Group.Duo.Top10 = stat.Value
			} else {
				results.Group.Squad.Top10 = stat.Value
			}
		} else if strings.Contains(stat.Name, fmt.Sprintf("placetop12_%v", platform)) {
			if mode == "solo" {
				results.Group.Solo.Top12 = stat.Value
			} else if mode == "duo" {
				results.Group.Duo.Top12 = stat.Value
			} else {
				results.Group.Squad.Top12 = stat.Value
			}
		} else if strings.Contains(stat.Name, fmt.Sprintf("placetop25_%v", platform)) {
			if mode == "solo" {
				results.Group.Solo.Top25 = stat.Value
			} else if mode == "duo" {
				results.Group.Duo.Top25 = stat.Value
			} else {
				results.Group.Squad.Top25 = stat.Value
			}
		} else if strings.Contains(stat.Name, fmt.Sprintf("matchesplayed_%v", platform)) {
			if mode == "solo" {
				results.Group.Solo.Matches = stat.Value
			} else if mode == "duo" {
				results.Group.Duo.Matches = stat.Value
			} else {
				results.Group.Squad.Matches = stat.Value
			}
		} else if strings.Contains(stat.Name, fmt.Sprintf("kills_%v", platform)) {
			if mode == "solo" {
				results.Group.Solo.Kills = stat.Value
			} else if mode == "duo" {
				results.Group.Duo.Kills = stat.Value
			} else {
				results.Group.Squad.Kills = stat.Value
			}
		} else if strings.Contains(stat.Name, fmt.Sprintf("score_%v", platform)) {
			if mode == "solo" {
				results.Group.Solo.Score = stat.Value
			} else if mode == "duo" {
				results.Group.Duo.Score = stat.Value
			} else {
				results.Group.Squad.Score = stat.Value
			}
		} else if strings.Contains(stat.Name, fmt.Sprintf("minutesplayed_%v", platform)) {
			if mode == "solo" {
				results.Group.Solo.TimePlayed = stat.Value
			} else if mode == "duo" {
				results.Group.Duo.TimePlayed = stat.Value
			} else {
				results.Group.Squad.TimePlayed = stat.Value
			}

			totalTime += stat.Value
		}
	}

	results.Group.Solo.KdRatio = math.Round(results.Group.Solo.Kills/(results.Group.Solo.Matches-results.Group.Solo.Wins)*100) / 100
	results.Group.Duo.KdRatio = math.Round(results.Group.Duo.Kills/(results.Group.Duo.Matches-results.Group.Duo.Wins)*100) / 100
	results.Group.Squad.KdRatio = math.Round(results.Group.Squad.Kills/(results.Group.Squad.Matches-results.Group.Squad.Wins)*100) / 100

	results.Group.Solo.WinPercentage = math.Round((results.Group.Solo.Wins/results.Group.Solo.Matches)*100) / 100
	results.Group.Duo.WinPercentage = math.Round((results.Group.Duo.Wins/results.Group.Duo.Matches)*100) / 100
	results.Group.Squad.WinPercentage = math.Round((results.Group.Squad.Wins/results.Group.Squad.Matches)*100) / 100

	results.Group.Solo.KillsPerMin = math.Round(results.Group.Solo.Kills/results.Group.Solo.TimePlayed*100) / 100
	results.Group.Duo.KillsPerMin = math.Round(results.Group.Duo.Kills/results.Group.Duo.TimePlayed*100) / 100
	results.Group.Squad.KillsPerMin = math.Round(results.Group.Squad.Kills/results.Group.Squad.TimePlayed*100) / 100

	results.Group.Solo.TimePlayedFormatted = formatTimeString(results.Group.Solo.TimePlayed)
	results.Group.Duo.TimePlayedFormatted = formatTimeString(results.Group.Duo.TimePlayed)
	results.Group.Squad.TimePlayedFormatted = formatTimeString(results.Group.Squad.TimePlayed)

	results.Group.Solo.KillsPerMatch = math.Round(results.Group.Solo.Kills/results.Group.Solo.Matches*100) / 100
	results.Group.Duo.KillsPerMatch = math.Round(results.Group.Duo.Kills/results.Group.Duo.Matches*100) / 100
	results.Group.Squad.KillsPerMatch = math.Round(results.Group.Squad.Kills/results.Group.Squad.Matches*100) / 100

	// <------------------------------------------------------------------->

	//Calculate lifetimeStats
	results.LifetimeStats.Wins = results.Group.Solo.Wins + results.Group.Duo.Wins + results.Group.Squad.Wins
	results.LifetimeStats.Top3 = results.Group.Solo.Top3 + results.Group.Duo.Top3 + results.Group.Squad.Top3
	results.LifetimeStats.Top5 = results.Group.Solo.Top5 + results.Group.Duo.Top5 + results.Group.Squad.Top5
	results.LifetimeStats.Top6 = results.Group.Solo.Top6 + results.Group.Duo.Top6 + results.Group.Squad.Top6
	results.LifetimeStats.Top10 = results.Group.Solo.Top10 + results.Group.Duo.Top10 + results.Group.Squad.Top10
	results.LifetimeStats.Top12 = results.Group.Solo.Top12 + results.Group.Duo.Top12 + results.Group.Squad.Top12
	results.LifetimeStats.Top25 = results.Group.Solo.Top25 + results.Group.Duo.Top25 + results.Group.Squad.Top25
	results.LifetimeStats.Matches = results.Group.Solo.Matches + results.Group.Duo.Matches + results.Group.Squad.Matches
	results.LifetimeStats.Kills = results.Group.Solo.Kills + results.Group.Duo.Kills + results.Group.Squad.Kills
	results.LifetimeStats.Score = results.Group.Solo.Score + results.Group.Duo.Score + results.Group.Squad.Score

	results.LifetimeStats.TimePlayed = totalTime

	results.LifetimeStats.KdRatio = math.Round(results.LifetimeStats.Kills/(results.LifetimeStats.Matches-results.LifetimeStats.Wins)*100) / 100
	results.LifetimeStats.WinPercentage = math.Round((results.LifetimeStats.Wins/results.LifetimeStats.Matches)*100) / 100
	results.LifetimeStats.TimePlayedFormatted = formatTimeString(results.LifetimeStats.TimePlayed)
	results.LifetimeStats.KillsPerMin = math.Round(results.LifetimeStats.Kills/results.LifetimeStats.TimePlayed*100) / 100
	results.LifetimeStats.KillsPerMatch = math.Round(results.LifetimeStats.Kills/results.LifetimeStats.Matches*100) / 100

	results.LifetimeStats.Wins = results.Group.Solo.Wins + results.Group.Duo.Wins + results.Group.Squad.Wins

	results.Info.AccountID = account.ID
	results.Info.Username = account.DisplayName
	results.Info.Platform = platform

	return results
}

func formatTimeString(time float64) string {
	result := ""
	days := math.Floor(time / 24 / 60)
	hours := math.Floor(math.Mod((time / 60), 24))
	mins := math.Mod(time, 60)

	if days > 0 {
		result += fmt.Sprintf("%vd ", days)
	}
	if hours > 0 {
		result += fmt.Sprintf("%vh ", hours)
	}
	if mins > 0 {
		result += fmt.Sprintf("%vm", mins)
	} else {
		result += "0m "
	}

	return result
}
