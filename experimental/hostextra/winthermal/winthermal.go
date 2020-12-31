// Copyright 2016 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Package winthermal reads the temperature sensors via WMI on Windows.
//
// This is an incomplete work-in-progress.
package winthermal

import (
	"errors"
	"time"

	"periph.io/x/conn/v3/driver"
	"periph.io/x/conn/v3/physic"
)

// Dev represents an handle to a WMI based sensor.
//
// Dev implements physic.SenseEnv.
type Dev struct {
	h obj
}

// String implements conn.Resource.
func (d *Dev) String() string {
	return d.h.InstanceName
}

// Halt implements conn.Resource.
func (d *Dev) Halt() error {
	return nil
}

// Sense implements physic.SenseEnv.
func (d *Dev) Sense(env *physic.Env) error {
	env.Temperature = physic.Temperature(d.h.CurrentTemperature)*100*physic.MilliCelsius + physic.ZeroCelsius
	return nil
}

func (d *Dev) SenseContinuous(interval time.Duration) (<-chan physic.Env, error) {
	return nil, errors.New("winthermal: not implemented yet")
}

func (d *Dev) Precision(e *physic.Env) {
}

//

// obj represents a MSAcpi_ThermalZoneTemperature instance. It intentionally
// leaves a lot of members out.
type obj struct {
	CurrentTemperature uint32
	InstanceName       string
	SamplingPeriod     int
}

type driverImpl struct {
}

func (d *driverImpl) String() string {
	return "winthermal"
}

func (d *driverImpl) After() []string {
	return nil
}

func (d *driverImpl) Prerequisites() []string {
	return nil
}

func (d *driverImpl) Init() (bool, error) {
	return true, initWindows()
}

var _ driver.Impl = &driverImpl{}
var _ physic.SenseEnv = &Dev{}
