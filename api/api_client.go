package api

type ApiClient struct {
	_settings *Settings
}

func NewApiClient() (*ApiClient, error) {
	client := &ApiClient{}
	settings, err := GetSettings()

	if(err != nil) {
		return nil, err
	}

	client._settings = settings
	return client, nil
}

func (this *ApiClient) Authorize() {

}

// TODO add more api's