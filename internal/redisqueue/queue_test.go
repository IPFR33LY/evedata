package redisqueue

import (
	"testing"

	"github.com/antihax/evedata/internal/redigohelper"
	"github.com/stretchr/testify/assert"
)

func TestHQ(t *testing.T) {
	pool := redigohelper.ConnectRedisTestPool()
	hq := NewRedisQueue(pool, "test-redisqueue")
	err := hq.QueueWork(
		[]Work{
			{"alliance",
				2,
			},
			{"alliance",
				3,
			},
			{"alliance",
				4,
			},
			{"alliance",
				5,
			},
			{"alliance",
				6,
			},
		},
		Priority_High,
	)
	assert.Nil(t, err)
	size, err := hq.Size()
	assert.Nil(t, err)
	assert.Equal(t, 5, size)

	size, err = hq.ScaledSize()
	assert.Nil(t, err)
	assert.Equal(t, 5120, size)

	var work []*Work
	for i := 0; i < 5; i++ {
		w, err := hq.GetWork()
		assert.Nil(t, err)
		work = append(work, w)
	}
	check := map[int]bool{2: true, 3: true, 4: true, 5: true, 6: true}

	for i := range work {
		delete(check, work[i].Parameter.(int))
	}
	assert.Empty(t, check)
}

func TestExpired(t *testing.T) {
	pool := redigohelper.ConnectRedisTestPool()
	hq := NewRedisQueue(pool, "test-redisqueue")
	err := hq.SetWorkExpire("testKey", 1, 9999)
	assert.Nil(t, err)

	b := hq.CheckWorkExpired("testKey", 1)
	assert.Equal(t, b, true)
}

func TestCompletion(t *testing.T) {
	pool := redigohelper.ConnectRedisTestPool()
	hq := NewRedisQueue(pool, "test-redisqueue")
	err := hq.SetWorkCompleted("testKeyComp", 1)
	assert.Nil(t, err)

	b := hq.CheckWorkCompleted("testKeyComp", 1)
	assert.Equal(t, b, true)
}

func TestCache(t *testing.T) {
	pool := redigohelper.ConnectRedisTestPool()
	hq := NewRedisQueue(pool, "test-redisqueue")
	err := hq.SetCache("testCache", int32(1), "hi")
	assert.Nil(t, err)

	s, err := hq.GetCache("testCache", int32(1))
	assert.Nil(t, err)
	assert.Equal(t, s, "hi")
}

func TestSetWorkCompletedInBulk(t *testing.T) {
	pool := redigohelper.ConnectRedisTestPool()
	hq := NewRedisQueue(pool, "test-redisqueue")
	err := hq.SetWorkCompletedInBulk("testKeyComp", []int64{1, 2, 3, 4, 5, 6, 7})
	assert.Nil(t, err)

	b := hq.CheckWorkCompleted("testKeyComp", 3)
	assert.Equal(t, b, true)
}

func TestExpiredInBulk(t *testing.T) {
	pool := redigohelper.ConnectRedisTestPool()
	hq := NewRedisQueue(pool, "test-redisqueue")
	err := hq.SetWorkExpire("testKey", 1, 9999)
	assert.Nil(t, err)

	err = hq.SetWorkExpire("testKey", 3, 9999)
	assert.Nil(t, err)

	b, err := hq.CheckWorkExpiredInBulk("testKey", []int64{1, 2, 3, 4})

	assert.Nil(t, err)
	assert.Equal(t, b[0], true)
	assert.Equal(t, b[2], true)
	assert.Equal(t, b[1], false)
	assert.Equal(t, b[3], false)
}

func TestCompletionInBulk(t *testing.T) {
	pool := redigohelper.ConnectRedisTestPool()
	hq := NewRedisQueue(pool, "test-redisqueue")
	err := hq.SetWorkCompleted("testKeyComp", 1)
	assert.Nil(t, err)

	err = hq.SetWorkCompleted("testKeyComp", 3)
	assert.Nil(t, err)

	b, err := hq.CheckWorkCompletedInBulk("testKeyComp", []int64{1, 2, 3, 4})

	assert.Nil(t, err)
	assert.Equal(t, b[0], true)
	assert.Equal(t, b[2], true)
	assert.Equal(t, b[1], false)
	assert.Equal(t, b[3], false)
}
