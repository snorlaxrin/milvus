package inspector

import (
	"context"

	"github.com/milvus-io/milvus/internal/streamingnode/server/wal/interceptors/timetick/mvcc"
	"github.com/milvus-io/milvus/internal/streamingnode/server/wal/interceptors/wab"
	"github.com/milvus-io/milvus/pkg/v2/streaming/util/types"
)

type TimeTickSyncOperator interface {
	// Channel returns the pchannel info.
	Channel() types.PChannelInfo

	// MVCCManager returns the related mvcc timestamp manager of current wal.
	MVCCManager() *mvcc.MVCCManager

	// WriteAheadBuffer get the related WriteAhead buffer.
	WriteAheadBuffer() wab.ROWriteAheadBuffer

	// Sync trigger a sync operation, try to send the timetick message into wal.
	// Sync operation is a blocking operation, and not thread-safe, will only call in one goroutine.
	Sync(ctx context.Context, forcePersisted bool)
}

// TimeTickSyncInspector is the inspector to sync time tick.
type TimeTickSyncInspector interface {
	// TriggerSync adds a pchannel info and notify the sync operation.
	// manually trigger the sync operation of pchannel.
	TriggerSync(pChannelInfo types.PChannelInfo, forcePersisted bool)

	// RegisterSyncOperator registers a sync operator.
	RegisterSyncOperator(operator TimeTickSyncOperator)

	// MustGetOperator gets the operator by pchannel info, otherwise panic.
	MustGetOperator(types.PChannelInfo) TimeTickSyncOperator

	// UnregisterSyncOperator unregisters a sync operator.
	UnregisterSyncOperator(operator TimeTickSyncOperator)

	// Close closes the inspector.
	Close()
}
