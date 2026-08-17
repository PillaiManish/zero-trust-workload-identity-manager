# Tasks Template — Gap Report

**Round**: 1
**Date**: 2026-08-17
**Derived from**: pattern-analysis.md pattern groups: multi-arch-support, operand-security-configuration, test-coverage

---

## Gap TK-1: Missing Multi-Architecture Validation Gate in Task Decomposition

**Affected section**: §4 Task Specifications — Acceptance criteria
**Class of missing information**: The template does not require that tasks involving external container image references include a multi-architecture validation gate as an acceptance criterion.
**Severity**: Patchable
**Pattern group**: multi-arch-support

**Current state**: Task acceptance criteria trace to spec requirements and include test commands, but do not mandate architecture-compatibility verification for image-referencing tasks.

**Required addition**: Tasks that introduce or modify external container image references (operand images, init containers, sidecars) must include an acceptance criterion verifying that all referenced images are available as multi-architecture manifest lists for all supported platform architectures.

---

## Gap TK-2: Missing Security Posture Verification Pairing

**Affected section**: §4 Task Specifications — Acceptance criteria; Core responsibility #4 (unit test co-generation)
**Class of missing information**: The template mandates unit test co-generation for Go tasks but does not require security posture verification as a specific acceptance criterion class for tasks modifying security configurations.
**Severity**: Patchable
**Pattern group**: operand-security-configuration

**Current state**: Unit test co-generation is mandatory but generic — it does not distinguish tasks that modify security configurations from other tasks.

**Required addition**: Tasks that modify security-related configurations (security policies, privilege levels, RBAC permissions, security contexts) must include acceptance criteria that verify the security configuration is correctly applied at runtime, not just that the code compiles.

---

## Gap TK-3: Missing Dependency Between Security Configuration and Verification Tasks

**Affected section**: §1 Task Dependency Graph; §3 Task Execution Manifest — Depends On
**Class of missing information**: The template does not require that security configuration tasks have explicit dependency links to verification tasks that assert the configuration's runtime behavior.
**Severity**: Eval-only
**Pattern group**: operand-security-configuration

**Current state**: Dependencies are based on file-level concerns and phase ordering, not on verification completeness requirements.

**Recommendation**: When a task modifies security configurations, the task graph should include an explicit verification dependency — the implementation is not "done" until runtime security posture is verified.

---

## Gap TK-4: Missing Operand-Level Task Decomposition Guidance

**Affected section**: Core responsibility #1 (Granular decomposition)
**Class of missing information**: The template decomposes tasks at file/package granularity but does not provide guidance for decomposing tasks at operand-component granularity when multiple operands have independent concerns.
**Severity**: Patchable
**Pattern group**: operand-security-configuration

**Current state**: Decomposition guidance focuses on file/package boundaries and agent routing, not on logical component boundaries.

**Required addition**: When the operator manages multiple independent operand components, task decomposition should separate concerns by operand where each operand has distinct security requirements, configuration surfaces, or lifecycle dependencies. This prevents combining unrelated operand changes into single tasks where one operand's issue masks another's.

---

## Gap TK-5: Missing Image Update Task Pattern

**Affected section**: §4 Task Specifications — Implementation notes
**Class of missing information**: The template does not provide a standard task pattern for operand image version updates, which have unique cascade requirements (image digest pinning, multi-arch verification, CSV relatedImages update, e2e smoke test).
**Severity**: Deferred
**Pattern group**: multi-arch-support

**Current state**: Image updates are treated like any other code change without specific cascade guidance.

**Recommendation**: Consider adding a standard task pattern for image version updates that includes architecture verification, digest pinning, and downstream manifest updates as mandatory sub-steps.
