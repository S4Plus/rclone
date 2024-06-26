package fs

import (
	"context"
	"time"
)

// DirWrapper wraps a Directory object so the Remote can be overridden
type DirWrapper struct {
	Directory        // Directory we are wrapping
	remote    string // name of the directory
}

// NewDirWrapper creates a wrapper for a directory object
//
// This passes through optional methods and should be used for
// wrapping backends to wrap native directories.
func NewDirWrapper(remote string, d Directory) *DirWrapper {
	return &DirWrapper{
		Directory: d,
		remote:    remote,
	}
}

// String returns the name
func (d *DirWrapper) String() string {
	return d.remote
}

// Remote returns the remote path
func (d *DirWrapper) Remote() string {
	return d.remote
}

// SetRemote sets the remote
func (d *DirWrapper) SetRemote(remote string) *DirWrapper {
	d.remote = remote
	return d
}

// Metadata returns metadata for an DirEntry
//
// It should return nil if there is no Metadata
func (d *DirWrapper) Metadata(ctx context.Context) (Metadata, error) {
	do, ok := d.Directory.(Metadataer)
	if !ok {
		return nil, nil
	}
	return do.Metadata(ctx)
}

// SetMetadata sets metadata for an DirEntry
//
// It should return fs.ErrorNotImplemented if it can't set metadata
func (d *DirWrapper) SetMetadata(ctx context.Context, metadata Metadata) error {
	do, ok := d.Directory.(SetMetadataer)
	if !ok {
		return ErrorNotImplemented
	}
	return do.SetMetadata(ctx, metadata)
}

// SetModTime sets the metadata on the DirEntry to set the modification date
//
// If there is any other metadata it does not overwrite it.
func (d *DirWrapper) SetModTime(ctx context.Context, t time.Time) error {
	do, ok := d.Directory.(SetModTimer)
	if !ok {
		return ErrorNotImplemented
	}
	return do.SetModTime(ctx, t)
}

// Check interfaces
var (
	_ DirEntry      = (*DirWrapper)(nil)
	_ Directory     = (*DirWrapper)(nil)
	_ FullDirectory = (*DirWrapper)(nil)
)
