package models

import (
	"strconv"
)

type QueryParameter interface {
	GetLimit() uint64
	GetOffset() uint64
	GetParameters() map[string]string
}

func NewQueryParameter(values map[string][]string) QueryParameter {
	qp := parameters{
		limit:  10,
		offset: 0,
		values: make(map[string]string, 10),
	}

	for key, val := range values {
		if key == "limit" {
			limit, err := strconv.ParseUint(val[0], 10, 64)
			if err != nil {
				continue
			}
			if limit > 0 {
				qp.limit = limit
			}
			if limit > 100 {
				qp.limit = 100
			}
			continue
		}
		if key == "offset" {
			if offset, err := strconv.ParseUint(val[0], 10, 64); err == nil {
				qp.offset = offset
			}
			continue
		}
		qp.values[key] = val[0]
	}

	return &qp
}

type parameters struct {
	limit  uint64
	offset uint64
	values map[string]string
}

func (p *parameters) GetParameters() map[string]string {
	params := make(map[string]string)
	for key, val := range p.values {
		if len(val) >= 1 {
			params[key] = val
		}
	}
	return params
}

func (p *parameters) GetLimit() uint64 {
	return p.limit
}

func (p *parameters) GetOffset() uint64 {
	return p.offset
}
