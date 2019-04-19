package device

import (
	"fmt"

	sctypes "github.com/Microsoft/KubeDevice-API/pkg/devicescheduler"
	"github.com/Microsoft/KubeDevice-API/pkg/types"
	"github.com/Microsoft/KubeDevice/device-scheduler/grpalloc"
	"github.com/golang/glog"
)

type GrpDevice struct {
}

func (d *GrpDevice) AddNode(nodeName string, nodeInfo *types.NodeInfo) {
}

func (d *GrpDevice) RemoveNode(nodeName string) {
}

func (d *GrpDevice) PodFitsDevice(nodeInfo *types.NodeInfo, podInfo *types.PodInfo, fillAllocateFrom bool) (bool, []sctypes.PredicateFailureReason, float64) {
	glog.V(5).Infof("Running group scheduler on device requests %+v", podInfo)
	return grpalloc.PodFitsGroupConstraints(nodeInfo, podInfo, fillAllocateFrom)
}

func (d *GrpDevice) PodAllocate(nodeInfo *types.NodeInfo, podInfo *types.PodInfo) error {
	fits, reasons, _ := grpalloc.PodFitsGroupConstraints(nodeInfo, podInfo, true)
	if !fits {
		return fmt.Errorf("Scheduler unable to allocate pod %s as pod no longer fits: %v", podInfo.Name, reasons)
	}
	return nil
}

func (d *GrpDevice) TakePodResources(nodeInfo *types.NodeInfo, podInfo *types.PodInfo) error {
	grpalloc.TakePodGroupResource(nodeInfo, podInfo)
	return nil
}

func (d *GrpDevice) ReturnPodResources(nodeInfo *types.NodeInfo, podInfo *types.PodInfo) error {
	grpalloc.ReturnPodGroupResource(nodeInfo, podInfo)
	return nil
}

func (d *GrpDevice) GetName() string {
	return "grpdevice"
}

func (d *GrpDevice) UsingGroupScheduler() bool {
	return true
}
