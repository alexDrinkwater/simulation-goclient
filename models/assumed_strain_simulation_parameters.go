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

// AssumedStrainSimulationParameters assumed strain simulation parameters
// swagger:model AssumedStrainSimulationParameters
type AssumedStrainSimulationParameters struct {

	// anisotropic strain coefficients parallel
	// Required: true
	AnisotropicStrainCoefficientsParallel *float64 `json:"anisotropicStrainCoefficientsParallel"`

	// anisotropic strain coefficients perpendicular
	// Required: true
	AnisotropicStrainCoefficientsPerpendicular *float64 `json:"anisotropicStrainCoefficientsPerpendicular"`

	// anisotropic strain coefficients z
	// Required: true
	AnisotropicStrainCoefficientsZ *float64 `json:"anisotropicStrainCoefficientsZ"`

	// assumed strain
	// Required: true
	AssumedStrain *float64 `json:"assumedStrain"`

	// should be a number between 0.5 and 1.5
	BladeCrashThreshold float64 `json:"bladeCrashThreshold,omitempty"`

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

	// Must be between 0.00001 to 0.001 meters
	// Required: true
	// Maximum: 0.001
	// Minimum: 1e-05
	HatchSpacing *float64 `json:"hatchSpacing"`

	// Must be between 10 to 1000 watts
	// Required: true
	// Maximum: 1000
	// Minimum: 10
	LaserWattage *float64 `json:"laserWattage"`

	// Must be between 0 to 180 degrees
	// Required: true
	// Maximum: 180
	// Minimum: 0
	LayerRotationAngle *float64 `json:"layerRotationAngle"`

	// Must be between 0.00001 to 0.0001 meters
	// Required: true
	// Maximum: 0.0001
	// Minimum: 1e-05
	LayerThickness *float64 `json:"layerThickness"`

	// Must be between 0 to 0.003 meters, Must be greater than minimumWallDistance
	// Required: true
	// Maximum: 0.003
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

	// output shrinkage
	// Required: true
	OutputShrinkage *bool `json:"outputShrinkage"`

	// if true, a VTK file of the support structure will be created
	OutputSupportsVtk bool `json:"outputSupportsVtk,omitempty"`

	// if true, a predistorted STL file will be created using the distortion simulated by the mechanics solver
	PerformDistortionCompensation bool `json:"performDistortionCompensation,omitempty"`

	// perform support optimization
	// Required: true
	PerformSupportOptimization *bool `json:"performSupportOptimization"`

	// poisson ratio
	// Required: true
	PoissonRatio *float64 `json:"poissonRatio"`

	// Must be between 0.01 to 10 meters/second
	// Required: true
	// Maximum: 10
	// Minimum: 0.01
	ScanSpeed *float64 `json:"scanSpeed"`

	// List of parts to simulate (current limit is one part, imposed by server)
	SimulationParts []*SimulationPart `json:"simulationParts"`

	// Must be between 0.001 to 0.1 meters
	// Required: true
	// Maximum: 0.1
	// Minimum: 0.001
	SlicingStripeWidth *float64 `json:"slicingStripeWidth"`

	// Must be between 0 to 180 degrees
	// Required: true
	// Maximum: 180
	// Minimum: 0
	StartingLayerAngle *float64 `json:"startingLayerAngle"`

	// strain scaling factor
	// Required: true
	StrainScalingFactor *float64 `json:"strainScalingFactor"`

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

	// should be a number between 0.5 and 1.5
	SupportFailureThreshold float64 `json:"supportFailureThreshold,omitempty"`

	// support optimization
	// Required: true
	SupportOptimization *bool `json:"supportOptimization"`

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

// Validate validates this assumed strain simulation parameters
func (m *AssumedStrainSimulationParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnisotropicStrainCoefficientsParallel(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAnisotropicStrainCoefficientsPerpendicular(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAnisotropicStrainCoefficientsZ(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAssumedStrain(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDetectBladeCrash(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateElasticModulus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHatchSpacing(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLaserWattage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLayerRotationAngle(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLayerThickness(formats); err != nil {
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

	if err := m.validateOutputShrinkage(formats); err != nil {
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

	if err := m.validateScanSpeed(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSimulationParts(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSlicingStripeWidth(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStartingLayerAngle(formats); err != nil {
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

	if err := m.validateSupportOptimization(formats); err != nil {
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

func (m *AssumedStrainSimulationParameters) validateAnisotropicStrainCoefficientsParallel(formats strfmt.Registry) error {

	if err := validate.Required("anisotropicStrainCoefficientsParallel", "body", m.AnisotropicStrainCoefficientsParallel); err != nil {
		return err
	}

	return nil
}

func (m *AssumedStrainSimulationParameters) validateAnisotropicStrainCoefficientsPerpendicular(formats strfmt.Registry) error {

	if err := validate.Required("anisotropicStrainCoefficientsPerpendicular", "body", m.AnisotropicStrainCoefficientsPerpendicular); err != nil {
		return err
	}

	return nil
}

func (m *AssumedStrainSimulationParameters) validateAnisotropicStrainCoefficientsZ(formats strfmt.Registry) error {

	if err := validate.Required("anisotropicStrainCoefficientsZ", "body", m.AnisotropicStrainCoefficientsZ); err != nil {
		return err
	}

	return nil
}

func (m *AssumedStrainSimulationParameters) validateAssumedStrain(formats strfmt.Registry) error {

	if err := validate.Required("assumedStrain", "body", m.AssumedStrain); err != nil {
		return err
	}

	return nil
}

func (m *AssumedStrainSimulationParameters) validateDetectBladeCrash(formats strfmt.Registry) error {

	if err := validate.Required("detectBladeCrash", "body", m.DetectBladeCrash); err != nil {
		return err
	}

	return nil
}

func (m *AssumedStrainSimulationParameters) validateElasticModulus(formats strfmt.Registry) error {

	if err := validate.Required("elasticModulus", "body", m.ElasticModulus); err != nil {
		return err
	}

	return nil
}

func (m *AssumedStrainSimulationParameters) validateHatchSpacing(formats strfmt.Registry) error {

	if err := validate.Required("hatchSpacing", "body", m.HatchSpacing); err != nil {
		return err
	}

	if err := validate.Minimum("hatchSpacing", "body", float64(*m.HatchSpacing), 1e-05, false); err != nil {
		return err
	}

	if err := validate.Maximum("hatchSpacing", "body", float64(*m.HatchSpacing), 0.001, false); err != nil {
		return err
	}

	return nil
}

func (m *AssumedStrainSimulationParameters) validateLaserWattage(formats strfmt.Registry) error {

	if err := validate.Required("laserWattage", "body", m.LaserWattage); err != nil {
		return err
	}

	if err := validate.Minimum("laserWattage", "body", float64(*m.LaserWattage), 10, false); err != nil {
		return err
	}

	if err := validate.Maximum("laserWattage", "body", float64(*m.LaserWattage), 1000, false); err != nil {
		return err
	}

	return nil
}

func (m *AssumedStrainSimulationParameters) validateLayerRotationAngle(formats strfmt.Registry) error {

	if err := validate.Required("layerRotationAngle", "body", m.LayerRotationAngle); err != nil {
		return err
	}

	if err := validate.Minimum("layerRotationAngle", "body", float64(*m.LayerRotationAngle), 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("layerRotationAngle", "body", float64(*m.LayerRotationAngle), 180, false); err != nil {
		return err
	}

	return nil
}

func (m *AssumedStrainSimulationParameters) validateLayerThickness(formats strfmt.Registry) error {

	if err := validate.Required("layerThickness", "body", m.LayerThickness); err != nil {
		return err
	}

	if err := validate.Minimum("layerThickness", "body", float64(*m.LayerThickness), 1e-05, false); err != nil {
		return err
	}

	if err := validate.Maximum("layerThickness", "body", float64(*m.LayerThickness), 0.0001, false); err != nil {
		return err
	}

	return nil
}

func (m *AssumedStrainSimulationParameters) validateMaximumWallDistance(formats strfmt.Registry) error {

	if err := validate.Required("maximumWallDistance", "body", m.MaximumWallDistance); err != nil {
		return err
	}

	if err := validate.Minimum("maximumWallDistance", "body", float64(*m.MaximumWallDistance), 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("maximumWallDistance", "body", float64(*m.MaximumWallDistance), 0.003, false); err != nil {
		return err
	}

	return nil
}

func (m *AssumedStrainSimulationParameters) validateMaximumWallThickness(formats strfmt.Registry) error {

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

func (m *AssumedStrainSimulationParameters) validateMinimumSupportHeight(formats strfmt.Registry) error {

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

func (m *AssumedStrainSimulationParameters) validateMinimumWallDistance(formats strfmt.Registry) error {

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

func (m *AssumedStrainSimulationParameters) validateMinimumWallThickness(formats strfmt.Registry) error {

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

func (m *AssumedStrainSimulationParameters) validateOutputDisplacementAfterCutoff(formats strfmt.Registry) error {

	if err := validate.Required("outputDisplacementAfterCutoff", "body", m.OutputDisplacementAfterCutoff); err != nil {
		return err
	}

	return nil
}

func (m *AssumedStrainSimulationParameters) validateOutputLayerVtk(formats strfmt.Registry) error {

	if err := validate.Required("outputLayerVtk", "body", m.OutputLayerVtk); err != nil {
		return err
	}

	return nil
}

func (m *AssumedStrainSimulationParameters) validateOutputShrinkage(formats strfmt.Registry) error {

	if err := validate.Required("outputShrinkage", "body", m.OutputShrinkage); err != nil {
		return err
	}

	return nil
}

func (m *AssumedStrainSimulationParameters) validatePerformSupportOptimization(formats strfmt.Registry) error {

	if err := validate.Required("performSupportOptimization", "body", m.PerformSupportOptimization); err != nil {
		return err
	}

	return nil
}

func (m *AssumedStrainSimulationParameters) validatePoissonRatio(formats strfmt.Registry) error {

	if err := validate.Required("poissonRatio", "body", m.PoissonRatio); err != nil {
		return err
	}

	return nil
}

func (m *AssumedStrainSimulationParameters) validateScanSpeed(formats strfmt.Registry) error {

	if err := validate.Required("scanSpeed", "body", m.ScanSpeed); err != nil {
		return err
	}

	if err := validate.Minimum("scanSpeed", "body", float64(*m.ScanSpeed), 0.01, false); err != nil {
		return err
	}

	if err := validate.Maximum("scanSpeed", "body", float64(*m.ScanSpeed), 10, false); err != nil {
		return err
	}

	return nil
}

func (m *AssumedStrainSimulationParameters) validateSimulationParts(formats strfmt.Registry) error {

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

func (m *AssumedStrainSimulationParameters) validateSlicingStripeWidth(formats strfmt.Registry) error {

	if err := validate.Required("slicingStripeWidth", "body", m.SlicingStripeWidth); err != nil {
		return err
	}

	if err := validate.Minimum("slicingStripeWidth", "body", float64(*m.SlicingStripeWidth), 0.001, false); err != nil {
		return err
	}

	if err := validate.Maximum("slicingStripeWidth", "body", float64(*m.SlicingStripeWidth), 0.1, false); err != nil {
		return err
	}

	return nil
}

func (m *AssumedStrainSimulationParameters) validateStartingLayerAngle(formats strfmt.Registry) error {

	if err := validate.Required("startingLayerAngle", "body", m.StartingLayerAngle); err != nil {
		return err
	}

	if err := validate.Minimum("startingLayerAngle", "body", float64(*m.StartingLayerAngle), 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("startingLayerAngle", "body", float64(*m.StartingLayerAngle), 180, false); err != nil {
		return err
	}

	return nil
}

func (m *AssumedStrainSimulationParameters) validateStrainScalingFactor(formats strfmt.Registry) error {

	if err := validate.Required("strainScalingFactor", "body", m.StrainScalingFactor); err != nil {
		return err
	}

	return nil
}

var assumedStrainSimulationParametersTypeStressModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["linearElastic","j2Plasticity"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		assumedStrainSimulationParametersTypeStressModePropEnum = append(assumedStrainSimulationParametersTypeStressModePropEnum, v)
	}
}

const (
	// AssumedStrainSimulationParametersStressModeLinearElastic captures enum value "linearElastic"
	AssumedStrainSimulationParametersStressModeLinearElastic string = "linearElastic"
	// AssumedStrainSimulationParametersStressModeJ2Plasticity captures enum value "j2Plasticity"
	AssumedStrainSimulationParametersStressModeJ2Plasticity string = "j2Plasticity"
)

// prop value enum
func (m *AssumedStrainSimulationParameters) validateStressModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, assumedStrainSimulationParametersTypeStressModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AssumedStrainSimulationParameters) validateStressMode(formats strfmt.Registry) error {

	if err := validate.Required("stressMode", "body", m.StressMode); err != nil {
		return err
	}

	// value enum
	if err := m.validateStressModeEnum("stressMode", "body", *m.StressMode); err != nil {
		return err
	}

	return nil
}

func (m *AssumedStrainSimulationParameters) validateSupportAngle(formats strfmt.Registry) error {

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

func (m *AssumedStrainSimulationParameters) validateSupportFactorOfSafety(formats strfmt.Registry) error {

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

func (m *AssumedStrainSimulationParameters) validateSupportOptimization(formats strfmt.Registry) error {

	if err := validate.Required("supportOptimization", "body", m.SupportOptimization); err != nil {
		return err
	}

	return nil
}

func (m *AssumedStrainSimulationParameters) validateSupportYieldStrength(formats strfmt.Registry) error {

	if err := validate.Required("supportYieldStrength", "body", m.SupportYieldStrength); err != nil {
		return err
	}

	return nil
}

func (m *AssumedStrainSimulationParameters) validateSupportYieldStrengthRatio(formats strfmt.Registry) error {

	if err := validate.Required("supportYieldStrengthRatio", "body", m.SupportYieldStrengthRatio); err != nil {
		return err
	}

	return nil
}

func (m *AssumedStrainSimulationParameters) validateVoxelSize(formats strfmt.Registry) error {

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
func (m *AssumedStrainSimulationParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssumedStrainSimulationParameters) UnmarshalBinary(b []byte) error {
	var res AssumedStrainSimulationParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
