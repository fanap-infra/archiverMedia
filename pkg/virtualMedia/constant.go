package virtualMedia

import "github.com/fanap-infra/archiverMedia/pkg/err"

const (
	FrameChunkIdentifier        = "BehFramChunk"
	FrameChunkIdentifierSize    = 12
	FrameChunkSizeStoreSize     = 4
	FrameChunkHeader            = FrameChunkIdentifierSize + FrameChunkSizeStoreSize
	ErrFileIsReadOnly           = err.Error("file is read-only")
	FrameChunkMinimumFrameCount = 10
)
