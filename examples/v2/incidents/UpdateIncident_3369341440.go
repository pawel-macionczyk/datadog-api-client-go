// Add commander to an incident returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/api/common"
	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	// there is a valid "incident" in the system
	IncidentDataID := os.Getenv("INCIDENT_DATA_ID")

	// there is a valid "user" in the system
	UserDataID := os.Getenv("USER_DATA_ID")

	body := datadog.IncidentUpdateRequest{
		Data: datadog.IncidentUpdateData{
			Id:   IncidentDataID,
			Type: datadog.INCIDENTTYPE_INCIDENTS,
			Relationships: &datadog.IncidentUpdateRelationships{
				CommanderUser: &datadog.NullableRelationshipToUser{
					Data: *datadog.NewNullableNullableRelationshipToUserData(&datadog.NullableRelationshipToUserData{
						Id:   UserDataID,
						Type: datadog.USERSTYPE_USERS,
					}),
				},
			},
		},
	}
	ctx := common.NewDefaultContext(context.Background())
	configuration := common.NewConfiguration()
	configuration.SetUnstableOperationEnabled("v2.UpdateIncident", true)
	apiClient := common.NewAPIClient(configuration)
	api := datadog.NewIncidentsApi(apiClient)
	resp, r, err := api.UpdateIncident(ctx, IncidentDataID, body, *datadog.NewUpdateIncidentOptionalParameters())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsApi.UpdateIncident`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `IncidentsApi.UpdateIncident`:\n%s\n", responseContent)
}
