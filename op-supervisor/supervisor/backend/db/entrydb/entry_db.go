package entrydb

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/ethereum/go-ethereum/log"
)

const (
	RecordSize = 25 // 24 bytes of user data plus a 1 byte entry ID
	EntrySize  = RecordSize - 1
)

type EntryIdx int64

// EntryID is a single byte identifier automatically added to each entry in the database.
// It is designed to detect when the entry at a given index has been changed by truncating and then rewriting new entries.
// It typically increments monotonically except if doing so would cause new entries written after a Truncate to be
// assigned the same IDs as previously.
type EntryID byte

type Entry [EntrySize]byte

// dataAccess defines a minimal API required to manipulate the actual stored data.
// It is a subset of the os.File API but could (theoretically) be satisfied by an in-memory implementation for testing.
type dataAccess interface {
	io.ReaderAt
	io.Writer
	io.Closer
	Truncate(size int64) error
}

type EntryDB struct {
	data         dataAccess
	lastEntryIdx EntryIdx
	nextEntryID  EntryID

	cleanupFailedWrite bool
}

// NewEntryDB creates an EntryDB. A new file will be created if the specified path does not exist,
// but parent directories will not be created.
// If the file exists it will be used as the existing data.
// Will automatically attempt to recover an invalid database by truncating any partially written entries.
func NewEntryDB(logger log.Logger, path string) (*EntryDB, error) {
	logger.Info("Opening entry database", "path", path)
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o644)
	if err != nil {
		return nil, fmt.Errorf("failed to open database at %v: %w", path, err)
	}
	info, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to stat database at %v: %w", path, err)
	}
	size := info.Size() / RecordSize
	db := &EntryDB{
		data:         file,
		lastEntryIdx: EntryIdx(size - 1),
		// Arbitrary starting point for this session.
		// Any recovery from a prior crash will be handled before writing new entries uses this anyway so it doesn't
		// matter what starting value we use.
		nextEntryID: 0,
	}
	if size*RecordSize != info.Size() {
		logger.Warn("File size is nut a multiple of entry size. Truncating to last complete entry", "fileSize", size, "entrySize", RecordSize)
		if err := db.recover(); err != nil {
			return nil, fmt.Errorf("failed to recover database at %v: %w", path, err)
		}
	}
	return db, nil
}

func (e *EntryDB) Size() int64 {
	return int64(e.lastEntryIdx) + 1
}

func (e *EntryDB) LastEntryIdx() EntryIdx {
	return e.lastEntryIdx
}

// Read an entry from the database by index. Returns io.EOF iff idx is after the last entry.
func (e *EntryDB) Read(idx EntryIdx) (Entry, EntryID, error) {
	if idx > e.lastEntryIdx {
		return Entry{}, 0, io.EOF
	}
	var out Entry
	read, err := e.data.ReadAt(out[:], int64(idx)*RecordSize)
	// Ignore io.EOF if we read the entire last entry as ReadAt may return io.EOF or nil when it reads the last byte
	if err != nil && !(errors.Is(err, io.EOF) && read == RecordSize) {
		return Entry{}, 0, fmt.Errorf("failed to read entry %v: %w", idx, err)
	}
	var id [1]byte
	read, err = e.data.ReadAt(id[:], int64(idx)*RecordSize+EntrySize)
	if err != nil && !(errors.Is(err, io.EOF) && read == 1) {
		return Entry{}, 0, fmt.Errorf("failed to read entry id %v: %w", idx, err)
	}
	return out, EntryID(id[0]), nil
}

// Append entries to the database.
// The entries are combined in memory and passed to a single Write invocation.
// If the write fails, it will attempt to truncate any partially written data.
// Subsequent writes to this instance will fail until partially written data is truncated.
func (e *EntryDB) Append(entries ...Entry) error {
	if e.cleanupFailedWrite {
		// Try to rollback partially written data from a previous Append
		if truncateErr := e.Truncate(e.lastEntryIdx); truncateErr != nil {
			return fmt.Errorf("failed to recover from previous write error: %w", truncateErr)
		}
	}
	id := e.nextEntryID
	data := make([]byte, 0, len(entries)*RecordSize)
	for _, entry := range entries {
		data = append(data, entry[:]...)
		data = append(data, byte(id))
		id++
	}
	if n, err := e.data.Write(data); err != nil {
		if n == 0 {
			// Didn't write any data, so no recovery required
			return err
		}
		// Try to rollback the partially written data
		if truncateErr := e.Truncate(e.lastEntryIdx); truncateErr != nil {
			// Failed to rollback, set a flag to attempt the clean up on the next write
			e.cleanupFailedWrite = true
			return errors.Join(err, fmt.Errorf("failed to remove partially written data: %w", truncateErr))
		}
		// Successfully rolled back the changes, still report the failed write
		return err
	}
	e.lastEntryIdx += EntryIdx(len(entries))
	e.nextEntryID += EntryID(len(entries))
	return nil
}

// Truncate the database so that the last retained entry is idx. Any entries after idx are deleted.
// If any complete entries are removed, it ensures that the next entry written uses a different ID to the one
// removed from that index.
func (e *EntryDB) Truncate(idx EntryIdx) error {
	_, prevID, err := e.Read(idx + 1)
	if errors.Is(err, io.EOF) {
		// There is no complete entry after idx so we don't need to avoid using its entry ID,
		// but we still truncate to remove any partially written entries.
	} else if err != nil {
		return fmt.Errorf("failed to read entry %v prior to truncating: %w", idx, err)
	} else {
		// Note that we don't wind back the nextEntryID when truncating.
		// Just make sure that the next entry we write does not have an ID equal to the entry that was present.
		if e.nextEntryID == prevID {
			e.nextEntryID++
		}
	}
	if err := e.data.Truncate((int64(idx) + 1) * RecordSize); err != nil {
		return fmt.Errorf("failed to truncate to entry %v: %w", idx, err)
	}
	// Update the lastEntryIdx cache
	e.lastEntryIdx = idx
	e.cleanupFailedWrite = false

	return nil
}

// recover an invalid database by truncating back to the last complete event.
func (e *EntryDB) recover() error {
	if err := e.data.Truncate((e.Size()) * RecordSize); err != nil {
		return fmt.Errorf("failed to truncate trailing partial entries: %w", err)
	}
	return nil
}

func (e *EntryDB) Close() error {
	return e.data.Close()
}
