package citrixclient

import (
	"context"
	"crypto/tls"
	"net/http"
	"time"

	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

type CitrixDaasClient struct {
	ApiClient    *openapiclient.APIClient
	AuthConfig   *AuthenticationConfiguration
	ClientConfig *ClientConfiguration
	AuthToken    *AuthTokenModel
}

type AuthTokenModel struct {
	Token     string `json:"access_token"`
	ExpiresAt string `json:"expires_in"`
}

func NewCitrixDaasClient(authUrl, hostname, customerId, clientId, clientSecret string, onPremise bool) (*CitrixDaasClient, error) {
	daasClient := &CitrixDaasClient{}

	/* ------ Setup API Client ------ */
	localCfg := openapiclient.NewConfiguration()
	localCfg.Host = hostname
	localCfg.Scheme = "https"
	// When running against on-prem, ignore SSL verification
	if onPremise {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		localCfg.HTTPClient = client
	}

	daasClient.ApiClient = openapiclient.NewAPIClient(localCfg)

	/* ------ Setup Authentication Configuration ------ */
	localAuthCfg := &AuthenticationConfiguration{}
	localAuthCfg.AuthUrl = authUrl
	localAuthCfg.ClientId = clientId
	localAuthCfg.ClientSecret = clientSecret
	localAuthCfg.OnPremise = onPremise

	daasClient.AuthConfig = localAuthCfg

	/* ------ Setup Client Configuration ------*/
	req := daasClient.ApiClient.MeTPApi.MeTPGetMe(context.Background())
	token, err := daasClient.SignIn()
	if err != nil {
		return nil, err
	}
	req = req.Authorization(token)
	resp, _, err := req.Execute()
	if err != nil {
		return nil, err
	}

	localClientCfg := &ClientConfiguration{}
	localClientCfg.CustomerId = customerId
	localClientCfg.SiteId = resp.Customers[0].Sites[0].Id

	daasClient.ClientConfig = localClientCfg

	return daasClient, nil
}

func (c *CitrixDaasClient) WaitForJob(ctx context.Context, jobId string, maxWaitTimeInMinutes int) (*openapiclient.JobResponseModel, error) {

	maxWaitTime := time.Now().UTC().Add(time.Minute * time.Duration(maxWaitTimeInMinutes))
	getJobIdRequest := c.ApiClient.JobsTPApi.JobsTPGetJob(ctx, jobId, c.ClientConfig.CustomerId, c.ClientConfig.SiteId)
	var jobResponseModel *openapiclient.JobResponseModel
	var err error

	for {

		if time.Now().UTC().After(maxWaitTime) {
			break
		}

		token, _ := c.SignIn()
		getJobIdRequest = getJobIdRequest.Authorization(token)

		jobResponseModel, _, err = getJobIdRequest.Execute()

		if err != nil {
			return jobResponseModel, err
		}

		if jobResponseModel.Status == openapiclient.JOBSTATUS_UNKNOWN || jobResponseModel.Status == openapiclient.JOBSTATUS_NOT_STARTED || jobResponseModel.Status == openapiclient.JOBSTATUS_IN_PROGRESS {
			time.Sleep(time.Second * time.Duration(30))
			continue
		}

		return jobResponseModel, err
	}

	return jobResponseModel, err
}
