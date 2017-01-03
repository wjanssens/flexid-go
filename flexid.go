package flexid

import (
	"time"
	"crypto/sha256"
	"unsafe"
)

/**
 * Generates 64-bit integer ids with a time component, sequence number component, and partition component.
 *
 * This system of generating IDs requires no central authority to generate key values.
 *
 * Since the time component of the ID is the first bits of the integer, IDs are inherently sortable.
 * This is often an advantage for storage and retrieval of rows in a database since it makes entities sortable and
 * captures the creation timestamp without the use of an extra column, but may be a disadvantage if it's not desirable
 * to expose the sorted identifiers in URLs.  To overcome this limitation it is recommended to translate IDs before
 * showing them to users by using HashIds (http://hashids.org/) or Knuth's Integer Hash.
 */
type FlexId struct {
	epoch         int64
	sequenceBits  uint8
	partitionBits uint8
	sequenceMask  int16
	partitionMask int16

	Sequence      int16
	Partition     int16
}

/*
 * Creates a new FlexId Generator.
 */
func NewFlexId(epoch int64, sequenceBits uint8, partitionBits uint8) *FlexId {
	if sequenceBits > 15 { panic("sequenceBits must be < 16") }
	if partitionBits > 15 { panic("partitionBits must by < 16") }

	var i uint8;
	var sequenceMask int16 = 0
	for i = 0; i < sequenceBits; i++ {
		sequenceMask = (sequenceMask << 1) | 1
	}
	var partitionMask int16 = 0
	for i = 0; i < partitionBits; i++ {
		partitionMask = (partitionMask << 1) | 1
	}

	return &FlexId{
		epoch: epoch,
		sequenceBits: sequenceBits,
		sequenceMask: sequenceMask,
		partitionBits: partitionBits,
		partitionMask: partitionMask,
	}
}

/*
 * Generates an ID with the supplied millis, supplied sequence, and supplied partition.
 * @param millis the number of milliseconds since epoch
 */
func (id *FlexId) Generate(args ...uint64) (result int64) {
	var millis int64
	var sequence = id.Sequence
	var partition = id.Partition

	switch len(args) {
	case 0:
		millis = time.Now().Unix() - id.epoch
		sequence = id.Sequence
		partition = id.Partition
	case 1:
		millis = int64(args[0]);
		sequence = id.Sequence
		partition = id.Partition
	case 2:
		millis = int64(args[0])
		sequence = int16(args[1])
		partition = id.Partition
	default:
		millis = int64(args[0])
		sequence = int16(args[1])
		partition = int16(args[2])
	}

	if len(args) < 3 {
		id.Sequence++
	}

	result =  millis << (id.sequenceBits + id.partitionBits)
	result |= int64(sequence & id.sequenceMask) << id.partitionBits
	result |= int64(partition & id.partitionMask)

	return result
}

/*
 * Extracts the millis component of an ID.
 * This is the raw value and is not adjusted for epoch.
 */
func (id FlexId) ExtractMillis(v int64) int64 {
	return v >> (id.sequenceBits + id.partitionBits)
}

/*
 * Extracts the date/time component of an ID.
 * This is the derived from the millis component of the ID with the configured epoch applied.
 */
func (id FlexId) ExtractTimestamp(v int64) time.Time {
	return time.Unix(id.ExtractMillis(v) + id.epoch, 0)
}

/*
 * Extracts the sequence component of an ID.
 */
func (id FlexId) ExtractSequence(v int64) int16 {
	return int16(v >> id.partitionBits) & id.sequenceMask
}

/*
 * Extracts the partition component of an ID.
 * If bits is 0 < b < id.partitionBits then only that many bits of the partition is returned,
 * which serves as simple method for mapping a large logical partition space onto a smaller physical partition space.
 */
func (id FlexId) ExtractPartition(v int64, bits uint8) int16 {
	mask := id.partitionMask;

	if bits > 0 && bits < id.partitionBits {
		mask = 0;
		for i := uint8(0); i < bits; i++ {
			mask = (mask << 1) | 1
		}
	}

	return int16(v) & mask
}

/*
 * A convenience method for computing a partition value from a string using an SHA-256 hash.
 * This would typically be used to compute a shard ID from a string identifier such as a username.
 */
func Sum256(text string) int16 {
	sum := sha256.Sum256([]byte(text));
	return *(*int16)(unsafe.Pointer(&sum[0]))
}

