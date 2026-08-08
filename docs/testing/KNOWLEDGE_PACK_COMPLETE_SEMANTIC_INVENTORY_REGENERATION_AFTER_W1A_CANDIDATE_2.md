# Candidate.2 — Complete Semantic Inventory Regeneration after W1A

| Field | Value |
|---|---|
| Record ID | `SEM-INV-REGEN-001` |
| Record Version | `0.1.0` |
| Status | Draft |
| Repository Commit | `efde0491fb0c0c49f3b4612fa0659855cbd221db` |
| Repository Tree | `145ce8321b9efac247f28c4138a5edc1ad847eca` |
| Numbered Knowledge Pack Documents | `30` |
| Supplemental Governance Sources | `1` |
| Supplemental Source | `OBDIA-CONST-001` |
| Semantic Source Documents | `31` |
| Knowledge Document 31 Created | **NO** |
| Semantic Regeneration | **PASS** |
| Wave 1 Complete | **NO** |
| Baseline Freeze | **BLOCKED** |

## Source Scope

The semantic source scope is:

- Knowledge documents `01–30`: 30 documents;
- supplemental governance source:
  `OBDIA-CONST-001`;
- total semantic source documents: 31.

`OBDIA-CONST-001` is the Documentation Constitution and is not a numbered Knowledge Pack document.

Its historical governance-anchor edge to `OBDIA-CANON-001` is reproduced.

## Historical Candidate.2 Inventory

- Raw relationship references: `800`
- Unique structural edges: `448`
- Advisory paths: `73`
- Cyclic components: `2`
- Reciprocal-edge rows: `184`
- Forward edges: `96`
- Backward edges: `322`
- Governance-anchor edges: `30`

## Regenerated Inventory

- Raw relationship references: `799`
- Unique structural edges: `448`
- Advisory paths: `73`
- Cyclic components: `2`
- Reciprocal-edge rows: `184`
- Forward edges: `96`
- Backward edges: `322`
- Governance-anchor edges: `30`

## Controlled W1A Delta

Exactly one existing relationship changed semantic classification:

- Review edge: `SEM-EDGE-0004`
- Relationship:
  `OBDIA-GLOSS-001 → OBDIA-RES-001`
- Historical fields: `Dependencies; Normative References`
- Current fields: `Cross-References`

No structural edge was added or removed.

The raw relationship-reference count changes from `800` to `799` because the former two-field normative relationship is now represented by one Cross-Reference field.

## Cycle Structure

- Component 1: `OBDIA-GLOSS-001`, `OBDIA-PUB-001`, `OBDIA-RES-001`
- Component 2: `OBDIA-AGENT-001`, `OBDIA-ASM-001`, `OBDIA-AUTH-001`, `OBDIA-CODE-001`, `OBDIA-COMP-001`, `OBDIA-CONN-001`, `OBDIA-DEC-001`, `OBDIA-DIAG-001`, `OBDIA-EVID-001`, `OBDIA-GH-001`, `OBDIA-IMP-001`, `OBDIA-NAME-001`, `OBDIA-REL-001`, `OBDIA-RISK-001`, `OBDIA-RSCH-001`, `OBDIA-SEC-001`, `OBDIA-STATE-001`, `OBDIA-TEST-001`, `OBDIA-VER-001`

Structural cycle membership is unchanged.

This is expected because the explicit informational edge remains. The dependency and normative-reference classifications removed by the human-ratified W1A correction are no longer present.

## Wave 1 State

- W1A completed rows: `5`
- Pending accountable-human review rows: `242`
- Wave 1 complete: **NO**
- Baseline freeze: **BLOCKED**

## Validation Conclusion

The complete semantic inventory was regenerated against the governed repository state at commit `efde0491fb0c0c49f3b4612fa0659855cbd221db`.

The regeneration reproduces the historical Constitution governance anchor, preserves the structural edge set, and reflects exactly the approved W1A semantic reclassification of `SEM-EDGE-0004`.

This evidence does not replace human review of the remaining 242 relationships and does not establish candidate approval or baseline freeze.

## Governance Boundary

This record does not establish:

- Knowledge document 31;
- Wave 1 completion;
- candidate.2 approval;
- baseline freeze;
- legal compliance or certification;
- implementation effectiveness;
- risk acceptance;
- independent external assurance;
- autonomous AI authority.
