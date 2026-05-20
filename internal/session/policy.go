package session

import "github.com/Shahzaibzah00r/zaibflow/internal/providers"

func RequiresClaudeSanitization(family providers.Family) bool {
	return family == providers.FamilyClaudeStrict
}
