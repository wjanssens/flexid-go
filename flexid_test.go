package flexid

import (
	"testing"
	"time"
)

const (
	UNIX_EPOCH int64 = 0
	DEFAULT_EPOCH int64 = 1420070400000
	INSTAGRAM_EPOCH int64 = 1293840000000
)
func TestFlexId_Millis(t *testing.T) {
	g := NewFlexId(DEFAULT_EPOCH, 8, 8);
	id := g.Generate(0x5A5A5A5A5A5A, 0x00, 0x00);
	if 0x5A5A5A5A5A5A0000 != id { t.Errorf("Incorrect id %d", id) }
	if 0x5A5A5A5A5A5A != g.ExtractMillis(id) { t.Errorf("Incorrect millis %d", g.ExtractMillis(id)) }
	if 0x00 != g.ExtractSequence(id) { t.Errorf("Incorrect sequence %d", g.ExtractSequence(id)) }
	if 0x00 != g.ExtractPartition(id, 0) { t.Errorf("Incorrect partition %d", g.ExtractPartition(id, 0)) }
}

func TestFlexId_Sequence(t *testing.T) {
	g := NewFlexId(DEFAULT_EPOCH, 8, 8);
	id := g.Generate(0x00, 0x5A, 0x00);
	if 0x5A00 != id { t.Errorf("Incorrect id %d", id) }
	if 0x00 != g.ExtractMillis(id) { t.Errorf("Incorrect millis %d", g.ExtractMillis(id)) }
	if 0x5A != g.ExtractSequence(id) { t.Errorf("Incorrect sequence %d", g.ExtractSequence(id)) }
	if 0x00 != g.ExtractPartition(id, 0) { t.Errorf("Incorrect partition %d", g.ExtractPartition(id, 0)) }

}

func TestFlexId_Partition(t *testing.T) {
	g := NewFlexId(DEFAULT_EPOCH, 8, 8);
	id := g.Generate(0x00, 0x00, 0x5A);
	if 0x5A != id { t.Errorf("Incorrect id %d", id) }
	if 0x00 != g.ExtractMillis(id) { t.Errorf("Incorrect millis %d", g.ExtractMillis(id)) }
	if 0x00 != g.ExtractSequence(id) { t.Errorf("Incorrect sequence %d", g.ExtractSequence(id)) }
	if 0x5A != g.ExtractPartition(id, 0) { t.Errorf("Incorrect partition %d", g.ExtractPartition(id, 0)) }
}

func TestFlexId_State(t *testing.T) {
	g := NewFlexId(DEFAULT_EPOCH, 8, 8);
	g.Sequence = 0x5A
	g.Partition = 0xA5
	id := g.Generate();
	if 0x5A != g.ExtractSequence(id) { t.Errorf("Incorrect sequence %d", g.ExtractSequence(id)) }
	if 0xA5 != g.ExtractPartition(id, 0) { t.Errorf("Incorrect partition %d", g.ExtractPartition(id, 0)) }

	id2 := g.Generate();
	if 0x5B != g.ExtractSequence(id2) { t.Errorf("Incorrect sequence %d", g.ExtractSequence(id2)) }
	if 0xA5 != g.ExtractPartition(id2, 0) { t.Errorf("Incorrect partition %d", g.ExtractPartition(id2, 0)) }
}

func TestFlexId_8_8(t *testing.T) {
	g := NewFlexId(DEFAULT_EPOCH, 8, 8);
	id := g.Generate(0x5A5A5A5A5A5A, 0x5A, 0x5A);
	if 0x5A5A5A5A5A5A5A5A != id { t.Errorf("Incorrect id %d", id) }
	if 0x5A5A5A5A5A5A != g.ExtractMillis(id) { t.Errorf("Incorrect millis %d", g.ExtractMillis(id)) }
	if 0x5A != g.ExtractSequence(id) { t.Errorf("Incorrect sequence %d", g.ExtractSequence(id)) }
	if 0x5A != g.ExtractPartition(id, 0) { t.Errorf("Incorrect partition %d", g.ExtractPartition(id, 0)) }
}

func TestFlexId_10_8(t *testing.T) {
	g := NewFlexId(DEFAULT_EPOCH, 10, 8);
	id := g.Generate(0x5A5A5A5A5, 0x25A, 0xA5);
	if 0x16969696965AA5 != id { t.Errorf("Incorrect id %d", id) }
	if 0x5A5A5A5A5 != g.ExtractMillis(id) { t.Errorf("Incorrect millis %d", g.ExtractMillis(id)) }
	if 0x25A != g.ExtractSequence(id) { t.Errorf("Incorrect sequence %d", g.ExtractSequence(id)) }
	if 0xA5 != g.ExtractPartition(id, 0) { t.Errorf("Incorrect partition %d", g.ExtractPartition(id, 0)) }
}

func TestFlexId_10_13(t *testing.T) {
	g := NewFlexId(INSTAGRAM_EPOCH, 10, 13);
	id := g.Generate(0x5A5A5A5A5, 0x25A, 0xA5);
	if 0x2D2D2D2D2CB40A5 != id { t.Errorf("Incorrect id %d", id) }
	if 0x5A5A5A5A5 != g.ExtractMillis(id) { t.Errorf("Incorrect millis %d", g.ExtractMillis(id)) }
	if 0x25A != g.ExtractSequence(id) { t.Errorf("Incorrect sequence %d", g.ExtractSequence(id)) }
	if 0xA5 != g.ExtractPartition(id, 0) { t.Errorf("Incorrect partition %d", g.ExtractPartition(id, 0)) }
}

func TestFlexId_0_0(t *testing.T) {
	g := NewFlexId(DEFAULT_EPOCH, 0, 0);
	id := g.Generate(0x5A5A5A5A5, 0x25A, 0xA5);
	if 0x5A5A5A5A5 != id { t.Errorf("Incorrect id %d", id) }
	if 0x5A5A5A5A5 != g.ExtractMillis(id) { t.Errorf("Incorrect millis %d", g.ExtractMillis(id)) }
	if 0x00 != g.ExtractSequence(id) { t.Errorf("Incorrect sequence %d", g.ExtractSequence(id)) }
	if 0x00 != g.ExtractPartition(id, 0) { t.Errorf("Incorrect partition %d", g.ExtractPartition(id, 0)) }
}

func TestFlexId_Timestamp_UnixEpoch(t *testing.T) {
	g := NewFlexId(UNIX_EPOCH, 8, 8);
	id := g.Generate();
	if time.Now().UnixNano() - g.ExtractTimestamp(id).UnixNano() < 1000 { t.Errorf("Incorrect timstamp %v", g.ExtractTimestamp(id)) }
}

func TestFlexId_Timestamp_CustomEpoch(t *testing.T) {
	g := NewFlexId(DEFAULT_EPOCH, 8, 8);
	id := g.Generate();
	if time.Now().UnixNano() - g.ExtractTimestamp(id).UnixNano() < 1000 { t.Errorf("Incorrect timstamp %v", g.ExtractTimestamp(id)) }
}
