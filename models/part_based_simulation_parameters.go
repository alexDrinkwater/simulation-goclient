// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PartBasedSimulationParameters part based simulation parameters
// swagger:model PartBasedSimulationParameters
type PartBasedSimulationParameters struct {

	// should be a number between 0.5 and 1.5
	BladeCrashThreshold float64 `json:"bladeCrashThreshold,omitempty"`

	// Id of build file being simulated, mutually exclusive with simulationParts
	BuildFileID int32 `json:"buildFileId,omitempty"`

	// detect blade crash
	// Required: true
	DetectBladeCrash *bool `json:"detectBladeCrash"`

	// detect support failure
	DetectSupportFailure bool `json:"detectSupportFailure,omitempty"`

	// a value that is used to scale the simulated distortion value
	DistortionScaleFactor float64 `json:"distortionScaleFactor,omitempty"`

	// elastic modulus
	// Required: true
	ElasticModulus *float64 `json:"elasticModulus"`

	// generate support voxels
	// Required: true
	GenerateSupportVoxels *bool `json:"generateSupportVoxels"`

	// Must be between 0 to 0.005 meters, Must be greater than minimumWallDistance
	// Required: true
	// Maximum: 0.005
	// Minimum: 0
	MaximumWallDistance *float64 `json:"maximumWallDistance"`

	// Must be between 0.00015 to 0.002 meters, Must be greater than minimumWallThickness
	// Required: true
	// Maximum: 0.002
	// Minimum: 0.00015
	MaximumWallThickness *float64 `json:"maximumWallThickness"`

	// Distance to move the part off the base plate for supports, Must be between 0 to 0.005 meters
	// Required: true
	// Maximum: 0.005
	// Minimum: 0
	MinimumSupportHeight *float64 `json:"minimumSupportHeight"`

	// Must be between 0 to 0.0003 meters, Must be less than maximumWallDistance
	// Required: true
	// Maximum: 0.0003
	// Minimum: 0
	MinimumWallDistance *float64 `json:"minimumWallDistance"`

	// Must be between 0.00005 to 0.0003 meters, Must be less than maximumWallThickness
	// Required: true
	// Maximum: 0.0003
	// Minimum: 5e-05
	MinimumWallThickness *float64 `json:"minimumWallThickness"`

	// output displacement after cutoff
	// Required: true
	OutputDisplacementAfterCutoff *bool `json:"outputDisplacementAfterCutoff"`

	// if true, mechanics solver output will include a zip file with the stress / distortion state at the end of each voxel layer
	// Required: true
	OutputLayerVtk *bool `json:"outputLayerVtk"`

	// if true, a VTK file of the support structure will be created
	OutputSupportsVtk bool `json:"outputSupportsVtk,omitempty"`

	// should be a number between 0.01 and 1.0
	PartFailureThreshold float64 `json:"partFailureThreshold,omitempty"`

	// part file location
	PartFileLocation string `json:"partFileLocation,omitempty"`

	// if true, a predistorted STL file will be created using the distortion simulated by the mechanics solver
	PerformDistortionCompensation bool `json:"performDistortionCompensation,omitempty"`

	// perform support optimization
	// Required: true
	PerformSupportOptimization *bool `json:"performSupportOptimization"`

	// poisson ratio
	// Required: true
	PoissonRatio *float64 `json:"poissonRatio"`

	// List of parts to simulate (current limit is one part, imposed by server)
	SimulationParts []*SimulationPart `json:"simulationParts"`

	// strain scaling factor
	// Required: true
	StrainScalingFactor *float64 `json:"strainScalingFactor"`

	// should be a number between 0.01 and 1.0
	StrainWarningThreshold float64 `json:"strainWarningThreshold,omitempty"`

	// stress mode
	// Required: true
	StressMode *string `json:"stressMode"`

	// Must be between 1 to 89 degrees
	// Required: true
	// Maximum: 89
	// Minimum: 1
	SupportAngle *float64 `json:"supportAngle"`

	// Multiplier for support calculations, Must be between 0.1 to 10
	// Required: true
	// Maximum: 10
	// Minimum: 0.1
	SupportFactorOfSafety *float64 `json:"supportFactorOfSafety"`

	// should be a number between 0.01 and 1.0
	SupportFailureThreshold float64 `json:"supportFailureThreshold,omitempty"`

	// support file location
	SupportFileLocation string `json:"supportFileLocation,omitempty"`

	// support yield strength
	// Required: true
	SupportYieldStrength *float64 `json:"supportYieldStrength"`

	// support yield strength ratio
	// Required: true
	SupportYieldStrengthRatio *float64 `json:"supportYieldStrengthRatio"`

	// Must be between 0.00002 to 0.002 meters
	// Required: true
	// Maximum: 0.002
	// Minimum: 2e-05
	VoxelSize *float64 `json:"voxelSize"`
}

// Validate validates this part based simulation parameters
func (m *PartBasedSimulationParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDetectBladeCrash(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateElasticModulus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGenerateSupportVoxels(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMaximumWallDistance(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMaximumWallThickness(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMinimumSupportHeight(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMinimumWallDistance(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMinimumWallThickness(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOutputDisplacementAfterCutoff(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOutputLayerVtk(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePerformSupportOptimization(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePoissonRatio(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSimulationParts(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStrainScalingFactor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStressMode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSupportAngle(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSupportFactorOfSafety(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSupportYieldStrength(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSupportYieldStrengthRatio(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVoxelSize(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PartBasedSimulationParameters) validateDetectBladeCrash(formats strfmt.Registry) error {

	if err := validate.Required("detectBladeCrash", "body", m.DetectBladeCrash); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateElasticModulus(formats strfmt.Registry) error {

	if err := validate.Required("elasticModulus", "body", m.ElasticModulus); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateGenerateSupportVoxels(formats strfmt.Registry) error {

	if err := validate.Required("generateSupportVoxels", "body", m.GenerateSupportVoxels); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateMaximumWallDistance(formats strfmt.Registry) error {

	if err := validate.Required("maximumWallDistance", "body", m.MaximumWallDistance); err != nil {
		return err
	}

	if err := validate.Minimum("maximumWallDistance", "body", float64(*m.MaximumWallDistance), 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("maximumWallDistance", "body", float64(*m.MaximumWallDistance), 0.005, false); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateMaximumWallThickness(formats strfmt.Registry) error {

	if err := validate.Required("maximumWallThickness", "body", m.MaximumWallThickness); err != nil {
		return err
	}

	if err := validate.Minimum("maximumWallThickness", "body", float64(*m.MaximumWallThickness), 0.00015, false); err != nil {
		return err
	}

	if err := validate.Maximum("maximumWallThickness", "body", float64(*m.MaximumWallThickness), 0.002, false); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateMinimumSupportHeight(formats strfmt.Registry) error {

	if err := validate.Required("minimumSupportHeight", "body", m.MinimumSupportHeight); err != nil {
		return err
	}

	if err := validate.Minimum("minimumSupportHeight", "body", float64(*m.MinimumSupportHeight), 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("minimumSupportHeight", "body", float64(*m.MinimumSupportHeight), 0.005, false); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateMinimumWallDistance(formats strfmt.Registry) error {

	if err := validate.Required("minimumWallDistance", "body", m.MinimumWallDistance); err != nil {
		return err
	}

	if err := validate.Minimum("minimumWallDistance", "body", float64(*m.MinimumWallDistance), 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("minimumWallDistance", "body", float64(*m.MinimumWallDistance), 0.0003, false); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateMinimumWallThickness(formats strfmt.Registry) error {

	if err := validate.Required("minimumWallThickness", "body", m.MinimumWallThickness); err != nil {
		return err
	}

	if err := validate.Minimum("minimumWallThickness", "body", float64(*m.MinimumWallThickness), 5e-05, false); err != nil {
		return err
	}

	if err := validate.Maximum("minimumWallThickness", "body", float64(*m.MinimumWallThickness), 0.0003, false); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateOutputDisplacementAfterCutoff(formats strfmt.Registry) error {

	if err := validate.Required("outputDisplacementAfterCutoff", "body", m.OutputDisplacementAfterCutoff); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateOutputLayerVtk(formats strfmt.Registry) error {

	if err := validate.Required("outputLayerVtk", "body", m.OutputLayerVtk); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validatePerformSupportOptimization(formats strfmt.Registry) error {

	if err := validate.Required("performSupportOptimization", "body", m.PerformSupportOptimization); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validatePoissonRatio(formats strfmt.Registry) error {

	if err := validate.Required("poissonRatio", "body", m.PoissonRatio); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateSimulationParts(formats strfmt.Registry) error {

	if swag.IsZero(m.SimulationParts) { // not required
		return nil
	}

	for i := 0; i < len(m.SimulationParts); i++ {

		if swag.IsZero(m.SimulationParts[i]) { // not required
			continue
		}

		if m.SimulationParts[i] != nil {

			if err := m.SimulationParts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("simulationParts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PartBasedSimulationParameters) validateStrainScalingFactor(formats strfmt.Registry) error {

	if err := validate.Required("strainScalingFactor", "body", m.StrainScalingFactor); err != nil {
		return err
	}

	return nil
}

var partBasedSimulationParametersTypeStressModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["LinearElastic","J2Plasticity"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		partBasedSimulationParametersTypeStressModePropEnum = append(partBasedSimulationParametersTypeStressModePropEnum, v)
	}
}

const (
	// PartBasedSimulationParametersStressModeLinearElastic captures enum value "LinearElastic"
	PartBasedSimulationParametersStressModeLinearElastic string = "LinearElastic"
	// PartBasedSimulationParametersStressModeJ2Plasticity captures enum value "J2Plasticity"
	PartBasedSimulationParametersStressModeJ2Plasticity string = "J2Plasticity"
)

// prop value enum
func (m *PartBasedSimulationParameters) validateStressModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, partBasedSimulationParametersTypeStressModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PartBasedSimulationParameters) validateStressMode(formats strfmt.Registry) error {

	if err := validate.Required("stressMode", "body", m.StressMode); err != nil {
		return err
	}

	// value enum
	if err := m.validateStressModeEnum("stressMode", "body", *m.StressMode); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateSupportAngle(formats strfmt.Registry) error {

	if err := validate.Required("supportAngle", "body", m.SupportAngle); err != nil {
		return err
	}

	if err := validate.Minimum("supportAngle", "body", float64(*m.SupportAngle), 1, false); err != nil {
		return err
	}

	if err := validate.Maximum("supportAngle", "body", float64(*m.SupportAngle), 89, false); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateSupportFactorOfSafety(formats strfmt.Registry) error {

	if err := validate.Required("supportFactorOfSafety", "body", m.SupportFactorOfSafety); err != nil {
		return err
	}

	if err := validate.Minimum("supportFactorOfSafety", "body", float64(*m.SupportFactorOfSafety), 0.1, false); err != nil {
		return err
	}

	if err := validate.Maximum("supportFactorOfSafety", "body", float64(*m.SupportFactorOfSafety), 10, false); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateSupportYieldStrength(formats strfmt.Registry) error {

	if err := validate.Required("supportYieldStrength", "body", m.SupportYieldStrength); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateSupportYieldStrengthRatio(formats strfmt.Registry) error {

	if err := validate.Required("supportYieldStrengthRatio", "body", m.SupportYieldStrengthRatio); err != nil {
		return err
	}

	return nil
}

func (m *PartBasedSimulationParameters) validateVoxelSize(formats strfmt.Registry) error {

	if err := validate.Required("voxelSize", "body", m.VoxelSize); err != nil {
		return err
	}

	if err := validate.Minimum("voxelSize", "body", float64(*m.VoxelSize), 2e-05, false); err != nil {
		return err
	}

	if err := validate.Maximum("voxelSize", "body", float64(*m.VoxelSize), 0.002, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PartBasedSimulationParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PartBasedSimulationParameters) UnmarshalBinary(b []byte) error {
	var res PartBasedSimulationParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
