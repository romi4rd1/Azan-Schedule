/*
	This package intended to generate file only
*/

package database

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path"
	"time"

	azan "github.com/trihatmaja/Azan-Schedule"
)

type Files struct {
	FileName  string
	OutputDir string
}

type OptionFiles struct {
	OutputDir string
	FileName  string
}

func NewFiles(opt OptionFiles) *Files {
	return &Files{
		FileName:  opt.FileName,
		OutputDir: opt.OutputDir,
	}
}

func (f *Files) Set(data azan.CalcResult) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path.Join(f.OutputDir, f.FileName), js, 0644)
}

func (f *Files) Validate(lat, long float64, city string) (bool, error) {
	return false, nil
}

func (f *Files) GetAll() ([]azan.CalcResult, error) {
	return []azan.CalcResult{}, errors.New("Not Implemented Yet")
}

func (f *Files) GetByCity(city string) (azan.CalcResult, error) {
	return azan.CalcResult{}, errors.New("Not Implemented Yet")
}

func (f *Files) GetByDate(date time.Time) (azan.CalcResult, error) {
	return azan.CalcResult{}, errors.New("Not Implemented Yet")
}

func (f *Files) GetByCityDate(city string, date time.Time) (azan.CalcResult, error) {
	return azan.CalcResult{}, errors.New("Not Implemented Yet")
}

func (f *Files) GetByCityMonth(city string, month int) (azan.CalcResult, error) {
	return azan.CalcResult{}, errors.New("Not Implemented Yet")
}

func (f *Files) GetCities() ([]azan.CalcResult, error) {
	return []azan.CalcResult{}, errors.New("Not Implemented Yet")
}
