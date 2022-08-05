module huawei.com/npu-exporter

go 1.16

require (
	github.com/agiledragon/gomonkey/v2 v2.8.0
	github.com/golang/protobuf v1.5.1
	github.com/patrickmn/go-cache v2.1.1-0.20191004192108-46f407853014+incompatible
	github.com/prometheus/client_golang v1.11.1
	github.com/smartystreets/goconvey v1.7.2
	github.com/stretchr/testify v1.7.0
	golang.org/x/crypto v0.0.0-20220314234659-1baeb1ce4c0b
	google.golang.org/grpc v1.28.0
	huawei.com/kmc v1.0.6
	huawei.com/mindx/common/hwlog v0.0.0
	huawei.com/mindx/common/limiter v0.0.0
	huawei.com/mindx/common/rand v0.0.0
	k8s.io/apimachinery v0.19.4
	k8s.io/client-go v0.19.4
	k8s.io/cri-api v0.19.4
)

replace (
	huawei.com/kmc => codehub-dg-y.huawei.com/it-edge-native/edge-native-core/coastguard.git v1.0.6
	huawei.com/mindx/common/hwlog => codehub-dg-y.huawei.com/MindX_DL/AtlasEnableWarehouse/common-utils.git/hwlog v0.0.1
	huawei.com/mindx/common/limiter => codehub-dg-y.huawei.com/MindX_DL/AtlasEnableWarehouse/common-utils.git/limiter v0.0.1
	huawei.com/mindx/common/rand => codehub-dg-y.huawei.com/MindX_DL/AtlasEnableWarehouse/common-utils.git/rand v0.0.1
	huawei.com/mindx/common/utils => codehub-dg-y.huawei.com/MindX_DL/AtlasEnableWarehouse/common-utils.git/utils v0.0.1
	k8s.io/api v0.0.0 => codehub-dg-y.huawei.com/OpenSourceCenter/kubernetes.git/staging/src/k8s.io/api v1.19.4-h4
	k8s.io/apimachinery => codehub-dg-y.huawei.com/OpenSourceCenter/kubernetes.git/staging/src/k8s.io/apimachinery v1.19.4-h4
	k8s.io/client-go => codehub-dg-y.huawei.com/OpenSourceCenter/kubernetes.git/staging/src/k8s.io/client-go v1.19.4-h4
	k8s.io/cri-api => codehub-dg-y.huawei.com/OpenSourceCenter/kubernetes.git/staging/src/k8s.io/cri-api v1.19.4-h4
)
