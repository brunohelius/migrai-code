package models

const (
	ProviderMigrAI ModelProvider = "migrai"

	// MigrAI models served via LiteLLM proxy
	MigrAIClaudeOpus46   ModelID = "migrai/claude-opus-4.6"
	MigrAIClaudeSonnet46 ModelID = "migrai/claude-sonnet-4.6"
	MigrAIClaudeHaiku45  ModelID = "migrai/claude-haiku-4.5"
	MigrAIMiniMaxM27     ModelID = "migrai/minimax-m2.7"
	MigrAIMiniMaxM25     ModelID = "migrai/minimax-m2.5"
	MigrAIDeepSeekV32    ModelID = "migrai/deepseek-v3.2"
	MigrAIQwen35         ModelID = "migrai/qwen3.5"
	MigrAIGPT54          ModelID = "migrai/gpt-5.4"
	MigrAIGemini3Flash   ModelID = "migrai/gemini-3-flash"
	MigrAIKimiK25        ModelID = "migrai/kimi-k2.5"
	MigrAIStepFlashFree  ModelID = "migrai/step-3.5-flash-free"
)

var MigrAIModels = map[ModelID]Model{
	MigrAIClaudeOpus46: {
		ID:                  MigrAIClaudeOpus46,
		Name:                "Claude Opus 4.6",
		Provider:            ProviderMigrAI,
		APIModel:            "claude-opus-4.6",
		CostPer1MIn:         15.0,
		CostPer1MOut:        75.0,
		ContextWindow:       200_000,
		DefaultMaxTokens:    16384,
		SupportsAttachments: true,
	},
	MigrAIClaudeSonnet46: {
		ID:                  MigrAIClaudeSonnet46,
		Name:                "Claude Sonnet 4.6",
		Provider:            ProviderMigrAI,
		APIModel:            "claude-sonnet-4.6",
		CostPer1MIn:         3.0,
		CostPer1MOut:        15.0,
		ContextWindow:       200_000,
		DefaultMaxTokens:    16384,
		SupportsAttachments: true,
	},
	MigrAIClaudeHaiku45: {
		ID:                  MigrAIClaudeHaiku45,
		Name:                "Claude Haiku 4.5",
		Provider:            ProviderMigrAI,
		APIModel:            "claude-haiku-4.5",
		CostPer1MIn:         0.8,
		CostPer1MOut:        4.0,
		ContextWindow:       200_000,
		DefaultMaxTokens:    8192,
		SupportsAttachments: true,
	},
	MigrAIMiniMaxM27: {
		ID:                  MigrAIMiniMaxM27,
		Name:                "MiniMax M2.7",
		Provider:            ProviderMigrAI,
		APIModel:            "minimax-m2.7",
		CostPer1MIn:         1.0,
		CostPer1MOut:        5.0,
		ContextWindow:       128_000,
		DefaultMaxTokens:    8192,
	},
	MigrAIMiniMaxM25: {
		ID:                  MigrAIMiniMaxM25,
		Name:                "MiniMax M2.5",
		Provider:            ProviderMigrAI,
		APIModel:            "minimax-m2.5",
		CostPer1MIn:         1.0,
		CostPer1MOut:        5.0,
		ContextWindow:       128_000,
		DefaultMaxTokens:    8192,
	},
	MigrAIDeepSeekV32: {
		ID:                  MigrAIDeepSeekV32,
		Name:                "DeepSeek V3.2",
		Provider:            ProviderMigrAI,
		APIModel:            "deepseek-v3.2",
		CostPer1MIn:         0.5,
		CostPer1MOut:        2.0,
		ContextWindow:       128_000,
		DefaultMaxTokens:    8192,
	},
	MigrAIQwen35: {
		ID:                  MigrAIQwen35,
		Name:                "Qwen 3.5",
		Provider:            ProviderMigrAI,
		APIModel:            "qwen3.5",
		CostPer1MIn:         0.5,
		CostPer1MOut:        2.0,
		ContextWindow:       128_000,
		DefaultMaxTokens:    8192,
	},
	MigrAIGPT54: {
		ID:                  MigrAIGPT54,
		Name:                "GPT 5.4",
		Provider:            ProviderMigrAI,
		APIModel:            "gpt-5.4",
		CostPer1MIn:         5.0,
		CostPer1MOut:        15.0,
		ContextWindow:       128_000,
		DefaultMaxTokens:    16384,
		SupportsAttachments: true,
	},
	MigrAIGemini3Flash: {
		ID:                  MigrAIGemini3Flash,
		Name:                "Gemini 3 Flash",
		Provider:            ProviderMigrAI,
		APIModel:            "gemini-3-flash",
		CostPer1MIn:         0.1,
		CostPer1MOut:        0.4,
		ContextWindow:       1_000_000,
		DefaultMaxTokens:    8192,
	},
	MigrAIKimiK25: {
		ID:                  MigrAIKimiK25,
		Name:                "Kimi K2.5",
		Provider:            ProviderMigrAI,
		APIModel:            "kimi-k2.5",
		CostPer1MIn:         0.5,
		CostPer1MOut:        2.0,
		ContextWindow:       128_000,
		DefaultMaxTokens:    8192,
	},
	MigrAIStepFlashFree: {
		ID:                  MigrAIStepFlashFree,
		Name:                "Step 3.5 Flash (Free)",
		Provider:            ProviderMigrAI,
		APIModel:            "step-3.5-flash-free",
		CostPer1MIn:         0,
		CostPer1MOut:        0,
		ContextWindow:       128_000,
		DefaultMaxTokens:    8192,
	},
}
