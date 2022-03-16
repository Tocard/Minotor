package data

import (
	"2miner-monitoring/es"
	"encoding/json"
	"time"

)

type Rentability struct {
	Timestamp string `json:"@timestamp"`
	NodesDefault              `json:"nodes_default"`
	FluxReward                float64 `json:"flux_reward,omitempty"`
	FluxInstantPaReward       float64 `json:"flux_instant_pa_reward,omitempty"`
	FluxLaterPaReward         float64 `json:"flux_later_pa_reward,omitempty"`
	DelayRewardMinutes        float64 `json:"delay_reward_minutes,omitempty"`
	DelayRewardDay            float64 `json:"delay_reward_day,omitempty"`
	FluxReward7Day            float64 `json:"flux_reward_7_day,omitempty"`
	FluxInstantPaReward7Day   float64 `json:"flux_instant_pa_reward_7_day,omitempty"`
	FluxLaterPaReward7Day     float64 `json:"flux_later_pa_reward_7_day,omitempty"`
	FluxTotalPaReward7Day     float64 `json:"flux_total_pa_reward_7_day,omitempty"`
	FluxTotalReward7Day       float64 `json:"flux_total_reward_7_day,omitempty"`
	FluxReward15Day           float64 `json:"flux_reward_15_day,omitempty"`
	FluxInstantPaReward15Day  float64 `json:"flux_instant_pa_reward_15_day,omitempty"`
	FluxLaterPaReward15Day    float64 `json:"flux_later_pa_reward_15_day,omitempty"`
	FluxTotalPaReward15Day    float64 `json:"flux_total_pa_reward_15_day,omitempty"`
	FluxTotalReward15Day      float64 `json:"flux_total_reward_15_day,omitempty"`
	FluxReward30Day           float64 `json:"flux_reward_30_day,omitempty"`
	FluxInstantPaReward30Day  float64 `json:"flux_instant_pa_reward_30_day,omitempty"`
	FluxLaterPaReward30Day    float64 `json:"flux_later_pa_reward_30_day,omitempty"`
	FluxTotalPaReward30Day    float64 `json:"flux_total_pa_reward_30_day,omitempty"`
	FluxTotalReward30Day      float64 `json:"flux_total_reward_30_day,omitempty"`
	FluxReward180Day          float64 `json:"flux_reward_180_day,omitempty"`
	FluxInstantPaReward180Day float64 `json:"flux_instant_pa_reward_180_day,omitempty"`
	FluxLaterPaReward180Day   float64 `json:"flux_later_pa_reward_180_day,omitempty"`
	FluxTotalPaReward180Day   float64 `json:"flux_total_pa_reward_180_day,omitempty"`
	FluxTotalReward180Day     float64 `json:"flux_total_reward_180_day,omitempty"`
	FluxReward365Day          float64 `json:"flux_reward_365_day,omitempty"`
	FluxInstantPaReward365Day float64 `json:"flux_instant_pa_reward_365_day,omitempty"`
	FluxLaterPaReward365Day   float64 `json:"flux_later_pa_reward_365_day,omitempty"`
	FluxTotalPaReward365Day   float64 `json:"flux_total_pa_reward_365_day,omitempty"`
	FluxTotalReward365Day     float64 `json:"flux_total_reward_365_day,omitempty"`
}

func getNodeCount(node NodesDefault, nodes Nodes) float64 {
	NodeCount := 0
	if node.Name == "Cumulus" {
		NodeCount = nodes.Data.CumulusEnabled
	} else if node.Name == "Nimbus" {
		NodeCount = nodes.Data.NimbusEnabled
	} else if node.Name == "Stratus" {
		NodeCount = nodes.Data.StratusEnabled
	}
	return float64(NodeCount)
}

func CalCulRentability(nodes Nodes, stats FluxBlocsStats) (int, string) {
	for _, NodeList := range NodesList {
		RentabilityNode := Rentability{}
		RentabilityNode.NodesDefault = NodeList
		RentabilityNode.FluxReward = FluxRewardPerBlock * NodeList.Reward / 100
		RentabilityNode.FluxInstantPaReward = FluxRewardPerBlock * NodeList.Reward / 100 / 2
		RentabilityNode.FluxLaterPaReward = RentabilityNode.FluxInstantPaReward
		RentabilityNode.DelayRewardMinutes = stats.TimeBetweenBlocks / 60 * getNodeCount(NodeList, nodes)
		RentabilityNode.DelayRewardDay = RentabilityNode.DelayRewardMinutes / 1440

		RentabilityNode.FluxInstantPaReward7Day = RentabilityNode.FluxInstantPaReward * 1440 * 7 / RentabilityNode.DelayRewardMinutes
		RentabilityNode.FluxReward7Day = RentabilityNode.FluxReward * 1440 * 7 / RentabilityNode.DelayRewardMinutes
		RentabilityNode.FluxLaterPaReward7Day = RentabilityNode.FluxLaterPaReward * 1440 * 7 / RentabilityNode.DelayRewardMinutes
		RentabilityNode.FluxTotalReward7Day = RentabilityNode.FluxLaterPaReward7Day + RentabilityNode.FluxReward7Day + RentabilityNode.FluxInstantPaReward7Day

		RentabilityNode.FluxInstantPaReward15Day = RentabilityNode.FluxInstantPaReward * 1440 * 15 / RentabilityNode.DelayRewardMinutes
		RentabilityNode.FluxReward15Day = RentabilityNode.FluxReward * 1440 * 15 / RentabilityNode.DelayRewardMinutes
		RentabilityNode.FluxLaterPaReward15Day = RentabilityNode.FluxLaterPaReward * 1440 * 15 / RentabilityNode.DelayRewardMinutes
		RentabilityNode.FluxTotalReward15Day = RentabilityNode.FluxLaterPaReward15Day + RentabilityNode.FluxReward15Day + RentabilityNode.FluxInstantPaReward15Day

		RentabilityNode.FluxInstantPaReward30Day = RentabilityNode.FluxInstantPaReward * 1440 * 30 / RentabilityNode.DelayRewardMinutes
		RentabilityNode.FluxReward30Day = RentabilityNode.FluxReward * 1440 * 30 / RentabilityNode.DelayRewardMinutes
		RentabilityNode.FluxLaterPaReward30Day = RentabilityNode.FluxLaterPaReward * 1440 * 30 / RentabilityNode.DelayRewardMinutes
		RentabilityNode.FluxTotalReward30Day = RentabilityNode.FluxLaterPaReward30Day + RentabilityNode.FluxReward30Day + RentabilityNode.FluxInstantPaReward30Day

		RentabilityNode.FluxInstantPaReward180Day = RentabilityNode.FluxInstantPaReward * 1440 * 180 / RentabilityNode.DelayRewardMinutes
		RentabilityNode.FluxReward180Day = RentabilityNode.FluxReward * 1440 * 180 / RentabilityNode.DelayRewardMinutes
		RentabilityNode.FluxLaterPaReward180Day = RentabilityNode.FluxLaterPaReward * 1440 * 180 / RentabilityNode.DelayRewardMinutes
		RentabilityNode.FluxTotalReward180Day = RentabilityNode.FluxLaterPaReward180Day + RentabilityNode.FluxReward180Day + RentabilityNode.FluxInstantPaReward180Day

		RentabilityNode.FluxInstantPaReward365Day = RentabilityNode.FluxInstantPaReward * 1440 * 365 / RentabilityNode.DelayRewardMinutes
		RentabilityNode.FluxReward365Day = RentabilityNode.FluxReward * 1440 * 365 / RentabilityNode.DelayRewardMinutes
		RentabilityNode.FluxLaterPaReward365Day = RentabilityNode.FluxLaterPaReward * 1440 * 365 / RentabilityNode.DelayRewardMinutes
		RentabilityNode.FluxTotalReward365Day = RentabilityNode.FluxLaterPaReward365Day + RentabilityNode.FluxReward365Day + RentabilityNode.FluxInstantPaReward365Day
		RentabilityNode.Timestamp = time.Now().Format(time.RFC3339)


		RentabilityNodeJson, err := json.Marshal(RentabilityNode)
		if err != nil {
			return 500, err.Error()
		}
		es.Bulk("2miners-flux-node", string(RentabilityNodeJson))
	}
	return 201, "All node Inseted"
}
