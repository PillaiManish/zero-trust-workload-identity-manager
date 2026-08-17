# Spec Template — Gap Report

**Round**: 1
**Date**: 2026-08-17
**Derived from**: pattern-analysis.md — reconciliation-safety, operand-security-configuration patterns

---

## Gap SP-1: Missing Edge Case Category for Security Policy Interactions

**Affected section**: Edge Cases guidance (within spec template quality self-check blocks)
**Class of missing information**: The spec template does not include guidance for documenting edge cases arising from interactions between component security configurations and platform-managed security policies.
**Severity**: Patchable

**Current state**: Edge case categories cover error states, boundary conditions, and integration failures generically but do not prompt spec authors to consider security policy interaction edge cases.

**Required addition**: Add to edge case guidance: "When components require custom security policies, document edge cases for: (1) platform policy conflicts with component requirements, (2) policy enforcement differences across platform versions, (3) policy injection vs explicit configuration semantics."

---

## Gap SP-2: Missing Functional Requirement Guidance for Multi-Architecture Deployments

**Affected section**: Functional Requirements guidance
**Class of missing information**: The spec template does not prompt spec authors to include architecture-support requirements when the feature deploys container images.
**Severity**: Patchable

**Current state**: Functional requirements guidance covers behavior, inputs, outputs, and constraints but does not include architecture-portability as a dimension.

**Required addition**: Add to functional requirements guidance: "When the feature deploys or references container images, specify which CPU architectures must be supported and how architecture availability is validated."

---

## Gap SP-3: Missing Assumption Category for Reconciliation Semantics

**Affected section**: Assumptions section guidance
**Class of missing information**: The spec template does not prompt spec authors to explicitly state assumptions about reconciliation behavior (idempotency, conflict handling, partial failure) when the feature involves controller-managed resources.
**Severity**: Patchable

**Current state**: Assumptions guidance covers dependencies, environment, and constraints but does not include reconciliation semantics as a category that must be explicitly resolved.

**Required addition**: Add assumption category: "For features involving controller-managed resources, explicitly state assumptions about: (1) reconciliation frequency, (2) conflict resolution strategy, (3) partial failure behavior, (4) ordering guarantees across resource types."

---

## Gap SP-4: Missing Testability Guidance for Security Configuration Verification

**Affected section**: Given/When/Then acceptance criteria guidance
**Class of missing information**: The spec template does not include guidance for writing testable acceptance criteria that verify runtime security posture (not just configuration correctness).
**Severity**: Eval-only

**Current state**: GWT guidance covers functional behavior but does not specifically address how to make security configurations testable at runtime.

**Recommendation**: Consider adding guidance: "For security-sensitive features, acceptance criteria should verify runtime state (actual privilege level, actual group membership) not just configuration state (YAML correctness)."
