# Plan Template — Gap Report

**Round**: 1
**Date**: 2026-08-17
**Derived from**: pattern-analysis.md pattern groups: reconciliation-safety, operand-security-configuration

---

## Gap PL-1: Missing Reconciliation Strategy Specification Requirement

**Affected section**: §5 Implementation phases — Phase template → Verification hooks
**Class of missing information**: The template does not require that plans specify the reconciliation strategy (imperative apply, SSA, create-or-update) and its conflict handling behavior for each workload type managed by the operator.
**Severity**: Patchable
**Pattern group**: reconciliation-safety

**Current state**: The phase template requires "Verification hooks" but does not mandate documenting how the reconciliation handles concurrent modifications, stale state, or version conflicts.

**Required addition**: When the plan involves controller reconciliation of multiple workload types (StatefulSet, DaemonSet, Deployment), the plan must specify: (1) the apply strategy per workload type, (2) conflict detection and retry behavior, (3) idempotency guarantees, and (4) partial failure handling.

---

## Gap PL-2: Missing Workload-Type-Specific Update Strategy Requirement

**Affected section**: §5 Implementation phases — Phase template
**Class of missing information**: The template does not require that plans differentiate update strategies by Kubernetes workload type when the operator manages multiple workload kinds.
**Severity**: Patchable
**Pattern group**: reconciliation-safety

**Current state**: Phases document target files and dependencies but do not mandate specifying how each workload type (StatefulSet rolling update, DaemonSet OnDelete vs RollingUpdate, Deployment rollout) handles spec changes.

**Required addition**: For operators managing multiple workload types, the plan must include a workload update strategy table in §5 specifying the update mechanism and its constraints for each managed workload kind.

---

## Gap PL-3: Missing Security Posture Verification in Phase Hooks

**Affected section**: §5 Implementation phases — Phase template → Verification hooks
**Class of missing information**: Plans do not require post-implementation verification that actual runtime security posture matches the designed security configuration.
**Severity**: Eval-only
**Pattern group**: operand-security-configuration

**Current state**: Verification hooks cover unit/integration/e2e/manual categories but do not explicitly require security posture assertions (e.g., verifying that running pods actually have the expected security context).

**Recommendation**: For phases that modify security configurations, verification hooks should include runtime security posture verification (checking actual container UID/GID, capabilities, volume permissions match design).

---

## Gap PL-4: Missing Non-Goals for Rejected Architectural Alternatives

**Affected section**: §1 Architectural strategy
**Class of missing information**: The template mentions "non-goals" only implicitly through the open questions mechanism. It does not require explicit documentation of rejected architectural alternatives and their rationale.
**Severity**: Patchable
**Pattern group**: reconciliation-safety

**Current state**: §1 focuses on the chosen strategy but does not mandate documenting WHY alternative approaches were rejected.

**Required addition**: §1 should include a "Rejected alternatives" subsection when multiple architectural approaches exist, documenting what was considered, why it was rejected, and what conditions might warrant revisiting.

---

## Gap PL-5: Missing Forward-Compatibility Planning Requirement

**Affected section**: §7 Risks, migrations, and operational follow-ups
**Class of missing information**: The template asks about upgrade/migration and compatibility but does not require explicit planning for forward-compatibility with platform evolution.
**Severity**: Eval-only
**Pattern group**: operand-security-configuration

**Current state**: §7 covers risks from upstream API drift but not from platform-level changes that affect operator behavior.

**Recommendation**: When the operator interacts with platform-managed security policies or APIs that evolve across platform versions, the plan should document the forward-compatibility strategy.
