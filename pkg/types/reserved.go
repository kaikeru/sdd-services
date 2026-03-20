package types

// ReservedFields holds the ktsu_* output fields that the orchestrator
// interprets for security signals, quality signals, flow control, and observability.
// These fields are reserved and may not be set by user pipeline steps.
type ReservedFields struct {
	// ktsu_injection_attempt signals that untrusted input attempted to hijack agent behavior.
	// Orchestrator action: fail the entire run immediately.
	InjectionAttempt bool `json:"ktsu_injection_attempt,omitempty"`
	// ktsu_untrusted_content signals suspicious content that does not rise to a clear injection attempt.
	// Orchestrator action: fail the step.
	UntrustedContent bool `json:"ktsu_untrusted_content,omitempty"`
	// ktsu_confidence is the agent's self-assessed confidence in its output (0.0–1.0).
	// Orchestrator action: fail the step if below the declared confidence_threshold.
	Confidence float64 `json:"ktsu_confidence,omitempty"`
	// ktsu_low_quality signals that the agent could not produce a reliable output.
	// Orchestrator action: fail the step.
	LowQuality bool `json:"ktsu_low_quality,omitempty"`
	// ktsu_skip_reason signals that there is legitimately nothing to do.
	// Orchestrator action: mark the step skipped with the provided reason.
	SkipReason string `json:"ktsu_skip_reason,omitempty"`
	// ktsu_needs_human signals that the case exceeds the agent's confidence or authorization.
	// Orchestrator action: fail the run with error code needs_human_review.
	NeedsHuman bool `json:"ktsu_needs_human,omitempty"`
	// ktsu_flags are soft labels for observability. No pipeline effect.
	Flags []string `json:"ktsu_flags,omitempty"`
	// ktsu_rationale is the agent's explanation of its reasoning. No pipeline effect.
	Rationale string `json:"ktsu_rationale,omitempty"`
}

// ReservedPrefix is the prefix for all reserved output fields
const ReservedPrefix = "ktsu_"
