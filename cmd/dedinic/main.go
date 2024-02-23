package main

import (
	"fmt"

	"github.com/kubeovn/kube-ovn/pkg/util"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/klog/v2"

	"github.com/multi-cluster-network/ovn-builder/pkg/dedinic"
)

func main() {

	defer klog.Flush()

	_ = dedinic.InitConfig()

	podInformerFactory := kubeinformers.NewSharedInformerFactoryWithOptions(dedinic.Conf.KubeClient, 0,
		kubeinformers.WithTweakListOptions(func(listOption *v1.ListOptions) {
			listOption.FieldSelector = fmt.Sprintf("spec.nodeName=%s", dedinic.Conf.NodeName)
			listOption.AllowWatchBookmarks = true
		}))
	nodeInformerFactory := kubeinformers.NewSharedInformerFactoryWithOptions(dedinic.Conf.KubeClient, 0,
		kubeinformers.WithTweakListOptions(func(listOption *v1.ListOptions) {
			listOption.AllowWatchBookmarks = true
		}))

	ctl, err := dedinic.NewController(dedinic.Conf, dedinic.StopCh, podInformerFactory, nodeInformerFactory)
	if err != nil {
		util.LogFatalAndExit(err, "failed to create controller")
	}
	klog.Info("start dedicnic controller")

	go dedinic.InitDelayQueue()

	go dedinic.InitNRIPlugin(dedinic.Conf, ctl)

	klog.Info("start nri dedicated plugin run")
	ctl.Run(dedinic.StopCh)
}
