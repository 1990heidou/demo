package core

import (
	"context"
	"encoding/json"
	"fmt"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	versionedclient "istio.io/client-go/pkg/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"log"
)

/*
istio 中有几个新的资源概念：
gateway: gw
virtual service: vs
destination rule: dr
service entry: se
*/

func NewIstioClient(config *rest.Config) *versionedclient.Clientset {
	ic, err := versionedclient.NewForConfig(config)
	if err != nil {
		log.Fatalf("Failed to create istio client: %s", err)
	}
	return ic
}

func CreateVirtualService(ic *versionedclient.Clientset, ns string, virtualServiceRequestData []byte) (*v1alpha3.VirtualService, error) {
	virtualService := &v1alpha3.VirtualService{}
	err := json.Unmarshal(virtualServiceRequestData, &virtualService)
	if err != nil {
		fmt.Println("反序列化失败")
		return nil, err
	}
	vs, err := ic.NetworkingV1alpha3().VirtualServices(ns).Create(context.TODO(), virtualService, metav1.CreateOptions{})
	if err != nil {
		fmt.Println("创建虚拟服务失败")
		return nil, err
	}
	return vs, nil
}

func ListVirtualService(ic *versionedclient.Clientset, ns string) (*v1alpha3.VirtualServiceList, error) {
	vsList, err := ic.NetworkingV1alpha3().VirtualServices(ns).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Failed to get VirtualService in %s namespace: %s", ns, err)
		return nil, err
	}

	return vsList, err
}

func ListDestinationRules(ic *versionedclient.Clientset, ns string) (*v1alpha3.DestinationRuleList, error) {
	drList, err := ic.NetworkingV1alpha3().DestinationRules(ns).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Failed to get DestinationRule in %s namespace: %s", ns, err)
		return nil, err
	}

	return drList, nil
}

func CreateGateway(ic *versionedclient.Clientset, ns string, gatewayRequestData []byte) (*v1alpha3.Gateway, error) {
	gateway := &v1alpha3.Gateway{}
	err := json.Unmarshal(gatewayRequestData, &gateway)
	if err != nil {
		fmt.Println("反序列化失败")
		return nil, err
	}
	vs, err := ic.NetworkingV1alpha3().Gateways(ns).Create(context.TODO(), gateway, metav1.CreateOptions{})
	if err != nil {
		fmt.Println("创建网关失败")
		return nil, err
	}
	return vs, nil
}

func ListGateways(ic *versionedclient.Clientset, ns string) (*v1alpha3.GatewayList, error) {
	gwList, err := ic.NetworkingV1alpha3().Gateways(ns).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Failed to get Gateway in %s namespace: %s", ns, err)
	}

	return gwList, nil
}

// ListServiceEntries 服务入口
func ListServiceEntries(ic *versionedclient.Clientset, ns string) (*v1alpha3.ServiceEntryList, error) {
	seList, err := ic.NetworkingV1alpha3().ServiceEntries(ns).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Failed to get ServiceEntry in %s namespace: %s", ns, err)
	}

	return seList, nil
}
