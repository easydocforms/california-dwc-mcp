// Package impairmentmcp connects to InjuryMgmt's public whole-person impairment
// MCP server. Send only synthetic or de-identified clinician-confirmed findings.
package impairmentmcp

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

const (
	// Endpoint serves the no-auth Streamable HTTP transport.
	Endpoint = "https://mcp.injurymgmt.com/mcp"
	// ToolName is the MCP tool for informational medical whole-person impairment.
	ToolName = "calculate_medical_wpi"
)

// Connect opens an MCP client session. The caller should close the session
// when finished and use a context with a deadline for network operations.
func Connect(ctx context.Context) (*mcp.ClientSession, error) {
	client := mcp.NewClient(&mcp.Implementation{
		Name:    "injurymgmt-impairment-go",
		Version: "0.1.0",
	}, nil)
	return client.Connect(ctx, &mcp.StreamableClientTransport{Endpoint: Endpoint}, nil)
}
