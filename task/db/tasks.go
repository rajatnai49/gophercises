package db

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

type Task struct {
	Key   int
	Value Value
}

type Value struct {
	Task   string
	Status bool
}

var taskBucket = []byte("new_tasks")
var db *bolt.DB

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		value := Value{
			Task:   task,
			Status: false,
		}
		buf, err := encodeFromJson(value)
		if err != nil {
			return err
		}
		return b.Put(itob(id), buf.Bytes())
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

func AllTask() ([]Task, error) {
	var tasks []Task

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			value, err := decodeToJson(v)
			if err == nil {
				task := Task{
					Key:   btoi(k),
					Value: value,
				}
				if !value.Status {
					tasks = append(tasks, task)
				}
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func DeleteTask(indexes []int) {
	tasks, err := AllTask()
	for _, index := range indexes {
		index = index - 1
		if err != nil || index < 0 || index >= len(tasks) {
			fmt.Printf("Not Valid Index: %d\n", index)
		}
		task := tasks[index]
		key := itob(task.Key)
		err = db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket(taskBucket)
			err := b.Delete(key)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			fmt.Printf("Not Valid Index: %d\n", index)
		}
	}
}

func CompleteTask(indexes []int) {
	tasks, err := AllTask()
	for _, index := range indexes {
		index = index - 1
		if err != nil || index < 0 || index >= len(tasks) {
			fmt.Printf("Not Valid Index: %d\n", index)
		}
		task := tasks[index]
		key := itob(task.Key)
		err = db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket(taskBucket)
			task.Value.Status = true
			buf, err := encodeFromJson(task.Value)
			if err != nil {
				return err
			}
			return b.Put(key, buf.Bytes())
		})
		if err != nil {
			fmt.Printf("Not Valid Index: %d\n", index)
		}
	}
}

func encodeFromJson(data Value) (bytes.Buffer, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(data); err != nil {
		return buf, err
	}
	return buf, nil
}

func decodeToJson(data []byte) (Value, error) {
	var value Value
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	if err := dec.Decode(&value); err != nil {
		return value, err
	}
	return value, nil
}

func itob(i int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
