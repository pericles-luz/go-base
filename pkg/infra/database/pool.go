package database

import "errors"

type Pool struct {
	connections map[string]*Database
}

func NewPool() *Pool {
	return &Pool{
		connections: make(map[string]*Database),
	}
}

func (p *Pool) GetConnection(name string) (*Database, error) {
	if conn, ok := p.connections[name]; ok {
		return conn, nil
	}
	return nil, errors.New("connection not found")
}

func (p *Pool) AddConnection(name string, conn *Database) error {
	if _, ok := p.connections[name]; ok {
		return errors.New("connection already exists")
	}
	p.connections[name] = conn
	return nil
}

func (p *Pool) RemoveConnection(name string) error {
	if _, ok := p.connections[name]; !ok {
		return errors.New("connection not found")
	}
	p.connections[name].Close()
	delete(p.connections, name)
	return nil
}
