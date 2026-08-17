# Patch Summaries — Round 1

**Date**: 2026-08-17
**Applied to**: eval-generation/output-refined-templates/

---

## Patch 1: repo-assessment-template.md — Security Constraint Matrix

**Gap**: RA-1, RA-3
**Section**: §10.1 Security Context & Permissions
**Change**: Added per-operand security constraint matrix requirement (mandatory when multiple operands exist) and RBAC minimum-privilege audit requirement.
**Lines added**: ~8

---

## Patch 2: repo-assessment-template.md — External Image Architecture Matrix

**Gap**: RA-2
**Section**: §10.4 Build & Compliance Constraints
**Change**: Added external image architecture matrix requirement (mandatory when operator references external images not built by this repo).
**Lines added**: ~5

---

## Patch 3: plan-template.md — Reconciliation Strategy in Phase Template

**Gap**: PL-1, PL-2, PL-4
**Section**: Phase template (§5)
**Change**: Added "Reconciliation strategy" bullet (mandatory when phase involves controller reconciliation) specifying apply mechanism, conflict detection, idempotency, and workload-type-specific update strategies. Added "Rejected alternatives" bullet.
**Lines added**: ~8

---

## Patch 4: tasks-template.md — Core Responsibilities Expansion

**Gap**: TK-1, TK-2, TK-4
**Section**: Core responsibilities
**Change**: (1) Added operand-component granularity guidance to responsibility #1. (2) Added responsibility #8 for security configuration verification pairing. (3) Added responsibility #9 for multi-architecture validation gate.
**Lines added**: ~10

---

## Patch 5: validation-template.md — Completeness Pillars

**Gap**: VL-1, VL-2
**Section**: Rubric — A) COMPLETENESS
**Change**: Added two new completeness pillars: "Security Constraint Completeness" (triggered when spec describes components with heterogeneous privilege needs) and "Architecture Support Requirements" (triggered when spec references container image deployment).
**Lines added**: ~8

---

## Patch 6: spec-template.md — Edge Case Categories

**Gap**: SP-1, SP-2, SP-3
**Section**: Edge Cases guidance comment block; Assumptions section guidance comment block
**Change**: (1) Added edge case categories for security policy interactions, architecture portability, and reconciliation conflicts. (2) Added assumption categories for reconciliation semantics, architecture support, and security policy baseline.
**Lines added**: ~14
