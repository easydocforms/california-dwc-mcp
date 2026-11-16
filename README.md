# InjuryMgmt WPI Calculator MCP

[InjuryMgmt's California workers' compensation impairment calculator](https://injurymgmt.com/california-dwc-impairment-calculator/) also has a free, hosted Model Context Protocol (MCP) endpoint. It returns informational medical whole-person impairment (WPI), status, calculator version and source references from clinician-confirmed, de-identified inputs for one injury.

The tool supports **spine Diagnosis Related Estimates (DRE)** with clinician-selected category and WPI, and **active mobile shoulder, elbow/forearm and wrist range of motion**. It does not select a spine DRE category, accept medical records, or convert WPI to a permanent disability percentage under the 2005 PDRS. InjuryMgmt is independent of the California Division of Workers' Compensation (DWC).

## Connect

- **Endpoint:** `https://mcp.injurymgmt.com/mcp`
- **Transport:** Streamable HTTP over HTTPS
- **Authentication:** None
- **Tool:** [`calculate_medical_wpi`](schemas/calculate_medical_wpi.json)

Add the endpoint as a remote server in an MCP client. For example, in Claude Code:

```sh
claude mcp add --transport http injurymgmt-wpi https://mcp.injurymgmt.com/mcp
```

Try a synthetic QA prompt:

> Use InjuryMgmt for a synthetic case. The clinician selected lumbar DRE Category II and 8% WPI. Return the medical WPI and cite the source reference.

The expected result is the clinician-selected 8% medical WPI with a reference to AMA Guides, Fifth Edition, Table 15-3.

## Go client

This repository holds a small Go client, an MCP Registry manifest, and a snapshot of the public tool schema. It connects to the hosted server; it is not the calculator implementation. Use Go 1.25 or later.

```sh
go get github.com/easydocforms/california-dwc-mcp@v0.1.0
```

Call `impairmentmcp.Connect(ctx)` to open an SDK session and `session.ListTools(ctx, nil)` to discover the current schema. Call `session.Close()` when finished.

Use only synthetic or de-identified inputs. Have the rating clinician review the inputs and results before use.

[InjuryMgmt](https://injurymgmt.com/) · [Web calculator and supported methods](https://injurymgmt.com/california-dwc-impairment-calculator/) · [MCP privacy](https://mcp.injurymgmt.com/privacy) · [MCP terms](https://mcp.injurymgmt.com/terms)
