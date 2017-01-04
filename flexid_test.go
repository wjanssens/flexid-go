package flexid

import (
	"testing"
	"time"
)

const (
	UNIX_EPOCH int64 = 0
	CUSTOM_EPOCH int64 = 1420070400000
)
func TestFlexId_Epoch(t *testing.T) {
	g := NewFlexId(CUSTOM_EPOCH, 4, 4, 4)
	g.Constant = 0x5A

	now := time.Now()

	id := g.Generate("test")

	if now.UnixNano() - g.ExtractMillis(id) < 2000 { t.Errorf("Incorrect millis %v", g.ExtractMillis(id)) }
	if g.ExtractTimestamp(id).Sub(now).Seconds() > 2 { t.Errorf("Incorrect timestamp %v", g.ExtractTimestamp(id)) }
	if 0x00 != g.ExtractSequence(id) { t.Errorf("Incorrect sequence %v", g.ExtractSequence(id)) }
	if 0x0B != g.ExtractShard(id, 0) { t.Errorf("Incorrect shard %v", g.ExtractShard(id, 0)) }
	if 0x0A != g.ExtractConstant(id) { t.Errorf("Incorrect constant %v", g.ExtractConstant(id)) }

	id2 := g.Generate("test");
	if 0x01 != g.ExtractSequence(id2) { t.Errorf("Incorrect sequence %v", g.ExtractSequence(id2)) }
}

func TestFlexId_Defaults(t *testing.T) {
	g := NewFlexId(UNIX_EPOCH, 8, 8, 0)
	g.Constant = 0x5A

	now := time.Now()

	id := g.Generate("test")

	if now.UnixNano() - g.ExtractMillis(id) < 2000 { t.Errorf("Incorrect millis %v", g.ExtractMillis(id)) }
	if g.ExtractTimestamp(id).Sub(now).Seconds() > 2 { t.Errorf("Incorrect timestamp %v", g.ExtractTimestamp(id)) }
	if 0x00 != g.ExtractSequence(id) { t.Errorf("Incorrect sequence %v", g.ExtractSequence(id)) }
	if 0x1B != g.ExtractShard(id, 0) { t.Errorf("Incorrect shard %v", g.ExtractShard(id, 0)) }
	if 0x00 != g.ExtractConstant(id) { t.Errorf("Incorrect constant %v", g.ExtractConstant(id)) }

	id2 := g.Generate("test");
	if 0x01 != g.ExtractSequence(id2) { t.Errorf("Incorrect sequence %v", g.ExtractSequence(id2)) }
}


func TestFlexId_Fours(t *testing.T) {
	g := NewFlexId(UNIX_EPOCH, 4, 4, 4)
	g.Constant = 0x5A

	now := time.Now()

	id := g.Generate("test")

	if now.UnixNano() - g.ExtractMillis(id) < 2000 { t.Errorf("Incorrect millis %v", g.ExtractMillis(id)) }
	if g.ExtractTimestamp(id).Sub(now).Seconds() > 2 { t.Errorf("Incorrect timestamp %v", g.ExtractTimestamp(id)) }
	if 0x00 != g.ExtractSequence(id) { t.Errorf("Incorrect sequence %v", g.ExtractSequence(id)) }
	if 0x0B != g.ExtractShard(id, 0) { t.Errorf("Incorrect shard %v", g.ExtractShard(id, 0)) }
	if 0x0A != g.ExtractConstant(id) { t.Errorf("Incorrect constant %v", g.ExtractConstant(id)) }

	id2 := g.Generate("test");
	if 0x01 != g.ExtractSequence(id2) { t.Errorf("Incorrect sequence %v", g.ExtractSequence(id2)) }
}

func TestFlexId_Fives(t *testing.T) {
	g := NewFlexId(UNIX_EPOCH, 5, 5, 5)
	g.Constant = 0x5A

	now := time.Now()

	id := g.Generate("test")

	if now.UnixNano() - g.ExtractMillis(id) < 2000 { t.Errorf("Incorrect millis %v", g.ExtractMillis(id)) }
	if g.ExtractTimestamp(id).Sub(now).Seconds() > 2 { t.Errorf("Incorrect timestamp %v", g.ExtractTimestamp(id)) }
	if 0x00 != g.ExtractSequence(id) { t.Errorf("Incorrect sequence %v", g.ExtractSequence(id)) }
	if 0x1B != g.ExtractShard(id, 0) { t.Errorf("Incorrect shard %v", g.ExtractShard(id, 0)) }
	if 0x1A != g.ExtractConstant(id) { t.Errorf("Incorrect constant %v", g.ExtractConstant(id)) }

	id2 := g.Generate("test");
	if 0x01 != g.ExtractSequence(id2) { t.Errorf("Incorrect sequence %v", g.ExtractSequence(id2)) }
}

func TestFlexId_Sixes(t *testing.T) {
	g := NewFlexId(UNIX_EPOCH, 6, 6, 6)
	g.Constant = 0x5A

	now := time.Now()

	id := g.Generate("test")

	if now.UnixNano() - g.ExtractMillis(id) < 2000 { t.Errorf("Incorrect millis %v", g.ExtractMillis(id)) }
	if g.ExtractTimestamp(id).Sub(now).Seconds() > 2 { t.Errorf("Incorrect timestamp %v", g.ExtractTimestamp(id)) }
	if 0x00 != g.ExtractSequence(id) { t.Errorf("Incorrect sequence %v", g.ExtractSequence(id)) }
	if 0x1B != g.ExtractShard(id, 0) { t.Errorf("Incorrect shard %v", g.ExtractShard(id, 0)) }
	if 0x1A != g.ExtractConstant(id) { t.Errorf("Incorrect constant %v", g.ExtractConstant(id)) }

	id2 := g.Generate("test");
	if 0x01 != g.ExtractSequence(id2) { t.Errorf("Incorrect sequence %v", g.ExtractSequence(id2)) }
}

func TestFlexId_Zeros(t *testing.T) {
	g := NewFlexId(UNIX_EPOCH, 0, 0, 0)
	g.Constant = 0x5A

	now := time.Now()

	id := g.Generate("test")

	if now.UnixNano() - g.ExtractMillis(id) < 2000 { t.Errorf("Incorrect millis %v", g.ExtractMillis(id)) }
	if g.ExtractTimestamp(id).Sub(now).Seconds() > 2 { t.Errorf("Incorrect timestamp %v", g.ExtractTimestamp(id)) }
	if 0x00 != g.ExtractSequence(id) { t.Errorf("Incorrect sequence %v", g.ExtractSequence(id)) }
	if 0x00 != g.ExtractShard(id, 0) { t.Errorf("Incorrect shard %v", g.ExtractShard(id, 0)) }
	if 0x00 != g.ExtractConstant(id) { t.Errorf("Incorrect constant %v", g.ExtractConstant(id)) }

	id2 := g.Generate("test");
	if 0x00 != g.ExtractSequence(id2) { t.Errorf("Incorrect sequence %v", g.ExtractSequence(id2)) }
}

