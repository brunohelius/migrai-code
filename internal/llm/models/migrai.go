package models

const (
	ProviderMigrAI ModelProvider = "migrai"

	// MigrAI models served via LiteLLM proxy
	MigrAIClaudeOpus46   ModelID = "migrai/claude-opus-4.6"
	MigrAIClaudeSonnet46 ModelID = "migrai/claude-sonnet-4.6"
	MigrAIMiniMaxM21     ModelID = "migrai/minimax-m2.1"
)

var MigrAIModels = map[ModelID]Model{
	MigrAIClaudeOpus46: {
		ID:                  MigrAIClaudeOpus46,
		Name:                "MigrAI - Claude Opus 4.6",
		Provider:            ProviderMigrAI,
		APIModel:            "claude-opus-4-6",
		CostPer1MIn:         15.0,
		CostPer1MOut:        75.0,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0,
		ContextWindow:       200_000,
		DefaultMaxTokens:    16384,
		SupportsAttachments: true,
	},
	MigrAIClaudeSonnet46: {
		ID:                  MigrAIClaudeSonnet46,
		Name:                "MigrAI - Claude Sonnet 4.6",
		Provider:            ProviderMigrAI,
		APIModel:            "claude-sonnet-4-6",
		CostPer1MIn:         3.0,
		CostPer1MOut:        15.0,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0,
		ContextWindow:       200_000,
		DefaultMaxTokens:    16384,
		SupportsAttachments: true,
	},
	MigrAIMiniMaxM21: {
		ID:                  MigrAIMiniMaxM21,
		Name:                "MigrAI - MiniMax M2.1",
		Provider:            ProviderMigrAI,
		APIModel:            "minimax-m2.1",
		CostPer1MIn:         1.0,
		CostPer1MOut:        5.0,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0,
		ContextWindow:       128_000,
		DefaultMaxTokens:    8192,
		SupportsAttachments: false,
	},
}
