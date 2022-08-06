package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-perfect-store-execution-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToStoreValuationCollection(raw []byte, l *logger.Logger) ([]StoreValuationCollection, error) {
	pm := &responses.StoreValuationCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to StoreValuationCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	storeValuationCollection := make([]StoreValuationCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		storeValuationCollection = append(storeValuationCollection, StoreValuationCollection{
			ObjectID:            data.ObjectID,
			ETag:                data.ETag,
			TotalScoreValue:     data.TotalScoreValue,
			ReferenceVisitUUID:  data.ReferenceVisitUUID,
			StoreLayoutName:     data.StoreLayoutName,
			ValidFromDate:       data.ValidFromDate,
			ValidToDate:         data.ValidToDate,
			ReferenceVisitID:    data.ReferenceVisitID,
			StoreLayoutID:       data.StoreLayoutID,
			StoreLayoutVersion:  data.StoreLayoutVersion,
			EntityLastChangedOn: data.EntityLastChangedOn,
		})
	}

	return storeValuationCollection, nil
}
