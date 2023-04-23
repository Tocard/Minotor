package data

import "time"

type ChiaPlotCheckSummary struct {
	Plots []struct {
		Name         string  `json:"name"`
		Size         int     `json:"size"`
		Folder       string  `json:"folder"`
		ContractPool string  `json:"contract_pool"`
		FarmerKey    string  `json:"farmer_key"`
		PlotQuality  float64 `json:"plot_quality,omitempty"`
		ProofFound   int     `json:"proof_found,omitempty"`
		ProofTested  int     `json:"proof_tested,omitempty"`
		Pseudo       string  `json:"pseudo,omitempty"`
	} `json:"plots"`
	InvalidPlots []string `json:"invalid_plots"`
	Pseudo       string   `json:"pseudo"`
	FarmSummary  struct {
		PlotsCount        int     `json:"plots_count"`
		PlotsSize         float64 `json:"plots_size"`
		PlotsSizeUnit     string  `json:"plots_size_unit"`
		InvalidPlotsCount int     `json:"invalid_plots_count"`
	} `json:"farm_summary"`
	Timestamp time.Time `json:"@timestamp"`
}
