package thirdapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"minotor/config"
	"net/http"
)

func harvestOperatorInfo(operatorID string) (map[string]interface{}, int) {
	result := make(map[string]interface{})

	payload := map[string]interface{}{
		"operationName": "getOperatorById",
		"variables": map[string]interface{}{
			"operatorId": operatorID,
		},
		"query": `your_graphql_query_here`,
	}

	payloadBytes, _ := json.Marshal(payload)
	response, err := http.Post(config.Cfg.StreamrApiUrl, "application/json", bytes.NewBuffer(payloadBytes))

	if err != nil {
		fmt.Println("Error for Operator:", err)
		return result, 500
	}

	defer response.Body.Close()

	if response.StatusCode == 200 {
		json.NewDecoder(response.Body).Decode(&result)
	} else {
		fmt.Println("Error for Operator")
	}

	return result, response.StatusCode
}

func HarvestAllOperatorsInfo() ([]byte, int) {

	graphqlQuery := `query getAllOperators($first: Int, $skip: Int) {
      operators(first: $first, skip: $skip) {
        ...OperatorFields
        __typename
      }
    }

    fragment OperatorFields on Operator {
      id
      stakes(first: 1000) {
        ...StakeFields
        sponsorship {
          ...SponsorshipFields
          __typename
        }
        __typename
      }
      delegations(first: 1000) {
        delegator {
          id
          __typename
        }
        valueDataWei
        operatorTokenBalanceWei
        id
        __typename
      }
      slashingEvents(first: 1000) {
        amount
        date
        sponsorship {
          id
          stream {
            id
            __typename
          }
          __typename
        }
        __typename
      }
      queueEntries(first: 1000) {
        amount
        date
        delegator {
          id
          __typename
        }
        id
        __typename
      }
      delegatorCount
      valueWithoutEarnings
      totalStakeInSponsorshipsWei
      dataTokenBalanceWei
      operatorTokenTotalSupplyWei
      metadataJsonString
      owner
      nodes
      cumulativeProfitsWei
      cumulativeOperatorsCutWei
      operatorsCutFraction
      __typename
    }

    fragment StakeFields on Stake {
      operator {
        id
        metadataJsonString
        __typename
      }
      amountWei
      earningsWei
      lockedWei
      joinTimestamp
      __typename
    }

    fragment SponsorshipFields on Sponsorship {
      id
      stream {
        id
        metadata
        __typename
      }
      metadata
      isRunning
      totalPayoutWeiPerSec
      stakes(first: 1000, orderBy: amountWei, orderDirection: desc) {
        ...StakeFields
        __typename
      }
      operatorCount
      maxOperators
      totalStakedWei
      remainingWei
      projectedInsolvency
      cumulativeSponsoring
      minimumStakingPeriodSeconds
      creator
      spotAPY
      __typename
    }`

	graphqlRequest := map[string]interface{}{
		"query": graphqlQuery,
		"variables": map[string]interface{}{
			"first": 1000, // Adjust as needed
			"skip":  0,
		},
	}

	payloadBytes, _ := json.Marshal(graphqlRequest)
	response, err := http.Post(config.Cfg.StreamrApiUrl, "application/json", bytes.NewBuffer(payloadBytes))

	if err != nil {
		return []byte(fmt.Sprintf("%s error on harvestAllOperatorsInfo", err)), 500
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte(fmt.Sprintf("%s error on harvestAllOperatorsInfo", err)), response.StatusCode
	}
	return body, 200

}
