// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package db

type ReadOpts struct {
	LoadDepth int
	OrderBy   map[string]bool
	GroupBy   string
	Limit     int
}

// DataReader describes an object which can read data from the database.
type DataReader interface {

	// Source attempts to find a data source with the given slug, taking into
	// account the given options.
	Source(string, *Source, ReadOpts) error

	// Survey attempts to find a dataset with the given slug, taking into
	// account the given options.
	Survey(string, *Survey, ReadOpts) error

	// Summary attempts to find a summary with the given slug, taking into
	// account the given options.
	Summary(string, *Summary, ReadOpts) error

	// Response attempts to load a survey response with the given id, taking
	// into account the given options.
	Response(string, *Response, ReadOpts) error

	// Respondent attempts to load a respondent with the given id, taking into
	// account the given options.
	Respondent(string, *Respondent, ReadOpts) error

	// Select attempts to run a a query agaist the database.
	// Args: query, query arguments, options, an interface to hold the expected
	// data.
	// Returns: error
	Select(string, []interface{}, ReadOpts, []interface{}) error
}

type WriteOpts struct {
	Upsert  bool
	Cascade bool
}

// DataWriter describes an object which can write data to the database.
type DataWriter interface {

	// WriteSource writes a source to the database.
	WriteSource(*Source, WriteOpts) error

	// LoadSource attempts to load a source into the database, including as
	// many of its component parts (surveys, respondents, etc) that are
	// provided.
	LoadSource(*Source, WriteOpts) error

	// WriteSurvey writes a survey to the database.
	WriteSurvey(*Survey, WriteOpts) error

	// LoadSurvey attempts to load a survey into the database, including as many
	// of its component parts (responses, etc) that are provided.
	LoadSurvey(*Survey, WriteOpts) error

	// WriteSummary writes a summary to the database.
	WriteSummary(*Summary, WriteOpts) error

	// WriteResponse writes a response to the database.
	WriteResponse(*Response, WriteOpts) error

	// LoadResponses loads a set of responses to the database.
	LoadResponses([]Response, WriteOpts) error

	// WriteRespondent writes a respondent to the database.
	WriteRespondent(*Respondent, WriteOpts) error

	// LoadRespondents loads a set of respondents to the database.
	LoadRespondents([]Respondent, WriteOpts) error

	// The following attempt to run statements against the database, using the
	// given string as the statement and the array of strings as the arguments.
	Insert(string, []interface{}, WriteOpts) error
	Update(string, []interface{}, WriteOpts) error
	Delete(string, []interface{}, WriteOpts) error
}
