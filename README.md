# sap-api-integrations-perfect-store-execution-reads
sap-api-integrations-perfect-store-execution-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API キャンペーンデータを取得するマイクロサービスです。  
sap-api-integrations-perfect-store-execution-reads には、サンプルのAPI Json フォーマットが含まれています。
sap-api-integrations-perfect-store-execution-reads は、オンプレミス版である（＝クラウド版ではない）SAPC4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/perfectstoreexecution/overview 

## 動作環境
sap-api-integrations-perfect-store-execution-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。
・ エッジ Kubernetes （推奨）
・ AION のリソース （推奨)
・ OS: LinuxOS （必須）
・ CPU: ARM/AMD/Intel（いずれか必須）

## クラウド環境での利用  
sap-api-integrations-perfect-store-execution-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-perfect-store-execution-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/perfectstoreexecution/overview 
* APIサービス名(=baseURL): c4codataapi

## 本レポジトリ に 含まれる API名
sap-api-integrations-perfect-store-execution-reads には、次の API をコールするためのリソースが含まれています。  

* StoreValuationCollection（ストアバリュエーションコレクション）

## API への 値入力条件 の 初期値
sap-api-integrations-perfect-store-execution-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.StoreValuationCollection.PerfectStoreExecutionID（パーフェクトストアエグゼキューションID）
* inoutSDC.StoreValuationCollection.PerfectStoreExecutionName（パーフェクトストアエグゼキューション名）


## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"StoreValuationCollection" が指定されています。    
  
```
	"api_schema": "PerfectStoreExecutionStoreValuationCollection",
	"accepter": ["StoreValuationCollection"],
	"campaign_code": "1",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "PerfectStoreExecutionStoreValuationCollection",
	"accepter": ["All"],
	"perfectstoreexecution_code": "1",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetPerfectStoreExecution(perfectstoreexecutionID, perfectstoreexecutionName string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "StoreValuationCollection":
			func() {
				c.StoreValuationCollection(perfectstoreexecutionID)
				wg.Done()
			}()
		case "StoreValuationCollectionName":
			func() {
				c.CampaignName(perfectstoreexecutionName)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP キャンペーン  の キャンペーンデータ が取得された結果の JSON の例です。  
以下の項目のうち、"ObjectID" ～ "EntityLastChangedOn" は、/SAP_API_Output_Formatter/type.go 内 の Type storevaluationcollection {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona5/bitbucket/sap-api-integrations-perfect-store-execution-reads/SAP_API_Caller/caller.go#L53",
	"function": "sap-api-integrations-perfect-store-execution-reads/SAP_API_Caller.(*SAPAPICaller).StoreValuationCollection",
	"level": "INFO",
	"message": [
		{
			"ObjectID": "00163E7930561EE996B9D4BEF50C0D85",
			"ETag": "2020-04-08T01:58:54+09:00",
			"TotalScoreValue": "50.00",
			"ReferenceVisitUUID": "00163E79-3056-1EE9-96B9-D24D2F71ED65",
			"StoreLayoutName": "Store Engagement Map.",
			"ValidFromDate": "",
			"ValidToDate": "9999-12-31T09:00:00+09:00",
			"ReferenceVisitID": "3466",
			"StoreLayoutID": "17",
			"StoreLayoutVersion": "1",
			"EntityLastChangedOn": "2019-04-09T16:02:39+09:00"
		}
	],
	"time": "2022-08-06T11:49:43+09:00"
}

```