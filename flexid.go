package flexid

import (
	"time"
	"crypto/sha256"
	"encoding/binary"
)

/**
 * Generates 64-bit integer ids with a time component, sequence number component, and shard component.
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
	epoch        int64

	sequenceBits uint8
	sequenceMask int16
	shardBits    uint8
	shardMask    int16

	Sequence     int16
}

/*
 * Creates a new FlexId Generator.
 */
func NewFlexId(epoch int64, sequenceBits uint8, shardBits uint8) *FlexId {
	if sequenceBits > 15 {
		panic("sequenceBits must be < 16")
	}
	if shardBits > 15 {
		panic("shardBits must be < 16")
	}

	return &FlexId{
		epoch: epoch,
		sequenceBits: sequenceBits,
		sequenceMask: makeMask(sequenceBits),
		shardBits: shardBits,
		shardMask: makeMask(shardBits),
	}
}

func makeMask(bits uint8) int16 {
	var mask int16 = 0
	for i := 0; i < int(bits); i++ {
		mask = (mask << 1) | 1
	}
	return mask
}

/*
 * Generates an ID using the supplied shard key.
 */
func (id *FlexId) Generate(shardKey string) (result int64) {
	var millis int64
	millis = time.Now().Unix() - id.epoch
	sequence := id.Sequence
	shard := Sum256(shardKey)

	result = millis << (id.sequenceBits + id.shardBits)
	result |= int64(sequence & id.sequenceMask) << (id.shardBits)
	result |= int64(shard & id.shardMask)

	id.Sequence++

	return result
}

/*
 * Extracts the millis component of an ID.
 * This is the raw value and is not adjusted for epoch.
 */
func (id FlexId) ExtractRawMillis(v int64) int64 {
	return v >> (id.sequenceBits + id.shardBits)
}

/*
 * Extracts the millis component of an ID.
 * This is the raw value and is not adjusted for epoch.
 */
func (id FlexId) ExtractMillis(v int64) int64 {
	return id.ExtractRawMillis(v) + id.epoch
}

/*
 * Extracts the date/time component of an ID.
 * This is the derived from the millis component of the ID with the configured epoch applied.
 */
func (id FlexId) ExtractTimestamp(v int64) time.Time {
	return time.Unix(id.ExtractMillis(v), 0)
}

/*
 * Extracts the sequence component of an ID.
 */
func (id FlexId) ExtractSequence(v int64) int16 {
	return int16(v >> id.shardBits) & id.sequenceMask
}

/*
 * Extracts the partition component of an ID.
 * If bits is 0 < bits < id.shardBits then only that many bits of the partition is returned,
 * which serves as simple method for mapping a large logical partition space onto a smaller physical partition space.
 */
func (id FlexId) ExtractShard(v int64, bits uint8) int16 {
	if (bits > 0 && bits < id.shardBits) {
		return int16(v) & makeMask(bits);
	} else {
		return int16(v) & id.shardMask;
	}
}

/*
 * A convenience method for computing a shard value from a string using an SHA-256 hash.
 * This would typically be used to compute a shard ID from a string identifier such as a username.
 */
func Sum256(text string) int16 {
	sum := sha256.Sum256([]byte(text));
	return int16(binary.BigEndian.Uint16(sum[18:]))
}

