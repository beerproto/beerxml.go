package beerXML

import (
	"encoding/json"
)

// A mash step is an internal record used within a mash profile to denote a separate step in a multi-step mash.
// A mash step is not intended for use outside of a mash profile.
type MashStep struct {
	// Name of the mash step – usually descriptive text such as “Dough In” or “Conversion”
	Name             string  `xml:"NAME" json:"name,omitempty"`
	// Version of the mash step record.  Should always be “1” for this version of the XML standard.
	Version          int32   `xml:"VERSION" json:"version,omitempty"`
	// May be “Infusion”, “Temperature” or “Decoction” depending on the type of step.
	// Infusion denotes adding hot water, Temperature denotes heating with an outside heat source,
	// and decoction denotes drawing off some mash for boiling.
	Type             string  `xml:"TYPE" json:"type,omitempty"`
	// The volume of water in liters to infuse in this step.  Required only for infusion steps, though one may also
	// add water for temperature mash steps.  One should not have an infusion amount for decoction steps.
	InfuseAmount     *float64 `xml:"INFUSE_AMOUNT" json:"infuse_amount,omitempty"`
	// The target temperature for this step in degrees Celsius.
	StepTemp         float64 `xml:"STEP_TEMP" json:"step_temp,omitempty"`
	// The number of minutes to spend at this step – i.e. the amount of time we are to hold this particular step
	// temperature.
	StepTime         int64   `xml:"STEP_TIME" json:"step_time,omitempty"`
	// Time in minutes to achieve the desired step temperature – useful particularly for temperature mashes where
	// it may take some time to achieve the step temperature.
	RampTime         *int64   `xml:"RAMP_TIME" json:"ramp_time,omitempty"`
	// the temperature you can expect the mash to fall to after a long mash step.  Measured in degrees Celsius.
	EndTemp          *float64 `xml:"END_TEMP" json:"end_temp,omitempty"`

	// Extensions

	// Textual description of this step such as “Infuse 4.5 gal of water at 170 F” – may be either generated by the
	// program or input by the user.
	Description      string  `xml:"DESCRIPTION" json:"description,omitempty"`
	// The total ratio of water to grain for this step AFTER the infusion along with the units, usually
	// expressed in qt/lb or l/kg.  Note this value must be consistent with the required infusion amount and amounts
	// added in earlier steps and is only relevant as part of a <MASH> profile.  For example “1.5 qt/lb” or “3.0 l/kg”
	WaterGrainRatio  string  `xml:"WATER_GRAIN_RATIO" json:"water_grain_ratio,omitempty"`
	// Calculated volume of mash to decoct.  Only applicable for a decoction step.
	// Includes the units as in “7.5 l” or “2.3 gal”
	DecoctionAmt     string  `xml:"DECOCTION_AMT" json:"decoction_amt,omitempty"`
	// The calculated infusion temperature based on the current step, grain, and other settings.
	// Applicable only for an infusion step.  Includes the units as in “154 F” or “68 C”
	InfuseTemp       string  `xml:"INFUSE_TEMP" json:"infuse_temp,omitempty"`
	// Step temperature in user defined temperature units.  For example “154F” or “68 C”
	DisplayStepTemp  string  `xml:"DISPLAY_STEP_TEMP" json:"display_step_temp,omitempty"`
	// Infusion amount along with the volume units as in “20 l” or “13 qt”
	DisplayInfuseAmt string  `xml:"DISPLAY_INFUSE_AMT" json:"display_infuse_amt,omitempty"`
}

type MashSteps struct {
	Mashstep []MashStep `xml:"MASH_STEP" json:"mash_step,omitempty"`
}

func (a MashSteps) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0)
	b = append(b, []byte("[")...)
	if len(a.Mashstep) > 0 {
		for _, hop := range a.Mashstep {
			bb, err := json.Marshal(hop)
			if err != nil {
				return nil, err
			}

			b = append(b, bb...)
			b = append(b, []byte(",")...)
		}

		// remove the trailing ','
		b = b[:len(b)-1]
	}
	b = append(b, []byte("]")...)

	return b, nil
}

func (a *MashSteps) UnmarshalJSON(b []byte) error {
	return nil
}
