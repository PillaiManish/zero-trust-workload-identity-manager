# Agents.md — Gap Report

**Round**: 1
**Date**: 2026-08-17
**Derived from**: pattern-analysis.md — all pattern groups; rca-summary.md gap type distribution

---

## Gap AG-1: Missing Domain Knowledge for Platform Security Policy Semantics

**Affected section**: Stage-Specific Agent Guidance → Repo-Assessment Stage Hints
**Class of missing information**: agents.md does not include domain knowledge about how platform security policies interact with operand security contexts — specifically the distinction between policies that *allow* configurations versus those that *inject* them.
**Severity**: Patchable

**Current state**: The agents.md file provides detailed guidance on controller patterns, API types, and reconciliation logic but lacks security policy semantic knowledge that repo-assessment and implementation agents need.

**Required addition**: Add domain knowledge section covering: (1) platform security policy enforcement semantics (allowlisting vs injection), (2) how operand pod specs must explicitly request configurations that the policy allows but does not inject, (3) the interaction between namespace-level allocations and pod-level declarations.

---

## Gap AG-2: Missing Agent Guidance for Multi-Architecture Image Validation

**Affected section**: Execution agent routing / Verification pairing
**Class of missing information**: No agent routing entry covers multi-architecture image validation as a verification step during operand image updates.
**Severity**: Patchable

**Current state**: Agent routing covers API, Controller, Manifests/Bindata, RBAC, OLM, Testing, and Docs agents but does not include image governance as a routing concern.

**Required addition**: Add verification pairing guidance: "Operand image version updates → pair with architecture manifest validation (verify all RELATED_IMAGE references resolve to multi-arch manifest lists for supported architectures)."

---

## Gap AG-3: Missing Test Strategy for Security Posture Verification

**Affected section**: Per-task testing during `/opsx-apply`
**Class of missing information**: The per-task testing table covers build verification, consistency checks, and unit tests but does not include security posture verification as a test strategy for tasks that modify privilege configurations.
**Severity**: Patchable

**Current state**: Task types map to verification strategies (go build, make verify, go test) but no entry exists for "security configuration changes → verify runtime security posture."

**Required addition**: Add test strategy entry: "Security context / privilege changes → verify actual UID/GID/capabilities match expected configuration (runtime assertion or dedicated security test)."

---

## Gap AG-4: Missing Controller Routing Rule for Reconciliation Safety

**Affected section**: Controller routing rules
**Class of missing information**: Routing rules distinguish between framework patterns (library-go vs controller-runtime) but do not include guidance on reconciliation safety patterns (conflict detection, retry strategy, idempotency verification).
**Severity**: Eval-only

**Current state**: Routing rules focus on which framework to use and API-before-controller ordering but do not guide agents on reconciliation safety requirements.

**Recommendation**: Consider adding routing guidance: "Tasks that implement read-modify-write reconciliation patterns must verify conflict handling (optimistic concurrency or SSA) — do not implement custom update logic when framework-provided patterns suffice."

---

## Gap AG-5: Missing Guidance for Incremental Security Hardening Anti-Pattern

**Affected section**: Common Mistakes (or new anti-patterns section)
**Class of missing information**: agents.md does not warn against incremental, ad-hoc security hardening across multiple PRs instead of comprehensive up-front security design.
**Severity**: Eval-only

**Current state**: Common mistakes cover implementation-level anti-patterns but not process-level anti-patterns about security design approach.

**Recommendation**: Consider adding: "Do NOT iterate security configurations across multiple PRs without a comprehensive security design document. Design all operand security requirements (SCC, RBAC, SecurityContext) holistically before implementation."
