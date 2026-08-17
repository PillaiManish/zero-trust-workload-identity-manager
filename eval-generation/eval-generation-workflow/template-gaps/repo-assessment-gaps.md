# Repo-Assessment Template — Gap Report

**Round**: 1
**Date**: 2026-08-17
**Derived from**: pattern-analysis.md pattern groups: operand-security-configuration, multi-arch-support

---

## Gap RA-1: Missing Security Constraint Matrix Requirement

**Affected section**: §10.1 Security Context & Permissions
**Class of missing information**: The template does not require a structured matrix mapping each operand component to its security requirements (privilege level, UID/GID, filesystem permissions, capabilities, platform security policy interactions).
**Severity**: Patchable
**Pattern group**: operand-security-configuration

**Current state**: §10.1 asks about SCC constraints and pod security context patterns generically, but does not mandate a per-operand enumeration when multiple operands have different privilege requirements.

**Required addition**: When the operator manages multiple operands with heterogeneous privilege needs, §10.1 must require a per-operand security constraint matrix documenting: security policy type (SCC/PSP/PSS), required UID/GID, supplemental groups, capabilities, volume mount permissions, and platform policy interaction behavior (e.g., injection vs explicit request semantics).

---

## Gap RA-2: Missing Architecture Support Matrix Requirement

**Affected section**: §10.4 Build & Compliance Constraints
**Class of missing information**: The template mentions "multi-arch support: build matrix, Dockerfile patterns" but does not require validation that all referenced container images (especially operand images not built by this repo) are available as multi-architecture manifests.
**Severity**: Patchable
**Pattern group**: multi-arch-support

**Current state**: §10.4 covers build-time multi-arch concerns but not runtime image reference validation for externally-sourced operand images.

**Required addition**: When the operator references external container images (e.g., via environment variables or CSV entries), the assessment must enumerate all image references and document their architecture availability. Flag any single-architecture images as risks.

---

## Gap RA-3: Missing RBAC Minimum-Privilege Audit Requirement

**Affected section**: §10.1 Security Context & Permissions (or §6 Architectural Guardrails — Security)
**Class of missing information**: The template does not require a per-controller RBAC audit mapping each controller's actual API operations to its declared permissions, identifying over-permissioned roles.
**Severity**: Patchable
**Pattern group**: operand-security-configuration

**Current state**: §6 mentions security guardrails generically but does not mandate minimum-privilege verification for RBAC.

**Required addition**: For operators with multiple controllers, the assessment must document whether each controller's RBAC permissions are scoped to the minimum verbs and resources required for its reconciliation logic. Flag wildcard verbs or broad resource grants.

---

## Gap RA-4: Missing Platform Security Policy Evolution Tracking

**Affected section**: §10.1 Security Context & Permissions
**Class of missing information**: The template does not require tracking how platform security policy evolution (new platform versions introducing new default policies) affects operand security configurations.
**Severity**: Eval-only
**Pattern group**: operand-security-configuration

**Current state**: The assessment focuses on current security requirements but does not mandate forward-compatibility analysis with upcoming platform security defaults.

**Recommendation**: Add guidance that when the operator runs on a platform with evolving security policies, the assessment should note version-specific security policy interactions and potential breaking changes in upcoming platform releases.

---

## Gap RA-5: Missing Dependency Vulnerability Surface Documentation

**Affected section**: §4.3 Image / Dependency Resolution
**Class of missing information**: The template does not require documentation of the dependency vulnerability surface — how many transitive dependencies exist, what the update blast radius is, and whether automated vulnerability scanning exists.
**Severity**: Deferred
**Pattern group**: cve-dependency-response

**Current state**: §4.3 covers how dependencies are resolved but not the vulnerability management posture.

**Recommendation**: Consider adding guidance for documenting dependency tree depth, vendor directory size, and automated scanning integration status.
