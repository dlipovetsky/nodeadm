package apis

import (
	"github.com/platform9/nodeadm/constants"
	kubeadmv1alpha1 "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1alpha1"
)

// SetInitDefaults sets defaults on the configuration used by init
func SetInitDefaults(config *InitConfiguration) {
	// First set Networking defaults
	SetNetworkingDefaults(&config.Networking)
	// Second set MasterConfiguration.Networking defaults
	SetMasterConfigurationNetworkingDefaultsWithNetworking(config)
	// Third use the remainder of MasterConfiguration defaults
	kubeadmv1alpha1.SetDefaults_MasterConfiguration(&config.MasterConfiguration)
	config.MasterConfiguration.KubernetesVersion = constants.KUBERNETES_VERSION
	config.MasterConfiguration.NoTaintMaster = true
}

// SetJoinDefaults sets defaults on the configuration used by join
func SetJoinDefaults(config *JoinConfiguration) {
	SetNetworkingDefaults(&config.Networking)
}

// SetNetworkingDefaults sets defaults for the network configuration
func SetNetworkingDefaults(netConfig *Networking) {
	if netConfig.ServiceSubnet == "" {
		netConfig.ServiceSubnet = constants.DefaultServiceSubnet
	}
	if netConfig.DNSDomain == "" {
		netConfig.DNSDomain = constants.DefaultDNSDomain
	}
}

// SetMasterConfigurationNetworkingDefaultsWithNetworking sets defaults with
// values from the top-level network configuration
func SetMasterConfigurationNetworkingDefaultsWithNetworking(config *InitConfiguration) {
	if config.MasterConfiguration.Networking.ServiceSubnet == "" {
		config.MasterConfiguration.Networking.ServiceSubnet = config.Networking.ServiceSubnet
	}
	if config.MasterConfiguration.Networking.PodSubnet == "" {
		config.MasterConfiguration.Networking.PodSubnet = config.Networking.PodSubnet
	}
	if config.MasterConfiguration.Networking.DNSDomain == "" {
		config.MasterConfiguration.Networking.DNSDomain = config.Networking.DNSDomain
	}
}