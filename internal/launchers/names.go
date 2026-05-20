package launchers

func LauncherName(profile string) string {
	switch profile {
	case "ollama":
		return "zf-local"
	default:
		return "zf-" + profile
	}
}

func LauncherCommand(profile string) []string {
	switch {
	case profile == "ollama":
		return []string{"run", "ollama"}
	case len(profile) > 3 && profile[:3] == "or-":
		return []string{"run", "openrouter", profile[3:]}
	default:
		return []string{"run", profile}
	}
}
