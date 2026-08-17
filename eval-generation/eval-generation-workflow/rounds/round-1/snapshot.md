# Round 1 Snapshot — Eval Generation

**Date**: 2026-08-17
**Status**: Completed

## Summary

Round 1 of eval generation transformed Epic Bug Analysis outputs (3 ZTWIM-specific bugs +
4 PR-derived issues across 5 pattern groups) into 24 generic, template-level eval cases
and 22 gap reports across 6 templates.

## Key Metrics

| Metric | Value |
|--------|-------|
| Total gaps identified | 22 |
| Patchable gaps (applied) | 14 |
| Eval-only gaps | 6 |
| Deferred gaps | 2 |
| Total eval cases | 24 |
| Templates refined | 5 |
| Patches applied | 6 |

## Pattern Groups → Eval Coverage

| Pattern Group | Gaps | Evals | Primary Stage |
|---------------|------|-------|---------------|
| operand-security-configuration | 8 | 10 | repo-assessment |
| multi-arch-support | 4 | 5 | repo-assessment + tasks |
| reconciliation-safety | 6 | 7 | plan |
| operand-configuration | 2 | 2 | tasks + code-generation |
| cve-dependency-response | 2 | 0 | deferred |

## Files Created/Modified

### Gap Reports (6 files)
- `template-gaps/repo-assessment-gaps.md` — 5 gaps (3 patchable, 1 eval-only, 1 deferred)
- `template-gaps/plan-gaps.md` — 5 gaps (3 patchable, 2 eval-only)
- `template-gaps/tasks-gaps.md` — 5 gaps (3 patchable, 1 eval-only, 1 deferred)
- `template-gaps/validation-gaps.md` — 3 gaps (2 patchable, 1 eval-only)
- `template-gaps/spec-gaps.md` — 4 gaps (3 patchable, 1 eval-only)
- `template-gaps/agents-gaps.md` — 5 gaps (2 patchable, 2 eval-only, 1 deferred)

### Eval Cases (4 files, 24 cases)
- `output-evals/repo-assessment/repo-assessment_eval.yaml` — 5 cases
- `output-evals/plan/plan_eval.yaml` — 5 cases
- `output-evals/tasks/tasks_eval.yaml` — 5 cases
- `output-evals/code-generation/code-generation_eval.yaml` — 9 cases

### Refined Templates (5 files patched)
- `output-refined-templates/repo-assessment-template.md` — §10.1 + §10.4 enhanced
- `output-refined-templates/plan-template.md` — Phase template expanded
- `output-refined-templates/tasks-template.md` — Core responsibilities #1, #8, #9 added
- `output-refined-templates/validation-template.md` — Completeness pillars added
- `output-refined-templates/spec-template.md` — Edge cases + Assumptions guidance expanded

## Generalization Verification

All 24 eval cases pass the SYSTEM_PROMPT self-check:
- [x] Reusability: Every eval applies to any future feature on this operator type
- [x] No proper nouns: No feature names, component names, or file paths from analyzed epic
- [x] Template-grounded: Each eval tests a requirement defined by the template schema
- [x] Comprehensible in isolation: Engineer unfamiliar with the analyzed feature would understand
- [x] ADR-independent: No eval assumes a specific architectural decision from the analyzed ADR

## Next Round Triggers

Round 2 should be triggered when:
1. New epic bug analysis outputs are available (new feature bundle analyzed)
2. Deferred gaps (RA-5, TK-5, AG-5) are prioritized for resolution
3. Eval case effectiveness data is available (pass/fail rates from actual pipeline runs)
