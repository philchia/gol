package gol

import (
	"errors"

	"github.com/philchia/gol/adapter"
)

// AddLogAdapter add a log adapter which implement the adapter.Adapter interface with give name key, return error if name already exists
func (l *gollog) AddLogAdapter(name string, adp adapter.Adapter) error {
	if adp == nil {
		return errors.New("nil adapter")
	}
	if _, ok := l.adapters[name]; ok {

		return errors.New("Adapter already exists")
	}

	tmpAdapters := make(map[string]adapter.Adapter, len(l.adapters)+1)

	for k, v := range l.adapters {
		tmpAdapters[k] = v
	}

	tmpAdapters[name] = adp
	l.adapters = tmpAdapters

	return nil
}

// RemoveAdapter remove a log adapter with give name key, return error in name not exists
func (l *gollog) RemoveAdapter(name string) error {
	if _, ok := l.adapters[name]; !ok {
		return errors.New("Adapter not exists")
	}

	tmpAdapters := make(map[string]adapter.Adapter, len(l.adapters)-1)

	for k, v := range l.adapters {
		if k != name {
			tmpAdapters[k] = v
		}
	}

	l.adapters = tmpAdapters

	return nil
}
