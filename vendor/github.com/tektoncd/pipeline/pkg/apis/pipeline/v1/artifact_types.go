package v1

// Algorithm Standard cryptographic hash algorithm
type Algorithm string

// Artifact represents an artifact within a system, potentially containing multiple values
// associated with it.
type Artifact struct {
	Name   string          `json:"name,omitempty"`   // The artifact's identifying category name
	Values []ArtifactValue `json:"values,omitempty"` // A collection of values related to the artifact
}

// ArtifactValue represents a specific value or data element within an Artifact.
type ArtifactValue struct {
	Digest map[Algorithm]string `json:"digest,omitempty"` // Algorithm-specific digests for verifying the content (e.g., SHA256)
	Uri    string               `json:"uri,omitempty"`    // Location where the artifact value can be retrieved
}

// TaskRunStepArtifact represents an artifact produced or used by a step within a task run.
// It directly uses the Artifact type for its structure.
type TaskRunStepArtifact = Artifact

// Artifacts represents the collection of input and output artifacts associated with
// a task run or a similar process. Artifacts in this context are units of data or resources
// that the process either consumes as input or produces as output.
type Artifacts struct {
	Inputs  []Artifact `json:"inputs,omitempty"`
	Outputs []Artifact `json:"outputs,omitempty"`
}