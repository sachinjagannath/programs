package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type GameEntry struct {
	Score     int       `json:"score"`
	Timestamp time.Time `json:"timestamp"`
	Duration  int       `json:"duration_seconds"`
}
type GameData struct {
	HighScore int         `json:"high_score"`
	History   []GameEntry `json:"history"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type JsonStorage struct {
	filepath string
	mu       sync.RWMutex
	data     *GameData
}

func NewJsonStorage(path string) (*JsonStorage, error) {
	dir := filepath.Dir(path)
	if dir != "" && dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, fmt.Errorf("failed to create directory: %w", err)
		}
	}

	s := &JsonStorage{
		filepath: path,
		data: &GameData{
			HighScore: 0,
			History:   make([]GameEntry, 0),
			UpdatedAt: time.Now(),
		},
	}
	if err := s.load(); err != nil {
		if os.IsNotExist(err) {
			if err := s.save(); err != nil {
				return nil, fmt.Errorf("failed to create storage file: %w", err)
			}
		} else {
			return nil, fmt.Errorf("failed to load storage: %w", err)
		}
	}
	return s, nil
}

func (s *JsonStorage) save() error {
	data, err := json.Marshal(s.data)
	if err != nil {
		return fmt.Errorf("failed to marshal the data: %w", err)
	}
	tmpFile := s.filepath + ".tmp"
	if err := os.WriteFile(tmpFile, data, 0644); err != nil {
		return fmt.Errorf("failed to write temp file: %w", err)
	}

	if err := os.Rename(tmpFile, s.filepath); err != nil {
		err := os.Remove(tmpFile)
		if err != nil {
			return err
		}
		return fmt.Errorf("failed to rename the file: %w", err)
	}
	return nil
}

func (s *JsonStorage) load() error {
	data, err := os.ReadFile(s.filepath)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, s.data)
}
