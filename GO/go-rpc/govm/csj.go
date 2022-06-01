package main

import (
	"context"
	"fmt"
	"net/url"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
)

const (
	ip       = "172.118.69.31"
	user     = "administrator@vsphere.local"
	password = "Huayun@123"
)

// 这里的 c 就是上面登录后 client 的 Client 属性
func findVMByName(ctx context.Context, c *vim25.Client, vmName string) {
	m := view.NewManager(c)

	v, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, []string{"VirtualMachine"}, true)
	if err != nil {
		panic(err)
	}

	defer v.Destroy(ctx)

	// Retrieve summary property for all machines
	// Reference: http://pubs.vmware.com/vsphere-60/topic/com.vmware.wssdk.apiref.doc/vim.VirtualMachine.html
	var vms []mo.VirtualMachine
	err = v.Retrieve(ctx, []string{"VirtualMachine"}, []string{"summary"}, &vms)
	if err != nil {
		panic(err)
	}

	// Print summary per vm (see also: govc/vm/info.go)

	for _, vm := range vms {
		// 判断虚拟机名称是否相同，相同的话，vm 就是查找到主机
		if vm.Summary.Config.Name == vmName {
			fmt.Printf("%s: %s\n", vm.Summary.Config.Name, vm.Summary.Config.GuestFullName)
			break
		}
	}
}

func main() {
	u := &url.URL{
		Scheme: "https",
		Host:   ip,
		Path:   "/sdk",
	}
	ctx := context.Background()
	u.User = url.UserPassword(user, password)
	client, err := govmomi.NewClient(ctx, u, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(client)

	findVMByName(ctx, client, "apidr")
}

// export GOVC_URL="172.118.69.31"
// export GOVC_USERNAME="administrator@vsphere.local"
// export GOVC_PASSWORD="Huayun@123"
// export GOVC_INSECURE="true"
