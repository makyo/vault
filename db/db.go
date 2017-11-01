// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package db

import "database/sql"

// DB holds connection information and allows querying of the database
type DB struct {
	Conn    *sql.DB
	Options *Options
}

func (db *DB) Source(slug string, source *Source, opts ReadOpts) error {
	return nil
}

func (db *DB) Survey(slug string, source *Source, opts ReadOpts) error {
	return nil
}

func (db *DB) Summary(slug string, summary *Summary, opts ReadOpts) error {
	return nil
}

func (db *DB) Response(id string, response *Response, opts ReadOpts) error {
	return nil
}

func (db *DB) Respondent(id string, respondent *Respondent, opts ReadOpts) error {
	return nil
}

func (db *DB) Select(query string, args []interface{}, opts ReadOpts, result []interface{}) error {
	return nil
}

func (db *DB) WriteSource(source *Source, opts WriteOpts) error {
	return nil
}

func (db *DB) LoadSource(source *Source, opts WriteOpts) error {
	return nil
}

func (db *DB) WriteSurvey(survey *Survey, opts WriteOpts) error {
	return nil
}

func (db *DB) LoadSurvey(survey *Survey, opts WriteOpts) error {
	return nil
}

func (db *DB) WriteSummary(summary *Summary, opts WriteOpts) error {
	return nil
}

func (db *DB) WriteResponse(response *Response, opts WriteOpts) error {
	return nil
}

func (db *DB) LoadResponses(responses []Response, opts WriteOpts) error {
	return nil
}

func (db *DB) WriteRespondent(respondent *Respondent, opts WriteOpts) error {
	return nil
}

func (db *DB) LoadRespondents(respondents []Respondent, opts WriteOpts) error {
	return nil
}

func (db *DB) Insert(query string, arg []interface{}, opts WriteOpts) error {
	return nil
}

func (db *DB) Update(query string, arg []interface{}, opts WriteOpts) error {
	return nil
}

func (db *DB) Delete(query string, arg []interface{}, opts WriteOpts) error {
	return nil
}
