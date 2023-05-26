package mq

import (
	"time"
)

func (mq *Mq) UserCollectCreate(body []byte) error {
	priority := uint32(1000)
	delay := 5 * time.Second
	ttr := 60 * time.Second
	_, err := mq.Put(body, priority, delay, ttr)
	if err != nil {
		return err
	}

	return nil
}
