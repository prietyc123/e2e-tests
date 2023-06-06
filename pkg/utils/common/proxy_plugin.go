package common

import (
	"context"

	toolchainv1alpha1 "github.com/codeready-toolchain/api/api/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *SuiteController) CreateProxyPlugin(proxyPluginName, proxyPluginNamespace, routeName, routeNamespace string) error {

	// Create the ProxyPlugin object
	proxy := &toolchainv1alpha1.ProxyPlugin{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: proxyPluginNamespace,
			Name:      proxyPluginName,
		},
		Spec: toolchainv1alpha1.ProxyPluginSpec{
			OpenShiftRouteTargetEndpoint: &toolchainv1alpha1.OpenShiftRouteTarget{
				Namespace: routeNamespace,
				Name:      routeName,
			},
		},
	}

	if err := s.KubeRest().Create(context.TODO(), proxy); err != nil {
		return err
	}
	return nil
}
