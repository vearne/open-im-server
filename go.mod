module Open_IM

go 1.18

require (
	firebase.google.com/go v3.13.0+incompatible
	github.com/OpenIMSDK/openKeeper v0.9.7
	github.com/OpenIMSDK/open_utils v1.0.8
	github.com/Shopify/sarama v1.32.0
	github.com/alibabacloud-go/darabonba-openapi v0.1.11
	github.com/alibabacloud-go/sts-20150401 v1.1.0
	github.com/alibabacloud-go/tea v1.1.17
	github.com/antonfisher/nested-logrus-formatter v1.3.1
	github.com/aws/aws-sdk-go-v2 v1.16.7
	github.com/aws/aws-sdk-go-v2/config v1.15.14
	github.com/aws/aws-sdk-go-v2/credentials v1.12.9
	github.com/aws/aws-sdk-go-v2/service/sts v1.16.9
	github.com/bwmarrin/snowflake v0.3.0
	github.com/dtm-labs/rockscache v0.0.11
	github.com/fatih/structs v1.1.0
	github.com/gin-gonic/gin v1.8.2
	github.com/go-playground/validator/v10 v10.11.1
	github.com/go-redis/redis/v8 v8.11.5
	github.com/gogo/protobuf v1.3.2
	github.com/golang-jwt/jwt/v4 v4.4.2
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/websocket v1.4.2
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/jinzhu/copier v0.3.5
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/minio/minio-go/v7 v7.0.22
	github.com/mitchellh/mapstructure v1.4.2
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
	github.com/olivere/elastic/v7 v7.0.23
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.13.0
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5
	github.com/robfig/cron/v3 v3.0.1
	github.com/sirupsen/logrus v1.9.0
	github.com/stretchr/testify v1.8.1
	github.com/swaggo/files v0.0.0-20220610200504-28940afbdbfe
	github.com/swaggo/gin-swagger v1.5.0
	github.com/swaggo/swag v1.8.3 // indirect
	github.com/tencentyun/qcloud-cos-sts-sdk v0.0.0-20210325043845-84a0811633ca
	go.mongodb.org/mongo-driver v1.8.3
	golang.org/x/image v0.3.0
	golang.org/x/net v0.5.0
	golang.org/x/tools v0.1.12
	google.golang.org/api v0.103.0
	google.golang.org/grpc v1.52.3
	google.golang.org/protobuf v1.28.1
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	gopkg.in/yaml.v3 v3.0.1
	gorm.io/driver/mysql v1.3.5
	gorm.io/gorm v1.23.8
)

require (
	github.com/go-openapi/spec v0.20.6 // indirect
	github.com/go-openapi/swag v0.21.1 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/goccy/go-json v0.10.0 // indirect
	github.com/jonboulle/clockwork v0.3.0 // indirect
	github.com/lestrrat-go/strftime v1.0.6 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/ugorji/go/codec v1.2.8 // indirect
	golang.org/x/crypto v0.5.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	gopkg.in/ini.v1 v1.66.2 // indirect
)

replace github.com/Shopify/sarama => github.com/Shopify/sarama v1.29.0
