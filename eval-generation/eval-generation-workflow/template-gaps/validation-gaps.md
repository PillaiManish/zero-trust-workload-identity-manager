# Validation Template — Gap Report

**Round**: 1
**Date**: 2026-08-17
**Derived from**: pattern-analysis.md — cross-cutting themes A (incremental security hardening) and B (upstream image governance)

---

## Gap VL-1: Missing Completeness Pillar for Security Constraint Enumeration

**Affected section**: Completeness scoring rubric
**Class of missing information**: The validator does not include a completeness pillar that checks whether a spec enumerates security constraints for each operand/component when the feature involves privileged operations or security-sensitive configurations.
**Severity**: Patchable

**Current state**: The validator checks structural completeness (sections present, requirements testable, scope defined) but does not have a specific pillar for "security requirement completeness" when the spec describes components with heterogeneous privilege needs.

**Required addition**: Add a completeness pillar: "When the spec describes components requiring elevated privileges or custom security policies, each component's security requirements (privilege level, identity, filesystem permissions, network access) must be explicitly enumerated."

---

## Gap VL-2: Missing Quality Issue Type for Architecture-Specific Requirements

**Affected section**: Quality issue classification
**Class of missing information**: The validator does not classify missing architecture-support requirements as a quality issue when the spec describes deployment of container images across heterogeneous infrastructure.
**Severity**: Patchable

**Current state**: Quality issues cover ambiguity, untestability, and inconsistency but do not include "missing architecture constraint" as a detectable issue class.

**Required addition**: Add quality issue type: "architecture-support-gap" — triggered when a spec references container images deployed to infrastructure but does not specify architecture support requirements or validation criteria.

---

## Gap VL-3: Missing Scoring Posture for Reconciliation Safety Requirements

**Affected section**: Quality scoring posture
**Class of missing information**: The validator does not have calibration for detecting when specs that describe controller reconciliation lack conflict handling, retry, and idempotency requirements.
**Severity**: Eval-only

**Current state**: Scoring posture evaluates testability and completeness generically but does not specifically flag the absence of reconciliation safety requirements for controller-based features.

**Recommendation**: Consider adding a few-shot calibration example that demonstrates BLOCKED status when a spec describes controller reconciliation without specifying conflict detection, retry strategy, or idempotency guarantees.
