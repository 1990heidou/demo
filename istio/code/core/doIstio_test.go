package core

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIstioApi(t *testing.T) {
	ns := "default"
	ic := NewIstioClient(GetRestConfigOutCluster())
	// Test virtualService
	virtualService, err := ListVirtualService(ic, ns)
	require.NoError(t, err)
	for _, vs := range virtualService.Items {
		fmt.Printf("VirtualService name:%s gateways:%+v hosts: %+v\n", vs.Name, vs.Spec.GetGateways(), vs.Spec.GetHosts())
	}

	// Test DestinationRules
	drList, err := ListDestinationRules(ic, ns)
	for i := range drList.Items {
		dr := drList.Items[i]
		fmt.Printf("Index: %d DestinationRule Host: %+v\n", i, dr.Spec.GetHost())
	}

	// Test Gateway
	gwList, err := ListGateways(ic, ns)
	for i := range gwList.Items {
		gw := gwList.Items[i]
		for _, s := range gw.Spec.GetServers() {
			fmt.Printf("Index: %d Gateway servers: %+v\n", i, s)
		}
	}

	// Test ServiceEntry
	seList, err := ListServiceEntries(ic, ns)
	for i := range seList.Items {
		se := seList.Items[i]
		for _, h := range se.Spec.GetHosts() {
			fmt.Printf("Index: %d ServiceEntry hosts: %+v\n", i, h)
		}
	}
}
